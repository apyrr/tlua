package modulespecifiers

import (
	"maps"
	"slices"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/module"
	"github.com/apyrr/tlua/internal/packagejson"
	"github.com/apyrr/tlua/internal/stringutil"
	"github.com/apyrr/tlua/internal/tspath"
)

func GetModuleSpecifiers(
	moduleSymbol *ast.Symbol,
	checker CheckerShape,
	compilerOptions *core.CompilerOptions,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
	forAutoImports bool,
) []string {
	result, _ := GetModuleSpecifiersWithInfo(
		moduleSymbol,
		checker,
		compilerOptions,
		importingSourceFile,
		host,
		userPreferences,
		options,
		forAutoImports,
	)
	return result
}

func GetModuleSpecifiersWithInfo(
	moduleSymbol *ast.Symbol,
	checker CheckerShape,
	compilerOptions *core.CompilerOptions,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
	forAutoImports bool,
) ([]string, ResultKind) {
	ambient := tryGetModuleNameFromAmbientModule(moduleSymbol, checker)
	if len(ambient) > 0 {
		if forAutoImports && IsExcludedByRegex(ambient, userPreferences.AutoImportSpecifierExcludeRegexes) {
			return nil, ResultKindAmbient
		}
		return []string{ambient}, ResultKindAmbient
	}

	moduleSourceFile := ast.GetSourceFileOfModule(moduleSymbol)
	if moduleSourceFile == nil {
		return nil, ResultKindNone
	}

	// Use original source file name when file is from project reference output
	moduleFileName := host.GetSourceOfProjectReferenceIfOutputIncluded(moduleSourceFile)

	return GetModuleSpecifiersForFileWithInfo(
		importingSourceFile,
		moduleFileName,
		compilerOptions,
		host,
		userPreferences,
		options,
		forAutoImports,
	)
}

func GetModuleSpecifiersForFileWithInfo(
	importingSourceFile SourceFileForSpecifierGeneration,
	moduleFileName string,
	compilerOptions *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
	forAutoImports bool,
) ([]string, ResultKind) {
	modulePaths := getAllModulePathsWorker(
		getInfo(host.GetSourceOfProjectReferenceIfOutputIncluded(importingSourceFile), host),
		moduleFileName,
		host,
		compilerOptions,
		options,
	)

	return computeModuleSpecifiers(
		modulePaths,
		compilerOptions,
		importingSourceFile,
		host,
		userPreferences,
		options,
		forAutoImports,
	)
}

func tryGetModuleNameFromAmbientModule(moduleSymbol *ast.Symbol, checker CheckerShape) string {
	for _, decl := range moduleSymbol.Declarations {
		if ast.IsModuleWithStringLiteralName(decl) && (!ast.IsModuleAugmentationExternal(decl) || !tspath.IsExternalModuleNameRelative(decl.Name().Text())) {
			return decl.Name().Text()
		}
	}

	// the module could be a namespace, which is export through "export=" from an ambient module.
	/**
	 * declare module "m" {
	 *     namespace ns {
	 *         class c {}
	 *     }
	 *     export = ns;
	 * }
	 */
	// `import {c} from "m";` is valid, in which case, `moduleSymbol` is "ns", but the module name should be "m"
	for _, d := range moduleSymbol.Declarations {
		if !ast.IsModuleDeclaration(d) {
			continue
		}

		possibleContainer := ast.FindAncestor(d, ast.IsModuleWithStringLiteralName)
		if possibleContainer == nil || possibleContainer.Parent == nil || !ast.IsSourceFile(possibleContainer.Parent) {
			continue
		}

		sym, ok := possibleContainer.Symbol().Exports[ast.InternalSymbolNameExportEquals]
		if !ok || sym == nil {
			continue
		}
		exportAssignmentDecl := sym.ValueDeclaration
		if exportAssignmentDecl == nil || exportAssignmentDecl.Kind != ast.KindExportAssignment {
			continue
		}
		exportSymbol := checker.GetSymbolAtLocation(exportAssignmentDecl.Expression())
		if exportSymbol == nil {
			continue
		}
		if exportSymbol.Flags&ast.SymbolFlagsAlias != 0 {
			exportSymbol = checker.GetAliasedSymbol(exportSymbol)
		}
		// TODO: Possible strada bug - isn't this insufficient in the presence of merge symbols?
		if exportSymbol == d.Symbol() {
			return possibleContainer.Name().Text()
		}
	}
	return ""
}

type Info struct {
	UseCaseSensitiveFileNames bool
	ImportingSourceFileName   string
	SourceDirectory           string
}

func getInfo(
	importingSourceFileName string,
	host ModuleSpecifierGenerationHost,
) Info {
	sourceDirectory := tspath.GetDirectoryPath(importingSourceFileName)
	return Info{
		ImportingSourceFileName:   importingSourceFileName,
		SourceDirectory:           sourceDirectory,
		UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
	}
}

func getAllModulePaths(
	info Info,
	importedFileName string,
	host ModuleSpecifierGenerationHost,
	compilerOptions *core.CompilerOptions,
	preferences UserPreferences,
	options ModuleSpecifierOptions,
) []ModulePath {
	// !!! use new cache model
	// importingFilePath := tspath.ToPath(info.ImportingSourceFileName, host.GetCurrentDirectory(), host.UseCaseSensitiveFileNames());
	// importedFilePath := tspath.ToPath(importedFileName, host.GetCurrentDirectory(), host.UseCaseSensitiveFileNames());
	// cache := host.getModuleSpecifierCache();
	// if (cache != nil) {
	//     cached := cache.get(importingFilePath, importedFilePath, preferences, options);
	//     if (cached.modulePaths) {return cached.modulePaths;}
	// }
	modulePaths := getAllModulePathsWorker(info, importedFileName, host, compilerOptions, options)
	// if (cache != nil) {
	//     cache.setModulePaths(importingFilePath, importedFilePath, preferences, options, modulePaths);
	// }
	return modulePaths
}

func getAllModulePathsWorker(
	info Info,
	importedFileName string,
	host ModuleSpecifierGenerationHost,
	compilerOptions *core.CompilerOptions,
	options ModuleSpecifierOptions,
) []ModulePath {
	allFileNames := make(map[string]ModulePath)
	paths := GetEachFileNameOfModule(info.ImportingSourceFileName, importedFileName, host, true)
	for _, p := range paths {
		allFileNames[p.FileName] = p
	}

	useCaseSensitiveFileNames := info.UseCaseSensitiveFileNames
	comparePaths := func(a, b ModulePath) int {
		return comparePathsByRedirect(a, b, useCaseSensitiveFileNames)
	}

	// Sort by paths closest to importing file Name directory
	sortedPaths := make([]ModulePath, 0, len(paths))
	for directory := info.SourceDirectory; len(allFileNames) != 0; {
		directoryStart := tspath.EnsureTrailingDirectorySeparator(directory)
		var pathsInDirectory []ModulePath
		for fileName, p := range allFileNames {
			if strings.HasPrefix(fileName, directoryStart) {
				pathsInDirectory = append(pathsInDirectory, p)
				delete(allFileNames, fileName)
			}
		}
		if len(pathsInDirectory) > 0 {
			slices.SortFunc(pathsInDirectory, comparePaths)
			sortedPaths = append(sortedPaths, pathsInDirectory...)
		}
		newDirectory := tspath.GetDirectoryPath(directory)
		if newDirectory == directory {
			break
		}
		directory = newDirectory
	}
	if len(allFileNames) > 0 {
		remainingPaths := slices.Collect(maps.Values(allFileNames))
		slices.SortFunc(remainingPaths, comparePaths)
		sortedPaths = append(sortedPaths, remainingPaths...)
	}
	return sortedPaths
}

// containsIgnoredPath checks if a path contains patterns that should be ignored.
// This is a local helper that duplicates tspath.ContainsIgnoredPath for performance.
func containsIgnoredPath(s string) bool {
	return strings.Contains(s, "/node_modules/.") ||
		strings.Contains(s, "/.git") ||
		strings.Contains(s, ".#")
}

// ContainsNodeModules checks if a path contains the node_modules directory.
func ContainsNodeModules(s string) bool {
	return strings.Contains(s, "/node_modules/")
}

// GetEachFileNameOfModule returns all possible file paths for a module, including symlink alternatives.
// This function handles symlink resolution and provides multiple path options for module resolution.
func GetEachFileNameOfModule(
	importingFileName string,
	importedFileName string,
	host ModuleSpecifierGenerationHost,
	preferSymlinks bool,
) []ModulePath {
	cwd := host.GetCurrentDirectory()
	importedPath := tspath.ToPath(importedFileName, cwd, host.UseCaseSensitiveFileNames())
	var referenceRedirect string
	outputAndReference := host.GetProjectReferenceFromSource(importedPath)
	if outputAndReference != nil && outputAndReference.OutputDts != "" {
		referenceRedirect = outputAndReference.OutputDts
	}

	redirects := host.GetRedirectTargets(importedPath)
	importedFileNames := make([]string, 0, 2+len(redirects))
	if len(referenceRedirect) > 0 {
		importedFileNames = append(importedFileNames, referenceRedirect)
	}
	importedFileNames = append(importedFileNames, importedFileName)
	importedFileNames = append(importedFileNames, redirects...)
	targets := core.Map(importedFileNames, func(f string) string { return tspath.GetNormalizedAbsolutePath(f, cwd) })
	shouldFilterIgnoredPaths := !core.Every(targets, containsIgnoredPath)

	results := make([]ModulePath, 0, 2)
	if !preferSymlinks {
		for _, p := range targets {
			if !(shouldFilterIgnoredPaths && containsIgnoredPath(p)) {
				results = append(results, ModulePath{
					FileName:        p,
					IsInNodeModules: ContainsNodeModules(p),
					IsRedirect:      referenceRedirect == p,
				})
			}
		}
	}

	symlinkCache := host.GetSymlinkCache()
	fullImportedFileName := tspath.GetNormalizedAbsolutePath(importedFileName, cwd)
	if symlinkCache != nil {
		tspath.ForEachAncestorDirectoryStoppingAtGlobalCache(
			host.GetGlobalTypingsCacheLocation(),
			tspath.GetDirectoryPath(fullImportedFileName),
			func(realPathDirectory string) (bool, bool) {
				symlinkSet, ok := symlinkCache.DirectoriesByRealpath().Load(tspath.ToPath(realPathDirectory, cwd, host.UseCaseSensitiveFileNames()).EnsureTrailingDirectorySeparator())
				if !ok {
					return false, false
				} // Continue to ancestor directory

				// Don't want to a package to globally import from itself (importNameCodeFix_symlink_own_package.ts)
				if tspath.StartsWithDirectory(importingFileName, realPathDirectory, host.UseCaseSensitiveFileNames()) {
					return false, true // Stop search, each ancestor directory will also hit this condition
				}

				for _, target := range targets {
					if !tspath.StartsWithDirectory(target, realPathDirectory, host.UseCaseSensitiveFileNames()) {
						continue
					}

					relative := tspath.GetRelativePathFromDirectory(
						realPathDirectory,
						target,
						tspath.ComparePathsOptions{
							UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
							CurrentDirectory:          cwd,
						},
					)
					symlinkSet.Range(func(symlinkDirectory string) bool {
						option := tspath.ResolvePath(symlinkDirectory, relative)
						results = append(results, ModulePath{
							FileName:        option,
							IsInNodeModules: ContainsNodeModules(option),
							IsRedirect:      target == referenceRedirect,
						})
						shouldFilterIgnoredPaths = true // We found a non-ignored path in symlinks, so we can reject ignored-path realpaths
						return true
					})
				}

				return false, false
			},
		)
	}

	if preferSymlinks {
		for _, p := range targets {
			if !(shouldFilterIgnoredPaths && containsIgnoredPath(p)) {
				results = append(results, ModulePath{
					FileName:        p,
					IsInNodeModules: ContainsNodeModules(p),
					IsRedirect:      referenceRedirect == p,
				})
			}
		}
	}

	return results
}

func computeModuleSpecifiers(
	modulePaths []ModulePath,
	compilerOptions *core.CompilerOptions,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
	forAutoImport bool,
) ([]string, ResultKind) {
	info := getInfo(importingSourceFile.FileName(), host)
	preferences := getModuleSpecifierPreferences(userPreferences, host, compilerOptions, importingSourceFile, "")

	var existingSpecifier string
	for _, modulePath := range modulePaths {
		targetPath := tspath.ToPath(modulePath.FileName, host.GetCurrentDirectory(), info.UseCaseSensitiveFileNames)
		var existingImport *ast.StringLiteralLike
		for _, importSpecifier := range importingSourceFile.Imports() {
			resolvedModule := host.GetResolvedModuleFromModuleSpecifier(importingSourceFile, importSpecifier)
			if resolvedModule.IsResolved() && tspath.ToPath(resolvedModule.ResolvedFileName, host.GetCurrentDirectory(), info.UseCaseSensitiveFileNames) == targetPath {
				existingImport = importSpecifier
				break
			}
		}
		if existingImport != nil {
			if preferences.relativePreference == RelativePreferenceNonRelative && tspath.PathIsRelative(existingImport.Text()) {
				// If the preference is for non-relative and the module specifier is relative, ignore it
				continue
			}
			existingMode := host.GetModeForUsageLocation(importingSourceFile, existingImport)
			targetMode := options.OverrideImportMode
			if targetMode == core.ResolutionModeNone {
				targetMode = host.GetDefaultResolutionModeForFile(importingSourceFile)
			}
			if existingMode != targetMode && existingMode != core.ResolutionModeNone && targetMode != core.ResolutionModeNone {
				// If the candidate import mode doesn't match the mode we're generating for, don't consider it
				continue
			}
			existingSpecifier = existingImport.Text()
			break
		}
	}

	if existingSpecifier != "" {
		return []string{existingSpecifier}, ResultKindNone
	}

	// A Lua module has one canonical require name; synthesize it from the
	// resolved path. When no name round-trips through resolveLua, emit
	// nothing: the resolver categorically rejects path-shaped specifiers
	// (relative './x', deep 'pkg/lib/util', '#imports'), so producing one
	// would be guaranteed-broken code.
	for _, modulePath := range modulePaths {
		name := tryGetLuaModuleSpecifier(modulePath, info, importingSourceFile, host, compilerOptions, userPreferences, options.OverrideImportMode)
		if name == "" || forAutoImport && IsExcludedByRegex(name, preferences.excludeRegexes) {
			continue
		}
		if modulePath.IsInNodeModules {
			return []string{name}, ResultKindNodeModules
		}
		return []string{name}, ResultKindRelative
	}
	return nil, ResultKindNone
}

// tryGetLuaModuleSpecifier is the single Lua name production for one resolved
// path. For a node_modules path the package-entrypoint name (bare "pkg" via
// package.json types/main) is tried first — it must run before the dotted
// derivation so an entrypoint file is named "pkg", not "pkg.lib.index" — and
// is only accepted when it is itself a Lua module name reachable from the
// search root (a scoped "@x/y" or deep "pkg/lib/util" result is path-shaped
// and unresolvable, so it is discarded). Package-internal files then fall
// back to the dotted path derivation ("pkg.lib.util").
func tryGetLuaModuleSpecifier(
	modulePath ModulePath,
	info Info,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	compilerOptions *core.CompilerOptions,
	userPreferences UserPreferences,
	overrideMode core.ResolutionMode,
) string {
	if modulePath.IsInNodeModules {
		name := tryGetModuleNameAsNodeModule(modulePath, info, importingSourceFile, host, compilerOptions, userPreferences /*packageNameOnly*/, false, overrideMode)
		if name != "" && tspath.IsLuaModuleName(name) && luaPathNodeModulesOnSearchChain(modulePath.FileName, compilerOptions, host) {
			return name
		}
	}
	return tryGetLuaModuleName(modulePath.FileName, compilerOptions, host)
}

// tryGetLuaModuleName maps a resolved module file path to its Lua dotted
// require name: the path relative to the search root (or to a node_modules
// directory on the resolver's ancestor chain), with the extension stripped,
// a trailing /init dropped when no sibling shadows it, and slashes joined
// with dots. It mirrors resolveLua in internal/module — synthesis and
// resolution must round-trip exactly. A file with no round-tripping name
// (outside both roots, a dotted path segment, a declaration file, a shadowed
// init file, a package-internal node_modules) returns "".
func tryGetLuaModuleName(moduleFileName string, compilerOptions *core.CompilerOptions, host ModuleSpecifierGenerationHost) string {
	// A declaration file is a global script, not a module: resolveLua never
	// probes .d.tlua. Checked before RemoveFileExtension, which would only
	// strip the ".tlua" half and leave a dotted ".d" segment by accident.
	if tspath.IsDeclarationFileName(moduleFileName) {
		return ""
	}
	// base is the directory the dotted name is resolved against — the search
	// root, or the node_modules directory the path goes through.
	var base string
	var rel string
	if idx := strings.LastIndex(moduleFileName, "/node_modules/"); idx >= 0 {
		// Only a node_modules directory on the resolver's search chain is
		// reachable; a nested (package-internal) node_modules is not.
		base = moduleFileName[:idx+len("/node_modules")]
		if !luaNodeModulesOnSearchChain(base, compilerOptions, host) {
			return ""
		}
		rel = moduleFileName[idx+len("/node_modules/"):]
	} else {
		base = compilerOptions.GetLuaSearchRoot(host.GetCurrentDirectory())
		rel = tspath.GetRelativePathFromDirectory(base, moduleFileName, tspath.ComparePathsOptions{
			UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
			CurrentDirectory:          host.GetCurrentDirectory(),
		})
		if rel == "" || strings.HasPrefix(rel, "..") || tspath.IsRootedDiskPath(rel) {
			return ""
		}
		rel = strings.TrimPrefix(rel, "./")
	}
	rel = tspath.RemoveFileExtension(rel)
	if withoutInit, hadInit := strings.CutSuffix(rel, "/init"); hadInit {
		// resolveLua probes `?.tlua` and `?.tsx` before `?/init.*`, so the
		// stripped name reaches this init file only when no sibling shadows
		// it. A shadowed init file has no round-tripping name.
		sibling := tspath.CombinePaths(base, withoutInit)
		if host.FileExists(sibling+tspath.ExtensionTs) || host.FileExists(sibling+tspath.ExtensionTsx) {
			return ""
		}
		rel = withoutInit
	}
	return tspath.LuaRelativePathToModuleName(rel)
}

// luaNodeModulesOnSearchChain reports whether nodeModulesDir is one of the
// node_modules directories resolveLua's fallback actually walks: its parent
// must be the Lua search root or an ancestor of the root.
func luaNodeModulesOnSearchChain(nodeModulesDir string, compilerOptions *core.CompilerOptions, host ModuleSpecifierGenerationHost) bool {
	root := compilerOptions.GetLuaSearchRoot(host.GetCurrentDirectory())
	return tspath.ContainsPath(tspath.GetDirectoryPath(nodeModulesDir), root, tspath.ComparePathsOptions{
		UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
		CurrentDirectory:          host.GetCurrentDirectory(),
	})
}

// luaPathNodeModulesOnSearchChain applies luaNodeModulesOnSearchChain to the
// node_modules segment of a file path (false when there is none).
func luaPathNodeModulesOnSearchChain(fileName string, compilerOptions *core.CompilerOptions, host ModuleSpecifierGenerationHost) bool {
	idx := strings.LastIndex(fileName, "/node_modules/")
	if idx < 0 {
		return false
	}
	return luaNodeModulesOnSearchChain(fileName[:idx+len("/node_modules")], compilerOptions, host)
}

func processEnding(
	fileName string,
	allowedEndings []ModuleSpecifierEnding,
	options *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
) string {
	if tspath.FileExtensionIsOneOf(fileName, []string{tspath.ExtensionJson}) {
		return fileName
	}

	noExtension := tspath.RemoveFileExtension(fileName)
	if fileName == noExtension {
		return fileName
	}

	jsPriority := slices.Index(allowedEndings, ModuleSpecifierEndingJsExtension)
	if !tspath.FileExtensionIsOneOf(fileName, []string{tspath.ExtensionDts}) && tspath.FileExtensionIsOneOf(fileName, []string{tspath.ExtensionTs}) && strings.Contains(fileName, ".d.") {
		// `foo.d.json.ts` and the like - remap back to `foo.json`
		if result := TryGetRealFileNameForNonJSDeclarationFileName(fileName); result != "" {
			return result
		}
	}

	switch allowedEndings[0] {
	case ModuleSpecifierEndingMinimal:
		withoutIndex := strings.TrimSuffix(noExtension, "/index")
		if host != nil && withoutIndex != noExtension && tryGetAnyFileFromPath(host, withoutIndex) {
			// Can't remove index if there's a file by the same name as the directory.
			// Probably more callers should pass `host` so we can determine this?
			return noExtension
		}
		return withoutIndex
	case ModuleSpecifierEndingIndex:
		return noExtension
	case ModuleSpecifierEndingJsExtension:
		return noExtension + getJSExtensionForFile(fileName, options)
	case ModuleSpecifierEndingTsExtension:
		// For now, we don't know if this import is going to be type-only, which means we don't
		// know if a .d.ts extension is valid, so use no extension or a .js extension
		if tspath.IsDeclarationFileName(fileName) {
			extensionlessPriority := -1
			for i, e := range allowedEndings {
				if e == ModuleSpecifierEndingMinimal || e == ModuleSpecifierEndingIndex {
					extensionlessPriority = i
					break
				}
			}
			if extensionlessPriority != -1 && extensionlessPriority < jsPriority {
				return noExtension
			}
			return noExtension + getJSExtensionForFile(fileName, options)
		}
		return fileName
	default:
		debug.AssertNever(allowedEndings[0])
		return ""
	}
}

func tryGetModuleNameAsNodeModule(
	pathObj ModulePath,
	info Info,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	options *core.CompilerOptions,
	userPreferences UserPreferences,
	packageNameOnly bool,
	overrideMode core.ResolutionMode,
) string {
	parts := GetNodeModulePathParts(pathObj.FileName)
	if parts == nil {
		return ""
	}

	// Simplify the full file path to something that can be resolved by Node.
	preferences := getModuleSpecifierPreferences(userPreferences, host, options, importingSourceFile, "")
	allowedEndings := preferences.getAllowedEndingsInPreferredOrder(core.ResolutionModeNone)

	caseSensitive := host.UseCaseSensitiveFileNames()
	moduleSpecifier := pathObj.FileName
	isPackageRootPath := false
	if !packageNameOnly {
		packageRootIndex := parts.PackageRootIndex
		var moduleFileName string
		for true {
			// If the module could be imported by a directory name, use that directory's name
			pkgJsonResults := tryDirectoryWithPackageJson(
				*parts,
				pathObj,
				importingSourceFile,
				host,
				overrideMode,
				options,
				allowedEndings,
			)
			moduleFileToTry := pkgJsonResults.moduleFileToTry
			packageRootPath := pkgJsonResults.packageRootPath
			blockedByExports := pkgJsonResults.blockedByExports
			verbatimFromExports := pkgJsonResults.verbatimFromExports
			if blockedByExports {
				return "" // File is under this package.json, but is not publicly exported - there's no way to name it via `node_modules` resolution
			}
			if verbatimFromExports {
				return moduleFileToTry
			}
			//}
			if len(packageRootPath) > 0 {
				moduleSpecifier = packageRootPath
				isPackageRootPath = true
				break
			}
			if len(moduleFileName) == 0 {
				moduleFileName = moduleFileToTry
			}
			// try with next level of directory
			packageRootIndex = core.IndexAfter(pathObj.FileName, "/", packageRootIndex+1)
			if packageRootIndex == -1 {
				moduleSpecifier = processEnding(moduleFileName, allowedEndings, options, host)
				break
			}
		}
	}

	if pathObj.IsRedirect && !isPackageRootPath {
		return ""
	}

	globalTypingsCacheLocation := host.GetGlobalTypingsCacheLocation()
	// Get a path that's relative to node_modules or the importing file's path
	// if node_modules folder is in this folder or any of its parent folders, no need to keep it.
	pathToTopLevelNodeModules := moduleSpecifier[0:parts.TopLevelNodeModulesIndex]

	if !stringutil.HasPrefix(info.SourceDirectory, pathToTopLevelNodeModules, caseSensitive) || len(globalTypingsCacheLocation) > 0 && stringutil.HasPrefix(globalTypingsCacheLocation, pathToTopLevelNodeModules, caseSensitive) {
		return ""
	}

	// If the module was found in @types, get the actual Node package name
	nodeModulesDirectoryName := moduleSpecifier[parts.TopLevelPackageNameIndex+1:]
	return module.GetPackageNameFromTypesPackageName(nodeModulesDirectoryName)
}

type pkgJsonDirAttemptResult struct {
	moduleFileToTry     string
	packageRootPath     string
	blockedByExports    bool
	verbatimFromExports bool
}

func tryDirectoryWithPackageJson(
	parts NodeModulePathParts,
	pathObj ModulePath,
	importingSourceFile SourceFileForSpecifierGeneration,
	host ModuleSpecifierGenerationHost,
	overrideMode core.ResolutionMode,
	options *core.CompilerOptions,
	allowedEndings []ModuleSpecifierEnding,
) pkgJsonDirAttemptResult {
	rootIdx := parts.PackageRootIndex
	if rootIdx == -1 {
		rootIdx = len(pathObj.FileName) // TODO: possible strada bug? -1 in js slice removes characters from the end, in go it panics - js behavior seems unwanted here?
	}
	packageRootPath := pathObj.FileName[0:rootIdx]
	packageJsonPath := tspath.CombinePaths(packageRootPath, "package.json")
	moduleFileToTry := pathObj.FileName
	maybeBlockedByTypesVersions := false
	packageJson := host.GetPackageJsonInfo(packageJsonPath)
	if packageJson == nil {
		// No package.json exists; an index.js will still resolve as the package name
		fileName := moduleFileToTry[parts.PackageRootIndex+1:]
		if fileName == "index"+tspath.ExtensionDts || fileName == "index"+tspath.ExtensionJs ||
			fileName == "index"+tspath.ExtensionTs || fileName == "index"+tspath.ExtensionTsx {
			return pkgJsonDirAttemptResult{moduleFileToTry: moduleFileToTry, packageRootPath: packageRootPath}
		} else {
			return pkgJsonDirAttemptResult{moduleFileToTry: moduleFileToTry}
		}
	}

	importMode := overrideMode
	if importMode == core.ResolutionModeNone {
		importMode = host.GetDefaultResolutionModeForFile(importingSourceFile)
	}

	packageJsonContent := packageJson.GetContents()
	if options.GetResolvePackageJsonExports() {
		// The package name that we found in node_modules could be different from the package
		// name in the package.json content via url/filepath dependency specifiers. We need to
		// use the actual directory name, so don't look at `packageJsonContent.name` here.
		nodeModulesDirectoryName := packageRootPath[parts.TopLevelPackageNameIndex+1:]
		packageName := module.GetPackageNameFromTypesPackageName(nodeModulesDirectoryName)

		conditions := module.GetConditions(options, importMode)

		var fromExports string
		if packageJsonContent != nil && packageJsonContent.Fields.Exports.Type != packagejson.JSONValueTypeNotPresent {
			fromExports = tryGetModuleNameFromExports(
				options,
				host,
				pathObj.FileName,
				packageRootPath,
				packageName,
				packageJsonContent.Fields.Exports,
				conditions,
			)
		}
		if len(fromExports) > 0 {
			return pkgJsonDirAttemptResult{
				moduleFileToTry:     fromExports,
				verbatimFromExports: true,
			}
		}
		if packageJsonContent != nil && packageJsonContent.Fields.Exports.Type != packagejson.JSONValueTypeNotPresent {
			return pkgJsonDirAttemptResult{
				moduleFileToTry:  pathObj.FileName,
				blockedByExports: true,
			}
		}
	}

	var versionPaths packagejson.VersionPaths
	if packageJsonContent != nil && packageJsonContent.TypesVersions.Type == packagejson.JSONValueTypeObject {
		versionPaths = packageJsonContent.GetVersionPaths(nil)
	}
	if versionPaths.GetPaths() != nil {
		subModuleName := pathObj.FileName[len(packageRootPath)+1:]
		fromPaths := tryGetModuleNameFromPaths(
			subModuleName,
			versionPaths.GetPaths(),
			allowedEndings,
			packageRootPath,
			host,
			options,
		)
		if len(fromPaths) == 0 {
			maybeBlockedByTypesVersions = true
		} else {
			moduleFileToTry = tspath.CombinePaths(packageRootPath, fromPaths)
		}
	}
	// If the file is the main module, it can be imported by the package name
	mainFileRelative := "index.js"
	if packageJsonContent != nil {
		if packageJsonContent.Typings.Valid {
			mainFileRelative = packageJsonContent.Typings.Value
		} else if packageJsonContent.Types.Valid {
			mainFileRelative = packageJsonContent.Types.Value
		} else if packageJsonContent.Main.Valid {
			mainFileRelative = packageJsonContent.Main.Value
		}
	}

	if len(mainFileRelative) > 0 && !(maybeBlockedByTypesVersions && module.MatchPatternOrExact(module.TryParsePatterns(versionPaths.GetPaths()), mainFileRelative) != core.Pattern{}) {
		// The 'main' file is also subject to mapping through typesVersions, and we couldn't come up with a path
		// explicitly through typesVersions, so if it matches a key in typesVersions now, it's not reachable.
		// (The only way this can happen is if some file in a package that's not resolvable from outside the
		// package got pulled into the program anyway, e.g. transitively through a file that *is* reachable. It
		// happens very easily in fourslash tests though, since every test file listed gets included. See
		// importNameCodeFix_typesVersions.ts for an example.)
		mainExportFile := tspath.ToPath(mainFileRelative, packageRootPath, host.UseCaseSensitiveFileNames())
		compareOpt := tspath.ComparePathsOptions{
			UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
			CurrentDirectory:          host.GetCurrentDirectory(),
		}
		if tspath.ComparePaths(tspath.RemoveFileExtension(string(mainExportFile)), tspath.RemoveFileExtension(moduleFileToTry), compareOpt) == 0 {
			// ^ An arbitrary removal of file extension for this comparison is almost certainly wrong
			return pkgJsonDirAttemptResult{packageRootPath: packageRootPath, moduleFileToTry: moduleFileToTry}
		} else if packageJsonContent == nil || packageJsonContent.Type.Value != "module" &&
			!tspath.FileExtensionIsOneOf(moduleFileToTry, tspath.ExtensionsNotSupportingExtensionlessResolution) &&
			stringutil.HasPrefix(moduleFileToTry, string(mainExportFile), host.UseCaseSensitiveFileNames()) &&
			tspath.ComparePaths(tspath.GetDirectoryPath(moduleFileToTry), tspath.RemoveTrailingDirectorySeparator(string(mainExportFile)), compareOpt) == 0 &&
			tspath.RemoveFileExtension(tspath.GetBaseFileName(moduleFileToTry)) == "index" {
			// if mainExportFile is a directory, which contains moduleFileToTry, we just try index file
			// example mainExportFile: `pkg/lib` and moduleFileToTry: `pkg/lib/index`, we can use packageRootPath
			// but this behavior is deprecated for packages with "type": "module", so we only do this for packages without "type": "module"
			// and make sure that the extension on index.{???} is something that supports omitting the extension
			return pkgJsonDirAttemptResult{packageRootPath: packageRootPath, moduleFileToTry: moduleFileToTry}
		}
	}

	return pkgJsonDirAttemptResult{moduleFileToTry: moduleFileToTry}
}

func tryGetModuleNameFromExports(
	options *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	targetFilePath string,
	packageDirectory string,
	packageName string,
	exports packagejson.ExportsOrImports,
	conditions []string,
) string {
	if exports.IsSubpaths() {
		// sub-mappings
		// 3 cases:
		// * directory mappings (legacyish, key ends with / (technically allows index/extension resolution under cjs mode))
		// * pattern mappings (contains a *)
		// * exact mappings (no *, does not end with /)
		for k, subk := range exports.AsObject().Entries() {
			subPackageName := tspath.GetNormalizedAbsolutePath(tspath.CombinePaths(packageName, k), "")
			mode := MatchingModeExact
			if strings.HasSuffix(k, "/") {
				mode = MatchingModeDirectory
			} else if strings.Contains(k, "*") {
				mode = MatchingModePattern
			}
			result := tryGetModuleNameFromExportsOrImports(options, host, targetFilePath, packageDirectory, subPackageName, subk, conditions, mode)
			if len(result) > 0 {
				return result
			}
		}
	}
	return tryGetModuleNameFromExportsOrImports(
		options,
		host,
		targetFilePath,
		packageDirectory,
		packageName,
		exports,
		conditions,
		MatchingModeExact,
	)
}

type specPair struct {
	ending ModuleSpecifierEnding
	value  string
}

func tryGetModuleNameFromPaths(
	relativeToBaseUrl string,
	paths *collections.OrderedMap[string, []string],
	allowedEndings []ModuleSpecifierEnding,
	baseDirectory string,
	host ModuleSpecifierGenerationHost,
	compilerOptions *core.CompilerOptions,
) string {
	caseSensitive := host.UseCaseSensitiveFileNames()
	for key, values := range paths.Entries() {
		for _, patternText := range values {
			normalized := tspath.NormalizePath(patternText)
			pattern := getRelativePathIfInSameVolume(normalized, baseDirectory, caseSensitive)
			if len(pattern) == 0 {
				pattern = normalized
			}
			prefix, suffix, ok := strings.Cut(pattern, "*")

			// In module resolution, if `pattern` itself has an extension, a file with that extension is looked up directly,
			// meaning a '.ts' or '.d.ts' extension is allowed to resolve. This is distinct from the case where a '*' substitution
			// causes a module specifier to have an extension, i.e. the extension comes from the module specifier in a JS/TS file
			// and matches the '*'. For example:
			//
			// Module Specifier      | Path Mapping (key: [pattern]) | Interpolation       | Resolution Action
			// ---------------------->------------------------------->--------------------->---------------------------------------------------------------
			// import "@app/foo"    -> "@app/*": ["./src/app/*.ts"] -> "./src/app/foo.ts" -> tryFile("./src/app/foo.ts") || [continue resolution algorithm]
			// import "@app/foo.ts" -> "@app/*": ["./src/app/*"]    -> "./src/app/foo.ts" -> [continue resolution algorithm]
			//
			// (https://github.com/microsoft/TypeScript/blob/ad4ded80e1d58f0bf36ac16bea71bc10d9f09895/src/compiler/moduleNameResolver.ts#L2509-L2516)
			//
			// The interpolation produced by both scenarios is identical, but only in the former, where the extension is encoded in
			// the path mapping rather than in the module specifier, will we prioritize a file lookup on the interpolation result.
			// (In fact, currently, the latter scenario will necessarily fail since no resolution mode recognizes '.ts' as a valid
			// extension for a module specifier.)
			//
			// Here, this means we need to be careful about whether we generate a match from the target filename (typically with a
			// .ts extension) or the possible relative module specifiers representing that file:
			//
			// Filename            | Relative Module Specifier Candidates         | Path Mapping                 | Filename Result    | Module Specifier Results
			// --------------------<----------------------------------------------<------------------------------<-------------------||----------------------------
			// dist/haha.d.ts      <- dist/haha, dist/haha.js                     <- "@app/*": ["./dist/*.d.ts"] <- @app/haha        || (none)
			// dist/haha.d.ts      <- dist/haha, dist/haha.js                     <- "@app/*": ["./dist/*"]      <- (none)           || @app/haha, @app/haha.js
			// dist/foo/index.d.ts <- dist/foo, dist/foo/index, dist/foo/index.js <- "@app/*": ["./dist/*.d.ts"] <- @app/foo/index   || (none)
			// dist/foo/index.d.ts <- dist/foo, dist/foo/index, dist/foo/index.js <- "@app/*": ["./dist/*"]      <- (none)           || @app/foo, @app/foo/index, @app/foo/index.js
			// dist/wow.js.js      <- dist/wow.js, dist/wow.js.js                 <- "@app/*": ["./dist/*.js"]   <- @app/wow.js      || @app/wow, @app/wow.js
			//
			// The "Filename Result" can be generated only if `pattern` has an extension. Care must be taken that the list of
			// relative module specifiers to run the interpolation (a) is actually valid for the module resolution mode, (b) takes
			// into account the existence of other files (e.g. 'dist/wow.js' cannot refer to 'dist/wow.js.js' if 'dist/wow.js'
			// exists) and (c) that they are ordered by preference. The last row shows that the filename result and module
			// specifier results are not mutually exclusive. Note that the filename result is a higher priority in module
			// resolution, but as long criteria (b) above is met, I don't think its result needs to be the highest priority result
			// in module specifier generation. I have included it last, as it's difficult to tell exactly where it should be
			// sorted among the others for a particular value of `importModuleSpecifierEnding`.

			var candidates []specPair
			for _, ending := range allowedEndings {
				result := processEnding(
					relativeToBaseUrl,
					[]ModuleSpecifierEnding{ending},
					compilerOptions,
					host,
				)
				candidates = append(candidates, specPair{
					ending: ending,
					value:  result,
				})
			}
			if len(tspath.TryGetExtensionFromPath(pattern)) > 0 {
				candidates = append(candidates, specPair{
					ending: ModuleSpecifierEndingJsExtension,
					value:  relativeToBaseUrl,
				})
			}

			if ok {
				for _, c := range candidates {
					value := c.value
					if len(value) >= len(prefix)+len(suffix) &&
						stringutil.HasPrefix(value, prefix, caseSensitive) && // TODO: possible strada bug: these are not case-switched in strada
						stringutil.HasSuffix(value, suffix, caseSensitive) &&
						validateEnding(c, relativeToBaseUrl, compilerOptions, host) {
						matchedStar := value[len(prefix) : len(value)-len(suffix)]
						if !tspath.PathIsRelative(matchedStar) {
							return replaceFirstStar(key, matchedStar)
						}
					}
				}
			} else if core.Some(candidates, func(c specPair) bool { return c.ending != ModuleSpecifierEndingMinimal && pattern == c.value }) ||
				core.Some(candidates, func(c specPair) bool {
					return c.ending == ModuleSpecifierEndingMinimal && pattern == c.value && validateEnding(c, relativeToBaseUrl, compilerOptions, host)
				}) {
				return key
			}
		}
	}
	return ""
}

func validateEnding(c specPair, relativeToBaseUrl string, compilerOptions *core.CompilerOptions, host ModuleSpecifierGenerationHost) bool {
	// Optimization: `removeExtensionAndIndexPostFix` can query the file system (a good bit) if `ending` is `Minimal`, the basename
	// is 'index', and a `host` is provided. To avoid that until it's unavoidable, we ran the function with no `host` above. Only
	// here, after we've checked that the minimal ending is indeed a match (via the length and prefix/suffix checks / `some` calls),
	// do we check that the host-validated result is consistent with the answer we got before. If it's not, it falls back to the
	// `ModuleSpecifierEnding.Index` result, which should already be in the list of candidates if `Minimal` was. (Note: the assumption here is
	// that every module resolution mode that supports dropping extensions also supports dropping `/index`. Like literally
	// everything else in this file, this logic needs to be updated if that's not true in some future module resolution mode.)
	return c.ending != ModuleSpecifierEndingMinimal || c.value == processEnding(relativeToBaseUrl, []ModuleSpecifierEnding{c.ending}, compilerOptions, host)
}

func tryGetModuleNameFromExportsOrImports(
	options *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	targetFilePath string,
	packageDirectory string,
	packageName string,
	exports packagejson.ExportsOrImports,
	conditions []string,
	mode MatchingMode,
) string {
	switch exports.Type {
	case packagejson.JSONValueTypeNotPresent:
		return ""
	case packagejson.JSONValueTypeString:
		strValue := exports.Value.(string)

		pathOrPattern := tspath.GetNormalizedAbsolutePath(tspath.CombinePaths(packageDirectory, strValue), "")
		var extensionSwappedTarget string
		if tspath.HasTSFileExtension(targetFilePath) {
			extensionSwappedTarget = tspath.RemoveFileExtension(targetFilePath) + module.TryGetJSExtensionForFile(targetFilePath, options)
		}

		compareOpts := tspath.ComparePathsOptions{
			UseCaseSensitiveFileNames: host.UseCaseSensitiveFileNames(),
			CurrentDirectory:          host.GetCurrentDirectory(),
		}

		switch mode {
		case MatchingModeExact:
			if len(extensionSwappedTarget) > 0 && tspath.ComparePaths(extensionSwappedTarget, pathOrPattern, compareOpts) == 0 ||
				tspath.ComparePaths(targetFilePath, pathOrPattern, compareOpts) == 0 {
				return packageName
			}
		case MatchingModeDirectory:
			if len(extensionSwappedTarget) > 0 && tspath.ContainsPath(pathOrPattern, extensionSwappedTarget, compareOpts) {
				fragment := tspath.GetRelativePathFromDirectory(pathOrPattern, extensionSwappedTarget, compareOpts)
				return tspath.GetNormalizedAbsolutePath(tspath.CombinePaths(tspath.CombinePaths(packageName, strValue), fragment), "")
			}
			if tspath.ContainsPath(pathOrPattern, targetFilePath, compareOpts) {
				fragment := tspath.GetRelativePathFromDirectory(pathOrPattern, targetFilePath, compareOpts)
				return tspath.GetNormalizedAbsolutePath(tspath.CombinePaths(tspath.CombinePaths(packageName, strValue), fragment), "")
			}
		case MatchingModePattern:
			leadingSlice, trailingSlice, _ := strings.Cut(pathOrPattern, "*")
			caseSensitive := host.UseCaseSensitiveFileNames()
			if len(extensionSwappedTarget) > 0 && stringutil.HasPrefixAndSuffixWithoutOverlap(extensionSwappedTarget, leadingSlice, trailingSlice, caseSensitive) {
				starReplacement := extensionSwappedTarget[len(leadingSlice) : len(extensionSwappedTarget)-len(trailingSlice)]
				return replaceFirstStar(packageName, starReplacement)
			}
			if stringutil.HasPrefixAndSuffixWithoutOverlap(targetFilePath, leadingSlice, trailingSlice, caseSensitive) {
				starReplacement := targetFilePath[len(leadingSlice) : len(targetFilePath)-len(trailingSlice)]
				return replaceFirstStar(packageName, starReplacement)
			}
		}
		return ""
	case packagejson.JSONValueTypeArray:
		arr := exports.AsArray()
		for _, e := range arr {
			result := tryGetModuleNameFromExportsOrImports(options, host, targetFilePath, packageDirectory, packageName, e, conditions, mode)
			if len(result) > 0 {
				return result
			}
		}
	case packagejson.JSONValueTypeObject:
		// conditional mapping
		obj := exports.AsObject()
		for key, value := range obj.Entries() {
			if key == "default" || slices.Contains(conditions, key) || slices.Contains(conditions, "types") && module.IsApplicableVersionedTypesKey(key) {
				result := tryGetModuleNameFromExportsOrImports(options, host, targetFilePath, packageDirectory, packageName, value, conditions, mode)
				if len(result) > 0 {
					return result
				}
			}
		}
	case packagejson.JSONValueTypeNull:
		return ""
	}
	return ""
}

// `importingSourceFile` and `importingSourceFileName`? Why not just use `importingSourceFile.path`?
// Because when this is called by the declaration emitter, `importingSourceFile` is the implementation
// file, but `importingSourceFileName` and `toFileName` refer to declaration files (the former to the
// one currently being produced; the latter to the one being imported). We need an implementation file
// just to get its `impliedNodeFormat` and to detect certain preferences from existing import module
// specifiers.
func GetModuleSpecifier(
	compilerOptions *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	importingSourceFile *ast.SourceFile, // !!! | FutureSourceFile
	importingSourceFileName string,
	oldImportSpecifier string, // used only in updatingModuleSpecifier
	toFileName string,
	options ModuleSpecifierOptions,
) string {
	return getModuleSpecifierWithPreferences(
		compilerOptions,
		host,
		importingSourceFile,
		importingSourceFileName,
		oldImportSpecifier,
		toFileName,
		UserPreferences{},
		options,
	)
}

func UpdateModuleSpecifier(
	compilerOptions *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	importingSourceFile *ast.SourceFile,
	importingSourceFileName string,
	oldImportSpecifier string,
	toFileName string,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
) string {
	return getModuleSpecifierWithPreferences(
		compilerOptions,
		host,
		importingSourceFile,
		importingSourceFileName,
		oldImportSpecifier,
		toFileName,
		userPreferences,
		options,
	)
}

func getModuleSpecifierWithPreferences(
	compilerOptions *core.CompilerOptions,
	host ModuleSpecifierGenerationHost,
	importingSourceFile *ast.SourceFile, // !!! | FutureSourceFile
	importingSourceFileName string,
	oldImportSpecifier string, // used only in updatingModuleSpecifier
	toFileName string,
	userPreferences UserPreferences,
	options ModuleSpecifierOptions,
) string {
	info := getInfo(importingSourceFileName, host)
	modulePaths := getAllModulePaths(info, toFileName, host, compilerOptions, userPreferences, options)

	// Same production as computeModuleSpecifiers: the canonical Lua name or
	// nothing — a relative or path-shaped specifier would never resolve.
	for _, modulePath := range modulePaths {
		if name := tryGetLuaModuleSpecifier(modulePath, info, importingSourceFile, host, compilerOptions, userPreferences, options.OverrideImportMode); name != "" {
			return name
		}
	}
	return ""
}

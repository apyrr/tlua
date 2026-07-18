package autoimport

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/binder"
	"github.com/apyrr/tlua/internal/checker"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/module"
	"github.com/apyrr/tlua/internal/packagejson"
	"github.com/apyrr/tlua/internal/symlinks"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tspath"
)

type pathAndFileName struct {
	path     tspath.Path
	fileName string
}

type aliasResolver struct {
	toPath         func(fileName string) tspath.Path
	host           RegistryCloneHost
	moduleResolver *module.Resolver

	rootFiles []*ast.SourceFile
	// symlinks maps from realpath to symlinked path and file name
	symlinks                    map[tspath.Path]pathAndFileName
	onFailedAmbientModuleLookup func(source ast.HasFileName, moduleName string)
	resolvedModules             collections.SyncMap[tspath.Path, *collections.SyncMap[module.ModeAwareCacheKey, *module.ResolvedModule]]
}

func newAliasResolver(
	rootFiles []*ast.SourceFile,
	symlinks map[tspath.Path]pathAndFileName,
	host RegistryCloneHost,
	moduleResolver *module.Resolver,
	toPath func(fileName string) tspath.Path,
	onFailedAmbientModuleLookup func(source ast.HasFileName, moduleName string),
) *aliasResolver {
	r := &aliasResolver{
		toPath:                      toPath,
		host:                        host,
		moduleResolver:              moduleResolver,
		rootFiles:                   rootFiles,
		symlinks:                    symlinks,
		onFailedAmbientModuleLookup: onFailedAmbientModuleLookup,
	}
	return r
}

// BindSourceFiles implements checker.Program.
func (r *aliasResolver) BindSourceFiles() {
	// We will bind as we parse
}

// SourceFiles implements checker.Program.
func (r *aliasResolver) SourceFiles() []*ast.SourceFile {
	return r.rootFiles
}

// Options implements checker.Program.
func (r *aliasResolver) Options() *core.CompilerOptions {
	return &core.CompilerOptions{
		NoCheck: core.TSTrue,
	}
}

// GetCurrentDirectory implements checker.Program.
func (r *aliasResolver) GetCurrentDirectory() string {
	return r.host.GetCurrentDirectory()
}

// UseCaseSensitiveFileNames implements checker.Program.
func (r *aliasResolver) UseCaseSensitiveFileNames() bool {
	return r.host.FS().UseCaseSensitiveFileNames()
}

// GetSourceFile implements checker.Program.
func (r *aliasResolver) GetSourceFile(fileName string) *ast.SourceFile {
	file := r.host.GetSourceFile(fileName, r.toPath(fileName))
	// file may be nil due to symlink/realpath mismatch; see TestAutoImportBuilderFS
	if file == nil {
		return nil
	}
	binder.BindSourceFile(file)
	return file
}

// GetDefaultResolutionModeForFile implements checker.Program.
func (r *aliasResolver) GetDefaultResolutionModeForFile(file ast.HasFileName) core.ResolutionMode {
	return core.ModuleKindESNext
}

// GetEmitModuleFormatOfFile implements checker.Program.
func (r *aliasResolver) GetEmitModuleFormatOfFile(sourceFile ast.HasFileName) core.ModuleKind {
	return core.ModuleKindESNext
}

// GetEmitSyntaxForUsageLocation implements checker.Program.
func (r *aliasResolver) GetEmitSyntaxForUsageLocation(sourceFile ast.HasFileName, usageLocation *ast.StringLiteralLike) core.ResolutionMode {
	return core.ModuleKindESNext
}

// GetImpliedNodeFormatForEmit implements checker.Program.
func (r *aliasResolver) GetImpliedNodeFormatForEmit(sourceFile ast.HasFileName) core.ModuleKind {
	return core.ModuleKindESNext
}

// GetModeForUsageLocation implements checker.Program.
func (r *aliasResolver) GetModeForUsageLocation(file ast.HasFileName, moduleSpecifier *ast.StringLiteralLike) core.ResolutionMode {
	return core.ModuleKindESNext
}

// GetResolvedModule implements checker.Program.
func (r *aliasResolver) GetResolvedModule(currentSourceFile ast.HasFileName, moduleReference string, mode core.ResolutionMode) *module.ResolvedModule {
	cache, _ := r.resolvedModules.LoadOrStore(currentSourceFile.Path(), &collections.SyncMap[module.ModeAwareCacheKey, *module.ResolvedModule]{})
	if resolved, ok := cache.Load(module.ModeAwareCacheKey{Name: moduleReference, Mode: mode}); ok {
		return resolved
	}
	// The aliasResolver is a synthetic program with no project-reference map, so
	// a fresh resolution would anchor a file that belongs to a referenced project
	// at the top program's root. Anchor at the file's owning program instead:
	// reuse its precomputed resolution when it already resolved this import, and
	// otherwise resolve fresh with that file's redirect.
	var resolved *module.ResolvedModule
	if _, program := r.host.GetDefaultProject(currentSourceFile.Path()); program != nil {
		resolved = program.GetResolvedModule(currentSourceFile, moduleReference, mode)
		if resolved == nil {
			var redirect module.ResolvedProjectReference
			if red := program.GetRedirectForResolution(currentSourceFile); red != nil {
				redirect = red
			}
			resolved, _ = r.moduleResolver.ResolveModuleName(moduleReference, currentSourceFile.FileName(), mode, redirect)
		}
	} else {
		resolved, _ = r.moduleResolver.ResolveModuleName(moduleReference, currentSourceFile.FileName(), mode, nil)
	}
	resolved, _ = cache.LoadOrStore(module.ModeAwareCacheKey{Name: moduleReference, Mode: mode}, resolved)
	if !resolved.IsResolved() && !tspath.PathIsRelative(moduleReference) {
		r.onFailedAmbientModuleLookup(currentSourceFile, moduleReference)
	}
	return resolved
}

// GetSourceFileForResolvedModule implements checker.Program.
func (r *aliasResolver) GetSourceFileForResolvedModule(fileName string) *ast.SourceFile {
	return r.GetSourceFile(fileName)
}

// GetResolvedModules implements checker.Program.
func (r *aliasResolver) GetResolvedModules() map[tspath.Path]module.ModeAwareCache[*module.ResolvedModule] {
	// only used when producing diagnostics, which hopefully the checker won't do
	return nil
}

// ---

// GetSymlinkCache implements checker.Program.
func (r *aliasResolver) GetSymlinkCache() *symlinks.KnownSymlinks {
	panic("unimplemented")
}

// GetSourceFileMetaData implements checker.Program.
func (r *aliasResolver) GetSourceFileMetaData(path tspath.Path) ast.SourceFileMetaData {
	panic("unimplemented")
}

// CommonSourceDirectory implements checker.Program.
func (r *aliasResolver) CommonSourceDirectory() string {
	panic("unimplemented")
}

// FileExists implements checker.Program.
func (r *aliasResolver) FileExists(fileName string) bool {
	panic("unimplemented")
}

// GetGlobalTypingsCacheLocation implements checker.Program.
func (r *aliasResolver) GetGlobalTypingsCacheLocation() string {
	panic("unimplemented")
}

// GetImportHelpersImportSpecifier implements checker.Program.
func (r *aliasResolver) GetImportHelpersImportSpecifier(path tspath.Path) *ast.Node {
	panic("unimplemented")
}

// GetJSXRuntimeImportSpecifier implements checker.Program.
func (r *aliasResolver) GetJSXRuntimeImportSpecifier(path tspath.Path) (moduleReference string, specifier *ast.Node) {
	panic("unimplemented")
}

// GetNearestAncestorDirectoryWithPackageJson implements checker.Program.
func (r *aliasResolver) GetNearestAncestorDirectoryWithPackageJson(dirname string) string {
	panic("unimplemented")
}

// GetPackageJsonInfo implements checker.Program.
func (r *aliasResolver) GetPackageJsonInfo(pkgJsonPath string) *packagejson.InfoCacheEntry {
	panic("unimplemented")
}

// GetProjectReferenceFromOutputDts implements checker.Program.
func (r *aliasResolver) GetProjectReferenceFromOutputDts(path tspath.Path) *tsoptions.SourceOutputAndProjectReference {
	panic("unimplemented")
}

// GetProjectReferenceFromSource implements checker.Program.
func (r *aliasResolver) GetProjectReferenceFromSource(path tspath.Path) *tsoptions.SourceOutputAndProjectReference {
	panic("unimplemented")
}

// GetRedirectForResolution implements checker.Program.
func (r *aliasResolver) GetRedirectForResolution(file ast.HasFileName) *tsoptions.ParsedCommandLine {
	panic("unimplemented")
}

// GetRedirectTargets implements checker.Program.
func (r *aliasResolver) GetRedirectTargets(path tspath.Path) []string {
	panic("unimplemented")
}

// GetResolvedModuleFromModuleSpecifier implements checker.Program.
func (r *aliasResolver) GetResolvedModuleFromModuleSpecifier(file ast.HasFileName, moduleSpecifier *ast.StringLiteralLike) *module.ResolvedModule {
	panic("unimplemented")
}

// GetSourceOfProjectReferenceIfOutputIncluded implements checker.Program.
func (r *aliasResolver) GetSourceOfProjectReferenceIfOutputIncluded(file ast.HasFileName) string {
	panic("unimplemented")
}

// IsSourceFileDefaultLibrary implements checker.Program.
func (r *aliasResolver) IsSourceFileDefaultLibrary(path tspath.Path) bool {
	return false
}

// IsSourceFromProjectReference implements checker.Program.
func (r *aliasResolver) IsSourceFromProjectReference(path tspath.Path) bool {
	panic("unimplemented")
}

// SourceFileMayBeEmitted implements checker.Program.
func (r *aliasResolver) SourceFileMayBeEmitted(sourceFile *ast.SourceFile, forceDtsEmit bool) bool {
	panic("unimplemented")
}

func (r *aliasResolver) GetPackagesMap() map[string]bool {
	return nil
}

var _ checker.Program = (*aliasResolver)(nil)

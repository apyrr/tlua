package modulespecifiers

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/module"
	"github.com/apyrr/tlua/internal/packagejson"
	"github.com/apyrr/tlua/internal/symlinks"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tspath"
)

// Mock host for testing
type mockModuleSpecifierGenerationHost struct {
	currentDir                string
	useCaseSensitiveFileNames bool
	symlinkCache              *symlinks.KnownSymlinks
	// existingFiles, when non-nil, is the exact set of files FileExists
	// reports; when nil every path exists (legacy behavior).
	existingFiles map[string]bool
}

func (h *mockModuleSpecifierGenerationHost) GetCurrentDirectory() string {
	return h.currentDir
}

func (h *mockModuleSpecifierGenerationHost) UseCaseSensitiveFileNames() bool {
	return h.useCaseSensitiveFileNames
}

func (h *mockModuleSpecifierGenerationHost) GetSymlinkCache() *symlinks.KnownSymlinks {
	return h.symlinkCache
}

func (h *mockModuleSpecifierGenerationHost) ResolveModuleName(moduleName string, containingFile string, resolutionMode core.ResolutionMode) *module.ResolvedModule {
	return nil
}

func (h *mockModuleSpecifierGenerationHost) GetGlobalTypingsCacheLocation() string {
	return ""
}

func (h *mockModuleSpecifierGenerationHost) CommonSourceDirectory() string {
	return h.currentDir
}

func (h *mockModuleSpecifierGenerationHost) GetProjectReferenceFromSource(path tspath.Path) *tsoptions.SourceOutputAndProjectReference {
	return nil
}

func (h *mockModuleSpecifierGenerationHost) GetRedirectTargets(path tspath.Path) []string {
	return nil
}

func (h *mockModuleSpecifierGenerationHost) GetSourceOfProjectReferenceIfOutputIncluded(file ast.HasFileName) string {
	return file.FileName()
}

func (h *mockModuleSpecifierGenerationHost) FileExists(path string) bool {
	if h.existingFiles != nil {
		return h.existingFiles[path]
	}
	return true
}

func (h *mockModuleSpecifierGenerationHost) GetPackageJsonInfo(pkgJsonPath string) *packagejson.InfoCacheEntry {
	return nil
}

func (h *mockModuleSpecifierGenerationHost) GetDefaultResolutionModeForFile(file ast.HasFileName) core.ResolutionMode {
	return core.ResolutionModeNone
}

func (h *mockModuleSpecifierGenerationHost) GetResolvedModuleFromModuleSpecifier(file ast.HasFileName, moduleSpecifier *ast.StringLiteralLike) *module.ResolvedModule {
	return nil
}

func (h *mockModuleSpecifierGenerationHost) GetModeForUsageLocation(file ast.HasFileName, moduleSpecifier *ast.StringLiteralLike) core.ResolutionMode {
	return core.ResolutionModeNone
}

// TestTryGetLuaModuleName locks the synthesis side of the name<->file
// round-trip: every non-empty result must be a name resolveLua maps back to
// exactly the input file, and files the resolver cannot reach (or that a
// probe earlier in the package.path ladder would shadow) get no name.
func TestTryGetLuaModuleName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		file     string
		rootDir  string   // CompilerOptions.RootDir; "" anchors at the current directory
		files    []string // files FileExists reports (the file itself is implied)
		expected string
	}{
		{
			name:     "plain root file",
			file:     "/project/foo.tlua",
			expected: "foo",
		},
		{
			name:     "nested directory",
			file:     "/project/a/b/c.tlua",
			expected: "a.b.c",
		},
		{
			name:     "tsx file",
			file:     "/project/a/b.tsx",
			expected: "a.b",
		},
		{
			name:     "init file without shadowing sibling",
			file:     "/project/a/b/init.tlua",
			expected: "a.b",
		},
		{
			name:     "init file shadowed by .tlua sibling",
			file:     "/project/a/b/init.tlua",
			files:    []string{"/project/a/b.tlua"},
			expected: "",
		},
		{
			name:     "init file shadowed by .tsx sibling",
			file:     "/project/a/b/init.tlua",
			files:    []string{"/project/a/b.tsx"},
			expected: "",
		},
		{
			name: "root-level init file keeps its own name",
			file: "/project/init.tlua",
			// No "/init" suffix to strip: "init" probes ?.tlua first and
			// finds this very file.
			expected: "init",
		},
		{
			name:     "dotted path segment cannot round-trip",
			file:     "/project/a.b/c.tlua",
			expected: "",
		},
		{
			name:     "dotted file name cannot round-trip",
			file:     "/project/a/b.c.tlua",
			expected: "",
		},
		{
			name:     "declaration file is not a module",
			file:     "/project/foo.d.tlua",
			expected: "",
		},
		{
			name:     "node_modules package file",
			file:     "/project/node_modules/pkg/util.tlua",
			expected: "pkg.util",
		},
		{
			name:     "node_modules package init file",
			file:     "/project/node_modules/pkg/init.tlua",
			expected: "pkg",
		},
		{
			name:     "node_modules init shadowed by sibling",
			file:     "/project/node_modules/pkg/init.tlua",
			files:    []string{"/project/node_modules/pkg.tlua"},
			expected: "",
		},
		{
			name:     "nested package-internal node_modules is unreachable",
			file:     "/project/node_modules/pkg/node_modules/dep/x.tlua",
			expected: "",
		},
		{
			name:     "node_modules off the search-root ancestor chain",
			file:     "/project/packages/a/node_modules/pkg/x.tlua",
			expected: "",
		},
		{
			name:     "node_modules at an ancestor of the search root",
			file:     "/project/node_modules/pkg/x.tlua",
			rootDir:  "/project/src",
			expected: "pkg.x",
		},
		{
			name:     "file outside the search root",
			file:     "/elsewhere/foo.tlua",
			expected: "",
		},
		{
			name:     "file above a rootDir search root",
			file:     "/project/foo.tlua",
			rootDir:  "/project/src",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			existing := map[string]bool{tt.file: true}
			for _, f := range tt.files {
				existing[f] = true
			}
			host := &mockModuleSpecifierGenerationHost{
				currentDir:                "/project",
				useCaseSensitiveFileNames: true,
				existingFiles:             existing,
			}
			options := &core.CompilerOptions{RootDir: tt.rootDir}
			if got := tryGetLuaModuleName(tt.file, options, host); got != tt.expected {
				t.Errorf("tryGetLuaModuleName(%q) = %q, expected %q", tt.file, got, tt.expected)
			}
		})
	}
}

func TestGetEachFileNameOfModule(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		importingFile  string
		importedFile   string
		preferSymlinks bool
		expectedCount  int
		expectedPaths  []string
	}{
		{
			name:           "basic file path",
			importingFile:  "/project/src/main.tlua",
			importedFile:   "/project/lib/utils.tlua",
			preferSymlinks: false,
			expectedCount:  1,
			expectedPaths:  []string{"/project/lib/utils.tlua"},
		},
		{
			name:           "symlink preference false",
			importingFile:  "/project/src/main.tlua",
			importedFile:   "/project/lib/utils.tlua",
			preferSymlinks: false,
			expectedCount:  1,
		},
		{
			name:           "symlink preference true",
			importingFile:  "/project/src/main.tlua",
			importedFile:   "/project/lib/utils.tlua",
			preferSymlinks: true,
			expectedCount:  1,
		},
		{
			name:           "ignored path with no alternatives",
			importingFile:  "/project/src/main.tlua",
			importedFile:   "/project/node_modules/.pnpm/file.tlua",
			preferSymlinks: false,
			expectedCount:  1, // Should return 1 because there's no better option (all paths are ignored)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			host := &mockModuleSpecifierGenerationHost{
				currentDir:                "/project",
				useCaseSensitiveFileNames: true,
				symlinkCache:              symlinks.NewKnownSymlink("/project", true),
			}

			result := GetEachFileNameOfModule(tt.importingFile, tt.importedFile, host, tt.preferSymlinks)

			if len(result) != tt.expectedCount {
				t.Errorf("Expected %d paths, got %d", tt.expectedCount, len(result))
			}

			if tt.expectedPaths != nil {
				for i, expectedPath := range tt.expectedPaths {
					if i >= len(result) {
						t.Errorf("Expected path %d: %s, but result has only %d paths", i, expectedPath, len(result))
						continue
					}
					if result[i].FileName != expectedPath {
						t.Errorf("Expected path %d to be %s, got %s", i, expectedPath, result[i].FileName)
					}
				}
			}

			for i, path := range result {
				if path.FileName == "" {
					t.Errorf("Path %d has empty FileName", i)
				}
			}
		})
	}
}

func TestGetEachFileNameOfModuleWithSymlinks(t *testing.T) {
	t.Parallel()
	host := &mockModuleSpecifierGenerationHost{
		currentDir:                "/project",
		useCaseSensitiveFileNames: true,
		symlinkCache:              symlinks.NewKnownSymlink("/project", true),
	}

	symlinkPath := tspath.ToPath("/project/symlink", "/project", true).EnsureTrailingDirectorySeparator()
	realDirectory := &symlinks.KnownDirectoryLink{
		Real:     "/real/path/",
		RealPath: tspath.ToPath("/real/path", "/project", true).EnsureTrailingDirectorySeparator(),
	}
	host.symlinkCache.SetDirectory("/project/symlink", symlinkPath, realDirectory)

	result := GetEachFileNameOfModule("/project/src/main.tlua", "/real/path/file.tlua", host, true)

	// Should find the symlink path
	found := false
	for _, path := range result {
		if path.FileName == "/project/symlink/file.tlua" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected to find symlink path /project/symlink/file.tlua")
	}
}

func TestContainsNodeModules(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "contains node_modules",
			path:     "/project/node_modules/lodash/index.lua",
			expected: true,
		},
		{
			name:     "does not contain node_modules",
			path:     "/project/src/utils.tlua",
			expected: false,
		},
		{
			name:     "node_modules in middle",
			path:     "/project/packages/node_modules/pkg/file.lua",
			expected: true,
		},
		{
			name:     "empty path",
			path:     "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := ContainsNodeModules(tt.path)
			if result != tt.expected {
				t.Errorf("ContainsNodeModules(%q) = %v, expected %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestContainsIgnoredPath(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "ignored path",
			path:     "/project/node_modules/.pnpm/file.tlua",
			expected: true,
		},
		{
			name:     "not ignored path",
			path:     "/project/src/file.tlua",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := containsIgnoredPath(tt.path)
			if result != tt.expected {
				t.Errorf("containsIgnoredPath(%q) = %v, expected %v", tt.path, result, tt.expected)
			}
		})
	}
}

func TestTryGetRealFileNameForNonJSDeclarationFileName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		fileName string
		expected string
	}{
		{
			name:     "json declaration file",
			fileName: "/project/foo.d.json.tlua",
			expected: "/project/foo.json",
		},
		{
			name:     "multi-dot source extension declaration file",
			fileName: "/project/foo.module.d.css.tlua",
			expected: "/project/foo.module.css",
		},
		{
			name:     "plain dts file ignored",
			fileName: "/project/foo.d.tlua",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := TryGetRealFileNameForNonJSDeclarationFileName(tt.fileName); got != tt.expected {
				t.Errorf("TryGetRealFileNameForNonJSDeclarationFileName(%q) = %q, expected %q", tt.fileName, got, tt.expected)
			}
		})
	}
}

func TestTryGetModuleNameFromExportsOrImports(t *testing.T) {
	t.Parallel()
	t.Run("with exports pattern", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			name           string
			targetFilePath string
			expected       string
		}{
			{
				name:           "match",
				targetFilePath: "/pkg/src/things/thing1/index.tlua",
				expected:       "./src/things/thing1",
			},
			{
				name:           "mismatch with matching leading and trailing strings",
				targetFilePath: "/pkg/src/things/index.tlua",
				expected:       "",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()
				result := tryGetModuleNameFromExportsOrImports(
					&core.CompilerOptions{},
					&mockModuleSpecifierGenerationHost{},
					tt.targetFilePath,
					"/pkg",
					"./src/things/*",
					packagejson.ExportsOrImports{
						JSONValue: packagejson.JSONValue{
							Type:  packagejson.JSONValueTypeString,
							Value: "./src/things/*/index.lua",
						},
					},
					[]string{},
					MatchingModePattern,
				)
				if result != tt.expected {
					t.Errorf("tryGetModuleNameFromExportsOrImports(targetFilePath = %q) = %v, expected %v", tt.targetFilePath, result, tt.expected)
				}
			})
		}
	})
}

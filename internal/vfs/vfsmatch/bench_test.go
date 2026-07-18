package vfsmatch

import (
	"testing"

	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/cachedvfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

// Benchmark test cases using the same hosts as the unit tests

func BenchmarkReadDirectory(b *testing.B) {
	benchCases := []struct {
		name       string
		host       func() vfs.FS
		path       string
		extensions []string
		excludes   []string
		includes   []string
	}{
		{
			name:       "LiteralIncludes",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"a.tlua", "b.tlua"},
		},
		{
			name:       "WildcardIncludes",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"z/*.tlua", "x/*.tlua"},
		},
		{
			name:       "RecursiveWildcard",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"**/a.tlua"},
		},
		{
			name:       "RecursiveWithExcludes",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			excludes:   []string{"**/b.tlua"},
			includes:   []string{"**/*.tlua"},
		},
		{
			name:       "ComplexPattern",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			excludes:   []string{"*.tlua", "z/??z.tlua", "*/b.tlua"},
			includes:   []string{"a.tlua", "b.tlua", "z/a.tlua", "z/abz.tlua", "z/aba.tlua", "x/b.tlua"},
		},
		{
			name:       "DottedFolders",
			host:       dottedFoldersHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"**/.*/*"},
		},
		{
			name:       "CommonPackageFolders",
			host:       commonFoldersHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"**/a.tlua"},
		},
		{
			name:       "NoIncludes",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
		},
		{
			name:       "MultipleRecursive",
			host:       caseInsensitiveHost,
			path:       "/dev",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"**/x/**/*"},
		},
		{
			name:       "LargeFileSystem",
			host:       largeFileSystemHost,
			path:       "/project",
			extensions: []string{".tlua", ".tsx", ".d.tlua"},
			includes:   []string{"src/**/*.tlua"},
			excludes:   []string{"**/node_modules/**", "**/*.test.tlua"},
		},
		{
			name:       "LargeAllFiles",
			host:       largeFileSystemHost,
			path:       "/project",
			extensions: []string{".tlua", ".tsx", ".lua"},
			excludes:   []string{"**/node_modules/**"},
			includes:   []string{"**/*"},
		},
	}

	for _, bc := range benchCases {
		b.Run(bc.name, func(b *testing.B) {
			host := cachedvfs.From(bc.host())
			b.ReportAllocs()
			for b.Loop() {
				matchFiles(bc.path, bc.extensions, bc.excludes, bc.includes, host.UseCaseSensitiveFileNames(), "/", UnlimitedDepth, host)
			}
		})
	}
}

// largeFileSystemHost creates a more realistic file system with many files
func largeFileSystemHost() vfs.FS {
	files := make(map[string]string)

	// Create a realistic project structure
	dirs := []string{
		"/project/src",
		"/project/src/components",
		"/project/src/utils",
		"/project/src/services",
		"/project/src/models",
		"/project/src/hooks",
		"/project/test",
		"/project/node_modules/react",
		"/project/node_modules/typescript",
		"/project/node_modules/@types/node",
	}

	// Add files to each directory
	for _, dir := range dirs {
		for j := range 20 {
			files[dir+"/file"+string(rune('a'+j))+".tlua"] = ""
			files[dir+"/file"+string(rune('a'+j))+".test.tlua"] = ""
		}
	}

	// Add some dotted directories
	files["/project/src/.hidden/secret.tlua"] = ""
	files["/project/.config/settings.tlua"] = ""

	return vfstest.FromMap(files, false)
}

// BenchmarkPatternCompilation benchmarks the pattern compilation step
func BenchmarkPatternCompilation(b *testing.B) {
	patterns := []struct {
		name string
		spec string
	}{
		{"Literal", "src/file.tlua"},
		{"SingleWildcard", "src/*.tlua"},
		{"QuestionMark", "src/?.tlua"},
		{"DoubleAsterisk", "**/file.tlua"},
		{"Complex", "src/**/components/*.tsx"},
		{"DottedPattern", "**/.*/*"},
	}

	for _, p := range patterns {
		b.Run(p.name, func(b *testing.B) {
			for b.Loop() {
				_, _ = compileGlobPattern(p.spec, "/project", UsageFiles, true)
			}
		})
	}
}

// BenchmarkPatternMatching benchmarks pattern matching against paths
func BenchmarkPatternMatching(b *testing.B) {
	testCases := []struct {
		name  string
		spec  string
		paths []string
	}{
		{
			name: "LiteralMatch",
			spec: "src/file.tlua",
			paths: []string{
				"/project/src/file.tlua",
				"/project/src/other.tlua",
				"/project/lib/file.tlua",
			},
		},
		{
			name: "WildcardMatch",
			spec: "src/*.tlua",
			paths: []string{
				"/project/src/file.tlua",
				"/project/src/component.tlua",
				"/project/src/deep/file.tlua",
				"/project/lib/file.tlua",
			},
		},
		{
			name: "RecursiveMatch",
			spec: "**/file.tlua",
			paths: []string{
				"/project/file.tlua",
				"/project/src/file.tlua",
				"/project/src/deep/nested/file.tlua",
				"/project/src/other.tlua",
			},
		},
		{
			name: "ComplexMatch",
			spec: "src/**/components/*.tsx",
			paths: []string{
				"/project/src/components/Button.tsx",
				"/project/src/features/auth/components/Login.tsx",
				"/project/src/components/Button.tlua",
				"/project/lib/components/Button.tsx",
			},
		},
	}

	for _, tc := range testCases {
		pattern, ok := compileGlobPattern(tc.spec, "/project", UsageFiles, true)
		if !ok {
			continue
		}

		b.Run(tc.name, func(b *testing.B) {
			for b.Loop() {
				for _, path := range tc.paths {
					pattern.matches(path)
				}
			}
		})
	}
}

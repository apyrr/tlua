package vfs_test

import (
	"testing"
	"testing/fstest"

	"github.com/apyrr/tlua/internal/repo"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/osvfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func BenchmarkReadFile(b *testing.B) {
	type bench struct {
		name string
		fs   vfs.FS
		path string
	}

	osFS := osvfs.FS()

	const smallData = "hello, world"
	tmpdir := tspath.NormalizeSlashes(b.TempDir())
	osSmallDataPath := tspath.CombinePaths(tmpdir, "foo.tlua")
	err := osFS.WriteFile(osSmallDataPath, smallData)
	assert.NilError(b, err)

	tests := []bench{
		{"MapFS small", vfstest.FromMap(fstest.MapFS{
			"/foo.tlua": &fstest.MapFile{
				Data: []byte(smallData),
			},
		}, true), "/foo.tlua"},
		{"OS small", osFS, osSmallDataPath},
	}

	if repo.TypeScriptSubmoduleExists() {
		checkerPath := tspath.CombinePaths(tspath.NormalizeSlashes(repo.TypeScriptSubmodulePath()), "src", "compiler", "checker.ts")

		checkerContents, ok := osFS.ReadFile(checkerPath)
		assert.Assert(b, ok)

		tests = append(tests, bench{
			"MapFS checker.tlua",
			vfstest.FromMap(fstest.MapFS{
				"/checker.tlua": &fstest.MapFile{
					Data: []byte(checkerContents),
				},
			}, true),
			"/checker.tlua",
		})
		tests = append(tests, bench{"OS checker.tlua", osFS, checkerPath})
	}

	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				_, _ = tt.fs.ReadFile(tt.path)
			}
		})
	}
}

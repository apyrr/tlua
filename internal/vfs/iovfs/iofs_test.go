package iovfs_test

import (
	"slices"
	"testing"
	"testing/fstest"

	"github.com/apyrr/tlua/internal/testutil"
	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/iovfs"
	"gotest.tools/v3/assert"
)

func TestIOFS(t *testing.T) {
	t.Parallel()

	testfs := fstest.MapFS{
		"foo.tlua": &fstest.MapFile{
			Data: []byte("hello, world"),
		},
		"dir1/file1.tlua": &fstest.MapFile{
			Data: []byte("export local foo = 42;"),
		},
		"dir1/file2.tlua": &fstest.MapFile{
			Data: []byte("export local foo = 42;"),
		},
		"dir2/file1.tlua": &fstest.MapFile{
			Data: []byte("export local foo = 42;"),
		},
	}

	fs := iovfs.From(testfs, true)

	t.Run("ReadFile", func(t *testing.T) {
		t.Parallel()

		content, ok := fs.ReadFile("/foo.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/does/not/exist.tlua")
		assert.Assert(t, !ok)
		assert.Equal(t, content, "")
	})

	t.Run("ReadFileUnrooted", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() { fs.ReadFile("bar") }, `vfs: path "bar" is not absolute`)
	})

	t.Run("FileExists", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, fs.FileExists("/foo.tlua"))
		assert.Assert(t, !fs.FileExists("/bar"))
	})

	t.Run("DirectoryExists", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, fs.DirectoryExists("/"))
		assert.Assert(t, fs.DirectoryExists("/dir1"))
		assert.Assert(t, fs.DirectoryExists("/dir1/"))
		assert.Assert(t, fs.DirectoryExists("/dir1/./"))
		assert.Assert(t, !fs.DirectoryExists("/bar"))
	})

	t.Run("GetAccessibleEntries", func(t *testing.T) {
		t.Parallel()

		entries := fs.GetAccessibleEntries("/")
		assert.DeepEqual(t, entries.Directories, []string{"dir1", "dir2"})
		assert.DeepEqual(t, entries.Files, []string{"foo.tlua"})
	})

	t.Run("WalkDir", func(t *testing.T) {
		t.Parallel()

		var files []string
		err := fs.WalkDir("/", func(path string, d vfs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				files = append(files, path)
			}
			return nil
		})
		assert.NilError(t, err)

		slices.Sort(files)

		assert.DeepEqual(t, files, []string{"/dir1/file1.tlua", "/dir1/file2.tlua", "/dir2/file1.tlua", "/foo.tlua"})
	})

	t.Run("WalkDirSkip", func(t *testing.T) {
		t.Parallel()

		var files []string
		err := fs.WalkDir("/", func(path string, d vfs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				files = append(files, path)
			}

			if path == "/" {
				return nil
			}

			return vfs.SkipDir
		})
		assert.NilError(t, err)

		slices.Sort(files)

		assert.DeepEqual(t, files, []string{"/foo.tlua"})
	})

	t.Run("Realpath", func(t *testing.T) {
		t.Parallel()

		realpath := fs.Realpath("/foo.tlua")
		assert.Equal(t, realpath, "/foo.tlua")
	})

	t.Run("UseCaseSensitiveFileNames", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, fs.UseCaseSensitiveFileNames())
	})
}

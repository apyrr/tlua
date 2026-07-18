package vfstest

import (
	"encoding/binary"
	"io/fs"
	"math/rand/v2"
	"runtime"
	"slices"
	"sync"
	"testing"
	"testing/fstest"
	"unicode/utf16"

	"github.com/apyrr/tlua/internal/testutil"
	"github.com/apyrr/tlua/internal/vfs"
	"gotest.tools/v3/assert"
)

func TestInsensitive(t *testing.T) {
	t.Parallel()

	contents := []byte("bar")

	vfs := convertMapFS(fstest.MapFS{
		"foo/bar/baz": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"foo/bar2/baz2": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"foo/bar3/baz3": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
	}, false /*useCaseSensitiveFileNames*/, nil)

	sensitive, err := fs.ReadFile(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.DeepEqual(t, sensitive, contents)
	sensitiveInfo, err := fs.Stat(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.Equal(t, sensitiveInfo.Sys(), 1234)
	sensitiveRealPath, err := vfs.Realpath("foo/bar/baz")
	assert.NilError(t, err)
	assert.Equal(t, sensitiveRealPath, "foo/bar/baz")
	entries, err := fs.ReadDir(vfs, "foo")
	assert.NilError(t, err)
	assert.DeepEqual(t, dirEntriesToNames(entries), []string{"bar", "bar2", "bar3"})

	_, err = vfs.Realpath("does/not/exist")
	assert.ErrorContains(t, err, "file does not exist")
	_, err = fs.Stat(vfs, "does/not/exist")
	assert.ErrorContains(t, err, "file does not exist")

	assert.NilError(t, fstest.TestFS(vfs, "foo/bar/baz"))

	insensitive, err := fs.ReadFile(vfs, "Foo/Bar/Baz")
	assert.NilError(t, err)
	assert.DeepEqual(t, insensitive, contents)
	insensitiveInfo, err := fs.Stat(vfs, "Foo/Bar/Baz")
	assert.NilError(t, err)
	assert.Equal(t, insensitiveInfo.Sys(), 1234)
	insensitiveRealPath, err := vfs.Realpath("Foo/Bar/Baz")
	assert.NilError(t, err)
	assert.Equal(t, insensitiveRealPath, "foo/bar/baz")
	entries, err = fs.ReadDir(vfs, "Foo")
	assert.NilError(t, err)
	assert.DeepEqual(t, dirEntriesToNames(entries), []string{"bar", "bar2", "bar3"})

	_, err = vfs.Realpath("Does/Not/Exist")
	assert.ErrorContains(t, err, "file does not exist")
	_, err = fs.Stat(vfs, "Does/Not/Exist")
	assert.ErrorContains(t, err, "file does not exist")

	// TODO: TestFS doesn't understand case-insensitive file systems.
	// This same thing would happen with an os.Dir on Windows.
	// assert.NilError(t, fstest.TestFS(vfs, "Foo/Bar/Baz"))
}

func TestInsensitiveUpper(t *testing.T) {
	t.Parallel()

	contents := []byte("bar")

	vfs := convertMapFS(fstest.MapFS{
		"Foo/Bar/Baz": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"Foo/Bar2/Baz2": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"Foo/Bar3/Baz3": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
	}, false /*useCaseSensitiveFileNames*/, nil)

	sensitive, err := fs.ReadFile(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.DeepEqual(t, sensitive, contents)
	sensitiveInfo, err := fs.Stat(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.Equal(t, sensitiveInfo.Sys(), 1234)
	entries, err := fs.ReadDir(vfs, "foo")
	assert.NilError(t, err)
	assert.DeepEqual(t, dirEntriesToNames(entries), []string{"Bar", "Bar2", "Bar3"})

	// assert.NilError(t, fstest.TestFS(vfs, "foo/bar/baz"))

	insensitive, err := fs.ReadFile(vfs, "Foo/Bar/Baz")
	assert.NilError(t, err)
	assert.DeepEqual(t, insensitive, contents)
	insensitiveInfo, err := fs.Stat(vfs, "Foo/Bar/Baz")
	assert.NilError(t, err)
	assert.Equal(t, insensitiveInfo.Sys(), 1234)
	entries, err = fs.ReadDir(vfs, "Foo")
	assert.NilError(t, err)
	assert.DeepEqual(t, dirEntriesToNames(entries), []string{"Bar", "Bar2", "Bar3"})

	assert.NilError(t, fstest.TestFS(vfs, "Foo/Bar/Baz"))
}

func TestSensitive(t *testing.T) {
	t.Parallel()

	contents := []byte("bar")

	vfs := convertMapFS(fstest.MapFS{
		"foo/bar/baz": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"foo/bar2/baz2": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
		"foo/bar3/baz3": &fstest.MapFile{
			Data: contents,
			Sys:  1234,
		},
	}, true /*useCaseSensitiveFileNames*/, nil)

	sensitive, err := fs.ReadFile(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.DeepEqual(t, sensitive, contents)
	sensitiveInfo, err := fs.Stat(vfs, "foo/bar/baz")
	assert.NilError(t, err)
	assert.Equal(t, sensitiveInfo.Sys(), 1234)

	assert.NilError(t, fstest.TestFS(vfs, "foo/bar/baz"))

	_, err = fs.ReadFile(vfs, "Foo/Bar/Baz")
	assert.ErrorContains(t, err, "file does not exist")
}

func TestSensitiveDuplicatePath(t *testing.T) {
	t.Parallel()

	testfs := fstest.MapFS{
		"foo": &fstest.MapFile{
			Data: []byte("bar"),
		},
		"Foo": &fstest.MapFile{
			Data: []byte("baz"),
		},
	}

	testutil.AssertPanics(t, func() {
		convertMapFS(testfs, false /*useCaseSensitiveFileNames*/, nil)
	}, `duplicate path: "Foo" and "foo" have the same canonical path`)
}

func TestInsensitiveDuplicatePath(t *testing.T) {
	t.Parallel()

	testfs := fstest.MapFS{
		"foo": &fstest.MapFile{
			Data: []byte("bar"),
		},
		"Foo": &fstest.MapFile{
			Data: []byte("baz"),
		},
	}

	convertMapFS(testfs, true /*useCaseSensitiveFileNames*/, nil)
}

func dirEntriesToNames(entries []fs.DirEntry) []string {
	names := make([]string, len(entries))
	for i, entry := range entries {
		names[i] = entry.Name()
	}
	return names
}

func TestWritableFS(t *testing.T) {
	t.Parallel()

	fs := FromMap[any](nil, false)

	err := fs.WriteFile("/foo/bar/baz", "hello, world")
	assert.NilError(t, err)

	content, ok := fs.ReadFile("/foo/bar/baz")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")

	err = fs.WriteFile("/foo/bar/baz", "goodbye, world")
	assert.NilError(t, err)

	content, ok = fs.ReadFile("/foo/bar/baz")
	assert.Assert(t, ok)
	assert.Equal(t, content, "goodbye, world")

	err = fs.WriteFile("/foo/bar/baz/oops", "goodbye, world")
	assert.ErrorContains(t, err, `mkdir "foo/bar/baz": path exists but is not a directory`)
}

func TestWritableFSDelete(t *testing.T) {
	t.Parallel()
	fs := FromMap[any](nil, false)

	_ = fs.WriteFile("/foo/bar/file.tlua", "remove")
	assert.Assert(t, fs.FileExists("/foo/bar/file.tlua"))
	err := fs.Remove("/foo/bar/file.tlua")
	assert.NilError(t, err)
	assert.Assert(t, !fs.FileExists("/foo/bar/file.tlua"))

	_ = fs.WriteFile("/foo/bar/test/remove2.tlua", "remove2")
	assert.Assert(t, fs.DirectoryExists("/foo/bar/test"))
	err = fs.Remove("/foo/bar/test")
	assert.NilError(t, err)
	assert.Assert(t, !fs.FileExists("/foo/bar/test/remove2.tlua"))
	assert.Assert(t, !fs.DirectoryExists("/foo/bar/test"))

	// no errors when removing file/dir that does not exist
	err = fs.Remove("/foo/bar/test")
	assert.NilError(t, err)
	err = fs.Remove("/foo/bar/file.tlua")
	assert.NilError(t, err)

	_ = fs.WriteFile("/foo/barbar", "remove2")
	_ = fs.Remove("/foo/bar")
	assert.Assert(t, fs.FileExists("/foo/barbar"))
}

func TestStress(t *testing.T) {
	t.Parallel()

	fs := FromMap[any](nil, false)

	ops := []func(){
		func() { _ = fs.WriteFile("/foo/bar/baz.txt", "hello, world") },
		func() { fs.ReadFile("/foo/bar/baz.txt") },
		func() { fs.DirectoryExists("/foo/bar") },
		func() { fs.FileExists("/foo/bar") },
		func() { fs.FileExists("/foo/bar/baz.txt") },
		func() { fs.GetAccessibleEntries("/foo/bar") },
		func() { fs.Realpath("/foo/bar/baz.txt") },
		func() {
			_ = fs.WalkDir("/", func(path string, d vfs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				_, err = d.Info()
				return err
			})
		},
	}

	var wg sync.WaitGroup
	for range runtime.GOMAXPROCS(0) {
		wg.Go(func() {
			randomOps := slices.Clone(ops)
			rand.Shuffle(len(randomOps), func(i, j int) {
				randomOps[i], randomOps[j] = randomOps[j], randomOps[i]
			})

			for i := range 10000 {
				randomOps[i%len(randomOps)]()
			}
		})
	}
	wg.Wait()
}

func TestParentDirFile(t *testing.T) {
	t.Parallel()

	testfs := fstest.MapFS{
		"foo": &fstest.MapFile{
			Data: []byte("bar"),
		},
		"foo/oops": &fstest.MapFile{
			Data: []byte("baz"),
		},
	}

	testutil.AssertPanics(t, func() {
		convertMapFS(testfs, false /*useCaseSensitiveFileNames*/, nil)
	}, `failed to create intermediate directories for "foo/oops": mkdir "foo": path exists but is not a directory`)
}

func TestFromMap(t *testing.T) {
	t.Parallel()

	t.Run("POSIX", func(t *testing.T) {
		t.Parallel()

		fs := FromMap(map[string]any{
			"/string": "hello, world",
			"/bytes":  []byte("hello, world"),
			"/mapfile": &fstest.MapFile{
				Data: []byte("hello, world"),
			},
		}, false)

		content, ok := fs.ReadFile("/string")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/bytes")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/mapfile")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")
	})

	t.Run("Windows", func(t *testing.T) {
		t.Parallel()

		fs := FromMap(map[string]any{
			"c:/string": "hello, world",
			"d:/bytes":  []byte("hello, world"),
			"e:/mapfile": &fstest.MapFile{
				Data: []byte("hello, world"),
			},
		}, false)

		content, ok := fs.ReadFile("c:/string")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("d:/bytes")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("e:/mapfile")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")
	})

	t.Run("Mixed", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() {
			FromMap(map[string]any{
				"/string":  "hello, world",
				"c:/bytes": []byte("hello, world"),
			}, false)
		}, `mixed posix and windows paths`)
	})

	t.Run("NonRooted", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() {
			FromMap(map[string]any{
				"string": "hello, world",
			}, false)
		}, `non-rooted path "string"`)
	})

	t.Run("NonNormalized", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() {
			FromMap(map[string]any{
				"/string/": "hello, world",
			}, false)
		}, `non-normalized path "/string/"`)
	})

	t.Run("NonNormalized2", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() {
			FromMap(map[string]any{
				"/string/../foo": "hello, world",
			}, false)
		}, `non-normalized path "/string/../foo"`)
	})

	t.Run("InvalidFile", func(t *testing.T) {
		t.Parallel()

		testutil.AssertPanics(t, func() {
			FromMap(map[string]any{
				"/string": 1234,
			}, false)
		}, `invalid file type int`)
	})
}

func TestVFSTestMapFS(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]string{
		"/foo.tlua":        "hello, world",
		"/dir1/file1.tlua": "export local foo = 42;",
		"/dir1/file2.tlua": "export local foo = 42;",
		"/dir2/file1.tlua": "export local foo = 42;",
	}, false /*useCaseSensitiveFileNames*/)

	t.Run("ReadFile", func(t *testing.T) {
		t.Parallel()

		content, ok := fs.ReadFile("/foo.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/does/not/exist.tlua")
		assert.Assert(t, !ok)
		assert.Equal(t, content, "")
	})

	t.Run("Realpath", func(t *testing.T) {
		t.Parallel()

		realpath := fs.Realpath("/foo.tlua")
		assert.Equal(t, realpath, "/foo.tlua")

		realpath = fs.Realpath("/Foo.tlua")
		assert.Equal(t, realpath, "/foo.tlua")

		realpath = fs.Realpath("/does/not/exist.tlua")
		assert.Equal(t, realpath, "/does/not/exist.tlua")
	})

	t.Run("UseCaseSensitiveFileNames", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, !fs.UseCaseSensitiveFileNames())
	})
}

func TestVFSTestMapFSWindows(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]string{
		"c:/foo.tlua":        "hello, world",
		"c:/dir1/file1.tlua": "export local foo = 42;",
		"c:/dir1/file2.tlua": "export local foo = 42;",
		"c:/dir2/file1.tlua": "export local foo = 42;",
	}, false)

	t.Run("ReadFile", func(t *testing.T) {
		t.Parallel()

		content, ok := fs.ReadFile("c:/foo.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("c:/does/not/exist.tlua")
		assert.Assert(t, !ok)
		assert.Equal(t, content, "")
	})

	t.Run("Realpath", func(t *testing.T) {
		t.Parallel()

		realpath := fs.Realpath("c:/foo.tlua")
		assert.Equal(t, realpath, "c:/foo.tlua")

		realpath = fs.Realpath("c:/Foo.tlua")
		assert.Equal(t, realpath, "c:/foo.tlua")

		realpath = fs.Realpath("c:/does/not/exist.tlua")
		assert.Equal(t, realpath, "c:/does/not/exist.tlua")
	})
}

func TestBOM(t *testing.T) {
	t.Parallel()

	const expected = "hello, world"

	tests := []struct {
		name  string
		order binary.ByteOrder
		bom   [2]byte
	}{
		{"BigEndian", binary.BigEndian, [2]byte{0xFE, 0xFF}},
		{"LittleEndian", binary.LittleEndian, [2]byte{0xFF, 0xFE}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var codePoints []uint16

			for _, r := range expected {
				codePoints = utf16.AppendRune(codePoints, r)
			}

			buf := tt.bom[:]

			for _, r := range codePoints {
				var err error
				buf, err = binary.Append(buf, tt.order, r)
				assert.NilError(t, err)
			}

			fs := FromMap(map[string][]byte{
				"/foo.tlua": buf,
			}, true)

			content, ok := fs.ReadFile("/foo.tlua")
			assert.Assert(t, ok)
			assert.Equal(t, content, expected)
		})
	}

	t.Run("UTF8", func(t *testing.T) {
		t.Parallel()

		fs := FromMap(map[string][]byte{
			"/foo.tlua": []byte("\xEF\xBB\xBF" + expected),
		}, true)

		content, ok := fs.ReadFile("/foo.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, expected)
	})
}

func TestSymlink(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]any{
		"/foo.tlua":           "hello, world",
		"/symlink.tlua":       Symlink("/foo.tlua"),
		"/some/dir/file.tlua": "hello, world",
		"/some/dirlink":       Symlink("/some/dir"),
		"/a":                  Symlink("/b"),
		"/b":                  Symlink("/c"),
		"/c":                  Symlink("/d"),
		"/d/existing.tlua":    "this is existing.tlua",
	}, false)

	t.Run("ReadFile", func(t *testing.T) {
		t.Parallel()

		content, ok := fs.ReadFile("/symlink.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/some/dirlink/file.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "hello, world")

		content, ok = fs.ReadFile("/a/existing.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "this is existing.tlua")
	})

	t.Run("Realpath", func(t *testing.T) {
		t.Parallel()

		realpath := fs.Realpath("/symlink.tlua")
		assert.Equal(t, realpath, "/foo.tlua")

		realpath = fs.Realpath("/some/dirlink")
		assert.Equal(t, realpath, "/some/dir")

		realpath = fs.Realpath("/some/dirlink/file.tlua")
		assert.Equal(t, realpath, "/some/dir/file.tlua")
	})

	t.Run("FileExists", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, fs.FileExists("/symlink.tlua"))
		assert.Assert(t, fs.FileExists("/some/dirlink/file.tlua"))
		assert.Assert(t, fs.FileExists("/a/existing.tlua"))
	})

	t.Run("DirectoryExists", func(t *testing.T) {
		t.Parallel()

		assert.Assert(t, fs.DirectoryExists("/some/dirlink"))
		assert.Assert(t, fs.DirectoryExists("/d"))
		assert.Assert(t, fs.DirectoryExists("/c"))
		assert.Assert(t, fs.DirectoryExists("/b"))
		assert.Assert(t, fs.DirectoryExists("/a"))
	})
}

func TestWritableFSSymlink(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]any{
		"/some/dir/other.tlua": "NOTHING",
		"/other.tlua":          Symlink("/some/dir/other.tlua"),
		"/some/dirlink":        Symlink("/some/dir"),
		"/brokenlink":          Symlink("/does/not/exist"),
		"/a":                   Symlink("/b"),
		"/b":                   Symlink("/c"),
		"/c":                   Symlink("/d"),
		"/d/existing.tlua":     "hello, world",
	}, false)

	err := fs.WriteFile("/some/dirlink/file.tlua", "hello, world")
	assert.NilError(t, err)

	content, ok := fs.ReadFile("/some/dirlink/file.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")

	content, ok = fs.ReadFile("/some/dir/file.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")

	err = fs.WriteFile("/some/dirlink/file.tlua", "goodbye, world")
	assert.NilError(t, err)

	content, ok = fs.ReadFile("/some/dirlink/file.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "goodbye, world")

	err = fs.WriteFile("/other.tlua", "hello, world")
	assert.NilError(t, err)

	content, ok = fs.ReadFile("/other.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")

	content, ok = fs.ReadFile("/some/dir/other.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")

	err = fs.WriteFile("/some/dirlink", "hello, world")
	assert.Error(t, err, `write "some/dirlink": path exists but is not a regular file`)

	// Can't write inside a broken dir symlink
	err = fs.WriteFile("/brokenlink/file.tlua", "hello, world")
	assert.Error(t, err, `broken symlink "brokenlink" -> "does/not/exist"`)

	err = fs.WriteFile("/brokenlink/also/wrong/file.tlua", "hello, world")
	assert.Error(t, err, `broken symlink "brokenlink" -> "does/not/exist"`)

	// But we can write to a broken file symlink
	err = fs.WriteFile("/brokenlink", "hello, world")
	assert.NilError(t, err)
	content, ok = fs.ReadFile("/brokenlink")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")
	content, ok = fs.ReadFile("/does/not/exist")
	assert.Assert(t, ok)
	assert.Equal(t, content, "hello, world")
}

func TestWritableFSSymlinkChain(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]any{
		"/a":               Symlink("/b"),
		"/b":               Symlink("/c"),
		"/c":               Symlink("/d"),
		"/d/existing.tlua": "hello, world",
	}, false)

	err := fs.WriteFile("/a/foo/bar/new.tlua", "this is new.tlua")
	assert.NilError(t, err)
	content, ok := fs.ReadFile("/a/foo/bar/new.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "this is new.tlua")
	content, ok = fs.ReadFile("/b/foo/bar/new.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "this is new.tlua")
	content, ok = fs.ReadFile("/d/foo/bar/new.tlua")
	assert.Assert(t, ok)
	assert.Equal(t, content, "this is new.tlua")
}

func TestWritableFSSymlinkChainNotDir(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]any{
		"/a": Symlink("/b"),
		"/b": Symlink("/c"),
		"/c": Symlink("/d"),
		"/d": "hello, world",
	}, false)

	err := fs.WriteFile("/a/foo/bar/new.tlua", "this is new.tlua")
	assert.Error(t, err, `mkdir "d": path exists but is not a directory`)
}

func TestWritableFSSymlinkDelete(t *testing.T) {
	t.Parallel()

	fs := FromMap(map[string]any{
		"/some/dir/other.tlua": "NOTHING",
		"/other.tlua":          Symlink("/some/dir/other.tlua"),
		"/some/dirlink":        Symlink("/some/dir"),
		"/brokenlink":          Symlink("/does/not/exist"),
		"/a":                   Symlink("/b"),
		"/b":                   Symlink("/c"),
		"/c":                   Symlink("/d"),
		"/d/existing.tlua":     "hello, world",
	}, false)

	err := fs.Remove("/a")
	assert.NilError(t, err)
	assert.Assert(t, !fs.DirectoryExists("/a"))
	assert.Assert(t, fs.DirectoryExists("/b"))
	assert.Assert(t, fs.DirectoryExists("/c"))
	assert.Assert(t, fs.FileExists("/d/existing.tlua"))

	// symlinks should still exist even if underlying file/dir is deleted
	err = fs.Remove("/d")
	assert.NilError(t, err)
	assert.Assert(t, !fs.DirectoryExists("/b"))
	assert.Assert(t, !fs.DirectoryExists("/c"))
	assert.Assert(t, !fs.DirectoryExists("/d"))
	assert.Assert(t, !fs.FileExists("/d/again.tlua"))
	err = fs.WriteFile("/d/again.tlua", "d exists again")
	assert.NilError(t, err)
	assert.Assert(t, fs.DirectoryExists("/b"))
	assert.Assert(t, fs.DirectoryExists("/c"))
	content, _ := fs.ReadFile("/b/again.tlua")
	assert.Equal(t, content, "d exists again")

	assert.Assert(t, !fs.FileExists("/brokenlink"))
	assert.Assert(t, !fs.DirectoryExists("/brokenlink"))
	err = fs.Remove("/does/not/exist") // should do nothing
	assert.NilError(t, err)
	assert.Assert(t, !fs.FileExists("/brokenlink"))
	assert.Assert(t, !fs.DirectoryExists("/brokenlink"))
	err = fs.WriteFile("/does/not/exist", "hello, world")
	assert.NilError(t, err)
	assert.Assert(t, fs.FileExists("/brokenlink"))
}

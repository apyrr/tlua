package project

import (
	"slices"
	"sync"
	"testing"

	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project/dirty"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestSnapshotFSBuilder(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	t.Run("builds directory tree on file add", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
		}, false /* useCaseSensitiveFileNames */)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Read the file to add it to the diskFiles
		fh := builder.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil, "file should exist")
		assert.Equal(t, fh.Content(), "local foo = 1;")

		// Finalize and check directories
		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// Check that directory structure was built
		// /src should contain /src/foo.tlua
		srcDir, ok := snapshot.diskDirectories[tspath.Path("/src")]
		assert.Assert(t, ok, "/src directory should exist")
		_, hasFoo := srcDir[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, hasFoo, "/src should contain /src/foo.tlua")

		// / should contain /src
		rootDir, ok := snapshot.diskDirectories[tspath.Path("/")]
		assert.Assert(t, ok, "/ directory should exist")
		_, hasSrc := rootDir[tspath.Path("/src")]
		assert.Assert(t, hasSrc, "/ should contain /src")
	})

	t.Run("builds nested directory tree", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/nested/deep/file.tlua": "export local x = 1;",
		}, false /* useCaseSensitiveFileNames */)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Read the file to add it to the diskFiles
		fh := builder.GetFile("/src/nested/deep/file.tlua")
		assert.Assert(t, fh != nil, "file should exist")

		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// Check the complete directory tree
		_, hasFile := snapshot.diskDirectories[tspath.Path("/src/nested/deep")][tspath.Path("/src/nested/deep/file.tlua")]
		assert.Assert(t, hasFile)
		_, hasDeep := snapshot.diskDirectories[tspath.Path("/src/nested")][tspath.Path("/src/nested/deep")]
		assert.Assert(t, hasDeep)
		_, hasNested := snapshot.diskDirectories[tspath.Path("/src")][tspath.Path("/src/nested")]
		assert.Assert(t, hasNested)
		_, hasSrc := snapshot.diskDirectories[tspath.Path("/")][tspath.Path("/src")]
		assert.Assert(t, hasSrc)
	})

	t.Run("removes directory entries on file delete", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
		}, false /* useCaseSensitiveFileNames */)

		// Start with existing diskFiles and directories
		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "local foo = 1;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/foo.tlua"): "foo.tlua",
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			existingDiskFiles,
			existingDirs,
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Mark the file for deletion by loading and deleting
		if entry, ok := builder.diskFiles.Load(tspath.Path("/src/foo.tlua")); ok {
			entry.Delete()
		}

		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// File should be deleted
		_, hasFile := snapshot.diskFiles[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, !hasFile, "file should be deleted")

		// Directory tree should be cleaned up
		_, hasSrcDir := snapshot.diskDirectories[tspath.Path("/src")]
		assert.Assert(t, !hasSrcDir, "/src directory should be removed")

		_, hasRootDir := snapshot.diskDirectories[tspath.Path("/")]
		assert.Assert(t, !hasRootDir, "root directory should be removed")
	})

	t.Run("removes only empty directories on file delete", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
			"/src/bar.tlua": "local bar = 2;",
		}, false /* useCaseSensitiveFileNames */)

		// Start with existing diskFiles and directories
		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "local foo = 1;"),
			tspath.Path("/src/bar.tlua"): newDiskFile("/src/bar.tlua", "local bar = 2;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/foo.tlua"): "foo.tlua",
				tspath.Path("/src/bar.tlua"): "bar.tlua",
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			existingDiskFiles,
			existingDirs,
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Delete only foo.tlua
		if entry, ok := builder.diskFiles.Load(tspath.Path("/src/foo.tlua")); ok {
			entry.Delete()
		}

		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// foo.tlua should be deleted
		_, hasFile := snapshot.diskFiles[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, !hasFile, "foo.tlua should be deleted")

		// bar.tlua should still exist
		_, hasBar := snapshot.diskFiles[tspath.Path("/src/bar.tlua")]
		assert.Assert(t, hasBar, "bar.tlua should still exist")

		// /src directory should still exist with bar.tlua
		srcDir, hasSrcDir := snapshot.diskDirectories[tspath.Path("/src")]
		assert.Assert(t, hasSrcDir, "/src directory should still exist")
		_, hasFoo := srcDir[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, !hasFoo, "/src should not contain foo.tlua")
		_, hasBarInDir := srcDir[tspath.Path("/src/bar.tlua")]
		assert.Assert(t, hasBarInDir, "/src should contain bar.tlua")

		// root should still contain /src
		rootDir, hasRootDir := snapshot.diskDirectories[tspath.Path("/")]
		assert.Assert(t, hasRootDir, "root directory should still exist")
		_, hasSrc := rootDir[tspath.Path("/src")]
		assert.Assert(t, hasSrc, "root should contain /src")
	})

	t.Run("adds file to existing directory", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
			"/src/bar.tlua": "local bar = 2;",
		}, false /* useCaseSensitiveFileNames */)

		// Start with existing file and directories
		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "local foo = 1;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/foo.tlua"): "foo.tlua",
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			existingDiskFiles,
			existingDirs,
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Read bar.tlua to add it
		fh := builder.GetFile("/src/bar.tlua")
		assert.Assert(t, fh != nil, "bar.tlua should exist")

		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// /src should contain both files
		srcDir := snapshot.diskDirectories[tspath.Path("/src")]
		_, hasFoo := srcDir[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, hasFoo, "/src should contain foo.tlua")
		_, hasBar := srcDir[tspath.Path("/src/bar.tlua")]
		assert.Assert(t, hasBar, "/src should contain bar.tlua")
	})

	t.Run("no change when no files added or deleted", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
		}, false /* useCaseSensitiveFileNames */)

		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "local foo = 1;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/foo.tlua"): "foo.tlua",
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			existingDiskFiles,
			existingDirs,
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Don't add or delete any files
		snapshot, changed := builder.Finalize()
		assert.Assert(t, !changed, "should not have changed")

		// Directories should remain the same
		srcDir := snapshot.diskDirectories[tspath.Path("/src")]
		_, hasFoo := srcDir[tspath.Path("/src/foo.tlua")]
		assert.Assert(t, hasFoo)
	})

	t.Run("overlay files are returned over disk files", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "local foo = 1;",
		}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/foo.tlua"): {
				fileBase: fileBase{fileName: "/src/foo.tlua", content: "local foo = 999;"},
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			overlays,
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Should return overlay content
		fh := builder.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), "local foo = 999;")
	})

	t.Run("multiple files added and deleted in single cycle", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/a.tlua":        "local a = 1;",
			"/src/b.tlua":        "local b = 2;",
			"/lib/utils.tlua":    "export local util = 1;",
			"/lib/helpers.tlua":  "export local helper = 1;",
			"/other/single.tlua": "local single = 1;",
		}, false /* useCaseSensitiveFileNames */)

		// Start with some existing files
		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/a.tlua"):        newDiskFile("/src/a.tlua", "local a = 1;"),
			tspath.Path("/other/single.tlua"): newDiskFile("/other/single.tlua", "local single = 1;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"):   "src",
				tspath.Path("/other"): "other",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/a.tlua"): "a.tlua",
			},
			tspath.Path("/other"): {
				tspath.Path("/other/single.tlua"): "single.tlua",
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			make(map[tspath.Path]*Overlay), // overlays
			existingDiskFiles,
			existingDirs,
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Add new files
		fh := builder.GetFile("/src/b.tlua")
		assert.Assert(t, fh != nil)
		fh = builder.GetFile("/lib/utils.tlua")
		assert.Assert(t, fh != nil)
		fh = builder.GetFile("/lib/helpers.tlua")
		assert.Assert(t, fh != nil)

		// Delete existing files
		if entry, ok := builder.diskFiles.Load(tspath.Path("/src/a.tlua")); ok {
			entry.Delete()
		}
		if entry, ok := builder.diskFiles.Load(tspath.Path("/other/single.tlua")); ok {
			entry.Delete()
		}

		snapshot, changed := builder.Finalize()
		assert.Assert(t, changed, "should have changed")

		// Verify deleted files are gone
		_, hasA := snapshot.diskFiles[tspath.Path("/src/a.tlua")]
		assert.Assert(t, !hasA, "/src/a.tlua should be deleted")
		_, hasSingle := snapshot.diskFiles[tspath.Path("/other/single.tlua")]
		assert.Assert(t, !hasSingle, "/other/single.tlua should be deleted")

		// Verify added files exist
		_, hasB := snapshot.diskFiles[tspath.Path("/src/b.tlua")]
		assert.Assert(t, hasB, "/src/b.tlua should exist")
		_, hasUtils := snapshot.diskFiles[tspath.Path("/lib/utils.tlua")]
		assert.Assert(t, hasUtils, "/lib/utils.tlua should exist")
		_, hasHelpers := snapshot.diskFiles[tspath.Path("/lib/helpers.tlua")]
		assert.Assert(t, hasHelpers, "/lib/helpers.tlua should exist")

		// Verify /other directory is cleaned up (was only entry deleted)
		_, hasOther := snapshot.diskDirectories[tspath.Path("/other")]
		assert.Assert(t, !hasOther, "/other directory should be removed")

		// Verify /src still exists with b.tlua (a.tlua deleted, b.tlua added)
		srcDir, hasSrc := snapshot.diskDirectories[tspath.Path("/src")]
		assert.Assert(t, hasSrc, "/src directory should exist")
		_, hasAInDir := srcDir[tspath.Path("/src/a.tlua")]
		assert.Assert(t, !hasAInDir, "/src should not contain a.tlua")
		_, hasBInDir := srcDir[tspath.Path("/src/b.tlua")]
		assert.Assert(t, hasBInDir, "/src should contain b.tlua")

		// Verify /lib was created with both files
		libDir, hasLib := snapshot.diskDirectories[tspath.Path("/lib")]
		assert.Assert(t, hasLib, "/lib directory should exist")
		_, hasUtilsInDir := libDir[tspath.Path("/lib/utils.tlua")]
		assert.Assert(t, hasUtilsInDir, "/lib should contain utils.tlua")
		_, hasHelpersInDir := libDir[tspath.Path("/lib/helpers.tlua")]
		assert.Assert(t, hasHelpersInDir, "/lib should contain helpers.tlua")

		// Verify root contains /src and /lib but not /other
		rootDir := snapshot.diskDirectories[tspath.Path("/")]
		_, hasSrcInRoot := rootDir[tspath.Path("/src")]
		assert.Assert(t, hasSrcInRoot, "root should contain /src")
		_, hasLibInRoot := rootDir[tspath.Path("/lib")]
		assert.Assert(t, hasLibInRoot, "root should contain /lib")
		_, hasOtherInRoot := rootDir[tspath.Path("/other")]
		assert.Assert(t, !hasOtherInRoot, "root should not contain /other")
	})

	t.Run("overlay directories are computed from overlays", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/overlay.tlua"): {
				fileBase: fileBase{fileName: "/src/overlay.tlua", content: "local x = 1;"},
			},
			tspath.Path("/src/nested/deep.tlua"): {
				fileBase: fileBase{fileName: "/src/nested/deep.tlua", content: "local y = 2;"},
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			overlays,
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Check overlayDirectories was built correctly
		srcDir, ok := builder.overlayDirectories[tspath.Path("/src")]
		assert.Assert(t, ok, "/src overlay directory should exist")
		_, hasOverlay := srcDir[tspath.Path("/src/overlay.tlua")]
		assert.Assert(t, hasOverlay, "/src should contain overlay.tlua")
		_, hasNested := srcDir[tspath.Path("/src/nested")]
		assert.Assert(t, hasNested, "/src should contain nested/")

		nestedDir, ok := builder.overlayDirectories[tspath.Path("/src/nested")]
		assert.Assert(t, ok, "/src/nested overlay directory should exist")
		_, hasDeep := nestedDir[tspath.Path("/src/nested/deep.tlua")]
		assert.Assert(t, hasDeep, "/src/nested should contain deep.tlua")

		rootDir, ok := builder.overlayDirectories[tspath.Path("/")]
		assert.Assert(t, ok, "/ overlay directory should exist")
		_, hasSrc := rootDir[tspath.Path("/src")]
		assert.Assert(t, hasSrc, "/ should contain /src")
	})

	t.Run("GetAccessibleEntries combines disk and overlay", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/disk.tlua": "local disk = 1;",
		}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/overlay.tlua"): {
				fileBase: fileBase{fileName: "/src/overlay.tlua", content: "local overlay = 1;"},
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			overlays,
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		entries := builder.GetAccessibleEntries("/src")

		// Should contain both disk file and overlay file (both as basenames)
		assert.Assert(t, slices.Contains(entries.Files, "disk.tlua"), "should contain disk.tlua")
		assert.Assert(t, slices.Contains(entries.Files, "overlay.tlua"), "should contain overlay.tlua")
	})

	t.Run("GetAccessibleEntries is safe under concurrent calls", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/a.tlua": "",
			"/src/b.tlua": "",
			"/src/c.tlua": "",
			"/src/d.tlua": "",
			"/src/e.tlua": "",
		}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/overlay.tlua"): {
				fileBase: fileBase{fileName: "/src/overlay.tlua", content: ""},
			},
		}

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay), // prevOverlays
			overlays,
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		_ = builder.fs.GetAccessibleEntries("/src")

		start := make(chan struct{})
		var wg sync.WaitGroup
		for range 50 {
			wg.Go(func() {
				<-start
				_ = builder.GetAccessibleEntries("/src")
			})
		}
		close(start)
		wg.Wait()

		// Sanity: the overlay is still reported after all the concurrent calls.
		entries := builder.GetAccessibleEntries("/src")
		assert.Assert(t, slices.Contains(entries.Files, "overlay.tlua"), "should contain overlay.tlua")
	})
}

func TestSnapshotFS(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	t.Run("GetFile returns overlay file", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "disk content",
		}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/foo.tlua"): {
				fileBase: fileBase{fileName: "/src/foo.tlua", content: "overlay content"},
			},
		}

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           overlays,
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		fh := snapshot.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), "overlay content")
	})

	t.Run("GetFile returns disk file when not in overlay", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "disk content",
		}, false /* useCaseSensitiveFileNames */)

		diskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "disk content"),
		}

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          diskFiles,
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		fh := snapshot.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), "disk content")
	})

	t.Run("GetFile reads from fs when not cached", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "fs content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		fh := snapshot.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), "fs content")
	})

	t.Run("GetFile returns nil for non-existent file", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		fh := snapshot.GetFile("/src/nonexistent.tlua")
		assert.Assert(t, fh == nil, "should return nil for non-existent file")
	})

	t.Run("isOpenFile returns true for overlays", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/foo.tlua"): {
				fileBase: fileBase{fileName: "/src/foo.tlua", content: "overlay content"},
			},
		}

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           overlays,
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		assert.Assert(t, snapshot.isOpenFile("/src/foo.tlua"), "overlay file should be open")
		assert.Assert(t, !snapshot.isOpenFile("/src/bar.tlua"), "non-overlay file should not be open")
	})

	t.Run("GetFileByPath uses provided path", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "disk content",
		}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/foo.tlua"): {
				fileBase: fileBase{fileName: "/src/foo.tlua", content: "overlay content"},
			},
		}

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           overlays,
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		// GetFileByPath should use the provided path directly
		fh := snapshot.GetFileByPath("/src/foo.tlua", tspath.Path("/src/foo.tlua"))
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), "overlay content")
	})

	t.Run("GetAccessibleEntries combines disk and overlay directories", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{}, false /* useCaseSensitiveFileNames */)

		overlays := map[tspath.Path]*Overlay{
			tspath.Path("/src/overlay.tlua"): {
				fileBase: fileBase{fileName: "/src/overlay.tlua", content: "overlay content"},
			},
		}
		overlayDirectories := map[tspath.Path]map[tspath.Path]string{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/overlay.tlua"): "overlay.tlua",
			},
		}
		diskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/disk.tlua"): newDiskFile("/src/disk.tlua", "disk content"),
		}
		diskDirectories := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"): {
				tspath.Path("/src"): "src",
			},
			tspath.Path("/src"): {
				tspath.Path("/src/disk.tlua"): "disk.tlua",
			},
		}

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           overlays,
			overlayDirectories: overlayDirectories,
			diskFiles:          diskFiles,
			diskDirectories:    diskDirectories,
		}

		entries := snapshot.GetAccessibleEntries("/src")

		// Should contain both disk file and overlay file (both as basenames)
		assert.Assert(t, slices.Contains(entries.Files, "disk.tlua"), "should contain disk.tlua")
		assert.Assert(t, slices.Contains(entries.Files, "overlay.tlua"), "should contain overlay.tlua")
	})
}

func TestSourceFS(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	t.Run("tracks files when tracking enabled", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		sourceFS := newSourceFS(true /* tracking */, snapshot, toPath)

		// File should not be seen yet
		assert.Assert(t, !sourceFS.SeenFile(tspath.Path("/src/foo.tlua")))

		// Read the file
		fh := sourceFS.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)

		// Now it should be seen
		assert.Assert(t, sourceFS.SeenFile(tspath.Path("/src/foo.tlua")))
	})

	t.Run("does not track files when tracking disabled", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		sourceFS := newSourceFS(false /* tracking */, snapshot, toPath)

		// Read the file
		fh := sourceFS.GetFile("/src/foo.tlua")
		assert.Assert(t, fh != nil)

		// Should not be seen since tracking is disabled
		assert.Assert(t, !sourceFS.SeenFile(tspath.Path("/src/foo.tlua")))
	})

	t.Run("DisableTracking stops tracking", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "content",
			"/src/bar.tlua": "content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		sourceFS := newSourceFS(true /* tracking */, snapshot, toPath)

		// Read foo while tracking
		sourceFS.GetFile("/src/foo.tlua")
		assert.Assert(t, sourceFS.SeenFile(tspath.Path("/src/foo.tlua")))

		// Disable tracking
		sourceFS.DisableTracking()

		// Read bar after tracking disabled
		sourceFS.GetFile("/src/bar.tlua")
		assert.Assert(t, !sourceFS.SeenFile(tspath.Path("/src/bar.tlua")))
	})

	t.Run("FileExists returns true for files in source", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		sourceFS := newSourceFS(false /* tracking */, snapshot, toPath)

		assert.Assert(t, sourceFS.FileExists("/src/foo.tlua"))
		assert.Assert(t, !sourceFS.FileExists("/src/nonexistent.tlua"))
	})

	t.Run("ReadFile returns content for files in source", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]string{
			"/src/foo.tlua": "file content",
		}, false /* useCaseSensitiveFileNames */)

		snapshot := &SnapshotFS{
			toPath:             toPath,
			fs:                 testFS,
			overlays:           make(map[tspath.Path]*Overlay),
			overlayDirectories: make(map[tspath.Path]map[tspath.Path]string),
			diskFiles:          make(map[tspath.Path]*diskFile),
			diskDirectories:    make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
		}

		sourceFS := newSourceFS(false /* tracking */, snapshot, toPath)

		content, ok := sourceFS.ReadFile("/src/foo.tlua")
		assert.Assert(t, ok)
		assert.Equal(t, content, "file content")

		_, ok = sourceFS.ReadFile("/src/nonexistent.tlua")
		assert.Assert(t, !ok)
	})
}

func TestAutoImportBuilderFS(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	// This test demonstrates that autoImportBuilderFS stores files in untrackedFiles keyed
	// by the path derived from the filename. When module resolution reads a file via its
	// symlink path, the file is cached at the symlink path key. If the file is subsequently
	// requested by its realpath (which is the typical case after module resolution resolves
	// the symlink), the cache is bypassed. If the file was deleted from disk between those
	// two operations, the realpath read fails and returns nil. This is the mechanism behind
	// the nil source file crash in aliasResolver.GetSourceFile.
	t.Run("symlink cache mismatch: file cached at symlink path, missed at realpath after deletion", func(t *testing.T) {
		t.Parallel()

		// Create a VFS with a real file and a symlinked directory pointing to it.
		// /real/pkg/index.d.tlua is the real file.
		// /project/node_modules/pkg is a symlink to /real/pkg.
		testFS := vfstest.FromMap(map[string]any{
			"/real/pkg/index.d.tlua":    "export declare local x: number;",
			"/project/node_modules/pkg": vfstest.Symlink("/real/pkg"),
		}, true /* useCaseSensitiveFileNames */)

		// Verify symlink works as expected
		symlinkPath := "/project/node_modules/pkg/index.d.tlua"
		realpathPath := testFS.Realpath(symlinkPath)
		assert.Equal(t, realpathPath, "/real/pkg/index.d.tlua", "Realpath should resolve the symlink to the real path")

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil, // nodeModulesRealpathAliases
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		autoImportFS := &autoImportBuilderFS{
			snapshotFSBuilder: builder,
		}

		// Step 1: Read the file via its symlink path (simulating what module resolution does
		// during FileExists). This caches the file in untrackedFiles at the symlink path key.
		fh := autoImportFS.GetFile(symlinkPath)
		assert.Assert(t, fh != nil, "File should be readable via symlink path")
		assert.Equal(t, fh.Content(), "export declare local x: number;")

		// Step 2: Simulate a file deletion from disk (e.g., npm install running concurrently).
		// This deletes both the real file and effectively breaks the symlink.
		err := testFS.Remove("/real/pkg/index.d.tlua")
		assert.NilError(t, err)

		// Step 3: Request the file by its realpath (simulating what GetSourceFile does after
		// the checker resolves the module). This bypasses the symlink-path cache entry in
		// untrackedFiles because the realpath has a different key.
		fh2 := autoImportFS.GetFile(realpathPath)
		// The file was cached at the symlink path, but the realpath lookup misses the cache
		// and goes to disk where the file is now deleted. This returns nil.
		assert.Assert(t, fh2 == nil, "File should be nil when accessed by realpath after deletion from disk")
	})
}

func TestRealpathAliasLifecycle(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	t.Run("alias recorded when reading symlinked node_modules file", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":               vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json":              `{"name": "mylib", "main": "index.lua"}`,
			"/packages/mylib/index.d.tlua":              `export declare local x: number;`,
			"/project/node_modules/nolink/package.json": `{"name": "nolink"}`,
		}, false)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Read a file through the symlink — should record an alias.
		fh := builder.GetFile("/project/node_modules/mylib/package.json")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), `{"name": "mylib", "main": "index.lua"}`)

		// Read a non-symlinked node_modules file — should NOT record an alias.
		fh2 := builder.GetFile("/project/node_modules/nolink/package.json")
		assert.Assert(t, fh2 != nil)

		snapshot, _ := builder.Finalize()

		// Alias exists for the symlinked file.
		aliases, ok := snapshot.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "alias should exist for realpath of symlinked file")
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))

		// No alias for the non-symlinked file.
		_, ok = snapshot.nodeModulesRealpathAliases[tspath.Path("/project/node_modules/nolink/package.json")]
		assert.Assert(t, !ok, "no alias should exist for non-symlinked file")
	})

	t.Run("no alias recorded for files outside node_modules", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/link":         vfstest.Symlink("/elsewhere"),
			"/elsewhere/index.tlua": `export local x = 1;`,
		}, false)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		fh := builder.GetFile("/project/link/index.tlua")
		assert.Assert(t, fh != nil)

		snapshot, _ := builder.Finalize()
		assert.Equal(t, len(snapshot.nodeModulesRealpathAliases), 0, "no aliases for non-node_modules symlinks")
	})

	t.Run("aliases carried over across snapshots", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		// Build first snapshot.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder1.GetFile("/project/node_modules/mylib/package.json")
		snapshot1, _ := builder1.Finalize()

		// Build second snapshot from the first, without reading the file again.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		snapshot2, _ := builder2.Finalize()

		// Alias should still be present.
		aliases, ok := snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "alias should survive across snapshots")
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))
	})

	t.Run("alias pruned when symlinked file is deleted", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
			"/packages/mylib/index.d.tlua": `export declare local x: number;`,
		}, false)

		// Build first snapshot — read both files.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder1.GetFile("/project/node_modules/mylib/package.json")
		builder1.GetFile("/project/node_modules/mylib/index.d.tlua")
		snapshot1, _ := builder1.Finalize()

		// Both should be aliased under the same realpath directory but separate files.
		_, ok := snapshot1.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok)
		_, ok = snapshot1.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/index.d.tlua")]
		assert.Assert(t, ok)

		// Build second snapshot — delete one file via markDirtyFiles.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Simulate deletion of index.d.tlua from the disk file cache.
		var entry *dirty.SyncMapEntry[tspath.Path, *diskFile]
		if entry, ok = builder2.diskFiles.Load(tspath.Path("/project/node_modules/mylib/index.d.tlua")); ok {
			entry.Delete()
		}

		snapshot2, _ := builder2.Finalize()

		// package.json alias should remain.
		aliases, ok := snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "package.json alias should survive")
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))

		// index.d.tlua alias should be fully pruned (empty set → removed from map).
		_, ok = snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/index.d.tlua")]
		assert.Assert(t, !ok, "index.d.tlua alias should be pruned after deletion")
	})

	t.Run("multiple symlinks to same realpath", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/project/node_modules/alias":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		// Read via both symlinks.
		fh1 := builder.GetFile("/project/node_modules/mylib/package.json")
		assert.Assert(t, fh1 != nil)
		fh2 := builder.GetFile("/project/node_modules/alias/package.json")
		assert.Assert(t, fh2 != nil)

		snapshot, _ := builder.Finalize()

		aliases, ok := snapshot.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "alias should exist")
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/alias/package.json")))
	})

	t.Run("multiple symlinks pruned individually", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/project/node_modules/alias":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		// Build first snapshot – read via both symlinks.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder1.GetFile("/project/node_modules/mylib/package.json")
		builder1.GetFile("/project/node_modules/alias/package.json")
		snapshot1, _ := builder1.Finalize()

		// Build second snapshot – delete ONE of the symlink disk entries.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		if entry, ok := builder2.diskFiles.Load(tspath.Path("/project/node_modules/alias/package.json")); ok {
			entry.Delete()
		}
		snapshot2, _ := builder2.Finalize()

		// The realpath alias set should still exist, but only contain the surviving symlink.
		aliases, ok := snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "alias set should still exist")
		assert.Assert(t, aliases.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")), "surviving symlink should remain")
		assert.Assert(t, !aliases.paths.Has(tspath.Path("/project/node_modules/alias/package.json")), "deleted symlink should be pruned")
	})

	t.Run("expandRealpathAliases expands change events", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder.GetFile("/project/node_modules/mylib/package.json")
		snapshot, _ := builder.Finalize()

		// Simulate a watch event on the REALPATH.
		change := FileChangeSummary{}
		change.Changed.Add("file:///packages/mylib/package.json")

		expanded := snapshot.expandRealpathAliases(change)

		// Should now also contain the symlink path.
		assert.Assert(t, expanded.Changed.Has("file:///packages/mylib/package.json"), "original event should remain")
		assert.Assert(t, expanded.Changed.Has("file:///project/node_modules/mylib/package.json"), "symlink event should be added")
	})

	t.Run("expandRealpathAliases expands delete events", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		builder := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder.GetFile("/project/node_modules/mylib/package.json")
		snapshot, _ := builder.Finalize()

		// Simulate a delete watch event on the REALPATH.
		change := FileChangeSummary{}
		change.Deleted.Add("file:///packages/mylib/package.json")

		expanded := snapshot.expandRealpathAliases(change)

		assert.Assert(t, expanded.Deleted.Has("file:///project/node_modules/mylib/package.json"), "symlink deletion should be added")
	})

	t.Run("expandRealpathAliases is a no-op with no aliases", func(t *testing.T) {
		t.Parallel()
		snapshot := &SnapshotFS{
			toPath:                     toPath,
			nodeModulesRealpathAliases: nil,
		}

		change := FileChangeSummary{}
		change.Changed.Add("file:///some/file.tlua")

		expanded := snapshot.expandRealpathAliases(change)
		assert.Equal(t, expanded.Changed.Len(), 1)
		assert.Assert(t, expanded.Changed.Has("file:///some/file.tlua"))
	})

	t.Run("markDirtyFiles invalidates symlinked file via realpath event", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib", "main": "index.lua"}`,
		}, false)

		// Build first snapshot — read the symlinked file.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		fh := builder1.GetFile("/project/node_modules/mylib/package.json")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), `{"name": "mylib", "main": "index.lua"}`)
		snapshot1, _ := builder1.Finalize()

		// Modify the real file on disk.
		err := testFS.WriteFile("/packages/mylib/package.json", `{"name": "mylib"}`)
		assert.NilError(t, err)

		// Build second snapshot — simulate realpath change event, expanded via aliases.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		change := FileChangeSummary{}
		change.Changed.Add("file:///packages/mylib/package.json")

		// Expand the realpath event to include the symlink path.
		change = snapshot1.expandRealpathAliases(change)
		// Now mark dirty — should find the file under the symlink key.
		builder2.markDirtyFiles(change)

		// Trigger reload by reading the file (simulates program construction).
		fh = builder2.GetFile("/project/node_modules/mylib/package.json")
		assert.Assert(t, fh != nil)
		assert.Equal(t, fh.Content(), `{"name": "mylib"}`, "builder should serve updated content after dirty marking")

		snapshot2, _ := builder2.Finalize()

		// The file should have been reloaded with new content.
		file, ok := snapshot2.diskFiles[tspath.Path("/project/node_modules/mylib/package.json")]
		assert.Assert(t, ok, "file should still be in diskFiles")
		assert.Equal(t, file.Content(), `{"name": "mylib"}`, "content should be updated")
	})

	t.Run("alias clone isolation between snapshots", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/project/node_modules/other":  vfstest.Symlink("/packages/other"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
			"/packages/other/package.json": `{"name": "other"}`,
		}, false)

		// Build first snapshot — read only mylib.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder1.GetFile("/project/node_modules/mylib/package.json")
		snapshot1, _ := builder1.Finalize()

		// Build second snapshot — also read other.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder2.GetFile("/project/node_modules/other/package.json")
		snapshot2, _ := builder2.Finalize()

		// snapshot1 should only have mylib alias.
		_, ok := snapshot1.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "snapshot1 should have mylib alias")
		_, ok = snapshot1.nodeModulesRealpathAliases[tspath.Path("/packages/other/package.json")]
		assert.Assert(t, !ok, "snapshot1 should NOT have other alias — it was added in a later snapshot")

		// snapshot2 should have both.
		_, ok = snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok, "snapshot2 should have mylib alias")
		_, ok = snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/other/package.json")]
		assert.Assert(t, ok, "snapshot2 should have other alias")
	})

	t.Run("adding symlink to inherited realpath key does not mutate previous snapshot", func(t *testing.T) {
		t.Parallel()
		testFS := vfstest.FromMap(map[string]any{
			"/project/node_modules/mylib":  vfstest.Symlink("/packages/mylib"),
			"/project/node_modules/alias":  vfstest.Symlink("/packages/mylib"),
			"/packages/mylib/package.json": `{"name": "mylib"}`,
		}, false)

		// Snapshot 1: read via one symlink only.
		builder1 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder1.GetFile("/project/node_modules/mylib/package.json")
		snapshot1, _ := builder1.Finalize()

		// Verify snapshot1 has exactly one alias for the realpath.
		aliases1, ok := snapshot1.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok)
		assert.Equal(t, aliases1.paths.Len(), 1)
		assert.Assert(t, aliases1.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))

		// Snapshot 2: read via the SECOND symlink, which maps to the same realpath.
		// This exercises the case where LoadOrStore finds the key in the base map
		// and must clone-on-write rather than mutating the shared set.
		builder2 := newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			snapshot1.diskFiles,
			snapshot1.diskDirectories,
			snapshot1.nodeModulesRealpathAliases,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
		builder2.GetFile("/project/node_modules/alias/package.json")
		snapshot2, _ := builder2.Finalize()

		// Snapshot 2 should have both symlinks.
		aliases2, ok := snapshot2.nodeModulesRealpathAliases[tspath.Path("/packages/mylib/package.json")]
		assert.Assert(t, ok)
		assert.Equal(t, aliases2.paths.Len(), 2)
		assert.Assert(t, aliases2.paths.Has(tspath.Path("/project/node_modules/mylib/package.json")))
		assert.Assert(t, aliases2.paths.Has(tspath.Path("/project/node_modules/alias/package.json")))

		// Snapshot 1 must NOT have been mutated — it should still have only one alias.
		assert.Equal(t, aliases1.paths.Len(), 1, "snapshot1 alias set must not be mutated by snapshot2")
		assert.Assert(t, !aliases1.paths.Has(tspath.Path("/project/node_modules/alias/package.json")),
			"snapshot1 must not contain alias added in snapshot2")
	})
}

func TestExpandAndFilterWatchEvents(t *testing.T) {
	t.Parallel()

	toPath := func(fileName string) tspath.Path {
		return tspath.Path(fileName)
	}

	newBuilder := func(testFS vfs.FS) *snapshotFSBuilder {
		return newSnapshotFSBuilder(
			testFS,
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*diskFile),
			make(map[tspath.Path]dirty.CloneableMap[tspath.Path, string]),
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)
	}

	t.Run("preserves node_modules directory deletion even when untracked", func(t *testing.T) {
		t.Parallel()
		// node_modules package files are read transiently and never tracked in
		// diskDirectories, so a node_modules directory deletion can neither be
		// expanded nor matched by extension. It must still be preserved.
		builder := newBuilder(vfstest.FromMap(map[string]string{
			"/project/index.tlua": "export local x = 1;",
		}, false))

		change := FileChangeSummary{}
		change.Deleted.Add("file:///project/node_modules")

		expanded := builder.expandAndFilterWatchEvents(change)
		assert.Assert(t, expanded.Deleted.Has("file:///project/node_modules"),
			"bare node_modules directory deletion should be preserved")
	})

	t.Run("preserves deletion of a package directory inside node_modules", func(t *testing.T) {
		t.Parallel()
		builder := newBuilder(vfstest.FromMap(map[string]string{
			"/project/index.tlua": "export local x = 1;",
		}, false))

		change := FileChangeSummary{}
		change.Deleted.Add("file:///project/node_modules/@scope/pkg")

		expanded := builder.expandAndFilterWatchEvents(change)
		assert.Assert(t, expanded.Deleted.Has("file:///project/node_modules/@scope/pkg"),
			"package directory deletion inside node_modules should be preserved")
	})

	t.Run("drops irrelevant untracked deletion outside node_modules", func(t *testing.T) {
		t.Parallel()
		builder := newBuilder(vfstest.FromMap(map[string]string{
			"/project/index.tlua": "export local x = 1;",
		}, false))

		change := FileChangeSummary{}
		change.Deleted.Add("file:///project/build")

		expanded := builder.expandAndFilterWatchEvents(change)
		assert.Equal(t, expanded.Deleted.Len(), 0,
			"untracked non-node_modules directory deletion should be dropped")
	})

	t.Run("expands tracked directory deletion into file deletions", func(t *testing.T) {
		t.Parallel()
		existingDiskFiles := map[tspath.Path]*diskFile{
			tspath.Path("/src/foo.tlua"): newDiskFile("/src/foo.tlua", "local foo = 1;"),
		}
		existingDirs := map[tspath.Path]dirty.CloneableMap[tspath.Path, string]{
			tspath.Path("/"):    {tspath.Path("/src"): "src"},
			tspath.Path("/src"): {tspath.Path("/src/foo.tlua"): "foo.tlua"},
		}
		builder := newSnapshotFSBuilder(
			vfstest.FromMap(map[string]string{"/src/foo.tlua": "local foo = 1;"}, false),
			make(map[tspath.Path]*Overlay),
			make(map[tspath.Path]*Overlay),
			existingDiskFiles,
			existingDirs,
			nil,
			lsproto.PositionEncodingKindUTF16,
			toPath,
		)

		change := FileChangeSummary{}
		change.Deleted.Add("file:///src")

		expanded := builder.expandAndFilterWatchEvents(change)
		assert.Assert(t, expanded.Deleted.Has("file:///src/foo.tlua"),
			"tracked directory deletion should expand to contained file deletions")
		assert.Assert(t, !expanded.Deleted.Has("file:///src"),
			"the directory URI itself should be replaced by its files")
	})
}

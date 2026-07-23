package project

import (
	"context"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestRefCountingCaches(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	setup := func(files map[string]any) *Session {
		fs := bundled.WrapFS(vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/))
		session := NewSession(&SessionInit{
			BackgroundCtx: context.Background(),
			Options: &SessionOptions{
				// The Lua resolver's search root for inferred projects is the
				// current directory; point it at the fixture project.
				CurrentDirectory:   "/user/username/projects/myproject",
				DefaultLibraryPath: bundled.LibPath(),
				TypingsLocation:    "/home/src/Library/Caches/typescript",
				PositionEncoding:   lsproto.PositionEncodingKindUTF8,
				WatchEnabled:       false,
				LoggingEnabled:     false,
			},
			FS: fs,
		})
		return session
	}

	t.Run("parseCache", func(t *testing.T) {
		t.Parallel()

		files := map[string]any{
			"/user/username/projects/myproject/src/main.tlua":  "local x = 1;",
			"/user/username/projects/myproject/src/utils.tlua": "function util() end",
		}

		t.Run("reuse unchanged file", func(t *testing.T) {
			t.Parallel()

			session := setup(files)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua", 1, files["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua", 1, files["/user/username/projects/myproject/src/utils.tlua"].(string), lsproto.LanguageKindTypeScript)
			snapshot := session.Snapshot()
			program := snapshot.ProjectCollection.InferredProject().Program
			main := program.GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			utils := program.GetSourceFile("/user/username/projects/myproject/src/utils.tlua")
			mainEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			utilsEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(utils.ParseOptions(), utils.Hash, utils.ScriptKind))
			assert.Equal(t, mainEntry.refCount, 1)
			assert.Equal(t, utilsEntry.refCount, 1)

			session.DidChangeFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: &lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{Line: 0, Character: 0},
							End:   lsproto.Position{Line: 0, Character: 12},
						},
						Text: "local x = 2;",
					},
				},
			})
			ls, err := session.GetLanguageService(context.Background(), "file:///user/username/projects/myproject/src/main.tlua")
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()
			newMain := ls.GetProgram().GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			newMainEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(newMain.ParseOptions(), newMain.Hash, newMain.ScriptKind))
			assert.Assert(t, newMain != main)
			assert.Assert(t, newMainEntry != mainEntry)
			assert.Equal(t, ls.GetProgram().GetSourceFile("/user/username/projects/myproject/src/utils.tlua"), utils)
			// Old snapshot is deref'd immediately when replaced by UpdateSnapshot,
			// so old mainEntry is already disposed and utils refCount is already 1.
			assert.Equal(t, mainEntry.refCount, 0)
			assert.Equal(t, newMainEntry.refCount, 1)
			assert.Equal(t, utilsEntry.refCount, 1)
		})

		t.Run("release file on close", func(t *testing.T) {
			t.Parallel()

			session := setup(files)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua", 1, files["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua", 1, files["/user/username/projects/myproject/src/utils.tlua"].(string), lsproto.LanguageKindTypeScript)
			snapshot := session.Snapshot()
			program := snapshot.ProjectCollection.InferredProject().Program
			main := program.GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			utils := program.GetSourceFile("/user/username/projects/myproject/src/utils.tlua")
			mainEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			utilsEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(utils.ParseOptions(), utils.Hash, utils.ScriptKind))
			assert.Equal(t, mainEntry.refCount, 1)
			assert.Equal(t, utilsEntry.refCount, 1)

			session.DidCloseFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua")
			_, err := session.GetLanguageService(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua")
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()
			assert.Equal(t, utilsEntry.refCount, 1)
			assert.Equal(t, mainEntry.refCount, 0)
			mainEntry, ok := session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			assert.Equal(t, ok, false)
		})

		t.Run("unchanged program does not over-ref", func(t *testing.T) {
			t.Parallel()

			// When a program is reused across snapshots without changes, we should
			// not accumulate extra refs. The ref count should stay at 1 per source file
			// until the program is finally disposed.
			session := setup(files)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua", 1, files["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua", 1, files["/user/username/projects/myproject/src/utils.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Get first snapshot and capture the program/entries
			snapshot1 := session.Snapshot()
			program1 := snapshot1.ProjectCollection.InferredProject().Program
			main := program1.GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			mainEntry, _ := session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			assert.Equal(t, mainEntry.refCount, 1, "initial refCount should be 1")

			// Change utils.tlua to trigger a new snapshot, but main.tlua stays the same
			// so main's source file should be reused.
			session.DidChangeFile(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: &lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{Line: 0, Character: 0},
							End:   lsproto.Position{Line: 0, Character: 18},
						},
						Text: "function util2() end",
					},
				},
			})

			// Get second snapshot - main.tlua should be reused (program is new but shares source files)
			ls, err := session.GetLanguageService(context.Background(), "file:///user/username/projects/myproject/src/main.tlua")
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()
			program2 := ls.GetProgram()
			main2 := program2.GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			assert.Equal(t, main, main2, "main.tlua source file should be reused")

			// main.tlua refCount should be 1: the old snapshot was immediately deref'd
			// when replaced, so only the new snapshot holds a ref.
			mainEntry, _ = session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			assert.Equal(t, mainEntry.refCount, 1, "refCount should be 1 (only new snapshot)")

			// Close files to trigger cleanup
			session.DidCloseFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua")
			session.DidCloseFile(context.Background(), "file:///user/username/projects/myproject/src/utils.tlua")
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()

			// Entry should now be gone (refCount 0, deleted)
			mainEntry, ok := session.parseCache.entries.Load(NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind))
			if ok {
				t.Logf("Entry still exists with refCount=%d", mainEntry.refCount)
			}
			assert.Assert(t, !ok, "entry should be deleted after program is disposed")
		})

		t.Run("fallback rebuild does not double-ref changed file", func(t *testing.T) {
			t.Parallel()

			testFiles := map[string]any{
				"/user/username/projects/myproject/src/main.tlua":  "local x = 1;",
				"/user/username/projects/myproject/src/utils.tlua": "local util = 1; return { util = util };",
			}
			session := setup(testFiles)
			mainURI := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
			session.DidOpenFile(context.Background(), mainURI, 1, testFiles["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)

			_, err := session.GetLanguageService(context.Background(), mainURI)
			assert.NilError(t, err)

			session.DidChangeFile(context.Background(), mainURI, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{
						Text: "local utils = require(\"src.utils\");\nlocal x = utils.util;",
					},
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), mainURI)
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()

			project := session.Snapshot().ProjectCollection.InferredProject()
			assert.Assert(t, project != nil)
			assert.Equal(t, project.ProgramUpdateKind, ProgramUpdateKindNewFiles)

			main := lsAfter.GetProgram().GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			mainKey := NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind)
			mainEntry, ok := session.parseCache.entries.Load(mainKey)
			assert.Assert(t, ok)
			assert.Equal(t, mainEntry.refCount, 1)

			session.DidCloseFile(context.Background(), mainURI)
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()

			_, ok = session.parseCache.entries.Load(mainKey)
			assert.Assert(t, !ok)
		})

		t.Run("case-only duplicate loads are released on dispose", func(t *testing.T) {
			t.Parallel()

			testFiles := map[string]any{
				"/user/username/projects/myproject/src/main.tlua":  "local a = require(\"src.utils\");\nlocal b = require(\"src.UTILS\");\nlocal x = a.util + b.util;",
				"/user/username/projects/myproject/src/utils.tlua": "local util = 1; return { util = util };",
			}
			session := setup(testFiles)
			mainURI := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
			session.DidOpenFile(context.Background(), mainURI, 1, testFiles["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), mainURI)
			assert.NilError(t, err)

			var projectEntries int
			session.parseCache.entries.Range(func(key ParseCacheKey, _ *refCountCacheEntry[*ast.SourceFile]) bool {
				if strings.HasPrefix(key.FileName, "/user/username/projects/myproject/src/") {
					projectEntries++
				}
				return true
			})
			assert.Equal(t, projectEntries, 3)

			utils := ls.GetProgram().GetSourceFile("/user/username/projects/myproject/src/utils.tlua")
			assert.Assert(t, utils != nil)

			session.DidCloseFile(context.Background(), mainURI)
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()

			projectEntries = 0
			session.parseCache.entries.Range(func(key ParseCacheKey, _ *refCountCacheEntry[*ast.SourceFile]) bool {
				if strings.HasPrefix(key.FileName, "/user/username/projects/myproject/src/") {
					projectEntries++
				}
				return true
			})
			assert.Equal(t, projectEntries, 0)
		})

		t.Run("case-only duplicate imported from multiple files is refcounted once", func(t *testing.T) {
			t.Parallel()

			// A file reached through a case-only-different file name from more than one
			// import site is parsed and acquired in the parse cache exactly once (same-casing
			// loads dedupe), but it must also be recorded as a duplicate exactly once.
			// Recording it once per import site would release it from the parse cache more
			// times than it was acquired, deleting the live entry out from under a program
			// that still references it and panicking the next time it is ref'd during a clone.
			testFiles := map[string]any{
				// entry.tlua requires the canonical casing first, then pulls in a.tlua and b.tlua,
				// which both require the same file through an upper-cased name.
				"/user/username/projects/myproject/src/entry.tlua":   "local dep = require(\"src.sub.dep\");\nlocal a = require(\"src.a\");\nlocal b = require(\"src.b\");\nlocal e = dep.dep;\nreturn { e = e };",
				"/user/username/projects/myproject/src/a.tlua":       "local dep = require(\"src.sub.DEP\");\nlocal a = dep.dep;\nreturn { a = a };",
				"/user/username/projects/myproject/src/b.tlua":       "local dep = require(\"src.sub.DEP\");\nlocal b = dep.dep;\nreturn { b = b };",
				"/user/username/projects/myproject/src/sub/dep.tlua": "local dep = 1; return { dep = dep };",
				"/user/username/projects/myproject/src/c.tlua":       "local c = 1; return { c = c };",
			}
			session := setup(testFiles)
			entryURI := lsproto.DocumentUri("file:///user/username/projects/myproject/src/entry.tlua")
			session.DidOpenFile(context.Background(), entryURI, 1, testFiles["/user/username/projects/myproject/src/entry.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), entryURI)
			assert.NilError(t, err)

			// The upper-cased name is recorded as a duplicate, and it should appear exactly once.
			program := ls.GetProgram()
			var dupKeys []ParseCacheKey
			for _, dup := range program.DuplicateSourceFiles() {
				if strings.HasSuffix(dup.ParseOptions.FileName, "/sub/DEP.tlua") {
					dupKeys = append(dupKeys, NewParseCacheKey(dup.ParseOptions, dup.Hash, dup.ScriptKind))
				}
			}
			assert.Equal(t, len(dupKeys), 1, "case-only duplicate should be recorded exactly once")
			dupEntry, ok := session.parseCache.entries.Load(dupKeys[0])
			assert.Assert(t, ok, "duplicate entry should exist in the parse cache")
			assert.Equal(t, dupEntry.refCount, 1)

			// Force a full program rebuild (adding a require changes the file's module
			// structure). The old snapshot is disposed, releasing each of its source and
			// duplicate files exactly once. If the duplicate were recorded twice, the
			// shared cache entry would be released to zero and deleted here even though
			// the new program still references it.
			session.DidChangeFile(context.Background(), entryURI, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{
						Text: "local dep = require(\"src.sub.dep\");\nlocal a = require(\"src.a\");\nlocal b = require(\"src.b\");\nlocal c = require(\"src.c\");\nlocal e = dep.dep;\nreturn { e = e };",
					},
				},
			})
			lsAfterRebuild, err := session.GetLanguageService(context.Background(), entryURI)
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()

			// Every parse-cache key referenced by the live program must still exist.
			rebuiltProgram := lsAfterRebuild.GetProgram()
			assertKeyAlive := func(key ParseCacheKey) {
				_, alive := session.parseCache.entries.Load(key)
				assert.Assert(t, alive, "live program references a deleted parse-cache entry: %s", key.FileName)
			}
			for _, file := range rebuiltProgram.SourceFiles() {
				assertKeyAlive(NewParseCacheKey(file.ParseOptions(), file.Hash, file.ScriptKind))
			}
			for _, dup := range rebuiltProgram.DuplicateSourceFiles() {
				assertKeyAlive(NewParseCacheKey(dup.ParseOptions, dup.Hash, dup.ScriptKind))
			}

			// An incremental (clone) update re-references the duplicate files; this must
			// not panic with "cache entry not found".
			session.DidChangeFile(context.Background(), entryURI, 3, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{
						Text: "local dep = require(\"src.sub.dep\");\nlocal a = require(\"src.a\");\nlocal b = require(\"src.b\");\nlocal c = require(\"src.c\");\nlocal e = dep.dep + 0;\nreturn { e = e };",
					},
				},
			})
			_, err = session.GetLanguageService(context.Background(), entryURI)
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()

			// Closing the project releases everything cleanly.
			// (The configured project is not disposed until another file in another project is opened,
			// so we open an untitled file to trigger that.)
			session.DidCloseFile(context.Background(), entryURI)
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()

			projectEntries := 0
			session.parseCache.entries.Range(func(key ParseCacheKey, _ *refCountCacheEntry[*ast.SourceFile]) bool {
				if strings.HasPrefix(key.FileName, "/user/username/projects/myproject/src/") {
					projectEntries++
				}
				return true
			})
			assert.Equal(t, projectEntries, 0)
		})
	})

	t.Run("extendedConfigCache", func(t *testing.T) {
		files := map[string]any{
			"/user/username/projects/myproject/tluaconfig.json": `{
				"extends": "./tluaconfig.base.json"
			}`,
			"/user/username/projects/myproject/tluaconfig.base.json": `{
				"compilerOptions": {}
			}`,
			"/user/username/projects/myproject/src/main.tlua": "local x = 1;",
		}

		t.Run("release extended configs with project close", func(t *testing.T) {
			t.Parallel()

			session := setup(files)
			session.DidOpenFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua", 1, files["/user/username/projects/myproject/src/main.tlua"].(string), lsproto.LanguageKindTypeScript)
			snapshot := session.Snapshot()
			config := snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tluaconfig.json")
			assert.Equal(t, config.ExtendedSourceFiles()[0], "/user/username/projects/myproject/tluaconfig.base.json")
			extendedConfigEntry, _ := session.extendedConfigCache.entries.Load("/user/username/projects/myproject/tluaconfig.base.json")
			assert.Equal(t, len(extendedConfigEntry.owners), 1)

			session.DidCloseFile(context.Background(), "file:///user/username/projects/myproject/src/main.tlua")
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()
			_, ok := session.extendedConfigCache.entries.Load("/user/username/projects/myproject/tluaconfig.base.json")
			assert.Equal(t, ok, false)
		})

		t.Run("release cache entries for unretained clone", func(t *testing.T) {
			t.Parallel()

			session := setup(files)
			uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
			baseSnapshot := session.Snapshot()
			extendedConfigPath := tspath.Path("/user/username/projects/myproject/tluaconfig.base.json")
			clone := baseSnapshot.Clone(context.Background(), SnapshotChange{
				reason: UpdateReasonRequestedLanguageServiceProjectNotLoaded,
				ResourceRequest: ResourceRequest{
					Documents: []lsproto.DocumentUri{uri},
				},
			}, baseSnapshot.fs.overlays, session)

			project := clone.GetDefaultProject(uri)
			assert.Assert(t, project != nil)
			assert.Equal(t, project.ProgramLastUpdate, clone.id)

			main := project.Program.GetSourceFile("/user/username/projects/myproject/src/main.tlua")
			mainKey := NewParseCacheKey(main.ParseOptions(), main.Hash, main.ScriptKind)
			mainEntry, ok := session.parseCache.entries.Load(mainKey)
			assert.Assert(t, ok)
			assert.Equal(t, mainEntry.refCount, 1)

			extendedConfigEntry, ok := session.extendedConfigCache.entries.Load(extendedConfigPath)
			assert.Assert(t, ok)
			assert.Equal(t, len(extendedConfigEntry.owners), 1)

			clone.Deref(session)

			_, ok = session.parseCache.entries.Load(mainKey)
			assert.Assert(t, !ok)

			_, ok = session.extendedConfigCache.entries.Load(extendedConfigPath)
			assert.Assert(t, !ok)
		})
	})
}

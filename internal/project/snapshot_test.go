package project

import (
	"context"
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestSnapshot(t *testing.T) {
	t.Parallel()
	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	setup := func(files map[string]any) *Session {
		fs := bundled.WrapFS(vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/))
		session := NewSession(&SessionInit{
			BackgroundCtx: context.Background(),
			Options: &SessionOptions{
				CurrentDirectory:   "/",
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

	t.Run("compilerHost gets frozen with snapshot's FS only once", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": "{}",
			"/home/projects/TS/p1/index.tlua":    "console.log('Hello, world!');",
		}
		session := setup(files)
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
		snapshotBefore := session.Snapshot()

		session.DidChangeFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
			{
				Partial: &lsproto.TextDocumentContentChangePartial{
					Text: "\n",
					Range: lsproto.Range{
						Start: lsproto.Position{Line: 0, Character: 24},
						End:   lsproto.Position{Line: 0, Character: 24},
					},
				},
			},
		})
		_, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/index.tlua")
		assert.NilError(t, err)
		snapshotAfter := session.Snapshot()

		// Configured project was updated by a clone
		assert.Equal(t, snapshotAfter.ProjectCollection.ConfiguredProject(tspath.Path("/home/projects/ts/p1/tsconfig.json")).ProgramUpdateKind, ProgramUpdateKindCloned)
		// Inferred project wasn't updated last snapshot change, so its program update kind is still NewFiles
		assert.Equal(t, snapshotBefore.ProjectCollection.InferredProject(), snapshotAfter.ProjectCollection.InferredProject())
		assert.Equal(t, snapshotAfter.ProjectCollection.InferredProject().ProgramUpdateKind, ProgramUpdateKindNewFiles)
		// host for inferred project should not change
		assert.Equal(t, snapshotAfter.ProjectCollection.InferredProject().host.sourceFS.source, snapshotBefore.fs)
	})

	t.Run("cached disk files are cleaned up", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": "{}",
			"/home/projects/TS/p1/index.tlua":    "import { a } from './a'; console.log(a);",
			"/home/projects/TS/p1/a.tlua":        "export local a = 1;",
			"/home/projects/TS/p2/tsconfig.json": "{}",
			"/home/projects/TS/p2/index.tlua":    "import { b } from './b'; console.log(b);",
			"/home/projects/TS/p2/b.tlua":        "export local b = 2;",
		}
		session := setup(files)
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p2/index.tlua", 1, files["/home/projects/TS/p2/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		snapshotBefore := session.Snapshot()

		// a.tlua and b.tlua are cached
		assert.Check(t, snapshotBefore.fs.diskFiles["/home/projects/ts/p1/a.tlua"] != nil)
		assert.Check(t, snapshotBefore.fs.diskFiles["/home/projects/ts/p2/b.tlua"] != nil)

		// Close p1's only open file
		session.DidCloseFile(context.Background(), "file:///home/projects/TS/p1/index.tlua")
		// Next open file is unrelated to p1, triggers p1 closing and file cache cleanup
		session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
		snapshotAfter := session.Snapshot()

		// a.tlua is cleaned up, b.tlua is still cached
		assert.Check(t, snapshotAfter.fs.diskFiles["/home/projects/ts/p1/a.tlua"] == nil)
		assert.Check(t, snapshotAfter.fs.diskFiles["/home/projects/ts/p2/b.tlua"] != nil)
	})

	t.Run("GetFile returns nil for non-existent files", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": "{}",
			"/home/projects/TS/p1/index.tlua":    "console.log('Hello, world!');",
		}
		session := setup(files)
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()

		handle := snapshot.GetFile("/home/projects/TS/p1/nonexistent.tlua")
		assert.Check(t, handle == nil, "GetFile should return nil for non-existent file")

		// Test that ReadFile returns false for non-existent file
		_, ok := snapshot.ReadFile("/home/projects/TS/p1/nonexistent.tlua")
		assert.Check(t, !ok, "ReadFile should return false for non-existent file")
	})

	t.Run("program change loads node_modules dependency and auto-imports includes it", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/otherproject/tsconfig.json": `{
				"compilerOptions": {
					"module": "commonjs"
				}
			}`,
			"/home/projects/otherproject/index.tlua": ``,
			"/home/projects/node_modules/foo/package.json": `{
				"types": "index.d.tlua",
				"typesVersions": {
					"*": {
						"bar/*": ["dist/*"],
						"exact-match": ["dist/index.d.tlua"],
						"foo/*": ["dist/*"],
						"*": ["dist/*"]
					}
				}
			}`,
			"/home/projects/node_modules/foo/nope.d.tlua":                     `export local nope = 0;`,
			"/home/projects/node_modules/foo/dist/index.d.tlua":               `export local index = 0;`,
			"/home/projects/node_modules/foo/dist/blah.d.tlua":                `export local blah = 0;`,
			"/home/projects/node_modules/foo/dist/foo/onlyInFooFolder.d.tlua": `export local foo = 0;`,
			"/home/projects/node_modules/foo/dist/subfolder/one.d.tlua":       `export local one = 0;`,
		}
		session := setup(files)
		t.Cleanup(session.Close)
		ctx := context.Background()
		otherIndexURI := lsproto.DocumentUri("file:///home/projects/otherproject/index.tlua")

		// Open the file
		session.DidOpenFile(ctx, otherIndexURI, 1, files["/home/projects/otherproject/index.tlua"].(string), lsproto.LanguageKindTypeScript)

		// Insert import statement:
		// This will trigger both a program rebuild which will include the node_modules files,
		// and an auto-import collection which should find the exports from those files.
		session.DidChangeFile(ctx, otherIndexURI, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
			{
				Partial: &lsproto.TextDocumentContentChangePartial{
					Text: `import {} from "foo/foo/subfolder/one";`,
					Range: lsproto.Range{
						Start: lsproto.Position{Line: 0, Character: 0},
						End:   lsproto.Position{Line: 0, Character: 0},
					},
				},
			},
		})

		// Now trigger snapshot clone with both program update and auto-imports registry building.
		_, err := session.GetCurrentLanguageServiceWithAutoImports(ctx, otherIndexURI)
		assert.NilError(t, err)
	})

	t.Run("fallback rebuild with recomputed parse options is safe for later clone", func(t *testing.T) {
		t.Parallel()

		testFiles := map[string]any{
			"/project/node_modules/pkg/index.tlua": `export local pkg = 0;`,
			"/project/src/other.tlua":              `export local other = 1;`,
		}
		session := setup(testFiles)
		pkgURI := lsproto.DocumentUri("file:///project/node_modules/pkg/index.tlua")
		otherURI := lsproto.DocumentUri("file:///project/src/other.tlua")

		session.DidOpenFile(context.Background(), pkgURI, 1, testFiles["/project/node_modules/pkg/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		session.DidOpenFile(context.Background(), otherURI, 1, testFiles["/project/src/other.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), pkgURI)
		assert.NilError(t, err)

		err = session.fs.fs.WriteFile("/project/node_modules/pkg/package.json", `{ "type": "module" }`)
		assert.NilError(t, err)
		session.DidChangeFile(context.Background(), pkgURI, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
			{
				WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{
					Text: `import "./missing"; export local pkg = 1;`,
				},
			},
		})
		_, err = session.GetLanguageService(context.Background(), pkgURI)
		assert.NilError(t, err)

		session.DidChangeFile(context.Background(), otherURI, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
			{
				WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{
					Text: `export local other = 2;`,
				},
			},
		})
		_, err = session.GetLanguageService(context.Background(), otherURI)
		assert.NilError(t, err)
	})

	t.Run("auto-import snapshot is adopted when session snapshot is unchanged", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": "{}",
			"/home/projects/TS/p1/index.tlua":    "local value = foo;",
			"/home/projects/TS/p1/foo.tlua":      "export local foo = 1;",
		}
		session := setup(files)
		t.Cleanup(session.Close)
		ctx := context.Background()
		uri := lsproto.DocumentUri("file:///home/projects/TS/p1/index.tlua")

		session.DidOpenFile(ctx, uri, 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(ctx, uri)
		assert.NilError(t, err)

		baseSnapshot := session.Snapshot()
		preparedSnapshot := session.GetSnapshotWithAutoImports(ctx, baseSnapshot, uri)
		defer preparedSnapshot.Deref(session)

		session.WaitForBackgroundTasks()
		assert.Equal(t, session.Snapshot(), preparedSnapshot)
	})

	t.Run("no-op watch change does not rebuild program", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": "{}",
			"/home/projects/TS/p1/index.tlua":    "import { a } from './a'; console.log(a);",
			"/home/projects/TS/p1/a.tlua":        "export local a = 1;",
		}
		session := setup(files)
		t.Cleanup(session.Close)
		ctx := context.Background()
		uri := lsproto.DocumentUri("file:///home/projects/TS/p1/index.tlua")
		configPath := tspath.Path("/home/projects/ts/p1/tsconfig.json")

		session.DidOpenFile(ctx, uri, 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(ctx, uri)
		assert.NilError(t, err)

		programBefore := session.Snapshot().ProjectCollection.ConfiguredProject(configPath).Program

		// Send a watch change event for a project file whose content on disk is unchanged.
		// This should not invalidate the program or trigger a recheck.
		session.pendingFileChangesMu.Lock()
		session.pendingFileChanges = append(session.pendingFileChanges, FileChange{
			Kind: FileChangeKindWatchChange,
			URI:  "file:///home/projects/TS/p1/a.tlua",
		})
		session.pendingFileChangesMu.Unlock()
		_, err = session.GetLanguageService(ctx, uri)
		assert.NilError(t, err)

		programAfter := session.Snapshot().ProjectCollection.ConfiguredProject(configPath).Program
		assert.Equal(t, programBefore, programAfter, "no-op watch change should not rebuild the program")

		// A watch change that reflects an actual content change on disk must still
		// rebuild the program.
		err = session.fs.fs.WriteFile("/home/projects/TS/p1/a.tlua", "export local a = 2;")
		assert.NilError(t, err)
		session.pendingFileChangesMu.Lock()
		session.pendingFileChanges = append(session.pendingFileChanges, FileChange{
			Kind: FileChangeKindWatchChange,
			URI:  "file:///home/projects/TS/p1/a.tlua",
		})
		session.pendingFileChangesMu.Unlock()
		_, err = session.GetLanguageService(ctx, uri)
		assert.NilError(t, err)

		programChanged := session.Snapshot().ProjectCollection.ConfiguredProject(configPath).Program
		assert.Assert(t, programBefore != programChanged, "real watch change should rebuild the program")
	})
}

func BenchmarkSnapshotCloneRefCost(b *testing.B) {
	if !bundled.Embedded {
		b.Skip("bundled files are not embedded")
	}

	for _, largeProjectSize := range []int{100, 1000, 10_000} {
		b.Run(fmt.Sprintf("largeProject_%d_files", largeProjectSize), func(b *testing.B) {
			files := map[string]any{
				// Small project: 100 files
				"/small/tsconfig.json": `{"compilerOptions": {"strict": true}}`,
				// Large project: variable number of files
				"/large/tsconfig.json": `{"compilerOptions": {"strict": true}}`,
			}

			// Generate small project files
			for i := range 100 {
				files[fmt.Sprintf("/small/file%d.tlua", i)] = fmt.Sprintf("export local small%d = %d;", i, i)
			}

			// Generate large project files
			for i := range largeProjectSize {
				files[fmt.Sprintf("/large/file%d.tlua", i)] = fmt.Sprintf("export local large%d = %d;", i, i)
			}

			fs := bundled.WrapFS(vfstest.FromMap(files, false /*useCaseSensitiveFileNames*/))
			session := NewSession(&SessionInit{
				BackgroundCtx: context.Background(),
				Options: &SessionOptions{
					CurrentDirectory:   "/",
					DefaultLibraryPath: bundled.LibPath(),
					TypingsLocation:    "/home/src/Library/Caches/typescript",
					PositionEncoding:   lsproto.PositionEncodingKindUTF8,
					WatchEnabled:       false,
					LoggingEnabled:     false,
				},
				FS: fs,
			})

			// Open one file from each project to initialize them
			session.DidOpenFile(context.Background(), "file:///small/file0.tlua", 1, files["/small/file0.tlua"].(string), lsproto.LanguageKindTypeScript)
			_, err := session.GetLanguageService(context.Background(), "file:///small/file0.tlua")
			if err != nil {
				b.Fatal(err)
			}

			session.DidOpenFile(context.Background(), "file:///large/file0.tlua", 1, files["/large/file0.tlua"].(string), lsproto.LanguageKindTypeScript)
			_, err = session.GetLanguageService(context.Background(), "file:///large/file0.tlua")
			if err != nil {
				b.Fatal(err)
			}

			session.WaitForBackgroundTasks()

			// Now benchmark: change the small project's tsconfig, then request a language service.
			// This forces a snapshot clone, which refs every file in both projects.
			strict := true
			b.ResetTimer()
			for b.Loop() {
				strict = !strict
				var tsconfigContent string
				if strict {
					tsconfigContent = `{"compilerOptions": {"strict": true}}`
				} else {
					tsconfigContent = `{"compilerOptions": {"strict": false}}`
				}
				err := session.fs.fs.WriteFile("/small/tsconfig.json", tsconfigContent)
				if err != nil {
					b.Fatal(err)
				}
				session.pendingFileChangesMu.Lock()
				session.pendingFileChanges = append(session.pendingFileChanges, FileChange{
					Kind: FileChangeKindWatchChange,
					URI:  "file:///small/tsconfig.json",
				})
				session.pendingFileChangesMu.Unlock()
				_, err = session.GetLanguageService(context.Background(), "file:///small/file0.tlua")
				if err != nil {
					b.Fatal(err)
				}
				session.WaitForBackgroundTasks()
			}
		})
	}
}

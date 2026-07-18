package project_test

import (
	"context"
	"io/fs"
	"maps"
	"slices"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestSession(t *testing.T) {
	t.Parallel()
	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	defaultFiles := map[string]any{
		"/home/projects/TS/p1/tsconfig.json": `{
			"compilerOptions": {
				"noLib": true,
				"module": "nodenext",
				"strict": true
			},
			"include": ["src"]
		}`,
		"/home/projects/TS/p1/src/index.tlua": `local x = require("src.x");`,
		"/home/projects/TS/p1/src/x.tlua":     `local x = 1; return { x = x };`,
		"/home/projects/TS/p1/config.tlua":    `local x = 1, y = 2;`,
	}

	t.Run("DidOpenFile", func(t *testing.T) {
		t.Parallel()
		t.Run("create configured project", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)
			snapshot := session.Snapshot()
			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, defaultFiles["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			snapshot = session.Snapshot()
			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)

			configuredProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/home/projects/ts/p1/tsconfig.json"))
			assert.Assert(t, configuredProject != nil)

			// Get language service to access the program
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Assert(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") != nil)
			assert.Equal(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua").Text(), "local x = 1; return { x = x };")
		})

		t.Run("create inferred project", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/config.tlua", 1, defaultFiles["/home/projects/TS/p1/config.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Find tsconfig, load, notice config.tlua is not included, create inferred project
			snapshot := session.Snapshot()
			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 2)

			// Should have both configured project (for tsconfig.json) and inferred project
			configuredProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/home/projects/ts/p1/tsconfig.json"))
			inferredProject := snapshot.ProjectCollection.InferredProject()
			assert.Assert(t, configuredProject != nil)
			assert.Assert(t, inferredProject != nil)
		})

		t.Run("inferred project for in-memory files", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/config.tlua", 1, defaultFiles["/home/projects/TS/p1/config.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "x", lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "untitled:Untitled-2", 1, "y", lsproto.LanguageKindTypeScript)

			snapshot := session.Snapshot()

			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
			assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)
		})

		t.Run("inferred project JS file", func(t *testing.T) {
			t.Parallel()
			jsFiles := map[string]any{
				"/home/projects/TS/p1/index.lua": `x = 1;`,
			}
			session, _ := projecttestutil.Setup(jsFiles)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.lua", 1, jsFiles["/home/projects/TS/p1/index.lua"].(string), lsproto.LanguageKindJavaScript)

			snapshot := session.Snapshot()
			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/index.lua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Assert(t, program.GetSourceFile("/home/projects/TS/p1/index.lua") != nil)
		})
	})

	t.Run("watchChange and didOpen in same batch rebuilds program", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/projects/TS/p1/tsconfig.json": `{
				"compilerOptions": {
					"noLib": true,
					"strict": true
				}
			}`,
			"/home/projects/TS/p1/src/a.tlua": "local a = 1;\n",
			"/home/projects/TS/p1/src/b.tlua": "local b = 1;\n",
		}
		session, utils := projecttestutil.Setup(files)
		oldContent := files["/home/projects/TS/p1/src/a.tlua"].(string)

		// Open b.tlua to create the project; a.tlua is included via tsconfig.
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/b.tlua", 1, files["/home/projects/TS/p1/src/b.tlua"].(string), lsproto.LanguageKindTypeScript)

		// Verify a.tlua is in the program with the original content.
		ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/b.tlua")
		assert.NilError(t, err)
		assert.Equal(t, ls.GetProgram().GetSourceFile("/home/projects/TS/p1/src/a.tlua").Text(), oldContent)

		// Modify a.tlua on disk (simulate a build tool or git checkout).
		newContent := "local a = 2;\nlocal extra = true;\n"
		err = utils.FS().WriteFile("/home/projects/TS/p1/src/a.tlua", newContent)
		assert.NilError(t, err)

		// Queue a watch event for the disk change (not flushed yet).
		session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
			{Type: lsproto.FileChangeTypeChanged, Uri: "file:///home/projects/TS/p1/src/a.tlua"},
		})

		// Open a.tlua in the editor—flushes both watch event and didOpen together.
		// Before the fix, processChanges would discard the watch event,
		// leaving the project with a stale SourceFile and a mismatched line map.
		session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/a.tlua", 1, newContent, lsproto.LanguageKindTypeScript)

		// The program's SourceFile must reflect the overlay (new) content.
		ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/a.tlua")
		assert.NilError(t, err)
		assert.Equal(t, ls.GetProgram().GetSourceFile("/home/projects/TS/p1/src/a.tlua").Text(), newContent)
	})

	t.Run("DidChangeFile", func(t *testing.T) {
		t.Parallel()
		t.Run("update file and program", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, defaultFiles["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()

			session.DidChangeFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: new(lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{
								Line:      0,
								Character: 10,
							},
							End: lsproto.Position{
								Line:      0,
								Character: 11,
							},
						},
						Text: "2",
					}),
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			programAfter := lsAfter.GetProgram()

			// Program should change due to the file content change
			assert.Check(t, programAfter != programBefore)
			assert.Equal(t, programAfter.GetSourceFile("/home/projects/TS/p1/src/x.tlua").Text(), "local x = 2; return { x = x };")
		})

		t.Run("update untitled file", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)

			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "local x = 1;", lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "untitled:Untitled-1")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()
			untitledFileName := lsproto.DocumentUri("untitled:Untitled-1").FileName()
			assert.Equal(t, programBefore.GetSourceFile(untitledFileName).Text(), "local x = 1;")

			session.DidChangeFile(context.Background(), "untitled:Untitled-1", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: new(lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{
								Line:      0,
								Character: 10,
							},
							End: lsproto.Position{
								Line:      0,
								Character: 11,
							},
						},
						Text: "2",
					}),
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "untitled:Untitled-1")
			assert.NilError(t, err)
			programAfter := lsAfter.GetProgram()

			assert.Check(t, programAfter != programBefore)
			assert.Equal(t, programAfter.GetSourceFile(untitledFileName).Text(), "local x = 2;")
		})

		t.Run("unchanged source files are reused", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, defaultFiles["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()
			indexFileBefore := programBefore.GetSourceFile("/home/projects/TS/p1/src/index.tlua")

			session.DidChangeFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: new(lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{
								Line:      0,
								Character: 0,
							},
							End: lsproto.Position{
								Line:      0,
								Character: 0,
							},
						},
						Text: ";",
					}),
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			programAfter := lsAfter.GetProgram()

			// Unchanged file should be reused
			assert.Equal(t, programAfter.GetSourceFile("/home/projects/TS/p1/src/index.tlua"), indexFileBefore)
		})

		t.Run("change can pull in new files", func(t *testing.T) {
			t.Parallel()
			files := maps.Clone(defaultFiles)
			files["/home/projects/TS/p1/y.tlua"] = `local y = 2; return { y = y };`
			session, _ := projecttestutil.Setup(files)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Verify y.tlua is not initially in the program
			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()
			assert.Check(t, programBefore.GetSourceFile("/home/projects/TS/p1/y.tlua") == nil)

			session.DidChangeFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: new(lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{
								Line:      0,
								Character: 0,
							},
							End: lsproto.Position{
								Line:      0,
								Character: 0,
							},
						},
						Text: "local y = require(\"y\");\n",
					}),
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programAfter := lsAfter.GetProgram()

			// y.tlua should now be included in the program
			assert.Assert(t, programAfter.GetSourceFile("/home/projects/TS/p1/y.tlua") != nil)
		})

		t.Run("single-file change followed by config change reloads program", func(t *testing.T) {
			t.Parallel()
			files := maps.Clone(defaultFiles)
			files["/home/projects/TS/p1/tsconfig.json"] = `{
				"compilerOptions": {
					"noLib": true,
					"module": "nodenext",
					"strict": true
				},
				"include": ["src/index.tlua"]
			}`
			session, utils := projecttestutil.Setup(files)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()
			assert.Equal(t, len(programBefore.GetSourceFiles()), 2)

			session.DidChangeFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{
				{
					Partial: new(lsproto.TextDocumentContentChangePartial{
						Range: lsproto.Range{
							Start: lsproto.Position{
								Line:      0,
								Character: 0,
							},
							End: lsproto.Position{
								Line:      0,
								Character: 0,
							},
						},
						Text: "\n",
					}),
				},
			})

			err = utils.FS().WriteFile("/home/projects/TS/p1/tsconfig.json", `{
				"compilerOptions": {
					"noLib": true,
					"module": "nodenext",
					"strict": true
				},
				"include": ["./**/*"]
			}`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/tsconfig.json",
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programAfter := lsAfter.GetProgram()
			assert.Equal(t, len(programAfter.GetSourceFiles()), 3)
		})
	})

	t.Run("DidCloseFile", func(t *testing.T) {
		t.Parallel()
		t.Run("Configured projects", func(t *testing.T) {
			t.Parallel()
			t.Run("delete a file, close it, recreate it", func(t *testing.T) {
				t.Parallel()
				files := maps.Clone(defaultFiles)
				session, utils := projecttestutil.Setup(files)

				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, files["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)
				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

				assert.NilError(t, utils.FS().Remove("/home/projects/TS/p1/src/x.tlua"))

				session.DidCloseFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
				ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
				assert.NilError(t, err)
				program := ls.GetProgram()
				assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") == nil)

				err = utils.FS().WriteFile("/home/projects/TS/p1/src/x.tlua", "")
				assert.NilError(t, err)

				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, "", lsproto.LanguageKindTypeScript)

				ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
				assert.NilError(t, err)
				program = ls.GetProgram()
				assert.Assert(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") != nil)
				assert.Equal(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua").Text(), "")
			})
		})

		t.Run("Inferred projects", func(t *testing.T) {
			t.Parallel()
			t.Run("delete a file, close it, recreate it", func(t *testing.T) {
				t.Parallel()
				files := maps.Clone(defaultFiles)
				delete(files, "/home/projects/TS/p1/tsconfig.json")
				session, utils := projecttestutil.Setup(files)

				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, files["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)
				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

				err := utils.FS().Remove("/home/projects/TS/p1/src/x.tlua")
				assert.NilError(t, err)

				session.DidCloseFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")

				ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
				assert.NilError(t, err)
				program := ls.GetProgram()
				assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") == nil)

				err = utils.FS().WriteFile("/home/projects/TS/p1/src/x.tlua", "")
				assert.NilError(t, err)

				session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, "", lsproto.LanguageKindTypeScript)

				ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
				assert.NilError(t, err)
				program = ls.GetProgram()
				assert.Assert(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") != nil)
				assert.Equal(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua").Text(), "")
			})

			t.Run("close untitled file", func(t *testing.T) {
				t.Parallel()
				session, _ := projecttestutil.Setup(defaultFiles)

				session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "local x = 1;", lsproto.LanguageKindTypeScript)
				session.DidCloseFile(context.Background(), "untitled:Untitled-1")
				session.DidOpenFile(context.Background(), "untitled:Untitled-2", 1, "", lsproto.LanguageKindTypeScript)
			})
		})
	})

	t.Run("DidSaveFile", func(t *testing.T) {
		t.Parallel()
		t.Run("save event first", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, defaultFiles["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			snapshot := session.Snapshot()
			assert.Equal(t, snapshot.ID(), uint64(1))

			session.DidSaveFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/src/index.tlua",
				},
			})

			session.WaitForBackgroundTasks()
			snapshot = session.Snapshot()
			// We didn't need a snapshot change, but the session overlays should be updated.
			assert.Equal(t, snapshot.ID(), uint64(1))

			// Open another file to force a snapshot update so we can see the changes.
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, defaultFiles["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)
			snapshot = session.Snapshot()
			assert.Equal(t, snapshot.GetFile("/home/projects/TS/p1/src/index.tlua").MatchesDiskText(), true)
		})

		t.Run("watch event first", func(t *testing.T) {
			t.Parallel()
			session, _ := projecttestutil.Setup(defaultFiles)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, defaultFiles["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			snapshot := session.Snapshot()
			assert.Equal(t, snapshot.ID(), uint64(1))

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/src/index.tlua",
				},
			})
			session.DidSaveFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")

			session.WaitForBackgroundTasks()
			snapshot = session.Snapshot()
			// We didn't need a snapshot change, but the session overlays should be updated.
			assert.Equal(t, snapshot.ID(), uint64(1))

			// Open another file to force a snapshot update so we can see the changes.
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, defaultFiles["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)
			snapshot = session.Snapshot()
			assert.Equal(t, snapshot.GetFile("/home/projects/TS/p1/src/index.tlua").MatchesDiskText(), true)
		})
	})

	t.Run("Source file sharing", func(t *testing.T) {
		t.Parallel()
		t.Run("projects with similar options share source files", func(t *testing.T) {
			t.Parallel()
			files := maps.Clone(defaultFiles)
			files["/home/projects/TS/p2/tsconfig.json"] = `{
				"compilerOptions": {
					"noLib": true,
					"module": "nodenext",
					"strict": true,
					"noCheck": true,
					"rootDir": ".."
				}
			}`
			files["/home/projects/TS/p2/src/index.tlua"] = `local x = require("p1.src.x");`
			session, _ := projecttestutil.Setup(files)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p2/src/index.tlua", 1, files["/home/projects/TS/p2/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			snapshot := session.Snapshot()
			assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 2)

			ls1, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program1 := ls1.GetProgram()

			ls2, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p2/src/index.tlua")
			assert.NilError(t, err)
			program2 := ls2.GetProgram()

			assert.Equal(
				t,
				program1.GetSourceFile("/home/projects/TS/p1/src/x.tlua"),
				program2.GetSourceFile("/home/projects/TS/p1/src/x.tlua"),
			)
		})
	})

	t.Run("DidChangeWatchedFiles", func(t *testing.T) {
		t.Parallel()

		t.Run("change open file", func(t *testing.T) {
			t.Parallel()
			files := maps.Clone(defaultFiles)
			session, utils := projecttestutil.Setup(files)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, files["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()

			err = utils.FS().WriteFile("/home/projects/TS/p1/src/x.tlua", `local x = 2; return { x = x };`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/src/x.tlua",
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			// Program should remain the same since the file is open and changes are handled through DidChangeTextDocument
			assert.Equal(t, programBefore, lsAfter.GetProgram())
		})

		t.Run("change closed program file", func(t *testing.T) {
			t.Parallel()
			files := maps.Clone(defaultFiles)
			session, utils := projecttestutil.Setup(files)

			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			programBefore := lsBefore.GetProgram()

			err = utils.FS().WriteFile("/home/projects/TS/p1/src/x.tlua", `local x = 2; return { x = x };`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/src/x.tlua",
				},
			})

			lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			assert.Check(t, lsAfter.GetProgram() != programBefore)
		})

		t.Run("change program file not in tsconfig root files", func(t *testing.T) {
			t.Parallel()
			for _, workspaceDir := range []string{"/", "/home/projects/TS/p1", "/somewhere/else/entirely"} {
				t.Run("workspaceDir="+strings.ReplaceAll(workspaceDir, "/", "_"), func(t *testing.T) {
					t.Parallel()
					files := map[string]any{
						"/home/projects/TS/p1/tsconfig.json": `{
							"compilerOptions": {
								"noLib": true,
								"module": "nodenext",
								"strict": true
							},
							"files": ["src/index.tlua"]
						}`,
						"/home/projects/TS/p1/src/index.tlua": `local x = require("x");`,
						"/home/projects/TS/p1/x.tlua":         `local x = 1; return { x = x };`,
					}

					session, utils := projecttestutil.SetupWithOptions(files, &project.SessionOptions{
						CurrentDirectory:   workspaceDir,
						DefaultLibraryPath: bundled.LibPath(),
						TypingsLocation:    projecttestutil.TestTypingsLocation,
						PositionEncoding:   lsproto.PositionEncodingKindUTF8,
						WatchEnabled:       true,
						LoggingEnabled:     true,
					})
					session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)
					lsBefore, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
					assert.NilError(t, err)
					programBefore := lsBefore.GetProgram()
					session.WaitForBackgroundTasks()

					assert.Check(t, utils.WatchesFile("/home/projects/ts/p1/x.tlua"))

					err = utils.FS().WriteFile("/home/projects/TS/p1/x.tlua", `local x = 2; return { x = x };`)
					assert.NilError(t, err)

					session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
						{
							Type: lsproto.FileChangeTypeChanged,
							Uri:  "file:///home/projects/TS/p1/x.tlua",
						},
					})

					lsAfter, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
					assert.NilError(t, err)
					assert.Check(t, lsAfter.GetProgram() != programBefore)
				})
			}
		})

		t.Run("change config file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true,
						"strict": false
					}
				}`,
				// Drives the diagnostic delta off noImplicitAny (still toggled by "strict"),
				// since tlua is strict-null always-on and can't use a strictNullChecks toggle.
				"/home/projects/TS/p1/src/index.tlua": `function f(a) { return a; }`,
			}

			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)

			err = utils.FS().WriteFile("/home/projects/TS/p1/tsconfig.json", `{
				"compilerOptions": {
					"noLib": false,
					"strict": true
				}
			}`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/tsconfig.json",
				},
			})

			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)
		})

		t.Run("delete explicitly included file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["src/index.tlua", "src/x.tlua", "src/require.d.tlua"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				// The typed use of x.x errors while the module resolves (number vs
				// string) and goes silent when the unresolved dotted require falls
				// back to `any` -- the delta proves re-resolution happened.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/x.tlua":         `local x = 1; return { x = x };`,
				"/home/projects/TS/p1/src/index.tlua": `local x = require("src.x");
local y: string = x.x;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/x.tlua"))
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)

			err = utils.FS().Remove("/home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/src/x.tlua",
				},
			})

			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			// File name is still in the command line, was explicitly included
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/x.tlua"))
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/x.tlua") == nil)

			// Open file to trigger cleanup
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			snapshot := session.Snapshot()
			assert.Check(t, snapshot.GetFile("/home/projects/TS/p1/src/x.tlua") == nil)
		})

		t.Run("delete wildcard included file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"include": ["src"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua":     `local x = 2; return { x = x };`,
				"/home/projects/TS/p1/src/x.tlua":         `local index = require("src.index"); local y: string = index.x;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/x.tlua", 1, files["/home/projects/TS/p1/src/x.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/index.tlua"))
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/x.tlua"))), 1)

			err = utils.FS().Remove("/home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/src/index.tlua",
				},
			})

			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/x.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			// File name is gone from the command line, was originally included via wildcard
			assert.Check(t, !slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/index.tlua"))
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/x.tlua"))), 0)

			// Open file to trigger cleanup
			session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
			snapshot := session.Snapshot()
			assert.Check(t, snapshot.GetFile("/home/projects/TS/p1/src/index.tlua") == nil)
		})

		t.Run("delete directory with wildcard included files", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"include": ["src"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua": `local x = require("src.sub.x");
local y: string = x.x;`,
				"/home/projects/TS/p1/src/sub/x.tlua": `local x = 1; return { x = x };`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/sub/x.tlua"))
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)

			// Delete the entire subdirectory from the file system.
			err = utils.FS().Remove("/home/projects/TS/p1/src/sub")
			assert.NilError(t, err)

			// When a directory is deleted, the client typically sends a single deletion
			// event for the directory itself. Because the registered glob pattern includes
			// file extensions (e.g. **/*.{ts,...}), the directory path does not match
			// and the event is filtered out, so the server is never notified.
			// Simulate this by sending a delete event for the directory URI.
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/src/sub",
				},
			})

			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			// The directory was deleted, so the file should no longer be in the program.
			assert.Check(t, !slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/sub/x.tlua"))
			// The unresolved dotted require now silently types as any, so the
			// deliberate type error in the typed use disappears.
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)
		})

		t.Run("delete directory with program-only files", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["src/index.tlua", "src/require.d.tlua"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua": `local x = require("src.sub.x");
local y: string = x.x;`,
				"/home/projects/TS/p1/src/sub/x.tlua": `local x = 1; return { x = x };`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/index.tlua"))
			// x.tlua is not in "files" but is pulled in via the require.
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/sub/x.tlua") != nil)
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)

			// Delete the entire subdirectory from the file system.
			err = utils.FS().Remove("/home/projects/TS/p1/src/sub")
			assert.NilError(t, err)

			// Send a delete event for the directory URI.
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/src/sub",
				},
			})

			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			// The directory was deleted, so the file should no longer be resolvable.
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/sub/x.tlua") == nil)
			// The unresolved dotted require now silently types as any, so the
			// deliberate type error in the typed use disappears.
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)
		})

		t.Run("delete sibling folder schedules diagnostics refresh", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["index.tlua"]
				}`,
				"/home/projects/TS/p1/index.tlua": `local content = require("f.content");
local value = content.content;
return { value = value };`,
				"/home/projects/TS/p1/f/content.tlua": `local content = 1; return { content = content };`,
			}
			session, utils := projecttestutil.Setup(files)
			contentURI := lsproto.DocumentUri("file:///home/projects/TS/p1/f/content.tlua")
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), contentURI, 1, files["/home/projects/TS/p1/f/content.tlua"].(string), lsproto.LanguageKindTypeScript)

			_, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/index.tlua")
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()

			baselineRefreshCount := len(utils.Client().RefreshDiagnosticsCalls())

			err = utils.FS().Remove("/home/projects/TS/p1/f")
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/f",
				},
			})
			session.DidCloseFile(context.Background(), contentURI)
			session.WaitForBackgroundTasks()

			refreshCount := len(utils.Client().RefreshDiagnosticsCalls())
			assert.Assert(t, refreshCount > baselineRefreshCount,
				"expected RefreshDiagnostics to be called after deleting /home/projects/TS/p1/f, got %d calls (baseline %d)",
				refreshCount, baselineRefreshCount)
		})

		t.Run("delete sibling folder schedules diagnostics refresh after opening third file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["index.tlua", "third.tlua"]
				}`,
				"/home/projects/TS/p1/index.tlua": `local content = require("f.content");
local value = content.content;
return { value = value };`,
				"/home/projects/TS/p1/f/content.tlua": `local content = 1; return { content = content };`,
				"/home/projects/TS/p1/third.tlua":     `local third = 3; return { third = third };`,
			}
			session, utils := projecttestutil.Setup(files)
			contentURI := lsproto.DocumentUri("file:///home/projects/TS/p1/f/content.tlua")
			thirdURI := lsproto.DocumentUri("file:///home/projects/TS/p1/third.tlua")
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/index.tlua", 1, files["/home/projects/TS/p1/index.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.DidOpenFile(context.Background(), contentURI, 1, files["/home/projects/TS/p1/f/content.tlua"].(string), lsproto.LanguageKindTypeScript)

			_, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/index.tlua")
			assert.NilError(t, err)
			session.WaitForBackgroundTasks()

			baselineRefreshCount := len(utils.Client().RefreshDiagnosticsCalls())

			err = utils.FS().Remove("/home/projects/TS/p1/f")
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeDeleted,
					Uri:  "file:///home/projects/TS/p1/f",
				},
			})
			session.DidOpenFile(context.Background(), thirdURI, 1, files["/home/projects/TS/p1/third.tlua"].(string), lsproto.LanguageKindTypeScript)
			session.WaitForBackgroundTasks()

			refreshCount := len(utils.Client().RefreshDiagnosticsCalls())
			assert.Assert(t, refreshCount > baselineRefreshCount,
				"expected RefreshDiagnostics to be called after deleting /home/projects/TS/p1/f and opening /home/projects/TS/p1/third.tlua, got %d calls (baseline %d)",
				refreshCount, baselineRefreshCount)
		})

		t.Run("create explicitly included file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["src/index.tlua", "src/y.tlua", "src/require.d.tlua"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua": `local y = require("src.y");
local s: string = y.y;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()

			// Initially the unresolved dotted require silently types as any,
			// so there is no diagnostic while y.tlua is missing.
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)

			// Add the missing file
			err = utils.FS().WriteFile("/home/projects/TS/p1/src/y.tlua", `local y = 1; return { y = y };`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/y.tlua",
				},
			})

			// The module now resolves: the deliberate type error in the typed use
			// (number vs string) proves re-resolution happened.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/y.tlua") != nil)
		})

		t.Run("create failed lookup location", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["src/index.tlua", "src/require.d.tlua"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua": `local z = require("src.z");
local s: string = z.z;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()

			// Initially the unresolved dotted require silently types as any,
			// so there is no diagnostic while z.tlua is missing.
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)

			// Add a new file through failed lookup watch
			err = utils.FS().WriteFile("/home/projects/TS/p1/src/z.tlua", `local z = 1; return { z = z };`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/z.tlua",
				},
			})

			// The failed lookup was invalidated: z resolves now, and the deliberate
			// type error in the typed use (number vs string) proves it.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/z.tlua") != nil)
		})

		t.Run("create wildcard included file", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"include": ["src"]
				}`,
				"/home/projects/TS/p1/src/index.tlua": `a;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()

			// Initially should have an error because declaration for 'a' is missing
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)

			// Add a new file through wildcard watch. Every .tlua file is a module,
			// so a global `a` can only come from a declaration file.
			err = utils.FS().WriteFile("/home/projects/TS/p1/src/a.d.tlua", `declare a: number;`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/a.d.tlua",
				},
			})

			// Error should be resolved and the new file should be included in the program
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/a.d.tlua") != nil)
		})

		t.Run("irrelevant extension changes are filtered out", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"include": ["src"]
				}`,
				"/home/projects/TS/p1/src/index.tlua": `local x = 1;`,
				"/home/projects/TS/p1/src/data.txt":   `some text`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)
			oldProgram := program

			// Modify an irrelevant file and send change/create events for files with
			// extensions that are not relevant to TypeScript compilation.
			err = utils.FS().WriteFile("/home/projects/TS/p1/src/data.txt", `updated text`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/TS/p1/src/data.txt",
				},
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/styles.css",
				},
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/image.png",
				},
			})

			// The program should not have been rebuilt since all events had irrelevant extensions.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, program, oldProgram, "program should not be rebuilt for irrelevant extension changes")
		})

		t.Run("pnpm install links local package", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/pnpm/pnpm-workspace.yaml": `packages:
  - 'packages/*'`,
				"/home/projects/pnpm/packages/alpha/package.json": `{ "name": "repo-alpha", "main": "index.tlua" }`,
				"/home/projects/pnpm/packages/alpha/tsconfig.json": `{
					"compilerOptions": { "noLib": true, "composite": true }
				}`,
				"/home/projects/pnpm/packages/alpha/index.tlua":  `local alpha = 1; return { alpha = alpha };`,
				"/home/projects/pnpm/packages/beta/package.json": `{ "name": "repo-beta" }`,
				"/home/projects/pnpm/packages/beta/tsconfig.json": `{
					"compilerOptions": { "noLib": true }
				}`,
				"/home/projects/pnpm/packages/beta/index.tlua": `local alpha = require("repo-alpha");`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/pnpm/packages/beta/index.tlua", 1, files["/home/projects/pnpm/packages/beta/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Before pnpm install: the require is unresolved because node_modules/repo-alpha
			// doesn't exist, so alpha's source is not in the program. (A bare-specifier require
			// is typed by the declared overloads and never produces a checker diagnostic, so
			// resolution is observed through program membership.)
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/pnpm/packages/beta/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Check(t, program.GetSourceFile("/home/projects/pnpm/packages/alpha/index.tlua") == nil)

			// Simulate pnpm install: create a symlink from beta's node_modules/@repo/alpha to packages/alpha.
			mapFS := utils.FsFromFileMap().FSys().(*vfstest.MapFS)
			err = mapFS.MkdirAll("home/projects/pnpm/packages/beta/node_modules", fs.ModePerm)
			assert.NilError(t, err)
			mapFS.AddSymlink("home/projects/pnpm/packages/beta/node_modules/repo-alpha", "home/projects/pnpm/packages/alpha")

			// Fire watch events mimicking what VS Code sends for a pnpm install.
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{Type: lsproto.FileChangeTypeCreated, Uri: "file:///home/projects/pnpm/packages/beta/node_modules"},
				{Type: lsproto.FileChangeTypeCreated, Uri: "file:///home/projects/pnpm/packages/beta/node_modules/repo-alpha"},
				{Type: lsproto.FileChangeTypeCreated, Uri: "file:///home/projects/pnpm/pnpm-lock.yaml"},
				{Type: lsproto.FileChangeTypeChanged, Uri: "file:///home/projects/pnpm/packages/beta/node_modules/.bin/tsc"},
				{Type: lsproto.FileChangeTypeChanged, Uri: "file:///home/projects/pnpm/packages/beta/node_modules/.bin/tsserver"},
			})

			// After pnpm install: the require should resolve, pulling alpha's source
			// (by its realpath through the symlink) into the program.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/pnpm/packages/beta/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Check(t, program.GetSourceFile("/home/projects/pnpm/packages/alpha/index.tlua") != nil)
		})

		t.Run("symlinked node_modules package.json change invalidates resolution", func(t *testing.T) {
			t.Parallel()
			// Set up a project that requires a symlinked node_modules package.
			// The package resolves via "main" in package.json to dist/index.lua.
			files := map[string]any{
				"/home/projects/myproject/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true,
						"module": "nodenext"
					},
					"files": ["src/index.tlua"]
				}`,
				"/home/projects/myproject/src/index.tlua": `local mylib = require("mylib");`,
				// The real package lives as a sibling directory
				"/home/projects/mylib/package.json": `{
					"name": "mylib",
					"main": "dist/index.tlua"
				}`,
				"/home/projects/mylib/dist/index.tlua": `local foo = 1; return { foo = foo };`,
				// node_modules/mylib is a symlink to the sibling
				"/home/projects/myproject/node_modules/mylib": vfstest.Symlink("/home/projects/mylib"),
			}

			session, utils := projecttestutil.SetupWithOptions(files, &project.SessionOptions{
				CurrentDirectory:   "/home/projects/myproject",
				DefaultLibraryPath: bundled.LibPath(),
				TypingsLocation:    projecttestutil.TestTypingsLocation,
				PositionEncoding:   lsproto.PositionEncodingKindUTF8,
				WatchEnabled:       true,
				LoggingEnabled:     true,
			})
			session.DidOpenFile(context.Background(), "file:///home/projects/myproject/src/index.tlua", 1, files["/home/projects/myproject/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Initial state: the require resolves successfully via package.json main ->
			// dist/index.tlua. (A bare-specifier require is typed by the declared
			// overloads and never produces a checker diagnostic, so resolution is
			// observed through program membership of the resolved module file.)
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/myproject/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			session.WaitForBackgroundTasks()
			assert.Check(t, program.GetSourceFile("/home/projects/mylib/dist/index.tlua") != nil, "require should resolve initially")

			// Assert: watched file globs cover the realpath of package.json and dist/index.d.tlua.
			// With a workspace dir set, watchers use RelativePattern with a base URI.
			assert.Check(t, utils.WatchesFile("/home/projects/mylib/package.json"), "realpath of package.json should be watched")
			assert.Check(t, utils.WatchesFile("/home/projects/mylib/dist/index.tlua"), "realpath of dist/index.tlua should be watched")

			// Edit package.json to remove "main" field
			err = utils.FS().WriteFile("/home/projects/mylib/package.json", `{
				"name": "mylib"
			}`)
			assert.NilError(t, err)

			// Fire watch event for the realpath of the changed package.json.
			// A real editor would fire this for the realpath since it watches realpaths.
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeChanged,
					Uri:  "file:///home/projects/mylib/package.json",
				},
			})

			// After removing "main" from package.json, the require should no longer
			// resolve, so the declaration file drops out of the program.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/myproject/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Assert(t, program.GetSourceFile("/home/projects/mylib/dist/index.tlua") == nil, "require should fail after removing main from package.json")
		})

		t.Run("create file in non-existent directory", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"files": ["src/index.tlua", "src/require.d.tlua"]
				}`,
				// noLib leaves `require` undeclared, so declare it in the fixture.
				"/home/projects/TS/p1/src/require.d.tlua": `declare function require(module: string): any;`,
				"/home/projects/TS/p1/src/index.tlua": `local helper = require("src.lib.helper");
local s: string = helper.helper;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Initially the unresolved dotted require silently types as any (the
			// src/lib/ directory doesn't even exist), so there is no diagnostic.
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 0)

			// Create the directory and file.
			err = utils.FS().WriteFile("/home/projects/TS/p1/src/lib/helper.tlua", `local helper = 1; return { helper = helper };`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/lib/helper.tlua",
				},
			})

			// The module resolves now; the deliberate type error in the typed use
			// (number vs string) proves the failed lookup was invalidated even
			// though its directory did not exist when the watch was set up.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Equal(t, len(program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), program.GetSourceFile("/home/projects/TS/p1/src/index.tlua"))), 1)
			assert.Check(t, program.GetSourceFile("/home/projects/TS/p1/src/lib/helper.tlua") != nil)
		})

		t.Run("create symlink directory matching include pattern", func(t *testing.T) {
			t.Parallel()
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true
					},
					"include": ["src"]
				}`,
				"/home/projects/TS/p1/src/index.tlua":   `local x = 1;`,
				"/home/projects/TS/shared/utils.tlua":   `local util = "hello";`,
				"/home/projects/TS/shared/helpers.tlua": `local helper = 42;`,
			}
			session, utils := projecttestutil.Setup(files)
			session.DidOpenFile(context.Background(), "file:///home/projects/TS/p1/src/index.tlua", 1, files["/home/projects/TS/p1/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program := ls.GetProgram()

			// Initially, project only has the one file in src/.
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/index.tlua"))
			assert.Check(t, !slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/linked/utils.tlua"))
			assert.Check(t, !slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/linked/helpers.tlua"))

			// Create a symlink directory inside src/ that points to the shared directory.
			mapFS := utils.FsFromFileMap().FSys().(*vfstest.MapFS)
			mapFS.AddSymlink("home/projects/TS/p1/src/linked", "home/projects/TS/shared")

			// Send directory creation event (what VS Code sends when a symlink directory appears).
			session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
				{
					Type: lsproto.FileChangeTypeCreated,
					Uri:  "file:///home/projects/TS/p1/src/linked",
				},
			})

			// After the symlink directory is created, the files inside it should be
			// picked up by the wildcard include pattern.
			ls, err = session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			program = ls.GetProgram()
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/index.tlua"))
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/linked/utils.tlua"))
			assert.Check(t, slices.Contains(program.CommandLine().ParsedConfig.FileNames, "/home/projects/TS/p1/src/linked/helpers.tlua"))
		})
	})

	t.Run("refreshes code lenses and inlay hints when relevant user preferences change", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/src/tsconfig.json": "{}",
			"/src/index.tlua":    "local x = 1;",
		}
		session, utils := projecttestutil.Setup(files)
		session.DidOpenFile(context.Background(), "file:///src/index.tlua", 1, files["/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), lsproto.DocumentUri("file:///src/index.tlua"))
		assert.NilError(t, err)

		session.Configure(lsutil.NewDefaultUserPreferences())
		// Change user preferences for code lens and inlay hints.
		newPrefs := session.Config()
		newPrefs.CodeLens.ReferencesCodeLensEnabled = core.TSTrue
		newPrefs.InlayHints.IncludeInlayFunctionLikeReturnTypeHints = core.TSTrue

		session.Configure(newPrefs)

		codeLensRefreshCalls := utils.Client().RefreshCodeLensCalls()
		inlayHintsRefreshCalls := utils.Client().RefreshInlayHintsCalls()
		assert.Equal(t, len(codeLensRefreshCalls), 1, "expected one RefreshCodeLens call after code lens preference change")
		assert.Equal(t, len(inlayHintsRefreshCalls), 1, "expected one RefreshInlayHints call after inlay hints preference change")
	})

	t.Run("schedules diagnostics refresh when reportStyleChecksAsWarnings changes", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/src/tsconfig.json": "{}",
			"/src/index.tlua":    "local x = 1;",
		}
		session, utils := projecttestutil.Setup(files)
		session.DidOpenFile(context.Background(), "file:///src/index.tlua", 1, files["/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), lsproto.DocumentUri("file:///src/index.tlua"))
		assert.NilError(t, err)
		session.WaitForBackgroundTasks()

		// Record the baseline count of RefreshDiagnostics calls.
		baselineRefreshCount := len(utils.Client().RefreshDiagnosticsCalls())

		// Toggle reportStyleChecksAsWarnings (default is true, so set it to false).
		prefs := lsutil.NewDefaultUserPreferences()
		prefs.ReportStyleChecksAsWarnings = core.TSFalse
		session.Configure(prefs)
		session.WaitForBackgroundTasks()

		refreshCount := len(utils.Client().RefreshDiagnosticsCalls())
		assert.Assert(t, refreshCount > baselineRefreshCount,
			"expected RefreshDiagnostics to be called after reportStyleChecksAsWarnings change, got %d calls (baseline %d)",
			refreshCount, baselineRefreshCount)
	})

	t.Run("config parsing", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/src/tsconfig.json": "{}",
			"/src/index.tlua":    "local x = 1;",
		}
		session, _ := projecttestutil.Setup(files)
		session.DidOpenFile(context.Background(), "file:///src/index.tlua", 1, files["/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), lsproto.DocumentUri("file:///src/index.tlua"))
		assert.NilError(t, err)

		configMap1 := map[string]any{
			"preferences": map[string]any{
				"useAliasesForRenames": true,
				"quoteStyle":           "single",
			},
			"unstable": map[string]any{
				"organizeImportsSort": "ordinalIgnoreCase",
			},
		}
		session.Configure(lsutil.ParseUserPreferences(map[string]any{"js/ts": configMap1}))
		actualConfig1 := session.Config()
		expectedPrefs1 := lsutil.NewDefaultUserPreferences()
		expectedPrefs1.UseAliasesForRename = core.TSTrue
		expectedPrefs1.QuotePreference = lsutil.QuotePreferenceSingle
		expectedPrefs1.OrganizeImportsSort = lsutil.OrganizeImportsSortOrdinalIgnoreCase

		assert.DeepEqual(t, actualConfig1, expectedPrefs1)

		configMap2 := map[string]any{
			"preferences": map[string]any{
				"useAliasesForRenames": false,
				"quoteStyle":           "double",
			},
			"unstable": map[string]any{
				"organizeImportsSort": "ordinal",
			},
		}
		session.Configure(lsutil.ParseUserPreferences(map[string]any{"js/ts": configMap2}))
		actualConfig2 := session.Config()
		expectedPrefs2 := lsutil.NewDefaultUserPreferences()
		expectedPrefs2.UseAliasesForRename = core.TSFalse
		expectedPrefs2.QuotePreference = lsutil.QuotePreferenceDouble
		expectedPrefs2.OrganizeImportsSort = lsutil.OrganizeImportsSortOrdinal

		assert.DeepEqual(t, actualConfig2, expectedPrefs2)
	})

	t.Run("language service for closed files", func(t *testing.T) {
		t.Parallel()

		t.Run("closed file in configured project not yet opened", func(t *testing.T) {
			t.Parallel()
			// Set up a project where a tsconfig exists but no files have been opened.
			// Requesting language service for a file covered by that tsconfig should
			// work even though the file was never opened via didOpen.
			files := map[string]any{
				"/home/projects/TS/p1/tsconfig.json": `{
					"compilerOptions": {
						"noLib": true,
						"strict": true
					},
					"include": ["src"]
				}`,
				"/home/projects/TS/p1/src/index.tlua": `local x: number = 1;`,
			}
			session, _ := projecttestutil.Setup(files)

			// Do NOT open any file. Directly request language service for a closed file
			// that belongs to the configured project.
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/p1/src/index.tlua")
			assert.NilError(t, err)
			assert.Assert(t, ls != nil)
			program := ls.GetProgram()
			assert.Assert(t, program != nil)
			sourceFile := program.GetSourceFile("/home/projects/TS/p1/src/index.tlua")
			assert.Assert(t, sourceFile != nil)
			assert.Equal(t, sourceFile.Text(), `local x: number = 1;`)
		})

		t.Run("closed file with no configured project creates inferred project", func(t *testing.T) {
			t.Parallel()
			// Set up a file that has no tsconfig. Requesting language service for it
			// should create an inferred project even though the file was never opened.
			files := map[string]any{
				"/home/projects/TS/loose/index.tlua": `local greeting: string = "hello";`,
			}
			session, _ := projecttestutil.Setup(files)

			// Do NOT open any file. Directly request language service for a closed file
			// that has no configured project.
			ls, err := session.GetLanguageService(context.Background(), "file:///home/projects/TS/loose/index.tlua")
			assert.NilError(t, err)
			assert.Assert(t, ls != nil)
			program := ls.GetProgram()
			assert.Assert(t, program != nil)
			sourceFile := program.GetSourceFile("/home/projects/TS/loose/index.tlua")
			assert.Assert(t, sourceFile != nil)
			assert.Equal(t, sourceFile.Text(), `local greeting: string = "hello";`)
		})
	})
}

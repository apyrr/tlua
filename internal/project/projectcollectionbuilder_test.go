package project_test

import (
	"context"
	"fmt"
	"maps"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
	"github.com/apyrr/tlua/internal/tspath"
	"gotest.tools/v3/assert"
)

func TestProjectCollectionBuilder(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	t.Run("when project found is solution referencing default project directly", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-src.json"}, "", nil)
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure configured project is found for open file
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json")) != nil)

		// Ensure request can use existing snapshot
		_, err := session.GetLanguageService(context.Background(), uri)
		assert.NilError(t, err)
		requestSnapshot := session.Snapshot()
		assert.Equal(t, requestSnapshot, snapshot)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") != nil, "direct reference should be present")

		// Close the file and open one in an inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should have been released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
	})

	t.Run("when project found is solution referencing default project indirectly", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-indirect1.json", "./tsconfig-indirect2.json"}, "", nil)
		applyIndirectProjectFiles(files, 1, "")
		applyIndirectProjectFiles(files, 2, "")
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure configured project is found for open file
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		srcProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json"))
		assert.Assert(t, srcProject != nil)

		// Verify the default project is the source project
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Equal(t, defaultProject, srcProject)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") != nil, "direct reference should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") != nil, "indirect reference should be present")

		// Close the file and open one in an inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect2.json") == nil)
	})

	t.Run("when project found is solution with disableReferencedProjectLoad referencing default project directly", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-src.json"}, `"disableReferencedProjectLoad": true`, nil)
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure no configured project is created due to disableReferencedProjectLoad
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json")) == nil)

		// Should use inferred project instead
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Assert(t, defaultProject != nil)
		assert.Equal(t, defaultProject.Kind, project.KindInferred)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil, "direct reference should not be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
	})

	t.Run("when project found is solution referencing default project indirectly through disableReferencedProjectLoad", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-indirect1.json"}, "", nil)
		applyIndirectProjectFiles(files, 1, `"disableReferencedProjectLoad": true`)
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure no configured project is created due to disableReferencedProjectLoad in indirect project
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json")) == nil)

		// Should use inferred project instead
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Assert(t, defaultProject != nil)
		assert.Equal(t, defaultProject.Kind, project.KindInferred)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") != nil, "solution direct reference should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil, "indirect reference should not be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") == nil)
	})

	t.Run("when project found is solution referencing default project indirectly through disableReferencedProjectLoad in one but without it in another", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-indirect1.json", "./tsconfig-indirect2.json"}, "", nil)
		applyIndirectProjectFiles(files, 1, `"disableReferencedProjectLoad": true`)
		applyIndirectProjectFiles(files, 2, "")
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure configured project is found through the indirect project without disableReferencedProjectLoad
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		srcProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json"))
		assert.Assert(t, srcProject != nil)

		// Verify the default project is the source project (found through indirect2, not indirect1)
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Equal(t, defaultProject, srcProject)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") != nil, "direct reference 1 should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect2.json") != nil, "direct reference 2 should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") != nil, "indirect reference should be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect1.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-indirect2.json") == nil)
	})

	t.Run("when project found is project with own files referencing the file from referenced project", func(t *testing.T) {
		t.Parallel()
		files := filesForSolutionConfigFile([]string{"./tsconfig-src.json"}, "", []string{`"./own/main.tlua"`})
		files["/user/username/projects/myproject/own/main.tlua"] = `
			local main = require("src.main");
			local bar = main.foo;
			return { bar = bar };
		`
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/src/main.tlua")
		content := files["/user/username/projects/myproject/src/main.tlua"].(string)

		// Ensure configured project is found for open file - should load both projects
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 2)
		srcProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig-src.json"))
		assert.Assert(t, srcProject != nil)
		ancestorProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/tsconfig.json"))
		assert.Assert(t, ancestorProject != nil)

		// Verify the default project is the source project
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Equal(t, defaultProject, srcProject)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") != nil, "solution config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") != nil, "direct reference should be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/user/username/projects/myproject/tsconfig-src.json") == nil)
	})

	t.Run("when file is not part of first config tree found, looks into ancestor folder and its references to find default project", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/src/projects/project/app/Component-demos.tlua": `
                local helpers = require("demos.helpers");
                local demo = helpers.foo;
                return { demo = demo };
            `,
			"/home/src/projects/project/app/Component.tlua": `local Component = 1; return { Component = Component };`,
			"/home/src/projects/project/app/tsconfig.json": `{
				"compilerOptions": {
					"composite": true,
					"outDir": "../app-dist/",
				},
				"include": ["**/*"],
				"exclude": ["**/*-demos.*"],
			}`,
			"/home/src/projects/project/demos/helpers.tlua": "local foo = 1; return { foo = foo };",
			"/home/src/projects/project/demos/tsconfig.json": `{
				"compilerOptions": {
					"composite": true,
					"rootDir": "../",
					"outDir": "../demos-dist/",
				},
				"include": [
					"**/*",
					"../app/**/*-demos.*",
				],
			}`,
			"/home/src/projects/project/tsconfig.json": `{
				"compilerOptions": {
					"outDir": "./dist/",
				},
				"references": [
					{ "path": "./demos/tsconfig.json" },
					{ "path": "./app/tsconfig.json" },
				],
				"files": []
			}`,
		}
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///home/src/projects/project/app/Component-demos.tlua")
		content := files["/home/src/projects/project/app/Component-demos.tlua"].(string)

		// Ensure configured project is found for open file
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 2)
		demoProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/home/src/projects/project/demos/tsconfig.json"))
		assert.Assert(t, demoProject != nil)
		solutionProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/home/src/projects/project/tsconfig.json"))
		assert.Assert(t, solutionProject != nil)

		// Verify the default project is the demos project (not the app project that excludes demos files)
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Equal(t, defaultProject, demoProject)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/app/tsconfig.json") != nil, "app config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/demos/tsconfig.json") != nil, "demos config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.json") != nil, "solution config should be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/app/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/demos/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.json") == nil)
	})

	t.Run("when dts file is next to ts file and included as root in referenced project", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/home/src/projects/project/src/index.d.tlua": `
                 declare global {
                    interface Window {
                        electron: ElectronAPI
                        api: unknown
                    }
                }
            `,
			"/home/src/projects/project/src/index.tlua": `local api = {}`,
			"/home/src/projects/project/tsconfig.json": `{
				"include": [
					"src/*.d.tlua",
				],
				"references": [{ "path": "./tsconfig.node.json" }],
			}`,
			"/home/src/projects/project/tsconfig.node.json": `{
				"include": ["src/**/*"],
                "compilerOptions": {
                    "composite": true,
                },
			}`,
		}
		session, _ := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///home/src/projects/project/src/index.d.tlua")
		content := files["/home/src/projects/project/src/index.d.tlua"].(string)

		// Ensure configured projects are found for open file
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 2)
		rootProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/home/src/projects/project/tsconfig.json"))
		assert.Assert(t, rootProject != nil)

		// Verify the default project is inferred
		defaultProject := snapshot.GetDefaultProject(uri)
		assert.Assert(t, defaultProject != nil)
		assert.Equal(t, defaultProject.Kind, project.KindInferred)

		// Searched configs should be present while file is open
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.json") != nil, "root config should be present")
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.node.json") != nil, "node config should be present")

		// Close the file and open another one in the inferred project
		session.DidCloseFile(context.Background(), uri)
		dummyUri := lsproto.DocumentUri("file:///user/username/workspaces/dummy/dummy.tlua")
		session.DidOpenFile(context.Background(), dummyUri, 1, "local x = 1;", lsproto.LanguageKindTypeScript)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil)

		// Config files should be released
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.json") == nil)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig("/home/src/projects/project/tsconfig.node.json") == nil)
	})

	t.Run("#1630", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/project/lib/tsconfig.json": `{
				"files": ["a.tlua"]
			}`,
			"/project/lib/a.tlua": `local a = 1;`,
			"/project/lib/b.tlua": `local b = 1;`,
			"/project/tsconfig.json": `{
				"files": [],
				"references": [{ "path": "./lib" }],
				"compilerOptions": {
					"disableReferencedProjectLoad": true
				}
			}`,
			"/project/index.tlua": ``,
		}

		session, _ := projecttestutil.Setup(files)

		// opening b.tlua puts /project/lib/tsconfig.json in the config file registry and creates the project,
		// but the project is ultimately not a match
		session.DidOpenFile(context.Background(), "file:///project/lib/b.tlua", 1, files["/project/lib/b.tlua"].(string), lsproto.LanguageKindTypeScript)
		// opening an unrelated file triggers cleanup of /project/lib/tsconfig.json since no open file is part of that project,
		// but will keep the config file in the registry since lib/b.tlua is still open
		session.DidOpenFile(context.Background(), "untitled:Untitled-1", 1, "", lsproto.LanguageKindTypeScript)
		// Opening index.tlua searches /project/tsconfig.json and then checks /project/lib/tsconfig.json without opening it.
		// No early return on config file existence means we try to find an already open project, which returns nil,
		// triggering a crash.
		session.DidOpenFile(context.Background(), "file:///project/index.tlua", 1, files["/project/index.tlua"].(string), lsproto.LanguageKindTypeScript)
	})

	t.Run("inferred project root files are in stable order", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/project/a.tlua": `local a = 1;`,
			"/project/b.tlua": `local b = 1;`,
			"/project/c.tlua": `local c = 1;`,
		}

		session, _ := projecttestutil.Setup(files)

		// b, c, a
		session.DidOpenFile(context.Background(), "file:///project/b.tlua", 1, files["/project/b.tlua"].(string), lsproto.LanguageKindTypeScript)
		session.DidOpenFile(context.Background(), "file:///project/c.tlua", 1, files["/project/c.tlua"].(string), lsproto.LanguageKindTypeScript)
		session.DidOpenFile(context.Background(), "file:///project/a.tlua", 1, files["/project/a.tlua"].(string), lsproto.LanguageKindTypeScript)

		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		inferredProject := snapshot.ProjectCollection.InferredProject()
		assert.Assert(t, inferredProject != nil)
		// It's more bookkeeping to maintain order of opening, since any file can move into or out of
		// the inferred project due to changes in other projects. Order shouldn't matter for correctness,
		// we just want it to be consistent, in case there are observable type ordering issues.
		assert.DeepEqual(t, inferredProject.Program.CommandLine().FileNames(), []string{
			"/project/a.tlua",
			"/project/b.tlua",
			"/project/c.tlua",
		})
	})

	t.Run("project lookup terminates", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/tsconfig.json": `{
				"files": [],
				"references": [
					{
						"path": "./packages/pkg1"
					},
					{
						"path": "./packages/pkg2"
					},
				]
			}`,
			"/packages/pkg1/tsconfig.json": `{
				"include": ["src/**/*.tlua"],
				"compilerOptions": {
					"composite": true,
				},
				"references": [
					{
						"path": "../pkg2"
					},
				]
			}`,
			"/packages/pkg2/tsconfig.json": `{
				"include": ["src/**/*.tlua"],
				"compilerOptions": {
					"composite": true,
				},
				"references": [
					{
						"path": "../pkg1"
					},
				]
			}`,
			"/script.tlua": `local a = 1;`,
		}
		session, _ := projecttestutil.Setup(files)
		session.DidOpenFile(context.Background(), "file:///script.tlua", 1, files["/script.tlua"].(string), lsproto.LanguageKindTypeScript)
		// Test should terminate
	})

	t.Run("file moves to inferred project after import is deleted", func(t *testing.T) {
		t.Parallel()
		// This test verifies that when a node_modules dependency file is open and the require
		// for it is deleted from the project root, requesting language service for the
		// dependency correctly moves it to an inferred project.
		files := map[string]any{
			"/project/tsconfig.json":              `{"compilerOptions": {"strict": true}}`,
			"/project/index.tlua":                 `local dep = require("dep");`,
			"/project/node_modules/dep/init.tlua": `local helper = 1; return { helper = helper };`,
		}
		session, _ := projecttestutil.Setup(files)

		// Step 1: Open the project root file
		rootUri := lsproto.DocumentUri("file:///project/index.tlua")
		session.DidOpenFile(context.Background(), rootUri, 1, files["/project/index.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), rootUri)
		assert.NilError(t, err)

		// Step 2: Open the node_modules dependency file - should be in the configured project
		depUri := lsproto.DocumentUri("file:///project/node_modules/dep/init.tlua")
		session.DidOpenFile(context.Background(), depUri, 1, files["/project/node_modules/dep/init.tlua"].(string), lsproto.LanguageKindTypeScript)

		snapshot := session.Snapshot()
		configuredProject := snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/project/tsconfig.json"))
		assert.Assert(t, configuredProject != nil, "configured project should exist")
		defaultProject := snapshot.GetDefaultProject(depUri)
		assert.Equal(t, defaultProject, configuredProject, "dependency should be in the configured project initially")

		// Step 3: Delete the require from the root file
		session.DidChangeFile(context.Background(), rootUri, 2, []lsproto.TextDocumentContentChangePartialOrWholeDocument{{
			WholeDocument: &lsproto.TextDocumentContentChangeWholeDocument{Text: `// require removed`},
		}})

		// Step 4: Request language service for the dependency - it should now be in an inferred project
		ls, err := session.GetLanguageService(context.Background(), depUri)
		assert.NilError(t, err)
		assert.Assert(t, ls != nil, "language service should be available for dependency")

		snapshot = session.Snapshot()
		defaultProject = snapshot.GetDefaultProject(depUri)
		assert.Assert(t, defaultProject != nil, "dependency should have a default project")
		assert.Equal(t, defaultProject.Kind, project.KindInferred, "dependency should be in an inferred project after import is deleted")
	})

	t.Run("should update project on package.json change", func(t *testing.T) {
		t.Parallel()
		// Set up a project that resolves a node_modules package through its package.json "main" field.
		// The package.json is not a program file, but it IS an affecting location.
		// When it changes, the project should be marked dirty and the program should be rebuilt.
		// Only index.tlua is a root file, so the package's module enters the program
		// purely through the "dep" resolution. (A bare-specifier require is typed by the
		// declared overloads and never produces a checker diagnostic, so resolution
		// is observed through program membership.)
		packageJsonFiles := map[string]any{
			"/home/projects/myproject/tsconfig.json": `{
				"compilerOptions": {
					"module": "nodenext",
					"noLib": true,
					"noEmit": true
				},
				"files": ["src/index.tlua"]
			}`,
			"/home/projects/myproject/node_modules/dep/package.json": `{
				"name": "dep",
				"main": "lib/utils.tlua"
			}`,
			"/home/projects/myproject/src/index.tlua": `local utils = require("dep");`,
			"/home/projects/myproject/node_modules/dep/lib/utils.tlua": `function add(a: number, b: number): number { return a + b; }
return { add = add };`,
		}

		session, utils := projecttestutil.Setup(packageJsonFiles)
		indexUri := lsproto.DocumentUri("file:///home/projects/myproject/src/index.tlua")
		session.DidOpenFile(context.Background(), indexUri, 1, packageJsonFiles["/home/projects/myproject/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

		// Verify initial state: "dep" resolves through main to lib/utils.tlua, so it is in the program
		ls, err := session.GetLanguageService(context.Background(), indexUri)
		assert.NilError(t, err)
		program := ls.GetProgram()
		assert.Check(t, program.GetSourceFile("/home/projects/myproject/node_modules/dep/lib/utils.tlua") != nil, "the main module of dep should be in the program with correct package.json")

		// Now change the package.json to point main at a non-existent file
		err = utils.FS().WriteFile("/home/projects/myproject/node_modules/dep/package.json", `{
			"name": "dep",
			"main": "lib/nonexistent.tlua"
		}`)
		assert.NilError(t, err)
		session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
			{
				Uri:  lsproto.DocumentUri("file:///home/projects/myproject/node_modules/dep/package.json"),
				Type: lsproto.FileChangeTypeChanged,
			},
		})

		ls, err = session.GetLanguageService(context.Background(), indexUri)
		assert.NilError(t, err)
		updatedProgram := ls.GetProgram()
		assert.Check(t, updatedProgram.GetSourceFile("/home/projects/myproject/node_modules/dep/lib/utils.tlua") == nil, "the main module of dep should drop out of the program after package.json change")
	})
}

func filesForSolutionConfigFile(solutionRefs []string, compilerOptions string, ownFiles []string) map[string]any {
	var compilerOptionsStr string
	if compilerOptions != "" {
		compilerOptionsStr = fmt.Sprintf(`"compilerOptions": {
			%s
		},`, compilerOptions)
	}
	var ownFilesStr string
	if len(ownFiles) > 0 {
		ownFilesStr = strings.Join(ownFiles, ",")
	}
	files := map[string]any{
		"/user/username/projects/myproject/tsconfig.json": fmt.Sprintf(`{
			%s
			"files": [%s],
			"references": [
				%s
			]
		}`, compilerOptionsStr, ownFilesStr, strings.Join(core.Map(solutionRefs, func(ref string) string {
			return fmt.Sprintf(`{ "path": "%s" }`, ref)
		}), ",")),
		"/user/username/projects/myproject/tsconfig-src.json": `{
			"compilerOptions": {
				"composite": true,
				"outDir": "./target",
			},
			"include": ["./src/**/*"]
		}`,
		"/user/username/projects/myproject/src/main.tlua": `
			local functions = require("src.helpers.functions");
			return { foo = functions.foo };`,
		"/user/username/projects/myproject/src/helpers/functions.tlua": `local foo = 1; return { foo = foo };`,
	}
	return files
}

func applyIndirectProjectFiles(files map[string]any, projectIndex int, compilerOptions string) {
	maps.Copy(files, filesForIndirectProject(projectIndex, compilerOptions))
}

func filesForIndirectProject(projectIndex int, compilerOptions string) map[string]any {
	files := map[string]any{
		fmt.Sprintf("/user/username/projects/myproject/tsconfig-indirect%d.json", projectIndex): fmt.Sprintf(`{
			"compilerOptions": {
				"composite": true,
				"outDir": "./target/",
				%s
			},
			"files": [
				"./indirect%d/main.tlua"
			],
			"references": [
				{
				"path": "./tsconfig-src.json"
				}
			]
		}`, compilerOptions, projectIndex),
		fmt.Sprintf("/user/username/projects/myproject/indirect%d/main.tlua", projectIndex): `local indirect = 1;`,
	}
	return files
}

package project_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/ls/lsconv"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

func TestProjectReferencesProgram(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	t.Run("program for referenced project", func(t *testing.T) {
		t.Parallel()
		files := filesForReferencedProjectProgram(false)
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/main/main.tlua")
		session.DidOpenFile(context.Background(), uri, 1, files["/user/username/projects/myproject/main/main.tlua"].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		file := p.Program.GetSourceFileByPath(tspath.Path("/user/username/projects/myproject/dependency/fns.tlua"))
		assert.Assert(t, file != nil)
		dtsFile := p.Program.GetSourceFileByPath(tspath.Path("/user/username/projects/myproject/decls/fns.d.tlua"))
		assert.Assert(t, dtsFile == nil)
	})

	t.Run("program with disableSourceOfProjectReferenceRedirect", func(t *testing.T) {
		t.Parallel()
		files := filesForReferencedProjectProgram(true)
		files["/user/username/projects/myproject/decls/fns.d.tlua"] = `
			declare function fn1(): void;
			declare function fn2(): void;
			declare function fn3(): void;
			declare function fn4(): void;
			declare function fn5(): void;
		`
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/main/main.tlua")
		session.DidOpenFile(context.Background(), uri, 1, files["/user/username/projects/myproject/main/main.tlua"].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		file := p.Program.GetSourceFileByPath(tspath.Path("/user/username/projects/myproject/dependency/fns.tlua"))
		assert.Assert(t, file == nil)
		dtsFile := p.Program.GetSourceFileByPath(tspath.Path("/user/username/projects/myproject/decls/fns.d.tlua"))
		assert.Assert(t, dtsFile != nil)
	})

	t.Run("references through symlink with package main", func(t *testing.T) {
		t.Parallel()
		files, aTest, bFoo, bBar := filesForSymlinkReferences(false)
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsconv.FileNameToDocumentURI(aTest)
		session.DidOpenFile(context.Background(), uri, 1, files[aTest].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		fooFile := p.Program.GetSourceFile(bFoo)
		assert.Assert(t, fooFile != nil)
		barFile := p.Program.GetSourceFile(bBar)
		assert.Assert(t, barFile != nil)
	})

	t.Run("references through symlink with package main with preserveSymlinks", func(t *testing.T) {
		t.Parallel()
		files, aTest, _, _ := filesForSymlinkReferences(true)
		// With preserveSymlinks the resolved files keep their node_modules
		// spelling instead of the realpath into packages/B.
		bFoo := "/user/username/projects/myproject/node_modules/b/src/index.tlua"
		bBar := "/user/username/projects/myproject/node_modules/b/src/bar.tlua"
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsconv.FileNameToDocumentURI(aTest)
		session.DidOpenFile(context.Background(), uri, 1, files[aTest].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		fooFile := p.Program.GetSourceFile(bFoo)
		assert.Assert(t, fooFile != nil)
		barFile := p.Program.GetSourceFile(bBar)
		assert.Assert(t, barFile != nil)
	})

	t.Run("references through symlink referencing from subFolder", func(t *testing.T) {
		t.Parallel()
		files, aTest, bFoo, bBar := filesForSymlinkReferencesInSubfolder(false)
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsconv.FileNameToDocumentURI(aTest)
		session.DidOpenFile(context.Background(), uri, 1, files[aTest].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		fooFile := p.Program.GetSourceFile(bFoo)
		assert.Assert(t, fooFile != nil)
		barFile := p.Program.GetSourceFile(bBar)
		assert.Assert(t, barFile != nil)
	})

	t.Run("references through symlink referencing from subFolder with preserveSymlinks", func(t *testing.T) {
		t.Parallel()
		files, aTest, _, _ := filesForSymlinkReferencesInSubfolder(true)
		// With preserveSymlinks the resolved files keep their node_modules
		// spelling instead of the realpath into packages/B.
		bFoo := "/user/username/projects/myproject/node_modules/b/src/foo.tlua"
		bBar := "/user/username/projects/myproject/node_modules/b/src/bar/foo.tlua"
		session, _ := projecttestutil.Setup(files)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 0)

		uri := lsconv.FileNameToDocumentURI(aTest)
		session.DidOpenFile(context.Background(), uri, 1, files[aTest].(string), lsproto.LanguageKindTypeScript)

		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		projects := snapshot.ProjectCollection.Projects()
		p := projects[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		fooFile := p.Program.GetSourceFile(bFoo)
		assert.Assert(t, fooFile != nil)
		barFile := p.Program.GetSourceFile(bBar)
		assert.Assert(t, barFile != nil)
	})

	t.Run("references through symlink with directory index subpath (issue 4373)", func(t *testing.T) {
		t.Parallel()
		files, aIndex, bFile := filesForDirectorySubpathSymlinkReferences()
		session, _ := projecttestutil.Setup(files)
		uri := lsconv.FileNameToDocumentURI(aIndex)
		session.DidOpenFile(context.Background(), uri, 1, files[aIndex].(string), lsproto.LanguageKindTypeScript)

		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		p := snapshot.ProjectCollection.Projects()[0]
		assert.Equal(t, p.Kind, project.KindConfigured)

		// The import must redirect to source, so the source file is part of the program...
		assert.Assert(t, p.Program.GetSourceFile(bFile) != nil)
		// ...and there must be no `TS2307: Cannot find module 'b/lib/File'` diagnostic.
		diagnostics := p.Program.GetSemanticDiagnostics(projecttestutil.WithRequestID(t.Context()), p.Program.GetSourceFile(aIndex))
		assert.Equal(t, len(diagnostics), 0)
	})

	t.Run("when new file is added to referenced project", func(t *testing.T) {
		t.Parallel()
		files := filesForReferencedProjectProgram(false)
		session, utils := projecttestutil.Setup(files)
		uri := lsproto.DocumentUri("file:///user/username/projects/myproject/main/main.tlua")
		session.DidOpenFile(context.Background(), uri, 1, files["/user/username/projects/myproject/main/main.tlua"].(string), lsproto.LanguageKindTypeScript)
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		programBefore := snapshot.ProjectCollection.Projects()[0].Program

		err := utils.FS().WriteFile("/user/username/projects/myproject/dependency/fns2.tlua", `local x = 2;`)
		assert.NilError(t, err)
		session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
			{
				Type: lsproto.FileChangeTypeCreated,
				Uri:  "file:///user/username/projects/myproject/dependency/fns2.tlua",
			},
		})

		_, err = session.GetLanguageService(context.Background(), uri)
		assert.NilError(t, err)
		snapshot = session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Check(t, snapshot.ProjectCollection.Projects()[0].Program != programBefore)
	})

	t.Run("dropped project reference does not crash on later change to the dropped config", func(t *testing.T) {
		t.Parallel()
		// Regression test for https://github.com/microsoft/typescript-go/issues/3942.
		// A project reference dropped from a tsconfig used to leave a stale entry in
		// the referenced config's retainingProjects. When the referencing project was
		// later deleted and the referenced config subsequently changed, the stale entry
		// named a project that no longer existed, crashing markProjectsAffectedByConfigChanges.
		files := map[string]any{
			"/user/username/projects/myproject/main/tsconfig.json": `{
				"compilerOptions": {
					"composite": true,
					"rootDir": ".."
				},
				"references": [{ "path": "../dependency" }]
			}`,
			"/user/username/projects/myproject/main/main.tlua": `
				local fns = require("dependency.fns");
				fns.fn1();
			`,
			"/user/username/projects/myproject/dependency/tsconfig.json": `{
				"compilerOptions": {
					"composite": true
				}
			}`,
			"/user/username/projects/myproject/dependency/fns.tlua": `
				function fn1() { }
				return { fn1 = fn1 };
			`,
			"/user/username/projects/myproject/other/tsconfig.json": `{
				"compilerOptions": {
					"composite": true
				}
			}`,
			"/user/username/projects/myproject/other/other.tlua": `local y = 1;`,
		}

		session, utils := projecttestutil.Setup(files)

		mainURI := lsproto.DocumentUri("file:///user/username/projects/myproject/main/main.tlua")
		mainContent := files["/user/username/projects/myproject/main/main.tlua"].(string)
		otherURI := lsproto.DocumentUri("file:///user/username/projects/myproject/other/other.tlua")
		otherContent := files["/user/username/projects/myproject/other/other.tlua"].(string)

		// 1. Open main.tlua. The main project's program resolves the `../dependency`
		//    project reference, so the dependency config's retainingProjects gains
		//    the main project path.
		session.DidOpenFile(context.Background(), mainURI, 1, mainContent, lsproto.LanguageKindTypeScript)
		session.WaitForBackgroundTasks()
		snapshot := session.Snapshot()
		assert.Equal(t, len(snapshot.ProjectCollection.Projects()), 1)
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig(tspath.Path("/user/username/projects/myproject/dependency/tsconfig.json")) != nil)

		// 2. Remove the project reference from main/tsconfig.json and rebuild.
		//    The new program no longer references dependency.
		err := utils.FS().WriteFile("/user/username/projects/myproject/main/tsconfig.json", `{
			"compilerOptions": {
				"composite": true,
				"rootDir": ".."
			}
		}`)
		assert.NilError(t, err)
		session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
			{
				Type: lsproto.FileChangeTypeChanged,
				Uri:  "file:///user/username/projects/myproject/main/tsconfig.json",
			},
		})
		_, err = session.GetLanguageService(context.Background(), mainURI)
		assert.NilError(t, err)

		// 3. Close main.tlua and open an unrelated file. Opening triggers cleanup of
		//    orphaned configured projects, deleting the main project.
		session.DidCloseFile(context.Background(), mainURI)
		session.DidOpenFile(context.Background(), otherURI, 1, otherContent, lsproto.LanguageKindTypeScript)
		session.WaitForBackgroundTasks()
		snapshot = session.Snapshot()
		assert.Assert(t, snapshot.ProjectCollection.ConfiguredProject(tspath.Path("/user/username/projects/myproject/main/tsconfig.json")) == nil)
		// Dropping the reference releases main from dependency's retainingProjects,
		// so the now-unreferenced dependency config is cleaned up and no stale entry
		// survives to crash a later config change.
		assert.Assert(t, snapshot.ConfigFileRegistry.GetConfig(tspath.Path("/user/username/projects/myproject/dependency/tsconfig.json")) == nil)

		// 4. Change dependency/tsconfig.json and flush. This used to copy the stale
		//    retainingProjects into affectedProjects and crash
		//    markProjectsAffectedByConfigChanges loading the deleted main project.
		err = utils.FS().WriteFile("/user/username/projects/myproject/dependency/tsconfig.json", `{
			"compilerOptions": {
				"composite": true,
				"strict": true
			}
		}`)
		assert.NilError(t, err)
		session.DidChangeWatchedFiles(context.Background(), []*lsproto.FileEvent{
			{
				Type: lsproto.FileChangeTypeChanged,
				Uri:  "file:///user/username/projects/myproject/dependency/tsconfig.json",
			},
		})
		_, err = session.GetLanguageService(context.Background(), otherURI)
		assert.NilError(t, err)
		session.WaitForBackgroundTasks()
	})
}

func filesForReferencedProjectProgram(disableSourceOfProjectReferenceRedirect bool) map[string]any {
	return map[string]any{
		"/user/username/projects/myproject/main/tsconfig.json": fmt.Sprintf(`{
			"compilerOptions": {
				"composite": true,
				"rootDir": ".."%s
			},
			"references": [{ "path": "../dependency" }]
		}`, core.IfElse(disableSourceOfProjectReferenceRedirect, `, "disableSourceOfProjectReferenceRedirect": true`, "")),
		"/user/username/projects/myproject/main/main.tlua": `
			local fns = require("dependency.fns");
			fns.fn1();
			fns.fn2();
			fns.fn3();
			fns.fn4();
			fns.fn5();
		`,
		"/user/username/projects/myproject/dependency/tsconfig.json": `{
			"compilerOptions": {
				"composite": true,
				"declarationDir": "../decls"
			},
		}`,
		"/user/username/projects/myproject/dependency/fns.tlua": `
			function fn1() { }
			function fn2() { }
			function fn3() { }
			function fn4() { }
			function fn5() { }
			return { fn1 = fn1, fn2 = fn2, fn3 = fn3, fn4 = fn4, fn5 = fn5 };
		`,
	}
}

func filesForSymlinkReferences(preserveSymlinks bool) (files map[string]any, aTest string, bFoo string, bBar string) {
	aTest = "/user/username/projects/myproject/packages/A/src/index.tlua"
	bFoo = "/user/username/projects/myproject/packages/B/src/index.tlua"
	bBar = "/user/username/projects/myproject/packages/B/src/bar.tlua"
	files = map[string]any{
		"/user/username/projects/myproject/packages/B/package.json": `{
			"main": "src/index.tlua"
		}`,
		aTest: `
			local b = require("b");
			local bar = require("b.src.bar");
			b.foo();
			bar.bar();
		`,
		bFoo: `function foo() { } return { foo = foo };`,
		bBar: `function bar() { } return { bar = bar };`,
		`/user/username/projects/myproject/node_modules/b`: vfstest.Symlink("/user/username/projects/myproject/packages/B"),
	}
	addConfigForPackage(files, "A", preserveSymlinks, []string{"../B"})
	addConfigForPackage(files, "B", preserveSymlinks, nil)
	return files, aTest, bFoo, bBar
}

func filesForSymlinkReferencesInSubfolder(preserveSymlinks bool) (files map[string]any, aTest string, bFoo string, bBar string) {
	aTest = "/user/username/projects/myproject/packages/A/src/test.tlua"
	bFoo = "/user/username/projects/myproject/packages/B/src/foo.tlua"
	bBar = "/user/username/projects/myproject/packages/B/src/bar/foo.tlua"
	files = map[string]any{
		"/user/username/projects/myproject/packages/B/package.json": `{}`,
		"/user/username/projects/myproject/packages/A/src/test.tlua": `
			local foo = require("b.src.foo");
			local bar = require("b.src.bar.foo");
			foo.foo();
			bar.bar();
		`,
		bFoo: `function foo() { } return { foo = foo };`,
		bBar: `function bar() { } return { bar = bar };`,
		`/user/username/projects/myproject/node_modules/b`: vfstest.Symlink("/user/username/projects/myproject/packages/B"),
	}
	addConfigForPackage(files, "A", preserveSymlinks, []string{"../B"})
	addConfigForPackage(files, "B", preserveSymlinks, nil)
	return files, aTest, bFoo, bBar
}

func filesForDirectorySubpathSymlinkReferences() (files map[string]any, aIndex string, bFile string) {
	aIndex = "/user/username/projects/myproject/packages/a/src/index.tlua"
	bFile = "/user/username/projects/myproject/packages/b/src/File/index.tlua"
	files = map[string]any{
		"/user/username/projects/myproject/packages/b/package.json": `{
			"main": "src/index.tlua"
		}`,
		aIndex: `
			local File = require("b.src.File");
			local result: number = File.helper();
			return { result = result };
		`,
		bFile: `function helper(): number { return 1; } return { helper = helper };`,
		"/user/username/projects/myproject/node_modules/b": vfstest.Symlink("/user/username/projects/myproject/packages/b"),
	}
	addConfigForPackage(files, "a", false, []string{"../b"})
	addConfigForPackage(files, "b", false, nil)
	return files, aIndex, bFile
}

func addConfigForPackage(files map[string]any, packageName string, preserveSymlinks bool, references []string) {
	compilerOptions := map[string]any{
		"outDir":    "lib",
		"rootDir":   "src",
		"composite": true,
	}
	if preserveSymlinks {
		compilerOptions["preserveSymlinks"] = true
	}
	var referencesToAdd []map[string]any
	for _, ref := range references {
		referencesToAdd = append(referencesToAdd, map[string]any{
			"path": ref,
		})
	}
	files[fmt.Sprintf("/user/username/projects/myproject/packages/%s/tsconfig.json", packageName)] = core.Must(core.StringifyJson(map[string]any{
		"compilerOptions": compilerOptions,
		"include":         []string{"src"},
		"references":      referencesToAdd,
	}, "    ", "  "))
}

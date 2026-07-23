package project_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/project"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
	"gotest.tools/v3/assert"
)

func TestBulkCacheInvalidation(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	// Base file structure for testing
	baseFiles := map[string]any{
		"/project/tluaconfig.json": `{
			"compilerOptions": {
				"strict": true,
				"target": "es2015",
				"types": ["node"]
			},
			"include": ["src/**/*"]
		}`,
		"/project/src/index.tlua":     `local helper = require("src.helper"); local h = helper.helper;`,
		"/project/src/helper.tlua":    `local helper = "test"; return { helper = helper };`,
		"/project/src/utils/lib.tlua": `function util() return "util"; end return { util = util };`,

		// Declaration files are global scripts, so siblings are pulled in through
		// reference directives rather than imports.
		"/project/node_modules/@types/node/index.d.tlua": `/// <reference path="./fs.d.tlua" />
/// <reference path="./console.d.tlua" />`,
		"/project/node_modules/@types/node/fs.d.tlua":      ``,
		"/project/node_modules/@types/node/console.d.tlua": ``,
	}

	t.Run("large number of node_modules changes invalidates only node_modules cache", func(t *testing.T) {
		t.Parallel()
		test := func(t *testing.T, fileEvents []*lsproto.FileEvent, expectNodeModulesInvalidation bool) {
			session, utils := projecttestutil.Setup(baseFiles)

			// Open a file to create the project
			session.DidOpenFile(context.Background(), "file:///project/src/index.tlua", 1, baseFiles["/project/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Get initial snapshot and verify config
			ls, err := session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
			assert.NilError(t, err)
			assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetES2015)

			snapshotBefore := session.Snapshot()
			configBefore := snapshotBefore.ConfigFileRegistry

			// Update tluaconfig.json on disk to test that configs don't get reloaded
			err = utils.FS().WriteFile("/project/tluaconfig.json", `{
			"compilerOptions": {
				"strict": true,
				"target": "esnext",
				"types": ["node"]
			},
			"include": ["src/**/*"]
		}`)
			assert.NilError(t, err)
			// Update fs.d.tlua in node_modules
			err = utils.FS().WriteFile("/project/node_modules/@types/node/fs.d.tlua", "new text")
			assert.NilError(t, err)

			// Process the excessive node_modules changes
			session.DidChangeWatchedFiles(context.Background(), fileEvents)

			// Get language service again to trigger snapshot update
			ls, err = session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
			assert.NilError(t, err)

			snapshotAfter := session.Snapshot()
			configAfter := snapshotAfter.ConfigFileRegistry

			// Config should NOT have been reloaded (target should remain ES2015, not esnext)
			assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetES2015, "Config should not have been reloaded for node_modules-only changes")

			// Config registry should be the same instance (no configs reloaded)
			assert.Equal(t, configBefore, configAfter, "Config registry should not have changed for node_modules-only changes")

			fsDtsText := snapshotAfter.GetFile("/project/node_modules/@types/node/fs.d.tlua").Content()
			if expectNodeModulesInvalidation {
				assert.Equal(t, fsDtsText, "new text")
			} else {
				assert.Equal(t, fsDtsText, "")
			}
		}

		t.Run("with file existing in cache", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/node_modules/generated/file%d.lua", lsproto.FileChangeTypeCreated)
			// Include two files in the program to trigger a full program creation.
			// Exclude fs.d.tlua to show that its content still gets invalidated.
			fileEvents = append(fileEvents, &lsproto.FileEvent{
				Uri:  "file:///project/node_modules/@types/node/index.d.tlua",
				Type: lsproto.FileChangeTypeChanged,
			}, &lsproto.FileEvent{
				Uri:  "file:///project/node_modules/@types/node/console.d.tlua",
				Type: lsproto.FileChangeTypeChanged,
			})

			test(t, fileEvents, true)
		})

		t.Run("without file existing in cache", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/node_modules/generated/file%d.lua", lsproto.FileChangeTypeCreated)
			test(t, fileEvents, false)
		})
	})

	t.Run("large number of changes outside node_modules", func(t *testing.T) {
		t.Parallel()
		test := func(t *testing.T, fileEvents []*lsproto.FileEvent, expectConfigReload bool) {
			session, utils := projecttestutil.Setup(baseFiles)

			// Open a file to create the project
			session.DidOpenFile(context.Background(), "file:///project/src/index.tlua", 1, baseFiles["/project/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			// Get initial state
			ls, err := session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
			assert.NilError(t, err)
			assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetES2015)

			// Update tluaconfig.json on disk
			err = utils.FS().WriteFile("/project/tluaconfig.json", `{
			"compilerOptions": {
				"strict": true,
				"target": "esnext",
				"types": ["node"]
			},
			"include": ["src/**/*"]
		}`)
			assert.NilError(t, err)
			// Add root file
			err = utils.FS().WriteFile("/project/src/rootFile.tlua", `console.log("root file")`)
			assert.NilError(t, err)

			session.DidChangeWatchedFiles(context.Background(), fileEvents)
			ls, err = session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
			assert.NilError(t, err)

			if expectConfigReload {
				assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetESNext, "Config should have been reloaded for changes outside node_modules")
				assert.Check(t, ls.GetProgram().GetSourceFile("/project/src/rootFile.tlua") != nil, "New root file should be present")
			} else {
				assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetES2015, "Config should not have been reloaded for changes outside node_modules")
				assert.Check(t, ls.GetProgram().GetSourceFile("/project/src/rootFile.tlua") == nil, "New root file should not be present")
			}
		}

		t.Run("with event matching include glob", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/generated/file%d.tlua", lsproto.FileChangeTypeCreated)
			fileEvents = append(fileEvents, &lsproto.FileEvent{
				Uri:  "file:///project/src/rootFile.tlua",
				Type: lsproto.FileChangeTypeCreated,
			})
			test(t, fileEvents, true)
		})

		t.Run("without event matching include glob", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/generated/file%d.tlua", lsproto.FileChangeTypeCreated)
			test(t, fileEvents, false)
		})
	})

	t.Run("large number of changes outside node_modules causes project reevaluation", func(t *testing.T) {
		t.Parallel()
		session, utils := projecttestutil.Setup(baseFiles)

		// Open a file that will initially use the root tsconfig
		session.DidOpenFile(context.Background(), "file:///project/src/utils/lib.tlua", 1, baseFiles["/project/src/utils/lib.tlua"].(string), lsproto.LanguageKindTypeScript)

		// Initially, the file should use the root project (strict mode)
		snapshot := session.Snapshot()
		initialProject := snapshot.GetDefaultProject("file:///project/src/utils/lib.tlua")
		assert.Equal(t, initialProject.Name(), "/project/tluaconfig.json", "Should initially use root tsconfig")

		// Get language service to verify initial strict mode
		ls, err := session.GetLanguageService(context.Background(), "file:///project/src/utils/lib.tlua")
		assert.NilError(t, err)
		assert.Equal(t, ls.GetProgram().Options().Strict, core.TSTrue, "Should initially use strict mode from root config")

		// Now create the nested tsconfig (this would normally be detected, but we'll simulate a missed event)
		err = utils.FS().WriteFile("/project/src/utils/tluaconfig.json", `{
			"compilerOptions": {
				"strict": false,
				"target": "esnext"
			}
		}`)
		assert.NilError(t, err)

		// Create excessive changes to trigger bulk invalidation
		fileEvents := generateFileEvents(1001, "file:///project/src/generated/file%d.tlua", lsproto.FileChangeTypeCreated)

		// Process the excessive changes - this should trigger project reevaluation
		session.DidChangeWatchedFiles(context.Background(), fileEvents)

		// Get language service - this should now find the nested config and switch projects
		ls, err = session.GetLanguageService(context.Background(), "file:///project/src/utils/lib.tlua")
		assert.NilError(t, err)

		snapshot = session.Snapshot()
		newProject := snapshot.GetDefaultProject("file:///project/src/utils/lib.tlua")

		// The file should now use the nested tsconfig
		assert.Equal(t, newProject.Name(), "/project/src/utils/tluaconfig.json", "Should now use nested tsconfig after bulk invalidation")
		assert.Equal(t, ls.GetProgram().Options().Strict, core.TSFalse, "Should now use non-strict mode from nested config")
		assert.Equal(t, ls.GetProgram().Options().Target, core.ScriptTargetESNext, "Should use esnext target from nested config")
	})

	t.Run("config file names cache", func(t *testing.T) {
		t.Parallel()
		test := func(t *testing.T, fileEvents []*lsproto.FileEvent, expectConfigDiscovery bool, testName string) {
			files := map[string]any{
				"/project/src/index.tlua": `console.log("test");`, // No tsconfig initially
			}
			session, utils := projecttestutil.Setup(files)

			// Open file without tsconfig - should create inferred project
			session.DidOpenFile(context.Background(), "file:///project/src/index.tlua", 1, files["/project/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

			snapshot := session.Snapshot()
			assert.Assert(t, snapshot.ProjectCollection.InferredProject() != nil, "Should have inferred project")
			assert.Equal(t, snapshot.GetDefaultProject("file:///project/src/index.tlua").Kind, project.KindInferred)

			// Create a tsconfig that would affect this file (simulating a missed creation event)
			err := utils.FS().WriteFile("/project/tluaconfig.json", `{
		"compilerOptions": {
			"strict": true
		},
		"include": ["src/**/*"]
	}`)
			assert.NilError(t, err)

			// Process the changes
			session.DidChangeWatchedFiles(context.Background(), fileEvents)

			// Get language service to trigger config discovery
			_, err = session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
			assert.NilError(t, err)

			snapshot = session.Snapshot()
			newProject := snapshot.GetDefaultProject("file:///project/src/index.tlua")

			// Check expected behavior
			if expectConfigDiscovery {
				// Should now use configured project instead of inferred
				assert.Equal(t, newProject.Kind, project.KindConfigured, "Should now use configured project after cache invalidation")
				assert.Equal(t, newProject.Name(), "/project/tluaconfig.json", "Should use the newly discovered tsconfig")
			} else {
				// Should still use inferred project (config file names cache not cleared)
				assert.Assert(t, newProject == snapshot.ProjectCollection.InferredProject(), "Should still use inferred project after node_modules-only changes")
			}
		}

		t.Run("excessive changes only in node_modules does not affect config file names cache", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/node_modules/generated/file%d.lua", lsproto.FileChangeTypeCreated)
			test(t, fileEvents, false, "node_modules changes should not clear config cache")
		})

		t.Run("excessive changes outside node_modules clears config file names cache", func(t *testing.T) {
			t.Parallel()
			fileEvents := generateFileEvents(1001, "file:///project/src/generated/file%d.tlua", lsproto.FileChangeTypeCreated)
			// Presence of any tluaconfig.json file event triggers rediscovery for config for all open files
			fileEvents = append(fileEvents, &lsproto.FileEvent{
				Uri:  lsproto.DocumentUri("file:///project/src/generated/tluaconfig.json"),
				Type: lsproto.FileChangeTypeCreated,
			})
			test(t, fileEvents, true, "non-node_modules changes should clear config cache")
		})
	})

	// Simulate external build tool changing files in dist/ (not included by any project)
	t.Run("excessive changes in dist folder do not invalidate", func(t *testing.T) {
		t.Parallel()
		files := map[string]any{
			"/project/src/index.tlua": `console.log("test");`, // No tsconfig initially
		}
		session, utils := projecttestutil.Setup(files)

		// Open file without tsconfig - should create inferred project
		session.DidOpenFile(context.Background(), "file:///project/src/index.tlua", 1, files["/project/src/index.tlua"].(string), lsproto.LanguageKindTypeScript)

		snapshot := session.Snapshot()
		assert.Equal(t, snapshot.GetDefaultProject("file:///project/src/index.tlua").Kind, project.KindInferred)

		// Create a tsconfig that would affect this file (simulating a missed creation event)
		// This should NOT be discovered after dist-folder changes
		err := utils.FS().WriteFile("/project/tluaconfig.json", `{
			"compilerOptions": {
				"strict": true
			},
			"include": ["src/**/*"]
		}`)
		assert.NilError(t, err)

		// Create excessive changes in dist folder only
		fileEvents := generateFileEvents(1001, "file:///project/dist/generated/file%d.lua", lsproto.FileChangeTypeCreated)
		session.DidChangeWatchedFiles(context.Background(), fileEvents)

		// File should still use inferred project (config file names cache NOT cleared for dist changes)
		_, err = session.GetLanguageService(context.Background(), "file:///project/src/index.tlua")
		assert.NilError(t, err)

		snapshot = session.Snapshot()
		newProject := snapshot.GetDefaultProject("file:///project/src/index.tlua")
		assert.Equal(t, newProject.Kind, project.KindInferred, "dist-folder changes should not cause config discovery")
		// This assertion will fail until we implement logic to ignore dist folder changes
	})

	// Regression test for https://github.com/microsoft/typescript-go/issues/4545
	//
	// A config file entry can be retained (here, by an open file whose default
	// project search fanned out to a referenced config) while its commandLine is
	// nil, because the referenced config file does not exist. A bulk cache
	// invalidation triggered by an excessive number of watch events ranges over
	// all config entries and used to dereference entry.commandLine.ConfigFile
	// without a nil check, crashing.
	//
	// The nil-check-avoiding short circuit `!ok || text != entry.commandLine...`
	// only protects the case where the file cannot be read, so the config file
	// must exist on disk (readable) while the entry's commandLine is still nil to
	// reach the crash.
	t.Run("bulk invalidation with retained config whose command line is nil", func(t *testing.T) {
		t.Parallel()
		appConfig := `{
			"compilerOptions": { "composite": true, "target": "esnext" },
			"include": ["**/*"]
		}`
		files := map[string]any{
			// Solution config references ./app, but app/tluaconfig.json does not exist.
			"/project/tluaconfig.json": `{
				"compilerOptions": { "composite": true },
				"files": [],
				"references": [{ "path": "./app" }]
			}`,
			"/project/app/main.tlua": `export local main = 1;`,
		}

		session, utils := projecttestutil.Setup(files)

		// Open a file in the (non-existent) referenced project. The default project
		// search fans out to app/tluaconfig.json, creating a retained config entry
		// with commandLine == nil and pendingReload == None.
		session.DidOpenFile(context.Background(), "file:///project/app/main.tlua", 1, files["/project/app/main.tlua"].(string), lsproto.LanguageKindTypeScript)
		_, err := session.GetLanguageService(context.Background(), "file:///project/app/main.tlua")
		assert.NilError(t, err)

		// Create the referenced config on disk WITHOUT notifying, so the
		// nil-commandLine entry is not reloaded (pendingReload stays None) but the
		// file becomes readable.
		err = utils.FS().WriteFile("/project/app/tluaconfig.json", appConfig)
		assert.NilError(t, err)

		// Trigger a bulk cache invalidation with an excessive number of watch events.
		// The creation of a config file drives the excessive-change path into
		// invalidateCache, which ranges over all config entries -- including the one
		// whose commandLine is nil -- and used to crash.
		fileEvents := generateFileEvents(1001, "file:///project/app/generated/file%d.tlua", lsproto.FileChangeTypeCreated)
		fileEvents = append(fileEvents, &lsproto.FileEvent{
			Uri:  "file:///project/newdir/tluaconfig.json",
			Type: lsproto.FileChangeTypeCreated,
		})
		session.DidChangeWatchedFiles(context.Background(), fileEvents)
		_, err = session.GetLanguageService(context.Background(), "file:///project/app/main.tlua")
		assert.NilError(t, err)
	})
}

// Helper function to generate excessive file change events
func generateFileEvents(count int, pathTemplate string, changeType lsproto.FileChangeType) []*lsproto.FileEvent {
	var events []*lsproto.FileEvent
	for i := range count {
		events = append(events, &lsproto.FileEvent{
			Uri:  lsproto.DocumentUri(fmt.Sprintf(pathTemplate, i)),
			Type: changeType,
		})
	}
	return events
}

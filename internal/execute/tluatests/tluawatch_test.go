package tluatests

import (
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

func TestWatch(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "watch with no tsconfig",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua": "",
			},
			commandLineArgs: []string{"index.tlua", "--watch"},
		},
		{
			subScenario: "watch with tsconfig and incremental",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      "",
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch", "--incremental"},
		},
		{
			subScenario: "watch skips build when no files change",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x: number = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				noChange,
			},
		},
		{
			subScenario: "watch rebuilds when file is modified",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x: number = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("modify file", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/index.tlua", `local x: number = 2;`)
				}),
			},
		},
		{
			subScenario: "watch rebuilds when source file is deleted",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua":          `local b = require("b");`,
				"/home/src/workspaces/project/b.tlua":          `local b = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "delete imported file",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/b.tlua")
					},
				},
			},
		},
		{
			subScenario: "watch detects new file resolving failed import",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua":          `local b = require("b");`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "create missing file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/b.tlua", `local b = 1;`)
					},
				},
			},
		},
		// Directory-level change detection via imports
		{
			subScenario: "watch detects imported file added in new directory",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local util = require("lib.util");`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "create directory and imported file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/lib/util.tlua", `local util = "hello";`)
					},
				},
			},
		},
		{
			subScenario: "watch detects imported directory removed",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local util = require("lib.util");`,
				"/home/src/workspaces/project/lib/util.tlua":   `local util = "hello";`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "remove directory with imported file",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/lib/util.tlua")
					},
				},
			},
		},
		{
			subScenario: "watch detects import path restructured",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local util = require("lib.util");`,
				"/home/src/workspaces/project/lib/util.tlua":   `local util = "v1";`,
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("move file to new path and update import", func(sys *TestSys) {
					sys.removeNoError("/home/src/workspaces/project/lib/util.tlua")
					sys.writeFileNoError("/home/src/workspaces/project/src/util.tlua", `local util = "v2";`)
					sys.writeFileNoError("/home/src/workspaces/project/index.tlua", `local util = require("src.util");`)
				}),
			},
		},
		// tsconfig include/exclude change detection
		{
			subScenario: "watch rebuilds when tsconfig include pattern adds file",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua": `local x = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("widen include pattern to add src dir", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/extra.tlua", `local extra = 2;`)
					sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", `{
	"compilerOptions": {},
	"include": ["*.tlua", "src/**/*.tlua"]
}`)
				}),
			},
		},
		{
			subScenario: "watch rebuilds when tsconfig is modified to change strict",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x = null; local y: string = x;`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("enable strict mode", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", `{"compilerOptions": {"strict": true}}`)
				}),
			},
		},
		// Path resolution: tsconfig include pointing to non-existent directory
		{
			subScenario: "watch detects file added to previously non-existent include path",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua": `local x = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["index.tlua", "src/**/*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("create src dir with ts file matching include", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/helper.tlua", `local helper = "added";`)
				}),
			},
		},
		{
			subScenario: "watch detects new file in existing include directory",
			files: FileMap{
				"/home/src/workspaces/project/src/a.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["src/**/*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("add new file to existing src directory", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/b.tlua", `local b = 2;`)
				}),
			},
		},
		// Wildcard include: nested subdirectory detection
		{
			subScenario: "watch detects file added in new nested subdirectory",
			files: FileMap{
				"/home/src/workspaces/project/src/a.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["src/**/*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("create nested dir with ts file", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/deep/nested/util.tlua", `local util = "nested";`)
				}),
			},
		},
		{
			subScenario: "watch detects file added in multiple new subdirectories simultaneously",
			files: FileMap{
				"/home/src/workspaces/project/src/a.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["src/**/*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("create multiple new subdirs with files", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/models/user.tlua", `interface User { name: string; }`)
					sys.writeFileNoError("/home/src/workspaces/project/src/utils/format.tlua", `function format(s: string): string return s.trim(); end`)
				}),
			},
		},
		{
			subScenario: "watch detects nested subdirectory removed and recreated",
			files: FileMap{
				"/home/src/workspaces/project/src/lib/helper.tlua": `local helper = "v1";`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"include": ["src/**/*.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption:      "remove nested dir",
					expectedDiff: "incremental has prior state and does not report no-inputs error",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/src/lib/helper.tlua")
					},
				},
				newTscEdit("recreate nested dir with new content", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/src/lib/helper.tlua", `local helper = "v2";`)
				}),
			},
		},
		// Path resolution: import from non-existent node_modules package
		{
			subScenario: "watch detects node modules package added",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local mylib = require("mylib");`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("install package in node_modules", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/package.json", `{"name": "mylib", "main": "index.lua", "types": "index.d.tlua"}`)
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/index.lua", `return { lib = "hello" };`)
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/index.d.tlua", `declare lib: string;`)
				}),
			},
		},
		// Path resolution: node_modules package removed
		{
			subScenario: "watch detects node modules package removed",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":                      `local mylib = require("mylib");`,
				"/home/src/workspaces/project/tluaconfig.json":                 `{}`,
				"/home/src/workspaces/project/node_modules/mylib/package.json": `{"name": "mylib", "main": "index.lua", "types": "index.d.tlua"}`,
				"/home/src/workspaces/project/node_modules/mylib/index.lua":    `return { lib = "hello" };`,
				"/home/src/workspaces/project/node_modules/mylib/index.d.tlua": `declare lib: string;`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "remove node_modules package",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/index.d.tlua")
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/index.lua")
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/package.json")
					},
				},
			},
		},
		// Path resolution: node_modules removed then reinstalled (npm ci after rm -rf)
		{
			subScenario: "watch detects node modules reinstalled after deletion",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":                      `local mylib = require("mylib");`,
				"/home/src/workspaces/project/tluaconfig.json":                 `{}`,
				"/home/src/workspaces/project/node_modules/mylib/package.json": `{"name": "mylib", "main": "index.lua", "types": "index.d.tlua"}`,
				"/home/src/workspaces/project/node_modules/mylib/index.lua":    `return { lib = "hello" };`,
				"/home/src/workspaces/project/node_modules/mylib/index.d.tlua": `declare lib: string;`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "delete node_modules entirely",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/index.d.tlua")
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/index.lua")
						sys.removeNoError("/home/src/workspaces/project/node_modules/mylib/package.json")
					},
				},
				newTscEdit("reinstall node_modules", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/package.json", `{"name": "mylib", "main": "index.lua", "types": "index.d.tlua"}`)
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/index.lua", `return { lib = "hello" };`)
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/index.d.tlua", `declare lib: string;`)
				}),
			},
		},
		// Config file lifecycle
		{
			subScenario: "watch handles tsconfig deleted",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption:      "delete tsconfig",
					expectedDiff: "incremental reports config read error while clean build without tsconfig prints usage help",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/tluaconfig.json")
					},
				},
			},
		},
		{
			subScenario: "watch handles tsconfig with extends base modified",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua": `local x = null; local y: string = x;`,
				"/home/src/workspaces/project/base.json": `{
	"compilerOptions": { "strict": false }
}`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"extends": "./base.json"
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("modify base config to enable strict", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/base.json", `{
	"compilerOptions": { "strict": true }
}`)
				}),
			},
		},
		{
			subScenario: "watch rebuilds when tsconfig is touched but content unchanged",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("touch tsconfig without changing content", func(sys *TestSys) {
					content := sys.readFileNoError("/home/src/workspaces/project/tluaconfig.json")
					sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", content)
				}),
			},
		},
		{
			subScenario: "watch with tsconfig files list entry deleted",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua": `local a = 1;`,
				"/home/src/workspaces/project/b.tlua": `local b = 2;`,
				"/home/src/workspaces/project/tluaconfig.json": `{
	"compilerOptions": {},
	"files": ["a.tlua", "b.tlua"]
}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("delete file listed in files array", func(sys *TestSys) {
					sys.removeNoError("/home/src/workspaces/project/b.tlua")
				}),
			},
		},
		// Module resolution & dependencies
		{
			subScenario: "watch detects module going missing then coming back",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local util = require("util");`,
				"/home/src/workspaces/project/util.tlua":       `local util = "v1";`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "delete util module",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/util.tlua")
					},
				},
				newTscEdit("recreate util module with new content", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/util.tlua", `local util = "v2";`)
				}),
			},
		},
		{
			subScenario: "watch detects nested package installed",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local mylib = require("scope.mylib");`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("install nested package", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/scope/mylib/init.tlua", `local lib: string = "s";`+"\n"+`return { lib = lib };`)
				}),
			},
		},
		{
			subScenario: "watch detects package json main field edited",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":                      `local mylib = require("mylib");` + "\n" + `local v: string = mylib.lib;`,
				"/home/src/workspaces/project/tluaconfig.json":                 `{}`,
				"/home/src/workspaces/project/node_modules/mylib/package.json": `{"name": "mylib", "main": "old.tlua"}`,
				"/home/src/workspaces/project/node_modules/mylib/old.tlua":     `local lib: number = 1;` + "\n" + `return { lib = lib };`,
				"/home/src/workspaces/project/node_modules/mylib/new.tlua":     `local lib: string = "s";` + "\n" + `return { lib = lib };`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "change package.json main field",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/node_modules/mylib/package.json", `{"name": "mylib", "main": "new.tlua"}`)
					},
					expectedDiff: "watch does not yet invalidate require resolutions when a package.json affecting location changes, so the stale main (old.tlua, number) keeps its type error; the clean build resolves the new main",
				},
			},
		},
		{
			subScenario: "watch detects at-types package installed later",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":                         `local lib = require("untyped-lib");`,
				"/home/src/workspaces/project/tluaconfig.json":                    `{}`,
				"/home/src/workspaces/project/node_modules/untyped-lib/index.lua": `return {};`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("install @types for the library", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/@types/untyped-lib/index.d.tlua", `declare value: string;`)
					sys.writeFileNoError("/home/src/workspaces/project/node_modules/@types/untyped-lib/package.json", `{"name": "@types/untyped-lib", "types": "index.d.tlua"}`)
				}),
			},
		},
		// File operations
		{
			subScenario: "watch detects file renamed and renamed back",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local helper = require("helper");`,
				"/home/src/workspaces/project/helper.tlua":     `local helper = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				{
					caption: "rename helper to helper2",
					edit: func(sys *TestSys) {
						sys.renameFileNoError("/home/src/workspaces/project/helper.tlua", "/home/src/workspaces/project/helper2.tlua")
					},
				},
				newTscEdit("rename back to helper", func(sys *TestSys) {
					sys.renameFileNoError("/home/src/workspaces/project/helper2.tlua", "/home/src/workspaces/project/helper.tlua")
				}),
			},
		},
		{
			subScenario: "watch detects file deleted and new file added simultaneously",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua":          `local b = require("b");`,
				"/home/src/workspaces/project/b.tlua":          `local b = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("delete b.tlua and create c.tlua with updated import", func(sys *TestSys) {
					sys.removeNoError("/home/src/workspaces/project/b.tlua")
					sys.writeFileNoError("/home/src/workspaces/project/c.tlua", `local c = 2;`)
					sys.writeFileNoError("/home/src/workspaces/project/a.tlua", `local c = require("c");`)
				}),
			},
		},
		{
			subScenario: "watch handles file rapidly recreated",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local data = require("data");`,
				"/home/src/workspaces/project/data.tlua":       `local val = "original";`,
				"/home/src/workspaces/project/tluaconfig.json": `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("delete and immediately recreate with new content", func(sys *TestSys) {
					sys.removeNoError("/home/src/workspaces/project/data.tlua")
					sys.writeFileNoError("/home/src/workspaces/project/data.tlua", `local val = "recreated";`)
				}),
			},
		},
		// Symlinks — only node_modules symlinks are resolved via Realpath,
		// matching the TypeScript compiler's behavior (see program.tlua:2119).
		{
			subScenario: "watch detects change in symlinked node_modules file",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":                     `local shared = require("shared");`,
				"/home/src/workspaces/shared/index.tlua":                      `local shared = "v1";`,
				"/home/src/workspaces/project/node_modules/shared/index.tlua": vfstest.Symlink("/home/src/workspaces/shared/index.tlua"),
				"/home/src/workspaces/project/tluaconfig.json":                `{}`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("modify symlink target", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/shared/index.tlua", `local shared = "v2";`)
				}),
			},
		},
		// Ancestor fallback stability — when a tsconfig include references a
		// directory that doesn't exist
		{
			subScenario: "watch stability with ancestor directory fallback",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x: number = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{ "include": ["*.tlua", "missing/**/*"] }`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("trivial file change", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/index.tlua", `local x: number = 2;`)
				}),
			},
		},
		// Ancestor fallback: creating deeply nested directories that didn't
		// exist at initial build time should trigger a rebuild and re-watch.
		{
			subScenario: "watch detects file added in deeply nested non-existent include path",
			files: FileMap{
				"/home/src/workspaces/project/index.tlua":      `local x: number = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": `{ "include": ["*.tlua", "deep/nested/dir/**/*"] }`,
			},
			commandLineArgs: []string{"--watch"},
			edits: []*tluaEdit{
				newTscEdit("create deeply nested file matching include", func(sys *TestSys) {
					sys.writeFileNoError("/home/src/workspaces/project/deep/nested/dir/added.tlua", `local added = 1;`)
				}),
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "commandLineWatch")
	}
}

func listToTsconfig(base string, tsconfigOpts ...string) (string, string) {
	optionString := strings.Join(tsconfigOpts, ",\n            ")
	tsconfigText := `{
	"compilerOptions": {
`
	after := "            "
	if base != "" {
		tsconfigText += "            " + base
		after = ",\n            "
	}
	if len(tsconfigOpts) != 0 {
		tsconfigText += after + optionString
	}
	tsconfigText += `
	}
}`
	return tsconfigText, optionString
}

func toTsconfig(base string, compilerOpts string) string {
	tsconfigText, _ := listToTsconfig(base, compilerOpts)
	return tsconfigText
}

func noEmitWatchTestInput(
	subScenario string,
	commandLineArgs []string,
	aText string,
	tsconfigOptions []string,
) *tluaInput {
	noEmitOpt := `"noEmit": true`
	tsconfigText, optionString := listToTsconfig(noEmitOpt, tsconfigOptions...)
	return &tluaInput{
		subScenario:     subScenario,
		commandLineArgs: commandLineArgs,
		files: FileMap{
			"/home/src/workspaces/project/a.tlua":          aText,
			"/home/src/workspaces/project/tluaconfig.json": tsconfigText,
		},
		edits: []*tluaEdit{
			newTscEdit("fix error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/a.tlua", `local a = "hello";`)
			}),
			newTscEdit("emit after fixing error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", toTsconfig("", optionString))
			}),
			newTscEdit("no emit run after fixing error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", toTsconfig(noEmitOpt, optionString))
			}),
			newTscEdit("introduce error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/a.tlua", aText)
			}),
			newTscEdit("emit when error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", toTsconfig("", optionString))
			}),
			newTscEdit("no emit run when error", func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", toTsconfig(noEmitOpt, optionString))
			}),
		},
	}
}

func newTscEdit(name string, edit func(sys *TestSys)) *tluaEdit {
	return &tluaEdit{caption: name, edit: edit}
}

func TestTscNoEmitWatch(t *testing.T) {
	t.Parallel()

	testCases := []*tluaInput{
		noEmitWatchTestInput(
			"syntax errors",
			[]string{"-w"},
			`local a = "hello`,
			nil,
		),
		noEmitWatchTestInput(
			"semantic errors",
			[]string{"-w"},
			`local a: number = "hello"`,
			nil,
		),
		noEmitWatchTestInput(
			"dts errors without dts enabled",
			[]string{"-w"},
			`local a = { private: 10 };`,
			nil,
		),
		noEmitWatchTestInput(
			"dts errors",
			[]string{"-w"},
			`local a = { private: 10 };`,
			[]string{`"declaration": true`},
		),
	}

	for _, test := range testCases {
		test.run(t, "noEmit")
	}
}

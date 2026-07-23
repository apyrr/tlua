package tluatests

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/testutil/stringtestutil"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
)

func TestTscCommandline(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "show help with ExitStatus.DiagnosticsPresent_OutputsSkipped",
			env: map[string]string{
				"TS_TEST_TERMINAL_WIDTH": "120",
			},
			commandLineArgs: nil,
		},
		{
			subScenario:     "show help with ExitStatus.DiagnosticsPresent_OutputsSkipped when host cannot provide terminal width",
			commandLineArgs: nil,
		},
		{
			subScenario: "does not add color when NO_COLOR is set",
			env: map[string]string{
				"NO_COLOR": "true",
			},
			commandLineArgs: nil,
		},
		{
			subScenario: "adds color when FORCE_COLOR is set",
			env: map[string]string{
				"FORCE_COLOR": "true",
			},
			commandLineArgs: nil,
		},
		{
			subScenario: "does not add color when NO_COLOR is set even if FORCE_COLOR is set",
			env: map[string]string{
				"NO_COLOR":    "true",
				"FORCE_COLOR": "true",
			},
			commandLineArgs: nil,
		},
		{
			subScenario:     "when build not first argument",
			commandLineArgs: []string{"--verbose", "--build"},
		},
		{
			subScenario:     "Initialized TSConfig with files options",
			commandLineArgs: []string{"--init", "file0.st", "file1.tlua", "file2.tlua"},
		},
		{
			subScenario:     "Initialized TSConfig with boolean value compiler options",
			commandLineArgs: []string{"--init", "--noUnusedLocals"},
		},
		{
			subScenario:     "Initialized TSConfig with enum value compiler options",
			commandLineArgs: []string{"--init", "--target", "es5", "--jsx", "react"},
		},
		{
			subScenario:     "Initialized TSConfig with list compiler options",
			commandLineArgs: []string{"--init", "--types", "jquery,mocha"},
		},
		{
			subScenario:     "Initialized TSConfig with list compiler options with enum value",
			commandLineArgs: []string{"--init", "--lib", "es5,es2015.core"},
		},
		{
			subScenario:     "Initialized TSConfig with incorrect compiler option",
			commandLineArgs: []string{"--init", "--someNonExistOption"},
		},
		{
			subScenario:     "Initialized TSConfig with incorrect compiler option value",
			commandLineArgs: []string{"--init", "--lib", "nonExistLib,es5,es2015.promise"},
		},
		{
			subScenario:     "Initialized TSConfig with advanced options",
			commandLineArgs: []string{"--init", "--declaration", "--declarationDir", "lib", "--skipLibCheck", "--noErrorTruncation"},
		},
		{
			subScenario:     "Initialized TSConfig with --help",
			commandLineArgs: []string{"--init", "--help"},
		},
		{
			subScenario:     "Initialized TSConfig with --watch",
			commandLineArgs: []string{"--init", "--watch"},
		},
		{
			subScenario:     "Initialized TSConfig with tluaconfig.json",
			commandLineArgs: []string{"--init"},
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
		},
		{
			subScenario:     "help",
			commandLineArgs: []string{"--help"},
		},
		{
			subScenario:     "help all",
			commandLineArgs: []string{"--help", "--all"},
		},
		{
			subScenario:     "Parse --lib option with file name",
			files:           FileMap{"/home/src/workspaces/project/first.tlua": `local Key = Symbol()`},
			commandLineArgs: []string{"--lib", "es6 ", "first.tlua"},
		},
		{
			subScenario:     "option diagnostics are suppressed when there are syntactic errors",
			files:           FileMap{"/home/src/workspaces/project/a.tlua": `local x: = 1;`},
			commandLineArgs: []string{"--strictPropertyInitialization", "--strictNullChecks", "false", "a.tlua"},
		},
		{
			subScenario: "Project is empty string",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
			commandLineArgs: []string{},
		},
		{
			subScenario: "Parse -p",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
			commandLineArgs: []string{"-p", "."},
		},
		{
			subScenario: "Parse -p with path to tsconfig file",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
			commandLineArgs: []string{"-p", "/home/src/workspaces/project/tluaconfig.json"},
		},
		{
			subScenario: "Parse -p with path to tsconfig folder",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
			commandLineArgs: []string{"-p", "/home/src/workspaces/project"},
		},
		{
			subScenario: "Parse -p with empty tsconfig file",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua":      `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": ``,
			},
			commandLineArgs: []string{"-p", "."},
		},
		{
			subScenario:     "Parse enum type options",
			commandLineArgs: []string{"first.tlua", "--module", "nodenext ", "--target", "esnext", "--jsx", "react", "--newLine", "crlf"},
		},
		{
			subScenario: "Parse watch interval option",
			files: FileMap{
				"/home/src/workspaces/project/first.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
			},
			commandLineArgs: []string{"-w", "--watchInterval", "1000"},
		},
		{
			subScenario:     "Parse watch interval option without tluaconfig.json",
			commandLineArgs: []string{"-w", "--watchInterval", "1000"},
		},
		{
			subScenario: "Config with references and empty file and refers to config with noEmit",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`{
					"files": [],
					"references": [
						{
							"path": "./packages/pkg1"
						},
					],
				}`),
				"/home/src/workspaces/project/packages/pkg1/tluaconfig.json": stringtestutil.Dedent(`{
					"compilerOptions": {
						"composite": true,
						"noEmit": true
					},
					"files": [
						"./index.tlua",
					],
				}`),
				"/home/src/workspaces/project/packages/pkg1/index.tlua": `local a = 1;`,
			},
			commandLineArgs: []string{"-p", "."},
		},
		{
			subScenario:     "locale",
			commandLineArgs: []string{"--locale", "cs", "--version"},
		},
		{
			subScenario:     "bad locale",
			commandLineArgs: []string{"--locale", "whoops", "--version"},
		},
	}

	for _, testCase := range testCases {
		testCase.run(t, "commandLine")
	}
}

func TestTscBareFileRootAnchoring(t *testing.T) {
	t.Parallel()
	// A bare file compiled by absolute path from an unrelated cwd must anchor
	// the Lua module search root at the source directory, so require("util")
	// resolves the sibling util.tlua rather than nothing under the process cwd.
	(&tluaInput{
		subScenario: "bare file resolves sibling require from unrelated cwd",
		cwd:         "/home/src/elsewhere",
		files: FileMap{
			"/home/src/app/util.tlua": "local value = 42;\nreturn { value = value };",
			"/home/src/app/main.tlua": "local util = require(\"util\");\nlocal n: number = util.value;\nn;",
		},
		commandLineArgs: []string{"/home/src/app/main.tlua"},
	}).run(t, "commandLine")
}

func TestTscMissingFiles(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "file in tsconfig does not exist",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(
					`{
					"files": ["./src/doesNotExist.tlua"]
					}`,
				),
			},
			commandLineArgs: []string{"-p", "./tluaconfig.json"},
		},
		{
			subScenario: "extensionless file in tsconfig does not exist",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(
					`{
					"files": ["./src/doesNotExist"]
					}`,
				),
			},
			commandLineArgs: []string{"-p", "./tluaconfig.json"},
		},
		{
			subScenario: "extensionless file in extended tsconfig in different folder does not exist",
			files: FileMap{
				"/home/src/workspaces/project/src/tluaconfig.json": stringtestutil.Dedent(
					`{
					"extends": "./../tluaconfig.base.json",
					}`,
				),
				"/home/src/workspaces/project/src/oops.tlua": "local abc = 10;",
				"/home/src/workspaces/project/tluaconfig.base.json": stringtestutil.Dedent(
					`{
					"files": ["./oops"],
					}`,
				),
			},
			commandLineArgs: []string{"-p", "./src/tluaconfig.json"},
		},
	}

	for _, testCase := range testCases {
		testCase.run(t, "commandLine")
	}
}

func TestTscComposite(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "when setting composite false on command line",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5",
						"module": "commonjs",
						"composite": true,
					},
					"include": [
						"src/**/*.tlua",
					],
				}`),
			},
			commandLineArgs: []string{"--composite", "false"},
		},
		{
			subScenario: "when setting composite null on command line",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5",
						"module": "commonjs",
						"composite": true,
					},
					"include": [
						"src/**/*.tlua",
					],
				}`),
			},
			commandLineArgs: []string{"--composite", "null"},
		},
		{
			subScenario: "when setting composite false on command line but has tsbuild info in config",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5",
						"module": "commonjs",
						"composite": true,
						"tsBuildInfoFile": "tluaconfig.json.tluabuildinfo",
					},
					"include": [
						"src/**/*.tlua",
					],
				}`),
			},
			commandLineArgs: []string{"--composite", "false"},
		},
		{
			subScenario: "when setting composite false and tsbuildinfo as null on command line but has tsbuild info in config",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5",
						"module": "commonjs",
						"composite": true,
						"tsBuildInfoFile": "tluaconfig.json.tluabuildinfo",
					},
					"include": [
						"src/**/*.tlua",
					],
				}`),
			},
			commandLineArgs: []string{"--composite", "false", "--tsBuildInfoFile", "null"},
		},
		{
			subScenario: "converting to modules",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"module": "none",
						"composite": true,
					},
				}`),
			},
			edits: []*tluaEdit{
				{
					caption: "convert to modules",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/tluaconfig.json", "none", "es2015")
					},
				},
			},
		},
		{
			subScenario: "synthetic jsx import of ESM module from CJS module no crash no jsx element",
			files: FileMap{
				"/home/src/projects/project/src/main.tlua": "return 42;",
				"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"module": "Node16",
						"jsx": "react-jsx",
						"jsxImportSource": "solid-js",
					},
				}`),
				"/home/src/projects/project/node_modules/solid-js/package.json": stringtestutil.Dedent(`
					{
						"name": "solid-js",
						"type": "module"
					}
				`),
				"/home/src/projects/project/node_modules/solid-js/jsx-runtime.d.tlua": stringtestutil.Dedent(`
					interface JsxIntrinsicElements { div: {}; }
				`),
			},
			cwd: "/home/src/projects/project",
		},
		{
			subScenario: "synthetic jsx import of ESM module from CJS module error on jsx element",
			files: FileMap{
				"/home/src/projects/project/src/main.tsx": "return <div/>;",
				"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"module": "Node16",
						"jsx": "react-jsx",
						"jsxImportSource": "solid-js",
					},
				}`),
				"/home/src/projects/project/node_modules/solid-js/package.json": stringtestutil.Dedent(`
					{
						"name": "solid-js",
						"type": "module"
					}
				`),
				"/home/src/projects/project/node_modules/solid-js/jsx-runtime.d.tlua": stringtestutil.Dedent(`
					interface JsxIntrinsicElements { div: {}; }
				`),
			},
			cwd: "/home/src/projects/project",
		},
	}

	for _, testCase := range testCases {
		testCase.run(t, "composite")
	}
}

func TestTscExtends(t *testing.T) {
	t.Parallel()
	getBuildConfigFileExtendsFileMap := func() FileMap {
		return FileMap{
			"/home/src/workspaces/solution/tluaconfig.json": stringtestutil.Dedent(`
				{
					"references": [
						{ "path": "./shared/tluaconfig.json" },
						{ "path": "./webpack/tluaconfig.json" },
					],
					"files": [],
				}`),
			"/home/src/workspaces/solution/shared/tluaconfig-base.json": stringtestutil.Dedent(`
				{
					"include": ["./typings-base/"],
				}`),
			"/home/src/workspaces/solution/shared/typings-base/globals.d.tlua": `type Unrestricted = any;`,
			"/home/src/workspaces/solution/shared/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "./tluaconfig-base.json",
					"compilerOptions": {
						"composite": true,
						"outDir": "../target-tlua-build/",
						"rootDir": "..",
					},
					"files": ["./index.tlua"],
				}`),
			"/home/src/workspaces/solution/shared/index.tlua": `local a: Unrestricted = 1;`,
			"/home/src/workspaces/solution/webpack/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../shared/tluaconfig-base.json",
					"compilerOptions": {
						"composite": true,
						"outDir": "../target-tlua-build/",
						"rootDir": "..",
					},
					"files": ["./index.tlua"],
					"references": [{ "path": "../shared/tluaconfig.json" }],
				}`),
			"/home/src/workspaces/solution/webpack/index.tlua": `local b: Unrestricted = 1;`,
		}
	}
	getTscExtendsWithSymlinkTestCase := func(builtType string) *tluaInput {
		return &tluaInput{
			subScenario: "resolves the symlink path",
			files: FileMap{
				"/users/user/projects/myconfigs/node_modules/@something/tluaconfig-node/tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "@something/tluaconfig-base/tluaconfig.json",
						"compilerOptions": {
							"removeComments": true
						}
					}
				`),
				"/users/user/projects/myconfigs/node_modules/@something/tluaconfig-base/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": { "composite": true }
					}
				`),
				"/users/user/projects/myproject/src/index.tlua": stringtestutil.Dedent(`
					// some comment
					local x = 10;
				`),
				"/users/user/projects/myproject/src/tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "@something/tluaconfig-node/tluaconfig.json"
					}`),
				"/users/user/projects/myproject/node_modules/@something/tluaconfig-node": vfstest.Symlink("/users/user/projects/myconfigs/node_modules/@something/tluaconfig-node"),
			},
			cwd:             "/users/user/projects/myproject",
			commandLineArgs: []string{builtType, "src", "--extendedDiagnostics"},
		}
	}
	getTscExtendsConfigDirTestCase := func(subScenarioSufix string, commandLineArgs []string, edits []*tluaEdit) *tluaInput {
		return &tluaInput{
			subScenario: "configDir template" + subScenarioSufix,
			files: FileMap{
				"/home/src/projects/configs/first/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../second/tluaconfig.json",
					"include": ["${configDir}/src"],
					"compilerOptions": {
						"typeRoots": ["root1", "${configDir}/root2", "root3"],
						"types": [],
					},
				}`),
				"/home/src/projects/configs/second/tluaconfig.json": stringtestutil.Dedent(`
				{
					"files": ["${configDir}/main.tlua"],
					"compilerOptions": {
						"declarationDir": "${configDir}/decls",
					},
					"watchOptions": {
						"excludeFiles": ["${configDir}/main.tlua"],
					},
				}`),
				"/home/src/projects/myproject/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../configs/first/tluaconfig.json",
					"compilerOptions": {
						"declaration": true,
						"outDir": "outDir",
						"traceResolution": true,
					},
				}`),
				"/home/src/projects/myproject/main.tlua": stringtestutil.Dedent(`
					// some comment
					local y = 10;
				`),
			},
			cwd:             "/home/src/projects/myproject",
			commandLineArgs: commandLineArgs,
			edits:           edits,
		}
	}
	getTscExtendsNonStringPathTestCase := func(propertyName string) *tluaInput {
		return &tluaInput{
			subScenario: "extends config with non-string " + propertyName,
			files: FileMap{
				"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "./base.json",
					}`),
				"/home/src/projects/project/base.json": stringtestutil.Dedent(`
					{
						"` + propertyName + `": [1],
					}`),
				"/home/src/projects/project/main.tlua": `local x = 1;`,
			},
			cwd:             "/home/src/projects/project",
			commandLineArgs: []string{"-p", "tluaconfig.json", "--pretty", "false"},
		}
	}
	getTscExtendsBase := func(baseContents string) FileMap {
		return FileMap{
			"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "./base.json",
				}`),
			"/home/src/projects/project/base.json": stringtestutil.Dedent(baseContents),
			"/home/src/projects/project/main.tlua": `local x = 1;`,
		}
	}
	testCases := []*tluaInput{
		{
			subScenario:     "when building solution with projects extends config with include",
			files:           getBuildConfigFileExtendsFileMap(),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "--v", "--listFiles"},
		},
		getTscExtendsNonStringPathTestCase("include"),
		getTscExtendsNonStringPathTestCase("exclude"),
		getTscExtendsNonStringPathTestCase("files"),
		{
			subScenario: "extends config with mixed valid and non-string include",
			files: getTscExtendsBase(`
				{
					"include": ["main.tlua", 1],
				}`),
			cwd:             "/home/src/projects/project",
			commandLineArgs: []string{"-p", "tluaconfig.json", "--pretty", "false"},
		},
		{
			subScenario:     "when building project uses reference and both extend config with include",
			files:           getBuildConfigFileExtendsFileMap(),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "webpack/tluaconfig.json", "--v", "--listFiles"},
		},
		getTscExtendsWithSymlinkTestCase("-p"),
		getTscExtendsWithSymlinkTestCase("-b"),
		getTscExtendsConfigDirTestCase("", []string{"--explainFiles"}, nil),
		getTscExtendsConfigDirTestCase(" showConfig", []string{"--showConfig"}, nil),
		getTscExtendsConfigDirTestCase(" with commandline", []string{"--explainFiles", "--outDir", "${configDir}/outDir"}, nil),
		getTscExtendsConfigDirTestCase("", []string{"--b", "--explainFiles", "--v"}, nil),
		getTscExtendsConfigDirTestCase("", []string{"--b", "-w", "--explainFiles", "--v"}, []*tluaEdit{
			{
				caption: "edit extended config file",
				edit: func(sys *TestSys) {
					sys.writeFileNoError(
						"/home/src/projects/configs/first/tluaconfig.json",
						stringtestutil.Dedent(`
						{
							"extends": "../second/tluaconfig.json",
							"include": ["${configDir}/src"],
							"compilerOptions": {
								"typeRoots": ["${configDir}/root2"],
								"types": [],
							},
						}`),
					)
				},
			},
		}),
	}

	for _, test := range testCases {
		test.run(t, "extends")
	}
}

func TestForceConsistentCasingInFileNames(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "with relative and non relative file resolutions",
			files: FileMap{
				// No tsconfig: the bare-file search root is the entry file's
				// own directory (src/), so the sibling casing variants
				// (Struct vs struct) resolve to the same file, and the
				// node_modules casing variants resolve via the ancestor walk.
				"/user/username/projects/myproject/src/struct.tlua": stringtestutil.Dedent(`
                    local xs1 = require("fp-ts.lib.Struct");
                    local xs2 = require("fp-ts.lib.struct");
                    local xs3 = require("Struct");
                    local xs4 = require("struct");
                `),
				"/user/username/projects/myproject/node_modules/fp-ts/lib/struct.tlua": `local foo = 10;`,
			},
			cwd:             "/user/username/projects/myproject",
			commandLineArgs: []string{"/user/username/projects/myproject/src/struct.tlua", "--forceConsistentCasingInFileNames", "--explainFiles"},
			ignoreCase:      true,
		},
		{
			subScenario: "when file is included from multiple places with different casing",
			files: FileMap{
				"/home/src/projects/project/src/struct.tlua": stringtestutil.Dedent(`
					local xs1 = require("fp-ts.lib.Struct");
					local xs2 = require("fp-ts.lib.struct");
					local xs3 = require("src.Struct");
					local xs4 = require("src.struct");
				`),
				"/home/src/projects/project/src/anotherFile.tlua": stringtestutil.Dedent(`
					local xs1 = require("fp-ts.lib.Struct");
					local xs2 = require("fp-ts.lib.struct");
					local xs3 = require("src.Struct");
					local xs4 = require("src.struct");
				`),
				"/home/src/projects/project/src/oneMore.tlua": stringtestutil.Dedent(`
					local xs1 = require("fp-ts.lib.Struct");
					local xs2 = require("fp-ts.lib.struct");
					local xs3 = require("src.Struct");
					local xs4 = require("src.struct");
				`),
				"/home/src/projects/project/tluaconfig.json":                    `{}`,
				"/home/src/projects/project/node_modules/fp-ts/lib/struct.tlua": `local foo = 10;`,
			},
			cwd:             "/home/src/projects/project",
			commandLineArgs: []string{"--explainFiles"},
			ignoreCase:      true,
		},
		{
			subScenario: "with type ref from file",
			files: FileMap{
				"/user/username/projects/myproject/src/fileOne.d.tlua": `interface c { }`,
				"/user/username/projects/myproject/src/file2.d.tlua": stringtestutil.Dedent(`
                    /// <reference types="./fileOne.d.tlua"/>
                    declare y: c;
                `),
				"/user/username/projects/myproject/tluaconfig.json": "{ }",
			},
			cwd:             "/user/username/projects/myproject",
			commandLineArgs: []string{"-p", "/user/username/projects/myproject", "--explainFiles", "--traceResolution"},
			ignoreCase:      true,
		},
		{
			subScenario: "with triple slash ref from file",
			files: FileMap{
				"/home/src/workspaces/project/src/c.tlua":      `/// <reference path="./D.tlua"/>`,
				"/home/src/workspaces/project/src/d.tlua":      `interface c { }`,
				"/home/src/workspaces/project/tluaconfig.json": "{ }",
			},
			ignoreCase: true,
		},
		{
			subScenario: "two files exist on disk that differs only in casing",
			files: FileMap{
				"/home/src/workspaces/project/c.tlua": `local D = require("D");`,
				"/home/src/workspaces/project/D.tlua": `local x = 10;`,
				"/home/src/workspaces/project/d.tlua": `local y = 20;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"files": ["c.tlua", "d.tlua"]
					}`),
			},
		},
	}
	for _, test := range testCases {
		test.run(t, "forceConsistentCasingInFileNames")
	}
}

func TestTscIgnoreConfig(t *testing.T) {
	t.Parallel()
	filesWithoutConfig := func() FileMap {
		return FileMap{
			"/home/src/workspaces/project/src/a.tlua": "local a = 10;",
			"/home/src/workspaces/project/src/b.tlua": "local b = 10;",
			"/home/src/workspaces/project/c.tlua":     "local c = 10;",
		}
	}
	filesWithConfig := func() FileMap {
		files := filesWithoutConfig()
		files["/home/src/workspaces/project/tluaconfig.json"] = stringtestutil.Dedent(`
			{
                "include": ["src"],
			}`)
		return files
	}
	getScenarios := func(subScenario string, commandLineArgs []string) []*tluaInput {
		commandLineArgsIgnoreConfig := append(commandLineArgs, "--ignoreConfig")
		return []*tluaInput{
			{
				subScenario:     subScenario,
				files:           filesWithConfig(),
				commandLineArgs: commandLineArgs,
			},
			{
				subScenario:     subScenario + " with --ignoreConfig",
				files:           filesWithConfig(),
				commandLineArgs: commandLineArgsIgnoreConfig,
			},
			{
				subScenario:     subScenario + " when config file absent",
				files:           filesWithoutConfig(),
				commandLineArgs: commandLineArgs,
			},
			{
				subScenario:     subScenario + " when config file absent with --ignoreConfig",
				files:           filesWithoutConfig(),
				commandLineArgs: commandLineArgsIgnoreConfig,
			},
		}
	}
	testCases := slices.Concat(
		getScenarios("without any options", nil),
		getScenarios("specifying files", []string{"src/a.tlua"}),
		getScenarios("specifying project", []string{"-p", "."}),
		getScenarios("mixing project and files", []string{"-p", ".", "src/a.tlua", "c.tlua"}),
	)
	for _, test := range testCases {
		test.run(t, "ignoreConfig")
	}
}

func TestTscIncremental(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "serializing error chain",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "incremental": true,
                        "strict": true,
                        "jsx": "react",
                        "module": "esnext",
                    },
                }`),
				"/home/src/workspaces/project/index.tsx": stringtestutil.Dedent(`
                    interface JsxElementChildrenAttribute { children: {}; }
                    interface JsxIntrinsicElements { div: {} }

                    declare React: any;

                    declare function Component(props: never): any;
                    declare function Component(props: { children?: number }): any;
                    (<Component>
                        <div />
                        <div />
                    </Component>)`),
			},
			edits: noChangeOnlyEdit,
		},
		{
			subScenario: "serializing composite project",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "composite": true,
                        "strict": true,
                        "module": "esnext",
                    },
                }`),
				"/home/src/workspaces/project/index.tsx":  `local a = 1;`,
				"/home/src/workspaces/project/other.tlua": `local b = 2;`,
			},
		},
		{
			subScenario: "change to modifier of class expression field with declaration emit enabled",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": {
						"module": "esnext",
						"declaration": true
					}
				}`),
				"/home/src/workspaces/project/main.tlua": stringtestutil.Dedent(`
                        function logMessage( person: MessageablePerson )
                            console.log( person.message );
                        end`),
				"/home/src/workspaces/project/MessageablePerson.tlua": stringtestutil.Dedent(`
                        local Messageable = function()
                            return {
                                message: 'hello' as string,
                            }
                        end;
                        local wrapper = function() return Messageable() end;
                        type MessageablePerson = ReturnType<typeof wrapper>;`),
				tluaLibPath + "/lib.d.tlua": tluaDefaultLibContent + "\n" + stringtestutil.Dedent(`
					type ReturnType<T extends (...args: any) => any> = T extends (...args: any) => infer R ? R : any;
                    type InstanceType<T extends abstract new (...args: any) => any> = T extends abstract new (...args: any) => infer R ? R : any;`),
			},
			commandLineArgs: []string{"--incremental"},
			edits: []*tluaEdit{
				noChange,
				{
					caption: "narrow message to a literal type",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/MessageablePerson.tlua", "as string", "as 'hello'")
					},
				},
				noChange,
				{
					caption: "widen message back to string",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/MessageablePerson.tlua", "as 'hello'", "as string")
					},
				},
				noChange,
			},
		},
		{
			subScenario: "change to modifier of class expression field",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": { 
						"module": "esnext"
					}
				}`),
				"/home/src/workspaces/project/main.tlua": stringtestutil.Dedent(`
					function logMessage( person: MessageablePerson )
						console.log( person.message );
					end`),
				"/home/src/workspaces/project/MessageablePerson.tlua": stringtestutil.Dedent(`
					local Messageable = function()
						return {
							message: 'hello' as string,
						}
					end;
					local wrapper = function() return Messageable() end;
					type MessageablePerson = ReturnType<typeof wrapper>;`),
				tluaLibPath + "/lib.d.tlua": tluaDefaultLibContent + "\n" + stringtestutil.Dedent(`
					type ReturnType<T extends (...args: any) => any> = T extends (...args: any) => infer R ? R : any;
                    type InstanceType<T extends abstract new (...args: any) => any> = T extends abstract new (...args: any) => infer R ? R : any;`),
			},
			commandLineArgs: []string{"--incremental"},
			edits: []*tluaEdit{
				noChange,
				{
					caption: "narrow message to a literal type",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/MessageablePerson.tlua", "as string", "as 'hello'")
					},
				},
				noChange,
				{
					caption: "widen message back to string",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/MessageablePerson.tlua", "as 'hello'", "as string")
					},
				},
				noChange,
			},
		},
		{
			subScenario: "when passing filename for buildinfo on commandline",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "target": "es5",
                        "module": "commonjs"
                    },
                    "include": [
                        "src/**/*.tlua"
                    ],
                }`),
			},
			commandLineArgs: []string{"--incremental", "--tsBuildInfoFile", ".tluabuildinfo", "--explainFiles"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario: "when passing rootDir from commandline",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "incremental": true,
                        "outDir": "dist"
                    }
                }`),
			},
			commandLineArgs: []string{"--rootDir", "src"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario: "with only dts files",
			files: FileMap{
				"/home/src/workspaces/project/src/main.d.tlua":    "declare x: number;",
				"/home/src/workspaces/project/src/another.d.tlua": "declare y: number;",
				"/home/src/workspaces/project/tluaconfig.json":    "{}",
			},
			commandLineArgs: []string{"--incremental"},
			edits: []*tluaEdit{
				noChange,
				{
					caption: "modify d.tlua file",
					edit: func(sys *TestSys) {
						sys.appendFile("/home/src/workspaces/project/src/main.d.tlua", "declare xy: number;")
					},
				},
			},
		},
		{
			subScenario: "when passing rootDir is in the tsconfig",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "incremental": true,
                        "outDir": "dist",
						"rootDir": "./"
                    }
                }`),
			},
			edits: noChangeOnlyEdit,
		},
		{
			subScenario: "tsbuildinfo has error",
			files: FileMap{
				"/home/src/workspaces/project/main.tlua":                "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json":          "{}",
				"/home/src/workspaces/project/tluaconfig.tluabuildinfo": "Some random string",
			},
			commandLineArgs: []string{"-i"},
			edits: []*tluaEdit{
				{
					caption: "tsbuildinfo written has error",
					edit: func(sys *TestSys) {
						sys.prependFile("/home/src/workspaces/project/tluaconfig.tluabuildinfo", "Some random string")
					},
				},
			},
		},
		{
			subScenario: "when global file is added, the signatures are updated",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": stringtestutil.Dedent(`
                    /// <reference path="./filePresent.tlua"/>
                    /// <reference path="./fileNotFound.tlua"/>
                    function main() end
                `),
				"/home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua": stringtestutil.Dedent(`
                    /// <reference path="./filePresent.tlua"/>
                    /// <reference path="./fileNotFound.tlua"/>
                    function anotherFileWithSameReferenes() end
                `),
				"/home/src/workspaces/project/src/filePresent.tlua": `function something() return 10; end`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": { "composite": true },
                    "include": ["src/**/*.tlua"],
                }`),
			},
			commandLineArgs: []string{},
			edits: []*tluaEdit{
				noChange,
				{
					caption: "Modify main file",
					edit: func(sys *TestSys) {
						sys.appendFile(`/home/src/workspaces/project/src/main.tlua`, `\nsomething();`)
					},
				},
				{
					caption: "Modify main file again",
					edit: func(sys *TestSys) {
						sys.appendFile(`/home/src/workspaces/project/src/main.tlua`, `\nsomething();`)
					},
				},
				{
					caption: "Add new file and update main file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/home/src/workspaces/project/src/newFile.tlua`, "function foo() return 20; end")
						sys.prependFile(
							`/home/src/workspaces/project/src/main.tlua`,
							`/// <reference path="./newFile.tlua"/>
`,
						)
						sys.appendFile(`/home/src/workspaces/project/src/main.tlua`, `foo();`)
					},
				},
				{
					caption: "Write file that could not be resolved",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/home/src/workspaces/project/src/fileNotFound.tlua`, "function something2() return 20; end")
					},
				},
				{
					caption: "Modify main file",
					edit: func(sys *TestSys) {
						sys.appendFile(`/home/src/workspaces/project/src/main.tlua`, `\nsomething();`)
					},
				},
			},
		},
		{
			subScenario: "react-jsx-emit-mode with no backing types found doesnt crash",
			files: FileMap{
				"/home/src/workspaces/project/node_modules/react/jsx-runtime.lua": "", // js needs to be present so there's a resolution result
				"/home/src/workspaces/project/node_modules/@types/react/index.d.tlua": stringtestutil.Dedent(`
					interface JsxElement {}
					interface JsxIntrinsicElements {
						div: {
							propA?: boolean;
						};
					}`), // doesn't contain a jsx-runtime definition
				"/home/src/workspaces/project/src/index.tsx": `local App = function() return <div propA={true}></div> end;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": { 
						"module": "commonjs",
						"jsx": "react-jsx", 
						"incremental": true, 
						"jsxImportSource": "react" 
					} 
				}`),
			},
		},
		{
			subScenario: "react-jsx-emit-mode with no backing types found doesnt crash under --strict",
			files: FileMap{
				"/home/src/workspaces/project/node_modules/react/jsx-runtime.lua": "", // js needs to be present so there's a resolution result
				"/home/src/workspaces/project/node_modules/@types/react/index.d.tlua": stringtestutil.Dedent(`
					interface JsxElement {}
					interface JsxIntrinsicElements {
						div: {
							propA?: boolean;
						};
					}`), // doesn't contain a jsx-runtime definition
				"/home/src/workspaces/project/src/index.tsx": `local App = function() return <div propA={true}></div> end;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": { 
						"module": "commonjs",
						"jsx": "react-jsx", 
						"incremental": true, 
						"jsxImportSource": "react" 
					} 
				}`),
			},
			commandLineArgs: []string{"--strict"},
		},
		{
			subScenario: "change to type that gets used as global through export in another file",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/home/src/workspaces/project/class1.tlua": stringtestutil.Dedent(`
					local constants = require("constants");
					local a: MagicNumber = 1;
					console.log(a);`),
				"/home/src/workspaces/project/constants.tlua": "local constantValue = 1 as const;\ntype MagicNumber = typeof constantValue;\nreturn { constantValue = constantValue };",
			},
			edits: []*tluaEdit{
				{
					caption: "Modify imports used in global file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/constants.tlua", "local constantValue = 2 as const;\ntype MagicNumber = typeof constantValue;\nreturn { constantValue = constantValue };")
					},
					expectedDiff: signaturePropagationDiff,
				},
			},
		},
		{
			subScenario: "change to type that gets used as global through export in another file through indirect import",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/home/src/workspaces/project/class1.tlua": stringtestutil.Dedent(`
					local a: MagicNumber = 1;
					console.log(a);`),
				"/home/src/workspaces/project/constants.tlua": "local constantValue = 1;\nreturn { constantValue = constantValue };",
				"/home/src/workspaces/project/reexport.tlua":  "local constants = require(\"constants\");\ntype MagicNumber = typeof constants.constantValue;",
			},
			edits: []*tluaEdit{
				{
					caption: "Modify imports used in global file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/constants.tlua", "local constantValue = 2;\nreturn { constantValue = constantValue };")
					},
				},
			},
		},
		{
			subScenario: "when file is deleted",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "outDir"
					}
				}`),
				"/home/src/workspaces/project/file1.tlua": `local C = { };`,
				"/home/src/workspaces/project/file2.tlua": `local D = { };`,
			},
			edits: []*tluaEdit{
				{
					caption: "delete file with imports",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/file2.tlua")
					},
				},
			},
		},
		{
			subScenario: "option changes with composite",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
					}
				}`),
				"/home/src/workspaces/project/a.tlua": `local a = 10;local aLocal = 10;return { a = a };`,
				"/home/src/workspaces/project/b.tlua": `local b = 10;local bLocal = 10;return { b = b };`,
				"/home/src/workspaces/project/c.tlua": `local a = require("a");local c = a.a;return { c = c };`,
				"/home/src/workspaces/project/d.tlua": `local b = require("b");local d = b.b;return { d = d };`,
			},
			edits: []*tluaEdit{
				{
					caption:         "with sourceMap",
					commandLineArgs: []string{"--sourceMap"},
					expectedDiff:    declarationEmitErrorReplayDiff,
				},
				{
					caption:      "should re-emit only js so they dont contain sourcemap",
					expectedDiff: declarationEmitErrorReplayDiff,
				},
				{
					caption:         "with declaration should not emit anything",
					commandLineArgs: []string{"--declaration"},
					// discrepancyExplanation: () => [
					// 	`Clean build tsbuildinfo will have compilerOptions with composite and ${option.replace(/-/g, "")}`,
					// 	`Incremental build will detect that it doesnt need to rebuild so tsbuild info is from before which has option composite only`,
					// ],
				},
				noChange,
				{
					caption:         "with declaration and declarationMap",
					commandLineArgs: []string{"--declaration", "--declarationMap"},
				},
				{
					caption: "should re-emit only dts so they dont contain sourcemap",
				},
				{
					caption:         "with emitDeclarationOnly should not emit anything",
					commandLineArgs: []string{"--emitDeclarationOnly"},
					// discrepancyExplanation: () => [
					// 	`Clean build tsbuildinfo will have compilerOptions with composite and ${option.replace(/-/g, "")}`,
					// 	`Incremental build will detect that it doesnt need to rebuild so tsbuild info is from before which has option composite only`,
					// ],
				},
				noChange,
				{
					caption: "local change",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/a.tlua", "Local = 1", "Local = 10")
					},
				},
				{
					caption:         "with declaration should not emit anything",
					commandLineArgs: []string{"--declaration"},
					// discrepancyExplanation: () => [
					// 	`Clean build tsbuildinfo will have compilerOptions with composite and ${option.replace(/-/g, "")}`,
					// 	`Incremental build will detect that it doesnt need to rebuild so tsbuild info is from before which has option composite only`,
					// ],
				},
				{
					caption:         "with inlineSourceMap",
					commandLineArgs: []string{"--inlineSourceMap"},
					expectedDiff:    declarationEmitErrorReplayDiff,
				},
				{
					caption:         "with sourceMap",
					commandLineArgs: []string{"--sourceMap"},
					expectedDiff:    declarationEmitErrorReplayDiff,
				},
				{
					caption: "declarationMap enabling",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/tluaconfig.json", `"composite": true,`, `"composite": true,        "declarationMap": true`)
					},
				},
				{
					caption:         "with sourceMap should not emit d.tlua",
					commandLineArgs: []string{"--sourceMap"},
					expectedDiff:    declarationEmitErrorReplayDiff,
				},
			},
		},
		{
			subScenario: "option changes with incremental",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"incremental": true,
					}
				}`),
				"/home/src/workspaces/project/a.tlua": `local a = 10;local aLocal = 10;return { a = a };`,
				"/home/src/workspaces/project/b.tlua": `local b = 10;local bLocal = 10;return { b = b };`,
				"/home/src/workspaces/project/c.tlua": `local a = require("a");local c = a.a;return { c = c };`,
				"/home/src/workspaces/project/d.tlua": `local b = require("b");local d = b.b;return { d = d };`,
			},
			edits: []*tluaEdit{
				{
					caption:         "with sourceMap",
					commandLineArgs: []string{"--sourceMap"},
				},
				{
					caption: "should re-emit only js so they dont contain sourcemap",
				},
				{
					caption:         "with declaration, emit Dts and should not emit js",
					commandLineArgs: []string{"--declaration"},
				},
				{
					caption:         "with declaration and declarationMap",
					commandLineArgs: []string{"--declaration", "--declarationMap"},
				},
				{
					caption:      "no change",
					expectedDiff: declarationEmitErrorReplayDiff,
					// discrepancyExplanation: () => [
					// 	`Clean build tsbuildinfo will have compilerOptions {}`,
					// 	`Incremental build will detect that it doesnt need to rebuild so tsbuild info is from before which has option declaration and declarationMap`,
					// ],
				},
				{
					caption: "local change",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/a.tlua", "Local = 1", "Local = 10")
					},
					expectedDiff: declarationEmitErrorReplayDiff,
				},
				{
					caption:         "with declaration and declarationMap",
					commandLineArgs: []string{"--declaration", "--declarationMap"},
				},
				{
					caption:      "no change",
					expectedDiff: declarationEmitErrorReplayDiff,
					// discrepancyExplanation: () => [
					// 	`Clean build tsbuildinfo will have compilerOptions {}`,
					// 	`Incremental build will detect that it doesnt need to rebuild so tsbuild info is from before which has option declaration and declarationMap`,
					// ],
				},
				{
					caption:         "with inlineSourceMap",
					commandLineArgs: []string{"--inlineSourceMap"},
				},
				{
					caption:         "with sourceMap",
					commandLineArgs: []string{"--sourceMap"},
				},
				{
					caption: "emit js files",
				},
				{
					caption:         "with declaration and declarationMap",
					commandLineArgs: []string{"--declaration", "--declarationMap"},
				},
				{
					caption:         "with declaration and declarationMap, should not re-emit",
					commandLineArgs: []string{"--declaration", "--declarationMap"},
				},
			},
		},
		{
			subScenario: "when there is bind diagnostics thats ignored",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"skipLibCheck": true,
						"incremental": true,
					}
				}`),
				"/home/src/workspaces/project/a.tlua": `local a = 10;`,
				"/home/src/workspaces/project/b.d.tlua": stringtestutil.Dedent(`
					interface NoName {
						Profiler: new ({ sampleInterval: number, maxBufferSize: number }) => {
							stop: () => Promise<any>;
						};
					}
				`),
			},
			commandLineArgs: []string{""},
			edits: []*tluaEdit{
				noChange,
				{
					caption:         "no change and tlua -b",
					commandLineArgs: []string{"-b", "-v"},
				},
			},
		},
		{
			subScenario:     "Compile incremental with case insensitive file names",
			commandLineArgs: []string{"-p", "."},
			files: FileMap{
				"/home/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"incremental": true
						},
					}`),
				"/home/project/src/index.tlua": stringtestutil.Dedent(`
					local lib1 = require('lib1');
					local lib2 = require('lib2');
					local someLib = require('someLib');
					local otherLib = require('otherlib');
					local foo1: string = lib1.foo1;
					local foo2: string = lib2.foo2;`),
				"/home/node_modules/lib1/index.tlua": "local someLib = require('someLib');\nlocal foo1: string = someLib.foo.foo;\nreturn { foo1 = foo1 };",
				"/home/node_modules/lib1/package.json": stringtestutil.Dedent(`
					{
						"name": "lib1"
					}`),
				"/home/node_modules/lib2/index.tlua": "local someLib = require('someLib');\nlocal foo2: string = someLib.foo.foo;\nreturn { foo2 = foo2 };",
				"/home/node_modules/lib2/package.json": stringtestutil.Dedent(`
					{
						"name": "lib2"
					}
					`),
				"/home/node_modules/someLib/index.tlua": "local other = require('otherlib');\nlocal foo = { foo = other.str };\nreturn { foo = foo };",
				"/home/node_modules/someLib/package.json": stringtestutil.Dedent(`
					{
						"name": "somelib"
					}`),
				"/home/node_modules/otherLib/index.tlua": `local str: string = "s";` + "\n" + `return { str = str };`,
				"/home/node_modules/otherLib/package.json": stringtestutil.Dedent(`
					{
						"name": "otherlib"
					}`),
			},
			cwd:        "/home/project",
			ignoreCase: true,
		},
		{
			subScenario: "internal symbolname in tsbuildInfo",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es2017",
						"strict": true
					}
				}`),
				"/home/src/workspaces/project/a.tlua": stringtestutil.Dedent(`
					local createFileListFromFiles = function(files: File[]): FileList
					local fileList: FileList = {
						length: files.length,
						item: function(index: number): File | null return files[index] || null end,
						[Symbol.iterator]: (0 as any) as () => ArrayIterator<File>,
					} as unknown as FileList;

					return fileList;
					end;
				`),
				getTestLibPathFor("es2015.iterable"): stringtestutil.Dedent(`
					interface SymbolConstructor {
						readonly iterator: unique symbol;
					}
					interface IteratorYieldResult<TYield> {
						done?: false;
						value: TYield;
					}
					interface IteratorReturnResult<TReturn> {
						done: true;
						value: TReturn;
					}
					type IteratorResult<T, TReturn = any> = IteratorYieldResult<T> | IteratorReturnResult<TReturn>;
					interface Iterator<T, TReturn = any, TNext = any> {
						// NOTE: 'next' is defined using a tuple to ensure we report the correct assignability errors in all places.
						next(...[value]: [] | [TNext]): IteratorResult<T, TReturn>;
						return?(value?: TReturn): IteratorResult<T, TReturn>;
						throw?(e?: any): IteratorResult<T, TReturn>;
					}
					interface Iterable<T, TReturn = any, TNext = any> {
						[Symbol.iterator](): Iterator<T, TReturn, TNext>;
					}
					interface IterableIterator<T, TReturn = any, TNext = any> extends Iterator<T, TReturn, TNext> {
						[Symbol.iterator](): IterableIterator<T, TReturn, TNext>;
					}
					interface IteratorObject<T, TReturn = unknown, TNext = unknown> extends Iterator<T, TReturn, TNext> {
						[Symbol.iterator](): IteratorObject<T, TReturn, TNext>;
					}
					type BuiltinIteratorReturn = intrinsic;
					interface ArrayIterator<T> extends IteratorObject<T, BuiltinIteratorReturn, unknown> {
						[Symbol.iterator](): ArrayIterator<T>;
					}
					interface Array<T> {
						[Symbol.iterator](): ArrayIterator<T>;
						entries(): ArrayIterator<[number, T]>;
						keys(): ArrayIterator<number>;
						values(): ArrayIterator<T>;
					}
				`),
				getTestLibPathFor("es2017.full"): stringtestutil.Dedent(`
					/// <reference lib="es2015.iterable"/>
					interface File {
					}
					interface FileList {
						readonly length: number;
						item(index: number): File | null;
						[index: number]: File;
						[Symbol.iterator](): ArrayIterator<File>;
					}
				`) + tluaDefaultLibContent,
			},
			commandLineArgs: []string{""},
			edits: []*tluaEdit{
				noChange,
				{
					caption:         "no change with incremental",
					commandLineArgs: []string{"--incremental"},
				},
				{
					caption:         "no change with incremental that reads buildInfo",
					commandLineArgs: []string{"--incremental"},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "incremental")
	}
}

func TestTscLibraryResolution(t *testing.T) {
	t.Parallel()
	getTscLibraryResolutionFileMap := func(libReplacement bool) FileMap {
		files := FileMap{
			"/home/src/workspace/projects/project1/utils.d.tlua": `declare y: number;`,
			"/home/src/workspace/projects/project1/file.tlua":    `local file = 10;`,
			"/home/src/workspace/projects/project1/core.d.tlua":  `declare core: number;`,
			"/home/src/workspace/projects/project1/index.tlua":   `local x = "type1";`,
			"/home/src/workspace/projects/project1/file2.tlua": stringtestutil.Dedent(`
				/// <reference lib="webworker"/>
				/// <reference lib="es5"/>
			`),
			"/home/src/workspace/projects/project1/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"composite": true,
						"typeRoots": ["./typeroot1"],
						"lib": ["es5", "dom"],
						"traceResolution": true,
						"libReplacement": %t
					}
				}
			`, libReplacement)),
			"/home/src/workspace/projects/project1/typeroot1/sometype/index.d.tlua": `type TheNum = "type1";`,
			"/home/src/workspace/projects/project2/utils.d.tlua":                    `declare y: number;`,
			"/home/src/workspace/projects/project2/index.tlua":                      `local y = 10`,
			"/home/src/workspace/projects/project2/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"composite": true,
						"lib": ["es5", "dom"],
						"traceResolution": true,
						"libReplacement": %t
					}
				}
			`, libReplacement)),
			"/home/src/workspace/projects/project3/utils.d.tlua": `declare y: number;`,
			"/home/src/workspace/projects/project3/index.tlua":   `local z = 10`,
			"/home/src/workspace/projects/project3/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"composite": true,
						"lib": ["es5", "dom"],
						"traceResolution": true,
						"libReplacement": %t
					}
				}
			`, libReplacement)),
			"/home/src/workspace/projects/project4/utils.d.tlua": `declare y: number;`,
			"/home/src/workspace/projects/project4/index.tlua":   `local z = 10`,
			"/home/src/workspace/projects/project4/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"composite": true,
						"lib": ["esnext", "dom", "webworker"],
						"traceResolution": true,
						"libReplacement": %t
					}
				}
			`, libReplacement)),
			getTestLibPathFor("dom"):       "interface DOMInterface { }",
			getTestLibPathFor("webworker"): "interface WebWorkerInterface { }",
			"/home/src/workspace/projects/node_modules/@typescript/unlreated/index.d.tlua": "declare unrelated: number;",
		}
		if libReplacement {
			files["/home/src/workspace/projects/node_modules/@typescript/lib-es5/index.d.tlua"] = tluaDefaultLibContent
			files["/home/src/workspace/projects/node_modules/@typescript/lib-esnext/index.d.tlua"] = tluaDefaultLibContent
			files["/home/src/workspace/projects/node_modules/@typescript/lib-dom/index.d.tlua"] = "interface DOMInterface { }"
			files["/home/src/workspace/projects/node_modules/@typescript/lib-webworker/index.d.tlua"] = "interface WebWorkerInterface { }"
		}
		return files
	}
	getTscLibResolutionTestCases := func(commandLineArgs []string) []*tluaInput {
		return []*tluaInput{
			{
				subScenario:     "with config",
				files:           getTscLibraryResolutionFileMap(false),
				cwd:             "/home/src/workspace/projects",
				commandLineArgs: commandLineArgs,
			},
			{
				subScenario:     "with config with libReplacement",
				files:           getTscLibraryResolutionFileMap(true),
				cwd:             "/home/src/workspace/projects",
				commandLineArgs: commandLineArgs,
			},
		}
	}
	getTscLibraryResolutionUnknown := func() FileMap {
		return FileMap{
			"/home/src/workspace/projects/project1/utils.d.tlua": `declare y: number;`,
			"/home/src/workspace/projects/project1/file.tlua":    `local file = 10;`,
			"/home/src/workspace/projects/project1/core.d.tlua":  `declare core: number;`,
			"/home/src/workspace/projects/project1/index.tlua":   `local x = "type1";`,
			"/home/src/workspace/projects/project1/file2.tlua": stringtestutil.Dedent(`
				/// <reference lib="webworker2"/>
				/// <reference lib="unknownlib"/>
				/// <reference lib="webworker"/>
			`),
			"/home/src/workspace/projects/project1/tluaconfig.json": stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
					"traceResolution": true,
					"libReplacement": true
				}
			}`),
			getTestLibPathFor("webworker"): "interface WebWorkerInterface { }",
		}
	}
	testCases := slices.Concat(
		getTscLibResolutionTestCases([]string{"-b", "project1", "project2", "project3", "project4", "--verbose", "--explainFiles"}),
		getTscLibResolutionTestCases([]string{"-p", "project1", "--explainFiles"}),
		getTscLibResolutionTestCases([]string{"-b", "-w", "project1", "project2", "project3", "project4", "--verbose", "--explainFiles"}),
		[]*tluaInput{
			{
				subScenario:     "unknown lib",
				files:           getTscLibraryResolutionUnknown(),
				cwd:             "/home/src/workspace/projects",
				commandLineArgs: []string{"-p", "project1", "--explainFiles"},
			},
			{
				subScenario: "when noLib toggles",
				files: FileMap{
					"/home/src/workspaces/project/a.d.tlua": `declare a: "hello";`,
					"/home/src/workspaces/project/b.tlua":   `local b = 10;`,
					"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": {
                            "declaration": true,
                            "incremental": true,
                            "lib": ["es6"],
                        },
                    }
                `),
				},
				edits: []*tluaEdit{
					{
						caption:         "with --noLib",
						commandLineArgs: []string{"--noLib"},
					},
				},
			},
		},
	)

	for _, test := range testCases {
		test.run(t, "libraryResolution")
	}
}

func TestTscListFilesOnly(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "loose file",
			files: FileMap{
				"/home/src/workspaces/project/test.tlua": "local x = 1;",
			},
			commandLineArgs: []string{"test.tlua", "--listFilesOnly"},
		},
		{
			subScenario: "combined with incremental",
			files: FileMap{
				"/home/src/workspaces/project/test.tlua":       "local x = 1;",
				"/home/src/workspaces/project/tluaconfig.json": "{}",
			},
			commandLineArgs: []string{"--incremental", "--listFilesOnly"},
			edits: []*tluaEdit{
				{
					caption:         "incremental actual build",
					commandLineArgs: []string{"--incremental"},
				},
				noChange,
				{
					caption:         "incremental should not build",
					commandLineArgs: []string{"--incremental"},
				},
			},
		},
	}

	for _, testCase := range testCases {
		testCase.run(t, "listFilesOnly")
	}
}

func TestTscNoCheck(t *testing.T) {
	t.Parallel()
	type noCheckScenario struct {
		subScenario string
		aText       string
	}
	getTscNoCheckTestCase := func(scenario *noCheckScenario, incremental bool, commandLineArgs []string) *tluaInput {
		noChangeWithCheck := &tluaEdit{
			caption:         "No Change run with checking",
			commandLineArgs: commandLineArgs,
		}
		fixErrorNoCheck := &tluaEdit{
			caption: "Fix `a` error with noCheck",
			edit: func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/a.tlua", `local a = "hello";`)
			},
		}
		addErrorNoCheck := &tluaEdit{
			caption: "Introduce error with noCheck",
			edit: func(sys *TestSys) {
				sys.writeFileNoError("/home/src/workspaces/project/a.tlua", scenario.aText)
			},
		}
		return &tluaInput{
			subScenario: scenario.subScenario + core.IfElse(incremental, " with incremental", ""),
			files: FileMap{
				"/home/src/workspaces/project/a.tlua": scenario.aText,
				"/home/src/workspaces/project/b.tlua": `local b = 10;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"declaration": true,
						"incremental": %t
					}
				}`, incremental)),
			},
			commandLineArgs: slices.Concat(commandLineArgs, []string{"--noCheck"}),
			edits: []*tluaEdit{
				noChange,
				fixErrorNoCheck,   // Fix error with noCheck
				noChange,          // Should be no op
				noChangeWithCheck, // Check errors - should not report any errors - update buildInfo
				noChangeWithCheck, // Should be no op
				noChange,          // Should be no op
				addErrorNoCheck,
				noChange,          // Should be no op
				noChangeWithCheck, // Should check errors and update buildInfo
				fixErrorNoCheck,   // Fix error with noCheck
				noChangeWithCheck, // Should check errors and update buildInfo
				{
					caption: "Add file with error",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/c.tlua", `local c: number = "hello";`)
					},
					commandLineArgs: commandLineArgs,
				},
				addErrorNoCheck,
				fixErrorNoCheck,
				noChangeWithCheck,
				noChange,          // Should be no op
				noChangeWithCheck, // Should be no op
			},
		}
	}

	cases := []noCheckScenario{
		{"syntax errors", `local a = "hello`},
		{"semantic errors", `local a: number = "hello";`},
		{"dts errors", `local a = { private: 10 };`},
	}
	testCases := core.FlatMap(cases, func(c noCheckScenario) []*tluaInput {
		return []*tluaInput{
			getTscNoCheckTestCase(&c, false, []string{}),
			getTscNoCheckTestCase(&c, true, []string{}),
			getTscNoCheckTestCase(&c, false, []string{"-b", "-v"}),
			getTscNoCheckTestCase(&c, true, []string{"-b", "-v"}),
		}
	})
	for _, test := range testCases {
		test.run(t, "noCheck")
	}
}

func TestTscNoEmit(t *testing.T) {
	t.Parallel()
	type tluaNoEmitScenario struct {
		subScenario string
		aText       string
		dtsEnabled  bool
	}
	noEmitScenarios := []*tluaNoEmitScenario{
		{
			subScenario: "syntax errors",
			aText:       `local a = "hello`,
		},
		{
			subScenario: "semantic errors",
			aText:       `local a: number = "hello"`,
		},
		{
			subScenario: "dts errors",
			aText:       `local a = { private: 10 };`,
			dtsEnabled:  true,
		},
		{
			subScenario: "dts errors without dts enabled",
			aText:       `local a = { private: 10 };`,
		},
	}
	getTscNoEmitAndErrorsFileMap := func(scenario *tluaNoEmitScenario, incremental bool, asModules bool, modify func(FileMap)) FileMap {
		files := FileMap{
			"/home/src/projects/project/a.tlua": scenario.aText + core.IfElse(asModules, "\nreturn { a = a };", ""),
			"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"incremental": %t,
						"declaration": %t
					}
				}
		`, incremental, scenario.dtsEnabled)),
		}
		if asModules {
			files["/home/src/projects/project/b.tlua"] = `local b = 10;`
		}
		if modify != nil {
			modify(files)
		}
		return files
	}
	getTscNoEmitAndErrorsTestCasesWorker := func(commandLineArgs []string, addNoEmitOnCommandLine bool, modify func(FileMap), edits func(scenario *tluaNoEmitScenario, commandLineArgs []string, asModules bool) []*tluaEdit) []*tluaInput {
		testingCases := make([]*tluaInput, 0, len(noEmitScenarios)*3)
		commandLineArgsForInput := commandLineArgs
		if addNoEmitOnCommandLine {
			commandLineArgsForInput = slices.Concat(commandLineArgs, []string{"--noEmit"})
		}
		for _, scenario := range noEmitScenarios {
			testingCases = append(
				testingCases,
				&tluaInput{
					subScenario:     scenario.subScenario,
					commandLineArgs: commandLineArgsForInput,
					files:           getTscNoEmitAndErrorsFileMap(scenario, false, false, modify),
					cwd:             "/home/src/projects/project",
					edits:           edits(scenario, commandLineArgs, false),
				},
				&tluaInput{
					subScenario:     scenario.subScenario + " with incremental",
					commandLineArgs: commandLineArgsForInput,
					files:           getTscNoEmitAndErrorsFileMap(scenario, true, false, modify),
					cwd:             "/home/src/projects/project",
					edits:           edits(scenario, commandLineArgs, false),
				},
				&tluaInput{
					subScenario:     scenario.subScenario + " with incremental as modules",
					commandLineArgs: commandLineArgsForInput,
					files:           getTscNoEmitAndErrorsFileMap(scenario, true, true, modify),
					cwd:             "/home/src/projects/project",
					edits:           edits(scenario, commandLineArgs, true),
				},
			)
		}
		return testingCases
	}
	getTscNoEmitAndErrorsTestCases := func(commandLineArgs []string) []*tluaInput {
		return getTscNoEmitAndErrorsTestCasesWorker(
			commandLineArgs,
			true,
			nil,
			func(scenario *tluaNoEmitScenario, commandLineArgs []string, asModules bool) []*tluaEdit {
				fixedATsContent := `local a = "hello";` + core.IfElse(asModules, "\nreturn { a = a };", "")
				return []*tluaEdit{
					noChange,
					{
						caption: "Fix error",
						edit: func(sys *TestSys) {
							sys.writeFileNoError("/home/src/projects/project/a.tlua", fixedATsContent)
						},
					},
					noChange,
					{
						caption:         "Emit after fixing error",
						commandLineArgs: commandLineArgs,
					},
					noChange,
					{
						caption: "Introduce error",
						edit: func(sys *TestSys) {
							sys.writeFileNoError("/home/src/projects/project/a.tlua", scenario.aText)
						},
					},
					{
						caption:         "Emit when error",
						commandLineArgs: commandLineArgs,
					},
					noChange,
				}
			},
		)
	}
	getTscNoEmitAndErrorsWatchTestCases := func(commandLineArgs []string) []*tluaInput {
		return getTscNoEmitAndErrorsTestCasesWorker(
			commandLineArgs,
			false,
			func(files FileMap) {
				files["/home/src/projects/project/tluaconfig.json"] = strings.Replace(files["/home/src/projects/project/tluaconfig.json"].(string), "}", `, "noEmit": true }`, 1)
			},
			func(scenario *tluaNoEmitScenario, commandLineArgs []string, asModules bool) []*tluaEdit {
				fixedATsContent := `local a = "hello";` + core.IfElse(asModules, "\nreturn { a = a };", "")
				return []*tluaEdit{
					{
						caption: "Fix error",
						edit: func(sys *TestSys) {
							sys.writeFileNoError("/home/src/projects/project/a.tlua", fixedATsContent)
						},
					},
					{
						caption: "Emit after fixing error",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/projects/project/tluaconfig.json", `"noEmit": true`, `"noEmit": false`)
						},
					},
					{
						caption: "no Emit run after fixing error",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/projects/project/tluaconfig.json", `"noEmit": false`, `"noEmit": true`)
						},
					},
					{
						caption: "Introduce error",
						edit: func(sys *TestSys) {
							sys.writeFileNoError("/home/src/projects/project/a.tlua", scenario.aText)
						},
					},
					{
						caption: "Emit when error",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/projects/project/tluaconfig.json", `"noEmit": true`, `"noEmit": false`)
						},
					},
					{
						caption: "no Emit run when error",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/projects/project/tluaconfig.json", `"noEmit": false`, `"noEmit": true`)
						},
					},
				}
			},
		)
	}
	getTscNoEmitChangesFileMap := func(optionsStr string) FileMap {
		return FileMap{
			"/home/src/workspaces/project/src/class.tlua": stringtestutil.Dedent(`
				local classC = {
					prop = 1,
				};
				return { classC = classC };`),
			"/home/src/workspaces/project/src/indirectClass.tlua": stringtestutil.Dedent(`
				local classMod = require('src.class');
				local indirectClass = {
					classC = classMod.classC,
				};
				return { indirectClass = indirectClass };`),
			"/home/src/workspaces/project/src/directUse.tlua": stringtestutil.Dedent(`
				local indirect = require('src.indirectClass');
				indirect.indirectClass.classC.prop;`),
			"/home/src/workspaces/project/src/indirectUse.tlua": stringtestutil.Dedent(`
				local indirect = require('src.indirectClass');
				indirect.indirectClass.classC.prop;`),
			"/home/src/workspaces/project/src/noChangeFile.tlua": stringtestutil.Dedent(`
				function writeLog(s: string) end`),
			"/home/src/workspaces/project/src/noChangeFileWithEmitSpecificError.tlua": stringtestutil.Dedent(`
				function someFunc(arguments: boolean, ...rest: any[]) end`),
			"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions":  { %s }
				}`, optionsStr)),
		}
	}

	type tluaNoEmitChangesScenario struct {
		subScenario   string
		optionsString string
	}
	noEmitChangesScenarios := []*tluaNoEmitChangesScenario{
		{
			// !!! sheetal missing initial reporting of Duplicate_identifier_arguments_Compiler_uses_arguments_to_initialize_rest_parameters is absent
			subScenario:   "composite",
			optionsString: `"composite": true`,
		},
		{
			subScenario:   "incremental declaration",
			optionsString: `"incremental": true, "declaration": true`,
		},
		{
			subScenario:   "incremental",
			optionsString: `"incremental": true`,
		},
	}
	getTscNoEmitChangesTestCases := func(commandLineArgs []string) []*tluaInput {
		noChangeWithNoEmit := &tluaEdit{
			caption:         "No Change run with noEmit",
			commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit"}),
		}
		noChangeWithEmit := &tluaEdit{
			caption:         "No Change run with emit",
			commandLineArgs: commandLineArgs,
		}
		// While the introduced error is live, incremental runs miss the
		// dependent-file errors that the clean build reports (see
		// signaturePropagationDiff), so the runs until the fix carry the
		// explanation.
		noChangeWithNoEmitWhileError := &tluaEdit{
			caption:         "No Change run with noEmit",
			commandLineArgs: noChangeWithNoEmit.commandLineArgs,
			expectedDiff:    signaturePropagationDiff,
		}
		noChangeWithEmitWhileError := &tluaEdit{
			caption:         "No Change run with emit",
			commandLineArgs: commandLineArgs,
			expectedDiff:    signaturePropagationDiff,
		}
		introduceError := func(sys *TestSys) {
			sys.replaceFileText("/home/src/workspaces/project/src/class.tlua", "prop", "prop1")
		}
		fixError := func(sys *TestSys) {
			sys.replaceFileText("/home/src/workspaces/project/src/class.tlua", "prop1", "prop")
		}
		testCases := make([]*tluaInput, 0, len(noEmitChangesScenarios))
		for _, scenario := range noEmitChangesScenarios {
			testCases = append(
				testCases,
				&tluaInput{
					subScenario:     "changes " + scenario.subScenario,
					commandLineArgs: commandLineArgs,
					files:           getTscNoEmitChangesFileMap(scenario.optionsString),
					edits: []*tluaEdit{
						noChangeWithNoEmit,
						noChangeWithNoEmit,
						{
							caption:         "Introduce error but still noEmit",
							commandLineArgs: noChangeWithNoEmit.commandLineArgs,
							edit:            introduceError,
							expectedDiff:    signaturePropagationDiff,
						},
						{
							caption: "Fix error and emit",
							edit:    fixError,
						},
						noChangeWithEmit,
						noChangeWithNoEmit,
						noChangeWithNoEmit,
						noChangeWithEmit,
						{
							caption:      "Introduce error and emit",
							edit:         introduceError,
							expectedDiff: signaturePropagationDiff,
						},
						noChangeWithEmitWhileError,
						noChangeWithNoEmitWhileError,
						noChangeWithNoEmitWhileError,
						noChangeWithEmitWhileError,
						{
							caption:         "Fix error and no emit",
							commandLineArgs: noChangeWithNoEmit.commandLineArgs,
							edit:            fixError,
						},
						noChangeWithEmit,
						noChangeWithNoEmit,
						noChangeWithNoEmit,
						noChangeWithEmit,
					},
				},
				&tluaInput{
					subScenario:     "changes with initial noEmit " + scenario.subScenario,
					commandLineArgs: noChangeWithNoEmit.commandLineArgs,
					files:           getTscNoEmitChangesFileMap(scenario.optionsString),
					edits: []*tluaEdit{
						noChangeWithEmit,
						{
							caption:         "Introduce error with emit",
							commandLineArgs: commandLineArgs,
							edit:            introduceError,
							expectedDiff:    signaturePropagationDiff,
						},
						{
							caption: "Fix error and no emit",
							edit:    fixError,
						},
						noChangeWithEmit,
					},
				},
			)
		}
		return testCases
	}
	getTscNoEmitDtsChangesFileMap := func(incremental bool, asModules bool) FileMap {
		files := FileMap{
			"/home/src/projects/project/a.tlua": `local a = { private: 10 };` + core.IfElse(asModules, "\nreturn { a = a };", ""),
			"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
				{
					"compilerOptions": {
						"incremental": %t,
					}
				}
		`, incremental)),
		}
		if asModules {
			files["/home/src/projects/project/b.tlua"] = `local b = 10;`
		}
		return files
	}
	getTscNoEmitDtsChangesEdits := func(commandLineArgs []string) []*tluaEdit {
		return []*tluaEdit{
			noChange,
			{
				caption:         "With declaration enabled noEmit - Should report errors",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit", "--declaration"}),
			},
			{
				caption:         "With declaration and declarationMap noEmit - Should report errors",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit", "--declaration", "--declarationMap"}),
			},
			noChange,
			{
				caption:         "Dts Emit with error",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--declaration"}),
			},
			{
				caption: "Fix the error",
				edit: func(sys *TestSys) {
					sys.replaceFileText("/home/src/projects/project/a.tlua", "private", "public")
				},
			},
			{
				caption:         "With declaration enabled noEmit",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit", "--declaration"}),
			},
			{
				caption:         "With declaration and declarationMap noEmit",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit", "--declaration", "--declarationMap"}),
			},
		}
	}
	getTscNoEmitDtsChangesTestCases := func() []*tluaInput {
		return []*tluaInput{
			{
				subScenario:     "dts errors with declaration enable changes",
				commandLineArgs: []string{"-b", "-v", "--noEmit"},
				files:           getTscNoEmitDtsChangesFileMap(false, false),
				cwd:             "/home/src/projects/project",
				edits:           getTscNoEmitDtsChangesEdits([]string{"-b", "-v"}),
			},
			{
				subScenario:     "dts errors with declaration enable changes with incremental",
				commandLineArgs: []string{"-b", "-v", "--noEmit"},
				files:           getTscNoEmitDtsChangesFileMap(true, false),
				cwd:             "/home/src/projects/project",
				edits:           getTscNoEmitDtsChangesEdits([]string{"-b", "-v"}),
			},
			{
				subScenario:     "dts errors with declaration enable changes with incremental as modules",
				commandLineArgs: []string{"-b", "-v", "--noEmit"},
				files:           getTscNoEmitDtsChangesFileMap(true, true),
				cwd:             "/home/src/projects/project",
				edits:           getTscNoEmitDtsChangesEdits([]string{"-b", "-v"}),
			},
		}
	}
	getTscNoEmitDtsChangesMultiFileErrorsTestCases := func(commandLineArgs []string) []*tluaInput {
		aContent := `local a = { private: 10 };`
		return []*tluaInput{
			{
				subScenario:     "dts errors with declaration enable changes with multiple files",
				commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit"}),
				files: FileMap{
					"/home/src/projects/project/a.tlua": aContent,
					"/home/src/projects/project/b.tlua": `local b = 10;`,
					"/home/src/projects/project/c.tlua": strings.Replace(aContent, "a", "c", 1),
					"/home/src/projects/project/d.tlua": strings.Replace(aContent, "a", "d", 1),
					"/home/src/projects/project/tluaconfig.json": stringtestutil.Dedent(`
						{
							"compilerOptions": {
								"incremental": true,
							}
						}
				`),
				},
				cwd: "/home/src/projects/project",
				edits: slices.Concat(
					getTscNoEmitDtsChangesEdits(commandLineArgs),
					[]*tluaEdit{
						{
							caption: "Fix the another ",
							edit: func(sys *TestSys) {
								sys.replaceFileText("/home/src/projects/project/c.tlua", "private", "public")
							},
							commandLineArgs: slices.Concat(commandLineArgs, []string{"--noEmit", "--declaration", "--declarationMap"}),
						},
					},
				),
			},
		}
	}
	getTscNoEmitLoopTestCase := func(suffix string, commandLineArgs []string) *tluaInput {
		return &tluaInput{
			subScenario: "does not go in loop when watching when no files are emitted" + suffix,
			files: FileMap{
				"/user/username/projects/myproject/a.tlua": "",
				"/user/username/projects/myproject/b.tlua": "",
				"/user/username/projects/myproject/tluaconfig.json": stringtestutil.Dedent(`
					{
                        "compilerOptions": {
                            "noEmit": true,
                        },
                    }`),
			},
			cwd:             "/user/username/projects/myproject",
			commandLineArgs: commandLineArgs,
			edits: []*tluaEdit{
				{
					caption: "No change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/user/username/projects/myproject/a.tlua`, sys.readFileNoError(`/user/username/projects/myproject/a.tlua`))
					},
				},
				{
					caption: "change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/user/username/projects/myproject/a.tlua`, "local x = 10;")
					},
				},
			},
		}
	}
	testCases := slices.Concat(
		[]*tluaInput{
			{
				subScenario: "when project has strict true",
				files: FileMap{
					"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
						{
							"compilerOptions": {
								"incremental": true,
								"strict": true
							}
						}`),
					"/home/src/workspaces/project/class1.tlua": `local class1 = {};`,
				},
				commandLineArgs: []string{"--noEmit"},
				edits:           noChangeOnlyEdit,
			},
			getTscNoEmitLoopTestCase("", []string{"-b", "-w", "-verbose"}),
			getTscNoEmitLoopTestCase(" with incremental", []string{"-b", "-w", "-verbose", "--incremental"}),
		},
		getTscNoEmitAndErrorsTestCases([]string{}),
		getTscNoEmitAndErrorsTestCases([]string{"-b", "-v"}),
		getTscNoEmitChangesTestCases([]string{}),
		getTscNoEmitChangesTestCases([]string{"-b", "-v"}),
		getTscNoEmitDtsChangesTestCases(),
		getTscNoEmitDtsChangesMultiFileErrorsTestCases([]string{}),
		getTscNoEmitDtsChangesMultiFileErrorsTestCases([]string{"-b", "-v"}),
		getTscNoEmitAndErrorsWatchTestCases([]string{"-b", "-verbose", "-w"}),
	)

	for _, test := range testCases {
		test.run(t, "noEmit")
	}
}

func TestTscNoEmitOnError(t *testing.T) {
	t.Parallel()
	type tluaNoEmitOnErrorScenario struct {
		subScenario       string
		mainErrorContent  string
		fixedErrorContent string
	}
	getTscNoEmitOnErrorFileMap := func(scenario *tluaNoEmitOnErrorScenario, declaration bool, incremental bool) FileMap {
		return FileMap{
			"/user/username/projects/noEmitOnError/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
			{
				"compilerOptions": {
					"outDir": "./dev-build",
					"declaration": %t,
					"incremental": %t,
					"noEmitOnError": true,
				},
			}`, declaration, incremental)),
			"/user/username/projects/noEmitOnError/shared/types/db.tlua": stringtestutil.Dedent(`
				interface A {
					name: string;
				}
			`),
			"/user/username/projects/noEmitOnError/src/main.tlua": scenario.mainErrorContent,
			"/user/username/projects/noEmitOnError/src/other.tlua": stringtestutil.Dedent(`
				console.log("hi");
			`),
		}
	}
	getTscNoEmitOnErrorTestCases := func(scenarios []*tluaNoEmitOnErrorScenario, commandLineArgs []string) []*tluaInput {
		testCases := make([]*tluaInput, 0, len(scenarios)*4)
		for _, scenario := range scenarios {
			edits := []*tluaEdit{
				noChange,
				{
					caption: "Fix error",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/noEmitOnError/src/main.tlua", scenario.fixedErrorContent)
					},
				},
				noChange,
			}
			testCases = append(
				testCases,
				&tluaInput{
					subScenario:     scenario.subScenario,
					files:           getTscNoEmitOnErrorFileMap(scenario, false, false),
					cwd:             "/user/username/projects/noEmitOnError",
					commandLineArgs: commandLineArgs,
					edits:           edits,
				},
				&tluaInput{
					subScenario:     scenario.subScenario + " with declaration",
					files:           getTscNoEmitOnErrorFileMap(scenario, true, false),
					cwd:             "/user/username/projects/noEmitOnError",
					commandLineArgs: commandLineArgs,
					edits:           edits,
				},
				&tluaInput{
					subScenario:     scenario.subScenario + " with incremental",
					files:           getTscNoEmitOnErrorFileMap(scenario, false, true),
					cwd:             "/user/username/projects/noEmitOnError",
					commandLineArgs: commandLineArgs,
					edits:           edits,
				},
				&tluaInput{
					subScenario:     scenario.subScenario + " with declaration with incremental",
					files:           getTscNoEmitOnErrorFileMap(scenario, true, true),
					cwd:             "/user/username/projects/noEmitOnError",
					commandLineArgs: commandLineArgs,
					edits:           edits,
				},
			)
		}
		return testCases
	}
	getTscWatchNoEmitOnErrorTestCases := func(scenarios []*tluaNoEmitOnErrorScenario, commandLineArgs []string) []*tluaInput {
		var edits []*tluaEdit
		for _, scenario := range scenarios {
			if edits != nil {
				edits = append(edits, &tluaEdit{
					caption: scenario.subScenario,
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/user/username/projects/noEmitOnError/src/main.tlua`, scenario.mainErrorContent)
					},
				})
			}
			edits = append(
				edits,
				&tluaEdit{
					caption: "No Change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/user/username/projects/noEmitOnError/src/main.tlua`, sys.readFileNoError(`/user/username/projects/noEmitOnError/src/main.tlua`))
					},
				},
				&tluaEdit{
					caption: "Fix " + scenario.subScenario,
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/noEmitOnError/src/main.tlua", scenario.fixedErrorContent)
					},
				},
				&tluaEdit{
					caption: "No Change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(`/user/username/projects/noEmitOnError/src/main.tlua`, sys.readFileNoError(`/user/username/projects/noEmitOnError/src/main.tlua`))
					},
				},
			)
		}
		return []*tluaInput{
			{
				subScenario:     "noEmitOnError",
				files:           getTscNoEmitOnErrorFileMap(scenarios[0], false, false),
				cwd:             "/user/username/projects/noEmitOnError",
				commandLineArgs: commandLineArgs,
				edits:           edits,
			},
			{
				subScenario:     "noEmitOnError with declaration",
				files:           getTscNoEmitOnErrorFileMap(scenarios[0], true, false),
				cwd:             "/user/username/projects/noEmitOnError",
				commandLineArgs: commandLineArgs,
				edits:           edits,
			},
			{
				subScenario:     "noEmitOnError with incremental",
				files:           getTscNoEmitOnErrorFileMap(scenarios[0], false, true),
				cwd:             "/user/username/projects/noEmitOnError",
				commandLineArgs: commandLineArgs,
				edits:           edits,
			},
			{
				subScenario:     "noEmitOnError with declaration with incremental",
				files:           getTscNoEmitOnErrorFileMap(scenarios[0], true, true),
				cwd:             "/user/username/projects/noEmitOnError",
				commandLineArgs: commandLineArgs,
				edits:           edits,
			},
		}
	}
	scenarios := []*tluaNoEmitOnErrorScenario{
		{
			subScenario: "syntax errors",
			mainErrorContent: stringtestutil.Dedent(`
                local a = {
                    lastName: 'sdsd'
                ;
            `),
			fixedErrorContent: stringtestutil.Dedent(`
                local a = {
                    lastName: 'sdsd'
                };`),
		},
		{
			subScenario: "semantic errors",
			mainErrorContent: stringtestutil.Dedent(`
                local a: string = 10;`),
			fixedErrorContent: stringtestutil.Dedent(`
                local a: string = "hello";`),
		},
		{
			subScenario: "dts errors",
			mainErrorContent: stringtestutil.Dedent(`
                local a = { private: 10 };
            `),
			fixedErrorContent: stringtestutil.Dedent(`
                local a = { p: 10 };
            `),
		},
	}
	testCases := slices.Concat(
		getTscNoEmitOnErrorTestCases(scenarios, []string{}),
		getTscNoEmitOnErrorTestCases(scenarios, []string{"-b", "-v"}),
		[]*tluaInput{
			{
				subScenario: `when declarationMap changes`,
				files: FileMap{
					"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
						{
							"compilerOptions": {
								"noEmitOnError": true,
								"declaration": true,
								"composite": true,
							},
						}`),
					"/home/src/workspaces/project/a.tlua": "local x = 10;",
					"/home/src/workspaces/project/b.tlua": "local y = 10;",
				},
				edits: []*tluaEdit{
					{
						caption: "error and enable declarationMap",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/workspaces/project/a.tlua", "x", "x: 20")
						},
						commandLineArgs: []string{"--declarationMap"},
					},
					{
						caption: "fix error declarationMap",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/workspaces/project/a.tlua", "x: 20", "x")
						},
						commandLineArgs: []string{"--declarationMap"},
					},
				},
			},
			{
				subScenario: "file deleted before fixing error with noEmitOnError",
				files: FileMap{
					"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
						{
							"compilerOptions": {
								"outDir": "outDir",
								"noEmitOnError": true,
							},
						}`),
					"/home/src/workspaces/project/file1.tlua": `local x: 30 = "hello";`,
					"/home/src/workspaces/project/file2.tlua": `local D = { };`,
				},
				commandLineArgs: []string{"-i"},
				edits: []*tluaEdit{
					{
						caption: "delete file without error",
						edit: func(sys *TestSys) {
							sys.removeNoError("/home/src/workspaces/project/file2.tlua")
						},
					},
				},
			},
		},
		getTscWatchNoEmitOnErrorTestCases(scenarios, []string{"-b", "-w", "-v"}),
	)

	for _, test := range testCases {
		test.run(t, "noEmitOnError")
	}
}

func TestTscProjectReferences(t *testing.T) {
	t.Parallel()
	cases := []tluaInput{
		{
			subScenario: "when project references composite project with noEmit",
			files: FileMap{
				"/home/src/workspaces/solution/utils/index.tlua": "local x = 10;\nreturn { x = x };",
				"/home/src/workspaces/solution/utils/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"noEmit": true
					}
				}`),
				"/home/src/workspaces/solution/project/index.tlua": "local utils = require(\"utils.index\");\nutils.x;",
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"rootDir": "..",
					},
					"references": [
						{ "path": "../utils" },
					],
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--p", "project"},
		},
		{
			subScenario: "when project references composite",
			files: FileMap{
				"/home/src/workspaces/solution/utils/index.tlua":   "local x = 10;\nreturn { x = x };",
				"/home/src/workspaces/solution/utils/index.d.tlua": "declare x: 10;",
				"/home/src/workspaces/solution/utils/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/home/src/workspaces/solution/project/index.tlua": "local utils = require(\"utils.index\");\nutils.x;",
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"rootDir": "..",
					},
					"references": [
						{ "path": "../utils" },
					],
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--p", "project"},
		},
		{
			subScenario: "when project reference is not built",
			files: FileMap{
				"/home/src/workspaces/solution/utils/index.tlua": "local x = 10;\nreturn { x = x };",
				"/home/src/workspaces/solution/utils/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/home/src/workspaces/solution/project/index.tlua": "local utils = require(\"utils.index\");\nutils.x;",
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"rootDir": "..",
					},
					"references": [
						{ "path": "../utils" },
					],
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--p", "project"},
		},
		{
			subScenario: "when project contains invalid project reference",
			files: FileMap{
				"/home/src/workspaces/solution/project/index.tlua": `local x = 10;`,
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"references": [
						{ "path": "../utils" },
					],
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--p", "project"},
		},
		{
			subScenario: "when project references have invalid fields",
			files: FileMap{
				"/home/src/workspaces/solution/project/index.tlua": `local x = 10;`,
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"noEmit": true
					},
					"files": ["index.tlua"],
					"references": [
						{ "path": true },
						{ "circular": true },
						{ "path": "../utils", "circular": "yes" },
						{ "path": "" },
						{ "path": "../valid", "circular": true }
					]
				}`),
				"/home/src/workspaces/solution/utils/index.tlua":   "local y = 10;",
				"/home/src/workspaces/solution/utils/index.d.tlua": "declare y: number;",
				"/home/src/workspaces/solution/utils/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/home/src/workspaces/solution/valid/index.tlua":   "local z = 10;",
				"/home/src/workspaces/solution/valid/index.d.tlua": "declare z: number;",
				"/home/src/workspaces/solution/valid/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--p", "project"},
		},
		{
			subScenario: "default setup was created correctly",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					}
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local primaryA = 0;",
				"/home/src/workspaces/project/secondary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"rootDir": "..",
					},
					"references": [{
						"path": "../primary"
					}]
				}`),
				"/home/src/workspaces/project/secondary/b.tlua": `local mod_1 = require("primary.a");`,
			},
			commandLineArgs: []string{"--p", "primary/tluaconfig.json"},
		},
		{
			subScenario: "errors when declaration = false",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"declaration": false
					}
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local primaryA = 0;",
			},
			commandLineArgs: []string{"--p", "primary/tluaconfig.json"},
		},
		{
			subScenario: "errors when the referenced project doesnt have composite",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": false,
						"outDir": "bin",
					}
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local primaryA = 0;",
				"/home/src/workspaces/project/reference/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"rootDir": "..",
					},
					"files": [ "b.tlua" ],
					"references": [ { "path": "../primary" } ]
				}`),
				"/home/src/workspaces/project/reference/b.tlua": `local mod_1 = require("primary.a");`,
			},
			commandLineArgs: []string{"--p", "reference/tluaconfig.json"},
		},
		{
			subScenario: "does not error when the referenced project doesnt have composite if its a container project",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": false,
						"outDir": "bin",
					}
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local primaryA = 0;",
				"/home/src/workspaces/project/reference/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"rootDir": "..",
					},
					"files": [ ],
					"references": [{
						"path": "../primary"
					}]
				}`),
				"/home/src/workspaces/project/reference/b.tlua": `local mod_1 = require("primary.a");`,
			},
			commandLineArgs: []string{"--p", "reference/tluaconfig.json"},
		},
		{
			subScenario: "errors when the file list is not exhaustive",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					},
					"files": [ "a.tlua" ]
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local b = require('b');",
				"/home/src/workspaces/project/primary/b.tlua": "local primaryB = 0;",
			},
			commandLineArgs: []string{"--p", "primary/tluaconfig.json"},
		},
		{
			subScenario: "errors when the referenced project doesnt exist",
			files: FileMap{
				"/home/src/workspaces/project/primary/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					},
					"references": [{
						"path": "../foo"
					}]
				}`),
				"/home/src/workspaces/project/primary/a.tlua": "local primaryA = 0;",
			},
			commandLineArgs: []string{"--p", "primary/tluaconfig.json"},
		},
		{
			subScenario: "redirects to the output dts file",
			files: FileMap{
				"/home/src/workspaces/project/alpha/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					}
				}`),
				"/home/src/workspaces/project/alpha/a.tlua":       "local m: number = 3;\nreturn { m = m };",
				"/home/src/workspaces/project/alpha/bin/a.d.tlua": "declare alphaBuilt: number;",
				"/home/src/workspaces/project/beta/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"rootDir": "..",
					},
					"references": [ { "path": "../alpha" } ]
				}`),
				"/home/src/workspaces/project/beta/b.tlua": "local a = require('alpha.a');",
			},
			commandLineArgs: []string{"--p", "beta/tluaconfig.json", "--explainFiles"},
		},
		{
			subScenario: "issues a nice error when the input file is missing",
			files: FileMap{
				"/home/src/workspaces/project/alpha/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					},
					"references": []
				}`),
				"/home/src/workspaces/project/alpha/a.tlua": "local m: number = 3;",
				"/home/src/workspaces/project/beta/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
						"rootDir": "..",
					},
					"references": [ { "path": "../alpha" } ]
				}`),
				"/home/src/workspaces/project/beta/b.tlua": "local a = require('alpha.a');",
			},
			commandLineArgs: []string{"--p", "beta/tluaconfig.json"},
		},
		{
			subScenario: "doesnt infer the rootDir from source paths",
			files: FileMap{
				"/home/src/workspaces/project/alpha/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					},
					"references": []
				}`),
				"/home/src/workspaces/project/alpha/src/a.tlua": "local m: number = 3;",
			},
			commandLineArgs: []string{"--p", "alpha/tluaconfig.json"},
		},
		{
			// !!! sheetal rootDir error not reported
			subScenario: "errors when a file is outside the rootdir",
			files: FileMap{
				"/home/src/workspaces/project/alpha/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "bin",
					},
					"references": []
				}`),
				"/home/src/workspaces/project/alpha/src/a.tlua": "/// <reference path=\"../../beta/b.tlua\"/>\nlocal m = 3;",
				"/home/src/workspaces/project/beta/b.tlua":      "local beta = 0;",
			},
			commandLineArgs: []string{"--p", "alpha/tluaconfig.json"},
		},
	}

	for _, c := range cases {
		c.run(t, "projectReferences")
	}
}

func TestTypeAcquisition(t *testing.T) {
	t.Parallel()
	(&tluaInput{
		subScenario: "parse tsconfig with typeAcquisition",
		files: FileMap{
			"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
					"noEmit": true,
				},
				"typeAcquisition": {
					"enable": true,
					"include": ["0.d.tlua", "1.d.tlua"],
					"exclude": ["0.lua", "1.lua"],
					"disableFilenameBasedTypeAcquisition": true,
				},
			}`),
		},
		commandLineArgs: []string{},
	}).run(t, "typeAcquisition")
}

func TestGenerateTrace(t *testing.T) {
	t.Parallel()
	cases := []*tluaInput{
		{
			subScenario: "generateTrace generates types file",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
				"/home/src/workspaces/project/a.tlua": stringtestutil.Dedent(`
				interface Person {
					name: string;
					age: number;
				}
				local p: Person = { name: "Alice", age: 30 };
				`),
			},
			commandLineArgs: []string{"--generateTrace", "/home/src/workspaces/project/trace", "--singleThreaded"},
		},
		{
			subScenario: "generateTrace with multiple files and complex types",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"noEmit": true
					}
				}`),
				"/home/src/workspaces/project/types.tlua": stringtestutil.Dedent(`
				interface Container<T> {
					value: T;
					map<U>(fn: (x: T) => U): Container<U>;
				}
				type Nullable<T> = T | null | undefined;
				`),
				"/home/src/workspaces/project/main.tlua": stringtestutil.Dedent(`
				local c: Container<number> = { value: 42, map: function(fn) return ({ value: fn(42), map: c.map }) end };
				local n: Nullable<string> = "hello";
				`),
			},
			commandLineArgs: []string{"--generateTrace", "/home/src/workspaces/project/trace", "--singleThreaded"},
		},
	}

	for _, c := range cases {
		c.run(t, "generateTrace")
	}
}

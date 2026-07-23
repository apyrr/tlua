package tluatests

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/testutil/harnessutil"
	"github.com/apyrr/tlua/internal/testutil/stringtestutil"
	"github.com/apyrr/tlua/internal/tsoptions"
	"gotest.tools/v3/assert"
)

func TestBuildCommandLine(t *testing.T) {
	t.Parallel()
	getBuildCommandLineDifferentOptionsMap := func(optionName string) FileMap {
		return FileMap{
			"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
			{
				"compilerOptions": {
					"%s": true
				}
			}`, optionName)),
			"/home/src/workspaces/project/a.tlua": `local a = 10;local aLocal = 10;return { a = a };`,
			"/home/src/workspaces/project/b.tlua": `local b = 10;local bLocal = 10;return { b = b };`,
			"/home/src/workspaces/project/c.tlua": `local a = require("a");local c = a.a;return { c = c };`,
			"/home/src/workspaces/project/d.tlua": `local b = require("b");local d = b.b;return { d = d };`,
		}
	}
	testCases := slices.Concat(
		[]*tluaInput{
			{
				subScenario:     "help",
				files:           FileMap{},
				commandLineArgs: []string{"--build", "--help"},
			},
			{
				subScenario:     "locale",
				files:           FileMap{},
				commandLineArgs: []string{"--build", "--help", "--locale", "en"},
			},
			{
				subScenario:     "bad locale",
				files:           FileMap{},
				commandLineArgs: []string{"--build", "--help", "--locale", "whoops"},
			},
			{
				subScenario:     "different options",
				files:           getBuildCommandLineDifferentOptionsMap("composite"),
				commandLineArgs: []string{"--build", "--verbose"},
				edits: []*tluaEdit{
					{
						caption:         "with sourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--sourceMap"},
						expectedDiff:    declarationEmitErrorReplayDiff,
					},
					{
						caption:      "should re-emit only js so they dont contain sourcemap",
						expectedDiff: declarationEmitErrorReplayDiff,
					},
					{
						caption:         "with declaration should not emit anything",
						commandLineArgs: []string{"--build", "--verbose", "--declaration"},
					},
					noChange,
					{
						caption:         "with declaration and declarationMap",
						commandLineArgs: []string{"--build", "--verbose", "--declaration", "--declarationMap"},
					},
					{
						caption: "should re-emit only dts so they dont contain sourcemap",
					},
					{
						caption:         "with emitDeclarationOnly should not emit anything",
						commandLineArgs: []string{"--build", "--verbose", "--emitDeclarationOnly"},
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
						commandLineArgs: []string{"--build", "--verbose", "--declaration"},
					},
					{
						caption:         "with inlineSourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--inlineSourceMap"},
						expectedDiff:    declarationEmitErrorReplayDiff,
					},
					{
						caption:         "with sourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--sourceMap"},
						expectedDiff:    declarationEmitErrorReplayDiff,
					},
				},
			},
			{
				subScenario:     "different options with incremental",
				files:           getBuildCommandLineDifferentOptionsMap("incremental"),
				commandLineArgs: []string{"--build", "--verbose"},
				edits: []*tluaEdit{
					{
						caption:         "with sourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--sourceMap"},
					},
					{
						caption: "should re-emit only js so they dont contain sourcemap",
					},
					{
						caption:         "with declaration, emit Dts and should not emit js",
						commandLineArgs: []string{"--build", "--verbose", "--declaration"},
					},
					{
						caption:         "with declaration and declarationMap",
						commandLineArgs: []string{"--build", "--verbose", "--declaration", "--declarationMap"},
					},
					noChange,
					{
						caption: "local change",
						edit: func(sys *TestSys) {
							sys.replaceFileText("/home/src/workspaces/project/a.tlua", "Local = 1", "Local = 10")
						},
						expectedDiff: declarationEmitErrorReplayDiff,
					},
					{
						caption:         "with declaration and declarationMap",
						commandLineArgs: []string{"--build", "--verbose", "--declaration", "--declarationMap"},
					},
					noChange,
					{
						caption:         "with inlineSourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--inlineSourceMap"},
					},
					{
						caption:         "with sourceMap",
						commandLineArgs: []string{"--build", "--verbose", "--sourceMap"},
					},
					{
						caption: "emit js files",
					},
					{
						caption:         "with declaration and declarationMap",
						commandLineArgs: []string{"--build", "--verbose", "--declaration", "--declarationMap"},
					},
					{
						caption:         "with declaration and declarationMap, should not re-emit",
						commandLineArgs: []string{"--build", "--verbose", "--declaration", "--declarationMap"},
					},
				},
			},
		},
	)

	for _, test := range testCases {
		test.run(t, "commandLine")
	}
}

func TestBuildClean(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "file name and output name clashing",
			files: FileMap{
				"/home/src/workspaces/solution/index.lua": "",
				"/home/src/workspaces/solution/bar.tlua":  "",
				"/home/src/workspaces/solution/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {}
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "--clean"},
		},
		{
			subScenario: "tsx with dts emit",
			files: FileMap{
				"/home/src/workspaces/solution/project/src/main.tsx": "local x = 10;",
				"/home/src/workspaces/solution/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": { "declaration": true },
					"include": ["src/**/*.tsx", "src/**/*.tlua"]
				}`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "project", "-v", "--explainFiles"},
			edits: []*tluaEdit{
				noChange,
				{
					caption:         "clean build",
					commandLineArgs: []string{"-b", "project", "--clean"},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "clean")
	}
}

func TestBuildConfigFileErrors(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "when tsconfig extends the missing file",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.first.json": stringtestutil.Dedent(`
					{
						"extends": "./foobar.json",
						"compilerOptions": {
							"composite": true
						}
					}`),
				"/home/src/workspaces/project/tluaconfig.second.json": stringtestutil.Dedent(`
					{
						"extends": "./foobar.json",
						"compilerOptions": {
							"composite": true
						}
					}`),
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true
						},
						"references": [
							{ "path": "./tluaconfig.first.json" },
							{ "path": "./tluaconfig.second.json" }
						]
					}`),
			},
			commandLineArgs: []string{"--b"},
		},
		{
			subScenario: "reports invalid project reference fields",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true
						},
						"files": ["index.tlua"],
						"references": [
							{ "path": true },
							{ "circular": true },
							{ "path": "./utils", "circular": "yes" },
							{ "path": "" },
							{ "path": "./valid", "circular": true }
						]
					}`),
				"/home/src/workspaces/project/index.tlua": "local x = 10;",
				"/home/src/workspaces/project/utils/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true
						},
						"files": ["index.tlua"]
					}`),
				"/home/src/workspaces/project/utils/index.tlua": "local y = 10;",
				"/home/src/workspaces/project/valid/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true
						},
						"files": ["index.tlua"]
					}`),
				"/home/src/workspaces/project/valid/index.tlua": "local z = 10;",
			},
			commandLineArgs: []string{"--b", "--dry"},
		},
		{
			subScenario: "reports syntax errors in config file",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua": "function foo() end",
				"/home/src/workspaces/project/b.tlua": "function bar() end",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
						},
						"files": [
							"a.tlua"
							"b.tlua"
						]
					}`),
			},
			commandLineArgs: []string{"--b"},
			edits: []*tluaEdit{
				{
					caption: "reports syntax errors after change to config file",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/tluaconfig.json", ",", `, "declaration": true`)
					},
				},
				{
					caption: "reports syntax errors after change to ts file",
					edit: func(sys *TestSys) {
						sys.appendFile("/home/src/workspaces/project/a.tlua", "\nfunction fooBar() end")
					},
				},
				noChange,
				{
					caption: "builds after fixing config file errors",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", stringtestutil.Dedent(`
							{
								"compilerOptions": {
									"composite": true, "declaration": true
								},
								"files": [
									"a.tlua",
									"b.tlua"
								]
							}`))
					},
				},
			},
		},
		{
			subScenario:     "missing config file",
			files:           FileMap{},
			commandLineArgs: []string{"--b", "bogus.json"},
		},
		{
			subScenario: "reports syntax errors in config file",
			files: FileMap{
				"/home/src/workspaces/project/a.tlua": "function foo() end",
				"/home/src/workspaces/project/b.tlua": "function bar() end",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
						},
						"files": [
							"a.tlua"
							"b.tlua"
						]
					}`),
			},
			commandLineArgs: []string{"--b", "-w"},
			edits: []*tluaEdit{
				{
					caption: "reports syntax errors after change to config file",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/tluaconfig.json", ",", `, "declaration": true`)
					},
				},
				{
					caption: "reports syntax errors after change to ts file",
					edit: func(sys *TestSys) {
						sys.appendFile("/home/src/workspaces/project/a.tlua", "\nfunction fooBar() end")
					},
				},
				{
					caption: "reports error when there is no change to tsconfig file",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/tluaconfig.json", "", "")
					},
				},
				{
					caption: "builds after fixing config file errors",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/tluaconfig.json", stringtestutil.Dedent(`
							{
								"compilerOptions": {
									"composite": true, "declaration": true
								},
								"files": [
									"a.tlua",
									"b.tlua"
								]
							}`))
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "configFileErrors")
	}
}

func TestBuildDemoProject(t *testing.T) {
	t.Parallel()

	getBuildDemoFileMap := func(modify func(FileMap)) FileMap {
		files := FileMap{
			"/user/username/projects/demo/animals/animal.tlua": stringtestutil.Dedent(`
				type Size = "small" | "medium" | "large";
				interface Animal {
					size: Size;
				}
			`),
			"/user/username/projects/demo/animals/dog.tlua": stringtestutil.Dedent(`
				local utilities = require('core.utilities');

				interface Dog extends Animal {
					woof(): void;
					name: string;
				}

				function createDog(): Dog
					return ({
						size: "medium",
						woof: function()
							console.log("Woof!");
						end,
						name: utilities.makeRandomName()
					});
				end

				return { createDog = createDog };
			`),
			"/user/username/projects/demo/animals/index.tlua": stringtestutil.Dedent(`
				local dog = require('animals.dog');

				return { createDog = dog.createDog };
			`),
			"/user/username/projects/demo/animals/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../tluaconfig-base.json",
					"compilerOptions": {
						"outDir": "../lib/animals",
						"rootDir": ".."
					},
					"references": [
						{ "path": "../core" }
					]
				}
			`),
			"/user/username/projects/demo/core/utilities.tlua": stringtestutil.Dedent(`
				function makeRandomName()
					return "Bob!?! ";
				end

				function lastElementOf<T>(arr: T[]): T | undefined
					if arr.length == 0 then return undefined end
					return arr[arr.length - 1];
				end

				return { makeRandomName = makeRandomName, lastElementOf = lastElementOf };
			`),
			"/user/username/projects/demo/core/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../tluaconfig-base.json",
					"compilerOptions": {
						"outDir": "../lib/core",
						"rootDir": ".."
					},
				}
			`),
			"/user/username/projects/demo/zoo/zoo.tlua": stringtestutil.Dedent(`
				local animals = require('animals.index');

				function createZoo(): Array<Dog>
					return {
						animals.createDog()
					};
				end

				return { createZoo = createZoo };
			`),
			"/user/username/projects/demo/zoo/tluaconfig.json": stringtestutil.Dedent(`
				{
					"extends": "../tluaconfig-base.json",
					"compilerOptions": {
						"outDir": "../lib/zoo",
						"rootDir": ".."
					},
					"references": [
						{
							"path": "../animals"
						}
					]
				}
			`),
			"/user/username/projects/demo/tluaconfig-base.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"declaration": true,
						"target": "es5",
						"module": "commonjs",
						"strict": true,
						"noUnusedLocals": true,
						"noUnusedParameters": true,
						"noImplicitReturns": true,
						"composite": true,
					},
				}
			`),
			"/user/username/projects/demo/tluaconfig.json": stringtestutil.Dedent(`
				{
					"files": [],
					"references": [
						{
							"path": "./core"
						},
						{
							"path": "./animals",
						},
						{
							"path": "./zoo",
						},
					],
				}
			`),
		}
		if modify != nil {
			modify(files)
		}
		return files
	}
	testCases := []*tluaInput{
		{
			subScenario:     "in master branch with everything setup correctly and reports no error",
			files:           getBuildDemoFileMap(nil),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "--verbose"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario: "in circular branch reports the error about it by stopping build",
			files: getBuildDemoFileMap(func(files FileMap) {
				files["/user/username/projects/demo/core/tluaconfig.json"] = stringtestutil.Dedent(`
					{
						"extends": "../tluaconfig-base.json",
						"compilerOptions": {
							"outDir": "../lib/core",
							"rootDir": ".."
						},
						"references": [
							{
								"path": "../zoo",
							}
						]
					}
				`)
			}),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "--verbose"},
		},
		{
			// !!! sheetal - this has missing errors from strada about files not in rootDir (3)
			subScenario: "in bad-ref branch reports the error about files not in rootDir at the import location",
			files: getBuildDemoFileMap(func(files FileMap) {
				files["/user/username/projects/demo/core/utilities.tlua"] = `local A = require('animals.index');
A;
` + files["/user/username/projects/demo/core/utilities.tlua"].(string)
			}),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "--verbose"},
		},
		{
			subScenario: "in circular is set in the reference",
			files: getBuildDemoFileMap(func(files FileMap) {
				files["/user/username/projects/demo/a/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"extends": "../tluaconfig-base.json",
					"compilerOptions": {
						"outDir": "../lib/a",
						"rootDir": "."
					},
					"references": [
						{
							"path": "../b",
							"circular": true
						}
					]
				}`)
				files["/user/username/projects/demo/b/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"extends": "../tluaconfig-base.json",
					"compilerOptions": {
						"outDir": "../lib/b",
						"rootDir": "."
					},
					"references": [
						{
							"path": "../a",
						}
					]
				}`)
				files["/user/username/projects/demo/a/index.tlua"] = "local a = 10;"
				files["/user/username/projects/demo/b/index.tlua"] = "local b = 10;"
				files["/user/username/projects/demo/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"files": [],
					"references": [
						{
							"path": "./core"
						},
						{
							"path": "./animals",
						},
						{
							"path": "./zoo",
						},
						{
							"path": "./a",
						},
						{
							"path": "./b",
						},
					],
				}`)
			}),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "--verbose"},
		},
		{
			subScenario: "updates with circular reference",
			files: getBuildDemoFileMap(func(files FileMap) {
				files["/user/username/projects/demo/core/tluaconfig.json"] = stringtestutil.Dedent(`
					{
						"extends": "../tluaconfig-base.json",
						"compilerOptions": {
							"outDir": "../lib/core",
							"rootDir": ".."
						},
						"references": [
							{
								"path": "../zoo",
							}
						]
					}
				`)
			}),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "-w", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "Fix error",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/demo/core/tluaconfig.json", stringtestutil.Dedent(`
							{
								"extends": "../tluaconfig-base.json",
								"compilerOptions": {
									"outDir": "../lib/core",
									"rootDir": ".."
								},
							}
						`))
					},
				},
			},
		},
		{
			// !!! sheetal - this has missing errors from strada about files not in rootDir (3)
			subScenario: "updates with bad reference",
			files: getBuildDemoFileMap(func(files FileMap) {
				files["/user/username/projects/demo/core/utilities.tlua"] = `local A = require('animals.index');
A;
` + files["/user/username/projects/demo/core/utilities.tlua"].(string)
			}),
			cwd:             "/user/username/projects/demo",
			commandLineArgs: []string{"--b", "-w", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "Prepend a line",
					edit: func(sys *TestSys) {
						sys.prependFile("/user/username/projects/demo/core/utilities.tlua", "\n")
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "demo")
	}
}

func TestBuildFileDelete(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "detects deleted file",
			files: FileMap{
				"/home/src/workspaces/solution/child/child.tlua": stringtestutil.Dedent(`
					local child2 = require("child2");
					function child()
						child2.child2();
					end
					return { child = child };
				`),
				"/home/src/workspaces/solution/child/child2.tlua": stringtestutil.Dedent(`
					function child2() end
					return { child2 = child2 };
				`),
				"/home/src/workspaces/solution/child/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": { "composite": true }
					}
				`),
				"/home/src/workspaces/solution/main/main.tlua": stringtestutil.Dedent(`
                    local child = require("child.child");
                    function main()
                        child.child();
                    end
                    return { main = main };
                `),
				"/home/src/workspaces/solution/main/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": { "composite": true, "rootDir": ".." },
						"references": [{ "path": "../child" }],
					}
				`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "main/tluaconfig.json", "-v", "--traceResolution", "--explainFiles"},
			edits: []*tluaEdit{
				{
					caption: "delete child2 file",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/solution/child/child2.tlua")
						sys.removeNoError("/home/src/workspaces/solution/child/child2.lua")
						sys.removeNoError("/home/src/workspaces/solution/child/child2.d.tlua")
					},
				},
			},
		},
		{
			subScenario: "deleted file without composite",
			files: FileMap{
				"/home/src/workspaces/solution/child/child.tlua": stringtestutil.Dedent(`
					local child2 = require("child2");
					function child()
						child2.child2();
					end
					return { child = child };
				`),
				"/home/src/workspaces/solution/child/child2.tlua": stringtestutil.Dedent(`
					function child2() end
					return { child2 = child2 };
				`),
				"/home/src/workspaces/solution/child/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": { }
					}
				`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "child/tluaconfig.json", "-v", "--traceResolution", "--explainFiles"},
			edits: []*tluaEdit{
				{
					caption: "delete child2 file",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/solution/child/child2.tlua")
						sys.removeNoError("/home/src/workspaces/solution/child/child2.lua")
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "fileDelete")
	}
}

func TestBuildDependencyUpdate(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			// https://github.com/microsoft/typescript-go/issues/2666
			subScenario: "rebuilds when dependency in node_modules is updated",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
							"outDir": "dist",
							"strict": true
						},
						"include": ["src/**/*"]
					}
				`),
				"/home/src/workspaces/project/src/index.tlua": stringtestutil.Dedent(`
					local myDep = require("my-dep");
					local value: string = myDep.myValue;
				`),
				"/home/src/workspaces/project/node_modules/my-dep/package.json": stringtestutil.Dedent(`
					{
						"name": "my-dep",
						"version": "1.0.0"
					}
				`),
				"/home/src/workspaces/project/node_modules/my-dep/init.tlua": `local myValue: string = "";` + "\n" + `return { myValue = myValue };`,
			},
			cwd:             "/home/src/workspaces/project",
			commandLineArgs: []string{"--b", "--verbose"},
			edits: []*tluaEdit{
				noChange,
				{
					caption: "update dependency module with breaking type change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/node_modules/my-dep/init.tlua", "local myValue: number = 1;\nreturn { myValue = myValue };")
					},
					expectedDiff: "Incremental build does not re-check dependents of a changed Lua module (declaration emit is unsupported, TS100054, so module signatures never change), so the type error appears only in the clean build",
				},
				{
					caption: "restore dependency module",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/node_modules/my-dep/init.tlua", `local myValue: string = "";`+"\n"+`return { myValue = myValue };`)
					},
				},
				{
					caption: "update dependency module timestamp without changing text",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/node_modules/my-dep/init.tlua", `local myValue: string = "";`+"\n"+`return { myValue = myValue };`)
					},
				},
				noChange,
				{
					caption: "delete dependency module",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/node_modules/my-dep/init.tlua")
					},
				},
			},
		},
		{
			subScenario: "rebuilds when missing dependency package json is added",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
							"outDir": "dist",
							"strict": true
						},
						"include": ["src/**/*"]
					}
				`),
				"/home/src/workspaces/project/src/index.tlua": stringtestutil.Dedent(`
					local myDep = require("my-dep");
					local value: string = myDep.myValue;
				`),
				"/home/src/workspaces/project/node_modules/my-dep/index.tlua": `local myValue: string = "";` + "\n" + `return { myValue = myValue };`,
				"/home/src/workspaces/project/node_modules/my-dep/alt.tlua":   `local myValue: number = 1;` + "\n" + `return { myValue = myValue };`,
			},
			cwd:             "/home/src/workspaces/project",
			commandLineArgs: []string{"--b", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "add package json redirecting main to a module with a breaking type change",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/home/src/workspaces/project/node_modules/my-dep/package.json", `{"main":"alt.tlua"}`)
					},
					expectedDiff: "Incremental build does not re-check dependents of a changed Lua module (declaration emit is unsupported, TS100054, so module signatures never change), so the type error appears only in the clean build",
				},
			},
		},
		{
			subScenario: "rebuilds when dependency package json redirects to a different module",
			files: FileMap{
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
							"outDir": "dist",
							"strict": true
						},
						"include": ["src/**/*"]
					}
				`),
				"/home/src/workspaces/project/src/index.tlua": stringtestutil.Dedent(`
					local myDep = require("my-dep");
					local value: string = myDep.myValue;
				`),
				"/home/src/workspaces/project/node_modules/my-dep/package.json": stringtestutil.Dedent(`
					{
						"name": "my-dep",
						"version": "1.0.0",
						"main": "index.tlua"
					}
				`),
				"/home/src/workspaces/project/node_modules/my-dep/index.tlua": `local myValue: string = "";` + "\n" + `return { myValue = myValue };`,
				"/home/src/workspaces/project/node_modules/my-dep/alt.tlua":   `local myValue: number = 1;` + "\n" + `return { myValue = myValue };`,
			},
			cwd:             "/home/src/workspaces/project",
			commandLineArgs: []string{"--b", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "redirect package main to a module with a breaking type change",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/node_modules/my-dep/package.json", "index.tlua", "alt.tlua")
					},
					expectedDiff: "Incremental build does not re-check dependents of a changed Lua module (declaration emit is unsupported, TS100054, so module signatures never change), so the type error appears only in the clean build",
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "dependencyUpdate")
	}
}

func TestBuildLateBoundSymbol(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "interface is merged and contains late bound member",
			files: FileMap{
				"/home/src/workspaces/project/src/globals.d.tlua": stringtestutil.Dedent(`
                    interface SymbolConstructor {
                        (description?: string | number): symbol;
                    }
                    declare Symbol: SymbolConstructor;
                `),
				"/home/src/workspaces/project/src/hkt.tlua": `interface HKT<T> { }`,
				"/home/src/workspaces/project/src/main.tlua": stringtestutil.Dedent(`
                    local sym = Symbol();

                    interface Local<T> extends HKT<T> { }
                    interface Local<T> {
                        [sym]: { a: T }
                    }
                    local x = 10;
                    type A = Local<number>[typeof sym];
                `),
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "rootDir": "src",
                        "incremental": true,
                    },
                }`),
			},
			commandLineArgs: []string{"--b", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "incremental-declaration-doesnt-change",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/project/src/main.tlua", "local x = 10;", "")
					},
				},
				{
					caption: "incremental-declaration-doesnt-change",
					edit: func(sys *TestSys) {
						sys.appendFile("/home/src/workspaces/project/src/main.tlua", "local x = 10;")
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "lateBoundSymbol")
	}
}

func TestBuildOutputPaths(t *testing.T) {
	t.Parallel()
	type tluaOutputPathScenario struct {
		subScenario      string
		files            FileMap
		expectedDtsNames []string
	}
	runOutputPaths := func(s *tluaOutputPathScenario) {
		t.Helper()
		input := &tluaInput{
			subScenario:     s.subScenario,
			files:           s.files,
			commandLineArgs: []string{"-b", "-v"},
			edits: []*tluaEdit{
				noChange,
				{
					caption:         "Normal build without change, that does not block emit on error to show files that get emitted",
					commandLineArgs: []string{"-p", "/home/src/workspaces/project/tluaconfig.json"},
				},
			},
		}
		input.run(t, "outputPaths")
		t.Run("GetOutputFileNames/"+s.subScenario, func(t *testing.T) {
			t.Parallel()
			sys := newTestSys(input, false)
			config, _ := tsoptions.GetParsedCommandLineOfConfigFile("/home/src/workspaces/project/tluaconfig.json", &core.CompilerOptions{}, nil, sys, nil)
			assert.DeepEqual(t, slices.Collect(config.GetOutputFileNames()), s.expectedDtsNames)
		})
	}
	testCases := []*tluaOutputPathScenario{
		{
			subScenario: "when rootDir is not specified",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "dist",
                    },
                }`),
			},
			expectedDtsNames: []string{
				"/home/src/workspaces/project/dist/src/index.lua",
			},
		},
		{
			subScenario: "when rootDir is not specified and is composite",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "dist",
						"composite": true,
                    },
                }`),
			},
			expectedDtsNames: []string{
				"/home/src/workspaces/project/dist/src/index.lua",
				"/home/src/workspaces/project/dist/src/index.d.tlua",
			},
		},
		{
			subScenario: "when rootDir is specified",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "dist",
						"rootDir": "src",
                    },
                }`),
			},
			expectedDtsNames: []string{
				"/home/src/workspaces/project/dist/index.lua",
			},
		},
		{
			// !!! sheetal error missing as not yet implemented
			subScenario: "when rootDir is specified but not all files belong to rootDir",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua":  "local x = 10;",
				"/home/src/workspaces/project/types/type.tlua": "type t = string;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "dist",
						"rootDir": "src",
                    },
                }`),
			},
			expectedDtsNames: []string{
				"/home/src/workspaces/project/dist/index.lua",
				"/home/src/workspaces/project/types/type.lua",
			},
		},
		{
			// !!! sheetal error missing as not yet implemented
			subScenario: "when rootDir is specified but not all files belong to rootDir and is composite",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua":  "local x = 10;",
				"/home/src/workspaces/project/types/type.tlua": "type t = string;",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "dist",
						"rootDir": "src",
						"composite": true
                    },
                }`),
			},
			expectedDtsNames: []string{
				"/home/src/workspaces/project/dist/index.lua",
				"/home/src/workspaces/project/dist/index.d.tlua",
				"/home/src/workspaces/project/types/type.lua",
				"/home/src/workspaces/project/types/type.d.tlua",
			},
		},
	}
	for _, test := range testCases {
		runOutputPaths(test)
	}
}

func TestBuildProgramUpdates(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "when referenced project change introduces error in the down stream project and then fixes it",
			files: FileMap{
				"/user/username/projects/sample1/Library/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": {
						"composite": true
					}
				}`),
				"/user/username/projects/sample1/Library/library.tlua": stringtestutil.Dedent(`
					interface SomeObject
					{
						message: string;
					}

					function createSomeObject(): SomeObject
						return {
							message = "new Object"
						};
					end

					return { createSomeObject = createSomeObject };
				`),
				"/user/username/projects/sample1/App/tluaconfig.json": stringtestutil.Dedent(`
				{ 
					"compilerOptions": { "rootDir": ".." },
					"references": [{ "path": "../Library" }]
					
				}`),
				"/user/username/projects/sample1/App/app.tlua": stringtestutil.Dedent(`
					local library = require("Library.library");
					library.createSomeObject().message;
				`),
			},
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"-b", "-w", "App"},
			edits: []*tluaEdit{
				{
					caption: "Introduce error",
					// Change message in library to message2
					edit: func(sys *TestSys) {
						sys.replaceFileTextAll("/user/username/projects/sample1/Library/library.tlua", "message", "message2")
					},
				},
				{
					caption: "Fix error",
					// Revert library changes
					edit: func(sys *TestSys) {
						sys.replaceFileTextAll("/user/username/projects/sample1/Library/library.tlua", "message2", "message")
					},
				},
			},
		},
		{
			subScenario: "declarationEmitErrors when fixing error files all files are emitted",
			files: FileMap{
				"/user/username/projects/solution/app/fileWithError.tlua": stringtestutil.Dedent(`
					local myClassWithError = {
						tags: function() end,
						p: 12
					};
				`),
				"/user/username/projects/solution/app/fileWithoutError.tlua": "local myClass = { };",
				"/user/username/projects/solution/app/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
			},
			cwd:             "/user/username/projects/solution",
			commandLineArgs: []string{"-b", "-w", "app"},
			edits: []*tluaEdit{
				{
					caption: "Fix error",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/solution/app/fileWithError.tlua", "p: 12", "")
					},
				},
			},
		},
		{
			subScenario: "declarationEmitErrors when file with no error changes",
			files: FileMap{
				"/user/username/projects/solution/app/fileWithError.tlua": stringtestutil.Dedent(`
					local myClassWithError = {
						tags: function() end,
						p: 12
					};
				`),
				"/user/username/projects/solution/app/fileWithoutError.tlua": "local myClass = { };",
				"/user/username/projects/solution/app/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
			},
			cwd:             "/user/username/projects/solution",
			commandLineArgs: []string{"-b", "-w", "app"},
			edits: []*tluaEdit{
				{
					caption: "Change fileWithoutError",
					edit: func(sys *TestSys) {
						sys.replaceFileTextAll("/user/username/projects/solution/app/fileWithoutError.tlua", "myClass", "myClass2")
					},
				},
			},
		},
		{
			subScenario: "declarationEmitErrors introduceError when fixing errors only changed file is emitted",
			files: FileMap{
				"/user/username/projects/solution/app/fileWithError.tlua": stringtestutil.Dedent(`
					local myClassWithError = {
						tags: function() end,
					};
				`),
				"/user/username/projects/solution/app/fileWithoutError.tlua": "local myClass = { };",
				"/user/username/projects/solution/app/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
			},
			cwd:             "/user/username/projects/solution",
			commandLineArgs: []string{"-b", "-w", "app"},
			edits: []*tluaEdit{
				{
					caption: "Introduce error",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/solution/app/fileWithError.tlua", stringtestutil.Dedent(`
							local myClassWithError = {
								tags: function() end,
								p: 12
							};
						`))
					},
				},
				{
					caption: "Fix error",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/solution/app/fileWithError.tlua", "p: 12", "")
					},
				},
			},
		},
		{
			subScenario: "declarationEmitErrors introduceError when file with no error changes",
			files: FileMap{
				"/user/username/projects/solution/app/fileWithError.tlua": stringtestutil.Dedent(`
					local myClassWithError = {
						tags: function() end,
					};
				`),
				"/user/username/projects/solution/app/fileWithoutError.tlua": "local myClass = { };",
				"/user/username/projects/solution/app/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true
					}
				}`),
			},
			cwd:             "/user/username/projects/solution",
			commandLineArgs: []string{"-b", "-w", "app"},
			edits: []*tluaEdit{
				{
					caption: "Introduce error",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/solution/app/fileWithError.tlua", stringtestutil.Dedent(`
							local myClassWithError = {
								tags: function() end,
								p: 12
							};
						`))
					},
				},
				{
					caption: "Change fileWithoutError",
					edit: func(sys *TestSys) {
						sys.replaceFileTextAll("/user/username/projects/solution/app/fileWithoutError.tlua", "myClass", "myClass2")
					},
				},
			},
		},
		{
			subScenario: "works when noUnusedParameters changes to false",
			files: FileMap{
				"/user/username/projects/myproject/index.tlua": `local fn = function(a: string, b: string) return b end;`,
				"/user/username/projects/myproject/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"noUnusedParameters": true,
					},
				}`),
			},
			cwd:             "/user/username/projects/myproject",
			commandLineArgs: []string{"-b", "-w"},

			edits: []*tluaEdit{
				{
					caption: "Change tsconfig to set noUnusedParameters to false",
					edit: func(sys *TestSys) {
						sys.writeFileNoError(
							`/user/username/projects/myproject/tluaconfig.json`,
							stringtestutil.Dedent(`
							{
								"compilerOptions": {
									"noUnusedParameters": false,
								},
							}`),
						)
					},
				},
			},
		},
		{
			subScenario: "works with extended source files",
			cwd:         "/user/username/projects/project",
			files: FileMap{
				"/user/username/projects/project/commonFile1.tlua":      "local x = 1",
				"/user/username/projects/project/commonFile2.tlua":      "local y = 1",
				"/user/username/projects/project/alpha.tluaconfig.json": "{}",
				"/user/username/projects/project/project1.tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "./alpha.tluaconfig.json",
						"compilerOptions": {
							"composite": true,
						},
						"files": ["commonFile1.tlua", "commonFile2.tlua"],
					}
				`),
				"/user/username/projects/project/bravo.tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "./alpha.tluaconfig.json",
					}
				`),
				"/user/username/projects/project/other.tlua": "local z = 0;",
				"/user/username/projects/project/project2.tluaconfig.json": stringtestutil.Dedent(`
					{
						"extends": "./bravo.tluaconfig.json",
						"compilerOptions": {
							"composite": true,
						},
						"files": ["other.tlua"],
					}
				`),
				"/user/username/projects/project/other2.tlua": "local k = 0;",
				"/user/username/projects/project/extendsConfig1.tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
						},
					}
				`),
				"/user/username/projects/project/extendsConfig2.tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"strictNullChecks": false,
						},
					}
				`),
				"/user/username/projects/project/extendsConfig3.tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"noImplicitAny": true,
						},
					}
				`),
				"/user/username/projects/project/project3.tluaconfig.json": stringtestutil.Dedent(`
				{
                    "extends": [
                        "./extendsConfig1.tluaconfig.json",
                        "./extendsConfig2.tluaconfig.json",
                        "./extendsConfig3.tluaconfig.json",
                    ],
                    "compilerOptions": {
                        "composite": false,
                    },
                    "files": ["other2.tlua"],
                }`),
			},
			commandLineArgs: []string{"-b", "-w", "-v", "project1.tluaconfig.json", "project2.tluaconfig.json", "project3.tluaconfig.json"},
			edits: []*tluaEdit{
				{
					caption: "Modify alpha config",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/alpha.tluaconfig.json", stringtestutil.Dedent(`
						{
                            "compilerOptions": {
								"strict": true
							}
                        }`))
					},
				},
				{
					caption: "change bravo config",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/bravo.tluaconfig.json", stringtestutil.Dedent(`
						{
                            "extends": "./alpha.tluaconfig.json",
                            "compilerOptions": { "strict": false }
                        }`))
					},
				},
				{
					caption: "project 2 extends alpha",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/project2.tluaconfig.json", stringtestutil.Dedent(`
						{
                            "extends": "./alpha.tluaconfig.json",
                            "files": ["other.tlua"]
                        }`))
					},
				},
				{
					caption: "update aplha config",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/alpha.tluaconfig.json", "{}")
					},
				},
				{
					caption: "Modify extendsConfigFile2",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/extendsConfig2.tluaconfig.json", stringtestutil.Dedent(`
						{
                            "compilerOptions": { "strictNullChecks": true }
                        }`))
					},
				},
				{
					caption: "Modify project 3",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/project3.tluaconfig.json", stringtestutil.Dedent(`
						{
                            "extends": ["./extendsConfig1.tluaconfig.json", "./extendsConfig2.tluaconfig.json"],
                            "compilerOptions": { "composite": false },
                            "files": ["other2.tlua"],
                        }`))
					},
				},
				{
					caption: "Delete extendedConfigFile2 and report error",
					edit: func(sys *TestSys) {
						sys.removeNoError("/user/username/projects/project/extendsConfig2.tluaconfig.json")
					},
				},
			},
		},
		{
			subScenario: "works correctly when project with extended config is removed",
			files: FileMap{
				"/user/username/projects/project/commonFile1.tlua": "local x = 1",
				"/user/username/projects/project/commonFile2.tlua": "local y = 1",
				"/user/username/projects/project/alpha.tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "strict": true,
                    },
                }`),
				"/user/username/projects/project/project1.tluaconfig.json": stringtestutil.Dedent(`
				{
                    "extends": "./alpha.tluaconfig.json",
                    "compilerOptions": {
                        "composite": true,
                    },
                    "files": ["commonFile1.tlua", "commonFile2.tlua"],
                }`),
				"/user/username/projects/project/bravo.tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "strict": true,
                    },
                }`),
				"/user/username/projects/project/other.tlua": "local z = 0;",
				"/user/username/projects/project/project2.tluaconfig.json": stringtestutil.Dedent(`
				{
                    "extends": "./bravo.tluaconfig.json",
                    "compilerOptions": {
                        "composite": true,
                    },
                    "files": ["other.tlua"],
                }`),
				"/user/username/projects/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "references": [
                        {
                            "path": "./project1.tluaconfig.json",
                        },
                        {
                            "path": "./project2.tluaconfig.json",
                        },
                    ],
                    "files": [],
                }`),
			},
			cwd:             "/user/username/projects/project",
			commandLineArgs: []string{"-b", "-w", "-v"},
			edits: []*tluaEdit{
				{
					caption: "Remove project2 from base config",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/project/tluaconfig.json", stringtestutil.Dedent(`
						{
                            "references": [
                                {
                                    "path": "./project1.tluaconfig.json",
                                },
                            ],
                            "files": [],
                        }`))
					},
				},
			},
		},
		{
			subScenario: "tsbuildinfo has error",
			files: FileMap{
				"/user/username/projects/project/main.tlua":                "local x = 10;",
				"/user/username/projects/project/tluaconfig.json":          "{}",
				"/user/username/projects/project/tluaconfig.tluabuildinfo": "Some random string",
			},
			cwd:             "/user/username/projects/project",
			commandLineArgs: []string{"--b", "-i", "-w"},
		},
		{
			subScenario: "tsbuildinfo has fewer fileInfos than fileNames",
			files: FileMap{
				"/user/username/projects/project/src/a.tlua":      "local a = 1;",
				"/user/username/projects/project/src/b.tlua":      "local b = 2;",
				"/user/username/projects/project/tluaconfig.json": `{"compilerOptions":{"composite":true,"outDir":"dist"},"files":["src/a.tlua","src/b.tlua"]}`,
				"/user/username/projects/project/dist/tluaconfig.tluabuildinfo": `{
					"version": "FakeTSVersion",
					"fileNames": ["lib.es2025.full.d.tlua", "../src/a.tlua", "../src/b.tlua"],
					"fileInfos": ["abc123"],
					"options": {"composite": true, "outDir": "./"},
					"root": [2, 3]
				}`,
			},
			cwd:             "/user/username/projects/project",
			commandLineArgs: []string{"--b", "-v"},
		},
		{
			subScenario: "when root is source from project reference",
			files: FileMap{
				"/home/src/workspaces/project/lib/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
							"outDir": "./dist"
						}
					}`),
				"/home/src/workspaces/project/lib/foo.tlua": "local FOO: string = 'THEFOOEXPORT';\nreturn { FOO = FOO };",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"references": [ { "path": "./lib" } ]
					}`),
				"/home/src/workspaces/project/index.tlua": "local foo = require(\"lib.foo\");\nfoo.FOO;",
			},
			commandLineArgs: []string{"--b"},
			edits: []*tluaEdit{
				{
					caption: "dts doesnt change",
					edit: func(sys *TestSys) {
						// foo.tlua ends in a top-level return, so add the statement at the top
						sys.prependFile("/home/src/workspaces/project/lib/foo.tlua", "local Bar = 10;\n")
					},
				},
			},
			cwd: "/home/src/workspaces/project",
		},
		{
			subScenario: "when root is source from project reference with composite",
			files: FileMap{
				"/home/src/workspaces/project/lib/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
							"outDir": "./dist"
						}
					}`),
				"/home/src/workspaces/project/lib/foo.tlua": "local FOO: string = 'THEFOOEXPORT';\nreturn { FOO = FOO };",
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
					{
						"compilerOptions": {
							"composite": true,
						},
						"references": [ { "path": "./lib" } ]
					}`),
				"/home/src/workspaces/project/index.tlua": "local foo = require(\"lib.foo\");\nfoo.FOO;",
			},
			commandLineArgs: []string{"--b"},
			edits: []*tluaEdit{
				{
					caption: "dts doesnt change",
					edit: func(sys *TestSys) {
						// foo.tlua ends in a top-level return, so add the statement at the top
						sys.prependFile("/home/src/workspaces/project/lib/foo.tlua", "local Bar = 10;\n")
					},
				},
			},
			cwd: "/home/src/workspaces/project",
		},
	}
	for _, test := range testCases {
		test.run(t, "programUpdates")
	}
}

func TestBuildProjectsBuilding(t *testing.T) {
	t.Parallel()
	addPackageFiles := func(files FileMap, index int) {
		files[fmt.Sprintf(`/user/username/projects/myproject/pkg%d/index.tlua`, index)] = fmt.Sprintf("local pkg%d = %d;\nreturn { pkg%d = pkg%d };", index, index, index, index)
		var references string
		if index > 0 {
			references = `"references": [{ "path": "../pkg0" }],`
		}
		files[fmt.Sprintf(`/user/username/projects/myproject/pkg%d/tluaconfig.json`, index)] = stringtestutil.Dedent(fmt.Sprintf(`
		{
			"compilerOptions": { "composite": true },
			%s
		}`, references))
	}
	addSolution := func(files FileMap, count int) {
		var pkgReferences []string
		for i := range count {
			pkgReferences = append(pkgReferences, fmt.Sprintf(`{ "path": "./pkg%d" }`, i))
		}
		files[`/user/username/projects/myproject/tluaconfig.json`] = stringtestutil.Dedent(fmt.Sprintf(`
		{
			"compilerOptions": { "composite": true },
			"references": [
				%s
			]
		}`, strings.Join(pkgReferences, ",\n\t\t\t\t")))
	}
	files := func(count int) FileMap {
		files := FileMap{}
		for i := range count {
			addPackageFiles(files, i)
		}
		addSolution(files, count)
		return files
	}

	getTestCases := func(pkgCount int, builders int) []*tluaInput {
		getEdits := func(watch bool) []*tluaEdit {
			// In watch mode the up-to-date incremental cycles skip re-reporting
			// the TS100054 declaration emit errors that a clean build repeats.
			noChangeEdit := core.IfElse(watch, noChangeWithDeclarationEmitErrorReplayDiff, noChange)
			return []*tluaEdit{
				{
					caption: "dts doesn't change",
					edit: func(sys *TestSys) {
						// index.tlua ends in a top-level return, so add the local-only statement at the top
						sys.prependFile(`/user/username/projects/myproject/pkg0/index.tlua`, "local someConst2 = 10;\n")
					},
				},
				noChangeEdit,
				{
					caption: "dts change",
					edit: func(sys *TestSys) {
						// changing the return table is the tlua equivalent of adding an export
						sys.replaceFileText(`/user/username/projects/myproject/pkg0/index.tlua`, "return { pkg0 = pkg0 };", "local someConst = 10;\nreturn { pkg0 = pkg0, someConst = someConst };")
					},
				},
				noChangeEdit,
			}
		}
		edits := getEdits(false)
		watchEdits := getEdits(true)
		return []*tluaInput{
			{
				subScenario:     fmt.Sprintf(`when there are %d projects in a solution`, pkgCount),
				files:           files(pkgCount),
				cwd:             "/user/username/projects/myproject",
				commandLineArgs: []string{"-b", "-v"},
				edits:           edits,
			},
			{
				subScenario:     fmt.Sprintf(`when there are %d projects in a solution with --builders %d`, pkgCount, builders),
				files:           files(pkgCount),
				cwd:             "/user/username/projects/myproject",
				commandLineArgs: []string{"-b", "-v", "--builders", strconv.Itoa(builders)},
				edits:           edits,
			},
			{
				subScenario:     fmt.Sprintf(`when there are %d projects in a solution`, pkgCount),
				files:           files(pkgCount),
				cwd:             "/user/username/projects/myproject",
				commandLineArgs: []string{"-b", "-w", "-v"},
				edits:           watchEdits,
			},
			{
				subScenario:     fmt.Sprintf(`when there are %d projects in a solution with --builders %d`, pkgCount, builders),
				files:           files(pkgCount),
				cwd:             "/user/username/projects/myproject",
				commandLineArgs: []string{"-b", "-w", "-v", "--builders", strconv.Itoa(builders)},
				edits:           watchEdits,
			},
		}
	}

	testCases := slices.Concat(
		getTestCases(3, 1),
		getTestCases(5, 2),
		getTestCases(8, 3),
		getTestCases(23, 3),
	)

	for _, test := range testCases {
		test.run(t, "projectsBuilding")
	}
}

func TestBuildProjectReferenceWithRootDirInParent(t *testing.T) {
	t.Parallel()
	getBuildProjectReferenceWithRootDirInParentFileMap := func(modify func(files FileMap)) FileMap {
		files := FileMap{
			"/home/src/workspaces/solution/src/main/a.tlua": stringtestutil.Dedent(`
				local b = require('main.b');
				local a = b.b;
			`),
			"/home/src/workspaces/solution/src/main/b.tlua": stringtestutil.Dedent(`
				local b = 0;
				return { b = b };
			`),
			"/home/src/workspaces/solution/src/main/tluaconfig.json": stringtestutil.Dedent(`
			{
				"extends": "../../tluaconfig.base.json",
				"references": [
					{ "path": "../other" },
				],
			}`),
			"/home/src/workspaces/solution/src/other/other.tlua": stringtestutil.Dedent(`
				local Other = 0;
			`),
			"/home/src/workspaces/solution/src/other/tluaconfig.json": stringtestutil.Dedent(`
			{
				"extends": "../../tluaconfig.base.json",
			}
			`),
			"/home/src/workspaces/solution/tluaconfig.base.json": stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
					"declaration": true,
					"rootDir": "./src/",
					"outDir": "./dist/",
					"skipDefaultLibCheck": true,
				},
				"exclude": [
					"node_modules",
				],
			}`),
		}
		if modify != nil {
			modify(files)
		}
		return files
	}
	testCases := []*tluaInput{
		{
			subScenario:     "builds correctly",
			files:           getBuildProjectReferenceWithRootDirInParentFileMap(nil),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/main", "/home/src/workspaces/solution/src/other"},
		},
		{
			subScenario: "reports error for same tsbuildinfo file because no rootDir in the base",
			files: getBuildProjectReferenceWithRootDirInParentFileMap(
				func(files FileMap) {
					text, _ := files["/home/src/workspaces/solution/tluaconfig.base.json"]
					files["/home/src/workspaces/solution/tluaconfig.base.json"] = strings.Replace(text.(string), `"rootDir": "./src/",`, "", 1)
				},
			),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/main", "--verbose"},
		},
		{
			subScenario: "reports error for same tsbuildinfo file",
			files: getBuildProjectReferenceWithRootDirInParentFileMap(
				func(files FileMap) {
					files["/home/src/workspaces/solution/src/main/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                        "references": [{ "path": "../other" }]
                    }`)
					files["/home/src/workspaces/solution/src/other/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                    }`)
				},
			),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/main", "--verbose"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario: "reports error for same tsbuildinfo file without incremental",
			files: getBuildProjectReferenceWithRootDirInParentFileMap(
				func(files FileMap) {
					files["/home/src/workspaces/solution/src/main/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "outDir": "../../dist/" },
                        "references": [{ "path": "../other" }]
                    }`)
					files["/home/src/workspaces/solution/src/other/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                    }`)
				},
			),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/main", "--verbose"},
		},
		{
			subScenario: "reports error for same tsbuildinfo file without incremental with tlua",
			files: getBuildProjectReferenceWithRootDirInParentFileMap(
				func(files FileMap) {
					files["/home/src/workspaces/solution/src/main/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "outDir": "../../dist/" },
                        "references": [{ "path": "../other" }]
                    }`)
					files["/home/src/workspaces/solution/src/other/tluaconfig.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                    }`)
				},
			),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/other", "--verbose"},
			edits: []*tluaEdit{
				{
					caption:         "Running tlua on main",
					commandLineArgs: []string{"-p", "src/main"},
				},
			},
		},
		{
			subScenario: "reports no error when tsbuildinfo differ",
			files: getBuildProjectReferenceWithRootDirInParentFileMap(
				func(files FileMap) {
					delete(files, "/home/src/workspaces/solution/src/main/tluaconfig.json")
					delete(files, "/home/src/workspaces/solution/src/other/tluaconfig.json")
					files["/home/src/workspaces/solution/src/main/tluaconfig.main.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                        "references": [{ "path": "../other/tluaconfig.other.json" }]
                    }`)
					files["/home/src/workspaces/solution/src/other/tluaconfig.other.json"] = stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true, "outDir": "../../dist/" },
                    }`)
				},
			),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "src/main/tluaconfig.main.json", "--verbose"},
			edits:           noChangeOnlyEdit,
		},
	}

	for _, test := range testCases {
		test.run(t, "projectReferenceWithRootDirInParent")
	}
}

func TestBuildReexport(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "Reports errors correctly",
			files: FileMap{
				"/user/username/projects/reexport/src/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "files": [],
                    "include": [],
                    "references": [{ "path": "./pure" }, { "path": "./main" }],
                }`),
				"/user/username/projects/reexport/src/main/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "outDir": "../../out",
                        "rootDir": "../",
                    },
                    "include": ["**/*.tlua"],
                    "references": [{ "path": "../pure" }],
                }`),
				"/user/username/projects/reexport/src/main/index.tlua": stringtestutil.Dedent(`
                    local pure = require("pure");
                    pure;

                    local session: Session = {
                        foo: 1
                    };
                `),
				"/user/username/projects/reexport/src/pure/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": {
                        "composite": true,
                        "outDir": "../../out",
                        "rootDir": "../",
                    },
                    "include": ["**/*.tlua"],
                }`),
				"/user/username/projects/reexport/src/pure/init.tlua": "local session = require(\"pure.session\");\nsession;",
				"/user/username/projects/reexport/src/pure/session.tlua": stringtestutil.Dedent(`
                    interface Session {
                        foo: number;
                        // bar: number;
                    }
                `),
			},
			cwd:             `/user/username/projects/reexport`,
			commandLineArgs: []string{"-b", "-w", "-verbose", "src"},
			edits: []*tluaEdit{
				{
					caption: "Introduce error",
					edit: func(sys *TestSys) {
						sys.replaceFileText(`/user/username/projects/reexport/src/pure/session.tlua`, "// ", "")
					},
				},
				{
					caption: "Fix error",
					edit: func(sys *TestSys) {
						sys.replaceFileText(`/user/username/projects/reexport/src/pure/session.tlua`, "bar: ", "// bar: ")
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test.run(t, "reexport")
	}
}

func TestBuildRoots(t *testing.T) {
	t.Parallel()
	getBuildRootsFromProjectReferencedProjectFileMap := func(serverFirst bool) FileMap {
		include := core.IfElse(serverFirst, `"src/**/*.tlua", "../shared/src/**/*.tlua"`, `"../shared/src/**/*.tlua", "src/**/*.tlua"`)
		return FileMap{
			"/home/src/workspaces/solution/tluaconfig.json": stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
				},
				"references": [
					{ "path": "projects/server" },
					{ "path": "projects/shared" },
				],
			}`),
			"/home/src/workspaces/solution/projects/shared/src/myClass.tlua": "local MyClass = { };\nreturn { MyClass = MyClass };",
			"/home/src/workspaces/solution/projects/shared/src/logging.tlua": stringtestutil.Dedent(`
				function log(str: string)
					console.log(str);
				end
			`),
			"/home/src/workspaces/solution/projects/shared/src/random.tlua": stringtestutil.Dedent(`
				function randomFn(str: string)
					console.log(str);
				end
			`),
			"/home/src/workspaces/solution/projects/shared/tluaconfig.json": stringtestutil.Dedent(`
			{
				"extends": "../../tluaconfig.json",
				"compilerOptions": {
					"outDir": "./dist",
				},
				"include": ["src/**/*.tlua"],
			}`),
			"/home/src/workspaces/solution/projects/server/src/server.tlua": stringtestutil.Dedent(`
				local myClass = require('shared.src.myClass');
				console.log('Hello, world!');
			`),
			"/home/src/workspaces/solution/projects/server/tluaconfig.json": stringtestutil.Dedent(fmt.Sprintf(`
			{
				"extends": "../../tluaconfig.json",
				"compilerOptions": {
					"rootDir": "..",
					"outDir": "./dist",
				},
				"include": [ %s ],
				"references": [
					{ "path": "../shared" },
				],
			}`, include)),
		}
	}
	getBuildRootsFromProjectReferencedProjectTestEdits := func(watch bool) []*tluaEdit {
		// In watch mode the up-to-date incremental cycles skip re-reporting the
		// TS100054 declaration emit errors that a clean build repeats.
		noChangeEdit := core.IfElse(watch, noChangeWithDeclarationEmitErrorReplayDiff, noChange)
		return []*tluaEdit{
			noChangeEdit,
			{
				caption: "edit logging file",
				edit: func(sys *TestSys) {
					sys.appendFile("/home/src/workspaces/solution/projects/shared/src/logging.tlua", "\nlocal x = 10;")
				},
			},
			noChangeEdit,
			{
				caption: "delete random file",
				edit: func(sys *TestSys) {
					sys.removeNoError("/home/src/workspaces/solution/projects/shared/src/random.tlua")
				},
			},
			noChangeEdit,
		}
	}
	testCases := []*tluaInput{
		{
			subScenario: `when two root files are consecutive`,
			files: FileMap{
				"/home/src/workspaces/project/file1.tlua": `local x = "hello";`,
				"/home/src/workspaces/project/file2.tlua": `local y = "world";`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": { "composite": true },
                    "include": ["*.tlua"],
                }`),
			},
			commandLineArgs: []string{"--b", "-v"},
			edits: []*tluaEdit{
				{
					caption: "delete file1",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/file1.tlua")
						sys.removeNoError("/home/src/workspaces/project/file1.lua")
						sys.removeNoError("/home/src/workspaces/project/file1.d.tlua")
					},
				},
			},
		},
		{
			subScenario: `when multiple root files are consecutive`,
			files: FileMap{
				"/home/src/workspaces/project/file1.tlua": `local x = "hello";`,
				"/home/src/workspaces/project/file2.tlua": `local y = "world";`,
				"/home/src/workspaces/project/file3.tlua": `local y = "world";`,
				"/home/src/workspaces/project/file4.tlua": `local y = "world";`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": { "composite": true },
                    "include": ["*.tlua"],
                }`),
			},
			commandLineArgs: []string{"--b", "-v"},
			edits: []*tluaEdit{
				{
					caption: "delete file1",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/file1.tlua")
						sys.removeNoError("/home/src/workspaces/project/file1.lua")
						sys.removeNoError("/home/src/workspaces/project/file1.d.tlua")
					},
				},
			},
		},
		{
			subScenario: `when files are not consecutive`,
			files: FileMap{
				"/home/src/workspaces/project/file1.tlua":    `local x = "hello";`,
				"/home/src/workspaces/project/random.d.tlua": `declare random: string;`,
				"/home/src/workspaces/project/file2.tlua": stringtestutil.Dedent(`
                    /// <reference path="./random.d.tlua"/>
                    local y = "world";
                `),
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": { "composite": true },
                    "include": ["file*.tlua"],
                }`),
			},
			commandLineArgs: []string{"--b", "-v"},
			edits: []*tluaEdit{
				{
					caption: "delete file1",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/file1.tlua")
						sys.removeNoError("/home/src/workspaces/project/file1.lua")
						sys.removeNoError("/home/src/workspaces/project/file1.d.tlua")
					},
				},
			},
		},
		{
			subScenario: `when consecutive and non consecutive are mixed`,
			files: FileMap{
				"/home/src/workspaces/project/file1.tlua":    `local x = "hello";`,
				"/home/src/workspaces/project/file2.tlua":    `local y = "world";`,
				"/home/src/workspaces/project/random.d.tlua": `declare random: string;`,
				"/home/src/workspaces/project/nonconsecutive.tlua": stringtestutil.Dedent(`
                /// <reference path="./random.d.tlua"/>
					local nonConsecutive = "hello";
				`),
				"/home/src/workspaces/project/random1.d.tlua": `declare random: string;`,
				"/home/src/workspaces/project/asArray1.tlua": stringtestutil.Dedent(`
					/// <reference path="./random1.d.tlua"/>
					local x = "hello";
				`),
				"/home/src/workspaces/project/asArray2.tlua":  `local x = "hello";`,
				"/home/src/workspaces/project/asArray3.tlua":  `local x = "hello";`,
				"/home/src/workspaces/project/random2.d.tlua": `declare random: string;`,
				"/home/src/workspaces/project/anotherNonConsecutive.tlua": stringtestutil.Dedent(`
					/// <reference path="./random2.d.tlua"/>
					local nonConsecutive = "hello";
				`),
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
                    "compilerOptions": { "composite": true },
                    "include": ["file*.tlua", "nonconsecutive*.tlua", "asArray*.tlua", "anotherNonConsecutive.tlua"],
                }`),
			},
			commandLineArgs: []string{"--b", "-v"},
			edits: []*tluaEdit{
				{
					caption: "delete file1",
					edit: func(sys *TestSys) {
						sys.removeNoError("/home/src/workspaces/project/file1.tlua")
						sys.removeNoError("/home/src/workspaces/project/file1.lua")
						sys.removeNoError("/home/src/workspaces/project/file1.d.tlua")
					},
				},
			},
		},
		{
			subScenario:     "when root file is from referenced project",
			files:           getBuildRootsFromProjectReferencedProjectFileMap(true),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "projects/server", "-v", "--traceResolution", "--explainFiles"},
			edits:           getBuildRootsFromProjectReferencedProjectTestEdits(false),
		},
		{
			subScenario:     "when root file is from referenced project and shared is first",
			files:           getBuildRootsFromProjectReferencedProjectFileMap(false),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "projects/server", "-v", "--traceResolution", "--explainFiles"},
			edits:           getBuildRootsFromProjectReferencedProjectTestEdits(false),
		},
		{
			subScenario:     "when root file is from referenced project",
			files:           getBuildRootsFromProjectReferencedProjectFileMap(true),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "-w", "projects/server", "-v", "--traceResolution", "--explainFiles"},
			edits:           getBuildRootsFromProjectReferencedProjectTestEdits(true),
		},
		{
			subScenario:     "when root file is from referenced project and shared is first",
			files:           getBuildRootsFromProjectReferencedProjectFileMap(false),
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "-w", "projects/server", "-v", "--traceResolution", "--explainFiles"},
			edits:           getBuildRootsFromProjectReferencedProjectTestEdits(true),
		},
	}

	for _, test := range testCases {
		test.run(t, "roots")
	}
}

func TestBuildSample(t *testing.T) {
	t.Parallel()

	getLogicConfig := func() string {
		return stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
					"declaration": true,
					"sourceMap": true,
					"skipDefaultLibCheck": true,
					"rootDir": "..",
				},
				"references": [
					{ "path": "../core" },
				],
			}`)
	}

	getBuildSampleFileMap := func(modify func(files FileMap)) FileMap {
		files := FileMap{
			"/user/username/projects/sample1/core/tluaconfig.json": stringtestutil.Dedent(`
			{
				"compilerOptions": {
					"composite": true,
					"declaration": true,
					"declarationMap": true,
					"skipDefaultLibCheck": true,
				},
			}`),
			"/user/username/projects/sample1/core/index.tlua": stringtestutil.Dedent(`
				local someString: string = "HELLO WORLD";
				function leftPad(s: string, n: number) return s + n; end
				function multiply(a: number, b: number) return a * b; end
				return { someString = someString, leftPad = leftPad, multiply = multiply };
			`),
			"/user/username/projects/sample1/core/some_decl.d.tlua":   `declare dts: any;`,
			"/user/username/projects/sample1/core/anotherModule.tlua": "local World = \"hello\";\nreturn { World = World };",
			"/user/username/projects/sample1/logic/tluaconfig.json":   getLogicConfig(),
			"/user/username/projects/sample1/logic/index.tlua": stringtestutil.Dedent(`
				local c = require('core.index');
				function getSecondsInDay()
					return c.multiply(10, 15);
				end
				local mod = require('core.anotherModule');
				return { getSecondsInDay = getSecondsInDay, m = mod };
			`),
			"/user/username/projects/sample1/tests/tluaconfig.json": stringtestutil.Dedent(`
			{
				"references": [
					{ "path": "../core" },
					{ "path": "../logic" },
				],
				"files": ["index.tlua"],
				"compilerOptions": {
					"composite": true,
					"declaration": true,
					"skipDefaultLibCheck": true,
					"rootDir": "..",
					"rootDir": "..",
				},
			}`),
			"/user/username/projects/sample1/tests/index.tlua": stringtestutil.Dedent(`
				local c = require('core.index');
				local logic = require('logic.index');

				c.leftPad("", 10);
				logic.getSecondsInDay();

				local mod = require('core.anotherModule');
				return { m = mod };
			`),
		}
		if modify != nil {
			modify(files)
		}
		return files
	}
	getStopBuildOnErrorTests := func(options []string) []*tluaInput {
		noChange := core.IfElse(options == nil, noChangeOnlyEdit, nil)
		return []*tluaInput{
			{
				subScenario: "skips builds downstream projects if upstream projects have errors with stopBuildOnErrors",
				files: getBuildSampleFileMap(func(files FileMap) {
					// index.tlua ends in a top-level return; prepend the erroneous call (function decls hoist)
					text, _ := files["/user/username/projects/sample1/core/index.tlua"]
					files["/user/username/projects/sample1/core/index.tlua"] = "multiply();\n" + text.(string)
				}),
				cwd:             "/user/username/projects/sample1",
				commandLineArgs: slices.Concat([]string{"--b", "tests", "--verbose", "--stopBuildOnErrors"}, options),
				edits: slices.Concat(
					noChange,
					[]*tluaEdit{
						{
							caption: "fix error",
							edit: func(sys *TestSys) {
								sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "multiply();\n", "")
							},
						},
					},
				),
			},
			{
				subScenario: "skips builds downstream projects if upstream projects have errors with stopBuildOnErrors when test does not reference core",
				files: getBuildSampleFileMap(func(files FileMap) {
					files["/user/username/projects/sample1/tests/tluaconfig.json"] = stringtestutil.Dedent(`
					{
						"references": [
							{ "path": "../logic" },
						],
						"files": ["index.tlua"],
						"compilerOptions": {
							"composite": true,
							"declaration": true,
							"skipDefaultLibCheck": true,
						},
					}`)
					// index.tlua ends in a top-level return; prepend the erroneous call (function decls hoist)
					text, _ := files["/user/username/projects/sample1/core/index.tlua"]
					files["/user/username/projects/sample1/core/index.tlua"] = "multiply();\n" + text.(string)
				}),
				cwd:             "/user/username/projects/sample1",
				commandLineArgs: slices.Concat([]string{"--b", "tests", "--verbose", "--stopBuildOnErrors"}, options),
				edits: slices.Concat(
					noChange,
					[]*tluaEdit{
						{
							caption: "fix error",
							edit: func(sys *TestSys) {
								sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "multiply();\n", "")
							},
						},
					},
				),
			},
		}
	}
	getBuildSampleCoreChangeEdits := func() []*tluaEdit {
		return []*tluaEdit{
			{
				caption: "incremental-declaration-changes",
				edit: func(sys *TestSys) {
					// adding someClass to the return table is the tlua equivalent of adding an export
					sys.replaceFileText(
						"/user/username/projects/sample1/core/index.tlua",
						"return { someString = someString",
						"local someClass = 10;\nreturn { someClass = someClass, someString = someString",
					)
				},
			},
			{
				caption: "incremental-declaration-doesnt-change",
				edit: func(sys *TestSys) {
					// index.tlua ends in a top-level return, so add the local-only statement at the top
					sys.prependFile("/user/username/projects/sample1/core/index.tlua", "local someClass2 = 10;\n")
				},
			},
			noChange,
		}
	}
	getBuildSampleWatchDtsChangingEdits := func() []*tluaEdit {
		return []*tluaEdit{
			{
				caption: "Make change to core",
				edit: func(sys *TestSys) {
					// adding someClass to the return table is the tlua equivalent of adding an export
					sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "return { someString = someString", "local someClass = 10;\nreturn { someClass = someClass, someString = someString")
				},
			},
			{
				caption: "Revert core file",
				edit: func(sys *TestSys) {
					sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "local someClass = 10;\nreturn { someClass = someClass, someString = someString", "return { someString = someString")
				},
			},
			{
				caption: "Make two changes",
				edit: func(sys *TestSys) {
					sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "return { someString = someString", "local someClass = 10;\nreturn { someClass = someClass, someString = someString")
					sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "return { someClass = someClass", "local someClass2 = 10;\nreturn { someClass2 = someClass2, someClass = someClass")
				},
			},
		}
	}
	getBuildSampleWatchNonDtsChangingEdits := func() []*tluaEdit {
		return []*tluaEdit{
			{
				caption: "Make local change to core",
				edit: func(sys *TestSys) {
					// index.tlua ends in a top-level return, so add the local-only function at the top
					sys.prependFile("/user/username/projects/sample1/core/index.tlua", "function foo() end\n")
				},
			},
		}
	}
	getBuildSampleWatchNewFileEdits := func() []*tluaEdit {
		return []*tluaEdit{
			{
				caption: "Change to new File and build core",
				edit: func(sys *TestSys) {
					sys.writeFileNoError("/user/username/projects/sample1/core/newfile.tlua", `local newFileConst = 30;`)
				},
			},
			{
				caption: "Change to new File and build core",
				edit: func(sys *TestSys) {
					sys.writeFileNoError("/user/username/projects/sample1/core/newfile.tlua", "local someClass2 = 10;")
				},
			},
		}
	}
	makeCircularReferences := func(files FileMap) {
		files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
		{
			"compilerOptions": {
				"composite": true,
				"declaration": true
			},
			"references": [
				{ "path": "../tests", "circular": true }
			],
		}`)
	}
	getIncrementalErrorTest := func(subScenario string, options []string) *tluaInput {
		return &tluaInput{
			subScenario:     "reportErrors " + subScenario,
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: slices.Concat([]string{"-b", "-w", "tests"}, options),
			edits: []*tluaEdit{
				{
					caption: "change logic",
					edit: func(sys *TestSys) {
						// index.tlua ends in a top-level return, so prepend the erroneous statement
						sys.prependFile("/user/username/projects/sample1/logic/index.tlua", "local y: string = 10;\n")
					},
				},
				{
					caption: "change core",
					edit: func(sys *TestSys) {
						// index.tlua ends in a top-level return, so prepend the erroneous statement
						sys.prependFile("/user/username/projects/sample1/core/index.tlua", "local x: string = 10;\n")
					},
				},
				{
					caption: "fix error in logic",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/logic/index.tlua", "local y: string = 10;\n", "")
					},
				},
			},
		}
	}
	testCases := slices.Concat([]*tluaInput{
		{
			subScenario: "builds correctly when outDir is specified",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/logic/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"declaration": true,
						"sourceMap": true,
						"outDir": "outDir",
						"rootDir": "..",
					},
					"references": [
						{ "path": "../core" },
					],
				}`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests"},
		},
		{
			subScenario: "builds correctly when declarationDir is specified",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/logic/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"declaration": true,
						"sourceMap": true,
						"declarationDir": "out/decls",
						"rootDir": "..",
					},
					"references": [
						{ "path": "../core" },
					],
				}`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests"},
		},
		{
			subScenario: "builds correctly when project is not composite or doesnt have any references",
			files: getBuildSampleFileMap(func(files FileMap) {
				text, _ := files["/user/username/projects/sample1/core/tluaconfig.json"]
				files["/user/username/projects/sample1/core/tluaconfig.json"] = strings.Replace(text.(string), `"composite": true,`, "", 1)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "core", "--verbose"},
		},
		{
			subScenario:     "does not write any files in a dry build",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--dry"},
		},
		{
			subScenario:     "removes all files it built",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests"},
			edits: []*tluaEdit{
				{
					caption:         "removes all files it built",
					commandLineArgs: []string{"--b", "tests", "--clean"},
				},
				{
					caption:         "no change --clean",
					commandLineArgs: []string{"--b", "tests", "--clean"},
				},
			},
		},
		{
			subScenario:     "cleaning project in not build order doesnt throw error",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "logic2", "--clean"},
		},
		{
			subScenario:     "always builds under with force option",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--force"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario:     "can detect when and what to rebuild",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: []*tluaEdit{
				noChange,
				{
					// Update a file in the leaf node (tests), only it should rebuild the last one
					caption: "Only builds the leaf node project",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/sample1/tests/index.tlua", "local m = 10;")
					},
				},
				{
					// Update a file in the parent (without affecting types), should get fast downstream builds
					caption: "Detects type-only changes in upstream projects",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/index.tlua", "HELLO WORLD", "WELCOME PLANET")
					},
				},
				{
					caption: "rebuilds when tsconfig changes",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/tests/tluaconfig.json", `"composite": true`, `"composite": true, "target": "es2020"`)
					},
				},
			},
		},
		{
			subScenario:     "when input file text does not change but its modified time changes",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "upstream project changes without changing file text",
					edit: func(sys *TestSys) {
						err := sys.FS().Chtimes("/user/username/projects/sample1/core/index.tlua", time.Time{}, sys.Now())
						if err != nil {
							panic(err)
						}
					},
				},
			},
		},
		{
			subScenario:     "when declarationMap changes",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "Disable declarationMap",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.json", `"declarationMap": true,`, `"declarationMap": false,`)
					},
				},
				{
					caption: "Enable declarationMap",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.json", `"declarationMap": false,`, `"declarationMap": true,`)
					},
				},
			},
		},
		{
			subScenario:     "indicates that it would skip builds during a dry build",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests"},
			edits: []*tluaEdit{
				{
					caption:         "--dry",
					commandLineArgs: []string{"--b", "tests", "--dry"},
				},
			},
		},
		{
			subScenario:     "rebuilds from start if force option is set",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests"},
			edits: []*tluaEdit{
				{
					caption:         "--force build",
					commandLineArgs: []string{"--b", "tests", "--verbose", "--force"},
				},
			},
		},
		{
			subScenario: "tsbuildinfo has error",
			files: FileMap{
				"/home/src/workspaces/project/main.tlua":                "local x = 10;",
				"/home/src/workspaces/project/tluaconfig.json":          "{}",
				"/home/src/workspaces/project/tluaconfig.tluabuildinfo": "Some random string",
			},
			commandLineArgs: []string{"--b", "-i", "-v"},
			edits: []*tluaEdit{
				{
					caption: "tsbuildinfo written has error",
					edit: func(sys *TestSys) {
						// This is to ensure the non incremental doesnt crash - as it wont have tsbuildInfo
						if !sys.forIncrementalCorrectness {
							sys.prependFile("/home/src/workspaces/project/tluaconfig.tluabuildinfo", "Some random string")
							sys.replaceFileText("/home/src/workspaces/project/tluaconfig.tluabuildinfo", fmt.Sprintf(`"version":"%s"`, core.Version()), fmt.Sprintf(`"version":"%s"`, harnessutil.FakeTSVersion)) // build info won't parse, need to manually sterilize for baseline
						}
					},
				},
			},
		},
		{
			subScenario:     "rebuilds completely when version in tsbuildinfo doesnt match ts version",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "convert tsbuildInfo version to something that is say to previous version",
					edit: func(sys *TestSys) {
						// This is to ensure the non incremental doesnt crash - as it wont have tsbuildInfo
						if !sys.forIncrementalCorrectness {
							sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.tluabuildinfo", fmt.Sprintf(`"version":"%s"`, harnessutil.FakeTSVersion), fmt.Sprintf(`"version":"%s"`, "FakeTsPreviousVersion"))
							sys.replaceFileText("/user/username/projects/sample1/logic/tluaconfig.tluabuildinfo", fmt.Sprintf(`"version":"%s"`, harnessutil.FakeTSVersion), fmt.Sprintf(`"version":"%s"`, "FakeTsPreviousVersion"))
							sys.replaceFileText("/user/username/projects/sample1/tests/tluaconfig.tluabuildinfo", fmt.Sprintf(`"version":"%s"`, harnessutil.FakeTSVersion), fmt.Sprintf(`"version":"%s"`, "FakeTsPreviousVersion"))
						}
					},
				},
			},
		},
		{
			subScenario: "rebuilds when extended config file changes",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/tests/tluaconfig.base.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5"
					}
				}`)
				text, _ := files["/user/username/projects/sample1/tests/tluaconfig.json"]
				files["/user/username/projects/sample1/tests/tluaconfig.json"] = strings.Replace(text.(string), `"references": [`, `"extends": "./tluaconfig.base.json", "references": [`, 1)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "change extended file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/sample1/tests/tluaconfig.base.json", stringtestutil.Dedent(`
						{
							"compilerOptions": { }
						}`))
					},
				},
			},
		},
		{
			subScenario:     "building project in not build order doesnt throw error",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "logic2/tluaconfig.json", "--verbose"},
		},
		{
			subScenario: "builds downstream projects even if upstream projects have errors",
			files: getBuildSampleFileMap(func(files FileMap) {
				text, _ := files["/user/username/projects/sample1/logic/index.tlua"]
				files["/user/username/projects/sample1/logic/index.tlua"] = strings.Replace(text.(string), "c.multiply(10, 15)", `c.muitply()`, 1)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario:     "listFiles",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--listFiles"},
			edits:           getBuildSampleCoreChangeEdits(),
		},
		{
			subScenario:     "listEmittedFiles",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--listEmittedFiles"},
			edits:           getBuildSampleCoreChangeEdits(),
		},
		{
			subScenario:     "explainFiles",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--explainFiles", "--v"},
			edits:           getBuildSampleCoreChangeEdits(),
		},
		{
			subScenario:     "sample",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
			edits: slices.Concat(
				getBuildSampleCoreChangeEdits(),
				[]*tluaEdit{
					{
						caption: "when logic config changes declaration dir",
						edit: func(sys *TestSys) {
							sys.replaceFileText(
								"/user/username/projects/sample1/logic/tluaconfig.json",
								`"declaration": true,`,
								`"declaration": true,
        "declarationDir": "decls",`,
							)
						},
					},
					{
						caption: "no change",
					},
				},
			),
		},
		{
			subScenario: "when logic specifies tsBuildInfoFile",
			files: getBuildSampleFileMap(func(files FileMap) {
				text, _ := files["/user/username/projects/sample1/logic/tluaconfig.json"]
				files["/user/username/projects/sample1/logic/tluaconfig.json"] = strings.Replace(
					text.(string),
					`"composite": true,`,
					`"composite": true,
    "tsBuildInfoFile": "ownFile.tluabuildinfo",`,
					1,
				)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
		},
		{
			subScenario: "when declaration option changes",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"incremental": true,
						"skipDefaultLibCheck": true,
					},
				}`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "core", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "incremental-declaration-changes",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.json", `"incremental": true,`, `"incremental": true, "declaration": true,`)
					},
				},
			},
		},
		{
			subScenario: "when target option changes",
			files: getBuildSampleFileMap(func(files FileMap) {
				files[getTestLibPathFor("esnext.full")] = `/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />`
				files[tluaLibPath+"/lib.d.tlua"] = `/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />`
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"incremental": true,
						"listFiles": true,
						"listEmittedFiles": true,
						"target": "esnext",
					},
				}`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "core", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "incremental-declaration-changes",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.json", `esnext`, `es5`)
					},
				},
			},
		},
		{
			subScenario: "when module option changes",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"incremental": true,
						"module": "node18",
					},
				}`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "core", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "incremental-declaration-changes",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/user/username/projects/sample1/core/tluaconfig.json", `node18`, `nodenext`)
					},
				},
			},
		},
		{
			// !!! sheetal this is not reporting error as file not found is not yet implemented
			subScenario: "reports error if input file is missing",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					 "compilerOptions": { "composite": true },
					 "files": ["anotherModule.tlua", "index.tlua", "some_decl.d.tlua"],
				}`)
				delete(files, "/user/username/projects/sample1/core/anotherModule.tlua")
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose"},
		},
		{
			// !!! sheetal this is not reporting error as file not found is not yet implemented
			subScenario: "reports error if input file is missing with force",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					 "compilerOptions": { "composite": true },
					 "files": ["anotherModule.tlua", "index.tlua", "some_decl.d.tlua"],
				}`)
				delete(files, "/user/username/projects/sample1/core/anotherModule.tlua")
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "tests", "--verbose", "--force"},
		},
		{
			subScenario:     "change builds changes and reports found errors message",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchDtsChangingEdits(),
		},
		{
			subScenario:     "non local change does not start build of referencing projects",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchNonDtsChangingEdits(),
		},
		{
			subScenario:     "builds when new file is added, and its subsequent updates",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchNewFileEdits(),
		},
		{
			subScenario:     "change builds changes and reports found errors message with circular references",
			files:           getBuildSampleFileMap(makeCircularReferences),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchDtsChangingEdits(),
		},
		{
			subScenario:     "non local change does not start build of referencing projects with circular references",
			files:           getBuildSampleFileMap(makeCircularReferences),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchNonDtsChangingEdits(),
		},
		{
			subScenario:     "builds when new file is added, and its subsequent updates with circular references",
			files:           getBuildSampleFileMap(makeCircularReferences),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits:           getBuildSampleWatchNewFileEdits(),
		},
		{
			subScenario: "watches config files that are not present",
			files: getBuildSampleFileMap(func(files FileMap) {
				delete(files, "/user/username/projects/sample1/logic/tluaconfig.json")
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests"},
			edits: []*tluaEdit{
				{
					caption: "Write logic",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/sample1/logic/tluaconfig.json", getLogicConfig())
					},
				},
			},
		},
		getIncrementalErrorTest("when preserveWatchOutput is not used", nil),
		getIncrementalErrorTest("when preserveWatchOutput is passed on command line", []string{"--preserveWatchOutput"}),
		getIncrementalErrorTest("when stopBuildOnErrors is passed on command line", []string{"--stopBuildOnErrors"}),
		{
			subScenario:     "incremental updates in verbose mode",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "tests", "--verbose"},
			edits: []*tluaEdit{
				{
					caption: "Make non dts change",
					edit: func(sys *TestSys) {
						// index.tlua ends in a top-level return, so add the local-only function at the top
						sys.prependFile("/user/username/projects/sample1/logic/index.tlua", "function someFn() end\n")
					},
				},
				{
					caption: "Make dts change",
					edit: func(sys *TestSys) {
						// adding someFn to the return table is the tlua equivalent of exporting it
						sys.replaceFileText("/user/username/projects/sample1/logic/index.tlua", "return { getSecondsInDay = getSecondsInDay", "return { someFn = someFn, getSecondsInDay = getSecondsInDay")
					},
				},
			},
		},
		{
			subScenario:     "should not trigger recompilation because of program emit",
			files:           getBuildSampleFileMap(nil),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "core", "--verbose"},
			edits: []*tluaEdit{
				noChangeWithDeclarationEmitErrorReplayDiff,
				{
					caption: "Add new file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/sample1/core/file3.tlua", `local y = 10;`)
					},
				},
				noChangeWithDeclarationEmitErrorReplayDiff,
			},
		},
		{
			subScenario: "should not trigger recompilation because of program emit with outDir specified",
			files: getBuildSampleFileMap(func(files FileMap) {
				files["/user/username/projects/sample1/core/tluaconfig.json"] = stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"outDir": "outDir"
					}
                }`)
			}),
			cwd:             "/user/username/projects/sample1",
			commandLineArgs: []string{"--b", "-w", "core", "--verbose"},
			edits: []*tluaEdit{
				noChangeWithDeclarationEmitErrorReplayDiff,
				{
					caption: "Add new file",
					edit: func(sys *TestSys) {
						sys.writeFileNoError("/user/username/projects/sample1/core/file3.tlua", `local y = 10;`)
					},
				},
				noChangeWithDeclarationEmitErrorReplayDiff,
			},
		},
	}, getStopBuildOnErrorTests(nil), getStopBuildOnErrorTests([]string{"--watch"}))

	for _, test := range testCases {
		test.run(t, "sample")
	}
}

func TestBuildTransitiveReferences(t *testing.T) {
	t.Parallel()

	getBuildTransitiveReferencesFileMap := func(modify func(files FileMap)) FileMap {
		files := FileMap{
			"/user/username/projects/transitiveReferences/refs/a.d.tlua": stringtestutil.Dedent(`
				declare X: { kind: string };
				declare A: { kind: string };
			`),
			"/user/username/projects/transitiveReferences/a.tlua": stringtestutil.Dedent(`
				local A = { kind: 'A' };
				return { A = A };
			`),
			"/user/username/projects/transitiveReferences/b.tlua": stringtestutil.Dedent(`
				local a = require('a');
				local b = a.A;
				return { b = b };
			`),
			"/user/username/projects/transitiveReferences/c.tlua": stringtestutil.Dedent(`
				/// <reference path="./refs/a.d.tlua"/>
				local b = require('b');
				b.b;
				X.kind;
			`),
			"/user/username/projects/transitiveReferences/tluaconfig.a.json": stringtestutil.Dedent(`
			{
				"files": ["a.tlua"],
				"compilerOptions": {
					"composite": true,
				},
			}`),
			"/user/username/projects/transitiveReferences/tluaconfig.b.json": stringtestutil.Dedent(`
			{
				"files": ["b.tlua"],
				"compilerOptions": {
					"composite": true,
				},
				"references": [{ "path": "tluaconfig.a.json" }],
			}`),
			"/user/username/projects/transitiveReferences/tluaconfig.c.json": stringtestutil.Dedent(`
			{
				"files": ["c.tlua"],
								"references": [{ "path": "tluaconfig.b.json" }],
			}`),
		}
		if modify != nil {
			modify(files)
		}
		return files
	}
	testCases := []*tluaInput{
		{
			subScenario:     "builds correctly",
			files:           getBuildTransitiveReferencesFileMap(nil),
			cwd:             "/user/username/projects/transitiveReferences",
			commandLineArgs: []string{"--b", "tluaconfig.c.json", "--listFiles"},
		},
	}

	for _, test := range testCases {
		test.run(t, "transitiveReferences")
	}
}

func TestBuildSolutionProject(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario: "verify that subsequent builds after initial build doesnt build anything",
			files: FileMap{
				"/home/src/workspaces/solution/src/folder/index.tlua": `local x = 10;`,
				"/home/src/workspaces/solution/src/folder/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "files": ["index.tlua"],
                        "compilerOptions": {
                            "composite": true
                        }
                    }
                `),
				"/home/src/workspaces/solution/src/folder2/index.tlua": `local x = 10;`,
				"/home/src/workspaces/solution/src/folder2/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "files": ["index.tlua"],
                        "compilerOptions": {
                            "composite": true
                        }
                    }
                `),
				"/home/src/workspaces/solution/src/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "files": [],
                        "compilerOptions": {
                            "composite": true
                        },
						"references": [
							{ "path": "./folder" },
							{ "path": "./folder2" },
						]
                }`),
				"/home/src/workspaces/solution/tests/index.tlua": `local x = 10;`,
				"/home/src/workspaces/solution/tests/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "files": ["index.tlua"],
                        "compilerOptions": {
                            "composite": true
                        },
                        "references": [
                            { "path": "../src" }
                        ]
                    }
                `),
				"/home/src/workspaces/solution/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "files": [],
                        "compilerOptions": {
                            "composite": true
                        },
                        "references": [
                            { "path": "./src" },
                            { "path": "./tests" }
                        ]
                    }
                `),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "--v"},
			edits:           noChangeOnlyEdit,
		},
		{
			subScenario: "when solution is referenced indirectly",
			files: FileMap{
				"/home/src/workspaces/solution/project1/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true },
                        "references": []
                    }
                `),
				"/home/src/workspaces/solution/project2/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true },
                        "references": []
                    }
                `),
				"/home/src/workspaces/solution/project2/src/b.tlua": "local b = 10;",
				"/home/src/workspaces/solution/project3/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true },
                        "references": [
							{ "path": "../project1" },
							{ "path": "../project2" }
						]
                    }
                `),
				"/home/src/workspaces/solution/project3/src/c.tlua": "local c = 10;",
				"/home/src/workspaces/solution/project4/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": { "composite": true },
                        "references": [{ "path": "../project3" }]
                    }
                `),
				"/home/src/workspaces/solution/project4/src/d.tlua": "local d = 10;",
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "project4", "--verbose", "--explainFiles"},
			edits: []*tluaEdit{
				{
					caption: "modify project3 file",
					edit: func(sys *TestSys) {
						sys.replaceFileText("/home/src/workspaces/solution/project3/src/c.tlua", "c = ", "cc = ")
					},
				},
			},
		},
		{
			subScenario: "has empty files diagnostic when files is empty and no references are provided",
			files: FileMap{
				"/home/src/workspaces/solution/no-references/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "references": [],
                        "files": [],
                        "compilerOptions": {
                            "composite": true,
                            "declaration": true,
                            "forceConsistentCasingInFileNames": true,
                            "skipDefaultLibCheck": true,
                        },
                    }`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "no-references"},
		},
		{
			subScenario: "does not have empty files diagnostic when files is empty and references are provided",
			files: FileMap{
				"/home/src/workspaces/solution/core/index.tlua": "function multiply(a: number, b: number) return a * b; end",
				"/home/src/workspaces/solution/core/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "compilerOptions": {
                            "composite": true,
                            "declaration": true,
                            "declarationMap": true,
                            "skipDefaultLibCheck": true,
                        },
                    }`),
				"/home/src/workspaces/solution/with-references/tluaconfig.json": stringtestutil.Dedent(`
                    {
                        "references": [
                            { "path": "../core" },
                        ],
                        "files": [],
                        "compilerOptions": {
                            "composite": true,
                            "declaration": true,
                            "forceConsistentCasingInFileNames": true,
                            "skipDefaultLibCheck": true,
                        },
                    }`),
			},
			cwd:             "/home/src/workspaces/solution",
			commandLineArgs: []string{"--b", "with-references"},
		},
	}

	for _, test := range testCases {
		test.run(t, "solution")
	}
}

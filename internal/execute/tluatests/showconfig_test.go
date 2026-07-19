package tluatests

import (
	"testing"

	"github.com/apyrr/tlua/internal/testutil/stringtestutil"
)

func TestShowConfig(t *testing.T) {
	t.Parallel()
	testCases := []*tluaInput{
		{
			subScenario:     "Default initialized TSConfig",
			commandLineArgs: []string{"--showConfig"},
		},
		{
			subScenario:     "Show TSConfig with files options",
			commandLineArgs: []string{"--showConfig", "file0.tlua", "file1.tlua", "file2.tlua"},
		},
		{
			subScenario:     "Show TSConfig with boolean value compiler options",
			commandLineArgs: []string{"--showConfig", "--noUnusedLocals"},
		},
		{
			subScenario:     "Show TSConfig with enum value compiler options",
			commandLineArgs: []string{"--showConfig", "--target", "es5", "--jsx", "react"},
		},
		{
			subScenario:     "Show TSConfig with list compiler options",
			commandLineArgs: []string{"--showConfig", "--types", "jquery,mocha"},
		},
		{
			subScenario:     "Show TSConfig with list compiler options with enum value",
			commandLineArgs: []string{"--showConfig", "--lib", "es5,es2015.core"},
		},
		{
			subScenario:     "Show TSConfig with incorrect compiler option",
			commandLineArgs: []string{"--showConfig", "--someNonExistOption"},
		},
		{
			subScenario:     "Show TSConfig with incorrect compiler option value",
			commandLineArgs: []string{"--showConfig", "--lib", "nonExistLib,es5,es2015.promise"},
		},
		{
			subScenario:     "Show TSConfig with advanced options",
			commandLineArgs: []string{"--showConfig", "--declaration", "--declarationDir", "lib", "--skipLibCheck", "--noErrorTruncation"},
		},
		{
			subScenario: "Show TSConfig with compileOnSave and more",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"target": "es5",
						"module": "commonjs",
						"strict": true
					},
					"compileOnSave": true,
					"exclude": [
						"dist"
					],
					"files": [],
					"include": [
						"src/*"
					],
					"references": [
						{ "path": "./test" }
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with more options",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"outDir": "./lib",
						"module": "commonjs",
						"target": "ES2017",
						"sourceMap": true,
						"experimentalDecorators": true,
						"emitDecoratorMetadata": true
					},
					"include": [
						"./src/**/*"
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with include filtering files",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": `local a = 1;`,
				"/home/src/workspaces/project/src/util.tlua": `local b = 2;`,
				"/home/src/workspaces/project/extra.tlua":    `local c = 3;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true
					},
					"include": [
						"src/**/*"
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with references",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"composite": true,
						"strict": true
					},
					"references": [
						{ "path": "./packages/a" },
						{ "path": "./packages/b" }
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with exclude",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua":  `local a = 1;`,
				"/home/src/workspaces/project/test/test1.tlua": `local src = require("src.index");`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true
					},
					"exclude": [
						"test"
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with files and include",
			files: FileMap{
				"/home/src/workspaces/project/src/main.tlua": `local a = 1;`,
				"/home/src/workspaces/project/extra.tlua":    `local c = 3;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true
					},
					"files": [
						"extra.tlua"
					],
					"include": [
						"src/**/*"
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with transitively implied options",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua": `local a = 1;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"module": "nodenext"
					}
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
		{
			subScenario: "Show TSConfig with exclude and outDir",
			files: FileMap{
				"/home/src/workspaces/project/src/index.tlua":    `local a = 1;`,
				"/home/src/workspaces/project/src/bin/tool.tlua": `local b = 2;`,
				"/home/src/workspaces/project/tluaconfig.json": stringtestutil.Dedent(`
				{
					"compilerOptions": {
						"strict": true,
						"outDir": "./build"
					},
					"exclude": [
						"build"
					]
				}`),
			},
			commandLineArgs: []string{"-p", "tluaconfig.json", "--showConfig"},
		},
	}

	for _, test := range testCases {
		test.run(t, "showConfig")
	}
}

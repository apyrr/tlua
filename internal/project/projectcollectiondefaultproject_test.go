package project_test

import (
	"context"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil/projecttestutil"
)

func TestProjectCollectionDefaultProject(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	// Project 1 references project 2, which does not have open files.
	// File project1/dist/index.d.tlua does not belong to any tsconfig.json, but is included in programs for
	// projects 3 and 4 via a triple-slash path reference to project 1's output.
	// When looking for a default project for project1/dist/index.d.tlua,
	// we should not try to unconditionally access project 2,
	// which isn't loaded because of `disableReferencedProjectLoad`.
	files := map[string]any{
		"/project1/tsconfig.json": `{
			"extends": "../tsconfig.json",
			"files": [],
			"include": ["src/**/*"],
			"references": [
				{
					"path": "../project2"
				}
			],
			"compilerOptions": {
				"composite": true,
				"outDir": "./dist",
				"rootDir": "./src",
			}
		}`,
		"/project1/src/index.tlua": `local foo = 42;
		return { foo = foo };`,
		"/project1/dist/index.d.tlua": `type Bar = {
				a: string;
			};`,
		"/project2/tsconfig.json": `{
			"extends": "../tsconfig.json",
			"files": [],
			"include": ["src/**/*"],
			"compilerOptions": {
				"composite": true,
				"outDir": "./dist",
				"rootDir": "./src"
			}
		}`,
		"/project3/tsconfig.json": `{
			"extends": "../tsconfig.json",
			"files": [],
			"include": ["src/**/*"],
			"references": [
				{
					"path": "../project1"
				}
			],
			"compilerOptions": {
				"composite": true,
				"outDir": "./dist",
				"rootDir": "./src",
			}
		}`,
		"/project3/src/index.tlua": `/// <reference path="../../project1/dist/index.d.tlua"/>
			local x: Bar = { a: "s" };`,
		"/project4/tsconfig.json": `{
			"extends": "../tsconfig.json",
			"files": [],
			"include": ["src/**/*"],
			"references": [
				{
					"path": "../project1"
				}
			],
			"compilerOptions": {
				"composite": true,
				"outDir": "./dist",
				"rootDir": "./src",
			}
		}`,
		"/project4/src/index.tlua": `/// <reference path="../../project1/dist/index.d.tlua"/>
local x: Bar = { a: "s" };`,
		"/tsconfig.json": `{
			"compilerOptions": {
				"disableReferencedProjectLoad": true,
				"disableSolutionSearching": true,
				"disableSourceOfProjectReferenceRedirect": true
			},
			"files": [],
			"references": [
				{
					"path": "./project1"
				},
				{
					"path": "./project2"
				},
				{
					"path": "./project3"
				},
				{
					"path": "./project4"
				}
			]
		}`,
	}
	uris := []lsproto.DocumentUri{
		"file:///project1/dist/index.d.tlua",
		"file:///project1/src/index.tlua",
		"file:///project3/src/index.tlua",
		"file:///project4/src/index.tlua",
	}
	session, _ := projecttestutil.Setup(files)
	// Should not crash.
	for _, uri := range uris {
		content := files[string(uri)[7:]].(string)
		session.DidOpenFile(context.Background(), uri, 1, content, lsproto.LanguageKindTypeScript)
	}
}

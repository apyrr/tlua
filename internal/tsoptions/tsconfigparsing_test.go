package tsoptions_test

import (
	"fmt"
	"io"
	"io/fs"
	"maps"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/diagnosticwriter"
	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/scanner"
	"github.com/apyrr/tlua/internal/testutil/baseline"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tsoptions/tsoptionstest"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs"
	"gotest.tools/v3/assert"
)

type testConfig struct {
	jsonText       string
	configFileName string
	basePath       string
	allFileList    map[string]string
}

var parseConfigFileTextToJsonTests = []struct {
	title string
	input []string
}{
	{
		title: "returns empty config for file with only whitespaces",
		input: []string{
			"",
			" ",
		},
	},
	{
		title: "returns empty config for file with comments only",
		input: []string{
			"// Comment",
			"/* Comment*/",
		},
	},
	{
		title: "returns empty config when config is empty object",
		input: []string{
			`{}`,
		},
	},
	{
		title: "returns config object without comments",
		input: []string{
			`{ // Excluded files
            "exclude": [
                // Exclude d.ts
                "file.d.ts"
            ]
        }`,
			`{
            /* Excluded
                    Files
            */
            "exclude": [
                /* multiline comments can be in the middle of a line */"file.d.ts"
            ]
        }`,
		},
	},
	{
		title: "keeps string content untouched",
		input: []string{
			`{
            "exclude": [
                "xx//file.d.ts"
            ]
        }`,
			`{
            "exclude": [
                "xx/*file.d.ts*/"
            ]
        }`,
		},
	},
	{
		title: "handles escaped characters in strings correctly",
		input: []string{
			`{
            "exclude": [
                "xx\"//files"
            ]
        }`,
			`{
            "exclude": [
                "xx\\" // end of line comment
            ]
        }`,
		},
	},
	{
		title: "returns object when users correctly specify library",
		input: []string{
			`{
            "compilerOptions": {
                "lib": ["es5"]
            }
        }`,
			`{
            "compilerOptions": {
                "lib": ["es5", "es6"]
            }
        }`,
		},
	},
}

func TestParseConfigFileTextToJson(t *testing.T) {
	t.Parallel()
	for _, rec := range parseConfigFileTextToJsonTests {
		t.Run(rec.title, func(t *testing.T) {
			t.Parallel()
			var baselineContent strings.Builder
			for i, jsonText := range rec.input {
				baselineContent.WriteString("Input::\n")
				baselineContent.WriteString(jsonText)
				baselineContent.WriteString("\n")
				parsed, errors := tsoptions.ParseConfigFileTextToJson("/apath/tluaconfig.json", "/apath", jsonText)
				baselineContent.WriteString("Config::\n")
				assert.NilError(t, writeJsonReadableText(&baselineContent, parsed), "Failed to write JSON text")
				baselineContent.WriteString("\n")
				baselineContent.WriteString("Errors::\n")
				diagnosticwriter.FormatDiagnosticsWithColorAndContext(&baselineContent, diagnosticwriter.FromASTDiagnostics(errors), &diagnosticwriter.FormattingOptions{
					NewLine: "\n",
					ComparePathsOptions: tspath.ComparePathsOptions{
						CurrentDirectory:          "/",
						UseCaseSensitiveFileNames: true,
					},
				})
				baselineContent.WriteString("\n")
				if i != len(rec.input)-1 {
					baselineContent.WriteString("\n")
				}
			}
			// The upstream jsonParse baselines lived in the TypeScript
			// submodule; these are now tracked as in-repo baselines.
			baseline.Run(t, rec.title+" jsonParse.js", baselineContent.String(), baseline.Options{Subfolder: "config/tsconfigParsing"})
		})
	}
}

type parseJsonConfigTestCase struct {
	title string
	input []testConfig
}

var parseJsonConfigFileTests = []parseJsonConfigTestCase{
	{
		title: "ignore dotted files and folders",
		input: []testConfig{{
			jsonText:       `{}`,
			configFileName: "tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/test.tlua": "", "/apath/.git/a.tlua": "", "/apath/.b.tlua": "", "/apath/..c.tlua": ""},
		}},
	},
	{
		title: "allow dotted files and folders when explicitly requested",
		input: []testConfig{{
			jsonText: `{
                    "files": ["/apath/.git/a.tlua", "/apath/.b.tlua", "/apath/..c.tlua"]
                }`,
			configFileName: "tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/test.tlua": "", "/apath/.git/a.tlua": "", "/apath/.b.tlua": "", "/apath/..c.tlua": ""},
		}},
	},
	{
		title: "implicitly exclude common package folders",
		input: []testConfig{{
			jsonText:       `{}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/node_modules/a.tlua": "", "/bower_components/b.tlua": "", "/jspm_packages/c.tlua": "", "/d.tlua": "", "/folder/e.tlua": ""},
		}},
	},
	{
		title: "generates errors for empty files list",
		input: []testConfig{{
			jsonText: `{
                "files": []
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "generates errors for empty files list when no references are provided",
		input: []testConfig{{
			jsonText: `{
                "files": [],
                "references": []
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "generates errors for directory with no .tlua files",
		input: []testConfig{{
			jsonText: `{
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.lua": ""},
		}},
	},
	{
		title: "generates errors for empty include",
		input: []testConfig{{
			jsonText: `{
                "include": []
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "tests/cases/unittests",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "generates errors for include with parent directory after recursive wildcard",
		input: []testConfig{{
			jsonText: `{
                "include": ["**/../*.tlua"]
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/main.tlua": ""},
		}},
	},
	{
		title: "parses tsconfig with compilerOptions, files, include, and exclude",
		input: []testConfig{{
			jsonText: `{
  "compilerOptions": {
    "outDir": "./dist",
    "strict": true,
    "noImplicitAny": true,
    "target": "ES2017",
    "module": "ESNext",
    "jsx": "react",
	"maxNodeModuleJsDepth": 1
  },
  "files": ["/apath/src/index.tlua", "/apath/src/app.tlua"],
  "include": ["/apath/src/**/*"],
  "exclude": ["/apath/node_modules", "/apath/dist"]
}`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/src/index.tlua": "", "/apath/src/app.tlua": "", "/apath/node_modules/module.tlua": "", "/apath/dist/output.lua": ""},
		}},
	},
	{
		title: "generates errors when commandline option is in tsconfig",
		input: []testConfig{{
			jsonText: `{
  "compilerOptions": {
    "help": true
  }
}`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "does not generate errors for empty files list when one or more references are provided",
		input: []testConfig{{
			jsonText: `{
                "files": [],
                "references": [{ "path": "/apath" }]
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "exclude outDir unless overridden",
		input: []testConfig{{
			jsonText: `{
                "compilerOptions": {
                    "outDir": "bin"
                }
            }`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/bin/a.tlua": "", "/b.tlua": ""},
		}, {
			jsonText: `{
                "compilerOptions": {
                    "outDir": "bin"
                },
                "exclude": [ "obj" ]
            }`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/bin/a.tlua": "", "/b.tlua": ""},
		}},
	},
	{
		title: "exclude declarationDir unless overridden",
		input: []testConfig{{
			jsonText: `{
                "compilerOptions": {
                    "declarationDir": "declarations"
                }
            }`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/declarations/a.d.tlua": "", "/a.tlua": ""},
		}, {
			jsonText: `{
                "compilerOptions": {
                    "declarationDir": "declarations"
                },
                "exclude": [ "types" ]
            }`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/declarations/a.d.tlua": "", "/a.tlua": ""},
		}},
	},
	{
		title: "generates errors for empty directory",
		input: []testConfig{{
			jsonText: `{
                "compilerOptions": {
                    "allowJs": true
                }
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{},
		}},
	},
	{
		title: "generates errors for includes with outDir",
		input: []testConfig{{
			jsonText: `{
                "compilerOptions": {
                    "outDir": "./"
                },
                "include": ["**/*"]
            }`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "generates errors when include is not string",
		input: []testConfig{{
			jsonText: `{
  "include": [
    [
      "./**/*.tlua"
    ]
  ]
}`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "generates errors when files is not string",
		input: []testConfig{{
			jsonText: `{
  "files": [
    {
      "compilerOptions": {
        "experimentalDecorators": true,
        "allowJs": true
      }
    }
  ]
}`,
			configFileName: "/apath/tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/a.tlua": ""},
		}},
	},
	{
		title: "with outDir from base tsconfig",
		input: []testConfig{
			{
				jsonText: `{
  "extends": "./tsconfigWithoutConfigDir.json"
}`,
				configFileName: "tluaconfig.json",
				basePath:       "/",
				allFileList: map[string]string{
					"/tsconfigWithoutConfigDir.json": tsconfigWithoutConfigDir,
					"/bin/a.tlua":                    "",
					"/b.tlua":                        "",
				},
			},
			{
				jsonText: `{
  "extends": "./tsconfigWithConfigDir.json"
}`,
				configFileName: "tluaconfig.json",
				basePath:       "/",
				allFileList: map[string]string{
					"/tsconfigWithConfigDir.json": tsconfigWithConfigDir,
					"/bin/a.tlua":                 "",
					"/b.tlua":                     "",
				},
			},
		},
	},
	{
		title: "returns error when tsconfig have excludes",
		input: []testConfig{{
			jsonText: `{
                    "compilerOptions": {
                        "lib": ["es5"]
                    },
                    "excludes": [
                        "foge.tlua"
                    ]
                }`,
			configFileName: "tluaconfig.json",
			basePath:       "/apath",
			allFileList:    map[string]string{"/apath/test.tlua": "", "/apath/foge.tlua": ""},
		}},
	},
	{
		title: "parses tsconfig with extends, files, include and other options",
		input: []testConfig{{
			jsonText: `{
				"extends": "./tsconfigWithExtends.json",
				"compilerOptions": {
				    "outDir": "./dist",
    				"strict": true,
    				"noImplicitAny": true,
				},
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/tsconfigWithExtends.json": tsconfigWithExtends, "/src/index.tlua": "", "/src/app.tlua": "", "/node_modules/module.tlua": "", "/dist/output.lua": ""},
		}},
	},
	{
		title: "parses tsconfig with extends and configDir",
		input: []testConfig{{
			jsonText: `{
				"extends": "./tluaconfig.base.json"
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/tluaconfig.base.json": tsconfigWithExtendsAndConfigDir, "/src/index.tlua": "", "/src/app.tlua": "", "/node_modules/module.tlua": "", "/dist/output.lua": ""},
		}},
	},
	{
		title: "reports error for an unknown option",
		input: []testConfig{{
			jsonText: `{
			    "compilerOptions": {
				"unknown": true
			    }
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/app.tlua": ""},
		}},
	},
	{
		title: "reports errors for wrong type option and invalid enum value",
		input: []testConfig{{
			jsonText: `{
			    "compilerOptions": {
				"target": "invalid value",
				"removeComments": "should be a boolean",
				"jsx": "invalid value"
			    }
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/app.tlua": ""},
		}},
	},
	{
		title: "reports errors for incorrectly cased option names",
		input: []testConfig{{
			jsonText: `{
			    "compilerOptions": {
				"sourcemap": true,
				"declarationmap": true,
				"nouncheckedindexedaccess": true,
				"exactoptionalpropertytypes": true,
				"verbatimmodulesyntax": true,
				"isolatedmodules": true,
				"nouncheckedsideeffectimports": true,
				"moduledetection": "force",
				"skiplibcheck": true,
				"checkjs": true
			    }
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/app.tlua": ""},
		}},
	},
	{
		title: "handles empty types array",
		input: []testConfig{{
			jsonText: `{
			    "compilerOptions": {
					"types": []
				}
			}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList:    map[string]string{"/app.tlua": ""},
		}},
	},
	{
		title: "issue 1267 scenario - extended files not picked up",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-base/backend.json",
  "compilerOptions": {
    "outDir": "dist",
    "rootDir": "src"
  },
  "exclude": ["node_modules", "dist"],
  "include": ["src/**/*"]
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-base/backend.json": `{
  "$schema": "https://json.schemastore.org/tsconfig",
  "display": "Backend",
  "compilerOptions": {
    "allowJs": true,
    "module": "nodenext",
    "removeComments": true,
    "emitDecoratorMetadata": true,
    "experimentalDecorators": true,
    "target": "esnext",
    "lib": ["ESNext"],
    "incremental": false,
    "noImplicitAny": true,
    "types": ["node", "vitest/globals"],
    "sourceMap": true
  },
  "files": [
    "types/ical2json.d.tlua",
    "types/express.d.tlua",
    "types/multer.d.tlua",
    "types/reset.d.tlua",
    "types/stripe-custom-typings.d.tlua",
    "types/nestjs-modules.d.tlua",
    "types/luxon.d.tlua",
    "types/nestjs-pino.d.tlua"
  ],
  "ts-node": {
    "files": true
  }
}`,
				"/tluaconfig-base/types/ical2json.d.tlua":             "export {}",
				"/tluaconfig-base/types/express.d.tlua":               "export {}",
				"/tluaconfig-base/types/multer.d.tlua":                "export {}",
				"/tluaconfig-base/types/reset.d.tlua":                 "export {}",
				"/tluaconfig-base/types/stripe-custom-typings.d.tlua": "export {}",
				"/tluaconfig-base/types/nestjs-modules.d.tlua":        "export {}",
				"/tluaconfig-base/types/luxon.d.tlua": `declare module 'luxon' {
  interface TSSettings {
    throwOnInvalid: true
  }
}
export {}`,
				"/tluaconfig-base/types/nestjs-pino.d.tlua": "export {}",
				"/src/main.tlua":  "export {}",
				"/src/utils.tlua": "export {}",
			},
		}},
	},
	{
		title: "null overrides in extended tsconfig - array fields",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": null,
    "lib": null,
    "typeRoots": null
  }
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-base.json": `{
  "compilerOptions": {
    "types": ["node", "@types/jest"],
    "lib": ["es2020", "dom"],
    "typeRoots": ["./types", "./node_modules/@types"]
  }
}`,
				"/app.tlua": "",
			},
		}},
	},
	{
		title: "null overrides in extended tsconfig - string fields",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "outDir": null,
    "rootDir": null
  }
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-base.json": `{
  "compilerOptions": {
    "outDir": "./dist",
    "rootDir": "./src"
  }
}`,
				"/app.tlua": "",
			},
		}},
	},
	{
		title: "null overrides in extended tsconfig - mixed field types",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": null,
    "outDir": null,
    "strict": false,
    "lib": ["es2022"],
    "allowJs": null
  }
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-base.json": `{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020", "dom"],
    "outDir": "./dist",
    "strict": true,
    "allowJs": true,
    "target": "es2020"
  }
}`,
				"/app.tlua": "",
			},
		}},
	},
	{
		title: "null overrides with multiple extends levels",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-middle.json",
  "compilerOptions": {
    "types": null,
    "lib": null
  }
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-middle.json": `{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": ["jest"],
    "outDir": "./build"
  }
}`,
				"/tluaconfig-base.json": `{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020"],
    "outDir": "./dist",
    "strict": true
  }
}`,
				"/app.tlua": "",
			},
		}},
	},
	{
		title: "null overrides in middle level of extends chain",
		input: []testConfig{{
			jsonText: `{
  "extends": "./tluaconfig-middle.json",
  "compilerOptions": {
    "outDir": "./final"
  }
}`,
			configFileName: "tluaconfig.json",
			basePath:       "/",
			allFileList: map[string]string{
				"/tluaconfig-middle.json": `{
  "extends": "./tluaconfig-base.json",
  "compilerOptions": {
    "types": null,
    "lib": null,
    "outDir": "./middle"
  }
}`,
				"/tluaconfig-base.json": `{
  "compilerOptions": {
    "types": ["node"],
    "lib": ["es2020"],
    "outDir": "./base",
    "strict": true
  }
}`,
				"/app.tlua": "",
			},
		}},
	},
}

var tsconfigWithExtends = `{
  "files": ["/src/index.tlua", "/src/app.tlua"],
  "include": ["/src/**/*"],
  "exclude": [],
  "ts-node": {
    "compilerOptions": {
      "module": "commonjs"
    },
    "transpileOnly": true
  }
}`

var tsconfigWithoutConfigDir = `{
  "compilerOptions": {
    "outDir": "bin"
  }
}`

var tsconfigWithConfigDir = `{
  "compilerOptions": {
    "outDir": "${configDir}/bin"
  }
}`

var tsconfigWithExtendsAndConfigDir = `{
  "compilerOptions": {
    "outFile": "${configDir}/outFile",
    "outDir": "${configDir}/outDir",
    "rootDir": "${configDir}/rootDir",
    "tsBuildInfoFile": "${configDir}/tsBuildInfoFile",
    "declarationDir": "${configDir}/declarationDir",
  }
}`

func TestParseJsonConfigFileContent(t *testing.T) {
	t.Parallel()
	for _, rec := range parseJsonConfigFileTests {
		t.Run(rec.title+" with json api", func(t *testing.T) {
			t.Parallel()
			baselineParseConfigWith(t, rec.title+" with json api.js", rec.input, getParsedWithJsonApi)
		})
	}
}

func getParsedWithJsonApi(config testConfig, host tsoptions.ParseConfigHost, basePath string) *tsoptions.ParsedCommandLine {
	configFileName := tspath.GetNormalizedAbsolutePath(config.configFileName, basePath)
	path := tspath.ToPath(config.configFileName, basePath, host.FS().UseCaseSensitiveFileNames())
	parsed, _ := tsoptions.ParseConfigFileTextToJson(configFileName, path, config.jsonText)
	return tsoptions.ParseJsonConfigFileContent(
		parsed,
		host,
		basePath,
		nil,
		configFileName,
		/*resolutionStack*/ nil,
		/*extraFileExtensions*/ nil,
		/*extendedConfigCache*/ nil,
	)
}

func TestParseJsonSourceFileConfigFileContent(t *testing.T) {
	t.Parallel()
	for _, rec := range parseJsonConfigFileTests {
		t.Run(rec.title+" with jsonSourceFile api", func(t *testing.T) {
			t.Parallel()
			baselineParseConfigWith(t, rec.title+" with jsonSourceFile api.js", rec.input, getParsedWithJsonSourceFileApi)
		})
	}
}

func TestParseJsonSourceFileConfigFileContentReportsInvalidExtendedConfig(t *testing.T) {
	t.Parallel()
	files := map[string]string{
		"/project/tluaconfig.json": `{
  "extends": "./bad.json"
}`,
		// The parser recovers from this as object-like JSON, producing expected-token
		// errors for ':', ',', and ':' (JSON members never take the shorthand
		// path, so the trailing `json` identifier is a member name still
		// awaiting its colon). `not` is tlua's word spelling of `!`, but not
		// of the postfix `!`, so it is not consumed after the property name
		// `is` and the ',' is reported at `not` rather than at `json`.
		"/project/bad.json":  "{ this is not json",
		"/project/main.tlua": "export local x = 1;",
	}
	host := tsoptionstest.NewVFSParseConfigHost(files, "/project", true /*useCaseSensitiveFileNames*/)
	configFileName := "/project/tluaconfig.json"
	configFile := tsoptions.NewTsconfigSourceFileFromFilePath(
		configFileName,
		tspath.ToPath(configFileName, host.GetCurrentDirectory(), host.FS().UseCaseSensitiveFileNames()),
		files[configFileName],
	)

	parsed := tsoptions.ParseJsonSourceFileConfigFileContent(
		configFile,
		host,
		host.GetCurrentDirectory(),
		nil,
		nil,
		configFileName,
		nil,
		nil,
		nil,
	)

	parseErrors := core.Filter(parsed.Errors, func(diagnostic *ast.Diagnostic) bool {
		return diagnostic.Code() == diagnostics.X_0_expected.Code()
	})
	expectedParseErrorMessages := []string{":", ",", ":"}
	expectedParseErrorPositions := []int{7, 10, 18}
	assert.Equal(t, len(expectedParseErrorMessages), len(parseErrors))
	assert.DeepEqual(t, core.Map(parseErrors, func(diagnostic *ast.Diagnostic) string {
		return diagnostic.MessageArgs()[0]
	}), expectedParseErrorMessages)
	assert.DeepEqual(t, core.Map(parseErrors, (*ast.Diagnostic).Pos), expectedParseErrorPositions)
	for _, diagnostic := range parseErrors {
		assert.Equal(t, diagnostic.File().FileName(), "/project/bad.json")
	}
}

// Extending an empty config file used to panic on nil Statements (#4265).
func TestParseJsonSourceFileConfigFileContentWithEmptyExtendedConfig(t *testing.T) {
	t.Parallel()
	files := map[string]string{
		"/project/tluaconfig.json": `{
  "extends": "./base.json"
}`,
		"/project/base.json": "",
		"/project/main.tlua": "export local x = 1;",
	}
	host := tsoptionstest.NewVFSParseConfigHost(files, "/project", true /*useCaseSensitiveFileNames*/)
	configFileName := "/project/tluaconfig.json"
	configFile := tsoptions.NewTsconfigSourceFileFromFilePath(
		configFileName,
		tspath.ToPath(configFileName, host.GetCurrentDirectory(), host.FS().UseCaseSensitiveFileNames()),
		files[configFileName],
	)

	parsed := tsoptions.ParseJsonSourceFileConfigFileContent(
		configFile,
		host,
		host.GetCurrentDirectory(),
		nil,
		nil,
		configFileName,
		nil,
		nil,
		nil,
	)

	assert.Assert(t, parsed != nil)
	assert.DeepEqual(t, parsed.FileNames(), []string{"/project/main.tlua"})
}

func TestParseJsonSourceFileConfigFileContentDoesNotDuplicateUnquotedKeyDiagnostics(t *testing.T) {
	t.Parallel()
	parsed := tsoptionstest.GetParsedCommandLine(t, `{
  compilerOptions: {
    strict: true
  }
}`, map[string]string{"/main.tlua": "export local x = 1;"}, "/", true /*useCaseSensitiveFileNames*/)

	diags := parsed.GetConfigFileParsingDiagnostics()
	assert.Equal(t, len(diags), 2)
	expectedLocations := []struct {
		line      int
		character int
	}{
		{line: 1, character: 2},
		{line: 2, character: 4},
	}
	for index, diagnostic := range diags {
		assert.Equal(t, diagnostic.Code(), diagnostics.String_literal_with_double_quotes_expected.Code())
		line, character := scanner.GetECMALineAndUTF16CharacterOfPosition(diagnostic.File(), diagnostic.Pos())
		assert.Equal(t, line, expectedLocations[index].line)
		assert.Equal(t, int(character), expectedLocations[index].character)
	}
}

func TestParseJsonSourceFileConfigFileContentReportsQuestionTokenDiagnostics(t *testing.T) {
	t.Parallel()
	parsed := tsoptionstest.GetParsedCommandLine(t, `{
  compilerOptions?: {
    strict?: true
  }
}`, map[string]string{"/main.tlua": "export local x = 1;"}, "/", true /*useCaseSensitiveFileNames*/)

	var questionTokenDiagnostics []*ast.Diagnostic
	for _, diagnostic := range parsed.GetConfigFileParsingDiagnostics() {
		if diagnostic.Code() == diagnostics.The_0_modifier_can_only_be_used_in_tlua_files.Code() {
			questionTokenDiagnostics = append(questionTokenDiagnostics, diagnostic)
		}
	}
	assert.Equal(t, len(questionTokenDiagnostics), 2)
	expectedLocations := []struct {
		line      int
		character int
	}{
		{line: 1, character: 17},
		{line: 2, character: 10},
	}
	for index, diagnostic := range questionTokenDiagnostics {
		line, character := scanner.GetECMALineAndUTF16CharacterOfPosition(diagnostic.File(), diagnostic.Pos())
		assert.Equal(t, line, expectedLocations[index].line)
		assert.Equal(t, int(character), expectedLocations[index].character)
	}
}

func TestParseNullEnumCompilerOptions(t *testing.T) {
	t.Parallel()

	config := testConfig{
		jsonText: `{
			"compilerOptions": {
				"target": null,
				"module": null
			}
		}`,
		configFileName: "tluaconfig.json",
		basePath:       "/",
		allFileList:    map[string]string{"/app.tlua": ""},
	}
	for name, getParsed := range map[string]func(testConfig, tsoptions.ParseConfigHost, string) *tsoptions.ParsedCommandLine{
		"json api":           getParsedWithJsonApi,
		"jsonSourceFile api": getParsedWithJsonSourceFileApi,
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			allFileLists := make(map[string]string, len(config.allFileList)+1)
			maps.Copy(allFileLists, config.allFileList)
			allFileLists["/tluaconfig.json"] = config.jsonText
			host := tsoptionstest.NewVFSParseConfigHost(allFileLists, config.basePath, true /*useCaseSensitiveFileNames*/)
			parsedConfigFileContent := getParsed(config, host, config.basePath)
			assert.Equal(t, len(parsedConfigFileContent.Errors), 0)
		})
	}
}

func getParsedWithJsonSourceFileApi(config testConfig, host tsoptions.ParseConfigHost, basePath string) *tsoptions.ParsedCommandLine {
	configFileName := tspath.GetNormalizedAbsolutePath(config.configFileName, basePath)
	path := tspath.ToPath(config.configFileName, basePath, host.FS().UseCaseSensitiveFileNames())
	parsed := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: configFileName,
		Path:     path,
	}, config.jsonText, core.ScriptKindJSON)
	tsConfigSourceFile := &tsoptions.TsConfigSourceFile{
		SourceFile: parsed,
	}
	return tsoptions.ParseJsonSourceFileConfigFileContent(
		tsConfigSourceFile,
		host,
		host.GetCurrentDirectory(),
		nil,
		nil,
		configFileName,
		/*resolutionStack*/ nil,
		/*extraFileExtensions*/ nil,
		/*extendedConfigCache*/ nil,
	)
}

func baselineParseConfigWith(t *testing.T, baselineFileName string, input []testConfig, getParsed func(config testConfig, host tsoptions.ParseConfigHost, basePath string) *tsoptions.ParsedCommandLine) {
	var baselineContent strings.Builder
	for i, config := range input {
		basePath := config.basePath
		if basePath == "" {
			basePath = tspath.GetNormalizedAbsolutePath(tspath.GetDirectoryPath(config.configFileName), "")
		}
		configFileName := tspath.CombinePaths(basePath, config.configFileName)
		allFileLists := make(map[string]string, len(config.allFileList)+1)
		maps.Copy(allFileLists, config.allFileList)
		allFileLists[configFileName] = config.jsonText
		host := tsoptionstest.NewVFSParseConfigHost(allFileLists, config.basePath, true /*useCaseSensitiveFileNames*/)
		parsedConfigFileContent := getParsed(config, host, basePath)

		baselineContent.WriteString("Fs::\n")
		if err := printFS(&baselineContent, host.FS(), "/"); err != nil {
			t.Fatal(err)
		}
		baselineContent.WriteString("\n")
		baselineContent.WriteString("configFileName:: ")
		baselineContent.WriteString(config.configFileName)
		baselineContent.WriteString("\n")
		baselineContent.WriteString("CompilerOptions::\n")
		assert.NilError(t, json.MarshalIndentWrite(&baselineContent, parsedConfigFileContent.ParsedConfig.CompilerOptions, "", "  "))
		baselineContent.WriteString("\n")
		baselineContent.WriteString("\n")

		if parsedConfigFileContent.ParsedConfig.TypeAcquisition != nil {
			baselineContent.WriteString("TypeAcquisition::\n")
			assert.NilError(t, json.MarshalIndentWrite(&baselineContent, parsedConfigFileContent.ParsedConfig.TypeAcquisition, "", "  "))
			baselineContent.WriteString("\n")
			baselineContent.WriteString("\n")
		}
		baselineContent.WriteString("FileNames::\n")
		baselineContent.WriteString(strings.Join(parsedConfigFileContent.ParsedConfig.FileNames, ","))
		baselineContent.WriteString("\n")
		baselineContent.WriteString("Errors::\n")
		diagnosticwriter.FormatDiagnosticsWithColorAndContext(&baselineContent, diagnosticwriter.FromASTDiagnostics(parsedConfigFileContent.Errors), &diagnosticwriter.FormattingOptions{
			NewLine: "\r\n",
			ComparePathsOptions: tspath.ComparePathsOptions{
				CurrentDirectory:          basePath,
				UseCaseSensitiveFileNames: true,
			},
		})
		baselineContent.WriteString("\n")
		if i != len(input)-1 {
			baselineContent.WriteString("\n")
		}
	}
	// Always a local baseline: file enumeration diverges from upstream by
	// design now that sources are .tlua rather than .ts.
	baseline.Run(t, baselineFileName, baselineContent.String(), baseline.Options{Subfolder: "config/tsconfigParsing"})
}

func writeJsonReadableText(output io.Writer, input any) error {
	return json.MarshalIndentWrite(output, input, "", "  ")
}

func TestParseTypeAcquisition(t *testing.T) {
	t.Parallel()
	cases := []struct {
		title      string
		configName string
		config     string
	}{
		{
			title: "Convert correctly format tluaconfig.json to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enable": true,
		"include": ["0.d.tlua", "1.d.tlua"],
		"exclude": ["0.lua", "1.lua"],
	},
}`,
			configName: "tluaconfig.json",
		},
		{
			title: "Convert incorrect format tluaconfig.json to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enableAutoDiscovy": true,
	}
}`, configName: "tluaconfig.json",
		},
		{
			title:  "Convert default tluaconfig.json to typeAcquisition ",
			config: `{}`, configName: "tluaconfig.json",
		},
		{
			title: "Convert tluaconfig.json with only enable property to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enable": true,
	},
}`, configName: "tluaconfig.json",
		},

		// jsconfig.json
		{
			title: "Convert jsconfig.json to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enable": false,
		"include": ["0.d.tlua"],
		"exclude": ["0.lua"],
	},
}`,
			configName: "jsconfig.json",
		},
		{title: "Convert default jsconfig.json to typeAcquisition ", config: `{}`, configName: "jsconfig.json"},
		{
			title: "Convert incorrect format jsconfig.json to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enableAutoDiscovy": true,
	},
}`,
			configName: "jsconfig.json",
		},
		{
			title: "Convert jsconfig.json with only enable property to typeAcquisition ",
			config: `{
	"typeAcquisition": {
		"enable": false,
	},
}`,
			configName: "jsconfig.json",
		},
	}
	for _, test := range cases {
		withJsonApiName := test.title + " with json api"
		input := []testConfig{
			{
				jsonText:       test.config,
				configFileName: test.configName,
				basePath:       "/apath",
				allFileList: map[string]string{
					"/apath/a.tlua": "",
					"/apath/b.tlua": "",
				},
			},
		}
		t.Run(withJsonApiName, func(t *testing.T) {
			t.Parallel()
			baselineParseConfigWith(t, withJsonApiName+".js", input, getParsedWithJsonApi)
		})
		withJsonSourceFileApiName := test.title + " with jsonSourceFile api"
		t.Run(withJsonSourceFileApiName, func(t *testing.T) {
			t.Parallel()
			baselineParseConfigWith(t, withJsonSourceFileApiName+".js", input, getParsedWithJsonSourceFileApi)
		})
	}
}

func printFS(output io.Writer, files vfs.FS, root string) error {
	return files.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Type().IsRegular() {
			if content, ok := files.ReadFile(path); !ok {
				return fmt.Errorf("failed to read file %s", path)
			} else {
				if _, err := fmt.Fprintf(output, "//// [%s]\r\n%s\r\n\r\n", path, content); err != nil {
					return err
				}
			}
		}
		return nil
	})
}

// BenchmarkParseTsconfigProject parses a representative tsconfig over a
// synthetic project tree, exercising option validation and include-glob
// file enumeration.
func BenchmarkParseTsconfigProject(b *testing.B) {
	compilerDir := "/project"
	tsconfigFileName := "/project/tluaconfig.json"
	jsonText := `{
    "compilerOptions": {
        "outDir": "../built/local",
        "rootDir": ".",
        "declaration": true,
        "sourceMap": true,
        "composite": true,
        "strict": true
    },
    "include": ["**/*.tlua"]
}`

	files := map[string]string{tsconfigFileName: jsonText}
	for i := range 100 {
		files[fmt.Sprintf("/project/dir%d/file%d.tlua", i%10, i)] = ""
	}
	host := tsoptionstest.NewVFSParseConfigHost(files, compilerDir, true /*useCaseSensitiveFileNames*/)

	tsconfigPath := tspath.ToPath(tsconfigFileName, compilerDir, host.FS().UseCaseSensitiveFileNames())
	parsed := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: tsconfigFileName,
		Path:     tsconfigPath,
	}, jsonText, core.ScriptKindJSON)
	assert.Equal(b, len(parsed.Diagnostics()), 0)

	b.ReportAllocs()

	for b.Loop() {
		tsoptions.ParseJsonSourceFileConfigFileContent(
			&tsoptions.TsConfigSourceFile{
				SourceFile: parsed,
			},
			host,
			host.GetCurrentDirectory(),
			nil,
			nil,
			tsconfigFileName,
			/*resolutionStack*/ nil,
			/*extraFileExtensions*/ nil,
			/*extendedConfigCache*/ nil,
		)
	}
}

// memoCache is a minimal memoizing ExtendedConfigCache used by tests to simulate
// cache hits across multiple parses of configs that extend a common base.
type memoCache struct {
	m map[tspath.Path]*tsoptions.ExtendedConfigCacheEntry
}

func (mc *memoCache) GetExtendedConfig(fileName string, path tspath.Path, resolutionStack []tspath.Path, host tsoptions.ParseConfigHost) *tsoptions.ExtendedConfigCacheEntry {
	if mc.m == nil {
		mc.m = make(map[tspath.Path]*tsoptions.ExtendedConfigCacheEntry)
	}
	if e, ok := mc.m[path]; ok {
		return e
	}
	e := tsoptions.ParseExtendedConfig(fileName, path, resolutionStack, host, mc)
	mc.m[path] = e
	return e
}

var _ tsoptions.ExtendedConfigCache = (*memoCache)(nil)

// TestExtendedConfigErrorsAppearOnCacheHit verifies that diagnostics produced while parsing an
// extended config are still reported when the extended config comes from the cache.
func TestExtendedConfigErrorsAppearOnCacheHit(t *testing.T) {
	t.Parallel()

	t.Run("single config parsed twice", func(t *testing.T) {
		t.Parallel()
		files := map[string]string{
			"/tluaconfig.json": `{
  "extends": "./base.json"
}`,
			// 'excludes' instead of 'exclude' triggers diagnostic
			"/base.json": `{
  "excludes": ["**/*.tlua"]
}`,
			"/app.tlua": "export {}",
		}

		host := tsoptionstest.NewVFSParseConfigHost(files, "/", true /*useCaseSensitiveFileNames*/)

		parseConfig := func(configFileName string, cache tsoptions.ExtendedConfigCache) *tsoptions.ParsedCommandLine {
			cfgPath := tspath.ToPath(configFileName, host.GetCurrentDirectory(), host.FS().UseCaseSensitiveFileNames())
			jsonText, ok := host.FS().ReadFile(configFileName)
			assert.Assert(t, ok, "missing %s in test fs", configFileName)
			tsConfigSourceFile := &tsoptions.TsConfigSourceFile{
				SourceFile: parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: configFileName, Path: cfgPath}, jsonText, core.ScriptKindJSON),
			}
			return tsoptions.ParseJsonSourceFileConfigFileContent(
				tsConfigSourceFile,
				host,
				host.GetCurrentDirectory(),
				nil,
				nil,
				configFileName,
				nil,
				nil,
				cache,
			)
		}

		cache := &memoCache{}
		first := parseConfig("/tluaconfig.json", cache)
		assert.Assert(t, len(first.Errors) > 0, "expected diagnostics on first parse, got 0")
		second := parseConfig("/tluaconfig.json", cache)
		assert.Assert(t, len(second.Errors) > 0, "expected diagnostics on second parse (cache hit), got 0")
	})

	t.Run("two configs share same base", func(t *testing.T) {
		t.Parallel()
		files := map[string]string{
			"/base.json": `{
  "excludes": ["**/*.tlua"]
}`,
			"/projA/tluaconfig.json": `{
  "extends": "../base.json"
}`,
			"/projB/tluaconfig.json": `{
  "extends": "../base.json"
}`,
			"/projA/app.tlua": "export {}",
			"/projB/app.tlua": "export {}",
		}

		host := tsoptionstest.NewVFSParseConfigHost(files, "/", true /*useCaseSensitiveFileNames*/)

		parseConfig := func(configFileName string, cache tsoptions.ExtendedConfigCache) *tsoptions.ParsedCommandLine {
			cfgPath := tspath.ToPath(configFileName, host.GetCurrentDirectory(), host.FS().UseCaseSensitiveFileNames())
			jsonText, ok := host.FS().ReadFile(configFileName)
			assert.Assert(t, ok, "missing %s in test fs", configFileName)
			tsConfigSourceFile := &tsoptions.TsConfigSourceFile{
				SourceFile: parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: configFileName, Path: cfgPath}, jsonText, core.ScriptKindJSON),
			}
			return tsoptions.ParseJsonSourceFileConfigFileContent(
				tsConfigSourceFile,
				host,
				host.GetCurrentDirectory(),
				nil,
				nil,
				configFileName,
				nil,
				nil,
				cache,
			)
		}

		cache := &memoCache{}
		first := parseConfig("/projA/tluaconfig.json", cache)
		assert.Assert(t, len(first.Errors) > 0, "expected diagnostics for projA parse, got 0")
		second := parseConfig("/projB/tluaconfig.json", cache)
		assert.Assert(t, len(second.Errors) > 0, "expected diagnostics for projB parse (cache hit on base), got 0")
	})
}

package compiler_test

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/bundled"
	"github.com/apyrr/tlua/internal/compiler"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/repo"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tspath"
	"github.com/apyrr/tlua/internal/vfs/osvfs"
	"github.com/apyrr/tlua/internal/vfs/vfstest"
	"gotest.tools/v3/assert"
)

type testFile struct {
	fileName string
	contents string
}

type programTest struct {
	testName      string
	files         []testFile
	expectedFiles []string
	target        core.ScriptTarget
}

var esnextLibs = []string{
	"lib.luajit.d.tlua",
}

var programTestCases = []programTest{
	{
		testName: "BasicFileOrdering",
		files: []testFile{
			{fileName: "c:/dev/src/index.tlua", contents: "/// <reference path='c:/dev/src2/a/5.tlua' />\n/// <reference path='c:/dev/src2/a/10.tlua' />"},
			{fileName: "c:/dev/src2/a/5.tlua", contents: "/// <reference path='4.tlua' />"},
			{fileName: "c:/dev/src2/a/4.tlua", contents: "/// <reference path='b/3.tlua' />"},
			{fileName: "c:/dev/src2/a/b/3.tlua", contents: "/// <reference path='2.tlua' />"},
			{fileName: "c:/dev/src2/a/b/2.tlua", contents: "/// <reference path='c/1.tlua' />"},
			{fileName: "c:/dev/src2/a/b/c/1.tlua", contents: "console.log('hello');"},
			{fileName: "c:/dev/src2/a/10.tlua", contents: "/// <reference path='b/c/d/9.tlua' />"},
			{fileName: "c:/dev/src2/a/b/c/d/9.tlua", contents: "/// <reference path='e/8.tlua' />"},
			{fileName: "c:/dev/src2/a/b/c/d/e/8.tlua", contents: "/// <reference path='7.tlua' />"},
			{fileName: "c:/dev/src2/a/b/c/d/e/7.tlua", contents: "/// <reference path='f/6.tlua' />"},
			{fileName: "c:/dev/src2/a/b/c/d/e/f/6.tlua", contents: "console.log('world!');"},
		},
		expectedFiles: slices.Concat(esnextLibs,
			[]string{
				"c:/dev/src2/a/b/c/1.tlua",
				"c:/dev/src2/a/b/2.tlua",
				"c:/dev/src2/a/b/3.tlua",
				"c:/dev/src2/a/4.tlua",
				"c:/dev/src2/a/5.tlua",
				"c:/dev/src2/a/b/c/d/e/f/6.tlua",
				"c:/dev/src2/a/b/c/d/e/7.tlua",
				"c:/dev/src2/a/b/c/d/e/8.tlua",
				"c:/dev/src2/a/b/c/d/9.tlua",
				"c:/dev/src2/a/10.tlua",
				"c:/dev/src/index.tlua",
			}),
		target: core.ScriptTargetESNext,
	},
	{
		testName: "FileOrderingImports",
		files: []testFile{
			{fileName: "c:/dev/src/index.tlua", contents: "local five = require('a.5');\nlocal ten = require('a.10');"},
			{fileName: "c:/dev/src/a/5.tlua", contents: "local four = require('a.4');"},
			{fileName: "c:/dev/src/a/4.tlua", contents: "local three = require('a.b.3');"},
			{fileName: "c:/dev/src/a/b/3.tlua", contents: "local two = require('a.b.2');"},
			{fileName: "c:/dev/src/a/b/2.tlua", contents: "local one = require('a.b.c.1');"},
			{fileName: "c:/dev/src/a/b/c/1.tlua", contents: "console.log('hello');"},
			{fileName: "c:/dev/src/a/10.tlua", contents: "local nine = require('a.b.c.d.9');"},
			{fileName: "c:/dev/src/a/b/c/d/9.tlua", contents: "local eight = require('a.b.c.d.e.8');"},
			{fileName: "c:/dev/src/a/b/c/d/e/8.tlua", contents: "local seven = require('a.b.c.d.e.7');"},
			{fileName: "c:/dev/src/a/b/c/d/e/7.tlua", contents: "local six = require('a.b.c.d.e.f.6');"},
			{fileName: "c:/dev/src/a/b/c/d/e/f/6.tlua", contents: "console.log('world!');"},
		},
		expectedFiles: slices.Concat(esnextLibs,
			[]string{
				"c:/dev/src/a/b/c/1.tlua",
				"c:/dev/src/a/b/2.tlua",
				"c:/dev/src/a/b/3.tlua",
				"c:/dev/src/a/4.tlua",
				"c:/dev/src/a/5.tlua",
				"c:/dev/src/a/b/c/d/e/f/6.tlua",
				"c:/dev/src/a/b/c/d/e/7.tlua",
				"c:/dev/src/a/b/c/d/e/8.tlua",
				"c:/dev/src/a/b/c/d/9.tlua",
				"c:/dev/src/a/10.tlua",
				"c:/dev/src/index.tlua",
			}),
		target: core.ScriptTargetESNext,
	},
	{
		testName: "FileOrderingCycles",
		files: []testFile{
			{fileName: "c:/dev/src/index.tlua", contents: "local five = require('a.5');\nlocal ten = require('a.10');"},
			{fileName: "c:/dev/src/a/5.tlua", contents: "local four = require('a.4');"},
			{fileName: "c:/dev/src/a/4.tlua", contents: "local three = require('a.b.3');"},
			{fileName: "c:/dev/src/a/b/3.tlua", contents: "local two = require('a.b.2');\nlocal cycle = require('index'); "},
			{fileName: "c:/dev/src/a/b/2.tlua", contents: "local one = require('a.b.c.1');"},
			{fileName: "c:/dev/src/a/b/c/1.tlua", contents: "console.log('hello');"},
			{fileName: "c:/dev/src/a/10.tlua", contents: "local nine = require('a.b.c.d.9');"},
			{fileName: "c:/dev/src/a/b/c/d/9.tlua", contents: "local eight = require('a.b.c.d.e.8');\nlocal cycle = require('index');"},
			{fileName: "c:/dev/src/a/b/c/d/e/8.tlua", contents: "local seven = require('a.b.c.d.e.7');"},
			{fileName: "c:/dev/src/a/b/c/d/e/7.tlua", contents: "local six = require('a.b.c.d.e.f.6');"},
			{fileName: "c:/dev/src/a/b/c/d/e/f/6.tlua", contents: "console.log('world!');"},
		},
		expectedFiles: slices.Concat(esnextLibs,
			[]string{
				"c:/dev/src/a/b/c/1.tlua",
				"c:/dev/src/a/b/2.tlua",
				"c:/dev/src/a/b/3.tlua",
				"c:/dev/src/a/4.tlua",
				"c:/dev/src/a/5.tlua",
				"c:/dev/src/a/b/c/d/e/f/6.tlua",
				"c:/dev/src/a/b/c/d/e/7.tlua",
				"c:/dev/src/a/b/c/d/e/8.tlua",
				"c:/dev/src/a/b/c/d/9.tlua",
				"c:/dev/src/a/10.tlua",
				"c:/dev/src/index.tlua",
			}),
		target: core.ScriptTargetESNext,
	},
}

func TestProgram(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		// Without embedding, we'd need to read all of the lib files out from disk into the MapFS.
		// Just skip this for now.
		t.Skip("bundled files are not embedded")
	}

	for _, testCase := range programTestCases {
		t.Run(testCase.testName, func(t *testing.T) {
			t.Parallel()
			libPrefix := bundled.LibPath() + "/"
			fs := vfstest.FromMap[any](nil, false /*useCaseSensitiveFileNames*/)
			fs = bundled.WrapFS(fs)

			for _, testFile := range testCase.files {
				_ = fs.WriteFile(testFile.fileName, testFile.contents)
			}

			opts := core.CompilerOptions{Target: testCase.target}

			program := compiler.NewProgram(compiler.ProgramOptions{
				Config: &tsoptions.ParsedCommandLine{
					ParsedConfig: &core.ParsedOptions{
						FileNames:       []string{"c:/dev/src/index.tlua"},
						CompilerOptions: &opts,
					},
				},
				Host: compiler.NewCompilerHost("c:/dev/src", fs, bundled.LibPath(), nil, nil),
			})

			actualFiles := []string{}
			for _, file := range program.GetSourceFiles() {
				actualFiles = append(actualFiles, strings.TrimPrefix(file.FileName(), libPrefix))
			}

			assert.DeepEqual(t, testCase.expectedFiles, actualFiles)
		})
	}
}

func TestIncludeProcessorDiagnosticsWithMissingFileCasing(t *testing.T) {
	t.Parallel()

	if !bundled.Embedded {
		t.Skip("bundled files are not embedded")
	}

	// Use case-sensitive file names so that /src/MyFile.tlua and /src/myFile.tlua
	// have different canonical paths but the same lower-case path, triggering
	// file casing diagnostics in the include processor.
	fs := vfstest.FromMap[any](nil, true /*useCaseSensitiveFileNames*/)
	fs = bundled.WrapFS(fs)

	// Only create the lowercase version; /src/MyFile.tlua does not exist.
	_ = fs.WriteFile("/src/myFile.tlua", `export local y = 2;`)

	opts := core.CompilerOptions{SkipDefaultLibCheck: core.TSTrue}

	// List both casings as root files. The first one (/src/MyFile.tlua) will fail
	// to load because it does not exist on the case-sensitive filesystem.
	program := compiler.NewProgram(compiler.ProgramOptions{
		Config: &tsoptions.ParsedCommandLine{
			ParsedConfig: &core.ParsedOptions{
				FileNames:       []string{"/src/MyFile.tlua", "/src/myFile.tlua"},
				CompilerOptions: &opts,
			},
		},
		Host: compiler.NewCompilerHost("/", fs, bundled.LibPath(), nil, nil),
	})

	// GetProgramDiagnostics triggers getDiagnostics which processes all
	// include processor diagnostics including the casing diagnostic whose
	// file path points to the missing /src/MyFile.tlua. Before the fix this
	// panicked with a nil pointer dereference.
	assert.NilError(t, func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("panic: %v", r)
			}
		}()
		program.GetProgramDiagnostics()
		return nil
	}())
}

func BenchmarkNewProgram(b *testing.B) {
	if !bundled.Embedded {
		// Without embedding, we'd need to read all of the lib files out from disk into the MapFS.
		// Just skip this for now.
		b.Skip("bundled files are not embedded")
	}

	for _, testCase := range programTestCases {
		b.Run(testCase.testName, func(b *testing.B) {
			fs := vfstest.FromMap[any](nil, false /*useCaseSensitiveFileNames*/)
			fs = bundled.WrapFS(fs)

			for _, testFile := range testCase.files {
				_ = fs.WriteFile(testFile.fileName, testFile.contents)
			}

			opts := core.CompilerOptions{Target: testCase.target}
			programOpts := compiler.ProgramOptions{
				Config: &tsoptions.ParsedCommandLine{
					ParsedConfig: &core.ParsedOptions{
						FileNames:       []string{"c:/dev/src/index.tlua"},
						CompilerOptions: &opts,
					},
				},
				Host: compiler.NewCompilerHost("c:/dev/src", fs, bundled.LibPath(), nil, nil),
			}

			for b.Loop() {
				compiler.NewProgram(programOpts)
			}
		})
	}

	b.Run("compiler", func(b *testing.B) {
		repo.SkipIfNoTypeScriptSubmodule(b)

		rootPath := tspath.NormalizeSlashes(filepath.Join(repo.TypeScriptSubmodulePath(), "src", "compiler"))

		fs := osvfs.FS()
		fs = bundled.WrapFS(fs)

		host := compiler.NewCompilerHost(rootPath, fs, bundled.LibPath(), nil, nil)

		parsed, errors := tsoptions.GetParsedCommandLineOfConfigFile(tspath.CombinePaths(rootPath, "tsconfig.json"), nil, nil, host, nil)
		assert.Equal(b, len(errors), 0, "Expected no errors in parsed command line")

		opts := compiler.ProgramOptions{
			Config: parsed,
			Host:   host,
		}

		for b.Loop() {
			compiler.NewProgram(opts)
		}
	})
}

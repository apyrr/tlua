package execute

import (
	"context"
	"fmt"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/compiler"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/execute/build"
	"github.com/apyrr/tlua/internal/execute/incremental"
	"github.com/apyrr/tlua/internal/execute/tlua"
	"github.com/apyrr/tlua/internal/format"
	"github.com/apyrr/tlua/internal/json"
	"github.com/apyrr/tlua/internal/locale"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/outputpaths"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/pprof"
	"github.com/apyrr/tlua/internal/tracing"
	"github.com/apyrr/tlua/internal/tsoptions"
	"github.com/apyrr/tlua/internal/tspath"
)

func startTracingIfNeeded(sys tlua.System, config *tsoptions.ParsedCommandLine, testing tlua.CommandLineTesting) *tracing.Tracing {
	traceDir := config.CompilerOptions().GenerateTrace
	if traceDir == "" {
		return nil
	}
	configFilePath := ""
	if config.ConfigFile != nil && config.ConfigFile.SourceFile != nil {
		configFilePath = config.ConfigFile.SourceFile.FileName()
	}
	tr, err := tracing.StartTracing(sys.FS(), traceDir, configFilePath, testing != nil)
	if err != nil {
		fmt.Fprintf(sys.Writer(), "Warning: Failed to start tracing: %v\n", err)
	}
	return tr
}

func stopTracing(sys tlua.System, tr *tracing.Tracing) {
	if tr == nil {
		return
	}
	if err := tr.StopTracing(); err != nil {
		fmt.Fprintf(sys.Writer(), "Warning: Failed to stop tracing: %v\n", err)
	}
}

func CommandLine(ctx context.Context, sys tlua.System, commandLineArgs []string, testing tlua.CommandLineTesting) tlua.CommandLineResult {
	if len(commandLineArgs) > 0 {
		switch strings.ToLower(commandLineArgs[0]) {
		case "-b", "--b", "-build", "--build":
			return tscBuildCompilation(ctx, sys, tsoptions.ParseBuildCommandLine(commandLineArgs, sys), testing)
			// case "-f":
			// 	return fmtMain(sys, commandLineArgs[1], commandLineArgs[1])
		}
	}

	return tscCompilation(ctx, sys, tsoptions.ParseCommandLine(commandLineArgs, sys), testing)
}

func fmtMain(sys tlua.System, input, output string) tlua.ExitStatus {
	ctx := format.WithFormatCodeSettings(context.Background(), lsutil.GetDefaultFormatCodeSettings(), "\n")
	input = string(tspath.ToPath(input, sys.GetCurrentDirectory(), sys.FS().UseCaseSensitiveFileNames()))
	output = string(tspath.ToPath(output, sys.GetCurrentDirectory(), sys.FS().UseCaseSensitiveFileNames()))
	fileContent, ok := sys.FS().ReadFile(input)
	if !ok {
		fmt.Fprintln(sys.Writer(), "File not found:", input)
		return tlua.ExitStatusNotImplemented
	}
	text := fileContent
	pathified := tspath.ToPath(input, sys.GetCurrentDirectory(), true)
	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: string(pathified),
		Path:     pathified,
	}, text, core.GetScriptKindFromFileName(string(pathified)))
	edits := format.FormatDocument(ctx, sourceFile)
	newText := core.ApplyBulkEdits(text, edits)

	if err := sys.FS().WriteFile(output, newText); err != nil {
		fmt.Fprintln(sys.Writer(), err.Error())
		return tlua.ExitStatusNotImplemented
	}
	return tlua.ExitStatusSuccess
}

func tscBuildCompilation(ctx context.Context, sys tlua.System, buildCommand *tsoptions.ParsedBuildCommandLine, testing tlua.CommandLineTesting) tlua.CommandLineResult {
	locale := buildCommand.Locale()
	reportDiagnostic := tlua.CreateDiagnosticReporter(sys, sys.Writer(), locale, buildCommand.CompilerOptions)

	if len(buildCommand.Errors) > 0 {
		for _, err := range buildCommand.Errors {
			reportDiagnostic(err)
		}
		return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
	}

	if pprofDir := buildCommand.CompilerOptions.PprofDir; pprofDir != "" {
		// !!! stderr?
		profileSession := pprof.BeginProfiling(pprofDir, sys.Writer())
		defer profileSession.Stop()
	}

	if buildCommand.CompilerOptions.Help.IsTrue() {
		tlua.PrintVersion(sys, locale)
		tlua.PrintBuildHelp(sys, locale, tsoptions.BuildOpts)
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess}
	}

	orchestrator := build.NewOrchestrator(build.Options{
		Sys:     sys,
		Command: buildCommand,
		Testing: testing,
	})
	return orchestrator.Start(ctx)
}

func tscCompilation(ctx context.Context, sys tlua.System, commandLine *tsoptions.ParsedCommandLine, testing tlua.CommandLineTesting) tlua.CommandLineResult {
	configFileName := ""
	locale := commandLine.Locale()
	reportDiagnostic := tlua.CreateDiagnosticReporter(sys, sys.Writer(), locale, commandLine.CompilerOptions())

	if len(commandLine.Errors) > 0 {
		for _, e := range commandLine.Errors {
			reportDiagnostic(e)
		}
		return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
	}

	if pprofDir := commandLine.CompilerOptions().PprofDir; pprofDir != "" {
		// !!! stderr?
		profileSession := pprof.BeginProfiling(pprofDir, sys.Writer())
		defer profileSession.Stop()
	}

	if commandLine.CompilerOptions().Init.IsTrue() {
		tlua.WriteConfigFile(sys, locale, reportDiagnostic, commandLine.Raw.(*collections.OrderedMap[string, any]))
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess}
	}

	if commandLine.CompilerOptions().Version.IsTrue() {
		tlua.PrintVersion(sys, locale)
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess}
	}

	if commandLine.CompilerOptions().Help.IsTrue() || commandLine.CompilerOptions().All.IsTrue() {
		tlua.PrintHelp(sys, locale, commandLine)
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess}
	}

	if commandLine.CompilerOptions().Watch.IsTrue() && commandLine.CompilerOptions().ListFilesOnly.IsTrue() {
		reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Options_0_and_1_cannot_be_combined, "watch", "listFilesOnly"))
		return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
	}

	if commandLine.CompilerOptions().Project != "" {
		if len(commandLine.FileNames()) != 0 {
			reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Option_project_cannot_be_mixed_with_source_files_on_a_command_line))
			return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
		}

		fileOrDirectory := tspath.NormalizePath(commandLine.CompilerOptions().Project)
		if sys.FS().DirectoryExists(fileOrDirectory) {
			configFileName = tspath.CombinePaths(fileOrDirectory, "tluaconfig.json")
			if !sys.FS().FileExists(configFileName) {
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Cannot_find_a_tluaconfig_json_file_at_the_current_directory_Colon_0, configFileName))
				return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
			}
		} else {
			configFileName = fileOrDirectory
			if !sys.FS().FileExists(configFileName) {
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.The_specified_path_does_not_exist_Colon_0, fileOrDirectory))
				return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
			}
		}
	} else if !commandLine.CompilerOptions().IgnoreConfig.IsTrue() || len(commandLine.FileNames()) == 0 {
		searchPath := tspath.NormalizePath(sys.GetCurrentDirectory())
		configFileName = findConfigFile(searchPath, sys.FS().FileExists, "tluaconfig.json")
		if len(commandLine.FileNames()) != 0 {
			if configFileName != "" {
				// Error to not specify config file
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.X_tluaconfig_json_is_present_but_will_not_be_loaded_if_files_are_specified_on_commandline_Use_ignoreConfig_to_skip_this_error))
				return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
			}
		} else if configFileName == "" {
			if commandLine.CompilerOptions().ShowConfig.IsTrue() {
				reportDiagnostic(ast.NewCompilerDiagnostic(diagnostics.Cannot_find_a_tluaconfig_json_file_at_the_current_directory_Colon_0, tspath.NormalizePath(sys.GetCurrentDirectory())))
			} else {
				tlua.PrintVersion(sys, locale)
				tlua.PrintHelp(sys, locale, commandLine)
			}
			return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsSkipped}
		}
	}

	// !!! convert to options with absolute paths is usually done here, but for ease of implementation, it's done in `tsoptions.ParseCommandLine()`
	compilerOptionsFromCommandLine := commandLine.CompilerOptions()
	configForCompilation := commandLine
	extendedConfigCache := &tlua.ExtendedConfigCache{}
	var compileTimes tlua.CompileTimes
	var commandLineRaw *collections.OrderedMap[string, any]
	if configFileName != "" {
		configStart := sys.Now()
		if raw, ok := commandLine.Raw.(*collections.OrderedMap[string, any]); ok {
			// Wrap command line options in a "compilerOptions" key to match tluaconfig.json structure
			wrapped := &collections.OrderedMap[string, any]{}
			wrapped.Set("compilerOptions", raw)
			commandLineRaw = wrapped
		}
		configParseResult, errors := tsoptions.GetParsedCommandLineOfConfigFile(configFileName, compilerOptionsFromCommandLine, commandLineRaw, sys, extendedConfigCache)
		compileTimes.ConfigTime = sys.Now().Sub(configStart)
		if len(errors) != 0 {
			// these are unrecoverable errors--exit to report them as diagnostics
			for _, e := range errors {
				reportDiagnostic(e)
			}
			return tlua.CommandLineResult{Status: tlua.ExitStatusDiagnosticsPresent_OutputsGenerated}
		}
		configForCompilation = configParseResult
		// Updater to reflect pretty
		reportDiagnostic = tlua.CreateDiagnosticReporter(sys, sys.Writer(), locale, commandLine.CompilerOptions())
	}

	reportErrorSummary := tlua.CreateReportErrorSummary(sys, locale, configForCompilation.CompilerOptions())
	if compilerOptionsFromCommandLine.ShowConfig.IsTrue() {
		showConfig(sys, configForCompilation, configFileName)
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess}
	}

	// Bare-file mode with no tsconfig: the Lua module search root must be where
	// the sources are, not where the process happens to run. Anchor it at the
	// command-line files' common parent directory so require("x") resolves the
	// sibling x.tlua regardless of the process cwd. (Applied after --showConfig
	// so that command only echoes the user's own options.)
	if configFileName == "" && len(configForCompilation.FileNames()) != 0 {
		options := configForCompilation.CompilerOptions()
		if options.RootDir == "" {
			options.RootDir = tspath.RemoveTrailingDirectorySeparator(outputpaths.GetComputedCommonSourceDirectory(
				configForCompilation.FileNames(),
				sys.GetCurrentDirectory(),
				sys.FS().UseCaseSensitiveFileNames(),
			))
		}
	}

	if configForCompilation.CompilerOptions().Watch.IsTrue() {
		watcher := createWatcher(
			sys,
			configForCompilation,
			compilerOptionsFromCommandLine,
			commandLineRaw,
			reportDiagnostic,
			reportErrorSummary,
			testing,
		)
		watcher.start(ctx)
		return tlua.CommandLineResult{Status: tlua.ExitStatusSuccess, Watcher: watcher}
	} else if configForCompilation.CompilerOptions().IsIncremental() {
		return performIncrementalCompilation(
			sys,
			configForCompilation,
			reportDiagnostic,
			reportErrorSummary,
			extendedConfigCache,
			&compileTimes,
			testing,
		)
	}
	return performCompilation(
		sys,
		configForCompilation,
		reportDiagnostic,
		reportErrorSummary,
		extendedConfigCache,
		&compileTimes,
		testing,
	)
}

func findConfigFile(searchPath string, fileExists func(string) bool, configName string) string {
	result, ok := tspath.ForEachAncestorDirectory(searchPath, func(ancestor string) (string, bool) {
		fullConfigName := tspath.CombinePaths(ancestor, configName)
		if fileExists(fullConfigName) {
			return fullConfigName, true
		}
		return fullConfigName, false
	})
	if !ok {
		return ""
	}
	return result
}

func getTraceFromSys(sys tlua.System, locale locale.Locale, testing tlua.CommandLineTesting) func(msg *diagnostics.Message, args ...any) {
	return tlua.GetTraceWithWriterFromSys(sys.Writer(), locale, testing)
}

func performIncrementalCompilation(
	sys tlua.System,
	config *tsoptions.ParsedCommandLine,
	reportDiagnostic tlua.DiagnosticReporter,
	reportErrorSummary tlua.DiagnosticsReporter,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	compileTimes *tlua.CompileTimes,
	testing tlua.CommandLineTesting,
) tlua.CommandLineResult {
	host := compiler.NewCachedFSCompilerHost(sys.GetCurrentDirectory(), sys.FS(), sys.DefaultLibraryPath(), extendedConfigCache, getTraceFromSys(sys, config.Locale(), testing))
	buildInfoReadStart := sys.Now()
	oldProgram := incremental.ReadBuildInfoProgram(config, incremental.NewBuildInfoReader(host), host)
	compileTimes.BuildInfoReadTime = sys.Now().Sub(buildInfoReadStart)

	tr := startTracingIfNeeded(sys, config, testing)

	parseStart := sys.Now()
	program := compiler.NewProgram(compiler.ProgramOptions{
		Config:  config,
		Host:    host,
		Tracing: tr,
	})
	compileTimes.ParseTime = sys.Now().Sub(parseStart)
	changesComputeStart := sys.Now()
	incrementalProgram := incremental.NewProgram(program, oldProgram, incremental.CreateHost(host), testing != nil)
	compileTimes.ChangesComputeTime = sys.Now().Sub(changesComputeStart)
	result, _ := tlua.EmitAndReportStatistics(tlua.EmitInput{
		Sys:                sys,
		ProgramLike:        incrementalProgram,
		Program:            incrementalProgram.GetProgram(),
		Config:             config,
		ReportDiagnostic:   reportDiagnostic,
		ReportErrorSummary: reportErrorSummary,
		Writer:             sys.Writer(),
		CompileTimes:       compileTimes,
		Testing:            testing,
		Tracing:            tr,
	})

	stopTracing(sys, tr)

	if testing != nil {
		testing.OnProgram(incrementalProgram)
	}
	return tlua.CommandLineResult{
		Status: result.Status,
	}
}

func performCompilation(
	sys tlua.System,
	config *tsoptions.ParsedCommandLine,
	reportDiagnostic tlua.DiagnosticReporter,
	reportErrorSummary tlua.DiagnosticsReporter,
	extendedConfigCache tsoptions.ExtendedConfigCache,
	compileTimes *tlua.CompileTimes,
	testing tlua.CommandLineTesting,
) tlua.CommandLineResult {
	host := compiler.NewCachedFSCompilerHost(sys.GetCurrentDirectory(), sys.FS(), sys.DefaultLibraryPath(), extendedConfigCache, getTraceFromSys(sys, config.Locale(), testing))

	tr := startTracingIfNeeded(sys, config, testing)

	parseStart := sys.Now()
	program := compiler.NewProgram(compiler.ProgramOptions{
		Config:  config,
		Host:    host,
		Tracing: tr,
	})
	compileTimes.ParseTime = sys.Now().Sub(parseStart)
	result, _ := tlua.EmitAndReportStatistics(tlua.EmitInput{
		Sys:                sys,
		ProgramLike:        program,
		Program:            program,
		Config:             config,
		ReportDiagnostic:   reportDiagnostic,
		ReportErrorSummary: reportErrorSummary,
		Writer:             sys.Writer(),
		CompileTimes:       compileTimes,
		Testing:            testing,
		Tracing:            tr,
	})

	stopTracing(sys, tr)

	return tlua.CommandLineResult{
		Status: result.Status,
	}
}

func showConfig(sys tlua.System, config *tsoptions.ParsedCommandLine, configFileName string) {
	tsConfig := tsoptions.ConvertToTSConfig(config, configFileName)
	_ = json.MarshalIndentWrite(sys.Writer(), tsConfig, "", "    ")
}

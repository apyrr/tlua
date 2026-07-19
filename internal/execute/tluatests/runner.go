package tluatests

import (
	"context"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/execute"
	"github.com/apyrr/tlua/internal/execute/tlua"
	"github.com/apyrr/tlua/internal/testutil/baseline"
	"github.com/apyrr/tlua/internal/tspath"
)

type tluaEdit struct {
	caption         string
	commandLineArgs []string
	edit            func(*TestSys)
	expectedDiff    string
}

var noChange = &tluaEdit{
	caption: "no change",
}

var noChangeOnlyEdit = []*tluaEdit{
	noChange,
}

// Declaration emit for a Lua module always fails with TS100054. That error is
// produced at emit time and is not persisted in the buildinfo, so an
// up-to-date incremental build and a from-scratch build report it at
// different points; edits where that asymmetry surfaces carry this
// explanation.
const declarationEmitErrorReplayDiff = "TS100054 declaration emit errors are emit-time diagnostics not stored in buildinfo, so incremental and clean builds report them at different times"

// Module signatures are derived from declaration emit, which TS100054
// suppresses for Lua modules, so an incremental build cannot see that a
// dependency's exported shape changed and skips re-checking dependents; a
// clean build reports the propagated errors.
const signaturePropagationDiff = "incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything"

// noChangeWithDeclarationEmitErrorReplayDiff is a no-change edit for tests
// where the TS100054 replay asymmetry shows up on the no-change build.
var noChangeWithDeclarationEmitErrorReplayDiff = &tluaEdit{
	caption:      "no change",
	expectedDiff: declarationEmitErrorReplayDiff,
}

type tluaInput struct {
	subScenario      string
	commandLineArgs  []string
	files            FileMap
	cwd              string
	edits            []*tluaEdit
	env              map[string]string
	ignoreCase       bool
	windowsStyleRoot string
}

func (test *tluaInput) executeCommand(sys *TestSys, baselineBuilder *strings.Builder, commandLineArgs []string) tlua.CommandLineResult {
	fmt.Fprint(baselineBuilder, "tlua ", strings.Join(commandLineArgs, " "), "\n")
	result := execute.CommandLine(context.Background(), sys, commandLineArgs, sys)
	switch result.Status {
	case tlua.ExitStatusSuccess:
		baselineBuilder.WriteString("ExitStatus:: Success")
	case tlua.ExitStatusDiagnosticsPresent_OutputsSkipped:
		baselineBuilder.WriteString("ExitStatus:: DiagnosticsPresent_OutputsSkipped")
	case tlua.ExitStatusDiagnosticsPresent_OutputsGenerated:
		baselineBuilder.WriteString("ExitStatus:: DiagnosticsPresent_OutputsGenerated")
	case tlua.ExitStatusInvalidProject_OutputsSkipped:
		baselineBuilder.WriteString("ExitStatus:: InvalidProject_OutputsSkipped")
	case tlua.ExitStatusProjectReferenceCycle_OutputsSkipped:
		baselineBuilder.WriteString("ExitStatus:: ProjectReferenceCycle_OutputsSkipped")
	case tlua.ExitStatusNotImplemented:
		baselineBuilder.WriteString("ExitStatus:: NotImplemented")
	default:
		panic(fmt.Sprintf("UnknownExitStatus %d", result.Status))
	}
	return result
}

func (test *tluaInput) run(t *testing.T, scenario string) {
	t.Helper()
	t.Run(test.getBaselineSubFolder()+"/"+test.subScenario, func(t *testing.T) {
		t.Parallel()
		// initial test tlua compile
		baselineBuilder := &strings.Builder{}
		sys := newTestSys(test, false)
		fmt.Fprint(
			baselineBuilder,
			"currentDirectory::",
			sys.GetCurrentDirectory(),
			"\nuseCaseSensitiveFileNames::",
			sys.FS().UseCaseSensitiveFileNames(),
			"\nInput::\n",
		)
		sys.baselineFSwithDiff(baselineBuilder)
		result := test.executeCommand(sys, baselineBuilder, test.commandLineArgs)
		sys.serializeState(baselineBuilder)
		if result.Watcher != nil && sys.mockWatchBackend.HasWatches() {
			baselineBuilder.WriteString(sys.mockWatchBackend.WatchState())
		}
		var unexpectedDiff strings.Builder
		unexpectedDiff.WriteString(sys.baselinePrograms(baselineBuilder, "Initial build"))

		for index, do := range test.edits {
			sys.clearOutput()
			wg := core.NewWorkGroup(false)
			var nonIncrementalSys *TestSys
			commandLineArgs := core.IfElse(do.commandLineArgs == nil, test.commandLineArgs, do.commandLineArgs)
			wg.Queue(func() {
				baselineBuilder.WriteString(fmt.Sprintf("\n\nEdit [%d]:: %s\n", index, do.caption))
				if do.edit != nil {
					do.edit(sys)
				}
				changedPaths := sys.fsDiffer.ChangedPaths()
				sys.baselineFSwithDiff(baselineBuilder)

				if result.Watcher == nil {
					test.executeCommand(sys, baselineBuilder, commandLineArgs)
				} else {
					sys.mockWatchBackend.SendChangedPaths(changedPaths)
					result.Watcher.DoCycle()
				}
				sys.serializeState(baselineBuilder)
				if result.Watcher != nil && sys.mockWatchBackend.HasWatches() {
					baselineBuilder.WriteString(sys.mockWatchBackend.WatchState())
				}
				unexpectedDiff.WriteString(sys.baselinePrograms(baselineBuilder, fmt.Sprintf("Edit [%d]:: %s\n", index, do.caption)))
			})
			wg.Queue(func() {
				// Compute build with all the edits
				nonIncrementalSys = newTestSys(test, true)
				for i := range index + 1 {
					if test.edits[i].edit != nil {
						test.edits[i].edit(nonIncrementalSys)
					}
				}
				execute.CommandLine(context.Background(), nonIncrementalSys, commandLineArgs, nonIncrementalSys)
			})
			wg.RunAndWait()

			diff := getDiffForIncremental(sys, nonIncrementalSys)
			if diff != "" {
				baselineBuilder.WriteString(fmt.Sprintf("\n\nDiff:: %s\n", core.IfElse(do.expectedDiff == "", "!!! Unexpected diff, please review and either fix or write explanation as expectedDiff !!!", do.expectedDiff)))
				baselineBuilder.WriteString(diff)
				if do.expectedDiff == "" {
					unexpectedDiff.WriteString(fmt.Sprintf("Edit [%d]:: %s\n!!! Unexpected diff, please review and either fix or write explanation as expectedDiff !!!\n%s\n", index, do.caption, diff))
				}
			} else if do.expectedDiff != "" {
				baselineBuilder.WriteString(fmt.Sprintf("\n\nDiff:: %s !!! Diff not found but explanation present, please review and remove the explanation !!!\n", do.expectedDiff))
				unexpectedDiff.WriteString(fmt.Sprintf("Edit [%d]:: %s\n!!! Diff not found but explanation present, please review and remove the explanation !!!\n", index, do.caption))
			}
		}
		baseline.Run(t, strings.ReplaceAll(test.subScenario, " ", "-")+".js", baselineBuilder.String(), baseline.Options{Subfolder: filepath.Join(test.getBaselineSubFolder(), scenario)})
		if unexpectedDiff.String() != "" {
			t.Errorf("Test %s has unexpected diff %s with incremental build, please review the baseline file", test.subScenario, unexpectedDiff.String())
		}
	})
}

func getDiffForIncremental(incrementalSys *TestSys, nonIncrementalSys *TestSys) string {
	var diffBuilder strings.Builder

	nonIncrementalOutputs := nonIncrementalSys.fs.writtenFiles.ToSlice()
	slices.Sort(nonIncrementalOutputs)
	for _, nonIncrementalOutput := range nonIncrementalOutputs {
		if tspath.FileExtensionIs(nonIncrementalOutput, tspath.ExtensionTsBuildInfo) ||
			strings.HasSuffix(nonIncrementalOutput, ".readable.baseline.txt") {
			// Just check existence
			if !incrementalSys.fsFromFileMap().FileExists(nonIncrementalOutput) {
				diffBuilder.WriteString(baseline.DiffText("nonIncremental "+nonIncrementalOutput, "incremental "+nonIncrementalOutput, "Exists", ""))
				diffBuilder.WriteString("\n")
			}
		} else {
			nonIncrementalText, ok := nonIncrementalSys.fsFromFileMap().ReadFile(nonIncrementalOutput)
			if !ok {
				panic("Written file not found " + nonIncrementalOutput)
			}
			incrementalText, ok := incrementalSys.fsFromFileMap().ReadFile(nonIncrementalOutput)
			if !ok || incrementalText != nonIncrementalText {
				diffBuilder.WriteString(baseline.DiffText("nonIncremental "+nonIncrementalOutput, "incremental "+nonIncrementalOutput, nonIncrementalText, incrementalText))
				diffBuilder.WriteString("\n")
			}
		}
	}

	incrementalOutput := incrementalSys.getOutput(true)
	nonIncrementalOutput := nonIncrementalSys.getOutput(true)
	if incrementalOutput != nonIncrementalOutput {
		diffBuilder.WriteString(baseline.DiffText("nonIncremental.output.txt", "incremental.output.txt", nonIncrementalOutput, incrementalOutput))
	}
	return diffBuilder.String()
}

func (test *tluaInput) getBaselineSubFolder() string {
	commandName := "tlua"
	if slices.ContainsFunc(test.commandLineArgs, func(arg string) bool {
		switch arg {
		case "-b", "--b", "-build", "--build":
			return true
		}
		return false
	}) {
		commandName = "tluabuild"
	}
	w := ""
	if slices.ContainsFunc(test.commandLineArgs, func(arg string) bool {
		switch arg {
		case "-w", "--w", "-watch", "--watch":
			return true
		}
		return false
	}) {
		w = "Watch"
	}
	return commandName + w
}

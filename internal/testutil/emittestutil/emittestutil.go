package emittestutil

import (
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/testutil/parsetestutil"
	"gotest.tools/v3/assert"
)

// Checks that pretty-printing the given file matches the expected output.
func CheckEmit(t *testing.T, emitContext *printer.EmitContext, file *ast.SourceFile, expected string) {
	t.Helper()
	text := checkEmitWorker(t, emitContext, file, expected)
	file2 := parsetestutil.ParseTypeScript(text, file.LanguageVariant == core.LanguageVariantJSX)
	parsetestutil.CheckDiagnosticsMessage(t, file2, "error on reparse: ")
}

// CheckEmitJS is CheckEmit without the reparse gate, for trees that print
// JS-only syntax. tlua has no conditional expression, but the downlevel
// lowerings (`?.`, `??`, destructuring defaults) still synthesize a ternary
// into their .js output, which tlua source syntax cannot read back.
func CheckEmitJS(t *testing.T, emitContext *printer.EmitContext, file *ast.SourceFile, expected string) {
	t.Helper()
	checkEmitWorker(t, emitContext, file, expected)
}

func checkEmitWorker(t *testing.T, emitContext *printer.EmitContext, file *ast.SourceFile, expected string) string {
	t.Helper()
	printer := printer.NewPrinter(
		printer.PrinterOptions{
			NewLine: core.NewLineKindLF,
		},
		printer.PrintHandlers{},
		emitContext,
	)
	text := printer.EmitSourceFile(file)
	actual := strings.TrimSuffix(text, "\n")
	assert.Equal(t, expected, actual)
	return text
}

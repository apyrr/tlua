package parser_test

import (
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/scanner"
	"gotest.tools/v3/assert"
)

// `null` is a JSON literal but an ordinary name in tlua source, so the scanner
// only reads it as the nil keyword in JSON mode. The language service and the
// formatter rescan file text independently of the parser; they have to scan a
// JSON file the same way, or their tokens disagree with the tree.
func TestJSONNullRescansAsParsed(t *testing.T) {
	t.Parallel()
	text := `{ "extends": null }`
	file := parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: "/tluaconfig.json", Path: "/tluaconfig.json"}, text, core.ScriptKindJSON)
	assert.Equal(t, scanner.ScanTokenAtPosition(file, strings.Index(text, "null")), ast.KindNilKeyword)

	source := `local extends = null`
	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: "/a.tlua", Path: "/a.tlua"}, source, core.ScriptKindTS)
	assert.Equal(t, scanner.ScanTokenAtPosition(sourceFile, strings.Index(source, "null")), ast.KindIdentifier)
}

// `nil` is tlua's spelling, not JSON's: in a JSON file it is a bare word, and
// JSON validation rejects it the way it rejects `undefined`.
func TestJSONNilIsNotALiteral(t *testing.T) {
	t.Parallel()
	text := `{ "outDir": nil }`
	file := parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: "/tluaconfig.json", Path: "/tluaconfig.json"}, text, core.ScriptKindJSON)
	assert.Equal(t, scanner.ScanTokenAtPosition(file, strings.Index(text, "nil")), ast.KindIdentifier)
	assert.Equal(t, len(file.Diagnostics()), 1)
	assert.Equal(t, file.Diagnostics()[0].Code(), diagnostics.Property_value_can_only_be_string_literal_numeric_literal_true_false_null_object_literal_or_array_literal.Code())
}

// JSON `null` parses to the nil keyword, which tlua prints as `nil`; emitting a
// JSON file (outDir copies it) must print it back as `null`.
func TestJSONNullPrintsAsNull(t *testing.T) {
	t.Parallel()
	text := `{ "a": null, "b": [null, 1] }`
	file := parser.ParseSourceFile(ast.SourceFileParseOptions{FileName: "/data.json", Path: "/data.json"}, text, core.ScriptKindJSON)
	assert.Equal(t, len(file.Diagnostics()), 0)
	p := printer.NewPrinter(printer.PrinterOptions{NewLine: core.NewLineKindLF}, printer.PrintHandlers{}, printer.NewEmitContext())
	assert.Equal(t, strings.TrimSpace(p.EmitSourceFile(file)), text)
}

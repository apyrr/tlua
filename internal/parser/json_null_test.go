package parser_test

import (
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/parser"
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

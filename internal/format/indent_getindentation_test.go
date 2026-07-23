package format_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/format"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/parser"
)

func TestGetIndentationForObjectTypeMemberPosition(t *testing.T) {
	t.Parallel()

	text := "local x: {\n    field: number,\n} = a;"
	// The member line "    field: number," starts after the "{\n" at position 11.

	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/test.tlua",
		Path:     "/test.tlua",
	}, text, core.ScriptKindTS)

	options := lsutil.GetDefaultFormatCodeSettings()

	// getAdjustedStartPosition with LeadingTriviaOptionNone returns the line start;
	// position 14 is inside the "    field" indentation of the object-type member.
	lineStart := format.GetLineStartPositionForPosition(14, sourceFile)

	indent := format.GetIndentation(lineStart, sourceFile, options, true)
	t.Logf("lineStart=%d, text[lineStart:]=%q", lineStart, text[lineStart:lineStart+10])
	t.Logf("GetIndentation at lineStart %d = %d", lineStart, indent)

	if indent != 4 {
		t.Errorf("Expected indentation 4, got %d", indent)
	}
}

package format_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/format"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/parser"
	"gotest.tools/v3/assert"
)

// The word-spelled operators share token kinds with `&&`/`||`/`!`, so the
// delete-space rules keyed on those kinds must not fire against the word form:
// removing the space in `a and b` yields the single identifier `aandb`, which
// no longer parses. The punctuation twins keep their usual treatment.
func TestFormatWordOperatorsKeepTheirSpaces(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  string
		output string
	}{
		// insertSpaceBeforeAndAfterBinaryOperators: false deletes around
		// punctuation operators (including the declaration `=`) but must leave
		// the words alone.
		{"and keeps spaces", "local x = a and b;", "local x=a and b;"},
		{"or keeps spaces", "local x = a or b;", "local x=a or b;"},
		{"and-or chain keeps spaces", "local x = a and b or c;", "local x=a and b or c;"},
		{"punctuation still tightens", "local x = a + b;", "local x=a+b;"},
		// The prefix rule deletes after `!` but must not merge `not` with its
		// operand.
		{"not keeps its space", "local y = not a;", "local y=not a;"},
		{"bang still tightens", "local y = ! a;", "local y=!a;"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ctx := format.WithFormatCodeSettings(t.Context(), lsutil.FormatCodeSettings{
				EditorSettings: lsutil.EditorSettings{
					TabSize:             4,
					IndentSize:          4,
					NewLineCharacter:    "\n",
					ConvertTabsToSpaces: core.TSTrue,
					IndentStyle:         lsutil.IndentStyleSmart,
				},
				InsertSpaceBeforeAndAfterBinaryOperators: core.TSFalse,
			}, "\n")
			sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
				FileName: "/test.tlua",
				Path:     "/test.tlua",
			}, tc.input, core.ScriptKindTS)
			edits := format.FormatDocument(ctx, sourceFile)
			assert.Equal(t, tc.output, applyBulkEdits(tc.input, edits))
		})
	}
}

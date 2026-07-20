package format_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/format"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/repo"
	"gotest.tools/v3/assert"
)

func applyBulkEdits(text string, edits []core.TextChange) string {
	b := strings.Builder{}
	b.Grow(len(text))
	lastEnd := 0
	for _, e := range edits {
		start := e.TextRange.Pos()
		if start != lastEnd {
			b.WriteString(text[lastEnd:e.TextRange.Pos()])
		}
		b.WriteString(e.NewText)

		lastEnd = e.TextRange.End()
	}
	b.WriteString(text[lastEnd:])

	return b.String()
}

func TestFormat(t *testing.T) {
	t.Parallel()

	t.Run("format lib.luajit.d.tlua", func(t *testing.T) {
		t.Parallel()
		ctx := format.WithFormatCodeSettings(t.Context(), lsutil.FormatCodeSettings{
			EditorSettings: lsutil.EditorSettings{
				TabSize:                4,
				IndentSize:             4,
				BaseIndentSize:         4,
				NewLineCharacter:       "\n",
				ConvertTabsToSpaces:    core.TSTrue,
				IndentStyle:            lsutil.IndentStyleSmart,
				TrimTrailingWhitespace: core.TSTrue,
			},
			InsertSpaceBeforeTypeAnnotation: core.TSTrue,
		}, "\n")
		// The upstream checker.ts corpus went away with the TypeScript
		// submodule; the largest bundled lib is the in-repo replacement. A
		// deliberately misformatted prefix keeps the must-edit assertion
		// below independent of whether the shipped lib happens to be
		// formatter-clean.
		text := "local   unformatted=1;\n" + readFormatCorpus(t)
		sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
			FileName: "/lib.luajit.d.tlua",
			Path:     "/lib.luajit.d.tlua",
		}, text, core.ScriptKindTS)
		edits := format.FormatDocument(ctx, sourceFile)
		// The corpus is known to need reformatting, so a silently inert
		// formatter (zero edits) must fail here, not pass by idempotency.
		assert.Assert(t, len(edits) > 0)
		newText := applyBulkEdits(text, edits)
		assert.Assert(t, len(newText) > 0)

		// Formatting must be idempotent: reformatting the formatted output
		// produces no further changes.
		reformattedFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
			FileName: "/lib.luajit.d.tlua",
			Path:     "/lib.luajit.d.tlua",
		}, newText, core.ScriptKindTS)
		reformatted := applyBulkEdits(newText, format.FormatDocument(ctx, reformattedFile))
		assert.Equal(t, reformatted, newText)
	})
}

// readFormatCorpus reads the large in-repo declaration file used as the
// formatting corpus.
func readFormatCorpus(tb testing.TB) string {
	tb.Helper()
	filePath := filepath.Join(repo.RootPath(), "internal", "bundled", "libs", "lib.luajit.d.tlua")
	fileContent, err := os.ReadFile(filePath)
	assert.NilError(tb, err)
	return string(fileContent)
}

func BenchmarkFormat(b *testing.B) {
	ctx := format.WithFormatCodeSettings(b.Context(), lsutil.FormatCodeSettings{
		EditorSettings: lsutil.EditorSettings{
			TabSize:                4,
			IndentSize:             4,
			BaseIndentSize:         4,
			NewLineCharacter:       "\n",
			ConvertTabsToSpaces:    core.TSTrue,
			IndentStyle:            lsutil.IndentStyleSmart,
			TrimTrailingWhitespace: core.TSTrue,
		},
		InsertSpaceBeforeTypeAnnotation: core.TSTrue,
	}, "\n")
	text := readFormatCorpus(b)
	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/lib.luajit.d.tlua",
		Path:     "/lib.luajit.d.tlua",
	}, text, core.ScriptKindTS)

	b.Run("format lib.luajit.d.tlua", func(b *testing.B) {
		for b.Loop() {
			edits := format.FormatDocument(ctx, sourceFile)
			newText := applyBulkEdits(text, edits)
			assert.Assert(b, len(newText) > 0)
		}
	})

	b.Run("format lib.luajit.d.tlua (no edit application)", func(b *testing.B) { // for comparison (how long does applying many edits take?)
		for b.Loop() {
			format.FormatDocument(ctx, sourceFile)
		}
	})

	p := printer.NewPrinter(printer.PrinterOptions{}, printer.PrintHandlers{}, printer.NewEmitContext())
	b.Run("pretty print lib.luajit.d.tlua", func(b *testing.B) { // for comparison
		for b.Loop() {
			newText := p.EmitSourceFile(sourceFile)
			assert.Assert(b, len(newText) > 0)
		}
	})
}

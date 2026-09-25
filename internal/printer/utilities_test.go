package printer

import (
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"gotest.tools/v3/assert"
)

func TestEscapeString(t *testing.T) {
	t.Parallel()
	data := []struct {
		s         string
		quoteChar QuoteChar
		expected  string
	}{
		{s: "", quoteChar: QuoteCharDoubleQuote, expected: ``},
		{s: "abc", quoteChar: QuoteCharDoubleQuote, expected: `abc`},
		{s: "ab\"c", quoteChar: QuoteCharDoubleQuote, expected: `ab\"c`},
		{s: "ab\tc", quoteChar: QuoteCharDoubleQuote, expected: `ab\tc`},
		{s: "ab\nc", quoteChar: QuoteCharDoubleQuote, expected: `ab\nc`},
		{s: "ab'c", quoteChar: QuoteCharDoubleQuote, expected: `ab'c`},
		{s: "ab'c", quoteChar: QuoteCharSingleQuote, expected: `ab\'c`},
		{s: "ab\"c", quoteChar: QuoteCharSingleQuote, expected: `ab"c`},
		{s: "ab`c", quoteChar: QuoteCharBacktick, expected: "ab\\`c"},
		// tlua: control chars use Lua's `\xHH`, never JS `\uXXXX`.
		{s: "\u001f", quoteChar: QuoteCharBacktick, expected: `\x1f`},
		{s: "a\x00b", quoteChar: QuoteCharDoubleQuote, expected: `a\x00b`},
		// `${` only starts a substitution inside a backtick template.
		{s: "a${b}", quoteChar: QuoteCharBacktick, expected: `a\${b}`},
		{s: "a${b}", quoteChar: QuoteCharDoubleQuote, expected: `a${b}`},
		// Non-ASCII (valid UTF-8, incl. astral) is written verbatim, not escaped.
		{s: "\u008f", quoteChar: QuoteCharDoubleQuote, expected: "\u008f"},
		{s: "𝟘𝟙", quoteChar: QuoteCharDoubleQuote, expected: "𝟘𝟙"},
		// A byte that is not valid UTF-8 reads back as itself through `\xHH`; bytes
		// that spell a surrogate are not valid UTF-8 either.
		{s: "\xffa", quoteChar: QuoteCharDoubleQuote, expected: `\xffa`},
		{s: "\xed\xa0\xbd", quoteChar: QuoteCharDoubleQuote, expected: `\xed\xa0\xbd`},
		// Unicode line breaks stay on one line.
		{s: "a\u2028b\u0085", quoteChar: QuoteCharDoubleQuote, expected: `a\xe2\x80\xa8b\xc2\x85`},
	}
	for i, rec := range data {
		t.Run(fmt.Sprintf("[%d] EscapeString(%q, %v)", i, rec.s, rec.quoteChar), func(t *testing.T) {
			t.Parallel()
			actual := EscapeString(rec.s, rec.quoteChar)
			assert.Equal(t, actual, rec.expected)
		})
	}
}

func TestIsRecognizedTripleSlashComment(t *testing.T) {
	t.Parallel()
	data := []struct {
		s            string
		commentRange ast.CommentRange
		expected     bool
	}{
		{s: "", commentRange: ast.CommentRange{Kind: ast.KindMultiLineCommentTrivia}, expected: false},
		{s: "", commentRange: ast.CommentRange{Kind: ast.KindSingleLineCommentTrivia}, expected: false},
		{s: "/a", expected: false},
		{s: "//", expected: false},
		{s: "//a", expected: false},
		{s: "///", expected: false},
		{s: "///a", expected: false},
		{s: "///<reference path=\"foo\" />", expected: true},
		{s: "///<reference types=\"foo\" />", expected: true},
		{s: "///<reference lib=\"foo\" />", expected: true},
		{s: "///<reference no-default-lib=\"foo\" />", expected: true},
		{s: "///<amd-dependency path=\"foo\" />", expected: true},
		{s: "///<amd-module />", expected: true},
		{s: "/// <reference path=\"foo\" />", expected: true},
		{s: "/// <reference types=\"foo\" />", expected: true},
		{s: "/// <reference lib=\"foo\" />", expected: true},
		{s: "/// <reference no-default-lib=\"foo\" />", expected: true},
		{s: "/// <amd-dependency path=\"foo\" />", expected: true},
		{s: "/// <amd-module />", expected: true},
		{s: "/// <reference path=\"foo\"/>", expected: true},
		{s: "/// <reference types=\"foo\"/>", expected: true},
		{s: "/// <reference lib=\"foo\"/>", expected: true},
		{s: "/// <reference no-default-lib=\"foo\"/>", expected: true},
		{s: "/// <amd-dependency path=\"foo\"/>", expected: true},
		{s: "/// <amd-module/>", expected: true},
		{s: "/// <reference path='foo' />", expected: true},
		{s: "/// <reference types='foo' />", expected: true},
		{s: "/// <reference lib='foo' />", expected: true},
		{s: "/// <reference no-default-lib='foo' />", expected: true},
		{s: "/// <amd-dependency path='foo' />", expected: true},
		{s: "/// <reference path=\"foo\" />  ", expected: true},
		{s: "/// <reference types=\"foo\" />  ", expected: true},
		{s: "/// <reference lib=\"foo\" />  ", expected: true},
		{s: "/// <reference no-default-lib=\"foo\" />  ", expected: true},
		{s: "/// <amd-dependency path=\"foo\" />  ", expected: true},
		{s: "/// <amd-module />  ", expected: true},
		{s: "/// <foo />", expected: false},
		{s: "/// <reference />", expected: false},
		{s: "/// <amd-dependency />", expected: false},
	}
	for i, rec := range data {
		t.Run(fmt.Sprintf("[%d] isRecognizedTripleSlashComment()", i), func(t *testing.T) {
			t.Parallel()
			commentRange := rec.commentRange
			if commentRange.Kind == ast.KindUnknown {
				commentRange.Kind = ast.KindSingleLineCommentTrivia
				commentRange.TextRange = core.NewTextRange(0, len(rec.s))
			}
			actual := IsRecognizedTripleSlashComment(rec.s, commentRange)
			assert.Equal(t, actual, rec.expected)
		})
	}
}

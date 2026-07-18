package printer

import (
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"gotest.tools/v3/assert"
)

func TestLuaEscapeString(t *testing.T) {
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
	}
	for i, rec := range data {
		t.Run(fmt.Sprintf("[%d] LuaEscapeString(%q, %v)", i, rec.s, rec.quoteChar), func(t *testing.T) {
			t.Parallel()
			actual := LuaEscapeString(rec.s, rec.quoteChar)
			assert.Equal(t, actual, rec.expected)
		})
	}
}

func TestEscapeJsxAttributeString(t *testing.T) {
	t.Parallel()
	data := []struct {
		s         string
		quoteChar QuoteChar
		expected  string
	}{
		{s: "", quoteChar: QuoteCharDoubleQuote, expected: ""},
		{s: "abc", quoteChar: QuoteCharDoubleQuote, expected: "abc"},
		{s: "ab\"c", quoteChar: QuoteCharDoubleQuote, expected: "ab&quot;c"},
		{s: "ab\tc", quoteChar: QuoteCharDoubleQuote, expected: "ab&#x9;c"},
		{s: "ab\nc", quoteChar: QuoteCharDoubleQuote, expected: "ab&#xA;c"},
		{s: "ab'c", quoteChar: QuoteCharDoubleQuote, expected: "ab'c"},
		{s: "ab'c", quoteChar: QuoteCharSingleQuote, expected: "ab&apos;c"},
		{s: "ab\"c", quoteChar: QuoteCharSingleQuote, expected: "ab\"c"},
		{s: "ab\u008fc", quoteChar: QuoteCharDoubleQuote, expected: "ab\u008Fc"},
		{s: "𝟘𝟙", quoteChar: QuoteCharDoubleQuote, expected: "𝟘𝟙"},
	}
	for i, rec := range data {
		t.Run(fmt.Sprintf("[%d] escapeJsxAttributeString(%q, %v)", i, rec.s, rec.quoteChar), func(t *testing.T) {
			t.Parallel()
			actual := escapeJsxAttributeString(rec.s, rec.quoteChar)
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

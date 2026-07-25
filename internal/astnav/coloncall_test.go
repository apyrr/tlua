package astnav_test

import (
	"strings"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/astnav"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/parser"
	"gotest.tools/v3/assert"
)

// The implicit `self` of a colon declaration wears the colon's range, which sits
// before the parameter list it was prepended to. GetTokenAtPosition binary-searches
// children by position, so an unmarked `self` made the list's first child start
// before the list itself: the search window inverted and every token from the open
// paren on became unreachable, answering with the whole declaration instead. That
// is brace matching, and any completion or signature help triggered exactly at `(`.
func TestGetTokenAtPositionInLuaMethodDeclaration(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input string
		token string
		kind  ast.Kind
	}{
		{"open paren after colon name", "function a.b:m( x: number) end", "(", ast.KindOpenParenToken},
		{"open paren, plain target", "function obj:m( x: number) end", "(", ast.KindOpenParenToken},
		// The dot form never had the problem; it is here so the two stay honest.
		{"open paren, dot form", "function a.b.m( x: number) end", "(", ast.KindOpenParenToken},
		{"parameter name", "function a.b:m( x: number) end", "x", ast.KindIdentifier},
		{"close paren", "function a.b:m( x: number) end", ")", ast.KindCloseParenToken},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
				FileName: "/test.tlua",
				Path:     "/test.tlua",
			}, tc.input, core.ScriptKindTS)

			pos := strings.Index(tc.input, tc.token)
			assert.Assert(t, pos >= 0)
			token := astnav.GetTokenAtPosition(sourceFile, pos)
			assert.Assert(t, token != nil)
			assert.Equal(t, token.Kind, tc.kind)
		})
	}
}

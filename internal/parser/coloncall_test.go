package parser_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/parser"
	"gotest.tools/v3/assert"
)

func parseLua(t *testing.T, input string) *ast.SourceFile {
	t.Helper()
	return parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/test.tlua",
		Path:     "/test.tlua",
	}, input, core.ScriptKindTS)
}

// A node ends where the current token begins, so the target of `function a.b:f()`
// has to be folded together while the colon is still unread. Finishing it
// afterwards stretched it over the colon, which put two children of the
// declaration on the same source -- an AST no walk in source order can follow.
func TestLuaMethodDeclarationTargetExcludesColon(t *testing.T) {
	t.Parallel()

	for _, input := range []string{
		"function obj:m() end",
		"function a.b:m() end",
		"function a.b.c:m() end",
	} {
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			decl := parseLua(t, input).Statements.Nodes[0].AsFunctionDeclaration()
			assert.Assert(t, decl.ColonToken != nil)
			assert.Equal(t, decl.Target.End(), decl.ColonToken.Pos())
		})
	}
}

// The implicit `self` is the declaration's first parameter, and it is marked
// reparsed because it wears the colon's range rather than source of its own. Every
// walker that reads the tree in source order keys on that flag, so losing it is how
// the formatter, the semantic-token collector and astnav all break at once.
func TestLuaMethodDeclarationSynthesizesReparsedSelfParameter(t *testing.T) {
	t.Parallel()

	decl := parseLua(t, "function a.b:m(x: number) end").Statements.Nodes[0].AsFunctionDeclaration()
	self := decl.Parameters.Nodes[0]
	assert.Equal(t, self.Name().Text(), "self")
	assert.Assert(t, self.Flags&ast.NodeFlagsReparsed != 0)
	// It borrows the colon, which is what makes the flag necessary.
	assert.Equal(t, self.Loc, decl.ColonToken.Loc)

	// The written parameter is real source and must stay visible.
	assert.Equal(t, decl.Parameters.Nodes[1].Name().Text(), "x")
	assert.Assert(t, decl.Parameters.Nodes[1].Flags&ast.NodeFlagsReparsed == 0)

	// A declaration written without the colon has no implicit parameter, so its
	// first written one must not be marked.
	dotted := parseLua(t, "function a.b.m(x: number) end").Statements.Nodes[0].AsFunctionDeclaration()
	assert.Assert(t, dotted.ColonToken == nil)
	assert.Assert(t, dotted.Parameters.Nodes[0].Flags&ast.NodeFlagsReparsed == 0)
}

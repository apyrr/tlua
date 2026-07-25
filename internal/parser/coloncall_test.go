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

// The implicit `self` is the declaration's first parameter, and the colon is what
// makes it implicit -- that pairing is what ast.LuaImplicitSelfParameter reports,
// and what every consumer that has to skip the parameter relies on.
func TestLuaMethodDeclarationSynthesizesSelfParameter(t *testing.T) {
	t.Parallel()

	decl := parseLua(t, "function a.b:m(x: number) end").Statements.Nodes[0].AsFunctionDeclaration()
	self := ast.LuaImplicitSelfParameter(decl)
	assert.Assert(t, self != nil)
	assert.Equal(t, self, decl.Parameters.Nodes[0])
	assert.Assert(t, ast.IsLuaImplicitSelfParameter(self))
	assert.Equal(t, self.Name().Text(), "self")
	// The written parameters keep their own identity behind it.
	assert.Assert(t, !ast.IsLuaImplicitSelfParameter(decl.Parameters.Nodes[1]))

	// A declaration written without the colon has no implicit parameter, so the
	// first written one must not be mistaken for it.
	dotted := parseLua(t, "function a.b.m(x: number) end").Statements.Nodes[0].AsFunctionDeclaration()
	assert.Assert(t, ast.LuaImplicitSelfParameter(dotted) == nil)
	assert.Assert(t, !ast.IsLuaImplicitSelfParameter(dotted.Parameters.Nodes[0]))
}

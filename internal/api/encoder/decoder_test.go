package encoder_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/apyrr/tlua/internal/api/encoder"
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/parser"
	"github.com/apyrr/tlua/internal/repo"
	"gotest.tools/v3/assert"
)

func parseSourceFile(code string) *ast.SourceFile {
	return parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/test.tlua",
		Path:     "/test.tlua",
	}, code, core.ScriptKindTS)
}

func TestDecodeSourceFile_Basic(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local x = 1;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)
	assert.Equal(t, decoded.AsNode().Kind, ast.KindSourceFile)
	assert.Equal(t, decoded.FileName(), "/test.tlua")
	assert.Equal(t, decoded.Text(), "local x = 1;")
	assert.Assert(t, decoded.Statements != nil)
	assert.Assert(t, decoded.EndOfFileToken != nil)
}

func TestDecodeSourceFile_Statements(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local a = 1;\nlocal b = 2;\nlocal c = 3;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)
	assert.Equal(t, len(decoded.Statements.Nodes), 3)
	for i, stmt := range decoded.Statements.Nodes {
		assert.Equal(t, stmt.Kind, ast.KindVariableStatement, "statement %d", i)
	}
}

func TestDecodeSourceFile_VariableDeclaration(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local x = 1;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	varStmt := decoded.Statements.Nodes[0].AsVariableStatement()
	assert.Assert(t, varStmt.DeclarationList != nil)
	declList := varStmt.DeclarationList.AsVariableDeclarationList()
	assert.Assert(t, declList.Declarations != nil)
	assert.Equal(t, len(declList.Declarations.Nodes), 1)

	decl := declList.Declarations.Nodes[0].AsVariableDeclaration()
	assert.Equal(t, decl.Name().Kind, ast.KindIdentifier)
	assert.Equal(t, decl.Name().AsIdentifier().Text, "x")
	assert.Assert(t, decl.Initializer != nil)
	assert.Equal(t, decl.Initializer.Kind, ast.KindNumericLiteral)
	assert.Equal(t, decl.Initializer.AsNumericLiteral().Text, "1")
}

func TestDecodeSourceFile_VariableDeclarationListFlags(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		code     string
		expected ast.NodeFlags
	}{
		{"local", "local x = 1;", ast.NodeFlagsLuaLocal},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			sf := parseSourceFile(tt.code)
			buf, _, err := encoder.EncodeSourceFile(sf)
			assert.NilError(t, err)

			decoded, err := encoder.DecodeSourceFile(buf)
			assert.NilError(t, err)

			declList := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList()
			got := declList.Flags & (ast.NodeFlagsLet | ast.NodeFlagsConst | ast.NodeFlagsLuaLocal)
			assert.Equal(t, got, tt.expected, "flags for %q: got %d, want %d", tt.code, got, tt.expected)
		})
	}
}

func TestDecodeSourceFile_FunctionDeclaration(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("function add(a: number, b: number): number { return a + b; }")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	funcDecl := decoded.Statements.Nodes[0].AsFunctionDeclaration()
	assert.Assert(t, funcDecl.Name() != nil)
	assert.Equal(t, funcDecl.Name().AsIdentifier().Text, "add")
	assert.Assert(t, funcDecl.Parameters != nil)
	assert.Equal(t, len(funcDecl.Parameters.Nodes), 2)
	assert.Assert(t, funcDecl.Type != nil)
	assert.Assert(t, funcDecl.Body != nil)

	param0 := funcDecl.Parameters.Nodes[0].AsParameterDeclaration()
	assert.Equal(t, param0.Name().AsIdentifier().Text, "a")
	assert.Assert(t, param0.Type != nil)
}

func TestDecodeSourceFile_IfStatement(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("if (true) { } else { }")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	ifStmt := decoded.Statements.Nodes[0].AsIfStatement()
	assert.Assert(t, ifStmt.Expression != nil)
	assert.Assert(t, ifStmt.ThenStatement != nil)
	assert.Assert(t, ifStmt.ElseStatement != nil)
	assert.Equal(t, ifStmt.ThenStatement.Kind, ast.KindBlock)
	assert.Equal(t, ifStmt.ElseStatement.Kind, ast.KindBlock)
}

func TestDecodeSourceFile_TemplateExpression(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local x = `hello ${name} world`;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	varDecl := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].AsVariableDeclaration()
	tmplExpr := varDecl.Initializer.AsTemplateExpression()
	assert.Assert(t, tmplExpr.Head != nil)
	assert.Equal(t, tmplExpr.Head.AsTemplateHead().Text, "hello ")
	assert.Assert(t, tmplExpr.TemplateSpans != nil)
	assert.Equal(t, len(tmplExpr.TemplateSpans.Nodes), 1)

	span := tmplExpr.TemplateSpans.Nodes[0].AsTemplateSpan()
	assert.Assert(t, span.Expression != nil)
	assert.Equal(t, span.Expression.Kind, ast.KindIdentifier)
	assert.Assert(t, span.Literal != nil)
	assert.Equal(t, span.Literal.AsTemplateTail().Text, " world")
}

func TestDecodeSourceFile_Positions(t *testing.T) {
	t.Parallel()
	code := "local x = 1;"
	sf := parseSourceFile(code)
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	assert.Equal(t, decoded.AsNode().Pos(), 0)
	assert.Equal(t, decoded.AsNode().End(), len(code))
}

func TestDecodeSourceFile_InterfaceDeclaration(t *testing.T) {
	t.Parallel()
	// classes are removed in tlua; round-trip an interface (same member-list shape)
	sf := parseSourceFile("interface Foo { bar(): void }")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	ifaceDecl := decoded.Statements.Nodes[0].AsInterfaceDeclaration()
	assert.Assert(t, ifaceDecl.Name() != nil)
	assert.Equal(t, ifaceDecl.Name().AsIdentifier().Text, "Foo")
	assert.Assert(t, ifaceDecl.Members != nil)
	assert.Equal(t, len(ifaceDecl.Members.Nodes), 1)
	assert.Equal(t, ifaceDecl.Members.Nodes[0].Kind, ast.KindMethodSignature)
}

func TestDecodeNodes_SubtreeRoundTrip(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("function greet(name: string) { return `Hello, ${name}!`; }")

	var funcNode *ast.Node
	visitor := &ast.NodeVisitor{}
	visitor.Visit = func(node *ast.Node) *ast.Node {
		if node.Kind == ast.KindFunctionDeclaration && funcNode == nil {
			funcNode = node
		}
		return node
	}
	visitor.VisitEachChild(sf.AsNode())
	assert.Assert(t, funcNode != nil)

	buf, _, err := encoder.EncodeNode(funcNode, sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeNodes(buf)
	assert.NilError(t, err)

	assert.Equal(t, decoded.Kind, ast.KindFunctionDeclaration)
	funcDecl := decoded.AsFunctionDeclaration()
	assert.Assert(t, funcDecl.Name() != nil)
	assert.Equal(t, funcDecl.Name().AsIdentifier().Text, "greet")
	assert.Assert(t, funcDecl.Parameters != nil)
	assert.Equal(t, len(funcDecl.Parameters.Nodes), 1)
	assert.Assert(t, funcDecl.Body != nil)
}

func TestDecodeSourceFile_BinaryExpression(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local x = 1 + 2;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	decl := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].AsVariableDeclaration()
	binExpr := decl.Initializer.AsBinaryExpression()
	assert.Assert(t, binExpr.Left != nil)
	assert.Assert(t, binExpr.Right != nil)
	assert.Assert(t, binExpr.OperatorToken != nil)
	assert.Equal(t, binExpr.Left.Kind, ast.KindNumericLiteral)
	assert.Equal(t, binExpr.Right.Kind, ast.KindNumericLiteral)
}

func TestDecodeSourceFile_KeywordExpressions(t *testing.T) {
	t.Parallel()
	// A keyword expression must decode as KeywordExpression, not Token, or the
	// printer panics. (`this` used to be the fixture here; it no longer parses.)
	sf := parseSourceFile("local x = true;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	// Navigate: const x = true -> VariableStatement -> declaration -> initializer
	decl := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].AsVariableDeclaration()
	keywordExpr := decl.Initializer
	assert.Equal(t, keywordExpr.Kind, ast.KindTrueKeyword)
	// This would panic if decoded as Token instead of KeywordExpression
	assert.Assert(t, keywordExpr.AsKeywordExpression() != nil)
}

func TestDecodeSourceFile_EmptyModuleBlock(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile(`declare global { }`)
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	// Navigate: declare global { } -> ModuleDeclaration -> ModuleBlock
	mod := decoded.Statements.Nodes[0].AsModuleDeclaration()
	assert.Assert(t, mod.Body != nil)
	block := mod.Body.AsModuleBlock()
	// Statements must be non-nil even when empty, otherwise the printer panics
	assert.Assert(t, block.Statements != nil)
	assert.Equal(t, len(block.Statements.Nodes), 0)
}

func TestDecodeSourceFile_EmptyBlockAndParams(t *testing.T) {
	t.Parallel()
	// Empty blocks and parameter lists must decode with non-nil NodeLists (not nil),
	// matching parser behavior. Previously the decoder left them nil, crashing the printer.
	sf := parseSourceFile("function foo() {}")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	funcDecl := decoded.Statements.Nodes[0].AsFunctionDeclaration()
	assert.Assert(t, funcDecl.Parameters != nil, "FunctionDeclaration.Parameters must be non-nil for foo()")
	assert.Equal(t, len(funcDecl.Parameters.Nodes), 0)
	assert.Assert(t, funcDecl.Body != nil)
	block := funcDecl.Body.AsBlock()
	assert.Assert(t, block.Statements != nil, "Block.Statements must be non-nil for empty blocks")
	assert.Equal(t, len(block.Statements.Nodes), 0)
}

func TestDecodeSourceFile_ArrowFunctionEmptyParams(t *testing.T) {
	t.Parallel()
	// `() => {}` must decode with non-nil Parameters (empty NodeList),
	// matching parser behavior. Previously the decoder left it nil, crashing the printer.
	sf := parseSourceFile("local f = () => {};")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	decl := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].AsVariableDeclaration()
	arrow := decl.Initializer.AsArrowFunction()
	assert.Assert(t, arrow.Parameters != nil, "ArrowFunction.Parameters must be non-nil for () => {}")
	assert.Equal(t, len(arrow.Parameters.Nodes), 0)
	assert.Assert(t, arrow.Body != nil)
	block := arrow.Body.AsBlock()
	assert.Assert(t, block.Statements != nil, "Block.Statements must be non-nil for empty body")
	assert.Equal(t, len(block.Statements.Nodes), 0)
}

func TestDecodeSourceFile_FunctionExpressionEmptyParams(t *testing.T) {
	t.Parallel()
	// `function() {}` must decode with non-nil Parameters (empty NodeList).
	sf := parseSourceFile("local f = function() {};")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	decl := decoded.Statements.Nodes[0].AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes[0].AsVariableDeclaration()
	funcExpr := decl.Initializer.AsFunctionExpression()
	assert.Assert(t, funcExpr.Parameters != nil, "FunctionExpression.Parameters must be non-nil for function() {}")
	assert.Equal(t, len(funcExpr.Parameters.Nodes), 0)
}

func TestDecodeSourceFile_PrefixUnaryOperator(t *testing.T) {
	t.Parallel()
	sf := parseSourceFile("local x = true; !x;")
	buf, _, err := encoder.EncodeSourceFile(sf)
	assert.NilError(t, err)

	decoded, err := encoder.DecodeSourceFile(buf)
	assert.NilError(t, err)

	exprStmt := decoded.Statements.Nodes[1].AsExpressionStatement()
	prefix := exprStmt.Expression.AsPrefixUnaryExpression()
	assert.Equal(t, prefix.Operator, ast.KindExclamationToken)
	assert.Equal(t, prefix.Operand.Kind, ast.KindIdentifier)
}

func BenchmarkDecodeSourceFile(b *testing.B) {
	repo.SkipIfNoTypeScriptSubmodule(b)
	filePath := filepath.Join(repo.TypeScriptSubmodulePath(), "src/compiler/checker.ts")
	fileContent, err := os.ReadFile(filePath)
	assert.NilError(b, err)
	code := string(fileContent)
	sourceFile := parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/checker.tlua",
		Path:     "/checker.tlua",
	}, code, core.ScriptKindTS)

	buf, _, err := encoder.EncodeSourceFile(sourceFile)
	assert.NilError(b, err)

	b.Run("parse", func(b *testing.B) {
		for b.Loop() {
			parser.ParseSourceFile(ast.SourceFileParseOptions{
				FileName: "/checker.tlua",
				Path:     "/checker.tlua",
			}, code, core.ScriptKindTS)
		}
	})

	b.Run("decode", func(b *testing.B) {
		for b.Loop() {
			_, decodeErr := encoder.DecodeSourceFile(buf)
			assert.NilError(b, decodeErr)
		}
	})
}

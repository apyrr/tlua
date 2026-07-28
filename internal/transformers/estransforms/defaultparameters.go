package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/transformers"
)

// luaDefaultParameterTransformer lowers parameter default values to Lua. Lua has
// no default parameters: `function f(x = 10)` is not Lua syntax at all, so the
// initializer cannot survive into the parameter list. It becomes a prologue
// assignment guarded on the parameter being absent:
//
//	function f(x)
//	  if x == nil then x = 10 end
//
// The guard is `x == nil`, not the idiomatic `x = x or 10`. Lua's only falsy
// values are `false` and `nil`, so the `or` spelling would wrongly replace an
// explicitly passed `false`. Testing against nil also matches the language rule
// the checker already applies: a parameter with an initializer is optional, and
// a missing argument arrives as nil.
//
// Prologue statements are emitted in parameter order, which is what makes
// defaults referring to earlier parameters (`function f(a, b = a)`) observe the
// already-defaulted value, and defaults referring to later ones observe nil.
//
// Signature-only declarations (overload signatures, ambient declarations) have no
// body to hold a prologue and emit nothing, so they are left alone.
//
// This runs unconditionally in the emit pipeline — Lua is always the target —
// rather than as part of the target-gated JS downlevel chain.
type luaDefaultParameterTransformer struct {
	transformers.Transformer
}

func (tx *luaDefaultParameterTransformer) visit(node *ast.Node) *ast.Node {
	node = tx.Visitor().VisitEachChild(node)
	if !ast.IsFunctionLike(node) || node.Body() == nil {
		return node
	}
	return tx.lower(node)
}

// lower strips initializers off the parameter list and rebuilds the body with a
// guarded assignment prepended for each one. It returns node unchanged when no
// parameter carries a default.
func (tx *luaDefaultParameterTransformer) lower(node *ast.Node) *ast.Node {
	params := node.Parameters()
	prologue := make([]*ast.Statement, 0, len(params))
	stripped := make([]*ast.Node, 0, len(params))
	changed := false

	for _, p := range params {
		param := p.AsParameterDeclaration()
		if param.Initializer == nil || !ast.IsIdentifier(param.Name()) {
			stripped = append(stripped, p)
			continue
		}
		changed = true
		prologue = append(prologue, tx.guardedAssignment(param.Name(), param.Initializer))
		stripped = append(stripped, tx.Factory().UpdateParameterDeclaration(
			param,
			param.Modifiers(),
			param.DotDotDotToken,
			param.Name(),
			param.QuestionToken,
			param.Type,
			nil, /*initializer*/
		))
	}
	if !changed {
		return node
	}

	body := node.Body()
	if !ast.IsBlock(body) {
		// A concise (expression) body has nowhere to put the prologue. Value
		// lambdas do not parse in tlua, so this is unreachable from source; leave
		// such a node alone rather than silently dropping the default.
		return node
	}
	block := body.AsBlock()
	newBody := tx.Factory().UpdateBlock(
		block,
		tx.Factory().NewNodeList(append(prologue, block.Statements.Nodes...)),
		block.MultiLine,
	)
	return tx.updateSignature(node, tx.Factory().NewNodeList(stripped), newBody)
}

// guardedAssignment builds `if <name> == nil then <name> = <initializer> end`.
func (tx *luaDefaultParameterTransformer) guardedAssignment(name *ast.Node, initializer *ast.Expression) *ast.Statement {
	f := tx.Factory()
	guard := f.NewBinaryExpression(
		nil, /*modifiers*/
		name.Clone(f),
		nil, /*typeNode*/
		f.NewToken(ast.KindEqualsEqualsToken),
		f.NewKeywordExpression(ast.KindNilKeyword),
	)
	assign := f.NewBinaryExpression(
		nil, /*modifiers*/
		name.Clone(f),
		nil, /*typeNode*/
		f.NewToken(ast.KindEqualsToken),
		initializer,
	)
	thenBlock := f.NewBlock(f.NewNodeList([]*ast.Node{f.NewExpressionStatement(assign)}), true /*multiLine*/)
	return f.NewIfStatement(guard, thenBlock, nil /*elseStatement*/)
}

// updateSignature rebuilds the one function-like form that reached us with the
// stripped parameter list and prologue-bearing body.
func (tx *luaDefaultParameterTransformer) updateSignature(node *ast.Node, params *ast.ParameterList, body *ast.Node) *ast.Node {
	f := tx.Factory()
	switch node.Kind {
	case ast.KindFunctionDeclaration:
		n := node.AsFunctionDeclaration()
		return f.UpdateFunctionDeclaration(n, n.Modifiers(), n.Target, n.ColonToken, n.Name(), n.TypeParameters, params, n.Type, n.FullSignature, body)
	case ast.KindFunctionExpression:
		n := node.AsFunctionExpression()
		return f.UpdateFunctionExpression(n, n.Modifiers(), n.Name(), n.TypeParameters, params, n.Type, n.FullSignature, body)
	case ast.KindArrowFunction:
		n := node.AsArrowFunction()
		return f.UpdateArrowFunction(n, n.Modifiers(), n.TypeParameters, params, n.Type, n.FullSignature, n.EqualsGreaterThanToken, body)
	}
	return node
}

// NewLuaDefaultParameterTransformer lowers parameter defaults to a nil-guarded
// prologue. It runs unconditionally in the emit pipeline (Lua is always the
// target), not as part of the target-gated JS downlevel chain.
func NewLuaDefaultParameterTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &luaDefaultParameterTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}

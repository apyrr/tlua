package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/transformers"
)

// luaOptionalChainTransformer lowers optional chaining (`?.`) to Lua. Lua has no
// `?.`, no ternary, and no assignment-expression, so the upstream JS downlevel
// (ternary + `(_a = e)` + `void 0`) cannot be reused. Instead:
//
//   - `base and access` guards the chain (Lua `and` returns the falsy operand,
//     and a member-bearing base can never be `false`, so truthiness == non-nil).
//   - side-effecting receivers are captured once into a preceding `local`.
//   - in a lazily-evaluated / non-order-safe position, the whole thing is wrapped
//     in a single-eval IIFE so nothing is hoisted out of its evaluation slot.
//   - a chain whose result is discarded (a bare `ExpressionStatement`) becomes
//     `if base then access end`, since `base and access` is not a Lua statement.
//
// This runs unconditionally (Lua is always the target), independent of the
// target-gated JS downlevel chain.
type luaOptionalChainTransformer struct {
	transformers.Transformer
	// preceding is the sink for hoisted `local _t = ...` statements of the
	// statement currently being lowered; nil when no hoistable statement is open.
	preceding *[]*ast.Statement
	// hoistTarget is the one outermost-chain node permitted to hoist into
	// `preceding` (the operative expression of an order-safe statement slot).
	hoistTarget *ast.Node
}

func (tx *luaOptionalChainTransformer) visit(node *ast.Node) *ast.Node {
	if node == nil || node.SubtreeFacts()&ast.SubtreeContainsOptionalChaining == 0 {
		return node
	}
	switch node.Kind {
	case ast.KindExpressionStatement:
		return tx.visitExpressionStatement(node.AsExpressionStatement())
	case ast.KindVariableStatement, ast.KindReturnStatement:
		return tx.visitHoistableStatement(node)
	case ast.KindPropertyAccessExpression, ast.KindElementAccessExpression, ast.KindCallExpression:
		if node.Flags&ast.NodeFlagsOptionalChain != 0 && !chainSpineHasTaggedTemplate(node) {
			return tx.placeChain(node)
		}
		return tx.Visitor().VisitEachChild(node)
	default:
		return tx.Visitor().VisitEachChild(node)
	}
}

// isOptionalChainExpr reports whether e is an outermost optional-chain access/call.
func isOptionalChainExpr(e *ast.Expression) bool {
	return e != nil && e.Flags&ast.NodeFlagsOptionalChain != 0 &&
		(e.Kind == ast.KindPropertyAccessExpression ||
			e.Kind == ast.KindElementAccessExpression ||
			e.Kind == ast.KindCallExpression)
}

// operativeChain returns the optional chain that is the whole operative
// expression of an order-safe statement slot (single-decl `local x = <chain>`
// or `return <chain>`), or nil. In these slots the chain base is the first thing
// the statement evaluates, so hoisting a `local` above it reorders nothing.
func operativeChain(stmt *ast.Node) *ast.Node {
	var e *ast.Expression
	switch stmt.Kind {
	case ast.KindReturnStatement:
		e = stmt.AsReturnStatement().Expression
	case ast.KindVariableStatement:
		list := stmt.AsVariableStatement().DeclarationList.AsVariableDeclarationList()
		if len(list.Declarations.Nodes) == 1 {
			e = list.Declarations.Nodes[0].AsVariableDeclaration().Initializer
		}
	}
	if e == nil {
		// No single operative expression: a multi-name `local a, b = ...`, a
		// bare `return`, or a `local` without an initializer.
		return nil
	}
	e = ast.SkipParentheses(e)
	if isOptionalChainExpr(e) {
		return e
	}
	return nil
}

// visitHoistableStatement opens a per-statement preceding-temp buffer and marks
// the operative chain (if any) hoistable, then splices any hoisted `local`s in
// front of the statement via a SyntaxList (flattened by VisitSlice).
func (tx *luaOptionalChainTransformer) visitHoistableStatement(node *ast.Node) *ast.Node {
	savedPre, savedTgt := tx.preceding, tx.hoistTarget
	buf := []*ast.Statement{}
	tx.preceding = &buf
	tx.hoistTarget = operativeChain(node)
	visited := tx.Visitor().VisitEachChild(node)
	tx.preceding, tx.hoistTarget = savedPre, savedTgt
	if len(buf) > 0 {
		return tx.Factory().NewSyntaxList(append(buf, visited))
	}
	return visited
}

// visitExpressionStatement rewrites a discarded chain `a?.b()` to
// `if a then a.b() end` (a bare `base and access` is not a legal Lua statement).
func (tx *luaOptionalChainTransformer) visitExpressionStatement(node *ast.ExpressionStatement) *ast.Node {
	inner := ast.SkipParentheses(node.Expression)
	if !isOptionalChainExpr(inner) || chainSpineHasTaggedTemplate(inner) {
		// The chain (if any) is nested in a subexpression, or is an invalid
		// tagged-template chain (TS1358); lower any inner chains in place.
		return tx.Visitor().VisitEachChild(node.AsNode())
	}
	savedPre, savedTgt := tx.preceding, tx.hoistTarget
	tx.preceding, tx.hoistTarget = nil, nil
	temps, guard, value := tx.lowerChain(inner)
	tx.preceding, tx.hoistTarget = savedPre, savedTgt

	// The discarded chain's base is the statement's first token, so its clone
	// carries the statement's leading comment. Suppress it on the guard/value
	// (which now sit inside the `if`) and let the synthesized `if` own the
	// original statement's position, so the comment prints once, ahead of `if`.
	tx.EmitContext().AddEmitFlags(guard, printer.EFNoComments)
	tx.EmitContext().AddEmitFlags(value, printer.EFNoComments)
	valueStmt := tx.Factory().NewExpressionStatement(value)
	thenBlock := tx.Factory().NewBlock(tx.Factory().NewNodeList([]*ast.Node{valueStmt}), true /*multiLine*/)
	thenBlock.Flags |= ast.NodeFlagsLuaBlock // `then ... end`, not braced
	ifStmt := tx.Factory().NewIfStatement(guard, thenBlock, nil)
	ifStmt.Loc = node.Loc
	tx.EmitContext().SetOriginal(ifStmt, node.AsNode())
	if len(temps) == 0 {
		return ifStmt
	}
	return tx.Factory().NewSyntaxList(append(temps, ifStmt))
}

// placeChain lowers an outermost optional chain and places the result according
// to its position: pure `and` (A), hoisted-temp `and` (B), or single-eval IIFE (C).
func (tx *luaOptionalChainTransformer) placeChain(node *ast.Node) *ast.Node {
	temps, guard, value := tx.lowerChain(node)
	expr := tx.Factory().NewLogicalANDExpression(guard, value)
	if len(temps) == 0 {
		return expr // (A) copiable base, no capture needed — valid in any position
	}
	if node == tx.hoistTarget && tx.preceding != nil {
		*tx.preceding = append(*tx.preceding, temps...)
		return expr // (B) order-safe slot — hoist the captures before the statement
	}
	// (C) lazy / non-order-safe slot — keep evaluation local inside an IIFE.
	stmts := append(temps, tx.Factory().NewReturnStatement(expr))
	return tx.Factory().NewImmediatelyInvokedArrowFunction(stmts)
}

// lowerChain lowers one optional-chain expression, returning the receiver-capture
// `local`s, the guard expression (truthy iff the receiver is present), and the
// full access built on the captured receiver. Only the receiver is ever
// duplicated (a plain identifier read); each non-copiable spine step is captured
// into its own `local`, since a Lua property read can fire an `__index`
// metamethod and must not be repeated.
func (tx *luaOptionalChainTransformer) lowerChain(node *ast.Node) (temps []*ast.Statement, guard *ast.Expression, value *ast.Expression) {
	r := flattenChain(node)
	baseNode := r.expression

	var baseExpr *ast.Expression
	if isOptionalChainExpr(baseNode) {
		// A nested `?.` link: lower it, then treat `g and v` as this chain's base.
		innerTemps, innerGuard, innerValue := tx.lowerChain(baseNode)
		temps = append(temps, innerTemps...)
		baseExpr = tx.Factory().NewLogicalANDExpression(innerGuard, innerValue)
	} else {
		// A plain base — visit so any chain nested inside it lowers too.
		baseExpr = tx.Visitor().VisitNode(baseNode)
	}

	// Capture the receiver unless it is trivially re-readable.
	var receiver *ast.Expression
	if transformers.IsSimpleCopiableExpression(baseExpr) {
		receiver = baseExpr
	} else {
		tmp := tx.Factory().NewTempVariable()
		temps = append(temps, tx.newLuaLocal(tmp, baseExpr))
		receiver = tmp
	}

	guard = tx.clone(receiver)
	value = tx.clone(receiver)
	for _, seg := range r.chain {
		value = tx.applySegment(value, seg)
	}
	return temps, guard, value
}

// applySegment rebuilds one chain link on top of expr, dropping its `?.` but
// preserving a Lua colon call's `:`.
func (tx *luaOptionalChainTransformer) applySegment(expr *ast.Expression, seg *ast.Node) *ast.Expression {
	switch seg.Kind {
	case ast.KindPropertyAccessExpression:
		p := seg.AsPropertyAccessExpression()
		return tx.Factory().NewPropertyAccessExpression(expr, nil /*questionDotToken*/, p.ColonToken, tx.Visitor().VisitNode(p.Name()), ast.NodeFlagsNone)
	case ast.KindElementAccessExpression:
		e := seg.AsElementAccessExpression()
		return tx.Factory().NewElementAccessExpression(expr, nil /*questionDotToken*/, tx.Visitor().VisitNode(e.ArgumentExpression), ast.NodeFlagsNone)
	case ast.KindCallExpression:
		c := seg.AsCallExpression()
		return tx.Factory().NewCallExpression(expr, nil /*questionDotToken*/, nil /*typeArguments*/, tx.Visitor().VisitNodes(c.Arguments), ast.NodeFlagsNone)
	default:
		panic("optional chain: unexpected segment kind")
	}
}

func (tx *luaOptionalChainTransformer) clone(node *ast.Node) *ast.Node {
	return node.Clone(tx.Factory())
}

// newLuaLocal builds `local <name> = <init>;`.
func (tx *luaOptionalChainTransformer) newLuaLocal(name *ast.IdentifierNode, init *ast.Expression) *ast.Statement {
	decl := tx.Factory().NewVariableDeclaration(name, nil /*exclamationToken*/, nil /*typeNode*/, init)
	list := tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{decl}), ast.NodeFlagsLuaLocal)
	return tx.Factory().NewVariableStatement(nil /*modifiers*/, list)
}

type flattenResult struct {
	expression *ast.Expression
	chain      []*ast.Node
}

func isNonNullChain(node *ast.Node) bool {
	return ast.IsNonNullExpression(node) && node.Flags&ast.NodeFlagsOptionalChain != 0
}

// chainSpineHasTaggedTemplate reports whether a tagged template sits in the
// access spine of node. A tagged template in an optional chain is a TS1358
// error, but emit still runs; `flattenChain` would then call `Node.Expression()`
// on the tagged template, which panics. Bailing out of lowering keeps the
// already-erroring program from crashing the compiler. The tagged-template kind
// is tested before descending, since its `Expression()` accessor is the panic.
func chainSpineHasTaggedTemplate(node *ast.Node) bool {
	for node != nil {
		if ast.IsTaggedTemplateExpression(node) {
			return true
		}
		switch node.Kind {
		case ast.KindPropertyAccessExpression,
			ast.KindElementAccessExpression,
			ast.KindCallExpression,
			ast.KindNonNullExpression,
			ast.KindParenthesizedExpression:
			node = node.Expression()
		default:
			return false
		}
	}
	return false
}

// flattenChain peels one optional-chain group: it returns the base expression
// before the group's first `?.`, and the links from that first optional access
// out to the outermost trailing access, in application order.
func flattenChain(chain *ast.Node) flattenResult {
	debug.Assert(!isNonNullChain(chain))
	links := []*ast.Node{chain}
	for !ast.IsTaggedTemplateExpression(chain) && chain.QuestionDotToken() == nil {
		chain = ast.SkipPartiallyEmittedExpressions(chain.Expression())
		debug.Assert(!isNonNullChain(chain))
		links = append([]*ast.Node{chain}, links...)
	}
	return flattenResult{chain.Expression(), links}
}

// NewLuaOptionalChainTransformer lowers `?.` to Lua. It runs unconditionally in
// the emit pipeline (Lua is always the target), not as part of the target-gated
// JS downlevel chain.
func NewLuaOptionalChainTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &luaOptionalChainTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}

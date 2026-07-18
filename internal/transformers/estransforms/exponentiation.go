package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/transformers"
)

type exponentiationTransformer struct {
	transformers.Transformer
}

func (ch *exponentiationTransformer) visit(node *ast.Node) *ast.Node {
	if node.SubtreeFacts()&ast.SubtreeContainsExponentiationOperator == 0 {
		return node
	}
	switch node.Kind {
	case ast.KindBinaryExpression:
		return ch.visitBinaryExpression(node.AsBinaryExpression())
	default:
		return ch.Visitor().VisitEachChild(node)
	}
}

func (ch *exponentiationTransformer) visitBinaryExpression(node *ast.BinaryExpression) *ast.Node {
	switch node.OperatorToken.Kind {
	case ast.KindAsteriskAsteriskToken:
		return ch.visitExponentiationExpression(node)
	}
	return ch.Visitor().VisitEachChild(node.AsNode())
}

func (ch *exponentiationTransformer) visitExponentiationExpression(node *ast.BinaryExpression) *ast.Node {
	left := ch.Visitor().VisitNode(node.Left)
	right := ch.Visitor().VisitNode(node.Right)
	result := ch.Factory().NewGlobalMethodCall("Math", "pow", []*ast.Node{left, right})
	result.Loc = node.Loc
	return result
}

func newExponentiationTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	tx := &exponentiationTransformer{}
	return tx.NewTransformer(tx.visit, opts.Context)
}

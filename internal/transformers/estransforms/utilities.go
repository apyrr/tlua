package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/printer"
)

func createNotNullCondition(emitContext *printer.EmitContext, left *ast.Node, right *ast.Node, invert bool) *ast.Node {
	token := ast.KindTildeEqualsToken
	op := ast.KindAmpersandAmpersandToken
	if invert {
		token = ast.KindEqualsEqualsToken
		op = ast.KindBarBarToken
	}

	return emitContext.Factory.NewBinaryExpression(
		nil,
		emitContext.Factory.NewBinaryExpression(
			nil,
			left,
			nil,
			emitContext.Factory.NewToken(token),
			emitContext.Factory.NewKeywordExpression(ast.KindNilKeyword),
		),
		nil,
		emitContext.Factory.NewToken(op),
		emitContext.Factory.NewBinaryExpression(
			nil,
			right,
			nil,
			emitContext.Factory.NewToken(token),
			emitContext.Factory.NewVoidZeroExpression(),
		),
	)
}

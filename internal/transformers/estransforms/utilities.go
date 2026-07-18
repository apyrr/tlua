package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/transformers"
)

func convertClassDeclarationToClassExpression(emitContext *printer.EmitContext, node *ast.ClassDeclaration) *ast.Expression {
	updated := emitContext.Factory.NewClassExpression(
		transformers.ExtractModifiers(emitContext, node.Modifiers(), ^ast.ModifierFlagsExportDefault),
		node.Name(),
		node.TypeParameters,
		node.HeritageClauses,
		node.Members,
	)
	emitContext.SetOriginal(updated, node.AsNode())
	updated.Loc = node.Loc
	return updated
}

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

// createAccessorPropertyBackingField creates a private backing field for an `accessor` PropertyDeclaration.
func createAccessorPropertyBackingField(f *printer.NodeFactory, node *ast.PropertyDeclaration, modifiers *ast.ModifierList, initializer *ast.Expression) *ast.Node {
	return f.UpdatePropertyDeclaration(
		node,
		modifiers,
		f.NewGeneratedPrivateNameForNodeEx(node.Name(), printer.AutoGenerateOptions{Suffix: "_accessor_storage"}),
		nil, /*postfixToken*/
		nil, /*typeNode*/
		initializer,
	)
}

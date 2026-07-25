package transformers

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/scanner"
)

func IsGeneratedIdentifier(emitContext *printer.EmitContext, name *ast.IdentifierNode) bool {
	return emitContext.HasAutoGenerateInfo(name)
}

func IsHelperName(emitContext *printer.EmitContext, name *ast.IdentifierNode) bool {
	return emitContext.EmitFlags(name)&printer.EFHelperName != 0
}

func IsLocalName(emitContext *printer.EmitContext, name *ast.IdentifierNode) bool {
	return emitContext.EmitFlags(name)&printer.EFLocalName != 0
}

func IsExportName(emitContext *printer.EmitContext, name *ast.IdentifierNode) bool {
	return emitContext.EmitFlags(name)&printer.EFExportName != 0
}

func IsIdentifierReference(name *ast.IdentifierNode, parent *ast.Node) bool {
	switch parent.Kind {
	case ast.KindBinaryExpression,
		ast.KindPrefixUnaryExpression,
		ast.KindAsExpression,
		ast.KindSatisfiesExpression,
		ast.KindElementAccessExpression,
		ast.KindNonNullExpression,
		ast.KindSpreadElement,
		ast.KindSpreadAssignment,
		ast.KindParenthesizedExpression,
		ast.KindArrayLiteralExpression,
		ast.KindDeleteExpression,
		ast.KindVoidExpression,
		ast.KindTypeAssertionExpression,
		ast.KindExpressionWithTypeArguments,
		ast.KindJsxSelfClosingElement,
		ast.KindJsxSpreadAttribute,
		ast.KindJsxExpression,
		ast.KindPartiallyEmittedExpression:
		// all immediate children that can be `Identifier` would be instances of `IdentifierReference`
		return true
	case ast.KindComputedPropertyName,
		ast.KindIfStatement,
		ast.KindWhileStatement,
		ast.KindReturnStatement,
		ast.KindThrowStatement,
		ast.KindExpressionStatement,
		ast.KindExportAssignment,
		ast.KindPropertyAccessExpression,
		ast.KindTemplateSpan:
		// only an `Expression()` child that can be `Identifier` would be an instance of `IdentifierReference`
		return parent.Expression() == name
	case ast.KindVariableDeclaration,
		ast.KindParameter,
		ast.KindBindingElement,
		ast.KindPropertySignature,
		ast.KindPropertyAssignment,
		ast.KindJsxAttribute:
		// only an `Initializer()` child that can be `Identifier` would be an instance of `IdentifierReference`
		return parent.Initializer() == name
	case ast.KindForOfStatement:
		return parent.Initializer() == name ||
			parent.Expression() == name
	case ast.KindNumericForStatement:
		return parent.AsNumericForStatement().From == name ||
			parent.AsNumericForStatement().To == name ||
			parent.AsNumericForStatement().Step == name
	case ast.KindRepeatStatement:
		return parent.AsRepeatStatement().Expression == name
	case ast.KindImportEqualsDeclaration:
		return parent.AsImportEqualsDeclaration().ModuleReference == name
	case ast.KindArrowFunction:
		return parent.Body() == name
	case ast.KindConditionalExpression:
		return parent.AsConditionalExpression().Condition == name ||
			parent.AsConditionalExpression().WhenTrue == name ||
			parent.AsConditionalExpression().WhenFalse == name
	case ast.KindCallExpression, ast.KindNewExpression:
		return parent.Expression() == name ||
			slices.Contains(parent.Arguments(), name)
	case ast.KindTaggedTemplateExpression:
		return parent.AsTaggedTemplateExpression().Tag == name
	case ast.KindJsxOpeningElement, ast.KindJsxClosingElement:
		return parent.TagName() == name
	default:
		return false
	}
}

func ConvertVariableDeclarationToAssignmentExpression(emitContext *printer.EmitContext, element *ast.VariableDeclaration) *ast.Expression {
	if element.Initializer == nil {
		return nil
	}
	// A Lua local declares a plain name: binding patterns parse only in
	// declaration files, which never reach a runtime-syntax transform.
	expression := element.Name()
	assignment := emitContext.Factory.NewAssignmentExpression(expression, element.Initializer)
	emitContext.SetOriginal(assignment, element.AsNode())
	emitContext.AssignCommentAndSourceMapRanges(assignment, element.AsNode())
	return assignment
}

func SingleOrMany(nodes []*ast.Node, factory *printer.NodeFactory) *ast.Node {
	if nodes == nil {
		return nil
	}
	if len(nodes) == 1 {
		return nodes[0]
	}
	return factory.NewSyntaxList(nodes)
}

// Used in the module transformer to check if an expression is reasonably without sideeffect,
//
//	and thus better to copy into multiple places rather than to cache in a temporary variable
//	- this is mostly subjective beyond the requirement that the expression not be sideeffecting
//
// Also used by the logical assignment downleveling transform to skip temp variables when they're
// not needed.
func IsSimpleCopiableExpression(expression *ast.Expression) bool {
	return ast.IsStringLiteralLike(expression) ||
		ast.IsNumericLiteral(expression) ||
		ast.IsKeywordKind(expression.Kind) ||
		ast.IsIdentifier(expression)
}

func IsOriginalNodeSingleLine(emitContext *printer.EmitContext, node *ast.Node) bool {
	if node == nil {
		return false
	}
	original := emitContext.MostOriginal(node)
	if original == nil {
		return false
	}
	source := ast.GetSourceFileOfNode(original)
	if source == nil {
		return false
	}
	startLine := scanner.GetECMALineOfPosition(source, original.Loc.Pos())
	endLine := scanner.GetECMALineOfPosition(source, original.Loc.End())
	return startLine == endLine
}

/**
 * A simple inlinable expression is an expression which can be copied into multiple locations
 * without risk of repeating any sideeffects and whose value could not possibly change between
 * any such locations
 */
func IsSimpleInlineableExpression(expression *ast.Expression) bool {
	return !ast.IsIdentifier(expression) && IsSimpleCopiableExpression(expression)
}

// MoveRangePastModifiers returns a text range that starts past any modifiers on the node.
func MoveRangePastModifiers(node *ast.Node) core.TextRange {
	var lastModifier *ast.Node
	if ast.CanHaveModifiers(node) {
		lastModifier = core.LastOrNil(node.ModifierNodes())
	}

	if lastModifier != nil && !ast.PositionIsSynthesized(lastModifier.End()) {
		return core.NewTextRange(lastModifier.End(), node.End())
	}
	return node.Loc
}

// GetNonAssignmentOperatorForCompoundAssignment returns the non-assignment operator for a compound assignment.
func GetNonAssignmentOperatorForCompoundAssignment(kind ast.Kind) ast.Kind {
	switch kind {
	case ast.KindPlusEqualsToken:
		return ast.KindPlusToken
	case ast.KindMinusEqualsToken:
		return ast.KindMinusToken
	case ast.KindAsteriskEqualsToken:
		return ast.KindAsteriskToken
	case ast.KindSlashEqualsToken:
		return ast.KindSlashToken
	case ast.KindPercentEqualsToken:
		return ast.KindPercentToken
	case ast.KindBarBarEqualsToken:
		return ast.KindBarBarToken
	case ast.KindAmpersandAmpersandEqualsToken:
		return ast.KindAmpersandAmpersandToken
	}
	return kind
}

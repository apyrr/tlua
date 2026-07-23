package estransforms

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/printer"
)

type anonymousFunctionDefinition = ast.Node // FunctionExpression | ArrowFunction

// Indicates whether an expression is an anonymous function definition.
//
// See https://tc39.es/ecma262/#sec-isanonymousfunctiondefinition
func isAnonymousFunctionDefinition(emitContext *printer.EmitContext, node *ast.Expression, cb func(*anonymousFunctionDefinition) bool) bool {
	node = ast.SkipOuterExpressions(node, ast.OEKAll)
	switch node.Kind {
	case ast.KindFunctionExpression:
		if node.AsFunctionExpression().Name() != nil {
			return false
		}
		break
	case ast.KindArrowFunction:
		break
	default:
		return false
	}
	if cb != nil {
		return cb(node)
	}
	return true
}

func isNamedEvaluation(emitContext *printer.EmitContext, node *ast.Node) bool {
	return isNamedEvaluationAnd(emitContext, node, nil)
}

func isNamedEvaluationAnd(emitContext *printer.EmitContext, node *ast.Node, cb func(*anonymousFunctionDefinition) bool) bool {
	if !ast.IsNamedEvaluationSource(node) {
		return false
	}
	switch node.Kind {
	case ast.KindShorthandPropertyAssignment:
		return isAnonymousFunctionDefinition(emitContext, node.AsShorthandPropertyAssignment().ObjectAssignmentInitializer, cb)
	case ast.KindPropertyAssignment, ast.KindVariableDeclaration, ast.KindParameter, ast.KindBindingElement:
		return isAnonymousFunctionDefinition(emitContext, node.Initializer(), cb)
	case ast.KindBinaryExpression:
		return isAnonymousFunctionDefinition(emitContext, node.AsBinaryExpression().Right, cb)
	case ast.KindExportAssignment:
		return isAnonymousFunctionDefinition(emitContext, node.Expression(), cb)
	default:
		debug.Fail("Unhandled case in isNamedEvaluation")
		return false
	}
}

// Gets a string literal to use as the assigned name of an anonymous class or function declaration.
func getAssignedNameOfIdentifier(emitContext *printer.EmitContext, name *ast.IdentifierNode, expression *ast.Node /*WrappedExpression<AnonymousFunctionDefinition>*/) *ast.StringLiteralNode {
	original := emitContext.MostOriginal(ast.SkipOuterExpressions(expression, ast.OEKAll))
	if ast.IsFunctionDeclaration(original) &&
		original.Name() == nil && ast.HasSyntacticModifier(original, ast.ModifierFlagsDefault) {
		return emitContext.Factory.NewStringLiteral("default", ast.TokenFlagsNone)
	}
	return emitContext.Factory.NewStringLiteralFromNode(name)
}

func getAssignedNameOfPropertyName(emitContext *printer.EmitContext, name *ast.PropertyName, assignedNameText string) (assignedName *ast.Expression, updatedName *ast.PropertyName) {
	factory := emitContext.Factory
	if len(assignedNameText) > 0 {
		assignedName := factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
		return assignedName, name
	}

	if ast.IsPropertyNameLiteral(name) || ast.IsPrivateIdentifier(name) {
		assignedName := factory.NewStringLiteralFromNode(name)
		return assignedName, name
	}

	expression := name.Expression()
	if ast.IsPropertyNameLiteral(expression) && !ast.IsIdentifier(expression) {
		assignedName := factory.NewStringLiteralFromNode(expression)
		return assignedName, name
	}

	debug.Assert(ast.IsComputedPropertyName(name), "Expected computed property name")

	assignedName = factory.NewGeneratedNameForNode(name)
	emitContext.AddVariableDeclaration(assignedName)

	key := factory.NewPropKeyHelper(expression)
	assignment := factory.NewAssignmentExpression(assignedName, key)
	updatedName = factory.UpdateComputedPropertyName(name.AsComputedPropertyName(), assignment)
	return assignedName, updatedName
}

// Creates a class `static {}` block used to dynamically set the name of a class.
//
// The assignedName parameter is the expression used to resolve the assigned name at runtime. This expression should not produce
// side effects.
// The thisExpression parameter overrides the expression to use for the actual `this` reference. This can be used to provide an
// expression that has already had its `EmitFlags` set or may have been tracked to prevent substitution.
func finishTransformNamedEvaluation(
	emitContext *printer.EmitContext,
	expression *ast.Node, // WrappedExpression<AnonymousFunctionDefinition>,
	assignedName *ast.Expression,
	ignoreEmptyStringLiteral bool,
) *ast.Expression {
	if ignoreEmptyStringLiteral && ast.IsStringLiteral(assignedName) && len(assignedName.Text()) == 0 {
		return expression
	}

	factory := emitContext.Factory
	innerExpression := ast.SkipOuterExpressions(expression, ast.OEKAll)

	// Classes are removed in tlua, so a named-evaluation target is never a class
	// expression; assign the name via the SetFunctionName helper.
	updatedExpression := factory.NewSetFunctionNameHelper(innerExpression, assignedName, "" /*prefix*/)

	return factory.RestoreOuterExpressions(expression, updatedExpression, ast.OEKAll)
}

func transformNamedEvaluationOfPropertyAssignment(context *printer.EmitContext, node *ast.PropertyAssignment /*NamedEvaluation & PropertyAssignment*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 13.2.5.5 RS: PropertyDefinitionEvaluation
	//   PropertyAssignment : PropertyName `:` AssignmentExpression
	//     ...
	//     5. If IsAnonymousFunctionDefinition(|AssignmentExpression|) is *true* and _isProtoSetter_ is *false*, then
	//        a. Let _popValue_ be ? NamedEvaluation of |AssignmentExpression| with argument _propKey_.
	//     ...

	factory := context.Factory
	assignedName, name := getAssignedNameOfPropertyName(context, node.Name(), assignedNameText)
	initializer := finishTransformNamedEvaluation(context, node.Initializer, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdatePropertyAssignment(node, nil /*modifiers*/, name, nil /*postfixToken*/, nil /*typeNode*/, initializer)
}

func transformNamedEvaluationOfShorthandAssignmentProperty(emitContext *printer.EmitContext, node *ast.ShorthandPropertyAssignment /*NamedEvaluation & ShorthandPropertyAssignment*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 13.15.5.3 RS: PropertyDestructuringAssignmentEvaluation
	//   AssignmentProperty : IdentifierReference Initializer?
	//     ...
	//     4. If |Initializer?| is present and _v_ is *undefined*, then
	//        a. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//           i. Set _v_ to ? NamedEvaluation of |Initializer| with argument _P_.
	//     ...

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else {
		assignedName = getAssignedNameOfIdentifier(emitContext, node.Name(), node.ObjectAssignmentInitializer)
	}
	objectAssignmentInitializer := finishTransformNamedEvaluation(emitContext, node.ObjectAssignmentInitializer, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateShorthandPropertyAssignment(
		node,
		nil, /*modifiers*/
		node.Name(),
		nil, /*postfixToken*/
		nil, /*typeNode*/
		node.EqualsToken,
		objectAssignmentInitializer,
	)
}

func transformNamedEvaluationOfVariableDeclaration(emitContext *printer.EmitContext, node *ast.VariableDeclaration /*NamedEvaluation & VariableDeclaration*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 14.3.1.2 RS: Evaluation
	//   LexicalBinding : BindingIdentifier Initializer
	//     ...
	//     3. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//        a. Let _value_ be ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...
	//
	// 14.3.2.1 RS: Evaluation
	//   VariableDeclaration : BindingIdentifier Initializer
	//     ...
	//     3. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//        a. Let _value_ be ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else {
		assignedName = getAssignedNameOfIdentifier(emitContext, node.Name(), node.Initializer)
	}
	initializer := finishTransformNamedEvaluation(emitContext, node.Initializer, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateVariableDeclaration(
		node,
		node.Name(),
		nil, /*exclamationToken*/
		nil, /*typeNode*/
		initializer,
	)
}

func transformNamedEvaluationOfParameterDeclaration(emitContext *printer.EmitContext, node *ast.ParameterDeclaration /*NamedEvaluation & ParameterDeclaration*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 8.6.3 RS: IteratorBindingInitialization
	//   SingleNameBinding : BindingIdentifier Initializer?
	//     ...
	//     5. If |Initializer| is present and _v_ is *undefined*, then
	//        a. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//           i. Set _v_ to ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...
	//
	// 14.3.3.3 RS: KeyedBindingInitialization
	//   SingleNameBinding : BindingIdentifier Initializer?
	//     ...
	//     4. If |Initializer| is present and _v_ is *undefined*, then
	//        a. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//           i. Set _v_ to ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else {
		assignedName = getAssignedNameOfIdentifier(emitContext, node.Name(), node.Initializer)
	}
	initializer := finishTransformNamedEvaluation(emitContext, node.Initializer, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateParameterDeclaration(
		node,
		nil, /*modifiers*/
		node.DotDotDotToken,
		node.Name(),
		nil, /*questionToken*/
		nil, /*typeNode*/
		initializer,
	)
}

func transformNamedEvaluationOfBindingElement(emitContext *printer.EmitContext, node *ast.BindingElement /*NamedEvaluation & BindingElement*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 8.6.3 RS: IteratorBindingInitialization
	//   SingleNameBinding : BindingIdentifier Initializer?
	//     ...
	//     5. If |Initializer| is present and _v_ is *undefined*, then
	//        a. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//           i. Set _v_ to ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...
	//
	// 14.3.3.3 RS: KeyedBindingInitialization
	//   SingleNameBinding : BindingIdentifier Initializer?
	//     ...
	//     4. If |Initializer| is present and _v_ is *undefined*, then
	//        a. If IsAnonymousFunctionDefinition(|Initializer|) is *true*, then
	//           i. Set _v_ to ? NamedEvaluation of |Initializer| with argument _bindingId_.
	//     ...

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else {
		assignedName = getAssignedNameOfIdentifier(emitContext, node.Name(), node.Initializer)
	}
	initializer := finishTransformNamedEvaluation(emitContext, node.Initializer, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateBindingElement(
		node,
		node.DotDotDotToken,
		node.PropertyName,
		node.Name(),
		initializer,
	)
}

func transformNamedEvaluationOfAssignmentExpression(emitContext *printer.EmitContext, node *ast.BinaryExpression /*NamedEvaluation & BinaryExpression*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 13.15.2 RS: Evaluation
	//   AssignmentExpression : LeftHandSideExpression `=` AssignmentExpression
	//     1. If |LeftHandSideExpression| is neither an |ObjectLiteral| nor an |ArrayLiteral|, then
	//        a. Let _lref_ be ? Evaluation of |LeftHandSideExpression|.
	//        b. If IsAnonymousFunctionDefinition(|AssignmentExpression|) and IsIdentifierRef of |LeftHandSideExpression| are both *true*, then
	//           i. Let _rval_ be ? NamedEvaluation of |AssignmentExpression| with argument _lref_.[[ReferencedName]].
	//     ...
	//
	//   AssignmentExpression : LeftHandSideExpression `&&=` AssignmentExpression
	//     ...
	//     5. If IsAnonymousFunctionDefinition(|AssignmentExpression|) is *true* and IsIdentifierRef of |LeftHandSideExpression| is *true*, then
	//        a. Let _rval_ be ? NamedEvaluation of |AssignmentExpression| with argument _lref_.[[ReferencedName]].
	//     ...
	//
	//   AssignmentExpression : LeftHandSideExpression `||=` AssignmentExpression
	//     ...
	//     5. If IsAnonymousFunctionDefinition(|AssignmentExpression|) is *true* and IsIdentifierRef of |LeftHandSideExpression| is *true*, then
	//        a. Let _rval_ be ? NamedEvaluation of |AssignmentExpression| with argument _lref_.[[ReferencedName]].
	//     ...
	//
	//   AssignmentExpression : LeftHandSideExpression `??=` AssignmentExpression
	//     ...
	//     4. If IsAnonymousFunctionDefinition(|AssignmentExpression|) is *true* and IsIdentifierRef of |LeftHandSideExpression| is *true*, then
	//        a. Let _rval_ be ? NamedEvaluation of |AssignmentExpression| with argument _lref_.[[ReferencedName]].
	//     ...

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else {
		assignedName = getAssignedNameOfIdentifier(emitContext, node.Left, node.Right)
	}
	right := finishTransformNamedEvaluation(emitContext, node.Right, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateBinaryExpression(
		node,
		nil, /*modifiers*/
		node.Left,
		nil, /*typeNode*/
		node.OperatorToken,
		right,
	)
}

func transformNamedEvaluationOfExportAssignment(emitContext *printer.EmitContext, node *ast.ExportAssignment /*NamedEvaluation & ExportAssignment*/, ignoreEmptyStringLiteral bool, assignedNameText string) *ast.Expression {
	// 16.2.3.7 RS: Evaluation
	//   ExportDeclaration : `export` `default` AssignmentExpression `;`
	//     1. If IsAnonymousFunctionDefinition(|AssignmentExpression|) is *true*, then
	//        a. Let _value_ be ? NamedEvaluation of |AssignmentExpression| with argument `"default"`.
	//     ...

	// NOTE: Since emit for `export =` translates to `module.exports = ...`, the assigned name of the class or function
	// is `""`.

	factory := emitContext.Factory
	var assignedName *ast.Expression
	if len(assignedNameText) > 0 {
		assignedName = factory.NewStringLiteral(assignedNameText, ast.TokenFlagsNone)
	} else if node.IsExportEquals {
		assignedName = factory.NewStringLiteral("", ast.TokenFlagsNone)
	} else {
		assignedName = factory.NewStringLiteral("default", ast.TokenFlagsNone)
	}
	expression := finishTransformNamedEvaluation(emitContext, node.Expression, assignedName, ignoreEmptyStringLiteral)
	return factory.UpdateExportAssignment(
		node,
		nil, /*modifiers*/
		node.IsExportEquals,
		nil, /*typeNode*/
		expression,
	)
}

// Performs a shallow transformation of a `NamedEvaluation` node, such that a valid name will be assigned.
func transformNamedEvaluation(context *printer.EmitContext, node *ast.Node /*NamedEvaluation*/, ignoreEmptyStringLiteral bool, assignedName string) *ast.Expression {
	switch node.Kind {
	case ast.KindPropertyAssignment:
		return transformNamedEvaluationOfPropertyAssignment(context, node.AsPropertyAssignment(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindShorthandPropertyAssignment:
		return transformNamedEvaluationOfShorthandAssignmentProperty(context, node.AsShorthandPropertyAssignment(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindVariableDeclaration:
		return transformNamedEvaluationOfVariableDeclaration(context, node.AsVariableDeclaration(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindParameter:
		return transformNamedEvaluationOfParameterDeclaration(context, node.AsParameterDeclaration(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindBindingElement:
		return transformNamedEvaluationOfBindingElement(context, node.AsBindingElement(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindBinaryExpression:
		return transformNamedEvaluationOfAssignmentExpression(context, node.AsBinaryExpression(), ignoreEmptyStringLiteral, assignedName)
	case ast.KindExportAssignment:
		return transformNamedEvaluationOfExportAssignment(context, node.AsExportAssignment(), ignoreEmptyStringLiteral, assignedName)
	default:
		debug.Fail("Unhandled case in transformNamedEvaluation")
		return node
	}
}

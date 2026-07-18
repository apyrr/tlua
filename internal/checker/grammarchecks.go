package checker

import (
	"fmt"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/binder"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/jsnum"
	"github.com/apyrr/tlua/internal/scanner"
)

func (c *Checker) grammarErrorOnFirstToken(node *ast.Node, message *diagnostics.Message, args ...any) bool {
	sourceFile := ast.GetSourceFileOfNode(node)
	if !c.hasParseDiagnostics(sourceFile) {
		span := scanner.GetRangeOfTokenAtPosition(sourceFile, node.Pos())
		c.addDiagnostic(ast.NewDiagnostic(sourceFile, span, message, args...))
		return true
	}
	return false
}

func (c *Checker) grammarErrorAtPos(nodeForSourceFile *ast.Node, start int, length int, message *diagnostics.Message, args ...any) bool {
	sourceFile := ast.GetSourceFileOfNode(nodeForSourceFile)
	if !c.hasParseDiagnostics(sourceFile) {
		c.addDiagnostic(ast.NewDiagnostic(sourceFile, core.NewTextRange(start, start+length), message, args...))
		return true
	}
	return false
}

func (c *Checker) grammarErrorOnNode(node *ast.Node, message *diagnostics.Message, args ...any) bool {
	sourceFile := ast.GetSourceFileOfNode(node)
	if !c.hasParseDiagnostics(sourceFile) {
		c.error(node, message, args...)
		return true
	}
	return false
}

func (c *Checker) grammarErrorOnNodeSkippedOnNoEmit(node *ast.Node, message *diagnostics.Message, args ...any) bool {
	sourceFile := ast.GetSourceFileOfNode(node)
	if !c.hasParseDiagnostics(sourceFile) {
		d := NewDiagnosticForNode(node, message, args...)
		d.SetSkippedOnNoEmit()
		c.addDiagnostic(d)
		return true
	}
	return false
}

func getIdentifierFromEntityNameExpression(node *ast.Node) *ast.Node {
	switch node.Kind {
	case ast.KindIdentifier:
		return node
	case ast.KindPropertyAccessExpression:
		return node.AsPropertyAccessExpression().Name()
	default:
		return nil
	}
}

func (c *Checker) checkGrammarRegularExpressionLiteral(node *ast.RegularExpressionLiteral) bool {
	sourceFile := ast.GetSourceFileOfNode(node.AsNode())
	if !c.hasParseDiagnostics(sourceFile) {
		var lastError *ast.Diagnostic
		if c.regExpScanner == nil {
			c.regExpScanner = scanner.NewScanner()
		}
		c.regExpScanner.SetScriptTarget(c.languageVersion)
		c.regExpScanner.SetLanguageVariant(sourceFile.LanguageVariant)
		c.regExpScanner.SetOnError(func(message *diagnostics.Message, start int, length int, args ...any) {
			if message.Category() == diagnostics.CategoryMessage && lastError != nil && start == lastError.Pos() && length == lastError.Len() {
				// For providing spelling suggestions.
				err := ast.NewDiagnostic(nil, core.NewTextRange(start, start+length), message, args...)
				lastError.AddRelatedInfo(err)
			} else if lastError == nil || start != lastError.Pos() {
				lastError = ast.NewDiagnostic(sourceFile, core.NewTextRange(start, start+length), message, args...)
				c.addDiagnostic(lastError)
			}
		})
		c.regExpScanner.SetText(sourceFile.Text())
		c.regExpScanner.ResetTokenState(node.AsNode().Pos())
		c.regExpScanner.Scan()
		tokenIsRegularExpressionLiteral := c.regExpScanner.ReScanSlashToken(true) == ast.KindRegularExpressionLiteral
		c.regExpScanner.SetText("")
		c.regExpScanner.SetOnError(nil)
		debug.Assert(tokenIsRegularExpressionLiteral)
		return lastError != nil
	}
	return false
}

func (c *Checker) checkGrammarPrivateIdentifierExpression(privId *ast.PrivateIdentifier) bool {
	privIdAsNode := privId.AsNode()
	if ast.GetContainingClass(privId.AsNode()) == nil {
		return c.grammarErrorOnNode(privId.AsNode(), diagnostics.Private_identifiers_are_not_allowed_outside_class_bodies)
	}

	if !ast.IsExpressionNode(privIdAsNode) {
		return c.grammarErrorOnNode(privIdAsNode, diagnostics.Private_identifiers_are_only_allowed_in_class_bodies_and_may_only_be_used_as_part_of_a_class_member_declaration_property_access_or_on_the_left_hand_side_of_an_in_expression)
	}

	isInOperation := ast.IsBinaryExpression(privId.Parent) && privId.Parent.AsBinaryExpression().OperatorToken.Kind == ast.KindInKeyword
	if c.getSymbolForPrivateIdentifierExpression(privIdAsNode) == nil && !isInOperation {
		return c.grammarErrorOnNode(privIdAsNode, diagnostics.Cannot_find_name_0, privId.Text)
	}

	return false
}

func (c *Checker) checkGrammarMappedType(node *ast.MappedTypeNode) bool {
	if len(node.Members.Nodes) > 0 {
		return c.grammarErrorOnNode(node.Members.Nodes[0], diagnostics.A_mapped_type_may_not_declare_properties_or_methods)
	}
	return false
}

func (c *Checker) checkGrammarModuleElementContext(node *ast.Statement, errorMessage *diagnostics.Message) bool {
	isInAppropriateContext := node.Parent.Kind == ast.KindSourceFile || node.Parent.Kind == ast.KindModuleBlock || node.Parent.Kind == ast.KindModuleDeclaration
	if !isInAppropriateContext {
		c.grammarErrorOnFirstToken(node, errorMessage)
	}
	return !isInAppropriateContext
}

func (c *Checker) checkGrammarModifiers(node *ast.Node /*Union[HasModifiers, HasIllegalModifiers]*/) bool {
	if node.Modifiers() == nil {
		return false
	}
	if c.reportObviousModifierErrors(node) {
		return true
	}
	var lastStatic *ast.Node
	var lastAsync *ast.Node
	var lastOverride *ast.Node
	flags := ast.ModifierFlagsNone
	modifiers := node.ModifierNodes()
	for _, modifier := range modifiers {
		if modifier.Kind != ast.KindReadonlyKeyword {
			if node.Kind == ast.KindPropertySignature || node.Kind == ast.KindMethodSignature {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_type_member, scanner.TokenToString(modifier.Kind))
			}
			if node.Kind == ast.KindIndexSignature && (modifier.Kind != ast.KindStaticKeyword || !ast.IsClassLike(node.Parent)) {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_an_index_signature, scanner.TokenToString(modifier.Kind))
			}
		}
		if modifier.Kind != ast.KindInKeyword && modifier.Kind != ast.KindOutKeyword {
			if node.Kind == ast.KindTypeParameter {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_type_parameter, scanner.TokenToString(modifier.Kind))
			}
		}
		switch modifier.Kind {
		case ast.KindOverrideKeyword:
			// If node.kind === SyntaxKind.Parameter, checkParameter reports an error if it's not a parameter property.
			if flags&ast.ModifierFlagsOverride != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "override")
			} else if flags&ast.ModifierFlagsAmbient != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "override", "declare")
			} else if flags&ast.ModifierFlagsReadonly != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "override", "readonly")
			} else if flags&ast.ModifierFlagsAccessor != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "override", "accessor")
			} else if flags&ast.ModifierFlagsAsync != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "override", "async")
			}
			flags |= ast.ModifierFlagsOverride
			lastOverride = modifier

		case ast.KindPublicKeyword,
			ast.KindProtectedKeyword,
			ast.KindPrivateKeyword:
			text := visibilityToString(ast.ModifierToFlag(modifier.Kind))

			if flags&ast.ModifierFlagsAccessibilityModifier != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.Accessibility_modifier_already_seen)
			} else if flags&ast.ModifierFlagsOverride != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "override")
			} else if flags&ast.ModifierFlagsStatic != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "static")
			} else if flags&ast.ModifierFlagsAccessor != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "accessor")
			} else if flags&ast.ModifierFlagsReadonly != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "readonly")
			} else if flags&ast.ModifierFlagsAsync != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "async")
			} else if node.Parent.Kind == ast.KindModuleBlock || node.Parent.Kind == ast.KindSourceFile {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_module_or_namespace_element, text)
			} else if flags&ast.ModifierFlagsAbstract != 0 {
				if modifier.Kind == ast.KindPrivateKeyword {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, text, "abstract")
				} else if modifier.Flags&ast.NodeFlagsReparsed == 0 {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, text, "abstract")
				}
			} else if ast.IsPrivateIdentifierClassElementDeclaration(node) {
				return c.grammarErrorOnNode(modifier, diagnostics.An_accessibility_modifier_cannot_be_used_with_a_private_identifier)
			}
			flags |= ast.ModifierToFlag(modifier.Kind)
		case ast.KindStaticKeyword:
			if flags&ast.ModifierFlagsStatic != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "static")
			} else if flags&ast.ModifierFlagsReadonly != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "static", "readonly")
			} else if flags&ast.ModifierFlagsAsync != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "static", "async")
			} else if flags&ast.ModifierFlagsAccessor != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "static", "accessor")
			} else if node.Parent.Kind == ast.KindModuleBlock || node.Parent.Kind == ast.KindSourceFile {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_module_or_namespace_element, "static")
			} else if node.Kind == ast.KindParameter {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_parameter, "static")
			} else if flags&ast.ModifierFlagsAbstract != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "static", "abstract")
			} else if flags&ast.ModifierFlagsOverride != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "static", "override")
			}
			flags |= ast.ModifierFlagsStatic
			lastStatic = modifier
		case ast.KindAccessorKeyword:
			if flags&ast.ModifierFlagsAccessor != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "accessor")
			} else if flags&ast.ModifierFlagsReadonly != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "accessor", "readonly")
			} else if flags&ast.ModifierFlagsAmbient != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "accessor", "declare")
			} else if node.Kind != ast.KindPropertyDeclaration {
				return c.grammarErrorOnNode(modifier, diagnostics.X_accessor_modifier_can_only_appear_on_a_property_declaration)
			}

			flags |= ast.ModifierFlagsAccessor
		case ast.KindReadonlyKeyword:
			if flags&ast.ModifierFlagsReadonly != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "readonly")
			} else if node.Kind != ast.KindPropertyDeclaration && node.Kind != ast.KindPropertySignature && node.Kind != ast.KindIndexSignature && node.Kind != ast.KindParameter {
				// If node.kind === SyntaxKind.Parameter, checkParameter reports an error if it's not a parameter property.
				return c.grammarErrorOnNode(modifier, diagnostics.X_readonly_modifier_can_only_appear_on_a_property_declaration_or_index_signature)
			} else if flags&ast.ModifierFlagsAccessor != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "readonly", "accessor")
			}
			flags |= ast.ModifierFlagsReadonly
		case ast.KindExportKeyword:
			if flags&ast.ModifierFlagsExport != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "export")
			} else if flags&ast.ModifierFlagsAmbient != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "export", "declare")
			} else if flags&ast.ModifierFlagsAbstract != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "export", "abstract")
			} else if flags&ast.ModifierFlagsAsync != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "export", "async")
			} else if ast.IsClassLike(node.Parent) && !ast.IsJSTypeAliasDeclaration(node) {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_class_elements_of_this_kind, "export")
			} else if node.Kind == ast.KindParameter {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_parameter, "export")
			}
			flags |= ast.ModifierFlagsExport
		case ast.KindDefaultKeyword:
			var container *ast.Node
			if node.Parent.Kind == ast.KindSourceFile {
				container = node.Parent
			} else {
				container = node.Parent.Parent
			}
			if container.Kind == ast.KindModuleDeclaration && !ast.IsAmbientModule(container) {
				return c.grammarErrorOnNode(modifier, diagnostics.A_default_export_can_only_be_used_in_an_ECMAScript_style_module)
			} else if flags&ast.ModifierFlagsExport == 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "export", "default")
			}

			flags |= ast.ModifierFlagsDefault
		case ast.KindDeclareKeyword:
			if flags&ast.ModifierFlagsAmbient != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "declare")
			} else if flags&ast.ModifierFlagsAsync != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_in_an_ambient_context, "async")
			} else if flags&ast.ModifierFlagsOverride != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_in_an_ambient_context, "override")
			} else if ast.IsClassLike(node.Parent) && !ast.IsPropertyDeclaration(node) {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_class_elements_of_this_kind, "declare")
			} else if node.Kind == ast.KindParameter {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_parameter, "declare")
			} else if (node.Parent.Flags&ast.NodeFlagsAmbient != 0) && node.Parent.Kind == ast.KindModuleBlock && !ast.IsVariableStatement(node) {
				// `declare name: Type` is the only spelling of an ambient value
				// (the var/let/const forms are deleted), so the declare modifier
				// it carries is legal even inside `declare global { ... }`.
				return c.grammarErrorOnNode(modifier, diagnostics.A_declare_modifier_cannot_be_used_in_an_already_ambient_context)
			} else if ast.IsPrivateIdentifierClassElementDeclaration(node) {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_a_private_identifier, "declare")
			} else if flags&ast.ModifierFlagsAccessor != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "declare", "accessor")
			}
			flags |= ast.ModifierFlagsAmbient
		case ast.KindAbstractKeyword:
			if flags&ast.ModifierFlagsAbstract != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "abstract")
			}
			if node.Kind != ast.KindClassDeclaration && node.Kind != ast.KindConstructorType {
				if node.Kind != ast.KindMethodDeclaration && node.Kind != ast.KindPropertyDeclaration && node.Kind != ast.KindGetAccessor && node.Kind != ast.KindSetAccessor {
					return c.grammarErrorOnNode(modifier, diagnostics.X_abstract_modifier_can_only_appear_on_a_class_method_or_property_declaration)
				}
				if !(node.Parent.Kind == ast.KindClassDeclaration && ast.HasSyntacticModifier(node.Parent, ast.ModifierFlagsAbstract)) {
					var message *diagnostics.Message
					if node.Kind == ast.KindPropertyDeclaration {
						message = diagnostics.Abstract_properties_can_only_appear_within_an_abstract_class
					} else {
						message = diagnostics.Abstract_methods_can_only_appear_within_an_abstract_class
					}
					return c.grammarErrorOnNode(modifier, message)
				}
				if flags&ast.ModifierFlagsStatic != 0 {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "static", "abstract")
				}
				if flags&ast.ModifierFlagsPrivate != 0 {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "private", "abstract")
				}
				if flags&ast.ModifierFlagsAsync != 0 && lastAsync != nil {
					return c.grammarErrorOnNode(lastAsync, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "async", "abstract")
				}
				if flags&ast.ModifierFlagsOverride != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "abstract", "override")
				}
				if flags&ast.ModifierFlagsAccessor != 0 && modifier.Flags&ast.NodeFlagsReparsed == 0 {
					return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "abstract", "accessor")
				}
			}
			if name := node.Name(); name != nil && name.Kind == ast.KindPrivateIdentifier {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_a_private_identifier, "abstract")
			}

			flags |= ast.ModifierFlagsAbstract
		case ast.KindAsyncKeyword:
			if flags&ast.ModifierFlagsAsync != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, "async")
			} else if node.Kind != ast.KindFunctionType && (flags&ast.ModifierFlagsAmbient != 0 || node.Parent.Flags&ast.NodeFlagsAmbient != 0) {
				// An `async (x) => T` function *type* is not an implementation,
				// so it is allowed in ambient/.d.ts contexts.
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_in_an_ambient_context, "async")
			} else if node.Kind == ast.KindParameter {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_appear_on_a_parameter, "async")
			}
			if flags&ast.ModifierFlagsAbstract != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_cannot_be_used_with_1_modifier, "async", "abstract")
			}
			flags |= ast.ModifierFlagsAsync
			lastAsync = modifier
		case ast.KindInKeyword,
			ast.KindOutKeyword:
			var inOutFlag ast.ModifierFlags
			if modifier.Kind == ast.KindInKeyword {
				inOutFlag = ast.ModifierFlagsIn
			} else {
				inOutFlag = ast.ModifierFlagsOut
			}
			var inOutText string
			if modifier.Kind == ast.KindInKeyword {
				inOutText = "in"
			} else {
				inOutText = "out"
			}
			parent := node.Parent
			if node.Kind != ast.KindTypeParameter || parent != nil && !(ast.IsInterfaceDeclaration(parent) || ast.IsClassLike(parent) || ast.IsTypeOrJSTypeAliasDeclaration(parent)) {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_can_only_appear_on_a_type_parameter_of_a_class_interface_or_type_alias, inOutText)
			}
			if flags&inOutFlag != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_already_seen, inOutText)
			}
			if inOutFlag&ast.ModifierFlagsIn != 0 && flags&ast.ModifierFlagsOut != 0 {
				return c.grammarErrorOnNode(modifier, diagnostics.X_0_modifier_must_precede_1_modifier, "in", "out")
			}
			flags |= inOutFlag
		}
	}

	if node.Kind == ast.KindConstructor {
		if flags&ast.ModifierFlagsStatic != 0 {
			return c.grammarErrorOnNode(lastStatic, diagnostics.X_0_modifier_cannot_appear_on_a_constructor_declaration, "static")
		}
		if flags&ast.ModifierFlagsOverride != 0 {
			return c.grammarErrorOnNode(lastOverride, diagnostics.X_0_modifier_cannot_appear_on_a_constructor_declaration, "override")
		}
		if flags&ast.ModifierFlagsAsync != 0 {
			return c.grammarErrorOnNode(lastAsync, diagnostics.X_0_modifier_cannot_appear_on_a_constructor_declaration, "async")
		}
		return false
	} else if node.Kind == ast.KindParameter && (flags&ast.ModifierFlagsParameterPropertyModifier != 0) && ast.IsBindingPattern(node.Name()) {
		return c.grammarErrorOnNode(node, diagnostics.A_parameter_property_may_not_be_declared_using_a_binding_pattern)
	} else if node.Kind == ast.KindParameter && (flags&ast.ModifierFlagsParameterPropertyModifier != 0) && node.AsParameterDeclaration().DotDotDotToken != nil {
		return c.grammarErrorOnNode(node, diagnostics.A_parameter_property_cannot_be_declared_using_a_rest_parameter)
	}
	if flags&ast.ModifierFlagsAsync != 0 {
		return c.checkGrammarAsyncModifier(node, lastAsync)
	}
	return false
}

func (c *Checker) reportObviousModifierErrors(node *ast.Node) bool {
	modifier := c.findFirstIllegalModifier(node)
	if modifier == nil {
		return false
	}
	return c.grammarErrorOnFirstToken(modifier, diagnostics.Modifiers_cannot_appear_here)
}

func (c *Checker) findFirstModifierExcept(node *ast.Node, allowedModifier ast.Kind) *ast.Node {
	modifier := core.Find(node.ModifierNodes(), ast.IsModifier)
	if modifier != nil && modifier.Kind != allowedModifier {
		return modifier
	}
	return nil
}

func (c *Checker) findFirstIllegalModifier(node *ast.Node) *ast.Node {
	switch node.Kind {
	case ast.KindGetAccessor,
		ast.KindSetAccessor,
		ast.KindConstructor,
		ast.KindPropertyDeclaration,
		ast.KindPropertySignature,
		ast.KindMethodDeclaration,
		ast.KindMethodSignature,
		ast.KindIndexSignature,
		ast.KindModuleDeclaration,
		ast.KindFunctionExpression,
		ast.KindArrowFunction,
		ast.KindParameter,
		ast.KindTypeParameter,
		ast.KindJSTypeAliasDeclaration:
		return nil
	case ast.KindClassStaticBlockDeclaration,
		ast.KindPropertyAssignment,
		ast.KindShorthandPropertyAssignment,
		ast.KindNamespaceExportDeclaration,
		ast.KindMissingDeclaration:
		return core.Find(node.ModifierNodes(), ast.IsModifier)
	default:
		if node.Parent.Kind == ast.KindModuleBlock || node.Parent.Kind == ast.KindSourceFile {
			return nil
		}
		switch node.Kind {
		case ast.KindFunctionDeclaration,
			ast.KindFunctionType:
			return c.findFirstModifierExcept(node, ast.KindAsyncKeyword)
		case ast.KindClassDeclaration,
			ast.KindConstructorType:
			return c.findFirstModifierExcept(node, ast.KindAbstractKeyword)
		case ast.KindClassExpression,
			ast.KindInterfaceDeclaration,
			ast.KindTypeAliasDeclaration:
			return core.Find(node.ModifierNodes(), ast.IsModifier)
		case ast.KindVariableStatement:
			return core.Find(node.ModifierNodes(), ast.IsModifier)
		default:
			panic("Unhandled case in findFirstIllegalModifier.")
		}
	}
}

func (c *Checker) checkGrammarAsyncModifier(node *ast.Node, asyncModifier *ast.Node) bool {
	switch node.Kind {
	case ast.KindMethodDeclaration,
		ast.KindFunctionDeclaration,
		ast.KindFunctionExpression,
		ast.KindArrowFunction,
		ast.KindFunctionType:
		return false
	}

	return c.grammarErrorOnNode(asyncModifier, diagnostics.X_0_modifier_cannot_be_used_here, "async")
}

func (c *Checker) checkGrammarForDisallowedTrailingComma(list *ast.NodeList, diag *diagnostics.Message) bool {
	if list != nil && list.HasTrailingComma() {
		return c.grammarErrorAtPos(list.Nodes[0], list.End()-len(","), len(","), diag)
	}
	return false
}

func (c *Checker) checkGrammarTypeParameterList(typeParameters *ast.NodeList, file *ast.SourceFile) bool {
	if typeParameters != nil && len(typeParameters.Nodes) == 0 {
		start := typeParameters.Pos() - len("<")
		end := scanner.SkipTrivia(file.Text(), typeParameters.End()) + len(">")
		return c.grammarErrorAtPos(file.AsNode(), start, end-start, diagnostics.Type_parameter_list_cannot_be_empty)
	}
	if typeParameters != nil {
		// A generic pack parameter (`<...A>`) binds a value pack whose length is not
		// fixed, so a plain parameter after it could never be positioned. Packs trail
		// plain parameters, matching where a value pack sits in a list (Lua, and
		// Luau's `<A, R...>`). Multiple trailing packs are allowed.
		seenPack := false
		for _, tp := range typeParameters.Nodes {
			if ast.IsPackTypeParameterDeclaration(tp) {
				seenPack = true
			} else if seenPack {
				return c.grammarErrorOnNode(tp, diagnostics.A_generic_type_parameter_cannot_follow_a_generic_pack_parameter)
			}
		}
	}
	return false
}

// checkGrammarMultiReturnType rejects an unparenthesized multiple return type
// (`function f(): number, string`). The parser accepts the bare list for
// recovery and leaves the diagnostic here, because a parse error would suppress
// every semantic diagnostic in the program.
//
// Parenthesized or not is structural, not a flag: parseParenthesizedReturnTypeList
// starts the node at the `(`, so the node begins before its first element, while
// the bare list starts at the element itself.
func (c *Checker) checkGrammarMultiReturnType(node *ast.Node) {
	elements := node.AsMultiReturnTypeNode().Elements
	if node.Flags&ast.NodeFlagsSynthesized != 0 || elements == nil || len(elements.Nodes) == 0 {
		return
	}
	// Reported with c.error rather than grammarErrorOnNode, for the same reason as
	// the named vararg: the parser builds these nodes in full, so they are not
	// recovery noise, and going quiet in every file that happens to have a parse
	// error elsewhere would hide the one rule the author needs to see.
	for _, element := range elements.Nodes {
		// A predicate narrows the one value a call is used for; a pack has several,
		// and no element of one can be the narrowed value.
		if ast.IsTypePredicateNode(element) {
			c.error(element, diagnostics.A_type_predicate_cannot_be_part_of_a_multiple_return_type)
		}
	}
	if node.Pos() >= elements.Nodes[0].Pos() {
		c.error(node, diagnostics.Multiple_return_types_must_be_parenthesized)
	}
}

func (c *Checker) checkGrammarParameterList(parameters *ast.NodeList) bool {
	seenOptionalParameter := false
	parameterCount := len(parameters.Nodes)

	for i := range parameterCount {
		parameter := parameters.Nodes[i].AsParameterDeclaration()
		if parameter.DotDotDotToken != nil {
			if !ast.IsVarargParameter(parameter.AsNode()) && !ast.NodeIsMissing(parameter.Name()) {
				// The legacy TS named rest (`...args: T[]`). The parser builds the node
				// in full and leaves the rejection here: as a *parse* error it would
				// suppress every semantic diagnostic in the program, so one un-migrated
				// parameter would silently disable type checking.
				//
				// Reported with c.error, not grammarErrorOnNode: the latter goes quiet
				// in any file that has a parse error, and this is exactly the diagnostic
				// a migrating file needs. It is safe to report there because the node is
				// well-formed -- the parser parsed the whole parameter -- so it is not
				// recovery noise. A *missing* name is recovery noise, and is skipped.
				//
				// No early return either: the remaining rest-parameter rules still
				// apply, exactly as they did when this was a parse error.
				c.error(parameter.Name(), diagnostics.A_vararg_parameter_cannot_have_a_name)
			}
			if i != parameterCount-1 {
				return c.grammarErrorOnNode(parameter.DotDotDotToken, diagnostics.A_rest_parameter_must_be_last_in_a_parameter_list)
			}
			if parameter.Flags&ast.NodeFlagsAmbient == 0 {
				c.checkGrammarForDisallowedTrailingComma(parameters, diagnostics.A_rest_parameter_or_binding_pattern_may_not_have_a_trailing_comma)
			}

			if parameter.QuestionToken != nil {
				return c.grammarErrorOnNode(parameter.QuestionToken, diagnostics.A_rest_parameter_cannot_be_optional)
			}

			if parameter.Initializer != nil {
				// Anchored on the parameter, not its name: a vararg's name is a
				// zero-width synthetic identifier, which would give no span.
				return c.grammarErrorOnNode(parameter.AsNode(), diagnostics.A_rest_parameter_cannot_have_an_initializer)
			}
		} else if isOptionalDeclaration(parameter.AsNode()) {
			seenOptionalParameter = true
			// A reparsed '?' token indicates a bracketed name in @param tag
			if parameter.QuestionToken != nil && parameter.QuestionToken.Flags&ast.NodeFlagsReparsed == 0 && parameter.Initializer != nil {
				return c.grammarErrorOnNode(parameter.Name(), diagnostics.Parameter_cannot_have_question_mark_and_initializer)
			}
		} else if seenOptionalParameter && parameter.Initializer == nil {
			return c.grammarErrorOnNode(parameter.Name(), diagnostics.A_required_parameter_cannot_follow_an_optional_parameter)
		}
	}

	return false
}

func (c *Checker) checkGrammarForUseStrictSimpleParameterList(node *ast.Node) bool {
	if c.languageVersion >= core.ScriptTargetES2016 {
		body := node.Body()
		var useStrictDirective *ast.Node
		if body != nil && ast.IsBlock(body) {
			useStrictDirective = binder.FindUseStrictPrologue(ast.GetSourceFileOfNode(node), body.Statements())
		}
		if useStrictDirective != nil {
			nonSimpleParameters := core.Filter(node.Parameters(), func(n *ast.Node) bool {
				parameter := n.AsParameterDeclaration()
				return parameter.Initializer != nil || ast.IsBindingPattern(parameter.Name()) || isRestParameter(parameter.AsNode())
			})
			if len(nonSimpleParameters) != 0 {
				for _, parameter := range nonSimpleParameters {
					err := c.error(parameter, diagnostics.This_parameter_is_not_allowed_with_use_strict_directive)
					err.AddRelatedInfo(createDiagnosticForNode(useStrictDirective, diagnostics.X_use_strict_directive_used_here))
				}

				err := c.error(useStrictDirective, diagnostics.X_use_strict_directive_cannot_be_used_with_non_simple_parameter_list)
				for index, parameter := range nonSimpleParameters {
					var relatedMessage *diagnostics.Message
					if index == 0 {
						relatedMessage = diagnostics.Non_simple_parameter_declared_here
					} else {
						relatedMessage = diagnostics.X_and_here
					}
					err.AddRelatedInfo(createDiagnosticForNode(parameter, relatedMessage))
				}

				return true
			}
		}
	}
	return false
}

func (c *Checker) checkGrammarFunctionLikeDeclaration(node *ast.Node) bool {
	// Prevent cascading error by short-circuit
	file := ast.GetSourceFileOfNode(node)
	funcData := node.FunctionLikeData()
	return c.checkGrammarModifiers(node) || c.checkGrammarTypeParameterList(funcData.TypeParameters, file) ||
		c.checkGrammarParameterList(funcData.Parameters) || c.checkGrammarArrowFunction(node, file) ||
		(ast.IsFunctionLikeDeclaration(node) && c.checkGrammarForUseStrictSimpleParameterList(node))
}

func (c *Checker) checkGrammarArrowFunction(node *ast.Node, file *ast.SourceFile) bool {
	if !ast.IsArrowFunction(node) {
		return false
	}

	arrowFunc := node.AsArrowFunction()
	equalsGreaterThanToken := arrowFunc.EqualsGreaterThanToken
	startLine := scanner.GetECMALineOfPosition(file, equalsGreaterThanToken.Pos())
	endLine := scanner.GetECMALineOfPosition(file, equalsGreaterThanToken.End())
	return startLine != endLine && c.grammarErrorOnNode(equalsGreaterThanToken, diagnostics.Line_terminator_not_permitted_before_arrow)
}

func (c *Checker) checkGrammarIndexSignatureParameters(node *ast.IndexSignatureDeclaration) bool {
	paramNodes := node.Parameters.Nodes

	if len(paramNodes) == 0 {
		return c.grammarErrorOnNode(node.AsNode(), diagnostics.An_index_signature_must_have_exactly_one_parameter)
	}

	parameter := paramNodes[0].AsParameterDeclaration()
	if len(paramNodes) != 1 {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_must_have_exactly_one_parameter)
	}

	c.checkGrammarForDisallowedTrailingComma(node.Parameters, diagnostics.An_index_signature_cannot_have_a_trailing_comma)
	if parameter.DotDotDotToken != nil {
		return c.grammarErrorOnNode(parameter.DotDotDotToken, diagnostics.An_index_signature_cannot_have_a_rest_parameter)
	}
	if parameter.Modifiers() != nil {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_parameter_cannot_have_an_accessibility_modifier)
	}
	if parameter.QuestionToken != nil {
		return c.grammarErrorOnNode(parameter.QuestionToken, diagnostics.An_index_signature_parameter_cannot_have_a_question_mark)
	}
	if parameter.Initializer != nil {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_parameter_cannot_have_an_initializer)
	}
	typeNode := parameter.Type
	if typeNode == nil {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_parameter_must_have_a_type_annotation)
	}
	t := c.getTypeFromTypeNode(typeNode)
	// A literal key names a single property, which a named member or a mapped type
	// states precisely. TypeScript rejects a *generic* key here too; tlua does not,
	// because `[key: K]` is how a Lua table is typed.
	hasLiteralKey := false
	c.forEachIndexKeyType(t, func(keyType *Type) {
		if isLiteralIndexKeyType(keyType) {
			hasLiteralKey = true
		}
	})
	if hasLiteralKey {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_parameter_type_cannot_be_a_literal_type_Consider_using_a_mapped_object_type_instead)
	}
	if !everyType(t, c.isValidIndexKeyType) {
		return c.grammarErrorOnNode(parameter.Name(), diagnostics.An_index_signature_parameter_type_cannot_include_nil)
	}
	if node.Type == nil {
		return c.grammarErrorOnNode(node.AsNode(), diagnostics.An_index_signature_must_have_a_type_annotation)
	}
	return false
}

func (c *Checker) checkGrammarIndexSignature(node *ast.IndexSignatureDeclaration) bool {
	// Prevent cascading error by short-circuit
	return c.checkGrammarModifiers(node.AsNode()) || c.checkGrammarIndexSignatureParameters(node)
}

func (c *Checker) checkGrammarForAtLeastOneTypeArgument(node *ast.Node, typeArguments *ast.NodeList) bool {
	if typeArguments != nil && len(typeArguments.Nodes) == 0 {
		sourceFile := ast.GetSourceFileOfNode(node)
		start := typeArguments.Pos() - len("<")
		end := scanner.SkipTrivia(sourceFile.Text(), typeArguments.End()) + len(">")
		return c.grammarErrorAtPos(sourceFile.AsNode(), start, end-start, diagnostics.Type_argument_list_cannot_be_empty)
	}
	return false
}

func (c *Checker) checkGrammarTypeArguments(node *ast.Node, typeArguments *ast.NodeList) bool {
	return c.checkGrammarForDisallowedTrailingComma(typeArguments, diagnostics.Trailing_comma_not_allowed) || c.checkGrammarForAtLeastOneTypeArgument(node, typeArguments)
}

func (c *Checker) checkGrammarTaggedTemplateChain(node *ast.TaggedTemplateExpression) bool {
	if node.QuestionDotToken != nil || node.Flags&ast.NodeFlagsOptionalChain != 0 {
		return c.grammarErrorOnNode(node.Template, diagnostics.Tagged_template_expressions_are_not_permitted_in_an_optional_chain)
	}
	return false
}

func (c *Checker) checkGrammarHeritageClause(node *ast.HeritageClause) bool {
	types := node.Types
	if c.checkGrammarForDisallowedTrailingComma(types, diagnostics.Trailing_comma_not_allowed) {
		return true
	}
	if types != nil && len(types.Nodes) == 0 {
		listType := scanner.TokenToString(node.Token)
		// TODO(danielr): why not error on the token?
		return c.grammarErrorAtPos(node.AsNode(), types.Pos(), 0, diagnostics.X_0_list_cannot_be_empty, listType)
	}

	for _, node := range types.Nodes { //nolint:modernize
		if c.checkGrammarExpressionWithTypeArguments(node) {
			return true
		}
	}
	return false
}

func (c *Checker) checkGrammarExpressionWithTypeArguments(node *ast.Node /*Union[ExpressionWithTypeArguments, TypeQuery]*/) bool {
	if ast.IsExpressionWithTypeArguments(node) && node.Expression().Kind == ast.KindImportKeyword && node.TypeArgumentList() != nil {
		return c.grammarErrorOnNode(node, diagnostics.This_use_of_import_is_invalid_import_calls_can_be_written_but_they_must_have_parentheses_and_cannot_have_type_arguments)
	}
	return c.checkGrammarTypeArguments(node, node.TypeArgumentList())
}

func (c *Checker) checkGrammarInterfaceDeclaration(node *ast.InterfaceDeclaration) bool {
	if node.HeritageClauses != nil {
		seenExtendsClause := false
		for _, heritageClauseNode := range node.HeritageClauses.Nodes {
			heritageClause := heritageClauseNode.AsHeritageClause()

			switch heritageClause.Token {
			case ast.KindExtendsKeyword:
				if seenExtendsClause {
					return c.grammarErrorOnFirstToken(heritageClauseNode, diagnostics.X_extends_clause_already_seen)
				}
				seenExtendsClause = true
			case ast.KindImplementsKeyword:
				return c.grammarErrorOnFirstToken(heritageClauseNode, diagnostics.Interface_declaration_cannot_have_implements_clause)
			default:
				panic(fmt.Sprintf("Unexpected token %q", heritageClause.Token.String()))
			}

			// Grammar checking heritageClause inside class declaration
			c.checkGrammarHeritageClause(heritageClause)
		}
	}

	return false
}

func (c *Checker) checkGrammarComputedPropertyName(node *ast.Node) bool {
	// If node is not a computedPropertyName, just skip the grammar checking
	if node.Kind != ast.KindComputedPropertyName {
		return false
	}

	computedPropertyName := node.AsComputedPropertyName()
	if computedPropertyName.Expression.Kind == ast.KindBinaryExpression && computedPropertyName.Expression.AsBinaryExpression().OperatorToken.Kind == ast.KindCommaToken {
		return c.grammarErrorOnNode(computedPropertyName.Expression, diagnostics.A_comma_expression_is_not_allowed_in_a_computed_property_name)
	}
	return false
}

func (c *Checker) checkGrammarForInvalidQuestionMark(postfixToken *ast.TokenNode, message *diagnostics.Message) bool {
	return postfixToken != nil && postfixToken.Kind == ast.KindQuestionToken && c.grammarErrorOnNode(postfixToken, message)
}

func (c *Checker) checkGrammarForInvalidExclamationToken(postfixToken *ast.TokenNode, message *diagnostics.Message) bool {
	return postfixToken != nil && postfixToken.Kind == ast.KindExclamationToken && c.grammarErrorOnNode(postfixToken, message)
}

func (c *Checker) checkGrammarObjectLiteralExpression(node *ast.ObjectLiteralExpression) bool {
	seen := make(map[string]DeclarationMeaning)

	var properties []*ast.Node
	if node.Properties != nil {
		properties = node.Properties.Nodes
	}
	for _, prop := range properties {
		if prop.Kind == ast.KindTableEntry {
			// Lua positional entries have no name and no member grammar to
			// check; duplicate-index detection is the checker's job (last
			// write wins, like `{[1]: a, [1]: b}`).
			continue
		}
		name := prop.Name()
		if name.Kind == ast.KindComputedPropertyName {
			// If the name is not a ComputedPropertyName, the grammar checking will skip it
			c.checkGrammarComputedPropertyName(name)
		}

		// tlua: `{ x = 1 }` is a Lua keyed table field, so the upstream TS1312
		// "did you mean to use a ':'" grammar error is gone. There is no
		// destructuring reading -- destructuring assignment is deleted.

		if name.Kind == ast.KindPrivateIdentifier {
			c.grammarErrorOnNode(name, diagnostics.Private_identifiers_are_not_allowed_outside_class_bodies)
		}

		// Modifiers are never allowed on properties except for 'async' on a method declaration
		if modifiers := prop.ModifierNodes(); len(modifiers) != 0 {
			if ast.CanHaveModifiers(prop) {
				for _, mod := range modifiers {
					if ast.IsModifier(mod) && (mod.Kind != ast.KindAsyncKeyword || prop.Kind != ast.KindMethodDeclaration) {
						c.grammarErrorOnNode(mod, diagnostics.X_0_modifier_cannot_be_used_here, scanner.GetTextOfNode(mod))
					}
				}
			} else if ast.CanHaveIllegalModifiers(prop) {
				for _, mod := range modifiers {
					if ast.IsModifier(mod) {
						c.grammarErrorOnNode(mod, diagnostics.X_0_modifier_cannot_be_used_here, scanner.GetTextOfNode(mod))
					}
				}
			}
		}

		// ECMA-262 11.1.5 Object Initializer
		// If previous is not undefined then throw a SyntaxError exception if any of the following conditions are true
		// a.This production is contained in strict code and IsDataDescriptor(previous) is true and
		// IsDataDescriptor(propId.descriptor) is true.
		//    b.IsDataDescriptor(previous) is true and IsAccessorDescriptor(propId.descriptor) is true.
		//    c.IsAccessorDescriptor(previous) is true and IsDataDescriptor(propId.descriptor) is true.
		//    d.IsAccessorDescriptor(previous) is true and IsAccessorDescriptor(propId.descriptor) is true
		// and either both previous and propId.descriptor have[[Get]] fields or both previous and propId.descriptor have[[Set]] fields
		var currentKind DeclarationMeaning
		switch prop.Kind {
		case ast.KindShorthandPropertyAssignment,
			ast.KindPropertyAssignment:
			var commonProp *ast.NamedMemberBase
			if prop.Kind == ast.KindShorthandPropertyAssignment {
				prop.ClassLikeData()
				commonProp = &prop.AsShorthandPropertyAssignment().NamedMemberBase
			} else {
				commonProp = &prop.AsPropertyAssignment().NamedMemberBase
			}

			// Grammar checking for computedPropertyName and shorthandPropertyAssignment
			c.checkGrammarForInvalidExclamationToken(commonProp.PostfixToken, diagnostics.A_definite_assignment_assertion_is_not_permitted_in_this_context)
			c.checkGrammarForInvalidQuestionMark(commonProp.PostfixToken, diagnostics.An_object_member_cannot_be_declared_optional)

			if name.Kind == ast.KindNumericLiteral {
				c.checkGrammarNumericLiteral(name.AsNumericLiteral())
			}

			currentKind = DeclarationMeaningPropertyAssignment
		case ast.KindMethodDeclaration:
			currentKind = DeclarationMeaningMethod
		case ast.KindGetAccessor:
			currentKind = DeclarationMeaningGetAccessor
		case ast.KindSetAccessor:
			currentKind = DeclarationMeaningSetAccessor
		default:
			panic(fmt.Sprintf("Unexpected node kind %q", prop.Kind))
		}

		{
			effectiveName, ok := c.getEffectivePropertyNameForPropertyNameNode(name)
			if !ok {
				continue
			}

			existingKind := seen[effectiveName]
			if existingKind == 0 {
				seen[effectiveName] = currentKind
			} else {
				if (currentKind&DeclarationMeaningMethod != 0) && (existingKind&DeclarationMeaningMethod != 0) {
					c.grammarErrorOnNode(name, diagnostics.Duplicate_identifier_0, scanner.GetTextOfNode(name))
				} else if (currentKind&DeclarationMeaningPropertyAssignment != 0) && (existingKind&DeclarationMeaningPropertyAssignment != 0) {
					c.grammarErrorOnNode(name, diagnostics.An_object_literal_cannot_have_multiple_properties_with_the_same_name, scanner.GetTextOfNode(name))
				} else if (currentKind&DeclarationMeaningGetOrSetAccessor != 0) && (existingKind&DeclarationMeaningGetOrSetAccessor != 0) {
					if existingKind != DeclarationMeaningGetOrSetAccessor && currentKind != existingKind {
						seen[effectiveName] = currentKind | existingKind
					} else {
						return c.grammarErrorOnNode(name, diagnostics.An_object_literal_cannot_have_multiple_get_Slashset_accessors_with_the_same_name)
					}
				} else {
					return c.grammarErrorOnNode(name, diagnostics.An_object_literal_cannot_have_property_and_accessor_with_the_same_name)
				}
			}
		}
	}

	return false
}

func (c *Checker) checkGrammarJsxElement(node *ast.Node) bool {
	c.checkGrammarJsxName(node.TagName())
	c.checkGrammarTypeArguments(node, node.TypeArgumentList())
	var seen collections.Set[string]
	for _, attrNode := range node.Attributes().Properties() {
		// Every attribute is a plain one: spread attributes no longer parse.
		attr := attrNode.AsJsxAttribute()
		name := attr.Name()
		initializer := attr.Initializer
		textOfName := name.Text()
		if !seen.Has(textOfName) {
			seen.Add(textOfName)
		} else {
			return c.grammarErrorOnNode(name, diagnostics.JSX_elements_cannot_have_multiple_attributes_with_the_same_name)
		}
		if initializer != nil && initializer.Kind == ast.KindJsxExpression && initializer.Expression() == nil {
			return c.grammarErrorOnNode(initializer, diagnostics.JSX_attributes_must_only_be_assigned_a_non_empty_expression)
		}
	}
	return false
}

func (c *Checker) checkGrammarJsxName(node *ast.JsxTagNameExpression) bool {
	if ast.IsPropertyAccessExpression(node) && ast.IsJsxNamespacedName(node.Expression()) {
		return c.grammarErrorOnNode(node.Expression(), diagnostics.JSX_property_access_expressions_cannot_include_JSX_namespace_names)
	}

	if ast.IsJsxNamespacedName(node) && c.compilerOptions.GetJSXTransformEnabled() && !scanner.IsIntrinsicJsxName(node.AsJsxNamespacedName().Namespace.Text()) {
		return c.grammarErrorOnNode(node, diagnostics.React_components_cannot_include_JSX_namespace_names)
	}

	return false
}

func (c *Checker) checkGrammarJsxExpression(node *ast.JsxExpression) bool {
	if node.Expression != nil && ast.IsCommaSequence(node.Expression) {
		return c.grammarErrorOnNode(node.Expression, diagnostics.JSX_expressions_may_not_use_the_comma_operator_Did_you_mean_to_write_an_array)
	}

	return false
}

// checkGrammarForOfStatement grammar-checks a Lua generic-for (the only form a
// ForOfStatement now takes). Its header allows multiple names and type
// annotations and its names never carry initializers, so the only grammar
// constraint is the ambient-context rule shared by every statement.
func (c *Checker) checkGrammarForOfStatement(forOfStatement *ast.ForOfStatement) bool {
	return c.checkGrammarStatementInAmbientContext(forOfStatement.AsNode())
}

func (c *Checker) checkGrammarAccessor(accessor *ast.AccessorDeclaration) bool {
	body := accessor.Body()
	if accessor.Flags&ast.NodeFlagsAmbient == 0 && (accessor.Parent.Kind != ast.KindTypeLiteral) && (accessor.Parent.Kind != ast.KindInterfaceDeclaration) {
		if body == nil && !ast.HasSyntacticModifier(accessor, ast.ModifierFlagsAbstract) {
			return c.grammarErrorAtPos(accessor, accessor.End()-1, len(";"), diagnostics.X_0_expected, "{")
		}
	}
	if body != nil {
		if ast.HasSyntacticModifier(accessor, ast.ModifierFlagsAbstract) {
			return c.grammarErrorOnNode(accessor, diagnostics.An_abstract_accessor_cannot_have_an_implementation)
		}
		if accessor.Parent.Kind == ast.KindTypeLiteral || accessor.Parent.Kind == ast.KindInterfaceDeclaration {
			return c.grammarErrorOnNode(body, diagnostics.An_implementation_cannot_be_declared_in_ambient_contexts)
		}
	}

	funcData := accessor.FunctionLikeData()
	var typeParameters *ast.NodeList
	if funcData != nil {
		typeParameters = funcData.TypeParameters
	}

	if typeParameters != nil {
		return c.grammarErrorOnNode(accessor.Name(), diagnostics.An_accessor_cannot_have_type_parameters)
	}
	if !c.doesAccessorHaveCorrectParameterCount(accessor) {
		return c.grammarErrorOnNode(accessor.Name(), core.IfElse(accessor.Kind == ast.KindGetAccessor, diagnostics.A_get_accessor_cannot_have_parameters, diagnostics.A_set_accessor_must_have_exactly_one_parameter))
	}
	if accessor.Kind == ast.KindSetAccessor {
		if funcData.Type != nil {
			return c.grammarErrorOnNode(accessor.Name(), diagnostics.A_set_accessor_cannot_have_a_return_type_annotation)
		}

		parameterNode := GetSetAccessorValueParameter(accessor)
		if parameterNode == nil {
			panic("Return value does not match parameter count assertion.")
		}
		parameter := parameterNode.AsParameterDeclaration()
		if parameter.DotDotDotToken != nil {
			return c.grammarErrorOnNode(parameter.DotDotDotToken, diagnostics.A_set_accessor_cannot_have_rest_parameter)
		}
		if parameter.QuestionToken != nil {
			return c.grammarErrorOnNode(parameter.QuestionToken, diagnostics.A_set_accessor_cannot_have_an_optional_parameter)
		}
		if parameter.Initializer != nil {
			return c.grammarErrorOnNode(accessor.Name(), diagnostics.A_set_accessor_parameter_cannot_have_an_initializer)
		}
	}

	return false
}

// Does the accessor have the right number of parameters?
//
//	A `get` accessor has no parameters or a single `this` parameter.
//	A `set` accessor has one parameter or a `this` parameter and one more parameter.
func (c *Checker) doesAccessorHaveCorrectParameterCount(accessor *ast.AccessorDeclaration) bool {
	// `getAccessorThisParameter` returns `nil` if the accessor's arity is incorrect,
	// even if there is a `this` parameter declared.
	return c.getAccessorThisParameter(accessor) != nil || len(accessor.Parameters()) == core.IfElse(accessor.Kind == ast.KindGetAccessor, 0, 1)
}

func (c *Checker) checkGrammarTypeOperatorNode(node *ast.TypeOperatorNode) bool {
	if node.Operator == ast.KindUniqueKeyword {
		innerType := node.Type
		if innerType.Kind != ast.KindSymbolKeyword {
			return c.grammarErrorOnNode(innerType, diagnostics.X_0_expected, scanner.TokenToString(ast.KindSymbolKeyword))
		}
		parent := ast.WalkUpParenthesizedTypes(node.Parent)
		switch parent.Kind {
		case ast.KindVariableDeclaration:
			decl := parent.AsVariableDeclaration()
			if decl.Name().Kind != ast.KindIdentifier {
				return c.grammarErrorOnNode(node.AsNode(), diagnostics.X_unique_symbol_types_may_not_be_used_on_a_variable_declaration_with_a_binding_name)
			}
			if !isVariableDeclarationInVariableStatement(decl.AsNode()) {
				return c.grammarErrorOnNode(node.AsNode(), diagnostics.X_unique_symbol_types_are_only_allowed_on_variables_in_a_variable_statement)
			}
			if decl.Parent.Flags&ast.NodeFlagsConst == 0 {
				return c.grammarErrorOnNode(parent.AsVariableDeclaration().Name(), diagnostics.A_variable_whose_type_is_a_unique_symbol_type_must_be_const)
			}
		case ast.KindPropertyDeclaration:
			if !ast.IsStatic(parent) || !hasReadonlyModifier(parent) {
				return c.grammarErrorOnNode(parent.AsPropertyDeclaration().Name(), diagnostics.A_property_of_a_class_whose_type_is_a_unique_symbol_type_must_be_both_static_and_readonly)
			}
		case ast.KindPropertySignature:
			if !ast.HasSyntacticModifier(parent, ast.ModifierFlagsReadonly) {
				return c.grammarErrorOnNode(parent.AsPropertySignatureDeclaration().Name(), diagnostics.A_property_of_an_interface_or_type_literal_whose_type_is_a_unique_symbol_type_must_be_readonly)
			}
		default:
			return c.grammarErrorOnNode(node.AsNode(), diagnostics.X_unique_symbol_types_are_not_allowed_here)
		}
	} else if node.Operator == ast.KindReadonlyKeyword {
		innerType := node.Type
		if innerType.Kind != ast.KindArrayType && innerType.Kind != ast.KindTupleType {
			return c.grammarErrorOnFirstToken(node.AsNode(), diagnostics.X_readonly_type_modifier_is_only_permitted_on_array_and_tuple_literal_types, scanner.TokenToString(ast.KindSymbolKeyword))
		}
	}

	return false
}

func (c *Checker) checkGrammarForInvalidDynamicName(node *ast.DeclarationName, message *diagnostics.Message) bool {
	if !c.isNonBindableDynamicName(node) {
		return false
	}
	var expression *ast.Node
	if ast.IsElementAccessExpression(node) {
		expression = ast.SkipParentheses(node.AsElementAccessExpression().ArgumentExpression)
	} else {
		expression = node.Expression()
	}

	if !ast.IsEntityNameExpression(expression) {
		return c.grammarErrorOnNode(node, message)
	}

	return false
}

// Indicates whether a declaration name is a dynamic name that cannot be late-bound.
func (c *Checker) isNonBindableDynamicName(node *ast.DeclarationName) bool {
	return ast.IsDynamicName(node) && !c.isLateBindableName(node)
}

func (c *Checker) checkGrammarMethod(node *ast.Node /*Union[MethodDeclaration, MethodSignature]*/) bool {
	if c.checkGrammarFunctionLikeDeclaration(node) {
		return true
	}

	if node.Kind == ast.KindMethodDeclaration {
		if node.Parent.Kind == ast.KindObjectLiteralExpression {
			// We only disallow modifier on a method declaration if it is a property of object-literal-expression
			if modifiers := node.Modifiers(); modifiers != nil && !(len(modifiers.Nodes) == 1 && modifiers.Nodes[0].Kind == ast.KindAsyncKeyword) {
				return c.grammarErrorOnFirstToken(node, diagnostics.Modifiers_cannot_appear_here)
			}

			methodDecl := node.AsMethodDeclaration()
			if c.checkGrammarForInvalidQuestionMark(methodDecl.PostfixToken, diagnostics.An_object_member_cannot_be_declared_optional) {
				return true
			}
			if c.checkGrammarForInvalidExclamationToken(methodDecl.PostfixToken, diagnostics.A_definite_assignment_assertion_is_not_permitted_in_this_context) {
				return true
			}
			if node.Body() == nil {
				return c.grammarErrorAtPos(node, node.End()-1, len(";"), diagnostics.X_0_expected, "{")
			}
		}
	}

	if ast.IsClassLike(node.Parent) {
		// Technically, computed properties in ambient contexts is disallowed
		// for property declarations and accessors too, not just methods.
		// However, property declarations disallow computed names in general,
		// and accessors are not allowed in ambient contexts in general,
		// so this error only really matters for methods.
		if node.Flags&ast.NodeFlagsAmbient != 0 {
			return c.checkGrammarForInvalidDynamicName(node.Name(), diagnostics.A_computed_property_name_in_an_ambient_context_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type)
		} else if node.Kind == ast.KindMethodDeclaration && node.Body() == nil {
			return c.checkGrammarForInvalidDynamicName(node.Name(), diagnostics.A_computed_property_name_in_a_method_overload_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type)
		}
	} else if node.Parent.Kind == ast.KindInterfaceDeclaration {
		return c.checkGrammarForInvalidDynamicName(node.Name(), diagnostics.A_computed_property_name_in_an_interface_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type)
	} else if node.Parent.Kind == ast.KindTypeLiteral {
		return c.checkGrammarForInvalidDynamicName(node.Name(), diagnostics.A_computed_property_name_in_a_type_literal_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type)
	}

	return false
}

// break and continue carry no label in tlua, so they always bind to the
// innermost enclosing loop. A `goto` is the only labeled jump; see
// checkGrammarGotoStatement.
func (c *Checker) checkGrammarBreakOrContinueStatement(node *ast.Node) bool {
	var current *ast.Node = node
	for current != nil {
		if ast.IsFunctionLikeOrClassStaticBlockDeclaration(current) {
			return c.grammarErrorOnNode(node, diagnostics.Jump_target_cannot_cross_function_boundary)
		}

		if ast.IsIterationStatement(current) {
			// break or continue within iteration statement - ok
			return false
		}

		current = current.Parent
	}

	message := diagnostics.A_continue_statement_can_only_be_used_within_an_enclosing_iteration_statement
	if node.Kind == ast.KindBreakStatement {
		message = diagnostics.A_break_statement_can_only_be_used_within_an_enclosing_iteration_statement
	}
	return c.grammarErrorOnNode(node, message)
}

// A label is visible throughout the block that declares it, nested blocks
// included, but never across a function or module boundary — each compiles to
// its own Lua function — so an unresolved goto and a goto that names a label
// outside that boundary report the same thing, as in Lua.
func (c *Checker) checkGrammarGotoStatement(node *ast.Node) bool {
	if c.checkGrammarStatementInAmbientContext(node) {
		return true
	}
	target := node.Label()
	label := ast.FindTargetLabel(node, target.Text())
	if label == nil {
		return c.grammarErrorOnNode(target, diagnostics.No_visible_label_0_for_goto, target.Text())
	}
	return c.checkGrammarGotoIntoLocalScope(node, label)
}

// Lua forbids jumping into the scope of a local. Only a forward jump can enter
// one; a backward jump leaves scopes. A local's scope ends at the last non-void
// statement of its block, so a label in the trailing run of labels and empty
// statements is always jumpable — that relaxation is what makes the
// `goto continue` idiom legal past a body's locals.
func (c *Checker) checkGrammarGotoIntoLocalScope(node *ast.Node, label *ast.Node) bool {
	// A label is a direct statement of its container, and the resolution walk
	// found it in one of the goto's ancestors, so the goto has an ancestor that
	// is a direct statement of the same list.
	container := label.Parent
	statements := ast.StatementsOfLabelContainer(container)
	debug.Assert(statements != nil, "a resolved label's parent must scope labels")
	child := node
	for child.Parent != container {
		child = child.Parent
	}
	gotoIndex, labelIndex := -1, -1
	for i, stmt := range statements {
		if stmt == child {
			gotoIndex = i
		}
		if stmt == label {
			labelIndex = i
		}
		if gotoIndex >= 0 && labelIndex >= 0 {
			break
		}
	}
	debug.Assert(gotoIndex >= 0 && labelIndex >= 0, "goto and label statements not found in their block")
	if labelIndex <= gotoIndex {
		return false
	}
	// A label followed only by labels and empty statements sits past the end of
	// every local's scope, so it is always jumpable.
	trailing := true
	for _, stmt := range statements[labelIndex+1:] {
		if stmt.Kind != ast.KindLabelStatement && stmt.Kind != ast.KindEmptyStatement {
			trailing = false
			break
		}
	}
	if trailing {
		return false
	}
	// Only locals declared directly in this block are still in scope at the
	// label; one inside a nested block between the two died with its block.
	for _, stmt := range statements[gotoIndex+1 : labelIndex] {
		if name := jumpedOverLocalName(stmt); name != nil {
			return c.grammarErrorOnNode(node, diagnostics.Cannot_jump_into_the_scope_of_local_0, name.Text())
		}
	}
	return false
}

// The name of the first local a statement declares, or nil if it declares none.
// A `local a, b` reports 'a' — all bindings' scopes are entered equally, and
// Lua's own message names the first — but the scan skips a nil name (which parse
// recovery can leave on a malformed declaration) so a later real name still
// triggers the error.
func jumpedOverLocalName(node *ast.Node) *ast.Node {
	switch node.Kind {
	case ast.KindVariableStatement:
		// Every variable statement counts, not just `local`: the transitional
		// let/const/var forms declare names that compile to Lua locals, so
		// jumping into their scope is the same error.
		list := node.AsVariableStatement().DeclarationList
		for _, decl := range list.AsVariableDeclarationList().Declarations.Nodes {
			if name := decl.Name(); name != nil {
				return name
			}
		}
		return nil
	case ast.KindFunctionDeclaration:
		if !ast.IsLuaLocal(node) {
			return nil
		}
		return node.Name()
	}
	return nil
}

func (c *Checker) checkGrammarBindingElement(node *ast.BindingElement) bool {
	if node.DotDotDotToken != nil {
		elements := node.Parent.ElementList()
		if node.AsNode() != core.LastOrNil(elements.Nodes) {
			return c.grammarErrorOnNode(&node.Node, diagnostics.A_rest_element_must_be_last_in_a_destructuring_pattern)
		}
		c.checkGrammarForDisallowedTrailingComma(elements, diagnostics.A_rest_parameter_or_binding_pattern_may_not_have_a_trailing_comma)

		if node.PropertyName != nil {
			return c.grammarErrorOnNode(node.Name(), diagnostics.A_rest_element_cannot_have_a_property_name)
		}
	}

	if node.DotDotDotToken != nil && node.Initializer != nil {
		// Error on equals token which immediately precedes the initializer
		return c.grammarErrorAtPos(node.AsNode(), node.Initializer.Pos()-1, 1, diagnostics.A_rest_element_cannot_have_an_initializer)
	}

	return false
}

func (c *Checker) checkGrammarVariableDeclaration(node *ast.VariableDeclaration) bool {
	nodeFlags := c.getCombinedNodeFlagsCached(node.AsNode())
	blockScopeKind := nodeFlags & ast.NodeFlagsBlockScoped
	if node.Parent.Parent.Kind != ast.KindForOfStatement {
		if nodeFlags&ast.NodeFlagsAmbient != 0 {
			c.checkAmbientInitializer(node.AsNode())
		} else if node.Initializer == nil {
			if ast.IsBindingPattern(node.Name()) && !ast.IsBindingPattern(node.Parent) {
				return c.grammarErrorOnNode(node.AsNode(), diagnostics.A_destructuring_declaration_must_have_an_initializer)
			}
			if blockScopeKind == ast.NodeFlagsConst {
				return c.grammarErrorOnNode(node.AsNode(), diagnostics.X_0_declarations_must_be_initialized, "const")
			}
		}
	}

	if c.program.GetEmitModuleFormatOfFile(ast.GetSourceFileOfNode(node.AsNode())) < core.ModuleKindSystem && (node.Parent.Parent.Flags&ast.NodeFlagsAmbient == 0) && ast.HasSyntacticModifier(node.Parent.Parent, ast.ModifierFlagsExport) {
		c.checkGrammarForEsModuleMarkerInBindingName(node.Name())
	}

	// The ES "let is not a valid LexicalDeclaration name" rule died with the
	// let/const surface syntax: `let` is an ordinary identifier, legal as any
	// binding name (including for-loop control variables, whose lists carry
	// the internal Let scope marker).
	return false
}

func (c *Checker) checkGrammarForEsModuleMarkerInBindingName(name *ast.Node) bool {
	if ast.IsIdentifier(name) {
		if name.Text() == "__esModule" {
			return c.grammarErrorOnNodeSkippedOnNoEmit(name, diagnostics.Identifier_expected_esModule_is_reserved_as_an_exported_marker_when_transforming_ECMAScript_modules)
		}
	} else {
		for _, element := range name.Elements() {
			if element.Name() != nil {
				return c.checkGrammarForEsModuleMarkerInBindingName(element.Name())
			}
		}
	}
	return false
}

func (c *Checker) checkGrammarVariableDeclarationList(declarationList *ast.VariableDeclarationList) bool {
	declarations := declarationList.Declarations
	if c.checkGrammarForDisallowedTrailingComma(declarations, diagnostics.Trailing_comma_not_allowed) {
		return true
	}

	if len(declarations.Nodes) == 0 {
		return c.grammarErrorAtPos(declarationList.AsNode(), declarations.Pos(), declarations.End()-declarations.Pos(), diagnostics.Variable_declaration_list_cannot_be_empty)
	}

	return false
}

func (c *Checker) checkGrammarForDisallowedBlockScopedVariableStatement(node *ast.VariableStatement) bool {
	if !c.containerAllowsBlockScopedVariable(node.Parent) {
		// Every parsed declaration form is position-restricted, and a plain
		// `local` list carries no block-scope bit (only NodeFlagsLuaLocal), so
		// the statement itself decides, not the flags. Without this a local
		// declared as an unbraced if/while/for body would bind into the
		// ENCLOSING scope: visible (and typed) after the statement while the
		// runtime value never escapes the body.
		blockScopeKind := c.getCombinedNodeFlagsCached(node.DeclarationList) & ast.NodeFlagsBlockScoped
		var keyword string
		switch {
		case ast.IsLuaLocal(node.DeclarationList) || blockScopeKind == ast.NodeFlagsLet:
			keyword = "local"
		case blockScopeKind == ast.NodeFlagsConst:
			// The ambient-global desugaring (`declare name: Type`).
			keyword = "declare"
		default:
			panic("Unknown BlockScope flag")
		}
		c.error(node.AsNode(), diagnostics.X_0_declarations_can_only_be_declared_inside_a_block, keyword)
	}

	return false
}

func (c *Checker) containerAllowsBlockScopedVariable(parent *ast.Node) bool {
	switch parent.Kind {
	case ast.KindIfStatement,
		ast.KindDoStatement,
		ast.KindWhileStatement,
		ast.KindWithStatement,
		ast.KindForOfStatement:
		return false
	}

	return true
}

func (c *Checker) checkGrammarMetaProperty(node *ast.MetaProperty) bool {
	nodeName := node.Name()
	nameText := nodeName.Text()

	switch node.KeywordToken {
	case ast.KindNewKeyword:
		if nameText != "target" {
			return c.grammarErrorOnNode(nodeName, diagnostics.X_0_is_not_a_valid_meta_property_for_keyword_1_Did_you_mean_2, nameText, scanner.TokenToString(node.KeywordToken), "target")
		}
	case ast.KindImportKeyword:
		if nameText != "meta" {
			isCallee := ast.IsCallExpression(node.Parent) && node.Parent.Expression() == node.AsNode()
			if nameText == "defer" {
				if !isCallee {
					return c.grammarErrorAtPos(node.AsNode(), node.AsNode().End(), 0, diagnostics.X_0_expected, "(")
				}
			} else {
				if isCallee {
					return c.grammarErrorOnNode(nodeName, diagnostics.X_0_is_not_a_valid_meta_property_for_keyword_import_Did_you_mean_meta_or_defer, nameText)
				}
				return c.grammarErrorOnNode(nodeName, diagnostics.X_0_is_not_a_valid_meta_property_for_keyword_1_Did_you_mean_2, nameText, scanner.TokenToString(node.KeywordToken), "meta")
			}
		}
	}

	return false
}

func (c *Checker) checkGrammarConstructorTypeParameters(node *ast.ConstructorDeclaration) bool {
	range_ := node.TypeParameters
	if range_ != nil {
		var pos int
		if range_.Pos() == range_.End() {
			pos = range_.Pos()
		} else {
			pos = scanner.SkipTrivia(ast.GetSourceFileOfNode(node.AsNode()).Text(), range_.Pos())
		}
		return c.grammarErrorAtPos(node.AsNode(), pos, range_.End()-pos, diagnostics.Type_parameters_cannot_appear_on_a_constructor_declaration)
	}

	return false
}

func (c *Checker) checkGrammarConstructorTypeAnnotation(node *ast.ConstructorDeclaration) bool {
	t := node.Type
	if t != nil {
		return c.grammarErrorOnNode(t, diagnostics.Type_annotation_cannot_appear_on_a_constructor_declaration)
	}
	return false
}

func (c *Checker) checkGrammarProperty(node *ast.Node /*Union[PropertyDeclaration, PropertySignature]*/) bool {
	propertyName := node.Name()
	if ast.IsComputedPropertyName(propertyName) && ast.IsBinaryExpression(propertyName.Expression()) && propertyName.Expression().AsBinaryExpression().OperatorToken.Kind == ast.KindInKeyword {
		return c.grammarErrorOnNode(node.Parent.Members()[0], diagnostics.A_mapped_type_may_not_declare_properties_or_methods)
	}
	if ast.IsClassLike(node.Parent) {
		if ast.IsStringLiteral(propertyName) && propertyName.Text() == "constructor" {
			return c.grammarErrorOnNode(propertyName, diagnostics.Classes_may_not_have_a_field_named_constructor)
		}
		if c.checkGrammarForInvalidDynamicName(propertyName, diagnostics.A_computed_property_name_in_a_class_property_declaration_must_have_a_simple_literal_type_or_a_unique_symbol_type) {
			return true
		}
		if ast.IsAutoAccessorPropertyDeclaration(node) && c.checkGrammarForInvalidQuestionMark(node.PostfixToken(), diagnostics.An_accessor_property_cannot_be_declared_optional) {
			return true
		}
	} else if ast.IsInterfaceDeclaration(node.Parent) {
		if c.checkGrammarForInvalidDynamicName(propertyName, diagnostics.A_computed_property_name_in_an_interface_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type) {
			return true
		}
		if !ast.IsPropertySignatureDeclaration(node) {
			// Interfaces cannot contain property declarations
			panic(fmt.Sprintf("Unexpected node kind %q", node.Kind))
		}
		if initializer := node.Initializer(); initializer != nil {
			return c.grammarErrorOnNode(initializer, diagnostics.An_interface_property_cannot_have_an_initializer)
		}
	} else if ast.IsTypeLiteralNode(node.Parent) {
		if c.checkGrammarForInvalidDynamicName(node.Name(), diagnostics.A_computed_property_name_in_a_type_literal_must_refer_to_an_expression_whose_type_is_a_literal_type_or_a_unique_symbol_type) {
			return true
		}
		if !ast.IsPropertySignatureDeclaration(node) {
			// Type literals cannot contain property declarations
			panic(fmt.Sprintf("Unexpected node kind %q", node.Kind))
		}
		if initializer := node.Initializer(); initializer != nil {
			return c.grammarErrorOnNode(initializer, diagnostics.A_type_literal_property_cannot_have_an_initializer)
		}
	}

	if node.Flags&ast.NodeFlagsAmbient != 0 {
		c.checkAmbientInitializer(node)
	}

	if ast.IsPropertyDeclaration(node) {
		propDecl := node.AsPropertyDeclaration()
		postfixToken := propDecl.PostfixToken
		if postfixToken != nil && postfixToken.Kind == ast.KindExclamationToken {
			switch {
			case propDecl.Initializer != nil:
				return c.grammarErrorOnNode(postfixToken, diagnostics.Declarations_with_initializers_cannot_also_have_definite_assignment_assertions)
			case propDecl.Type == nil:
				return c.grammarErrorOnNode(postfixToken, diagnostics.Declarations_with_definite_assignment_assertions_must_also_have_type_annotations)
			case !ast.IsClassLike(node.Parent) || node.Flags&ast.NodeFlagsAmbient != 0 || ast.IsStatic(node) || ast.HasAbstractModifier(node):
				return c.grammarErrorOnNode(postfixToken, diagnostics.A_definite_assignment_assertion_is_not_permitted_in_this_context)
			}
		}
	}

	return false
}

func (c *Checker) checkAmbientInitializer(node *ast.Node) bool {
	var initializer *ast.Expression
	var typeNode *ast.TypeNode
	switch node.Kind {
	case ast.KindVariableDeclaration:
		varDecl := node.AsVariableDeclaration()
		initializer = varDecl.Initializer
		typeNode = varDecl.Type
	case ast.KindPropertyDeclaration:
		propDecl := node.AsPropertyDeclaration()
		initializer = propDecl.Initializer
		typeNode = propDecl.Type
	case ast.KindPropertySignature:
		propSig := node.AsPropertySignatureDeclaration()
		initializer = propSig.Initializer
		typeNode = propSig.Type
	default:
		panic(fmt.Sprintf("Unexpected node kind %q", node.Kind))
	}

	if initializer != nil {
		isInvalidInitializer := !(isInitializerStringOrNumberLiteralExpression(initializer) || initializer.Kind == ast.KindTrueKeyword || initializer.Kind == ast.KindFalseKeyword)
		isConstOrReadonly := isDeclarationReadonly(node) || ast.IsVariableDeclaration(node) && c.isVarConst(node)
		if isConstOrReadonly && (typeNode == nil) {
			if isInvalidInitializer {
				return c.grammarErrorOnNode(initializer, diagnostics.A_const_initializer_in_an_ambient_context_must_be_a_string_numeric_or_boolean_literal)
			}
		} else {
			return c.grammarErrorOnNode(initializer, diagnostics.Initializers_are_not_allowed_in_ambient_contexts)
		}
	}

	return false
}

func isInitializerStringOrNumberLiteralExpression(expr *ast.Expression) bool {
	return ast.IsStringOrNumericLiteralLike(expr) ||
		expr.Kind == ast.KindPrefixUnaryExpression && expr.AsPrefixUnaryExpression().Operator == ast.KindMinusToken && expr.AsPrefixUnaryExpression().Operand.Kind == ast.KindNumericLiteral
}

func (c *Checker) checkGrammarTopLevelElementForRequiredDeclareModifier(node *ast.Node) bool {
	// A declare modifier is required for any top level .d.ts declaration except export=, export default, export as namespace
	// interfaces and imports categories:
	//
	//  DeclarationElement:
	//     ExportAssignment
	//     export_opt   InterfaceDeclaration
	//     export_opt   TypeAliasDeclaration
	//     export_opt   ImportDeclaration
	//     export_opt   ExternalImportDeclaration
	//     export_opt   AmbientDeclaration
	//
	// TODO: The spec needs to be amended to reflect this grammar.
	if node.Kind == ast.KindInterfaceDeclaration || node.Kind == ast.KindTypeAliasDeclaration || node.Kind == ast.KindNamespaceExportDeclaration || ast.HasSyntacticModifier(node, ast.ModifierFlagsAmbient|ast.ModifierFlagsExport|ast.ModifierFlagsDefault) {
		return false
	}

	return c.grammarErrorOnFirstToken(node, diagnostics.Top_level_declarations_in_d_ts_files_must_start_with_either_a_declare_or_export_modifier)
}

func (c *Checker) checkGrammarTopLevelElementsForRequiredDeclareModifier(file *ast.SourceFile) bool {
	for _, decl := range file.Statements.Nodes {
		if ast.IsDeclarationNode(decl) || decl.Kind == ast.KindVariableStatement {
			if c.checkGrammarTopLevelElementForRequiredDeclareModifier(decl) {
				return true
			}
		}
	}
	return false
}

func (c *Checker) checkGrammarSourceFile(node *ast.SourceFile) bool {
	return node.Flags&ast.NodeFlagsAmbient != 0 && c.checkGrammarTopLevelElementsForRequiredDeclareModifier(node)
}

func (c *Checker) checkGrammarStatementInAmbientContext(node *ast.Node) bool {
	if node.Flags&ast.NodeFlagsAmbient != 0 {
		// Find containing block which is either Block, ModuleBlock, SourceFile
		links := c.nodeLinks.Get(node)
		if !links.hasReportedStatementInAmbientContext && (ast.IsFunctionLike(node.Parent) || ast.IsAccessor(node.Parent)) {
			links.hasReportedStatementInAmbientContext = c.grammarErrorOnFirstToken(node, diagnostics.An_implementation_cannot_be_declared_in_ambient_contexts)
			return links.hasReportedStatementInAmbientContext
		}

		// We are either parented by another statement, or some sort of block.
		// If we're in a block, we only want to really report an error once
		// to prevent noisiness.  So use a bit on the block to indicate if
		// this has already been reported, and don't report if it has.
		//
		if node.Parent.Kind == ast.KindBlock || node.Parent.Kind == ast.KindModuleBlock || node.Parent.Kind == ast.KindSourceFile {
			links := c.nodeLinks.Get(node.Parent)
			// Check if the containing block ever report this error
			if !links.hasReportedStatementInAmbientContext {
				links.hasReportedStatementInAmbientContext = c.grammarErrorOnFirstToken(node, diagnostics.Statements_are_not_allowed_in_ambient_contexts)
				return links.hasReportedStatementInAmbientContext
			}
		} else {
			// We must be parented by a statement.  If so, there's no need
			// to report the error as our parent will have already done it.
			// debug.Assert(ast.IsStatement(node.Parent)) // !!! commented out in strada - fails if uncommented
		}
	}
	return false
}

func (c *Checker) checkGrammarNumericLiteral(node *ast.NumericLiteral) {
	nodeText := scanner.GetTextOfNode(node.AsNode())

	// Realism (size) checking
	// We should test against `getTextOfNode(node)` rather than `node.text`, because `node.text` for large numeric literals can contain "."
	// e.g. `node.text` for numeric literal `1100000000000000000000` is `1.1e21`.
	isFractional := strings.ContainsRune(nodeText, '.')
	isScientific := node.TokenFlags&ast.TokenFlagsScientific != 0

	// Scientific notation (e.g. 2e54 and 1e00000000010) and fractional numbers
	// (e.g. 9000000000000000.001) are inherently imprecise, so skip the size check.
	if isFractional || isScientific {
		return
	}

	// Here `node` is guaranteed to be a numeric literal representing an integer.
	// We need to judge whether the integer `node` represents is <= 2 ** 53 - 1, which can be accomplished by comparing to `value` defined below because:
	// 1) when `node` represents an integer <= 2 ** 53 - 1, `node.text` is its exact string representation and thus `value` precisely represents the integer.
	// 2) otherwise, although `node.text` may be imprecise string representation, its mathematical value and consequently `value` cannot be less than 2 ** 53,
	//    thus the result of the predicate won't be affected.
	value := jsnum.FromString(node.Text)
	if value <= jsnum.MaxSafeInteger {
		return
	}

	c.addErrorOrSuggestion(false, createDiagnosticForNode(node.AsNode(), diagnostics.Numeric_literals_with_absolute_values_equal_to_2_53_or_greater_are_too_large_to_be_represented_accurately_as_integers))
}

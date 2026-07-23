package parser

import (
	"strconv"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/scanner"
)

func (p *Parser) finishReparsedNode(node *ast.Node, locationNode *ast.Node) {
	node.Flags = p.contextFlags | ast.NodeFlagsReparsed
	node.Loc = locationNode.Loc
	p.overrideParentInImmediateChildren(node)
}

func (p *Parser) finishMutatedNode(node *ast.Node) {
	p.overrideParentInImmediateChildren(node)
}

// Deep-clone the given node and add the clone to the reparsed clone list. The list is used by ast.GetReparsedNodeForNode
// to locate reparsed clones of JSDoc nodes. Since the binder attaches symbols to reparsed nodes and not to JSDoc nodes, we
// need the mapping when obtaining symbols and types from JSDoc nodes.
func (p *Parser) addDeepCloneReparse(node *ast.Node) *ast.Node {
	clone := p.factory.DeepCloneReparse(node)
	if clone != nil {
		p.reparsedClones = append(p.reparsedClones, clone)
	}
	return clone
}

func (p *Parser) addTransformedReparse(newNode *ast.Node, old *ast.Node) *ast.Node {
	p.finishReparsedNode(newNode, old)
	newNode.Flags |= ast.NodeFlagsReparserTransformedLiteral
	p.reparsedClones = append(p.reparsedClones, newNode)
	return newNode
}

func (p *Parser) checkNonIdentifierName(name *ast.Node) *ast.Node {
	if ast.IsIdentifier(name) && !scanner.IsValidIdentifier(name.AsIdentifier().Text) {
		errLoc := name.Loc
		if errLoc.Len() == 0 { // missing name, emit error on the character before the missing name node
			errLoc = core.NewTextRange(name.Loc.Pos()-1, name.Loc.Pos())
		}
		p.parseErrorAtRange(errLoc, diagnostics.Identifier_expected)
	}
	return name
}

// Hosted tags find a host and add their children to the correct location under the host.
// Unhosted tags add synthetic nodes to the reparse list.
func (p *Parser) reparseTags(parent *ast.Node, jsDoc []*ast.Node) {
	for _, j := range jsDoc {
		isLast := j == jsDoc[len(jsDoc)-1]
		tags := j.AsJSDoc().Tags
		if tags == nil {
			continue
		}
		for _, tag := range tags.Nodes {
			p.reparseUnhosted(tag, parent, j)
			if isLast {
				p.reparseHosted(tag, parent, j)
			}
		}
	}
}

func (p *Parser) reparseUnhosted(tag *ast.Node, parent *ast.Node, jsDoc *ast.Node) {
	switch tag.Kind {
	case ast.KindJSDocTypedefTag:
		typeExpression := tag.TypeExpression()
		if typeExpression == nil {
			break
		}
		fullName := tag.Name()
		isNamespace := fullName != nil && ast.IsModuleDeclaration(fullName)
		var modifiers *ast.ModifierList
		if isNamespace {
			modifiers = p.createExportModifier(tag)
		}
		typeAlias := p.factory.NewJSTypeAliasDeclaration(modifiers, p.addDeepCloneReparse(p.checkNonIdentifierName(p.getInnermostNameOfJSDocNamespace(fullName))), nil, nil)
		typeAlias.AsTypeAliasDeclaration().TypeParameters = p.gatherTypeParameters(jsDoc, true /*typedefOrCallback*/)
		var t *ast.Node
		switch typeExpression.Kind {
		case ast.KindJSDocTypeExpression:
			t = p.addDeepCloneReparse(typeExpression.Type())
		case ast.KindJSDocTypeLiteral:
			t = p.reparseJSDocTypeLiteral(typeExpression)
		default:
			panic("typedef tag type expression should be a name reference or a type expression" + typeExpression.Kind.String())
		}
		typeAlias.AsTypeAliasDeclaration().Type = t
		p.finishReparsedNode(typeAlias, tag)
		p.jsdocInfos = append(p.jsdocInfos, JSDocInfo{parent: typeAlias, jsDocs: []*ast.Node{jsDoc}})
		typeAlias.Flags |= ast.NodeFlagsHasJSDoc
		result := p.wrapInJSDocNamespace(fullName, typeAlias, false /*nested*/)
		p.reparseList = append(p.reparseList, result)
	case ast.KindJSDocCallbackTag:
		typeExpression := tag.TypeExpression()
		if typeExpression == nil {
			break
		}
		fullName := tag.Name()
		isNamespace := fullName != nil && ast.IsModuleDeclaration(fullName)
		var modifiers *ast.ModifierList
		if isNamespace {
			modifiers = p.createExportModifier(tag)
		}
		functionType := p.reparseJSDocSignature(typeExpression, tag, jsDoc, tag, nil)
		typeAlias := p.factory.NewJSTypeAliasDeclaration(modifiers, p.addDeepCloneReparse(p.getInnermostNameOfJSDocNamespace(fullName)), nil, functionType)
		typeAlias.AsTypeAliasDeclaration().TypeParameters = p.gatherTypeParameters(jsDoc, true /*typedefOrCallback*/)
		p.finishReparsedNode(typeAlias, tag)
		p.jsdocInfos = append(p.jsdocInfos, JSDocInfo{parent: typeAlias, jsDocs: []*ast.Node{jsDoc}})
		typeAlias.Flags |= ast.NodeFlagsHasJSDoc
		result := p.wrapInJSDocNamespace(fullName, typeAlias, false /*nested*/)
		p.reparseList = append(p.reparseList, result)
	case ast.KindJSDocOverloadTag:
		if ast.IsFunctionDeclaration(parent) && p.parsingContexts&(1<<PCObjectLiteralMembers) == 0 {
			p.reparseList = append(p.reparseList, p.reparseJSDocSignature(tag.AsJSDocOverloadTag().TypeExpression, parent, jsDoc, tag, parent.Modifiers()))
		}
	}
}

func (p *Parser) reparseJSDocSignature(jsSignature *ast.Node, fun *ast.Node, jsDoc *ast.Node, tag *ast.Node, modifiers *ast.ModifierList) *ast.Node {
	var signature *ast.Node
	clonedModifiers := p.factory.DeepCloneReparseModifiers(modifiers)
	switch fun.Kind {
	case ast.KindFunctionDeclaration:
		signature = p.factory.NewFunctionDeclaration(clonedModifiers, nil /*target*/, nil /*colonToken*/, p.factory.DeepCloneReparse(p.checkNonIdentifierName(fun.Name())), nil, nil, nil, nil, nil)
	case ast.KindJSDocCallbackTag:
		signature = p.factory.NewFunctionTypeNode(nil, nil, nil, p.factory.NewKeywordTypeNode(ast.KindAnyKeyword))
	default:
		panic("Unexpected kind " + fun.Kind.String())
	}

	if tag.Kind != ast.KindJSDocCallbackTag {
		signature.FunctionLikeData().TypeParameters = p.gatherTypeParameters(jsDoc, false /*typedefOrCallback*/)
	}
	parameters := p.nodeSliceArena.NewSlice(0)
	for pi, param := range jsSignature.Parameters() {
		var parameter *ast.Node
		if param.Kind == ast.KindJSDocThisTag {
			// `this` parameters are removed from tlua; a JSDoc `@this` entry no
			// longer synthesizes a receiver parameter.
			continue
		} else if param.Kind == ast.KindJSDocParameterTag || param.Kind == ast.KindJSDocPropertyTag {
			jsparam := param.AsJSDocParameterOrPropertyTag()
			// Skip sub-property parameters (e.g., @param x.y) - these have QualifiedNames
			// and describe properties of a parent parameter, not standalone parameters.
			if ast.IsQualifiedName(jsparam.Name()) {
				continue
			}
			var dotDotDotToken *ast.Node
			var paramType *ast.TypeNode

			if jsparam.TypeExpression != nil {
				if jsparam.TypeExpression.Type().Kind == ast.KindJSDocVariadicType {
					dotDotDotToken = p.factory.NewToken(ast.KindDotDotDotToken)
					dotDotDotToken.Loc = jsparam.Loc
					dotDotDotToken.Flags = p.contextFlags | ast.NodeFlagsReparsed

					variadicType := jsparam.TypeExpression.Type().AsJSDocVariadicType()
					paramType = p.reparseJSDocTypeLiteral(variadicType.Type)
				} else {
					paramType = p.reparseJSDocTypeLiteral(jsparam.TypeExpression.Type())
				}
			}
			name := jsparam.Name()
			switch {
			case dotDotDotToken != nil:
				// A tlua vararg has no name, and its annotation is the pack's
				// *element* type -- exactly what stripping the JSDocVariadicType
				// above leaves behind. Keeping the JSDoc name would mint a named
				// rest, a form no longer accepted, and declaration emit would then
				// write a .d.ts this compiler cannot read back.
				name = p.addTransformedReparse(p.factory.NewIdentifier(ast.VarargParameterName), name)
			case ast.IsIdentifier(name) && !scanner.IsValidIdentifier(name.AsIdentifier().Text):
				// drop invalid chars for _, if empty, write _0, etc., so we have a valid param name to emit later
				result := strings.Builder{}
				for i, ch := range name.AsIdentifier().Text {
					if i == 0 {
						if !scanner.IsIdentifierStart(ch) {
							result.WriteRune('_')
						} else {
							result.WriteRune(ch)
						}
						continue
					} else if !scanner.IsIdentifierPart(ch) {
						result.WriteRune('_')
					} else {
						result.WriteRune(ch)
					}
				}
				if result.Len() == 0 {
					result.WriteRune('_')
					result.WriteString(strconv.Itoa(pi))
				}
				name = p.addTransformedReparse(p.factory.NewIdentifier(result.String()), name)
			default:
				name = p.addDeepCloneReparse(name)
			}
			parameter = p.factory.NewParameterDeclaration(nil, dotDotDotToken, name, p.makeQuestionIfOptional(jsparam), paramType, nil)
		}
		p.finishReparsedNode(parameter, param)
		parameters = append(parameters, parameter)
		p.reparseJSDocComment(parameter, param)
	}
	signature.FunctionLikeData().Parameters = p.newNodeList(jsSignature.AsJSDocSignature().Parameters.Loc, parameters)

	if jsSignature.Type() != nil && jsSignature.Type().TypeExpression() != nil {
		signature.FunctionLikeData().Type = p.addDeepCloneReparse(jsSignature.Type().TypeExpression().Type())
	}
	loc := jsSignature
	if tag.Kind == ast.KindJSDocOverloadTag {
		loc = tag.TagName()
	}
	p.finishReparsedNode(signature, loc)
	return signature
}

func (p *Parser) reparseJSDocTypeLiteral(t *ast.TypeNode) *ast.Node {
	if t == nil {
		return nil
	}
	if t.Kind == ast.KindJSDocTypeLiteral {
		jstypeliteral := t.AsJSDocTypeLiteral()
		isArrayType := jstypeliteral.IsArrayType
		properties := p.nodeSliceArena.NewSlice(0)
		for _, prop := range jstypeliteral.JSDocPropertyTags {
			if prop.Kind != ast.KindJSDocPropertyTag && prop.Kind != ast.KindJSDocParameterTag {
				continue
			}
			jsprop := prop.AsJSDocParameterOrPropertyTag()
			name := prop.Name()
			if name.Kind == ast.KindQualifiedName {
				name = name.AsQualifiedName().Right
			}
			if ast.IsIdentifier(name) && !scanner.IsValidIdentifier(name.AsIdentifier().Text) {
				name = p.addTransformedReparse(p.factory.NewStringLiteral(name.AsIdentifier().Text, ast.TokenFlagsNone), name)
			} else {
				name = p.addDeepCloneReparse(name)
			}
			property := p.factory.NewPropertySignatureDeclaration(nil, name, p.makeQuestionIfOptional(jsprop), nil, nil)
			if jsprop.TypeExpression != nil {
				property.AsPropertySignatureDeclaration().Type = p.reparseJSDocTypeLiteral(jsprop.TypeExpression.Type())
			}
			p.finishReparsedNode(property, prop)
			properties = append(properties, property)
			p.reparseJSDocComment(property, prop)
		}
		t = p.factory.NewTypeLiteralNode(p.newNodeList(jstypeliteral.Loc, properties))
		if isArrayType {
			p.finishReparsedNode(t, jstypeliteral.AsNode())
			t = p.factory.NewArrayTypeNode(t)
		}
		p.finishReparsedNode(t, jstypeliteral.AsNode())
		return t
	}
	return p.addDeepCloneReparse(t)
}

func (p *Parser) reparseJSDocComment(node *ast.Node, tag *ast.Node) {
	if comment := tag.CommentList(); comment != nil {
		newComment := p.factory.NewNodeList(core.Map(comment.Nodes, p.factory.DeepCloneReparse))
		newComment.Loc = comment.Loc
		propJSDoc := p.factory.NewJSDoc(newComment, nil)
		p.finishReparsedNode(propJSDoc, tag)
		propJSDoc.Parent = node
		p.jsdocInfos = append(p.jsdocInfos, JSDocInfo{parent: node, jsDocs: []*ast.Node{propJSDoc}})
		node.Flags |= ast.NodeFlagsHasJSDoc
	}
}

func (p *Parser) gatherTypeParameters(j *ast.Node, typedefOrCallback bool) *ast.NodeList {
	var typeParameters []*ast.Node
	pos := -1
	endPos := -1
	firstTemplate := true
	for _, tag := range j.AsJSDoc().Tags.Nodes {
		// When a JSDoc comment contains an `@typedef` or `@callback` tag, `@template` type parameter
		// declarations apply to the type being defined.
		if !typedefOrCallback && (ast.IsJSDocTypedefTag(tag) || ast.IsJSDocCallbackTag(tag)) {
			return nil
		}
		if !ast.IsJSDocTemplateTag(tag) {
			continue
		}
		if firstTemplate {
			pos = tag.Pos()
			firstTemplate = false
		}
		endPos = tag.End()
		constraint := tag.AsJSDocTemplateTag().Constraint
		firstTypeParameter := true
		for _, tp := range tag.TypeParameters() {
			var reparse *ast.Node
			if constraint != nil && firstTypeParameter {
				reparse = p.factory.NewTypeParameterDeclaration(
					p.factory.DeepCloneReparseModifiers(tp.Modifiers()),
					nil, // dotDotDotToken (JSDoc type parameters are never packs)
					p.addDeepCloneReparse(p.checkNonIdentifierName(tp.Name())),
					p.addDeepCloneReparse(constraint.Type()),
					nil, // expression
					p.addDeepCloneReparse(tp.AsTypeParameterDeclaration().DefaultType),
				)
				p.finishReparsedNode(reparse, tp)
			} else {
				reparse = p.addDeepCloneReparse(tp)
			}
			if typeParameters == nil {
				typeParameters = p.nodeSliceArena.NewSlice(0)
			}
			typeParameters = append(typeParameters, reparse)
			firstTypeParameter = false
		}
	}
	if len(typeParameters) == 0 {
		return nil
	} else {
		return p.newNodeList(core.NewTextRange(pos, endPos), typeParameters)
	}
}

func (p *Parser) reparseHosted(tag *ast.Node, parent *ast.Node, jsDoc *ast.Node) {
	switch tag.Kind {
	case ast.KindJSDocTypeTag:
		switch parent.Kind {
		case ast.KindVariableStatement:
			if parent.AsVariableStatement().DeclarationList != nil {
				for _, declaration := range parent.AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes {
					if declaration.Type() == nil && tag.TypeExpression() != nil {
						declaration.AsMutable().SetType(p.addDeepCloneReparse(tag.TypeExpression().Type()))
						p.finishMutatedNode(declaration)
						return
					}
				}
			}
		case ast.KindVariableDeclaration, ast.KindExportAssignment, ast.KindPropertyAssignment,
			ast.KindShorthandPropertyAssignment:
			if parent.Type() == nil && tag.TypeExpression() != nil {
				parent.AsMutable().SetType(p.addDeepCloneReparse(tag.TypeExpression().Type()))
				p.finishMutatedNode(parent)
				return
			}
		case ast.KindParameter:
			if parent.Type() == nil && tag.TypeExpression() != nil {
				parent.AsMutable().SetType(p.reparseJSDocTypeLiteral(tag.TypeExpression().Type()))
				p.finishMutatedNode(parent)
				return
			}
		case ast.KindExpressionStatement:
			if parent.Expression().Kind == ast.KindBinaryExpression {
				bin := parent.Expression().AsBinaryExpression()
				if kind := ast.GetAssignmentDeclarationKind(bin.AsNode()); kind != ast.JSDeclarationKindNone && tag.TypeExpression() != nil {
					bin.AsMutable().SetType(p.addDeepCloneReparse(tag.TypeExpression().Type()))
					p.finishMutatedNode(bin.AsNode())
					return
				}
			}
		case ast.KindReturnStatement, ast.KindParenthesizedExpression:
			if parent.Expression() != nil && tag.TypeExpression() != nil {
				parent.AsMutable().SetExpression(p.makeNewCast(
					p.addDeepCloneReparse(tag.TypeExpression().Type()),
					parent.Expression(),
					true, /*isAssertion*/
				))
				p.finishMutatedNode(parent)
				return
			}
		}
		if fun := getFunctionLikeHost(parent); fun != nil {
			noTypedParams := core.Every(fun.Parameters(), func(param *ast.Node) bool { return param.Type() == nil })
			if fun.TypeParameterList() == nil && fun.Type() == nil && noTypedParams && tag.TypeExpression() != nil {
				fun.FunctionLikeData().FullSignature = p.addDeepCloneReparse(tag.TypeExpression().Type())
				p.finishMutatedNode(fun)
			}
		}
	case ast.KindJSDocSatisfiesTag:
		switch parent.Kind {
		case ast.KindVariableStatement:
			if parent.AsVariableStatement().DeclarationList != nil {
				for _, declaration := range parent.AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes {
					if declaration.Initializer() != nil && tag.TypeExpression() != nil {
						declaration.AsMutable().SetInitializer(p.makeNewCast(
							p.addDeepCloneReparse(tag.TypeExpression().Type()),
							declaration.Initializer(),
							false, /*isAssertion*/
						))
						p.finishMutatedNode(declaration)
						break
					}
				}
			}
		case ast.KindVariableDeclaration, ast.KindPropertyAssignment:
			if parent.Initializer() != nil && tag.TypeExpression() != nil {
				parent.AsMutable().SetInitializer(p.makeNewCast(
					p.addDeepCloneReparse(tag.TypeExpression().Type()),
					parent.Initializer(),
					false, /*isAssertion*/
				))
				p.finishMutatedNode(parent)
			}
		case ast.KindShorthandPropertyAssignment:
			shorthand := parent.AsShorthandPropertyAssignment()
			if shorthand.ObjectAssignmentInitializer != nil && tag.AsJSDocSatisfiesTag().TypeExpression != nil {
				shorthand.ObjectAssignmentInitializer = p.makeNewCast(
					p.addDeepCloneReparse(tag.AsJSDocSatisfiesTag().TypeExpression.Type()),
					shorthand.ObjectAssignmentInitializer,
					false, /*isAssertion*/
				)
				p.finishMutatedNode(parent)
			}
		case ast.KindReturnStatement, ast.KindParenthesizedExpression, ast.KindExportAssignment:
			if parent.Expression() != nil && tag.TypeExpression() != nil {
				parent.AsMutable().SetExpression(p.makeNewCast(
					p.addDeepCloneReparse(tag.TypeExpression().Type()),
					parent.Expression(),
					false, /*isAssertion*/
				))
				p.finishMutatedNode(parent)
			}
		case ast.KindExpressionStatement:
			if parent.Expression().Kind == ast.KindBinaryExpression {
				bin := parent.Expression().AsBinaryExpression()
				if kind := ast.GetAssignmentDeclarationKind(bin.AsNode()); kind != ast.JSDeclarationKindNone && tag.TypeExpression() != nil {
					bin.Right = p.makeNewCast(
						p.addDeepCloneReparse(tag.TypeExpression().Type()),
						bin.Right,
						false, /*isAssertion*/
					)
					p.finishMutatedNode(bin.AsNode())
				}
			}
		}
	case ast.KindJSDocTemplateTag:
		if fun := getFunctionLikeHost(parent); fun != nil {
			if fun.TypeParameters() == nil && fun.FunctionLikeData().FullSignature == nil {
				fun.FunctionLikeData().TypeParameters = p.gatherTypeParameters(jsDoc, false /*typedefOrCallback*/)
				p.finishMutatedNode(fun)
			}
		}
	case ast.KindJSDocParameterTag:
		if fun := getFunctionLikeHost(parent); fun != nil && fun.FunctionLikeData().FullSignature == nil {
			parameterTag := tag.AsJSDocParameterOrPropertyTag()
			if param, ok := findMatchingParameter(fun, parameterTag, jsDoc); ok {
				if param.Type == nil && parameterTag.TypeExpression != nil {
					param.AsParameterDeclaration().Type = p.reparseJSDocTypeLiteral(parameterTag.TypeExpression.Type())
				}
				if param.QuestionToken == nil {
					if question := p.makeQuestionIfOptional(parameterTag); question != nil {
						param.QuestionToken = question
					}
				}
				p.finishMutatedNode(param.AsNode())
			}
		}
	// KindJSDocThisTag: `this` parameters are removed from tlua, so a `@this`
	// tag no longer synthesizes a receiver parameter; the tag stays a plain
	// JSDoc comment node.
	case ast.KindJSDocReturnTag:
		if fun := getFunctionLikeHost(parent); fun != nil && fun.FunctionLikeData().FullSignature == nil {
			if fun.Type() == nil && tag.TypeExpression() != nil {
				fun.FunctionLikeData().Type = p.addDeepCloneReparse(tag.TypeExpression().Type())
				p.finishMutatedNode(fun)
			}
		}
	case ast.KindJSDocReadonlyTag, ast.KindJSDocPrivateTag, ast.KindJSDocPublicTag, ast.KindJSDocProtectedTag, ast.KindJSDocOverrideTag:
		if parent.Kind == ast.KindExpressionStatement {
			parent = parent.Expression()
		}
		switch parent.Kind {
		case ast.KindBinaryExpression:
			var keyword ast.Kind
			switch tag.Kind {
			case ast.KindJSDocReadonlyTag:
				keyword = ast.KindReadonlyKeyword
			case ast.KindJSDocPrivateTag:
				keyword = ast.KindPrivateKeyword
			case ast.KindJSDocPublicTag:
				keyword = ast.KindPublicKeyword
			case ast.KindJSDocProtectedTag:
				keyword = ast.KindProtectedKeyword
			case ast.KindJSDocOverrideTag:
				keyword = ast.KindOverrideKeyword
			}
			modifier := p.factory.NewModifier(keyword)
			modifier.Loc = tag.Loc
			modifier.Flags = p.contextFlags | ast.NodeFlagsReparsed
			var nodes []*ast.Node
			var loc core.TextRange
			if parent.Modifiers() == nil {
				nodes = p.nodeSliceArena.NewSlice(1)
				nodes[0] = modifier
				loc = tag.Loc
			} else {
				nodes = append(parent.ModifierNodes(), modifier)
				loc = parent.Modifiers().Loc
			}
			parent.AsMutable().SetModifiers(p.newModifierList(loc, nodes))
			p.finishMutatedNode(parent)
		}
	}
}

func (p *Parser) makeQuestionIfOptional(parameter *ast.JSDocParameterOrPropertyTag) *ast.Node {
	var questionToken *ast.Node
	if parameter.IsBracketed || parameter.TypeExpression != nil && parameter.TypeExpression.Type().Kind == ast.KindJSDocOptionalType {
		questionToken = p.factory.NewToken(ast.KindQuestionToken)
		questionToken.Loc = parameter.Loc
		questionToken.Flags = p.contextFlags | ast.NodeFlagsReparsed
	}
	return questionToken
}

func findMatchingParameter(fun *ast.Node, parameterTag *ast.JSDocParameterOrPropertyTag, jsDoc *ast.Node) (*ast.ParameterDeclaration, bool) {
	tagIndex := -1
	paramCount := -1
	for _, tag := range jsDoc.AsJSDoc().Tags.Nodes {
		if tag.Kind == ast.KindJSDocParameterTag {
			paramCount++
			if tag.AsJSDocParameterOrPropertyTag() == parameterTag {
				tagIndex = paramCount
				break
			}
		}
	}
	for parameterIndex, parameter := range fun.Parameters() {
		if parameter.Name().Kind == ast.KindIdentifier {
			if parameterTag.Name().Kind == ast.KindIdentifier &&
				((parameter.Name().Text() == parameterTag.Name().Text()) || (parameterIndex == tagIndex && len(parameterTag.Name().Text()) == 0)) {
				return parameter.AsParameterDeclaration(), true
			}
		} else if parameterIndex == tagIndex {
			return parameter.AsParameterDeclaration(), true
		}
	}
	return nil, false
}

func skipSatisfiesExpressions(node *ast.Node) *ast.Node {
	for node != nil && node.Kind == ast.KindSatisfiesExpression {
		node = node.Expression()
	}
	return node
}

func getFunctionLikeHost(host *ast.Node) *ast.Node {
	fun := host
	switch host.Kind {
	case ast.KindVariableStatement:
		if nodes := host.AsVariableStatement().DeclarationList.AsVariableDeclarationList().Declarations.Nodes; len(nodes) != 0 {
			fun = nodes[0].Initializer()
		}
	case ast.KindPropertyAssignment:
		fun = host.Initializer()
	case ast.KindExportAssignment, ast.KindReturnStatement:
		fun = host.Expression()
	case ast.KindExpressionStatement:
		fun = ast.GetRightMostAssignedExpression(host.Expression())
	}
	fun = skipSatisfiesExpressions(fun)
	if ast.IsFunctionLike(fun) {
		return fun
	}
	return nil
}

func (p *Parser) makeNewCast(t *ast.TypeNode, e *ast.Node, isAssertion bool) *ast.Node {
	var assert *ast.Node
	if isAssertion {
		assert = p.factory.NewAsExpression(e, t)
	} else {
		assert = p.factory.NewSatisfiesExpression(e, t)
	}
	p.finishNodeWithEnd(assert, e.Pos(), e.End())
	return assert
}

func (p *Parser) createExportModifier(locationNode *ast.Node) *ast.ModifierList {
	exportModifier := p.factory.NewModifier(ast.KindExportKeyword)
	exportModifier.Loc = locationNode.Loc
	exportModifier.Flags = p.contextFlags | ast.NodeFlagsReparsed
	nodes := p.nodeSliceArena.NewSlice1(exportModifier)
	return p.newModifierList(locationNode.Loc, nodes)
}

// getInnermostNameOfJSDocNamespace returns the innermost identifier from a
// JSDoc namespace chain (ModuleDeclaration). For a simple identifier, it returns
// the identifier itself. For "A.B.C", it returns the identifier "C".
func (p *Parser) getInnermostNameOfJSDocNamespace(fullName *ast.Node) *ast.Node {
	if fullName == nil {
		return nil
	}
	for fullName.Kind == ast.KindModuleDeclaration {
		body := fullName.AsModuleDeclaration().Body
		if body == nil {
			return fullName.Name()
		}
		fullName = body
	}
	return fullName
}

// wrapInJSDocNamespace wraps a statement (typically a type alias) in namespace
// declarations corresponding to a JSDoc dotted name. For example, given name
// "A.B.C" and a type alias for C, this produces:
//
//	namespace A { namespace B { type C = ... } }
//
// If the name is a simple identifier (not a ModuleDeclaration), it returns the
// statement as-is.
func (p *Parser) wrapInJSDocNamespace(fullName *ast.Node, statement *ast.Node, nested bool) *ast.Node {
	if fullName == nil || !ast.IsModuleDeclaration(fullName) {
		return statement
	}
	// Recursively wrap from outermost to innermost. Inner namespaces always get an export modifier
	// so members are accessible via dotted access from outside. The outermost namespace is treated as
	// exported only in module files via IsImplicitlyExportedJSDocDeclaration (in the binder), so it
	// does not get an explicit export modifier here.
	wrapped := p.wrapInJSDocNamespace(fullName.Body(), statement, true /*nested*/)
	block := p.factory.NewModuleBlock(p.newNodeList(fullName.Loc, p.nodeSliceArena.NewSlice1(wrapped)))
	p.finishReparsedNode(block, fullName)
	var modifiers *ast.ModifierList
	if nested {
		modifiers = p.createExportModifier(fullName)
	}
	result := p.factory.NewModuleDeclaration(modifiers, ast.KindNamespaceKeyword, p.addDeepCloneReparse(fullName.Name()), block)
	p.finishReparsedNode(result, fullName)
	p.reparsedClones = append(p.reparsedClones, result)
	return result
}

package tstransforms

// !!! SourceMaps and Comments need to be validated

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/binder"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/transformers"
)

// Transforms TypeScript-specific runtime syntax into JavaScript-compatible syntax.
type RuntimeSyntaxTransformer struct {
	transformers.Transformer
	compilerOptions                     *core.CompilerOptions
	parentNode                          *ast.Node
	currentNode                         *ast.Node
	currentSourceFile                   *ast.Node
	currentScope                        *ast.Node // SourceFile | Block | ModuleBlock | CaseBlock
	currentScopeFirstDeclarationsOfName map[string]*ast.Node
	currentNamespace                    *ast.ModuleDeclarationNode
	resolver                            binder.ReferenceResolver
}

func NewRuntimeSyntaxTransformer(opt *transformers.TransformOptions) *transformers.Transformer {
	compilerOptions := opt.CompilerOptions
	emitContext := opt.Context
	tx := &RuntimeSyntaxTransformer{compilerOptions: compilerOptions, resolver: opt.Resolver}
	return tx.NewTransformer(tx.visit, emitContext)
}

// Pushes a new child node onto the ancestor tracking stack, returning the grandparent node to be restored later via `popNode`.
func (tx *RuntimeSyntaxTransformer) pushNode(node *ast.Node) (grandparentNode *ast.Node) {
	grandparentNode = tx.parentNode
	tx.parentNode = tx.currentNode
	tx.currentNode = node
	return grandparentNode
}

// Pops the last child node off the ancestor tracking stack, restoring the grandparent node.
func (tx *RuntimeSyntaxTransformer) popNode(grandparentNode *ast.Node) {
	tx.currentNode = tx.parentNode
	tx.parentNode = grandparentNode
}

func (tx *RuntimeSyntaxTransformer) pushScope(node *ast.Node) (savedCurrentScope *ast.Node, savedCurrentScopeFirstDeclarationsOfName map[string]*ast.Node) {
	savedCurrentScope = tx.currentScope
	savedCurrentScopeFirstDeclarationsOfName = tx.currentScopeFirstDeclarationsOfName
	switch node.Kind {
	case ast.KindSourceFile:
		tx.currentScope = node
		tx.currentSourceFile = node
		tx.currentScopeFirstDeclarationsOfName = nil
	case ast.KindModuleBlock, ast.KindBlock:
		tx.currentScope = node
		tx.currentScopeFirstDeclarationsOfName = nil
	case ast.KindFunctionDeclaration, ast.KindVariableStatement:
		tx.recordDeclarationInScope(node)
	}
	return savedCurrentScope, savedCurrentScopeFirstDeclarationsOfName
}

func (tx *RuntimeSyntaxTransformer) popScope(savedCurrentScope *ast.Node, savedCurrentScopeFirstDeclarationsOfName map[string]*ast.Node) {
	if tx.currentScope != savedCurrentScope {
		// only reset the first declaration for a name if we are exiting the scope in which it was declared
		tx.currentScopeFirstDeclarationsOfName = savedCurrentScopeFirstDeclarationsOfName
	}

	tx.currentScope = savedCurrentScope
}

// Visits each node in the AST
func (tx *RuntimeSyntaxTransformer) visit(node *ast.Node) *ast.Node {
	grandparentNode := tx.pushNode(node)
	defer tx.popNode(grandparentNode)

	savedCurrentScope, savedCurrentScopeFirstDeclarationsOfName := tx.pushScope(node)
	defer tx.popScope(savedCurrentScope, savedCurrentScopeFirstDeclarationsOfName)

	if node.SubtreeFacts()&ast.SubtreeContainsTypeScript == 0 && (tx.currentNamespace == nil || node.SubtreeFacts()&ast.SubtreeContainsIdentifier == 0) {
		return node
	}

	switch node.Kind {
	// TypeScript parameter property modifiers are elided
	case ast.KindPublicKeyword,
		ast.KindPrivateKeyword,
		ast.KindProtectedKeyword,
		ast.KindReadonlyKeyword,
		ast.KindOverrideKeyword:
		node = nil
	case ast.KindModuleDeclaration:
		node = tx.visitModuleDeclaration(node.AsModuleDeclaration())
	case ast.KindFunctionDeclaration:
		node = tx.visitFunctionDeclaration(node.AsFunctionDeclaration())
	case ast.KindVariableStatement:
		node = tx.visitVariableStatement(node.AsVariableStatement())
	case ast.KindExportDeclaration, ast.KindImportDeclaration, ast.KindImportClause:
		if tx.currentNamespace != nil && tx.currentScope != nil && tx.currentScope.Kind != ast.KindBlock {
			// do not emit ES6 imports and exports since they are illegal inside a namespace
			node = nil
		} else {
			node = tx.Visitor().VisitEachChild(node)
		}
	case ast.KindImportEqualsDeclaration:
		if tx.currentNamespace != nil && tx.currentScope != nil && tx.currentScope.Kind != ast.KindBlock && node.AsImportEqualsDeclaration().ModuleReference.Kind == ast.KindExternalModuleReference {
			// do not emit ES6 imports and exports since they are illegal inside a namespace
			node = nil
		} else if tx.currentNamespace != nil && tx.currentScope != nil && tx.currentScope.Kind == ast.KindBlock && node.AsImportEqualsDeclaration().ModuleReference.Kind != ast.KindExternalModuleReference {
			// inside a block within a namespace, elide internal import aliases
			node = nil
		} else {
			node = tx.visitImportEqualsDeclaration(node.AsImportEqualsDeclaration())
		}
	case ast.KindIdentifier:
		node = tx.visitIdentifier(node)
	case ast.KindShorthandPropertyAssignment:
		node = tx.visitShorthandPropertyAssignment(node.AsShorthandPropertyAssignment())
	default:
		node = tx.Visitor().VisitEachChild(node)
	}
	return node
}

// Records that a declaration was emitted in the current scope, if it was the first declaration for the provided symbol.
func (tx *RuntimeSyntaxTransformer) recordDeclarationInScope(node *ast.Node) {
	switch node.Kind {
	case ast.KindVariableStatement:
		tx.recordDeclarationInScope(node.AsVariableStatement().DeclarationList)
		return
	case ast.KindVariableDeclarationList:
		for _, decl := range node.AsVariableDeclarationList().Declarations.Nodes {
			tx.recordDeclarationInScope(decl)
		}
		return
	case ast.KindArrayBindingPattern, ast.KindObjectBindingPattern:
		for _, element := range node.Elements() {
			tx.recordDeclarationInScope(element)
		}
		return
	}
	name := node.Name()
	if name != nil {
		if ast.IsIdentifier(name) {
			if tx.currentScopeFirstDeclarationsOfName == nil {
				tx.currentScopeFirstDeclarationsOfName = make(map[string]*ast.Node)
			}
			text := name.Text()
			if _, found := tx.currentScopeFirstDeclarationsOfName[text]; !found {
				tx.currentScopeFirstDeclarationsOfName[text] = node
			}
		} else if ast.IsBindingPattern(name) {
			tx.recordDeclarationInScope(name)
		}
	}
}

// Determines whether a declaration is the first declaration with the same name emitted in the current scope.
func (tx *RuntimeSyntaxTransformer) isFirstDeclarationInScope(node *ast.Node) bool {
	name := node.Name()
	if name != nil && ast.IsIdentifier(name) {
		text := name.Text()
		if firstDeclaration, found := tx.currentScopeFirstDeclarationsOfName[text]; found {
			return firstDeclaration == node
		}
	}
	return false
}

func (tx *RuntimeSyntaxTransformer) isExportOfNamespace(node *ast.Node) bool {
	// Namespaces are deleted from the grammar; nothing is ever transformed
	// inside a namespace body, so no declaration is a namespace export.
	return false
}

// Gets an expression used to refer to a namespace from within its declaration.
func (tx *RuntimeSyntaxTransformer) getNamespaceContainerName(node *ast.Node) *ast.IdentifierNode {
	return tx.Factory().NewGeneratedNameForNode(node)
}

// Gets an expression used to refer to an export of a namespace by property name.
func (tx *RuntimeSyntaxTransformer) getNamespaceQualifiedProperty(ns *ast.IdentifierNode, name *ast.IdentifierNode) *ast.Expression {
	return tx.Factory().GetNamespaceMemberName(ns, name, printer.NameOptions{AllowSourceMaps: true})
}

// Gets an expression used within the provided node's container for any exported references.
func (tx *RuntimeSyntaxTransformer) getExportQualifiedReferenceToDeclaration(node *ast.Declaration) *ast.Expression {
	if tx.isExportOfNamespace(node.AsNode()) {
		return tx.Factory().GetExternalModuleOrNamespaceExportName(tx.getNamespaceContainerName(tx.currentNamespace), node, false /*allowComments*/, true /*allowSourceMaps*/)
	}
	return tx.Factory().GetDeclarationNameEx(node.AsNode(), printer.NameOptions{AllowSourceMaps: true})
}

func (tx *RuntimeSyntaxTransformer) addVarForDeclaration(statements []*ast.Statement, node *ast.Declaration) ([]*ast.Statement, bool) {
	tx.recordDeclarationInScope(node)
	if !tx.isFirstDeclarationInScope(node) {
		return statements, false
	}

	// var name;
	name := tx.Factory().GetLocalNameEx(node, printer.AssignedNameOptions{AllowSourceMaps: true})
	varDecl := tx.Factory().NewVariableDeclaration(name, nil, nil, nil)
	varFlags := core.IfElse(tx.currentScope == tx.currentSourceFile, ast.NodeFlagsNone, ast.NodeFlagsLet)
	varDecls := tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), varFlags)
	// Replicate modifierVisitor: strip TypeScript modifiers, and export when in namespace.
	modifierMask := ^ast.ModifierFlagsTypeScriptModifier
	if tx.currentNamespace != nil {
		modifierMask &^= ast.ModifierFlagsExport
	}
	modifiers := transformers.ExtractModifiers(tx.EmitContext(), node.Modifiers(), modifierMask)
	varStatement := tx.Factory().NewVariableStatement(modifiers, varDecls)

	tx.EmitContext().SetOriginal(varDecl, node)
	// !!! synthetic comments
	tx.EmitContext().SetOriginal(varStatement, node)

	tx.EmitContext().SetSourceMapRange(varStatement, node.Loc)

	// Emit a namespace's trailing comments after its closure, not after this declaration.
	tx.EmitContext().SetCommentRange(varStatement, node.Loc)
	tx.EmitContext().AddEmitFlags(varStatement, printer.EFNoTrailingComments)
	statements = append(statements, varStatement)

	return statements, true
}

// visitModuleDeclaration elides every module declaration: ambient modules and
// `declare global` have no runtime form, and identifier namespaces are deleted
// from the grammar (a recovered parse already carries the error), so the old
// namespace-IIFE emission is gone.
func (tx *RuntimeSyntaxTransformer) visitModuleDeclaration(node *ast.ModuleDeclaration) *ast.Node {
	return tx.EmitContext().NewNotEmittedStatement(node.AsNode())
}

func (tx *RuntimeSyntaxTransformer) visitImportEqualsDeclaration(node *ast.ImportEqualsDeclaration) *ast.Node {
	if node.ModuleReference.Kind == ast.KindExternalModuleReference {
		return tx.Visitor().VisitEachChild(node.AsNode())
	}

	moduleReference := tx.Factory().CreateExpressionFromEntityName(node.ModuleReference)
	tx.EmitContext().SetEmitFlags(moduleReference, printer.EFNoComments|printer.EFNoNestedComments)
	if !tx.isExportOfNamespace(node.AsNode()) {
		//  export var ${name} = ${moduleReference};
		//  var ${name} = ${moduleReference};
		varDecl := tx.Factory().NewVariableDeclaration(node.Name(), nil /*exclamationToken*/, nil /*type*/, moduleReference)
		tx.EmitContext().SetOriginal(varDecl, node.AsNode())
		varList := tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), ast.NodeFlagsNone)
		varModifiers := transformers.ExtractModifiers(tx.EmitContext(), node.Modifiers(), ast.ModifierFlagsExport)
		varStatement := tx.Factory().NewVariableStatement(varModifiers, varList)
		tx.EmitContext().SetOriginal(varStatement, node.AsNode())
		tx.EmitContext().AssignCommentAndSourceMapRanges(varStatement, node.AsNode())
		return varStatement
	} else {
		// exports.${name} = ${moduleReference};
		statement := tx.createExportStatement(node.Name(), moduleReference, node.Loc, node.Loc, node.AsNode())
		statement.Loc = node.Loc
		return statement
	}
}

func (tx *RuntimeSyntaxTransformer) visitVariableStatement(node *ast.VariableStatement) *ast.Node {
	if tx.isExportOfNamespace(node.AsNode()) {
		expressions := []*ast.Expression{}
		for _, declaration := range node.DeclarationList.AsVariableDeclarationList().Declarations.Nodes {
			v := declaration.AsVariableDeclaration()
			if v.Initializer == nil {
				continue
			}
			// Binding-pattern declaration names no longer parse, so the
			// declaration-to-assignment conversion is the only form left.
			expression := transformers.ConvertVariableDeclarationToAssignmentExpression(tx.EmitContext(), v)
			if expression != nil {
				expressions = append(expressions, expression)
			}
		}
		if len(expressions) == 0 {
			return nil
		}
		expression := tx.Factory().InlineExpressions(expressions)
		statement := tx.Factory().NewExpressionStatement(expression)
		tx.EmitContext().SetOriginal(statement, node.AsNode())
		tx.EmitContext().AssignCommentAndSourceMapRanges(statement, node.AsNode())

		// re-visit as the new node
		savedCurrent := tx.currentNode
		tx.currentNode = statement
		statement = tx.Visitor().VisitEachChild(statement)
		tx.currentNode = savedCurrent
		return statement
	}
	return tx.Visitor().VisitEachChild(node.AsNode())
}

// createNamespaceExportExpression creates an assignment to a namespace member for use as a
// callback during destructuring flattening.
func (tx *RuntimeSyntaxTransformer) createNamespaceExportExpression(exportName *ast.IdentifierNode, exportValue *ast.Expression, location *core.TextRange) *ast.Expression {
	memberName := tx.getNamespaceQualifiedProperty(tx.getNamespaceContainerName(tx.currentNamespace), exportName)
	expression := tx.Factory().NewAssignmentExpression(memberName, exportValue)
	if location != nil {
		expression.Loc = *location
	}
	return expression
}

func (tx *RuntimeSyntaxTransformer) visitFunctionDeclaration(node *ast.FunctionDeclaration) *ast.Node {
	if tx.isExportOfNamespace(node.AsNode()) {
		updated := tx.Factory().UpdateFunctionDeclaration(
			node,
			tx.Visitor().VisitModifiers(transformers.ExtractModifiers(tx.EmitContext(), node.Modifiers(), ^ast.ModifierFlagsExport)),
			tx.Visitor().VisitNode(node.Target),
			node.ColonToken,
			tx.Visitor().VisitNode(node.Name()),
			nil, /*typeParameters*/
			tx.Visitor().VisitNodes(node.Parameters),
			nil, /*returnType*/
			nil, /*fullSignature*/
			tx.Visitor().VisitNode(node.Body),
		)
		export := tx.createExportStatementForDeclaration(node.AsNode())
		if export != nil {
			return tx.Factory().NewSyntaxList([]*ast.Node{updated, export})
		}
		return updated
	}
	return tx.Visitor().VisitEachChild(node.AsNode())
}

func (tx *RuntimeSyntaxTransformer) transformConstructorBodyWorker(statementsIn []*ast.Statement, superPath []int, initializerStatements []*ast.Statement) []*ast.Statement {
	var statementsOut []*ast.Statement
	superStatementIndex := superPath[0]
	superStatement := statementsIn[superStatementIndex]

	// visit up to the statement containing `super`
	statementsOut = append(statementsOut, core.FirstResult(tx.Visitor().VisitSlice(statementsIn[:superStatementIndex]))...)

	// if the statement containing `super` is a `try` statement, transform the body of the `try` block
	if ast.IsTryStatement(superStatement) {
		tryStatement := superStatement.AsTryStatement()
		tryBlock := tryStatement.TryBlock.AsBlock()

		// keep track of hierarchy as we descend
		grandparentOfTryStatement := tx.pushNode(tryStatement.AsNode())
		grandparentOfTryBlock := tx.pushNode(tryBlock.AsNode())
		savedCurrentScope, savedCurrentScopeFirstDeclarationsOfName := tx.pushScope(tryBlock.AsNode())

		// visit the `try` block
		tryBlockStatements := tx.transformConstructorBodyWorker(
			tryBlock.Statements.Nodes,
			superPath[1:],
			initializerStatements,
		)

		// restore hierarchy as we ascend to the `try` statement
		tx.popScope(savedCurrentScope, savedCurrentScopeFirstDeclarationsOfName)
		tx.popNode(grandparentOfTryBlock)

		tryBlockStatementList := tx.Factory().NewNodeList(tryBlockStatements)
		tryBlockStatementList.Loc = tryBlock.Statements.Loc
		statementsOut = append(statementsOut, tx.Factory().UpdateTryStatement(
			tryStatement,
			tx.Factory().UpdateBlock(tryBlock, tryBlockStatementList, tryBlock.MultiLine),
			tx.Visitor().VisitNode(tryStatement.CatchClause),
			tx.Visitor().VisitNode(tryStatement.FinallyBlock),
		))

		// restore hierarchy as we ascend to the parent of the `try` statement
		tx.popNode(grandparentOfTryStatement)
	} else {
		// visit the statement containing `super`
		statementsOut = append(statementsOut, core.FirstResult(tx.Visitor().VisitSlice(statementsIn[superStatementIndex:superStatementIndex+1]))...)

		// insert the initializer statements
		statementsOut = append(statementsOut, initializerStatements...)
	}

	// visit the statements after `super`
	statementsOut = append(statementsOut, core.FirstResult(tx.Visitor().VisitSlice(statementsIn[superStatementIndex+1:]))...)
	return statementsOut
}

func (tx *RuntimeSyntaxTransformer) visitShorthandPropertyAssignment(node *ast.ShorthandPropertyAssignment) *ast.Node {
	name := node.Name()
	exportedOrImportedName := tx.visitExpressionIdentifier(name)
	if exportedOrImportedName != name {
		expression := exportedOrImportedName
		if node.ObjectAssignmentInitializer != nil {
			equalsToken := node.EqualsToken
			if equalsToken == nil {
				equalsToken = tx.Factory().NewToken(ast.KindEqualsToken)
			}
			expression = tx.Factory().NewBinaryExpression(
				nil, /*modifiers*/
				expression,
				nil, /*typeNode*/
				equalsToken,
				tx.Visitor().VisitNode(node.ObjectAssignmentInitializer),
			)
		}

		updated := tx.Factory().NewPropertyAssignment(nil /*modifiers*/, node.Name(), nil /*postfixToken*/, nil /*typeNode*/, expression)
		updated.Loc = node.Loc
		tx.EmitContext().SetOriginal(updated, node.AsNode())
		tx.EmitContext().AssignCommentAndSourceMapRanges(updated, node.AsNode())
		return updated
	}
	return tx.Factory().UpdateShorthandPropertyAssignment(
		node,
		nil, /*modifiers*/
		exportedOrImportedName,
		nil, /*postfixToken*/
		nil, /*typeNode*/
		node.EqualsToken,
		tx.Visitor().VisitNode(node.ObjectAssignmentInitializer),
	)
}

func (tx *RuntimeSyntaxTransformer) visitIdentifier(node *ast.IdentifierNode) *ast.Node {
	if transformers.IsIdentifierReference(node, tx.parentNode) {
		return tx.visitExpressionIdentifier(node)
	}
	return node
}

func (tx *RuntimeSyntaxTransformer) visitExpressionIdentifier(node *ast.IdentifierNode) *ast.Node {
	if tx.currentNamespace != nil && !transformers.IsGeneratedIdentifier(tx.EmitContext(), node) && !transformers.IsLocalName(tx.EmitContext(), node) {
		location := tx.EmitContext().MostOriginal(node.AsNode())
		container := tx.resolver.GetReferencedExportContainer(location, false /*prefixLocals*/)
		if container != nil && ast.IsModuleDeclaration(container) {
			containerName := tx.getNamespaceContainerName(container)

			memberName := node.Clone(tx.Factory())
			tx.EmitContext().SetEmitFlags(memberName, printer.EFNoComments|printer.EFNoSourceMap)

			expression := tx.Factory().GetNamespaceMemberName(containerName, memberName, printer.NameOptions{AllowSourceMaps: true})
			tx.EmitContext().AssignCommentAndSourceMapRanges(expression, node.AsNode())
			return expression
		}
	}
	return node
}

func (tx *RuntimeSyntaxTransformer) createExportStatementForDeclaration(node *ast.Declaration) *ast.Statement {
	exportName := tx.Factory().GetExternalModuleOrNamespaceExportName(tx.getNamespaceContainerName(tx.currentNamespace), node, false /*allowComments*/, true /*allowSourceMaps*/)
	localName := tx.Factory().GetLocalName(node)
	expression := tx.Factory().NewAssignmentExpression(exportName, localName)
	exportAssignmentSourceMapRange := node.Loc
	if node.Name() != nil {
		exportAssignmentSourceMapRange = exportAssignmentSourceMapRange.WithPos(node.Name().Pos())
	}
	tx.EmitContext().SetSourceMapRange(expression, exportAssignmentSourceMapRange)

	statement := tx.Factory().NewExpressionStatement(expression)
	exportStatementSourceMapRange := node.Loc.WithPos(-1)
	tx.EmitContext().SetSourceMapRange(statement, exportStatementSourceMapRange)
	return statement
}

func (tx *RuntimeSyntaxTransformer) createExportAssignment(name *ast.IdentifierNode, expression *ast.Expression, exportAssignmentSourceMapRange core.TextRange, original *ast.Node) *ast.Expression {
	exportName := tx.getNamespaceQualifiedProperty(tx.getNamespaceContainerName(tx.currentNamespace), name)
	exportAssignment := tx.Factory().NewAssignmentExpression(exportName, expression)
	tx.EmitContext().SetOriginal(exportAssignment, original)
	tx.EmitContext().SetSourceMapRange(exportAssignment, exportAssignmentSourceMapRange)
	return exportAssignment
}

func (tx *RuntimeSyntaxTransformer) createExportStatement(name *ast.IdentifierNode, expression *ast.Expression, exportAssignmentSourceMapRange core.TextRange, exportStatementSourceMapRange core.TextRange, original *ast.Node) *ast.Statement {
	exportStatement := tx.Factory().NewExpressionStatement(tx.createExportAssignment(name, expression, exportAssignmentSourceMapRange, original))
	tx.EmitContext().SetOriginal(exportStatement, original)
	tx.EmitContext().SetSourceMapRange(exportStatement, exportStatementSourceMapRange)
	return exportStatement
}

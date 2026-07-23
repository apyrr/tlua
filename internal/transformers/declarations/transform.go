package declarations

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/modulespecifiers"
	"github.com/apyrr/tlua/internal/nodebuilder"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/scanner"
	"github.com/apyrr/tlua/internal/transformers"
	"github.com/apyrr/tlua/internal/tspath"
)

type ReferencedFilePair struct {
	file *ast.SourceFile
	ref  *ast.FileReference
}

type OutputPaths interface {
	DeclarationFilePath() string
	JsFilePath() string
}

// Used to be passed in the TransformationContext, which is now just an EmitContext
type DeclarationEmitHost interface {
	modulespecifiers.ModuleSpecifierGenerationHost
	GetCurrentDirectory() string
	UseCaseSensitiveFileNames() bool
	GetSourceFileFromReference(origin *ast.SourceFile, ref *ast.FileReference) *ast.SourceFile

	GetOutputPathsFor(file *ast.SourceFile, forceDtsPaths bool) OutputPaths
	GetEffectiveDeclarationFlags(node *ast.Node, flags ast.ModifierFlags) ast.ModifierFlags
	GetEmitResolver() printer.EmitResolver
}

type DeclarationTransformer struct {
	transformers.Transformer
	host                DeclarationEmitHost
	compilerOptions     *core.CompilerOptions
	tracker             *SymbolTrackerImpl
	state               *SymbolTrackerSharedState
	resolver            printer.EmitResolver
	declarationFilePath string
	declarationMapPath  string

	needsDeclare                     bool
	needsScopeFixMarker              bool
	resultHasScopeMarker             bool
	enclosingDeclaration             *ast.Node
	resultHasExternalModuleIndicator bool
	suppressNewDiagnosticContexts    bool
	witnessedCjsExports              collections.Set[string]
	lateStatementReplacementMap      map[ast.NodeId]*ast.Node
	expandoHosts                     map[ast.NodeId]*ast.Node               // store the result of transforming expando hosts so they can be inserted later if the host is actually referenced
	expandoMembers                   map[ast.NodeId][]*ast.Node             // store any found expando _members_ after transforming them so *if* the host is referenced, they can be emitted alongside it
	deferredExpandoAssignments       map[ast.NodeId][]*ast.BinaryExpression // expando assignments whose host wasn't visible when collected, processed if the host is late-marked visible
	rawReferencedFiles               []ReferencedFilePair
	rawTypeReferenceDirectives       []*ast.FileReference
	rawLibReferenceDirectives        []*ast.FileReference
	bindingNameVisitor               *ast.NodeVisitor
	expressionVisitor                *ast.NodeVisitor
	cjsExportAssignmentVisitor       *ast.NodeVisitor
	exportStrippingVisitor           *ast.NodeVisitor

	cjsExportAssignment     *ast.Node
	cjsExportMembers        []*ast.Node
	cjsExportAssignmentName *ast.Node // tracks the name node used for `export =` in CJS module.exports assignments
	declareStrippingVisitor *ast.NodeVisitor
}

// TODO: Convert to transformers.TransformerFactory signature to allow more automatic composition with other transforms
func NewDeclarationTransformer(host DeclarationEmitHost, context *printer.EmitContext, compilerOptions *core.CompilerOptions, declarationFilePath string, declarationMapPath string) *DeclarationTransformer {
	resolver := host.GetEmitResolver()
	state := &SymbolTrackerSharedState{isolatedDeclarations: compilerOptions.IsolatedDeclarations.IsTrue(), stripInternal: compilerOptions.StripInternal.IsTrue(), resolver: resolver}
	tracker := NewSymbolTracker(host, resolver, state)
	// TODO: Use new host GetOutputPathsFor method instead of passing in entrypoint paths (which will also better support bundled emit)
	tx := &DeclarationTransformer{
		host:                host,
		compilerOptions:     compilerOptions,
		tracker:             tracker,
		state:               state,
		resolver:            resolver,
		declarationFilePath: declarationFilePath,
		declarationMapPath:  declarationMapPath,
	}
	tx.state.reportExpandoFunctionErrors = func(node *ast.Node) {
		if !tx.state.isolatedDeclarations {
			return
		}
		props := resolver.GetPropertiesOfContainerFunction(node)
		for _, p := range props {
			if ast.IsExpandoPropertyDeclaration(p.ValueDeclaration) {
				errorTarget := p.ValueDeclaration
				if ast.IsBinaryExpression(errorTarget) {
					errorTarget = errorTarget.AsBinaryExpression().Left
				}
				tx.state.addDiagnostic(createDiagnosticForNode(errorTarget, diagnostics.Assigning_properties_to_functions_without_declaring_them_is_not_supported_with_isolatedDeclarations_Add_an_explicit_declaration_for_the_properties_assigned_to_this_function))
			}
		}
	}
	tx.NewTransformer(tx.visit, context)
	tx.bindingNameVisitor = tx.EmitContext().NewNodeVisitor(tx.visitBindingName)
	tx.expressionVisitor = tx.EmitContext().NewNodeVisitor(tx.visitNestedExpression)
	tx.exportStrippingVisitor = tx.EmitContext().NewNodeVisitor(tx.stripExportModifiers)
	tx.cjsExportAssignmentVisitor = tx.EmitContext().NewNodeVisitor(tx.visitCJSExportAssignments)
	tx.declareStrippingVisitor = tx.EmitContext().NewNodeVisitor(tx.stripDeclareModifiers)
	return tx
}

func (tx *DeclarationTransformer) GetDiagnostics() []*ast.Diagnostic {
	return tx.state.diagnostics
}

func (tx *DeclarationTransformer) shouldStripInternal(node *ast.Node) bool {
	return tx.state.stripInternal && node != nil && tx.isInternalDeclaration(node, tx.state.currentSourceFile)
}

func (tx *DeclarationTransformer) isInternalDeclaration(node *ast.Node, sourceFile *ast.SourceFile) bool {
	if node == nil {
		return false
	}
	parseTreeNode := tx.EmitContext().MostOriginal(node)
	if !ast.IsParseTreeNode(parseTreeNode) {
		return false
	}
	if parseTreeNode.Kind == ast.KindParameter {
		params := parseTreeNode.Parent.Parameters()
		paramIdx := slices.IndexFunc(params, func(p *ast.ParameterDeclarationNode) bool {
			return p.AsNode() == parseTreeNode
		})
		var previousSibling *ast.Node
		if paramIdx > 0 {
			previousSibling = params[paramIdx-1].AsNode()
		}

		text := sourceFile.Text()
		var commentRanges []ast.CommentRange

		if previousSibling != nil {
			// to handle
			// ... parameters, /** @internal */
			// public param: string
			trailingPos := scanner.SkipTriviaEx(text, previousSibling.End()+1, &scanner.SkipTriviaOptions{StopAtComments: true})
			for comment := range scanner.GetTrailingCommentRanges(tx.Factory().AsNodeFactory(), text, trailingPos) {
				commentRanges = append(commentRanges, comment)
			}
			for comment := range scanner.GetLeadingCommentRanges(tx.Factory().AsNodeFactory(), text, node.Pos()) {
				commentRanges = append(commentRanges, comment)
			}
		} else {
			trailingPos := scanner.SkipTriviaEx(text, node.Pos(), &scanner.SkipTriviaOptions{StopAtComments: true})
			for comment := range scanner.GetTrailingCommentRanges(tx.Factory().AsNodeFactory(), text, trailingPos) {
				commentRanges = append(commentRanges, comment)
			}
		}

		if len(commentRanges) > 0 {
			return hasInternalAnnotation(commentRanges[len(commentRanges)-1], sourceFile)
		}
		return false
	}

	for commentRange := range tx.getLeadingCommentRangesOfNode(parseTreeNode, sourceFile) {
		if hasInternalAnnotation(commentRange, sourceFile) {
			return true
		}
	}
	return false
}

func (tx *DeclarationTransformer) getLeadingCommentRangesOfNode(node *ast.Node, sourceFile *ast.SourceFile) iter.Seq[ast.CommentRange] {
	if node == nil || node.Kind == ast.KindJsxText {
		return nil
	}
	return scanner.GetLeadingCommentRanges(tx.Factory().AsNodeFactory(), sourceFile.Text(), node.Pos())
}

func hasInternalAnnotation(commentRange ast.CommentRange, sourceFile *ast.SourceFile) bool {
	comment := sourceFile.Text()[commentRange.Pos():commentRange.End()]
	return strings.Contains(comment, "@internal")
}

const declarationEmitNodeBuilderFlags = nodebuilder.FlagsMultilineObjectLiterals |
	nodebuilder.FlagsUseTypeOfFunction |
	nodebuilder.FlagsUseStructuralFallback |
	nodebuilder.FlagsAllowEmptyTuple |
	nodebuilder.FlagsGenerateNamesForShadowedTypeParams |
	nodebuilder.FlagsNoTruncation

const declarationEmitInternalNodeBuilderFlags = nodebuilder.InternalFlagsAllowUnresolvedNames

// functions as both `visitDeclarationStatements` and `transformRoot`, utilitzing SyntaxList nodes
func (tx *DeclarationTransformer) visit(node *ast.Node) *ast.Node {
	if node == nil {
		return nil
	}
	switch node.Kind {
	case ast.KindSourceFile:
		return tx.visitSourceFile(node.AsSourceFile())
	// statements we keep but do something to
	case ast.KindFunctionDeclaration,
		ast.KindModuleDeclaration,
		ast.KindImportEqualsDeclaration,
		ast.KindInterfaceDeclaration,
		ast.KindJSTypeAliasDeclaration,
		ast.KindTypeAliasDeclaration,
		ast.KindVariableStatement,
		ast.KindImportDeclaration,
		ast.KindJSImportDeclaration,
		ast.KindExportDeclaration,
		ast.KindExportAssignment:
		return tx.visitDeclarationStatements(node)
	// statements we elide
	case ast.KindBreakStatement,
		ast.KindContinueStatement,
		ast.KindDebuggerStatement,
		ast.KindEmptyStatement,
		ast.KindForOfStatement,
		ast.KindNumericForStatement,
		ast.KindRepeatStatement,
		ast.KindIfStatement,
		ast.KindLabelStatement,
		ast.KindGotoStatement,
		ast.KindReturnStatement,
		ast.KindThrowStatement,
		ast.KindWhileStatement,
		ast.KindNotEmittedStatement,
		ast.KindBlock,
		ast.KindMissingDeclaration,
		ast.KindExpressionStatement:
		return nil
	// parts of things, things we just visit children of
	default:
		return tx.visitDeclarationSubtree(node)
	}
}

func throwDiagnostic(result printer.SymbolAccessibilityResult) *SymbolAccessibilityDiagnostic {
	panic("Diagnostic emitted without context")
}

func (tx *DeclarationTransformer) visitSourceFile(node *ast.SourceFile) *ast.Node {
	tx.cjsExportAssignmentName = nil
	if node.IsDeclarationFile {
		return node.AsNode()
	}

	// A Lua module's shape is its chunk's return value, and no ambient form can
	// describe that yet. Declaration files describe globals, which still work.
	if ast.IsExternalModule(node) {
		tx.state.addDiagnostic(createDiagnosticForNode(node.AsNode(), diagnostics.Declaration_emit_is_not_supported_for_a_Lua_module_yet))
		return tx.Factory().UpdateSourceFile(node, tx.Factory().NewNodeList(nil), node.EndOfFileToken).AsNode()
	}

	tx.needsDeclare = true
	tx.needsScopeFixMarker = false
	tx.resultHasScopeMarker = false
	tx.enclosingDeclaration = node.AsNode()
	tx.state.getSymbolAccessibilityDiagnostic = throwDiagnostic
	tx.resultHasExternalModuleIndicator = false
	tx.suppressNewDiagnosticContexts = false
	tx.state.lateMarkedStatements = make([]*ast.Node, 0)
	tx.lateStatementReplacementMap = make(map[ast.NodeId]*ast.Node)
	tx.expandoHosts = make(map[ast.NodeId]*ast.Node)
	tx.expandoMembers = make(map[ast.NodeId][]*ast.Node)
	tx.deferredExpandoAssignments = make(map[ast.NodeId][]*ast.BinaryExpression)
	tx.rawReferencedFiles = make([]ReferencedFilePair, 0)
	tx.rawTypeReferenceDirectives = make([]*ast.FileReference, 0)
	tx.rawLibReferenceDirectives = make([]*ast.FileReference, 0)
	tx.witnessedCjsExports.Clear()
	tx.state.currentSourceFile = node
	tx.collectFileReferences(node)
	tx.resolver.PrecalculateDeclarationEmitVisibility(tx.EmitContext().MostOriginal(node.AsNode()).AsSourceFile())
	updated := tx.transformSourceFile(node)
	tx.state.currentSourceFile = nil
	return updated
}

func (tx *DeclarationTransformer) collectFileReferences(sourceFile *ast.SourceFile) {
	tx.rawReferencedFiles = append(tx.rawReferencedFiles, core.Map(sourceFile.ReferencedFiles, func(ref *ast.FileReference) ReferencedFilePair { return ReferencedFilePair{file: sourceFile, ref: ref} })...)
	tx.rawTypeReferenceDirectives = append(tx.rawTypeReferenceDirectives, sourceFile.TypeReferenceDirectives...)
	tx.rawLibReferenceDirectives = append(tx.rawLibReferenceDirectives, sourceFile.LibReferenceDirectives...)
}

func nodeOrSyntaxListChildren(node *ast.Node) []*ast.Node {
	if ast.IsSyntaxList(node) {
		return node.AsSyntaxList().Children
	}
	return []*ast.Node{node}
}

func flattenSyntaxLists(nodes []*ast.Node) []*ast.Node {
	return core.FlatMap(nodes, nodeOrSyntaxListChildren)
}

func (tx *DeclarationTransformer) appendCjsExports(combinedStatements *ast.StatementList) *ast.StatementList {
	result := []*ast.Node{}
	if tx.cjsExportAssignment != nil {
		result = append(result, tx.cjsExportAssignment)
	}
	result = append(result, tx.cjsExportMembers...)
	result = append(result, combinedStatements.Nodes...)
	statementNodes := flattenSyntaxLists(result)
	if len(statementNodes) != len(combinedStatements.Nodes) {
		combinedStatements = tx.Factory().NewNodeList(statementNodes)
	}
	return combinedStatements
}

func (tx *DeclarationTransformer) transformSourceFile(node *ast.SourceFile) *ast.Node {
	tx.cjsExportAssignment = nil
	tx.cjsExportAssignmentName = nil
	tx.cjsExportMembers = nil
	defer func() {
		tx.cjsExportAssignment = nil
		tx.cjsExportAssignmentName = nil
		tx.cjsExportMembers = nil
	}()
	tx.cjsExportAssignmentVisitor.VisitNode(node.AsNode()) // collect nested module.exports= assignments
	tx.expressionVisitor.VisitNode(node.AsNode())          // collect expando members (requires any export assignment be located in advance)
	var combinedStatements *ast.StatementList
	statements := tx.Visitor().VisitNodes(node.Statements)
	combinedStatements = tx.transformAndReplaceLatePaintedStatements(statements)
	combinedStatements = tx.appendCjsExports(combinedStatements)
	combinedStatements.Loc = statements.Loc // setTextRange
	if ast.IsExternalOrCommonJSModule(node) {
		if ast.IsInJSFile(node.AsNode()) {
			if exportEquals := node.Symbol.Exports[ast.InternalSymbolNameExportEquals]; exportEquals != nil && len(exportEquals.Declarations) > 1 {
				for _, node := range exportEquals.Declarations {
					tx.state.addDiagnostic(createDiagnosticForNode(node, diagnostics.Multiple_module_exports_assignments_cannot_be_serialized_for_declaration_emit))
				}
			}
		}
		if !tx.resultHasExternalModuleIndicator || (tx.needsScopeFixMarker && !tx.resultHasScopeMarker) {
			marker := createEmptyExports(tx.Factory().AsNodeFactory())
			newList := append(combinedStatements.Nodes, marker)
			withMarker := tx.Factory().NewNodeList(newList)
			withMarker.Loc = combinedStatements.Loc
			combinedStatements = withMarker
		}
	}
	outputFilePath := tspath.GetDirectoryPath(tspath.NormalizeSlashes(tx.declarationFilePath))
	result := tx.Factory().UpdateSourceFile(node, combinedStatements, node.EndOfFileToken)
	result.AsSourceFile().LibReferenceDirectives = tx.getLibReferences()
	result.AsSourceFile().TypeReferenceDirectives = tx.getTypeReferences()
	result.AsSourceFile().IsDeclarationFile = true
	result.AsSourceFile().ReferencedFiles = tx.getReferencedFiles(outputFilePath)
	return result.AsNode()
}

func createEmptyExports(factory *ast.NodeFactory) *ast.Node {
	return factory.NewExportDeclaration(nil /*modifiers*/, false, factory.NewNamedExports(factory.NewNodeList([]*ast.Node{})), nil)
}

func (tx *DeclarationTransformer) transformAndReplaceLatePaintedStatements(statements *ast.StatementList) *ast.StatementList {
	// This is a `while` loop because `handleSymbolAccessibilityError` can see additional import aliases marked as visible during
	// error handling which must now be included in the output and themselves checked for errors.
	// For example:
	// ```
	// module A {
	//   export module Q {}
	//   import B = Q;
	//   import C = B;
	//   export import D = C;
	// }
	// ```
	// In such a scenario, only Q and D are initially visible, but we don't consider imports as private names - instead we say they if they are referenced they must
	// be recorded. So while checking D's visibility we mark C as visible, then we must check C which in turn marks B, completing the chain of
	// dependent imports and allowing a valid declaration file output. Today, this dependent alias marking only happens for internal import aliases.
	for true {
		if len(tx.state.lateMarkedStatements) == 0 {
			break
		}

		next := tx.state.lateMarkedStatements[0]
		tx.state.lateMarkedStatements = tx.state.lateMarkedStatements[1:]

		saveNeedsDeclare := tx.needsDeclare
		tx.needsDeclare = next.Parent != nil && ast.IsSourceFile(next.Parent)

		result := tx.transformTopLevelDeclaration(next)

		tx.needsDeclare = saveNeedsDeclare
		original := tx.EmitContext().MostOriginal(next)
		id := ast.GetNodeId(original)
		tx.lateStatementReplacementMap[id] = result
	}

	// And lastly, we need to get the final form of all those indetermine import declarations from before and add them to the output list
	// (and remove them from the set to examine for outter declarations)
	results := make([]*ast.Node, 0, len(statements.Nodes))
	for _, statement := range statements.Nodes {
		if !ast.IsLateVisibilityPaintedStatement(statement) {
			results = append(results, statement)
			continue
		}
		original := tx.EmitContext().MostOriginal(statement)
		id := ast.GetNodeId(original)
		replacement, ok := tx.lateStatementReplacementMap[id]
		if !ok {
			results = append(results, statement)
			continue // not replaced
		}
		if replacement == nil {
			continue // deleted
		}
		if replacement.Kind == ast.KindSyntaxList {
			if !tx.needsScopeFixMarker || !tx.resultHasExternalModuleIndicator {
				for _, elem := range replacement.AsSyntaxList().Children {
					if needsScopeMarker(elem) {
						tx.needsScopeFixMarker = true
					}
					if ast.IsSourceFile(statement.Parent) && ast.IsExternalModuleIndicator(elem) {
						tx.resultHasExternalModuleIndicator = true
					}
				}
			}
			results = append(results, replacement.AsSyntaxList().Children...)
		} else {
			if needsScopeMarker(replacement) {
				tx.needsScopeFixMarker = true
			}
			if ast.IsSourceFile(statement.Parent) && ast.IsExternalModuleIndicator(replacement) {
				tx.resultHasExternalModuleIndicator = true
			}
			results = append(results, replacement)
		}
	}

	return tx.Factory().NewNodeList(results)
}

func (tx *DeclarationTransformer) getReferencedFiles(outputFilePath string) (results []*ast.FileReference) {
	// Handle path rewrites for triple slash ref comments
	for _, pair := range tx.rawReferencedFiles {
		sourceFile := pair.file
		ref := pair.ref

		if !ref.Preserve {
			continue
		}

		file := tx.host.GetSourceFileFromReference(sourceFile, ref)
		if file == nil {
			continue
		}

		var declFileName string
		if file.IsDeclarationFile {
			declFileName = file.FileName()
		} else {
			paths := tx.host.GetOutputPathsFor(file, true)
			// Try to use output path for referenced file, or output js path if that doesn't exist, or the input path if all else fails
			declFileName = paths.DeclarationFilePath()
			if len(declFileName) == 0 {
				declFileName = paths.JsFilePath()
			}
			if len(declFileName) == 0 {
				declFileName = file.FileName()
			}
		}
		// Should only be missing if the source file is missing a fileName (at which point we can't name a reference to it anyway)
		// TODO: Shouldn't this be a crash or assert instead of a silent continue?
		if len(declFileName) == 0 {
			continue
		}

		fileName := tspath.GetRelativePathToDirectoryOrUrl(
			outputFilePath,
			declFileName,
			false, // TODO: Probably unsafe to assume this isn't a URL, but that's what strada does
			tspath.ComparePathsOptions{
				CurrentDirectory:          tx.host.GetCurrentDirectory(),
				UseCaseSensitiveFileNames: tx.host.UseCaseSensitiveFileNames(),
			},
		)

		results = append(results, &ast.FileReference{
			TextRange:      core.NewTextRange(-1, -1),
			FileName:       fileName,
			ResolutionMode: ref.ResolutionMode,
			Preserve:       ref.Preserve,
		})
	}
	return results
}

func (tx *DeclarationTransformer) getLibReferences() (result []*ast.FileReference) {
	// clone retained references
	for _, ref := range tx.rawLibReferenceDirectives {
		if !ref.Preserve {
			continue
		}
		result = append(result, &ast.FileReference{
			TextRange:      core.NewTextRange(-1, -1),
			FileName:       ref.FileName,
			ResolutionMode: ref.ResolutionMode,
			Preserve:       ref.Preserve,
		})
	}
	return result
}

func (tx *DeclarationTransformer) getTypeReferences() (result []*ast.FileReference) {
	// clone retained references
	for _, ref := range tx.rawTypeReferenceDirectives {
		if !ref.Preserve {
			continue
		}
		result = append(result, &ast.FileReference{
			TextRange:      core.NewTextRange(-1, -1),
			FileName:       ref.FileName,
			ResolutionMode: ref.ResolutionMode,
			Preserve:       ref.Preserve,
		})
	}
	return result
}

func (tx *DeclarationTransformer) setupDiagnosticContext(input *ast.Node) (bool, func()) {
	canProdiceDiagnostic := canProduceDiagnostics(input)
	oldWithinObjectLiteralType := tx.suppressNewDiagnosticContexts
	shouldEnterSuppressNewDiagnosticsContextContext := (input.Kind == ast.KindTypeLiteral || input.Kind == ast.KindMappedType) && !(input.Parent.Kind == ast.KindTypeAliasDeclaration || input.Parent.Kind == ast.KindJSTypeAliasDeclaration)

	oldDiag := tx.state.getSymbolAccessibilityDiagnostic
	if canProdiceDiagnostic && !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(input)
	}
	oldName := tx.state.errorNameNode

	if shouldEnterSuppressNewDiagnosticsContextContext {
		tx.suppressNewDiagnosticContexts = true
	}

	return canProdiceDiagnostic, func() {
		tx.state.getSymbolAccessibilityDiagnostic = oldDiag
		tx.state.errorNameNode = oldName
		tx.suppressNewDiagnosticContexts = oldWithinObjectLiteralType
	}
}

func (tx *DeclarationTransformer) visitDeclarationSubtree(input *ast.Node) *ast.Node {
	if tx.shouldStripInternal(input) {
		return nil
	}
	if ast.IsDeclaration(input) {
		if isDeclarationAndNotVisible(tx.EmitContext(), tx.resolver, input) {
			return nil
		}
		if ast.HasDynamicName(input) {
			if tx.state.isolatedDeclarations {
				// Classes and object literals usually elide properties with computed names that are not of a literal type
				// In isolated declarations TSC needs to error on these as we don't know the type in a DTE.
				if !tx.resolver.IsDefinitelyReferenceToGlobalSymbolObject(input.Name().Expression()) {
					if ast.IsObjectLiteralExpression(input.Parent) {
						tx.state.addDiagnostic(createDiagnosticForNode(input, diagnostics.Computed_property_names_on_class_or_object_literals_cannot_be_inferred_with_isolatedDeclarations))
						return nil
					} else if (ast.IsInterfaceDeclaration(input.Parent) || ast.IsTypeLiteralNode(input.Parent)) && !ast.IsEntityNameExpression(input.Name().Expression()) {
						// Type declarations just need to double-check that the input computed name is an entity name expression
						tx.state.addDiagnostic(createDiagnosticForNode(input, diagnostics.Computed_properties_must_be_number_or_string_literals_variables_or_dotted_expressions_with_isolatedDeclarations))
						return nil
					}
				}
			} else if !tx.resolver.IsLateBound(tx.EmitContext().ParseNode(input)) || !ast.IsEntityNameExpression(input.Name().Expression()) {
				return nil
			}
		}
	}

	// Elide implementation signatures from overload sets
	if ast.IsFunctionLike(input) && tx.resolver.IsImplementationOfOverload(input) {
		return nil
	}

	if ast.IsHeritageClause(input) && (len(input.AsHeritageClause().Types.Nodes) == 0 || (len(input.AsHeritageClause().Types.Nodes) == 1 && ast.NodeIsMissing(input.AsHeritageClause().Types.Nodes[0]))) {
		return nil
	}

	previousEnclosingDeclaration := tx.enclosingDeclaration
	if isEnclosingDeclaration(input) {
		tx.enclosingDeclaration = input
	}

	canProdiceDiagnostic, cleanupDiagnosticContext := tx.setupDiagnosticContext(input)
	defer cleanupDiagnosticContext()

	var result *ast.Node

	switch input.Kind {
	case ast.KindMappedType:
		result = tx.transformMappedTypeNode(input.AsMappedTypeNode())
	case ast.KindHeritageClause:
		result = tx.transformHeritageClause(input.AsHeritageClause())
	case ast.KindMethodSignature:
		result = tx.transformMethodSignatureDeclaration(input.AsMethodSignatureDeclaration())
	case ast.KindConstructSignature:
		result = tx.transformConstructSignatureDeclaration(input.AsConstructSignatureDeclaration())
	case ast.KindPropertySignature:
		result = tx.transformPropertySignatureDeclaration(input.AsPropertySignatureDeclaration())
	case ast.KindCallSignature:
		result = tx.transformCallSignatureDeclaration(input.AsCallSignatureDeclaration())
	case ast.KindIndexSignature:
		result = tx.transformIndexSignatureDeclaration(input.AsIndexSignatureDeclaration())
	case ast.KindVariableDeclaration:
		result = tx.transformVariableDeclaration(input.AsVariableDeclaration())
	case ast.KindTypeParameter:
		result = tx.transformTypeParameterDeclaration(input.AsTypeParameterDeclaration())
	case ast.KindExpressionWithTypeArguments:
		result = tx.transformExpressionWithTypeArguments(input.AsExpressionWithTypeArguments())
	case ast.KindTypeReference:
		result = tx.transformTypeReference(input.AsTypeReferenceNode())
	case ast.KindConditionalType:
		result = tx.transformConditionalTypeNode(input.AsConditionalTypeNode())
	case ast.KindFunctionType:
		result = tx.transformFunctionTypeNode(input.AsFunctionTypeNode())
	case ast.KindConstructorType:
		result = tx.transformConstructorTypeNode(input.AsConstructorTypeNode())
	case ast.KindImportType:
		result = tx.transformImportTypeNode(input.AsImportTypeNode())
	case ast.KindTypeQuery:
		tx.checkEntityNameVisibility(input.AsTypeQueryNode().ExprName, tx.enclosingDeclaration)
		result = tx.Visitor().VisitEachChild(input)
	case ast.KindTupleType:
		result = tx.Visitor().VisitEachChild(input)
		if result != nil {
			if transformers.IsOriginalNodeSingleLine(tx.EmitContext(), input) {
				tx.EmitContext().AddEmitFlags(result, printer.EFSingleLine)
			}
		}
	case ast.KindJSDocTypeExpression:
		result = tx.transformJSDocTypeExpression(input.AsJSDocTypeExpression())
	case ast.KindJSDocTypeLiteral:
		result = tx.transformJSDocTypeLiteral(input.AsJSDocTypeLiteral())
	case ast.KindJSDocPropertyTag:
		result = tx.transformJSDocPropertyTag(input.AsJSDocParameterOrPropertyTag())
	case ast.KindJSDocAllType:
		result = tx.transformJSDocAllType(input.AsJSDocAllType())
	case ast.KindJSDocNullableType:
		result = tx.transformJSDocNullableType(input.AsJSDocNullableType())
	case ast.KindJSDocNonNullableType:
		result = tx.transformJSDocNonNullableType(input.AsJSDocNonNullableType())
	case ast.KindJSDocOptionalType:
		result = tx.transformJSDocOptionalType(input.AsJSDocOptionalType())
	case ast.KindJSDocVariadicType:
		result = tx.transformJSDocVariadicType(input.AsJSDocVariadicType())
	default:
		result = tx.Visitor().VisitEachChild(input)
	}

	if result != nil && canProdiceDiagnostic && ast.HasDynamicName(input) {
		tx.checkName(input)
	}

	tx.enclosingDeclaration = previousEnclosingDeclaration
	return result
}

func (tx *DeclarationTransformer) checkName(node *ast.Node) {
	oldDiag := tx.state.getSymbolAccessibilityDiagnostic
	if !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNodeName(node)
	}
	tx.state.errorNameNode = node.Name()
	debug.Assert(ast.HasDynamicName(node)) // Should only be called with dynamic names
	entityName := node.Name().Expression()
	tx.checkEntityNameVisibility(entityName, tx.enclosingDeclaration)
	if !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = oldDiag
	}
	tx.state.errorNameNode = nil
}

func (tx *DeclarationTransformer) transformMappedTypeNode(input *ast.MappedTypeNode) *ast.Node {
	// handle missing template type nodes, since the printer does not
	var typeNode *ast.Node
	if input.Type == nil {
		typeNode = tx.Factory().NewKeywordTypeNode(ast.KindAnyKeyword)
	} else {
		typeNode = tx.Visitor().Visit(input.Type)
	}
	return tx.Factory().UpdateMappedTypeNode(
		input,
		input.ReadonlyToken,
		tx.Visitor().Visit(input.TypeParameter),
		tx.Visitor().Visit(input.NameType),
		input.QuestionToken,
		typeNode,
		nil,
	)
}

func (tx *DeclarationTransformer) transformHeritageClause(clause *ast.HeritageClause) *ast.Node {
	retainedClauses := core.Filter(clause.Types.Nodes, func(t *ast.Node) bool {
		return ast.IsEntityNameExpression(t.AsExpressionWithTypeArguments().Expression) ||
			(clause.Token == ast.KindExtendsKeyword && t.Expression().Kind == ast.KindNilKeyword)
	})
	if len(retainedClauses) == 0 {
		return nil // elide empty clause
	}
	if len(retainedClauses) == len(clause.Types.Nodes) {
		return tx.Visitor().VisitEachChild(clause.AsNode())
	}
	return tx.Factory().UpdateHeritageClause(
		clause,
		clause.Token,
		tx.Visitor().VisitNodes(tx.Factory().NewNodeList(retainedClauses)),
	)
}

func (tx *DeclarationTransformer) transformImportTypeNode(input *ast.ImportTypeNode) *ast.Node {
	if !ast.IsLiteralImportTypeNode(input.AsNode()) {
		return input.AsNode()
	}
	return tx.Factory().UpdateImportTypeNode(
		input,
		input.IsTypeOf,
		tx.Factory().UpdateLiteralTypeNode(
			input.Argument.AsLiteralTypeNode(),
			tx.rewriteModuleSpecifier(input.AsNode(), input.Argument.AsLiteralTypeNode().Literal),
		),
		input.Qualifier,
		tx.Visitor().VisitNodes(input.TypeArguments),
	)
}

func (tx *DeclarationTransformer) transformConstructorTypeNode(input *ast.ConstructorTypeNode) *ast.Node {
	return tx.Factory().UpdateConstructorTypeNode(
		input,
		tx.ensureModifiers(input.AsNode()),
		tx.Visitor().VisitNodes(input.TypeParameters),
		tx.updateParamList(input.AsNode(), input.Parameters),
		tx.Visitor().Visit(input.Type),
	)
}

func (tx *DeclarationTransformer) transformFunctionTypeNode(input *ast.FunctionTypeNode) *ast.Node {
	return tx.Factory().UpdateFunctionTypeNode(
		input,
		input.Modifiers(),
		tx.Visitor().VisitNodes(input.TypeParameters),
		tx.updateParamList(input.AsNode(), input.Parameters),
		tx.Visitor().Visit(input.Type),
	)
}

func (tx *DeclarationTransformer) transformConditionalTypeNode(input *ast.ConditionalTypeNode) *ast.Node {
	checkType := tx.Visitor().Visit(input.CheckType)
	extendsType := tx.Visitor().Visit(input.ExtendsType)
	oldEnclosingDecl := tx.enclosingDeclaration
	tx.enclosingDeclaration = input.TrueType
	trueType := tx.Visitor().Visit(input.TrueType)
	tx.enclosingDeclaration = oldEnclosingDecl
	falseType := tx.Visitor().Visit(input.FalseType)

	return tx.Factory().UpdateConditionalTypeNode(
		input,
		checkType,
		extendsType,
		trueType,
		falseType,
	)
}

func (tx *DeclarationTransformer) transformTypeReference(input *ast.TypeReferenceNode) *ast.Node {
	tx.checkEntityNameVisibility(input.TypeName, tx.enclosingDeclaration)
	return tx.Visitor().VisitEachChild(input.AsNode())
}

func (tx *DeclarationTransformer) transformExpressionWithTypeArguments(input *ast.ExpressionWithTypeArguments) *ast.Node {
	if ast.IsEntityName(input.Expression) || ast.IsEntityNameExpression(input.Expression) {
		tx.checkEntityNameVisibility(input.Expression, tx.enclosingDeclaration)
	}
	return tx.Visitor().VisitEachChild(input.AsNode())
}

func (tx *DeclarationTransformer) transformTypeParameterDeclaration(input *ast.TypeParameterDeclaration) *ast.Node {
	return tx.Visitor().VisitEachChild(input.AsNode())
}

func (tx *DeclarationTransformer) transformVariableDeclaration(input *ast.VariableDeclaration) *ast.Node {
	if tx.state.currentSourceFile.CommonJSModuleIndicator != nil && ast.IsVariableDeclarationInitializedToRequire(input.AsNode()) {
		return tx.transformCjsRequireVariableDeclaration(input)
	}
	if ast.IsBindingPattern(input.Name()) {
		return tx.recreateBindingPattern(input.Name().AsBindingPattern())
	}
	// Variable declaration types also suppress new diagnostic contexts, provided the contexts wouldn't be made for binding pattern types
	tx.suppressNewDiagnosticContexts = true
	return tx.Factory().UpdateVariableDeclaration(
		input,
		input.Name(),
		nil,
		tx.ensureType(input.AsNode(), false),
		tx.ensureNoInitializer(input.AsNode()),
	)
}

func (tx *DeclarationTransformer) transformCjsRequireVariableDeclaration(input *ast.VariableDeclaration) *ast.Node {
	specifier := tx.rewriteModuleSpecifier(input.AsNode(), input.Initializer.AsCallExpression().Arguments.Nodes[0])
	if ast.IsIdentifier(input.Name()) {
		// `const x = require("something")` -> `import x = require("something")`
		return tx.Factory().NewImportEqualsDeclaration(nil, false, input.Name(), tx.Factory().NewExternalModuleReference(specifier))
	} else if ast.IsArrayBindingPattern(input.Name()) {
		// TODO: Is this actually reachable? should we error on this?
		return nil
	} else { // object binding pattern

		// `const {x, y: z} = require("something")` -> `import {x, y as z} from "something"`
		b := input.Name().AsBindingPattern()
		var importSpecifiers []*ast.Node
		for _, elem := range b.Elements.Nodes {
			if !ast.IsIdentifier(elem.Name()) {
				continue // nested destructuring, bail
			}
			importSpecifiers = append(importSpecifiers, tx.Factory().NewImportSpecifier(false, elem.PropertyName(), elem.Name()))
		}
		return tx.Factory().NewImportDeclaration(
			nil,
			tx.Factory().NewImportClause(
				ast.KindUnknown,
				nil,
				tx.Factory().NewNamedImports(tx.Factory().NewNodeList(importSpecifiers)),
			),
			specifier,
		)
	}
}

func (tx *DeclarationTransformer) recreateBindingPattern(input *ast.BindingPattern) *ast.Node {
	var results []*ast.Node
	for _, elem := range input.Elements.Nodes {
		result := tx.recreateBindingElement(elem.AsBindingElement())
		if result == nil {
			continue
		}
		if result.Kind == ast.KindSyntaxList {
			results = append(results, result.AsSyntaxList().Children...)
		} else {
			results = append(results, result)
		}
	}
	if len(results) == 0 {
		return nil
	}
	if len(results) == 1 {
		return results[0]
	}
	return tx.Factory().NewSyntaxList(results)
}

func (tx *DeclarationTransformer) recreateBindingElement(e *ast.BindingElement) *ast.Node {
	if e.Name() == nil {
		return nil
	}
	if !getBindingNameVisible(tx.resolver, e.AsNode()) {
		return nil
	}
	if ast.IsBindingPattern(e.Name()) {
		return tx.recreateBindingPattern(e.Name().AsBindingPattern())
	}
	return tx.Factory().NewVariableDeclaration(
		e.Name(),
		nil,
		tx.ensureType(e.AsNode(), false),
		nil, // TODO: possible strada bug - not emitting const initialized binding pattern elements?
	)
}

func (tx *DeclarationTransformer) transformIndexSignatureDeclaration(input *ast.IndexSignatureDeclaration) *ast.Node {
	t := tx.Visitor().Visit(input.Type)
	if t == nil {
		t = tx.Factory().NewKeywordTypeNode(ast.KindAnyKeyword)
	}
	return tx.Factory().UpdateIndexSignatureDeclaration(
		input,
		tx.ensureModifiers(input.AsNode()),
		tx.updateParamList(input.AsNode(), input.Parameters),
		t,
	)
}

func (tx *DeclarationTransformer) transformCallSignatureDeclaration(input *ast.CallSignatureDeclaration) *ast.Node {
	return tx.Factory().UpdateCallSignatureDeclaration(
		input,
		tx.ensureTypeParams(input.AsNode(), input.TypeParameters),
		tx.updateParamList(input.AsNode(), input.Parameters),
		tx.ensureType(input.AsNode(), false),
	)
}

func (tx *DeclarationTransformer) transformPropertySignatureDeclaration(input *ast.PropertySignatureDeclaration) *ast.Node {
	if ast.IsPrivateIdentifier(input.Name()) {
		return nil
	}
	result := tx.Factory().UpdatePropertySignatureDeclaration(
		input,
		tx.ensureModifiers(input.AsNode()),
		input.Name(),
		input.PostfixToken,
		tx.ensureType(input.AsNode(), false),
		tx.ensureNoInitializer(input.AsNode()), // TODO: possible strada bug (fixed here) - const property signatures never initialized
	)
	tx.preservePartialJsDoc(result, input.AsNode())
	return result
}

func (tx *DeclarationTransformer) transformConstructSignatureDeclaration(input *ast.ConstructSignatureDeclaration) *ast.Node {
	return tx.Factory().UpdateConstructSignatureDeclaration(
		input,
		tx.ensureTypeParams(input.AsNode(), input.TypeParameters),
		tx.updateParamList(input.AsNode(), input.Parameters),
		tx.ensureType(input.AsNode(), false),
	)
}

func (tx *DeclarationTransformer) transformMethodSignatureDeclaration(input *ast.MethodSignatureDeclaration) *ast.Node {
	if ast.IsPrivateIdentifier(input.Name()) {
		return nil
	} else {
		return tx.Factory().UpdateMethodSignatureDeclaration(
			input,
			tx.ensureModifiers(input.AsNode()),
			input.Name(),
			input.PostfixToken,
			tx.ensureTypeParams(input.AsNode(), input.TypeParameters),
			tx.updateParamList(input.AsNode(), input.Parameters),
			tx.ensureType(input.AsNode(), false),
		)
	}
}

func (tx *DeclarationTransformer) visitDeclarationStatements(input *ast.Node) *ast.Node {
	if tx.shouldStripInternal(input) {
		return nil
	}
	switch input.Kind {
	case ast.KindExportDeclaration:
		if ast.IsSourceFile(input.Parent) {
			tx.resultHasExternalModuleIndicator = true
		}
		tx.resultHasScopeMarker = true
		// Rewrite external module names if necessary
		return tx.Factory().UpdateExportDeclaration(
			input.AsExportDeclaration(),
			input.Modifiers(),
			input.IsTypeOnly(),
			input.AsExportDeclaration().ExportClause,
			tx.rewriteModuleSpecifier(input, input.ModuleSpecifier()),
		)
	case ast.KindExportAssignment:
		return tx.transformExportAssignment(input, input, input.Expression(), input.AsExportAssignment().IsExportEquals)
	default:
		id := ast.GetNodeId(tx.EmitContext().MostOriginal(input))
		if tx.lateStatementReplacementMap[id] == nil {
			// Don't actually transform yet; just leave as original node - will be elided/swapped by late pass
			tx.lateStatementReplacementMap[id] = tx.transformTopLevelDeclaration(input)
		}
		return input
	}
}

func (tx *DeclarationTransformer) tryGetNameOfAssignedExpression(unwrapped *ast.Node) *ast.Node {
	var nameNode *ast.Node
	var nameText string
	if !ast.IsPropertyAccessExpression(unwrapped) && unwrapped.Name() != nil {
		nameText = unwrapped.Name().Text()
	} else if ast.IsIdentifier(unwrapped) {
		nameText = unwrapped.Text()
	}
	if nameText != "" && nameText != "default" {
		if tx.resolver.IsNameResolvable(tx.enclosingDeclaration, nameText) {
			// create a unique name that shares the same text as its' base
			nameNode = tx.Factory().NewUniqueNameEx(nameText, printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsOptimistic})
		} else {
			// use the node's name as-is, since it's not otherwise in-scope
			nameNode = tx.Factory().NewIdentifier(nameText)
		}
	}
	return nameNode
}

func (tx *DeclarationTransformer) getNameOfExportedAssignedExpression(unwrapped *ast.Node, isExportEquals bool) *ast.Node {
	nameNode := tx.tryGetNameOfAssignedExpression(unwrapped)
	if nameNode == nil {
		// fallback to a default name
		if isExportEquals && ast.IsSourceFileJS(tx.state.currentSourceFile) {
			// only JS files prefer to use `_exports` for export assignments - TS has always used `_default` for both `export=` and `export default`
			nameNode = tx.Factory().NewUniqueNameEx("_exports", printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsOptimistic})
		} else {
			nameNode = tx.Factory().NewUniqueNameEx("_default", printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsOptimistic})
		}
	}
	tx.cjsExportAssignmentName = nameNode
	return nameNode
}

func (tx *DeclarationTransformer) transformExportAssignment(input *ast.Node, assignment *ast.Node, expression *ast.Node, isExportEquals bool) *ast.Node {
	if ast.IsSourceFile(input.Parent) {
		tx.resultHasExternalModuleIndicator = true
	}
	tx.resultHasScopeMarker = true
	if ast.IsIdentifier(expression) && (ast.IsSourceFile(input.Parent) || ast.IsModuleBlock(input.Parent)) {
		exportAssignment := tx.Factory().NewExportAssignment(nil, isExportEquals, nil, expression)
		tx.preserveJsDoc(exportAssignment, input)
		return exportAssignment
	}

	unwrapped := ast.SkipOuterExpressions(expression, ast.OEKExpressionTypePassthrough)
	newId := tx.getNameOfExportedAssignedExpression(unwrapped, isExportEquals)
	if ast.IsFunctionLike(unwrapped) {
		// Promote function or arrow function expressions to a function declaration
		var mods []*ast.Node
		if tx.needsDeclare {
			mods = append(mods, tx.Factory().NewModifier(ast.KindDeclareKeyword))
		}
		fullSignatureType := assignment.Type()
		funcDecl := tx.transformFunctionLikeToDeclaration(unwrapped, newId, tx.Factory().NewModifierList(mods), fullSignatureType)
		tx.preserveJsDoc(funcDecl, input)
		// Reuse the same name node for the export so unique names resolve consistently
		exportAssignment := tx.Factory().NewExportAssignment(nil, isExportEquals, nil, newId)
		tx.removeAllComments(exportAssignment)
		return tx.Factory().NewSyntaxList([]*ast.Node{exportAssignment, funcDecl})
	}

	// expression is non-identifier, create _default typed variable to reference
	tx.state.getSymbolAccessibilityDiagnostic = func(_ printer.SymbolAccessibilityResult) *SymbolAccessibilityDiagnostic {
		return &SymbolAccessibilityDiagnostic{
			diagnosticMessage: diagnostics.Default_export_of_the_module_has_or_is_using_private_name_0,
			errorNode:         input,
		}
	}
	tx.cjsExportAssignmentName = newId
	tx.tracker.PushErrorFallbackNode(assignment)
	// `declare const x = <literal>` is deleted from the grammar; keep the
	// literal type as the annotation instead of an initializer.
	var type_ *ast.Node
	if ast.IsPrimitiveLiteralValue(unwrapParenthesizedExpression(expression)) {
		if lit := tx.resolver.CreateLiteralConstValue(tx.EmitContext(), tx.EmitContext().ParseNode(assignment), tx.tracker); lit != nil {
			type_ = tx.Factory().NewLiteralTypeNode(lit)
		}
	}
	if type_ == nil {
		type_ = tx.ensureType(assignment, false)
	}
	varDecl := tx.Factory().NewVariableDeclaration(newId, nil, type_, nil /*initializer*/)
	tx.tracker.PopErrorFallbackNode()
	var modList *ast.ModifierList
	if tx.needsDeclare {
		modList = tx.Factory().NewModifierList([]*ast.Node{tx.Factory().NewModifier(ast.KindDeclareKeyword)})
	} else {
		modList = tx.Factory().NewModifierList([]*ast.Node{})
	}
	statement := tx.Factory().NewVariableStatement(modList, tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), ast.NodeFlagsConst))
	exportAssignment := tx.Factory().NewExportAssignment(nil, isExportEquals, nil, newId)
	// Remove comments from the export declaration and copy them onto the synthetic _default declaration
	tx.preserveJsDoc(statement, input)
	return tx.Factory().NewSyntaxList([]*ast.Node{statement, exportAssignment})
}

func (tx *DeclarationTransformer) transformFunctionLikeToDeclaration(unwrapped *ast.Node, funcName *ast.Node, mods *ast.ModifierList, fullSignatureType *ast.Node) *ast.Node {
	d := unwrapped.FunctionLikeData()
	sig := d.FullSignature
	if sig == nil {
		sig = fullSignatureType
	}
	if sig == nil {
		return tx.Factory().NewFunctionDeclaration(
			mods,
			nil, /*target*/
			nil, /*colonToken*/
			funcName,
			tx.ensureTypeParams(unwrapped, d.TypeParameters),
			tx.updateParamList(unwrapped, d.Parameters),
			tx.ensureType(unwrapped, false),
			tx.Visitor().VisitNode(sig),
			nil,
		)
	}
	// If a full signature type node is present, emit as a variable statement to reuse it.
	return tx.Factory().NewVariableStatement(
		mods,
		tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{tx.Factory().NewVariableDeclaration(funcName, nil, tx.Visitor().VisitNode(sig), nil)}), ast.NodeFlagsConst),
	)
}

func (tx *DeclarationTransformer) transformBinaryExpressionToExportDeclaration(input *ast.Node, name *ast.Node) *ast.Node {
	propertyName := input.AsBinaryExpression().Right

	// track alias target so referenced declarations are included in the output
	tx.tracker.handleSymbolAccessibilityError(tx.resolver.IsEntityNameVisible(propertyName, tx.enclosingDeclaration))

	if ast.IsIdentifier(name) && propertyName.Text() == name.Text() {
		propertyName = nil
	}

	return tx.Factory().NewExportDeclaration(
		nil,
		false,
		tx.Factory().NewNamedExports(tx.Factory().NewNodeList([]*ast.Node{tx.Factory().NewExportSpecifier(false, propertyName, name)})),
		nil,
	)
}

func (tx *DeclarationTransformer) transformCommonJSExport(input *ast.Node, name *ast.Node) *ast.Node {
	res := tx.transformCommonJSExportWorker(input, name)
	if res == nil {
		return res
	}
	return tx.wrapInCJSExportNamespace(res)
}

func (tx *DeclarationTransformer) transformCommonJSExportWorker(input *ast.Node, name *ast.Node) *ast.Node {
	var nameText string
	if ast.IsIdentifier(name) || ast.IsStringLiteral(name) {
		nameText = name.Text()
	}
	if tx.witnessedCjsExports.Has(nameText) && nameText != "" {
		return nil // Already emitted this export name
	}
	tx.witnessedCjsExports.Add(nameText)
	tx.resultHasExternalModuleIndicator = true
	tx.resultHasScopeMarker = true
	// only transform cjs exports to shorthand at the top-level of a source file, otherwise we uniformly emit nested exports with a type annotation
	if isCommonJSAliasExport(input) && ast.IsExpressionStatement(input.Parent) && ast.IsSourceFile(input.Parent.Parent) {
		// export { name }
		// export { source as name }
		return tx.transformBinaryExpressionToExportDeclaration(input, name)
	}

	if ast.IsIdentifier(name) {
		if name.Text() == "default" {
			// const _default: Type; export default _default;
			newId := tx.Factory().NewUniqueNameEx("_default", printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsOptimistic})
			tx.state.getSymbolAccessibilityDiagnostic = func(_ printer.SymbolAccessibilityResult) *SymbolAccessibilityDiagnostic {
				return &SymbolAccessibilityDiagnostic{
					diagnosticMessage: diagnostics.Default_export_of_the_module_has_or_is_using_private_name_0,
					errorNode:         input,
				}
			}
			tx.tracker.PushErrorFallbackNode(input)
			type_ := tx.ensureType(input, false)
			varDecl := tx.Factory().NewVariableDeclaration(newId, nil, type_, nil)
			tx.tracker.PopErrorFallbackNode()
			var modList *ast.ModifierList
			if tx.needsDeclare {
				modList = tx.Factory().NewModifierList([]*ast.Node{tx.Factory().NewModifier(ast.KindDeclareKeyword)})
			} else {
				modList = tx.Factory().NewModifierList([]*ast.Node{})
			}
			statement := tx.Factory().NewVariableStatement(modList, tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), ast.NodeFlagsConst))

			assignment := tx.Factory().NewExportAssignment(input.Modifiers(), false, nil, newId)
			// Remove comments from the export declaration and copy them onto the synthetic _default declaration
			tx.preserveJsDoc(statement, input)
			tx.removeAllComments(assignment)
			return tx.Factory().NewSyntaxList([]*ast.Node{statement, assignment})
		} else if tx.host.GetEmitResolver().GetReferencedValueDeclaration(name) == input || tx.host.GetEmitResolver().GetReferencedValueDeclaration(name) == nil {
			// only inline to a export var if the `name` lookup points at this assignment or nothing - if it points at something else, we must use a temp name
			// export var name: Type
			tx.tracker.PushErrorFallbackNode(input)
			type_ := tx.ensureType(input, false)
			varDecl := tx.Factory().NewVariableDeclaration(name, nil, type_, nil)
			tx.tracker.PopErrorFallbackNode()
			var modList *ast.ModifierList
			if tx.needsDeclare {
				modList = tx.Factory().NewModifierList([]*ast.Node{tx.Factory().NewModifier(ast.KindExportKeyword), tx.Factory().NewModifier(ast.KindDeclareKeyword)})
			} else {
				modList = tx.Factory().NewModifierList([]*ast.Node{tx.Factory().NewModifier(ast.KindExportKeyword)})
			}
			// Const, not None: declaration-emit lists are ambient, and the
			// printer's Luau `declare name: Type;` form is keyed on Const.
			return tx.Factory().NewVariableStatement(modList, tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), ast.NodeFlagsConst))
		}
	}
	// const _exported: Type; export {_exported as "name"};
	newId := tx.Factory().NewUniqueNameEx("_exported", printer.AutoGenerateOptions{Flags: printer.GeneratedIdentifierFlagsOptimistic})
	tx.state.getSymbolAccessibilityDiagnostic = func(_ printer.SymbolAccessibilityResult) *SymbolAccessibilityDiagnostic {
		return &SymbolAccessibilityDiagnostic{
			diagnosticMessage: diagnostics.Default_export_of_the_module_has_or_is_using_private_name_0,
			errorNode:         input,
		}
	}
	tx.tracker.PushErrorFallbackNode(input)
	type_ := tx.ensureType(input, false)
	varDecl := tx.Factory().NewVariableDeclaration(newId, nil, type_, nil)
	tx.tracker.PopErrorFallbackNode()
	var modList *ast.ModifierList
	if tx.needsDeclare {
		modList = tx.Factory().NewModifierList([]*ast.Node{tx.Factory().NewModifier(ast.KindDeclareKeyword)})
	} else {
		modList = tx.Factory().NewModifierList([]*ast.Node{})
	}
	statement := tx.Factory().NewVariableStatement(modList, tx.Factory().NewVariableDeclarationList(tx.Factory().NewNodeList([]*ast.Node{varDecl}), ast.NodeFlagsConst))

	assignment := tx.Factory().NewExportDeclaration(nil, false, tx.Factory().NewNamedExports(tx.Factory().NewNodeList([]*ast.Node{tx.Factory().NewExportSpecifier(false, newId, name)})), nil)
	// Remove comments from the export declaration and copy them onto the synthetic _default declaration
	tx.preserveJsDoc(statement, input)
	tx.removeAllComments(assignment)
	return tx.Factory().NewSyntaxList([]*ast.Node{statement, assignment})
}

func (tx *DeclarationTransformer) wrapInCJSExportNamespace(content *ast.Node) *ast.Node {
	if tx.cjsExportAssignmentName == nil {
		return content
	}
	// Reuse the same name node so unique names resolve consistently with the class/export
	nsName := tx.cjsExportAssignmentName
	var members []*ast.Node
	if content.Kind == ast.KindSyntaxList {
		members = content.AsSyntaxList().Children
	} else {
		members = []*ast.Node{content}
	}
	var nsMods []*ast.Node
	if tx.needsDeclare {
		nsMods = append(nsMods, tx.Factory().NewModifier(ast.KindDeclareKeyword))
	}
	members, _ = tx.declareStrippingVisitor.VisitSlice(members)
	return tx.Factory().NewModuleDeclaration(
		tx.Factory().NewModifierList(nsMods),
		ast.KindNamespaceKeyword,
		nsName,
		tx.Factory().NewModuleBlock(tx.Factory().NewNodeList(members)),
	)
}

func isCommonJSAliasExport(node *ast.Node) bool {
	if ast.IsBinaryExpression(node) && ast.IsIdentifier(node.AsBinaryExpression().Right) {
		if symbol := node.Symbol(); symbol != nil && len(symbol.Declarations) == 1 {
			return true
		}
	}
	return false
}

func (tx *DeclarationTransformer) rewriteModuleSpecifier(parent *ast.Node, input *ast.Node) *ast.Node {
	if input == nil {
		return nil
	}
	tx.resultHasExternalModuleIndicator = tx.resultHasExternalModuleIndicator || (parent.Kind != ast.KindModuleDeclaration && parent.Kind != ast.KindImportType)
	return input
}

func (tx *DeclarationTransformer) preserveJsDoc(updated *ast.Node, original *ast.Node) {
	// Copy comment range from original to updated node so JSDoc comments are preserved
	tx.EmitContext().AssignCommentRange(updated, original)
}

func (tx *DeclarationTransformer) preservePartialJsDoc(updated *ast.Node, original *ast.Node) {
	if original.Flags&ast.NodeFlagsReparsed == 0 {
		return
	}
	jsdoc := core.FirstOrNil(original.EagerJSDoc(ast.GetSourceFileOfNode(original)))
	if jsdoc == nil {
		return
	}
	description := scanner.GetTextOfJSDocComment(jsdoc.AsJSDoc().Comment)
	if description == "" {
		return
	}
	comment := "*\n * " + strings.ReplaceAll(description, "\n", "\n * ") + "\n "
	tx.EmitContext().AddSyntheticLeadingComment(updated, ast.KindMultiLineCommentTrivia, comment, true /*hasTrailingNewLine*/)
}

func (tx *DeclarationTransformer) removeAllComments(node *ast.Node) {
	tx.EmitContext().AddEmitFlags(node, printer.EFNoComments)
	// !!! TODO: Also remove synthetic trailing/leading comments added by transforms
	// emitNode.leadingComments = undefined;
	// emitNode.trailingComments = undefined;
}

func (tx *DeclarationTransformer) ensureType(node *ast.Node, ignorePrivate bool) *ast.Node {
	if !ignorePrivate && tx.host.GetEffectiveDeclarationFlags(tx.EmitContext().ParseNode(node), ast.ModifierFlagsPrivate) != 0 {
		// Private nodes emit no types.
		return nil
	}

	if tx.shouldPrintWithInitializer(node) {
		if ast.IsVariableDeclaration(node) {
			// `declare const x = "lit"` is deleted from the grammar, so a
			// literal const variable keeps its literal TYPE as the annotation:
			// `declare x: "lit"`.
			if lit := tx.resolver.CreateLiteralConstValue(tx.EmitContext(), tx.EmitContext().ParseNode(node), tx.tracker); lit != nil {
				return tx.Factory().NewLiteralTypeNode(lit)
			}
		}
		// Literal const parameters and properties have an initializer ensured
		// rather than a type.
		return nil
	}

	if ast.IsVarargParameter(node) && node.Type() == nil {
		// A bare `...` already says everything: it is an `any` pack. Serializing
		// the symbol's type instead would print the array it carries (`...: any[]`),
		// which reparses as a pack of *arrays*.
		return nil
	}

	// Should be removed createTypeOfDeclaration will actually now reuse the existing annotation so there is no real need to duplicate type walking
	// Left in for now to minimize diff during syntactic type node builder refactor
	if !ast.IsExportAssignment(node) && !ast.IsBindingElement(node) && node.Type() != nil && (!ast.IsParameterDeclaration(node) || !tx.resolver.RequiresAddingImplicitUndefined(node, nil)) {
		if tx.state.currentSourceFile.IsJS() {
			// JS types have a heap of constructs we can't directly emit into .d.ts files; the node builder contains logic to remap those where possible, so we invoke it here
			// In strada we always built js declarations symbolically, so all js type nodes went through this postprocessing
			res := tx.resolver.TryJSTypeNodeToTypeNode(tx.EmitContext(), node.Type(), tx.enclosingDeclaration, declarationEmitNodeBuilderFlags, declarationEmitInternalNodeBuilderFlags, tx.tracker)
			if res != nil {
				return res
			}
			// otherwise, fall back to full serialization
		} else {
			return tx.Visitor().Visit(node.Type())
		}
	}

	oldErrorNameNode := tx.state.errorNameNode
	tx.state.errorNameNode = node.Name()
	var oldDiag GetSymbolAccessibilityDiagnostic
	if !tx.suppressNewDiagnosticContexts {
		oldDiag = tx.state.getSymbolAccessibilityDiagnostic
		if canProduceDiagnostics(node) {
			tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(node)
		}
	}
	var typeNode *ast.Node

	flags := declarationEmitNodeBuilderFlags
	if ast.HasInferredType(node) {
		typeNode = tx.resolver.CreateTypeOfDeclaration(tx.EmitContext(), node, tx.enclosingDeclaration, flags, declarationEmitInternalNodeBuilderFlags, tx.tracker)
	} else if ast.IsFunctionLike(node) {
		typeNode = tx.resolver.CreateReturnTypeOfSignatureDeclaration(tx.EmitContext(), node, tx.enclosingDeclaration, flags, declarationEmitInternalNodeBuilderFlags, tx.tracker)
	} else {
		debug.AssertNever(node)
	}

	tx.state.errorNameNode = oldErrorNameNode
	if !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = oldDiag
	}
	if typeNode == nil {
		return tx.Factory().NewKeywordTypeNode(ast.KindAnyKeyword)
	}
	return typeNode
}

func (tx *DeclarationTransformer) shouldPrintWithInitializer(node *ast.Node) bool {
	return canHaveLiteralInitializer(tx.host, node) && node.Initializer() != nil && tx.resolver.IsLiteralConstDeclaration(tx.EmitContext().MostOriginal(node))
}

func (tx *DeclarationTransformer) checkEntityNameVisibility(entityName *ast.Node, enclosingDeclaration *ast.Node) {
	visibilityResult := tx.resolver.IsEntityNameVisible(entityName, enclosingDeclaration)
	tx.tracker.handleSymbolAccessibilityError(visibilityResult)
}

// Transforms the direct child of a source file into zero or more replacement statements
func (tx *DeclarationTransformer) transformTopLevelDeclaration(input *ast.Node) *ast.Node {
	if len(tx.state.lateMarkedStatements) > 0 {
		// Remove duplicates of the current statement from the deferred work queue (this was done via orderedRemoveItem in strada - why? to ensure the same backing array? microop?)
		tx.state.lateMarkedStatements = core.Filter(tx.state.lateMarkedStatements, func(node *ast.Node) bool { return node != input })
	}
	if tx.shouldStripInternal(input) {
		return nil
	}
	if input.Kind == ast.KindImportEqualsDeclaration {
		return tx.transformImportEqualsDeclaration(input.AsImportEqualsDeclaration())
	}
	if input.Kind == ast.KindImportDeclaration || input.Kind == ast.KindJSImportDeclaration {
		res := tx.transformImportDeclaration(input.AsImportDeclaration())
		if res != nil && res.Kind != ast.KindImportDeclaration {
			res := res.Clone(tx.Factory())
			res.Kind = ast.KindImportDeclaration
			return res
		}
		return res
	}
	if ast.IsDeclaration(input) && isDeclarationAndNotVisible(tx.EmitContext(), tx.resolver, input) {
		return nil
	}

	// !!! TODO: JSDoc support
	// if (isJSDocImportTag(input)) return;

	// Elide implementation signatures from overload sets
	if ast.IsFunctionLike(input) && tx.resolver.IsImplementationOfOverload(input) {
		return nil
	}
	original := tx.EmitContext().MostOriginal(input)
	id := ast.GetNodeId(original)
	_, isExpandoHost := tx.expandoHosts[id]
	_, hasDeferredExpandoAssignments := tx.deferredExpandoAssignments[id]
	if isExpandoHost || hasDeferredExpandoAssignments {
		return tx.createFullExpandoBlock(id)
	}

	previousEnclosingDeclaration := tx.enclosingDeclaration
	if isEnclosingDeclaration(input) {
		tx.enclosingDeclaration = input
	}

	canProdiceDiagnostic := canProduceDiagnostics(input)
	oldDiag := tx.state.getSymbolAccessibilityDiagnostic
	oldName := tx.state.errorNameNode
	if canProdiceDiagnostic {
		tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(input)
	}
	saveNeedsDeclare := tx.needsDeclare

	var result *ast.Node
	switch input.Kind {
	case ast.KindTypeAliasDeclaration, ast.KindJSTypeAliasDeclaration:
		result = tx.transformTypeAliasDeclaration(input.AsTypeAliasDeclaration())
	case ast.KindInterfaceDeclaration:
		result = tx.transformInterfaceDeclaration(input.AsInterfaceDeclaration())
	case ast.KindFunctionDeclaration:
		result = tx.transformFunctionDeclaration(input.AsFunctionDeclaration())
	case ast.KindModuleDeclaration:
		result = tx.transformModuleDeclaration(input.AsModuleDeclaration())
	case ast.KindVariableStatement:
		result = tx.transformVariableStatement(input.AsVariableStatement())
	default:
		// Anything left unhandled is an error, so this should be unreachable
		panic(fmt.Sprintf("Unhandled top-level node in declaration emit: %q", input.Kind))
	}

	tx.enclosingDeclaration = previousEnclosingDeclaration
	tx.state.getSymbolAccessibilityDiagnostic = oldDiag
	tx.needsDeclare = saveNeedsDeclare
	tx.state.errorNameNode = oldName
	return result
}

func (tx *DeclarationTransformer) transformTypeAliasDeclaration(input *ast.TypeAliasDeclaration) *ast.Node {
	tx.needsDeclare = false
	return tx.Factory().UpdateTypeAliasDeclaration(
		input,
		tx.ensureModifiers(input.AsNode()),
		input.Name(),
		tx.Visitor().VisitNodes(input.TypeParameters),
		tx.Visitor().Visit(input.Type),
	)
}

func (tx *DeclarationTransformer) transformInterfaceDeclaration(input *ast.InterfaceDeclaration) *ast.Node {
	return tx.Factory().UpdateInterfaceDeclaration(
		input,
		tx.ensureModifiers(input.AsNode()),
		input.Name(),
		tx.Visitor().VisitNodes(input.TypeParameters),
		tx.Visitor().VisitNodes(input.HeritageClauses),
		tx.Visitor().VisitNodes(input.Members),
	)
}

func (tx *DeclarationTransformer) transformFunctionDeclaration(input *ast.FunctionDeclaration) *ast.Node {
	// A dotted declaration (`function M.f() ... end`) contributes a member to M's
	// type rather than a standalone declaration; it has nothing to emit on its own.
	if input.Target != nil {
		return nil
	}
	if tx.resolver.IsExpandoFunctionDeclaration(input.AsNode()) {
		tx.state.reportExpandoFunctionErrors(input.AsNode())
	}
	return tx.Factory().UpdateFunctionDeclaration(
		input,
		tx.ensureModifiers(input.AsNode()),
		nil, /*target*/
		nil, /*colonToken*/
		input.Name(),
		tx.ensureTypeParams(input.AsNode(), input.TypeParameters),
		tx.updateParamList(input.AsNode(), input.Parameters),
		tx.ensureType(input.AsNode(), false),
		nil, /*fullSignature*/
		nil,
	)
}

func (tx *DeclarationTransformer) transformModuleDeclaration(input *ast.ModuleDeclaration) *ast.Node {
	// !!! TODO: module declarations are now parsed into nested module objects with export modifiers
	// It'd be good to collapse those back in the declaration output, but the AST can't represent the
	// `namespace a.b.c` shape for the printer (without using invalid identifier names).
	mods := tx.ensureModifiers(input.AsNode())
	saveNeedsDeclare := tx.needsDeclare
	tx.needsDeclare = false
	inner := input.Body
	keyword := input.Keyword
	if keyword != ast.KindGlobalKeyword && (input.Name() == nil || !ast.IsStringLiteral(input.Name())) {
		keyword = ast.KindNamespaceKeyword
	}

	if inner != nil && inner.Kind == ast.KindModuleBlock {
		oldNeedsScopeFix := tx.needsScopeFixMarker
		oldHasScopeFix := tx.resultHasScopeMarker
		tx.resultHasScopeMarker = false
		tx.needsScopeFixMarker = false
		statements := tx.Visitor().VisitNodes(inner.StatementList())
		lateStatements := tx.transformAndReplaceLatePaintedStatements(statements)
		if input.Flags&ast.NodeFlagsAmbient != 0 {
			tx.needsScopeFixMarker = false // If it was `declare`'d everything is implicitly exported already, ignore late printed "privates"
		}
		// With the final list of statements, there are 3 possibilities:
		// 1. There's an export assignment or export declaration in the namespace - do nothing
		// 2. Everything is exported and there are no export assignments or export declarations - strip all export modifiers
		// 3. Some things are exported, some are not, and there's no marker - add an empty marker
		if !ast.IsGlobalScopeAugmentation(input.AsNode()) && !tx.resultHasScopeMarker && !hasScopeMarker(lateStatements) {
			if tx.needsScopeFixMarker {
				lateStatements = tx.Factory().NewNodeList(append(lateStatements.Nodes, createEmptyExports(tx.Factory().AsNodeFactory())))
			} else {
				lateStatements = tx.exportStrippingVisitor.VisitNodes(lateStatements)
			}
		}

		body := tx.Factory().UpdateModuleBlock(inner.AsModuleBlock(), lateStatements)
		tx.needsDeclare = saveNeedsDeclare
		tx.needsScopeFixMarker = oldNeedsScopeFix
		tx.resultHasScopeMarker = oldHasScopeFix

		return tx.Factory().UpdateModuleDeclaration(
			input,
			mods,
			keyword,
			input.Name(),
			body,
		)
	}
	if inner != nil {
		// trigger visit. ignore result (is deferred, so is just inner unless elided)
		tx.Visitor().Visit(inner)
		// eagerly transform nested namespaces (the nesting doesn't need any elision or painting done)
		original := tx.EmitContext().MostOriginal(inner)
		id := ast.GetNodeId(original)
		body, _ := tx.lateStatementReplacementMap[id]
		delete(tx.lateStatementReplacementMap, id)
		return tx.Factory().UpdateModuleDeclaration(
			input,
			mods,
			keyword,
			input.Name(),
			body,
		)
	}
	return tx.Factory().UpdateModuleDeclaration(
		input,
		mods,
		keyword,
		input.Name(),
		nil,
	)
}

func (tx *DeclarationTransformer) stripExportModifiers(statement *ast.Node) *ast.Node {
	if statement == nil {
		return nil
	}
	parseNode := tx.EmitContext().ParseNode(statement)
	if ast.IsImportEqualsDeclaration(statement) || (parseNode != nil && tx.host.GetEffectiveDeclarationFlags(parseNode, ast.ModifierFlagsDefault) != 0) || !ast.CanHaveModifiers(statement) {
		// `export import` statements should remain as-is, as imports are _not_ implicitly exported in an ambient namespace
		// Likewise, `export default` classes and the like and just be `default`, so we preserve their `export` modifiers, too
		return statement
	}

	oldFlags := ast.GetCombinedModifierFlags(statement)
	if oldFlags&ast.ModifierFlagsExport == 0 {
		return statement
	}
	newFlags := oldFlags & (ast.ModifierFlagsAll ^ ast.ModifierFlagsExport)
	modifiers := ast.CreateModifiersFromModifierFlags(newFlags, tx.Factory().NewModifier)
	return ast.ReplaceModifiers(tx.Factory().AsNodeFactory(), statement, tx.Factory().NewModifierList(modifiers))
}

func (tx *DeclarationTransformer) transformVariableStatement(input *ast.VariableStatement) *ast.Node {
	visible := false
	for _, decl := range input.DeclarationList.AsVariableDeclarationList().Declarations.Nodes {
		visible = getBindingNameVisible(tx.resolver, decl)
		if visible {
			break
		}
	}
	if !visible {
		return nil
	}

	inputNodes := input.DeclarationList.AsVariableDeclarationList().Declarations.Nodes
	var extraImports []*ast.Node
	if tx.state.currentSourceFile.CommonJSModuleIndicator != nil {
		var normalDeclarations []*ast.Node
		var imports []*ast.Node
		for _, n := range inputNodes {
			if ast.IsVariableDeclarationInitializedToRequire(n) {
				imports = append(imports, n)
			} else {
				normalDeclarations = append(normalDeclarations, n)
			}
		}
		inputNodes = normalDeclarations
		extraImports, _ = tx.Visitor().VisitSlice(imports)
	}

	nodes, _ := tx.Visitor().VisitSlice(inputNodes)
	if len(nodes) == 0 {
		if len(extraImports) > 0 {
			return tx.Factory().NewSyntaxList(extraImports)
		}
		return nil
	}
	nodeList := tx.Factory().NewNodeList(nodes)

	modifiers := tx.ensureModifiers(input.AsNode())

	// Rebuild as a Const list regardless of the source flags: the emitted
	// statement is ambient (ensureModifiers adds `declare`), and a LuaLocal
	// list under a declare modifier has no parseable spelling.
	var declList *ast.Node
	if input.DeclarationList.Flags&(ast.NodeFlagsBlockScoped|ast.NodeFlagsLuaLocal) == ast.NodeFlagsConst {
		declList = tx.Factory().UpdateVariableDeclarationList(input.DeclarationList.AsVariableDeclarationList(), nodeList, input.DeclarationList.Flags)
	} else {
		declList = tx.Factory().NewVariableDeclarationList(nodeList, ast.NodeFlagsConst)
		tx.EmitContext().SetOriginal(declList, input.DeclarationList)
		tx.EmitContext().SetCommentRange(declList, input.DeclarationList.Loc)
		declList.Loc = input.DeclarationList.Loc
	}
	res := tx.Factory().UpdateVariableStatement(input, modifiers, declList)
	if len(extraImports) > 0 {
		return tx.Factory().NewSyntaxList(append(extraImports, res))
	}
	return res
}

func (tx *DeclarationTransformer) ensureModifiers(node *ast.Node) *ast.ModifierList {
	currentFlags := ast.GetCombinedModifierFlags(tx.EmitContext().ParseNode(node)) & ast.ModifierFlagsAll
	newFlags := tx.ensureModifierFlags(node)
	if currentFlags == newFlags {
		mods := node.Modifiers()
		if mods == nil {
			return mods
		}
		if canReuseModifierNodes(mods.Nodes) {
			return tx.Factory().NewModifierList(mods.Nodes)
		}
	}
	result := ast.CreateModifiersFromModifierFlags(newFlags, tx.Factory().NewModifier)
	if len(result) == 0 {
		return nil
	}
	return tx.Factory().NewModifierList(result)
}

func (tx *DeclarationTransformer) ensureModifierFlags(node *ast.Node) ast.ModifierFlags {
	mask := ast.ModifierFlagsAll ^ (ast.ModifierFlagsPublic | ast.ModifierFlagsAsync | ast.ModifierFlagsOverride) // No async and override modifiers in declaration files
	additions := ast.ModifierFlagsNone
	if tx.needsDeclare && !isAlwaysType(node) {
		additions = ast.ModifierFlagsAmbient
	}
	parentIsFile := node.Parent.Kind == ast.KindSourceFile
	if !parentIsFile {
		mask ^= ast.ModifierFlagsAmbient
		additions = ast.ModifierFlagsNone
	}
	if ast.IsImplicitlyExportedJSDocDeclaration(node) {
		additions |= ast.ModifierFlagsExport
	}
	return maskModifierFlags(node, mask, additions)
}

func (tx *DeclarationTransformer) ensureTypeParams(node *ast.Node, params *ast.TypeParameterList) *ast.TypeParameterList {
	if tx.host.GetEffectiveDeclarationFlags(tx.EmitContext().ParseNode(node), ast.ModifierFlagsPrivate) != 0 {
		return nil
	}
	var typeParameters *ast.TypeParameterList
	if typeParameters = tx.Visitor().VisitNodes(params); typeParameters != nil {
		return typeParameters
	}
	oldErrorNameNode := tx.state.errorNameNode
	tx.state.errorNameNode = node.Name()
	var oldDiag GetSymbolAccessibilityDiagnostic
	if !tx.suppressNewDiagnosticContexts {
		oldDiag = tx.state.getSymbolAccessibilityDiagnostic
		if canProduceDiagnostics(node) {
			tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(node)
		}
	}

	if data := node.FunctionLikeData(); data != nil && data.FullSignature != nil {
		if nodes := tx.resolver.CreateTypeParametersOfSignatureDeclaration(tx.EmitContext(), node, tx.enclosingDeclaration, declarationEmitNodeBuilderFlags, declarationEmitInternalNodeBuilderFlags, tx.tracker); nodes != nil {
			typeParameters = &ast.TypeParameterList{
				Loc:   node.Loc,
				Nodes: nodes,
			}
		}
	}

	tx.state.errorNameNode = oldErrorNameNode
	if !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = oldDiag
	}
	return typeParameters
}

func (tx *DeclarationTransformer) updateParamList(node *ast.Node, params *ast.ParameterList) *ast.ParameterList {
	if tx.host.GetEffectiveDeclarationFlags(tx.EmitContext().ParseNode(node), ast.ModifierFlagsPrivate) != 0 || len(params.Nodes) == 0 {
		return tx.Factory().NewNodeList([]*ast.Node{})
	}
	results := make([]*ast.Node, len(params.Nodes))
	for i, p := range params.Nodes {
		results[i] = tx.ensureParameter(p.AsParameterDeclaration())
	}
	return tx.Factory().NewNodeList(results)
}

func (tx *DeclarationTransformer) ensureParameter(p *ast.ParameterDeclaration) *ast.Node {
	oldDiag := tx.state.getSymbolAccessibilityDiagnostic
	if !tx.suppressNewDiagnosticContexts {
		tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(p.AsNode())
	}
	var questionToken *ast.TokenNode
	if tx.resolver.IsOptionalParameter(p.AsNode()) {
		if p.QuestionToken != nil {
			questionToken = p.QuestionToken
		} else {
			questionToken = tx.Factory().NewToken(ast.KindQuestionToken)
		}
	}
	result := tx.Factory().UpdateParameterDeclaration(
		p,
		nil,
		p.DotDotDotToken,
		tx.bindingNameVisitor.VisitNode(p.Name()),
		questionToken,
		tx.ensureType(p.AsNode(), true),
		tx.ensureNoInitializer(p.AsNode()),
	)
	tx.state.getSymbolAccessibilityDiagnostic = oldDiag
	return result
}

func (tx *DeclarationTransformer) ensureNoInitializer(node *ast.Node) *ast.Node {
	if tx.shouldPrintWithInitializer(node) && !ast.IsVariableDeclaration(node) {
		unwrappedInitializer := unwrapParenthesizedExpression(node.Initializer())
		if !ast.IsPrimitiveLiteralValue(unwrappedInitializer) {
			tx.tracker.ReportInferenceFallback(node)
		}
		return tx.resolver.CreateLiteralConstValue(tx.EmitContext(), tx.EmitContext().ParseNode(node), tx.tracker)
	}
	return nil
}

func (tx *DeclarationTransformer) visitBindingName(node *ast.Node) *ast.Node {
	switch node.Kind {
	case ast.KindIdentifier, ast.KindOmittedExpression:
		return node
	case ast.KindArrayBindingPattern, ast.KindObjectBindingPattern:
		return node.VisitEachChild(tx.bindingNameVisitor)
	case ast.KindBindingElement:
		if node.PropertyName() != nil && ast.IsComputedPropertyName(node.PropertyName()) && ast.IsEntityNameExpression(node.PropertyName().Expression()) {
			tx.checkEntityNameVisibility(node.PropertyName().Expression(), tx.enclosingDeclaration)
		}
		return tx.Factory().UpdateBindingElement(node.AsBindingElement(), node.AsBindingElement().DotDotDotToken, node.PropertyName(), tx.bindingNameVisitor.VisitNode(node.Name()), nil /*initializer*/)
	default:
		return node
	}
}

func (tx *DeclarationTransformer) transformImportEqualsDeclaration(decl *ast.ImportEqualsDeclaration) *ast.Node {
	if !tx.resolver.IsDeclarationVisible(decl.AsNode()) {
		return nil
	}
	if decl.ModuleReference.Kind == ast.KindExternalModuleReference {
		// Rewrite external module names if necessary
		specifier := ast.GetExternalModuleImportEqualsDeclarationExpression(decl.AsNode())
		return tx.Factory().UpdateImportEqualsDeclaration(
			decl,
			decl.Modifiers(),
			decl.IsTypeOnly,
			decl.Name(),
			tx.Factory().UpdateExternalModuleReference(decl.ModuleReference.AsExternalModuleReference(), tx.rewriteModuleSpecifier(decl.AsNode(), specifier)),
		)
	} else {
		oldDiag := tx.state.getSymbolAccessibilityDiagnostic
		tx.state.getSymbolAccessibilityDiagnostic = createGetSymbolAccessibilityDiagnosticForNode(decl.AsNode())
		tx.checkEntityNameVisibility(decl.ModuleReference, tx.enclosingDeclaration)
		tx.state.getSymbolAccessibilityDiagnostic = oldDiag
		return decl.AsNode()
	}
}

func (tx *DeclarationTransformer) transformImportDeclaration(decl *ast.ImportDeclaration) *ast.Node {
	if decl.ImportClause == nil {
		// import "mod" - possibly needed for side effects? (global interface patches, module augmentations, etc)
		return tx.Factory().UpdateImportDeclaration(
			decl,
			decl.Modifiers(),
			decl.ImportClause,
			tx.rewriteModuleSpecifier(decl.AsNode(), decl.ModuleSpecifier),
		)
	}
	phaseModifier := decl.ImportClause.AsImportClause().PhaseModifier
	if phaseModifier == ast.KindDeferKeyword {
		phaseModifier = ast.KindUnknown
	}
	// The `importClause` visibility corresponds to the default's visibility.
	var visibleDefaultBinding *ast.Node
	if decl.ImportClause != nil && decl.ImportClause.Name() != nil && tx.resolver.IsDeclarationVisible(decl.ImportClause) {
		visibleDefaultBinding = decl.ImportClause.Name()
	}
	if decl.ImportClause.AsImportClause().NamedBindings == nil {
		// No named bindings (either namespace or list), meaning the import is just default or should be elided
		if visibleDefaultBinding == nil {
			return nil
		}
		return tx.Factory().UpdateImportDeclaration(
			decl,
			decl.Modifiers(),
			tx.Factory().UpdateImportClause(
				decl.ImportClause.AsImportClause(),
				phaseModifier,
				visibleDefaultBinding,
				/*namedBindings*/ nil,
			),
			tx.rewriteModuleSpecifier(decl.AsNode(), decl.ModuleSpecifier),
		)
	}
	if decl.ImportClause.AsImportClause().NamedBindings.Kind == ast.KindNamespaceImport {
		// Namespace import (optionally with visible default)
		var namedBindings *ast.Node
		if tx.resolver.IsDeclarationVisible(decl.ImportClause.AsImportClause().NamedBindings) {
			namedBindings = decl.ImportClause.AsImportClause().NamedBindings
		}
		if visibleDefaultBinding == nil && namedBindings == nil {
			return nil
		}
		return tx.Factory().UpdateImportDeclaration(
			decl,
			decl.Modifiers(),
			tx.Factory().UpdateImportClause(
				decl.ImportClause.AsImportClause(),
				phaseModifier,
				visibleDefaultBinding,
				namedBindings,
			),
			tx.rewriteModuleSpecifier(decl.AsNode(), decl.ModuleSpecifier),
		)
	}
	// Named imports (optionally with visible default)
	bindingList := core.Filter(
		decl.ImportClause.AsImportClause().NamedBindings.Elements(),
		func(b *ast.Node) bool {
			return tx.resolver.IsDeclarationVisible(b)
		},
	)
	if len(bindingList) > 0 || visibleDefaultBinding != nil {
		var namedImports *ast.Node
		if len(bindingList) > 0 {
			namedImports = tx.Factory().UpdateNamedImports(
				decl.ImportClause.AsImportClause().NamedBindings.AsNamedImports(),
				tx.Factory().NewNodeList(bindingList),
			)
		}
		return tx.Factory().UpdateImportDeclaration(
			decl,
			decl.Modifiers(),
			tx.Factory().UpdateImportClause(
				decl.ImportClause.AsImportClause(),
				phaseModifier,
				visibleDefaultBinding,
				namedImports,
			),
			tx.rewriteModuleSpecifier(decl.AsNode(), decl.ModuleSpecifier),
		)
	}
	// Augmentation of export depends on import
	if tx.resolver.IsImportRequiredByAugmentation(decl) {
		if tx.state.isolatedDeclarations {
			tx.state.addDiagnostic(createDiagnosticForNode(decl.AsNode(), diagnostics.Declaration_emit_for_this_file_requires_preserving_this_import_for_augmentations_This_is_not_supported_with_isolatedDeclarations))
		}
		return tx.Factory().UpdateImportDeclaration(
			decl,
			decl.Modifiers(),
			/*importClause*/ nil,
			tx.rewriteModuleSpecifier(decl.AsNode(), decl.ModuleSpecifier),
		)
	}
	// Nothing visible
	return nil
}

func (tx *DeclarationTransformer) transformJSDocTypeExpression(input *ast.JSDocTypeExpression) *ast.Node {
	return tx.Visitor().Visit(input.Type)
}

func (tx *DeclarationTransformer) transformJSDocTypeLiteral(input *ast.JSDocTypeLiteral) *ast.Node {
	members, _ := tx.Visitor().VisitSlice(input.JSDocPropertyTags)
	replacement := tx.Factory().NewTypeLiteralNode(tx.Factory().NewNodeList(members))
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) transformJSDocPropertyTag(input *ast.JSDocParameterOrPropertyTag) *ast.Node {
	replacement := tx.Factory().NewPropertySignatureDeclaration(
		nil,
		tx.Visitor().Visit(input.TagName),
		nil,
		tx.Visitor().Visit(input.TypeExpression),
		nil,
	)
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) transformJSDocAllType(input *ast.JSDocAllType) *ast.Node {
	replacement := tx.Factory().NewKeywordTypeNode(ast.KindAnyKeyword)
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) transformJSDocNullableType(input *ast.JSDocNullableType) *ast.Node {
	replacement := tx.Factory().NewUnionTypeNode(tx.Factory().NewNodeList([]*ast.Node{
		tx.Visitor().Visit(input.Type),
		tx.Factory().NewLiteralTypeNode(tx.Factory().NewKeywordExpression(ast.KindNilKeyword)),
	}))
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) transformJSDocNonNullableType(input *ast.JSDocNonNullableType) *ast.Node {
	return tx.Visitor().Visit(input.Type)
}

func (tx *DeclarationTransformer) transformJSDocVariadicType(input *ast.JSDocVariadicType) *ast.Node {
	replacement := tx.Factory().NewArrayTypeNode(tx.Visitor().Visit(input.Type))
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) transformJSDocOptionalType(input *ast.JSDocOptionalType) *ast.Node {
	replacement := tx.Factory().NewUnionTypeNode(tx.Factory().NewNodeList([]*ast.Node{
		tx.Visitor().Visit(input.Type),
		tx.Factory().NewKeywordTypeNode(ast.KindNilKeyword),
	}))
	tx.EmitContext().SetOriginal(replacement, input.AsNode())
	return replacement
}

func (tx *DeclarationTransformer) getNameExpressionPreferringIdentifier(nameExpr *ast.Node) *ast.Node {
	if ast.IsNumericLiteral(nameExpr) {
		// Numeric property names are string properties in JS; convert to string literal
		nameExpr = tx.Factory().NewStringLiteral(nameExpr.Text(), ast.TokenFlagsNone)
	}
	if ast.IsStringLiteralLike(nameExpr) && scanner.IsIdentifierText(nameExpr.Text(), core.LanguageVariantStandard) {
		result := tx.Factory().NewIdentifier(nameExpr.Text()) // prefer non-string literal names where possible
		kwKind := scanner.IdentifierToKeywordKind(result.AsIdentifier())
		// keep keywords as strings, except `default`, which has special reformulations in the transformer
		if kwKind == ast.KindUnknown || kwKind == ast.KindDefaultKeyword {
			// fake this into a parse tree node so the reference resolver resolves the node via `resolveName`
			result.Parent = nameExpr.Parent
			result.Flags &^= ast.NodeFlagsSynthesized
			// intentionally leave Loc unset so the string isn't used as the text source of the identifier
			return result
		}
	}
	return nameExpr
}

func isNotDeclareModifier(mod *ast.Modifier) bool {
	return mod.Kind != ast.KindDeclareKeyword
}

func (tx *DeclarationTransformer) stripDeclareModifiers(node *ast.Node) *ast.Node {
	if node == nil {
		return nil
	}
	mods := node.Modifiers()
	if mods != nil {
		flags := node.ModifierFlags()
		if flags&ast.ModifierFlagsAmbient != 0 {
			filtered := core.Filter(mods.Nodes, isNotDeclareModifier)
			node.AsMutable().SetModifiers(tx.Factory().NewModifierList(filtered))
		}
	}
	return node // no need to recur into children, only strip at top-level
}

func (tx *DeclarationTransformer) visitCJSExportAssignments(expression *ast.Node) *ast.Node {
	if expression != nil {
		_, cleanupDiagnosticContext := tx.setupDiagnosticContext(expression)
		defer cleanupDiagnosticContext()
		switch ast.GetAssignmentDeclarationKind(expression) {
		case ast.JSDeclarationKindModuleExports:
			if tx.state.currentSourceFile.CommonJSModuleIndicator != nil {
				result := tx.transformExportAssignment(expression.Parent, expression, expression.AsBinaryExpression().Right, true /*isExportEquals*/)
				if result != nil {
					tx.cjsExportAssignment = result
					tx.resultHasScopeMarker = true
					tx.resultHasExternalModuleIndicator = true
				}
			}
		}
		return tx.cjsExportAssignmentVisitor.VisitEachChild(expression) // recur through the whole tree, looking for module.exports=
	}
	return nil
}

func (tx *DeclarationTransformer) visitNestedExpression(expression *ast.Node) *ast.Node {
	if expression != nil {
		_, cleanupDiagnosticContext := tx.setupDiagnosticContext(expression)
		defer cleanupDiagnosticContext()
		switch ast.GetAssignmentDeclarationKind(expression) {
		case ast.JSDeclarationKindProperty:
			tx.transformExpandoAssignment(expression.AsBinaryExpression())
		case ast.JSDeclarationKindExportsProperty:
			if tx.state.currentSourceFile.CommonJSModuleIndicator != nil {
				result := tx.transformCommonJSExport(expression, tx.getNameExpressionPreferringIdentifier(ast.GetElementOrPropertyAccessName(expression.AsBinaryExpression().Left)))
				if result != nil {
					tx.cjsExportMembers = append(tx.cjsExportMembers, result)
				}
			}
		case ast.JSDeclarationKindObjectDefinePropertyExports:
			if tx.state.currentSourceFile.CommonJSModuleIndicator != nil {
				result := tx.transformCommonJSExport(expression, tx.getNameExpressionPreferringIdentifier(expression.Arguments()[1]))
				if result != nil {
					tx.cjsExportMembers = append(tx.cjsExportMembers, result)
				}
			}
		}
		return tx.expressionVisitor.VisitEachChild(expression) // recur through the whole tree, looking for special assignments
	}
	return nil
}

func (tx *DeclarationTransformer) transformExpandoAssignment(node *ast.BinaryExpression) {
	left := node.Left

	symbol := node.Symbol
	// This reads the raw binder symbol. Lua augmentation candidates never carry
	// the Assignment flag there (it lives on the checker-local merged member),
	// so every Lua member write returns here. That is the intended emit for
	// now: augmented members reach the .d.ts through the printed type of their
	// base variable, and implicit-global serialization is deferred to the `_G`
	// task. Do not widen this gate to
	// symbol presence: the binder stamps a symbol on every Lua member write,
	// declaring or not.
	if symbol == nil || symbol.Flags&ast.SymbolFlagsAssignment == 0 {
		return
	}

	ns := ast.GetLeftmostAccessExpression(left)
	if ns == nil || ns.Kind != ast.KindIdentifier {
		return
	}

	declaration := tx.resolver.GetReferencedValueDeclaration(ns)
	if declaration == nil {
		return
	}

	if tx.shouldStripInternal(declaration) {
		return
	}

	if ast.IsVariableDeclaration(declaration) && declaration.Type() != nil {
		return
	}

	if ast.IsFunctionDeclaration(declaration) && declaration.FunctionLikeData().FullSignature != nil {
		return
	}

	if ast.IsVariableDeclaration(declaration) && !ast.IsFunctionLike(declaration.Initializer()) {
		return // We're going to add a type, no need to dupe members with a namespace
	}

	host := declaration.Symbol()
	if host == nil {
		return
	}

	name := tx.Factory().NewIdentifier(ns.Text())
	property := tx.tryGetPropertyName(left)
	// A word-operator name like `A["and"] = ...` is skipped: the exported alias
	// would print as the bare `export { C as and }`, which does not parse.
	if property == "" || !scanner.IsBareWritableName(property, core.LanguageVariantStandard) {
		return
	}

	hostId := tx.getExpandoHostId(declaration)

	if ast.IsDeclaration(declaration) && isDeclarationAndNotVisible(tx.EmitContext(), tx.resolver, declaration) {
		// The host isn't visible (yet) - printing the type of a visible declaration may still
		// late-mark it as visible (e.g. an exported variable whose type prints as `typeof host`),
		// so defer the assignment to be processed if and when that happens.
		tx.deferredExpandoAssignments[hostId] = append(tx.deferredExpandoAssignments[hostId], node)
		return
	}

	if ast.IsFunctionDeclaration(declaration) && !shouldEmitFunctionProperties(declaration.AsFunctionDeclaration()) {
		return
	}

	tx.transformExpandoHost(name, declaration)

	exportName := tx.Factory().NewIdentifier(property)
	localName := tx.tryGetNameOfAssignedExpression(node.AsNode())
	if localName == nil && !tx.resolver.IsNameResolvable(tx.enclosingDeclaration, property) && !ast.IsNonContextualKeyword(scanner.StringToToken(exportName.Text())) {
		// use exportName as localName if there won't be any conflicts or keyword issues
		localName = exportName
	}
	if localName == nil || ast.IsNonContextualKeyword(scanner.StringToToken(localName.Text())) {
		// fallback to a generated name if the localName doesn't exist or is a keyword
		localName = tx.Factory().NewGeneratedNameForNode(node.AsNode())
	}

	_, cleanupDiagnosticContext := tx.setupDiagnosticContext(node.AsNode())
	defer cleanupDiagnosticContext()

	if ast.IsIdentifier(node.Right) {
		// alias-like, emit an `export {name}` or `export {name as alias}`
		result := tx.transformBinaryExpressionToExportDeclaration(node.AsNode(), exportName)
		tx.expandoMembers[hostId] = append(tx.expandoMembers[hostId], result)
		return
	}

	preexistingExpandoHasExport := core.Some(tx.expandoMembers[hostId], ast.IsExportDeclaration)
	var varModifiers *ast.ModifierList

	if preexistingExpandoHasExport {
		varModifiers = tx.Factory().NewModifierList(ast.CreateModifiersFromModifierFlags(ast.ModifierFlagsExport, tx.Factory().NewModifier))
	}

	synthesizedNamespace := tx.Factory().NewModuleDeclaration(nil /*modifiers*/, ast.KindNamespaceKeyword, name, tx.Factory().NewModuleBlock(tx.Factory().NewNodeList([]*ast.Node{})))
	synthesizedNamespace.Parent = tx.enclosingDeclaration
	declarationData := synthesizedNamespace.DeclarationData()
	declarationData.Symbol = host
	containerData := synthesizedNamespace.LocalsContainerData()
	containerData.Locals = make(ast.SymbolTable, 0)
	containerData.Locals[localName.Text()] = symbol

	oldEnclosing := tx.enclosingDeclaration
	tx.enclosingDeclaration = synthesizedNamespace
	defer func() {
		tx.enclosingDeclaration = oldEnclosing
	}()

	statements := []*ast.Statement{
		tx.Factory().NewVariableStatement(
			varModifiers,
			tx.Factory().NewVariableDeclarationList(
				tx.Factory().NewNodeList([]*ast.Node{
					tx.Factory().NewVariableDeclaration(localName, nil /*exclamationToken*/, tx.ensureType(node.AsNode(), false), nil /*initializer*/),
				}),
				ast.NodeFlagsNone,
			),
		),
	}

	if localName.Text() != exportName.Text() {
		namedExports := tx.Factory().NewNamedExports(tx.Factory().NewNodeList(
			[]*ast.Node{
				tx.Factory().NewExportSpecifier(false /*isTypeOnly*/, localName, exportName),
			},
		))
		statements = append(statements, tx.Factory().NewExportDeclaration(nil /*modifiers*/, false /*isTypeOnly*/, namedExports, nil /*moduleSpecifier*/))
	}

	if len(statements) > 1 && !preexistingExpandoHasExport {
		// Add an `export` modifier to all existing expando members so they remain exported after the `export {}` is added
		for _, decl := range tx.expandoMembers[hostId] {
			modifierFlags := ast.ModifierFlagsExport | ast.GetCombinedModifierFlags(decl)
			decl.AsMutable().SetModifiers(tx.Factory().NewModifierList(ast.CreateModifiersFromModifierFlags(modifierFlags, tx.Factory().NewModifier)))
		}
	}
	tx.expandoMembers[hostId] = append(tx.expandoMembers[hostId], statements...)
}

func (tx *DeclarationTransformer) getExpandoHostId(declaration *ast.Declaration) ast.NodeId {
	root := core.IfElse(ast.IsVariableDeclaration(declaration), declaration.Parent.Parent, declaration)
	id := ast.GetNodeId(tx.EmitContext().MostOriginal(root))
	return id
}

func (tx *DeclarationTransformer) transformExpandoHost(name *ast.Node, declaration *ast.Declaration) {
	root := core.IfElse(ast.IsVariableDeclaration(declaration), declaration.Parent.Parent, declaration)
	id := tx.getExpandoHostId(declaration)

	if _, ok := tx.expandoHosts[id]; ok {
		return
	}

	saveNeedsDeclare := tx.needsDeclare
	tx.needsDeclare = true

	modifierFlags := tx.ensureModifierFlags(root)
	defaultExport := modifierFlags&ast.ModifierFlagsExport != 0 && modifierFlags&ast.ModifierFlagsDefault != 0

	tx.needsDeclare = saveNeedsDeclare

	if defaultExport {
		modifierFlags |= ast.ModifierFlagsAmbient
		modifierFlags ^= ast.ModifierFlagsDefault
		modifierFlags ^= ast.ModifierFlagsExport
	}

	_, cleanupDiagnosticContext := tx.setupDiagnosticContext(declaration)
	defer cleanupDiagnosticContext()

	modifiers := tx.Factory().NewModifierList(ast.CreateModifiersFromModifierFlags(modifierFlags, tx.Factory().NewModifier))
	replacement := make([]*ast.Node, 0)

	if ast.IsFunctionDeclaration(declaration) {
		typeParameters, parameters := extractExpandoHostParams(declaration)
		replacement = append(replacement, tx.Factory().UpdateFunctionDeclaration(declaration.AsFunctionDeclaration(), modifiers, nil /*target*/, nil /*colonToken*/, declaration.Name(), tx.ensureTypeParams(declaration, typeParameters), tx.updateParamList(declaration, parameters), tx.ensureType(declaration, false), nil /*fullSignature*/, nil /*body*/))
	} else if ast.IsVariableDeclaration(declaration) && ast.IsFunctionExpressionOrArrowFunction(declaration.Initializer()) {
		fn := declaration.Initializer()
		typeParameters, parameters := extractExpandoHostParams(fn)
		replacement = append(replacement, tx.Factory().NewFunctionDeclaration(modifiers, nil /*target*/, nil /*colonToken*/, tx.Factory().NewIdentifier(name.Text()), tx.ensureTypeParams(fn, typeParameters), tx.updateParamList(fn, parameters), tx.ensureType(fn, false), nil /*fullSignature*/, nil /*body*/))
	} else {
		tx.expandoHosts[id] = tx.transformTopLevelDeclaration(declaration)
		return
	}

	tx.state.reportExpandoFunctionErrors(declaration)

	if defaultExport {
		if ast.IsSourceFile(declaration.Parent) {
			tx.resultHasExternalModuleIndicator = true
		}
		tx.resultHasScopeMarker = true
		replacement = append(replacement, tx.Factory().NewExportAssignment(nil /*modifiers*/, false /*isExportEquals*/, nil /*typeNode*/, name))
	}

	// store host result to be added to the output when it's actually visited
	tx.expandoHosts[id] = tx.Factory().NewSyntaxList(replacement)
	if _, ok := tx.lateStatementReplacementMap[id]; ok {
		tx.lateStatementReplacementMap[id] = tx.createFullExpandoBlock(id)
	}
}

func (tx *DeclarationTransformer) createFullExpandoBlock(id ast.NodeId) *ast.Node {
	// Process any expando assignments on this host that were skipped because it wasn't
	// visible when they were collected - if it's still not visible, they simply get
	// re-deferred, and are dropped if the host is never late-marked visible.
	if deferred, ok := tx.deferredExpandoAssignments[id]; ok {
		delete(tx.deferredExpandoAssignments, id)
		for _, assignment := range deferred {
			tx.transformExpandoAssignment(assignment)
		}
	}
	n := tx.expandoHosts[id]
	if addOns, ok := tx.expandoMembers[id]; ok {
		var modifiers *ast.ModifierList
		var name *ast.Node
		var host []*ast.Node
		if n != nil && n.Kind == ast.KindSyntaxList {
			// find the first named syntax list element and use its' name & modifiers
			for c := range n.AsSyntaxList().IterChildren() {
				if c.Name() != nil {
					name = c.Name().Clone(tx.Factory())
					if c.Modifiers() != nil {
						modifiers = c.Modifiers().Clone(tx.Factory().AsNodeFactory())
					}
					break
				}
			}
			host = n.AsSyntaxList().Children
		} else if n != nil {
			name = n.Name().Clone(tx.Factory())
			if n.Modifiers() != nil {
				modifiers = n.Modifiers().Clone(tx.Factory().AsNodeFactory())
			}
			host = []*ast.Node{n}
		}
		if name != nil {
			moduleDecl := tx.Factory().NewModuleDeclaration(
				modifiers,
				ast.KindNamespaceKeyword,
				name,
				tx.Factory().NewModuleBlock(tx.Factory().NewNodeList(addOns)),
			)
			members := append(host, moduleDecl)
			return tx.Factory().NewSyntaxList(members)
		}
	}
	return n
}

func extractExpandoHostParams(node *ast.Node) (typeParameters *ast.TypeParameterList, parameters *ast.ParameterList) {
	switch node.Kind {
	case ast.KindFunctionExpression:
		fn := node.AsFunctionExpression()
		return fn.TypeParameters, fn.Parameters
	case ast.KindArrowFunction:
		fn := node.AsArrowFunction()
		return fn.TypeParameters, fn.Parameters
	default:
		fn := node.AsFunctionDeclaration()
		return fn.TypeParameters, fn.Parameters
	}
}

func (tx *DeclarationTransformer) tryGetPropertyName(node *ast.Node) string {
	if ast.IsElementAccessExpression(node) {
		return tx.resolver.GetElementAccessExpressionName(node.AsElementAccessExpression())
	}
	if ast.IsPropertyAccessExpression(node) {
		return node.Name().Text()
	}
	return ""
}

package ls

import (
	"context"
	"strings"
	"unicode"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/astnav"
	"github.com/apyrr/tlua/internal/checker"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/ls/lsconv"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/nodebuilder"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/scanner"
	"github.com/apyrr/tlua/internal/stringutil"
)

func (l *LanguageService) ProvideInlayHint(
	ctx context.Context,
	params *lsproto.InlayHintParams,
) (lsproto.InlayHintResponse, error) {
	userPreferences := l.UserPreferences()
	inlayHintPreferences := userPreferences.InlayHints
	if !isAnyInlayHintEnabled(inlayHintPreferences) {
		return lsproto.InlayHintsOrNull{InlayHints: nil}, nil
	}

	program, file := l.getProgramAndFile(params.TextDocument.Uri)
	quotePreference := lsutil.GetQuotePreference(file, userPreferences)

	checker, done := program.GetTypeCheckerForFile(ctx, file)
	defer done()
	inlayHintState := &inlayHintState{
		ctx:             ctx,
		span:            l.converters.FromLSPRange(file, params.Range),
		preferences:     inlayHintPreferences,
		quotePreference: quotePreference,
		file:            file,
		checker:         checker,
		converters:      l.converters,
	}
	inlayHintState.visit(file.AsNode())
	return lsproto.InlayHintsOrNull{InlayHints: &inlayHintState.result}, nil
}

type inlayHintState struct {
	ctx             context.Context
	span            core.TextRange
	preferences     lsutil.InlayHintsPreferences
	quotePreference lsutil.QuotePreference
	file            *ast.SourceFile
	checker         *checker.Checker
	converters      *lsconv.Converters
	result          []*lsproto.InlayHint
}

func (s *inlayHintState) visit(node *ast.Node) bool {
	if node == nil || node.End()-node.Pos() == 0 || node.Flags&ast.NodeFlagsReparsed != 0 {
		return false
	}

	switch node.Kind {
	case ast.KindModuleDeclaration, ast.KindInterfaceDeclaration,
		ast.KindFunctionDeclaration, ast.KindFunctionExpression, ast.KindArrowFunction:
		if s.ctx.Err() != nil {
			return true
		}
	}

	if !s.span.Intersects(node.Loc) {
		return false
	}

	if ast.IsTypeNode(node) && !ast.IsExpressionWithTypeArguments(node) {
		return false
	}

	if s.preferences.IncludeInlayVariableTypeHints.IsTrue() && ast.IsVariableDeclaration(node) {
		s.visitVariableLikeDeclaration(node)
	} else if shouldShowParameterNameHints(s.preferences) && ast.IsCallExpression(node) {
		s.visitCallExpression(node)
	} else {
		if s.preferences.IncludeInlayFunctionParameterTypeHints.IsTrue() &&
			ast.IsFunctionLikeDeclaration(node) &&
			ast.HasContextSensitiveParameters(node) {
			s.visitFunctionLikeForParameterType(node)
		}
		if s.preferences.IncludeInlayFunctionLikeReturnTypeHints.IsTrue() &&
			isSignatureSupportingReturnAnnotation(node) {
			s.visitFunctionDeclarationLikeForReturnType(node)
		}
	}
	return node.ForEachChild(s.visit)
}

// FunctionDeclaration | FunctionExpression | ArrowFunction
func (s *inlayHintState) visitFunctionDeclarationLikeForReturnType(decl *ast.FunctionLikeDeclaration) {
	if ast.IsArrowFunction(decl) {
		if astnav.FindChildOfKind(decl, ast.KindOpenParenToken, s.file) == nil {
			return
		}
	}

	typeAnnotation := decl.Type()
	if typeAnnotation != nil || decl.Body() == nil {
		return
	}

	signature := s.checker.GetSignatureFromDeclaration(decl)
	if signature == nil {
		return
	}

	typePredicate := s.checker.GetTypePredicateOfSignature(signature)

	if typePredicate != nil && typePredicate.Type() != nil {
		hintParts := s.typePredicateToInlayHintParts(typePredicate)
		s.addTypeHints(hintParts, s.getTypeAnnotationPosition(decl))
		return
	}

	returnType := s.checker.GetReturnTypeOfSignature(signature)
	if isModuleReferenceType(returnType) {
		return
	}

	hintParts := s.typeToInlayHintParts(returnType)
	s.addTypeHints(hintParts, s.getTypeAnnotationPosition(decl))
}

func (s *inlayHintState) visitCallExpression(expr *ast.Node) {
	args := expr.Arguments()
	if len(args) == 0 {
		return
	}

	signature := s.checker.GetResolvedSignature(expr)
	if signature == nil {
		return
	}

	// The written arguments start after the parameters the receiver fills.
	signatureParamPos := ast.LuaImplicitArgumentCount(expr)
	for _, originalArg := range args {
		arg := ast.SkipParentheses(originalArg)
		if shouldShowLiteralParameterNameHintsOnly(s.preferences) && !isHintableLiteral(arg) {
			signatureParamPos++
			continue
		}

		identifierInfo := s.getParameterIdentifierInfoAtPosition(signature, signatureParamPos)
		signatureParamPos++
		if identifierInfo == nil {
			return
		}

		parameter := identifierInfo.parameter
		parameterName := identifierInfo.name
		isFirstVariadicArgument := identifierInfo.isRestParameter
		parameterNameNotSameAsArgument := s.preferences.IncludeInlayParameterNameHintsWhenArgumentMatchesName.IsTrue() ||
			!identifierOrAccessExpressionPostfixMatchesParameterName(arg, parameterName)
		if !parameterNameNotSameAsArgument && !isFirstVariadicArgument {
			continue
		}

		if s.leadingCommentsContainsParameterName(arg, parameterName) {
			continue
		}

		s.addParameterHints(
			parameterName,
			parameter,
			astnav.GetStartOfNode(originalArg, s.file, false /*includeJSDoc*/),
			isFirstVariadicArgument,
		)
	}
}

func (s *inlayHintState) visitVariableLikeDeclaration(decl *ast.Node) {
	if decl.Initializer() == nil || ast.IsBindingPattern(decl.Name()) || !isHintableDeclaration(decl) {
		return
	}

	typeAnnotation := decl.Type()
	if typeAnnotation != nil {
		return
	}

	declarationType := s.checker.GetTypeAtLocation(decl)
	if isModuleReferenceType(declarationType) {
		return
	}

	hintParts := s.typeToInlayHintParts(declarationType)
	var hintText string
	if hintParts.String != nil {
		hintText = *hintParts.String
	} else if hintParts.InlayHintLabelParts != nil {
		var b strings.Builder
		for _, part := range *hintParts.InlayHintLabelParts {
			b.WriteString(part.Value)
		}
		hintText = b.String()
	}
	if !s.preferences.IncludeInlayVariableTypeHintsWhenTypeMatchesName.IsTrue() &&
		!ast.IsComputedPropertyName(decl.Name()) &&
		stringutil.EquateStringCaseInsensitive(decl.Name().Text(), hintText) {
		return
	}
	s.addTypeHints(hintParts, decl.Name().End())
}

func (s *inlayHintState) visitFunctionLikeForParameterType(node *ast.FunctionLikeDeclaration) {
	signature := s.checker.GetSignatureFromDeclaration(node)
	if signature == nil {
		return
	}

	pos := 0
	for _, param := range node.Parameters() {
		if isHintableDeclaration(param) {
			var symbol *ast.Symbol
			if ast.IsThisParameter(param) {
				symbol = signature.ThisParameter()
			} else {
				symbol = signature.Parameters()[pos]
			}
			s.addParameterTypeHint(param, symbol)
		}
		if ast.IsThisParameter(param) {
			continue
		}
		pos++
	}
}

func (s *inlayHintState) addParameterTypeHint(node *ast.ParameterDeclarationNode, symbol *ast.Symbol) {
	typeAnnotation := node.Type()
	if typeAnnotation != nil || symbol == nil {
		return
	}
	typeHints := s.getParameterDeclarationTypeHints(symbol)
	if typeHints == nil {
		return
	}
	var pos int
	if node.QuestionToken() != nil {
		pos = node.QuestionToken().End()
	} else {
		pos = node.Name().End()
	}
	s.addTypeHints(*typeHints, pos)
}

func (s *inlayHintState) getParameterDeclarationTypeHints(symbol *ast.Symbol) *lsproto.StringOrInlayHintLabelParts {
	valueDeclaration := symbol.ValueDeclaration
	if valueDeclaration == nil || !ast.IsParameterDeclaration(valueDeclaration) {
		return nil
	}

	signatureParamType := s.checker.GetTypeOfSymbolAtLocation(symbol, valueDeclaration)
	if isModuleReferenceType(signatureParamType) {
		return nil
	}

	return new(s.typeToInlayHintParts(signatureParamType))
}

func (s *inlayHintState) typeToInlayHintParts(t *checker.Type) lsproto.StringOrInlayHintLabelParts {
	return s.printInlayHintParts(func(nb *checker.NodeBuilder, flags nodebuilder.Flags) *ast.Node {
		return nb.TypeToTypeNode(t, nil /*enclosingDeclaration*/, flags, nodebuilder.InternalFlagsNone, nil /*tracker*/)
	})
}

func (s *inlayHintState) typePredicateToInlayHintParts(typePredicate *checker.TypePredicate) lsproto.StringOrInlayHintLabelParts {
	return s.printInlayHintParts(func(nb *checker.NodeBuilder, flags nodebuilder.Flags) *ast.Node {
		return nb.TypePredicateToTypePredicateNode(typePredicate, nil /*enclosingDeclaration*/, flags, nodebuilder.InternalFlagsNone, nil /*tracker*/)
	})
}

// printInlayHintParts prints the node build returns with the printer hover and signature help
// use, so a hint spells every type the way the rest of the language service does. The printer
// reports each name with its symbol, which becomes a part linking to the declaration.
func (s *inlayHintState) printInlayHintParts(build func(nb *checker.NodeBuilder, flags nodebuilder.Flags) *ast.Node) lsproto.StringOrInlayHintLabelParts {
	flags := nodebuilder.FlagsIgnoreErrors | nodebuilder.FlagsAllowUniqueESSymbolType |
		nodebuilder.FlagsUseAliasDefinedOutsideCurrentScope
	if s.quotePreference == lsutil.QuotePreferenceSingle {
		flags |= nodebuilder.FlagsUseSingleQuotesForStringLiteralType
	}
	// The printer reads emit flags (single-line object types) from the context the node
	// builder set them in, so the two share one.
	emitContext := printer.NewEmitContext()
	idToSymbol := make(map[*ast.IdentifierNode]*ast.Symbol)
	// !!! Avoid type node reuse so we collect identifier symbols.
	node := build(checker.NewNodeBuilderEx(s.checker, emitContext, idToSymbol), flags)
	debug.Assert(node != nil, "should always get a type node")
	p := printer.NewPrinter(printer.PrinterOptions{NewLine: core.NewLineKindLF}, printer.PrintHandlers{}, emitContext)
	p.IdToSymbol = idToSymbol
	writer := &inlayHintPartsWriter{state: s}
	p.Write(node, nil /*sourceFile*/, writer, nil /*sourceMapGenerator*/)
	return lsproto.StringOrInlayHintLabelParts{InlayHintLabelParts: new(writer.labelParts())}
}

func (s *inlayHintState) addTypeHints(hint lsproto.StringOrInlayHintLabelParts, position int) {
	if hint.String != nil {
		hint.String = new(": " + *hint.String)
	} else {
		hint.InlayHintLabelParts = new(append([]*lsproto.InlayHintLabelPart{{Value: ": "}}, *hint.InlayHintLabelParts...))
	}
	s.result = append(s.result, &lsproto.InlayHint{
		Label:       hint,
		Position:    s.converters.PositionToLineAndCharacter(s.file, core.TextPos(position)),
		Kind:        new(lsproto.InlayHintKindType),
		PaddingLeft: new(true),
	})
}

func (s *inlayHintState) addParameterHints(text string, parameter *ast.IdentifierNode, position int, isFirstVariadicArgument bool) {
	hintText := core.IfElse(isFirstVariadicArgument, "...", "") + text
	displayParts := []*lsproto.InlayHintLabelPart{
		s.getNodeDisplayPart(hintText, parameter),
		{
			Value: ":",
		},
	}
	labelParts := lsproto.StringOrInlayHintLabelParts{InlayHintLabelParts: &displayParts}

	s.result = append(s.result, &lsproto.InlayHint{
		Label:        labelParts,
		Position:     s.converters.PositionToLineAndCharacter(s.file, core.TextPos(position)),
		Kind:         new(lsproto.InlayHintKindParameter),
		PaddingRight: new(true),
	})
}

func shouldShowParameterNameHints(preferences lsutil.InlayHintsPreferences) bool {
	return (preferences.IncludeInlayParameterNameHints == lsutil.IncludeInlayParameterNameHintsLiterals ||
		preferences.IncludeInlayParameterNameHints == lsutil.IncludeInlayParameterNameHintsAll)
}

func shouldShowLiteralParameterNameHintsOnly(preferences lsutil.InlayHintsPreferences) bool {
	return preferences.IncludeInlayParameterNameHints == lsutil.IncludeInlayParameterNameHintsLiterals
}

// node is FunctionDeclaration | ArrowFunction | FunctionExpression.
func isSignatureSupportingReturnAnnotation(node *ast.Node) bool {
	return ast.IsArrowFunction(node) || ast.IsFunctionExpression(node) || ast.IsFunctionDeclaration(node)
}

func isHintableDeclaration(node *ast.VariableOrParameterDeclaration) bool {
	if (ast.IsPartOfParameterDeclaration(node) || ast.IsVariableDeclaration(node) && ast.IsVarConst(node)) &&
		node.Initializer() != nil {
		initializer := ast.SkipParentheses(node.Initializer())
		return !(isHintableLiteral(initializer) ||
			ast.IsObjectLiteralExpression(initializer) || ast.IsAssertionExpression(initializer))
	}
	return true
}

func isHintableLiteral(node *ast.Node) bool {
	switch node.Kind {
	case ast.KindPrefixUnaryExpression:
		return ast.IsLiteralExpression(node.AsPrefixUnaryExpression().Operand)
	case ast.KindTrueKeyword, ast.KindFalseKeyword, ast.KindNilKeyword,
		ast.KindNoSubstitutionTemplateLiteral, ast.KindTemplateExpression:
		return true
	}
	return ast.IsLiteralExpression(node)
}

func isModuleReferenceType(t *checker.Type) bool {
	symbol := t.Symbol()
	return symbol != nil && symbol.Flags&ast.SymbolFlagsModule != 0
}

func (s *inlayHintState) getNodeDisplayPart(text string, node *ast.Node) *lsproto.InlayHintLabelPart {
	file := ast.GetSourceFileOfNode(node)
	pos := astnav.GetStartOfNode(node, file, false /*includeJSDoc*/)
	end := node.End()
	return &lsproto.InlayHintLabelPart{
		Value: text,
		Location: &lsproto.Location{
			Uri:   lsconv.FileNameToDocumentURI(file.FileName()),
			Range: s.converters.ToLSPRange(file, core.NewTextRange(pos, end)),
		},
	}
}

type parameterInfo struct {
	parameter       *ast.IdentifierNode
	name            string
	isRestParameter bool
}

func (s *inlayHintState) getParameterIdentifierInfoAtPosition(signature *checker.Signature, pos int) *parameterInfo {
	parameters := signature.Parameters()
	paramCount := len(parameters) - core.IfElse(signature.HasRestParameter(), 1, 0)
	if pos < paramCount {
		param := parameters[pos]
		paramId := getParameterDeclarationIdentifier(param)
		if paramId == nil {
			return nil
		}
		return &parameterInfo{
			parameter:       paramId,
			name:            paramId.Text(),
			isRestParameter: false,
		}
	}

	var restParameter *ast.Symbol
	var restId *ast.IdentifierNode
	if paramCount < len(parameters) {
		restParameter = parameters[paramCount]
		restId = getParameterDeclarationIdentifier(restParameter)
	}
	if restId == nil {
		return nil
	}

	restType := s.checker.GetTypeOfSymbol(restParameter)
	if restType.IsTupleType() {
		associatedNames := make([]*ast.Node, 0, len(restType.Target().AsTupleType().ElementInfos()))
		for _, elementInfo := range restType.Target().AsTupleType().ElementInfos() {
			labeledElement := elementInfo.LabeledDeclaration()
			associatedNames = append(associatedNames, labeledElement)
		}
		index := pos - paramCount
		if index < len(associatedNames) {
			associatedName := associatedNames[index]
			if associatedName != nil {
				debug.Assert(ast.IsIdentifier(associatedName.Name()))
				var isRestTupleElement bool
				if ast.IsNamedTupleMember(associatedName) {
					isRestTupleElement = associatedName.AsNamedTupleMember().DotDotDotToken != nil
				} else {
					isRestTupleElement = associatedName.AsParameterDeclaration().DotDotDotToken != nil
				}
				return &parameterInfo{
					parameter:       associatedName.Name(),
					name:            associatedName.Name().Text(),
					isRestParameter: isRestTupleElement,
				}
			}
		}

		return nil
	}

	if pos == paramCount {
		if ast.IsVarargSymbolName(restParameter.Name) {
			// A vararg has no parameter name to hint with. Its synthetic `...` name
			// would render as `......:`, since the hint prefixes its own `...`.
			return nil
		}
		return &parameterInfo{
			parameter:       restId,
			name:            restParameter.Name,
			isRestParameter: true,
		}
	}
	return nil
}

func getParameterDeclarationIdentifier(symbol *ast.Symbol) *ast.IdentifierNode {
	if symbol.ValueDeclaration != nil && ast.IsParameterDeclaration(symbol.ValueDeclaration) && ast.IsIdentifier(symbol.ValueDeclaration.Name()) {
		return symbol.ValueDeclaration.Name()
	}
	return nil
}

func identifierOrAccessExpressionPostfixMatchesParameterName(expr *ast.Expression, parameterName string) bool {
	if ast.IsIdentifier(expr) {
		return expr.Text() == parameterName
	}
	if ast.IsPropertyAccessExpression(expr) {
		return expr.Name().Text() == parameterName
	}
	return false
}

func (s *inlayHintState) leadingCommentsContainsParameterName(node *ast.Node, name string) bool {
	if !scanner.IsIdentifierText(name) {
		return false
	}

	ranges := getLeadingCommentRangesOfNode(node, s.file)
	fileText := s.file.Text()
	for r := range ranges {
		commentText := strings.TrimFunc(fileText[r.Pos():r.End()], func(r rune) bool {
			return unicode.IsSpace(r) || r == '/' || r == '*'
		})
		if commentText == name {
			return true
		}
	}

	return false
}

func (s *inlayHintState) getTypeAnnotationPosition(decl *ast.FunctionLikeDeclaration) int {
	closeParenToken := astnav.FindChildOfKind(decl, ast.KindCloseParenToken, s.file)
	if closeParenToken != nil {
		return closeParenToken.End()
	}
	return decl.ParameterList().End()
}

func isAnyInlayHintEnabled(preferences lsutil.InlayHintsPreferences) bool {
	return preferences.IncludeInlayParameterNameHints != lsutil.IncludeInlayParameterNameHintsNone ||
		preferences.IncludeInlayFunctionParameterTypeHints.IsTrue() ||
		preferences.IncludeInlayVariableTypeHints.IsTrue() ||
		preferences.IncludeInlayPropertyDeclarationTypeHints.IsTrue() ||
		preferences.IncludeInlayFunctionLikeReturnTypeHints.IsTrue()
}

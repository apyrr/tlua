package parser

import (
	"slices"
	"strings"
	"sync"
	"unicode/utf8"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/debug"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/scanner"
	"github.com/apyrr/tlua/internal/stringutil"
	"github.com/apyrr/tlua/internal/tspath"
)

type ParsingContext int

const (
	PCSourceElements        ParsingContext = iota // Elements in source file
	PCBlockStatements                             // Statements in block
	PCLuaBlockStatements                          // Statements in an `end`-terminated Lua block
	PCLuaIfClauseStatements                       // Statements in a Lua if/elseif arm (ends at elseif/else/end)
	PCLuaRepeatStatements                         // Statements in a Lua repeat body (ends at until)
	PCTypeMembers                                 // Members in interface or type literal
	PCHeritageClauseElement                       // Elements in a heritage clause
	PCVariableDeclarations                        // Variable declarations in variable statement
	PCObjectBindingElements                       // Binding elements in object binding list
	PCArrayBindingElements                        // Binding elements in array binding list
	PCArgumentExpressions                         // Expressions in argument list
	PCObjectLiteralMembers                        // Members in object literal
	PCJsxAttributes                               // Attributes in jsx element
	PCJsxChildren                                 // Things between opening and closing JSX tags
	PCArrayLiteralMembers                         // Members in array literal
	PCParameters                                  // Parameters in parameter list
	PCJSDocParameters                             // JSDoc parameters in parameter list of JSDoc function type
	PCRestProperties                              // Property names in a rest type list
	PCTypeParameters                              // Type parameters in type parameter list
	PCTypeArguments                               // Type arguments in type argument list
	PCTupleElementTypes                           // Element types in tuple element type list
	PCHeritageClauses                             // Heritage clauses for a class or interface declaration.
	PCJSDocComment                                // Parsing via JSDocParser
	PCCount                                       // Number of parsing contexts
)

type ParsingContexts int

type JSDocInfo struct {
	parent *ast.Node
	jsDocs []*ast.Node
}

type jsdocScannerInfo uint8

const (
	jsdocScannerInfoHasJSDoc jsdocScannerInfo = 1 << iota
	jsdocScannerInfoHasDeprecated
	jsdocScannerInfoHasSeeOrLink
)

type Parser struct {
	scanner *scanner.Scanner
	factory ast.NodeFactory

	opts       ast.SourceFileParseOptions
	sourceText string

	scriptKind        core.ScriptKind
	isDeclarationFile bool
	languageVariant   core.LanguageVariant
	diagnostics       []*ast.Diagnostic
	jsDiagnostics     []*ast.Diagnostic
	jsdocDiagnostics  []*ast.Diagnostic

	token            ast.Kind
	sourceFlags      ast.NodeFlags
	contextFlags     ast.NodeFlags
	parsingContexts  ParsingContexts
	hasDeprecatedTag bool
	hasParseError    bool

	// packsInReturnUnion lets the OUTERMOST union of a return type recognize a
	// parenthesized pack as a constituent (`nil | (number, string)`), so a pack may
	// appear in any position, not only leading. parseUnionOrIntersectionType reads and
	// immediately clears it, so nested unions and intersections are unaffected.
	packsInReturnUnion bool

	identifiers                map[string]string
	identifierCount            int
	nodeSliceArena             core.Arena[*ast.Node]
	stringSliceArena           core.Arena[string]
	jsdocInfos                 []JSDocInfo
	jsdocCommentsSpace         []string
	jsdocCommentRangesSpace    []ast.CommentRange
	jsdocTagCommentsSpace      []string
	jsdocTagCommentsPartsSpace []*ast.Node
	reparseList                []*ast.Node
	commonJSModuleIndicator    *ast.Node

	currentParent        *ast.Node
	setParentFromContext ast.Visitor
	reparsedClones       []*ast.Node
}

func newParser() *Parser {
	res := &Parser{}
	res.initializeClosures()
	return res
}

var viableKeywordSuggestions = scanner.GetViableKeywordSuggestions()

var parserPool = sync.Pool{
	New: func() any {
		return newParser()
	},
}

func getParser() *Parser {
	return parserPool.Get().(*Parser)
}

func putParser(p *Parser) {
	*p = Parser{scanner: p.scanner, setParentFromContext: p.setParentFromContext}
	parserPool.Put(p)
}

func ParseSourceFile(opts ast.SourceFileParseOptions, sourceText string, scriptKind core.ScriptKind) *ast.SourceFile {
	p := getParser()
	defer putParser(p)
	p.initializeState(opts, sourceText, scriptKind)
	p.nextToken()
	if p.scriptKind == core.ScriptKindJSON {
		return p.parseJSONText()
	}
	return p.parseSourceFileWorker()
}

func (p *Parser) initializeClosures() {
	p.setParentFromContext = func(n *ast.Node) bool {
		n.Parent = p.currentParent
		return false
	}
}

func (p *Parser) isJavaScript() bool {
	return p.scriptKind == core.ScriptKindJS || p.scriptKind == core.ScriptKindJSX
}

func (p *Parser) parseJSONText() *ast.SourceFile {
	pos := p.nodePos()
	var statements *ast.NodeList
	var eof *ast.TokenNode

	if p.token == ast.KindEndOfFile {
		statements = p.newNodeList(core.NewTextRange(pos, p.nodePos()), nil)
		eof = p.parseTokenNode()
	} else {
		var expressions any // []*ast.Expression | *ast.Expression

		for p.token != ast.KindEndOfFile {
			var expression *ast.Expression
			switch p.token {
			case ast.KindOpenBracketToken:
				expression = p.parseArrayLiteralExpression()
			case ast.KindTrueKeyword, ast.KindFalseKeyword, ast.KindNilKeyword:
				expression = p.parseTokenNode()
			case ast.KindMinusToken:
				if p.lookAhead(func(p *Parser) bool {
					return p.nextToken() == ast.KindNumericLiteral && p.nextToken() != ast.KindColonToken
				}) {
					expression = p.parsePrefixUnaryExpression()
				} else {
					expression = p.parseObjectLiteralExpression()
				}
			case ast.KindNumericLiteral, ast.KindStringLiteral:
				if p.lookAhead(func(p *Parser) bool { return p.nextToken() != ast.KindColonToken }) {
					expression = p.parseLiteralExpression(false /*intern*/)
					break
				}
				fallthrough
			default:
				expression = p.parseObjectLiteralExpression()
			}

			// Error recovery: collect multiple top-level expressions
			if expressions != nil {
				if es, ok := expressions.([]*ast.Expression); ok {
					expressions = append(es, expression)
				} else {
					expressions = []*ast.Expression{expressions.(*ast.Expression), expression}
				}
			} else {
				expressions = expression
				if p.token != ast.KindEndOfFile {
					p.parseErrorAtCurrentToken(diagnostics.Unexpected_token)
				}
			}
		}

		var expression *ast.Expression
		if es, ok := expressions.([]*ast.Expression); ok {
			expression = p.finishNode(p.factory.NewArrayLiteralExpression(p.newNodeList(core.NewTextRange(pos, p.nodePos()), es), false), pos)
		} else {
			expression = expressions.(*ast.Expression)
		}
		statement := p.finishNode(p.factory.NewExpressionStatement(expression), pos)
		statements = p.newNodeList(core.NewTextRange(pos, p.nodePos()), []*ast.Node{statement})
		eof = p.parseExpectedToken(ast.KindEndOfFile)
	}
	node := p.finishNode(p.factory.NewSourceFile(p.opts, p.sourceText, statements, eof), pos)
	result := node.AsSourceFile()
	if len(result.Statements.Nodes) > 0 {
		p.validateJsonValue(result, result.Statements.Nodes[0].Expression())
	}
	p.finishSourceFile(result, false)
	return result
}

func getErrorSpanForNode(sourceText string, node *ast.Node) core.TextRange {
	pos := scanner.SkipTrivia(sourceText, node.Pos())
	return core.NewTextRange(pos, node.End())
}

func (p *Parser) validateJsonValue(sourceFile *ast.SourceFile, valueExpression *ast.Expression) {
	if valueExpression == nil {
		return
	}
	switch valueExpression.Kind {
	case ast.KindTrueKeyword, ast.KindFalseKeyword, ast.KindNilKeyword, ast.KindNumericLiteral:
		return
	case ast.KindStringLiteral:
		if !isDoubleQuotedString(valueExpression) {
			p.diagnostics = append(p.diagnostics, ast.NewDiagnostic(sourceFile, getErrorSpanForNode(p.sourceText, valueExpression), diagnostics.String_literal_with_double_quotes_expected))
		}
		return
	case ast.KindPrefixUnaryExpression:
		if valueExpression.AsPrefixUnaryExpression().Operator != ast.KindMinusToken || valueExpression.AsPrefixUnaryExpression().Operand.Kind != ast.KindNumericLiteral {
			break // not valid JSON syntax
		}
		return
	case ast.KindObjectLiteralExpression:
		p.validateJsonObjectLiteral(sourceFile, valueExpression.AsObjectLiteralExpression())
		return
	case ast.KindArrayLiteralExpression:
		for _, element := range valueExpression.Elements() {
			p.validateJsonValue(sourceFile, element)
		}
		return
	}
	p.diagnostics = append(p.diagnostics, ast.NewDiagnostic(sourceFile, getErrorSpanForNode(p.sourceText, valueExpression), diagnostics.Property_value_can_only_be_string_literal_numeric_literal_true_false_null_object_literal_or_array_literal))
}

func isDoubleQuotedString(node *ast.Node) bool {
	return ast.IsStringLiteral(node) && node.AsStringLiteral().TokenFlags&ast.TokenFlagsSingleQuote == 0
}

// validateJsonObjectLiteral validates properties of a JSON object literal.
func (p *Parser) validateJsonObjectLiteral(sourceFile *ast.SourceFile, node *ast.ObjectLiteralExpression) {
	for _, element := range node.Properties.Nodes {
		if element.Kind != ast.KindPropertyAssignment {
			p.diagnostics = append(p.diagnostics, ast.NewDiagnostic(sourceFile, getErrorSpanForNode(p.sourceText, element), diagnostics.Property_assignment_expected))
			continue
		}
		if element.Name() != nil && !isDoubleQuotedString(element.Name()) {
			p.diagnostics = append(p.diagnostics, ast.NewDiagnostic(sourceFile, getErrorSpanForNode(p.sourceText, element.Name()), diagnostics.String_literal_with_double_quotes_expected))
		}
		p.validateJsonValue(sourceFile, element.AsPropertyAssignment().Initializer)
	}
}

func ParseIsolatedEntityName(text string) *ast.EntityName {
	p := getParser()
	defer putParser(p)
	p.initializeState(ast.SourceFileParseOptions{}, text, core.ScriptKindJS)
	p.nextToken()
	entityName := p.parseEntityName(true, nil)
	return core.IfElse(p.token == ast.KindEndOfFile && len(p.diagnostics) == 0, entityName, nil)
}

func (p *Parser) initializeState(opts ast.SourceFileParseOptions, sourceText string, scriptKind core.ScriptKind) {
	if scriptKind == core.ScriptKindUnknown {
		panic("ScriptKind must be specified when parsing source file: " + opts.FileName)
	}

	if p.scanner == nil {
		p.scanner = scanner.NewScanner()
	} else {
		p.scanner.Reset()
	}
	p.opts = opts
	p.sourceText = sourceText
	p.scriptKind = scriptKind
	p.isDeclarationFile = tspath.IsDeclarationFileName(opts.FileName)
	p.languageVariant = getLanguageVariant(p.scriptKind)
	switch p.scriptKind {
	case core.ScriptKindJS, core.ScriptKindJSX:
		p.contextFlags = ast.NodeFlagsJavaScriptFile
	case core.ScriptKindJSON:
		p.contextFlags = ast.NodeFlagsJavaScriptFile | ast.NodeFlagsJsonFile
	default:
		p.contextFlags = ast.NodeFlagsNone
	}
	p.scanner.SetText(p.sourceText)
	p.scanner.SetOnError(p.scanError)
	p.scanner.SetLanguageVariant(p.languageVariant)
}

func (p *Parser) scanError(message *diagnostics.Message, pos int, length int, args ...any) {
	p.parseErrorAtRange(core.NewTextRange(pos, pos+length), message, args...)
}

func (p *Parser) parseErrorAt(pos int, end int, message *diagnostics.Message, args ...any) *ast.Diagnostic {
	return p.parseErrorAtRange(core.NewTextRange(pos, end), message, args...)
}

func (p *Parser) parseErrorAtCurrentToken(message *diagnostics.Message, args ...any) *ast.Diagnostic {
	return p.parseErrorAtRange(p.scanner.TokenRange(), message, args...)
}

func (p *Parser) parseErrorAtRange(loc core.TextRange, message *diagnostics.Message, args ...any) *ast.Diagnostic {
	// Don't report another error if it would just be at the same location as the last error
	var result *ast.Diagnostic
	if len(p.diagnostics) == 0 || p.diagnostics[len(p.diagnostics)-1].Pos() != loc.Pos() {
		result = ast.NewDiagnostic(nil, loc, message, args...)
		p.diagnostics = append(p.diagnostics, result)
	}
	p.hasParseError = true
	return result
}

type ParserState struct {
	scannerState       scanner.ScannerState
	contextFlags       ast.NodeFlags
	diagnosticsLen     int
	jsDiagnosticsLen   int
	jsdocInfosLen      int
	reparsedClonesLen  int
	hasParseError      bool
	packsInReturnUnion bool
}

func (p *Parser) mark() ParserState {
	return ParserState{
		scannerState:       p.scanner.Mark(),
		contextFlags:       p.contextFlags,
		diagnosticsLen:     len(p.diagnostics),
		jsDiagnosticsLen:   len(p.jsDiagnostics),
		jsdocInfosLen:      len(p.jsdocInfos),
		reparsedClonesLen:  len(p.reparsedClones),
		hasParseError:      p.hasParseError,
		packsInReturnUnion: p.packsInReturnUnion,
	}
}

func (p *Parser) rewind(state ParserState) {
	p.scanner.Rewind(state.scannerState)
	p.token = p.scanner.Token()
	p.contextFlags = state.contextFlags
	p.diagnostics = p.diagnostics[0:state.diagnosticsLen]
	p.jsDiagnostics = p.jsDiagnostics[0:state.jsDiagnosticsLen]
	p.jsdocInfos = p.jsdocInfos[0:state.jsdocInfosLen]
	p.reparsedClones = p.reparsedClones[0:state.reparsedClonesLen]
	p.hasParseError = state.hasParseError
	// Restored across a speculative parse so a lookAhead/tryParse that runs while the
	// flag is set (return-union pack recognition) cannot leak it past a rewind.
	p.packsInReturnUnion = state.packsInReturnUnion
}

func (p *Parser) lookAhead(callback func(p *Parser) bool) bool {
	state := p.mark()
	result := callback(p)
	p.rewind(state)
	return result
}

func (p *Parser) nextToken() ast.Kind {
	// if the keyword had an escape
	if ast.IsKeyword(p.token) && (p.scanner.HasUnicodeEscape() || p.scanner.HasExtendedUnicodeEscape()) {
		// issue a parse error for the escape
		p.parseErrorAtCurrentToken(diagnostics.Keywords_cannot_contain_escape_characters)
	}
	p.token = p.scanner.Scan()
	return p.token
}

func (p *Parser) nextTokenWithoutCheck() ast.Kind {
	p.token = p.scanner.Scan()
	return p.token
}

func (p *Parser) nextTokenJSDoc() ast.Kind {
	p.token = p.scanner.ScanJSDocToken()
	return p.token
}

func (p *Parser) nextJSDocCommentTextToken(inBackticks bool) ast.Kind {
	p.token = p.scanner.ScanJSDocCommentTextToken(inBackticks)
	return p.token
}

func (p *Parser) nodePos() int {
	return p.scanner.TokenFullStart()
}

func (p *Parser) hasPrecedingLineBreak() bool {
	return p.scanner.HasPrecedingLineBreak()
}

func (p *Parser) jsdocScannerInfo() jsdocScannerInfo {
	if !p.scanner.HasPrecedingJSDocComment() {
		return 0
	}
	info := jsdocScannerInfoHasJSDoc
	if p.scanner.HasPrecedingJSDocWithDeprecatedTag() {
		info |= jsdocScannerInfoHasDeprecated
	}
	if p.scanner.HasPrecedingJSDocWithSeeOrLink() {
		info |= jsdocScannerInfoHasSeeOrLink
	}
	return info
}

func (p *Parser) parseSourceFileWorker() *ast.SourceFile {
	isDeclarationFile := p.isDeclarationFile
	if isDeclarationFile {
		p.contextFlags |= ast.NodeFlagsAmbient
	}
	pos := p.nodePos()
	statements := p.parseListIndex(PCSourceElements, (*Parser).parseToplevelStatement)
	end := p.nodePos()
	endJSDoc := p.jsdocScannerInfo()
	eof := p.parseTokenNode()
	p.withJSDoc(eof, endJSDoc)
	if eof.Kind != ast.KindEndOfFile {
		panic("Expected end of file token from scanner.")
	}
	if len(p.reparseList) != 0 {
		statements = append(statements, p.reparseList...)
		p.reparseList = nil
	}
	node := p.finishNode(p.factory.NewSourceFile(p.opts, p.sourceText, p.newNodeList(core.NewTextRange(pos, end), statements), eof), pos)
	result := node.AsSourceFile()
	p.finishSourceFile(result, isDeclarationFile)
	collectExternalModuleReferences(result)
	if ast.IsInJSFile(node) {
		result.SetJSDiagnostics(attachFileToDiagnostics(p.jsDiagnostics, result))
	}
	return result
}

func (p *Parser) finishSourceFile(result *ast.SourceFile, isDeclarationFile bool) {
	result.CommentDirectives = p.scanner.CommentDirectives()
	result.Pragmas = getCommentPragmas(&p.factory, p.sourceText)
	p.processPragmasIntoFields(result)
	result.SetDiagnostics(attachFileToDiagnostics(p.diagnostics, result))
	result.SetJSDocDiagnostics(attachFileToDiagnostics(p.jsdocDiagnostics, result))
	result.CommonJSModuleIndicator = p.commonJSModuleIndicator
	result.IsDeclarationFile = isDeclarationFile
	result.LanguageVariant = p.languageVariant
	result.ScriptKind = p.scriptKind
	result.Flags |= p.sourceFlags
	result.Identifiers = p.identifiers
	result.NodeCount = p.factory.NodeCount()
	result.TextCount = p.factory.TextCount()
	result.IdentifierCount = p.identifierCount
	result.SetJSDocCache(p.createJSDocCache())
	// For non-JS files, enable lazy JSDoc parsing on demand
	if !p.isJavaScript() {
		result.SetHasLazyJSDoc(true)
	}
	slices.SortFunc(p.reparsedClones, ast.CompareNodePositions)
	result.ReparsedClones = slices.Clone(p.reparsedClones)
	ast.SetExternalModuleIndicator(result, p.opts.ExternalModuleIndicatorOptions)
}

func (p *Parser) createJSDocCache() map[*ast.Node][]*ast.Node {
	if len(p.jsdocInfos) == 0 {
		return nil
	}
	result := make(map[*ast.Node][]*ast.Node, len(p.jsdocInfos))
	for _, info := range p.jsdocInfos {
		result[info.parent] = info.jsDocs
	}
	return result
}

// parseToplevelStatement adapts parseStatement to the indexed element callback
// parseListIndex expects; the index is unused now that top-level-await
// reparsing is gone.
func (p *Parser) parseToplevelStatement(_ int) *ast.Node {
	return p.parseStatement()
}

func (p *Parser) parseListIndex(kind ParsingContext, parseElement func(p *Parser, index int) *ast.Node) []*ast.Node {
	saveParsingContexts := p.parsingContexts
	p.parsingContexts |= 1 << kind
	outerReparseList := p.reparseList
	p.reparseList = nil
	list := make([]*ast.Node, 0, 16)
	for i := 0; !p.isListTerminator(kind); i++ {
		if p.isListElement(kind, false /*inErrorRecovery*/) {
			elementPos := p.nodePos()
			elt := parseElement(p, len(list))
			// Error recovery must always make progress, even when a removed
			// construct is still recognized as the start of a list element.
			if p.nodePos() == elementPos && p.token != ast.KindEndOfFile {
				p.nextToken()
			}
			if len(p.reparseList) != 0 {
				for _, e := range p.reparseList {
					// Propagate @typedef type alias declarations outwards to a context that permits them.
					if (ast.IsJSTypeAliasDeclaration(e) || ast.IsJSImportDeclaration(e)) && kind != PCSourceElements && kind != PCBlockStatements {
						outerReparseList = append(outerReparseList, e)
					} else {
						list = append(list, e)
					}
				}
				p.reparseList = nil
			}
			list = append(list, elt)
			continue
		}
		if p.abortParsingListOrMoveToNextToken(kind) {
			break
		}
	}
	p.reparseList = outerReparseList
	p.parsingContexts = saveParsingContexts
	return p.nodeSliceArena.Clone(list)
}

func (p *Parser) parseList(kind ParsingContext, parseElement func(p *Parser) *ast.Node) *ast.NodeList {
	pos := p.nodePos()
	nodes := p.parseListIndex(kind, func(p *Parser, _ int) *ast.Node { return parseElement(p) })
	return p.newNodeList(core.NewTextRange(pos, p.nodePos()), nodes)
}

// Return a non-nil (but possibly empty) slice if parsing was successful, or nil if parseElement returned nil
func (p *Parser) parseDelimitedList(kind ParsingContext, parseElement func(p *Parser) *ast.Node) *ast.NodeList {
	pos := p.nodePos()
	saveParsingContexts := p.parsingContexts
	p.parsingContexts |= 1 << kind
	list := make([]*ast.Node, 0, 16)
	for {
		if p.isListElement(kind, false /*inErrorRecovery*/) {
			startPos := p.nodePos()
			element := parseElement(p)
			if element == nil {
				p.parsingContexts = saveParsingContexts
				// Return nil to indicate parseElement failed
				return nil
			}
			list = append(list, element)
			if p.parseOptional(ast.KindCommaToken) {
				// No need to check for a zero length node since we know we parsed a comma
				continue
			}
			// Lua allows `;` as a table field separator (including trailing).
			// tsconfig JSON parsing reuses this parsing context, so gate on
			// script kind to keep JSON strict.
			if kind == PCObjectLiteralMembers && p.scriptKind != core.ScriptKindJSON && p.parseOptional(ast.KindSemicolonToken) {
				continue
			}
			if p.isListTerminator(kind) {
				break
			}
			// We didn't get a comma, and the list wasn't terminated, explicitly parse
			// out a comma so we give a good error message.
			p.parseExpected(ast.KindCommaToken)
			// If the token was a semicolon, and the caller allows that, then skip it and
			// continue.  This ensures we get back on track and don't result in tons of
			// parse errors.  For example, this can happen when people do things like use
			// a semicolon to delimit object literal members.   Note: we'll have already
			// reported an error when we called parseExpected above.
			if kind == PCObjectLiteralMembers && p.token == ast.KindSemicolonToken && !p.hasPrecedingLineBreak() {
				p.nextToken()
			}
			if startPos == p.nodePos() {
				// What we're parsing isn't actually remotely recognizable as a element and we've consumed no tokens whatsoever
				// Consume a token to advance the parser in some way and avoid an infinite loop
				// This can happen when we're speculatively parsing parenthesized expressions which we think may be arrow functions,
				// or when a modifier keyword which is disallowed as a parameter name (ie, `static` in strict mode) is supplied
				p.nextToken()
			}
			continue
		}
		if p.isListTerminator(kind) {
			break
		}
		if p.abortParsingListOrMoveToNextToken(kind) {
			break
		}
	}
	p.parsingContexts = saveParsingContexts
	return p.newNodeList(core.NewTextRange(pos, p.nodePos()), p.nodeSliceArena.Clone(list))
}

// Return a non-nil (but possibly empty) NodeList. If the opening token is
// missing, return an empty list at the current position.
func (p *Parser) parseBracketedList(kind ParsingContext, parseElement func(p *Parser) *ast.Node, opening ast.Kind, closing ast.Kind) *ast.NodeList {
	if p.parseExpected(opening) {
		result := p.parseDelimitedList(kind, parseElement)
		p.parseExpected(closing)
		return result
	}
	return p.parseEmptyNodeList()
}

func (p *Parser) parseEmptyNodeList() *ast.NodeList {
	return p.newNodeList(core.NewTextRange(p.nodePos(), p.nodePos()), nil)
}

// Returns true if we should abort parsing.
func (p *Parser) abortParsingListOrMoveToNextToken(kind ParsingContext) bool {
	p.parsingContextErrors(kind)
	if p.isInSomeParsingContext() {
		return true
	}
	p.nextToken()
	return false
}

// True if positioned at element or terminator of the current list or any enclosing list
func (p *Parser) isInSomeParsingContext() bool {
	// We should be in at least one parsing context, be it SourceElements while parsing
	// a SourceFile, or JSDocComment when lazily parsing JSDoc.
	debug.Assert(p.parsingContexts != 0, "Missing parsing context")
	for kind := range PCCount {
		if p.parsingContexts&(1<<kind) != 0 {
			if p.isListElement(kind, true /*inErrorRecovery*/) || p.isListTerminator(kind) {
				return true
			}
		}
	}
	return false
}

func (p *Parser) parsingContextErrors(context ParsingContext) {
	switch context {
	case PCSourceElements:
		p.parseErrorAtCurrentToken(diagnostics.Declaration_or_statement_expected)
	case PCBlockStatements:
		p.parseErrorAtCurrentToken(diagnostics.Declaration_or_statement_expected)
	case PCLuaBlockStatements, PCLuaIfClauseStatements, PCLuaRepeatStatements:
		p.parseErrorAtCurrentToken(diagnostics.Declaration_or_statement_expected)
	case PCRestProperties, PCTypeMembers:
		p.parseErrorAtCurrentToken(diagnostics.Property_or_signature_expected)
	case PCHeritageClauseElement:
		p.parseErrorAtCurrentToken(diagnostics.Expression_expected)
	case PCVariableDeclarations:
		if ast.IsKeyword(p.token) {
			p.parseErrorAtCurrentToken(diagnostics.X_0_is_not_allowed_as_a_variable_declaration_name, scanner.TokenToString(p.token))
		} else {
			p.parseErrorAtCurrentToken(diagnostics.Variable_declaration_expected)
		}
	case PCObjectBindingElements:
		p.parseErrorAtCurrentToken(diagnostics.Property_destructuring_pattern_expected)
	case PCArrayBindingElements:
		p.parseErrorAtCurrentToken(diagnostics.Array_element_destructuring_pattern_expected)
	case PCArgumentExpressions:
		p.parseErrorAtCurrentToken(diagnostics.Argument_expression_expected)
	case PCObjectLiteralMembers:
		p.parseErrorAtCurrentToken(diagnostics.Property_assignment_expected)
	case PCArrayLiteralMembers:
		p.parseErrorAtCurrentToken(diagnostics.Expression_or_comma_expected)
	case PCJSDocParameters:
		p.parseErrorAtCurrentToken(diagnostics.Parameter_declaration_expected)
	case PCParameters:
		if ast.IsKeyword(p.token) {
			p.parseErrorAtCurrentToken(diagnostics.X_0_is_not_allowed_as_a_parameter_name, scanner.TokenToString(p.token))
		} else {
			p.parseErrorAtCurrentToken(diagnostics.Parameter_declaration_expected)
		}
	case PCTypeParameters:
		p.parseErrorAtCurrentToken(diagnostics.Type_parameter_declaration_expected)
	case PCTypeArguments:
		p.parseErrorAtCurrentToken(diagnostics.Type_argument_expected)
	case PCTupleElementTypes:
		p.parseErrorAtCurrentToken(diagnostics.Type_expected)
	case PCHeritageClauses:
		p.parseErrorAtCurrentToken(diagnostics.Unexpected_token_expected)
	case PCJsxAttributes, PCJsxChildren, PCJSDocComment:
		p.parseErrorAtCurrentToken(diagnostics.Identifier_expected)
	default:
		panic("Unhandled case in parsingContextErrors")
	}
}

func (p *Parser) isListElement(parsingContext ParsingContext, inErrorRecovery bool) bool {
	switch parsingContext {
	case PCSourceElements, PCBlockStatements, PCLuaBlockStatements, PCLuaIfClauseStatements, PCLuaRepeatStatements:
		// If we're in error recovery, then we don't want to treat ';' as an empty statement.
		// The problem is that ';' can show up in far too many contexts, and if we see one
		// and assume it's a statement, then we may bail out inappropriately from whatever
		// we're parsing.  For example, if we have a semicolon in the middle of a class, then
		// we really don't want to assume the class is over and we're on a statement in the
		// outer module.  We just want to consume and move on.
		return !(p.token == ast.KindSemicolonToken && inErrorRecovery) && p.isStartOfStatement()
	case PCTypeMembers:
		return p.lookAhead((*Parser).scanTypeMemberStart)
	case PCObjectLiteralMembers:
		switch p.token {
		// Generators are removed in tlua: `*` is not kept as a member start, so
		// list recovery consumes it with a natural error instead of looping.
		case ast.KindOpenBracketToken, ast.KindDotDotDotToken, ast.KindDotToken: // Not an object literal member, but don't want to close the object (see `tests/cases/fourslash/completionsDotInObjectLiteral.ts`)
			return true
		default:
			// Lua positional table entries admit any expression as a member,
			// on top of the keyed forms (whose names include non-expression
			// keywords, e.g. `{ default: 1 }`); JSON keeps the strict
			// property-name grammar.
			if p.scriptKind != core.ScriptKindJSON && p.isStartOfExpression() {
				return true
			}
			return p.isLiteralPropertyName()
		}
	case PCRestProperties:
		return p.isLiteralPropertyName()
	case PCObjectBindingElements:
		// No `...`: binding-pattern rests are gone, and parseObjectBindingElement can
		// no longer consume the token -- claiming it starts an element would spin the
		// list parser.
		return p.token == ast.KindOpenBracketToken || p.isLiteralPropertyName()
	case PCHeritageClauseElement:
		// If we see `{ ... }` then only consume it as an expression if it is followed by `,` or `{`
		// That way we won't consume the body of a class in its heritage clause.
		if p.token == ast.KindOpenBraceToken {
			return p.isValidHeritageClauseObjectLiteral()
		}
		if !inErrorRecovery {
			return p.isStartOfLeftHandSideExpression() && !p.isHeritageClauseExtendsOrImplementsKeyword()
		}
		// If we're in error recovery we tighten up what we're willing to match.
		// That way we don't treat something like "this" as a valid heritage clause
		// element during recovery.
		return p.isIdentifier() && !p.isHeritageClauseExtendsOrImplementsKeyword()
	case PCVariableDeclarations:
		return p.isBindingIdentifierOrPrivateIdentifierOrPattern()
	case PCArrayBindingElements:
		// Likewise: parseArrayBindingElement cannot consume a `...`.
		return p.token == ast.KindCommaToken || p.isBindingIdentifierOrPrivateIdentifierOrPattern()
	case PCTypeParameters:
		// A leading `...` starts a Lua generic pack parameter (`<...A>`).
		return p.token == ast.KindInKeyword || p.token == ast.KindDotDotDotToken || p.isIdentifier()
	case PCArrayLiteralMembers:
		// Not an array literal member, but don't want to close the array (see `tests/cases/fourslash/completionsDotInArrayLiteralInObjectLiteral.ts`)
		if p.token == ast.KindCommaToken || p.token == ast.KindDotToken {
			return true
		}
		fallthrough
	case PCArgumentExpressions:
		// `...` is covered by isStartOfExpression now that it is the vararg.
		return p.isStartOfExpression()
	case PCParameters:
		return p.isStartOfParameter(false /*isJSDocParameter*/)
	case PCJSDocParameters:
		return p.isStartOfParameter(true /*isJSDocParameter*/)
	case PCTypeArguments, PCTupleElementTypes:
		return p.token == ast.KindCommaToken || p.isStartOfType(false /*inStartOfParameter*/)
	case PCHeritageClauses:
		return p.isHeritageClause()
	case PCJsxAttributes:
		// `{` no longer starts an attribute: spread attributes are gone. It must not
		// be reported as an element start either, or the list would never advance
		// past it -- parseJsxAttribute cannot consume it, so parseList would spin.
		return p.tokenIsJsxName()
	case PCJsxChildren:
		return true
	case PCJSDocComment:
		return true
	}
	panic("Unhandled case in isListElement")
}

func (p *Parser) isListTerminator(kind ParsingContext) bool {
	if p.token == ast.KindEndOfFile {
		return true
	}
	switch kind {
	case PCBlockStatements, PCTypeMembers, PCObjectLiteralMembers,
		PCObjectBindingElements:
		return p.token == ast.KindCloseBraceToken
	case PCLuaBlockStatements:
		return p.token == ast.KindEndKeyword
	case PCLuaIfClauseStatements:
		return p.token == ast.KindEndKeyword || p.token == ast.KindElseIfKeyword || p.token == ast.KindElseKeyword
	case PCLuaRepeatStatements:
		return p.token == ast.KindUntilKeyword
	case PCHeritageClauseElement:
		return p.token == ast.KindOpenBraceToken || p.token == ast.KindExtendsKeyword || p.token == ast.KindImplementsKeyword
	case PCVariableDeclarations:
		// If we can consume a semicolon (either explicitly, or with ASI), then consider us done
		// with parsing the list of variable declarators.
		// In the case where we're parsing the variable declarator of a 'for-in' statement, we
		// are done if we see an 'in' keyword in front of us. Same with for-of
		// ERROR RECOVERY TWEAK:
		// For better error recovery, if we see an '=>' then we just stop immediately.  We've got an
		// arrow function here and it's going to be very unlikely that we'll resynchronize and get
		// another variable declaration.
		return p.canParseSemicolon() || p.token == ast.KindInKeyword || p.token == ast.KindOfKeyword || p.token == ast.KindEqualsGreaterThanToken
	case PCTypeParameters:
		// Tokens other than '>' are here for better error recovery
		return p.token == ast.KindGreaterThanToken || p.token == ast.KindOpenParenToken || p.token == ast.KindOpenBraceToken || p.token == ast.KindExtendsKeyword || p.token == ast.KindImplementsKeyword
	case PCArgumentExpressions:
		// Tokens other than ')' are here for better error recovery
		return p.token == ast.KindCloseParenToken || p.token == ast.KindSemicolonToken
	case PCArrayLiteralMembers, PCTupleElementTypes, PCArrayBindingElements:
		return p.token == ast.KindCloseBracketToken
	case PCJSDocParameters, PCParameters, PCRestProperties:
		// Tokens other than ')' and ']' (the latter for index signatures) are here for better error recovery
		return p.token == ast.KindCloseParenToken || p.token == ast.KindCloseBracketToken /*|| token == ast.KindOpenBraceToken*/
	case PCTypeArguments:
		// All other tokens should cause the type-argument to terminate except comma token
		return p.token != ast.KindCommaToken
	case PCHeritageClauses:
		return p.token == ast.KindOpenBraceToken || p.token == ast.KindCloseBraceToken
	case PCJsxAttributes:
		return p.token == ast.KindGreaterThanToken || p.token == ast.KindSlashToken
	case PCJsxChildren:
		return p.token == ast.KindLessThanToken && p.lookAhead((*Parser).nextTokenIsSlash)
	}
	return false
}

func (p *Parser) parseExpectedJSDoc(kind ast.Kind) bool {
	if p.token == kind {
		p.nextTokenJSDoc()
		return true
	}
	if !isKeywordOrPunctuation(kind) {
		panic("Invalid JSDoc kind: expected keyword or punctuation")
	}
	p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(kind))
	return false
}

func (p *Parser) parseExpectedMatchingBrackets(openKind ast.Kind, closeKind ast.Kind, openParsed bool, openPosition int) {
	if p.token == closeKind {
		p.nextToken()
		return
	}
	lastError := p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(closeKind))
	if !openParsed {
		return
	}
	if lastError != nil {
		related := ast.NewDiagnostic(nil, core.NewTextRange(openPosition, openPosition), diagnostics.The_parser_expected_to_find_a_1_to_match_the_0_token_here, scanner.TokenToString(openKind), scanner.TokenToString(closeKind))
		lastError.AddRelatedInfo(related)
	}
}

func (p *Parser) parseOptional(token ast.Kind) bool {
	if p.token == token {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) parseExpected(kind ast.Kind) bool {
	return p.parseExpectedWithDiagnostic(kind, nil, true)
}

func (p *Parser) parseExpectedWithoutAdvancing(kind ast.Kind) bool {
	return p.parseExpectedWithDiagnostic(kind, nil, false)
}

func (p *Parser) parseExpectedWithDiagnostic(kind ast.Kind, message *diagnostics.Message, shouldAdvance bool) bool {
	if p.token == kind {
		if shouldAdvance {
			p.nextToken()
		}
		return true
	}
	// Report specific message if provided with one.  Otherwise, report generic fallback message.
	if message != nil {
		p.parseErrorAtCurrentToken(message)
	} else {
		p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(kind))
	}
	return false
}

func (p *Parser) parseTokenNode() *ast.Node {
	pos := p.nodePos()
	kind := p.token
	p.nextToken()
	return p.finishNode(p.factory.NewToken(kind), pos)
}

func (p *Parser) parseExpectedToken(kind ast.Kind) *ast.Node {
	token := p.parseOptionalToken(kind)
	if token == nil {
		p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(kind))
		token = p.finishNode(p.factory.NewToken(kind), p.nodePos())
	}
	return token
}

func (p *Parser) parseOptionalToken(kind ast.Kind) *ast.Node {
	if p.token == kind {
		return p.parseTokenNode()
	}
	return nil
}

func (p *Parser) parseExpectedTokenJSDoc(kind ast.Kind) *ast.Node {
	optional := p.parseOptionalTokenJSDoc(kind)
	if optional == nil {
		if !isKeywordOrPunctuation(kind) {
			panic("expected keyword or punctuation")
		}
		p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(kind))
		optional = p.finishNode(p.factory.NewToken(kind), p.nodePos())
	}
	return optional
}

func (p *Parser) parseOptionalTokenJSDoc(kind ast.Kind) *ast.Node {
	if p.token == kind {
		return p.parseTokenNode()
	}
	return nil
}

func (p *Parser) parseStatement() *ast.Statement {
	switch p.token {
	case ast.KindSemicolonToken:
		return p.parseEmptyStatement()
	case ast.KindOpenBraceToken:
		return p.parseRemovedBlockStatement()
	case ast.KindLocalKeyword:
		return p.parseLocalStatement(p.nodePos(), p.jsdocScannerInfo())
	case ast.KindFunctionKeyword:
		return p.parseFunctionDeclaration(p.nodePos(), p.jsdocScannerInfo(), nil /*modifiers*/)
	case ast.KindIfKeyword:
		return p.parseIfStatement()
	case ast.KindDoKeyword:
		return p.parseDoBlock()
	case ast.KindRepeatKeyword:
		return p.parseRepeatStatement()
	case ast.KindWhileKeyword:
		return p.parseWhileStatement()
	case ast.KindForKeyword:
		return p.parseForStatement()
	case ast.KindContinueKeyword:
		return p.parseContinueStatement()
	case ast.KindBreakKeyword:
		return p.parseBreakStatement()
	case ast.KindColonColonToken:
		return p.parseLabelStatement()
	case ast.KindGotoKeyword:
		return p.parseGotoStatement()
	case ast.KindReturnKeyword:
		return p.parseReturnStatement()
	case ast.KindThrowKeyword:
		return p.parseThrowStatement()
	case ast.KindDebuggerKeyword:
		return p.parseDebuggerStatement()
	case ast.KindAsyncKeyword, ast.KindInterfaceKeyword, ast.KindTypeKeyword, ast.KindModuleKeyword, ast.KindNamespaceKeyword,
		ast.KindDeclareKeyword,
		ast.KindPrivateKeyword, ast.KindProtectedKeyword, ast.KindPublicKeyword, ast.KindAbstractKeyword, ast.KindAccessorKeyword,
		ast.KindStaticKeyword, ast.KindReadonlyKeyword, ast.KindGlobalKeyword:
		if p.isStartOfDeclaration() {
			return p.parseDeclaration()
		}
	}
	return p.parseExpressionStatement()
}

func (p *Parser) parseDeclaration() *ast.Statement {
	// `parseListElement` attempted to get the reused node at this position,
	// but the ambient context flag was not yet set, so the node appeared
	// not reusable in that context.
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	modifiers := p.parseModifiersEx(false /*stopOnStartOfClassStaticBlock*/)
	isAmbient := modifiers != nil && core.Some(modifiers.Nodes, isDeclareModifier)
	if isAmbient {
		// !!! incremental parsing
		// node := p.tryReuseAmbientDeclaration(pos)
		// if node {
		// 	return node
		// }
		for _, m := range modifiers.Nodes {
			m.Flags |= ast.NodeFlagsAmbient
		}
		saveContextFlags := p.contextFlags
		p.setContextFlags(ast.NodeFlagsAmbient, true)
		result := p.parseDeclarationWorker(pos, jsdoc, modifiers)
		p.contextFlags = saveContextFlags
		return result
	} else {
		return p.parseDeclarationWorker(pos, jsdoc, modifiers)
	}
}

func (p *Parser) parseDeclarationWorker(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Statement {
	switch p.token {
	case ast.KindFunctionKeyword:
		return p.parseFunctionDeclaration(pos, jsdoc, modifiers)
	case ast.KindInterfaceKeyword:
		return p.parseInterfaceDeclaration(pos, jsdoc, modifiers)
	case ast.KindTypeKeyword:
		return p.parseTypeAliasDeclaration(pos, jsdoc, modifiers)
	case ast.KindGlobalKeyword:
		return p.parseGlobalScopeAugmentation(pos, jsdoc, modifiers)
	}
	// The Luau-style ambient global form `declare name: Type;`. Placed after the
	// switch so every keyword-introduced declaration form keeps its meaning; any
	// remaining identifier-or-keyword name is accepted.
	if modifiers != nil && core.Some(modifiers.Nodes, isDeclareModifier) && isAmbientGlobalNameToken(p.token) {
		return p.parseAmbientGlobalDeclaration(pos, jsdoc, modifiers)
	}
	if modifiers != nil {
		// We reached this point because we encountered modifiers and assumed a declaration
		// would follow. For recovery and error reporting purposes, return an incomplete declaration.
		p.parseErrorAt(p.nodePos(), p.nodePos(), diagnostics.Declaration_expected)
		return p.finishNode(p.factory.NewMissingDeclaration(modifiers), pos)
	}
	panic("Unhandled case in parseDeclarationWorker")
}

// parseAmbientGlobalDeclaration parses the Luau-style ambient global form
// `declare name: Type;` as sugar for `declare const name: Type;`: it builds
// the same const VariableStatement, so the binder and checker treat the two
// forms identically (read-only global via NodeFlagsConst).
// isAmbientGlobalNameToken reports whether token can be the name of the
// Luau-style ambient global form `declare name: Type;`. Keywords that begin
// other declaration forms (or are modifiers) keep their existing meaning
// after `declare`; any other identifier-or-keyword token — including the
// reserved words `string` and `package`, which name Lua standard libraries —
// is usable as the global's name.
func isAmbientGlobalNameToken(token ast.Kind) bool {
	switch token {
	case ast.KindLocalKeyword,
		ast.KindFunctionKeyword, ast.KindClassKeyword, ast.KindInterfaceKeyword, ast.KindTypeKeyword, ast.KindDeferKeyword,
		ast.KindModuleKeyword, ast.KindNamespaceKeyword, ast.KindGlobalKeyword,
		ast.KindAbstractKeyword, ast.KindAccessorKeyword, ast.KindAsyncKeyword, ast.KindDeclareKeyword,
		ast.KindPrivateKeyword, ast.KindProtectedKeyword, ast.KindPublicKeyword, ast.KindReadonlyKeyword, ast.KindStaticKeyword:
		return false
	}
	return tokenIsIdentifierOrKeyword(token)
}

func (p *Parser) parseAmbientGlobalDeclaration(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Statement {
	declPos := p.nodePos()
	name := p.parseIdentifierName()
	p.parseExpected(ast.KindColonToken)
	typeNode := p.parseType()
	declaration := p.finishNode(p.factory.NewVariableDeclaration(name, nil /*exclamationToken*/, typeNode, nil /*initializer*/), declPos)
	list := p.newNodeList(core.NewTextRange(declaration.Pos(), declaration.End()), []*ast.Node{declaration})
	declarationList := p.finishNode(p.factory.NewVariableDeclarationList(list, ast.NodeFlagsConst), declPos)
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewVariableStatement(modifiers, declarationList), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func isDeclareModifier(modifier *ast.Node) bool {
	return modifier.Kind == ast.KindDeclareKeyword
}

func (p *Parser) parseBlock(diagnosticMessage *diagnostics.Message) *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	openBracePosition := p.scanner.TokenStart()
	openBraceParsed := p.parseExpectedWithDiagnostic(ast.KindOpenBraceToken, diagnosticMessage, true /*shouldAdvance*/)
	multiline := false
	if openBraceParsed {
		multiline = p.hasPrecedingLineBreak()
		statements := p.parseList(PCBlockStatements, (*Parser).parseStatement)
		p.parseExpectedMatchingBrackets(ast.KindOpenBraceToken, ast.KindCloseBraceToken, openBraceParsed, openBracePosition)
		result := p.finishNode(p.factory.NewBlock(statements, multiline), pos)
		p.withJSDoc(result, jsdoc)
		if p.token == ast.KindEqualsToken {
			p.parseErrorAtCurrentToken(diagnostics.Declaration_or_statement_expected_This_follows_a_block_of_statements_so_if_you_intended_to_write_a_destructuring_assignment_you_might_need_to_wrap_the_whole_assignment_in_parentheses)
			p.nextToken()
		}
		return result
	}
	result := p.finishNode(p.factory.NewBlock(p.parseEmptyNodeList(), multiline), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseEmptyStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindSemicolonToken)
	result := p.finishNode(p.factory.NewEmptyStatement(), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseIfStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindIfKeyword)
	// tlua has only the Lua form `if c then ... end`. A parenthesized condition
	// (`if (x) then`) is just a parenthesized expression; the C-style
	// `if (c) <stmt>` form is removed, so the condition always leads to `then`.
	return p.parseLuaIfRest(pos, jsdoc)
}

// parseRemovedBlockStatement reports and recovers from a `{` in statement
// position. Block statements are removed: a Lua block is `do ... end`, and a
// bare `{ ... }` is neither a statement nor (in statement position) a table
// literal. We report at the brace but still parse the balanced region as a
// Block, rather than a token-spanning leaf node that would crash
// language-service position queries. For a balanced `{ ... }` -- e.g. a stale
// C-style function body -- recovery resumes after `}`; an unbalanced `{` (no
// closing `}` before EOF) greedily absorbs the trailing statements into the
// Block, matching upstream block recovery.
func (p *Parser) parseRemovedBlockStatement() *ast.Node {
	p.parseErrorAtCurrentToken(diagnostics.Declaration_or_statement_expected)
	return p.parseBlock(nil)
}

// parseLuaIfRest parses `<cond> then <arm> [elseif ...|else ...] end` — the
// part after `if` or `elseif`.
func (p *Parser) parseLuaIfRest(pos int, jsdoc jsdocScannerInfo) *ast.Node {
	return p.parseLuaIfRestAfterCondition(pos, jsdoc, p.parseExpressionAllowIn())
}

func (p *Parser) parseLuaIfRestAfterCondition(pos int, jsdoc jsdocScannerInfo, expression *ast.Expression) *ast.Node {
	p.parseExpected(ast.KindThenKeyword)
	thenStatement := p.parseLuaClauseBlock(PCLuaIfClauseStatements)
	var elseStatement *ast.Statement
	switch p.token {
	case ast.KindElseIfKeyword:
		// An elseif arm is a nested IfStatement in the else slot — the
		// `else if` shape the printer chains. The single shared `end` is
		// consumed at the innermost recursion.
		elsePos := p.nodePos()
		elseJSDoc := p.jsdocScannerInfo()
		p.nextToken()
		elseStatement = p.parseLuaIfRest(elsePos, elseJSDoc)
	case ast.KindElseKeyword:
		p.nextToken()
		elseStatement = p.parseLuaClauseBlock(PCLuaBlockStatements)
		p.parseExpected(ast.KindEndKeyword)
	default:
		p.parseExpected(ast.KindEndKeyword)
	}
	result := p.finishNode(p.factory.NewIfStatement(expression, thenStatement, elseStatement), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

// parseRepeatStatement parses Lua `repeat ... until c`. Body statements live
// directly on the node (no nested Block): the until-condition is a child of
// the same locals container, so body locals resolve in it positionally.
func (p *Parser) parseRepeatStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindRepeatKeyword)
	statements := p.parseList(PCLuaRepeatStatements, (*Parser).parseStatement)
	p.parseExpected(ast.KindUntilKeyword)
	expression := p.parseExpressionAllowIn()
	p.parseOptional(ast.KindSemicolonToken) // do-while precedent
	result := p.finishNode(p.factory.NewRepeatStatement(statements, expression), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseDoBlock() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindDoKeyword)
	// tlua has only Lua `do ... end` standalone blocks; the TS `do { } while (c)`
	// form is removed, so the statements always run to `end`.
	block := p.finishLuaEndBlock(pos)
	p.withJSDoc(block, jsdoc)
	return block
}

func (p *Parser) parseWhileStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindWhileKeyword)
	// tlua has only the Lua form `while c do ... end`. A parenthesized condition
	// (`while (x) do`) is just a parenthesized expression; the C-style
	// `while (c) <stmt>` form is removed.
	return p.parseLuaWhileRest(pos, jsdoc, p.parseExpressionAllowIn())
}

func (p *Parser) parseLuaWhileRest(pos int, jsdoc jsdocScannerInfo, expression *ast.Expression) *ast.Node {
	p.parseExpected(ast.KindDoKeyword)
	statement := p.parseLuaClauseBlock(PCLuaBlockStatements)
	p.parseExpected(ast.KindEndKeyword)
	result := p.finishNode(p.factory.NewWhileStatement(expression, statement), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseForStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindForKeyword)
	// Lua has only two for-forms, both beginning with a control name:
	// `for i = e1, e2[, e3] do ... end` (numeric) and
	// `for k, v in explist do ... end` (generic). The JS C-style `for (;;)` and
	// `for (x in obj)` forms are deleted.
	return p.parseLuaForStatement(pos, jsdoc)
}

// parseLuaForStatement parses the Lua for-forms after the `for` keyword: one
// control name (numeric) or a comma-separated name list (generic), each with
// an optional type annotation.
func (p *Parser) parseLuaForStatement(pos int, jsdoc jsdocScannerInfo) *ast.Node {
	declPos := p.nodePos()
	name := p.parseBindingIdentifier()
	typeNode := p.parseTypeAnnotation()
	declaration := p.finishNode(p.factory.NewVariableDeclaration(name, nil /*exclamationToken*/, typeNode, nil /*initializer*/), declPos)
	p.checkJSSyntax(declaration)
	if p.token == ast.KindEqualsToken {
		return p.parseLuaNumericForRest(pos, jsdoc, declaration, declPos)
	}
	return p.parseLuaGenericForRest(pos, jsdoc, declaration, declPos)
}

// parseLuaNumericForRest parses `= from, to[, step] do ... end`. The control
// variable is block-scoped to the loop (NodeFlagsLet, like `for (let i ...)`)
// and takes no initializer — the bounds live in dedicated statement slots.
func (p *Parser) parseLuaNumericForRest(pos int, jsdoc jsdocScannerInfo, declaration *ast.Node, declPos int) *ast.Node {
	p.parseExpected(ast.KindEqualsToken)
	from := p.parseAssignmentExpressionOrHigher()
	p.parseExpected(ast.KindCommaToken)
	to := p.parseAssignmentExpressionOrHigher()
	var step *ast.Expression
	if p.parseOptional(ast.KindCommaToken) {
		step = p.parseAssignmentExpressionOrHigher()
	}
	if p.token == ast.KindCommaToken {
		p.parseErrorAtCurrentToken(diagnostics.A_Lua_numeric_for_statement_has_at_most_three_control_expressions)
		for p.parseOptional(ast.KindCommaToken) {
			p.parseAssignmentExpressionOrHigher() // consume extras for recovery
		}
	}
	declarations := p.newNodeList(core.NewTextRange(declPos, declaration.End()), []*ast.Node{declaration})
	initializer := p.finishNodeWithEnd(p.factory.NewVariableDeclarationList(declarations, ast.NodeFlagsLet), declPos, declaration.End())
	p.parseExpected(ast.KindDoKeyword)
	body := p.parseLuaClauseBlock(PCLuaBlockStatements)
	p.parseExpected(ast.KindEndKeyword)
	result := p.finishNode(p.factory.NewNumericForStatement(initializer, from, to, step, body), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

// parseLuaGenericForRest parses `[, name]* in explist do ... end`. The name
// list is block-scoped (Let) and LuaLocal-flagged — the flag is the
// generic-for discriminator on the reused ForOfStatement kind.
func (p *Parser) parseLuaGenericForRest(pos int, jsdoc jsdocScannerInfo, firstDeclaration *ast.Node, declPos int) *ast.Node {
	declarations := []*ast.Node{firstDeclaration}
	for p.parseOptional(ast.KindCommaToken) {
		namePos := p.nodePos()
		name := p.parseBindingIdentifier()
		typeNode := p.parseTypeAnnotation()
		declaration := p.finishNode(p.factory.NewVariableDeclaration(name, nil /*exclamationToken*/, typeNode, nil /*initializer*/), namePos)
		p.checkJSSyntax(declaration)
		declarations = append(declarations, declaration)
	}
	list := p.newNodeList(core.NewTextRange(declPos, declarations[len(declarations)-1].End()), declarations)
	initializer := p.finishNodeWithEnd(p.factory.NewVariableDeclarationList(list, ast.NodeFlagsLet|ast.NodeFlagsLuaLocal), declPos, declarations[len(declarations)-1].End())
	p.parseExpected(ast.KindInKeyword)
	expression := p.parseCommaValueList()
	p.parseExpected(ast.KindDoKeyword)
	body := p.parseLuaClauseBlock(PCLuaBlockStatements)
	p.parseExpected(ast.KindEndKeyword)
	result := p.finishNode(p.factory.NewForOfStatement(initializer, expression, body), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseBreakStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindBreakKeyword)
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewBreakStatement(), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseContinueStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindContinueKeyword)
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewContinueStatement(), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

// `::name::` declares a label, a goto target. Unlike upstream's labeled
// statement it wraps nothing: it is a standalone statement, as in Lua.
func (p *Parser) parseLabelStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindColonColonToken)
	label := p.parseIdentifier()
	p.parseExpected(ast.KindColonColonToken)
	result := p.finishNode(p.factory.NewLabelStatement(label), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

// `goto name` jumps to a visible label. Resolution and the jump-into-local-scope
// rule are checked in the checker; the binder wires the flow edge.
func (p *Parser) parseGotoStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindGotoKeyword)
	label := p.parseIdentifier()
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewGotoStatement(label), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseReturnStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindReturnKeyword)
	var expression *ast.Expression
	if !p.canParseSemicolon() {
		expression = p.parseReturnExpressionList()
	}
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewReturnStatement(expression), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseThrowStatement() *ast.Node {
	// ThrowStatement[Yield] :
	//      throw [no LineTerminator here]Expression[In, ?Yield];
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindThrowKeyword)
	// Because of automatic semicolon insertion, we need to report error if this
	// throw could be terminated with a semicolon.  Note: we can't call 'parseExpression'
	// directly as that might consume an expression on the following line.
	// Instead, we create a "missing" identifier, but don't report an error. The actual error
	// will be reported in the grammar walker.
	var expression *ast.Expression
	if !p.hasPrecedingLineBreak() {
		expression = p.parseExpressionAllowIn()
	} else {
		expression = p.createMissingIdentifier()
	}
	if !p.tryParseSemicolon() {
		p.parseErrorForMissingSemicolonAfter(expression)
	}
	result := p.finishNode(p.factory.NewThrowStatement(expression), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseDebuggerStatement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindDebuggerKeyword)
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewDebuggerStatement(), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

// Upstream also produced a labeled statement here, when the expression turned
// out to be a bare identifier followed by `:`. tlua labels are `::name::`, a
// statement of their own, so `name:` no longer starts anything.
func (p *Parser) parseExpressionStatement() *ast.Statement {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	hasParen := p.token == ast.KindOpenParenToken
	expression := p.parseExpression()

	if !p.tryParseSemicolon() {
		p.parseErrorForMissingSemicolonAfter(expression)
	}
	result := p.finishNode(p.factory.NewExpressionStatement(expression), pos)
	if hasParen {
		jsdoc &^= jsdocScannerInfoHasJSDoc
	}
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseLocalStatement(pos int, jsdoc jsdocScannerInfo) *ast.Node {
	if p.lookAhead(func(p *Parser) bool { return p.nextToken() == ast.KindFunctionKeyword }) {
		return p.parseLocalFunctionDeclaration(pos, jsdoc)
	}
	// `local type X = ...` / `local interface I {}` — a module/block-private type
	// declaration. The trailing identifier disambiguates it from a variable literally
	// named `type`/`interface` (`local type = 5`, which is a binding).
	if p.lookAhead(func(p *Parser) bool {
		next := p.nextToken()
		return (next == ast.KindTypeKeyword || next == ast.KindInterfaceKeyword) && p.nextTokenIsIdentifierOnSameLine()
	}) {
		return p.parseLocalTypeDeclaration(pos, jsdoc)
	}
	declarationList := p.parseLocalDeclarationList()
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewVariableStatement(nil /*modifiers*/, declarationList), pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

func (p *Parser) parseLocalFunctionDeclaration(pos int, jsdoc jsdocScannerInfo) *ast.Node {
	p.parseExpected(ast.KindLocalKeyword)
	p.parseExpected(ast.KindFunctionKeyword)
	name := p.parseBindingIdentifier()
	typeParameters := p.parseTypeParameters()
	parameters := p.parseParameters(ParseFlagsNone)
	returnType := p.parseReturnTypeList(ast.KindColonToken)
	body := p.parseLuaFunctionBlock()
	result := p.finishNode(p.factory.NewFunctionDeclaration(nil /*modifiers*/, nil /*target*/, nil /*colonToken*/, name, typeParameters, parameters, returnType, nil /*fullSignature*/, body), pos)
	result.Flags |= ast.NodeFlagsLuaLocal
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

// parseLocalTypeDeclaration parses `local type X = ...` / `local interface I {}`.
// The `local` prefix marks the declaration NodeFlagsLuaLocal so the checker keeps it
// out of the global type table; bare top-level `type`/`interface` are hoisted to global.
func (p *Parser) parseLocalTypeDeclaration(pos int, jsdoc jsdocScannerInfo) *ast.Node {
	p.parseExpected(ast.KindLocalKeyword)
	var result *ast.Node
	switch p.token {
	case ast.KindInterfaceKeyword:
		result = p.parseInterfaceDeclaration(pos, jsdoc, nil /*modifiers*/)
	case ast.KindTypeKeyword:
		result = p.parseTypeAliasDeclaration(pos, jsdoc, nil /*modifiers*/)
	default:
		// The sole caller's lookahead guarantees `type`/`interface` follows `local`.
		panic("parseLocalTypeDeclaration requires `local type` or `local interface`")
	}
	result.Flags |= ast.NodeFlagsLuaLocal
	return result
}

func (p *Parser) parseLocalDeclarationList() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindLocalKeyword)
	// `local a, b, c = e1, e2` — comma-separated names (each with an optional
	// type annotation), then one optional value list assigning all names
	// positionally. The value list hangs off the LAST declaration so sibling
	// node ranges stay ordered and non-overlapping.
	var declarations []*ast.Node
	for {
		declPos := p.nodePos()
		jsdoc := p.jsdocScannerInfo()
		name := p.parseBindingIdentifier()
		typeNode := p.parseTypeAnnotation()
		last := p.token != ast.KindCommaToken
		var initializer *ast.Expression
		if last {
			initializer = p.parseLocalValueList()
		}
		declaration := p.finishNode(p.factory.NewVariableDeclaration(name, nil /*exclamationToken*/, typeNode, initializer), declPos)
		p.withJSDoc(declaration, jsdoc)
		p.checkJSSyntax(declaration)
		declarations = append(declarations, declaration)
		if last {
			break
		}
		p.nextToken() // the comma
	}
	list := p.newNodeList(core.NewTextRange(declarations[0].Pos(), declarations[len(declarations)-1].End()), declarations)
	return p.finishNode(p.factory.NewVariableDeclarationList(list, ast.NodeFlagsLuaLocal), pos)
}

// parseLocalValueList parses the optional `= e1, e2, ...` value list of a
// local declaration statement.
func (p *Parser) parseLocalValueList() *ast.Expression {
	if !p.parseOptional(ast.KindEqualsToken) {
		return nil
	}
	return p.parseCommaValueList()
}

func (p *Parser) nextTokenIsIdentifier() bool {
	p.nextToken()
	return p.isIdentifier()
}

// nextTokensStartLuaColonCall scans past a `:` and reports whether a colon
// call follows: a method name and its argument list's `(` (or the `<` of
// explicit type arguments). Requiring the `(`/`<` keeps error recovery from
// committing on annotation-shaped text like `id(value: T)` and swallowing
// whatever follows into a bogus argument list.
func (p *Parser) nextTokensStartLuaColonCall() bool {
	if !tokenIsLuaMethodName(p.nextToken()) {
		return false
	}
	switch p.nextToken() {
	case ast.KindOpenParenToken, ast.KindLessThanToken:
		return true
	}
	return false
}

func (p *Parser) parseVariableDeclaration() *ast.Node {
	// The definite-assignment assertion (`let x!: T`) died with the TS
	// declaration forms; a local without an initializer is definitely nil.
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	name := p.parseIdentifierOrPatternWithDiagnostic(diagnostics.Private_identifiers_are_not_allowed_in_variable_declarations)
	typeNode := p.parseTypeAnnotation()
	var initializer *ast.Expression
	if p.token != ast.KindInKeyword {
		initializer = p.parseInitializer()
	}
	result := p.finishNode(p.factory.NewVariableDeclaration(name, nil /*exclamationToken*/, typeNode, initializer), pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

func (p *Parser) parseIdentifierOrPattern() *ast.Node {
	return p.parseIdentifierOrPatternWithDiagnostic(nil)
}

// isArrayBindingPatternStart reports whether the current `[` token begins an
// array binding pattern. Array binding patterns (`local [a, b] = x`) are
// removed from tlua source — Lua uses a name list (`local a, b = f()`) — but
// still parse in declaration files so ported .d.ts function types with
// non-vararg pattern parameters (`([k, v]: [string, number]) => ...`) keep
// their shape. (No bundled lib uses them; vararg patterns are TLUA100034
// even in declaration files.) Keeping the rule in one predicate ensures
// every binding-start lookahead and the parse entry point agree.
func (p *Parser) isArrayBindingPatternStart() bool {
	return p.token == ast.KindOpenBracketToken && p.isDeclarationFile
}

func (p *Parser) parseIdentifierOrPatternWithDiagnostic(privateIdentifierDiagnosticMessage *diagnostics.Message) *ast.Node {
	if p.isArrayBindingPatternStart() {
		return p.parseArrayBindingPattern()
	}
	if p.token == ast.KindOpenBraceToken {
		return p.parseObjectBindingPattern()
	}
	return p.parseBindingIdentifierWithDiagnostic(privateIdentifierDiagnosticMessage)
}

func (p *Parser) parseArrayBindingPattern() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindOpenBracketToken)
	saveContextFlags := p.contextFlags
	p.setContextFlags(ast.NodeFlagsDisallowInContext, false)
	elements := p.parseDelimitedList(PCArrayBindingElements, (*Parser).parseArrayBindingElement)
	p.contextFlags = saveContextFlags
	p.parseExpected(ast.KindCloseBracketToken)
	return p.finishNode(p.factory.NewBindingPattern(ast.KindArrayBindingPattern, elements), pos)
}

func (p *Parser) parseArrayBindingElement() *ast.Node {
	pos := p.nodePos()
	var dotDotDotToken *ast.Node
	var name *ast.Node
	var initializer *ast.Expression
	if p.token != ast.KindCommaToken {
		// These are all nil for a missing element. No rest element: tlua has no
		// binding-pattern rest, so a `...` here is left for the natural error.
		name = p.parseIdentifierOrPattern()
		initializer = p.parseInitializer()
	}
	return p.finishNode(p.factory.NewBindingElement(dotDotDotToken, nil /*propertyName*/, name, initializer), pos)
}

func (p *Parser) parseObjectBindingPattern() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindOpenBraceToken)
	saveContextFlags := p.contextFlags
	p.setContextFlags(ast.NodeFlagsDisallowInContext, false)
	elements := p.parseDelimitedList(PCObjectBindingElements, (*Parser).parseObjectBindingElement)
	p.contextFlags = saveContextFlags
	p.parseExpected(ast.KindCloseBraceToken)
	return p.finishNode(p.factory.NewBindingPattern(ast.KindObjectBindingPattern, elements), pos)
}

func (p *Parser) parseObjectBindingElement() *ast.Node {
	pos := p.nodePos()
	// No rest element: `{a, ...rest}` is gone, so `...` is left unconsumed and
	// reported by ordinary recovery.
	var dotDotDotToken *ast.Node
	tokenIsIdentifier := p.isBindingIdentifier()
	propertyName := p.parsePropertyName()
	var name *ast.Node
	if tokenIsIdentifier && p.token != ast.KindColonToken {
		name = propertyName
		propertyName = nil
	} else {
		p.parseExpected(ast.KindColonToken)
		name = p.parseIdentifierOrPattern()
	}
	initializer := p.parseInitializer()
	return p.finishNode(p.factory.NewBindingElement(dotDotDotToken, propertyName, name, initializer), pos)
}

func (p *Parser) parseInitializer() *ast.Expression {
	if p.parseOptional(ast.KindEqualsToken) {
		return p.parseAssignmentExpressionOrHigher()
	}
	return nil
}

func (p *Parser) parseTypeAnnotation() *ast.TypeNode {
	if p.parseOptional(ast.KindColonToken) {
		return p.parseType()
	}
	return nil
}

func (p *Parser) parseFunctionDeclaration(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	p.parseExpected(ast.KindFunctionKeyword)
	// Generators are removed in tlua: `function*` is not parsed. The `*` is left
	// unconsumed and reported as an unexpected token (like other removed syntax).
	var target *ast.Expression
	var colonToken *ast.TokenNode
	var selfParameter *ast.Node
	var name *ast.Node
	if modifiers == nil || modifiers.ModifierFlags&ast.ModifierFlagsDefault == 0 || p.isBindingIdentifier() {
		namePos := p.nodePos()
		name = p.parseBindingIdentifier()
		// `function a.b.f(...)` declares `f` as a member of the table `a.b`. Each dot
		// folds the segment parsed so far into the target chain; the final segment
		// stays behind as the declaration's own name.
		for p.token == ast.KindDotToken {
			if target == nil {
				target = name
			} else {
				target = p.finishNode(p.factory.NewPropertyAccessExpression(target, nil /*questionDotToken*/, nil /*colonToken*/, name, ast.NodeFlagsNone), namePos)
			}
			p.nextToken() // the dot
			name = p.parseIdentifierName()
		}
		// `function a.b:f(...)` is the same declaration with an implicit `self` first
		// parameter. A colon here is never a return type: those are parsed after the
		// parameter list. At most one colon segment is allowed, so `a:b:c` and `a:b.c`
		// fall out as the natural `'(' expected`.
		if colonToken = p.parseOptionalToken(ast.KindColonToken); colonToken != nil {
			if target == nil {
				target = name
			} else {
				target = p.finishNode(p.factory.NewPropertyAccessExpression(target, nil /*questionDotToken*/, nil /*colonToken*/, name, ast.NodeFlagsNone), namePos)
			}
			name = p.parseIdentifierName()
			selfParameter = p.newLuaSelfParameter(target, colonToken)
		}
	}
	typeParameters := p.parseTypeParameters()
	parameters := p.parseParameters(ParseFlagsNone)
	if selfParameter != nil {
		// The colon form is sugar for an explicit first parameter, so the synthesized
		// `self` joins the parsed parameters rather than living in a side channel.
		parameters = p.newNodeList(parameters.Loc, append([]*ast.Node{selfParameter}, parameters.Nodes...))
	}
	returnType := p.parseReturnTypeList(ast.KindColonToken)
	body := p.parseFunctionDeclarationBody()
	result := p.finishNode(p.factory.NewFunctionDeclaration(modifiers, target, colonToken, name, typeParameters, parameters, returnType, nil /*fullSignature*/, body), pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

// newLuaSelfParameter builds the parameter that `function M:f(...)` leaves implicit:
// an ordinary first parameter `self: typeof M`.
//
// The synthesized nodes need non-empty ranges: an empty one makes `ast.NodeIsMissing`
// report them as missing, and the checker then resolves the type query to `any` without
// saying why. The parameter and its name take the colon, the one character of source
// that stands for them. The type query mirrors the target's own range instead, so that
// any error it reports about the target — `Cannot find name`, say — lands on the target
// and collapses into the identical error reported there.
func (p *Parser) newLuaSelfParameter(target *ast.Expression, colonToken *ast.TokenNode) *ast.Node {
	loc := colonToken.Loc
	name := p.finishSelfNode(p.newIdentifier("self"), loc)
	typeNode := p.finishSelfNode(p.factory.NewTypeQueryNode(p.newLuaSelfEntityName(target), nil /*typeArguments*/), target.Loc)
	parameter := p.factory.NewParameterDeclaration(nil /*modifiers*/, nil /*dotDotDotToken*/, name, nil /*questionToken*/, typeNode, nil /*initializer*/)
	return p.finishSelfNode(parameter, loc)
}

// newLuaSelfEntityName mirrors the target chain as an entity name so that a type query
// can name it. parseFunctionDeclaration builds that chain out of identifiers alone, so
// only the two entity-name kinds occur here.
func (p *Parser) newLuaSelfEntityName(target *ast.Expression) *ast.EntityName {
	if ast.IsIdentifier(target) {
		return p.finishSelfNode(p.newIdentifier(target.Text()), target.Loc)
	}
	left := p.newLuaSelfEntityName(target.Expression())
	right := p.finishSelfNode(p.newIdentifier(target.Name().Text()), target.Name().Loc)
	return p.finishSelfNode(p.factory.NewQualifiedName(left, right), target.Loc)
}

func (p *Parser) finishSelfNode(node *ast.Node, loc core.TextRange) *ast.Node {
	return p.finishNodeWithEnd(node, loc.Pos(), loc.End())
}

func (p *Parser) parseFunctionDeclarationBody() *ast.Node {
	// Ambient declarations carry no body -- just a signature terminated by `;`.
	if p.contextFlags&ast.NodeFlagsAmbient != 0 {
		p.parseSemicolon()
		return nil
	}
	if p.token == ast.KindSemicolonToken {
		p.parseSemicolon()
		return nil
	}
	return p.parseLuaFunctionBlock()
}

func isAsyncModifier(modifier *ast.Node) bool {
	return modifier.Kind == ast.KindAsyncKeyword
}

func (p *Parser) parseHeritageClauses() *ast.NodeList {
	// ClassTail[Yield,Await] : (Modified) See 14.5
	//      ClassHeritage[?Yield,?Await]opt { ClassBody[?Yield,?Await]opt }
	if p.isHeritageClause() {
		return p.parseList(PCHeritageClauses, (*Parser).parseHeritageClause)
	}
	return nil
}

func (p *Parser) parseHeritageClause() *ast.Node {
	pos := p.nodePos()
	kind := p.token
	p.nextToken()
	types := p.parseDelimitedList(PCHeritageClauseElement, (*Parser).parseExpressionWithTypeArguments)
	return p.checkJSSyntax(p.finishNode(p.factory.NewHeritageClause(kind, types), pos))
}

func (p *Parser) parseExpressionWithTypeArguments() *ast.Node {
	pos := p.nodePos()
	expression := p.parseLeftHandSideExpressionOrHigher()
	if ast.IsExpressionWithTypeArguments(expression) {
		return expression
	}
	typeArguments := p.parseTypeArguments()
	return p.finishNode(p.factory.NewExpressionWithTypeArguments(expression, typeArguments), pos)
}

func (p *Parser) nextTokenIsOpenParen() bool {
	return p.nextToken() == ast.KindOpenParenToken
}

func (p *Parser) parseErrorForMissingSemicolonAfter(node *ast.Node) {
	// Tagged template literals are sometimes used in places where only simple strings are allowed, i.e.:
	//   module `M1` {
	//   ^^^^^^^^^^^ This block is parsed as a template literal like module`M1`.
	if node.Kind == ast.KindTaggedTemplateExpression {
		p.parseErrorAtRange(p.skipRangeTrivia(node.AsTaggedTemplateExpression().Template.Loc), diagnostics.Module_declaration_names_may_only_use_or_quoted_strings)
		return
	}
	// Otherwise, if this isn't a well-known keyword-like identifier, give the generic fallback message.
	var expressionText string
	if node.Kind == ast.KindIdentifier {
		expressionText = node.Text()
	}
	if expressionText == "" {
		p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(ast.KindSemicolonToken))
		return
	}
	pos := scanner.SkipTrivia(p.sourceText, node.Pos())
	// Some known keywords are likely signs of syntax being used improperly.
	// (const/let/var deliberately get no hint: they are ordinary identifiers
	// and the deleted declaration forms fall out as generic parse errors.)
	switch expressionText {
	case "declare":
		// If a declared node failed to parse, it would have emitted a diagnostic already.
		return
	case "interface":
		p.parseErrorForInvalidName(diagnostics.Interface_name_cannot_be_0, diagnostics.Interface_must_be_given_a_name, ast.KindOpenBraceToken)
		return
	case "is":
		p.parseErrorAt(pos, p.scanner.TokenStart(), diagnostics.A_type_predicate_is_only_allowed_in_return_type_position_for_functions_and_methods)
		return
	case "module", "namespace":
		p.parseErrorForInvalidName(diagnostics.Namespace_name_cannot_be_0, diagnostics.Namespace_must_be_given_a_name, ast.KindOpenBraceToken)
		return
	case "type":
		p.parseErrorForInvalidName(diagnostics.Type_alias_name_cannot_be_0, diagnostics.Type_alias_must_be_given_a_name, ast.KindEqualsToken)
		return
	}
	// The user alternatively might have misspelled or forgotten to add a space after a common keyword.
	suggestion := core.GetSpellingSuggestionForStrings(expressionText, slices.Values(viableKeywordSuggestions))
	if suggestion == "" {
		suggestion = getSpaceSuggestion(expressionText)
	}
	if suggestion != "" {
		p.parseErrorAt(pos, node.End(), diagnostics.Unknown_keyword_or_identifier_Did_you_mean_0, suggestion)
		return
	}
	// Unknown tokens are handled with their own errors in the scanner
	if p.token == ast.KindUnknown {
		return
	}
	// Otherwise, we know this some kind of unknown word, not just a missing expected semicolon.
	p.parseErrorAt(pos, node.End(), diagnostics.Unexpected_keyword_or_identifier)
}

func getSpaceSuggestion(expressionText string) string {
	for _, keyword := range viableKeywordSuggestions {
		if len(expressionText) > len(keyword)+2 && strings.HasPrefix(expressionText, keyword) {
			return keyword + " " + expressionText[len(keyword):]
		}
	}
	return ""
}

func (p *Parser) parseErrorForInvalidName(nameDiagnostic *diagnostics.Message, blankDiagnostic *diagnostics.Message, tokenIfBlankName ast.Kind) {
	if p.token == tokenIfBlankName {
		p.parseErrorAtCurrentToken(blankDiagnostic)
	} else {
		p.parseErrorAtCurrentToken(nameDiagnostic, p.scanner.TokenValue())
	}
}

func (p *Parser) parseInterfaceDeclaration(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	p.parseExpected(ast.KindInterfaceKeyword)
	name := p.parseIdentifier()
	typeParameters := p.parseTypeParameters()
	heritageClauses := p.parseHeritageClauses()
	members := p.parseObjectTypeMembers()
	result := p.finishNode(p.factory.NewInterfaceDeclaration(modifiers, name, typeParameters, heritageClauses, members), pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

func (p *Parser) parseTypeAliasDeclaration(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	p.parseExpected(ast.KindTypeKeyword)
	if p.hasPrecedingLineBreak() {
		p.parseErrorAtCurrentToken(diagnostics.Line_break_not_permitted_here)
	}
	name := p.parseIdentifier()
	typeParameters := p.parseTypeParameters()
	p.parseExpected(ast.KindEqualsToken)
	var typeNode *ast.TypeNode
	if p.token == ast.KindIntrinsicKeyword && p.lookAhead((*Parser).nextIsNotDot) {
		typeNode = p.parseKeywordTypeNode()
	} else {
		typeNode = p.parseType()
	}
	p.parseSemicolon()
	result := p.finishNode(p.factory.NewTypeAliasDeclaration(modifiers, name, typeParameters, typeNode), pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

func (p *Parser) nextIsNotDot() bool {
	return p.nextToken() != ast.KindDotToken
}

// parseGlobalScopeAugmentation parses `declare global { ... }`, the only
// surviving ModuleDeclaration form (string-named ambient modules are deleted;
// `require` is typed through overloads instead).
func (p *Parser) parseGlobalScopeAugmentation(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	// parse 'global' as name of global scope augmentation
	name := p.parseIdentifier()
	var body *ast.Node
	if p.token == ast.KindOpenBraceToken {
		body = p.parseModuleBlock()
	} else {
		p.parseSemicolon()
	}
	result := p.finishNode(p.factory.NewModuleDeclaration(modifiers, ast.KindGlobalKeyword, name, body), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseModuleBlock() *ast.Node {
	pos := p.nodePos()
	var statements *ast.NodeList
	if p.parseExpected(ast.KindOpenBraceToken) {
		statements = p.parseList(PCBlockStatements, (*Parser).parseStatement)
		p.parseExpected(ast.KindCloseBraceToken)
	} else {
		statements = p.parseEmptyNodeList()
	}
	return p.finishNode(p.factory.NewModuleBlock(statements), pos)
}

// TYPES

func (p *Parser) parseType() *ast.TypeNode {
	saveContextFlags := p.contextFlags
	p.setContextFlags(ast.NodeFlagsTypeExcludesFlags, false)
	// packsInReturnUnion applies only to the OUTERMOST union of a return type. Consume
	// it across this dispatch and re-arm it just for the union branch, so it cannot
	// leak past a function-type or conditional return into a nested union (a parameter
	// type, a type-parameter constraint) and wrongly recognize a pack there.
	packsInReturnUnion := p.packsInReturnUnion
	p.packsInReturnUnion = false
	var typeNode *ast.TypeNode
	if p.isStartOfFunctionTypeOrConstructorType() {
		typeNode = p.parseFunctionOrConstructorType()
	} else {
		pos := p.nodePos()
		p.packsInReturnUnion = packsInReturnUnion
		typeNode = p.parseUnionTypeOrHigher()
		if !p.inDisallowConditionalTypesContext() && !p.hasPrecedingLineBreak() && p.parseOptional(ast.KindExtendsKeyword) {
			// The type following 'extends' is not permitted to be another conditional type
			extendsType := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, true, (*Parser).parseType)
			p.parseExpected(ast.KindQuestionToken)
			trueType := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseType)
			p.parseExpected(ast.KindColonToken)
			falseType := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseType)
			conditionalType := p.factory.NewConditionalTypeNode(typeNode, extendsType, trueType, falseType)
			p.finishNode(conditionalType, pos)
			typeNode = conditionalType
		}
	}
	p.contextFlags = saveContextFlags
	return typeNode
}

func (p *Parser) parseUnionTypeOrHigher() *ast.TypeNode {
	return p.parseUnionOrIntersectionType(ast.KindBarToken, (*Parser).parseIntersectionTypeOrHigher)
}

func (p *Parser) parseIntersectionTypeOrHigher() *ast.TypeNode {
	return p.parseUnionOrIntersectionType(ast.KindAmpersandToken, (*Parser).parseTypeOperatorOrHigher)
}

func (p *Parser) parseUnionOrIntersectionType(operator ast.Kind, parseConstituentType func(p *Parser) *ast.TypeNode) *ast.TypeNode {
	pos := p.nodePos()
	isUnionType := operator == ast.KindBarToken
	// A return type's outermost union recognizes a parenthesized pack as a constituent
	// (`nil | (number, string)`). Consume the flag here so nested unions and every
	// intersection see it cleared -- a pack is only valid at the top of the return type.
	packsInUnion := isUnionType && p.packsInReturnUnion
	p.packsInReturnUnion = false
	parseConstituent := func(p *Parser) *ast.TypeNode {
		if packsInUnion && p.isParenthesizedReturnTypeList() {
			return p.parseParenthesizedReturnTypeList()
		}
		return p.parseFunctionOrConstructorTypeToError(isUnionType, parseConstituentType)
	}
	hasLeadingOperator := p.parseOptional(operator)
	var typeNode *ast.TypeNode
	if hasLeadingOperator {
		typeNode = parseConstituent(p)
	} else {
		typeNode = parseConstituentType(p)
	}
	if p.token == operator || hasLeadingOperator {
		types := make([]*ast.Node, 1, 8)
		types[0] = typeNode
		for p.parseOptional(operator) {
			types = append(types, parseConstituent(p))
		}
		typeNode = p.createUnionOrIntersectionTypeNode(operator, p.newNodeList(core.NewTextRange(pos, p.nodePos()), p.nodeSliceArena.Clone(types)))
		p.finishNode(typeNode, pos)
	}
	return typeNode
}

func (p *Parser) createUnionOrIntersectionTypeNode(operator ast.Kind, types *ast.NodeList) *ast.Node {
	switch operator {
	case ast.KindBarToken:
		return p.factory.NewUnionTypeNode(types)
	case ast.KindAmpersandToken:
		return p.factory.NewIntersectionTypeNode(types)
	default:
		panic("Unhandled case in createUnionOrIntersectionType")
	}
}

func (p *Parser) parseTypeOperatorOrHigher() *ast.TypeNode {
	operator := p.token
	switch operator {
	case ast.KindKeyOfKeyword, ast.KindUniqueKeyword, ast.KindReadonlyKeyword:
		return p.parseTypeOperator(operator)
	case ast.KindInferKeyword:
		return p.parseInferType()
	}
	return doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parsePostfixTypeOrHigher)
}

func (p *Parser) parseTypeOperator(operator ast.Kind) *ast.Node {
	pos := p.nodePos()
	p.parseExpected(operator)
	return p.finishNode(p.factory.NewTypeOperatorNode(operator, p.parseTypeOperatorOrHigher()), pos)
}

func (p *Parser) parseInferType() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindInferKeyword)
	return p.finishNode(p.factory.NewInferTypeNode(p.parseTypeParameterOfInferType()), pos)
}

func (p *Parser) parseTypeParameterOfInferType() *ast.Node {
	pos := p.nodePos()
	name := p.parseIdentifier()
	constraint := p.tryParseConstraintOfInferType()
	return p.finishNode(p.factory.NewTypeParameterDeclaration(nil /*modifiers*/, nil /*dotDotDotToken*/, name, constraint, nil /*expression*/, nil /*defaultType*/), pos)
}

func (p *Parser) tryParseConstraintOfInferType() *ast.Node {
	state := p.mark()
	if p.parseOptional(ast.KindExtendsKeyword) {
		constraint := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, true, (*Parser).parseType)
		if p.inDisallowConditionalTypesContext() || p.token != ast.KindQuestionToken {
			return constraint
		}
	}
	p.rewind(state)
	return nil
}

func (p *Parser) parsePostfixTypeOrHigher() *ast.Node {
	pos := p.nodePos()
	typeNode := p.parseNonArrayType()
	for !p.hasPrecedingLineBreak() {
		switch p.token {
		case ast.KindExclamationToken:
			// Only the punctuation `!` is the JSDoc non-nullable postfix; the word
			// `not` in type position is left for stock recovery to reject.
			if !p.tokenIsExclamationPunctuation() {
				return typeNode
			}
			p.nextToken()
			typeNode = p.finishNode(p.factory.NewJSDocNonNullableType(typeNode), pos)
		case ast.KindQuestionToken:
			// If next token is start of a type we have a conditional type
			if p.lookAhead((*Parser).nextIsStartOfType) {
				return typeNode
			}
			p.nextToken()
			typeNode = p.finishNode(p.factory.NewJSDocNullableType(typeNode), pos)
		case ast.KindOpenBracketToken:
			p.parseExpected(ast.KindOpenBracketToken)
			if p.isStartOfType(false /*isStartOfParameter*/) {
				indexType := p.parseType()
				p.parseExpected(ast.KindCloseBracketToken)
				typeNode = p.finishNode(p.factory.NewIndexedAccessTypeNode(typeNode, indexType), pos)
			} else {
				p.parseExpected(ast.KindCloseBracketToken)
				typeNode = p.finishNode(p.factory.NewArrayTypeNode(typeNode), pos)
			}
		default:
			return typeNode
		}
	}
	return typeNode
}

func (p *Parser) nextIsStartOfType() bool {
	p.nextToken()
	return p.isStartOfType(false /*inStartOfParameter*/)
}

func (p *Parser) parseNonArrayType() *ast.Node {
	switch p.token {
	case ast.KindAnyKeyword, ast.KindUnknownKeyword, ast.KindStringKeyword, ast.KindNumberKeyword,
		ast.KindSymbolKeyword, ast.KindBooleanKeyword, ast.KindNilKeyword, ast.KindNeverKeyword, ast.KindObjectKeyword,
		ast.KindThreadKeyword, ast.KindUserdataKeyword, ast.KindCDataKeyword,
		ast.KindFunctionKeyword:
		state := p.mark()
		keywordTypeNode := p.parseKeywordTypeNode()
		// If these are followed by a dot then parse these out as a dotted type reference instead
		if p.token != ast.KindDotToken {
			return keywordTypeNode
		}
		p.rewind(state)
		return p.parseTypeReference()
	case ast.KindAsteriskEqualsToken:
		// If there is '*=', treat it as * followed by postfix =
		p.scanner.ReScanAsteriskEqualsToken()
		fallthrough
	case ast.KindAsteriskToken:
		return p.parseJSDocAllType()
	case ast.KindQuestionToken:
		return p.parseJSDocNullableType()
	case ast.KindExclamationToken:
		// Only the punctuation `!` is the JSDoc non-nullable prefix; the word
		// `not` in type position falls through to stock recovery.
		if p.tokenIsExclamationPunctuation() {
			return p.parseJSDocNonNullableType()
		}
		return p.parseTypeReference()
	case ast.KindNoSubstitutionTemplateLiteral, ast.KindStringLiteral, ast.KindNumericLiteral, ast.KindTrueKeyword,
		ast.KindFalseKeyword:
		return p.parseLiteralTypeNode(false /*negative*/)
	case ast.KindMinusToken:
		if p.lookAhead((*Parser).nextTokenIsNumericLiteral) {
			return p.parseLiteralTypeNode(true /*negative*/)
		}
		return p.parseTypeReference()
	case ast.KindVoidKeyword:
		return p.parseKeywordTypeNode()
	case ast.KindTypeOfKeyword:
		if p.lookAhead((*Parser).nextIsStartOfTypeOfImportType) {
			return p.parseImportType()
		}
		return p.parseTypeQuery()
	case ast.KindOpenBraceToken:
		if p.lookAhead((*Parser).nextIsStartOfMappedType) {
			return p.parseMappedType()
		}
		return p.parseTypeLiteral()
	case ast.KindOpenBracketToken:
		// Tuple type syntax `[A, B]`: a fixed-shape table whose elements live
		// at the 1-based number keys. Values are ordinary table constructors
		// (`{a, b}`); only the type syntax uses brackets.
		return p.parseTupleType()
	case ast.KindOpenParenToken:
		return p.parseParenthesizedType()
	case ast.KindImportKeyword:
		return p.parseImportType()
	case ast.KindAssertsKeyword:
		// `asserts this is T` is removed with the `this` type. The predicate
		// production can only consume an ordinary identifier, so the lookahead
		// requires one; committing on a reserved word like `this` would mint a
		// predicate with an empty missing name and cascade.
		if p.lookAhead((*Parser).nextTokenIsIdentifierOnSameLine) {
			return p.parseAssertsTypePredicate()
		}
		return p.parseTypeReference()
	case ast.KindTemplateHead:
		return p.parseTemplateType()
	default:
		return p.parseTypeReference()
	}
}

func (p *Parser) parseKeywordTypeNode() *ast.Node {
	pos := p.nodePos()
	result := p.factory.NewKeywordTypeNode(p.token)
	p.nextToken()
	return p.finishNode(result, pos)
}

func (p *Parser) parseJSDocAllType() *ast.Node {
	pos := p.nodePos()
	p.nextToken()
	return p.finishNode(p.factory.NewJSDocAllType(), pos)
}

func (p *Parser) parseJSDocNonNullableType() *ast.TypeNode {
	pos := p.nodePos()
	p.nextToken()
	return p.finishNode(p.factory.NewJSDocNonNullableType(p.parseTypeOperatorOrHigher()), pos)
}

func (p *Parser) parseJSDocNullableType() *ast.Node {
	pos := p.nodePos()
	// skip the ?
	p.nextToken()
	return p.finishNode(p.factory.NewJSDocNullableType(p.parseTypeOperatorOrHigher()), pos)
}

func (p *Parser) parseJSDocType() *ast.TypeNode {
	p.scanner.SetSkipJSDocLeadingAsterisks(true)
	pos := p.nodePos()

	hasDotDotDot := p.parseOptional(ast.KindDotDotDotToken)
	t := p.parseTypeOrTypePredicate()
	p.scanner.SetSkipJSDocLeadingAsterisks(false)
	if hasDotDotDot {
		t = p.finishNode(p.factory.NewJSDocVariadicType(t), pos)
	}
	if p.token == ast.KindEqualsToken {
		p.nextToken()
		return p.finishNode(p.factory.NewJSDocOptionalType(t), pos)
	}
	return t
}

func (p *Parser) parseLiteralTypeNode(negative bool) *ast.Node {
	pos := p.nodePos()
	if negative {
		p.nextToken()
	}
	var expression *ast.Expression
	if p.token == ast.KindTrueKeyword || p.token == ast.KindFalseKeyword || p.token == ast.KindNilKeyword {
		expression = p.parseKeywordExpression()
	} else {
		expression = p.parseLiteralExpression(false /*intern*/)
	}
	if negative {
		expression = p.finishNode(p.factory.NewPrefixUnaryExpression(ast.KindMinusToken, expression), pos)
	}
	return p.finishNode(p.factory.NewLiteralTypeNode(expression), pos)
}

func (p *Parser) parseTypeReference() *ast.Node {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewTypeReferenceNode(p.parseEntityNameOfTypeReference(), p.parseTypeArgumentsOfTypeReference()), pos)
}

func (p *Parser) parseEntityNameOfTypeReference() *ast.Node {
	return p.parseEntityName(true /*allowReservedWords*/, diagnostics.Type_expected)
}

func (p *Parser) parseEntityName(allowReservedWords bool, diagnosticMessage *diagnostics.Message) *ast.Node {
	pos := p.nodePos()
	var entity *ast.Node
	if allowReservedWords {
		entity = p.parseIdentifierNameWithDiagnostic(diagnosticMessage)
	} else {
		entity = p.parseIdentifierWithDiagnostic(diagnosticMessage, nil)
	}
	for p.parseOptional(ast.KindDotToken) {
		if p.token == ast.KindLessThanToken {
			// The entity is part of a JSDoc-style generic. We will use the gap between `typeName` and
			// `typeArguments` to report it as a grammar error in the checker.
			break
		}
		entity = p.finishNode(p.factory.NewQualifiedName(entity, p.parseRightSideOfDot(allowReservedWords, false /*allowPrivateIdentifiers*/, true /*allowUnicodeEscapeSequenceInIdentifierName*/)), pos)
	}
	return entity
}

func (p *Parser) parseRightSideOfDot(allowIdentifierNames bool, allowPrivateIdentifiers bool, allowUnicodeEscapeSequenceInIdentifierName bool) *ast.Node {
	// Technically a keyword is valid here as all identifiers and keywords are identifier names.
	// However, often we'll encounter this in error situations when the identifier or keyword
	// is actually starting another valid construct.
	//
	// So, we check for the following specific case:
	//
	//      name.
	//      identifierOrKeyword identifierNameOrKeyword
	//
	// Note: the newlines are important here.  For example, if that above code
	// were rewritten into:
	//
	//      name.identifierOrKeyword
	//      identifierNameOrKeyword
	//
	// Then we would consider it valid.  That's because ASI would take effect and
	// the code would be implicitly: "name.identifierOrKeyword; identifierNameOrKeyword".
	// In the first case though, ASI will not take effect because there is not a
	// line terminator after the identifier or keyword.
	if p.hasPrecedingLineBreak() && tokenIsIdentifierOrKeyword(p.token) && p.lookAhead((*Parser).nextTokenIsIdentifierOrKeywordOnSameLine) {
		// Report that we need an identifier.  However, report it right after the dot,
		// and not on the next token.  This is because the next token might actually
		// be an identifier and the error would be quite confusing.
		p.parseErrorAt(p.nodePos(), p.nodePos(), diagnostics.Identifier_expected)
		return p.createMissingIdentifier()
	}
	if p.token == ast.KindPrivateIdentifier {
		node := p.parsePrivateIdentifier()
		if allowPrivateIdentifiers {
			return node
		}
		p.parseErrorAt(p.nodePos(), p.nodePos(), diagnostics.Identifier_expected)
		return p.createMissingIdentifier()
	}
	if allowIdentifierNames {
		if allowUnicodeEscapeSequenceInIdentifierName {
			return p.parseIdentifierName()
		}
		return p.parseIdentifierNameErrorOnUnicodeEscapeSequence()
	}
	id := p.parseIdentifier()
	return id
}

func (p *Parser) newIdentifier(text string) *ast.Node {
	p.identifierCount++
	id := p.factory.NewIdentifier(text)
	return id
}

func (p *Parser) createMissingIdentifier() *ast.Node {
	return p.finishNode(p.newIdentifier(""), p.nodePos())
}

func (p *Parser) parsePrivateIdentifier() *ast.Node {
	pos := p.nodePos()
	text := p.scanner.TokenValue()
	p.nextToken()
	return p.finishNode(p.factory.NewPrivateIdentifier(p.internIdentifier(text)), pos)
}

func (p *Parser) reScanLessThanToken() ast.Kind {
	p.token = p.scanner.ReScanLessThanToken()
	return p.token
}

func (p *Parser) reScanGreaterThanToken() ast.Kind {
	p.token = p.scanner.ReScanGreaterThanToken()
	return p.token
}

func (p *Parser) reScanSlashToken() ast.Kind {
	p.token = p.scanner.ReScanSlashToken()
	return p.token
}

func (p *Parser) reScanTemplateToken(isTaggedTemplate bool) ast.Kind {
	p.token = p.scanner.ReScanTemplateToken(isTaggedTemplate)
	return p.token
}

func (p *Parser) parseTypeArgumentsOfTypeReference() *ast.NodeList {
	if !p.hasPrecedingLineBreak() && p.reScanLessThanToken() == ast.KindLessThanToken {
		return p.parseTypeArguments()
	}
	return nil
}

func (p *Parser) parseTypeArguments() *ast.NodeList {
	if p.token == ast.KindLessThanToken {
		return p.parseBracketedList(PCTypeArguments, (*Parser).parseTypeArgument, ast.KindLessThanToken, ast.KindGreaterThanToken)
	}
	return nil
}

// parseTypeArgument recognizes a value pack where a generic pack parameter can
// consume one. Multi-element and variadic packs share the return-pack grammar;
// the empty pack needs its own spelling because `()` is not otherwise a type.
func (p *Parser) parseTypeArgument() *ast.TypeNode {
	if p.isParenthesizedReturnTypeList() {
		// Deliberately NO union continuation here (unlike return position): the
		// checker instantiates a `<...A>` pack parameter by wrapping a non-pack
		// argument in a one-value pack, so a union-of-packs argument would
		// silently mistype every call site ("Expected 1 arguments"). Until that
		// lifting distributes, the parse error is the honest failure.
		return p.parseParenthesizedReturnTypeList()
	}
	// `()` is the empty pack, but only when it is not the parameter list of a
	// zero-argument function type `() => R` -- that stays an ordinary type.
	if p.token == ast.KindOpenParenToken && p.lookAhead(func(p *Parser) bool {
		return p.nextToken() == ast.KindCloseParenToken && p.nextToken() != ast.KindEqualsGreaterThanToken
	}) {
		pos := p.nodePos()
		p.nextToken()
		elements := p.parseEmptyNodeList()
		p.parseExpected(ast.KindCloseParenToken)
		return p.finishNode(p.factory.NewMultiReturnTypeNode(elements), pos)
	}
	return p.parseType()
}

func (p *Parser) nextIsStartOfTypeOfImportType() bool {
	p.nextToken()
	return p.token == ast.KindImportKeyword
}

func (p *Parser) parseImportType() *ast.Node {
	p.sourceFlags |= ast.NodeFlagsPossiblyContainsDynamicImport
	pos := p.nodePos()
	isTypeOf := p.parseOptional(ast.KindTypeOfKeyword)
	p.parseExpected(ast.KindImportKeyword)
	p.parseExpected(ast.KindOpenParenToken)
	typeNode := p.parseType()
	p.parseExpected(ast.KindCloseParenToken)
	var qualifier *ast.Node
	if p.parseOptional(ast.KindDotToken) {
		qualifier = p.parseEntityNameOfTypeReference()
	}
	typeArguments := p.parseTypeArgumentsOfTypeReference()
	return p.finishNode(p.factory.NewImportTypeNode(isTypeOf, typeNode, qualifier, typeArguments), pos)
}

func (p *Parser) parseTypeQuery() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindTypeOfKeyword)
	entityName := p.parseEntityName(true /*allowReservedWords*/, nil)
	// Make sure we perform ASI to prevent parsing the next line's type arguments as part of an instantiation expression
	var typeArguments *ast.NodeList
	if !p.hasPrecedingLineBreak() {
		typeArguments = p.parseTypeArguments()
	}
	return p.finishNode(p.factory.NewTypeQueryNode(entityName, typeArguments), pos)
}

func (p *Parser) nextIsStartOfMappedType() bool {
	p.nextToken()
	if p.token == ast.KindPlusToken || p.token == ast.KindMinusToken {
		return p.nextToken() == ast.KindReadonlyKeyword
	}
	if p.token == ast.KindReadonlyKeyword {
		p.nextToken()
	}
	return p.token == ast.KindOpenBracketToken && p.nextTokenIsIdentifier() && p.nextToken() == ast.KindInKeyword
}

func (p *Parser) parseMappedType() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindOpenBraceToken)
	var readonlyToken *ast.Node // ReadonlyKeyword | PlusToken | MinusToken
	if p.token == ast.KindReadonlyKeyword || p.token == ast.KindPlusToken || p.token == ast.KindMinusToken {
		readonlyToken = p.parseTokenNode()
		if readonlyToken.Kind != ast.KindReadonlyKeyword {
			p.parseExpected(ast.KindReadonlyKeyword)
		}
	}
	p.parseExpected(ast.KindOpenBracketToken)
	typeParameter := p.parseMappedTypeParameter()
	var nameType *ast.TypeNode
	if p.parseOptional(ast.KindAsKeyword) {
		nameType = p.parseType()
	}
	p.parseExpected(ast.KindCloseBracketToken)
	var questionToken *ast.Node // QuestionToken | PlusToken | MinusToken
	if p.token == ast.KindQuestionToken || p.token == ast.KindPlusToken || p.token == ast.KindMinusToken {
		questionToken = p.parseTokenNode()
		if questionToken.Kind != ast.KindQuestionToken {
			p.parseExpected(ast.KindQuestionToken)
		}
	}
	typeNode := p.parseTypeAnnotation()
	p.parseSemicolon()
	members := p.parseList(PCTypeMembers, (*Parser).parseTypeMember)
	p.parseExpected(ast.KindCloseBraceToken)
	return p.finishNode(p.factory.NewMappedTypeNode(readonlyToken, typeParameter, nameType, questionToken, typeNode, members), pos)
}

func (p *Parser) parseMappedTypeParameter() *ast.Node {
	pos := p.nodePos()
	name := p.parseIdentifierName()
	p.parseExpected(ast.KindInKeyword)
	typeNode := p.parseType()
	return p.finishNode(p.factory.NewTypeParameterDeclaration(nil /*modifiers*/, nil /*dotDotDotToken*/, name, typeNode, nil /*expression*/, nil /*defaultType*/), pos)
}

func (p *Parser) parseTypeMember() *ast.Node {
	if p.token == ast.KindOpenParenToken || p.token == ast.KindLessThanToken {
		return p.parseSignatureMember(ast.KindCallSignature)
	}
	if p.token == ast.KindNewKeyword && p.lookAhead((*Parser).nextTokenIsOpenParenOrLessThan) {
		return p.parseSignatureMember(ast.KindConstructSignature)
	}
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	modifiers := p.parseModifiers()
	if p.isIndexSignature() {
		return p.parseIndexSignatureDeclaration(pos, jsdoc, modifiers)
	}
	return p.parsePropertyOrMethodSignature(pos, jsdoc, modifiers)
}

func (p *Parser) nextTokenIsOpenParenOrLessThan() bool {
	p.nextToken()
	return p.token == ast.KindOpenParenToken || p.token == ast.KindLessThanToken
}

func (p *Parser) parseSignatureMember(kind ast.Kind) *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	if kind == ast.KindConstructSignature {
		p.parseExpected(ast.KindNewKeyword)
	}
	typeParameters := p.parseTypeParameters()
	parameters := p.parseParameters(ParseFlagsType)
	var typeNode *ast.TypeNode
	if kind == ast.KindCallSignature {
		typeNode = p.parseCallableReturnType(ast.KindColonToken, true /*isType*/)
	} else {
		typeNode = p.parseReturnType(ast.KindColonToken, true /*isType*/)
	}
	p.parseTypeMemberSemicolon()
	var result *ast.Node
	if kind == ast.KindCallSignature {
		result = p.factory.NewCallSignatureDeclaration(typeParameters, parameters, typeNode)
	} else {
		result = p.factory.NewConstructSignatureDeclaration(typeParameters, parameters, typeNode)
	}
	p.finishNode(result, pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseTypeParameters() *ast.NodeList {
	if p.token == ast.KindLessThanToken {
		return p.parseBracketedList(PCTypeParameters, (*Parser).parseTypeParameter, ast.KindLessThanToken, ast.KindGreaterThanToken)
	}
	return nil
}

func (p *Parser) parseTypeParameter() *ast.Node {
	pos := p.nodePos()
	modifiers := p.parseModifiersEx(false /*stopOnStartOfClassStaticBlock*/)
	// A leading `...` declares a Lua generic pack parameter (`<...A>`): its value
	// is a whole value pack, and `...A` / `...: A` spread it. The grammar checks
	// enforce that pack parameters trail plain ones and appear only in pack
	// positions (see checkGrammarTypeParameterList).
	dotDotDotToken := p.parseOptionalToken(ast.KindDotDotDotToken)
	name := p.parseIdentifier()
	var constraint *ast.TypeNode
	var expression *ast.Expression
	if p.parseOptional(ast.KindExtendsKeyword) {
		// It's not uncommon for people to write improper constraints to a generic.  If the
		// user writes a constraint that is an expression and not an actual type, then parse
		// it out as an expression (so we can recover well), but report that a type is needed
		// instead.
		if p.isStartOfType(false /*inStartOfParameter*/) || !p.isStartOfExpression() {
			constraint = p.parseType()
		} else {
			// It was not a type, and it looked like an expression.  Parse out an expression
			// here so we recover well.  Note: it is important that we call parseUnaryExpression
			// and not parseExpression here.  If the user has:
			//
			//      <T extends "">
			//
			// We do *not* want to consume the `>` as we're consuming the expression for "".
			expression = p.parseUnaryExpressionOrHigher()
		}
	}
	var defaultType *ast.TypeNode
	if p.parseOptional(ast.KindEqualsToken) {
		defaultType = p.parseType()
	}
	result := p.factory.NewTypeParameterDeclaration(modifiers, dotDotDotToken, name, constraint, expression, defaultType)
	return p.finishNode(result, pos)
}

func (p *Parser) parseParameters(flags ParseFlags) *ast.NodeList {
	// FormalParameters [Yield,Await]: (modified)
	//      [empty]
	//      FormalParameterList[?Yield,Await]
	//
	// FormalParameter[Yield,Await]: (modified)
	//      BindingElement[?Yield,Await]
	//
	// BindingElement [Yield,Await]: (modified)
	//      SingleNameBinding[?Yield,?Await]
	//      BindingPattern[?Yield,?Await]Initializer [In, ?Yield,?Await] opt
	//
	// SingleNameBinding [Yield,Await]:
	//      BindingIdentifier[?Yield,?Await]Initializer [In, ?Yield,?Await] opt
	if p.parseExpected(ast.KindOpenParenToken) {
		parameters := p.parseParametersWorker(flags, true /*allowAmbiguity*/)
		p.parseExpected(ast.KindCloseParenToken)
		return parameters
	}
	return p.parseEmptyNodeList()
}

func (p *Parser) parseParametersWorker(flags ParseFlags, allowAmbiguity bool) *ast.NodeList {
	// FormalParameters [Yield,Await]: (modified)
	//      [empty]
	//      FormalParameterList[?Yield,Await]
	//
	// FormalParameter[Yield,Await]: (modified)
	//      BindingElement[?Yield,Await]
	//
	// BindingElement [Yield,Await]: (modified)
	//      SingleNameBinding[?Yield,?Await]
	//      BindingPattern[?Yield,?Await]Initializer [In, ?Yield,?Await] opt
	//
	// SingleNameBinding [Yield,Await]:
	//      BindingIdentifier[?Yield,?Await]Initializer [In, ?Yield,?Await] opt
	parameters := p.parseDelimitedList(PCParameters, func(p *Parser) *ast.Node {
		parameter := p.parseParameterEx(allowAmbiguity)
		if parameter != nil && flags&ParseFlagsType == 0 {
			p.checkJSSyntax(parameter)
		}
		return parameter
	})
	return parameters
}

func (p *Parser) parseParameter() *ast.Node {
	return p.parseParameterEx(true /*allowAmbiguity*/)
}

func (p *Parser) parseParameterEx(allowAmbiguity bool) *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	// FormalParameter [Yield,Await]:
	//      BindingElement[?Yield,?Await]
	modifiers := p.parseModifiersEx(false /*stopOnStartOfClassStaticBlock*/)
	// `this` parameters are removed from tlua. A `this` token here is rejected
	// as a reserved word (no parameter named `this` is ever created) and list
	// recovery skips the token.
	dotDotDotToken := p.parseOptionalToken(ast.KindDotDotDotToken)
	isVararg := dotDotDotToken != nil && p.isVarargParameterFollow()
	if !allowAmbiguity && !isVararg && !p.isParameterNameStart() {
		return nil
	}
	var name *ast.Node
	if isVararg {
		name = p.createVarargParameterName()
	} else {
		// A legacy TS named rest parameter (`...args: T[]`) parses into a complete
		// node and is rejected by checkGrammarParameterList, not here: a *syntactic*
		// diagnostic suppresses every semantic diagnostic in the whole program, so
		// one un-migrated parameter would silently turn off type checking.
		name = p.parseNameOfParameter(modifiers)
	}
	// The `?` and `=` are parsed here so they reach the existing rest-parameter
	// grammar checks rather than generic parse recovery.
	result := p.factory.NewParameterDeclaration(
		modifiers,
		dotDotDotToken,
		name,
		p.parseOptionalToken(ast.KindQuestionToken),
		p.parseTypeAnnotation(),
		p.parseInitializer(),
	)
	p.withJSDoc(p.finishNode(result, pos), jsdoc)
	return result
}

// isVarargParameterFollow reports whether the current token continues a vararg
// *parameter* -- that is, whether the `...` (or `...name`) just consumed sits in
// a parameter list rather than in an expression or a return pack. It is the one
// definition of that follow set: parseParameterEx, the arrow lookahead's
// named-rest check, and isParenthesizedReturnTypeList key off it, and they must
// agree or the same text parses differently depending on which speculative path
// runs first. The arrow lookahead's *direct*-follow case cannot use it
// wholesale: there `,` and `)`-then-`:` are genuinely ambiguous with expressions
// (`(..., x)`, `case (...):`) and resolve by speculative parse instead.
func (p *Parser) isVarargParameterFollow() bool {
	switch p.token {
	case ast.KindColonToken, ast.KindCommaToken, ast.KindCloseParenToken, ast.KindQuestionToken, ast.KindEqualsToken:
		return true
	}
	return false
}

// createVarargParameterName synthesizes the zero-width `...` identifier that
// stands in for the vararg parameter's absent binding name. See
// ast.VarargParameterName.
func (p *Parser) createVarargParameterName() *ast.Node {
	return p.finishNode(p.newIdentifier(ast.VarargParameterName), p.nodePos())
}

func (p *Parser) isParameterNameStart() bool {
	// Be permissive about await and yield by calling isBindingIdentifier instead of isIdentifier; disallowing
	// them during a speculative parse leads to many more follow-on errors than allowing the function to parse then later
	// complaining about the use of the keywords.
	return p.isBindingIdentifier() || p.token == ast.KindOpenBraceToken || p.isArrayBindingPatternStart()
}

func (p *Parser) parseNameOfParameter(modifiers *ast.ModifierList) *ast.Node {
	// FormalParameter [Yield,Await]:
	//      BindingElement[?Yield,?Await]
	name := p.parseIdentifierOrPatternWithDiagnostic(diagnostics.Private_identifiers_cannot_be_used_as_parameters)
	if name.Loc.Len() == 0 && modifiers == nil && ast.IsModifierKind(p.token) {
		// in cases like
		// 'use strict'
		// function foo(static)
		// isParameter('static') == true, because of isModifier('static')
		// however 'static' is not a legal identifier in a strict mode.
		// so result of this function will be Parameter (flags = 0, name = missing, type = undefined, initializer = undefined)
		// and current token will not change => parsing of the enclosing parameter list will last till the end of time (or OOM)
		// to avoid this we'll advance cursor to the next token.
		p.nextToken()
	}
	return name
}

func (p *Parser) parseReturnType(returnToken ast.Kind, isType bool) *ast.TypeNode {
	if p.shouldParseReturnType(returnToken, isType) {
		return doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTypeOrTypePredicate)
	}
	return nil
}

// parseCallableReturnType accepts an explicitly parenthesized return pack,
// avoiding collisions with commas in the surrounding type grammar.
func (p *Parser) parseCallableReturnType(returnToken ast.Kind, isType bool) *ast.TypeNode {
	if !p.shouldParseReturnType(returnToken, isType) {
		return nil
	}
	pos := p.nodePos()
	if p.isParenthesizedReturnTypeList() {
		// A pack-LEADING return type: claim the pack before parseType's function-type
		// lookahead could mistake it for a parameter list, then continue into a union
		// (packs allowed in later positions) when `|` follows.
		return p.parseReturnTypeUnionRest(p.parseParenthesizedReturnTypeList(), pos)
	}
	// Otherwise an ordinary type, a conditional (`() => T extends A ? 1 : 2`), or a
	// union whose LATER constituents may be packs (`nil | (number, string)`). parseType
	// handles conditionals and the leading union exactly as elsewhere; the
	// packsInReturnUnion flag lets its outermost union recognize a parenthesized pack.
	return p.parseReturnTypeAllowingPacks()
}

// parseReturnTypeAllowingPacks parses a non-pack-leading return type with conditional
// types allowed (a return type is not a conditional-disallowing position) and its
// outermost union taught to recognize a parenthesized pack in any constituent.
func (p *Parser) parseReturnTypeAllowingPacks() *ast.TypeNode {
	return doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, func(p *Parser) *ast.TypeNode {
		saved := p.packsInReturnUnion
		p.packsInReturnUnion = true
		result := p.parseTypeOrTypePredicate()
		p.packsInReturnUnion = saved
		return result
	})
}

// parseReturnTypeUnionRest continues an already-parsed pack-leading return constituent
// into a union when a `|` follows, so each later constituent may itself be a pack:
//
//	function f(): (number, nil) | (nil, string)
//
// The non-pack-leading spelling (`nil | (number, string)`) is handled by
// parseReturnTypeAllowingPacks instead. With no `|` following, first is returned
// unwrapped.
func (p *Parser) parseReturnTypeUnionRest(first *ast.TypeNode, pos int) *ast.TypeNode {
	if p.token != ast.KindBarToken {
		return first
	}
	types := make([]*ast.Node, 1, 8)
	types[0] = first
	for p.parseOptional(ast.KindBarToken) {
		if p.isParenthesizedReturnTypeList() {
			types = append(types, p.parseParenthesizedReturnTypeList())
		} else {
			// An ordinary constituent, through parseFunctionOrConstructorTypeToError as
			// the union grammar does: an unparenthesized function type gets its own
			// diagnostic instead of a parse-error cascade. Packs are claimed above.
			types = append(types, doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, func(p *Parser) *ast.TypeNode {
				return p.parseFunctionOrConstructorTypeToError(true /*isUnionType*/, (*Parser).parseIntersectionTypeOrHigher)
			}))
		}
	}
	list := p.newNodeList(core.NewTextRange(pos, p.nodePos()), p.nodeSliceArena.Clone(types))
	return p.finishNode(p.factory.NewUnionTypeNode(list), pos)
}

// isParenthesizedReturnTypeList distinguishes a parenthesized return pack from a
// parenthesized or callable type in return position. A leading `...` is the
// ambiguous case: `(...) => R` and `(...: T) => R` are callable types with a
// vararg parameter, while `(...T)` is a pure-variadic return pack, which -- being
// a single element -- carries no comma to key off.
func (p *Parser) isParenthesizedReturnTypeList() bool {
	return p.token == ast.KindOpenParenToken && p.lookAhead(func(p *Parser) bool {
		p.nextToken()
		if p.token == ast.KindDotDotDotToken {
			p.nextToken()
			if p.isVarargParameterFollow() {
				// `(...) => R`, `(...: T) => R`, and the malformed parameter lists
				// (`(...?: T)`, `(..., x)`) that the parameter grammar reports
				// precisely -- reading those as a pack would bury the real
				// diagnostic under parse recovery.
				return false
			}
			// `(...x: T) => R`: a named rest, which is rejected, but still a
			// parameter list. A pack tail's type never has a `:` after it.
			if p.isIdentifier() && p.nextToken() == ast.KindColonToken {
				return false
			}
			return true
		}
		doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTypeOrTypePredicate)
		return p.token == ast.KindCommaToken
	})
}

func (p *Parser) parseParenthesizedReturnTypeList() *ast.TypeNode {
	pos := p.nodePos()
	p.nextToken()
	var elements []*ast.TypeNode
	if p.token == ast.KindDotDotDotToken {
		// A pure-variadic pack `(...T)`: no leading fixed element, so no comma.
		rest := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTupleElementType)
		elements = append(elements, p.normalizeReturnListElement(rest))
	} else {
		elements = append(elements, p.parseFirstReturnTypeListElement())
	}
	for p.parseOptional(ast.KindCommaToken) {
		element := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTupleElementType)
		elements = append(elements, p.normalizeReturnListElement(element))
	}
	p.parseExpected(ast.KindCloseParenToken)
	list := p.newNodeList(core.NewTextRange(elements[0].Pos(), elements[len(elements)-1].End()), elements)
	return p.finishNode(p.factory.NewMultiReturnTypeNode(list), pos)
}

// parseFirstReturnTypeListElement parses the leading fixed element of a return
// pack, which is reached only when a comma follows it.
func (p *Parser) parseFirstReturnTypeListElement() *ast.TypeNode {
	first := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTypeOrTypePredicate)
	debug.Assert(p.token == ast.KindCommaToken, "parenthesized return pack must contain a comma")
	// A type predicate element is rejected by checkGrammarMultiReturnType, not
	// here: as a parse error it would suppress every semantic diagnostic in the
	// program -- and, in this file, the grammar checks for the pack itself.
	return p.normalizeReturnListElement(first)
}

// parseReturnTypeList parses a declaration-header return type. A multiple
// return type must be explicitly parenthesized:
//
//	function f(): (number, string)
//
// A bare comma-separated list is still parsed for error recovery, with a
// diagnostic on the list. Callable type positions go through
// parseCallableReturnType.
func (p *Parser) parseReturnTypeList(returnToken ast.Kind) *ast.TypeNode {
	if !p.shouldParseReturnType(returnToken, false /*isType*/) {
		return nil
	}
	pos := p.nodePos()
	if p.isParenthesizedReturnTypeList() {
		return p.parseReturnTypeUnionRest(p.parseParenthesizedReturnTypeList(), pos)
	}
	// An ordinary type, a conditional, or a union whose later constituents may be
	// packs (`nil | (number, string)`); a bare comma after it is the unparenthesized
	// multi-return recovery below.
	first := p.parseReturnTypeAllowingPacks()
	if p.token != ast.KindCommaToken {
		return first
	}
	elements := []*ast.TypeNode{p.normalizeReturnListElement(first)}
	for p.parseOptional(ast.KindCommaToken) {
		element := doInContext(p, ast.NodeFlagsDisallowConditionalTypesContext, false, (*Parser).parseTupleElementType)
		elements = append(elements, p.normalizeReturnListElement(element))
	}
	// The missing parentheses are reported by checkGrammarMultiReturnType, which
	// recognizes the bare form because the node starts at its first element rather
	// than at a `(`. Reporting here would make it a syntactic diagnostic, which
	// suppresses every semantic diagnostic in the program.
	list := p.newNodeList(core.NewTextRange(elements[0].Pos(), elements[len(elements)-1].End()), elements)
	return p.finishNode(p.factory.NewMultiReturnTypeNode(list), pos)
}

// normalizeReturnListElement consumes a trailing same-line `?` that
// parsePostfixTypeOrHigher left unconsumed (it defers `?` to a potential
// conditional type whenever the next token can start a type — e.g. the `{` of
// a brace body). Inside a declaration-header return list no conditional type
// can enclose the element, so the `?` is always the optional-element marker.
func (p *Parser) normalizeReturnListElement(typeNode *ast.TypeNode) *ast.TypeNode {
	// A postfix `T?` parsed before the list was recognized comes back as a
	// JSDocNullableType; rewrite it exactly like parseTupleElementType does.
	if ast.IsJSDocNullableType(typeNode) && typeNode.Pos() == typeNode.Type().Pos() {
		node := p.factory.NewOptionalTypeNode(typeNode.Type())
		node.Flags = typeNode.Flags
		node.Loc = typeNode.Loc
		typeNode.Type().Parent = node
		return node
	}
	if p.token == ast.KindQuestionToken && !p.hasPrecedingLineBreak() {
		pos := typeNode.Pos()
		p.nextToken()
		return p.finishNode(p.factory.NewOptionalTypeNode(typeNode), pos)
	}
	return typeNode
}

func (p *Parser) shouldParseReturnType(returnToken ast.Kind, isType bool) bool {
	if returnToken == ast.KindEqualsGreaterThanToken {
		p.parseExpected(returnToken)
		return true
	} else if p.parseOptional(ast.KindColonToken) {
		return true
	} else if isType && p.token == ast.KindEqualsGreaterThanToken {
		// This is easy to get backward, especially in type contexts, so parse the type anyway
		p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(ast.KindColonToken))
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) parseTypeOrTypePredicate() *ast.TypeNode {
	if p.isIdentifier() {
		state := p.mark()
		pos := p.nodePos()
		id := p.parseIdentifier()
		if p.token == ast.KindIsKeyword && !p.hasPrecedingLineBreak() {
			p.nextToken()
			// A predicate's asserted type (`x is T`) is a single-value narrowing type,
			// never the outermost return union, so a pack must not be recognized inside
			// it. Consume packsInReturnUnion here so `x is nil | (number, string)` does
			// not parse a pack; parseReturnTypeAllowingPacks restores it afterward.
			p.packsInReturnUnion = false
			return p.finishNode(p.factory.NewTypePredicateNode(nil /*assertsModifier*/, id, p.parseType()), pos)
		}
		p.rewind(state)
	}
	return p.parseType()
}

func (p *Parser) parseTypeMemberSemicolon() {
	// We allow type members to be separated by commas or (possibly ASI) semicolons.
	// First check if it was a comma.  If so, we're done with the member.
	if p.parseOptional(ast.KindCommaToken) {
		return
	}
	// Didn't have a comma.  We must have a (possible ASI) semicolon.
	p.parseSemicolon()
}

func (p *Parser) parsePropertyName() *ast.Node {
	prop := p.parsePropertyNameWorker(true /*allowComputedPropertyNames*/)
	return prop
}

func (p *Parser) parsePropertyNameWorker(allowComputedPropertyNames bool) *ast.Node {
	if p.token == ast.KindStringLiteral || p.token == ast.KindNumericLiteral {
		literal := p.parseLiteralExpression(true /*intern*/)
		return literal
	}
	if allowComputedPropertyNames && p.token == ast.KindOpenBracketToken {
		return p.parseComputedPropertyName()
	}
	if p.token == ast.KindPrivateIdentifier {
		return p.parsePrivateIdentifier()
	}
	return p.parseIdentifierName()
}

func (p *Parser) parseComputedPropertyName() *ast.Node {
	// PropertyName [Yield]:
	//      LiteralPropertyName
	//      ComputedPropertyName[?Yield]
	pos := p.nodePos()
	p.parseExpected(ast.KindOpenBracketToken)
	// We parse any expression (including a comma expression). But the grammar
	// says that only an assignment expression is allowed, so the grammar checker
	// will error if it sees a comma expression.
	expression := p.parseExpressionAllowIn()
	p.parseExpected(ast.KindCloseBracketToken)
	return p.finishNode(p.factory.NewComputedPropertyName(expression), pos)
}

// finishLuaEndBlock parses statements up to and consuming `end`, wrapping them in
// a Lua-flagged Block whose range starts at pos. Shared by function bodies and
// standalone `do ... end` blocks so the `end`-block shape stays in one place.
func (p *Parser) finishLuaEndBlock(pos int) *ast.Node {
	statements := p.parseList(PCLuaBlockStatements, (*Parser).parseStatement)
	p.parseExpected(ast.KindEndKeyword)
	block := p.finishNode(p.factory.NewBlock(statements, true /*multiLine*/), pos)
	block.Flags |= ast.NodeFlagsLuaBlock
	return block
}

func (p *Parser) parseLuaFunctionBlock() *ast.Node {
	return p.finishLuaEndBlock(p.nodePos())
}

// parseLuaClauseBlock parses statements up to (not including) the clause
// terminator of a Lua control-flow construct and wraps them in a Lua-flagged
// Block. The caller consumes the terminator so it stays inside the enclosing
// statement's range.
func (p *Parser) parseLuaClauseBlock(context ParsingContext) *ast.Node {
	pos := p.nodePos()
	statements := p.parseList(context, (*Parser).parseStatement)
	block := p.finishNode(p.factory.NewBlock(statements, true /*multiLine*/), pos)
	block.Flags |= ast.NodeFlagsLuaBlock
	return block
}

func (p *Parser) isIndexSignature() bool {
	return p.token == ast.KindOpenBracketToken && p.lookAhead((*Parser).nextIsUnambiguouslyIndexSignature)
}

func (p *Parser) nextIsUnambiguouslyIndexSignature() bool {
	// The only allowed sequence is:
	//
	//   [id:
	//
	// However, for error recovery, we also check the following cases:
	//
	//   [...
	//   [id,
	//   [id?,
	//   [id?:
	//   [id?]
	//   [public id
	//   [private id
	//   [protected id
	//   []
	//
	p.nextToken()
	if p.token == ast.KindDotDotDotToken || p.token == ast.KindCloseBracketToken {
		return true
	}
	if ast.IsModifierKind(p.token) {
		p.nextToken()
		if p.isIdentifier() {
			return true
		}
	} else if !p.isIdentifier() {
		return false
	} else {
		// Skip the identifier
		p.nextToken()
	}
	// A colon signifies a well formed indexer
	// A comma should be a badly formed indexer because comma expressions are not allowed
	// in computed properties.
	if p.token == ast.KindColonToken || p.token == ast.KindCommaToken {
		return true
	}
	// Question mark could be an indexer with an optional property,
	// or it could be a conditional expression in a computed property.
	if p.token != ast.KindQuestionToken {
		return false
	}
	// If any of the following tokens are after the question mark, it cannot
	// be a conditional expression, so treat it as an indexer.
	p.nextToken()
	return p.token == ast.KindColonToken || p.token == ast.KindCommaToken || p.token == ast.KindCloseBracketToken
}

func (p *Parser) parseIndexSignatureDeclaration(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	parameters := p.parseBracketedList(PCParameters, (*Parser).parseParameter, ast.KindOpenBracketToken, ast.KindCloseBracketToken)
	typeNode := p.parseTypeAnnotation()
	p.parseTypeMemberSemicolon()
	result := p.finishNode(p.factory.NewIndexSignatureDeclaration(modifiers, parameters, typeNode), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parsePropertyOrMethodSignature(pos int, jsdoc jsdocScannerInfo, modifiers *ast.ModifierList) *ast.Node {
	name := p.parsePropertyName()
	questionToken := p.parseOptionalToken(ast.KindQuestionToken)
	var result *ast.Node
	if p.token == ast.KindOpenParenToken || p.token == ast.KindLessThanToken {
		// Method signatures don't exist in expression contexts.  So they have neither
		// [Yield] nor [Await]
		typeParameters := p.parseTypeParameters()
		parameters := p.parseParameters(ParseFlagsType)
		returnType := p.parseCallableReturnType(ast.KindColonToken, true /*isType*/)
		result = p.factory.NewMethodSignatureDeclaration(modifiers, name, questionToken, typeParameters, parameters, returnType)
	} else {
		typeNode := p.parseTypeAnnotation()
		// Although type literal properties cannot not have initializers, we attempt
		// to parse an initializer so we can report in the checker that an interface
		// property or type literal property cannot have an initializer.
		var initializer *ast.Expression
		if p.token == ast.KindEqualsToken {
			initializer = p.parseInitializer()
		}
		result = p.factory.NewPropertySignatureDeclaration(modifiers, name, questionToken, typeNode, initializer)
	}
	p.parseTypeMemberSemicolon()
	p.withJSDoc(p.finishNode(result, pos), jsdoc)
	return result
}

func (p *Parser) parseTypeLiteral() *ast.Node {
	pos := p.nodePos()
	result := p.finishNode(p.factory.NewTypeLiteralNode(p.parseObjectTypeMembers()), pos)
	return result
}

func (p *Parser) parseObjectTypeMembers() *ast.NodeList {
	if p.parseExpected(ast.KindOpenBraceToken) {
		members := p.parseList(PCTypeMembers, (*Parser).parseTypeMember)
		p.parseExpected(ast.KindCloseBraceToken)
		return members
	}
	return p.parseEmptyNodeList()
}

func (p *Parser) parseTupleType() *ast.Node {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewTupleTypeNode(p.parseBracketedList(PCTupleElementTypes, (*Parser).parseTupleElementNameOrTupleElementType, ast.KindOpenBracketToken, ast.KindCloseBracketToken)), pos)
}

func (p *Parser) parseTupleElementNameOrTupleElementType() *ast.Node {
	if p.lookAhead((*Parser).scanStartOfNamedTupleElement) {
		pos := p.nodePos()
		jsdoc := p.jsdocScannerInfo()
		dotDotDotToken := p.parseOptionalToken(ast.KindDotDotDotToken)
		name := p.parseIdentifierName()
		questionToken := p.parseOptionalToken(ast.KindQuestionToken)
		p.parseExpected(ast.KindColonToken)
		typeNode := p.parseTupleElementType()
		result := p.finishNode(p.factory.NewNamedTupleMember(dotDotDotToken, name, questionToken, typeNode), pos)
		p.withJSDoc(result, jsdoc)
		return result
	}
	return p.parseTupleElementType()
}

func (p *Parser) scanStartOfNamedTupleElement() bool {
	if p.token == ast.KindDotDotDotToken {
		return tokenIsIdentifierOrKeyword(p.nextToken()) && p.nextTokenIsColonOrQuestionColon()
	}
	return tokenIsIdentifierOrKeyword(p.token) && p.nextTokenIsColonOrQuestionColon()
}

func (p *Parser) nextTokenIsColonOrQuestionColon() bool {
	return p.nextToken() == ast.KindColonToken || p.token == ast.KindQuestionToken && p.nextToken() == ast.KindColonToken
}

func (p *Parser) parseTupleElementType() *ast.TypeNode {
	pos := p.nodePos()
	if p.parseOptional(ast.KindDotDotDotToken) {
		return p.finishNode(p.factory.NewRestTypeNode(p.parseType()), pos)
	}
	typeNode := p.parseType()
	if ast.IsJSDocNullableType(typeNode) && typeNode.Pos() == typeNode.Type().Pos() {
		node := p.factory.NewOptionalTypeNode(typeNode.Type())
		node.Flags = typeNode.Flags
		node.Loc = typeNode.Loc
		typeNode.Type().Parent = node
		return node
	}
	return typeNode
}

func (p *Parser) parseParenthesizedType() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindOpenParenToken)
	typeNode := p.parseType()
	p.parseExpected(ast.KindCloseParenToken)
	return p.finishNode(p.factory.NewParenthesizedTypeNode(typeNode), pos)
}

func (p *Parser) parseAssertsTypePredicate() *ast.TypeNode {
	pos := p.nodePos()
	assertsModifier := p.parseExpectedToken(ast.KindAssertsKeyword)
	// `asserts this is T` is removed with the `this` type; only `asserts x is T`
	// over an ordinary identifier remains.
	parameterName := p.parseIdentifier()
	var typeNode *ast.TypeNode
	if p.parseOptional(ast.KindIsKeyword) {
		typeNode = p.parseType()
	}
	return p.finishNode(p.factory.NewTypePredicateNode(assertsModifier, parameterName, typeNode), pos)
}

func (p *Parser) parseTemplateType() *ast.Node {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewTemplateLiteralTypeNode(p.parseTemplateHead(false /*isTaggedTemplate*/), p.parseTemplateTypeSpans()), pos)
}

func (p *Parser) parseTemplateHead(isTaggedTemplate bool) *ast.Node {
	if !isTaggedTemplate && p.scanner.TokenFlags()&ast.TokenFlagsIsInvalid != 0 {
		p.reScanTemplateToken(false /*isTaggedTemplate*/)
	}
	pos := p.nodePos()
	result := p.factory.NewTemplateHead(p.scanner.TokenValue(), p.getTemplateLiteralRawText(2 /*endLength*/), p.scanner.TokenFlags())
	p.nextToken()
	return p.finishNode(result, pos)
}

func (p *Parser) getTemplateLiteralRawText(endLength int) string {
	tokenText := p.scanner.TokenText()
	if p.scanner.TokenFlags()&ast.TokenFlagsUnterminated != 0 {
		endLength = 0
	}
	return tokenText[1 : len(tokenText)-endLength]
}

func (p *Parser) parseTemplateTypeSpans() *ast.NodeList {
	pos := p.nodePos()
	var list []*ast.Node
	for {
		span := p.parseTemplateTypeSpan()
		list = append(list, span)
		if span.AsTemplateLiteralTypeSpan().Literal.Kind != ast.KindTemplateMiddle {
			break
		}
	}
	return p.newNodeList(core.NewTextRange(pos, p.nodePos()), list)
}

func (p *Parser) parseTemplateTypeSpan() *ast.Node {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewTemplateLiteralTypeSpan(p.parseType(), p.parseLiteralOfTemplateSpan(false /*isTaggedTemplate*/)), pos)
}

func (p *Parser) parseLiteralOfTemplateSpan(isTaggedTemplate bool) *ast.Node {
	if p.token == ast.KindCloseBraceToken {
		p.reScanTemplateToken(isTaggedTemplate)
		return p.parseTemplateMiddleOrTail()
	}
	p.parseErrorAtCurrentToken(diagnostics.X_0_expected, scanner.TokenToString(ast.KindCloseBraceToken))
	return p.finishNode(p.factory.NewTemplateTail("", "", ast.TokenFlagsNone), p.nodePos())
}

func (p *Parser) parseTemplateMiddleOrTail() *ast.Node {
	pos := p.nodePos()
	var result *ast.Node
	if p.token == ast.KindTemplateMiddle {
		result = p.factory.NewTemplateMiddle(p.scanner.TokenValue(), p.getTemplateLiteralRawText(2 /*endLength*/), p.scanner.TokenFlags())
	} else {
		result = p.factory.NewTemplateTail(p.scanner.TokenValue(), p.getTemplateLiteralRawText(1 /*endLength*/), p.scanner.TokenFlags())
	}
	p.nextToken()
	return p.finishNode(result, pos)
}

func (p *Parser) parseFunctionOrConstructorTypeToError(isInUnionType bool, parseConstituentType func(p *Parser) *ast.TypeNode) *ast.TypeNode {
	// the function type and constructor type shorthand notation
	// are not allowed directly in unions and intersections, but we'll
	// try to parse them gracefully and issue a helpful message.
	if p.isStartOfFunctionTypeOrConstructorType() {
		typeNode := p.parseFunctionOrConstructorType()
		var diagnostic *diagnostics.Message
		if typeNode.Kind == ast.KindFunctionType {
			diagnostic = core.IfElse(isInUnionType,
				diagnostics.Function_type_notation_must_be_parenthesized_when_used_in_a_union_type,
				diagnostics.Function_type_notation_must_be_parenthesized_when_used_in_an_intersection_type)
		} else {
			diagnostic = core.IfElse(isInUnionType,
				diagnostics.Constructor_type_notation_must_be_parenthesized_when_used_in_a_union_type,
				diagnostics.Constructor_type_notation_must_be_parenthesized_when_used_in_an_intersection_type)
		}
		p.parseErrorAtRange(typeNode.Loc, diagnostic)
		return typeNode
	}
	return parseConstituentType(p)
}

func (p *Parser) isStartOfFunctionTypeOrConstructorType() bool {
	return p.token == ast.KindLessThanToken ||
		p.token == ast.KindOpenParenToken && p.lookAhead((*Parser).nextIsUnambiguouslyStartOfFunctionType) ||
		p.token == ast.KindNewKeyword ||
		p.token == ast.KindAbstractKeyword && p.lookAhead((*Parser).nextTokenIsNewKeyword) ||
		// tlua async function type: `async (params) => T`. Only the parenthesized
		// form is unambiguous; `async <T>(...)` collides with the `async<T>` type
		// reference and is deferred, so `async` alone stays a type reference.
		p.token == ast.KindAsyncKeyword && p.lookAhead((*Parser).nextIsStartOfAsyncFunctionType)
}

func (p *Parser) parseFunctionOrConstructorType() *ast.TypeNode {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	modifiers := p.parseModifiersForConstructorType()
	isConstructorType := p.parseOptional(ast.KindNewKeyword)
	if !isConstructorType && modifiers == nil {
		// A function type may carry an `async` modifier (the tlua coroutine
		// contract); constructor types only take `abstract`, parsed above.
		modifiers = p.parseModifiersForFunctionType()
	}
	typeParameters := p.parseTypeParameters()
	parameters := p.parseParameters(ParseFlagsType)
	var returnType *ast.TypeNode
	if isConstructorType {
		returnType = p.parseReturnType(ast.KindEqualsGreaterThanToken, false /*isType*/)
	} else {
		returnType = p.parseCallableReturnType(ast.KindEqualsGreaterThanToken, false /*isType*/)
	}
	var result *ast.TypeNode
	if isConstructorType {
		result = p.factory.NewConstructorTypeNode(modifiers, typeParameters, parameters, returnType)
	} else {
		result = p.factory.NewFunctionTypeNode(modifiers, typeParameters, parameters, returnType)
	}
	p.finishNode(result, pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseModifiersForConstructorType() *ast.ModifierList {
	if p.token == ast.KindAbstractKeyword {
		pos := p.nodePos()
		modifier := p.factory.NewModifier(p.token)
		p.nextToken()
		p.finishNode(modifier, pos)
		return p.newModifierList(modifier.Loc, p.nodeSliceArena.NewSlice1(modifier))
	}
	return nil
}

func (p *Parser) nextTokenIsNewKeyword() bool {
	return p.nextToken() == ast.KindNewKeyword
}

// nextIsStartOfAsyncFunctionType reports whether `async` (already the current
// token) begins an async function type `async (params) => T`. Only the
// parenthesized form is accepted; a following `<` is left to the `async<T>`
// type-reference reading, and a line break after `async` rules it out.
func (p *Parser) nextIsStartOfAsyncFunctionType() bool {
	p.nextToken()
	if p.hasPrecedingLineBreak() || p.token != ast.KindOpenParenToken {
		return false
	}
	return p.nextIsUnambiguouslyStartOfFunctionType()
}

func (p *Parser) nextIsUnambiguouslyStartOfFunctionType() bool {
	p.nextToken()
	if p.token == ast.KindCloseParenToken || p.token == ast.KindDotDotDotToken {
		// ( )
		// ( ...
		return true
	}
	if p.skipParameterStart() {
		// We successfully skipped modifiers (if any) and an identifier or binding pattern,
		// now see if we have something that indicates a parameter declaration
		if p.token == ast.KindColonToken || p.token == ast.KindCommaToken || p.token == ast.KindQuestionToken || p.token == ast.KindEqualsToken {
			// ( xxx :
			// ( xxx ,
			// ( xxx ?
			// ( xxx =
			return true
		}
		if p.token == ast.KindCloseParenToken && p.nextToken() == ast.KindEqualsGreaterThanToken {
			// ( xxx ) =>
			return true
		}
	}
	return false
}

func (p *Parser) skipParameterStart() bool {
	if ast.IsModifierKind(p.token) {
		// Skip modifiers
		p.parseModifiers()
	}
	// A bare `...` needs no carve-out here: the sole caller,
	// nextIsUnambiguouslyStartOfFunctionType, already accepts a leading `...`
	// before it gets this far.
	p.parseOptional(ast.KindDotDotDotToken)
	if p.isIdentifier() {
		p.nextToken()
		return true
	}
	if p.token == ast.KindOpenBraceToken || p.isArrayBindingPatternStart() {
		// Return true if we can parse an object/array binding pattern with no errors
		previousErrorCount := len(p.diagnostics)
		p.parseIdentifierOrPattern()
		return previousErrorCount == len(p.diagnostics)
	}
	return false
}

func (p *Parser) parseModifiers() *ast.ModifierList {
	return p.parseModifiersEx(false /*stopOnStartOfClassStaticBlock*/)
}

func (p *Parser) parseModifiersEx(stopOnStartOfClassStaticBlock bool) *ast.ModifierList {
	var hasStaticModifier bool
	pos := p.nodePos()
	list := make([]*ast.Node, 0, 16)
	for {
		modifier := p.tryParseModifier(hasStaticModifier, stopOnStartOfClassStaticBlock)
		if modifier == nil {
			break
		}
		if modifier.Kind == ast.KindStaticKeyword {
			hasStaticModifier = true
		}
		list = append(list, modifier)
	}
	if len(list) != 0 {
		return p.newModifierList(core.NewTextRange(pos, p.nodePos()), p.nodeSliceArena.Clone(list))
	}
	return nil
}

func (p *Parser) tryParseModifier(hasSeenStaticModifier bool, stopOnStartOfClassStaticBlock bool) *ast.Node {
	pos := p.nodePos()
	kind := p.token
	if stopOnStartOfClassStaticBlock && p.token == ast.KindStaticKeyword && p.lookAhead((*Parser).nextTokenIsOpenBrace) {
		return nil
	} else if hasSeenStaticModifier && p.token == ast.KindStaticKeyword {
		return nil
	} else {
		if !p.parseAnyContextualModifier() {
			return nil
		}
	}
	return p.finishNode(p.factory.NewModifier(kind), pos)
}

func (p *Parser) parseContextualModifier(t ast.Kind) bool {
	state := p.mark()
	if p.token == t && p.nextTokenCanFollowModifier() {
		return true
	}
	p.rewind(state)
	return false
}

func (p *Parser) parseAnyContextualModifier() bool {
	state := p.mark()
	if ast.IsModifierKind(p.token) && p.nextTokenCanFollowModifier() {
		return true
	}
	p.rewind(state)
	return false
}

func (p *Parser) nextTokenCanFollowModifier() bool {
	switch p.token {
	case ast.KindStaticKeyword:
		p.nextToken()
		return p.canFollowModifier()
	case ast.KindGetKeyword, ast.KindSetKeyword:
		p.nextToken()
		return p.canFollowGetOrSetKeyword()
	default:
		return p.nextTokenIsOnSameLineAndCanFollowModifier()
	}
}

func (p *Parser) nextTokenIsIdentifierOrKeyword() bool {
	return tokenIsIdentifierOrKeyword(p.nextToken())
}

// nextTokenIsIdentifierOrKeywordOrGreaterThan decides whether a `<` opens a JSX
// element, so it must admit every JSX tag name — including a word-spelled
// operator, which the scan has already resolved to its punctuation kind.
func (p *Parser) nextTokenIsIdentifierOrKeywordOrGreaterThan() bool {
	p.nextToken()
	return p.token == ast.KindGreaterThanToken || p.tokenIsJsxName()
}

func (p *Parser) nextTokenIsIdentifierOrKeywordOnSameLine() bool {
	return p.nextTokenIsIdentifierOrKeyword() && !p.hasPrecedingLineBreak()
}

func (p *Parser) nextTokenIsIdentifierOrKeywordOrLiteralOnSameLine() bool {
	return (p.nextTokenIsIdentifierOrKeyword() || p.token == ast.KindNumericLiteral || p.token == ast.KindStringLiteral) && !p.hasPrecedingLineBreak()
}

func (p *Parser) nextTokenIsClassKeywordOnSameLine() bool {
	return p.nextToken() == ast.KindClassKeyword && !p.hasPrecedingLineBreak()
}

func (p *Parser) nextTokenIsFunctionKeywordOnSameLine() bool {
	return p.nextToken() == ast.KindFunctionKeyword && !p.hasPrecedingLineBreak()
}

func (p *Parser) canFollowModifier() bool {
	// Generators are removed in tlua: `*` cannot follow a modifier (`async *m()`).
	return p.token == ast.KindOpenBracketToken || p.token == ast.KindOpenBraceToken || p.token == ast.KindDotDotDotToken || p.isLiteralPropertyName()
}

func (p *Parser) canFollowGetOrSetKeyword() bool {
	return p.token == ast.KindOpenBracketToken || p.isLiteralPropertyName()
}

func (p *Parser) nextTokenIsOnSameLineAndCanFollowModifier() bool {
	p.nextToken()
	if p.hasPrecedingLineBreak() {
		return false
	}
	return p.canFollowModifier()
}

func (p *Parser) nextTokenIsOpenBrace() bool {
	return p.nextToken() == ast.KindOpenBraceToken
}

func (p *Parser) parseExpression() *ast.Expression {
	// Expression[in]:
	//      AssignmentExpression[in]
	//      Expression[in] , AssignmentExpression[in]

	pos := p.nodePos()
	expr := p.parseAssignmentExpressionOrHigher()
	for {
		operatorToken := p.parseOptionalToken(ast.KindCommaToken)
		if operatorToken == nil {
			break
		}
		expr = p.makeBinaryExpression(expr, operatorToken, p.parseAssignmentExpressionOrHigher(), pos)
	}
	return expr
}

func (p *Parser) parseExpressionAllowIn() *ast.Expression {
	return doInContext(p, ast.NodeFlagsDisallowInContext, false, (*Parser).parseExpression)
}

// parseCommaValueList parses a Lua value list: one or more comma-separated
// assignment expressions. A single value stays a plain expression; two or
// more become an ExpressionList. Lua has no comma operator, so a bare comma
// in a value-list position always separates values (a parenthesized comma
// expression still parses as before).
func (p *Parser) parseCommaValueList() *ast.Expression {
	pos := p.nodePos()
	expr := p.parseAssignmentExpressionOrHigher()
	if p.token == ast.KindCommaToken {
		elements := []*ast.Expression{expr}
		for p.parseOptional(ast.KindCommaToken) {
			elements = append(elements, p.parseAssignmentExpressionOrHigher())
		}
		list := p.newNodeList(core.NewTextRange(elements[0].Pos(), elements[len(elements)-1].End()), elements)
		expr = p.finishNode(p.factory.NewExpressionList(list), pos)
	}
	return expr
}

// parseReturnExpressionList parses the value list of a return statement:
//
//	return a, b
func (p *Parser) parseReturnExpressionList() *ast.Expression {
	// Allow `in`, mirroring parseExpression / parseExpressionAllowIn for the old
	// single-operand parse.
	saveContextFlags := p.contextFlags
	p.contextFlags &^= ast.NodeFlagsDisallowInContext
	expr := p.parseCommaValueList()
	p.contextFlags = saveContextFlags
	return expr
}

func (p *Parser) parseAssignmentExpressionOrHigher() *ast.Expression {
	//  AssignmentExpression[in,yield]:
	//      1) ConditionalExpression[?in,?yield]
	//      2) LeftHandSideExpression = AssignmentExpression[?in,?yield]
	//      3) LeftHandSideExpression AssignmentOperator AssignmentExpression[?in,?yield]
	//      4) ArrowFunctionExpression[?in,?yield]
	//      5) AsyncArrowFunctionExpression[in,yield,await]
	//      6) [+Yield] YieldExpression[?In]
	//
	// Note: for ease of implementation we treat productions '2' and '3' as the same thing.
	// (i.e. they're both BinaryExpressions with an assignment operator in it).
	// Generators are removed in tlua: `yield` is an ordinary identifier, never a
	// YieldExpression (production '6').
	// tlua has no arrow function expressions (productions '4' and '5'); an
	// anonymous function value is written `function(...) ... end`. So an
	// assignment expression reduces to productions '1'-'3'.
	//
	// We parse out a BinaryExpression. If we get something that is a LeftHandSide
	// or higher, we then try to parse out the assignment expression part. We pass
	// in the 'lowest' precedence so that it matches and consumes anything.
	pos := p.nodePos()
	expr := p.parseBinaryExpressionOrHigher(ast.OperatorPrecedenceLowest)
	// Now see if we might be in cases '2' or '3'.
	// If the expression was a LHS expression, and we have an assignment operator, then
	// we're in '2' or '3'. Consume the assignment and return.
	//
	// Note: we call reScanGreaterToken so that we get an appropriately merged token
	// for cases like `> > =` becoming `>>=`
	if ast.IsLeftHandSideExpression(expr) && ast.IsAssignmentOperator(p.reScanGreaterThanToken()) {
		return p.makeBinaryExpression(expr, p.parseTokenNode(), p.parseAssignmentExpressionOrHigher(), pos)
	}
	// It wasn't an assignment or a lambda. tlua has no conditional expression:
	// a stray ternary `?` is left to stock recovery.
	return expr
}

// parseModifiersForFunctionType parses the leading modifiers of a function
// type. In tlua the only such modifier is `async` (`async (…) => T`); arrow
// function values are removed, so this no longer feeds a value form.
func (p *Parser) parseModifiersForFunctionType() *ast.ModifierList {
	if p.token == ast.KindAsyncKeyword {
		pos := p.nodePos()
		p.nextToken()
		modifier := p.finishNode(p.factory.NewModifier(ast.KindAsyncKeyword), pos)
		return p.newModifierList(modifier.Loc, p.nodeSliceArena.NewSlice1(modifier))
	}
	return nil
}

func (p *Parser) parseBinaryExpressionOrHigher(precedence ast.OperatorPrecedence) *ast.Expression {
	pos := p.nodePos()
	leftOperand := p.parseUnaryExpressionOrHigher()
	return p.parseBinaryExpressionRest(precedence, leftOperand, pos)
}

func (p *Parser) parseBinaryExpressionRest(precedence ast.OperatorPrecedence, leftOperand *ast.Expression, pos int) *ast.Expression {
	lastOperand := leftOperand
	for {
		// We either have a binary operator here, or we're finished.  We call
		// reScanGreaterToken so that we merge token sequences like > and = into >=
		operator := p.reScanGreaterThanToken()
		newPrecedence := ast.GetBinaryOperatorPrecedence(operator)
		// Check the precedence to see if we should "take" this operator
		// - For left associative operators (all operators but `^` and `..`), consume the operator,
		//   recursively call the function below, and parse binaryExpression as a rightOperand
		//   of the caller if the new precedence of the operator is greater then or equal to the current precedence.
		//   For example:
		//      a - b - c;
		//            ^token; leftOperand = b. Return b to the caller as a rightOperand
		//      a * b - c
		//            ^token; leftOperand = b. Return b to the caller as a rightOperand
		//      a - b * c;
		//            ^token; leftOperand = b. Return b * c to the caller as a rightOperand
		// - For right associative operators (`^` and `..`), consume the operator, recursively call the function
		//   and parse binaryExpression as a rightOperand of the caller if the new precedence of
		//   the operator is strictly grater than the current precedence
		//   For example:
		//      a ^ b ^ c;
		//            ^token; leftOperand = b. Return b ^ c to the caller as a rightOperand
		//      a - b ^ c;
		//            ^token; leftOperand = b. Return b ^ c to the caller as a rightOperand
		//      a ^ b - c
		//             ^token; leftOperand = b. Return b to the caller as a rightOperand
		var consumeCurrentOperator bool
		if operator == ast.KindAsteriskAsteriskToken || operator == ast.KindDotDotToken {
			// Lua's exponentiation and concatenation operators are right-associative.
			consumeCurrentOperator = newPrecedence >= precedence
		} else {
			consumeCurrentOperator = newPrecedence > precedence
		}
		if !consumeCurrentOperator {
			break
		}
		if operator == ast.KindInKeyword && p.inDisallowInContext() {
			break
		}
		if operator == ast.KindAsKeyword || operator == ast.KindSatisfiesKeyword {
			// Make sure we *do* perform ASI for constructs like this:
			//    var x = foo
			//    as (Bar)
			// This should be parsed as an initialized variable, followed
			// by a function call to 'as' with the argument 'Bar'
			if p.hasPrecedingLineBreak() {
				break
			} else {
				p.nextToken()
				// When we have 'a ## b as SomeType' or 'a ## b satisfies SomeType', where ## is some binary
				// operator, we want to stop parsing on any following operator with a higher precedence than ##
				// because continuing would make it impossible to erase the `as` or `satisfies` without changing
				// the meaning of the expression. See https://github.com/microsoft/TypeScript/issues/63527.
				lastPrecedence := ast.OperatorPrecedenceHighest
				if ast.IsBinaryExpression(lastOperand) {
					lastPrecedence = ast.GetBinaryOperatorPrecedence(lastOperand.AsBinaryExpression().OperatorToken.Kind)
				}
				if operator == ast.KindSatisfiesKeyword {
					leftOperand = p.makeSatisfiesExpression(leftOperand, p.parseType())
				} else {
					leftOperand = p.makeAsExpression(leftOperand, p.parseType())
				}
				// Stop if the precedence of the next operator is too high.
				if ast.GetBinaryOperatorPrecedence(p.reScanGreaterThanToken()) > lastPrecedence {
					break
				}
			}
		} else {
			leftOperand = p.makeBinaryExpression(leftOperand, p.parseTokenNode(), p.parseBinaryExpressionOrHigher(newPrecedence), pos)
			lastOperand = leftOperand
		}
	}
	return leftOperand
}

func (p *Parser) makeSatisfiesExpression(expression *ast.Expression, typeNode *ast.TypeNode) *ast.Node {
	return p.checkJSSyntax(p.finishNode(p.factory.NewSatisfiesExpression(expression, typeNode), expression.Pos()))
}

func (p *Parser) makeAsExpression(left *ast.Expression, right *ast.TypeNode) *ast.Node {
	return p.checkJSSyntax(p.finishNode(p.factory.NewAsExpression(left, right), left.Pos()))
}

func (p *Parser) makeBinaryExpression(left *ast.Expression, operatorToken *ast.Node, right *ast.Expression, pos int) *ast.Node {
	return p.finishNode(p.factory.NewBinaryExpression(nil /*modifiers*/, left, nil /*typeNode*/, operatorToken, right), pos)
}

func (p *Parser) parseUnaryExpressionOrHigher() *ast.Expression {
	if p.isUpdateExpression() {
		pos := p.nodePos()
		updateExpression := p.parseUpdateExpression()
		if p.token == ast.KindAsteriskAsteriskToken {
			return p.parseBinaryExpressionRest(ast.GetBinaryOperatorPrecedence(p.token), updateExpression, pos)
		}
		return updateExpression
	}
	return p.parseSimpleUnaryExpression()
}

// The Lua vararg expression. Legal only inside a vararg function or the main
// chunk; checkVarargExpression reports 100035 otherwise.
func (p *Parser) parseVarargExpression() *ast.Expression {
	pos := p.nodePos()
	p.parseExpected(ast.KindDotDotDotToken)
	return p.finishNode(p.factory.NewVarargExpression(), pos)
}

func (p *Parser) isUpdateExpression() bool {
	switch p.token {
	case ast.KindPlusToken, ast.KindMinusToken, ast.KindExclamationToken, ast.KindHashToken, ast.KindDeleteKeyword, ast.KindVoidKeyword:
		return false
	case ast.KindDotDotDotToken:
		// The Lua vararg is not a prefixexp: it cannot be called, indexed, or
		// member-accessed. Declining it here routes it through
		// parseSimpleUnaryExpression instead of parseUpdateExpression, which is
		// what denies it left-hand-side suffixes -- and it reaches that route
		// both bare (`local a = ...`) and as the operand of a prefix unary
		// operator (`not ...`, `-...`), which Lua allows.
		return false
	case ast.KindLessThanToken:
		return p.languageVariant == core.LanguageVariantJSX
	}
	return true
}

func (p *Parser) parseUpdateExpression() *ast.Expression {
	if p.languageVariant == core.LanguageVariantJSX && p.token == ast.KindLessThanToken && p.lookAhead((*Parser).nextTokenIsIdentifierOrKeywordOrGreaterThan) {
		// JSXElement is part of primaryExpression
		return p.parseJsxElementOrSelfClosingElementOrFragment(true /*inExpressionContext*/, -1 /*topInvalidNodePosition*/, nil /*openingTag*/, false /*mustBeUnary*/)
	}
	return p.parseLeftHandSideExpressionOrHigher()
}

func (p *Parser) parseJsxElementOrSelfClosingElementOrFragment(inExpressionContext bool, topInvalidNodePosition int, openingTag *ast.Node, mustBeUnary bool) *ast.Expression {
	pos := p.nodePos()
	opening := p.parseJsxOpeningOrSelfClosingElementOrOpeningFragment(inExpressionContext)
	var result *ast.Expression
	switch opening.Kind {
	case ast.KindJsxOpeningElement:
		children := p.parseJsxChildren(opening)
		var closingElement *ast.Node
		lastChild := core.LastOrNil(children.Nodes)
		if lastChild != nil && lastChild.Kind == ast.KindJsxElement &&
			!ast.TagNamesAreEquivalent(lastChild.AsJsxElement().OpeningElement.TagName(), lastChild.AsJsxElement().ClosingElement.TagName()) &&
			ast.TagNamesAreEquivalent(opening.TagName(), lastChild.AsJsxElement().ClosingElement.TagName()) {
			// when an unclosed JsxOpeningElement incorrectly parses its parent's JsxClosingElement,
			// restructure (<div>(...<span>...</div>)) --> (<div>(...<span>...</>)</div>)
			// (no need to error; the parent will error)
			end := lastChild.Children().End()
			missingIdentifier := p.finishNodeWithEnd(p.newIdentifier(""), end, end)
			newClosingElement := p.finishNodeWithEnd(p.factory.NewJsxClosingElement(missingIdentifier), end, end)
			newLast := p.finishNodeWithEnd(
				p.factory.NewJsxElement(lastChild.AsJsxElement().OpeningElement, lastChild.Children(), newClosingElement),
				lastChild.AsJsxElement().OpeningElement.Pos(),
				end,
			)
			// force reset parent pointers from discarded parse result
			if lastChild.AsJsxElement().OpeningElement != nil {
				lastChild.AsJsxElement().OpeningElement.Parent = newLast
			}
			if lastChild.Children() != nil {
				for _, c := range lastChild.Children().Nodes {
					c.Parent = newLast
				}
			}
			newClosingElement.Parent = newLast
			children = p.newNodeList(core.NewTextRange(children.Pos(), newLast.End()), append(children.Nodes[0:len(children.Nodes)-1], newLast))
			closingElement = lastChild.AsJsxElement().ClosingElement
		} else {
			closingElement = p.parseJsxClosingElement(opening, inExpressionContext)
			if !ast.TagNamesAreEquivalent(opening.TagName(), closingElement.TagName()) {
				if openingTag != nil && ast.IsJsxOpeningElement(openingTag) && ast.TagNamesAreEquivalent(closingElement.TagName(), openingTag.TagName()) {
					// opening incorrectly matched with its parent's closing -- put error on opening
					p.parseErrorAtRange(opening.TagName().Loc, diagnostics.JSX_element_0_has_no_corresponding_closing_tag, scanner.GetTextOfNodeFromSourceText(p.sourceText, opening.TagName(), false /*includeTrivia*/))
				} else {
					// other opening/closing mismatches -- put error on closing
					p.parseErrorAtRange(closingElement.TagName().Loc, diagnostics.Expected_corresponding_JSX_closing_tag_for_0, scanner.GetTextOfNodeFromSourceText(p.sourceText, opening.TagName(), false /*includeTrivia*/))
				}
			}
		}
		result = p.finishNode(p.factory.NewJsxElement(opening, children, closingElement), pos)
		closingElement.Parent = result // force reset parent pointers from possibly discarded parse result
	case ast.KindJsxOpeningFragment:
		result = p.finishNode(p.factory.NewJsxFragment(opening, p.parseJsxChildren(opening), p.parseJsxClosingFragment(inExpressionContext)), pos)
	case ast.KindJsxSelfClosingElement:
		// Nothing else to do for self-closing elements
		result = opening
	default:
		panic("Unhandled case in parseJsxElementOrSelfClosingElementOrFragment")
	}
	// If the user writes the invalid code '<div></div><div></div>' in an expression context (i.e. not wrapped in
	// an enclosing tag), we'll naively try to parse   ^ this as a 'less than' operator and the remainder of the tag
	// as garbage, which will cause the formatter to badly mangle the JSX. Perform a speculative parse of a JSX
	// element if we see a < token so that we can wrap it in a synthetic binary expression so the formatter
	// does less damage and we can report a better error.
	// Since JSX elements are invalid < operands anyway, this lookahead parse will only occur in error scenarios
	// of one sort or another.
	// If we are in a unary context, we can't do this recovery; the binary expression we return here is not
	// a valid UnaryExpression and will cause problems later.
	if !mustBeUnary && inExpressionContext && p.token == ast.KindLessThanToken {
		topBadPos := topInvalidNodePosition
		if topBadPos < 0 {
			topBadPos = result.Pos()
		}
		invalidElement := p.parseJsxElementOrSelfClosingElementOrFragment( /*inExpressionContext*/ true, topBadPos, nil, false)
		operatorToken := p.factory.NewToken(ast.KindCommaToken)
		operatorToken.Loc = core.NewTextRange(invalidElement.Pos(), invalidElement.Pos())
		p.parseErrorAt(scanner.SkipTrivia(p.sourceText, topBadPos), invalidElement.End(), diagnostics.JSX_expressions_must_have_one_parent_element)
		result = p.finishNode(p.factory.NewBinaryExpression(nil /*modifiers*/, result, nil /*typeNode*/, operatorToken, invalidElement), pos)
	}
	return result
}

func (p *Parser) parseJsxChildren(openingTag *ast.Expression) *ast.NodeList {
	pos := p.nodePos()
	saveParsingContexts := p.parsingContexts
	p.parsingContexts |= 1 << PCJsxChildren
	var list []*ast.Node
	for {
		currentToken := p.scanner.ReScanJsxToken(true /*allowMultilineJsxText*/)
		child := p.parseJsxChild(openingTag, currentToken)
		if child == nil {
			break
		}
		list = append(list, child)
		if ast.IsJsxOpeningElement(openingTag) && child.Kind == ast.KindJsxElement &&
			!ast.TagNamesAreEquivalent(child.AsJsxElement().OpeningElement.TagName(), child.AsJsxElement().ClosingElement.TagName()) &&
			ast.TagNamesAreEquivalent(openingTag.TagName(), child.AsJsxElement().ClosingElement.TagName()) {
			// stop after parsing a mismatched child like <div>...(<span></div>) in order to reattach the </div> higher
			break
		}
	}
	p.parsingContexts = saveParsingContexts
	return p.newNodeList(core.NewTextRange(pos, p.nodePos()), list)
}

func (p *Parser) parseJsxChild(openingTag *ast.Node, token ast.Kind) *ast.Expression {
	switch token {
	case ast.KindEndOfFile:
		// If we hit EOF, issue the error at the tag that lacks the closing element
		// rather than at the end of the file (which is useless)
		if ast.IsJsxOpeningFragment(openingTag) {
			p.parseErrorAtRange(openingTag.Loc, diagnostics.JSX_fragment_has_no_corresponding_closing_tag)
		} else {
			// We want the error span to cover only 'Foo.Bar' in < Foo.Bar >
			// or to cover only 'Foo' in < Foo >
			tag := openingTag.TagName()
			start := min(scanner.SkipTrivia(p.sourceText, tag.Pos()), tag.End())
			p.parseErrorAt(start, tag.End(), diagnostics.JSX_element_0_has_no_corresponding_closing_tag,
				scanner.GetTextOfNodeFromSourceText(p.sourceText, openingTag.TagName(), false /*includeTrivia*/))
		}
		return nil
	case ast.KindLessThanSlashToken, ast.KindConflictMarkerTrivia:
		return nil
	case ast.KindJsxText, ast.KindJsxTextAllWhiteSpaces:
		return p.parseJsxText()
	case ast.KindOpenBraceToken:
		return p.parseJsxExpression(false /*inExpressionContext*/)
	case ast.KindLessThanToken:
		return p.parseJsxElementOrSelfClosingElementOrFragment(false /*inExpressionContext*/, -1 /*topInvalidNodePosition*/, openingTag, false)
	}
	panic("Unhandled case in parseJsxChild")
}

func (p *Parser) parseJsxText() *ast.Node {
	pos := p.nodePos()
	result := p.factory.NewJsxText(p.scanner.TokenValue(), p.token == ast.KindJsxTextAllWhiteSpaces)
	p.scanJsxText()
	return p.finishNode(result, pos)
}

func (p *Parser) parseJsxExpression(inExpressionContext bool) *ast.Node {
	pos := p.nodePos()
	if !p.parseExpected(ast.KindOpenBraceToken) {
		return nil
	}
	// No children spread: `<div>{...items}</div>` is gone. The token is left for
	// ordinary expression parsing, where `...` is the Lua vararg.
	var dotDotDotToken *ast.Node
	var expression *ast.Expression
	if p.token != ast.KindCloseBraceToken {
		// Only an AssignmentExpression is valid here per the JSX spec,
		// but we can unambiguously parse a comma sequence and provide
		// a better error message in grammar checking.
		expression = p.parseExpression()
	}
	if inExpressionContext {
		p.parseExpected(ast.KindCloseBraceToken)
	} else if p.parseExpectedWithoutAdvancing(ast.KindCloseBraceToken) {
		p.scanJsxText()
	}
	return p.finishNode(p.factory.NewJsxExpression(dotDotDotToken, expression), pos)
}

func (p *Parser) scanJsxText() ast.Kind {
	p.token = p.scanner.ScanJsxToken()
	return p.token
}

func (p *Parser) scanJsxIdentifier() ast.Kind {
	p.token = p.scanner.ScanJsxIdentifier()
	return p.token
}

func (p *Parser) scanJsxAttributeValue() ast.Kind {
	p.token = p.scanner.ScanJsxAttributeValue()
	return p.token
}

func (p *Parser) parseJsxClosingElement(open *ast.Node, inExpressionContext bool) *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindLessThanSlashToken)
	tagName := p.parseJsxElementName()
	if p.parseExpectedWithDiagnostic(ast.KindGreaterThanToken, nil /*diagnosticMessage*/, false /*shouldAdvance*/) {
		// manually advance the scanner in order to look for jsx text inside jsx
		if inExpressionContext || !ast.TagNamesAreEquivalent(open.TagName(), tagName) {
			p.nextToken()
		} else {
			p.scanJsxText()
		}
	}
	return p.finishNode(p.factory.NewJsxClosingElement(tagName), pos)
}

func (p *Parser) parseJsxOpeningOrSelfClosingElementOrOpeningFragment(inExpressionContext bool) *ast.Expression {
	pos := p.nodePos()
	p.parseExpected(ast.KindLessThanToken)
	if p.token == ast.KindGreaterThanToken {
		// See below for explanation of scanJsxText
		p.scanJsxText()
		return p.finishNode(p.factory.NewJsxOpeningFragment(), pos)
	}
	tagName := p.parseJsxElementName()
	var typeArguments *ast.NodeList
	if p.contextFlags&ast.NodeFlagsJavaScriptFile == 0 {
		typeArguments = p.parseTypeArguments()
	}
	attributes := p.parseJsxAttributes()
	var result *ast.Expression
	if p.token == ast.KindGreaterThanToken {
		// Closing tag, so scan the immediately-following text with the JSX scanning instead
		// of regular scanning to avoid treating illegal characters (e.g. '#') as immediate
		// scanning errors
		p.scanJsxText()
		result = p.factory.NewJsxOpeningElement(tagName, typeArguments, attributes)
	} else {
		p.parseExpected(ast.KindSlashToken)
		if p.parseExpectedWithoutAdvancing(ast.KindGreaterThanToken) {
			if inExpressionContext {
				p.nextToken()
			} else {
				p.scanJsxText()
			}
		}
		result = p.factory.NewJsxSelfClosingElement(tagName, typeArguments, attributes)
	}
	return p.finishNode(result, pos)
}

func (p *Parser) parseJsxElementName() *ast.Expression {
	pos := p.nodePos()
	// JsxElement can have name in the form of
	//      propertyAccessExpression
	//      primaryExpression in the form of an identifier and "this" keyword
	// We can't just simply use parseLeftHandSideExpressionOrHigher because then we will start consider class,function etc as a keyword
	// We only want to consider "this" as a primaryExpression
	initialExpression := p.parseJsxTagName()
	if ast.IsJsxNamespacedName(initialExpression) {
		return initialExpression // `a:b.c` is invalid syntax, don't even look for the `.` if we parse `a:b`, and let `parseAttribute` report "unexpected :" instead.
	}
	expression := initialExpression
	for p.parseOptional(ast.KindDotToken) {
		expression = p.finishNode(p.factory.NewPropertyAccessExpression(expression, nil, nil /*colonToken*/, p.parseRightSideOfDot(true /*allowIdentifierNames*/, false /*allowPrivateIdentifiers*/, false /*allowUnicodeEscapeSequenceInIdentifierName*/), ast.NodeFlagsNone), pos)
	}
	return expression
}

func (p *Parser) parseJsxTagName() *ast.Expression {
	pos := p.nodePos()
	p.scanJsxIdentifier()
	// `this` is removed from tlua, so `<this.Foo/>` no longer produces a
	// this-expression tag name; the keyword is consumed as an ordinary
	// identifier name and resolves (and errors) like any other unknown tag.
	tagName := p.parseIdentifierNameErrorOnUnicodeEscapeSequence()
	if p.parseOptional(ast.KindColonToken) {
		p.scanJsxIdentifier()
		return p.finishNode(p.factory.NewJsxNamespacedName(tagName, p.parseIdentifierNameErrorOnUnicodeEscapeSequence()), pos)
	}
	return tagName
}

func (p *Parser) parseJsxAttributes() *ast.Node {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewJsxAttributes(p.parseList(PCJsxAttributes, (*Parser).parseJsxAttribute)), pos)
}

func (p *Parser) parseJsxAttribute() *ast.Node {
	// No spread attribute: `<Foo {...props} />` is gone with every other spread
	// form. A `{` in attribute position is now a parse error.
	pos := p.nodePos()
	return p.finishNode(p.factory.NewJsxAttribute(p.parseJsxAttributeName(), p.parseJsxAttributeValue()), pos)
}

func (p *Parser) parseJsxAttributeName() *ast.Node {
	pos := p.nodePos()
	p.scanJsxIdentifier()
	attrName := p.parseIdentifierNameErrorOnUnicodeEscapeSequence()
	if p.parseOptional(ast.KindColonToken) {
		p.scanJsxIdentifier()
		return p.finishNode(p.factory.NewJsxNamespacedName(attrName, p.parseIdentifierNameErrorOnUnicodeEscapeSequence()), pos)
	}
	return attrName
}

func (p *Parser) parseJsxAttributeValue() *ast.Expression {
	if p.token == ast.KindEqualsToken {
		if p.scanJsxAttributeValue() == ast.KindStringLiteral {
			return p.parseLiteralExpression(false /*intern*/)
		}
		if p.token == ast.KindOpenBraceToken {
			return p.parseJsxExpression( /*inExpressionContext*/ true)
		}
		if p.token == ast.KindLessThanToken {
			return p.parseJsxElementOrSelfClosingElementOrFragment(true /*inExpressionContext*/, -1, nil, false)
		}
		p.parseErrorAtCurrentToken(diagnostics.X_or_JSX_element_expected)
	}
	return nil
}

func (p *Parser) parseJsxClosingFragment(inExpressionContext bool) *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindLessThanSlashToken)
	if p.parseExpectedWithDiagnostic(ast.KindGreaterThanToken, diagnostics.Expected_corresponding_closing_tag_for_JSX_fragment, false /*shouldAdvance*/) {
		// manually advance the scanner in order to look for jsx text inside jsx
		if inExpressionContext {
			p.nextToken()
		} else {
			p.scanJsxText()
		}
	}
	return p.finishNode(p.factory.NewJsxClosingFragment(), pos)
}

func (p *Parser) parseSimpleUnaryExpression() *ast.Expression {
	switch p.token {
	case ast.KindDotDotDotToken:
		// See isUpdateExpression: the vararg lands here, below the left-hand-side
		// suffix machinery, whether it is bare or the operand of a prefix unary.
		return p.parseVarargExpression()
	case ast.KindPlusToken, ast.KindMinusToken, ast.KindExclamationToken, ast.KindHashToken:
		return p.parsePrefixUnaryExpression()
	case ast.KindDeleteKeyword:
		return p.parseDeleteExpression()
	case ast.KindVoidKeyword:
		return p.parseVoidExpression()
	case ast.KindLessThanToken:
		// Just like in parseUpdateExpression, we need to avoid parsing type assertions when
		// in JSX and we see an expression like "+ <foo> bar".
		if p.languageVariant == core.LanguageVariantJSX {
			return p.parseJsxElementOrSelfClosingElementOrFragment(true /*inExpressionContext*/, -1 /*topInvalidNodePosition*/, nil /*openingTag*/, true /*mustBeUnary*/)
		}
		// // This is modified UnaryExpression grammar in TypeScript
		// //  UnaryExpression (modified):
		// //      < type > UnaryExpression
		return p.parseTypeAssertion()
	default:
		// `await` is an ordinary identifier in tlua: no AwaitExpression.
		return p.parseUpdateExpression()
	}
}

func (p *Parser) parsePrefixUnaryExpression() *ast.Node {
	pos := p.nodePos()
	operator := p.token
	p.nextToken()
	// Lua exponentiation binds tighter than prefix unary operators. Starting the
	// operand at unary precedence admits `^`, but stops before multiplicative operators.
	operand := p.parseBinaryExpressionOrHigher(ast.OperatorPrecedenceUnary)
	return p.finishNode(p.factory.NewPrefixUnaryExpression(operator, operand), pos)
}

func (p *Parser) parseDeleteExpression() *ast.Node {
	pos := p.nodePos()
	p.nextToken()
	return p.finishNode(p.factory.NewDeleteExpression(p.parseSimpleUnaryExpression()), pos)
}

func (p *Parser) parseVoidExpression() *ast.Node {
	pos := p.nodePos()
	p.nextToken()
	return p.finishNode(p.factory.NewVoidExpression(p.parseSimpleUnaryExpression()), pos)
}

func (p *Parser) parseTypeAssertion() *ast.Node {
	debug.Assert(p.languageVariant != core.LanguageVariantJSX, "Type assertions should never be parsed in JSX; they should be parsed as comparisons or JSX elements/fragments.")
	pos := p.nodePos()
	p.parseExpected(ast.KindLessThanToken)
	typeNode := p.parseType()
	p.parseExpected(ast.KindGreaterThanToken)
	expression := p.parseSimpleUnaryExpression()
	return p.finishNode(p.factory.NewTypeAssertion(typeNode, expression), pos)
}

func (p *Parser) parseLeftHandSideExpressionOrHigher() *ast.Expression {
	// Original Ecma:
	// LeftHandSideExpression: See 11.2
	//      NewExpression
	//      CallExpression
	//
	// Our simplification:
	//
	// LeftHandSideExpression: See 11.2
	//      MemberExpression
	//      CallExpression
	//
	// See comment in parseMemberExpressionOrHigher on how we replaced NewExpression with
	// MemberExpression to make our lives easier.
	//
	// to best understand the below code, it's important to see how CallExpression expands
	// out into its own productions:
	//
	// CallExpression:
	//      MemberExpression Arguments
	//      CallExpression Arguments
	//      CallExpression[Expression]
	//      CallExpression.IdentifierName
	//      import (AssignmentExpression)
	//      super Arguments
	//      super.IdentifierName
	//
	// Because of the recursion in these calls, we need to bottom out first. There are three
	// bottom out states we can run into: 1) We see 'super' which must start either of
	// the last two CallExpression productions. 2) We see 'import' which must start import call.
	// 3)we have a MemberExpression which either completes the LeftHandSideExpression,
	// or starts the beginning of the first four CallExpression productions.
	pos := p.nodePos()
	var expression *ast.Expression
	if p.token == ast.KindSuperKeyword {
		expression = p.parseSuperExpression()
	} else {
		expression = p.parseMemberExpressionOrHigher()
	}
	// Now, we *may* be complete.  However, we might have consumed the start of a
	// CallExpression or OptionalExpression.  As such, we need to consume the rest
	// of it here to be complete.
	return p.parseCallExpressionRest(pos, expression)
}

func (p *Parser) nextTokenIsDot() bool {
	return p.nextToken() == ast.KindDotToken
}

func (p *Parser) parseSuperExpression() *ast.Expression {
	pos := p.nodePos()
	expression := p.parseKeywordExpression()
	if p.token == ast.KindLessThanToken {
		startPos := p.nodePos()
		typeArguments := p.tryParseTypeArgumentsInExpression()
		if typeArguments != nil {
			p.parseErrorAt(startPos, p.nodePos(), diagnostics.X_super_may_not_use_type_arguments)
			if !p.isTemplateStartOfTaggedTemplate() {
				expression = p.finishNode(p.factory.NewExpressionWithTypeArguments(expression, typeArguments), pos)
			}
		}
	}
	if p.token == ast.KindOpenParenToken || p.token == ast.KindDotToken || p.token == ast.KindOpenBracketToken {
		return expression
	}
	// If we have seen "super" it must be followed by '(' or '.'.
	// If it wasn't then just try to parse out a '.' and report an error.
	p.parseErrorAtCurrentToken(diagnostics.X_super_must_be_followed_by_an_argument_list_or_member_access)
	// private names will never work with `super` (`super.#foo`), but that's a semantic error, not syntactic
	return p.finishNode(p.factory.NewPropertyAccessExpression(expression, nil /*questionDotToken*/, nil /*colonToken*/, p.parseRightSideOfDot(true /*allowIdentifierNames*/, true /*allowPrivateIdentifiers*/, true /*allowUnicodeEscapeSequenceInIdentifierName*/), ast.NodeFlagsNone), pos)
}

func (p *Parser) isTemplateStartOfTaggedTemplate() bool {
	return p.token == ast.KindNoSubstitutionTemplateLiteral || p.token == ast.KindTemplateHead
}

func (p *Parser) tryParseTypeArgumentsInExpression() *ast.NodeList {
	// TypeArguments must not be parsed in JavaScript files to avoid ambiguity with binary operators.
	// Check the cheap precondition before saving the parser state: unless the current token is `<`,
	// there is nothing to speculatively parse and the mark/rewind would be a no-op.
	if p.contextFlags&ast.NodeFlagsJavaScriptFile != 0 || p.token != ast.KindLessThanToken {
		return nil
	}
	state := p.mark()
	if p.reScanLessThanToken() == ast.KindLessThanToken {
		p.nextToken()
		typeArguments := p.parseDelimitedList(PCTypeArguments, (*Parser).parseTypeArgument)
		// If it doesn't have the closing `>` then it's definitely not an type argument list.
		if p.reScanGreaterThanToken() == ast.KindGreaterThanToken {
			p.nextToken()
			// We successfully parsed a type argument list. The next token determines whether we want to
			// treat it as such. If the type argument list is followed by `(` or a template literal, as in
			// `f<number>(42)`, we favor the type argument interpretation even though JavaScript would view
			// it as a relational expression.
			if p.canFollowTypeArgumentsInExpression() {
				return typeArguments
			}
		}
	}
	p.rewind(state)
	return nil
}

func (p *Parser) canFollowTypeArgumentsInExpression() bool {
	switch p.token {
	// These tokens can follow a type argument list in a call expression:
	// foo<x>(
	// foo<T> `...`
	// foo<T> `...${100}...`
	case ast.KindOpenParenToken, ast.KindNoSubstitutionTemplateLiteral, ast.KindTemplateHead:
		return true
	// A type argument list followed by `<` never makes sense, and a type argument list followed
	// by `>` is ambiguous with a (re-scanned) `>>` operator, so we disqualify both. Also, in
	// this context, `+` and `-` are unary operators, not binary operators.
	case ast.KindLessThanToken, ast.KindGreaterThanToken, ast.KindPlusToken, ast.KindMinusToken:
		return false
	}
	// We favor the type argument list interpretation when it is immediately followed by
	// a line break, a binary operator, or something that can't start an expression.
	return p.hasPrecedingLineBreak() || p.isBinaryOperator() || !p.isStartOfExpression()
}

func (p *Parser) parseMemberExpressionOrHigher() *ast.Node {
	// Note: to make our lives simpler, we decompose the NewExpression productions and
	// place ObjectCreationExpression and FunctionExpression into PrimaryExpression.
	// like so:
	//
	//   PrimaryExpression : See 11.1
	//      this
	//      Identifier
	//      Literal
	//      ArrayLiteral
	//      ObjectLiteral
	//      (Expression)
	//      FunctionExpression
	//      new MemberExpression Arguments?
	//
	//   MemberExpression : See 11.2
	//      PrimaryExpression
	//      MemberExpression[Expression]
	//      MemberExpression.IdentifierName
	//
	//   CallExpression : See 11.2
	//      MemberExpression
	//      CallExpression Arguments
	//      CallExpression[Expression]
	//      CallExpression.IdentifierName
	//
	// Technically this is ambiguous.  i.e. CallExpression defines:
	//
	//   CallExpression:
	//      CallExpression Arguments
	//
	// If you see: "new Foo()"
	//
	// Then that could be treated as a single ObjectCreationExpression, or it could be
	// treated as the invocation of "new Foo".  We disambiguate that in code (to match
	// the original grammar) by making sure that if we see an ObjectCreationExpression
	// we always consume arguments if they are there. So we treat "new Foo()" as an
	// object creation only, and not at all as an invocation.  Another way to think
	// about this is that for every "new" that we see, we will consume an argument list if
	// it is there as part of the *associated* object creation node.  Any additional
	// argument lists we see, will become invocation expressions.
	//
	// Because there are no other places in the grammar now that refer to FunctionExpression
	// or ObjectCreationExpression, it is safe to push down into the PrimaryExpression
	// production.
	//
	// Because CallExpression and MemberExpression are left recursive, we need to bottom out
	// of the recursion immediately.  So we parse out a primary expression to start with.
	pos := p.nodePos()
	expression := p.parsePrimaryExpression()
	return p.parseMemberExpressionRest(pos, expression, true /*allowOptionalChain*/)
}

func (p *Parser) parseMemberExpressionRest(pos int, expression *ast.Expression, allowOptionalChain bool) *ast.Expression {
	for {
		var questionDotToken *ast.Node
		isPropertyAccess := false
		if allowOptionalChain && p.isStartOfOptionalPropertyOrElementAccessChain() {
			questionDotToken = p.parseExpectedToken(ast.KindQuestionDotToken)
			isPropertyAccess = tokenIsIdentifierOrKeyword(p.token)
		} else {
			isPropertyAccess = p.parseOptional(ast.KindDotToken)
		}
		if isPropertyAccess {
			expression = p.parsePropertyAccessExpressionRest(pos, expression, questionDotToken, nil /*colonToken*/)
			continue
		}
		if p.parseOptional(ast.KindOpenBracketToken) {
			expression = p.parseElementAccessExpressionRest(pos, expression, questionDotToken)
			continue
		}
		if p.isTemplateStartOfTaggedTemplate() {
			// Absorb type arguments into TemplateExpression when preceding expression is ExpressionWithTypeArguments
			if questionDotToken == nil && ast.IsExpressionWithTypeArguments(expression) {
				original := expression.AsExpressionWithTypeArguments()
				expression = p.parseTaggedTemplateRest(pos, original.Expression, questionDotToken, original.TypeArguments)
				p.unparseExpressionWithTypeArguments(original.Expression, original.TypeArguments, expression)
			} else {
				expression = p.parseTaggedTemplateRest(pos, expression, questionDotToken, nil /*typeArguments*/)
			}
			continue
		}
		if questionDotToken == nil {
			if p.tokenIsExclamationPunctuation() && !p.hasPrecedingLineBreak() {
				p.nextToken()
				expression = p.checkJSSyntax(p.finishNode(p.factory.NewNonNullExpression(expression, ast.NodeFlagsNone), pos))
				continue
			}
			typeArguments := p.tryParseTypeArgumentsInExpression()
			if typeArguments != nil {
				expression = p.finishNode(p.factory.NewExpressionWithTypeArguments(expression, typeArguments), pos)
				continue
			}
		}
		return expression
	}
}

func (p *Parser) isStartOfOptionalPropertyOrElementAccessChain() bool {
	return p.token == ast.KindQuestionDotToken && p.lookAhead((*Parser).nextTokenIsIdentifierOrKeywordOrOpenBracketOrTemplate)
}

func (p *Parser) nextTokenIsIdentifierOrKeywordOrOpenBracketOrTemplate() bool {
	p.nextToken()
	return tokenIsIdentifierOrKeyword(p.token) || p.token == ast.KindOpenBracketToken || p.isTemplateStartOfTaggedTemplate()
}

func (p *Parser) parsePropertyAccessExpressionRest(pos int, expression *ast.Expression, questionDotToken *ast.Node, colonToken *ast.Node) *ast.Node {
	// Private identifiers are member names, never Lua method names.
	name := p.parseRightSideOfDot(true /*allowIdentifierNames*/, colonToken == nil /*allowPrivateIdentifiers*/, true /*allowUnicodeEscapeSequenceInIdentifierName*/)
	isOptionalChain := questionDotToken != nil || p.tryReparseOptionalChain(expression)
	propertyAccess := p.factory.NewPropertyAccessExpression(expression, questionDotToken, colonToken, name, core.IfElse(isOptionalChain, ast.NodeFlagsOptionalChain, ast.NodeFlagsNone))
	if isOptionalChain && ast.IsPrivateIdentifier(name) {
		p.parseErrorAtRange(p.skipRangeTrivia(name.Loc), diagnostics.An_optional_chain_cannot_contain_private_identifiers)
	}
	if ast.IsExpressionWithTypeArguments(expression) {
		typeArguments := expression.TypeArgumentList()
		if typeArguments != nil {
			loc := core.NewTextRange(typeArguments.Pos()-1, scanner.SkipTrivia(p.sourceText, typeArguments.End())+1)
			p.parseErrorAtRange(loc, diagnostics.An_instantiation_expression_cannot_be_followed_by_a_property_access)
		}
	}
	return p.finishNode(propertyAccess, pos)
}

func (p *Parser) tryReparseOptionalChain(node *ast.Expression) bool {
	if node.Flags&ast.NodeFlagsOptionalChain != 0 {
		return true
	}
	// check for an optional chain in a non-null expression
	if ast.IsNonNullExpression(node) {
		expr := node.Expression()
		for ast.IsNonNullExpression(expr) && expr.Flags&ast.NodeFlagsOptionalChain == 0 {
			expr = expr.Expression()
		}
		if expr.Flags&ast.NodeFlagsOptionalChain != 0 {
			// this is part of an optional chain. Walk down from `node` to `expression` and set the flag.
			for ast.IsNonNullExpression(node) {
				node.Flags |= ast.NodeFlagsOptionalChain
				node = node.Expression()
			}
			return true
		}
	}
	return false
}

func (p *Parser) parseElementAccessExpressionRest(pos int, expression *ast.Expression, questionDotToken *ast.Node) *ast.Node {
	var argumentExpression *ast.Expression
	if p.token == ast.KindCloseBracketToken {
		p.parseErrorAt(p.nodePos(), p.nodePos(), diagnostics.An_element_access_expression_should_take_an_argument)
		argumentExpression = p.createMissingIdentifier()
	} else {
		argument := p.parseExpressionAllowIn()
		switch argument.Kind {
		case ast.KindStringLiteral:
			argument.AsStringLiteral().Text = p.internIdentifier(argument.Text())
		case ast.KindNoSubstitutionTemplateLiteral:
			argument.AsNoSubstitutionTemplateLiteral().Text = p.internIdentifier(argument.Text())
		case ast.KindNumericLiteral:
			argument.AsNumericLiteral().Text = p.internIdentifier(argument.Text())
		}
		argumentExpression = argument
	}
	p.parseExpected(ast.KindCloseBracketToken)
	isOptionalChain := questionDotToken != nil || p.tryReparseOptionalChain(expression)
	return p.finishNode(p.factory.NewElementAccessExpression(expression, questionDotToken, argumentExpression, core.IfElse(isOptionalChain, ast.NodeFlagsOptionalChain, ast.NodeFlagsNone)), pos)
}

func (p *Parser) parseCallExpressionRest(pos int, expression *ast.Expression) *ast.Expression {
	for {
		expression = p.parseMemberExpressionRest(pos, expression /*allowOptionalChain*/, true)
		// `obj:name(args)` is Lua's method-call sugar for `obj.name(obj, args)`.
		// A colon access exists only as an immediately-called callee, so the
		// branch commits only when the lookahead sees the method name and its
		// argument list's `(` (or the `<` of explicit type arguments); anything
		// else -- a bare `obj:name`, annotation-shaped text like `(value: T)`
		// under error recovery -- is left for the enclosing construct to
		// diagnose. A label's `::` lexes as one ColonColonToken, so it never
		// reaches here.
		if p.token == ast.KindColonToken && p.lookAhead((*Parser).nextTokensStartLuaColonCall) {
			colonToken := p.parseTokenNode()
			expression = p.parsePropertyAccessExpressionRest(pos, expression, nil /*questionDotToken*/, colonToken)
			typeArguments := p.tryParseTypeArgumentsInExpression()
			argumentList := p.parseArgumentList()
			expression = p.finishNode(p.factory.NewCallExpression(expression, nil /*questionDotToken*/, typeArguments, argumentList, expression.Flags&ast.NodeFlagsOptionalChain), pos)
			continue
		}
		var typeArguments *ast.NodeList
		questionDotToken := p.parseOptionalToken(ast.KindQuestionDotToken)
		if questionDotToken != nil {
			typeArguments = p.tryParseTypeArgumentsInExpression()
			if p.isTemplateStartOfTaggedTemplate() {
				expression = p.parseTaggedTemplateRest(pos, expression, questionDotToken, typeArguments)
				continue
			}
		}
		if typeArguments != nil || p.token == ast.KindOpenParenToken {
			// Absorb type arguments into CallExpression when preceding expression is ExpressionWithTypeArguments
			if questionDotToken == nil && expression.Kind == ast.KindExpressionWithTypeArguments {
				typeArguments = expression.TypeArgumentList()
				expression = expression.AsExpressionWithTypeArguments().Expression
			}
			inner := expression
			argumentList := p.parseArgumentList()
			isOptionalChain := questionDotToken != nil || p.tryReparseOptionalChain(expression)
			expression = p.checkJSSyntax(p.finishNode(p.factory.NewCallExpression(expression, questionDotToken, typeArguments, argumentList, core.IfElse(isOptionalChain, ast.NodeFlagsOptionalChain, ast.NodeFlagsNone)), pos))
			p.unparseExpressionWithTypeArguments(inner, typeArguments, expression)
			continue
		}
		if questionDotToken != nil {
			// We parsed `?.` but then failed to parse anything, so report a missing identifier here.
			p.parseErrorAtCurrentToken(diagnostics.Identifier_expected)
			name := p.createMissingIdentifier()
			expression = p.finishNode(p.factory.NewPropertyAccessExpression(expression, questionDotToken, nil /*colonToken*/, name, ast.NodeFlagsOptionalChain), pos)
		}
		break
	}
	return expression
}

func (p *Parser) parseArgumentList() *ast.NodeList {
	p.parseExpected(ast.KindOpenParenToken)
	result := p.parseDelimitedList(PCArgumentExpressions, (*Parser).parseArgumentExpression)
	p.parseExpected(ast.KindCloseParenToken)
	return result
}

func (p *Parser) parseArgumentExpression() *ast.Expression {
	return doInContext(p, ast.NodeFlagsDisallowInContext, false, (*Parser).parseArgumentOrArrayLiteralElement)
}

func (p *Parser) parseArgumentOrArrayLiteralElement() *ast.Expression {
	switch p.token {
	case ast.KindCommaToken:
		return p.finishNode(p.factory.NewOmittedExpression(), p.nodePos())
	}
	// `...` is no longer a spread: it flows through the normal expression path as
	// the Lua vararg, so `f(...x)` recovers as a vararg followed by "',' expected".
	return p.parseAssignmentExpressionOrHigher()
}

func (p *Parser) parseTaggedTemplateRest(pos int, tag *ast.Expression, questionDotToken *ast.Node, typeArguments *ast.NodeList) *ast.Node {
	var template *ast.Expression
	if p.token == ast.KindNoSubstitutionTemplateLiteral {
		p.reScanTemplateToken(true /*isTaggedTemplate*/)
		template = p.parseLiteralExpression(false /*intern*/)
	} else {
		template = p.parseTemplateExpression(true /*isTaggedTemplate*/)
	}
	isOptionalChain := questionDotToken != nil || tag.Flags&ast.NodeFlagsOptionalChain != 0
	return p.checkJSSyntax(p.finishNode(p.factory.NewTaggedTemplateExpression(tag, questionDotToken, typeArguments, template, core.IfElse(isOptionalChain, ast.NodeFlagsOptionalChain, ast.NodeFlagsNone)), pos))
}

func (p *Parser) parseTemplateExpression(isTaggedTemplate bool) *ast.Expression {
	pos := p.nodePos()
	return p.finishNode(p.factory.NewTemplateExpression(p.parseTemplateHead(isTaggedTemplate), p.parseTemplateSpans(isTaggedTemplate)), pos)
}

func (p *Parser) parseTemplateSpans(isTaggedTemplate bool) *ast.NodeList {
	pos := p.nodePos()
	var list []*ast.Node
	for {
		span := p.parseTemplateSpan(isTaggedTemplate)
		list = append(list, span)
		if span.AsTemplateSpan().Literal.Kind != ast.KindTemplateMiddle {
			break
		}
	}
	return p.newNodeList(core.NewTextRange(pos, p.nodePos()), list)
}

func (p *Parser) parseTemplateSpan(isTaggedTemplate bool) *ast.Node {
	pos := p.nodePos()
	expression := p.parseExpressionAllowIn()
	literal := p.parseLiteralOfTemplateSpan(isTaggedTemplate)
	return p.finishNode(p.factory.NewTemplateSpan(expression, literal), pos)
}

func (p *Parser) parsePrimaryExpression() *ast.Expression {
	switch p.token {
	case ast.KindNoSubstitutionTemplateLiteral:
		if p.scanner.TokenFlags()&ast.TokenFlagsIsInvalid != 0 {
			p.reScanTemplateToken(false /*isTaggedTemplate*/)
		}
		fallthrough
	case ast.KindNumericLiteral, ast.KindStringLiteral:
		return p.parseLiteralExpression(false /*intern*/)
	// `this` is removed from tlua (Lua methods take an explicit `self`); it is
	// no longer a primary expression and falls through to the Expression_expected
	// recovery below. `super`/`nil`/`true`/`false` remain keyword expressions.
	case ast.KindSuperKeyword, ast.KindNilKeyword, ast.KindTrueKeyword, ast.KindFalseKeyword:
		return p.parseKeywordExpression()
	case ast.KindOpenParenToken:
		return p.parseParenthesizedExpression()
	case ast.KindOpenBracketToken:
		// Array literals exist only in JSON (tluaconfig.json etc.); tlua source
		// uses table constructors `{...}`. In TS source `[` falls through to
		// the Expression_expected recovery below.
		if p.scriptKind == core.ScriptKindJSON {
			return p.parseArrayLiteralExpression()
		}
	case ast.KindOpenBraceToken:
		return p.parseObjectLiteralExpression()
	case ast.KindAsyncKeyword:
		// Async arrow function values are removed in tlua, so `async` here can only
		// introduce an async function expression. If we encounter `async [no
		// LineTerminator here] function` then this is an async function; otherwise,
		// it is an identifier.
		if !p.lookAhead((*Parser).nextTokenIsFunctionKeywordOnSameLine) {
			break
		}
		return p.parseFunctionExpression()
	case ast.KindFunctionKeyword:
		return p.parseFunctionExpression()
	case ast.KindNewKeyword:
		return p.parseNewExpressionOrNewDotTarget()
	case ast.KindSlashToken, ast.KindSlashEqualsToken:
		if p.reScanSlashToken() == ast.KindRegularExpressionLiteral {
			return p.parseLiteralExpression(false /*intern*/)
		}
	case ast.KindTemplateHead:
		return p.parseTemplateExpression(false /*isTaggedTemplate*/)
	case ast.KindPrivateIdentifier:
		return p.parsePrivateIdentifier()
	}
	return p.parseIdentifierWithDiagnostic(diagnostics.Expression_expected, nil)
}

func (p *Parser) parseParenthesizedExpression() *ast.Expression {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	p.parseExpected(ast.KindOpenParenToken)
	expression := p.parseExpressionAllowIn()
	p.parseExpected(ast.KindCloseParenToken)
	result := p.finishNode(p.factory.NewParenthesizedExpression(expression), pos)
	p.withJSDoc(result, jsdoc)
	return result
}

func (p *Parser) parseArrayLiteralExpression() *ast.Expression {
	pos := p.nodePos()
	openBracketPosition := p.scanner.TokenStart()
	openBracketParsed := p.parseExpected(ast.KindOpenBracketToken)
	multiLine := p.hasPrecedingLineBreak()
	elements := p.parseDelimitedList(PCArrayLiteralMembers, (*Parser).parseArgumentOrArrayLiteralElement)
	p.parseExpectedMatchingBrackets(ast.KindOpenBracketToken, ast.KindCloseBracketToken, openBracketParsed, openBracketPosition)
	return p.finishNode(p.factory.NewArrayLiteralExpression(elements, multiLine), pos)
}

func (p *Parser) parseObjectLiteralExpression() *ast.Expression {
	pos := p.nodePos()
	openBracePosition := p.scanner.TokenStart()
	openBraceParsed := p.parseExpected(ast.KindOpenBraceToken)
	multiLine := p.hasPrecedingLineBreak()
	properties := p.parseDelimitedList(PCObjectLiteralMembers, (*Parser).parseObjectLiteralElement)
	p.parseExpectedMatchingBrackets(ast.KindOpenBraceToken, ast.KindCloseBraceToken, openBraceParsed, openBracePosition)
	return p.finishNode(p.factory.NewObjectLiteralExpression(properties, multiLine), pos)
}

func (p *Parser) parseObjectLiteralElement() *ast.Node {
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	// No SpreadAssignment: `...` is the Lua vararg, so `{...}` is a positional
	// entry holding a vararg expression, and `{...obj}` decomposes into that
	// entry followed by a natural error at `obj`.
	// Lua positional table entry: anything that is not a keyed member form
	// (`[k] = v`/`[k]: v`, or a literal property name followed by `=`/`:`).
	// TS object-literal methods, accessors, shorthand, and member modifiers
	// no longer parse: `{m() {}}`, `{get p() {}}`, and `{x}` decompose into
	// positional expressions with natural errors where they are not valid
	// expressions. JSON keeps the strict keyed grammar.
	if p.scriptKind != core.ScriptKindJSON &&
		p.token != ast.KindOpenBracketToken &&
		!p.lookAhead((*Parser).nextIsKeyedTableMember) {
		expression := doInContext(p, ast.NodeFlagsDisallowInContext, false, (*Parser).parseAssignmentExpressionOrHigher)
		result := p.finishNode(p.factory.NewTableEntry(expression), pos)
		p.withJSDoc(result, jsdoc)
		return result
	}
	tokenIsIdentifier := p.isIdentifier()
	name := p.parsePropertyName()
	// Disallowing of optional property assignments and definite assignment assertion happens in the grammar checker.
	postfixToken := p.parseOptionalToken(ast.KindQuestionToken)
	// questionToken and exclamationToken are not supported by property assignments and are reported in the grammar checker
	if postfixToken == nil && p.tokenIsExclamationPunctuation() {
		postfixToken = p.parseTokenNode()
	}
	// Both keyed forms are ordinary Lua table fields (`x = v` and `[k] = v`):
	// a PropertyAssignment carrying NodeFlagsLuaTableField, which drives the
	// `=` printing. There is no ShorthandPropertyAssignment -- the TS
	// name-as-value shorthand died with destructuring assignment, and giving
	// the identifier field the same node shape as every other keyed field
	// means no consumer can mistake the key for a value reference.
	var node *ast.Node
	if (tokenIsIdentifier || name.Kind == ast.KindComputedPropertyName) && p.scriptKind != core.ScriptKindJSON {
		p.parseExpected(ast.KindEqualsToken)
		initializer := doInContext(p, ast.NodeFlagsDisallowInContext, false, (*Parser).parseAssignmentExpressionOrHigher)
		node = p.factory.NewPropertyAssignment(nil /*modifiers*/, name, postfixToken, nil /*typeNode*/, initializer)
		node.Flags |= ast.NodeFlagsLuaTableField
	} else {
		// JSON's `"key": value` members — the `:` delimiter survives only here.
		p.parseExpected(ast.KindColonToken)
		initializer := doInContext(p, ast.NodeFlagsDisallowInContext, false, (*Parser).parseAssignmentExpressionOrHigher)
		node = p.factory.NewPropertyAssignment(nil /*modifiers*/, name, postfixToken, nil /*typeNode*/, initializer)
	}
	p.finishNode(node, pos)
	p.withJSDoc(node, jsdoc)
	return node
}

// nextIsKeyedTableMember classifies the next object-literal member as keyed
// (an identifier name followed by `=`, allowing the `?`/`!` postfix recovery
// forms) vs a Lua positional entry. Runs under lookAhead. Computed `[k]` keys
// are handled by the caller before this runs; the TS `:` delimiter and
// literal string/number names are deleted syntax (Lua spells those with
// brackets), so both decompose as positional entries with natural errors.
func (p *Parser) nextIsKeyedTableMember() bool {
	if !p.isIdentifier() {
		return false
	}
	p.nextToken()
	// The `!` recovery form is punctuation-only, matching the actual parse in
	// parseObjectLiteralElement: the word `not` after a property name is not a
	// postfix.
	if p.token == ast.KindQuestionToken || p.tokenIsExclamationPunctuation() {
		p.nextToken()
	}
	return p.token == ast.KindEqualsToken
}

func (p *Parser) parseFunctionExpression() *ast.Expression {
	// Lua's function constructor, the expression form of a function:
	//
	//      function (FormalParameters) Block end
	//
	// It is always anonymous -- a name here is TS syntax that no longer
	// parses, and falls out as the natural `'(' expected` -- and the body is
	// an `end`-terminated Lua block, like every other tlua function body.
	pos := p.nodePos()
	jsdoc := p.jsdocScannerInfo()
	modifiers := p.parseModifiers()
	p.parseExpected(ast.KindFunctionKeyword)
	typeParameters := p.parseTypeParameters()
	parameters := p.parseParameters(ParseFlagsNone)
	returnType := p.parseReturnTypeList(ast.KindColonToken)
	body := p.parseLuaFunctionBlock()
	result := p.factory.NewFunctionExpression(modifiers, nil /*name*/, typeParameters, parameters, returnType, nil /*fullSignature*/, body)
	p.finishNode(result, pos)
	p.withJSDoc(result, jsdoc)
	p.checkJSSyntax(result)
	return result
}

func (p *Parser) unparseExpressionWithTypeArguments(expression *ast.Node, typeArguments *ast.NodeList, result *ast.Node) {
	// force overwrite the `.Parent` of the expression and type arguments to erase the fact that they may have originally been parsed as an ExpressionWithTypeArguments and be parented to such
	if expression != nil {
		expression.Parent = result
	}
	if typeArguments != nil {
		for _, a := range typeArguments.Nodes {
			a.Parent = result
		}
	}
}

func (p *Parser) parseNewExpressionOrNewDotTarget() *ast.Node {
	pos := p.nodePos()
	p.parseExpected(ast.KindNewKeyword)
	if p.parseOptional(ast.KindDotToken) {
		name := p.parseIdentifierName()
		return p.finishNode(p.factory.NewMetaProperty(ast.KindNewKeyword, name), pos)
	}
	expressionPos := p.nodePos()
	expression := p.parseMemberExpressionRest(expressionPos, p.parsePrimaryExpression(), false /*allowOptionalChain*/)
	var typeArguments *ast.NodeList
	// Absorb type arguments into NewExpression when preceding expression is ExpressionWithTypeArguments
	if expression.Kind == ast.KindExpressionWithTypeArguments {
		typeArguments = expression.TypeArgumentList()
		expression = expression.AsExpressionWithTypeArguments().Expression
	}
	if p.token == ast.KindQuestionDotToken {
		p.parseErrorAtCurrentToken(diagnostics.Invalid_optional_chain_from_new_expression_Did_you_mean_to_call_0, scanner.GetTextOfNodeFromSourceText(p.sourceText, expression, false /*includeTrivia*/))
	}
	var argumentList *ast.NodeList
	if p.token == ast.KindOpenParenToken {
		argumentList = p.parseArgumentList()
	}
	result := p.checkJSSyntax(p.finishNode(p.factory.NewNewExpression(expression, typeArguments, argumentList), pos))
	p.unparseExpressionWithTypeArguments(expression, typeArguments, result)
	return result
}

func (p *Parser) parseKeywordExpression() *ast.Node {
	pos := p.nodePos()
	result := p.factory.NewKeywordExpression(p.token)
	p.nextToken()
	return p.finishNode(result, pos)
}

func (p *Parser) parseLiteralExpression(intern bool) *ast.Node {
	pos := p.nodePos()
	text := p.scanner.TokenValue()
	if intern {
		text = p.internIdentifier(text)
	}
	tokenFlags := p.scanner.TokenFlags()
	var result *ast.Node
	switch p.token {
	case ast.KindStringLiteral:
		result = p.factory.NewStringLiteral(text, tokenFlags)
	case ast.KindNumericLiteral:
		result = p.factory.NewNumericLiteral(text, tokenFlags)
	case ast.KindRegularExpressionLiteral:
		result = p.factory.NewRegularExpressionLiteral(text, tokenFlags)
	case ast.KindNoSubstitutionTemplateLiteral:
		result = p.factory.NewNoSubstitutionTemplateLiteral(text, tokenFlags)
	default:
		panic("Unhandled case in parseLiteralExpression")
	}
	p.nextToken()
	return p.finishNode(result, pos)
}

func (p *Parser) parseIdentifierNameErrorOnUnicodeEscapeSequence() *ast.Node {
	if p.scanner.HasUnicodeEscape() || p.scanner.HasExtendedUnicodeEscape() {
		p.parseErrorAtCurrentToken(diagnostics.Unicode_escape_sequence_cannot_appear_here)
	}
	return p.createIdentifier(tokenIsIdentifierOrKeyword(p.token))
}

func (p *Parser) parseBindingIdentifier() *ast.Node {
	return p.parseBindingIdentifierWithDiagnostic(nil)
}

func (p *Parser) parseBindingIdentifierWithDiagnostic(privateIdentifierDiagnosticMessage *diagnostics.Message) *ast.Node {
	id := p.createIdentifierWithDiagnostic(p.isBindingIdentifier(), nil /*diagnosticMessage*/, privateIdentifierDiagnosticMessage)
	return id
}

func (p *Parser) parseIdentifierName() *ast.Node {
	return p.parseIdentifierNameWithDiagnostic(nil)
}

func (p *Parser) parseIdentifierNameWithDiagnostic(diagnosticMessage *diagnostics.Message) *ast.Node {
	return p.createIdentifierWithDiagnostic(tokenIsIdentifierOrKeyword(p.token), diagnosticMessage, nil)
}

func (p *Parser) parseIdentifier() *ast.Node {
	return p.parseIdentifierWithDiagnostic(nil, nil)
}

func (p *Parser) parseIdentifierWithDiagnostic(diagnosticMessage *diagnostics.Message, privateIdentifierDiagnosticMessage *diagnostics.Message) *ast.Node {
	return p.createIdentifierWithDiagnostic(p.isIdentifier(), diagnosticMessage, privateIdentifierDiagnosticMessage)
}

func (p *Parser) createIdentifier(isIdentifier bool) *ast.Node {
	return p.createIdentifierWithDiagnostic(isIdentifier, nil, nil)
}

func (p *Parser) createIdentifierWithDiagnostic(isIdentifier bool, diagnosticMessage *diagnostics.Message, privateIdentifierDiagnosticMessage *diagnostics.Message) *ast.Node {
	if isIdentifier {
		var pos int
		if p.scanner.HasPrecedingJSDocLeadingAsterisks() {
			pos = p.scanner.TokenStart()
		} else {
			pos = p.nodePos()
		}
		text := p.scanner.TokenValue()
		p.nextTokenWithoutCheck()
		return p.finishNode(p.newIdentifier(p.internIdentifier(text)), pos)
	}
	if p.token == ast.KindPrivateIdentifier {
		if privateIdentifierDiagnosticMessage != nil {
			p.parseErrorAtCurrentToken(privateIdentifierDiagnosticMessage)
		} else {
			p.parseErrorAtCurrentToken(diagnostics.Private_identifiers_are_not_allowed_outside_class_bodies)
		}
		return p.createIdentifier(true /*isIdentifier*/)
	}
	// Only for end of file because the error gets reported incorrectly on embedded script tags.
	reportAtCurrentPosition := p.token == ast.KindEndOfFile
	if diagnosticMessage != nil {
		if reportAtCurrentPosition {
			pos := p.scanner.TokenFullStart()
			p.parseErrorAt(pos, pos, diagnosticMessage)
		} else {
			p.parseErrorAtCurrentToken(diagnosticMessage)
		}
	} else if isReservedWord(p.token) {
		if reportAtCurrentPosition {
			pos := p.scanner.TokenFullStart()
			p.parseErrorAt(pos, pos, diagnostics.Identifier_expected_0_is_a_reserved_word_that_cannot_be_used_here, p.scanner.TokenText())
		} else {
			p.parseErrorAtCurrentToken(diagnostics.Identifier_expected_0_is_a_reserved_word_that_cannot_be_used_here, p.scanner.TokenText())
		}
	} else {
		if reportAtCurrentPosition {
			pos := p.scanner.TokenFullStart()
			p.parseErrorAt(pos, pos, diagnostics.Identifier_expected)
		} else {
			p.parseErrorAtCurrentToken(diagnostics.Identifier_expected)
		}
	}
	return p.createMissingIdentifier()
}

func (p *Parser) internIdentifier(text string) string {
	if identifier, ok := p.identifiers[text]; ok {
		return identifier
	}
	identifier := text
	if p.identifiers == nil {
		p.identifiers = make(map[string]string)
	}
	p.identifiers[identifier] = identifier
	return identifier
}

func (p *Parser) newNodeList(loc core.TextRange, nodes []*ast.Node) *ast.NodeList {
	list := p.factory.NewNodeList(nodes)
	list.Loc = loc
	return list
}

func (p *Parser) newModifierList(loc core.TextRange, nodes []*ast.Node) *ast.ModifierList {
	list := p.factory.NewModifierList(nodes)
	list.Loc = loc
	return list
}

func (p *Parser) finishNode(node *ast.Node, pos int) *ast.Node {
	return p.finishNodeWithEnd(node, pos, p.nodePos())
}

func (p *Parser) finishNodeWithEnd(node *ast.Node, pos int, end int) *ast.Node {
	node.Loc = core.NewTextRange(pos, end)
	node.Flags |= p.contextFlags
	if p.hasParseError {
		node.Flags |= ast.NodeFlagsThisNodeHasError
		p.hasParseError = false
	}
	p.overrideParentInImmediateChildren(node)
	return node
}

func (p *Parser) overrideParentInImmediateChildren(node *ast.Node) {
	p.currentParent = node
	node.ForEachChild(p.setParentFromContext)
	p.currentParent = nil
}

func (p *Parser) nextTokenIsSlash() bool {
	return p.nextToken() == ast.KindSlashToken
}

func (p *Parser) scanTypeMemberStart() bool {
	// Return true if we have the start of a signature member
	if p.token == ast.KindOpenParenToken || p.token == ast.KindLessThanToken || p.token == ast.KindGetKeyword || p.token == ast.KindSetKeyword {
		return true
	}
	idToken := false
	// Eat up all modifiers, but hold on to the last one in case it is actually an identifier
	for ast.IsModifierKind(p.token) {
		idToken = true
		p.nextToken()
	}
	// Index signatures and computed property names are type members
	if p.token == ast.KindOpenBracketToken {
		return true
	}
	// Try to get the first property-like token following all modifiers
	if p.isLiteralPropertyName() {
		idToken = true
		p.nextToken()
	}
	// If we were able to get any potential identifier, check that it is
	// the start of a member declaration
	if idToken {
		return p.token == ast.KindOpenParenToken || p.token == ast.KindLessThanToken || p.token == ast.KindQuestionToken || p.token == ast.KindColonToken || p.token == ast.KindCommaToken || p.canParseSemicolon()
	}
	return false
}

func (p *Parser) canParseSemicolon() bool {
	// If there's a real semicolon, then we can always parse it out.
	// We can parse out an optional semicolon in ASI cases in the following cases.
	// The Lua clause terminators make one-liners parse (`while c do break end`).
	return p.token == ast.KindSemicolonToken || p.token == ast.KindCloseBraceToken || p.token == ast.KindEndOfFile ||
		p.token == ast.KindEndKeyword || p.token == ast.KindUntilKeyword || p.token == ast.KindElseIfKeyword || p.token == ast.KindElseKeyword ||
		p.hasPrecedingLineBreak()
}

func (p *Parser) tryParseSemicolon() bool {
	if !p.canParseSemicolon() {
		return false
	}
	if p.token == ast.KindSemicolonToken {
		// consume the semicolon if it was explicitly provided.
		p.nextToken()
	}
	return true
}

func (p *Parser) parseSemicolon() bool {
	return p.tryParseSemicolon() || p.parseExpected(ast.KindSemicolonToken)
}

func (p *Parser) isLiteralPropertyName() bool {
	return tokenIsIdentifierOrKeyword(p.token) || p.token == ast.KindStringLiteral || p.token == ast.KindNumericLiteral
}

func (p *Parser) isStartOfStatement() bool {
	switch p.token {
	case ast.KindSemicolonToken, ast.KindOpenBraceToken,
		ast.KindLocalKeyword, ast.KindFunctionKeyword, ast.KindClassKeyword, ast.KindIfKeyword,
		ast.KindDoKeyword, ast.KindWhileKeyword, ast.KindForKeyword, ast.KindRepeatKeyword, ast.KindContinueKeyword, ast.KindBreakKeyword,
		ast.KindColonColonToken, ast.KindGotoKeyword,
		ast.KindReturnKeyword, ast.KindThrowKeyword,
		ast.KindDebuggerKeyword:
		return true
	case ast.KindAsyncKeyword, ast.KindDeclareKeyword, ast.KindInterfaceKeyword, ast.KindModuleKeyword, ast.KindNamespaceKeyword,
		ast.KindTypeKeyword, ast.KindGlobalKeyword, ast.KindDeferKeyword:
		// When these don't start a declaration, they're an identifier in an expression statement
		return true
	case ast.KindAccessorKeyword, ast.KindPublicKeyword, ast.KindPrivateKeyword, ast.KindProtectedKeyword, ast.KindStaticKeyword,
		ast.KindReadonlyKeyword:
		// When these don't start a declaration, they may be the start of a class member if an identifier
		// immediately follows. Otherwise they're an identifier in an expression statement.
		return p.isStartOfDeclaration() || !p.lookAhead((*Parser).nextTokenIsIdentifierOrKeywordOnSameLine)

	default:
		return p.isStartOfExpression()
	}
}

func (p *Parser) isStartOfDeclaration() bool {
	return p.lookAhead((*Parser).scanStartOfDeclaration)
}

func (p *Parser) scanStartOfDeclaration() bool {
	for {
		switch p.token {
		case ast.KindLocalKeyword, ast.KindFunctionKeyword, ast.KindClassKeyword:
			return true
		// 'declare', 'module', 'namespace', 'interface'* and 'type' are all legal JavaScript identifiers;
		// however, an identifier cannot be followed by another identifier on the same line. This is what we
		// count on to parse out the respective declarations. For instance, we exploit this to say that
		//
		//    namespace n
		//
		// can be none other than the beginning of a namespace declaration, but need to respect that JavaScript sees
		//
		//    namespace
		//    n
		//
		// as the identifier 'namespace' on one line followed by the identifier 'n' on another.
		// We need to look one token ahead to see if it permissible to try parsing a declaration.
		//
		// *Note*: 'interface' is actually a strict mode reserved word. So while
		//
		//   "use strict"
		//   interface
		//   I {}
		//
		// could be legal, it would add complexity for very little gain.
		case ast.KindInterfaceKeyword, ast.KindTypeKeyword, ast.KindDeferKeyword:
			return p.nextTokenIsIdentifierOnSameLine()
		case ast.KindAbstractKeyword, ast.KindAccessorKeyword, ast.KindAsyncKeyword, ast.KindDeclareKeyword, ast.KindPrivateKeyword,
			ast.KindProtectedKeyword, ast.KindPublicKeyword, ast.KindReadonlyKeyword:
			previousToken := p.token
			p.nextToken()
			// ASI takes effect for this modifier.
			if p.hasPrecedingLineBreak() {
				return false
			}
			if previousToken == ast.KindDeclareKeyword && p.token == ast.KindTypeKeyword {
				// If we see 'declare type', then commit to parsing a type alias. parseTypeAliasDeclaration will
				// report Line_break_not_permitted_here if needed.
				return true
			}
			if previousToken == ast.KindDeclareKeyword && isAmbientGlobalNameToken(p.token) {
				// The Luau-style ambient global form `declare name: Type;`. Only the
				// trailing colon distinguishes it from `declare` used as an identifier
				// expression, so require it before committing to a declaration.
				p.nextToken()
				return p.token == ast.KindColonToken
			}
			continue
		case ast.KindGlobalKeyword:
			p.nextToken()
			return p.token == ast.KindOpenBraceToken || p.token == ast.KindIdentifier
		case ast.KindStaticKeyword:
			p.nextToken()
			continue
		}
		return false
	}
}

func (p *Parser) isStartOfExpression() bool {
	if p.isStartOfLeftHandSideExpression() {
		return true
	}
	switch p.token {
	// `typeof` is absent: it starts a type query, never an expression. Admitting it
	// here would make the statement list re-enter on a token no expression parser
	// consumes, which never terminates.
	case ast.KindPlusToken, ast.KindMinusToken, ast.KindExclamationToken, ast.KindHashToken, ast.KindDeleteKeyword,
		ast.KindVoidKeyword, ast.KindLessThanToken,
		ast.KindAwaitKeyword, ast.KindYieldKeyword, ast.KindPrivateIdentifier:
		// Yield/await always starts an expression.  Either it is an identifier (in which case
		// it is definitely an expression).  Or it's a keyword (either because we're in
		// a generator or async function, or in strict mode (or both)) and it started a yield or await expression.
		return true
	}
	// Error tolerance.  If we see the start of some binary operator, we consider
	// that the start of an expression.  That way we'll parse out a missing identifier,
	// give a good message about an identifier being missing, and then consume the
	// rest of the binary expression.
	if p.isBinaryOperator() {
		return true
	}
	return p.isIdentifier()
}

func (p *Parser) isStartOfLeftHandSideExpression() bool {
	switch p.token {
	// `this` is not an expression start in tlua. It must stay out of this set:
	// `this` is a reserved keyword, so `parseIdentifier` will not consume it, and
	// treating it as an expression start would spin the enclosing list parser.
	// Excluded here, a stray `this` is skipped by list recovery instead.
	case ast.KindSuperKeyword, ast.KindNilKeyword, ast.KindTrueKeyword, ast.KindFalseKeyword,
		ast.KindNumericLiteral, ast.KindStringLiteral, ast.KindNoSubstitutionTemplateLiteral, ast.KindTemplateHead,
		ast.KindOpenParenToken, ast.KindOpenBraceToken, ast.KindFunctionKeyword, ast.KindClassKeyword,
		ast.KindNewKeyword, ast.KindSlashToken, ast.KindSlashEqualsToken, ast.KindIdentifier,
		// `...` is the Lua vararg, a primary expression.
		ast.KindDotDotDotToken:
		return true
	case ast.KindOpenBracketToken:
		// `[` starts an expression only in JSON (array literals); keeps
		// nested-array list recovery working there without making `[` an
		// expression start in tlua source.
		return p.scriptKind == core.ScriptKindJSON
	}
	return p.isIdentifier()
}

func (p *Parser) isStartOfType(inStartOfParameter bool) bool {
	switch p.token {
	case ast.KindAnyKeyword, ast.KindUnknownKeyword, ast.KindStringKeyword, ast.KindNumberKeyword,
		ast.KindBooleanKeyword, ast.KindReadonlyKeyword, ast.KindSymbolKeyword, ast.KindUniqueKeyword, ast.KindVoidKeyword,
		ast.KindNilKeyword, ast.KindTypeOfKeyword, ast.KindNeverKeyword,
		ast.KindOpenBraceToken, ast.KindLessThanToken, ast.KindBarToken, ast.KindAmpersandToken,
		ast.KindNewKeyword, ast.KindStringLiteral, ast.KindNumericLiteral, ast.KindTrueKeyword,
		ast.KindFalseKeyword, ast.KindObjectKeyword, ast.KindThreadKeyword, ast.KindUserdataKeyword, ast.KindCDataKeyword,
		ast.KindAsteriskToken, ast.KindQuestionToken,
		ast.KindDotDotDotToken, ast.KindInferKeyword, ast.KindImportKeyword, ast.KindAssertsKeyword, ast.KindNoSubstitutionTemplateLiteral,
		ast.KindTemplateHead:
		return true
	case ast.KindExclamationToken:
		// The punctuation `!` starts a JSDoc non-nullable type; the word `not`
		// starts no type at all.
		return p.tokenIsExclamationPunctuation()
	case ast.KindOpenBracketToken:
		// `[` starts a tuple type -- except when deciding whether a token
		// begins a PARAMETER: array binding patterns are declaration-file
		// only (isArrayBindingPatternStart handles them there), so treating
		// `[` as a parameter start in source would manufacture a phantom
		// nameless parameter out of recovery on removed syntax.
		return !inStartOfParameter
	case ast.KindFunctionKeyword:
		return !inStartOfParameter
	case ast.KindMinusToken:
		return !inStartOfParameter && p.lookAhead((*Parser).nextTokenIsNumericLiteral)
	case ast.KindOpenParenToken:
		// Only consider '(' the start of a type if followed by ')', '...', an identifier, a modifier,
		// or something that starts a type. We don't want to consider things like '(1)' a type.
		return !inStartOfParameter && p.lookAhead((*Parser).nextIsParenthesizedOrFunctionType)
	}
	return p.isIdentifier()
}

func (p *Parser) nextTokenIsNumericLiteral() bool {
	p.nextToken()
	return p.token == ast.KindNumericLiteral
}

func (p *Parser) nextIsParenthesizedOrFunctionType() bool {
	p.nextToken()
	return p.token == ast.KindCloseParenToken || p.isStartOfParameter(false /*isJSDocParameter*/) || p.isStartOfType(false /*inStartOfParameter*/)
}

func (p *Parser) isStartOfParameter(isJSDocParameter bool) bool {
	return p.token == ast.KindDotDotDotToken ||
		p.isBindingIdentifierOrPrivateIdentifierOrPattern() ||
		ast.IsModifierKind(p.token) ||
		p.isStartOfType(!isJSDocParameter /*inStartOfParameter*/)
}

func (p *Parser) isBindingIdentifierOrPrivateIdentifierOrPattern() bool {
	return p.token == ast.KindOpenBraceToken || p.token == ast.KindPrivateIdentifier || p.isBindingIdentifier() || p.isArrayBindingPatternStart()
}

func (p *Parser) isNextTokenOpenParenOrLessThanOrDot() bool {
	return p.lookAhead((*Parser).nextTokenIsOpenParenOrLessThanOrDot)
}

func (p *Parser) nextTokenIsOpenParenOrLessThanOrDot() bool {
	switch p.nextToken() {
	case ast.KindOpenParenToken, ast.KindLessThanToken, ast.KindDotToken:
		return true
	}
	return false
}

func (p *Parser) nextTokenIsIdentifierOnSameLine() bool {
	p.nextToken()
	return p.isIdentifier() && !p.hasPrecedingLineBreak()
}

// Ignore strict mode flag because we will report an error in type checker instead.
func (p *Parser) isIdentifier() bool {
	if p.token == ast.KindIdentifier {
		return true
	}
	// `yield` and `await` are always ordinary identifiers in tlua: generators
	// and await expressions are removed, so there are no [Yield]/[Await]
	// parsing contexts.
	return p.token > ast.KindLastReservedWord
}

func (p *Parser) isBindingIdentifier() bool {
	// `let await`/`let yield` in [Yield] or [Await] are allowed here and disallowed in the binder.
	return p.token == ast.KindIdentifier || p.token > ast.KindLastReservedWord
}

func (p *Parser) isBinaryOperator() bool {
	if p.inDisallowInContext() && p.token == ast.KindInKeyword {
		return false
	}
	return ast.GetBinaryOperatorPrecedence(p.token) != ast.OperatorPrecedenceInvalid
}

func (p *Parser) isValidHeritageClauseObjectLiteral() bool {
	return p.lookAhead((*Parser).nextIsValidHeritageClauseObjectLiteral)
}

func (p *Parser) nextIsValidHeritageClauseObjectLiteral() bool {
	if p.nextToken() == ast.KindCloseBraceToken {
		// if we see "extends {}" then only treat the {} as what we're extending (and not
		// the class body) if we have:
		//
		//      extends {} {
		//      extends {},
		//      extends {} extends
		//      extends {} implements
		next := p.nextToken()
		return next == ast.KindCommaToken || next == ast.KindOpenBraceToken || next == ast.KindExtendsKeyword || next == ast.KindImplementsKeyword
	}
	return true
}

func (p *Parser) isHeritageClause() bool {
	return p.token == ast.KindExtendsKeyword || p.token == ast.KindImplementsKeyword
}

func (p *Parser) isHeritageClauseExtendsOrImplementsKeyword() bool {
	return p.isHeritageClause() && p.lookAhead((*Parser).nextIsStartOfExpression)
}

func (p *Parser) nextIsStartOfExpression() bool {
	p.nextToken()
	return p.isStartOfExpression()
}

func (p *Parser) nextTokenIsTokenStringLiteral() bool {
	return p.nextToken() == ast.KindStringLiteral
}

func (p *Parser) setContextFlags(flags ast.NodeFlags, value bool) {
	if value {
		p.contextFlags |= flags
	} else {
		p.contextFlags &^= flags
	}
}

func doInContext[T any](p *Parser, flags ast.NodeFlags, value bool, f func(p *Parser) T) T {
	saveContextFlags := p.contextFlags
	p.setContextFlags(flags, value)
	result := f(p)
	p.contextFlags = saveContextFlags
	return result
}

func (p *Parser) inDisallowInContext() bool {
	return p.contextFlags&ast.NodeFlagsDisallowInContext != 0
}

// tokenIsExclamationPunctuation reports whether the current token is `!` written
// as punctuation rather than as the word `not`. The two share
// KindExclamationToken, but only the punctuation spelling is the non-null,
// definite-assignment and non-nullable-type operator: `x not` and
// `local x not: T` are not Lua, and accepting them would silently suppress nil
// checks. The scanner records the spelling when it reads the word.
func (p *Parser) tokenIsExclamationPunctuation() bool {
	return p.token == ast.KindExclamationToken && !p.scanner.TokenIsWordOperator()
}

// tokenIsJsxName reports whether the current token can start a JSX tag or
// attribute name. A JSX name is an IdentifierName, so keywords qualify; so do
// the word-spelled operators, which the ordinary scan has already resolved to
// their punctuation kind — scanJsxIdentifier turns them back into identifiers,
// but the list-element test runs first.
func (p *Parser) tokenIsJsxName() bool {
	return tokenIsIdentifierOrKeyword(p.token) || p.scanner.TokenIsWordOperator()
}

func (p *Parser) inDisallowConditionalTypesContext() bool {
	return p.contextFlags&ast.NodeFlagsDisallowConditionalTypesContext != 0
}

func (p *Parser) skipRangeTrivia(textRange core.TextRange) core.TextRange {
	return core.NewTextRange(scanner.SkipTrivia(p.sourceText, textRange.Pos()), textRange.End())
}

func isReservedWord(token ast.Kind) bool {
	return ast.KindFirstReservedWord <= token && token <= ast.KindLastReservedWord
}

func attachFileToDiagnostics(diagnostics []*ast.Diagnostic, file *ast.SourceFile) []*ast.Diagnostic {
	for _, d := range diagnostics {
		d.SetFile(file)
		for _, r := range d.RelatedInformation() {
			r.SetFile(file)
		}
	}
	return diagnostics
}

func getCommentPragmas(f *ast.NodeFactory, sourceText string) (pragmas []ast.Pragma) {
	for commentRange := range scanner.GetLeadingCommentRanges(f, sourceText, 0) {
		comment := sourceText[commentRange.Pos():commentRange.End()]
		pragmas = append(pragmas, extractPragmas(commentRange, comment)...)
	}
	return pragmas
}

func extractPragmas(commentRange ast.CommentRange, text string) []ast.Pragma {
	if commentRange.Kind == ast.KindSingleLineCommentTrivia {
		pos := 2
		tripleSlash := match(text, pos, "/")
		if tripleSlash {
			pos++
		}
		pos = skipBlanks(text, pos)
		if tripleSlash && match(text, pos, "<") {
			tagName := extractName(text, pos+1)
			if tagName != "reference" {
				return nil
			}
			pos += 10
			args := make(map[string]ast.PragmaArgument)
			for {
				pos = skipBlanks(text, pos)
				if match(text, pos, "/>") {
					break
				}
				argName := extractName(text, pos)
				if argName == "" {
					break
				}
				pos = skipBlanks(text, pos+len(argName))
				if !match(text, pos, "=") {
					break
				}
				pos = skipBlanks(text, pos+1)
				value, ok := extractQuotedString(text, pos)
				if !ok {
					break
				}
				args[argName] = ast.PragmaArgument{
					Name:      argName,
					Value:     value,
					TextRange: core.NewTextRange(commentRange.Pos()+pos+1, commentRange.Pos()+pos+1+len(value)),
				}
				pos += len(value) + 2
			}
			return []ast.Pragma{{
				CommentRange: commentRange,
				Name:         "reference",
				Args:         args,
			}}
		}
		if match(text, pos, "@") {
			pos++
			pragmaName := extractName(text, pos)
			if !(pragmaName == "ts-check" || pragmaName == "ts-nocheck") {
				return nil
			}
			return []ast.Pragma{{
				CommentRange: commentRange,
				Name:         pragmaName,
			}}
		}
	}
	if commentRange.Kind == ast.KindMultiLineCommentTrivia {
		text = strings.TrimSuffix(text, "*/")
		pos := 2
		var pragmas []ast.Pragma
		for {
			if pos = skipTo(text, pos, "@"); pos < 0 {
				break
			}
			// Mirrors the /@(\S+)(\s+(?:\S.*)?)?$/gm pragma regex used by TypeScript: the '@'
			// must be immediately followed by a non-whitespace pragma name, and the remainder
			// of the line is consumed as that pragma's arguments. As a consequence, only the
			// first '@'-token on a line is considered, so an unrelated '@token' earlier on the
			// line (e.g. an email address) prevents a later '@jsx' on the same line from being
			// treated as a pragma.
			namePos := pos + 1
			nameEnd := skipNonBlanks(text, namePos)
			if nameEnd == namePos {
				pos++
				continue
			}
			lineEnd := lineEndPos(text, pos)
			pragmaName := strings.ToLower(text[namePos:nameEnd])
			if pragmaName == "jsx" || pragmaName == "jsxfrag" || pragmaName == "jsximportsource" || pragmaName == "jsxruntime" {
				start := skipBlanks(text, nameEnd)
				argEnd := skipNonBlanks(text, start)
				if argEnd != start {
					args := make(map[string]ast.PragmaArgument, 1)
					args["factory"] = ast.PragmaArgument{
						Name:      "factory",
						Value:     text[start:argEnd],
						TextRange: core.NewTextRange(commentRange.Pos()+start, commentRange.Pos()+argEnd),
					}
					pragmas = append(pragmas, ast.Pragma{
						CommentRange: commentRange,
						Name:         pragmaName,
						Args:         args,
					})
				}
			}
			pos = lineEnd
		}
		return pragmas
	}
	return nil
}

func match(text string, pos int, s string) bool {
	return strings.HasPrefix(text[pos:], s)
}

func skipBlanks(text string, pos int) int {
	for pos < len(text) && (text[pos] == ' ' || text[pos] == '\t') {
		pos++
	}
	return pos
}

func skipNonBlanks(text string, pos int) int {
	for pos < len(text) && (text[pos] != ' ' && text[pos] != '\t' && text[pos] != '\r' && text[pos] != '\n') {
		pos++
	}
	return pos
}

func skipTo(text string, pos int, s string) int {
	if pos >= len(text) {
		return -1
	}
	i := strings.Index(text[pos:], s)
	if i < 0 {
		return -1
	}
	return pos + i
}

func lineEndPos(text string, pos int) int {
	for pos < len(text) {
		ch, size := utf8.DecodeRuneInString(text[pos:])
		if stringutil.IsLineBreak(ch) {
			return pos
		}
		pos += size
	}
	return len(text)
}

func extractName(text string, pos int) string {
	start := pos
	for pos < len(text) && (text[pos] >= 'A' && text[pos] <= 'Z' || text[pos] >= 'a' && text[pos] <= 'z' || text[pos] == '-') {
		pos++
	}
	return strings.ToLower(text[start:pos])
}

func extractQuotedString(text string, pos int) (string, bool) {
	if pos == len(text) {
		return "", false
	}
	quote := text[pos]
	if quote != '\'' && quote != '"' {
		return "", false
	}
	pos++
	start := pos
	for pos < len(text) && text[pos] != quote {
		pos++
	}
	if pos == len(text) {
		return "", false
	}
	return text[start:pos], true
}

func (p *Parser) processPragmasIntoFields(context *ast.SourceFile) {
	context.CheckJsDirective = nil
	context.ReferencedFiles = nil
	context.TypeReferenceDirectives = nil
	context.LibReferenceDirectives = nil
	// context.AmdDependencies = nil
	for _, pragma := range context.Pragmas {
		switch pragma.Name {
		case "reference":
			types, typesOk := pragma.Args["types"]
			lib, libOk := pragma.Args["lib"]
			path, pathOk := pragma.Args["path"]
			resolutionMode, resolutionModeOk := pragma.Args["resolution-mode"]
			preserve, preserveOk := pragma.Args["preserve"]
			noDefaultLib, noDefaultLibOk := pragma.Args["no-default-lib"]
			switch {
			case noDefaultLibOk && noDefaultLib.Value == "true":
				// Ignored.
			case typesOk:
				var parsed core.ResolutionMode
				if resolutionModeOk {
					parsed = p.parseResolutionMode(resolutionMode.Value, resolutionMode.Pos(), resolutionMode.End())
				}
				context.TypeReferenceDirectives = append(context.TypeReferenceDirectives, &ast.FileReference{
					TextRange:      types.TextRange,
					FileName:       types.Value,
					ResolutionMode: parsed,
					Preserve:       preserveOk && preserve.Value == "true",
				})
			case libOk:
				context.LibReferenceDirectives = append(context.LibReferenceDirectives, &ast.FileReference{
					TextRange: lib.TextRange,
					FileName:  lib.Value,
					Preserve:  preserveOk && preserve.Value == "true",
				})
			case pathOk:
				context.ReferencedFiles = append(context.ReferencedFiles, &ast.FileReference{
					TextRange: path.TextRange,
					FileName:  path.Value,
					Preserve:  preserveOk && preserve.Value == "true",
				})
			default:
				p.parseErrorAtRange(pragma.TextRange, diagnostics.Invalid_reference_directive_syntax)
			}
		case "ts-check", "ts-nocheck":
			// _last_ of either nocheck or check in a file is the "winner"
			if context.CheckJsDirective == nil || pragma.TextRange.Pos() > context.CheckJsDirective.Range.Pos() {
				context.CheckJsDirective = &ast.CheckJsDirective{
					Enabled: pragma.Name == "ts-check",
					Range:   pragma.CommentRange,
				}
			}
		case "jsx", "jsxfrag", "jsximportsource", "jsxruntime":
			// Nothing to do here
		default:
			panic("Unhandled pragma kind: " + pragma.Name)
		}
	}
}

func (p *Parser) parseResolutionMode(mode string, pos int, end int) (resolutionKind core.ResolutionMode) {
	if mode == "import" {
		resolutionKind = core.ModuleKindESNext
		return resolutionKind
	}
	if mode == "require" {
		resolutionKind = core.ModuleKindCommonJS
		return resolutionKind
	}
	p.parseErrorAt(pos, end, diagnostics.X_resolution_mode_should_be_either_require_or_import)
	return resolutionKind
}

func (p *Parser) jsErrorAtRange(loc core.TextRange, message *diagnostics.Message, args ...any) {
	p.jsDiagnostics = append(p.jsDiagnostics, ast.NewDiagnostic(nil, core.NewTextRange(scanner.SkipTrivia(p.sourceText, loc.Pos()), loc.End()), message, args...))
}

func (p *Parser) checkJSSyntax(node *ast.Node) *ast.Node {
	if node.Flags&ast.NodeFlagsJavaScriptFile == 0 || node.Flags&(ast.NodeFlagsJSDoc|ast.NodeFlagsReparsed) != 0 {
		return node
	}
	switch node.Kind {
	case ast.KindParameter:
		if token := node.QuestionToken(); token != nil && token.Flags&ast.NodeFlagsReparsed == 0 && ast.IsQuestionToken(token) {
			p.jsErrorAtRange(token.Loc, diagnostics.The_0_modifier_can_only_be_used_in_tlua_files, "?")
		}
		fallthrough
	case ast.KindMethodSignature, ast.KindFunctionExpression,
		ast.KindFunctionDeclaration, ast.KindArrowFunction, ast.KindVariableDeclaration, ast.KindIndexSignature:
		if ast.IsFunctionLike(node) && node.Body() == nil {
			p.jsErrorAtRange(node.Loc, diagnostics.Signature_declarations_can_only_be_used_in_tlua_files)
		} else if t := node.Type(); t != nil && t.Flags&ast.NodeFlagsReparsed == 0 {
			p.jsErrorAtRange(t.Loc, diagnostics.Type_annotations_can_only_be_used_in_tlua_files)
		}
	case ast.KindImportDeclaration:
		if clause := node.ImportClause(); clause != nil && clause.IsTypeOnly() {
			p.jsErrorAtRange(node.Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, "import type")
		}
	case ast.KindExportDeclaration:
		if node.IsTypeOnly() {
			p.jsErrorAtRange(node.Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, "export type")
		}
	case ast.KindImportSpecifier:
		if node.IsTypeOnly() {
			p.jsErrorAtRange(node.Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, "import...type")
		}
	case ast.KindExportSpecifier:
		if node.IsTypeOnly() {
			p.jsErrorAtRange(node.Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, "export...type")
		}
	case ast.KindImportEqualsDeclaration:
		p.jsErrorAtRange(node.Loc, diagnostics.X_import_can_only_be_used_in_tlua_files)
	case ast.KindExportAssignment:
		if node.AsExportAssignment().IsExportEquals {
			p.jsErrorAtRange(node.Loc, diagnostics.X_export_can_only_be_used_in_tlua_files)
		}
	case ast.KindHeritageClause:
		if node.AsHeritageClause().Token == ast.KindImplementsKeyword {
			p.jsErrorAtRange(node.Loc, diagnostics.X_implements_clauses_can_only_be_used_in_tlua_files)
		}
	case ast.KindInterfaceDeclaration:
		p.jsErrorAtRange(node.Name().Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, "interface")
	case ast.KindModuleDeclaration:
		p.jsErrorAtRange(node.Name().Loc, diagnostics.X_0_declarations_can_only_be_used_in_tlua_files, scanner.TokenToString(node.AsModuleDeclaration().Keyword))
	case ast.KindTypeAliasDeclaration:
		p.jsErrorAtRange(node.Name().Loc, diagnostics.Type_aliases_can_only_be_used_in_tlua_files)
	case ast.KindNonNullExpression:
		p.jsErrorAtRange(node.Loc, diagnostics.Non_null_assertions_can_only_be_used_in_tlua_files)
	case ast.KindAsExpression:
		p.jsErrorAtRange(node.Type().Loc, diagnostics.Type_assertion_expressions_can_only_be_used_in_tlua_files)
	case ast.KindSatisfiesExpression:
		p.jsErrorAtRange(node.Type().Loc, diagnostics.Type_satisfaction_expressions_can_only_be_used_in_tlua_files)
	}
	// Check absence of type parameters, type arguments and non-JavaScript modifiers
	switch node.Kind {
	case ast.KindFunctionExpression, ast.KindFunctionDeclaration, ast.KindArrowFunction:
		if list := node.TypeParameterList(); list != nil && core.Some(list.Nodes, func(n *ast.Node) bool { return n.Flags&ast.NodeFlagsReparsed == 0 }) {
			p.jsErrorAtRange(list.Loc, diagnostics.Type_parameter_declarations_can_only_be_used_in_tlua_files)
		}
		fallthrough
	case ast.KindVariableStatement:
		for _, modifier := range node.ModifierNodes() {
			if modifier.Flags&ast.NodeFlagsReparsed == 0 && ast.ModifierToFlag(modifier.Kind)&ast.ModifierFlagsJavaScript == 0 {
				p.jsErrorAtRange(modifier.Loc, diagnostics.The_0_modifier_can_only_be_used_in_tlua_files, scanner.TokenToString(modifier.Kind))
			}
		}
	case ast.KindParameter:
		if core.Some(node.ModifierNodes(), ast.IsModifier) {
			p.jsErrorAtRange(node.Modifiers().Loc, diagnostics.Parameter_modifiers_can_only_be_used_in_tlua_files)
		}
	case ast.KindCallExpression, ast.KindNewExpression, ast.KindExpressionWithTypeArguments, ast.KindJsxSelfClosingElement,
		ast.KindJsxOpeningElement, ast.KindTaggedTemplateExpression:
		if list := node.TypeArgumentList(); list != nil && core.Some(list.Nodes, func(n *ast.Node) bool { return n.Flags&ast.NodeFlagsReparsed == 0 }) {
			p.jsErrorAtRange(list.Loc, diagnostics.Type_arguments_can_only_be_used_in_tlua_files)
		}
	}
	return node
}

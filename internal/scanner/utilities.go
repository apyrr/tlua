package scanner

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/debug"
)

func IdentifierToKeywordKind(node *ast.Identifier) ast.Kind {
	return textToKeyword[node.Text]
}

func GetSourceTextOfNodeFromSourceFile(sourceFile *ast.SourceFile, node *ast.Node, includeTrivia bool) string {
	return GetTextOfNodeFromSourceText(sourceFile.Text(), node, includeTrivia)
}

func GetTextOfNodeFromSourceText(sourceText string, node *ast.Node, includeTrivia bool) string {
	if ast.NodeIsMissing(node) {
		return ""
	}
	pos := node.Pos()
	if !includeTrivia {
		pos = SkipTrivia(sourceText, pos)
	}
	text := sourceText[pos:node.End()]
	if node.Flags&ast.NodeFlagsReparserTransformedLiteral != 0 {
		// This is similar to `getLiteralTextOfNode` in the printer, but without the context of an `emitContext` to provide overrides
		if ast.IsStringLiteral(node) {
			if node.AsStringLiteral().TokenFlags&ast.TokenFlagsSingleQuote != 0 {
				return "'" + text + "'"
			}
			return "\"" + text + "\""
		} else if ast.IsIdentifier(node) {
			return node.Text()
		}
		// Only the above node kinds are currently transformed into one another by the reparser, requiring the textual remapping.
		// (Any reamppings done by emit transforms are handled by `getLiteralTextOfNode` in the printer)
		// Fail on any other kinds.
		debug.FailBadSyntaxKind(node, "Unexpected reparser-transformed node kind")
	}
	// if (isJSDocTypeExpressionOrChild(node)) {
	//     // strip space + asterisk at line start
	//     text = text.split(/\r\n|\n|\r/).map(line => line.replace(/^\s*\*/, "").trimStart()).join("\n");
	// }
	return text
}

func GetTextOfNode(node *ast.Node) string {
	return GetSourceTextOfNodeFromSourceFile(ast.GetSourceFileOfNode(node), node, false /*includeTrivia*/)
}

func GetTextOfJSDocComment(comment *ast.NodeList) string {
	if comment == nil {
		return ""
	}
	var b strings.Builder
	for _, n := range comment.Nodes {
		switch n.Kind {
		case ast.KindJSDocText:
			b.WriteString(n.Text())
		case ast.KindJSDocLink, ast.KindJSDocLinkCode, ast.KindJSDocLinkPlain:
			b.WriteString(GetTextOfNode(n))
		}
	}
	return strings.TrimRightFunc(b.String(), unicode.IsSpace)
}

func DeclarationNameToString(name *ast.Node) string {
	if name == nil || name.Pos() == name.End() {
		return "(Missing)"
	}
	return GetTextOfNode(name)
}

// IsBareWritableName reports whether name can be written back into source as a
// bare identifier. Code that emits or inserts a name must ask this rather than
// IsIdentifierText alone: `and`, `or` and `not` are identifier-shaped but scan
// as operators, so a bare `t.and`, `and: number` or `import { and }` does not
// parse.
func IsBareWritableName(name string) bool {
	return IsIdentifierText(name) && !IsWordOperatorText(name)
}

// TokenIsLuaMethodName reports whether token can name a Lua colon-call method.
// Lua's Name grammar excludes only Lua's own reserved words; TS-only keywords
// (`new`, `type`, `delete`, ...) are valid method names, matching what the dot
// path accepts via allowIdentifierNames. Keeping Lua's reserved words out is
// what makes error recovery safe: committing a colon call on `t:end(` or
// `t:until(` would swallow the enclosing block terminator. (`and`/`or`/`not`
// scan as operator tokens and never reach here.)
//
// The keywords that open a statement whose next token can be `(` are excluded
// too: in `foo: with (x) { ... }` — a deleted TS label — committing the colon
// call `foo:with(x)` would detach the statement from its body and warp
// everything after it. Unlike `new`/`delete`, neither is a plausible Lua
// method name.
func TokenIsLuaMethodName(token ast.Kind) bool {
	switch token {
	case ast.KindBreakKeyword, ast.KindDoKeyword, ast.KindElseKeyword, ast.KindElseIfKeyword,
		ast.KindEndKeyword, ast.KindFalseKeyword, ast.KindForKeyword, ast.KindFunctionKeyword,
		ast.KindGotoKeyword, ast.KindIfKeyword, ast.KindInKeyword, ast.KindLocalKeyword,
		ast.KindNilKeyword, ast.KindRepeatKeyword, ast.KindReturnKeyword, ast.KindThenKeyword,
		ast.KindTrueKeyword, ast.KindUntilKeyword, ast.KindWhileKeyword:
		return false
	}
	return token >= ast.KindIdentifier
}

// IsLuaMethodName reports whether text can name the method of a Lua colon call:
// an identifier-shaped word that is not one of Lua's reserved words.
func IsLuaMethodName(text string) bool {
	return IsIdentifierText(text) && TokenIsLuaMethodName(GetIdentifierToken(text))
}

func IsIdentifierText(name string) bool {
	ch, size := utf8.DecodeRuneInString(name)
	if !IsIdentifierStart(ch) {
		return false
	}
	for i := size; i < len(name); {
		ch, size = utf8.DecodeRuneInString(name[i:])
		if !IsIdentifierPart(ch) {
			return false
		}
		i += size
	}
	return true
}

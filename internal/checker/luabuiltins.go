package checker

import (
	"slices"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// Shared plumbing for the Lua member-call builtins (string.format, the
// pattern-capture family, io.read/io.lines). Each builtin refines a call by
// SYMBOL IDENTITY -- the namespace global or the bundled interface member --
// so shadowed and aliased names never refine.

// luaRefiningMemberNames pre-filters the quick-path guard: only calls to
// members with these names can carry a type-changing Lua refinement.
// (format is absent: it only adds diagnostics, never changes the type.)
var luaRefiningMemberNames = []string{"match", "find", "gmatch", "read", "lines"}

// isLuaRefiningCallSyntax reports whether node LOOKS like a refining builtin
// call, by syntax alone. The quick path must not resolve symbols -- its whole
// point is to skip checking -- so it trades precision for zero resolution: a
// false positive only routes one call through the ordinary check path.
func isLuaRefiningCallSyntax(node *ast.Node) bool {
	if !ast.IsCallExpression(node) || node.QuestionDotToken() != nil {
		return false
	}
	callee := ast.SkipParentheses(node.Expression())
	return ast.IsPropertyAccessExpression(callee) && slices.Contains(luaRefiningMemberNames, callee.Name().Text())
}

type luaBuiltinCall struct {
	name          string
	callee        *ast.Node // the property-access callee, parentheses skipped
	namespaceForm bool      // string.match(s, ...) as opposed to s.match(...)
}

// resolveLuaBuiltinCall is the shared detector: a call to one of names,
// either on the namespace global or as a bundled-interface member. The
// resolvers are passed lazily so non-matching names never touch symbols.
func (c *Checker) resolveLuaBuiltinCall(node *ast.Node, names []string, globalSymbol func() *ast.Symbol, interfaceType func() *Type) *luaBuiltinCall {
	if !ast.IsCallExpression(node) || node.QuestionDotToken() != nil {
		return nil
	}
	callee := ast.SkipParentheses(node.Expression())
	if !ast.IsPropertyAccessExpression(callee) {
		return nil
	}
	name := callee.Name().Text()
	if !slices.Contains(names, name) {
		return nil
	}
	base := ast.SkipParentheses(callee.Expression())
	if ast.IsIdentifier(base) && c.isLuaGlobalReference(base, globalSymbol()) {
		return &luaBuiltinCall{name: name, callee: callee, namespaceForm: true}
	}
	member := c.getSymbolOfNameOrPropertyAccessExpression(callee)
	ifaceType := interfaceType()
	if member != nil && ifaceType != nil && c.getParentOfSymbol(c.getMergedSymbol(member)) == ifaceType.symbol && c.isLuaLibMember(member) {
		return &luaBuiltinCall{name: name, callee: callee}
	}
	return nil
}

// isLuaLibMember reports whether the symbol is declared in the bundled LuaJIT
// lib. A standard lib's String declares regex-based members of the same names
// (match, find), which must keep their declared types. The default-library
// check keeps a user file that merely shares the lib's name from spoofing it.
func (c *Checker) isLuaLibMember(symbol *ast.Symbol) bool {
	for _, declaration := range symbol.Declarations {
		if file := ast.GetSourceFileOfNode(declaration); file != nil &&
			strings.HasSuffix(file.FileName(), "lib.luajit.d.tlua") && c.program.IsSourceFileDefaultLibrary(file.Path()) {
			return true
		}
	}
	return false
}

// getLuaRefinedCallType computes the pattern/io refinement for a call ONCE and
// caches it (nil negatives included) on the node's links. getCallPackType asks
// on every flow walk that re-derives a value list, so without the cache the
// literal re-parse and pack construction would run per reference.
func (c *Checker) getLuaRefinedCallType(node *ast.Node, checkMode CheckMode) *Type {
	links := c.luaBuiltinLinks.Get(node)
	if !links.resolved {
		links.refined = core.OrElse(c.checkLuaPatternCall(node, checkMode), c.checkLuaIOCall(node, checkMode))
		links.resolved = true
	}
	return links.refined
}

// createLuaValuePack wraps element types as a closed pack; a trailing open
// tail of any models values the builtin could not analyze.
func (c *Checker) createLuaValuePack(types []*Type, openTail bool) *Type {
	if openTail {
		types = append(slices.Clip(types), c.anyType)
	}
	infos := make([]TupleElementInfo, len(types))
	for i := range infos {
		infos[i] = TupleElementInfo{flags: ElementFlagsRequired}
	}
	if openTail {
		infos[len(infos)-1] = TupleElementInfo{flags: ElementFlagsRest}
	}
	return c.createPackTypeEx(types, infos, false /*collapse*/)
}

// createLuaIteratorType wraps a result pack as the zero-parameter closure the
// generic for consumes (gmatch, io.lines).
func (c *Checker) createLuaIteratorType(returnPack *Type) *Type {
	signature := c.newSignature(SignatureFlagsNone, nil, nil, nil, nil, returnPack, nil, 0)
	return c.newAnonymousType(nil, nil, []*Signature{signature}, nil, nil)
}

package checker

import "github.com/apyrr/tlua/internal/ast"

var luaPatternNames = []string{"match", "find", "gmatch"}

type luaPatternCall struct {
	name         string
	patternIndex int
	plainIndex   int
}

func (c *Checker) getLuaPatternCall(node *ast.Node) *luaPatternCall {
	call := c.resolveLuaBuiltinCall(node, luaPatternNames, c.getLuaStringGlobalSymbol, func() *Type { return c.globalStringType })
	if call == nil {
		return nil
	}
	if call.namespaceForm {
		return &luaPatternCall{name: call.name, patternIndex: 1, plainIndex: 3}
	}
	return &luaPatternCall{name: call.name, patternIndex: 0, plainIndex: 2}
}

func (c *Checker) checkLuaPatternCall(node *ast.Node, checkMode CheckMode) *Type {
	call := c.getLuaPatternCall(node)
	if call == nil {
		return nil
	}
	args := node.Arguments()
	if call.patternIndex >= len(args) {
		return nil
	}

	if call.name == "find" && call.plainIndex < len(args) {
		plainType := c.checkExpressionCachedEx(args[call.plainIndex], checkMode)
		switch {
		case plainType.flags&TypeFlagsBooleanLiteral != 0 && getBooleanLiteralValue(plainType):
			// Plain search finds no captures, so the pattern needs no parsing.
			return c.createLuaValuePack([]*Type{c.numberOrNilType, c.numberOrNilType}, false)
		case plainType.flags&(TypeFlagsBooleanLiteral|TypeFlagsNil) != 0:
			// A literal false or nil keeps pattern semantics.
		default:
			// A plain flag only known at runtime decides whether captures
			// exist, so no refined pack is honest.
			return nil
		}
	}

	patternType := c.checkExpressionCachedEx(args[call.patternIndex], checkMode)
	if patternType.flags&TypeFlagsStringLiteral == 0 {
		return nil
	}
	types, hasCaptures, ok := c.parseLuaPatternString(getStringLiteralValue(patternType))
	if !ok {
		return nil
	}

	switch call.name {
	case "match":
		return c.createLuaValuePack(types, false)
	case "find":
		// Lua find returns no whole-match value when there are no captures.
		if !hasCaptures {
			types = nil
		}
		return c.createLuaValuePack(append([]*Type{c.numberOrNilType, c.numberOrNilType}, types...), false)
	case "gmatch":
		return c.createLuaIteratorType(c.createLuaValuePack(types, false))
	}
	return nil
}

func (c *Checker) parseLuaPatternString(pattern string) (types []*Type, hasCaptures bool, ok bool) {
	depth := 0
	parsingSet := false

	for i := 0; i < len(pattern); i++ {
		switch {
		case pattern[i] == '%':
			i++
			if !parsingSet && i < len(pattern) && pattern[i] == 'b' {
				i += 2
			}
		case !parsingSet && pattern[i] == '[':
			parsingSet = true
			if i+1 < len(pattern) && pattern[i+1] == ']' {
				i++
			}
		case parsingSet && pattern[i] == ']':
			parsingSet = false
		case pattern[i] == '(':
			if parsingSet {
				continue
			}
			hasCaptures = true
			if i+1 < len(pattern) && pattern[i+1] == ')' {
				i++
				types = append(types, c.numberOrNilType)
				continue
			}
			depth++
			types = append(types, c.stringOrNilType)
		case pattern[i] == ')':
			if parsingSet {
				continue
			}
			depth--
			if depth < 0 {
				return nil, false, false
			}
		}
	}

	if depth != 0 || parsingSet {
		return nil, false, false
	}
	if len(types) == 0 {
		types = append(types, c.stringOrNilType)
	}
	return types, hasCaptures, true
}

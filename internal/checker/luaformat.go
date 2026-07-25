package checker

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/diagnostics"
)

var luaFormatNames = []string{"format"}

// getLuaFormatCall returns the literal format expression and the number of
// effective call arguments preceding its substitutions.
func (c *Checker) getLuaFormatCall(node *ast.Node) (*ast.Node, int) {
	call := c.resolveLuaBuiltinCall(node, luaFormatNames, c.getLuaStringGlobalSymbol, func() *Type { return c.globalStringType })
	if call == nil {
		return nil, 0
	}
	// The format string is the function's first parameter in every spelling: the
	// namespace form writes it (`string.format(f, x)`), and the member form IS it
	// (`f:format(x)` -- the receiver and the format string are the same text). So a
	// colon call takes it off the written list, and the receiver supplies it.
	//
	// The returned offset counts EFFECTIVE arguments before the substitutions, and
	// getEffectiveCallArguments prepends a colon call's receiver. Either way the
	// format string lands at effective 0, so the substitutions always start at 1.
	if index := ast.LuaWrittenArgumentIndex(node, 0); index >= 0 {
		args := node.Arguments()
		if index >= len(args) {
			return nil, 0
		}
		return args[index], 1
	}
	return call.callee.Expression(), 1
}

func (c *Checker) checkLuaFormatCall(node *ast.Node, checkMode CheckMode) {
	formatExpression, argumentOffset := c.getLuaFormatCall(node)
	if formatExpression == nil {
		return
	}
	formatType := c.checkExpressionCachedEx(formatExpression, checkMode)
	if formatType.flags&TypeFlagsStringLiteral == 0 {
		return
	}
	specifiers := c.parseLuaFormatString(getStringLiteralValue(formatType))
	// Recomputing the effective arguments (resolveCall built them once
	// already) is bounded to literal-format calls, and the expensive part --
	// a tail call's pack resolution -- is cached on the tail's own links.
	args := c.getEffectiveCallArguments(node)
	expectedCount := len(specifiers) + argumentOffset
	actualCount := len(args)
	spreadIndex := c.getSpreadArgumentIndex(args)
	openTail := spreadIndex >= 0 && spreadIndex == actualCount-1
	if actualCount != expectedCount && !(actualCount < expectedCount && openTail) {
		// Both counts are effective, so a colon call's receiver has to come back out
		// of each before the reader sees it: `("%d"):format(1, 2)` wrote two.
		unwritten := ast.LuaImplicitArgumentCount(node)
		c.error(node, diagnostics.Expected_0_arguments_but_got_1, expectedCount-unwritten, actualCount-unwritten)
	}
	for i, specifierType := range specifiers {
		argumentIndex := i + argumentOffset
		if argumentIndex >= actualCount || isSpreadArgument(args[argumentIndex]) {
			break
		}
		argument := args[argumentIndex]
		argumentType := c.checkExpressionCachedEx(argument, checkMode)
		c.checkTypeAssignableTo(argumentType, specifierType, argument, diagnostics.Argument_of_type_0_is_not_assignable_to_parameter_of_type_1)
	}
}

func (c *Checker) parseLuaFormatString(format string) []*Type {
	var result []*Type
	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			continue
		}
		i++
		if i < len(format) && format[i] == '%' {
			continue
		}
		for i < len(format) && !isLuaFormatLetter(format[i]) && format[i] != '*' {
			i++
		}
		if i == len(format) {
			break
		}
		switch format[i] {
		case 's':
			// LuaJIT's %s applies tostring (and __tostring) to any value, so
			// unlike Luau it constrains nothing.
			result = append(result, c.anyType)
		case 'q':
			// %q quotes a string; a number coerces, everything else errors.
			result = append(result, c.getUnionType([]*Type{c.stringType, c.numberType}))
		case 'c', 'd', 'i', 'o', 'u', 'x', 'X', 'e', 'E', 'f', 'g', 'G':
			result = append(result, c.numberType)
		default:
			result = append(result, c.anyType)
		}
	}
	return result
}

func isLuaFormatLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

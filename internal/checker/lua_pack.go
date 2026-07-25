package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// Lua value packs. A call or vararg yields a list of values whose arity is part
// of its type; every other expression yields exactly one. These helpers project
// and adjust that arity without disturbing the element types.

func (c *Checker) adjustLuaParenthesizedValue(t *Type) *Type {
	// Parentheses are an arity barrier in Lua: even a zero-result call becomes
	// exactly one nil value when it is parenthesized.
	return c.mapType(t, func(arm *Type) *Type {
		return core.IfElse(arm.flags&TypeFlagsVoid != 0, c.nilType, arm)
	})
}

// luaPackProducerAndWrappers separates erased, arity-transparent wrappers from
// their producer. Parentheses deliberately remain an arity barrier.
func luaPackProducerAndWrappers(expr *ast.Node) (*ast.Node, []*ast.Node) {
	var wrappers []*ast.Node
	for {
		switch expr.Kind {
		case ast.KindNonNullExpression, ast.KindTypeAssertionExpression, ast.KindAsExpression, ast.KindSatisfiesExpression:
			wrappers = append(wrappers, expr)
			expr = expr.Expression()
		default:
			return expr, wrappers
		}
	}
}

// applyLuaPackWrappers maps only slot zero while preserving the producer tail.
// Argument probing checks wrappers because synthetics replace the original AST.
func (c *Checker) applyLuaPackWrappers(pack *Type, wrappers []*ast.Node, checkMode CheckMode, ensureChecked bool) *Type {
	if pack == nil {
		return nil
	}
	if ensureChecked && len(wrappers) != 0 {
		c.checkExpressionCachedEx(wrappers[0], checkMode)
	}
	for i := len(wrappers) - 1; i >= 0; i-- {
		wrapper := wrappers[i]
		switch wrapper.Kind {
		case ast.KindNonNullExpression:
			pack = c.mapLuaPackFirstType(pack, c.GetNonNullableType)
		case ast.KindTypeAssertionExpression, ast.KindAsExpression:
			asserted := c.checkExpressionCachedEx(wrapper, checkMode)
			pack = c.mapLuaPackFirstType(pack, func(*Type) *Type { return asserted })
		case ast.KindSatisfiesExpression:
			// `satisfies` checks the scalar view but does not change it.
		}
	}
	return pack
}

// mapLuaPackFirstType preserves a producer's arity and tail while changing the
// scalar value an erased assertion exposes at position zero.
func (c *Checker) mapLuaPackFirstType(pack *Type, mapFirst func(*Type) *Type) *Type {
	if pack == nil {
		return nil
	}
	return c.mapType(pack, func(arm *Type) *Type {
		if arm.flags&TypeFlagsVoid != 0 {
			return arm
		}
		if !isPackType(arm) {
			return mapFirst(arm)
		}
		elementTypes := slices.Clone(c.getTypeArguments(arm))
		elementInfos := slices.Clone(arm.TargetTupleType().elementInfos)
		if len(elementTypes) == 0 {
			return arm
		}
		switch {
		case elementInfos[0].flags&ElementFlagsRest != 0:
			// A homogeneous open pack may still be empty. Split its asserted
			// first value from the unchanged remainder without forcing arity.
			return c.createPackTypeEx(
				[]*Type{mapFirst(elementTypes[0]), elementTypes[0]},
				[]TupleElementInfo{{flags: ElementFlagsOptional}, elementInfos[0]},
				false, /*collapse*/
			)
		case elementInfos[0].flags&ElementFlagsVariadic != 0:
			// Tail<A> is not representable directly. Conservatively retain a
			// zero-or-more shape with an optional refined head and a homogeneous
			// tail; do not prefix A, which would add a fictitious runtime value.
			first := c.packElementForIndex(arm, 0)
			tail := c.getIndexedAccessType(elementTypes[0], c.numberType)
			return c.createPackTypeEx(
				[]*Type{mapFirst(first), tail},
				[]TupleElementInfo{{flags: ElementFlagsOptional}, {flags: ElementFlagsRest}},
				false, /*collapse*/
			)
		default:
			elementTypes[0] = mapFirst(elementTypes[0])
			return c.createPackTypeEx(elementTypes, elementInfos, false /*collapse*/)
		}
	})
}

package checker

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/jsnum"
)

// select(n, ...) returns the values of its vararg from position n onward. Which
// values those are depends on the *value* of a literal index, not on any type the
// declared signature can name (Luau special-cases it the same way, as a magic
// function). The lib deliberately declares select non-generically -- `(index:
// number, ...: any): (any?, ...any)` -- as the safe fallback for a shadowed or
// aliased select; this refinement replaces that type for calls to the real global.
//
// Only a call to the *global* `select` is refined: a shadowing local, or a call
// through an alias, resolves to a different symbol and keeps the declared type, as
// the metatable and type() builtins do. The `"#"` form is a separate overload
// returning `number`, so a non-numeric index is left alone here.

// isLuaSelectCall reports whether the refinement interprets this call. The quick
// path of getQuickTypeOfExpression must skip such a call: it reads the signature's
// return type without running checkCallExpression's special cases. (Today select's
// two overloads already defeat that quick path's getSingleSignature, but that is an
// accident of the lib's shape, not a guarantee.)
func (c *Checker) isLuaSelectCall(node *ast.Node) bool {
	if !ast.IsCallExpression(node) || node.QuestionDotToken() != nil {
		return false
	}
	callee := ast.SkipParentheses(node.Expression())
	return ast.IsIdentifier(callee) && callee.Text() == "select" && c.isLuaGlobalReference(callee, c.getLuaSelectGlobalSymbol())
}

// checkLuaSelectCall returns the refined pack of a select call, or nil when the
// declared type stands. The caller has already checked the call's arguments, so
// the cached expression types are used; checkMode is threaded through so a select
// inside an inferential pass does not fix types in the normal-mode cache.
func (c *Checker) checkLuaSelectCall(node *ast.Node, checkMode CheckMode) *Type {
	if !c.isLuaSelectCall(node) {
		return nil
	}
	args := node.Arguments()
	if len(args) < 1 {
		return nil
	}
	indexType := c.checkExpressionCachedEx(args[0], checkMode)
	// `select("#", ...)` is the other overload; only a numeric index selects values.
	if indexType.flags&TypeFlagsNumberLike == 0 {
		return nil
	}
	values := args[1:]
	// A trailing `...` or multi-return call contributes several values at once. A
	// void tail contributes none and drops out here. The tail may also be a UNION
	// of packs (a call through a union-typed function), which only
	// luaSelectTailValues below can read -- TargetTupleType on a union is a panic.
	var tailPack *Type
	if len(values) != 0 {
		tailPack = c.getCallPackType(values[len(values)-1], checkMode)
		if tailPack != nil && tailPack.flags&TypeFlagsVoid != 0 {
			values = values[:len(values)-1]
			tailPack = nil
		}
	}
	fixedCount := len(values)
	if tailPack != nil {
		fixedCount--
	}

	if index, ok := literalIntegerValue(indexType); ok && index >= 1 {
		if index-1 <= fixedCount {
			// The slice starts in the fixed values: those keep their positions, and
			// the tail (if any) follows in full.
			var elementTypes []*Type
			var elementInfos []TupleElementInfo
			for i := index - 1; i < fixedCount; i++ {
				elementTypes = append(elementTypes, c.checkExpressionCachedEx(values[i], checkMode))
				elementInfos = append(elementInfos, TupleElementInfo{flags: ElementFlagsRequired})
			}
			if tailPack != nil {
				if isPackType(tailPack) {
					// A single pack tail keeps its exact shape.
					elementTypes = append(elementTypes, c.getElementTypes(tailPack)...)
					elementInfos = append(elementInfos, tailPack.TargetTupleType().elementInfos...)
				} else {
					// A union of packs: position-blind past the fixed slice, but sound.
					elementTypes = append(elementTypes, c.luaSelectTailValues(tailPack, 0))
					elementInfos = append(elementInfos, TupleElementInfo{flags: ElementFlagsRest})
				}
			}
			if len(elementTypes) == 0 {
				// Past the end: no values, i.e. `void`.
				return c.voidType
			}
			// A real pack, never collapsed to a scalar: this result flows through
			// getCallPackType (a value list, a return), whose consumers require a
			// pack even for a single value. A single-value slice collapsed to a
			// scalar would be wrapped as a variadic element of a non-pack type and
			// normalize to `...any`.
			return c.createPackTypeEx(elementTypes, elementInfos, false /*collapse*/)
		}
		if tailPack == nil {
			// Past the end of a closed value list: no values.
			return c.voidType
		}
		// The index lands inside the tail: the selected values are the tail's,
		// from that offset on. Positions inside the tail are not re-derivable
		// (its own fixed prefix may repeat under a union), so the sound open
		// pack of the remaining values stands in.
		selected := c.luaSelectTailValues(tailPack, index-1-fixedCount)
		if selected.flags&TypeFlagsNever != 0 {
			return c.voidType
		}
		return c.newRestPack(selected)
	}

	// A non-literal index cannot pin positions, so the result is the sound open
	// pack of whatever the values could be: index 1 selects them all, so any value
	// may occupy any slot. `nil` past the end is the open pack's ordinary read.
	var members []*Type
	for i := range fixedCount {
		members = append(members, c.checkExpressionCachedEx(values[i], checkMode))
	}
	if tailPack != nil {
		members = append(members, c.luaSelectTailValues(tailPack, 0))
	}
	if len(members) == 0 {
		return c.voidType
	}
	return c.newRestPack(c.getUnionType(members))
}

// luaSelectTailValues is the union of every value a call-pack tail can produce at
// or after offset, distributed over a union of packs. A void arm, or a single
// value the offset skips past, contributes nothing.
func (c *Checker) luaSelectTailValues(pack *Type, offset int) *Type {
	return c.mapType(pack, func(s *Type) *Type {
		switch {
		case isPackType(s):
			return c.getPackValuesFromOffset(s, offset)
		case s.flags&TypeFlagsVoid != 0:
			return c.neverType
		case offset == 0:
			return s
		default:
			return c.neverType
		}
	})
}

// literalIntegerValue returns the value of an integer number-literal type. Bounds
// (Lua's select requires a non-zero index) are the caller's business.
func literalIntegerValue(t *Type) (int, bool) {
	if t.flags&TypeFlagsNumberLiteral == 0 {
		return 0, false
	}
	n := getNumberLiteralValue(t)
	i := int(n)
	if jsnum.Number(i) != n {
		return 0, false // not an integer
	}
	return i, true
}

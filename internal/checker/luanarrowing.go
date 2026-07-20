package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// Lua discriminates values with the `type` function instead of a `typeof`
// operator, and file handles with `io.type`. Both are ordinary global functions,
// so a guard only narrows when its callee resolves to the real global.
type luaGuardKind int32

const (
	luaGuardNone luaGuardKind = iota
	luaGuardType
	luaGuardIOType
)

// The tags `type()` can return, in the order they are declared in lualibs/global.d.ts.
var luaTypeTagNames = []string{"nil", "boolean", "number", "string", "table", "function", "thread", "userdata", "cdata"}

// The string tags `io.type()` can return. It returns nil for a non-file value.
var luaFileTagNames = []string{"file", "closed file"}

// getLuaTypeGuardCall returns the guarded argument of a `type(x)` or `io.type(x)`
// call, provided the callee resolves to the corresponding global. A shadowing
// `local type` or an alias `local t = type` therefore never narrows.
func (c *Checker) getLuaTypeGuardCall(expr *ast.Node) (luaGuardKind, *ast.Node) {
	callee, argument := ast.LuaTypeGuardCall(expr)
	if callee == nil {
		return luaGuardNone, nil
	}
	if ast.IsIdentifier(callee) {
		if c.isLuaGlobalReference(callee, c.getLuaTypeGlobalSymbol()) {
			return luaGuardType, argument
		}
		return luaGuardNone, nil
	}
	if c.isLuaGlobalReference(ast.SkipParentheses(callee.Expression()), c.getLuaIOGlobalSymbol()) {
		return luaGuardIOType, argument
	}
	return luaGuardNone, nil
}

// isLuaGlobalReference reports whether name resolves to the given global. A global
// that the program does not declare matches nothing, including an unresolved name.
// Only an identifier can reference a global: any other expression (for example the
// `f()` in `f().setmetatable`) matches nothing rather than being resolved.
func (c *Checker) isLuaGlobalReference(name *ast.Node, global *ast.Symbol) bool {
	return global != nil && ast.IsIdentifier(name) && c.getMergedSymbol(c.getResolvedSymbol(name)) == global
}

// isLuaTypeGuardLiteral reports whether a comparison operand is one a type guard
// can be tested against: a tag string, or `nil` for the absent result of io.type.
func isLuaTypeGuardLiteral(node *ast.Node) bool {
	return ast.IsStringLiteralLike(node) || node.Kind == ast.KindNilKeyword
}

// isLuaAbsentGuardResult reports whether literal is the guard's result for a nil
// argument: `type(nil)` is the string "nil", `io.type(nil)` is nil itself.
func isLuaAbsentGuardResult(guard luaGuardKind, literal *ast.Node) bool {
	if guard == luaGuardIOType {
		return literal.Kind == ast.KindNilKeyword
	}
	return ast.IsStringLiteralLike(literal) && literal.Text() == "nil"
}

// narrowTypeByLuaGuardLiteral narrows t by the result of a type guard, which is a
// tag string or (for io.type) nil.
func (c *Checker) narrowTypeByLuaGuardLiteral(t *Type, guard luaGuardKind, literal *ast.Node, assumeTrue bool) *Type {
	if literal.Kind == ast.KindNilKeyword {
		if guard == luaGuardIOType {
			// io.type returns nil exactly for the values that are not file handles.
			return c.narrowTypeByLuaFile(t, !assumeTrue)
		}
		// `type()` always returns a string, so the equality can never hold.
		return core.IfElse(assumeTrue, c.neverType, t)
	}
	return c.narrowTypeByLuaTag(t, guard, literal.Text(), assumeTrue)
}

// narrowTypeByLuaTag narrows t by a tag string the guard reported (assumeTrue) or
// did not report (!assumeTrue).
func (c *Checker) narrowTypeByLuaTag(t *Type, guard luaGuardKind, tag string, assumeTrue bool) *Type {
	if guard == luaGuardIOType {
		if !slices.Contains(luaFileTagNames, tag) {
			return core.IfElse(assumeTrue, c.neverType, t)
		}
		if !assumeTrue {
			// A handle that is not open may still be a closed file, so a single
			// negative tag rules nothing out.
			return t
		}
		return c.narrowTypeByLuaFile(t, true /*assumeTrue*/)
	}
	switch tag {
	case "nil":
		if assumeTrue {
			return c.narrowTypeByTypeFacts(t, c.nilType, TypeFactsEQUndefined)
		}
		return c.getAdjustedTypeWithFacts(t, TypeFactsNEUndefined)
	case "boolean":
		if assumeTrue {
			return c.narrowTypeByTypeFacts(t, c.booleanType, TypeFactsTypeofEQBoolean)
		}
		return c.getAdjustedTypeWithFacts(t, TypeFactsTypeofNEBoolean)
	case "number":
		if assumeTrue {
			return c.narrowTypeByTypeFacts(t, c.numberType, TypeFactsTypeofEQNumber)
		}
		return c.getAdjustedTypeWithFacts(t, TypeFactsTypeofNENumber)
	case "string":
		if assumeTrue {
			return c.narrowTypeByTypeFacts(t, c.stringType, TypeFactsTypeofEQString)
		}
		return c.getAdjustedTypeWithFacts(t, TypeFactsTypeofNEString)
	case "function":
		if assumeTrue {
			if t.flags&TypeFlagsAny != 0 {
				return t
			}
			// The facts filter keeps a union's own function constituents with their
			// declared signatures; the `function` top type only stands in where nothing
			// declared can be kept -- unknown says nothing to filter, and a bare
			// instantiable (a type parameter) gains callability only by intersection.
			// narrowTypeByTypeFacts is unusable here: its candidate-substitution arm
			// replaces precise signatures with the all-any top type.
			if t.flags&TypeFlagsUnknown != 0 {
				return c.getLuaFunctionType()
			}
			return c.mapType(t, func(s *Type) *Type {
				if !c.hasTypeFacts(s, TypeFactsTypeofEQFunction) {
					return c.neverType
				}
				if s.flags&TypeFlagsInstantiable != 0 && !c.isTypeSubtypeOf(s, c.getLuaFunctionType()) {
					return c.getIntersectionType([]*Type{s, c.getLuaFunctionType()})
				}
				return s
			})
		}
		return c.getAdjustedTypeWithFacts(t, TypeFactsTypeofNEFunction)
	case "table":
		return c.narrowTypeByLuaTable(t, assumeTrue)
	case "thread":
		return c.narrowTypeByLuaBrand(t, c.getLuaThreadType(), assumeTrue)
	case "userdata":
		return c.narrowTypeByLuaBrand(t, c.getLuaUserdataType(), assumeTrue)
	case "cdata":
		return c.narrowTypeByLuaBrand(t, c.getLuaCDataType(), assumeTrue)
	}
	// `type()` returns no other tag, so the equality can never hold.
	return core.IfElse(assumeTrue, c.neverType, t)
}

// narrowTypeByLuaTable narrows by the "table" tag. Type facts cannot express it:
// every object type carries the same object-ish facts, including the branded
// thread/userdata/cdata types that `type()` reports under their own tag, so those
// are filtered out by subtype relation before the object facts apply. Unlike the
// JavaScript "object" tag, "table" does not include nil.
func (c *Checker) narrowTypeByLuaTable(t *Type, assumeTrue bool) *Type {
	if t.flags&TypeFlagsAny != 0 {
		return t
	}
	if assumeTrue {
		return c.narrowTypeByTypeFacts(c.filterType(t, func(s *Type) bool { return !c.isLuaBrandedType(s) }), c.nonPrimitiveType, TypeFactsTypeofEQObject)
	}
	return c.mapType(t, func(s *Type) *Type {
		if c.isLuaBrandedType(s) {
			return s
		}
		return c.getAdjustedTypeWithFacts(s, TypeFactsTypeofNEObject)
	})
}

// narrowTypeByLuaBrand narrows by one of the branded object tags. The facts route
// is unusable for them (see narrowTypeByLuaTable), so both polarities go through
// the subtype filter that backs type predicates -- in disjoint mode, because the
// runtime tags partition values: a type unrelated to the brand narrows to never
// instead of an intersection, so a provably false guard is a dead branch. An
// unresolvable brand - its declaration is not in the program - narrows nothing.
func (c *Checker) narrowTypeByLuaBrand(t *Type, brand *Type, assumeTrue bool) *Type {
	if brand == nil {
		return t
	}
	return c.getNarrowedTypeEx(t, brand, assumeTrue, false /*checkDerived*/, true /*disjoint*/)
}

func (c *Checker) narrowTypeByLuaFile(t *Type, assumeTrue bool) *Type {
	return c.narrowTypeByLuaBrand(t, c.getLuaFileType(), assumeTrue)
}

// isLuaTableType reports whether an object- or intersection-flagged type is a plain Lua
// table -- what the `table` keyword type admits. A setmetatable pairing always is: a real
// table by construction, even when __call makes it callable. Anything else must be
// non-callable (a function is not a table) and carry no brand member (a thread, userdata or
// cdata is not a table either), matching what `type(x) == "table"` narrowing believes. An
// intersection answers as a whole: one plain-table constituent must not vouch for a callable
// or branded partner. The verdict is cached per type because a repeat probe would re-pay the
// member resolution.
func (c *Checker) isLuaTableType(source *Type) bool {
	if isMetatableType(source) {
		return true
	}
	// A reference instantiation shares its generic target's verdict: instantiation maps
	// members one-to-one and cannot add call/construct signatures or brand members, so one
	// probe per generic serves every instantiation. A generic interface's declared type is
	// a reference to ITSELF -- that one probes directly.
	if source.objectFlags&ObjectFlagsReference != 0 {
		if target := source.Target(); target != source {
			return c.isLuaTableType(target)
		}
	}
	if source.flags&TypeFlagsIntersection != 0 {
		// A setmetatable pairing constituent makes the intersection a real table by
		// construction, whatever else it is intersected with...
		if core.Some(source.Types(), isMetatableType) {
			return true
		}
		// ...while a primitive constituent (the branded-string idiom `string & { tag }`)
		// means the value is that primitive at runtime -- `type()` never says "table".
		if core.Some(source.Types(), func(t *Type) bool { return t.flags&TypeFlagsPrimitive != 0 }) {
			return false
		}
	}
	if cached, ok := c.luaTableTypeResults[source]; ok {
		return cached
	}
	result := !c.typeHasCallOrConstructSignatures(source) && !c.hasLuaBrandMember(source)
	// A type queried mid-resolution answers from a partial member snapshot; caching that
	// would make the transient answer permanent (the hazard PropertiesTypesKey folds into
	// its cache key).
	if source.objectFlags&ObjectFlagsUnresolvedMembers == 0 {
		c.luaTableTypeResults[source] = result
	}
	return result
}

// hasLuaBrandMember reports whether the type carries one of the unique-symbol brand members
// that mark the branded non-table object kinds. Exclusion works by member IDENTITY: the
// interned unique-symbol type survives inheritance, while a lookalike field of another type
// stays a plain table. Unlike isLuaBrandedType's subtype checks, a property lookup plus
// pointer compare cannot re-enter the relation machinery this is called from.
func (c *Checker) hasLuaBrandMember(source *Type) bool {
	for name, brandType := range c.getLuaBrandMembers() {
		prop := c.getPropertyOfTypeEx(source, name, false /*includeTypeOnlyMembers*/)
		if prop != nil && c.getTypeOfSymbol(prop) == brandType {
			return true
		}
	}
	return false
}

// isLuaBrandedType reports whether s is one of the branded non-table object kinds
// that `type()` reports as "thread", "userdata", or "cdata" rather than "table".
func (c *Checker) isLuaBrandedType(s *Type) bool {
	base := c.getBaseConstraintOrType(s)
	return slices.ContainsFunc(c.getLuaBrandTypes(), func(brand *Type) bool {
		return c.isTypeSubtypeOf(base, brand)
	})
}

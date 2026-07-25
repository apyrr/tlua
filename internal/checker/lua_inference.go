package checker

import (
	"maps"
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// Typing of the values Lua assignments install. A multiple assignment captures
// its whole value list before storing any of it, so a target reads its siblings
// from the pre-store snapshot rather than from the finished program.

func (c *Checker) getLuaAugmentationInitializerType(assignment luaAugmentation) *Type {
	return c.getLuaAugmentationInitializerTypeEx(assignment, make(map[luaSnapshotKey]bool))
}

func (c *Checker) getLuaAugmentationTargetInitializerType(symbol *ast.Symbol, target *ast.Node) (*Type, bool) {
	for _, assignment := range c.luaAssignmentAugmentations[c.getMergedSymbol(symbol)] {
		if assignment.Target == target {
			return c.getLuaAugmentationInitializerType(assignment), true
		}
	}
	return nil, false
}

type luaSnapshotKey struct {
	symbol *ast.Symbol
	source *ast.Node
}

func (c *Checker) getLuaAugmentationInitializerTypeEx(assignment luaAugmentation, resolving map[luaSnapshotKey]bool) *Type {
	binary := assignment.Source.AsBinaryExpression()
	if initializer := luaExplicitAssignmentValueAt(binary.Right, assignment.ValueIndex); initializer != nil {
		if defaulted := c.luaDefaultedAugmentationInitializer(assignment.Target, initializer); defaulted != initializer {
			t := c.checkExpressionWithContextualType(defaulted, nil, nil, CheckModeNormal)
			t = c.packElementForIndex(t, 0)
			t = c.getWidenedLiteralLikeTypeForContextualType(t, nil)
			return c.finalizeLuaAugmentationInitializerType(assignment, defaulted, t)
		}
		if t, ok := c.getLuaCapturedSnapshotValueType(assignment, initializer, resolving); ok {
			return c.finalizeLuaAugmentationInitializerType(assignment, initializer, t)
		}
	}
	t := c.getLuaAssignmentValueType(binary.Right, assignment.ValueIndex, CheckModeNormal)
	return c.finalizeLuaAugmentationInitializerType(assignment, luaExplicitAssignmentValueAt(binary.Right, assignment.ValueIndex), t)
}

func luaConstructorTypesShareMembers(left *Type, right *Type) bool {
	if left.flags&TypeFlagsObject == 0 || right.flags&TypeFlagsObject == 0 ||
		left.objectFlags&ObjectFlagsMembersResolved == 0 || right.objectFlags&ObjectFlagsMembersResolved == 0 {
		return false
	}
	leftStructured := left.AsStructuredType()
	rightStructured := right.AsStructuredType()
	return maps.Equal(leftStructured.members, rightStructured.members) &&
		slices.Equal(leftStructured.signatures, rightStructured.signatures) &&
		slices.Equal(leftStructured.indexInfos, rightStructured.indexInfos)
}

func (c *Checker) appendLuaConstructorType(types []*Type, candidate *Type, dedupeMembers bool) []*Type {
	if dedupeMembers && core.Some(types, func(existing *Type) bool {
		return luaConstructorTypesShareMembers(existing, candidate)
	}) {
		return types
	}
	return core.AppendIfUnique(types, candidate)
}

func luaObjectLiteralConstructor(initializer *ast.Node) *ast.Node {
	if initializer == nil {
		return nil
	}
	constructor := skipLuaRuntimeTransparentWrappers(initializer)
	if ast.IsObjectLiteralExpression(constructor) {
		return constructor
	}
	return nil
}

func hasLuaTypeAssertionWrapper(initializer *ast.Node) bool {
	for ast.IsOuterExpression(initializer, ast.OEKParentheses|ast.OEKAssertions) {
		if initializer.Kind == ast.KindTypeAssertionExpression || ast.IsAsExpression(initializer) {
			return true
		}
		initializer = initializer.Expression()
	}
	return false
}

// getLuaCapturedSnapshotValueType projects a direct target read from the
// transaction's pre-store snapshot. This breaks declaration cycles for swaps
// without pretending the assignments execute sequentially.
func (c *Checker) getLuaCapturedSnapshotValueType(assignment luaAugmentation, value *ast.Node, resolving map[luaSnapshotKey]bool) (*Type, bool) {
	slot, ok := luaAssignmentSlotForNode(assignment.Target)
	if !ok {
		return nil, false
	}
	if c.isLuaCapturedReceiverValue(assignment.Target, value) {
		return nil, false
	}
	var symbols []*ast.Symbol
	if target, captured := slot.capturedTargetForValue(c, value); captured {
		if symbol := c.luaAugmentationTargets[target]; symbol != nil {
			symbols = append(symbols, symbol)
		}
	} else {
		reference := skipLuaRuntimeTransparentWrappers(value)
		var known bool
		symbols, known = c.resolveLuaConstructors().referenceSymbols(reference)
		if !known {
			return nil, false
		}
	}
	if len(symbols) == 0 {
		return nil, false
	}
	var types []*Type
	for _, symbol := range symbols {
		if symbol == nil || symbol == c.unknownSymbol {
			return nil, false
		}
		t, ok := c.getLuaAugmentationSnapshotType(symbol, assignment.Source, resolving)
		if !ok {
			return nil, false
		}
		types = core.AppendIfUnique(types, t)
	}
	t := c.getWidenedType(c.getUnionType(types))
	return c.applyLuaSnapshotValueWrappers(value, t), true
}

func (c *Checker) getLuaAugmentationSnapshotType(symbol *ast.Symbol, source *ast.Node, resolving map[luaSnapshotKey]bool) (*Type, bool) {
	symbol = c.getMergedSymbol(symbol)
	assignments := c.luaAssignmentAugmentations[symbol]
	if len(assignments) == 0 {
		return nil, false
	}
	key := luaSnapshotKey{symbol: symbol, source: source}
	if resolving[key] {
		// A cycle wholly predating this transaction has no more precise snapshot.
		return c.anyType, true
	}
	resolving[key] = true
	defer delete(resolving, key)

	types := c.appendLuaAssignmentDeclaredTypes(nil, symbol, assignments, source, resolving)
	if len(types) == 0 {
		if !hasLuaValueDeclarationOutsideAssignments(symbol, assignments) {
			// An implicit global exists as nil before its first runtime store.
			return c.nilType, true
		}
		return nil, false
	}
	return c.getWidenedType(c.getUnionType(types)), true
}

// appendLuaAssignmentDeclaredTypes appends the constructor type that every
// effective store to symbol installs. excludedSource drops one transaction's own
// stores: Lua evaluates a whole value list before storing any of it, so the value
// that statement captured cannot include them. Dropping them is what makes a swap
// resolve instead of looking circular; a cycle across statements still reaches
// the guard in getLuaAugmentationSnapshotType.
func (c *Checker) appendLuaAssignmentDeclaredTypes(types []*Type, symbol *ast.Symbol, assignments []luaAugmentation, excludedSource *ast.Node, resolving map[luaSnapshotKey]bool) []*Type {
	dedupeConstructorTypes := c.shouldDedupeLuaConstructorTypes(symbol, assignments)
	for _, assignment := range c.effectiveLuaAssignmentAugmentations(assignments) {
		// A self-store changes flow only when its erased wrappers narrow the
		// captured value; it never contributes a new declared type.
		if assignment.Source == excludedSource || c.isSelfPreservingLuaCapturedTarget(assignment.Target) {
			continue
		}
		types = c.appendLuaConstructorType(types, c.getLuaAugmentationInitializerTypeEx(assignment, resolving), dedupeConstructorTypes)
	}
	return types
}

func hasLuaValueDeclarationOutsideAssignments(symbol *ast.Symbol, assignments []luaAugmentation) bool {
	assignmentDeclarations := luaAssignmentDeclarationSet(assignments)
	for _, declaration := range symbol.Declarations {
		if assignmentDeclarations.Has(declaration) {
			continue
		}
		if raw := declaration.Symbol(); raw != nil && raw.Flags&ast.SymbolFlagsValue != 0 {
			return true
		}
	}
	return false
}

func (c *Checker) applyLuaSnapshotValueWrappers(value *ast.Node, t *Type) *Type {
	switch value.Kind {
	case ast.KindParenthesizedExpression, ast.KindSatisfiesExpression:
		return c.applyLuaSnapshotValueWrappers(value.Expression(), t)
	case ast.KindNonNullExpression:
		return c.GetNonNullableType(c.applyLuaSnapshotValueWrappers(value.Expression(), t))
	case ast.KindTypeAssertionExpression, ast.KindAsExpression:
		if ast.IsConstAssertion(value) {
			return c.applyLuaSnapshotValueWrappers(value.Expression(), t)
		}
		return c.getTypeFromTypeNode(value.Type())
	default:
		return t
	}
}

func (c *Checker) checkLuaAugmentationEmptyArrayType(assignment luaAugmentation, t *Type) *Type {
	if c.isEmptyArrayLiteralType(t) && !c.hasParentWithTypeAnnotation(assignment.Symbol) {
		c.reportImplicitAny(assignment.Source, c.anyArrayType, WideningKindNormal)
		return c.anyArrayType
	}
	return t
}

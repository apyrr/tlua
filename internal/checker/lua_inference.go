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
	return c.getLuaAugmentationInitializerTypeEx(assignment, c.luaSnapshotResolving)
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

// getLuaSelfReadSnapshotType projects a reference that reads its own store's
// target while that target's declared type is still resolving. The statement
// captures its whole value list before storing any of it, so the read sees the
// pre-store snapshot — what the other stores install — rather than re-entering
// its own resolution and reporting a false circularity. The whole-value
// capture and the exact defaulted guard already sidestep such reads without
// checking them; this covers a self-read nested anywhere else in the value,
// such as `a = (a or 0) + 1`. A resolved target keeps the ordinary flow path,
// which respects statement order.
//
// The snapshot itself is statement-ordered: only stores that can have executed
// before this one contribute, and nil stays in the union unless one of them
// definitely has. `x = x + 1; x = 0` therefore reads nil (the later store has
// not run), while `total = 0; total = total + 1` reads number — mirroring the
// order sensitivity the ordinary flow path applies to reads outside a store.
func (c *Checker) getLuaSelfReadSnapshotType(reference *ast.Node, symbol *ast.Symbol) (*Type, bool) {
	merged := c.getMergedSymbol(symbol)
	source := c.luaSelfStoreForRead(reference, merged)
	if source == nil {
		return nil, false
	}
	// A resolved target keeps the ordinary flow path. Gating on the cached
	// declared type rather than on a resolution cycle matters when the read
	// sits inside an immediately invoked value: the first entry arrives
	// through the IIFE's return-type resolution, before the target's own
	// resolution is on the stack, and falling through would complete a
	// return-type/declared-type cycle instead of reading the snapshot.
	if c.valueSymbolLinks.Get(symbol).resolvedType != nil || c.valueSymbolLinks.Get(merged).resolvedType != nil {
		return nil, false
	}
	return c.getLuaOrderedSelfSnapshotType(merged, source, c.luaSnapshotResolving)
}

// getLuaOrderedSelfSnapshotType is getLuaAugmentationSnapshotType restricted to
// the stores that can precede `source` at runtime, nil-extended when none of
// them definitely does. The unordered variant stays in use for sibling captures
// inside one transaction, where the whole value list is taken before any store
// and order between statements is not in question.
func (c *Checker) getLuaOrderedSelfSnapshotType(symbol *ast.Symbol, source *ast.Node, resolving map[luaSnapshotKey]bool) (*Type, bool) {
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

	prior, definite := c.luaStoresBeforeSelfRead(source, assignments)
	types := c.appendLuaAssignmentDeclaredTypes(nil, symbol, prior, source, resolving)
	implicit := !hasLuaValueDeclarationOutsideAssignments(symbol, assignments)
	if implicit && !definite {
		// An implicit global exists as nil until a store definitely ran.
		types = core.AppendIfUnique(types, c.nilType)
	}
	if len(types) == 0 {
		if implicit {
			return c.nilType, true
		}
		return nil, false
	}
	return c.getWidenedType(c.getUnionType(types)), true
}

// luaStoresBeforeSelfRead filters symbol's stores down to those that can have
// executed before `source` runs, and reports whether one of them definitely
// has. A store inside a function body is ordered by the position of the
// function (it must exist before a call can run the store), and never counts
// as definite — the call may not have happened — except within the reading
// store's own body, where one invocation runs its statements in order.
// Cross-file stores follow load order, which the checker does not track, so
// they mirror the leniency of ordinary cross-file reads: included, and treated
// as definite. A self-preserving store (`x = x`) re-stores whatever was there,
// nil included; it contributes no declared type and cannot discharge nil.
func (c *Checker) luaStoresBeforeSelfRead(source *ast.Node, assignments []luaAugmentation) (prior []luaAugmentation, definite bool) {
	sourceFile := ast.GetSourceFileOfNode(source)
	sourceFn := ast.FindAncestor(source, ast.IsFunctionLike)
	for _, assignment := range assignments {
		if assignment.Source == source {
			continue
		}
		discharges := !c.isSelfPreservingLuaCapturedTarget(assignment.Target)
		if ast.GetSourceFileOfNode(assignment.Source) != sourceFile {
			prior = append(prior, assignment)
			definite = definite || discharges
			continue
		}
		fn := ast.FindAncestor(assignment.Source, ast.IsFunctionLike)
		switch {
		case sourceFn != nil && fn == sourceFn:
			// Same body: the statement-order rules apply within one invocation.
			if assignment.Source.Pos() < source.Pos() {
				prior = append(prior, assignment)
				if discharges && luaStoreDominatesRead(assignment.Source, source) {
					definite = true
				}
			}
		case sourceFn != nil:
			// A store elsewhere is not ordered against this invocation; mirror
			// the leniency ordinary closure reads get.
			prior = append(prior, assignment)
			definite = definite || discharges
		case fn != nil:
			if fn.Pos() < source.Pos() {
				prior = append(prior, assignment)
			}
		case assignment.Source.Pos() < source.Pos():
			prior = append(prior, assignment)
			if discharges && luaStoreDominatesRead(assignment.Source, source) {
				definite = true
			}
		}
	}
	return prior, definite
}

// luaStoreDominatesRead reports whether a lexically earlier store has
// definitely executed by the time the reading store runs: the statement block
// holding the store also holds the read, so control that reached the read
// entered that block and passed the store. Two sibling stores inside one
// conditional therefore order against each other, while a store whose block
// the read is outside of may have been skipped. (goto can in principle jump
// over the store within the block; the lexical approximation ignores it.)
func luaStoreDominatesRead(store *ast.Node, read *ast.Node) bool {
	if store.Parent == nil || !ast.IsExpressionStatement(store.Parent) {
		return false
	}
	container := store.Parent.Parent
	if container == nil {
		return false
	}
	return ast.FindAncestor(read, func(parent *ast.Node) bool { return parent == container }) != nil
}

// isLuaSelfStoreRead reports whether reference reads symbol from inside one of
// symbol's own store value lists. Such a read participates in the pre-store
// snapshot, so flow analysis must not flag it as used before being assigned —
// the idiom `a = (a or 0) + 1` handles the missing value itself, and a store
// that does not is diagnosed against the snapshot instead.
func (c *Checker) isLuaSelfStoreRead(reference *ast.Node, symbol *ast.Symbol) bool {
	return c.luaSelfStoreForRead(reference, c.getMergedSymbol(symbol)) != nil
}

func (c *Checker) luaSelfStoreForRead(reference *ast.Node, merged *ast.Symbol) *ast.Node {
	assignments := c.luaAssignmentAugmentations[merged]
	if len(assignments) == 0 {
		return nil
	}
	source := luaEnclosingSelfStore(reference, assignments)
	if source == nil {
		return nil
	}
	// A value declaration outside the stores (say an annotated local) types the
	// symbol without inference over its stores; ordinary flow analysis then
	// tracks its statement order, including used-before-assigned.
	if hasLuaValueDeclarationOutsideAssignments(merged, assignments) {
		return nil
	}
	return source
}

// luaEnclosingSelfStore finds the assignment whose value list lexically holds
// reference, provided that assignment is one of the symbol's registered
// stores. A function boundary on the way up declines — a closure in the value
// runs after the store, so its reads are not part of the captured snapshot —
// unless the function is immediately invoked: an IIFE runs while the value
// list is being evaluated, before anything is stored.
func luaEnclosingSelfStore(reference *ast.Node, assignments []luaAugmentation) *ast.Node {
	for node, parent := reference, reference.Parent; parent != nil; node, parent = parent, parent.Parent {
		if ast.IsFunctionLike(parent) && ast.GetImmediatelyInvokedFunctionExpression(parent) == nil {
			return nil
		}
		if ast.IsBinaryExpression(parent) && parent.AsBinaryExpression().Right == node {
			for _, assignment := range assignments {
				if assignment.Source == parent {
					return parent
				}
			}
		}
	}
	return nil
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

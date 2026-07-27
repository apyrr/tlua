package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
)

// The typed metatable protocol. A Lua table's failed reads fall through to its metatable's
// __index, so setmetatable(t, mt) does not return the table it was given: it returns a table
// that also carries __index's shape. The result is a MetatableType, the pair of the two
// operands, and getmetatable reads the metatable back out of it.
//
// The pairing is only as precise as the metatable's type. A metatable whose type says no
// more than `LuaMetatable<T> | nil` -- which is exactly what the parameter declares -- cannot
// say which fallback applies, so such a call returns the plain table rather than a wrong
// augmentation. Precision comes from the metatable literal a call passes, or from a variable
// that has one as its type.

// MetatableTypeKey interns a setmetatable result. The pair of operands determines the result:
// everything else the type stores is derived from the metatable.
type MetatableTypeKey struct {
	tableID     TypeId
	metatableID TypeId
}

type luaMetatableCallKind int

const (
	luaMetatableCallNone luaMetatableCallKind = iota
	luaMetatableCallSet
	luaMetatableCallGet
	luaMetatableCallDebugSet
	luaMetatableCallDebugGet
)

func (k luaMetatableCallKind) isSet() bool {
	return k == luaMetatableCallSet || k == luaMetatableCallDebugSet
}

func (k luaMetatableCallKind) isGet() bool {
	return k == luaMetatableCallGet || k == luaMetatableCallDebugGet
}

// getLuaMetatableCall reports which metatable global a call invokes. Only the real globals carry
// the protocol: a shadowing local, or a call through an alias, resolves to a different symbol and
// keeps whatever it declares, as it does for the type() narrowing builtins. The debug library's
// setmetatable and getmetatable are recognized as member calls, the way io.type is: the callee's
// base must be the debug global.
func (c *Checker) getLuaMetatableCall(node *ast.Node) luaMetatableCallKind {
	if !ast.IsCallExpression(node) {
		return luaMetatableCallNone
	}
	callee := ast.SkipParentheses(node.Expression())
	if ast.IsIdentifier(callee) {
		switch callee.Text() {
		case "setmetatable":
			if c.isLuaGlobalReference(callee, c.getLuaSetmetatableGlobalSymbol()) {
				return luaMetatableCallSet
			}
		case "getmetatable":
			if c.isLuaGlobalReference(callee, c.getLuaGetmetatableGlobalSymbol()) {
				return luaMetatableCallGet
			}
		}
		return luaMetatableCallNone
	}
	if ast.IsPropertyAccessExpression(callee) && c.isLuaGlobalReference(ast.SkipParentheses(callee.Expression()), c.getLuaDebugGlobalSymbol()) {
		switch callee.Name().Text() {
		case "setmetatable":
			return luaMetatableCallDebugSet
		case "getmetatable":
			return luaMetatableCallDebugGet
		}
	}
	return luaMetatableCallNone
}

// isLuaMetatableCall reports whether the metatable protocol interprets this call. The quick
// path of getQuickTypeOfExpression must skip such a call: it reads the signature's return type
// without running checkCallExpression's special cases.
func (c *Checker) isLuaMetatableCall(node *ast.Node) bool {
	return c.getLuaMetatableCall(node) != luaMetatableCallNone
}

// checkLuaMetatableCall returns the type of a setmetatable or getmetatable call, or nil when
// the call's declared return type stands.
func (c *Checker) checkLuaMetatableCall(node *ast.Node, returnType *Type) *Type {
	args := node.Arguments()
	kind := c.getLuaMetatableCall(node)
	switch {
	case kind.isSet():
		if len(args) < 2 {
			return nil
		}
		// A protected table refuses a new metatable: the runtime raises, so the old pairing
		// stands unchanged -- the one place setmetatable does not replace. debug.setmetatable
		// bypasses the protection.
		if kind == luaMetatableCallSet && c.luaSetmetatableSeesProtected(node, returnType) {
			c.error(node, diagnostics.Cannot_change_a_protected_metatable)
			return returnType
		}
		// The table is the signature's return type: setmetatable is declared to return its
		// first parameter, so inference has already folded and widened a literal argument. The
		// metatable is read from the argument itself -- its declared parameter type is the
		// imprecise union by construction -- and widened, so a fresh literal never reaches a
		// type that outlives the expression.
		return c.getSetmetatableResultType(returnType, c.getLuaMetatableArgumentType(args[1]))
	case kind.isGet():
		if len(args) < 1 {
			return nil
		}
		return c.getGetmetatableResultType(c.checkExpressionCached(args[0]), kind == luaMetatableCallDebugGet)
	}
	return nil
}

// getLuaMetatableArgumentType is the metatable operand as the pairing sees it: read from the
// argument -- the declared parameter type is the imprecise union by construction -- and widened,
// so a fresh literal never reaches a type that outlives the expression.
func (c *Checker) getLuaMetatableArgumentType(arg *ast.Node) *Type {
	return c.getWidenedType(c.getRegularTypeOfObjectLiteral(c.checkExpressionCached(arg)))
}

// attachLuaMetatablePairings maps each augmentation symbol to the statement-form
// setmetatable calls whose table operand names its storage. Flow analysis already
// re-pairs references downstream of such a statement, but flow never leaves the
// statement's own file or function, while the storage's declared type is what
// every other file reads -- so the pairing must reach the declared type too.
// Statement placement follows the leniency of the write machinery: like a store,
// a call inside a function body or another file still contributes.
func (c *Checker) attachLuaMetatablePairings() {
	resolver := c.resolveLuaConstructors()
	for _, file := range c.files {
		for _, node := range file.LuaSetmetatableCandidates {
			if !c.getLuaMetatableCall(node).isSet() {
				continue
			}
			symbols, known := resolver.referenceSymbols(skipLuaRuntimeTransparentWrappers(node.Arguments()[0]))
			if !known {
				continue
			}
			isStatement := ast.IsLuaStatementFormCall(node)
			for _, symbol := range symbols {
				symbol = c.getMergedSymbol(symbol)
				// Only inferred augmentation storage folds the pairing in; a declared
				// contract keeps its annotation, exactly as it does for stores. A
				// dotted method's symbol resolves through the function path, which
				// never applies pairings -- and a metatable on a function is per-type
				// at runtime, not per-value -- so function storage stays flow-only.
				if symbol == nil || symbol.Flags&ast.SymbolFlagsAssignment == 0 ||
					symbol.Flags&(ast.SymbolFlagsFunction|ast.SymbolFlagsMethod) != 0 {
					continue
				}
				// Every real call anchors the protected check at its own program
				// point; only the statement form -- the fire-and-forget install
				// idiom -- contributes to the declared type. An expression form
				// feeds its result somewhere, and that value's flow or the store
				// machinery carries it from there.
				c.luaMetatablePairingCalls[node] = core.AppendIfUnique(c.luaMetatablePairingCalls[node], symbol)
				if !isStatement {
					continue
				}
				if len(c.luaMetatablePairings[symbol]) == 0 {
					// First pairing wins the symbol's place in the resolution
					// order ensureLuaPairedSymbolsResolved replays.
					c.luaOrderedPairingSymbols = append(c.luaOrderedPairingSymbols, symbol)
				}
				c.luaMetatablePairings[symbol] = append(c.luaMetatablePairings[symbol], node)
			}
		}
	}
}

// ensureLuaPairedSymbolsResolved resolves every paired symbol's type in program
// order the first time any of them is about to resolve. Metatable arguments may
// reference each other in a cycle; whichever symbol resolves first captures its
// neighbors fully paired, while the read that closes the cycle captures the
// first symbol's temporarily published unpaired base. That break is inherent --
// no finite type holds a mutual pairing -- but where it lands must not depend on
// which symbol a checker, a hover, or a pool ordering happens to touch first.
// Called before the triggering symbol publishes anything, so the replay below
// resolves it in its canonical position like any other. Runs on first use rather
// than at attach time: attachLuaMetatablePairings runs before the checker's
// global types exist, too early to resolve anything.
func (c *Checker) ensureLuaPairedSymbolsResolved(symbol *ast.Symbol) {
	// The map-emptiness test keeps the common pairing-free program at one bool
	// and one length read, skipping the symbol lookups entirely.
	if c.luaPairedSymbolsResolved || len(c.luaMetatablePairings) == 0 ||
		len(c.luaMetatablePairings[c.getMergedSymbol(symbol)]) == 0 {
		return
	}
	// Set first: every resolution below re-enters the hook that called us.
	c.luaPairedSymbolsResolved = true
	for _, paired := range c.luaOrderedPairingSymbols {
		c.getTypeOfSymbol(paired)
	}
}

// applyLuaDeclaredMetatablePairings pairs an augmentation symbol's declared type
// with the metatables the program's setmetatable statements install on it, in
// program order -- pairing replaces, so with several calls the last one stands,
// as at runtime. Called after the unpaired type is published to the symbol's
// links: computing the metatable argument may read the symbol back (the class
// idiom's self-referential metatables), and that read must see the plain table
// rather than a false circularity. The pre-pairing base is recorded so the
// protected check can rebuild the type a given call saw (luaSetmetatableSeesProtected).
func (c *Checker) applyLuaDeclaredMetatablePairings(symbol *ast.Symbol, t *Type) *Type {
	if len(c.luaMetatablePairings) == 0 {
		return t
	}
	merged := c.getMergedSymbol(symbol)
	if len(c.luaMetatablePairings[merged]) == 0 {
		return t
	}
	c.luaDeclaredPairingBases[merged] = t
	return c.replayLuaDeclaredPairings(merged, t, nil /*stopAt*/, nil /*anchor*/)
}

// replayLuaDeclaredPairings walks a symbol's effect timeline in program order,
// stopping before stopAt when given, and folds the pairings onto the base type.
// A rebinding store puts a fresh table in the storage, so the accumulated
// pairing -- protection included -- does not survive it; but a store only
// matters when something after it observes the storage: the next pairing, or
// the anchored call itself. The declared surface (anchor nil) therefore follows
// the latest pairing even past a trailing store, while a protected-check prefix
// (anchor = the call being checked) sees a rebind just before its own position.
// An anchored replay only honors a reset that definitely ran before the anchor:
// a conditional rebind may not have executed, and on the path where it did not,
// the protection still refuses the call.
func (c *Checker) replayLuaDeclaredPairings(merged *ast.Symbol, base *Type, stopAt *ast.Node, anchor *ast.Node) *Type {
	var stop *luaProgramOrderKey
	if stopAt != nil {
		key := c.luaProgramOrder(stopAt)
		stop = &key
	}
	t := base
	resetPending := false
	for _, effect := range c.getLuaSymbolEffectTimeline(merged) {
		// The stop point is compared by program order, not identity: an
		// expression-form anchor is not on the timeline, but still stops the
		// replay at its own position.
		if stop != nil && !effect.key.less(*stop) {
			break
		}
		// An effect in the opposite branch of a run-once `if` cannot have
		// happened on any path that reaches the anchor.
		if anchor != nil && luaEffectExcludesAnchor(effect.node, anchor) {
			continue
		}
		switch effect.kind {
		case luaEffectRebindingStore:
			if anchor == nil || c.luaEffectDefinitelyPrecedes(effect.node, anchor) {
				resetPending = true
			}
		case luaEffectPairing:
			if resetPending {
				t = base
				resetPending = false
			}
			// Only debug's bypass can strip protection, and for an anchored
			// prefix stripping counts only when the call definitely ran: on
			// the path where it did not, the protection remains and still
			// refuses the anchored call. Installing protection stays ungated
			// -- any path that installed it refuses -- and an ordinary
			// pairing over a protected type is a refused no-op either way.
			if anchor != nil && someType(t, c.isProtectedMetatableType) && !c.luaEffectDefinitelyPrecedes(effect.node, anchor) {
				continue
			}
			t = c.applyLuaDeclaredMetatablePairing(t, effect.node)
		}
	}
	if resetPending && stop != nil {
		t = base
	}
	return t
}

// luaEffectDefinitelyPrecedes reports whether an effect -- a rebinding store or a
// protection-stripping pairing call -- has definitely run by the time the anchor call
// does: the same ordering questions luaStoresBeforeSelfRead answers for a self-read,
// asked of a statement already known to precede the anchor in program order. An effect
// in a function body needs a call nothing orders, so only the anchor's own body orders
// it, by dominance within one invocation. An effect outside the anchor's body is
// reached by the leniency ordinary cross-body and cross-file reads get, provided it
// runs unconditionally where it sits; a conditional one may have been skipped, and a
// skipped rebind or bypass leaves the old table -- protection included -- in place.
// (An assertion-wrapped call statement conservatively fails both placement tests --
// its parent is the wrapper -- which keeps protection, the safe direction.)
func (c *Checker) luaEffectDefinitelyPrecedes(source *ast.Node, anchor *ast.Node) bool {
	sourceFn := ast.FindAncestor(source, ast.IsFunctionLike)
	anchorFn := ast.FindAncestor(anchor, ast.IsFunctionLike)
	switch {
	case sourceFn != nil:
		// A store in a function body needs a call nothing here orders. Only the
		// anchor's own body orders it: control there passed the store when the
		// store's block holds the anchor, or when nothing in the body skips it.
		return sourceFn == anchorFn && (luaStoreDominatesRead(source, anchor) || luaStoreRunsUnconditionally(source))
	case anchorFn == nil && ast.GetSourceFileOfNode(source) == ast.GetSourceFileOfNode(anchor):
		// Both run in the same chunk: control that reached the call either
		// entered the store's own block on the way, or could not skip the store.
		return luaStoreDominatesRead(source, anchor) || luaStoreRunsUnconditionally(source)
	default:
		// A top-level store in another file, or one the anchor's body closes over:
		// reached by the leniency ordinary cross-body and cross-file reads get,
		// provided nothing can skip it.
		return luaStoreRunsUnconditionally(source)
	}
}

// luaEffectExcludesAnchor reports whether an effect and the anchor sit in mutually
// exclusive branches of one run-once `if`: a chunk executes once, so only one branch of
// a top-level `if` ever runs, and an effect in the other branch cannot have happened on
// any path that reaches the anchor. Inside a function body or a loop the same `if` runs
// again with the condition free to flip, so the branches are not exclusive across
// invocations and no effect is excluded. A function boundary between the effect and the
// branch likewise defers the effect past the branch's own run, ending the walk.
func luaEffectExcludesAnchor(effect *ast.Node, anchor *ast.Node) bool {
	child := effect
	for parent := child.Parent; parent != nil; child, parent = parent, parent.Parent {
		if ast.IsFunctionLike(parent) {
			return false
		}
		if !ast.IsIfStatement(parent) {
			continue
		}
		ifStatement := parent.AsIfStatement()
		var other *ast.Node
		switch child {
		case ifStatement.ThenStatement:
			other = ifStatement.ElseStatement
		case ifStatement.ElseStatement:
			other = ifStatement.ThenStatement
		}
		// An effect in the condition runs on both paths; keep climbing.
		if other != nil && ast.FindAncestor(anchor, func(node *ast.Node) bool { return node == other }) != nil {
			return luaIfRunsAtMostOnce(parent)
		}
	}
	return false
}

// luaIfRunsAtMostOnce reports whether an `if` statement executes at most once per
// program run: nothing above it is a function body or a loop.
func luaIfRunsAtMostOnce(node *ast.Node) bool {
	return ast.FindAncestor(node.Parent, func(parent *ast.Node) bool {
		return ast.IsFunctionLike(parent) || ast.IsIterationStatement(parent)
	}) == nil
}

// luaStoreRunsUnconditionally reports whether a store runs whenever the body holding it
// runs: its statement sits in that chunk or function body rather than inside a
// conditional or a loop. A bare block (a do-block) runs exactly once when its holder
// does, so the climb sees through any nesting of them; a block belonging to an if or a
// loop hangs off the control statement, not another block, and stops the climb.
func luaStoreRunsUnconditionally(source *ast.Node) bool {
	if source.Parent == nil || !ast.IsExpressionStatement(source.Parent) {
		return false
	}
	container := source.Parent.Parent
	for container != nil && ast.IsBlock(container) && container.Parent != nil && ast.IsBlock(container.Parent) {
		container = container.Parent
	}
	if container == nil {
		return false
	}
	return ast.IsSourceFile(container) ||
		ast.IsBlock(container) && container.Parent != nil &&
			(ast.IsFunctionLike(container.Parent) || ast.IsSourceFile(container.Parent))
}

// applyLuaDeclaredMetatablePairing applies one recorded call to a declared type.
// A protected pairing refuses replacement (the call site reports it); debug.setmetatable
// bypasses the protection, as on the flow path. Union storage pairs arm by arm: the
// call pairs whichever table the storage holds, and a pairing itself never forms over
// a union (canCarryMetatableMembers).
func (c *Checker) applyLuaDeclaredMetatablePairing(t *Type, node *ast.Node) *Type {
	if c.getLuaMetatableCall(node) == luaMetatableCallSet && someType(t, c.isProtectedMetatableType) {
		return t
	}
	metatableType := c.getLuaDeclaredPairingMetatableType(node, t)
	return c.mapType(t, func(arm *Type) *Type {
		if paired := c.getSetmetatableResultType(arm, metatableType); paired != nil {
			return paired
		}
		return arm
	})
}

// getLuaDeclaredPairingMetatableType checks a recorded call's metatable operand for the
// declared-type replay. The replay can be what checks the operand first -- the call's own
// table argument may have triggered the symbol resolution that runs it -- and at that
// moment the call's signature is still resolving, so contextual typing through the call
// finds nothing and a context-sensitive metatable member would lose its parameter types.
// The context the call would provide is supplied directly: setmetatable declares the
// operand as LuaMetatable<T> with T the table being paired. Memoized per call so prefix
// rebuilds reuse the same interned type.
func (c *Checker) getLuaDeclaredPairingMetatableType(node *ast.Node, tableType *Type) *Type {
	if cached, ok := c.luaPairingMetatableArgTypes[node]; ok {
		return cached
	}
	arg := node.Arguments()[1]
	// A statement that names several storage symbols (disjoint constructor-arm
	// groups sharing a member name) replays under several table types, but the
	// literal's members freeze on their first check. Contextualize such a call
	// against the bare table type instead, so the memo's content cannot depend
	// on which storage resolves first.
	operandType := c.getUnpairedTableType(tableType)
	if len(c.luaMetatablePairingCalls[node]) > 1 {
		operandType = c.nonPrimitiveType
	}
	if typeArguments := node.TypeArguments(); len(typeArguments) != 0 {
		// An explicit instantiation names T itself, so the call contextualizes
		// the literal with the written type; the replay must match it or the
		// literal freezes under a contract the call then contradicts.
		operandType = c.getTypeFromTypeNode(typeArguments[0])
	}
	var t *Type
	if globalType := c.getGlobalType("LuaMetatable", 1 /*arity*/, false /*reportErrors*/); globalType != c.emptyGenericType {
		contextualType := c.createTypeFromGenericGlobalType(globalType, []*Type{operandType})
		t = c.checkExpressionWithContextualType(arg, contextualType, nil /*inferenceContext*/, CheckModeNormal)
		t = c.getWidenedType(c.getRegularTypeOfObjectLiteral(t))
	} else {
		t = c.getLuaMetatableArgumentType(arg)
	}
	c.luaPairingMetatableArgTypes[node] = t
	return t
}

// luaSetmetatableSeesProtected reports whether a setmetatable call's table operand is
// protected when the call runs. A recorded pairing call answers from its prefix alone:
// the base with the pairings and rebinding stores that precede it replayed. The
// flow-derived operand type is wrong in both directions there -- outside the call's own
// file it falls back to the final declared type, which contains the pairing this very
// call installs (a false report on the installer) and reflects replacements that run
// only after this call, debug.setmetatable's bypass included (a masked report on a
// refused replacement). An unrecorded call keeps the flow answer, which is exact for
// the local narrowing that produced it.
func (c *Checker) luaSetmetatableSeesProtected(node *ast.Node, returnType *Type) bool {
	symbols := c.luaMetatablePairingCalls[node]
	if len(symbols) == 0 {
		return someType(returnType, c.isProtectedMetatableType)
	}
	return core.Some(symbols, func(symbol *ast.Symbol) bool {
		return someType(c.getLuaDeclaredPairingTypeBefore(symbol, node), c.isProtectedMetatableType)
	})
}

// getLuaDeclaredPairingTypeBefore rebuilds a symbol's declared type as of just before
// `call`: the recorded pre-pairing base with only the pairings preceding the call
// applied. Pairing results are interned, so the rebuild redoes no structural work.
func (c *Checker) getLuaDeclaredPairingTypeBefore(symbol *ast.Symbol, call *ast.Node) *Type {
	merged := c.getMergedSymbol(symbol)
	base, ok := c.luaDeclaredPairingBases[merged]
	if !ok {
		// Resolving the symbol records the base on the way to the paired type.
		t := c.getTypeOfSymbol(merged)
		if base, ok = c.luaDeclaredPairingBases[merged]; !ok {
			// A resolution that never applied pairings (a collision's pre-set
			// any, say) has nothing call-order-sensitive in it.
			return t
		}
	}
	return c.replayLuaDeclaredPairings(merged, base, call /*stopAt*/, call /*anchor*/)
}

type luaProgramOrderKey struct {
	file int
	pos  int
}

// luaSymbolEffectKind classifies one entry on a symbol's effect timeline.
type luaSymbolEffectKind int8

const (
	// luaEffectRebindingStore is an effective store that replaces the table
	// outright, clearing whatever pairing the storage had accumulated.
	luaEffectRebindingStore luaSymbolEffectKind = iota
	// luaEffectPairing is a statement-form setmetatable call that installs a
	// metatable on the storage.
	luaEffectPairing
)

// luaSymbolEffect is one program-ordered effect on an augmentation symbol's
// declared metatable state. The node lets an anchored replay ask whether a
// store definitely ran before the call it anchors on.
type luaSymbolEffect struct {
	key  luaProgramOrderKey
	kind luaSymbolEffectKind
	node *ast.Node
}

func (key luaProgramOrderKey) less(other luaProgramOrderKey) bool {
	return key.file < other.file || key.file == other.file && key.pos < other.pos
}

func (c *Checker) luaProgramOrder(node *ast.Node) luaProgramOrderKey {
	return luaProgramOrderKey{file: c.fileIndexMap[ast.GetSourceFileOfNode(node)], pos: node.Pos()}
}

// getLuaSymbolEffectTimeline merges a symbol's statement-form pairings with its
// rebinding stores into one program-ordered list. A defaulted guard (`X = X or
// {}`) keeps the existing table -- and its metatable -- whenever one exists,
// and a self-preserving store re-stores what was there, so neither appears. A
// store inside a function body does, because a pairing inside one contributes:
// the two must agree, or a body that rebinds and re-pairs would replay the old
// protection against its own fresh table. Whether a store definitely ran is the
// replay's question (luaEffectDefinitelyPrecedes), not the timeline's.
func (c *Checker) getLuaSymbolEffectTimeline(merged *ast.Symbol) []luaSymbolEffect {
	if effects, ok := c.luaSymbolEffectTimelines[merged]; ok {
		return effects
	}
	effects := []luaSymbolEffect{}
	for _, assignment := range c.effectiveLuaAssignmentAugmentations(c.luaAssignmentAugmentations[merged]) {
		if c.isSelfPreservingLuaCapturedTarget(assignment.Target) {
			continue
		}
		if initializer := luaExplicitAssignmentValueAt(assignment.Source.AsBinaryExpression().Right, assignment.ValueIndex); initializer != nil &&
			(c.isLuaDefaultedAugmentationGuard(assignment.Target, initializer) ||
				c.luaStoreMayPreserveTarget(assignment.Target, initializer)) {
			// A value that may evaluate to the target itself -- `X = cond and X
			// or {}` -- keeps the table and its metatable on that path, so the
			// store is no rebinding.
			continue
		}
		effects = append(effects, luaSymbolEffect{key: c.luaProgramOrder(assignment.Source), kind: luaEffectRebindingStore, node: assignment.Source})
	}
	for _, node := range c.luaMetatablePairings[merged] {
		effects = append(effects, luaSymbolEffect{key: c.luaProgramOrder(node), kind: luaEffectPairing, node: node})
	}
	// Stores and pairings are whole statements, so no two effects share a key.
	slices.SortFunc(effects, func(left luaSymbolEffect, right luaSymbolEffect) int {
		if left.key.less(right.key) {
			return -1
		}
		if right.key.less(left.key) {
			return 1
		}
		return 0
	})
	c.luaSymbolEffectTimelines[merged] = effects
	return effects
}

// getGetmetatableResultType returns the metatable paired with t, or nil when t carries no
// pairing the checker knows about and the declared return type stands. A protected pairing shows
// its __metatable sentinel instead of the real metatable, unless the caller bypasses it as
// debug.getmetatable does.
func (c *Checker) getGetmetatableResultType(t *Type, bypassProtection bool) *Type {
	if t.flags&TypeFlagsUnion != 0 {
		metatables := make([]*Type, 0, len(t.Types()))
		for _, u := range t.Types() {
			metatable := c.getGetmetatableResultType(u, bypassProtection)
			if metatable == nil {
				// One arm with no known metatable makes the whole union unknown: the value may
				// be that arm, and no answer covers both.
				return nil
			}
			metatables = append(metatables, metatable)
		}
		return c.getUnionType(metatables)
	}
	if isMetatableType(t) {
		d := t.AsMetatableType()
		if !bypassProtection && d.protectedType != nil {
			// __metatable is what a protected table shows in place of its metatable.
			return d.protectedType
		}
		return d.metatableType
	}
	return nil
}

func isMetatableType(t *Type) bool {
	return t.flags&TypeFlagsObject != 0 && t.objectFlags&ObjectFlagsMetatable != 0
}

func (c *Checker) isProtectedMetatableType(t *Type) bool {
	return isMetatableType(t) && t.AsMetatableType().protectedType != nil
}

// getUnpairedTableType strips a pairing from a table: what setmetatable returned before is the
// table it was given, and the metatable it was given is gone.
func (c *Checker) getUnpairedTableType(t *Type) *Type {
	return c.mapType(t, func(u *Type) *Type {
		for isMetatableType(u) {
			u = u.AsMetatableType().tableType
		}
		return u
	})
}

// luaOperatorMetamethods are the operator metamethods the checker dispatches. __tostring, __gc
// and __mode never change a value's type, so none of the three is worth a pairing on its own.
var luaOperatorMetamethods = []string{
	"__eq", "__lt", "__le", "__unm", "__len",
	"__add", "__sub", "__mul", "__div", "__mod", "__pow", "__concat",
}

// getMetatableHandlerType returns the committed, nil-stripped type of a metamethod, or nil. A
// declared-optional member commits to nothing: LuaMetatable<T> declares every metamethod
// optionally, so reading through its optionality would make every LuaMetatable<T>-typed argument
// callable and protected. A metatable literal's members are not optional, so the idiomatic path
// is unaffected. __index is the exception -- its optional form types the defaults idiom -- and
// reads through optionality on its own path.
func (c *Checker) getMetatableHandlerType(metatableType *Type, name string) *Type {
	// Metamethods never live on the global Object/Function augmentations, so the lookup skips
	// them -- and a merged `interface Object { __add: ... }` cannot inject an operator everywhere.
	prop := c.getPropertyOfTypeEx(metatableType, name, false /*includeTypeOnlyMembers*/)
	if prop == nil || prop.Flags&ast.SymbolFlagsOptional != 0 {
		return nil
	}
	t := c.removeMissingOrUndefinedType(c.getTypeOfSymbol(prop))
	if c.isErrorType(t) {
		return nil
	}
	return t
}

// getMetatableSource returns a metamethod that is a table or a function -- __index and
// __newindex both take that shape -- along with which one it is, or nil when the metamethod is
// absent, ambiguous, or not table-or-function shaped. A callable table is dropped as ambiguous:
// which half of it answers a lookup is not worth guessing at. __index reads through optionality
// (the defaults idiom); every other metamethod does not.
func (c *Checker) getMetatableSource(metatableType *Type, name string, allowOptional bool) (*Type, bool) {
	prop := c.getPropertyOfTypeEx(metatableType, name, false /*includeTypeOnlyMembers*/)
	if prop == nil || !allowOptional && prop.Flags&ast.SymbolFlagsOptional != 0 {
		return nil, false
	}
	source := c.removeMissingOrUndefinedType(c.getTypeOfSymbol(prop))
	if c.isErrorType(source) || source.flags&(TypeFlagsObject|TypeFlagsIntersection|TypeFlagsInstantiableNonPrimitive) == 0 {
		return nil, false
	}
	isFunction := len(c.getSignaturesOfType(source, SignatureKindCall)) != 0
	if isFunction && (len(c.getPropertiesOfType(source)) != 0 || len(c.getIndexInfosOfType(source)) != 0) {
		return nil, false
	}
	return source, isFunction
}

// getMetatableCallSource returns __call when it commits to a callable, or nil. Its signatures
// make the pairing callable.
func (c *Checker) getMetatableCallSource(metatableType *Type) *Type {
	handler := c.getMetatableHandlerType(metatableType, "__call")
	if handler == nil || len(c.getSignaturesOfType(handler, SignatureKindCall)) == 0 {
		return nil
	}
	return handler
}

// metatableHasOperatorMetamethod reports whether the metatable commits to any operator
// metamethod: a committed, callable handler is what an operator can dispatch to.
func (c *Checker) metatableHasOperatorMetamethod(metatableType *Type) bool {
	return core.Some(luaOperatorMetamethods, func(name string) bool {
		handler := c.getMetatableHandlerType(metatableType, name)
		return handler != nil && len(c.getSignaturesOfType(handler, SignatureKindCall)) != 0
	})
}

// getSetmetatableResultType pairs a table with the metatable it is given. setmetatable
// *replaces* a metatable, so the table is unpaired first: whatever its old metamethods answered,
// they do not answer now. The new pairing forms only when the metatable commits to a metamethod
// the checker interprets; otherwise the unpaired table stands, which is also how
// `setmetatable(t, nil)` detaches one. Returns nil only when the call is already in error and its
// declared type should stand.
func (c *Checker) getSetmetatableResultType(tableType *Type, metatableType *Type) *Type {
	if c.isErrorType(tableType) {
		return nil
	}
	tableType = c.getUnpairedTableType(tableType)
	if c.isErrorType(metatableType) || !c.canCarryMetatableMembers(metatableType) {
		return tableType
	}
	// __index is declared optional, so a nil arm says only that there may be no fallback at all.
	// What remains has to name a single fallback for the read augmentation to mean anything.
	indexSource, indexIsFunction := c.getMetatableSource(metatableType, "__index", true /*allowOptional*/)
	if c.isGenericObjectType(tableType) || indexSource != nil && c.isGenericObjectType(indexSource) {
		// A pairing resolves its members eagerly, so it can only hold concrete types. A generic
		// one falls back to the intersection of table and __index, as a generic spread does:
		// less precise about collisions, blind to every other metamethod, but it instantiates.
		// The imprecision cuts the other way too: as a plain intersection it exposes the __index
		// source's metamethod-named members to ambient operator dispatch, which a real pairing
		// refuses (Lua does not inherit metamethods through __index). Accepted with the rest.
		if indexSource == nil || indexIsFunction {
			return tableType
		}
		return c.getIntersectionType([]*Type{tableType, indexSource})
	}
	if !c.canCarryMetatableMembers(tableType) {
		return tableType
	}
	newindexSource, newindexIsFunction := c.getMetatableSource(metatableType, "__newindex", false /*allowOptional*/)
	callSource := c.getMetatableCallSource(metatableType)
	protectedType := c.getMetatableHandlerType(metatableType, "__metatable")
	if indexSource == nil && newindexSource == nil && callSource == nil && protectedType == nil &&
		!c.metatableHasOperatorMetamethod(metatableType) {
		// The metatable commits to nothing the checker interprets, so the plain table stands.
		return tableType
	}
	key := MetatableTypeKey{tableID: tableType.id, metatableID: metatableType.id}
	if cached := c.metatableTypes[key]; cached != nil {
		return cached
	}
	t := c.newObjectType(ObjectFlagsMetatable, nil /*symbol*/)
	// The merged members come from both operands, so a non-inferrable one makes the pair
	// non-inferrable too.
	t.objectFlags |= (tableType.objectFlags | metatableType.objectFlags) & ObjectFlagsNonInferrableType
	d := t.AsMetatableType()
	d.tableType = tableType
	d.metatableType = metatableType
	d.indexSource = indexSource
	d.indexIsFunction = indexIsFunction
	d.newindexSource = newindexSource
	d.newindexIsFunction = newindexIsFunction
	d.callSource = callSource
	d.protectedType = protectedType
	c.metatableTypes[key] = t
	return t
}

// canCarryMetatableMembers reports whether a pairing can read members out of t. An intersection
// can: it resolves properties and index signatures like the object types it is made of. A union
// cannot -- which of its arms answers a read is exactly what it does not say.
func (c *Checker) canCarryMetatableMembers(t *Type) bool {
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0
}

// canCarryAmbientMetamethods reports whether an unpaired operand type can carry ambient
// metamethod members: the same structured shapes getMetatableSource reads from. Primitives never
// carry one -- the number/string operator fast paths stay free of property walks -- and a union
// answers nothing here as everywhere (canCarryMetatableMembers): which arm would dispatch is
// exactly what it does not say.
func canCarryAmbientMetamethods(t *Type) bool {
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection|TypeFlagsInstantiableNonPrimitive) != 0
}

// instantiateMetatableType instantiates a pairing's operands and re-pairs them. The pairing is
// not structural -- its members are already merged -- so it cannot be instantiated in place; and
// if the instantiated metatable no longer says what the fallback is, the augmentation goes with
// it and the table stands alone.
func (c *Checker) instantiateMetatableType(t *Type, m *TypeMapper) *Type {
	d := t.AsMetatableType()
	tableType := c.instantiateType(d.tableType, m)
	metatableType := c.instantiateType(d.metatableType, m)
	if tableType == d.tableType && metatableType == d.metatableType {
		return t
	}
	if result := c.getSetmetatableResultType(tableType, metatableType); result != nil {
		return result
	}
	return tableType
}

// resolveMetatableTypeMembers merges a table with the metamethods of its metatable. A read the
// table misses falls through to __index; the table wins a collision, because a read hits the
// table first -- but only when the table answers with a value: a key that is absent or nil is
// exactly what makes Lua consult the metatable, so a member that admits nil reads as its own
// non-nil half or the fallback. __call makes the pairing callable. __index never contributes call
// signatures: a callable __index table does not make the table itself callable, which is __call's
// job.
func (c *Checker) resolveMetatableTypeMembers(t *Type) {
	d := t.AsMetatableType()
	// The table's own shape. Its property symbols are shared rather than cloned, exactly as an
	// interface shares the symbols of the members it inherits.
	members := createSymbolTable(c.getPropertiesOfType(d.tableType))
	if members == nil {
		members = make(ast.SymbolTable)
	}
	callSignatures := c.getSignaturesOfType(d.tableType, SignatureKindCall)
	constructSignatures := c.getSignaturesOfType(d.tableType, SignatureKindConstruct)
	indexInfos := slices.Clone(c.getIndexInfosOfType(d.tableType))
	// Publish the table's shape before reading the metatable, so a lookup that re-enters this
	// type while a fallback resolves sees a resolved, if incomplete, type instead of recurring.
	c.setStructuredTypeMembers(t, members, callSignatures, constructSignatures, indexInfos)
	t.objectFlags |= ObjectFlagsUnresolvedMembers
	if d.indexSource != nil {
		fallbackInfos := c.getMetatableFallbackIndexInfos(d)
		for name, prop := range members {
			if merged := c.getMetatableFallthroughSymbol(d, fallbackInfos, prop); merged != nil {
				members[name] = merged
			}
		}
		if !d.indexIsFunction {
			// The fallback's own members fill the names the table does not have at all.
			members = c.addInheritedMembers(members, c.getPropertiesOfType(d.indexSource))
		}
		for _, info := range fallbackInfos {
			if findIndexInfo(indexInfos, info.keyType) == nil {
				indexInfos = append(indexInfos, info)
			}
		}
	}
	if d.callSource != nil {
		// __call makes the pairing callable: a call passes the table itself as the metamethod's
		// first argument, so a call site's arguments begin at its second parameter.
		callSignatures = slices.Clone(callSignatures)
		for _, signature := range c.getSignaturesOfType(d.callSource, SignatureKindCall) {
			callSignatures = append(callSignatures, c.getLuaCallMetamethodSignature(signature))
		}
	}
	t.objectFlags &^= ObjectFlagsUnresolvedMembers
	c.setStructuredTypeMembers(t, members, callSignatures, constructSignatures, indexInfos)
}

// getLuaCallMetamethodSignature drops __call's receiver parameter: Lua passes the table itself as
// the metamethod's first argument, so a call site supplies the rest. A signature whose position 0
// is already its rest parameter keeps it -- the rest absorbed the receiver and absorbs the
// arguments the same way.
func (c *Checker) getLuaCallMetamethodSignature(sig *Signature) *Signature {
	if len(sig.parameters) == 0 || signatureHasRestParameter(sig) && len(sig.parameters) == 1 {
		return sig
	}
	result := c.newSignature(sig.flags&SignatureFlagsPropagatingFlags, sig.declaration, sig.typeParameters,
		sig.thisParameter, sig.parameters[1:], c.getReturnTypeOfSignature(sig), nil,
		max(int(sig.minArgumentCount)-1, 0))
	result.target = sig.target
	result.mapper = sig.mapper
	return result
}

// getMetatableFallbackIndexInfos returns the index signatures __index contributes. A table
// contributes the ones it declares. A function contributes one per key domain its parameter
// admits: it answers a read of *any* key there, which is an index signature and not a property,
// since nothing names the keys it will be asked for. The domains are the ones a computed table
// key synthesizes -- `any` keys both strings and numbers, and a literal key widens, because an
// index may not be keyed by one.
func (c *Checker) getMetatableFallbackIndexInfos(d *MetatableType) []*IndexInfo {
	if !d.indexIsFunction {
		return c.getIndexInfosOfType(d.indexSource)
	}
	var infos []*IndexInfo
	for _, signature := range c.getSignaturesOfType(d.indexSource, SignatureKindCall) {
		keyType := c.getTypeAtPosition(signature, 1)
		if !c.isValidIndexArgumentType(keyType) {
			continue
		}
		valueType := c.adjustMultiReturn(c.getReturnTypeOfSignature(signature))
		c.forEachObjectLiteralIndexKeyType(keyType, func(indexKeyType *Type) {
			// Overloads that answer the same key domain both apply: a read of such a key gets
			// whichever one Lua dispatches to. The infos are ours until they are published, so
			// widening one in place is safe.
			if existing := findIndexInfo(infos, indexKeyType); existing != nil {
				existing.valueType = c.getUnionType([]*Type{existing.valueType, valueType})
				return
			}
			infos = append(infos, c.newIndexInfo(indexKeyType, valueType, false /*isReadonly*/, nil /*declaration*/, nil /*components*/))
		})
	}
	return infos
}

// getMetatableFallthroughSymbol merges a table member that may be nil with what __index answers
// for the same key, or returns nil when the member stands as it is. A raw read that comes back
// nil is what makes Lua run the metatable, so such a member is not the whole answer: it reads as
// its own non-nil half or the fallback's value. A member that cannot be nil never falls through.
func (c *Checker) getMetatableFallthroughSymbol(d *MetatableType, fallbackInfos []*IndexInfo, prop *ast.Symbol) *ast.Symbol {
	propType := c.getTypeOfSymbol(prop)
	nonNilType := c.GetNonNullableType(propType)
	if nonNilType == propType && prop.Flags&ast.SymbolFlagsOptional == 0 {
		return nil
	}
	fallbackType := c.getMetatableFallbackType(d, fallbackInfos, prop.Name)
	if fallbackType == nil {
		return nil
	}
	merged := c.createSymbolWithType(prop, c.getUnionType([]*Type{nonNilType, fallbackType}))
	// The key now always answers, whether the table has it or the metatable does.
	merged.Flags &^= ast.SymbolFlagsOptional
	return merged
}

// getMetatableFallbackType returns what __index answers for a key named name: the fallback
// table's own member, or the index signature that covers the key's domain.
func (c *Checker) getMetatableFallbackType(d *MetatableType, fallbackInfos []*IndexInfo, name string) *Type {
	if !d.indexIsFunction {
		if prop := c.getPropertyOfType(d.indexSource, name); prop != nil {
			return c.getTypeOfSymbol(prop)
		}
	}
	keyType := c.esSymbolType
	if !isLateBoundName(name) {
		keyType = c.keyTypeForPropertyName(name)
	}
	if info := c.findApplicableIndexInfo(fallbackInfos, keyType); info != nil {
		return info.valueType
	}
	return nil
}

// luaMetamethodForOperator maps a binary operator token to the metamethod Lua runs for it and
// whether the operands swap first: a > b runs __lt(b, a) and a >= b runs __le(b, a).
func luaMetamethodForOperator(operator ast.Kind) (string, bool) {
	switch operator {
	case ast.KindPlusToken:
		return "__add", false
	case ast.KindMinusToken:
		return "__sub", false
	case ast.KindAsteriskToken:
		return "__mul", false
	case ast.KindSlashToken:
		return "__div", false
	case ast.KindPercentToken:
		return "__mod", false
	case ast.KindAsteriskAsteriskToken:
		return "__pow", false
	case ast.KindDotDotToken:
		return "__concat", false
	case ast.KindLessThanToken:
		return "__lt", false
	case ast.KindGreaterThanToken:
		return "__lt", true
	case ast.KindLessThanEqualsToken:
		return "__le", false
	case ast.KindGreaterThanEqualsToken:
		return "__le", true
	}
	return "", false
}

// getLuaMetamethodSignatures returns the call signatures of the metamethod an operand's type
// commits to. A pairing reads the handler off its metatable and off nothing else -- setmetatable
// told the checker exactly where dispatch goes. Any other type that can carry ambient
// metamethods reads a metamethod-NAMED member off the type itself: the ambient-operator
// convention, by which a declared userdata type (GMod's Vector) states its metatable behavior as
// `__add`-style members that no visible setmetatable ever attaches. Optionality commits to
// nothing on either path, so a LuaMetatable<T>-typed value stays inert. Only OPERATOR
// metamethods dispatch ambiently -- __index, __newindex and __call keep meaning only what a
// real pairing's metatable says.
func (c *Checker) getLuaMetamethodSignatures(operandType *Type, name string) []*Signature {
	carrier := operandType
	if isMetatableType(operandType) {
		carrier = operandType.AsMetatableType().metatableType
	} else if !canCarryAmbientMetamethods(operandType) {
		return nil
	}
	handler := c.getMetatableHandlerType(carrier, name)
	if handler == nil {
		return nil
	}
	return c.getSignaturesOfType(handler, SignatureKindCall)
}

// hasLuaEqualityMetamethod reports whether either operand carries an __eq handler against a
// peer that could dispatch it. Lua only consults __eq between two tables (or two userdata), so a
// handler suppresses the no-overlap error only when the other operand could hold a table --
// `span == "hello"` stays flagged. A union peer counts when some arm could dispatch. The
// equality paths already yield boolean, so presence is all that matters -- no parameter check.
func (c *Checker) hasLuaEqualityMetamethod(leftType *Type, rightType *Type) bool {
	return len(c.getLuaMetamethodSignatures(leftType, "__eq")) != 0 && someType(rightType, c.typeMayBeLuaTable) ||
		len(c.getLuaMetamethodSignatures(rightType, "__eq")) != 0 && someType(leftType, c.typeMayBeLuaTable)
}

// typeMayBeLuaTable reports whether a value of this type could be a table (or userdata) at
// runtime -- the shapes an __eq dispatch needs on BOTH sides. A type variable answers by its
// base constraint, not its own non-primitive flag: `T extends string` never holds a table, while
// an unconstrained T might.
func (c *Checker) typeMayBeLuaTable(t *Type) bool {
	if t.flags&TypeFlagsInstantiableNonPrimitive != 0 {
		if constraint := c.getBaseConstraintOfType(t); constraint != nil {
			return someType(constraint, c.typeMayBeLuaTable)
		}
		return true
	}
	return t.flags&(TypeFlagsObject|TypeFlagsIntersection) != 0
}

// checkLuaBinaryMetamethod dispatches a binary operator to a paired operand's metamethod, or
// returns nil when neither operand carries one and ordinary checking stands. Lua consults the
// left operand's handler first, then the right's, and always passes the operands in dispatch
// order -- so a handler found on the right still receives the left operand first, and its
// parameter types say whether that is allowed. The first signature that admits both operands
// wins; when none does, the first names the mismatch. The result is the handler's first return
// value, as at any call boundary.
func (c *Checker) checkLuaBinaryMetamethod(operator ast.Kind, left *ast.Node, right *ast.Node, leftType *Type, rightType *Type) *Type {
	name, swapped := luaMetamethodForOperator(operator)
	if name == "" {
		return nil
	}
	if swapped {
		left, right = right, left
		leftType, rightType = rightType, leftType
	}
	signatures := c.getLuaMetamethodSignatures(leftType, name)
	if len(signatures) == 0 {
		signatures = c.getLuaMetamethodSignatures(rightType, name)
	}
	// A handler receives both operands whatever it declares, but a signature that does not name
	// an operand's position cannot check it: an under-arity handler commits to nothing (the
	// out-of-range position would read as `any` and admit every operand vacuously), and ordinary
	// operand checking stands. Deliberately AFTER the left-then-right selection: Lua would still
	// run the left operand's handler at runtime, so an under-arity left refuses dispatch outright
	// rather than falling through to a right handler that would never be consulted.
	signatures = core.Filter(signatures, func(s *Signature) bool {
		return c.tryGetTypeAtPosition(s, 0) != nil && c.tryGetTypeAtPosition(s, 1) != nil
	})
	if len(signatures) == 0 {
		return nil
	}
	signature := core.Find(signatures, func(s *Signature) bool {
		return c.isTypeAssignableTo(leftType, c.getTypeAtPosition(s, 0)) &&
			c.isTypeAssignableTo(rightType, c.getTypeAtPosition(s, 1))
	})
	if signature == nil {
		signature = signatures[0]
		c.checkTypeAssignableTo(leftType, c.getTypeAtPosition(signature, 0), left, nil)
		c.checkTypeAssignableTo(rightType, c.getTypeAtPosition(signature, 1), right, nil)
	}
	return c.adjustMultiReturn(c.getReturnTypeOfSignature(signature))
}

// checkLuaUnaryMetamethod dispatches a unary operator to an operand's metamethod. An
// under-arity handler commits to nothing, as at the binary sites.
func (c *Checker) checkLuaUnaryMetamethod(operandType *Type, name string) *Type {
	signatures := core.Filter(c.getLuaMetamethodSignatures(operandType, name), func(s *Signature) bool {
		return c.tryGetTypeAtPosition(s, 0) != nil
	})
	if len(signatures) == 0 {
		return nil
	}
	return c.adjustMultiReturn(c.getReturnTypeOfSignature(signatures[0]))
}

// getMetatableNewindexWriteType returns the type __newindex accepts for a write of a key the
// merged shape misses, or nil when the pairing commits to no __newindex or the handler does not
// answer this key. The table form receives the write, so its own member or index signature types
// it; the function form is called as (table, key, value), so its value parameter does. name is
// the written property's name, or "" for a computed key.
func (c *Checker) getMetatableNewindexWriteType(t *Type, keyType *Type, name string) *Type {
	if !isMetatableType(t) {
		return nil
	}
	d := t.AsMetatableType()
	if d.newindexSource == nil {
		return nil
	}
	if d.newindexIsFunction {
		if !c.isValidIndexArgumentType(keyType) {
			return nil
		}
		for _, signature := range c.getSignaturesOfType(d.newindexSource, SignatureKindCall) {
			if c.isTypeAssignableTo(keyType, c.getTypeAtPosition(signature, 1)) {
				return c.getTypeAtPosition(signature, 2)
			}
		}
		return nil
	}
	if name != "" {
		if prop := c.getPropertyOfType(d.newindexSource, name); prop != nil {
			return c.getWriteTypeOfSymbol(prop)
		}
	}
	if info := c.getApplicableIndexInfo(d.newindexSource, keyType); info != nil && !info.isReadonly {
		return info.valueType
	}
	return nil
}

// getMetatableNewindexWriteTypeForName is the named-member entry point for the write fallback.
func (c *Checker) getMetatableNewindexWriteTypeForName(t *Type, name string) *Type {
	keyType := c.esSymbolType
	if !isLateBoundName(name) {
		keyType = c.keyTypeForPropertyName(name)
	}
	return c.getMetatableNewindexWriteType(t, keyType, name)
}

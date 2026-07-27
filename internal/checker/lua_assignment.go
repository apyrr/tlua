package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
)

// skipLuaTypeOnlyWrappers removes assertions erased by emission without
// changing Lua's value arity. Parentheses are intentionally not included.
func skipLuaTypeOnlyWrappers(node *ast.Node) *ast.Node {
	return ast.SkipOuterExpressions(node, ast.OEKAssertions)
}

// skipLuaRuntimeTransparentWrappers is for identity and source-shape checks
// where parentheses and every erased assertion preserve the runtime value.
func skipLuaRuntimeTransparentWrappers(node *ast.Node) *ast.Node {
	return ast.SkipOuterExpressions(node, ast.OEKParentheses|ast.OEKAssertions)
}

// skipLuaEvolvingArrayWrappers removes wrappers that preserve both the runtime
// value and its inferred static type. An `as` assertion is intentionally not
// transparent: its annotated type must prevent automatic array evolution.
func skipLuaEvolvingArrayWrappers(node *ast.Node) *ast.Node {
	return ast.SkipOuterExpressions(node, ast.OEKParentheses|ast.OEKNonNullAssertions|ast.OEKSatisfies)
}

// outermostLuaWrapper walks from an expression to the outer erased wrapper
// that still denotes that expression. Callers choose whether parentheses are
// significant by excluding or including OEKParentheses.
func outermostLuaWrapper(node *ast.Node, kinds ast.OuterExpressionKinds) *ast.Node {
	for node.Parent != nil && ast.IsOuterExpression(node.Parent, kinds) && node.Parent.Expression() == node {
		node = node.Parent
	}
	return node
}

// luaAssignmentSlot names one positional store in a multiple-assignment commit.
// Every target expression is captured before any of the commit's stores occur.
// Targets and values are derived from the assignment on demand, so a slot stays
// a two-word value on the flow and checking hot paths.
type luaAssignmentSlot struct {
	assignment *ast.Node
	index      int
}

func (slot luaAssignmentSlot) targetList() *ast.Node {
	return slot.assignment.AsBinaryExpression().Left
}

func (slot luaAssignmentSlot) targetAt(index int) *ast.Node {
	return luaExplicitAssignmentValueAt(slot.targetList(), index)
}

func (slot luaAssignmentSlot) target() *ast.Node {
	return slot.targetAt(slot.index)
}

func (slot luaAssignmentSlot) explicitValue(index int) *ast.Node {
	return luaExplicitAssignmentValueAt(slot.assignment.AsBinaryExpression().Right, index)
}

// targetCount reports how many positional stores the commit performs.
func (slot luaAssignmentSlot) targetCount() int {
	if list := slot.targetList(); list.Kind == ast.KindExpressionList {
		return len(list.Elements())
	}
	return 1
}

// luaAssignmentSlotForNode normalizes scalar binaries and positional targets
// into the same slot consumed by checking, inference, and flow.
func luaAssignmentSlotForNode(node *ast.Node) (luaAssignmentSlot, bool) {
	if node == nil {
		return luaAssignmentSlot{}, false
	}
	var assignment *ast.Node
	var target *ast.Node
	if ast.IsBinaryExpression(node) {
		if node.AsBinaryExpression().Left.Kind == ast.KindExpressionList {
			return luaAssignmentSlot{}, false
		}
		assignment = node
		target = node.AsBinaryExpression().Left
	} else {
		target = node
		assignment = ast.GetAssignmentTarget(target)
		if assignment == nil || !ast.IsBinaryExpression(assignment) {
			return luaAssignmentSlot{}, false
		}
	}
	binary := assignment.AsBinaryExpression()
	if ast.IsInJSFile(assignment) || binary.OperatorToken.Kind != ast.KindEqualsToken {
		return luaAssignmentSlot{}, false
	}
	index := 0
	if binary.Left.Kind == ast.KindExpressionList {
		if index = ast.IndexOfNode(binary.Left.Elements(), target); index < 0 {
			return luaAssignmentSlot{}, false
		}
	} else if target != binary.Left {
		return luaAssignmentSlot{}, false
	}
	if ast.GetLuaAssignmentTargetReference(target) == nil {
		return luaAssignmentSlot{}, false
	}
	return luaAssignmentSlot{assignment: assignment, index: index}, true
}

// luaExplicitAssignmentValueAt returns the expression syntactically aligned
// with a target. It is for source-shape questions only; semantic values must be
// projected from the full value-list pack.
func luaExplicitAssignmentValueAt(valueList *ast.Node, index int) *ast.Node {
	if valueList.Kind == ast.KindExpressionList {
		elements := valueList.Elements()
		if index < len(elements) {
			return elements[index]
		}
		return nil
	}
	if index == 0 {
		return valueList
	}
	return nil
}

// luaAssignmentTargetForExplicitValue inverts luaExplicitAssignmentValueAt for
// direct RHS expressions. Surplus values have no assignment target.
func luaAssignmentTargetForExplicitValue(value *ast.Node) *ast.Node {
	valueList, index, ok := luaValueListPosition(value)
	if !ok || valueList.Parent == nil || !ast.IsBinaryExpression(valueList.Parent) {
		return nil
	}
	assignment := valueList.Parent.AsBinaryExpression()
	if assignment.OperatorToken.Kind != ast.KindEqualsToken || assignment.Right != valueList {
		return nil
	}
	return luaExplicitAssignmentValueAt(assignment.Left, index)
}

// luaValueListPosition inverts luaExplicitAssignmentValueAt: it maps a direct
// value expression to the list that holds it and its position in that list. A
// value outside a list is its own single-element list.
func luaValueListPosition(value *ast.Node) (valueList *ast.Node, index int, ok bool) {
	if value == nil || value.Parent == nil {
		return nil, 0, false
	}
	if value.Parent.Kind != ast.KindExpressionList {
		return value, 0, true
	}
	index = slices.Index(value.Parent.Elements(), value)
	return value.Parent, index, index >= 0
}

// getLuaAssignmentValueType projects one source expression without checking
// unrelated RHS arms. Assertions affect slot zero but preserve a call's arity.
func (c *Checker) getLuaAssignmentValueType(valueList *ast.Node, index int, checkMode CheckMode) *Type {
	expression, relativeIndex := luaAssignmentValueSourceAt(valueList, index)
	scalar := c.checkExpressionForMutableLocation(expression, checkMode)
	pack := c.getCallPackType(expression, checkMode)
	return c.getLuaAdjustedValueTypeAt(scalar, pack, relativeIndex)
}

// getLuaFlowAssignmentValueType preserves literal types used by control-flow
// narrowing while sharing Lua's positional and pack-adjustment rules.
func (c *Checker) getLuaFlowAssignmentValueType(valueList *ast.Node, index int) *Type {
	expression, relativeIndex := luaAssignmentValueSourceAt(valueList, index)
	scalar := c.getTypeOfExpression(expression)
	pack := c.getCallPackType(expression, CheckModeNormal)
	t := c.getLuaAdjustedValueTypeAt(scalar, pack, relativeIndex)
	if relativeIndex == 0 {
		t = c.finalizeLuaConstructorInitializerType(expression, t)
	}
	return t
}

func luaAssignmentValueSourceAt(valueList *ast.Node, index int) (*ast.Node, int) {
	if valueList.Kind == ast.KindExpressionList {
		elements := valueList.Elements()
		last := len(elements) - 1
		if index < last {
			return elements[index], 0
		}
		return elements[last], index - last
	}
	return valueList, index
}

type luaValuePresence uint8

const (
	luaValueAbsent luaValuePresence = iota
	luaValueMaybePresent
	luaValuePresent
)

// luaFirstValuePresence separates pack arity from value type at slot zero: a
// returned nil is still present, while a missing result is the nil Lua uses for
// padding.
func (c *Checker) luaFirstValuePresence(t *Type) luaValuePresence {
	if t.flags&TypeFlagsUnion != 0 {
		var result luaValuePresence
		for i, arm := range t.Types() {
			presence := c.luaFirstValuePresence(arm)
			if i == 0 {
				result = presence
			} else if result != presence {
				return luaValueMaybePresent
			}
		}
		return result
	}
	if t.flags&TypeFlagsVoid != 0 {
		return luaValueAbsent
	}
	if !isPackType(t) {
		return luaValuePresent
	}
	presence := luaValueAbsent
	for _, info := range t.TargetTupleType().elementInfos {
		if info.flags&ElementFlagsRequired != 0 {
			return luaValuePresent
		}
		if info.flags&(ElementFlagsFixed|ElementFlagsVariable) != 0 {
			presence = luaValueMaybePresent
		}
	}
	return presence
}

// getLuaAdjustedValueTypeAt combines a scalar view with its producer pack.
// This keeps wrapper-refined slot zero even when a generic tail is opaque.
func (c *Checker) getLuaAdjustedValueTypeAt(scalar *Type, pack *Type, index int) *Type {
	if pack == nil {
		return core.IfElse(index == 0, scalar, c.nilType)
	}
	if index != 0 {
		return c.packElementForIndex(pack, index)
	}
	switch c.luaFirstValuePresence(pack) {
	case luaValueAbsent:
		return c.nilType
	case luaValueMaybePresent:
		return c.getUnionType([]*Type{scalar, c.nilType})
	default:
		return scalar
	}
}

// canonicalLuaAliasSymbol follows identity-preserving local aliases to the
// binding that names one constructor. A binding assigned anywhere is not a
// stable alias, so the walk stops there rather than reasoning about where in
// the program the writes occur.
func (c *Checker) canonicalLuaAliasSymbol(symbol *ast.Symbol) *ast.Symbol {
	original := c.getMergedSymbol(symbol)
	symbol = original
	var seen collections.Set[*ast.Symbol]
	for symbol != nil {
		if !seen.AddIfAbsent(symbol) {
			return original
		}
		if len(symbol.Declarations) != 1 || !c.isLuaStableIdentityBinding(symbol) {
			return symbol
		}
		declaration := symbol.Declarations[0]
		if !ast.IsVariableDeclaration(declaration) || !ast.IsLuaLocal(declaration) {
			return symbol
		}
		initializer := luaExplicitVariableInitializer(declaration)
		if initializer == nil {
			return symbol
		}
		if initializer = skipLuaRuntimeTransparentWrappers(initializer); !ast.IsIdentifier(initializer) {
			return symbol
		}
		// Augmentation discovery may still create an implicit global with this
		// name, so do not cache an unresolved expression symbol here.
		next := c.getMergedSymbol(c.resolveName(
			initializer,
			initializer.Text(),
			ast.SymbolFlagsValue,
			nil,
			false, /*isUse*/
			false, /*excludeGlobals*/
		))
		if !c.isParameterOrMutableLocalVariable(next) {
			return symbol
		}
		if c.isSymbolAssigned(next) {
			return symbol
		}
		symbol = next
	}
	return original
}

// luaExplicitVariableInitializer returns only the source expression aligned
// with this local name. It never invents a value from a trailing call's pack.
func luaExplicitVariableInitializer(declaration *ast.Node) *ast.Node {
	if !ast.HasLuaLocalValueList(declaration) {
		return declaration.Initializer()
	}
	declarations := declaration.Parent.AsVariableDeclarationList().Declarations.Nodes
	index := ast.IndexOfNode(declarations, declaration)
	return luaExplicitAssignmentValueAt(ast.LuaLocalValueList(declaration.Parent), index)
}

func (c *Checker) isLuaStableIdentityBinding(symbol *ast.Symbol) bool {
	symbol = c.getMergedSymbol(symbol)
	return symbol != nil && c.isParameterOrMutableLocalVariable(symbol) && !c.isSymbolAssigned(symbol)
}

type luaStableAccessKey struct {
	root *ast.Symbol
	path []luaStableAccessPart
}

type luaStableAccessPart struct {
	name     string
	key      *ast.Symbol
	computed bool
}

type luaStableAccessKeyResult struct {
	key luaStableAccessKey
	ok  bool
}

// getLuaStableAccessKey proves storage identity for named accesses and computed
// accesses whose key binding cannot change. Bare identifiers remain distinct.
//
// Every input is a function of the bound program, so the answer is memoized per
// node. The flow engine asks for it on each visit of every assignment node, and
// recomputing walks the access chain and re-resolves the root name each time.
func (c *Checker) getLuaStableAccessKey(reference *ast.Node) (luaStableAccessKey, bool) {
	if cached, hit := c.luaStableAccessKeys[reference]; hit {
		return cached.key, cached.ok
	}
	key, ok := c.computeLuaStableAccessKey(reference)
	c.luaStableAccessKeys[reference] = luaStableAccessKeyResult{key: key, ok: ok}
	return key, ok
}

func (c *Checker) computeLuaStableAccessKey(reference *ast.Node) (luaStableAccessKey, bool) {
	reference = skipLuaRuntimeTransparentWrappers(reference)
	if !ast.IsAccessExpression(reference) {
		return luaStableAccessKey{}, false
	}

	var reversed []luaStableAccessPart
	expression := reference
	for ast.IsAccessExpression(expression) {
		if name, ok := c.getAccessedPropertyName(expression); ok {
			reversed = append(reversed, luaStableAccessPart{name: name})
		} else if ast.IsElementAccessExpression(expression) {
			index := skipLuaRuntimeTransparentWrappers(expression.AsElementAccessExpression().ArgumentExpression)
			if !ast.IsIdentifier(index) {
				return luaStableAccessKey{}, false
			}
			symbol := c.getMergedSymbol(c.getResolvedSymbol(index))
			if !c.isConstantVariable(symbol) &&
				!(c.isParameterOrMutableLocalVariable(symbol) && !c.isSymbolAssigned(symbol)) {
				return luaStableAccessKey{}, false
			}
			reversed = append(reversed, luaStableAccessPart{key: symbol, computed: true})
		} else {
			return luaStableAccessKey{}, false
		}
		expression = skipLuaRuntimeTransparentWrappers(expression.Expression())
	}
	if !ast.IsIdentifier(expression) {
		return luaStableAccessKey{}, false
	}
	path := make([]luaStableAccessPart, len(reversed))
	for i := range reversed {
		path[len(reversed)-1-i] = reversed[i]
	}
	root := c.getMergedSymbol(c.resolveLuaRoot(reference, expression.Text()))
	if root == nil {
		return luaStableAccessKey{}, false
	}
	root = c.canonicalLuaAliasSymbol(root)
	if root == c.luaGlobalsSymbol {
		for len(path) != 0 && !path[0].computed && path[0].name == c.luaGlobalsSymbol.Name {
			path = path[1:]
		}
	}
	return luaStableAccessKey{root: root, path: path}, true
}

func (c *Checker) isSameLuaStableAccess(left *ast.Node, right *ast.Node) bool {
	leftKey, leftOK := c.getLuaStableAccessKey(left)
	if !leftOK {
		// Most callers pass something that is not a static access at all, so do
		// not key the other side before this one has proven it can match.
		return false
	}
	rightKey, rightOK := c.getLuaStableAccessKey(right)
	return rightOK && leftKey.root == rightKey.root &&
		slices.EqualFunc(leftKey.path, rightKey.path, func(left luaStableAccessPart, right luaStableAccessPart) bool {
			return left.computed == right.computed && core.IfElse(left.computed, left.key == right.key, left.name == right.name)
		})
}

func (c *Checker) isSameLuaCapturedTargetStorage(left *ast.Node, right *ast.Node) bool {
	if c.isSameLuaStableAccess(left, right) {
		return true
	}
	left = skipLuaRuntimeTransparentWrappers(left)
	right = skipLuaRuntimeTransparentWrappers(right)
	if ast.IsAccessExpression(left) && ast.IsAccessExpression(right) {
		leftName, leftNamed := c.getAccessedPropertyName(left)
		rightName, rightNamed := c.getAccessedPropertyName(right)
		if leftNamed || rightNamed {
			if !leftNamed || !rightNamed || leftName != rightName {
				return false
			}
		} else {
			if !ast.IsElementAccessExpression(left) || !ast.IsElementAccessExpression(right) ||
				!c.isSameLuaCapturedTargetStorage(
					left.AsElementAccessExpression().ArgumentExpression,
					right.AsElementAccessExpression().ArgumentExpression,
				) {
				return false
			}
		}
		return c.isSameLuaCapturedTargetStorage(left.Expression(), right.Expression())
	}
	return c.isSameLuaCapturedReference(left, right)
}

// winningTargetStorageIndex returns the leftmost matching target because LuaJIT
// commits stores right-to-left, making that store the observable winner.
func (slot luaAssignmentSlot) winningTargetStorageIndex(c *Checker, reference *ast.Node) (int, bool) {
	for index := range slot.targetCount() {
		target := slot.targetAt(index)
		targetReference := ast.GetLuaAssignmentTargetReference(target)
		if targetReference != nil && c.isSameLuaCapturedTargetStorage(reference, targetReference) {
			return index, true
		}
	}
	return -1, false
}

func (slot luaAssignmentSlot) capturedTargetForValue(c *Checker, value *ast.Node) (*ast.Node, bool) {
	index, ok := slot.winningTargetStorageIndex(c, value)
	if !ok {
		return nil, false
	}
	return slot.targetAt(index), true
}

// isLuaCapturedReceiverValue recognizes recursive stores such as x.self = x.
// Snapshot inference cannot make an ancestor receiver more precise while one
// of its own members is being inferred, so these use ordinary cycle recovery.
func (c *Checker) isLuaCapturedReceiverValue(target *ast.Node, value *ast.Node) bool {
	if _, ok := luaAssignmentSlotForNode(target); !ok {
		return false
	}
	reference := ast.GetLuaAssignmentTargetReference(target)
	for ast.IsAccessExpression(reference) {
		reference = skipLuaRuntimeTransparentWrappers(reference.Expression())
		if c.isSameLuaCapturedTargetStorage(reference, value) {
			return true
		}
	}
	return false
}

func (c *Checker) isOverwrittenLuaCapturedTarget(target *ast.Node) bool {
	slot, ok := luaAssignmentSlotForNode(target)
	if !ok || slot.targetCount() < 2 {
		// Only a sibling store can supersede this one, so a lone target is never
		// overwritten. The flow engine asks on every assignment it visits, and
		// proving it the long way costs a storage-identity walk per question.
		return false
	}
	winner, ok := slot.winningTargetStorageIndex(c, slot.target())
	return ok && winner != slot.index
}

// A store such as a = a restores the exact value captured before the
// transaction. Treating it as a new flow type would hide sibling mutations.
func (c *Checker) isSelfPreservingLuaCapturedTarget(target *ast.Node) bool {
	slot, ok := luaAssignmentSlotForNode(target)
	if !ok {
		return false
	}
	value := slot.explicitValue(slot.index)
	return value != nil && c.isSameLuaCapturedTargetStorage(slot.target(), value)
}

// luaStoreMayPreserveTarget reports whether a store's value can evaluate to the
// stored target itself. `and`/`or` yield one of their operands, so any operand
// in result position that names the target's storage keeps the table -- and its
// metatable -- on that path. Deliberately over-approximate: a left `and`
// operand only results when falsy, but treating it as preserving merely retains
// a pairing, the conservative direction.
func (c *Checker) luaStoreMayPreserveTarget(target *ast.Node, value *ast.Node) bool {
	value = skipLuaTypeOnlyWrappers(value)
	if ast.IsBinaryExpression(value) {
		binary := value.AsBinaryExpression()
		if binary.OperatorToken.Kind == ast.KindAmpersandAmpersandToken || binary.OperatorToken.Kind == ast.KindBarBarToken {
			return c.luaStoreMayPreserveTarget(target, binary.Left) || c.luaStoreMayPreserveTarget(target, binary.Right)
		}
	}
	return c.isSameLuaCapturedTargetStorage(target, value)
}

// isSelfPreservingLuaAssignmentValue recognizes the RHS read paired with a
// target such as a in a, a[1] = a, value. Reading an otherwise-empty evolving
// array here must not force it to any[] before the sibling mutation is applied.
func (c *Checker) isSelfPreservingLuaAssignmentValue(node *ast.Node) bool {
	target := luaAssignmentTargetForExplicitValue(outermostLuaWrapper(node, ast.OEKParentheses|ast.OEKAssertions))
	return target != nil && c.isSelfPreservingLuaCapturedTarget(target)
}

// A direct local alias does not consume an empty evolving array. Both locals
// still point at the same runtime table, and a later indexed write may provide
// the element type.
func isLuaDirectAliasInitializerValue(node *ast.Node) bool {
	value := outermostLuaWrapper(node, ast.OEKParentheses|ast.OEKAssertions)
	reference := skipLuaRuntimeTransparentWrappers(value)
	if !ast.IsIdentifier(reference) || value.Parent == nil {
		return false
	}
	valueList, index, ok := luaValueListPosition(value)
	if !ok {
		return false
	}
	holder := valueList.Parent
	if !ast.IsVariableDeclaration(holder) || !ast.IsLuaLocal(holder) {
		return false
	}
	if valueList == holder.Initializer() && value.Parent.Kind != ast.KindExpressionList {
		return true
	}
	list := holder.Parent
	if list == nil || list.Kind != ast.KindVariableDeclarationList {
		return false
	}
	declarations := list.AsVariableDeclarationList().Declarations.Nodes
	return index < len(declarations) && luaExplicitVariableInitializer(declarations[index]) == value
}

// isSameLuaCapturedReference compares addresses within one LHS-capture phase.
// Unlike general flow matching, a mutable index such as i is stable between
// the two plain reads in a[i], a[i], even if i was assigned earlier.
func (c *Checker) isSameLuaCapturedReference(left *ast.Node, right *ast.Node) bool {
	if left == nil || right == nil {
		return false
	}
	left = skipLuaRuntimeTransparentWrappers(left)
	right = skipLuaRuntimeTransparentWrappers(right)
	if left == right {
		return true
	}

	leftGlobal := c.luaGlobalReferenceSymbol(left)
	rightGlobal := c.luaGlobalReferenceSymbol(right)
	if leftGlobal != nil && leftGlobal == rightGlobal ||
		c.luaGlobalReferenceMatches(leftGlobal, right) ||
		c.luaGlobalReferenceMatches(rightGlobal, left) {
		return true
	}

	switch left.Kind {
	case ast.KindIdentifier, ast.KindPrivateIdentifier:
		return right.Kind == left.Kind && c.getResolvedSymbol(left) == c.getResolvedSymbol(right)
	case ast.KindThisKeyword, ast.KindSuperKeyword:
		return right.Kind == left.Kind
	}
	return false
}

// isLuaCapturedReceiverRebound finds an effective store that makes an access
// target's captured receiver differ from the same syntactic path afterward.
// Assigning a receiver back to its captured value is deliberately preserved.
func (c *Checker) isLuaCapturedReceiverRebound(target *ast.Node) bool {
	slot, ok := luaAssignmentSlotForNode(target)
	if !ok || slot.targetCount() < 2 {
		// A store to `x.a.b` cannot change what `x` or `x.a` name, so only a
		// sibling can rebind a receiver. Without this the walk proves that again
		// for every receiver level of every single-target member assignment.
		return false
	}
	reference := ast.GetLuaAssignmentTargetReference(target)
	if !ast.IsAccessExpression(reference) {
		return false
	}

	for ast.IsAccessExpression(reference) {
		receiver := skipLuaRuntimeTransparentWrappers(reference.Expression())
		if !ast.IsIdentifier(receiver) && !ast.IsAccessExpression(receiver) {
			break
		}
		if index, assigned := slot.winningTargetStorageIndex(c, receiver); assigned {
			value := slot.explicitValue(index)
			if value == nil || !c.isSameLuaCapturedTargetStorage(receiver, value) {
				return true
			}
		}
		reference = receiver
	}
	return false
}

// checkLuaAssignment checks every target against its adjusted positional value.
// Missing values become nil and a trailing call contributes its full pack.
func (c *Checker) checkLuaAssignment(left *ast.Node, right *ast.Node, checkMode CheckMode) *Type {
	targets := []*ast.Node{left}
	if left.Kind == ast.KindExpressionList {
		targets = left.Elements()
	}
	pack := c.getPackTypeOfValueList(right, checkMode)
	for i, target := range targets {
		symbol := c.getMergedSymbol(c.luaAugmentationTargets[target])
		var targetType *Type
		// A recovered or non-lvalue target has no reference; it is still checked
		// as an ordinary expression so the assignment reports its own diagnostic.
		reference := ast.GetLuaAssignmentTargetReference(target)
		isAccessTarget := reference != nil && ast.IsAccessExpression(reference)
		isAugmentation := symbol != nil && symbol.Flags&ast.SymbolFlagsAssignment != 0
		// Reading the receiver's contract resolves its type, so only a write that
		// declares a member asks. Every other store ignores the answer.
		usesReceiverContract := isAugmentation && isAccessTarget &&
			c.luaAugmentationUsesReceiverContract(target, reference)
		if isAugmentation && isAccessTarget &&
			(c.isLuaCapturedReceiverRebound(target) || !usesReceiverContract) {
			targetType = c.checkLuaAugmentationTarget(target, symbol, checkMode)
		} else {
			targetType = c.checkExpressionEx(target, checkMode)
		}
		if isAugmentation && len(c.luaMetatablePairings[symbol]) != 0 {
			// A store replaces the table, so it checks against the storage the
			// declared-type pairing wrapped -- the pairing describes what the
			// setmetatable statements install, not what a store must supply.
			targetType = c.getUnpairedTableType(targetType)
		}
		overwritten := c.isOverwrittenLuaCapturedTarget(target)
		if overwritten {
			// Every Lua store executes, even when a later store to the same captured
			// address determines the final value. Always validate the reference.
			c.checkAssignmentReference(target)
			// A newly inferred augmentation has no pre-existing contract for the
			// transient value. Its declaration and post-statement flow use only the
			// winning store, while declared targets continue through compatibility.
			if isAugmentation && !usesReceiverContract {
				continue
			}
		}
		rightType := c.getLuaAssignmentValueType(right, i, checkMode)
		if symbol != nil {
			// Declaration inference may intentionally see through an erased
			// assertion to the table literal installed at runtime. Check that
			// declaration against the same positional type it was inferred from.
			if inferredType, ok := c.getLuaAugmentationTargetInitializerType(symbol, target); ok {
				rightType = inferredType
			}
		}
		// Elaborate against the value actually aligned with this target; the whole
		// list is not an object or array literal, so it cannot carry elaboration.
		valueExpression := luaExplicitAssignmentValueAt(right, i)
		if valueExpression == nil {
			valueExpression = right
		}
		c.checkAssignmentOperator(target, ast.KindEqualsToken, valueExpression, targetType, rightType)
	}
	return c.packElementForIndex(pack, 0)
}

func (c *Checker) luaAugmentationUsesReceiverContract(target *ast.Node, reference *ast.Node) bool {
	if c.luaNumericContractTargets.Has(target) {
		return true
	}
	name, named := c.getAccessedPropertyName(reference)
	if !named {
		return false
	}
	receiverType := c.getTypeOfExpression(reference.Expression())
	if property := c.getPropertyOfType(receiverType, name); property != nil &&
		c.getMergedSymbol(property) == c.getMergedSymbol(c.luaAugmentationTargets[target]) {
		return false
	}
	return c.luaContractOwnsProperty(receiverType, name)
}

// A synthesized member belongs to the constructor arms its path reaches, which
// the receiver's own declaration type need not describe. Validate the receiver
// normally, then use the member's checker-local declaration type.
func (c *Checker) checkLuaAugmentationTarget(target *ast.Node, symbol *ast.Symbol, checkMode CheckMode) *Type {
	reference := ast.GetLuaAssignmentTargetReference(target)
	if reference != nil && ast.IsAccessExpression(reference) {
		c.checkExpressionEx(reference.Expression(), checkMode)
		if ast.IsElementAccessExpression(reference) {
			c.checkExpressionEx(reference.AsElementAccessExpression().ArgumentExpression, checkMode)
		}
	}
	return c.getTypeOfSymbol(symbol)
}

// Lua assignment targets are declaration spans but not declaration-shaped AST
// nodes, so their owning symbol carries the parent directly.
func getDeclarationParentOfSymbol(symbol *ast.Symbol) *ast.Symbol {
	// A synthesized union or intersection property inherits its constituent's
	// value declaration without inheriting the assignment flag, so a Lua target
	// reaches here on a plain property too. The symbol's own parent answers both
	// shapes; nothing about this is worth panicking the compiler over.
	if declaration := symbol.ValueDeclaration; declaration != nil {
		if declarationSymbol := declaration.Symbol(); declarationSymbol != nil {
			return declarationSymbol.Parent
		}
	}
	return symbol.Parent
}

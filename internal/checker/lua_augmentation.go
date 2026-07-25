package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
)

type luaAugmentation struct {
	ast.LuaWriteCandidate
	name string
}

func (item luaAugmentation) declaration() *ast.Node {
	if item.Target != nil {
		return item.Target
	}
	return item.Source
}

// initializeLuaAugmentations resolves Lua assignment declarations only after all
// files' ordinary globals exist. It operates on checker-local transient symbols.
func (c *Checker) initializeLuaAugmentations() {
	var candidates []ast.LuaWriteCandidate
	for _, file := range c.files {
		candidates = append(candidates, file.LuaWriteCandidates...)
	}

	implicit := make(map[string][]luaAugmentation)
	for _, candidate := range candidates {
		node := candidate.Source
		// Discover every implicit global before stable-alias analysis walks and
		// resolves arbitrary assignment targets in the source file.
		rootName, root, path, ok := c.resolveLexicalLuaAugmentationPath(candidate)
		if !ok {
			continue
		}
		name := rootName
		bare := len(path) == 0 && candidate.Target != nil && ast.IsIdentifier(candidate.Target)
		switch {
		case bare && root == nil:
			// Bare unresolved assignments create implicit globals.
		case bare && root != nil && root.ValueDeclaration != nil &&
			ast.IsVariableDeclaration(root.ValueDeclaration) && ast.IsLuaLocal(root.ValueDeclaration):
			// A bare store to a local declares no member, so it joins no attachment
			// group. It still contributes its constructor to that local's storage,
			// which is what lets a rebound local keep naming the tables it can hold.
			// Deliberately no luaAugmentationTargets entry: the store declares
			// nothing, so it must keep checking against the local's own declared
			// type rather than an inferred augmentation contract.
			c.luaAssignmentAugmentations[root] = append(
				c.luaAssignmentAugmentations[root],
				luaAugmentation{LuaWriteCandidate: candidate},
			)
			continue
		case root == c.luaGlobalsSymbol && len(path) == 1:
			// Static writes through the built-in environment bypass lexical shadows
			// of the member name and create the same global as a bare assignment.
			name = path[0]
			if ast.IsNumberKeyName(name) {
				continue
			}
			if existing := c.getMergedSymbol(c.globals[name]); existing != nil && existing.Flags&ast.SymbolFlagsValue != 0 {
				if ast.IsFunctionDeclaration(node) {
					// The function name is not itself an access expression, so map its
					// binder symbol when the target is an actual environment property.
					if isLuaEnvironmentExport(existing) {
						c.recordMergedSymbol(existing, candidate.Symbol)
					}
				}
				continue
			}
		default:
			continue
		}
		// Recovered assignments in a syntactically invalid file are not reliable
		// global declarations and can hide the primary diagnostics.
		if len(ast.GetSourceFileOfNode(node).Diagnostics()) != 0 {
			continue
		}
		implicit[name] = append(implicit[name], luaAugmentation{LuaWriteCandidate: candidate})
	}
	for name, declarations := range implicit {
		symbol, collision := c.newLuaAugmentationSymbol(
			declarations,
			name,
			ast.SymbolFlagsFunctionScopedVariable,
			ast.SymbolFlagsFunction,
		)
		// A type-only global (an interface, say) may already own the name; merge
		// rather than clobber, as declared globals do.
		c.mergeGlobalSymbol(symbol)
		merged := c.globals[name]
		// mergeGlobalSymbol may clone a pre-existing type-only global, so collision
		// recovery and declaration mappings belong on the final canonical symbol.
		c.finishLuaAugmentationSymbol(merged, declarations, collision)
	}

	// Order candidates by path depth so a shallower member is attached before a
	// deeper path descends through it. Constructor-arm identity, not lexical
	// spelling, decides which writes form one group.
	type depthOrdered struct {
		item  luaAugmentation
		depth int
	}
	var ordered []depthOrdered
	for _, candidate := range candidates {
		_, root, path, ok := c.resolveLexicalLuaAugmentationPath(candidate)
		if !ok || len(path) == 0 {
			continue
		}
		if root == nil {
			continue
		}
		if root == c.luaGlobalsSymbol {
			// One-segment environment writes were handled as globals above. Deeper
			// writes augment the canonical global named by the first segment.
			if len(path) == 1 {
				continue
			}
			// Only the remaining path matters here; arms are resolved from the
			// candidate's own target expression when the group is attached.
			_, path, ok = c.rerootLuaEnvironmentPath(root, path)
			if !ok {
				continue
			}
		}
		ordered = append(ordered, depthOrdered{
			item: luaAugmentation{
				LuaWriteCandidate: candidate,
				name:              path[len(path)-1],
			},
			depth: len(path),
		})
	}
	slices.SortStableFunc(ordered, func(left, right depthOrdered) int {
		return left.depth - right.depth
	})
	pending := make([]luaAugmentation, 0, len(ordered))
	for _, entry := range ordered {
		pending = append(pending, entry.item)
	}
	for len(pending) != 0 {
		// An alias can point through a deeper augmentation path. Retry unresolved
		// groups after each attachment pass with a fresh topology cache. Resolve a
		// whole pass before mutating any arm so map order cannot choose a winner.
		constructorArms := newLuaConstructorResolver(c)
		deferred, attached := c.attachLuaAugmentations(pending, constructorArms)
		if !attached {
			break
		}
		pending = deferred
	}
}

// isLuaEnvironmentExport mirrors the filter used to build `typeof _G`.
func isLuaEnvironmentExport(symbol *ast.Symbol) bool {
	return symbol != nil && symbol.Flags&ast.SymbolFlagsBlockScoped == 0 &&
		!(symbol.Flags&ast.SymbolFlagsValueModule != 0 && len(symbol.Declarations) != 0 && core.Every(symbol.Declarations, ast.IsAmbientModule))
}

func (c *Checker) luaEnvironmentExport(name string) *ast.Symbol {
	symbol := c.getMergedSymbol(c.globals[name])
	if !isLuaEnvironmentExport(symbol) {
		return nil
	}
	return symbol
}

// resolveLexicalLuaAugmentationPath deliberately leaves local aliases distinct.
// The implicit-global phase runs before alias canonicalization, so it must not
// resolve otherwise-unknown names through a binding it has yet to create.
func (c *Checker) resolveLexicalLuaAugmentationPath(candidate ast.LuaWriteCandidate) (rootName string, root *ast.Symbol, path []string, ok bool) {
	rootName, path, ok = luaAugmentationPath(candidate)
	if !ok {
		return "", nil, nil, false
	}
	reference := candidate.Target
	if reference == nil {
		reference = candidate.Source
	}
	root, path = c.resolveLexicalLuaPathRoot(reference, rootName, path)
	return rootName, root, path, true
}

func (c *Checker) resolveLexicalLuaPathRoot(reference *ast.Node, rootName string, path []string) (*ast.Symbol, []string) {
	root := c.resolveLuaRoot(reference, rootName)
	if root != nil {
		root = c.getMergedSymbol(root)
		for root == c.luaGlobalsSymbol && len(path) != 0 && path[0] == c.luaGlobalsSymbol.Name {
			path = path[1:]
		}
	}
	return root, path
}

func (c *Checker) rerootLuaEnvironmentPath(root *ast.Symbol, path []string) (*ast.Symbol, []string, bool) {
	if root != c.luaGlobalsSymbol || len(path) == 0 {
		return root, path, root != nil
	}
	root = c.luaEnvironmentExport(path[0])
	if root == nil {
		return nil, nil, false
	}
	return root, path[1:], true
}

func (c *Checker) resolveLuaRoot(reference *ast.Node, name string) *ast.Symbol {
	// resolveName's scope walk already consults ast.LookupLuaLocal position-sensitively.
	return c.resolveName(reference, name, ast.SymbolFlagsValue, nil, false /*isUse*/, false /*excludeGlobals*/)
}

// luaAccessLevelName returns the mangled member name for one access-expression
// level (dot, string, or numeric key), or false when the level has no static
// name. Dot and string spellings collapse to the same name; numeric keys mangle
// into their own namespace.
func luaAccessLevelName(access *ast.Node) (string, bool) {
	name := ast.GetElementOrPropertyAccessName(access)
	if name == nil {
		return "", false
	}
	return ast.GetPropertyNameForPropertyNameNode(name), true
}

func luaAugmentationPath(candidate ast.LuaWriteCandidate) (string, []string, bool) {
	node := candidate.Source
	expression := candidate.Target
	var methodName string
	if ast.IsFunctionDeclaration(node) {
		expression = node.AsFunctionDeclaration().Target
		if expression == nil {
			// A bare function that writes a lexical local has a standalone
			// implementation symbol but is not a structural augmentation.
			return "", nil, false
		}
		name := ast.GetNameOfDeclaration(node)
		if name == nil {
			return "", nil, false
		}
		methodName = ast.GetPropertyNameForPropertyNameNode(name)
	} else if !ast.IsBinaryExpression(node) || expression == nil {
		return "", nil, false
	}
	rootName, path, ok := luaEntityNamePath(expression, luaAccessLevelName)
	if !ok {
		return "", nil, false
	}
	if methodName != "" {
		path = append(path, methodName)
	}
	return rootName, path, true
}

// accessName keeps syntactic augmentation eligibility distinct from the
// checker-resolved static names used for guard identity.
func luaEntityNamePath(expression *ast.Node, accessName func(*ast.Node) (string, bool)) (string, []string, bool) {
	var reversed []string
	// A postfix assertion is erased, but whole-expression parentheses remain
	// significant for exact default-guard recognition.
	expression = skipLuaTypeOnlyWrappers(expression)
	for ast.IsAccessExpression(expression) {
		name, ok := accessName(expression)
		if !ok {
			return "", nil, false
		}
		reversed = append(reversed, name)
		// Parentheses and assertions around a receiver do not change the entity
		// denoted by the enclosing access.
		expression = skipLuaRuntimeTransparentWrappers(expression.Expression())
	}
	if !ast.IsIdentifier(expression) {
		return "", nil, false
	}
	path := make([]string, 0, len(reversed))
	for i := len(reversed) - 1; i >= 0; i-- {
		path = append(path, reversed[i])
	}
	return expression.Text(), path, true
}

func (c *Checker) attachLuaAugmentations(pending []luaAugmentation, constructorArms *luaConstructorResolver) (deferred []luaAugmentation, changed bool) {
	type resolvedItem struct {
		item luaAugmentation
		arms []*ast.Symbol
	}
	var items []resolvedItem
	for _, item := range pending {
		itemArms, known := constructorArms.parentArms(item)
		if !known {
			deferred = append(deferred, item)
			continue
		}
		arms := appendLuaConstructorArms(c, nil, itemArms)
		if len(arms) == 0 {
			// A deeper path descends through a member that a shallower pass has yet
			// to attach. Defer instead of dropping the write; the loop stops once a
			// pass attaches nothing new.
			deferred = append(deferred, item)
			continue
		}
		items = append(items, resolvedItem{item: item, arms: arms})
	}
	if len(items) == 0 {
		return deferred, false
	}

	// Merge same-name writes with transitively overlapping constructor arms, even
	// when lexical aliases placed them in different forest branches. Writes onto
	// disjoint arm sets stay separate groups and declare separate members.
	parent := make([]int, len(items))
	for index := range parent {
		parent[index] = index
	}
	var find func(int) int
	find = func(index int) int {
		if parent[index] != index {
			parent[index] = find(parent[index])
		}
		return parent[index]
	}
	union := func(left int, right int) {
		left = find(left)
		right = find(right)
		if left != right {
			parent[right] = left
		}
	}
	type ownerKey struct {
		name string
		arm  *ast.Symbol
	}
	owners := make(map[ownerKey]int)
	for index, item := range items {
		name := item.item.name
		for _, arm := range item.arms {
			key := ownerKey{name: name, arm: arm}
			if owner, exists := owners[key]; exists {
				union(index, owner)
			} else {
				owners[key] = index
			}
		}
	}

	type component struct {
		items []luaAugmentation
		arms  []*ast.Symbol
		seen  collections.Set[*ast.Symbol]
	}
	var components []*component
	byRoot := make(map[int]*component)
	for index, item := range items {
		root := find(index)
		current := byRoot[root]
		if current == nil {
			current = &component{}
			byRoot[root] = current
			components = append(components, current)
		}
		current.items = append(current.items, item.item)
		for _, arm := range item.arms {
			if current.seen.AddIfAbsent(arm) {
				current.arms = append(current.arms, arm)
			}
		}
	}

	for _, component := range components {
		changed = c.attachLuaAugmentationComponent(component.items, component.arms) || changed
	}
	return deferred, changed
}

// A member is installed into every arm of its component, so reaching several
// arms never makes it optional.
func (c *Checker) attachLuaAugmentationComponent(group []luaAugmentation, arms []*ast.Symbol) bool {
	name := group[0].name
	var synthesized []*ast.Symbol
	for _, arm := range arms {
		arm = c.getMergedSymbol(arm)
		if existing := core.OrElse(arm.Exports[name], arm.Members[name]); existing != nil {
			existing = c.getMergedSymbol(existing)
			if len(c.luaAssignmentAugmentations[existing]) != 0 {
				synthesized = core.AppendIfUnique(synthesized, existing)
				continue
			}
			// Existing members keep their declared contract, so an ordinary write is
			// checked against it. A function body is a declaration, not a write, so
			// it always duplicates a member the constructor already declares.
			var methods []*ast.Node
			for _, declaration := range existing.Declarations {
				if ast.IsFunctionDeclaration(declaration) {
					methods = append(methods, declaration)
				}
			}
			for _, item := range group {
				if ast.IsFunctionDeclaration(item.Source) {
					methods = append(methods, item.declaration())
				}
			}
			if len(methods) != 0 {
				c.reportLuaMethodCollision(methods, name)
			}
			return false
		}
	}
	if ast.IsNumberKeyName(name) && core.Some(arms, luaConstructorBlocksNumericEvolution) {
		// Keep the runtime constructor member for navigation and aliases, but
		// check this write through the asserted surface that blocks array evolution.
		// Marked only once the component is known to attach: this set lives for the
		// checker's lifetime, so an abandoned component must leave nothing behind.
		for _, item := range group {
			if item.Target != nil {
				c.luaNumericContractTargets.Add(item.Target)
			}
		}
	}
	if len(synthesized) != 0 {
		var combined []luaAugmentation
		for _, existing := range synthesized {
			combined = append(combined, c.luaAssignmentAugmentations[existing]...)
			arms = appendLuaConstructorArms(c, arms, c.luaAugmentationMemberArms[existing])
		}
		group = dedupeLuaAugmentations(append(combined, group...))
	}
	member, collision := c.newLuaAugmentationSymbol(group, name, ast.SymbolFlagsProperty, ast.SymbolFlagsMethod)
	c.finishLuaAugmentationSymbol(member, group, collision)
	for _, existing := range synthesized {
		c.recordMergedSymbol(member, existing)
	}
	for _, arm := range arms {
		mergedArm := c.getMergedSymbol(arm)
		if mergedArm.Flags&ast.SymbolFlagsTransient == 0 {
			mergedArm = c.cloneSymbol(mergedArm)
		}
		ast.GetExports(mergedArm)[name] = member
		// The first arm is the canonical parent; later arms only host the member.
		if member.Parent == nil {
			member.Parent = mergedArm
		}
		c.luaAugmentationMemberArms[member] = appendLuaConstructorArms(c, c.luaAugmentationMemberArms[member], []*ast.Symbol{mergedArm})
	}
	return true
}

func dedupeLuaAugmentations(items []luaAugmentation) []luaAugmentation {
	result := make([]luaAugmentation, 0, len(items))
	seen := make(map[*ast.Node]struct{}, len(items))
	for _, item := range items {
		declaration := item.declaration()
		if _, exists := seen[declaration]; exists {
			continue
		}
		seen[declaration] = struct{}{}
		result = append(result, item)
	}
	return result
}

func (c *Checker) newLuaAugmentationSymbol(group []luaAugmentation, name string, valueFlags ast.SymbolFlags, methodFlags ast.SymbolFlags) (*ast.Symbol, bool) {
	methodCount := 0
	for _, item := range group {
		if ast.IsFunctionDeclaration(item.Source) {
			methodCount++
		}
	}
	hasMethod := methodCount != 0
	// A dotted function is Lua assignment sugar, so ordinary writes are checked
	// against its method type. Only two declaration-shaped bodies are duplicates.
	collision := methodCount > 1
	if collision {
		c.reportLuaAugmentationCollision(group, name)
	}
	flags := valueFlags | ast.SymbolFlagsAssignment
	if hasMethod && !collision {
		flags = methodFlags | ast.SymbolFlagsAssignment
	}
	symbol := c.newSymbol(flags, name)
	for _, item := range group {
		symbol.Declarations = append(symbol.Declarations, item.declaration())
		if symbol.ValueDeclaration == nil && item.Symbol.ValueDeclaration != nil {
			symbol.ValueDeclaration = item.Symbol.ValueDeclaration
		}
	}
	return symbol, collision
}

func (c *Checker) finishLuaAugmentationSymbol(symbol *ast.Symbol, group []luaAugmentation, collision bool) {
	if collision {
		// Keep recovery neutral and source-order independent. The declarations are
		// still one navigation group, but neither incompatible declaration wins.
		c.valueSymbolLinks.Get(symbol).resolvedType = c.anyType
	}
	for _, item := range group {
		c.recordMergedSymbol(symbol, item.Symbol)
		if item.Target != nil {
			c.luaAssignmentAugmentations[symbol] = append(c.luaAssignmentAugmentations[symbol], item)
			c.luaAugmentationTargets[item.Target] = symbol
		}
	}
}

func (c *Checker) reportLuaAugmentationCollision(group []luaAugmentation, name string) {
	var methods []*ast.Node
	for _, item := range group {
		if ast.IsFunctionDeclaration(item.Source) {
			methods = append(methods, item.declaration())
		}
	}
	c.reportLuaMethodCollision(methods, name)
}

func (c *Checker) reportLuaMethodCollision(methods []*ast.Node, name string) {
	displayName := ast.NumberKeyDisplayName(name)
	for _, declaration := range methods {
		if !c.luaReportedMethodCollisions.AddIfAbsent(declaration) {
			continue
		}
		related := make([]*ast.Node, 0, len(methods)-1)
		for _, other := range methods {
			if other != declaration {
				related = append(related, other)
			}
		}
		c.addDuplicateDeclarationError(declaration, diagnostics.Duplicate_identifier_0, displayName, related)
	}
}

// luaDefaultedAugmentationInitializer recognizes exactly `X = X or E`.
// Parenthesized and chained guards intentionally do not match.
func (c *Checker) luaDefaultedAugmentationInitializer(target *ast.Node, initializer *ast.Node) *ast.Node {
	if !c.isLuaDefaultedAugmentationGuard(target, initializer) {
		return initializer
	}
	return skipLuaTypeOnlyWrappers(initializer).AsBinaryExpression().Right
}

func (c *Checker) isLuaDefaultedAugmentationGuard(target *ast.Node, initializer *ast.Node) bool {
	initializer = skipLuaTypeOnlyWrappers(initializer)
	if !ast.IsBinaryExpression(initializer) {
		return false
	}
	binary := initializer.AsBinaryExpression()
	return binary.OperatorToken.Kind == ast.KindBarBarToken && c.luaSameEntityName(target, binary.Left)
}

// isLuaEnvironmentExpression recognizes `_G` and repeated self-access such as
// `_G._G`, while respecting lexical shadows of the root identifier.
func (c *Checker) isLuaEnvironmentExpression(node *ast.Node) bool {
	root := luaEntityRoot(node)
	if !ast.IsIdentifier(root) || root.Text() != c.luaGlobalsSymbol.Name ||
		c.getMergedSymbol(c.getResolvedSymbol(root)) != c.luaGlobalsSymbol {
		return false
	}
	node = skipLuaRuntimeTransparentWrappers(node)
	for ast.IsAccessExpression(node) {
		name, ok := c.getAccessedPropertyName(node)
		if !ok || name != c.luaGlobalsSymbol.Name {
			return false
		}
		node = skipLuaRuntimeTransparentWrappers(node.Expression())
	}
	return true
}

// isLuaParenthesizedEntity reports whether parentheses wrap a whole entity
// name. luaEntityNamePath keeps them significant so `X = (X) or E` is not the
// exact defaulted-guard idiom; every entity comparison must agree on that.
func isLuaParenthesizedEntity(node *ast.Node) bool {
	return ast.IsParenthesizedExpression(skipLuaTypeOnlyWrappers(node))
}

func luaEntityRoot(node *ast.Node) *ast.Node {
	node = skipLuaRuntimeTransparentWrappers(node)
	for ast.IsAccessExpression(node) {
		node = skipLuaRuntimeTransparentWrappers(node.Expression())
	}
	return node
}

// resolveLuaEntityPath canonicalizes only stable local aliases. The identity is
// for table storage, never for the local variable slots.
func (c *Checker) resolveLuaEntityPath(reference *ast.Node, rootName string, path []string) (*ast.Symbol, []string, bool) {
	root, path := c.resolveLexicalLuaPathRoot(reference, rootName, path)
	if root == nil {
		return nil, nil, false
	}
	if len(path) != 0 {
		root = c.canonicalLuaAliasSymbol(root)
	}
	return c.rerootLuaEnvironmentPath(root, path)
}

func (c *Checker) luaSameEntityName(left *ast.Node, right *ast.Node) bool {
	// A stable alias names one entity under two roots, so the storage identity
	// decides names that differ textually. It skips parentheses, which stay
	// significant here, so an entity wrapped in them never takes this path.
	if !isLuaParenthesizedEntity(left) && !isLuaParenthesizedEntity(right) &&
		c.isSameLuaStableAccess(left, right) {
		return true
	}
	leftName, leftPath, leftOK := luaEntityNamePath(left, c.getAccessedPropertyName)
	rightName, rightPath, rightOK := luaEntityNamePath(right, c.getAccessedPropertyName)
	if !leftOK || !rightOK {
		return false
	}
	if leftName == rightName && slices.Equal(leftPath, rightPath) {
		return true
	}
	// Only the environment spelling can re-root a textually different name.
	if leftName != c.luaGlobalsSymbol.Name && rightName != c.luaGlobalsSymbol.Name {
		return false
	}
	leftRoot, leftPath, leftOK := c.resolveLuaEntityPath(left, leftName, leftPath)
	rightRoot, rightPath, rightOK := c.resolveLuaEntityPath(right, rightName, rightPath)
	return leftOK && rightOK && leftRoot == rightRoot && slices.Equal(leftPath, rightPath)
}

// effectiveLuaAssignmentAugmentations applies the transaction's single winner
// rule to inference and constructor discovery as well as flow.
func (c *Checker) effectiveLuaAssignmentAugmentations(assignments []luaAugmentation) []luaAugmentation {
	effective := make([]luaAugmentation, 0, len(assignments))
	for _, assignment := range assignments {
		if c.isOverwrittenLuaCapturedTarget(assignment.Target) {
			continue
		}
		effective = append(effective, assignment)
	}
	return effective
}

// luaAssignmentDeclarationSet collects the declaration nodes an assignment group
// owns, so a symbol's remaining declarations can be told apart from them.
func luaAssignmentDeclarationSet(assignments []luaAugmentation) collections.Set[*ast.Node] {
	var declarations collections.Set[*ast.Node]
	for _, assignment := range assignments {
		declarations.Add(assignment.declaration())
	}
	return declarations
}

func hasOnlyLuaConstructorAssignmentDeclarations(symbol *ast.Symbol, assignments []luaAugmentation) bool {
	declarations := luaAssignmentDeclarationSet(assignments)
	for _, declaration := range symbol.Declarations {
		if !declarations.Has(declaration) {
			return false
		}
	}
	return true
}

package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
)

// Resolution of the table constructors a Lua storage location can hold. Arm
// resolution is a function of the program alone, never of a query point, which
// is what lets one cache per checker serve every caller.

func luaConstructorBlocksNumericEvolution(arm *ast.Symbol) bool {
	for _, declaration := range arm.Declarations {
		if !ast.IsObjectLiteralExpression(declaration) || len(declaration.Properties()) != 0 {
			continue
		}
		initializer := outermostLuaWrapper(declaration, ast.OEKParentheses|ast.OEKAssertions)
		return !isEmptyEvolvingArrayInitializer(initializer)
	}
	return false
}

// parentArms resolves one write's receiver to the constructor arms it can
// reach. The member is declared on every one of them.
func (r *luaConstructorResolver) parentArms(item luaAugmentation) ([]*ast.Symbol, bool) {
	var receiver *ast.Node
	if ast.IsFunctionDeclaration(item.Source) {
		receiver = item.Source.AsFunctionDeclaration().Target
	} else if reference := ast.GetLuaAssignmentTargetReference(item.Target); reference != nil && ast.IsAccessExpression(reference) {
		receiver = reference.Expression()
	}
	if receiver == nil {
		return nil, false
	}
	return r.initializerArms(receiver)
}

// resolveLuaConstructors returns the shared resolver. Arm resolution is a
// function of the program alone, so one cache serves every caller once
// initializeLuaAugmentations has finished attaching members.
func (c *Checker) resolveLuaConstructors() *luaConstructorResolver {
	if c.luaConstructorResolver == nil {
		c.luaConstructorResolver = newLuaConstructorResolver(c)
	}
	return c.luaConstructorResolver
}

// shouldDedupeLuaConstructorTypes reports whether a symbol's inferred union may
// contain structurally identical arms reached through different aliases.
func (c *Checker) shouldDedupeLuaConstructorTypes(symbol *ast.Symbol, assignments []luaAugmentation) bool {
	return len(assignments) > 1 && c.hasLuaConstructorArms(symbol)
}

// hasLuaConstructorArms reports whether a symbol's storage resolves to at least
// one known table constructor.
func (c *Checker) hasLuaConstructorArms(symbol *ast.Symbol) bool {
	arms, known := c.resolveLuaConstructors().armsAt(symbol)
	return known && len(arms) != 0
}

type luaConstructorArms struct {
	arms  []*ast.Symbol
	known bool
}

// luaConstructorResolver maps a symbol to the constructors its storage can
// hold. Resolution is a function of the program, never of a query point, so one
// cache per checker serves every caller. It is checker-local because attached
// augmentation members must never mutate shared binder symbols.
type luaConstructorResolver struct {
	checker   *Checker
	cache     map[*ast.Symbol]luaConstructorArms
	resolving collections.Set[*ast.Symbol]
}

func newLuaConstructorResolver(c *Checker) *luaConstructorResolver {
	return &luaConstructorResolver{
		checker: c,
		cache:   make(map[*ast.Symbol]luaConstructorArms),
	}
}

func (r *luaConstructorResolver) armsAt(symbol *ast.Symbol) ([]*ast.Symbol, bool) {
	c := r.checker
	symbol = c.getMergedSymbol(symbol)
	if symbol == nil {
		return nil, false
	}
	if resolved, ok := r.cache[symbol]; ok {
		return resolved.arms, resolved.known
	}
	if !r.resolving.AddIfAbsent(symbol) {
		return nil, false
	}
	arms, known := r.computeArms(symbol)
	r.resolving.Delete(symbol)
	if known {
		// Only a resolved answer is a function of the program alone. An unknown
		// one means a member this pass has yet to attach, so caching it would let
		// the shared resolver answer from candidate order rather than the program.
		r.cache[symbol] = luaConstructorArms{arms: arms, known: known}
	}
	return arms, known
}

func (r *luaConstructorResolver) computeArms(symbol *ast.Symbol) ([]*ast.Symbol, bool) {
	c := r.checker
	assignments := c.luaAssignmentAugmentations[symbol]
	if len(assignments) != 0 && hasOnlyLuaConstructorAssignmentDeclarations(symbol, assignments) {
		return r.applyConstructorAssignments(nil, true, assignments)
	}

	// A local that is also assigned holds its initializer's constructor plus every
	// constructor assigned to it. An assignment whose value is not a constructor
	// resolves as unknown and seals the local.
	if declaration := symbol.ValueDeclaration; len(assignments) != 0 && declaration != nil &&
		ast.IsVariableDeclaration(declaration) && ast.IsLuaLocal(declaration) {
		if declaration.Type() != nil {
			return nil, false
		}
		var initialArms []*ast.Symbol
		initialKnown := true
		if initializer := luaExplicitVariableInitializer(declaration); initializer != nil {
			initialArms, initialKnown = r.initializerArms(initializer)
		}
		return r.applyConstructorAssignments(initialArms, initialKnown, assignments)
	}

	declarations := symbol.Declarations
	if symbol.Flags&ast.SymbolFlagsAssignment == 0 && symbol.ValueDeclaration != nil {
		declarations = []*ast.Node{symbol.ValueDeclaration}
	}
	var arms []*ast.Symbol
	for _, declaration := range declarations {
		var initializer *ast.Node
		switch {
		case ast.IsVariableDeclaration(declaration):
			if declaration.Type() != nil || !ast.IsLuaLocal(declaration) && declaration.Parent.Flags&ast.NodeFlagsConst == 0 {
				return nil, false
			}
			// A local's initializer denotes its constructor only while the binding is
			// stable. A rebound local is sealed, so an old `local a = t` alias never
			// leaks a later member onto t's constructor.
			if ast.IsLuaLocal(declaration) && !c.isLuaStableIdentityBinding(symbol) {
				return nil, false
			}
			initializer = luaExplicitVariableInitializer(declaration)
		case ast.IsPropertyAssignment(declaration):
			initializer = declaration.Initializer()
		case ast.IsBinaryExpression(declaration):
			binary := declaration.AsBinaryExpression()
			initializer = c.luaDefaultedAugmentationInitializer(binary.Left, binary.Right)
		default:
			return nil, false
		}
		if initializer == nil {
			return nil, false
		}
		initializerArms, known := r.initializerArms(initializer)
		if !known {
			return nil, false
		}
		arms = appendLuaConstructorArms(c, arms, initializerArms)
	}
	return arms, true
}

// assignmentInitializerArms resolves the constructors one assignment installs.
// A defaulted guard (`X = X or {}`) contributes its default, not the self-read.
func (r *luaConstructorResolver) assignmentInitializerArms(assignment luaAugmentation) ([]*ast.Symbol, bool) {
	initializer := luaExplicitAssignmentValueAt(assignment.Source.AsBinaryExpression().Right, assignment.ValueIndex)
	if initializer == nil {
		return nil, false
	}
	return r.initializerArms(r.checker.luaDefaultedAugmentationInitializer(assignment.Target, initializer))
}

func (r *luaConstructorResolver) initializerArms(initializer *ast.Node) ([]*ast.Symbol, bool) {
	c := r.checker
	if constructor := luaObjectLiteralConstructor(initializer); constructor != nil {
		if constructor.Symbol() == nil {
			return nil, false
		}
		return []*ast.Symbol{c.getMergedSymbol(constructor.Symbol())}, true
	}
	return r.referenceArms(skipLuaRuntimeTransparentWrappers(initializer))
}

// applyConstructorAssignments unions every constructor the storage can hold.
// Assignments are declarations, not ordered stores, so a write in any file or
// branch contributes its constructor.
func (r *luaConstructorResolver) applyConstructorAssignments(initialArms []*ast.Symbol, initialKnown bool, assignments []luaAugmentation) ([]*ast.Symbol, bool) {
	c := r.checker
	if !initialKnown {
		return nil, false
	}
	arms := appendLuaConstructorArms(c, nil, initialArms)
	for _, assignment := range c.effectiveLuaAssignmentAugmentations(assignments) {
		if c.isSelfPreservingLuaCapturedTarget(assignment.Target) {
			continue
		}
		initializerArms, known := r.assignmentInitializerArms(assignment)
		if !known {
			return nil, false
		}
		arms = appendLuaConstructorArms(c, arms, initializerArms)
	}
	return arms, true
}

// referenceArms resolves a static entity path -- bare name or dotted access --
// to the constructor arms it can hold. Access paths go through referenceSymbols
// too, so an `_G.x` receiver reroots to the global named x.
func (r *luaConstructorResolver) referenceArms(reference *ast.Node) ([]*ast.Symbol, bool) {
	reference = skipLuaRuntimeTransparentWrappers(reference)
	symbols, known := r.referenceSymbols(reference)
	if !known {
		return nil, false
	}
	var arms []*ast.Symbol
	for _, symbol := range symbols {
		symbolArms, known := r.armsAt(symbol)
		if !known {
			return nil, false
		}
		arms = appendLuaConstructorArms(r.checker, arms, symbolArms)
	}
	return arms, true
}

// referenceSymbols resolves a static entity path without asking for its type,
// so snapshot inference cannot recurse through the symbol being inferred.
func (r *luaConstructorResolver) referenceSymbols(reference *ast.Node) ([]*ast.Symbol, bool) {
	c := r.checker
	rootName, path, ok := luaEntityNamePath(reference, c.getAccessedPropertyName)
	if !ok {
		return nil, false
	}
	root, path, ok := c.resolveLuaEntityPath(reference, rootName, path)
	if !ok {
		return nil, false
	}
	symbols := []*ast.Symbol{root}
	for _, name := range path {
		var next []*ast.Symbol
		for _, symbol := range symbols {
			arms, known := r.armsAt(symbol)
			if !known {
				return nil, false
			}
			for _, arm := range arms {
				arm = c.getMergedSymbol(arm)
				member := c.getMergedSymbol(core.OrElse(arm.Exports[name], arm.Members[name]))
				if member == nil {
					return nil, false
				}
				if !slices.Contains(next, member) {
					next = append(next, member)
				}
			}
		}
		symbols = next
	}
	return symbols, true
}

func appendLuaConstructorArms(c *Checker, arms []*ast.Symbol, additions []*ast.Symbol) []*ast.Symbol {
	for _, arm := range additions {
		arm = c.getMergedSymbol(arm)
		if arm != nil && !slices.Contains(arms, arm) {
			arms = append(arms, arm)
		}
	}
	return arms
}

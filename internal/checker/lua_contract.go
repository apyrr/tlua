package checker

import (
	"slices"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
)

// Composition of an `as`-asserted surface with the constructor a Lua assignment
// actually installs. An assertion states the contract callers see; the runtime
// table may still carry members the contract does not govern, and those are
// added as a separate intersection arm rather than overriding the contract.

type luaContractConstructorPair struct {
	contract    *Type
	constructor *Type
}

// luaContractOwnsProperty is deliberately existential for unions. A property
// present in only one contract arm is not safe to read through the union, but
// the constructor must not manufacture a writable declaration that bypasses
// that arm's property or index constraint.
func (c *Checker) luaContractOwnsProperty(contract *Type, name string) bool {
	if contract.flags&TypeFlagsUnion != 0 {
		return core.Some(contract.Types(), func(part *Type) bool {
			return c.luaContractOwnsProperty(part, name)
		})
	}
	return c.getPropertyOfType(contract, name) != nil ||
		c.getApplicableIndexInfoForName(contract, name) != nil
}

// Contract surfaces detach an inferred member from its expression-shaped
// assignment declaration, which has no declaration symbol for generic
// intersection synthesis to inspect. Their types are computed when first asked
// for: composing a table's type must not resolve its members, because a
// member's initializer may read a sibling member (`T.b = T.a * 2`) and a
// member function's parameters take their types from the table's contract,
// either of which reads the table's type again.

type luaNestedContractSurface struct {
	contract *Type
	source   *ast.Symbol
}

// newLuaDeferredContractSurfaceProperty is a surface for a member the contract
// does not govern; its types are its source member's.
func (c *Checker) newLuaDeferredContractSurfaceProperty(source *ast.Symbol) *ast.Symbol {
	prop := c.newLuaContractSurfacePropertySymbol(source)
	c.luaContractSurfaceSources[prop] = source
	return prop
}

// newLuaNestedContractSurfaceProperty is a surface for additions below a
// member the contract governs; its type holds the additions alone, so it
// intersects with the contract's member rather than replacing it.
func (c *Checker) newLuaNestedContractSurfaceProperty(contractProp *ast.Symbol, contract *Type, source *ast.Symbol) *ast.Symbol {
	prop := c.newLuaContractSurfacePropertySymbol(contractProp)
	c.luaNestedContractSurfaces[prop] = &luaNestedContractSurface{contract: contract, source: source}
	return prop
}

func (c *Checker) getTypeOfLuaNestedContractSurface(surface *luaNestedContractSurface) *Type {
	constructor := c.getTypeOfSymbol(surface.source)
	const objectLike = TypeFlagsObject | TypeFlagsUnionOrIntersection | TypeFlagsInstantiableNonPrimitive
	if constructor.flags&objectLike != 0 {
		if nested := c.getLuaConstructorContractExtras(surface.contract, constructor, make(map[luaContractConstructorPair]bool), false /*excludeNumeric*/); nested != nil {
			return nested
		}
	}
	return c.unknownType
}

// hasLuaMemberAdditions reports whether a member holds a table with members of
// its own, written in its constructor (`T.config = { extra = 1 }`) or added by
// writes below it (`T.config.extra = 1`), which may go beyond the contract. It
// reads the member's constructor arms, not its type.
func (c *Checker) hasLuaMemberAdditions(member *ast.Symbol) bool {
	arms, known := c.resolveLuaConstructors().armsAt(member)
	return known && core.Some(arms, func(arm *ast.Symbol) bool {
		arm = c.getMergedSymbol(arm)
		return len(arm.Members) != 0 || len(arm.Exports) != 0
	})
}

func (c *Checker) newLuaContractSurfacePropertySymbol(source *ast.Symbol) *ast.Symbol {
	flags := ast.SymbolFlagsProperty | (source.Flags & ast.SymbolFlagsOptional)
	checkFlags := ast.CheckFlagsSyntheticProperty | (source.CheckFlags & ast.CheckFlagsLate)
	if c.isReadonlySymbol(source) {
		checkFlags |= ast.CheckFlagsReadonly
	}
	prop := c.newSymbolEx(flags, source.Name, checkFlags)
	prop.Declarations = source.Declarations
	return prop
}

// getLuaConstructorContractExtras keeps only constructor members that the
// asserted contract does not already govern. In particular, a later structural
// write must not replace the contract's type or readonly modifier with the
// writable member synthesized for the runtime constructor.
func (c *Checker) getLuaConstructorContractExtras(contract *Type, constructor *Type, resolving map[luaContractConstructorPair]bool, excludeNumeric bool) *Type {
	key := luaContractConstructorPair{contract: contract, constructor: constructor}
	if resolving[key] {
		return nil
	}
	resolving[key] = true
	defer delete(resolving, key)

	var members ast.SymbolTable
	for _, constructorProp := range c.getPropertiesOfType(constructor) {
		name := constructorProp.Name
		if excludeNumeric && ast.IsNumberKeyName(name) {
			continue
		}
		contractProp := c.getPropertyOfType(contract, name)
		if contractProp == nil {
			if !c.luaContractOwnsProperty(contract, name) {
				if members == nil {
					members = make(ast.SymbolTable)
				}
				members[name] = c.newLuaDeferredContractSurfaceProperty(constructorProp)
			}
			continue
		}

		// Preserve additions below a property already described by the contract
		// without adding a second, writable copy of that property. Whether there
		// are additions is read from the member's constructor, not its type.
		const objectLike = TypeFlagsObject | TypeFlagsUnionOrIntersection | TypeFlagsInstantiableNonPrimitive
		if !c.hasLuaMemberAdditions(constructorProp) {
			continue
		}
		contractPropType := c.getTypeOfSymbol(contractProp)
		if contractPropType.flags&objectLike == 0 {
			continue
		}
		if members == nil {
			members = make(ast.SymbolTable)
		}
		members[name] = c.newLuaNestedContractSurfaceProperty(contractProp, contractPropType, constructorProp)
	}
	if len(members) == 0 {
		return nil
	}
	return c.newAnonymousType(nil, members, nil, nil, nil)
}

// A table's contract is what its declared type says about a member: the seed of
// a host environment's table, the type an `as` assertion gives a constructor,
// or a closed table's declared type. `function T:m(...) ... end` is Lua sugar
// for `T.m = function(self, ...) ... end`, so both spellings take their
// contextual type from the contract member and are checked against it. A write
// that declares a member the contract does not declare is the table's own, and
// has no contract to follow.

// isLuaMemberFunctionDeclaration reports a dotted or colon function
// declaration, which stores its function into a table member.
func isLuaMemberFunctionDeclaration(node *ast.Node) bool {
	return ast.IsFunctionDeclaration(node) && node.AsFunctionDeclaration().Target != nil
}

// luaContractAssertion returns the assertion that gives a table constructor
// its contract: the outermost `as T` around it. A const assertion states no
// contract.
func luaContractAssertion(constructor *ast.Node) *ast.Node {
	for node := outermostLuaWrapper(constructor, ast.OEKParentheses|ast.OEKAssertions); node != constructor; node = node.Expression() {
		if node.Kind == ast.KindTypeAssertionExpression || ast.IsAsExpression(node) {
			if ast.IsConstAssertion(node) {
				return nil
			}
			return node
		}
	}
	return nil
}

// isLuaContractArm reports whether a constructor carries a contract its members
// may implement: a host environment's seeded table, or a table constructor
// under an `as` assertion. It reads syntax only, so attachment may ask.
func (c *Checker) isLuaContractArm(arm *ast.Symbol) bool {
	if c.isLuaEnvironmentArm(arm) {
		return true
	}
	for _, declaration := range c.getMergedSymbol(arm).Declarations {
		if ast.IsObjectLiteralExpression(declaration) && luaContractAssertion(declaration) != nil {
			return true
		}
	}
	return false
}

// getLuaContractOfArm returns a constructor's contract type, or nil.
func (c *Checker) getLuaContractOfArm(arm *ast.Symbol) *Type {
	arm = c.getMergedSymbol(arm)
	if global := c.luaEnvironmentArmGlobals[arm]; global != nil {
		return c.getLuaEnvironmentSeedType(global)
	}
	for _, declaration := range arm.Declarations {
		if !ast.IsObjectLiteralExpression(declaration) {
			continue
		}
		if assertion := luaContractAssertion(declaration); assertion != nil {
			return c.getTypeFromTypeNode(assertion.Type())
		}
	}
	return nil
}

// getLuaMemberContract returns the contract member that a write to receiver's
// member `name` implements, or nil when there is none.
//
// A write that declares a member (`declaresMember`) of an open table reads the
// contract of the constructors the receiver holds, never the table's composed
// type: that type includes the member being written, so reading it while the
// write is typed would re-enter the write. A write to a member the table
// already has reads the receiver's type, which is declared.
func (c *Checker) getLuaMemberContract(receiver *ast.Node, name string, declaresMember bool) *ast.Symbol {
	if !declaresMember {
		receiverType := c.getTypeOfExpression(receiver)
		if c.isErrorType(receiverType) {
			return nil
		}
		return c.getPropertyOfType(receiverType, name)
	}
	arms, known := c.resolveLuaConstructors().referenceArms(receiver)
	if !known {
		return nil
	}
	var contracts []*Type
	for _, arm := range arms {
		if contract := c.getLuaContractOfArm(arm); contract != nil {
			contracts = append(contracts, contract)
		}
	}
	if len(contracts) == 0 {
		return nil
	}
	return c.getPropertyOfType(c.getUnionType(contracts), name)
}

// getLuaConstructorTypeForAssertion is the table an assertion converts: the
// constructor with the members later writes attached to it, since a Lua table
// is built in steps. An attached member the asserted contract declares is
// checked against it on its own, so here it stands as the contract's member:
// comparing its body again would report a mismatch twice, or report a colon
// body whose self carries the table's other members.
func (c *Checker) getLuaConstructorTypeForAssertion(expression *ast.Node, t *Type, contract *Type) *Type {
	constructor := ast.SkipParentheses(expression)
	if !ast.IsObjectLiteralExpression(constructor) || t.flags&TypeFlagsObject == 0 {
		return t
	}
	literal := c.getMergedSymbol(constructor.Symbol())
	if literal == nil || len(literal.Exports) == 0 {
		return t
	}
	members := make(ast.SymbolTable)
	for _, property := range c.getPropertiesOfType(t) {
		if literal.Exports[property.Name] != nil {
			if contractProperty := c.getPropertyOfType(contract, property.Name); contractProperty != nil {
				members[property.Name] = contractProperty
				continue
			}
		}
		members[property.Name] = property
	}
	// Keep the literal's flags, so widening treats the result as the object
	// literal it is (an empty `{}` member still widens).
	result := c.newAnonymousType(t.symbol, members, nil, nil, c.getIndexInfosOfType(t))
	result.objectFlags |= t.objectFlags
	return result
}

// getLuaMemberFunctionContract returns the contract member a dotted or colon
// function declaration implements.
func (c *Checker) getLuaMemberFunctionContract(fn *ast.Node) *ast.Symbol {
	name := fn.Name()
	if name == nil {
		return nil
	}
	declaresMember := c.getMergedSymbol(c.getSymbolOfDeclaration(fn)).Flags&ast.SymbolFlagsAssignment != 0
	return c.getLuaMemberContract(fn.AsFunctionDeclaration().Target, ast.GetPropertyNameForPropertyNameNode(name), declaresMember)
}

// getContextualTypeForLuaMemberFunction is the contextual type of a dotted or
// colon function declaration: the contract member it implements, exactly as
// the member is the contextual type of the value in `T.m = function ... end`.
func (c *Checker) getContextualTypeForLuaMemberFunction(fn *ast.Node) *Type {
	contract := c.getLuaMemberFunctionContract(fn)
	if contract == nil {
		return nil
	}
	return c.getWriteTypeOfSymbol(contract)
}

// getContextualTypeForLuaMemberDeclaration is the contextual type of the value
// in `T.m = value` when the write declares member m: the contract member m, if
// T's contract declares one.
func (c *Checker) getContextualTypeForLuaMemberDeclaration(left *ast.Node) *Type {
	name, ok := c.getAccessedPropertyName(left)
	if !ok {
		return nil
	}
	contract := c.getLuaMemberContract(left.Expression(), name, true /*declaresMember*/)
	if contract == nil {
		return nil
	}
	return c.getWriteTypeOfSymbol(contract)
}

// getLuaReceiverCheckSignature returns the signature a colon body is checked
// against its contract with. The body's implicit self is the table it is
// declared on, which carries members the contract does not know, so it takes
// the contract's receiver type for the check, as a TypeScript method's
// implicit `this` is never compared. An explicit self keeps its annotation.
func (c *Checker) getLuaReceiverCheckSignature(fn *ast.Node, contract *Type) *Signature {
	signature := c.getSignatureFromDeclaration(fn)
	parameters := fn.Parameters()
	if len(parameters) == 0 || parameters[0].Flags&ast.NodeFlagsReparsed == 0 || len(signature.parameters) == 0 {
		return signature
	}
	contractSignatures := c.getSignaturesOfType(c.GetNonNullableType(contract), SignatureKindCall)
	// A contract without a receiver has nothing to substitute; the colon body
	// then fails the check with its own self.
	if len(contractSignatures) != 1 || len(contractSignatures[0].parameters) == 0 ||
		contractSignatures[0].parameters[0].Name != ast.InternalSymbolNameSelf {
		return signature
	}
	receiverType := c.getTypeAtPosition(contractSignatures[0], 0)
	// cloneSignature leaves the return type unresolved; the copy must share the
	// body's, not infer its own.
	result := c.cloneSignature(signature)
	result.resolvedReturnType = c.getReturnTypeOfSignature(signature)
	result.resolvedTypePredicate = c.getTypePredicateOfSignature(signature)
	result.parameters = slices.Clone(signature.parameters)
	result.parameters[0] = c.createSymbolWithType(signature.parameters[0], receiverType)
	return result
}

// finalizeLuaAugmentationInitializerType composes the expression's asserted
// contract with the checker-local constructor shape augmented later.
func (c *Checker) finalizeLuaAugmentationInitializerType(assignment luaAugmentation, initializer *ast.Node, t *Type) *Type {
	t = c.finalizeLuaConstructorInitializerType(initializer, t)
	if initializer != nil && isEmptyEvolvingArrayInitializer(initializer) {
		return c.checkLuaAugmentationEmptyArrayType(assignment, t)
	}
	return t
}

// composeLuaLocalConstructorContract gives a local the composition an assignment
// declaration already gets. A local's type comes from its initializer rather
// than from luaAssignmentAugmentations, so without this it accepts the
// augmenting write and then cannot read the member back. An annotation still
// seals: the type is then the annotation, not the constructor.
func (c *Checker) composeLuaLocalConstructorContract(declaration *ast.Node, t *Type) *Type {
	if t == nil || !ast.IsVariableDeclaration(declaration) || !ast.IsLuaLocal(declaration) ||
		declaration.Type() != nil {
		return t
	}
	initializer := ast.LuaExplicitVariableInitializer(declaration)
	if initializer == nil {
		return t
	}
	return c.finalizeLuaConstructorInitializerType(initializer, t)
}

func (c *Checker) finalizeLuaConstructorInitializerType(initializer *ast.Node, t *Type) *Type {
	if constructor := luaObjectLiteralConstructor(initializer); constructor != nil && hasLuaTypeAssertionWrapper(initializer) {
		constructorType := c.checkExpressionForMutableLocation(constructor, CheckModeNormal)
		excludeNumeric := len(constructor.Properties()) == 0 && !isEmptyEvolvingArrayInitializer(initializer)
		if extras := c.getLuaConstructorContractExtras(
			t,
			constructorType,
			make(map[luaContractConstructorPair]bool),
			excludeNumeric,
		); extras != nil {
			t = c.getIntersectionType([]*Type{t, extras})
		}
	}
	return t
}

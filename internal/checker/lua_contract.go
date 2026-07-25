package checker

import (
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

// newLuaContractSurfaceProperty detaches an inferred surface from its
// expression-shaped assignment declaration, which has no declaration symbol
// for generic intersection synthesis to inspect.
func (c *Checker) newLuaContractSurfaceProperty(source *ast.Symbol, readType *Type, writeType *Type) *ast.Symbol {
	flags := ast.SymbolFlagsProperty | (source.Flags & ast.SymbolFlagsOptional)
	checkFlags := ast.CheckFlagsSyntheticProperty | (source.CheckFlags & ast.CheckFlagsLate)
	if c.isReadonlySymbol(source) {
		checkFlags |= ast.CheckFlagsReadonly
	}
	prop := c.newSymbolEx(flags, source.Name, checkFlags)
	prop.Declarations = source.Declarations
	links := c.valueSymbolLinks.Get(prop)
	links.resolvedType = readType
	links.writeType = writeType
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
				members[name] = c.newLuaContractSurfaceProperty(
					constructorProp,
					c.getTypeOfSymbol(constructorProp),
					c.getWriteTypeOfSymbol(constructorProp),
				)
			}
			continue
		}

		// Preserve additions below a property already described by the contract
		// without adding a second, writable copy of that property.
		contractPropType := c.getTypeOfSymbol(contractProp)
		constructorPropType := c.getTypeOfSymbol(constructorProp)
		const objectLike = TypeFlagsObject | TypeFlagsUnionOrIntersection | TypeFlagsInstantiableNonPrimitive
		if contractPropType.flags&objectLike == 0 || constructorPropType.flags&objectLike == 0 {
			continue
		}
		nested := c.getLuaConstructorContractExtras(
			contractPropType,
			constructorPropType,
			resolving,
			false, /*excludeNumeric*/
		)
		if nested == nil {
			continue
		}
		if members == nil {
			members = make(ast.SymbolTable)
		}
		members[name] = c.newLuaContractSurfaceProperty(contractProp, nested, nested)
	}
	if len(members) == 0 {
		return nil
	}
	return c.newAnonymousType(nil, members, nil, nil, nil)
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
	initializer := luaExplicitVariableInitializer(declaration)
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

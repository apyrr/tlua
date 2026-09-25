package checker

import (
	"maps"
	"slices"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
)

// Host environments (the luaEnvironments option). A Lua host may install
// globals before it runs a chunk -- load(chunk, name, mode, env) -- and hosts
// commonly install a fresh table per group of files: every file of one plugin
// folder sees the same table, a sibling folder sees another. Each group's
// global is an open table seeded with a declared type, so writes in the group's
// files declare members on that group's table and nowhere else.

type luaEnvironmentGlobal struct {
	// seedName names the global type the host's table starts as. Only the name
	// is recorded: attachment runs before types exist and must not resolve any.
	seedName string
	// arm is the node-less constructor the group's writes attach members to.
	arm *ast.Symbol
}

func (c *Checker) isLuaEnvironmentArm(arm *ast.Symbol) bool {
	return c.luaEnvironmentArmGlobals[c.getMergedSymbol(arm)] != nil
}

func (c *Checker) getLuaFileEnvironment(file *ast.SourceFile) ast.SymbolTable {
	return c.luaFileEnvironments[file]
}

// initializeLuaEnvironments creates each group's globals. The program assigns
// files to groups once; the symbols are checker-local.
func (c *Checker) initializeLuaEnvironments() {
	environments := c.compilerOptions.LuaEnvironments
	fileGroups := c.program.GetLuaEnvironmentGroups()
	if len(environments) == 0 || len(fileGroups) == 0 {
		return
	}
	tables := make(map[core.LuaEnvironmentGroup]ast.SymbolTable)
	for _, file := range c.files {
		group, ok := fileGroups[file.Path()]
		if !ok {
			continue
		}
		table := tables[group]
		if table == nil {
			table = make(ast.SymbolTable)
			globals := environments[group.Entry].Globals
			// Sorted, so symbol creation order is a function of the options alone.
			for _, name := range slices.Sorted(maps.Keys(globals)) {
				// A host-installed global is assignment-declared, like an implicit
				// global, so statement setmetatable pairs with it.
				global := c.newSymbol(ast.SymbolFlagsFunctionScopedVariable|ast.SymbolFlagsAssignment, name)
				arm := c.newSymbol(ast.SymbolFlagsObjectLiteral, ast.InternalSymbolNameObject)
				environmentGlobal := &luaEnvironmentGlobal{seedName: globals[name], arm: arm}
				c.luaEnvironmentGlobals[global] = environmentGlobal
				c.luaEnvironmentArmGlobals[arm] = environmentGlobal
				table[name] = global
			}
			tables[group] = table
		}
		c.luaFileEnvironments[file] = table
	}
	c.initializeLuaEnvironmentRegistries(environments, tables)
}

// initializeLuaEnvironmentRegistries makes each group a member of its entry's
// registry interface, named by the group's last path segment: a host that
// registers classes by folder name (`scripted_ents.Register(ENT, "door")`) can
// then declare its lookups as `R[K]`. The members merge into a checker-local
// clone of the declared interface, so the shared binder symbol stays frozen.
// A name the interface already declares keeps its declaration. Option parsing
// guarantees a registry entry has exactly one global.
func (c *Checker) initializeLuaEnvironmentRegistries(environments []*core.LuaEnvironment, tables map[core.LuaEnvironmentGroup]ast.SymbolTable) {
	groups := slices.SortedFunc(maps.Keys(tables), func(left core.LuaEnvironmentGroup, right core.LuaEnvironmentGroup) int {
		if left.Entry != right.Entry {
			return left.Entry - right.Entry
		}
		return strings.Compare(left.Path, right.Path)
	})
	members := make(map[string]ast.SymbolTable)
	for _, group := range groups {
		environment := environments[group.Entry]
		if environment.Registry == "" {
			continue
		}
		var global *ast.Symbol
		for _, symbol := range tables[group] {
			global = symbol
		}
		table := members[environment.Registry]
		if table == nil {
			table = make(ast.SymbolTable)
			members[environment.Registry] = table
		}
		className := group.Name()
		member := table[className]
		if member == nil {
			member = c.newSymbol(ast.SymbolFlagsProperty, className)
			table[className] = member
		}
		// Two groups registering one name (two addons shipping the same class)
		// are both possible at run time, so the member is their union.
		c.luaEnvironmentRegistryMembers[member] = append(c.luaEnvironmentRegistryMembers[member], global)
	}
	for _, registryName := range slices.Sorted(maps.Keys(members)) {
		existing := c.globals[registryName]
		if existing == nil || existing.Flags&ast.SymbolFlagsInterface == 0 {
			// checkLuaEnvironmentTypes reports the missing interface.
			continue
		}
		registry := c.newSymbol(ast.SymbolFlagsInterface, registryName)
		for name, member := range members[registryName] {
			if existing.Members[name] == nil {
				ast.GetMembers(registry)[name] = member
			}
		}
		c.mergeGlobalSymbol(registry)
		merged := c.globals[registryName]
		for _, member := range registry.Members {
			member.Parent = merged
		}
	}
}

// checkLuaEnvironmentTypes reports seeds and registries that name no global
// type. It runs once global types exist, whether or not any file reads the
// globals, so a misnamed type is reported even before a group is used.
func (c *Checker) checkLuaEnvironmentTypes() {
	for _, environment := range c.compilerOptions.LuaEnvironments {
		for _, name := range slices.Sorted(maps.Keys(environment.Globals)) {
			if seedName := environment.Globals[name]; c.getGlobalSymbol(seedName, ast.SymbolFlagsType, nil) == nil {
				c.addDiagnostic(ast.NewCompilerDiagnostic(diagnostics.The_luaEnvironments_entry_for_0_names_1_which_is_not_a_global_type, environment.Root, seedName))
			}
		}
		if environment.Registry != "" {
			if registry := c.globals[environment.Registry]; registry == nil || registry.Flags&ast.SymbolFlagsInterface == 0 {
				c.addDiagnostic(ast.NewCompilerDiagnostic(diagnostics.The_luaEnvironments_entry_for_0_names_1_which_is_not_a_global_type, environment.Root, environment.Registry))
			}
		}
	}
}

// getTypeOfLuaEnvironmentRegistryMember is the table of every group that
// registered the member's name.
func (c *Checker) getTypeOfLuaEnvironmentRegistryMember(globals []*ast.Symbol) *Type {
	types := make([]*Type, 0, len(globals))
	for _, global := range globals {
		types = append(types, c.getTypeOfSymbol(global))
	}
	return c.getUnionType(types)
}

// getLuaEnvironmentSeedType resolves the seed contract, or nil when no global
// type has the configured name.
func (c *Checker) getLuaEnvironmentSeedType(global *luaEnvironmentGlobal) *Type {
	seedSymbol := c.getGlobalSymbol(global.seedName, ast.SymbolFlagsType, nil)
	if seedSymbol == nil {
		return nil
	}
	return c.getDeclaredTypeOfSymbol(seedSymbol)
}

// getTypeOfLuaEnvironmentGlobal composes the seed contract with the members the
// group's writes declared, exactly as `{} as T` composes an asserted contract
// with its constructor: members the seed already governs keep the seed's type.
func (c *Checker) getTypeOfLuaEnvironmentGlobal(symbol *ast.Symbol, global *luaEnvironmentGlobal) *Type {
	seed := c.getLuaEnvironmentSeedType(global)
	if seed == nil {
		// checkLuaEnvironmentTypes reports the missing seed.
		return c.errorType
	}
	constructor := c.newAnonymousType(global.arm, global.arm.Exports, nil, nil, nil)
	extras := c.getLuaConstructorContractExtras(seed, constructor, make(map[luaContractConstructorPair]bool), false /*excludeNumeric*/)
	if extras == nil {
		return seed
	}
	return c.getIntersectionType([]*Type{seed, extras})
}

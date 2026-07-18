package checker

import (
	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/tspath"
)

// `require("a.b.c")` loads a tlua module by its Lua dotted name: the chunk's
// top-level return value is what it yields (bindTopLevelReturn declares that
// as the file's export). A dotted name that resolves to no project file keeps
// the declared overloads in require.d.tlua -- a LuaJIT lib ("bit", "ffi"), a
// runtime-provided package, a C module -- whose last one returns any. A
// path-shaped specifier (./x, a/b) is not a Lua module name and can never
// work at runtime, so it reports Cannot-find-module instead of silently
// typing as any.
func (c *Checker) checkLuaRequireCall(node *ast.Node) *Type {
	if !ast.IsRequireCall(node, true /*requireStringLiteralLikeArgument*/) {
		return nil
	}
	if !c.isLuaRequireReference(node.Expression()) {
		return nil
	}
	specifier := node.Arguments()[0]
	dotted := tspath.IsLuaModuleName(specifier.Text())
	moduleSymbol := c.resolveExternalModuleName(specifier, specifier, dotted /*ignoreErrors*/)
	if moduleSymbol == nil {
		if dotted {
			// Unresolved dotted names fall back to the ambient overloads.
			return nil
		}
		// resolveExternalModuleName reported why.
		return c.anyType
	}
	c.checkLuaModuleCanonicalName(specifier, moduleSymbol)
	// A chunk that returns nothing still loads: Lua stores `true` for it in
	// package.loaded, and that is what require reports.
	if moduleSymbol.Exports[ast.InternalSymbolNameExportEquals] == nil {
		return c.trueType
	}
	resolved := c.resolveExternalModuleSymbol(moduleSymbol, false /*dontResolveAlias*/)
	if resolved == nil {
		return c.anyType
	}
	return c.getTypeOfSymbol(resolved)
}

// isLuaRequireReference reports whether callee is the global require, either
// directly or through the standard Lua caching idiom `local require =
// require` (a local whose initializer chain reaches the global). A local
// bound to anything else — `local require = mock` — is an ordinary call.
func (c *Checker) isLuaRequireReference(callee *ast.Node) bool {
	global := c.getLuaRequireGlobalSymbol()
	if global == nil || !ast.IsIdentifier(callee) {
		return false
	}
	symbol := c.getMergedSymbol(c.getResolvedSymbol(callee))
	for range 10 { // alias chains are short; the bound guards cycles
		if symbol == global {
			return true
		}
		decl := symbol.ValueDeclaration
		if decl == nil || !ast.IsVariableDeclaration(decl) {
			return false
		}
		initializer := decl.Initializer()
		if initializer == nil || !ast.IsIdentifier(initializer) {
			return false
		}
		next := c.getMergedSymbol(c.getResolvedSymbol(initializer))
		if next == symbol {
			return false
		}
		symbol = next
	}
	return false
}

// checkLuaModuleCanonicalName enforces one name per module: package.loaded is
// keyed by the require() string, so the same file reached under two names
// would load twice at runtime. The canonical name is derived from the whole
// program's resolution table, so every checker in the pool computes the same
// answer regardless of which files it owns or their check order.
func (c *Checker) checkLuaModuleCanonicalName(specifier *ast.Node, moduleSymbol *ast.Symbol) {
	moduleFile := ast.GetSourceFileOfModule(moduleSymbol)
	if moduleFile == nil {
		return
	}
	name := specifier.Text()
	if canonical := c.getLuaCanonicalModuleName(moduleFile.Path()); canonical != "" && canonical != name {
		c.error(specifier, diagnostics.Module_name_0_resolves_to_the_same_file_as_1_a_module_has_one_canonical_name, name, canonical)
	}
}

// getLuaCanonicalModuleName returns the canonical require() name for a
// resolved module file: the shortest (then lexicographically first) name any
// file in the program resolves it under. Derived once per checker from the
// program's resolution table.
func (c *Checker) getLuaCanonicalModuleName(path tspath.Path) string {
	c.luaCanonicalNamesOnce.Do(func() {
		names := make(map[tspath.Path]string)
		for _, cache := range c.program.GetResolvedModules() {
			for key, resolved := range cache {
				if resolved == nil || !resolved.IsResolved() || !tspath.IsLuaModuleName(key.Name) {
					continue
				}
				resolvedPath := tspath.ToPath(resolved.ResolvedFileName, c.program.GetCurrentDirectory(), c.program.UseCaseSensitiveFileNames())
				if current, ok := names[resolvedPath]; !ok || luaModuleNameLess(key.Name, current) {
					names[resolvedPath] = key.Name
				}
			}
		}
		c.luaCanonicalNames = names
	})
	return c.luaCanonicalNames[path]
}

// luaModuleNameLess orders candidate canonical names: shorter first ("pkg"
// beats "pkg.init"), then lexicographic.
func luaModuleNameLess(a, b string) bool {
	if len(a) != len(b) {
		return len(a) < len(b)
	}
	return a < b
}

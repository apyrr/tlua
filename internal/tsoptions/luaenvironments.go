package tsoptions

import (
	"maps"
	"path"
	"strings"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/diagnostics"
	"github.com/apyrr/tlua/internal/tspath"
)

// normalizeLuaEnvironmentRoot returns a luaEnvironments root in the form the
// checker matches, or false when it is not a glob relative to the Lua search
// root. A leading `./` is accepted and dropped.
func normalizeLuaEnvironmentRoot(root string) (string, bool) {
	root = tspath.NormalizeSlashes(strings.TrimSpace(root))
	for strings.HasPrefix(root, "./") {
		root = root[2:]
	}
	root = strings.TrimSuffix(root, "/")
	if root == "" || tspath.IsRootedDiskPath(root) {
		return "", false
	}
	for segment := range strings.SplitSeq(root, "/") {
		if segment == "" || segment == "." || segment == ".." {
			return "", false
		}
		if _, err := path.Match(segment, ""); err != nil { //nolint:forbidigo // glob syntax check, not path handling
			return "", false
		}
	}
	return root, true
}

// validateLuaEnvironment checks one luaEnvironments entry as tluaconfig parsing
// produced it, reporting each problem at the key that has it. An entry with a
// problem is dropped, so the diagnostics are the only trace it leaves.
func validateLuaEnvironment(value any, valueExpression *ast.Expression, sourceFile *ast.SourceFile) []*ast.Diagnostic {
	entries, ok := value.(*collections.OrderedMap[string, any])
	if !ok {
		return nil
	}
	var errors []*ast.Diagnostic
	report := func(key string, message *diagnostics.Message, args ...any) {
		errors = append(errors, CreateDiagnosticForNodeInSourceFileOrCompilerDiagnostic(sourceFile, luaEnvironmentKeyNode(valueExpression, key), message, args...))
	}
	for key := range entries.Keys() {
		if key != "root" && key != "globals" && key != "registry" {
			report(key, diagnostics.Unknown_key_0_in_a_luaEnvironments_entry_Expected_root_globals_or_registry, key)
		}
	}
	rootValue, _ := entries.Get("root")
	root, _ := rootValue.(string)
	if _, validRoot := normalizeLuaEnvironmentRoot(root); !validRoot {
		report("root", diagnostics.X_root_must_be_a_glob_relative_to_the_Lua_search_root)
	}
	globalsValue, _ := entries.Get("globals")
	globals, ok := globalsValue.(*collections.OrderedMap[string, any])
	validGlobals := ok && globals.Size() != 0
	if ok {
		for name, typeName := range globals.Entries() {
			if typeName, isString := typeName.(string); name == "" || !isString || typeName == "" {
				validGlobals = false
			}
		}
	}
	if !validGlobals {
		report("globals", diagnostics.X_globals_must_map_at_least_one_global_name_to_the_name_of_a_global_type)
	}
	if registryValue, exists := entries.Get("registry"); exists {
		if registry, isString := registryValue.(string); !isString || registry == "" {
			report("registry", diagnostics.X_registry_must_be_the_name_of_a_global_interface)
		} else if validGlobals && globals.Size() != 1 {
			// The registry maps a name to one table; with several globals there
			// is no single table to name.
			report("registry", diagnostics.The_luaEnvironments_entry_for_0_has_a_registry_so_it_must_declare_exactly_one_global, root)
		}
	}
	return errors
}

// luaEnvironmentKeyNode returns the property of an entry's object literal that
// holds key, or the entry itself when the key is absent.
func luaEnvironmentKeyNode(valueExpression *ast.Expression, key string) *ast.Node {
	if valueExpression == nil || !ast.IsObjectLiteralExpression(valueExpression) {
		return valueExpression
	}
	for _, property := range valueExpression.Properties() {
		if ast.IsPropertyAssignment(property) && property.Name() != nil && property.Name().Text() == key {
			return property
		}
	}
	return valueExpression
}

// parseLuaEnvironments reads the option as tluaconfig parsing produces it
// (ordered maps), as generic decoding produces it (buildinfo), or already
// typed. Config parsing reports invalid entries; any that reach here from
// another source are dropped.
func parseLuaEnvironments(value any) []*core.LuaEnvironment {
	if environments, ok := value.([]*core.LuaEnvironment); ok {
		return environments
	}
	list, ok := value.([]any)
	if !ok {
		return nil
	}
	var result []*core.LuaEnvironment
	for _, item := range list {
		entries := jsonObjectEntries(item)
		root, _ := entries["root"].(string)
		root, validRoot := normalizeLuaEnvironmentRoot(root)
		registry, _ := entries["registry"].(string)
		globals := make(map[string]string)
		validGlobals := true
		for name, typeName := range jsonObjectEntries(entries["globals"]) {
			typeName, isString := typeName.(string)
			if name == "" || !isString || typeName == "" {
				validGlobals = false
				break
			}
			globals[name] = typeName
		}
		if !validRoot || !validGlobals || len(globals) == 0 || registry != "" && len(globals) != 1 {
			continue
		}
		result = append(result, &core.LuaEnvironment{Root: root, Globals: globals, Registry: registry})
	}
	return result
}

// jsonObjectEntries reads a JSON object as tluaconfig parsing produces it
// (an ordered map) or as generic decoding produces it (a Go map, as when
// options are read back from buildinfo).
func jsonObjectEntries(json any) map[string]any {
	switch object := json.(type) {
	case *collections.OrderedMap[string, any]:
		return maps.Collect(object.Entries())
	case map[string]any:
		return object
	}
	return nil
}

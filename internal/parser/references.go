package parser

import (
	"github.com/apyrr/tlua/internal/ast"
)

func collectExternalModuleReferences(file *ast.SourceFile) {
	for _, node := range file.Statements.Nodes {
		collectModuleReferences(file, node)
	}

	// `require("./m")` is Lua's module load, so every file's require calls
	// contribute to the program graph. The walk is guarded by a text scan for
	// "require"/"import", so files that mention neither pay nothing.
	ast.ForEachDynamicImportOrRequireCall(file /*includeTypeSpaceImports*/, true /*requireStringLiteralLikeArgument*/, true, func(node *ast.Node, moduleSpecifier *ast.Expression) bool {
		ast.SetImportsOfSourceFile(file, append(file.Imports(), moduleSpecifier))
		return false
	})
}

func collectModuleReferences(file *ast.SourceFile, node *ast.Statement) {
	// Only `declare global` augmentations survive (string-named ambient modules
	// are deleted); they merge into the global scope from external modules.
	if ast.IsModuleDeclaration(node) && ast.IsAmbientModule(node) &&
		(ast.HasSyntacticModifier(node, ast.ModifierFlagsAmbient) || file.IsDeclarationFile) &&
		ast.IsExternalModule(file) {
		file.ModuleAugmentations = append(file.ModuleAugmentations, node.AsModuleDeclaration().Name())
	}
}

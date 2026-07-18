package ast

import (
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/tspath"
)

type SourceFileParseOptions struct {
	FileName                       string
	Path                           tspath.Path
	ExternalModuleIndicatorOptions ExternalModuleIndicatorOptions
}

type ExternalModuleIndicatorOptions struct {
	JSX   bool
	Force bool
}

func GetExternalModuleIndicatorOptions(fileName string, options *core.CompilerOptions, metadata SourceFileMetaData) ExternalModuleIndicatorOptions {
	// Every non-declaration file is a module (see getExternalModuleIndicator),
	// so the options no longer influence module detection.
	return ExternalModuleIndicatorOptions{}
}

func SetExternalModuleIndicator(file *SourceFile, opts ExternalModuleIndicatorOptions) {
	file.ExternalModuleIndicator = getExternalModuleIndicator(file, opts)
}

// getExternalModuleIndicator reports the node that makes file a module. In Lua
// every file is a chunk that `require` loads on its own, so every source file
// is a module and the file itself is the indicator. Declaration files stay
// global scripts: they describe ambient declarations that merge into the global
// scope, which is what the lualibs rely on.
func getExternalModuleIndicator(file *SourceFile, opts ExternalModuleIndicatorOptions) *Node {
	if file.ScriptKind == core.ScriptKindJSON || file.IsDeclarationFile {
		return nil
	}
	return file.AsNode()
}

package tsoptions

import (
	"slices"

	"github.com/apyrr/tlua/internal/collections"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/tspath"
)

var LibMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	// Lua runtimes. Add new runtime variants here as they are supported.
	{Key: "luajit", Value: "lib.luajit.d.tlua"},
})

var (
	Libs        = slices.Collect(LibMap.Keys())
	LibFilesSet = collections.NewSetFromItems(core.Map(slices.Collect(LibMap.Values()), func(s any) string { return s.(string) })...)
)

func GetLibFileName(libName string) (string, bool) {
	// checks if the libName is a valid lib name or file name and converts the lib name to the filename if needed
	libName = tspath.ToFileNameLowerCase(libName)
	if LibFilesSet.Has(libName) {
		return libName, true
	}
	lib, ok := LibMap.Get(libName)
	if !ok {
		return "", false
	}
	return lib.(string), true
}

var targetOptionMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "es5", Value: core.ScriptTargetES5},
	{Key: "es6", Value: core.ScriptTargetES2015},
	{Key: "es2015", Value: core.ScriptTargetES2015},
	{Key: "es2016", Value: core.ScriptTargetES2016},
	{Key: "es2017", Value: core.ScriptTargetES2017},
	{Key: "es2018", Value: core.ScriptTargetES2018},
	{Key: "es2019", Value: core.ScriptTargetES2019},
	{Key: "es2020", Value: core.ScriptTargetES2020},
	{Key: "es2021", Value: core.ScriptTargetES2021},
	{Key: "es2022", Value: core.ScriptTargetES2022},
	{Key: "es2023", Value: core.ScriptTargetES2023},
	{Key: "es2024", Value: core.ScriptTargetES2024},
	{Key: "es2025", Value: core.ScriptTargetES2025},
	{Key: "esnext", Value: core.ScriptTargetESNext},
})

var moduleOptionMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "commonjs", Value: core.ModuleKindCommonJS},
	{Key: "amd", Value: core.ModuleKindAMD},
	{Key: "system", Value: core.ModuleKindSystem},
	{Key: "umd", Value: core.ModuleKindUMD},
	{Key: "es6", Value: core.ModuleKindES2015},
	{Key: "es2015", Value: core.ModuleKindES2015},
	{Key: "es2020", Value: core.ModuleKindES2020},
	{Key: "es2022", Value: core.ModuleKindES2022},
	{Key: "esnext", Value: core.ModuleKindESNext},
	{Key: "node16", Value: core.ModuleKindNode16},
	{Key: "node18", Value: core.ModuleKindNode18},
	{Key: "node20", Value: core.ModuleKindNode20},
	{Key: "nodenext", Value: core.ModuleKindNodeNext},
	{Key: "preserve", Value: core.ModuleKindPreserve},
})

var jsxOptionMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "preserve", Value: core.JsxEmitPreserve},
	{Key: "react-native", Value: core.JsxEmitReactNative},
	{Key: "react-jsx", Value: core.JsxEmitReactJSX},
	{Key: "react-jsxdev", Value: core.JsxEmitReactJSXDev},
	{Key: "react", Value: core.JsxEmitReact},
})

var newLineOptionMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "crlf", Value: core.NewLineKindCRLF},
	{Key: "lf", Value: core.NewLineKindLF},
})

// The LuaJIT lib is the default for every target: tlua has exactly one
// runtime library surface (per Lua runtime variant), not a per-ES-target
// lattice.
func GetDefaultLibFileName(options *core.CompilerOptions) string {
	return "lib.luajit.d.tlua"
}

var watchFileEnumMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "fixedpollinginterval", Value: core.WatchFileKindFixedPollingInterval},
	{Key: "prioritypollinginterval", Value: core.WatchFileKindPriorityPollingInterval},
	{Key: "dynamicprioritypolling", Value: core.WatchFileKindDynamicPriorityPolling},
	{Key: "fixedchunksizepolling", Value: core.WatchFileKindFixedChunkSizePolling},
	{Key: "usefsevents", Value: core.WatchFileKindUseFsEvents},
	{Key: "usefseventsonparentdirectory", Value: core.WatchFileKindUseFsEventsOnParentDirectory},
})

var watchDirectoryEnumMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "usefsevents", Value: core.WatchDirectoryKindUseFsEvents},
	{Key: "fixedpollinginterval", Value: core.WatchDirectoryKindFixedPollingInterval},
	{Key: "dynamicprioritypolling", Value: core.WatchDirectoryKindDynamicPriorityPolling},
	{Key: "fixedchunksizepolling", Value: core.WatchDirectoryKindFixedChunkSizePolling},
})

var fallbackEnumMap = collections.NewOrderedMapFromList([]collections.MapEntry[string, any]{
	{Key: "fixedinterval", Value: core.PollingKindFixedInterval},
	{Key: "priorityinterval", Value: core.PollingKindPriorityInterval},
	{Key: "dynamicpriority", Value: core.PollingKindDynamicPriority},
	{Key: "fixedchunksize", Value: core.PollingKindFixedChunkSize},
})

package estransforms

import (
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/transformers"
)

var (
	// Classes, decorators, and (with the whole let/const/var/using surface)
	// `using` declarations are removed in tlua, so ESNext has no downlevel:
	// class fields, class static blocks, decorator, and using lowering are gone.
	NewESNextTransformer transformers.TransformerFactory = nil
	// 2025: only module system syntax (import attributes, json modules), untransformed regex modifiers
	// 2024: no new downlevel syntax
	// 2023: no new downlevel syntax
	NewES2021Transformer = transformers.Chain(newLogicalAssignmentTransformer)
	// Optional chaining is lowered to Lua by NewLuaOptionalChainTransformer, which
	// runs unconditionally in the emit pipeline (Lua is always the target), so it
	// is intentionally absent from this target-gated JS downlevel chain.
	NewES2020Transformer = transformers.Chain(NewES2021Transformer, newNullishCoalescingTransformer)
	NewES2019Transformer = transformers.Chain(NewES2020Transformer, newOptionalCatchTransformer)
	// The object rest/spread downlevel transform is removed in tlua: spread
	// elements, spread assignments, and binding-pattern rests no longer parse,
	// so there is nothing to lower.
	NewES2018Transformer = transformers.Chain(NewES2019Transformer, newTaggedTemplateLiftRestrictionTransformer)
	// The async downlevel transform (__awaiter + generator state machine) is
	// removed in tlua: suspension is the Lua coroutine library's job. The
	// throwaway JS emit therefore keeps the native `async` keyword at ES2017+
	// instead of lowering it; Lua — the real target — emits a plain function.
	NewES2017Transformer = NewES2018Transformer
	NewES2016Transformer = transformers.Chain(NewES2017Transformer, newExponentiationTransformer)
)

func GetESTransformer(opts *transformers.TransformOptions) *transformers.Transformer {
	options := opts.CompilerOptions
	switch options.GetEmitScriptTarget() {
	case core.ScriptTargetESNext:
		return nil
	case core.ScriptTargetES2025, core.ScriptTargetES2024, core.ScriptTargetES2023, core.ScriptTargetES2022, core.ScriptTargetES2021:
		return nil
	case core.ScriptTargetES2020:
		return NewES2021Transformer(opts)
	case core.ScriptTargetES2019:
		return NewES2020Transformer(opts)
	case core.ScriptTargetES2018:
		return NewES2019Transformer(opts)
	case core.ScriptTargetES2017:
		return NewES2018Transformer(opts)
	case core.ScriptTargetES2016:
		return NewES2017Transformer(opts)
	default: // other, older, option, transform maximally
		return NewES2016Transformer(opts)
	}
}

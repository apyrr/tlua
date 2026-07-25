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
	// ES2025 through ES2019 add no downlevel syntax retained by tlua: `||=` and
	// `&&=` are removed from the language, so logical-assignment lowering is gone
	// with them, as nullish-coalescing lowering went with `??`.
	NewES2021Transformer transformers.TransformerFactory = nil
	// Optional chaining is lowered to Lua by NewLuaOptionalChainTransformer, which
	// runs unconditionally in the emit pipeline (Lua is always the target), so it
	// is intentionally absent from this target-gated JS downlevel chain.
	NewES2020Transformer = NewES2021Transformer
	NewES2019Transformer = NewES2020Transformer
	// The object rest/spread downlevel transform is removed in tlua: spread
	// elements, spread assignments, and binding-pattern rests no longer parse,
	// so there is nothing to lower.
	NewES2018Transformer transformers.TransformerFactory = newTaggedTemplateLiftRestrictionTransformer
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
	case core.ScriptTargetES2025, core.ScriptTargetES2024, core.ScriptTargetES2023, core.ScriptTargetES2022, core.ScriptTargetES2021,
		core.ScriptTargetES2020, core.ScriptTargetES2019, core.ScriptTargetES2018:
		// ES2021 down to ES2019 have no downlevel left, so their factories are nil
		// and must not be invoked; ES2018's own lowering starts at ES2017.
		return nil
	case core.ScriptTargetES2017:
		return NewES2018Transformer(opts)
	case core.ScriptTargetES2016:
		return NewES2017Transformer(opts)
	default: // other, older, option, transform maximally
		return NewES2016Transformer(opts)
	}
}

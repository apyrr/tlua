package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTsconfigComputedPropertyError(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @filename: tsconfig.json
{
    [|["oops!" + 42]|]: "true",
    "compilerOptions": { "lib": ["es5"] },
    "files": [
        "nonexistentfile.tlua"
    ],
    "compileOnSave": true
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.MarkTestAsStradaServer()
	// JSON keeps `:`, so a computed key `[...]` where a quoted name is
	// expected recovers cleanly to a single "string literal expected" over
	// the bracket expression -- no `=`-vs-`:` cascade.
	f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{
		{
			Range: lsproto.Range{
				Start: lsproto.Position{Line: 1, Character: 4},
				End:   lsproto.Position{Line: 1, Character: 18},
			},
			Message: lsproto.StringOrMarkupContent{String: new("String literal with double quotes expected.")},
			Code:    &lsproto.IntegerOrString{Integer: new(int32(1327))},
		},
	})
}

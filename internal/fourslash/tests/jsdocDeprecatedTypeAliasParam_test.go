package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestJsdocDeprecatedTypeAliasParam(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: a.tlua
/** @deprecated */
type Options = {}
/** @deprecated */
local function deprecatedFunction(options: [|Options|]) return options end
[|deprecatedFunction|]({})
/** @deprecated */
interface Props { a: number }
local function f(p: [|Props|]): [|Options|] return p end
declare v: [|Props|]
f(v)`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "a.tlua")
	diag := func(name string, i int) *lsproto.Diagnostic {
		return &lsproto.Diagnostic{
			Message: lsproto.StringOrMarkupContent{String: new("'" + name + "' is deprecated.")},
			Code:    &lsproto.IntegerOrString{Integer: new(int32(6385))},
			Range:   f.Ranges()[i].LSRange,
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagDeprecated},
		}
	}
	f.VerifySuggestionDiagnostics(t, []*lsproto.Diagnostic{
		diag("Options", 0),
		{
			Message: lsproto.StringOrMarkupContent{String: new("The signature '(options: Options): Options' of 'deprecatedFunction' is deprecated.")},
			Code:    &lsproto.IntegerOrString{Integer: new(int32(6387))},
			Range:   f.Ranges()[1].LSRange,
			Tags:    &[]lsproto.DiagnosticTag{lsproto.DiagnosticTagDeprecated},
		},
		diag("Props", 2),
		diag("Options", 3),
		diag("Props", 4),
	})
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestNavigateItemsLet(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @noLib: true
local [|c|] = 10;
function foo() {
    local [|d|] = 10;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyWorkspaceSymbol(t, []*fourslash.VerifyWorkspaceSymbolCase{
		{
			Pattern:     "c",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:     "c",
					Kind:     lsproto.SymbolKindVariable,
					Location: f.Ranges()[0].LSLocation(),
				},
			}),
		}, {
			Pattern:     "d",
			Preferences: nil,
			Exact: new([]*lsproto.SymbolInformation{
				{
					Name:          "d",
					Kind:          lsproto.SymbolKindVariable,
					Location:      f.Ranges()[1].LSLocation(),
					ContainerName: new("foo"),
				},
			}),
		},
	})
}

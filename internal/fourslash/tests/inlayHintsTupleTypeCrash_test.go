package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestInlayHintsTupleTypeCrash(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function iterateTuples(tuples: [string][]): void
  tuples.forEach(function(l) end)
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayFunctionParameterTypeHints: core.TSTrue,
		},
	})
}

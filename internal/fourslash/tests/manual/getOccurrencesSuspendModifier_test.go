package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// Highlighting the `suspend` modifier marks the modifier on the containing
// function and nothing else: tlua has no `await`, so a suspend function body
// holds no companion keyword to highlight alongside it.
func TestGetOccurrencesSuspendModifier(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|su/**/spend|] function f()
 return suspend function ()
   return 300;
 end
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}

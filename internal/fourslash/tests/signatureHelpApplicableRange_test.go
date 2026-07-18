package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpApplicableRange(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local obj = {
    foo(s: string): string {
        return s;
    }
};

local s =/*a*/ obj.foo("Hello, world!")/*b*/  
  /*c*/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	// Markers a, b, c should NOT show signature help (outside the call)
	f.VerifyNoSignatureHelpForMarkers(t, "a", "b", "c")
}

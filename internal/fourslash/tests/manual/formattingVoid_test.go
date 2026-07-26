package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// Adapted from the generated corpus, which formatted `void` in both of its
// positions. tlua deleted the void EXPRESSION and kept void as a TYPE, so the
// markers that spaced out `void a` and `void (0)` are gone and the ones that
// space out an annotation stay. Quarantining the whole test would have taken the
// surviving half with it -- this is the only formatting coverage void's type has.
func TestFormattingVoid(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/  local x: () =>           void    ;
/*2*/  local y:     void    ;
/*3*/  function test(a:void,b:string) end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `local x: () => void;`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `local y: void;`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `function test(a: void, b: string) end`)
}

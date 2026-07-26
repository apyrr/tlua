package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTypeAssertionsFormatting(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local _ = ( <  any   >      publisher);/*1*/
local _ =  <  any  >      3;/*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `local _ = (<any>publisher);`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `local _ = <any>3;`)
}

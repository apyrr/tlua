package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormattingSpaceAfterCommaBeforeOpenParen(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `foo(a,(b))/*1*/
foo(a,(c as b).d)/*2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.Insert(t, ";")
	f.VerifyCurrentLineContent(t, `foo(a, (b));`)
	f.GoToMarker(t, "2")
	f.Insert(t, ";")
	f.VerifyCurrentLineContent(t, `foo(a, (c as b).d);`)
}

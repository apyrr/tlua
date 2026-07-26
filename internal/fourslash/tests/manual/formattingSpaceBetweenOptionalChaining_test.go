package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormattingSpaceBetweenOptionalChaining(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/local _ = a    ?.    b   ?.   c   .   d;
/*2*/local _ = o    .  m()   ?.   length;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `local _ = a?.b?.c.d;`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `local _ = o.m()?.length;`)
}

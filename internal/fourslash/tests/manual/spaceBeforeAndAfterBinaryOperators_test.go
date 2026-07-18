package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSpaceBeforeAndAfterBinaryOperators(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local i = 0;
/*1*/(i=i+1,i=i+1);
/*2*/(i=i+1,i=i+1);
/*3*/(1,2);
/*4*/(i=i+1,2);
/*5*/(i=i+1,i=i+1,i=i+1,i=i-1,2);
local s = 'foo';
/*6*/for (local i = 0,ii = 2; i < s.length; ii=ii+1,i=i+1) {
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `(i = i + 1, i = i + 1);`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `(i = i + 1, i = i + 1);`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `(1, 2);`)
	f.GoToMarker(t, "4")
	f.VerifyCurrentLineContent(t, `(i = i + 1, 2);`)
	f.GoToMarker(t, "5")
	f.VerifyCurrentLineContent(t, `(i = i + 1, i = i + 1, i = i + 1, i = i - 1, 2);`)
	f.GoToMarker(t, "6")
	f.VerifyCurrentLineContent(t, `for (local i = 0, ii = 2; i < s.length; ii = ii + 1, i = i + 1) {`)
}

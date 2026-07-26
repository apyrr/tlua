package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormattingNonNullAssertionOperator(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/ local  _  =  'bar' ! ;
/*2*/ local  _  =  ( 'bar' ) ! ;
/*3*/ local  _  =  'bar' [ 1 ] ! ;
/*4*/ local  bar  =  'bar' . foo ! ;
/*5*/ local  foo  =  bar ! ;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `local _ = 'bar'!;`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `local _ = ('bar')!;`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `local _ = 'bar'[1]!;`)
	f.GoToMarker(t, "4")
	f.VerifyCurrentLineContent(t, `local bar = 'bar'.foo!;`)
	f.GoToMarker(t, "5")
	f.VerifyCurrentLineContent(t, `local foo = bar!;`)
}

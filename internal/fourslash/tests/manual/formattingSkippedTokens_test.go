package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormattingSkippedTokens(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/foo(): Bar { }
/*2*/function Foo      () #   { }
/*3*/4+:5
 namespace M {
function a(
/*4*/    : T) { }
}
/*5*/local x       =`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `foo(): Bar { }`)
	f.GoToMarker(t, "2")
	// `#` is tlua's length operator now, not a skipped token, so the parser recovers
	// differently here: `#   { }` parses as a bare expression statement, which Lua
	// does not allow (TLUA100057). The formatter stays conservative over that
	// erroneous span and leaves it verbatim, while the well-formed head of the line
	// still collapses `Foo      ()` to `Foo()`.
	f.VerifyCurrentLineContent(t, `function Foo() #   { }`)
	f.GoToMarker(t, "3")
	// `4+` is a bare expression statement (TLUA100057) — its missing-operand
	// error sits on the unconsumed `:`, outside the statement, so it does not
	// defer. The formatter leaves the errored `4+` span verbatim and re-spaces
	// only the skipped-token tail.
	f.VerifyCurrentLineContent(t, `4+: 5`)
	f.GoToMarker(t, "4")
	f.VerifyCurrentLineContent(t, `    : T) { }`)
	f.GoToMarker(t, "5")
	f.VerifyCurrentLineContent(t, `    local x =`)
}

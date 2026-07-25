package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// Signature help drops the parameter a colon call's receiver fills, because the
// caller cannot write into it. The receiver fills one argument SLOT though, and
// that is a whole parameter only when the parameter takes one: a lone rest
// parameter goes on absorbing what the caller writes, so it has to stay. Dropping
// it left the popup advertising no parameters at all for a call that takes any
// number of them.
func TestTluaColonCallSignatureHelpRestParameter(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
interface String {
  joinAll(...: string): string;
}
local s: string = "x";
s:joinAll(/*colon*/);
s.joinAll(/*dot*/);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	// The rest parameter survives the elision and stays the active one.
	f.GoToMarker(t, "colon")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:           "joinAll(...: string): string",
		ParameterCount: 1,
	})

	// The dot form writes into the same rest parameter; nothing is elided there
	// either, so the two spellings agree.
	f.GoToMarker(t, "dot")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:           "joinAll(...: string): string",
		ParameterCount: 1,
	})
}

// A leading ordinary parameter ahead of a rest parameter IS wholly filled by the
// receiver, so it is dropped and the rest parameter shifts up.
func TestTluaColonCallSignatureHelpRestAfterSelf(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
interface String {
  joinFrom(self: string, ...: string): string;
}
local s: string = "x";
s:joinFrom(/*1*/);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:           "joinFrom(...: string): string",
		ParameterCount: 1,
	})
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

// The bundled lib declares `self` as parameter 0 of every receiver-bearing
// member, and signature help compensates for the colon form by starting the
// written arguments at parameter 1. The two have to agree: this pins the active
// parameter for both call forms into the same lib signature.
func TestSignatureHelpLuaLibColonCall(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local s: string = "hello";

s:sub(/*1*/);
s:sub(1, /*2*/);

s.sub(/*3*/);
s.sub(s, /*4*/);
s.sub(s, 1, /*5*/);

declare h: LuaFile;

h:seek(/*6*/);
h:seek("set", /*7*/);`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

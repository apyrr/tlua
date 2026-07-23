package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpAfterParameterVS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `type Type = (a, b, c) => void
local a: Type = function(a/*1*/, b/*2*/) end
local b: Type = function (a/*3*/, b/*4*/) end
local c: Type = function({ /*5*/a: { b/*6*/ }}/*7*/ = { }/*8*/, [b/*9*/]/*10*/, .../*11*/c/*12*/) end`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

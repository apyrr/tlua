package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpRestArgs1VS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function fn(a: number, b: number, c: number) end
local a = [1, 2] as const;
local b = [1] as const;

fn(...a, /*1*/);
fn(/*2*/, ...a);

fn(...b, /*3*/);
fn(/*4*/, ...b, /*5*/);`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

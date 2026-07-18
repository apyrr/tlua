package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// Calling a union of signatures resolves a combined signature whose rest
// parameter the checker mints itself, with a name of its own ("args"). Signature
// help is where that synthetic parameter is displayed, and in tlua it must still
// read as a vararg: `...args: string[]` is not tlua syntax, and it shows the
// array where the pack's element type belongs.
func TestTluaVarargSignatureHelp(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare cond: boolean;
declare function fixedArity(a: number, b: boolean): void;
declare function varargArity(...: string): void;

local combined = cond and fixedArity or varargArity;
combined(/**/);
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

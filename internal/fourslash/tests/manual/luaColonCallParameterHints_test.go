package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/testutil"
)

// A colon call binds signature parameter 0 (self) to the receiver, so
// parameter-indexed LS features must start the written arguments at
// parameter 1: hints label `x`/`y`, never `self`.
func TestLuaColonCallInlayParameterHints(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local T = { n = 0 };
function T:add(x: number, y: number): number
  return self.n + x + y;
end
T:add(1, 2);
T.add(T, 3, 4);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayParameterNameHints: "all",
		},
	})
}

// Signature help inside a colon call's argument list must highlight the
// written parameter (`x` at the first argument), not `self`.
func TestLuaColonCallSignatureHelp(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local T = { n = 0 };
function T:add(x: number, y: number): number
  return self.n + x + y;
end
T:add(/*1*/1, /*2*/2);
T.add(T, /*3*/3, 4);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

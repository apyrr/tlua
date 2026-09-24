package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/testutil"
)

// An assertion that also returns values (`R asserts x`, Lua's assert) shows its
// returned values next to the clause in signature help, not just `asserts x`.
func TestLuaAssertsWithReturnSignatureHelp(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local function expect(v: unknown): string asserts v is string
  return v
end
expect(/*1*/"x");
assert(/*2*/1);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

// Inlay type hints print the returned values of such an assertion too.
func TestLuaAssertsWithReturnInlayHints(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `local check = assert;
local expect = function(v: unknown): string asserts v is string
  return v
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayVariableTypeHints: core.TSTrue,
		},
	})
}

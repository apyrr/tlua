package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTluaMultiReturnQuickInfo(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
function pair/*1*/(): (number, string)
  return 1, "a";
end

function inferred/*2*/()
  return 1, "a";
end

local a/*3*/, b/*4*/, c/*5*/ = pair();

local single/*6*/ = pair();
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "function pair(): (number, string)", "")
	f.VerifyQuickInfoAt(t, "2", "function inferred(): (number, string)", "")
	f.VerifyQuickInfoAt(t, "3", "local a: number", "")
	f.VerifyQuickInfoAt(t, "4", "local b: string", "")
	f.VerifyQuickInfoAt(t, "5", "local c: nil", "")
	f.VerifyQuickInfoAt(t, "6", "local single: number", "")
}

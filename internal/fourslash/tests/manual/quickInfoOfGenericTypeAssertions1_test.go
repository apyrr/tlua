package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoOfGenericTypeAssertions1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function f<T>(x: T): T return null; end
local /*1*/r = function<T>(x: T) return x end;
local /*2*/r2 = f as <T>(x: T) => T;
local a;
local /*3*/r3 = a as <T>(x: <A>(y: A) => A) => T;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "local r: <T>(x: T) => T", "")
	f.VerifyQuickInfoAt(t, "2", "local r2: <T>(x: T) => T", "")
	f.VerifyQuickInfoAt(t, "3", "local r3: <T>(x: <A>(y: A) => A) => T", "")
}

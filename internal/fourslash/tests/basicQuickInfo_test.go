package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestBasicQuickInfo(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
/**
 * Some var
 */
local someVar/*1*/ = 123;

/**
 * Other var
 * See {@link someVar}
 */
local otherVar/*2*/ = someVar;

class Foo/*3*/ {
	#bar: string;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "local someVar: number", "Some var")
	f.VerifyQuickInfoAt(t, "2", "local otherVar: number", "Other var\nSee [someVar](file:///basicQuickInfo.tlua#5,7-5,14)")
}

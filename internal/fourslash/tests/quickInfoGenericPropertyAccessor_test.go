package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoGenericPropertyAccessor(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare o: {
    f: <T>(x: T) => T
    get g(): <T>(x: T) => T
}

declare x: number

o.f/*1*/(x)
o.g/*2*/(x)
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "(property) f: <number>(x: number) => number", "")
	f.VerifyQuickInfoAt(t, "2", "(accessor) g: <number>(x: number) => number", "")
}

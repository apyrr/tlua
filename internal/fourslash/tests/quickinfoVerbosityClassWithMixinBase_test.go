package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickinfoVerbosityClassWithMixinBase(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `
class Base {}

declare Mixin: new () => Base & { mixed: string };

class Derived/*1*/ extends Mixin {}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"1": {0, 1}})
}

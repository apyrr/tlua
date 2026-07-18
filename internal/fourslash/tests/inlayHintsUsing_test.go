package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestInlayHintsUsing(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	// Object-literal methods are removed in tlua; the disposable member is a
	// Lua-keyed function-valued field instead.
	const content = `// @target: esnext
using _defer = {
	[Symbol.dispose] = function() {},
};`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineInlayHints(t, nil /*span*/, &lsutil.UserPreferences{InlayHints: lsutil.InlayHintsPreferences{
		IncludeInlayVariableTypeHints: core.TSTrue,
	}})
}

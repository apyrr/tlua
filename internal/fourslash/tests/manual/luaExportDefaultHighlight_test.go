package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// With the case/default clause arm gone from syntactic highlights, the
// `default` keyword of `export default` reaches the modifier path. Lock that
// behavior (upstream TypeScript treats export-default's `default` as a
// modifier for highlighting too).
func TestLuaExportDefaultHighlight(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `export /*default*/default function f() {}
f();`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, "default")
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestRangeFormatStartingInsideJSDocComment(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	content := "// @Filename: /a.tlua\n" +
		"/**\n" +
		" * @a\n" +
		" * `\n" +
		"/*s*/ * @b\n" +
		" */\n" +
		"function f() {}/*e*/"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatSelection(t, "s", "e")
}

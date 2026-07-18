package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoJSDocCodefenceAtSign(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * text
 * @example Foo
 * ` + "```" + `
 * @Embed[asfasdfasf]
 * ` + "```" + `
 * becomes
 * ` + "```html" + `
 * <div></div>
 * ` + "```" + `
 */
local /*1*/x = 1;

/**
 * Some text
 * ` + "```" + `
 * @tag inside code
 * ` + "```" + `
 * @param y - a number
 */
function /*2*/foo(y: number) {}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

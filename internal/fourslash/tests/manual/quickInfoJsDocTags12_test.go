package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoJsDocTags12(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/**
 * @param {Object} options the args object
 * @param {number} options.a first number
 * @param {number} options.b second number
 * @param {Function} callback the callback function
 * @returns {number}
 */
function /**/f(options, callback = nil) end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

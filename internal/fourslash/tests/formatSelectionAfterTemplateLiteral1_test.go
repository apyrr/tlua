package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormatSelectionAfterTemplateLiteral1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = "local a = `head${\"x\"};\n`;\n\n/*begin*/export local f = function()\n    return `world`;\n/*end*/end\n"
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatSelection(t, "begin", "end")
	f.VerifyCurrentFileContent(t, "local a = `head${\"x\"};\n`;\n\nexport local f = function()\n    return `world`;\nend\n")
}

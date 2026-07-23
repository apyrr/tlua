package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormatDocumentZeroTabSize(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo()
    if (true) {
        local x = 1;
    }
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	opts := f.GetOptions()
	opts.FormatCodeSettings.TabSize = 0
	opts.FormatCodeSettings.IndentSize = 0
	opts.FormatCodeSettings.ConvertTabsToSpaces = core.TSTrue
	f.Configure(t, opts)
	f.FormatDocument(t, "")
	f.VerifyCurrentFileContent(t, "function foo()\nif (true) {\nlocal x = 1;\n}\nend")
}

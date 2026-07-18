package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestRenameBuiltinTypes(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
local arr: /*1*/Array<number> = [];
local map1: /*2*/Map<string, number> = new Map();
local prom: /*3*/Promise<void> = Promise.resolve();
local str: /*4*/string = "hello";
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	// All of these should fail because they're library/builtin types
	for _, marker := range []string{"1", "2", "3", "4"} {
		f.GoToMarker(t, marker)
		f.VerifyRenameFailed(t, nil /*preferences*/)
	}
}

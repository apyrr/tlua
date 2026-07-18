package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestLuaGlobalTableQuickInfo(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /write.tlua
_G.shared = 1;
// @Filename: /read.tlua
local bare = shared/*bare*/;
local qualified = _G.shared/*qualified*/;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "bare", "local shared: number", "")
	f.VerifyQuickInfoAt(t, "qualified", "local shared: number", "")
}

func TestLuaGlobalTableFindAllReferences(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /write.tlua
_G.[|shared|] = 1;
// @Filename: /read.tlua
local bare = [|shared|];
local qualified = _G.[|shared|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t)
}

func TestLuaGlobalTableGoToDefinition(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /write.tlua
_G.shared = 1;
// @Filename: /read.tlua
local bare = /*use*/shared;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, false /*includeOriginalSelectionRange*/, "use")
}

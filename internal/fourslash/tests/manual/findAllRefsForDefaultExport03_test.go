package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFindAllRefsForDefaultExport03(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/function /*2*/f()
    return 100;
end

/*3*/export default /*4*/f;

local x: typeof /*5*/f;

local y = /*6*/f();

/*7*/namespace /*8*/f {
    local localValue = 100;
}`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t, "1", "2", "3", "4", "5", "6", "7", "8")
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestRenameForDefaultExport03(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|function /*1*/[|{| "contextRangeIndex": 0 |}f|]() {
    return 100;
}|]

[|export default /*2*/[|{| "contextRangeIndex": 2 |}f|];|]

local x: typeof /*3*/[|f|];

local y = /*4*/[|f|]();

/**
 *  Commenting [|{| "inComment": true |}f|]
 */
[|namespace /*5*/[|{| "contextRangeIndex": 7 |}f|] {
    local localValue = 100;
}|]`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, ToAny(core.Filter(f.GetRangesByText().Get("f"), func(r *fourslash.RangeMarker) bool { return r.Marker == nil || r.Marker.Data["inComment"] == nil }))...)
}

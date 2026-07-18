package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestRenameForDefaultExport01(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|export default class [|{| "contextRangeIndex": 0 |}DefaultExportedClass|] {
}|]
/*
 *  Commenting [|{| "inComment": true |}DefaultExportedClass|]
 */

local x: [|DefaultExportedClass|];

local y = new [|DefaultExportedClass|];`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	ranges := f.GetRangesByText().Get("DefaultExportedClass")

	var markerOrRanges []fourslash.MarkerOrRangeOrName
	for _, r := range ranges {
		if !(r.Marker != nil && r.Marker.Data != nil && r.Marker.Data["inComment"] == true) {
			markerOrRanges = append(markerOrRanges, r)
		}
	}

	f.VerifyBaselineRename(t, nil /*preferences*/, markerOrRanges...)
}

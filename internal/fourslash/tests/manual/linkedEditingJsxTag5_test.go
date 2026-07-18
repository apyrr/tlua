package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestLinkedEditingJsxTag5(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @FileName: /unclosedElement.tsx
local jsx = (
    <div/*0*/>
        </*1start*/div/*1*/>
    <//*2start*/div/*2*/>/*3*/
);/*4*/
// @FileName: /mismatchedElement.tsx
local jsx = (
    /*5*/</*6start*/div/*6*/>
        <//*7start*/div/*7*/>
    </*8*//div/*9*/>/*10*/
);
// @Filename: /invalidClosing.tsx
local jsx = (
   <di/*11*/v>
   </*12*/ //*13*/div>
);`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	linkedCursors1 := []lsproto.Range{
		{Start: f.MarkerByName(t, "1start").LSPosition, End: f.MarkerByName(t, "1").LSPosition},
		{Start: f.MarkerByName(t, "2start").LSPosition, End: f.MarkerByName(t, "2").LSPosition},
	}
	linkedCursors2 := []lsproto.Range{
		{Start: f.MarkerByName(t, "6start").LSPosition, End: f.MarkerByName(t, "6").LSPosition},
		{Start: f.MarkerByName(t, "7start").LSPosition, End: f.MarkerByName(t, "7").LSPosition},
	}

	f.VerifyLinkedEditing(t, map[string][]lsproto.Range{
		"0":  nil,
		"1":  linkedCursors1,
		"2":  linkedCursors1,
		"3":  nil,
		"4":  nil,
		"5":  nil,
		"6":  linkedCursors2,
		"7":  linkedCursors2,
		"8":  nil,
		"9":  nil,
		"10": nil,
		"11": nil, // this tag does not parse as a closing tag
		"12": nil,
		"13": nil,
	})
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestFormatSuspendKeyword(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/*1*/local x = suspend         function() return 1 end;
/*2*/local y = suspend function() return 1 end;
/*3*/local z = suspend    function   () return 1; end;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.FormatDocument(t, "")
	f.GoToMarker(t, "1")
	f.VerifyCurrentLineContent(t, `local x = suspend function() return 1 end;`)
	f.GoToMarker(t, "2")
	f.VerifyCurrentLineContent(t, `local y = suspend function() return 1 end;`)
	f.GoToMarker(t, "3")
	f.VerifyCurrentLineContent(t, `local z = suspend function() return 1; end;`)
}

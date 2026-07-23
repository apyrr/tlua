package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestInlayHintsInteractiveParameterNamesInSpan1(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo1 (a: number, b: number) end
function foo2 (c: number, d: number) end
function foo3 (e: number, f: number) end
function foo4 (g: number, h: number) end
function foo5 (i: number, j: number) end
function foo6 (k: number, i: number) end

function c1 () foo1(/*a*/1, /*b*/2); end
function c2 () foo2(/*c*/1, /*d*/2); end
function c3 () foo3(/*e*/1, /*f*/2); end
function c4 () foo4(/*g*/1, /*h*/2); end
function c5 () foo5(/*i*/1, /*j*/2); end
function c6 () foo6(/*k*/1, /*l*/2); end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	start := f.MarkerByName(t, "c")
	end := f.MarkerByName(t, "h")
	span := &lsproto.Range{Start: start.LSPosition, End: end.LSPosition}
	f.VerifyBaselineInlayHints(t, span, &lsutil.UserPreferences{
		InlayHints: lsutil.InlayHintsPreferences{
			IncludeInlayParameterNameHints: "literals",
		},
	})
}

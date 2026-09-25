package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

// The node builder reuses a declaration's annotation nodes with their source positions, so
// a printer handed the real source file reads the comments inside the annotation back out
// and copies them into the display -- twice on a one-line type literal, once as the trailing
// comment of the member before and once more ahead of the member after. Every display path
// prints with comments removed, the way checker.TypeToString does.
const tluaDisplayOmitsSourceCommentsContent = `// @Filename: /a.tlua
local function cm(x: { p: number, --[[cp]] q: string }): { a: number, --[[ca]] b: string }
  return nil as any
end
cm(/*call*/nil as any)
local /*v*/v: { a: number, -- line
  --[[cv]] b: string } = nil as any
local function gen<T extends { a: number, --[[ct]] b: string }>(x: T): T return x end
gen</*typeArgs*/{ a: number, b: string }>(nil as any)
type Box<T extends { a: number, --[[cb]] b: string }> = T
local bx: Box</*typeRef*/`

func TestTluaSignatureHelpOmitsSourceComments(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, tluaDisplayOmitsSourceCommentsContent)
	defer done()

	// Parameter and return type both reuse the declared annotations.
	f.GoToMarker(t, "call")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:          "cm(x: { p: number; q: string; }): { a: number; b: string; }",
		ParameterSpan: "x: { p: number; q: string; }",
	})

	// A type parameter's constraint prints on one line, with its comment removed.
	f.GoToMarker(t, "typeArgs")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:          "gen<T extends { a: number; b: string; }>(x: { a: number; b: string; }): { a: number; b: string; }",
		ParameterSpan: "T extends { a: number; b: string; }",
	})
	f.GoToMarker(t, "typeRef")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:          "Box<T extends { a: number; b: string; }>",
		ParameterSpan: "T extends { a: number; b: string; }",
	})
}

// The classified (Visual Studio) hover prints through the node builder and printer rather
// than checker.TypeToString, so it is the path that needs the same guarantee. A `--` line
// comment copied into one-line hover text would read as commenting out the rest of the type.
func TestTluaClassifiedHoverOmitsSourceComments(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, tluaDisplayOmitsSourceCommentsContent)
	defer done()
	f.VerifyQuickInfoAt(t, "v", "local v: { a: number; b: string; }", "")
}

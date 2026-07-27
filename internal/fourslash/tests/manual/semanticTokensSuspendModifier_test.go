package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// A suspend function carries tlua's custom `suspend` semantic-token modifier,
// never the LSP predefined `async` one — tlua has no `async`, so reporting it
// would name a runtime model this language does not have. The harness client
// advertises only the predefined modifier set, exactly like vscode-languageclient,
// so this test also pins that the server keeps its custom modifiers in the
// legend regardless of advertisement (the VS Code `semanticTokenModifiers`
// contribution feeds theming, not LSP capabilities — filtering by advertisement
// would strip `suspend` for every real client). A lost modifier is silent in
// the editor, so it is pinned here.
func TestSemanticTokensSuspendModifier(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `suspend function fetchOk(): boolean
  return true;
end

function plain(): boolean
  return true;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "function.declaration.suspend", Text: "fetchOk"},
		{Type: "function.declaration", Text: "plain"},
	})
}

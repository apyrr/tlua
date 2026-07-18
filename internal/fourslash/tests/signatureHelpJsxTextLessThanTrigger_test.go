package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpJsxTextLessThanTrigger(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `//@Filename: test.tsx
//@jsx: react
declare React: any;
declare function Text(props: { children?: any }): any;

local text = () => {
	return <Text>/*m*/</Text>;
};`

	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.GoToMarker(t, "m")
	f.Insert(t, "<")
	f.VerifyNoSignatureHelpWithContext(t, &lsproto.SignatureHelpContext{
		TriggerKind:      lsproto.SignatureHelpTriggerKindTriggerCharacter,
		TriggerCharacter: new("<"),
		IsRetrigger:      false,
	})
}

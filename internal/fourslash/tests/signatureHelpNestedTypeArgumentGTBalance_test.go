package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// Regression for getPossibleTypeArgumentsInfo (see PR #3222): when scanning backward through
// explicit type arguments, the < / > balance must accumulate across adjacent closing angles.
func TestSignatureHelpNestedTypeArgumentGTBalance(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare function f<T, U>(): void;
type A<T> = T;
type B<T> = T;
type C<T> = T;
f<A<B<C<number>>>, /*nested*/;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.GoToMarker(t, "nested")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:           "f<T, U>(): void",
		ParameterName:  "U",
		ParameterSpan:  "U",
		ParameterCount: 2,
	})
}

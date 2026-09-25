package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestGenericFunctionSignatureHelp3(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `function foo1<T>(x: number, callback: (y1: T) => number) end
function foo2<T>(x: number, callback: (y2: T) => number) end
function foo3<T>(x: number, callback: (y3: T) => number) end
function foo4<T>(x: number, callback: (y4: T) => number) end
function foo5<T>(x: number, callback: (y5: T) => number) end
 IDE shows the results on the right of each line, fourslash says different
foo1(/*1*/               // signature help shows y as T
foo2(1,/*2*/             // signature help shows y as {}
foo3(1, (/*3*/           // signature help shows y as T
foo4<string>(1,/*4*/     // signature help shows y as string
foo5<string>(1, (/*5*/   // signature help shows y as T`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToMarker(t, "1")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "foo1(x: number, callback: (y1: unknown) => number): void"})
	f.GoToMarker(t, "2")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "foo2(x: number, callback: (y2: unknown) => number): void"})
	f.GoToMarker(t, "3")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "callback(y3: unknown): number"})
	f.GoToMarker(t, "4")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "foo4(x: number, callback: (y4: string) => number): void"})
	f.GoToMarker(t, "5")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{Text: "callback(y5: string): number"})
}

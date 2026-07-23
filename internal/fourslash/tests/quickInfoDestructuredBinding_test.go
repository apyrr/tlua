package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoDestructuredBinding(t *testing.T) {
	t.Parallel()
	fourslash.SkipIfFailing(t)
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `
function f({ /*1*/x }: { x: number }) end
function g([/*2*/y]: number[]) end
function h({ a: { /*3*/b } }: { a: { b: string } }) end
local { /*4*/c } = { c: 42 };
local { /*5*/d } = { d: "hello" };
local { /*6*/e } = { e: true };
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	// Destructured object binding parameters should show "(parameter)" not "var"
	f.VerifyQuickInfoAt(t, "1", "(parameter) x: number", "")
	// Destructured array binding parameters should show "(parameter)" not "var"
	f.VerifyQuickInfoAt(t, "2", "(parameter) y: number", "")
	// Nested destructured parameters should also show "(parameter)"
	f.VerifyQuickInfoAt(t, "3", "(parameter) b: string", "")
	// Destructured const/let/var bindings should show their proper keyword
	f.VerifyQuickInfoAt(t, "4", "local c: number", "")
	f.VerifyQuickInfoAt(t, "5", "local d: string", "")
	f.VerifyQuickInfoAt(t, "6", "local e: boolean", "")
}

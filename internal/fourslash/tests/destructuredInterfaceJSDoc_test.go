package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestDestructuredInterfaceJSDoc(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
interface FooBar {
    /** foo comment */
    foo: number;
    /** bar comment */
    bar: string;
    /** baz comment */
    baz: string;
}

declare fubar: FooBar;

local {/*1*/foo, /*2*/bar, /*3*/baz: /*4*/biz} = fubar;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

func TestDestructuredInterfaceJSDocWithRename(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
interface FooBar {
    /** foo comment */
    foo: number;
    /** bar comment */
    bar: string;
}

declare fubar: FooBar;

local {foo: /*1*/myFoo, bar: /*2*/myBar} = fubar;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

func TestDestructuredWithOwnJSDoc(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
interface Foo {
    /** This is bar from the interface */
    bar: string;
    /** This is baz from the interface */
    baz: number;
}

declare foo: Foo;

/** Comment on the variable statement. */
local {
    /** Comment on bar destructuring. */ /*1*/bar,
    /** Comment on baz destructuring. */ /*2*/baz
} = foo;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

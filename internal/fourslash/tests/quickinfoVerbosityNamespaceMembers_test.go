package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickinfoVerbosityNamespaceMembers(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare namespace NS/*1*/ {
    type StringAlias = string;
    type Pair<T> = { first: T; second: T };

    class MyClass {
        name: string;
        greet(): void;
    }

    interface MyInterface {
        id: number;
        label: string;
    }

    local value: number;
    function doSomething(x: string): boolean;
}
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHoverWithVerbosity(t, map[string][]int{"1": {0, 1}})
}

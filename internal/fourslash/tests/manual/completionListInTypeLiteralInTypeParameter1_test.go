package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionListInTypeLiteralInTypeParameter1(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Foo {
    one: string;
    two: number;
    333: symbol;
    '4four': boolean;
    '5 five': object;
    number: string;
    Object: number;
}

interface Bar<T extends Foo> {
    foo: T;
}

local foobar: Bar<{/**/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Unsorted: []fourslash.CompletionsExpectedItem{
				"one",
				"two",
				// tlua: `333` is a number key (t[333] != t["333"]), so it
				// completes as the bare numeric member `333`, not the quoted
				// string key `"333"`. The '4four'/'5 five' string keys stay quoted.
				"333",
				"\"4four\"",
				"\"5 five\"",
				"number",
				"Object",
			},
		},
	})
}

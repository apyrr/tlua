package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestArgumentCompletions(t *testing.T) {
	t.Parallel()
	fourslash.SkipIfFailing(t)

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
function foo(a: "a", b: "b") end
foo("a", /*1*/);


local t3 = ['x', 'y', 'z'] as const;
local x: [string, string, string, 'a' | 'b'] = [...t3, /*2*/];
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{`"b"`},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{`"b"`},
		},
	})
}

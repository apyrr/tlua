package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// In an object literal, a number-key member is written `1: v` (a numeric
// literal name) while the disjoint string key is written `"1": v`. Their
// completions are therefore distinct entries with distinct labels.
func TestTluaNumberKeyObjectLiteralCompletion(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface Opts { 1: string; "1": number; }
local o: Opts = { /*both*/ };
interface NegOpts { [-1]: string; }
local n: NegOpts = { /*neg*/ };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "both", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"1",
				"\"1\"",
			},
		},
	})
	// A negative number key cannot be written `-1:`; its object-literal form is
	// the computed field `[-1] = v`, so the completion label is `[-1]`.
	f.VerifyCompletions(t, "neg", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"[-1]",
			},
		},
	})
}

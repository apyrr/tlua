package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// A number key has no string-literal spelling, so it is excluded from
// string-literal (t["|"]) completions; and an already-written number-key
// member is filtered from object-literal completions by its identity name.
func TestTluaNumberKeyStringAndFilterCompletion(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I3 { 1: number; 2: number; foo: string; }
declare i: I3;
local s = i["/*str*/"];
local o: I3 = { [1] = 5, /*filt*/ };`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	// String-literal completion offers only string keys; number keys 1 and 2
	// have no "..." spelling and are excluded.
	f.VerifyCompletions(t, "str", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"foo",
			},
			Excludes: []string{"1", "2"},
		},
	})
	// The already-written number key 1 (`[1] = 5`) is filtered; 2 and foo remain.
	f.VerifyCompletions(t, "filt", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				"2",
				"foo",
			},
			Excludes: []string{"1"},
		},
	})
}

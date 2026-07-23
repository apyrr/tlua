package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionRemovedStatementKeywords(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")

	const content = `/*top*/
function f()
  /*body*/
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	f.VerifyCompletions(t, []string{"top", "body"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"if", "local", "function"},
			Excludes: []string{"class", "try", "catch", "finally", "with", "export", "default", "from"},
		},
	})
}

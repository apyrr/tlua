package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionsLiteralDirectlyInArgumentWithNullableConstraint(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true

declare function func<
  T extends 'a' | 'b' | nil = nil,
>(arg?: T): string;

func('/*1*/');`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"1"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"a",
				"b",
			},
		},
	})
}

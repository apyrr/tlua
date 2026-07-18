package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

// `and`, `or` and `not` are identifier-shaped but scan as operators, so a
// property so named cannot be completed as a bare `t.and` -- that does not parse.
// Completing one rewrites the dot into the bracket form, as for any other
// non-identifier name. An ordinary keyword such as `while` is a legal member name
// and stays bare.
func TestCompletionWordOperatorMemberName(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `declare t: { ["and"]: number, ["or"]: number, ["not"]: number, while: number, plain: number };
t[|./*1*/|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()

	bracket := func(name string) *lsproto.CompletionItem {
		return &lsproto.CompletionItem{
			Label:      name,
			InsertText: new("[\"" + name + "\"]"),
			FilterText: new("." + name),
			TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
				TextEdit: &lsproto.TextEdit{
					NewText: "[\"" + name + "\"]",
					Range:   f.Ranges()[0].LSRange,
				},
			},
		}
	}

	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				bracket("and"),
				bracket("or"),
				bracket("not"),
				&lsproto.CompletionItem{Label: "while"},
				&lsproto.CompletionItem{Label: "plain"},
			},
		},
	})
}

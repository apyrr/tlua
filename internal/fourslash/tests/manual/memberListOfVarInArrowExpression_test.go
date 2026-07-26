package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestMemberListOfVarInArrowExpression(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface IMap<T> {
    [key: string]: T;
}
local map: IMap<{ a1: string; }[]>;
local categories: string[];
each(categories, function(category)
    local changes = map[category];
    changes[0][|./*1*/a1|];
    return each(changes, function(change)
    end);
end);
function each<T>(items: T[], handler: (item: T) => void) end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "1", "(property) a1: string", "")
	// `changes[0]` reads through the number index, so the element may not be there
	// and the completion offers the access as an optional chain.
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a1",
					Detail:     new("(property) a1: string"),
					InsertText: new("?.a1"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.a1",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
}

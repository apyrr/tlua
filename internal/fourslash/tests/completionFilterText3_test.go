package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/ls"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionFilterText3(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strict: true
declare foo1: { b: number; "a bc": string; };
if true then
    foo1[|.|]/*1*/
else
    foo1[|.a|]/*2*/
end

declare foo2: { b: number; "a bc": string; } | undefined;
if true then
    foo2[|.|]/*3*/
elseif false then
    foo2[|.a|]/*4*/
else
    foo2[|?.|]/*5*/
end
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "1", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "[\"a bc\"]",
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "[\"a bc\"]",
							Range:   f.Ranges()[1].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "3", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[2].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "4", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new(".a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[3].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "5", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      "a bc",
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new("?.[\"a bc\"]"),
					FilterText: new("?.a bc"),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: "?.[\"a bc\"]",
							Range:   f.Ranges()[4].LSRange,
						},
					},
				},
			},
		},
	})
}

package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/ls"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTsxCompletion12(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `//@Filename: file.tsx
// @jsx: preserve
// @noLib: true
interface JsxElement { }
interface JsxIntrinsicElements {
}
interface JsxElementAttributesProperty { props; }
interface OptionPropBag {
    propx: number
    propString: "hell"
    optional?: boolean
}
declare function Opt(attributes: OptionPropBag): JsxElement;
local opt = <Opt /*1*/ />;
local opt1 = <Opt [|prop|]/*2*/ />;
local opt2 = <Opt propx={100} /*3*/ />;
local opt3 = <Opt propx={100} optional /*4*/ />;
local opt4 = <Opt wrong /*5*/ />;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"1", "5"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"propString",
				"propx",
				&lsproto.CompletionItem{
					Label:      "optional?",
					InsertText: new("optional"),
					FilterText: new("optional"),
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextOptionalMember)),
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
			Exact: []fourslash.CompletionsExpectedItem{
				"propString",
				"propx",
				&lsproto.CompletionItem{
					Label:      "optional?",
					FilterText: new("optional"),
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextOptionalMember)),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						InsertReplaceEdit: &lsproto.InsertReplaceEdit{
							NewText: "optional",
							Insert:  f.Ranges()[0].LSRange,
							Replace: f.Ranges()[0].LSRange,
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
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"propString",
				&lsproto.CompletionItem{
					Label:      "optional?",
					InsertText: new("optional"),
					FilterText: new("optional"),
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextOptionalMember)),
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
			Exact: []fourslash.CompletionsExpectedItem{
				"propString",
			},
		},
	})
}

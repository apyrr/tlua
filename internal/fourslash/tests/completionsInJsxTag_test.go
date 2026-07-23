package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionsInJsxTag(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: preserve
// @Filename: /a.tsx
interface JsxElement {}
interface JsxIntrinsicElements {
    div: {
        /** Doc */
        foo: string
        /** Label docs */
        "aria-label": string
    }
}
<div /*1*/ ></div>;
<div  /*2*/ />`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"1", "2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "aria-label",
					Kind:   new(lsproto.CompletionItemKindField),
					Detail: new("(property) \"aria-label\": string"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "Label docs",
						},
					},
				},
				&lsproto.CompletionItem{
					Label:  "foo",
					Kind:   new(lsproto.CompletionItemKindField),
					Detail: new("(property) foo: string"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "Doc",
						},
					},
				},
			},
		},
	})
}

func TestCompletionsInJsxNamespacedIntrinsicTag(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @Filename: /a.tsx
declare React: any;
interface JsxElement {}
interface JsxIntrinsicElements {
    /** Element docs */
    "foo:bar": {
        /** Foo docs */
        foo: boolean
        /** Bar docs */
        bar: string
    }
}
<foo:bar /*1*/ />
<foo:bar  /*2*/></foo:bar>`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"1", "2"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:  "bar",
					Kind:   new(lsproto.CompletionItemKindField),
					Detail: new("(property) bar: string"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "Bar docs",
						},
					},
				},
				&lsproto.CompletionItem{
					Label:  "foo",
					Kind:   new(lsproto.CompletionItemKindField),
					Detail: new("(property) foo: boolean"),
					Documentation: &lsproto.StringOrMarkupContent{
						MarkupContent: &lsproto.MarkupContent{
							Kind:  lsproto.MarkupKindMarkdown,
							Value: "Foo docs",
						},
					},
				},
			},
		},
	})
}

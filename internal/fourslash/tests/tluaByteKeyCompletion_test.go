package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/ls"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

// A Lua string key is a byte string. A key that is not valid UTF-8 is labeled
// and inserted with `\xHH` escapes, which read back as the same bytes; that
// includes a key that begins with 0xFE, the internal-name byte. Inside a string
// literal the key is spelled with that literal's escapes, quote included.
func TestTluaByteKeyCompletion(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `interface I { "\xffa": number; "\xfeb": string; "it's": boolean; }
declare i: I;
i[|.|]/*keys*/
local s = i['[|/*str*/|]']`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "keys", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label:      `\xffa`,
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new(`["\xffa"]`),
					FilterText: new(`.\xffa`),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: `["\xffa"]`,
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
				&lsproto.CompletionItem{
					Label:      `\xfeb`,
					Kind:       new(lsproto.CompletionItemKindField),
					SortText:   new(string(ls.SortTextLocationPriority)),
					InsertText: new(`["\xfeb"]`),
					FilterText: new(`.\xfeb`),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: `["\xfeb"]`,
							Range:   f.Ranges()[0].LSRange,
						},
					},
				},
			},
		},
	})
	f.VerifyCompletions(t, "str", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				`\xffa`,
				`\xfeb`,
				// The escaped label still resolves to its property.
				&lsproto.CompletionItem{
					Label:  `it\'s`,
					Detail: new(`(property) I["it's"]: boolean`),
					TextEdit: &lsproto.TextEditOrInsertReplaceEdit{
						TextEdit: &lsproto.TextEdit{
							NewText: `it\'s`,
							Range:   f.Ranges()[1].LSRange,
						},
					},
				},
			},
		},
	})
}

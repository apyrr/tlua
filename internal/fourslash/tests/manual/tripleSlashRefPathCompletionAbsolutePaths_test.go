package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTripleSlashRefPathCompletionAbsolutePaths(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /tests/cases/fourslash/tests/test0.tlua
/// <reference path="/tests/cases/f/*0*/
// @Filename: /tests/cases/fourslash/tests/test1.tlua
/// <reference path="/tests/cases/fourslash/*1*/
// @Filename: /tests/cases/fourslash/tests/test2.tlua
/// <reference path="/tests/cases/fourslash//*2*/
// @Filename: /tests/cases/fourslash/f1.tlua
/*f1*/
// @Filename: /tests/cases/fourslash/f2.tsx
/*f2*/
// @Filename: /tests/cases/fourslash/folder/f1.tlua
/*subf1*/
// @Filename: /tests/cases/fourslash/f3.lua
/*f3*/
// @Filename: /tests/cases/fourslash/f4.jsx
/*f4*/
// @Filename: /tests/cases/fourslash/e1.tlua
/*e1*/
// @Filename: /tests/cases/fourslash/e2.lua
/*e2*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, []string{"0", "1"}, &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"fourslash",
			},
		},
	})
	f.VerifyCompletions(t, "2", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"e1.tlua",
				"f1.tlua",
				"f2.tsx",
				"folder",
				"tests",
			},
		},
	})
}

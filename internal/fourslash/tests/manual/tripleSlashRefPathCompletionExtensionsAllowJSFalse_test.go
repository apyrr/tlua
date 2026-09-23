package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestTripleSlashRefPathCompletionExtensionsAllowJSFalse(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: test0.tlua
/// <reference path="/*0*/
/// <reference path=".//*1*/
/// <reference path="./f/*2*/
// @Filename: f1.tlua

// @Filename: f1.lua

// @Filename: f1.d.tlua

// @Filename: f1.cs
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, f.Markers(), &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &[]string{},
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"f1.d.tlua",
				"f1.tlua",
			},
		},
	})
}

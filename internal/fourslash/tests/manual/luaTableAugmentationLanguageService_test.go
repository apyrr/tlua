package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

// Implicit globals and augmented table members are synthesized at checker
// initialization from the per-file candidate lists, so the language service
// must see one merged symbol across files: hover unions the arms, and
// references and rename reach every contributing assignment.
func TestLuaTableAugmentationQuickInfo(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
Config = {};
Config.debug/*member*/ = true;
Config.debug = 1;
// @Filename: /b.tlua
local d = Config/*global*/.debug;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "member", "(property) debug: number | boolean", "")
	f.VerifyQuickInfoAt(t, "global", "local Config: {\n    debug: number | boolean;\n}", "")
}

func TestLuaTableAugmentationFindAllReferences(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
Config = {};
Config.[|debug|] = true;
// @Filename: /b.tlua
local d = Config.[|debug|];`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t)
}

func TestLuaTableAugmentationRename(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
Config = {};
Config./*member*/debug = true;
// @Filename: /b.tlua
local d = Config.debug;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, "member")
}

// An augmented member declared in one file is offered as a member completion
// on the merged global from another file.
func TestLuaTableAugmentationCompletion(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
Config = {};
Config.debug = true;
// @Filename: /b.tlua
local d = Config./*complete*/`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "complete", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{
				&lsproto.CompletionItem{
					Label: "debug",
					Kind:  new(lsproto.CompletionItemKindField),
				},
			},
		},
	})
}

// The dot spelling and the string-key spelling of an augmented member share one
// symbol, so renaming at the declaration reaches both. The baseline records
// whether the `["item"]` spelling is included.
func TestLuaTableAugmentationRenameKeySpellings(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
Store = {};
Store./*member*/item = 1;
// @Filename: /b.tlua
local x = Store["item"];
// @Filename: /c.tlua
local y = Store.item;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, "member")
}

func TestLuaTableAugmentationGoToDefinition(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
G = {};
G.value = 1;
// @Filename: /b.tlua
local v = G./*use*/value;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, false /*includeOriginalSelectionRange*/, "use")
}

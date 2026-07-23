package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// The tags are completed from `type`'s declared return type, so the guard needs no
// language-service special case.
func TestLuaTypeGuardTagCompletions(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: esnext,luajit
// @strict: true
declare v: string | number;
local matched = type(v) == "/**/";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"boolean",
				"cdata",
				"function",
				"nil",
				"number",
				"string",
				"table",
				"thread",
				"userdata",
			},
		},
	})
}

func TestLuaTypeGuardNarrowedQuickInfo(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @lib: esnext,luajit
// @strict: true
function f(v: string | number)
  if type(v) == "string" then
    use(v/*narrowed*/);
  else
    use(v/*rest*/);
  end
end
declare function use(x: unknown): void;`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "narrowed", "(parameter) v: string", "")
	f.VerifyQuickInfoAt(t, "rest", "(parameter) v: number", "")
}

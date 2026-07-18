package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// A Lua label is a `::name::` statement and `goto name` is its only jump, so the
// label helpers that used to walk labeled statements and labeled break/continue
// must resolve through the new nodes: goto-definition on the goto, references
// and rename in the block that scopes the label.
func TestLuaLabelGoToDefinition(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
function f(n: number): number
    ::/*def*/top::
    n = n - 1;
    if n > 0 then goto /*use*/top end
    return n;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineGoToDefinition(t, false /*includeOriginalSelectionRange*/, "use")
}

func TestLuaLabelFindAllReferences(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
function f(flag: boolean): number
    ::[|top|]::
    if flag then goto [|top|] end
    do
        if flag then goto [|top|] end
    end
    return 0;
end

// A same-named label in another function is a different label.
function g(flag: boolean): number
    ::top::
    if flag then goto top end
    return 1;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineFindAllReferences(t)
}

func TestLuaLabelRename(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
function f(flag: boolean): number
    ::/*label*/top::
    if flag then goto top end
    return 0;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineRename(t, nil /*preferences*/, "label")
}

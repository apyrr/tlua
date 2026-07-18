package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// Break/continue bind to the innermost enclosing loop; occurrences highlight
// the owning loop's keywords and only its own jumps. (The upstream
// getOccurrencesLoopBreakContinue suite was deleted with the switch statement,
// which every one of its cases wrapped around the loops.)
func TestLuaLoopBreakContinueOccurrences(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `[|while|] true do
    if false then
        [|break|];
    end
    for i = 1, 10 do
        if i > 5 then
            break;
        end
        continue;
    end
    [|continue|];
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineDocumentHighlights(t, nil /*preferences*/, ToAny(f.Ranges())...)
}

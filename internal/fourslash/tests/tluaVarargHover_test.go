package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// A vararg parameter carries a synthetic `...` name so that it owns a symbol.
// Hover is the display choke point where that name would leak, and where the
// symbol's array type would be printed instead of the pack's element type.
func TestTluaVarargHover(t *testing.T) {
	fourslash.SkipIfFailing(t)
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
declare function /*1*/annotated(...: number): number;
declare function /*2*/bare(...): void;
declare function /*3*/withFixed(label: string, ...: string): string;

type /*4*/Callable = (...: boolean) => void;

// Each arm of a union of function types prints its own vararg. (This does NOT
// reach combineUnionOrIntersectionParameters -- hover shows the union, not a
// combined signature. The checker-minted rest, which must also print as ` + "`...`" + `
// rather than ` + "`...args: string[]`" + `, is locked by tluaVarargSignatureHelp and by
// the tluaVarargFunctionTypes compiler baseline.)
declare cond: boolean;
declare function fixedArity(a: number, b: boolean): void;
declare function varargArity(...: string): void;
local /*5*/combined = cond and fixedArity or varargArity;
`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyBaselineHover(t)
}

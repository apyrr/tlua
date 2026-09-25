package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestCompletionListInObjectLiteral4(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @strictNullChecks: true
interface Thing {
    hello: number;
    world: string;
}

declare function funcA(x : Thing): void;
declare function funcB(x?: Thing): void;
declare function funcC(x : Thing | nil): void;
declare function funcD(x : Thing | nil): void;
declare function funcE(x : Thing | nil): void;
declare function funcF(x?: Thing | nil): void;

funcA({ /*A*/ });
funcB({ /*B*/ });
funcC({ /*C*/ });
funcD({ /*D*/ });
funcE({ /*E*/ });
funcF({ /*F*/ });`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, f.Markers(), &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Exact: []fourslash.CompletionsExpectedItem{
				"hello",
				"world",
			},
		},
	})
}

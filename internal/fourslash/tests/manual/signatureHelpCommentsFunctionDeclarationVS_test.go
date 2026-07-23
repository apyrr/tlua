package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpCommentsFunctionDeclarationVS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `/** This comment should appear for foo*/
function foo() end
foo(/*4*/);
/** This is comment for function signature*/
function fooWithParameters(/** this is comment about a*/a: string,
    /** this is comment for b*/
    b: number) {
    local d = a;
}
fooWithParameters(/*10*/"a",/*11*/10);
/**
* Does something
* @param a a string
*/
declare function fn(a: string);
fn(/*12*/"hello");`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

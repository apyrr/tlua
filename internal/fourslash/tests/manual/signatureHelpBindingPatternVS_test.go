package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/lsp/lsproto"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSignatureHelpBindingPatternVS(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `
/**
 * @param options An empty object binding pattern.
 */
function emptyObj({}) end
emptyObj(/*emptyObj*/)

/**
 * @param items An empty array binding pattern.
 */
function emptyArr([]) end
emptyArr(/*emptyArr*/)

/**
 * @param param An object with a and b properties.
 */
function nonEmptyObj({a, b}: {a: number, b: string}) end
nonEmptyObj(/*nonEmptyObj*/)

/**
 * @param tuple A tuple with two elements.
 */
function nonEmptyArr([x, y]: [number, string]) end
nonEmptyArr(/*nonEmptyArr*/)

/**
 * @param first The first number parameter.
 * @param second An object with a and b properties.
 */
function idLeading(first: number, {a, b}: {a: number, b: string}) end
idLeading(123/*idLeading*/, { a: 1, b: 2 }/*bindingTrailing*/)

/**
 * @param first An object with a and b properties.
 * @param last The last number parameter.
 */
function bindingLeading({a, b}: {a: number, b: string}, last: number) end
bindingLeading(/*bindingLeading*/{ a: 1, b: 2 }, 123 /*idTrailing*/)

/**
 * @param param1 {Object} The first parameter
 * @param param1.a {number} Comment a
 * @param param1.b {string} Comment b
 * @param param2 {Object} The second parameter
 * @param param2.c {boolean} Comment c
 * @param param2.d {unknown} Comment d
 */
function multipleBindings({ a, b }, { c, d }) end
multipleBindings({ a: 0, b: "" }/*firstObjParam*/, { c: true, d: "" }/*secondObjParam*/)
`
	f, done := fourslash.NewFourslash(t, &lsproto.ClientCapabilities{VSSupportsVisualStudioExtensions: new(true)}, content)
	defer done()
	f.VerifyBaselineSignatureHelp(t)
}

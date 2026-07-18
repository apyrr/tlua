package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestQuickInfoJsxNamespacedIntrinsic(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @jsx: react
// @Filename: /a.tsx
declare React: any;
interface JsxElement {}
interface JsxIntrinsicElements {
    /** Element docs */
    "foo:bar": {
        /** Foo docs */
        foo: boolean
        /** Bar docs */
        bar: string
    }
}
<foo:ba/*tag*/r fo/*attr*/o />`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyQuickInfoAt(t, "tag", "(property) JsxIntrinsicElements[\"foo:bar\"]: {\n    foo: boolean;\n    bar: string;\n}", "Element docs")
	f.VerifyQuickInfoAt(t, "attr", "(property) foo: boolean", "Foo docs")
}

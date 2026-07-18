package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

func TestSemanticClassificationJSX(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tsx
local Component = () => <div>Hello</div>;
local afterJSX = 42;
local alsoAfterJSX = "test";`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/a.tsx")
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "function.declaration.local", Text: "Component"},
		{Type: "variable.declaration.local", Text: "afterJSX"},
		{Type: "variable.declaration.local", Text: "alsoAfterJSX"},
	})
}

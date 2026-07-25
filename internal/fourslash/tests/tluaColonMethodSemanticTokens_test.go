package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	"github.com/apyrr/tlua/internal/testutil"
)

// `function obj:f()` synthesizes its `self` parameter over the colon and the
// target -- source that precedes the parameter list. The token walk reaches that
// parameter after it has already emitted the name, so classifying it sends the
// encoder backwards, and encodeSemanticTokens panics rather than emitting a
// descending delta. That killed semantic highlighting for every file containing a
// colon method, so the implicit parameter contributes no tokens at all.
func TestTluaColonMethodSemanticTokens(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
local obj = {};
function obj:f(x: number): number
  return x;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/a.tlua")
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "variable.declaration.local", Text: "obj"},
		{Type: "variable.local", Text: "obj"},
		{Type: "function.declaration", Text: "f"},
		{Type: "parameter.declaration", Text: "x"},
		{Type: "parameter", Text: "x"},
	})
}

// The dotted target is the same hazard one level deeper: the synthesized type
// query mirrors the whole `a.b` chain, so skipping the parameter has to skip that
// subtree with it.
func TestTluaDottedColonMethodSemanticTokens(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
local a = { b = {} };
function a.b:f(x: number): number
  return x;
end`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/a.tlua")
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "variable.declaration.local", Text: "a"},
		{Type: "property.declaration", Text: "b"},
		{Type: "variable.local", Text: "a"},
		{Type: "property", Text: "b"},
		{Type: "function.declaration", Text: "f"},
		{Type: "parameter.declaration", Text: "x"},
		{Type: "parameter", Text: "x"},
	})
}

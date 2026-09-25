package fourslash_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/fourslash"
	. "github.com/apyrr/tlua/internal/fourslash/tests/util"
	"github.com/apyrr/tlua/internal/testutil"
)

// An optional colon call `b?:m()` stores its `?` in the access's optional-link
// slot, where `?.` otherwise sits, and its `:` in the colon slot. The language
// service reads both: it has to offer the receiver's methods after `?:`, elide the
// parameter the receiver fills from signature help as it does for `b:m()`, and
// describe and colour the method like any other.
const tluaOptionalColonCallContent = `// @Filename: /a.tlua
interface Box {
  value: number;
  read(self: Box): number;
  set(self: Box, v: number): nil;
}
declare function maybeBox(): Box | nil;
local b = maybeBox()
local r = b?:/*name*/re/*hover*/ad()
b?:set(/*arg*/)`

func TestTluaOptionalColonCallCompletions(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, tluaOptionalColonCallContent)
	defer done()
	f.VerifyCompletions(t, "name", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"read", "set"},
		},
	})
}

// A plain colon call completes the receiver's members the same way.
func TestTluaColonCallCompletions(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
interface Box {
  read(self: Box): number;
}
declare b: Box;
local r = b:/*name*/read()`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.VerifyCompletions(t, "name", &fourslash.CompletionsExpectedList{
		IsIncomplete: false,
		ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
			CommitCharacters: &DefaultCommitCharacters,
			EditRange:        Ignored,
		},
		Items: &fourslash.CompletionsExpectedItems{
			Includes: []fourslash.CompletionsExpectedItem{"read"},
		},
	})
}

// Completion has to work while the call is being typed, before any argument list
// exists: the parser builds an incomplete colon call where the line ends at the
// `:` or right after the name. Only members the call can name are offered --
// callable, and spelled as a Lua name -- and a receiver that may be nil still
// offers its members.
func TestTluaColonCallCompletionsWhileTyping(t *testing.T) {
	t.Parallel()
	cases := []struct{ name, source string }{
		{"colon ends the line", "local r = b:/*m*/\nlocal z = 1"},
		{"optional colon ends the line", "local r = c?:/*m*/\nlocal z = 1"},
		{"partial name", "local r = b:re/*m*/\nlocal z = 1"},
		{"end of file", "local r = b:/*m*/"},
		{"receiver may be nil", "local r = c:/*m*/read()"},
		{"inside a block", "local function f()\n  b:/*m*/\nend"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			defer testutil.RecoverAndFail(t, "Panic on fourslash test")
			content := "// @Filename: /a.tlua\ninterface Box {\n  value: number;\n  [\"a b\"]: (self: Box) => number;\n  read(self: Box): number;\n}\ndeclare b: Box;\ndeclare c: Box | nil;\n" + tc.source
			f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
			defer done()
			f.VerifyCompletions(t, "m", &fourslash.CompletionsExpectedList{
				IsIncomplete: false,
				ItemDefaults: &fourslash.CompletionsExpectedItemDefaults{
					CommitCharacters: &DefaultCommitCharacters,
					EditRange:        Ignored,
				},
				Items: &fourslash.CompletionsExpectedItems{
					Includes: []fourslash.CompletionsExpectedItem{"read"},
					Excludes: []string{"value", "a b"},
				},
			})
		})
	}
}

func TestTluaOptionalColonCallSignatureHelp(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, tluaOptionalColonCallContent)
	defer done()
	f.GoToMarker(t, "arg")
	f.VerifySignatureHelp(t, fourslash.VerifySignatureHelpOptions{
		Text:           "set(v: number): nil",
		ParameterCount: 1,
	})
}

func TestTluaOptionalColonCallHover(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, tluaOptionalColonCallContent)
	defer done()
	f.VerifyBaselineHover(t)
}

func TestTluaOptionalColonCallSemanticTokens(t *testing.T) {
	t.Parallel()
	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: /a.tlua
interface Box {
  read(self: Box): number;
}
declare function maybeBox(): Box | nil;
local b = maybeBox()
local r = b?:read()`
	f, done := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	defer done()
	f.GoToFile(t, "/a.tlua")
	f.VerifySemanticTokens(t, []fourslash.SemanticToken{
		{Type: "interface.declaration", Text: "Box"},
		{Type: "method.declaration", Text: "read"},
		{Type: "parameter.declaration", Text: "self"},
		{Type: "interface", Text: "Box"},
		{Type: "function.declaration", Text: "maybeBox"},
		{Type: "interface", Text: "Box"},
		{Type: "variable.declaration.local", Text: "b"},
		{Type: "function", Text: "maybeBox"},
		{Type: "variable.declaration.local", Text: "r"},
		{Type: "variable.local", Text: "b"},
		{Type: "method", Text: "read"},
	})
}

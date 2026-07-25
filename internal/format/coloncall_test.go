package format_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/format"
	"github.com/apyrr/tlua/internal/ls/lsutil"
	"github.com/apyrr/tlua/internal/parser"
	"gotest.tools/v3/assert"
)

func luaFormatContext(t *testing.T, spaceBeforeTypeAnnotation core.Tristate) context.Context {
	t.Helper()
	return format.WithFormatCodeSettings(t.Context(), lsutil.FormatCodeSettings{
		EditorSettings: lsutil.EditorSettings{
			TabSize:             4,
			IndentSize:          4,
			NewLineCharacter:    "\n",
			ConvertTabsToSpaces: core.TSTrue,
			IndentStyle:         lsutil.IndentStyleSmart,
		},
		InsertSpaceBeforeAndAfterBinaryOperators: core.TSTrue,
		InsertSpaceBeforeTypeAnnotation:          spaceBeforeTypeAnnotation,
	}, "\n")
}

func parseLua(t *testing.T, input string) *ast.SourceFile {
	t.Helper()
	return parser.ParseSourceFile(ast.SourceFileParseOptions{
		FileName: "/test.tlua",
		Path:     "/test.tlua",
	}, input, core.ScriptKindTS)
}

func formatLua(t *testing.T, input string, spaceBeforeTypeAnnotation core.Tristate) string {
	t.Helper()
	ctx := luaFormatContext(t, spaceBeforeTypeAnnotation)
	return applyBulkEdits(input, format.FormatDocument(ctx, parseLua(t, input)))
}

// The method colon binds a receiver to a name the way `.` does, so it takes no
// padding on either side. The type-annotation colon it shares a token kind with
// does, which is what pads it if nothing distinguishes the two: `obj:m()` comes
// back as `obj: m()`, and a declaration carries both colons at once.
func TestFormatLuaMethodColonStaysTight(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{"colon call keeps no space", "local isAdmin = LocalPlayer():IsAdmin()", "local isAdmin = LocalPlayer():IsAdmin()"},
		{"colon call loses trailing space", "local isAdmin = LocalPlayer(): IsAdmin()", "local isAdmin = LocalPlayer():IsAdmin()"},
		{"colon call loses surrounding space", "local isAdmin = LocalPlayer() : IsAdmin()", "local isAdmin = LocalPlayer():IsAdmin()"},
		{"chained colon calls stay tight", "local x = a:b():c():d()", "local x = a:b():c():d()"},
		{"method declaration stays tight", "function obj:m() end", "function obj:m() end"},
		{"method declaration loses space", "function obj : m() end", "function obj:m() end"},
		{"dotted method declaration", "function a.b.c:m() end", "function a.b.c:m() end"},
		// The return-type colon of the same declaration is a type annotation and
		// keeps its trailing space.
		{"return type still padded", "function obj:m(a: number): number return a end", "function obj:m(a: number): number return a end"},
		{"return type gains space", "function obj:m(a:number):number return a end", "function obj:m(a: number): number return a end"},
		// Colons the method rules must not reach.
		{"variable annotation untouched", "local x: number = 1", "local x: number = 1"},
		{"plain function annotations untouched", "function f(a: number): number return a end", "function f(a: number): number return a end"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.output, formatLua(t, tc.input, core.TSUnknown))
		})
	}
}

// insertSpaceBeforeTypeAnnotation pads the colon of an annotation. The method
// colon is not one, so the option must leave it alone. `function obj :m()` still
// parses -- the space is trivia -- but it reads as an annotation on `obj`, which
// is the confusion the tight spelling exists to avoid.
func TestFormatLuaMethodColonIgnoresTypeAnnotationOption(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  string
		output string
	}{
		{"method colon unpadded", "function obj:m(a: number): number return a end", "function obj:m(a : number) : number return a end"},
		{"method colon tightened", "function obj : m() end", "function obj:m() end"},
		{"colon call unpadded", "local isAdmin = LocalPlayer():IsAdmin()", "local isAdmin = LocalPlayer():IsAdmin()"},
		{"annotation padded", "local x: number = 1", "local x : number = 1"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.output, formatLua(t, tc.input, core.TSTrue))
		})
	}
}

// `function M:f()` synthesizes a `self` parameter over the colon, source that the
// formatter's forward token scan has already passed by the time it walks the
// parameter list. Descending into it rewound the scanner onto the colon, which
// re-emitted edits for tokens behind it -- inverted spans that corrupted the
// document rather than merely misformatting it.
func TestFormatLuaMethodDeclarationEditsAreWellFormed(t *testing.T) {
	t.Parallel()

	inputs := []string{
		"function obj:m() end",
		"function obj:m(a: number): number return a end",
		"function a.b.c:m(a: number) end",
		"function obj : m ( a : number ) : number return a end",
	}
	// Both settings, because padding the annotations is what makes the colon rules
	// fire hardest around a declaration's two colons -- and the text assertions
	// above cannot see an overlapping pair that happens to compose to the right
	// string.
	for _, spaceBeforeTypeAnnotation := range []core.Tristate{core.TSUnknown, core.TSTrue} {
		for _, input := range inputs {
			t.Run(fmt.Sprintf("%s (spaceBeforeTypeAnnotation=%v)", input, spaceBeforeTypeAnnotation), func(t *testing.T) {
				t.Parallel()
				ctx := luaFormatContext(t, spaceBeforeTypeAnnotation)

				end := 0
				for _, edit := range format.FormatDocument(ctx, parseLua(t, input)) {
					assert.Assert(t, edit.TextRange.Pos() <= edit.TextRange.End(), "edit %v is inverted", edit.TextRange)
					assert.Assert(t, edit.TextRange.Pos() >= end, "edit %v overlaps the previous edit", edit.TextRange)
					end = edit.TextRange.End()
				}
			})
		}
	}
}

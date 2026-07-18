package tstransforms_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/testutil/emittestutil"
	"github.com/apyrr/tlua/internal/testutil/parsetestutil"
	"github.com/apyrr/tlua/internal/transformers"
	"github.com/apyrr/tlua/internal/transformers/tstransforms"
)

func TestTypeEraser(t *testing.T) {
	t.Parallel()
	data := []struct {
		title  string
		input  string
		output string
		jsx    bool
	}{
		{title: "InterfaceDeclaration", input: "interface I { }", output: ""},
		{title: "TypeAliasDeclaration", input: "type T = U;", output: ""},
		{title: "ExpressionWithTypeArguments", input: "F<T>", output: "F;"},
		{title: "VariableDeclaration2", input: "local a: number", output: "local a;"},
		{title: "FunctionDeclaration1", input: "declare function f() {}", output: ""},
		{title: "FunctionDeclaration2", input: "function f();", output: ""},
		{title: "FunctionDeclaration3", input: "function f<T>(): U {}", output: "function f() { }"},
		{title: "FunctionExpression", input: "(function<T>(): U end)", output: "(function()\nend);"},
		{title: "ArrowFunction", input: "(<T>(): U => {})", output: "(function()\nend);"},
		{title: "ParameterDeclaration", input: "function f(a: number, b?: boolean) {}", output: "function f(a, b) { }"},
		{title: "CallExpression", input: "f<T>()", output: "f();"},
		{title: "NewExpression1", input: "new f<T>()", output: "new f();"},
		{title: "NewExpression2", input: "new f<T>", output: "new f;"},
		{title: "NonNullExpression", input: "x!", output: "x;"},
		{title: "TypeAssertionExpression#1", input: "<T>x", output: "x;"},
		{title: "TypeAssertionExpression#2", input: "(<T>x).c", output: "x.c;"},
		{title: "AsExpression#1", input: "x as T", output: "x;"},
		{title: "AsExpression#2", input: "(x as T).c", output: "x.c;"},
		{title: "SatisfiesExpression#1", input: "x satisfies T", output: "x;"},
		{title: "SatisfiesExpression#2", input: "(x satisfies T).c", output: "x.c;"},
		{title: "JsxSelfClosingElement", input: "<x<T> />", output: "<x />;", jsx: true},
		{title: "JsxOpeningElement", input: "<x<T>></x>", output: "<x></x>;", jsx: true},
	}

	for _, rec := range data {
		t.Run(rec.title, func(t *testing.T) {
			t.Parallel()
			file := parsetestutil.ParseTypeScript(rec.input, rec.jsx)
			parsetestutil.CheckDiagnostics(t, file)
			compilerOptions := &core.CompilerOptions{}
			emittestutil.CheckEmit(t, nil, tstransforms.NewTypeEraserTransformer(&transformers.TransformOptions{CompilerOptions: compilerOptions, Context: printer.NewEmitContext()}).TransformSourceFile(file), rec.output)
		})
	}
}

// A tagged template is a non-Lua construct: its lowered emit `f ""` (Lua string-call sugar)
// is not re-parseable tlua, so it can't live in the reparse-gated table above. It has its own
// test via CheckEmitJS (no reparse) to keep covering the type-eraser's TaggedTemplateExpression
// branch, which must erase the `<T>` type arguments.
func TestTypeEraserTaggedTemplate(t *testing.T) {
	t.Parallel()
	file := parsetestutil.ParseTypeScript("f<T>``", false)
	parsetestutil.CheckDiagnostics(t, file)
	compilerOptions := &core.CompilerOptions{}
	transformed := tstransforms.NewTypeEraserTransformer(&transformers.TransformOptions{CompilerOptions: compilerOptions, Context: printer.NewEmitContext()}).TransformSourceFile(file)
	emittestutil.CheckEmitJS(t, nil, transformed, "f \"\";")
}

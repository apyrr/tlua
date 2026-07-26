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
		{title: "ExpressionWithTypeArguments", input: "local _ = F<T>", output: "local _ = F;"},
		{title: "VariableDeclaration2", input: "local a: number", output: "local a;"},
		{title: "FunctionDeclaration1", input: "declare function f();", output: ""},
		{title: "FunctionDeclaration2", input: "function f();", output: ""},
		{title: "FunctionDeclaration3", input: "function f<T>(): U end", output: "function f()\nend"},
		{title: "FunctionExpression", input: "local _ = (function<T>(): U end)", output: "local _ = (function()\nend);"},
		{title: "ParameterDeclaration", input: "function f(a: number, b?: boolean) end", output: "function f(a, b)\nend"},
		{title: "CallExpression", input: "f<T>()", output: "f();"},
		{title: "NonNullExpression", input: "local _ = x!", output: "local _ = x;"},
		{title: "TypeAssertionExpression#1", input: "local _ = <T>x", output: "local _ = x;"},
		{title: "TypeAssertionExpression#2", input: "local _ = (<T>x).c", output: "local _ = x.c;"},
		{title: "AsExpression#1", input: "local _ = x as T", output: "local _ = x;"},
		{title: "AsExpression#2", input: "local _ = (x as T).c", output: "local _ = x.c;"},
		{title: "SatisfiesExpression#1", input: "local _ = x satisfies T", output: "local _ = x;"},
		{title: "SatisfiesExpression#2", input: "local _ = (x satisfies T).c", output: "local _ = x.c;"},
		{title: "JsxSelfClosingElement", input: "local _ = <x<T> />", output: "local _ = <x />;", jsx: true},
		{title: "JsxOpeningElement", input: "local _ = <x<T>></x>", output: "local _ = <x></x>;", jsx: true},
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

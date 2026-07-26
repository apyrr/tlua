package printer_test

import (
	"testing"

	"github.com/apyrr/tlua/internal/ast"
	"github.com/apyrr/tlua/internal/core"
	"github.com/apyrr/tlua/internal/printer"
	"github.com/apyrr/tlua/internal/testutil/emittestutil"
	"github.com/apyrr/tlua/internal/testutil/parsetestutil"
	"github.com/apyrr/tlua/internal/transformers"
	"github.com/apyrr/tlua/internal/transformers/tstransforms"
)

func TestEmit(t *testing.T) {
	t.Parallel()
	data := []struct {
		title  string
		input  string
		output string
		jsx    bool
	}{
		{title: "StringLiteral#1", input: `;local _ = "test"`, output: ";\nlocal _ = \"test\";"},
		// tlua: a control char emits as Lua `\xHH` (not JS `\uXXXX` nor octal `\ddd`) and
		// re-scans cleanly — the reparse gate proves the emitted escape round-trips.
		{title: "StringLiteral#3", input: "local _ = \"\\x1b\"", output: "local _ = \"\\x1b\";"},
		{title: "StringLiteral#2", input: `;local _ = 'test'`, output: ";\nlocal _ = 'test';"},
		{title: "NumericLiteral#1", input: `local _ = 0`, output: `local _ = 0;`},
		{title: "NumericLiteral#2", input: `local _ = 10_000`, output: `local _ = 10000;`},
		{title: "BooleanLiteral#1", input: `local _ = true`, output: `local _ = true;`},
		{title: "BooleanLiteral#2", input: `local _ = false`, output: `local _ = false;`},
		// tlua: backtick templates lower to Lua strings/concatenation (no `${}` interpolation).
		{title: "NoSubstitutionTemplateLiteral", input: "local _ = ``", output: "local _ = \"\";"},
		{title: "NoSubstitutionTemplateLiteral#2", input: "local _ = `\n`", output: "local _ = \"\\n\";"},

		{title: "RegularExpressionLiteral#1", input: `local _ = /a/`, output: `local _ = /a/;`},
		{title: "RegularExpressionLiteral#2", input: `local _ = /a/g`, output: `local _ = /a/g;`},
		// `null` is an accepted alias that canonicalizes to `nil`.
		{title: "NullLiteral", input: `local _ = null`, output: `local _ = nil;`},
		{title: "SuperExpression", input: `super()`, output: `super();`},
		{title: "PropertyAccess#1", input: `local _ = a.b`, output: `local _ = a.b;`},
		// `#` is the Lua length operator now, not a private-identifier sigil, so
		// `a.#b` no longer parses; the PrivateIdentifier node survives only for emit.
		{title: "PropertyAccess#3", input: `local _ = a?.b`, output: `local _ = a?.b;`},
		{title: "PropertyAccess#4", input: `local _ = a?.b.c`, output: `local _ = a?.b.c;`},
		{title: "PropertyAccess#5", input: `local _ = 1..b`, output: `local _ = 1..b;`},
		{title: "PropertyAccess#6", input: `local _ = 1.0.b`, output: `local _ = 1.0.b;`},
		{title: "PropertyAccess#7", input: `local _ = 0x1.b`, output: `local _ = 0x1.b;`},
		{title: "PropertyAccess#8", input: `local _ = 0b1.b`, output: `local _ = 0b1.b;`},
		{title: "PropertyAccess#9", input: `local _ = 0o1.b`, output: `local _ = 0o1.b;`},
		{title: "PropertyAccess#10", input: `local _ = 10e1.b`, output: `local _ = 10e1.b;`},
		{title: "PropertyAccess#11", input: `local _ = 10E1.b`, output: `local _ = 10E1.b;`},
		{title: "PropertyAccess#12", input: `local _ = a.b?.c`, output: `local _ = a.b?.c;`},
		{title: "PropertyAccess#13", input: "local _ = a\n.b", output: "local _ = a\n    .b;"},
		{title: "PropertyAccess#14", input: "local _ = a.\nb", output: "local _ = a.\n    b;"},
		{title: "ElementAccess#1", input: `local _ = a[b]`, output: `local _ = a[b];`},
		{title: "ElementAccess#2", input: `local _ = a?.[b]`, output: `local _ = a?.[b];`},
		{title: "ElementAccess#3", input: `local _ = a?.[b].c`, output: `local _ = a?.[b].c;`},
		{title: "CallExpression#1", input: `a()`, output: `a();`},
		{title: "CallExpression#2", input: `a<T>()`, output: `a<T>();`},
		{title: "CallExpression#3", input: `a(b)`, output: `a(b);`},
		{title: "CallExpression#4", input: `a<T>(b)`, output: `a<T>(b);`},
		{title: "CallExpression#5", input: `local _ = a(b).c`, output: `local _ = a(b).c;`},
		{title: "CallExpression#6", input: `local _ = a<T>(b).c`, output: `local _ = a<T>(b).c;`},
		{title: "CallExpression#7", input: `a?.(b)`, output: `a?.(b);`},
		{title: "CallExpression#8", input: `a?.<T>(b)`, output: `a?.<T>(b);`},
		{title: "CallExpression#9", input: `local _ = a?.(b).c`, output: `local _ = a?.(b).c;`},
		{title: "CallExpression#10", input: `local _ = a?.<T>(b).c`, output: `local _ = a?.<T>(b).c;`},
		{title: "CallExpression#11", input: `a<T, U>()`, output: `a<T, U>();`},
		// {title: "CallExpression#12", input: `a<T,>()`, output: `a<T,>();`}, // TODO: preserve trailing comma after Strada migration
		{title: "CallExpression#13", input: `a?.b()`, output: `a?.b();`},
		// Tagged templates are a non-Lua construct: the tag applied to the lowered
		// string emits as `tag ""` (Lua string-call sugar) which tlua does not parse.
		{title: "TypeAssertionExpression#1", input: `local _ = <T>a`, output: `local _ = <T>a;`},
		{title: "FunctionExpression#1", input: "local _ = (function() end)", output: "local _ = (function()\nend);"},
		{title: "FunctionExpression#2", input: "local _ = (function(a) return a end)", output: "local _ = (function(a)\n    return a;\nend);"},
		{title: "FunctionExpression#4", input: "local _ = (async function() end)", output: "local _ = (async function()\nend);"},
		{title: "FunctionExpression#6", input: "local _ = (function<T>() end)", output: "local _ = (function<T>()\nend);"},
		{title: "FunctionExpression#7", input: "local _ = (function(a) end)", output: "local _ = (function(a)\nend);"},
		{title: "FunctionExpression#8", input: "local _ = (function(): T end)", output: "local _ = (function(): T\nend);"},
		{title: "PrefixUnaryExpression#1", input: `local _ = +a`, output: `local _ = +a;`},
		{title: "PrefixUnaryExpression#3", input: `local _ = + +a`, output: `local _ = + +a;`},
		{title: "PrefixUnaryExpression#5", input: `local _ = -a`, output: `local _ = -a;`},
		{title: "PrefixUnaryExpression#7", input: `local _ = - -a`, output: `local _ = - -a;`},
		{title: "PrefixUnaryExpression#9", input: `local _ = +-a`, output: `local _ = +-a;`},
		{title: "PrefixUnaryExpression#11", input: `local _ = -+a`, output: `local _ = -+a;`},
		{title: "PrefixUnaryExpression#14", input: `local _ = !a`, output: `local _ = !a;`},
		// Lua length operator.
		{title: "PrefixUnaryExpression#len", input: `local _ = #a`, output: `local _ = #a;`},
		{title: "PrefixUnaryExpression#len2", input: `local _ = #a + #b`, output: `local _ = #a + #b;`},
		// `#` before `!` keeps a space, else the emitted `#!a` would re-scan as a shebang.
		{title: "PrefixUnaryExpression#lenNot", input: `local _ = # !a`, output: `local _ = # !a;`},
		// The comma operator has no bare-statement spelling, and `local _ = a, b` is a
		// value list (an ExpressionList node), so the source parens are what keep this a
		// comma BinaryExpression. They are a real node and round-trip verbatim.
		{title: "BinaryExpression#1", input: `local _ = (a,b)`, output: `local _ = (a, b);`},
		{title: "BinaryExpression#2", input: `local _ = a+b`, output: `local _ = a + b;`},
		{title: "BinaryExpression#3", input: `local _ = a^b`, output: `local _ = a ^ b;`},
		// Lua concatenation is right-associative: `a .. b .. c` needs no parens, but a
		// left-nested `(a .. b) .. c` keeps them.
		{title: "BinaryExpression#concat", input: `local _ = a..b`, output: `local _ = a .. b;`},
		{title: "BinaryExpression#concatChain", input: `local _ = a..b..c`, output: `local _ = a .. b .. c;`},
		{title: "BinaryExpression#concatLeftParen", input: `local _ = (a..b)..c`, output: `local _ = (a .. b) .. c;`},
		{title: "BinaryExpression#5", input: `local _ = a in b`, output: `local _ = a in b;`},
		// `&&`/`||` are aliases of `and`/`or` and print with the canonical Lua
		// spelling, whichever the source used.
		{title: "BinaryExpression#6", input: "local _ = a\n&& b", output: "local _ = a\n    and b;"},
		{title: "BinaryExpression#7", input: "local _ = a &&\nb", output: "local _ = a and\n    b;"},
		{title: "BinaryExpression#8", input: "local _ = a and b", output: "local _ = a and b;"},
		{title: "BinaryExpression#9", input: "local _ = a or b", output: "local _ = a or b;"},
		// `not` is spelled `!` on the way out: the token kind is shared with the
		// non-null and definite-assignment `!`, which must stay punctuation.
		{title: "PrefixUnaryExpression#not", input: "local _ = not a", output: "local _ = !a;"},
		// tlua has no conditional expression in source; the node survives only in
		// emit-time synthesis (`?.`/`??` lowering), covered by the factory-built
		// TestParenthesizeConditional tests below.
		{title: "TemplateExpression#1", input: "local _ = `a${b}c`", output: `local _ = ("a" .. tostring(b) .. "c");`},
		{title: "TemplateExpression#2", input: "local _ = `a${b}c${d}e`", output: `local _ = ("a" .. tostring(b) .. "c" .. tostring(d) .. "e");`},
		{title: "VarargExpression", input: `f(...)`, output: `f(...);`},
		{title: "ExpressionWithTypeArguments", input: `local _ = a<T>`, output: `local _ = a<T>;`},
		{title: "AsExpression", input: `local _ = a as T`, output: `local _ = a as T;`},
		{title: "SatisfiesExpression", input: `local _ = a satisfies T`, output: `local _ = a satisfies T;`},
		{title: "NonNullExpression", input: `local _ = a!`, output: `local _ = a!;`},
		{title: "MetaProperty#1", input: `local _ = new.target`, output: `local _ = new.target;`},
		{title: "ObjectLiteralExpression#1", input: `local _ = ({})`, output: `local _ = ({});`},
		{title: "ObjectLiteralExpression#2", input: `local _ = ({a,})`, output: `local _ = ({ a, });`},
		{title: "PropertyAssignment", input: "local _ = ({a = b})", output: "local _ = ({ a = b });"},
		{title: "PropertyAssignment#2", input: "local _ = ({[a] = b})", output: "local _ = ({ [a] = b });"},
		{title: "VariableStatement#1", input: `local a`, output: `local a;`},
		{title: "VariableStatement#3", input: `local a = b`, output: `local a = b;`},
		{title: "EmptyStatement", input: `;`, output: `;`},
		{title: "ContinueStatement", input: `continue`, output: "continue;"},
		{title: "BreakStatement", input: `break`, output: "break;"},
		{title: "ReturnStatement#1", input: `return`, output: "return;"},
		{title: "ReturnStatement#2", input: `return a`, output: "return a;"},
		{title: "LabelStatement", input: `::a::`, output: "::a::"},
		{title: "GotoStatement", input: `::a:: goto a`, output: "::a::\ngoto a;"},
		{title: "FunctionDeclaration#2", input: `function f() end`, output: "function f()\nend"},
		{title: "FunctionDeclaration#4", input: `async function f() end`, output: "async function f()\nend"},
		{title: "FunctionDeclaration#6", input: `function f<T>() end`, output: "function f<T>()\nend"},
		{title: "FunctionDeclaration#7", input: `function f(a) end`, output: "function f(a)\nend"},
		{title: "FunctionDeclaration#8", input: `function f():T end`, output: "function f(): T\nend"},
		{title: "FunctionDeclaration#9", input: `function f();`, output: `function f();`},
		{title: "InterfaceDeclaration#1", input: `interface a {}`, output: "interface a {\n}"},
		{title: "InterfaceDeclaration#2", input: `interface a<T>{}`, output: "interface a<T> {\n}"},
		{title: "InterfaceDeclaration#3", input: `interface a extends b {}`, output: "interface a extends b {\n}"},
		{title: "InterfaceDeclaration#4", input: `interface a extends b, c {}`, output: "interface a extends b, c {\n}"},
		{title: "TypeAliasDeclaration#1", input: `type a = b`, output: "type a = b;"},
		{title: "TypeAliasDeclaration#2", input: `type a<T> = b`, output: "type a<T> = b;"},
		{title: "ModuleDeclaration#7", input: `local _ = global;`, output: "local _ = global;"},
		{title: "ModuleDeclaration#8", input: `global{}`, output: "global { }"},
		{title: "KeywordTypeNode#1", input: `type T = any`, output: `type T = any;`},
		{title: "KeywordTypeNode#2", input: `type T = unknown`, output: `type T = unknown;`},
		{title: "KeywordTypeNode#3", input: `type T = never`, output: `type T = never;`},
		{title: "KeywordTypeNode#4", input: `type T = void`, output: `type T = void;`},
		{title: "KeywordTypeNode#5", input: `type T = undefined`, output: `type T = nil;`},
		{title: "KeywordTypeNode#6", input: `type T = null`, output: `type T = nil;`},
		{title: "KeywordTypeNode#7", input: `type T = object`, output: `type T = object;`},
		{title: "KeywordTypeNode#8", input: `type T = string`, output: `type T = string;`},
		{title: "KeywordTypeNode#9", input: `type T = symbol`, output: `type T = symbol;`},
		{title: "KeywordTypeNode#10", input: `type T = number`, output: `type T = number;`},
		{title: "KeywordTypeNode#11", input: `type T = bigint`, output: `type T = bigint;`},
		{title: "KeywordTypeNode#12", input: `type T = boolean`, output: `type T = boolean;`},
		{title: "KeywordTypeNode#13", input: `type T = intrinsic`, output: `type T = intrinsic;`},
		{title: "TypePredicateNode#1", input: `function f(): asserts a;`, output: `function f(): asserts a;`},
		{title: "TypePredicateNode#2", input: `function f(): asserts a is b;`, output: `function f(): asserts a is b;`},
		{title: "TypeReferenceNode#1", input: `type T = a`, output: `type T = a;`},
		{title: "TypeReferenceNode#2", input: `type T = a.b`, output: `type T = a.b;`},
		{title: "TypeReferenceNode#3", input: `type T = a<U>`, output: `type T = a<U>;`},
		{title: "TypeReferenceNode#4", input: `type T = a.b<U>`, output: `type T = a.b<U>;`},
		{title: "FunctionTypeNode#1", input: `type T = () => a`, output: `type T = () => a;`},
		{title: "FunctionTypeNode#2", input: `type T = <T>() => a`, output: `type T = <T>() => a;`},
		{title: "FunctionTypeNode#3", input: `type T = (a) => b`, output: `type T = (a) => b;`},
		{title: "TypeQueryNode#1", input: `type T = typeof a`, output: `type T = typeof a;`},
		{title: "TypeQueryNode#2", input: `type T = typeof a.b`, output: `type T = typeof a.b;`},
		{title: "TypeQueryNode#3", input: `type T = typeof a<U>`, output: `type T = typeof a<U>;`},
		{title: "TypeLiteralNode#1", input: `type T = {}`, output: `type T = {};`},
		{title: "TypeLiteralNode#2", input: `type T = {a}`, output: "type T = {\n    a;\n};"},
		{title: "ArrayTypeNode", input: `type T = a[]`, output: "type T = a[];"},
		{title: "UnionTypeNode#1", input: `type T = a | b`, output: "type T = a | b;"},
		{title: "UnionTypeNode#2", input: `type T = a | b | c`, output: "type T = a | b | c;"},
		{title: "UnionTypeNode#3", input: `type T = | a | b`, output: "type T = a | b;"},
		{title: "IntersectionTypeNode#1", input: `type T = a & b`, output: "type T = a & b;"},
		{title: "IntersectionTypeNode#2", input: `type T = a & b & c`, output: "type T = a & b & c;"},
		{title: "IntersectionTypeNode#3", input: `type T = & a & b`, output: "type T = a & b;"},
		{title: "ConditionalTypeNode", input: `type T = a extends b ? c : d`, output: "type T = a extends b ? c : d;"},
		{title: "InferTypeNode#1", input: `type T = a extends infer b ? c : d`, output: "type T = a extends infer b ? c : d;"},
		{title: "InferTypeNode#2", input: `type T = a extends infer b extends c ? d : e`, output: "type T = a extends infer b extends c ? d : e;"},
		{title: "ParenthesizedTypeNode", input: `type T = (U)`, output: "type T = (U);"},
		{title: "SelfKeywordType", input: `interface T { value: self }`, output: "interface T {\n    value: self;\n}"},
		{title: "TypeOperatorNode#1", input: `type T = keyof U`, output: "type T = keyof U;"},
		{title: "TypeOperatorNode#2", input: `type T = readonly U[]`, output: "type T = readonly U[];"},
		{title: "TypeOperatorNode#3", input: `type T = unique symbol`, output: "type T = unique symbol;"},
		{title: "IndexedAccessTypeNode", input: `type T = a[b]`, output: "type T = a[b];"},
		{title: "MappedTypeNode#1", input: `type T = { [a in b]: c }`, output: "type T = {\n    [a in b]: c;\n};"},
		{title: "MappedTypeNode#2", input: `type T = { [a in b as c]: d }`, output: "type T = {\n    [a in b as c]: d;\n};"},
		{title: "MappedTypeNode#3", input: `type T = { readonly [a in b]: c }`, output: "type T = {\n    readonly [a in b]: c;\n};"},
		{title: "MappedTypeNode#4", input: `type T = { +readonly [a in b]: c }`, output: "type T = {\n    +readonly [a in b]: c;\n};"},
		{title: "MappedTypeNode#5", input: `type T = { -readonly [a in b]: c }`, output: "type T = {\n    -readonly [a in b]: c;\n};"},
		{title: "MappedTypeNode#6", input: `type T = { [a in b]?: c }`, output: "type T = {\n    [a in b]?: c;\n};"},
		{title: "MappedTypeNode#7", input: `type T = { [a in b]+?: c }`, output: "type T = {\n    [a in b]+?: c;\n};"},
		{title: "MappedTypeNode#8", input: `type T = { [a in b]-?: c }`, output: "type T = {\n    [a in b]-?: c;\n};"},
		{title: "MappedTypeNode#9", input: `type T = { [a in b]: c; d }`, output: "type T = {\n    [a in b]: c;\n    d;\n};"},
		{title: "LiteralTypeNode#1", input: `type T = null`, output: "type T = nil;"},
		{title: "LiteralTypeNode#2", input: `type T = true`, output: "type T = true;"},
		{title: "LiteralTypeNode#3", input: `type T = false`, output: "type T = false;"},
		{title: "LiteralTypeNode#4", input: `type T = ""`, output: "type T = \"\";"},
		{title: "LiteralTypeNode#5", input: "type T = ''", output: "type T = '';"},
		// tlua: a no-substitution backtick literal type collapses to an equivalent string literal type.
		{title: "LiteralTypeNode#6", input: "type T = ``", output: "type T = \"\";"},
		{title: "LiteralTypeNode#7", input: `type T = 0`, output: "type T = 0;"},
		{title: "LiteralTypeNode#9", input: `type T = -0`, output: "type T = -0;"},
		{title: "TemplateTypeNode#1", input: "type T = `a${b}c`", output: "type T = `a${b}c`;"},
		{title: "TemplateTypeNode#2", input: "type T = `a${b}c${d}e`", output: "type T = `a${b}c${d}e`;"},
		{title: "ImportTypeNode#1", input: `type T = import(a)`, output: "type T = import(a);"},
		{title: "ImportTypeNode#2", input: `type T = import(a).b`, output: "type T = import(a).b;"},
		{title: "ImportTypeNode#3", input: `type T = import(a).b<U>`, output: "type T = import(a).b<U>;"},
		{title: "ImportTypeNode#4", input: `type T = typeof import(a)`, output: "type T = typeof import(a);"},
		{title: "ImportTypeNode#5", input: `type T = typeof import(a).b`, output: "type T = typeof import(a).b;"},
		{title: "PropertySignature#1", input: "interface I {a}", output: "interface I {\n    a;\n}"},
		{title: "PropertySignature#2", input: "interface I {readonly a}", output: "interface I {\n    readonly a;\n}"},
		{title: "PropertySignature#3", input: "interface I {\"a\"}", output: "interface I {\n    \"a\";\n}"},
		{title: "PropertySignature#4", input: "interface I {'a'}", output: "interface I {\n    'a';\n}"},
		{title: "PropertySignature#5", input: "interface I {0}", output: "interface I {\n    0;\n}"},
		{title: "PropertySignature#7", input: "interface I {[a]}", output: "interface I {\n    [a];\n}"},
		{title: "PropertySignature#8", input: "interface I {a?}", output: "interface I {\n    a?;\n}"},
		{title: "PropertySignature#9", input: "interface I {a: b}", output: "interface I {\n    a: b;\n}"},
		{title: "MethodSignature#1", input: "interface I {a()}", output: "interface I {\n    a();\n}"},
		{title: "MethodSignature#2", input: "interface I {\"a\"()}", output: "interface I {\n    \"a\"();\n}"},
		{title: "MethodSignature#3", input: "interface I {'a'()}", output: "interface I {\n    'a'();\n}"},
		{title: "MethodSignature#4", input: "interface I {0()}", output: "interface I {\n    0();\n}"},
		{title: "MethodSignature#6", input: "interface I {[a]()}", output: "interface I {\n    [a]();\n}"},
		{title: "MethodSignature#7", input: "interface I {a?()}", output: "interface I {\n    a?();\n}"},
		{title: "MethodSignature#8", input: "interface I {a<T>()}", output: "interface I {\n    a<T>();\n}"},
		{title: "MethodSignature#9", input: "interface I {a(): b}", output: "interface I {\n    a(): b;\n}"},
		{title: "MethodSignature#10", input: "interface I {a(b): c}", output: "interface I {\n    a(b): c;\n}"},
		{title: "CallSignature#1", input: "interface I {()}", output: "interface I {\n    ();\n}"},
		{title: "CallSignature#2", input: "interface I {():a}", output: "interface I {\n    (): a;\n}"},
		{title: "CallSignature#3", input: "interface I {(p)}", output: "interface I {\n    (p);\n}"},
		{title: "CallSignature#4", input: "interface I {<T>()}", output: "interface I {\n    <T>();\n}"},
		{title: "IndexSignatureDeclaration#1", input: "interface I {[a]}", output: "interface I {\n    [a];\n}"},
		{title: "IndexSignatureDeclaration#2", input: "interface I {[a: b]}", output: "interface I {\n    [a: b];\n}"},
		{title: "IndexSignatureDeclaration#3", input: "interface I {[a: b]: c}", output: "interface I {\n    [a: b]: c;\n}"},
		{title: "ParameterDeclaration#1", input: "function f(a);", output: "function f(a);"},
		{title: "ParameterDeclaration#2", input: "function f(a: b);", output: "function f(a: b);"},
		{title: "ParameterDeclaration#3", input: "function f(a = b);", output: "function f(a = b);"},
		{title: "ParameterDeclaration#4", input: "function f(a?);", output: "function f(a?);"},
		{title: "ParameterDeclaration#5", input: "function f(...);", output: "function f(...);"},
		{title: "ParameterDeclaration#6", input: "function f(...: a);", output: "function f(...: a);"},
		// {title: "ParameterDeclaration#7", input: "function f(a,)", output: "function f(a,);"}, // TODO: preserve trailing comma after Strada migration
		{title: "ObjectBindingPattern#1", input: "function f({});", output: "function f({});"},
		{title: "ObjectBindingPattern#2", input: "function f({a});", output: "function f({ a });"},
		{title: "ObjectBindingPattern#3", input: "function f({a = b});", output: "function f({ a = b });"},
		{title: "ObjectBindingPattern#4", input: "function f({a: b});", output: "function f({ a: b });"},
		{title: "ObjectBindingPattern#5", input: "function f({a: b = c});", output: "function f({ a: b = c });"},
		{title: "ObjectBindingPattern#6", input: "function f({\"a\": b});", output: "function f({ \"a\": b });"},
		{title: "ObjectBindingPattern#7", input: "function f({'a': b});", output: "function f({ 'a': b });"},
		{title: "ObjectBindingPattern#8", input: "function f({0: b});", output: "function f({ 0: b });"},
		{title: "ObjectBindingPattern#9", input: "function f({[a]: b});", output: "function f({ [a]: b });"},
		{title: "ObjectBindingPattern#11", input: "function f({a: {}});", output: "function f({ a: {} });"},
		{title: "TypeParameterDeclaration#1", input: "function f<T>();", output: "function f<T>();"},
		{title: "TypeParameterDeclaration#2", input: "function f<in T>();", output: "function f<in T>();"},
		{title: "TypeParameterDeclaration#3", input: "function f<T extends U>();", output: "function f<T extends U>();"},
		{title: "TypeParameterDeclaration#4", input: "function f<T = U>();", output: "function f<T = U>();"},
		{title: "TypeParameterDeclaration#5", input: "function f<T extends U = V>();", output: "function f<T extends U = V>();"},
		{title: "TypeParameterDeclaration#6", input: "function f<T, U>();", output: "function f<T, U>();"},
		// {title: "TypeParameterDeclaration#7", input: "function f<T,>();", output: "function f<T,>();"}, // TODO: preserve trailing comma after Strada migration
		{title: "JsxElement1", input: "local _ = <a></a>", output: "local _ = <a></a>;", jsx: true},
		{title: "JsxElement2", input: "local _ = <this></this>", output: "local _ = <this></this>;", jsx: true},
		{title: "JsxElement3", input: "local _ = <a:b></a:b>", output: "local _ = <a:b></a:b>;", jsx: true},
		{title: "JsxElement4", input: "local _ = <a.b></a.b>", output: "local _ = <a.b></a.b>;", jsx: true},
		{title: "JsxElement5", input: "local _ = <a<b>></a>", output: "local _ = <a<b>></a>;", jsx: true},
		{title: "JsxElement6", input: "local _ = <a b></a>", output: "local _ = <a b></a>;", jsx: true},
		{title: "JsxElement7", input: "local _ = <a>b</a>", output: "local _ = <a>b</a>;", jsx: true},
		{title: "JsxElement8", input: "local _ = <a>{b}</a>", output: "local _ = <a>{b}</a>;", jsx: true},
		{title: "JsxElement9", input: "local _ = <a><b></b></a>", output: "local _ = <a><b></b></a>;", jsx: true},
		{title: "JsxElement10", input: "local _ = <a><b /></a>", output: "local _ = <a><b /></a>;", jsx: true},
		{title: "JsxElement11", input: "local _ = <a><></></a>", output: "local _ = <a><></></a>;", jsx: true},
		{title: "JsxElement12", input: "local _ = <a>\n    {/* missing */}\n    {\n        // foo\n    }\n</a>", output: "local _ = <a>\n    {--[[ missing ]]}\n    {\n    -- foo\n    }\n</a>;", jsx: true},
		{title: "JsxSelfClosingElement1", input: "local _ = <a />", output: "local _ = <a />;", jsx: true},
		{title: "JsxSelfClosingElement2", input: "local _ = <this />", output: "local _ = <this />;", jsx: true},
		{title: "JsxSelfClosingElement3", input: "local _ = <a:b />", output: "local _ = <a:b />;", jsx: true},
		{title: "JsxSelfClosingElement4", input: "local _ = <a.b />", output: "local _ = <a.b />;", jsx: true},
		{title: "JsxSelfClosingElement5", input: "local _ = <a<b> />", output: "local _ = <a<b> />;", jsx: true},
		{title: "JsxSelfClosingElement6", input: "local _ = <a b/>", output: "local _ = <a b/>;", jsx: true},
		{title: "JsxFragment1", input: "local _ = <></>", output: "local _ = <></>;", jsx: true},
		{title: "JsxFragment2", input: "local _ = <>b</>", output: "local _ = <>b</>;", jsx: true},
		{title: "JsxFragment3", input: "local _ = <>{b}</>", output: "local _ = <>{b}</>;", jsx: true},
		{title: "JsxFragment4", input: "local _ = <><b></b></>", output: "local _ = <><b></b></>;", jsx: true},
		{title: "JsxFragment5", input: "local _ = <><b /></>", output: "local _ = <><b /></>;", jsx: true},
		{title: "JsxFragment6", input: "local _ = <><></></>", output: "local _ = <><></></>;", jsx: true},
		{title: "JsxAttribute1", input: "local _ = <a b/>", output: "local _ = <a b/>;", jsx: true},
		{title: "JsxAttribute2", input: "local _ = <a b:c/>", output: "local _ = <a b:c/>;", jsx: true},
		{title: "JsxAttribute3", input: "local _ = <a b=\"c\"/>", output: "local _ = <a b=\"c\"/>;", jsx: true},
		{title: "JsxAttribute4", input: "local _ = <a b='c'/>", output: "local _ = <a b='c'/>;", jsx: true},
		{title: "JsxAttribute5", input: "local _ = <a b={c}/>", output: "local _ = <a b={c}/>;", jsx: true},
		{title: "JsxAttribute6", input: "local _ = <a b=<c></c>/>", output: "local _ = <a b=<c></c>/>;", jsx: true},
		{title: "JsxAttribute7", input: "local _ = <a b=<c />/>", output: "local _ = <a b=<c />/>;", jsx: true},
		{title: "JsxAttribute8", input: "local _ = <a b=<></>/>", output: "local _ = <a b=<></>/>;", jsx: true},
	}

	for _, rec := range data {
		t.Run(rec.title, func(t *testing.T) {
			t.Parallel()
			file := parsetestutil.ParseTypeScript(rec.input, rec.jsx)
			parsetestutil.CheckDiagnostics(t, file)
			emittestutil.CheckEmit(t, nil, file, rec.output)
		})
	}
}

// localDeclOf wraps an expression as the initializer of `local _ = <expr>`. Lua
// admits only a call or an assignment as a statement, so a synthesized
// expression under test needs a declaration to sit in for the bare expression
// statement these tests used to build.
func localDeclOf(factory *ast.NodeFactory, expression *ast.Expression) *ast.Node {
	return factory.NewVariableStatement(
		nil, /*modifiers*/
		factory.NewVariableDeclarationList(
			factory.NewNodeList([]*ast.Node{
				factory.NewVariableDeclaration(
					factory.NewIdentifier("_"),
					nil, /*exclamationToken*/
					nil, /*type*/
					expression,
				),
			}),
			ast.NodeFlagsLuaLocal,
		),
	)
}

func TestParenthesizePropertyAccess1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewPropertyAccessExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					nil, /*questionDotToken*/
					nil, /*colonToken*/
					factory.NewIdentifier("c"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b).c;")
}

func TestParenthesizePropertyAccess2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewPropertyAccessExpression(
					// will be parenthesized on emit:
					factory.NewPropertyAccessExpression(
						factory.NewIdentifier("a"),
						factory.NewToken(ast.KindQuestionDotToken),
						nil, /*colonToken*/
						factory.NewIdentifier("b"),
						ast.NodeFlagsOptionalChain,
					),
					nil, /*questionDotToken*/
					nil, /*colonToken*/
					factory.NewIdentifier("c"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a?.b).c;")
}

func TestParenthesizeElementAccess1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewElementAccessExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					nil, /*questionDotToken*/
					factory.NewIdentifier("c"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b)[c];")
}

func TestParenthesizeElementAccess2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewElementAccessExpression(
					// will be parenthesized on emit:
					factory.NewPropertyAccessExpression(
						factory.NewIdentifier("a"),
						factory.NewToken(ast.KindQuestionDotToken),
						nil, /*colonToken*/
						factory.NewIdentifier("b"),
						ast.NodeFlagsOptionalChain,
					),
					nil, /*questionDotToken*/
					factory.NewIdentifier("c"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a?.b)[c];")
}

func TestParenthesizeCall1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewCallExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					nil, /*questionDotToken*/
					nil, /*typeArguments*/
					factory.NewNodeList([]*ast.Node{}),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b)();")
}

func TestParenthesizeCall2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewCallExpression(
					// will be parenthesized on emit:
					factory.NewPropertyAccessExpression(
						factory.NewIdentifier("a"),
						factory.NewToken(ast.KindQuestionDotToken),
						nil, /*colonToken*/
						factory.NewIdentifier("b"),
						ast.NodeFlagsOptionalChain,
					),
					nil, /*questionDotToken*/
					nil, /*typeArguments*/
					factory.NewNodeList([]*ast.Node{}),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a?.b)();")
}

func TestParenthesizeCall4(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewCallExpression(
					factory.NewIdentifier("a"),
					nil, /*questionDotToken*/
					nil, /*typeArguments*/
					factory.NewNodeList([]*ast.Node{
						factory.NewBinaryExpression(
							nil, /*modifiers*/
							factory.NewIdentifier("b"),
							nil, /*typeNode*/
							factory.NewToken(ast.KindCommaToken),
							factory.NewIdentifier("c"),
						),
					}),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "a((b, c));")
}

func TestParenthesizeTypeAssertion1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewTypeAssertion(
					factory.NewTypeReferenceNode(
						factory.NewIdentifier("T"),
						nil, /*typeArguments*/
					),
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindPlusToken),
						factory.NewIdentifier("b"),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = <T>(a + b);")
}

func TestParenthesizeArrowFunction1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewArrowFunction(
					nil, /*modifiers*/
					nil, /*typeParameters*/
					factory.NewNodeList([]*ast.Node{}),
					nil, /*returnType*/
					nil, /*fullSignature*/
					factory.NewToken(ast.KindEqualsGreaterThanToken),
					factory.NewObjectLiteralExpression(
						factory.NewNodeList([]*ast.Node{}),
						false, /*multiLine*/
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The parens under test are statement-position parens: the lowered arrow prints as a
	// `function`-initial expression. Initializer position drops them, so this stays a bare
	// expression statement, whose printed form cannot reparse clean (TLUA100057).
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(function() return {} end);")
}

func TestParenthesizeArrowFunction2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewArrowFunction(
					nil, /*modifiers*/
					nil, /*typeParameters*/
					factory.NewNodeList([]*ast.Node{}),
					nil, /*returnType*/
					nil, /*fullSignature*/
					factory.NewToken(ast.KindEqualsGreaterThanToken),
					factory.NewPropertyAccessExpression(
						factory.NewObjectLiteralExpression(
							factory.NewNodeList([]*ast.Node{}),
							false, /*multiLine*/
						),
						nil, /*questionDotToken*/
						nil, /*colonToken*/
						factory.NewIdentifier("a"),
						ast.NodeFlagsNone,
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The parens under test are statement-position parens: the lowered arrow prints as a
	// `function`-initial expression. Initializer position drops them, so this stays a bare
	// expression statement, whose printed form cannot reparse clean (TLUA100057).
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(function() return {}.a end);")
}

func isBinaryOperator(token ast.Kind) bool {
	switch token {
	case ast.KindCommaToken,
		ast.KindLessThanToken,
		ast.KindGreaterThanToken,
		ast.KindLessThanEqualsToken,
		ast.KindGreaterThanEqualsToken,
		ast.KindEqualsEqualsToken,
		ast.KindTildeEqualsToken,
		ast.KindPlusToken,
		ast.KindMinusToken,
		ast.KindAsteriskToken,
		ast.KindAsteriskAsteriskToken,
		ast.KindSlashToken,
		ast.KindPercentToken,
		ast.KindAmpersandAmpersandToken,
		ast.KindBarBarToken,
		ast.KindEqualsToken,
		ast.KindInKeyword:
		return true
	}
	return false
}

func makeSide(label string, kind ast.Kind, factory *ast.NodeFactory) *ast.Node {
	switch {
	case kind == ast.KindIdentifier || kind == ast.KindUnknown:
		return factory.NewIdentifier(label)
	case kind == ast.KindArrowFunction:
		return factory.NewArrowFunction(
			nil, /*modifiers*/
			nil, /*typeParameters*/
			factory.NewNodeList([]*ast.Node{}),
			nil, /*returnType*/
			nil, /*fullSignature*/
			factory.NewToken(ast.KindEqualsGreaterThanToken),
			factory.NewBlock(factory.NewNodeList([]*ast.Node{}), false /*multiLine*/),
		)
	case isBinaryOperator(kind):
		return factory.NewBinaryExpression(
			nil, /*modifiers*/
			factory.NewIdentifier(label+"l"),
			nil, /*typeNode*/
			factory.NewToken(kind),
			factory.NewIdentifier(label+"r"),
		)
	default:
		panic("unsupported kind")
	}
}

func TestParenthesizeBinary(t *testing.T) {
	t.Parallel()

	data := []struct {
		left     ast.Kind
		operator ast.Kind
		right    ast.Kind
		output   string
		// bare keeps the expression in an expression statement rather than as the
		// initializer of `local _`. An initializer parenthesizes a comma operand
		// (OperatorPrecedenceDisallowComma), which would swap the rule under test
		// for the disallowed-comma rule; a bare statement keeps the operator at
		// lowest precedence, at the cost of the reparse gate (see CheckEmitJS).
		bare bool
	}{
		{operator: ast.KindCommaToken, output: "l, r", bare: true},
		{operator: ast.KindCommaToken, left: ast.KindPlusToken, output: "ll + lr, r", bare: true},
		{operator: ast.KindAsteriskToken, left: ast.KindPlusToken, output: "(ll + lr) * r"},
		{operator: ast.KindAsteriskToken, right: ast.KindPlusToken, output: "l * (rl + rr)"},
		{operator: ast.KindPlusToken, left: ast.KindAsteriskToken, output: "ll * lr + r"},
		{operator: ast.KindPlusToken, right: ast.KindAsteriskToken, output: "l + rl * rr"},
		{operator: ast.KindSlashToken, left: ast.KindAsteriskToken, output: "ll * lr / r"},
		{operator: ast.KindSlashToken, left: ast.KindAsteriskAsteriskToken, output: "ll ^ lr / r"},
		{operator: ast.KindAsteriskAsteriskToken, left: ast.KindAsteriskToken, output: "(ll * lr) ^ r"},
		{operator: ast.KindAsteriskAsteriskToken, left: ast.KindAsteriskAsteriskToken, output: "(ll ^ lr) ^ r"},
		{operator: ast.KindAsteriskToken, right: ast.KindAsteriskToken, output: "l * rl * rr"},
		{operator: ast.KindAmpersandAmpersandToken, right: ast.KindArrowFunction, output: "l and (function()\nend)"},
	}
	for _, rec := range data {
		t.Run(rec.output, func(t *testing.T) {
			t.Parallel()

			var factory ast.NodeFactory
			expression := factory.NewBinaryExpression(
				nil, /*modifiers*/
				makeSide("l", rec.left, &factory),
				nil, /*typeNode*/
				factory.NewToken(rec.operator),
				makeSide("r", rec.right, &factory),
			)
			statement := localDeclOf(&factory, expression)
			expected := "local _ = " + rec.output + ";"
			if rec.bare {
				statement = factory.NewExpressionStatement(expression)
				expected = rec.output + ";"
			}
			file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
				[]*ast.Node{statement},
			), factory.NewToken(ast.KindEndOfFile))

			parsetestutil.MarkSyntheticRecursive(file)
			if rec.bare {
				emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), expected)
			} else {
				emittestutil.CheckEmit(t, nil, file.AsSourceFile(), expected)
			}
		})
	}
}

func TestParenthesizeSpreadElement2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewCallExpression(
					factory.NewIdentifier("a"),
					nil, /*questionDotToken*/
					nil, /*typeArguments*/
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewSpreadElement(
								// will be parenthesized on emit:
								factory.NewBinaryExpression(
									nil, /*modifiers*/
									factory.NewIdentifier("b"),
									nil, /*typeNode*/
									factory.NewToken(ast.KindCommaToken),
									factory.NewIdentifier("c"),
								),
							),
						},
					),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// Spread is emit-only JS syntax: tlua reads `...` as the vararg, which is not
	// a prefixexp and so cannot take the `(b, c)` suffix. No reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "a(...(b, c));")
}

func TestParenthesizeExpressionWithTypeArguments(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewExpressionWithTypeArguments(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewTypeReferenceNode(
								factory.NewIdentifier("c"),
								nil,
							),
						},
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b)<c>;")
}

func TestParenthesizeAsExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewAsExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					factory.NewTypeReferenceNode(
						factory.NewIdentifier("c"),
						nil,
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b) as c;")
}

func TestParenthesizeSatisfiesExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewSatisfiesExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					factory.NewTypeReferenceNode(
						factory.NewIdentifier("c"),
						nil,
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b) satisfies c;")
}

func TestParenthesizeNonNullExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			localDeclOf(&factory,
				factory.NewNonNullExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "local _ = (a, b)!;")
}

func TestParenthesizeExpressionStatement1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewObjectLiteralExpression(
					factory.NewNodeList(
						[]*ast.Node{},
					),
					false, /*multiLine*/
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// Bare expression statements only arise from error recovery or synthesis, so the printed form cannot reparse clean (TLUA100057).
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "({});")
}

func TestParenthesizeExpressionStatement2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	fn := factory.NewFunctionExpression(
		nil, /*modifiers*/
		nil, /*name*/
		nil, /*typeParameters*/
		factory.NewNodeList(
			[]*ast.Node{},
		),
		nil, /*returnType*/
		nil, /*fullSignature*/
		factory.NewBlock(
			factory.NewNodeList([]*ast.Node{}),
			false, /*multiLine*/
		),
	)
	// Function expressions always print an `end`-terminated Lua body; the printer
	// no longer gates on a body flag, so a flag-less synthesized node prints valid
	// Lua too.
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(fn),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// Bare expression statements only arise from error recovery or synthesis, so the printed form cannot reparse clean (TLUA100057).
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(function()\nend);")
}

func TestParenthesizeArrayType(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewArrayTypeNode(
					// will be parenthesized on emit:
					factory.NewUnionTypeNode(
						factory.NewNodeList(
							[]*ast.Node{
								factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
								factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
							},
						),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = (a | b)[];")
}

func TestParenthesizeUnionType1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewUnionTypeNode(
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
							// will be parenthesized on emit:
							factory.NewFunctionTypeNode(
								nil, /*modifiers*/
								nil, /*typeParameters*/
								factory.NewNodeList(
									[]*ast.Node{},
								),
								factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
							),
						},
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = a | (() => b);")
}

func TestParenthesizeUnionType2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewUnionTypeNode(
					factory.NewNodeList(
						[]*ast.Node{
							// will be parenthesized on emit:
							factory.NewInferTypeNode(
								factory.NewTypeParameterDeclaration(
									nil,
									nil, /*dotDotDotToken*/
									factory.NewIdentifier("a"),
									factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
									nil, /*expression*/
									nil, /*defaultType*/
								),
							),
							factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
						},
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = (infer a extends b) | c;")
}

func TestParenthesizeIntersectionType(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewIntersectionTypeNode(
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
							// will be parenthesized on emit:
							factory.NewUnionTypeNode(
								factory.NewNodeList(
									[]*ast.Node{
										factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
										factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
									},
								),
							),
						},
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = a & (b | c);")
}

func TestParenthesizeReadonlyTypeOperator1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewTypeOperatorNode(
					ast.KindReadonlyKeyword,
					// will be parenthesized on emit:
					factory.NewUnionTypeNode(
						factory.NewNodeList(
							[]*ast.Node{
								factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
								factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
							},
						),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = readonly (a | b);")
}

func TestParenthesizeReadonlyTypeOperator2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewTypeOperatorNode(
					ast.KindReadonlyKeyword,
					// will be parenthesized on emit:
					factory.NewTypeOperatorNode(
						ast.KindKeyOfKeyword,
						factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = readonly (keyof a);")
}

func TestParenthesizeKeyofTypeOperator(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewTypeOperatorNode(
					ast.KindKeyOfKeyword,
					// will be parenthesized on emit:
					factory.NewUnionTypeNode(
						factory.NewNodeList(
							[]*ast.Node{
								factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
								factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
							},
						),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = keyof (a | b);")
}

func TestParenthesizeIndexedAccessType(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewIndexedAccessTypeNode(
					// will be parenthesized on emit:
					factory.NewUnionTypeNode(
						factory.NewNodeList(
							[]*ast.Node{
								factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
								factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
							},
						),
					),
					factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = (a | b)[c];")
}

func TestParenthesizeConditionalType1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewConditionalTypeNode(
					// will be parenthesized on emit:
					factory.NewFunctionTypeNode(
						nil, /*modifiers*/
						nil, /*typeParameters*/
						factory.NewNodeList(
							[]*ast.Node{},
						),
						factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
					),
					factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
					factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
					factory.NewTypeReferenceNode(factory.NewIdentifier("d"), nil /*typeArguments*/),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = (() => a) extends b ? c : d;")
}

func TestParenthesizeConditionalType2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewConditionalTypeNode(
					factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
					// will be parenthesized on emit:
					factory.NewConditionalTypeNode(
						factory.NewTypeReferenceNode(factory.NewIdentifier("b"), nil /*typeArguments*/),
						factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
						factory.NewTypeReferenceNode(factory.NewIdentifier("d"), nil /*typeArguments*/),
						factory.NewTypeReferenceNode(factory.NewIdentifier("e"), nil /*typeArguments*/),
					),
					factory.NewTypeReferenceNode(factory.NewIdentifier("f"), nil /*typeArguments*/),
					factory.NewTypeReferenceNode(factory.NewIdentifier("g"), nil /*typeArguments*/),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = a extends (b extends c ? d : e) ? f : g;")
}

func TestParenthesizeConditionalType3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewConditionalTypeNode(
					factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
					factory.NewFunctionTypeNode(
						nil, /*modifiers*/
						nil, /*typeParameters*/
						factory.NewNodeList(
							[]*ast.Node{},
						),
						// will be parenthesized on emit:
						factory.NewInferTypeNode(
							factory.NewTypeParameterDeclaration(
								nil,
								nil, /*dotDotDotToken*/
								factory.NewIdentifier("b"),
								factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
								nil, /*expression*/
								nil, /*defaultType*/
							),
						),
					),
					factory.NewTypeReferenceNode(factory.NewIdentifier("d"), nil /*typeArguments*/),
					factory.NewTypeReferenceNode(factory.NewIdentifier("e"), nil /*typeArguments*/),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = a extends () => (infer b extends c) ? d : e;")
}

func TestParenthesizeConditionalType4(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList([]*ast.Node{
		factory.NewTypeAliasDeclaration(
			nil,                        /*modifiers*/
			factory.NewIdentifier("_"), /*name*/
			nil,                        /*typeParameters*/
			factory.NewConditionalTypeNode(
				factory.NewTypeReferenceNode(factory.NewIdentifier("a"), nil /*typeArguments*/),
				factory.NewFunctionTypeNode(
					nil, /*modifiers*/
					nil, /*typeParameters*/
					factory.NewNodeList(
						[]*ast.Node{},
					),
					// will be parenthesized on emit:
					factory.NewUnionTypeNode(
						factory.NewNodeList(
							[]*ast.Node{
								factory.NewInferTypeNode(
									factory.NewTypeParameterDeclaration(
										nil,
										nil, /*dotDotDotToken*/
										factory.NewIdentifier("b"),
										factory.NewTypeReferenceNode(factory.NewIdentifier("c"), nil /*typeArguments*/),
										nil, /*expression*/
										nil, /*defaultType*/
									),
								),
								factory.NewTypeReferenceNode(factory.NewIdentifier("d"), nil /*typeArguments*/),
							},
						),
					),
				),
				factory.NewTypeReferenceNode(factory.NewIdentifier("e"), nil /*typeArguments*/),
				factory.NewTypeReferenceNode(factory.NewIdentifier("f"), nil /*typeArguments*/),
			),
		),
	}), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = a extends () => (infer b extends c) | d ? e : f;")
}

func TestNameGeneration(t *testing.T) {
	t.Parallel()
	ec := printer.NewEmitContext()
	file := ec.Factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", ec.Factory.NewNodeList([]*ast.Node{
		ec.Factory.NewVariableStatement(nil, ec.Factory.NewVariableDeclarationList(
			ec.Factory.NewNodeList([]*ast.Node{
				ec.Factory.NewVariableDeclaration(ec.Factory.NewTempVariable(), nil, nil, nil),
			}),
			ast.NodeFlagsNone,
		)),
		ec.Factory.NewFunctionDeclaration(
			nil,
			nil,
			nil,
			ec.Factory.NewIdentifier("f"),
			nil,
			ec.Factory.NewNodeList([]*ast.Node{}),
			nil,
			nil,
			ec.Factory.NewBlock(ec.Factory.NewNodeList([]*ast.Node{
				ec.Factory.NewVariableStatement(nil, ec.Factory.NewVariableDeclarationList(
					ec.Factory.NewNodeList([]*ast.Node{
						ec.Factory.NewVariableDeclaration(ec.Factory.NewTempVariable(), nil, nil, nil),
					}),
					ast.NodeFlagsNone,
				)),
			}), true),
		),
	}), ec.Factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, ec, file.AsSourceFile(), "local _a;\nfunction f()\n    local _a;\nend")
}

func TestPartiallyEmittedExpression(t *testing.T) {
	t.Parallel()

	compilerOptions := &core.CompilerOptions{}

	file := parsetestutil.ParseTypeScript(`return ((container.parent
    .left as PropertyAccessExpression)
    .expression as PropertyAccessExpression)
    .expression;`, false /*jsx*/)

	emitContext := printer.NewEmitContext()
	file = tstransforms.NewTypeEraserTransformer(&transformers.TransformOptions{CompilerOptions: compilerOptions, Context: emitContext}).TransformSourceFile(file)
	emittestutil.CheckEmit(t, emitContext, file.AsSourceFile(), `return container.parent
    .left
    .expression
    .expression;`)
}

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
		// removed marks a case this round-trip test can't run: either the syntax
		// tlua removed from source (so the input no longer parses as .tlua), or a
		// non-Lua construct whose best-effort Lua emit is intentionally not
		// re-parseable (e.g. a tagged template lowers to call-sugar tlua doesn't
		// accept). Emit of these node kinds — which survive for JSON, declaration,
		// and best-effort value emit — is covered by the compiler baselines
		// instead. The flag is per-entry (not a title match) so renaming or adding
		// a case can never silently skip surviving syntax.
		removed bool
	}{
		{title: "StringLiteral#1", input: `;"test"`, output: ";\n\"test\";"},
		// tlua: a control char emits as Lua `\xHH` (not JS `\uXXXX` nor octal `\ddd`) and
		// re-scans cleanly — the reparse gate proves the emitted escape round-trips.
		{title: "StringLiteral#3", input: "\"\\x1b\"", output: "\"\\x1b\";"},
		{title: "StringLiteral#2", input: `;'test'`, output: ";\n'test';"},
		{title: "NumericLiteral#1", input: `0`, output: `0;`},
		{title: "NumericLiteral#2", input: `10_000`, output: `10000;`},
		{title: "BooleanLiteral#1", input: `true`, output: `true;`},
		{title: "BooleanLiteral#2", input: `false`, output: `false;`},
		// tlua: backtick templates lower to Lua strings/concatenation (no `${}` interpolation).
		{title: "NoSubstitutionTemplateLiteral", input: "``", output: "\"\";"},
		{title: "NoSubstitutionTemplateLiteral#2", input: "`\n`", output: "\"\\n\";"},

		{title: "RegularExpressionLiteral#1", input: `/a/`, output: `/a/;`},
		{title: "RegularExpressionLiteral#2", input: `/a/g`, output: `/a/g;`},
		// `null` is an accepted alias that canonicalizes to `nil`.
		{title: "NullLiteral", input: `null`, output: `nil;`},
		{title: "SuperExpression", input: `super()`, output: `super();`},
		{title: "PropertyAccess#1", input: `a.b`, output: `a.b;`},
		// `#` is the Lua length operator now, not a private-identifier sigil, so
		// `a.#b` no longer parses; the PrivateIdentifier node survives only for emit.
		{title: "PropertyAccess#2", input: `a.#b`, output: `a.#b;`, removed: true},
		{title: "PropertyAccess#3", input: `a?.b`, output: `a?.b;`},
		{title: "PropertyAccess#4", input: `a?.b.c`, output: `a?.b.c;`},
		{title: "PropertyAccess#5", input: `1..b`, output: `1..b;`},
		{title: "PropertyAccess#6", input: `1.0.b`, output: `1.0.b;`},
		{title: "PropertyAccess#7", input: `0x1.b`, output: `0x1.b;`},
		{title: "PropertyAccess#8", input: `0b1.b`, output: `0b1.b;`},
		{title: "PropertyAccess#9", input: `0o1.b`, output: `0o1.b;`},
		{title: "PropertyAccess#10", input: `10e1.b`, output: `10e1.b;`},
		{title: "PropertyAccess#11", input: `10E1.b`, output: `10E1.b;`},
		{title: "PropertyAccess#12", input: `a.b?.c`, output: `a.b?.c;`},
		{title: "PropertyAccess#13", input: "a\n.b", output: "a\n    .b;"},
		{title: "PropertyAccess#14", input: "a.\nb", output: "a.\n    b;"},
		{title: "ElementAccess#1", input: `a[b]`, output: `a[b];`},
		{title: "ElementAccess#2", input: `a?.[b]`, output: `a?.[b];`},
		{title: "ElementAccess#3", input: `a?.[b].c`, output: `a?.[b].c;`},
		{title: "CallExpression#1", input: `a()`, output: `a();`},
		{title: "CallExpression#2", input: `a<T>()`, output: `a<T>();`},
		{title: "CallExpression#3", input: `a(b)`, output: `a(b);`},
		{title: "CallExpression#4", input: `a<T>(b)`, output: `a<T>(b);`},
		{title: "CallExpression#5", input: `a(b).c`, output: `a(b).c;`},
		{title: "CallExpression#6", input: `a<T>(b).c`, output: `a<T>(b).c;`},
		{title: "CallExpression#7", input: `a?.(b)`, output: `a?.(b);`},
		{title: "CallExpression#8", input: `a?.<T>(b)`, output: `a?.<T>(b);`},
		{title: "CallExpression#9", input: `a?.(b).c`, output: `a?.(b).c;`},
		{title: "CallExpression#10", input: `a?.<T>(b).c`, output: `a?.<T>(b).c;`},
		{title: "CallExpression#11", input: `a<T, U>()`, output: `a<T, U>();`},
		// {title: "CallExpression#12", input: `a<T,>()`, output: `a<T,>();`}, // TODO: preserve trailing comma after Strada migration
		{title: "CallExpression#13", input: `a?.b()`, output: `a?.b();`},
		{title: "NewExpression#1", input: `new a`, output: `new a;`},
		{title: "NewExpression#2", input: `new a.b`, output: `new a.b;`},
		{title: "NewExpression#3", input: `new a()`, output: `new a();`},
		{title: "NewExpression#4", input: `new a.b()`, output: `new a.b();`},
		{title: "NewExpression#5", input: `new a<T>()`, output: `new a<T>();`},
		{title: "NewExpression#6", input: `new a.b<T>()`, output: `new a.b<T>();`},
		{title: "NewExpression#7", input: `new a(b)`, output: `new a(b);`},
		{title: "NewExpression#8", input: `new a.b(c)`, output: `new a.b(c);`},
		{title: "NewExpression#9", input: `new a<T>(b)`, output: `new a<T>(b);`},
		{title: "NewExpression#10", input: `new a.b<T>(c)`, output: `new a.b<T>(c);`},
		{title: "NewExpression#11", input: `new a(b).c`, output: `new a(b).c;`},
		{title: "NewExpression#12", input: `new a<T>(b).c`, output: `new a<T>(b).c;`},
		// Tagged templates are a non-Lua construct: the tag applied to the lowered
		// string emits as `tag ""` (Lua string-call sugar) which tlua does not parse.
		{title: "TaggedTemplateExpression#1", input: "tag``", output: "tag \"\";", removed: true},
		{title: "TaggedTemplateExpression#2", input: "tag<T>``", output: "tag<T> \"\";", removed: true},
		{title: "TypeAssertionExpression#1", input: `<T>a`, output: `<T>a;`},
		{title: "FunctionExpression#1", input: "(function() end)", output: "(function()\nend);"},
		{title: "FunctionExpression#2", input: "(function(a) return a end)", output: "(function(a)\n    return a;\nend);"},
		{title: "FunctionExpression#4", input: "(async function() end)", output: "(async function()\nend);"},
		{title: "FunctionExpression#6", input: "(function<T>() end)", output: "(function<T>()\nend);"},
		{title: "FunctionExpression#7", input: "(function(a) end)", output: "(function(a)\nend);"},
		{title: "FunctionExpression#8", input: "(function(): T end)", output: "(function(): T\nend);"},
		{title: "ArrowFunction#1", input: `a=>{}`, output: "(function(a)\nend);"},
		{title: "ArrowFunction#2", input: `()=>{}`, output: "(function()\nend);"},
		{title: "ArrowFunction#3", input: `(a)=>{}`, output: "(function(a)\nend);"},
		{title: "ArrowFunction#4", input: `<T>(a)=>{}`, output: "(function<T>(a)\nend);"},
		{title: "ArrowFunction#5", input: `async a=>{}`, output: "(async function(a)\nend);"},
		{title: "ArrowFunction#6", input: `async()=>{}`, output: "(async function()\nend);"},
		{title: "ArrowFunction#7", input: `async<T>()=>{}`, output: "(async function<T>()\nend);"},
		{title: "ArrowFunction#8", input: `():T=>{}`, output: "(function(): T\nend);"},
		{title: "ArrowFunction#9", input: `()=>a`, output: `(function() return a end);`},
		{title: "DeleteExpression", input: `delete a`, output: `delete a;`},
		{title: "VoidExpression", input: `void a`, output: `void a;`},
		{title: "PrefixUnaryExpression#1", input: `+a`, output: `+a;`},
		{title: "PrefixUnaryExpression#3", input: `+ +a`, output: `+ +a;`},
		{title: "PrefixUnaryExpression#5", input: `-a`, output: `-a;`},
		{title: "PrefixUnaryExpression#7", input: `- -a`, output: `- -a;`},
		{title: "PrefixUnaryExpression#9", input: `+-a`, output: `+-a;`},
		{title: "PrefixUnaryExpression#11", input: `-+a`, output: `-+a;`},
		{title: "PrefixUnaryExpression#14", input: `!a`, output: `!a;`},
		// Lua length operator.
		{title: "PrefixUnaryExpression#len", input: `#a`, output: `#a;`},
		{title: "PrefixUnaryExpression#len2", input: `#a + #b`, output: `#a + #b;`},
		// `#` before `!` keeps a space, else the emitted `#!a` would re-scan as a shebang.
		{title: "PrefixUnaryExpression#lenNot", input: `# !a`, output: `# !a;`},
		{title: "BinaryExpression#1", input: `a,b`, output: `a, b;`},
		{title: "BinaryExpression#2", input: `a+b`, output: `a + b;`},
		{title: "BinaryExpression#3", input: `a^b`, output: `a ^ b;`},
		// Lua concatenation is right-associative: `a .. b .. c` needs no parens, but a
		// left-nested `(a .. b) .. c` keeps them.
		{title: "BinaryExpression#concat", input: `a..b`, output: `a .. b;`},
		{title: "BinaryExpression#concatChain", input: `a..b..c`, output: `a .. b .. c;`},
		{title: "BinaryExpression#concatLeftParen", input: `(a..b)..c`, output: `(a .. b) .. c;`},
		{title: "BinaryExpression#4", input: `a instanceof b`, output: `a instanceof b;`},
		{title: "BinaryExpression#5", input: `a in b`, output: `a in b;`},
		// `&&`/`||` are aliases of `and`/`or` and print with the canonical Lua
		// spelling, whichever the source used.
		{title: "BinaryExpression#6", input: "a\n&& b", output: "a\n    and b;"},
		{title: "BinaryExpression#7", input: "a &&\nb", output: "a and\n    b;"},
		{title: "BinaryExpression#8", input: "a and b", output: "a and b;"},
		{title: "BinaryExpression#9", input: "a or b", output: "a or b;"},
		// `not` is spelled `!` on the way out: the token kind is shared with the
		// non-null and definite-assignment `!`, which must stay punctuation.
		{title: "PrefixUnaryExpression#not", input: "not a", output: "!a;"},
		// tlua has no conditional expression in source; the node survives only in
		// emit-time synthesis (`?.`/`??` lowering), covered by the factory-built
		// TestParenthesizeConditional tests below.
		{title: "ConditionalExpression#1", input: `a?b:c`, output: `a ? b : c;`, removed: true},
		{title: "ConditionalExpression#2", input: "a\n?b:c", output: "a\n    ? b : c;", removed: true},
		{title: "ConditionalExpression#3", input: "a?\nb:c", output: "a ?\n    b : c;", removed: true},
		{title: "ConditionalExpression#4", input: "a?b\n:c", output: "a ? b\n    : c;", removed: true},
		{title: "ConditionalExpression#5", input: "a?b:\nc", output: "a ? b :\n    c;", removed: true},
		{title: "TemplateExpression#1", input: "`a${b}c`", output: `("a" .. tostring(b) .. "c");`},
		{title: "TemplateExpression#2", input: "`a${b}c${d}e`", output: `("a" .. tostring(b) .. "c" .. tostring(d) .. "e");`},
		{title: "SpreadElement", input: `[...a]`, output: `[...a];`, removed: true},
		{title: "VarargExpression", input: `f(...)`, output: `f(...);`},
		{title: "OmittedExpression", input: `[,]`, output: `[,];`, removed: true},
		{title: "ExpressionWithTypeArguments", input: `a<T>`, output: `a<T>;`},
		{title: "AsExpression", input: `a as T`, output: `a as T;`},
		{title: "SatisfiesExpression", input: `a satisfies T`, output: `a satisfies T;`},
		{title: "NonNullExpression", input: `a!`, output: `a!;`},
		{title: "MetaProperty#1", input: `new.target`, output: `new.target;`},
		{title: "ArrayLiteralExpression#1", input: `[]`, output: `[];`, removed: true},
		{title: "ArrayLiteralExpression#2", input: `[a]`, output: `[a];`, removed: true},
		{title: "ArrayLiteralExpression#3", input: `[a,]`, output: `[a,];`, removed: true},
		{title: "ArrayLiteralExpression#4", input: `[,a]`, output: `[, a];`, removed: true},
		{title: "ArrayLiteralExpression#5", input: `[...a]`, output: `[...a];`, removed: true},
		{title: "ArrayLiteralExpression#6", input: `local array = [/* comment */];`, output: `local array = [ /* comment */];`, removed: true},
		{title: "ObjectLiteralExpression#1", input: `({})`, output: `({});`},
		{title: "ObjectLiteralExpression#2", input: `({a,})`, output: `({ a, });`},
		{title: "ShorthandPropertyAssignment", input: `({a})`, output: `({ a });`},
		{title: "PropertyAssignment", input: "({a = b})", output: "({ a = b });"},
		{title: "PropertyAssignment#2", input: "({[a] = b})", output: "({ [a] = b });"},
		{title: "SpreadAssignment", input: `({...a})`, output: `({ ...a });`, removed: true},
		{title: "Block", input: `{}`, output: `{ }`},
		{title: "VariableStatement#1", input: `local a`, output: `local a;`},
		{title: "VariableStatement#3", input: `local a = b`, output: `local a = b;`},
		{title: "EmptyStatement", input: `;`, output: `;`},
		{title: "IfStatement#1", input: `if(a);`, output: "if (a)\n    ;"},
		{title: "IfStatement#2", input: `if(a);else;`, output: "if (a)\n    ;\nelse\n    ;"},
		{title: "IfStatement#3", input: `if(a);else{}`, output: "if (a)\n    ;\nelse { }"},
		{title: "IfStatement#4", input: `if(a);else if(b);`, output: "if (a)\n    ;\nelse if (b)\n    ;"},
		{title: "IfStatement#5", input: `if(a);else if(b) {}`, output: "if (a)\n    ;\nelse if (b) { }"},
		{title: "IfStatement#6", input: `if(a) {}`, output: "if (a) { }"},
		{title: "IfStatement#7", input: `if(a) {} else;`, output: "if (a) { }\nelse\n    ;"},
		{title: "IfStatement#8", input: `if(a) {} else {}`, output: "if (a) { }\nelse { }"},
		{title: "IfStatement#9", input: `if(a) {} else if(b);`, output: "if (a) { }\nelse if (b)\n    ;"},
		{title: "IfStatement#10", input: `if(a) {} else if(b){}`, output: "if (a) { }\nelse if (b) { }"},
		{title: "DoStatement#2", input: `do {} while(a);`, output: "do { } while (a);"},
		{title: "WhileStatement#1", input: `while(a);`, output: "while (a)\n    ;"},
		{title: "WhileStatement#2", input: `while(a) {}`, output: "while (a) { }"},
		{title: "ForOfStatement#1", input: `for(a of b);`, output: "for (a of b)\n    ;", removed: true},
		{title: "ForOfStatement#2", input: `for(local a of b);`, output: "for (local a of b)\n    ;", removed: true},
		{title: "ForOfStatement#3", input: `for(a of b){}`, output: "for (a of b) { }", removed: true},
		{title: "ContinueStatement", input: `continue`, output: "continue;"},
		{title: "BreakStatement", input: `break`, output: "break;"},
		{title: "ReturnStatement#1", input: `return`, output: "return;"},
		{title: "ReturnStatement#2", input: `return a`, output: "return a;"},
		{title: "WithStatement#1", input: `with(a);`, output: "with (a)\n    ;"},
		{title: "WithStatement#2", input: `with(a){}`, output: "with (a) { }"},
		{title: "LabelStatement", input: `::a::`, output: "::a::"},
		{title: "GotoStatement", input: `::a:: goto a`, output: "::a::\ngoto a;"},
		{title: "ThrowStatement", input: `throw a`, output: "throw a;"},
		{title: "TryStatement#1", input: `try {} catch {}`, output: "try { }\ncatch { }"},
		{title: "TryStatement#2", input: `try {} finally {}`, output: "try { }\nfinally { }"},
		{title: "TryStatement#3", input: `try {} catch {} finally {}`, output: "try { }\ncatch { }\nfinally { }"},
		{title: "DebuggerStatement", input: `debugger`, output: "debugger;"},
		{title: "FunctionDeclaration#2", input: `function f(){}`, output: `function f() { }`},
		{title: "FunctionDeclaration#4", input: `async function f(){}`, output: `async function f() { }`},
		{title: "FunctionDeclaration#6", input: `function f<T>(){}`, output: `function f<T>() { }`},
		{title: "FunctionDeclaration#7", input: `function f(a){}`, output: `function f(a) { }`},
		{title: "FunctionDeclaration#8", input: `function f():T{}`, output: `function f(): T { }`},
		{title: "FunctionDeclaration#9", input: `function f();`, output: `function f();`},
		{title: "InterfaceDeclaration#1", input: `interface a {}`, output: "interface a {\n}"},
		{title: "InterfaceDeclaration#2", input: `interface a<T>{}`, output: "interface a<T> {\n}"},
		{title: "InterfaceDeclaration#3", input: `interface a extends b {}`, output: "interface a extends b {\n}"},
		{title: "InterfaceDeclaration#4", input: `interface a extends b, c {}`, output: "interface a extends b, c {\n}"},
		{title: "TypeAliasDeclaration#1", input: `type a = b`, output: "type a = b;"},
		{title: "TypeAliasDeclaration#2", input: `type a<T> = b`, output: "type a<T> = b;"},
		{title: "ModuleDeclaration#7", input: `global;`, output: "global;"},
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
		{title: "ConstructorTypeNode#1", input: `type T = new () => a`, output: `type T = new () => a;`},
		{title: "ConstructorTypeNode#2", input: `type T = new <T>() => a`, output: `type T = new <T>() => a;`},
		{title: "ConstructorTypeNode#3", input: `type T = new (a) => b`, output: `type T = new (a) => b;`},
		{title: "ConstructorTypeNode#4", input: `type T = abstract new () => a`, output: `type T = abstract new () => a;`},
		{title: "TypeQueryNode#1", input: `type T = typeof a`, output: `type T = typeof a;`},
		{title: "TypeQueryNode#2", input: `type T = typeof a.b`, output: `type T = typeof a.b;`},
		{title: "TypeQueryNode#3", input: `type T = typeof a<U>`, output: `type T = typeof a<U>;`},
		{title: "TypeLiteralNode#1", input: `type T = {}`, output: `type T = {};`},
		{title: "TypeLiteralNode#2", input: `type T = {a}`, output: "type T = {\n    a;\n};"},
		{title: "ArrayTypeNode", input: `type T = a[]`, output: "type T = a[];"},
		{title: "TupleTypeNode#1", input: `type T = []`, output: "type T = [\n];", removed: true},
		{title: "TupleTypeNode#2", input: `type T = [a]`, output: "type T = [\n    a\n];", removed: true},
		{title: "TupleTypeNode#3", input: `type T = [a,]`, output: "type T = [\n    a\n];", removed: true},
		{title: "RestTypeNode", input: `type T = [...a]`, output: "type T = [\n    ...a\n];", removed: true},
		{title: "OptionalTypeNode", input: `type T = [a?]`, output: "type T = [\n    a?\n];", removed: true},
		{title: "NamedTupleMember#1", input: `type T = [a: b]`, output: "type T = [\n    a: b\n];", removed: true},
		{title: "NamedTupleMember#2", input: `type T = [a?: b]`, output: "type T = [\n    a?: b\n];", removed: true},
		{title: "NamedTupleMember#3", input: `type T = [...a: b]`, output: "type T = [\n    ...a: b\n];", removed: true},
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
		{title: "ImportTypeNode#6", input: `type T = import(a, { with: { } })`, output: "type T = import(a, { with: {} });"},
		{title: "ImportTypeNode#6", input: `type T = import(a, { with: { b: "c" } })`, output: "type T = import(a, { with: { b: \"c\" } });"},
		{title: "ImportTypeNode#7", input: `type T = import(a, { with: { "b": "c" } })`, output: "type T = import(a, { with: { \"b\": \"c\" } });"},
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
		{title: "ConstructSignature#1", input: "interface I {new ()}", output: "interface I {\n    new ();\n}"},
		{title: "ConstructSignature#2", input: "interface I {new ():a}", output: "interface I {\n    new (): a;\n}"},
		{title: "ConstructSignature#3", input: "interface I {new (p)}", output: "interface I {\n    new (p);\n}"},
		{title: "ConstructSignature#4", input: "interface I {new <T>()}", output: "interface I {\n    new <T>();\n}"},
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
		{title: "ObjectBindingPattern#10", input: "function f({...a});", output: "function f({ ...a });", removed: true},
		{title: "ObjectBindingPattern#11", input: "function f({a: {}});", output: "function f({ a: {} });"},
		{title: "ArrayBindingPattern#1", input: "function f([]);", output: "function f([]);", removed: true},
		{title: "ArrayBindingPattern#2", input: "function f([,]);", output: "function f([,]);", removed: true},
		{title: "ArrayBindingPattern#3", input: "function f([a]);", output: "function f([a]);", removed: true},
		{title: "ArrayBindingPattern#4", input: "function f([a, b]);", output: "function f([a, b]);", removed: true},
		{title: "ArrayBindingPattern#5", input: "function f([a, , b]);", output: "function f([a, , b]);", removed: true},
		{title: "ArrayBindingPattern#6", input: "function f([a = b]);", output: "function f([a = b]);", removed: true},
		{title: "ArrayBindingPattern#7", input: "function f([...a]);", output: "function f([...a]);", removed: true},
		{title: "ArrayBindingPattern#8", input: "function f([{}]);", output: "function f([{}]);", removed: true},
		{title: "ArrayBindingPattern#9", input: "function f([[]]);", output: "function f([[]]);", removed: true},
		{title: "TypeParameterDeclaration#1", input: "function f<T>();", output: "function f<T>();"},
		{title: "TypeParameterDeclaration#2", input: "function f<in T>();", output: "function f<in T>();"},
		{title: "TypeParameterDeclaration#3", input: "function f<T extends U>();", output: "function f<T extends U>();"},
		{title: "TypeParameterDeclaration#4", input: "function f<T = U>();", output: "function f<T = U>();"},
		{title: "TypeParameterDeclaration#5", input: "function f<T extends U = V>();", output: "function f<T extends U = V>();"},
		{title: "TypeParameterDeclaration#6", input: "function f<T, U>();", output: "function f<T, U>();"},
		// {title: "TypeParameterDeclaration#7", input: "function f<T,>();", output: "function f<T,>();"}, // TODO: preserve trailing comma after Strada migration
		{title: "JsxElement1", input: "<a></a>", output: "<a></a>;", jsx: true},
		{title: "JsxElement2", input: "<this></this>", output: "<this></this>;", jsx: true},
		{title: "JsxElement3", input: "<a:b></a:b>", output: "<a:b></a:b>;", jsx: true},
		{title: "JsxElement4", input: "<a.b></a.b>", output: "<a.b></a.b>;", jsx: true},
		{title: "JsxElement5", input: "<a<b>></a>", output: "<a<b>></a>;", jsx: true},
		{title: "JsxElement6", input: "<a b></a>", output: "<a b></a>;", jsx: true},
		{title: "JsxElement7", input: "<a>b</a>", output: "<a>b</a>;", jsx: true},
		{title: "JsxElement8", input: "<a>{b}</a>", output: "<a>{b}</a>;", jsx: true},
		{title: "JsxElement9", input: "<a><b></b></a>", output: "<a><b></b></a>;", jsx: true},
		{title: "JsxElement10", input: "<a><b /></a>", output: "<a><b /></a>;", jsx: true},
		{title: "JsxElement11", input: "<a><></></a>", output: "<a><></></a>;", jsx: true},
		{title: "JsxElement12", input: "<a>\n    {/* missing */}\n    {\n        // foo\n    }\n</a>", output: "<a>\n    {--[[ missing ]]}\n    {\n    -- foo\n    }\n</a>;", jsx: true},
		{title: "JsxSelfClosingElement1", input: "<a />", output: "<a />;", jsx: true},
		{title: "JsxSelfClosingElement2", input: "<this />", output: "<this />;", jsx: true},
		{title: "JsxSelfClosingElement3", input: "<a:b />", output: "<a:b />;", jsx: true},
		{title: "JsxSelfClosingElement4", input: "<a.b />", output: "<a.b />;", jsx: true},
		{title: "JsxSelfClosingElement5", input: "<a<b> />", output: "<a<b> />;", jsx: true},
		{title: "JsxSelfClosingElement6", input: "<a b/>", output: "<a b/>;", jsx: true},
		{title: "JsxFragment1", input: "<></>", output: "<></>;", jsx: true},
		{title: "JsxFragment2", input: "<>b</>", output: "<>b</>;", jsx: true},
		{title: "JsxFragment3", input: "<>{b}</>", output: "<>{b}</>;", jsx: true},
		{title: "JsxFragment4", input: "<><b></b></>", output: "<><b></b></>;", jsx: true},
		{title: "JsxFragment5", input: "<><b /></>", output: "<><b /></>;", jsx: true},
		{title: "JsxFragment6", input: "<><></></>", output: "<><></></>;", jsx: true},
		{title: "JsxAttribute1", input: "<a b/>", output: "<a b/>;", jsx: true},
		{title: "JsxAttribute2", input: "<a b:c/>", output: "<a b:c/>;", jsx: true},
		{title: "JsxAttribute3", input: "<a b=\"c\"/>", output: "<a b=\"c\"/>;", jsx: true},
		{title: "JsxAttribute4", input: "<a b='c'/>", output: "<a b='c'/>;", jsx: true},
		{title: "JsxAttribute5", input: "<a b={c}/>", output: "<a b={c}/>;", jsx: true},
		{title: "JsxAttribute6", input: "<a b=<c></c>/>", output: "<a b=<c></c>/>;", jsx: true},
		{title: "JsxAttribute7", input: "<a b=<c />/>", output: "<a b=<c />/>;", jsx: true},
		{title: "JsxAttribute8", input: "<a b=<></>/>", output: "<a b=<></>/>;", jsx: true},
		{title: "JsxSpreadAttribute", input: "<a {...b}/>", output: "<a {...b}/>;", jsx: true, removed: true},
	}

	for _, rec := range data {
		t.Run(rec.title, func(t *testing.T) {
			t.Parallel()
			if rec.removed {
				t.Skip("emit of removed array/tuple/iteration syntax; covered by compiler baselines")
			}
			file := parsetestutil.ParseTypeScript(rec.input, rec.jsx)
			parsetestutil.CheckDiagnostics(t, file)
			emittestutil.CheckEmit(t, nil, file, rec.output)
		})
	}
}

func TestParenthesizeArrayLiteral(t *testing.T) {
	t.Parallel()
	t.Skip("emit of removed array/tuple/spread syntax; covered by compiler baselines")

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewArrayLiteralExpression(
					factory.NewNodeList(
						[]*ast.Node{
							// will be parenthesized on emit:
							factory.NewBinaryExpression(
								nil, /*modifiers*/
								factory.NewIdentifier("a"),
								nil, /*typeNode*/
								factory.NewToken(ast.KindCommaToken),
								factory.NewIdentifier("b"),
							),
						},
					),
					false, /*multiLine*/
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "[(a, b)];")
}

func TestParenthesizePropertyAccess1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b).c;")
}

func TestParenthesizePropertyAccess2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a?.b).c;")
}

func TestParenthesizePropertyAccess3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewPropertyAccessExpression(
					// will be parenthesized on emit:
					factory.NewNewExpression(
						factory.NewIdentifier("a"),
						nil, /*typeArguments*/
						nil, /*arguments*/
					),
					nil, /*questionDotToken*/
					nil, /*colonToken*/
					factory.NewIdentifier("b"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(new a).b;")
}

func TestParenthesizeElementAccess1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b)[c];")
}

func TestParenthesizeElementAccess2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a?.b)[c];")
}

func TestParenthesizeElementAccess3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewElementAccessExpression(
					// will be parenthesized on emit:
					factory.NewNewExpression(
						factory.NewIdentifier("a"),
						nil, /*typeArguments*/
						nil, /*arguments*/
					),
					nil, /*questionDotToken*/
					factory.NewIdentifier("b"),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(new a)[b];")
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

func TestParenthesizeCall3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewCallExpression(
					// will be parenthesized on emit:
					factory.NewNewExpression(
						factory.NewIdentifier("C"),
						nil, /*typeArguments*/
						nil, /*arguments*/
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(new C)();")
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

func TestParenthesizeNew1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewNewExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					nil, /*typeArguments*/
					factory.NewNodeList([]*ast.Node{}),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "new (a, b)();")
}

func TestParenthesizeNew2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewNewExpression(
					// will be parenthesized on emit:
					factory.NewCallExpression(
						factory.NewIdentifier("C"),
						nil, /*questionDotToken*/
						nil, /*typeArguments*/
						factory.NewNodeList([]*ast.Node{}),
						ast.NodeFlagsNone,
					),
					nil, /*typeArguments*/
					nil, /*arguments*/
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "new (C());")
}

func TestParenthesizeNew3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewNewExpression(
					factory.NewIdentifier("C"),
					nil, /*typeArguments*/
					factory.NewNodeList([]*ast.Node{
						factory.NewBinaryExpression(
							nil, /*modifiers*/
							factory.NewIdentifier("a"),
							nil, /*typeNode*/
							factory.NewToken(ast.KindCommaToken),
							factory.NewIdentifier("b"),
						),
					}),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "new C((a, b));")
}

func TestParenthesizeTaggedTemplate1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewTaggedTemplateExpression(
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
					factory.NewNoSubstitutionTemplateLiteral("", ast.TokenFlagsNone),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// tlua: the no-substitution template lowers to a Lua string; the parenthesized tag is
	// preserved. CheckEmitJS (no reparse) because `tag ""` is Lua string-call sugar tlua
	// doesn't parse back — a non-Lua tagged template's best-effort emit.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(a, b) \"\";")
}

func TestParenthesizeTaggedTemplate2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewTaggedTemplateExpression(
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
					factory.NewNoSubstitutionTemplateLiteral("", ast.TokenFlagsNone),
					ast.NodeFlagsNone,
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// See TestParenthesizeTaggedTemplate1: CheckEmitJS because the tagged-template emit isn't
	// re-parseable tlua.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(a?.b) \"\";")
}

func TestParenthesizeTypeAssertion1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "<T>(a + b);")
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
					// will be parenthesized on emit:
					factory.NewObjectLiteralExpression(
						factory.NewNodeList([]*ast.Node{}),
						false, /*multiLine*/
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(function() return {} end);")
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
					// will be parenthesized on emit:
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(function() return {}.a end);")
}

func TestParenthesizeDelete(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewDeleteExpression(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "delete (a + b);")
}

func TestParenthesizeVoid(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewVoidExpression(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "void (a + b);")
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
		ast.KindPlusEqualsToken,
		ast.KindMinusEqualsToken,
		ast.KindAsteriskEqualsToken,
		ast.KindSlashEqualsToken,
		ast.KindPercentEqualsToken,
		ast.KindBarBarEqualsToken,
		ast.KindAmpersandAmpersandEqualsToken,
		ast.KindInKeyword,
		ast.KindInstanceOfKeyword:
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
	}{
		{operator: ast.KindCommaToken, output: "l, r"},
		{operator: ast.KindCommaToken, left: ast.KindPlusToken, output: "ll + lr, r"},
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
			file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
				[]*ast.Node{
					factory.NewExpressionStatement(
						factory.NewBinaryExpression(
							nil, /*modifiers*/
							makeSide("l", rec.left, &factory),
							nil, /*typeNode*/
							factory.NewToken(rec.operator),
							makeSide("r", rec.right, &factory),
						),
					),
				},
			), factory.NewToken(ast.KindEndOfFile))

			parsetestutil.MarkSyntheticRecursive(file)
			emittestutil.CheckEmit(t, nil, file.AsSourceFile(), rec.output+";")
		})
	}
}

func TestParenthesizeConditional1(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewConditionalExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("b"),
					),
					factory.NewToken(ast.KindQuestionToken),
					factory.NewIdentifier("c"),
					factory.NewToken(ast.KindColonToken),
					factory.NewIdentifier("d"),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The ternary is emit-only JS syntax: no reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(a, b) ? c : d;")
}

func TestParenthesizeConditional2(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewConditionalExpression(
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("a"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindEqualsToken),
						factory.NewIdentifier("b"),
					),
					factory.NewToken(ast.KindQuestionToken),
					factory.NewIdentifier("c"),
					factory.NewToken(ast.KindColonToken),
					factory.NewIdentifier("d"),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The ternary is emit-only JS syntax: no reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "(a = b) ? c : d;")
}

func TestParenthesizeConditional3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewConditionalExpression(
					// will be parenthesized on emit:
					factory.NewArrowFunction(
						nil, /*modifiers*/
						nil, /*typeParameters*/
						factory.NewNodeList([]*ast.Node{}),
						nil, /*returnType*/
						nil, /*fullSignature*/
						factory.NewToken(ast.KindEqualsGreaterThanToken),
						factory.NewBlock(
							factory.NewNodeList([]*ast.Node{}),
							false, /*multiLine*/
						),
					),
					factory.NewToken(ast.KindQuestionToken),
					factory.NewIdentifier("a"),
					factory.NewToken(ast.KindColonToken),
					factory.NewIdentifier("b"),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The ternary is emit-only JS syntax: no reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "((function()\nend) ? a : b);")
}

func TestParenthesizeConditional5(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewConditionalExpression(
					factory.NewIdentifier("a"),
					factory.NewToken(ast.KindQuestionToken),
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("b"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("c"),
					),
					factory.NewToken(ast.KindColonToken),
					factory.NewIdentifier("d"),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The ternary is emit-only JS syntax: no reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "a ? (b, c) : d;")
}

func TestParenthesizeConditional6(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewConditionalExpression(
					factory.NewIdentifier("a"),
					factory.NewToken(ast.KindQuestionToken),
					factory.NewIdentifier("b"),
					factory.NewToken(ast.KindColonToken),
					// will be parenthesized on emit:
					factory.NewBinaryExpression(
						nil, /*modifiers*/
						factory.NewIdentifier("c"),
						nil, /*typeNode*/
						factory.NewToken(ast.KindCommaToken),
						factory.NewIdentifier("d"),
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// The ternary is emit-only JS syntax: no reparse gate.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "a ? b : (c, d);")
}

func TestParenthesizeSpreadElement1(t *testing.T) {
	t.Parallel()
	t.Skip("emit of removed array/tuple/spread syntax; covered by compiler baselines")

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewArrayLiteralExpression(
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewSpreadElement(
								// will be parenthesized on emit:
								factory.NewBinaryExpression(
									nil, /*modifiers*/
									factory.NewIdentifier("a"),
									nil, /*typeNode*/
									factory.NewToken(ast.KindCommaToken),
									factory.NewIdentifier("b"),
								),
							),
						},
					),
					false, /*multiLine*/
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "[...(a, b)];")
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

func TestParenthesizeSpreadElement3(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
				factory.NewNewExpression(
					factory.NewIdentifier("a"),
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
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	// Emit-only, as in TestParenthesizeSpreadElement2.
	emittestutil.CheckEmitJS(t, nil, file.AsSourceFile(), "new a(...(b, c));")
}

func TestParenthesizeExpressionWithTypeArguments(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b)<c>;")
}

func TestParenthesizeAsExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b) as c;")
}

func TestParenthesizeSatisfiesExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b) satisfies c;")
}

func TestParenthesizeNonNullExpression(t *testing.T) {
	t.Parallel()

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(a, b)!;")
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
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "({});")
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
	// Synthesized function expressions must be Lua-flagged like parsed ones,
	// or they print with the unparseable braced body.
	fn.Flags |= ast.NodeFlagsLuaBlock
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewExpressionStatement(fn),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "(function()\nend);")
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

func TestParenthesizeOptionalType(t *testing.T) {
	t.Parallel()
	t.Skip("emit of removed array/tuple/spread syntax; covered by compiler baselines")

	var factory ast.NodeFactory
	file := factory.NewSourceFile(ast.SourceFileParseOptions{FileName: "/file.tlua", Path: "/file.tlua"}, "", factory.NewNodeList(
		[]*ast.Node{
			factory.NewTypeAliasDeclaration(
				nil,                        /*modifiers*/
				factory.NewIdentifier("_"), /*name*/
				nil,                        /*typeParameters*/
				factory.NewTupleTypeNode(
					factory.NewNodeList(
						[]*ast.Node{
							factory.NewOptionalTypeNode(
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
						},
					),
				),
			),
		},
	), factory.NewToken(ast.KindEndOfFile))

	parsetestutil.MarkSyntheticRecursive(file)
	emittestutil.CheckEmit(t, nil, file.AsSourceFile(), "type _ = [\n    (a | b)?\n];")
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
	emittestutil.CheckEmit(t, ec, file.AsSourceFile(), "local _a;\nfunction f() {\n    local _a;\n}")
}

func TestNoTrailingCommaAfterTransform(t *testing.T) {
	t.Parallel()
	t.Skip("emit of removed array/tuple/spread syntax; covered by compiler baselines")

	file := parsetestutil.ParseTypeScript("[a!]", false /*jsx*/)
	emitContext := printer.NewEmitContext()

	var visitor *ast.NodeVisitor
	visitor = emitContext.NewNodeVisitor(func(node *ast.Node) *ast.Node {
		switch node.Kind {
		case ast.KindNonNullExpression:
			node = node.Expression()
		default:
			node = node.VisitEachChild(visitor)
		}
		return node
	})
	file = visitor.VisitSourceFile(file)

	emittestutil.CheckEmit(t, emitContext, file.AsSourceFile(), "[a];")
}

func TestTrailingCommaAfterTransform(t *testing.T) {
	t.Parallel()
	t.Skip("emit of removed array/tuple/spread syntax; covered by compiler baselines")

	file := parsetestutil.ParseTypeScript("[a!,]", false /*jsx*/)
	emitContext := printer.NewEmitContext()

	var visitor *ast.NodeVisitor
	visitor = emitContext.NewNodeVisitor(func(node *ast.Node) *ast.Node {
		switch node.Kind {
		case ast.KindNonNullExpression:
			node = node.Expression()
		default:
			node = node.VisitEachChild(visitor)
		}
		return node
	})
	file = visitor.VisitSourceFile(file)

	emittestutil.CheckEmit(t, emitContext, file.AsSourceFile(), "[a,];")
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

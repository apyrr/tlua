//// [tests/cases/compiler/tluaPrefixExpressionParentheses.tlua] ////

//// [tluaPrefixExpressionParentheses.tlua]
// Lua lets `.`, `[`, `:` and an argument list follow only a prefixexp: a name, an
// index, a call, or a parenthesized expression. Anything else must be
// parenthesized first. The emit parenthesizes it either way, so the output is
// valid Lua even while the error stands.

local a = "x":upper()
local b = {}.x
local c = { 1 }[1]
local d = 1 .x
local e = `t`:upper()
local f = "x"?:upper()
local l = "x"(1)
local m = {}(2)

// Parentheses make any expression a prefixexp; wrappers that erase at emit are
// looked through.
local g = ("x"):upper()
local h = ({ x = 1 }).x
local i = (function() return 1 end)()
local j = ("x" as string):upper()
declare t: { f: () => number } | nil;
local k = t!.f()

// After anything but a prefixexp, a `(` on the next line does not continue the
// expression: Lua ends it there and the `(` starts the next statement. (On the
// same line it is the call above: statements there need a `;` between them.)
declare function report(s: string): nil;
local afterString = "abc"
(report)("x")
local afterTable = {}
(report)("y")

// The error is a grammar check, not a syntax error, so type errors elsewhere in
// the file are still reported.
local wrong: number = "not a number"


//// [tluaPrefixExpressionParentheses.lua]
-- Lua lets `.`, `[`, `:` and an argument list follow only a prefixexp: a name, an
-- index, a call, or a parenthesized expression. Anything else must be
-- parenthesized first. The emit parenthesizes it either way, so the output is
-- valid Lua even while the error stands.
local a = ("x"):upper();
local b = ({}).x;
local c = ({ 1 })[1];
local d = (1).x;
local e = ("t"):upper();
local f = "x" and ("x"):upper();
local l = ("x")(1);
local m = ({})(2);
-- Parentheses make any expression a prefixexp; wrappers that erase at emit are
-- looked through.
local g = ("x"):upper();
local h = ({ x = 1 }).x;
local i = (function()
  return 1;
end)();
local j = ("x"):upper();
local k = t.f();
local afterString = "abc";
(report)("x");
local afterTable = {};
(report)("y");
-- The error is a grammar check, not a syntax error, so type errors elsewhere in
-- the file are still reported.
local wrong = "not a number";

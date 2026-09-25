//// [tests/cases/compiler/tluaTemplateStringReceiver.tlua] ////

//// [tluaTemplateStringReceiver.tlua]
// A template lowers to a `..` chain, which is not a Lua prefixexp: as a receiver or
// operand it needs exactly one pair of parentheses, wherever those come from.

local t: number = 5
local u: string = "u"
local f: (s: string) => string = function(s: string) return s end

// Parenthesized in source (the valid spelling).
local a1 = (`a${t}`):upper()
local a2 = (`a${t}`).len
local a3 = (`a${t}`)[1]
local a4 = (`${t}`):upper()
local a5 = (`a`):upper()
local a6 = ("a"):upper()

// Bare receivers (TLUA100060); the printer still has to parenthesize once.
local b1 = `a${t}`:upper()
local b2 = `a${t}`.len
local b3 = `a${t}`[1]
local b4 = `a`:upper()
local b5 = "a":upper()

// Operands.
local c1 = #`a${t}`
local c2 = `a${t}` .. "b"
local c3 = "b" .. `a${t}`
local c4 = `a${t}` == u
local c5 = #`${t}`
local c6 = f(`a${t}`)
local c7 = `${t}`
local c8 = `a${t}` .. `b${t}`
local c9 = { `a${t}`, k = `b${t}` }
local c10 = (`a${t}`)()


//// [tluaTemplateStringReceiver.lua]
-- A template lowers to a `..` chain, which is not a Lua prefixexp: as a receiver or
-- operand it needs exactly one pair of parentheses, wherever those come from.
local t = 5;
local u = "u";
local f = function(s)
  return s;
end;
-- Parenthesized in source (the valid spelling).
local a1 = ("a" .. tostring(t)):upper();
local a2 = ("a" .. tostring(t)).len;
local a3 = ("a" .. tostring(t))[1];
local a4 = (tostring(t)):upper();
local a5 = ("a"):upper();
local a6 = ("a"):upper();
-- Bare receivers (TLUA100060); the printer still has to parenthesize once.
local b1 = ("a" .. tostring(t)):upper();
local b2 = ("a" .. tostring(t)).len;
local b3 = ("a" .. tostring(t))[1];
local b4 = ("a"):upper();
local b5 = ("a"):upper();
-- Operands.
local c1 = #("a" .. tostring(t));
local c2 = ("a" .. tostring(t)) .. "b";
local c3 = "b" .. "a" .. tostring(t);
local c4 = "a" .. tostring(t) == u;
local c5 = #tostring(t);
local c6 = f("a" .. tostring(t));
local c7 = tostring(t);
local c8 = ("a" .. tostring(t)) .. "b" .. tostring(t);
local c9 = { "a" .. tostring(t), k = "b" .. tostring(t) };
local c10 = ("a" .. tostring(t))();

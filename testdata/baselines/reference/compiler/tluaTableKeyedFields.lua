//// [tests/cases/compiler/tluaTableKeyedFields.tlua] ////

//// [tluaTableKeyedFields.tlua]
// Lua keyed fields: `name = value` is equivalent to `name: value`.
local t = { x = 1, y = "s" };
local tx: number = t.x;
local ty: string = t.y;

// Bracket keys with `=`: literal, dynamic, and numeric.
local key = "dyn";
local u = { ["lit"] = 1, [key] = 2, [1] = "one" };
local ulit: number = u.lit;
local uone: string = u[1];

// Nested tables.
local nested = { a = { b = 2 } };
local nb: number = nested.a.b;

// Identifier and bracket keys coexist in one literal.
local both = { a = 1, b = 2, ["c"] = 3, ["d"] = 4 };
local bsum: number = both.a + both.b + both.c + both.d;

// Inference and widening behave exactly like `:` fields.
local widened = { n = 0 };
widened.n = 42;

// Table fields holding functions (brace-bodied function expressions).
local fns = { add = function(a: number, b: number): number return a + b; end };
local three: number = fns.add(1, 2);


//// [tluaTableKeyedFields.lua]
-- Lua keyed fields: `name = value` is equivalent to `name: value`.
local t = { x = 1, y = "s" };
local tx = t.x;
local ty = t.y;
-- Bracket keys with `=`: literal, dynamic, and numeric.
local key = "dyn";
local u = { ["lit"] = 1, [key] = 2, [1] = "one" };
local ulit = u.lit;
local uone = u[1];
-- Nested tables.
local nested = { a = { b = 2 } };
local nb = nested.a.b;
-- Identifier and bracket keys coexist in one literal.
local both = { a = 1, b = 2, ["c"] = 3, ["d"] = 4 };
local bsum = both.a + both.b + both.c + both.d;
-- Inference and widening behave exactly like `:` fields.
local widened = { n = 0 };
widened.n = 42;
-- Table fields holding functions (brace-bodied function expressions).
local fns = { add = function(a, b)
    return a + b;
  end };
local three = fns.add(1, 2);

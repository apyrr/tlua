//// [tests/cases/compiler/tluaLocalMultiNameDeclarations.tlua] ////

//// [tluaLocalMultiNameDeclarations.tlua]
function pair(): (number, string)
  return 1, "a";
end

function variadic(): (number, ...string)
  return 1, "a", "b";
end

function nothing(): void
end

// Exact arity.
local a, b = pair();

// More names than values: the extras read as nil.
local c, d, e = pair();

// A non-tail call truncates to one value; the tail expands.
local f, g, h = pair(), pair();

// Plain positional values.
local i, j = 1, "x";

// More values than names: extra values are evaluated and dropped.
local k = 1, "dropped";

// A variadic tail covers all remaining names.
local l, m, n = variadic();

// A void call provides no values.
local o, q = nothing();

// No value list at all.
local r, s;

function useAll(): number
  r = a + c + f + l;
  return r;
end


//// [tluaLocalMultiNameDeclarations.lua]
function pair()
    return 1, "a";
end
function variadic()
    return 1, "a", "b";
end
function nothing()
end
-- Exact arity.
local a, b = pair();
-- More names than values: the extras read as nil.
local c, d, e = pair();
-- A non-tail call truncates to one value; the tail expands.
local f, g, h = pair(), pair();
-- Plain positional values.
local i, j = 1, "x";
-- More values than names: extra values are evaluated and dropped.
local k = 1, "dropped";
-- A variadic tail covers all remaining names.
local l, m, n = variadic();
-- A void call provides no values.
local o, q = nothing();
-- No value list at all.
local r, s;
function useAll()
    r = a + c + f + l;
    return r;
end

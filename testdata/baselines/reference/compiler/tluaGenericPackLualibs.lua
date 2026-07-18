//// [tests/cases/compiler/tluaGenericPackLualibs.tlua] ////

//// [tluaGenericPackLualibs.tlua]
// The standard library correlates a call's arguments with its results now that
// generic packs exist: pcall/xpcall/assert forward the callee's pack.

declare function compute(a: number, b: string): (boolean, number);
declare function pure(x: number): string;

// pcall: A is checked against compute's parameters, R becomes compute's results.
local ok, flag, count = pcall(compute, 1, "x");
local okType: boolean = ok;
local flagType: boolean = flag;
local countType: number = count;

// A wrong argument does NOT error: it fails the generic overload and falls
// through to the erased fallback, whose results are `(boolean, ...any)`. That is
// the deliberate trade-off for keeping every previously-legal pcall working
// (overloaded callees, below); precise results reward the calls that match.
local wok2, wres = pcall(compute, "wrong", "x"); // wres: any, not number

// A single-result callee correlates too.
local pok, s = pcall(pure, 3);
local sType: string = s;

// xpcall forwards trailing arguments (LuaJIT) and correlates the result pack.
local xok, xs = xpcall(pure, function(e: any) return e; end, 7);
local xsType: string = xs;

// assert returns everything it was passed; the first value is narrowed non-nil.
declare maybe: number | nil;
local asserted, extra = assert(maybe, "message");
local assertedType: number = asserted;
local extraType: string = extra;

// coroutine.wrap stays deliberately erased: a wrapped call exchanges values with
// `yield` (the first call's results are the first yield's arguments), so packs
// correlated with fn's parameters/returns would be wrong for every coroutine that
// yields. The result is a plain LuaFunction.
local wrapped = coroutine.wrap(compute);
local wok = wrapped(2, "y");

// An OVERLOADED callee cannot pin A to one signature, so pcall falls through to
// the erased overload instead of mis-checking against the wrong one: select's
// "#" form is legal here.
local sok, count = pcall(select, "#", 1, 2);
local sokType: boolean = sok;

// The error-path approximation, pinned: pcall's result is `(boolean, ...R)`, so
// on failure the second value reads as R's first element, not as the error value
// (the true union `(true, ...R) | (false, error)` is inexpressible).
local cok, cerr = pcall(compute, 1, "x");
local cerrIsFirstResult: boolean = cerr;

// select with a literal index slices the argument pack: select(2, ...) drops the
// first value, so `first` is the second argument (a boolean), not the string.
declare str: string;
declare bool: boolean;
declare num: number;
local first, rest = select(2, str, bool, num);
local firstIsBoolean: boolean = first;
local restIsNumber: number = rest;

// select("#", ...) counts the arguments.
local argCount: number = select("#", str, bool, num);

// A non-literal index cannot pin positions, so each selected slot is the sound
// union of the values (or nil past the end).
declare dynamicIndex: number;
local dyn = select(dynamicIndex, str, num);
local dynType: string | number | nil = dyn;

// A slice of exactly one value in the tail of a value list stays a pack (not a
// collapsed scalar): `select(3, ...)` here selects just `num`, and the list keeps
// its arity.
local one, two = str, select(3, str, bool, num);
local oneIsString: string = one;
local twoIsNumber: number = two;

// The same as a multiple-return tail.
function sliceTail(): (string, number)
  return str, select(3, str, bool, num);
end

// A literal index with an open tail keeps the fixed positions and splices the
// tail: values before the index never leak into the result.
function literalWithTail(...: boolean): void
  local m, b = select(2, str, num, ...);
  local mIsNumber: number = m;
  local bIsBoolean: boolean | nil = b;
end

// A literal index landing inside the tail selects the tail's values only.
function literalInsideTail(...: boolean): void
  local t = select(3, str, num, ...);
  local tIsBoolean: boolean | nil = t;
end

// The tail may be a call through a union-typed function -- a union of packs.
// (This crashed the checker before it distributed over the union.)
declare function gPack(): (number, string);
declare function hPack(): (boolean, nil);
declare cond: boolean;
local packFn = cond and gPack or hPack;
declare dynamic: number;
local u = select(dynamic, packFn());
local uType: number | string | boolean | nil = u;

// A literal slice keeps literal freshness: the annotation still accepts it.
local exact: 5 = select(1, 5);


//// [tluaGenericPackLualibs.lua]
-- The standard library correlates a call's arguments with its results now that
-- generic packs exist: pcall/xpcall/assert forward the callee's pack.
-- pcall: A is checked against compute's parameters, R becomes compute's results.
local ok, flag, count = pcall(compute, 1, "x");
local okType = ok;
local flagType = flag;
local countType = count;
-- A wrong argument does NOT error: it fails the generic overload and falls
-- through to the erased fallback, whose results are `(boolean, ...any)`. That is
-- the deliberate trade-off for keeping every previously-legal pcall working
-- (overloaded callees, below); precise results reward the calls that match.
local wok2, wres = pcall(compute, "wrong", "x"); -- wres: any, not number
-- A single-result callee correlates too.
local pok, s = pcall(pure, 3);
local sType = s;
-- xpcall forwards trailing arguments (LuaJIT) and correlates the result pack.
local xok, xs = xpcall(pure, function(e)
    return e;
end, 7);
local xsType = xs;
local asserted, extra = assert(maybe, "message");
local assertedType = asserted;
local extraType = extra;
-- coroutine.wrap stays deliberately erased: a wrapped call exchanges values with
-- `yield` (the first call's results are the first yield's arguments), so packs
-- correlated with fn's parameters/returns would be wrong for every coroutine that
-- yields. The result is a plain LuaFunction.
local wrapped = coroutine.wrap(compute);
local wok = wrapped(2, "y");
-- An OVERLOADED callee cannot pin A to one signature, so pcall falls through to
-- the erased overload instead of mis-checking against the wrong one: select's
-- "#" form is legal here.
local sok, count = pcall(select, "#", 1, 2);
local sokType = sok;
-- The error-path approximation, pinned: pcall's result is `(boolean, ...R)`, so
-- on failure the second value reads as R's first element, not as the error value
-- (the true union `(true, ...R) | (false, error)` is inexpressible).
local cok, cerr = pcall(compute, 1, "x");
local cerrIsFirstResult = cerr;
local first, rest = select(2, str, bool, num);
local firstIsBoolean = first;
local restIsNumber = rest;
-- select("#", ...) counts the arguments.
local argCount = select("#", str, bool, num);
local dyn = select(dynamicIndex, str, num);
local dynType = dyn;
-- A slice of exactly one value in the tail of a value list stays a pack (not a
-- collapsed scalar): `select(3, ...)` here selects just `num`, and the list keeps
-- its arity.
local one, two = str, select(3, str, bool, num);
local oneIsString = one;
local twoIsNumber = two;
-- The same as a multiple-return tail.
function sliceTail()
    return str, select(3, str, bool, num);
end
-- A literal index with an open tail keeps the fixed positions and splices the
-- tail: values before the index never leak into the result.
function literalWithTail(...)
    local m, b = select(2, str, num, ...);
    local mIsNumber = m;
    local bIsBoolean = b;
end
-- A literal index landing inside the tail selects the tail's values only.
function literalInsideTail(...)
    local t = select(3, str, num, ...);
    local tIsBoolean = t;
end
local packFn = cond and gPack or hPack;
local u = select(dynamic, packFn());
local uType = u;
-- A literal slice keeps literal freshness: the annotation still accepts it.
local exact = select(1, 5);

//// [tests/cases/compiler/tluaTableNoMethods.tlua] ////

//// [tluaTableNoMethods.tlua]
// Object-literal methods and accessors are removed: they decompose into
// positional expressions with natural errors.
local m = { m() { return 1; } };
local g = { get p() { return 1; } };
local s = { set p(v: number) { } };

// The replacement: functions are table values.
local fns = { run = function(): number return 1; end };
local one: number = fns.run();

// Literal string/number keys use the bracket form.
local keyed = { [1] = "a", ["k"] = 2 };
local ka: string = keyed[1];
local kk: number = keyed.k;


//// [tluaTableNoMethods.lua]
-- Object-literal methods and accessors are removed: they decompose into
-- positional expressions with natural errors.
local m = { m(), { , 1, } };
local g = { get, p(), { , 1, } };
local s = { set, p(v, number), {} };
-- The replacement: functions are table values.
local fns = { run = function()
        return 1;
    end };
local one = fns.run();
-- Literal string/number keys use the bracket form.
local keyed = { [1] = "a", ["k"] = 2 };
local ka = keyed[1];
local kk = keyed.k;

//// [tests/cases/compiler/tluaGenericPackBasics.tlua] ////

//// [tluaGenericPackBasics.tlua]
// A generic pack parameter, declared `<...A>`, binds a whole value pack. It is
// spread with `...A` in a multiple-return position and `...: A` as a vararg
// parameter, so a function can forward its arguments and correlate results.

// The canonical forwarder: pass the pack straight through.
declare function apply<...A, ...R>(f: (...: A) => (...R), ...: A): (...R);

declare function f2(a: number, b: string): (boolean, string);
declare function f0(): number;

// R is inferred from f's return pack; A is inferred from the trailing arguments.
local ok, msg = apply(f2, 1, "x");
local okType: boolean = ok;
local msgType: string = msg;

// A single-result callee: the pack collapses to one value.
local single: number = apply(f0);

// A pack parameter reused across the function and callback correlates the two.
declare function forward<...A>(f: (...: A) => void, ...: A): void;
forward(f2, 1, "x");


//// [tluaGenericPackBasics.lua]
-- A generic pack parameter, declared `<...A>`, binds a whole value pack. It is
-- spread with `...A` in a multiple-return position and `...: A` as a vararg
-- parameter, so a function can forward its arguments and correlate results.
-- R is inferred from f's return pack; A is inferred from the trailing arguments.
local ok, msg = apply(f2, 1, "x");
local okType = ok;
local msgType = msg;
-- A single-result callee: the pack collapses to one value.
local single = apply(f0);
forward(f2, 1, "x");

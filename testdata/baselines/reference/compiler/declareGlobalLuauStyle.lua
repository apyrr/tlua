//// [tests/cases/compiler/declareGlobalLuauStyle.tlua] ////

//// [globals.d.tlua]
declare version: string;
declare limits: { min: number, max: number };

//// [main.tlua]
local n: number = limits.min;
local s: string = version;
version = "overwritten"; // error: constant
limits = { min = 0, max = 1 }; // error: constant
local bad: boolean = version; // error: string not boolean


//// [main.lua]
local n = limits.min;
local s = version;
version = "overwritten"; -- error: constant
limits = { min = 0, max = 1 }; -- error: constant
local bad = version; -- error: string not boolean

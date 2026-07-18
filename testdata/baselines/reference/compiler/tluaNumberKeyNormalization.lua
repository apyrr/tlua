//// [tests/cases/compiler/tluaNumberKeyNormalization.tlua] ////

//// [tluaNumberKeyNormalization.tlua]
// Every spelling of a numeric value names the same key.
local t = { [0x10] = "hex", [1e2] = "exp", [1.5] = "frac", [-1] = "neg" };
local hex: string = t[16];
local exp: string = t[100];
local frac: string = t[1.5];
local neg: string = t[-1];
local es = {hex, exp, frac, neg};


//// [tluaNumberKeyNormalization.lua]
-- Every spelling of a numeric value names the same key.
local t = { [0x10] = "hex", [1e2] = "exp", [1.5] = "frac", [-1] = "neg" };
local hex = t[16];
local exp = t[100];
local frac = t[1.5];
local neg = t[-1];
local es = { hex, exp, frac, neg };

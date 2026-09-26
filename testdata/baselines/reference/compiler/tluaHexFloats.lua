//// [tests/cases/compiler/tluaHexFloats.tlua] ////

//// [tluaHexFloats.tlua]
// Lua hexadecimal numerals carry an optional fraction and binary exponent:
// 0xA.8p1 is (10 + 8/16) * 2^1.

local exponent: 16 = 0x1p4
local both: 21 = 0xA.8p1
local fraction: 0.5 = 0x.8
local negative: 0.25 = 0x1P-2
local positive: 4 = 0x1p+2
local trailingDot: 1 = 0x1.
local integer: 255 = 0xff
local concat = 0x10 .. "a"
local tightConcat = 0x10.."a"
-- An identifier cannot follow a hexadecimal literal, integer or float.
local afterInteger = 0x1g
local afterFloat = 0x1.8g
local afterLeadingZero = 012g
-- A binary exponent is exponent notation, not a digit run too long to be exact.
local large = 0x1p60

-- A binary exponent needs digits.
local missing = 0x1p


//// [tluaHexFloats.lua]
-- Lua hexadecimal numerals carry an optional fraction and binary exponent:
-- 0xA.8p1 is (10 + 8/16) * 2^1.
local exponent = 0x1p4;
local both = 0xA.8p1;
local fraction = 0x.8;
local negative = 0x1P-2;
local positive = 0x1p+2;
local trailingDot = 0x1.;
local integer = 0xff;
local concat = 0x10 .. "a";
local tightConcat = 0x10 .. "a";
-- An identifier cannot follow a hexadecimal literal, integer or float.
local afterInteger = 0x1;
g;
local afterFloat = 0x1.8;
g;
local afterLeadingZero = 12;
g;
-- A binary exponent is exponent notation, not a digit run too long to be exact.
local large = 0x1p60;
-- A binary exponent needs digits.
local missing = 0x1p;

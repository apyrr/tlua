//// [tests/cases/compiler/tluaStringDecimalEscapeRange.tlua] ////

//// [tluaStringDecimalEscapeRange.tlua]
// Lua's decimal escape names one byte, so it cannot exceed 255.
local ok = "\255"
local tooLarge = "\256"


//// [tluaStringDecimalEscapeRange.lua]
-- Lua's decimal escape names one byte, so it cannot exceed 255.
local ok = "\xff";
local tooLarge = "\\256";

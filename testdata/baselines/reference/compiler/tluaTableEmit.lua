//// [tests/cases/compiler/tluaTableEmit.tlua] ////

//// [tluaTableEmit.tlua]
// Keyed fields round-trip in emit: `x = 1` and `[k] = v` print with `=`;
// `;` separators normalize to `,`.
local key = "dyn";
local t = {
    x = 1;
    y = 2,
    ["lit"] = 3;
    [key] = 4,
    [5] = "five";
};


//// [tluaTableEmit.lua]
-- Keyed fields round-trip in emit: `x = 1` and `[k] = v` print with `=`;
-- `;` separators normalize to `,`.
local key = "dyn";
local t = {
    x = 1,
    y = 2,
    ["lit"] = 3,
    [key] = 4,
    [5] = "five",
};

//// [tests/cases/compiler/tluaTableSeparators.tlua] ////

//// [tluaTableSeparators.tlua]
// `;` is a field separator alongside `,`; mixing and trailing separators
// are allowed.
local a = { a = 1; b = 2, c = 3; };
local b = { x = 1; y = 2 };
local c = {
    first = 1;
    second = 2,
    third = 3;
};
local d = { only = 1; };
local e = { trailingComma = 1, };
local all = {a, b, c, d, e};


//// [tluaTableSeparators.lua]
-- `;` is a field separator alongside `,`; mixing and trailing separators
-- are allowed.
local a = { a = 1, b = 2, c = 3, };
local b = { x = 1, y = 2 };
local c = {
  first = 1,
  second = 2,
  third = 3,
};
local d = { only = 1, };
local e = { trailingComma = 1, };
local all = { a, b, c, d, e };

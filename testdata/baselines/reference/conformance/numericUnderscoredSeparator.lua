//// [tests/cases/conformance/ported/numericUnderscoredSeparator.tlua] ////

//// [numericUnderscoredSeparator.tlua]
-- ported from tests/cases/compiler/numericUnderscoredSeparator.ts
-- dropped: @target vary-list directive (tlua defaults to latest target)

local _ = 1_000_000_000_000
local _ = 0b1010_0001_1000_0101
local _ = 0b1010_0001_1000_0101
local _ = 0xA0_B0_C0


//// [numericUnderscoredSeparator.lua]
-- ported from tests/cases/compiler/numericUnderscoredSeparator.ts
-- dropped: @target vary-list directive (tlua defaults to latest target)
local _ = 1000000000000;
local _ = 41349;
local _ = 41349;
local _ = 10531008;

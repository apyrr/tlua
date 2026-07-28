//// [tests/cases/conformance/ported/parser.numericSeparators.binary.tlua] ////

//// [parser.numericSeparators.binary.tlua]
-- ported from tests/cases/conformance/parser/ecmascript2021/numericSeparators/parser.numericSeparators.binary.ts
-- dropped: nothing of substance; bare literal expression statements rebound to `local _ =` (tlua disallows bare non-call/assignment expression statements)


local _ = 0b00_11;
local _ = 0B0_1;
local _ = 0b1100_0011;
local _ = 0B0_11_0101;


//// [parser.numericSeparators.binary.lua]
-- ported from tests/cases/conformance/parser/ecmascript2021/numericSeparators/parser.numericSeparators.binary.ts
-- dropped: nothing of substance; bare literal expression statements rebound to `local _ =` (tlua disallows bare non-call/assignment expression statements)
local _ = 0b00_11;
local _ = 0B0_1;
local _ = 0b1100_0011;
local _ = 0B0_11_0101;

//// [tests/cases/conformance/ported/parser.numericSeparators.hex.tlua] ////

//// [parser.numericSeparators.hex.tlua]
-- ported from tests/cases/conformance/parser/ecmascript2021/numericSeparators/parser.numericSeparators.hex.ts
-- dropped: the es2015 target directive (tlua uses its latest target)
-- dropped: bare literal expression statements rebound to `local _ =` (tlua disallows bare non-call/assignment expression statements)

local _ = 0x00_11
local _ = 0X0_1
local _ = 0x1100_0011
local _ = 0X0_11_0101


//// [parser.numericSeparators.hex.lua]
-- ported from tests/cases/conformance/parser/ecmascript2021/numericSeparators/parser.numericSeparators.hex.ts
-- dropped: the es2015 target directive (tlua uses its latest target)
-- dropped: bare literal expression statements rebound to `local _ =` (tlua disallows bare non-call/assignment expression statements)
local _ = 0x00_11;
local _ = 0X0_1;
local _ = 0x1100_0011;
local _ = 0X0_11_0101;

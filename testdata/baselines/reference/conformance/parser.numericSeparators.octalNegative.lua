//// [tests/cases/conformance/ported/parser.numericSeparators.octalNegative.tlua] ////

//// [1.tlua]
local _ = 0o00_

//// [2.tlua]
local _ = 0o_110

//// [3.tlua]
local _ = 0_O0101

//// [4.tlua]
local _ = 0o01__11

//// [5.tlua]
local _ = 0O0110_0110__

//// [6.tlua]
local _ = 0o___0111010_0101_1


//// [1.lua]
local _ = 0o00_;
//// [2.lua]
local _ = 0o_110;
//// [3.lua]
local _ = 0;
O0101;
//// [4.lua]
local _ = 0o01__11;
//// [5.lua]
local _ = 0O0110_0110__;
//// [6.lua]
local _ = 0o___0111010_0101_1;

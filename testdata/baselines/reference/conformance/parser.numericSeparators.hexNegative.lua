//// [tests/cases/conformance/ported/parser.numericSeparators.hexNegative.tlua] ////

//// [1.tlua]
local _ = 0x00_

//// [2.tlua]
local _ = 0x_110

//// [3.tlua]
local _ = 0_X0101

//// [4.tlua]
local _ = 0x01__11

//// [5.tlua]
local _ = 0X0110_0110__

//// [6.tlua]
local _ = 0x___0111010_0101_1


//// [1.lua]
local _ = 0x00_;
//// [2.lua]
local _ = 0x_110;
//// [3.lua]
local _ = 0;
X0101;
//// [4.lua]
local _ = 0x01__11;
//// [5.lua]
local _ = 0X0110_0110__;
//// [6.lua]
local _ = 0x___0111010_0101_1;

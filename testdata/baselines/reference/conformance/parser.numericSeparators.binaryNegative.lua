//// [tests/cases/conformance/ported/parser.numericSeparators.binaryNegative.tlua] ////

//// [1.tlua]
local _ = 0b00_;

//// [2.tlua]
local _ = 0b_110;

//// [3.tlua]
local _ = 0_B0101;

//// [4.tlua]
local _ = 0b01__11;

//// [5.tlua]
local _ = 0B0110_0110__;

//// [6.tlua]
local _ = 0b___0111010_0101_1;


//// [1.lua]
local _ = 0b00_;
//// [2.lua]
local _ = 0b_110;
//// [3.lua]
local _ = 0;
B0101;
//// [4.lua]
local _ = 0b01__11;
//// [5.lua]
local _ = 0B0110_0110__;
//// [6.lua]
local _ = 0b___0111010_0101_1;

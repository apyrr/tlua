//// [tests/cases/compiler/tluaEnumIsIdentifier.tlua] ////

//// [tluaEnumIsIdentifier.tlua]
local enum = 1
enum = enum + 1

local value: number = enum


//// [tluaEnumIsIdentifier.lua]
local enum = 1;
enum = enum + 1;
local value = enum;

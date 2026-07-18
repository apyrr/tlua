//// [tests/cases/compiler/simpleTestMultiFile.tlua] ////

//// [foo.tlua]
local x: number = "";

//// [bar.tlua]
local y: string = 1;

//// [foo.lua]
local x = "";
//// [bar.lua]
local y = 1;

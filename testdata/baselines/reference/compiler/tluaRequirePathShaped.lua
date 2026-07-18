//// [tests/cases/compiler/tluaRequirePathShaped.tlua] ////

//// [m.tlua]
local x = 1;
return { x = x };

//// [main.tlua]
local a = require("./m");
local b = require("../up");
local c = require("dir/m");
local ok = require("m");
ok;


//// [m.lua]
local x = 1;
return { x = x };
//// [main.lua]
local a = require("./m");
local b = require("../up");
local c = require("dir/m");
local ok = require("m");
ok;

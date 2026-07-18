//// [tests/cases/compiler/tluaRequireCanonicalName.tlua] ////

//// [init.tlua]
local v = 1;
return { v = v };

//// [main.tlua]
local a = require("pkg");
local b = require("pkg.init");
a; b;


//// [init.lua]
local v = 1;
return { v = v };
//// [main.lua]
local a = require("pkg");
local b = require("pkg.init");
a;
b;

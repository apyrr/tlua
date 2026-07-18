//// [tests/cases/compiler/tluaRequireInit.tlua] ////

//// [init.tlua]
local version = 1;
return { version = version };

//// [main.tlua]
local pkg = require("pkg");
local v: number = pkg.version;
v;


//// [init.lua]
local version = 1;
return { version = version };
//// [main.lua]
local pkg = require("pkg");
local v = pkg.version;
v;

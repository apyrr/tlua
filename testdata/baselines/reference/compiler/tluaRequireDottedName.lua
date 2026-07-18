//// [tests/cases/compiler/tluaRequireDottedName.tlua] ////

//// [util.tlua]
local value = 42;
return { value = value };

//// [main.tlua]
local util = require("app.util");
local n: number = util.value;
n;


//// [util.lua]
local value = 42;
return { value = value };
//// [main.lua]
local util = require("app.util");
local n = util.value;
n;

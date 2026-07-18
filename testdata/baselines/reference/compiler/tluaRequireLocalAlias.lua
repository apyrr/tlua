//// [tests/cases/compiler/tluaRequireLocalAlias.tlua] ////

//// [m.tlua]
local value = 42;
return { value = value };

//// [main.tlua]
// The standard Lua caching idiom keeps require's module typing.
local require = require;
local m = require("m");
local n: number = m.value;
n;


//// [m.lua]
local value = 42;
return { value = value };
//// [main.lua]
-- The standard Lua caching idiom keeps require's module typing.
local require = require;
local m = require("m");
local n = m.value;
n;

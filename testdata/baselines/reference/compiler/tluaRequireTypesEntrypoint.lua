//// [tests/cases/compiler/tluaRequireTypesEntrypoint.tlua] ////

//// [package.json]
{ "name": "pkg", "types": "typed.tlua", "main": "impl.lua" }
//// [typed.tlua]
local answer = 42;
return { answer = answer };
//// [main.tlua]
// A package's typed entrypoint (package.json "types") types the require.
local pkg = require("pkg");
local n: number = pkg.answer;
n;


//// [main.lua]
-- A package's typed entrypoint (package.json "types") types the require.
local pkg = require("pkg");
local n = pkg.answer;
n;

//// [tests/cases/compiler/tluaModuleWordsAreIdentifiers.tlua] ////

//// [tluaModuleWordsAreIdentifiers.tlua]
// With the module syntax deleted, its words are ordinary Lua names.
local export = 1;
local from = "f";
local default = { n = 1 };
local t = { export = export, from = from, default = default.n };

function M_from(self: {}, mode: string): string
  return mode;
end

local obj = { from = M_from, default = 2, export = 3 };
local a = obj:from("on");
local b = obj.default;
local c = obj.export;

// `require` is an ordinary global function, so it can be shadowed or aliased.
local alias = require;
local viaAlias = alias("bit");

return t;


//// [tluaModuleWordsAreIdentifiers.lua]
-- With the module syntax deleted, its words are ordinary Lua names.
local export = 1;
local from = "f";
local default = { n = 1 };
local t = { export = export, from = from, default = default.n };
function M_from(self, mode)
    return mode;
end
local obj = { from = M_from, default = 2, export = 3 };
local a = obj:from("on");
local b = obj.default;
local c = obj.export;
-- `require` is an ordinary global function, so it can be shadowed or aliased.
local alias = require;
local viaAlias = alias("bit");
return t;

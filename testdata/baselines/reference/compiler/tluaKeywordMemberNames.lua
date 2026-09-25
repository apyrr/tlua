//// [tests/cases/compiler/tluaKeywordMemberNames.tlua] ////

//// [access.tlua]
local t: { [k: string]: number } = {}
local a1 = _G.nil
local a2 = t.end
local a3 = t.then
local a4 = t.local
local a5 = t.true
local a6 = t?.until
t.end = 1
t.nil = 2
local ok1 = t.continue
local ok2 = t.type
local ok3 = t.self
local ok4 = t["end"]

//// [tableKeys.tlua]
local k1 = { nil = 1, self = 2, type = 3 }
local k2 = { ["end"] = 1 }

//// [funcDecl.tlua]
local m = {}
function m.end() end
function m:nil() end
function m.ok() end
m.while = {}
function m.while.f() end
m["end"] = function() end

//// [iface.tlua]
interface K { end: number; nil: number }
local k: K = { ["end"] = 1, ["nil"] = 2 }
local r1 = k["end"]
local r2 = k.end


//// [access.lua]
local t = {};
local a1 = _G["nil"];
local a2 = t["end"];
local a3 = t["then"];
local a4 = t["local"];
local a5 = t["true"];
local a6 = t and t["until"];
t["end"] = 1;
t["nil"] = 2;
local ok1 = t.continue;
local ok2 = t.type;
local ok3 = t.self;
local ok4 = t["end"];
//// [tableKeys.lua]
local k1 = { ["nil"] = 1, self = 2, type = 3 };
local k2 = { ["end"] = 1 };
//// [funcDecl.lua]
local m = {};
function m.end()
end
function m:nil()
end
function m.ok()
end
m["while"] = {};
function m["while"].f()
end
m["end"] = function()
end;
//// [iface.lua]
local k = { ["end"] = 1, ["nil"] = 2 };
local r1 = k["end"];
local r2 = k["end"];

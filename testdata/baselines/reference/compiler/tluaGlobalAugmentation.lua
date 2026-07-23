//// [tests/cases/compiler/tluaGlobalAugmentation.tlua] ////

//// [a.tlua]
G.x = 1;
local gx: number | string = G.x;
local n: number | string = N;

//// [b.tlua]
G = {};
G.x = "two";
N = 1;
N = "two";

//// [c.tlua]
local G = {};
G.localOnly = true;
local localOnly: boolean = G.localOnly;

function G.identity(value: number): number
  return value;
end

local localIdentity: number = G.identity(1);

//// [d.tlua]
function G.double(value: number): number
  return value * 2;
end

local doubled: number = G.double(2);

function G:triple(value: number): number
  return value * 3;
end

local tripled: number = G.triple(G, 3);

//// [e.tlua]
Shadow = {};
Shadow.globalValue = 1;
local globalValue: number = Shadow.globalValue;

local Shadow = {};
Shadow.localValue = true;
local localValue: boolean = Shadow.localValue;

ModuleGlobal = {};
ModuleGlobal.value = "module";

//// [f.tlua]
local moduleValue: string = ModuleGlobal.value;
local crossFileGlobalValue: number = Shadow.globalValue;


//// [a.lua]
G.x = 1;
local gx = G.x;
local n = N;
//// [b.lua]
G = {};
G.x = "two";
N = 1;
N = "two";
//// [c.lua]
local G = {};
G.localOnly = true;
local localOnly = G.localOnly;
function G.identity(value)
  return value;
end
local localIdentity = G.identity(1);
//// [d.lua]
function G.double(value)
  return value * 2;
end
local doubled = G.double(2);
function G:triple(value)
  return value * 3;
end
local tripled = G.triple(G, 3);
//// [e.lua]
Shadow = {};
Shadow.globalValue = 1;
local globalValue = Shadow.globalValue;
local Shadow = {};
Shadow.localValue = true;
local localValue = Shadow.localValue;
ModuleGlobal = {};
ModuleGlobal.value = "module";
//// [f.lua]
local moduleValue = ModuleGlobal.value;
local crossFileGlobalValue = Shadow.globalValue;

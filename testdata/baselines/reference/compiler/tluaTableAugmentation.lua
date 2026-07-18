//// [tests/cases/compiler/tluaTableAugmentation.tlua] ////

//// [tluaTableAugmentation.tlua]
local M = {};
local before: number | string = M.x;
M.x = 1;
M.x = "two";
local x: number | string = M.x;

local nested = {};
nested.sub = {};
nested.sub.value = 1;
local value: number = nested.sub.value;

local keyed = {};
keyed["name"] = true;
keyed[1] = "one";
local name: boolean = keyed.name;
local one: string = keyed[1];

local fixed = { x = 1 };
fixed.x = "no";

local annotated: { x: number } = { x = 1 };
annotated.y = 2;

local dynamic = {};
local key = "key";
dynamic[key] = 1;

local closed = true and {} or {};
closed.x = 1;

local collision = {};
collision.member = 1;
function collision.member(): void
end

local reverseCollision = {};
function reverseCollision.member(): void
end
reverseCollision.member = 1;


//// [tluaTableAugmentation.lua]
local M = {};
local before = M.x;
M.x = 1;
M.x = "two";
local x = M.x;
local nested = {};
nested.sub = {};
nested.sub.value = 1;
local value = nested.sub.value;
local keyed = {};
keyed["name"] = true;
keyed[1] = "one";
local name = keyed.name;
local one = keyed[1];
local fixed = { x = 1 };
fixed.x = "no";
local annotated = { x = 1 };
annotated.y = 2;
local dynamic = {};
local key = "key";
dynamic[key] = 1;
local closed = true and {} or {};
closed.x = 1;
local collision = {};
collision.member = 1;
function collision.member()
end
local reverseCollision = {};
function reverseCollision.member()
end
reverseCollision.member = 1;

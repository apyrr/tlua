//// [tests/cases/compiler/tluaDottedFunctionShadowing.tlua] ////

//// [tluaDottedFunctionShadowing.tlua]
local M = {};

function M.first(): number
  return 1;
end

local firstOnly = M.first();

local M = {};

function M.second(): number
  return 2;
end

local secondOnly = M.second();
local gone = M.first();


//// [tluaDottedFunctionShadowing.lua]
local M = {};
function M.first()
  return 1;
end
local firstOnly = M.first();
local M = {};
function M.second()
  return 2;
end
local secondOnly = M.second();
local gone = M.first();

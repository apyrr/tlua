//// [tests/cases/compiler/tluaDottedFunctionDeclarations.tlua] ////

//// [tluaDottedFunctionDeclarations.tlua]
local M = {};

function M.identity(a: number): number
  return a;
end

function M.add(a: number, b: number): number
  return M.identity(a) + b;
end

local early = M.add(1, 2);
local late = M.identity(3);

local bad = M.add("nope", 2);
local missing = M.nonexistent(1);


//// [tluaDottedFunctionDeclarations.lua]
local M = {};
function M.identity(a)
    return a;
end
function M.add(a, b)
    return M.identity(a) + b;
end
local early = M.add(1, 2);
local late = M.identity(3);
local bad = M.add("nope", 2);
local missing = M.nonexistent(1);

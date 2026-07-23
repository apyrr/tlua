//// [tests/cases/compiler/tluaRequireModuleBasic.tlua] ////

//// [util.tlua]
local M = {};

function M.clamp(n: number, lo: number, hi: number): number
  if n < lo then return lo end
  if n > hi then return hi end
  return n;
end

M.name = "util";

return M;

//// [main.tlua]
// A chunk's top-level return is its module value; require reports it.
local util = require("util");
local n = util.clamp(5, 1, 3);
local s: string = util.name;

// The extension is implicit, as in Lua's module names.
local missing = util.nope;


//// [util.lua]
local M = {};
function M.clamp(n, lo, hi)
  if n < lo then
    return lo;
  end
  if n > hi then
    return hi;
  end
  return n;
end
M.name = "util";
return M;
//// [main.lua]
-- A chunk's top-level return is its module value; require reports it.
local util = require("util");
local n = util.clamp(5, 1, 3);
local s = util.name;
-- The extension is implicit, as in Lua's module names.
local missing = util.nope;

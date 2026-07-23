//// [tests/cases/compiler/tluaLocalMultiNameEmit.tlua] ////

//// [tluaLocalMultiNameEmit.tlua]
function pair(): (number, string)
  return 1, "a";
end

local a, b = pair();
local c, d = 1, "x";
local e, f, g = pair(), true;
local h, i;
local j: number, k: string = pair();

// A parenthesized comma expression keeps its parens (it is one value).
local l, m = (0, 1), 2;


//// [tluaLocalMultiNameEmit.lua]
function pair()
  return 1, "a";
end
local a, b = pair();
local c, d = 1, "x";
local e, f, g = pair(), true;
local h, i;
local j, k = pair();
-- A parenthesized comma expression keeps its parens (it is one value).
local l, m = (0, 1), 2;

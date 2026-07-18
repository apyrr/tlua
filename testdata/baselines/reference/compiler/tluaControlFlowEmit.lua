//// [tests/cases/compiler/tluaControlFlowEmit.tlua] ////

//// [m.tlua]
local N = 3;
local done = false;
local counter = 0;
return { N = N, done = done, counter = counter };

//// [main.tlua]
local m = require("m");
local N, done = m.N, m.done;

// Required values substitute in numeric-for bounds and until-conditions.
function f(): number
  local total = 0;
  for i = 1, N do
    total = total + i;
  end
  repeat
    total = total + 1;
  until done;
  while done do
    total = total - 1;
  end
  return total;
end

// Top-level Lua statements emit verbatim (no panic, no module wrapper).
for i = 1, 2 do
  f();
end
repeat
  f();
until true;
do
  f();
end


//// [m.lua]
local N = 3;
local done = false;
local counter = 0;
return { N = N, done = done, counter = counter };
//// [main.lua]
local m = require("m");
local N, done = m.N, m.done;
-- Required values substitute in numeric-for bounds and until-conditions.
function f()
    local total = 0;
    for i = 1, N do
        total = total + i;
    end
    repeat
        total = total + 1;
    until done;
    while done do
        total = total - 1;
    end
    return total;
end
-- Top-level Lua statements emit verbatim (no panic, no module wrapper).
for i = 1, 2 do
    f();
end
repeat
    f();
until true;
do
    f();
end

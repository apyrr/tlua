//// [tests/cases/compiler/tluaDottedFunctionNonEmptyTable.tlua] ////

//// [tluaDottedFunctionNonEmptyTable.tlua]
local Counter = { count = 0, label = "hits" };

function Counter.bump(by: number): number
  Counter.count = Counter.count + by;
  return Counter.count;
end

local total = Counter.bump(2) + Counter.count;
local name = Counter.label;

// A dotted declaration may not redeclare a field written in the table itself.
local Clash = { run = 1 };

function Clash.run(): void
end


//// [tluaDottedFunctionNonEmptyTable.lua]
local Counter = { count = 0, label = "hits" };
function Counter.bump(by)
    Counter.count = Counter.count + by;
    return Counter.count;
end
local total = Counter.bump(2) + Counter.count;
local name = Counter.label;
-- A dotted declaration may not redeclare a field written in the table itself.
local Clash = { run = 1 };
function Clash.run()
end

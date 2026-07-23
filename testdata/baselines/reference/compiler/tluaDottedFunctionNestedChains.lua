//// [tests/cases/compiler/tluaDottedFunctionNestedChains.tlua] ////

//// [tluaDottedFunctionNestedChains.tlua]
local a = { b = {} };

function a.b.deep(x: number): number
  return x;
end

local viaChain = a.b.deep(1);

local outer = {};

function outer.mid(): void
end

// The intermediate link is a function, not a table, so it cannot be extended.
local host = { fn = 0 };

function host.fn.nested(): void
end


//// [tluaDottedFunctionNestedChains.lua]
local a = { b = {} };
function a.b.deep(x)
  return x;
end
local viaChain = a.b.deep(1);
local outer = {};
function outer.mid()
end
-- The intermediate link is a function, not a table, so it cannot be extended.
local host = { fn = 0 };
function host.fn.nested()
end

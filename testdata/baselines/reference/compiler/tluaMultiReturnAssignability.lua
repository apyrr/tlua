//// [tests/cases/compiler/tluaMultiReturnAssignability.tlua] ////

//// [tluaMultiReturnAssignability.tlua]
function pair(): (number, string)
  return 1, "a";
end

function triple(): (number, string, boolean)
  return 1, "a", true;
end

function one(): number
  return 1;
end

function onePlusOptional(): (number, string?)
  return 1;
end

function allOptional(): (number?, string?)
end

function nothing(): void
end

// Extra returned values are dropped at call sites: assigning a many-valued
// function where fewer values are consumed is sound.
local toSingle: (() => number) = pair;
local toPair: typeof pair = triple;

// A one-valued function satisfies a pack whose tail is optional...
local optOk: typeof onePlusOptional = one;

// ...but not a pack that requires two values.
local needTwo: typeof pair = one;

// Zero values only satisfy packs that require none.
local voidOk: typeof allOptional = nothing;
local voidBad: typeof pair = nothing;
local voidBadRequired: typeof onePlusOptional = nothing;

// The first value must still match.
local firstMismatch: (() => string) = pair;


//// [tluaMultiReturnAssignability.lua]
function pair()
    return 1, "a";
end
function triple()
    return 1, "a", true;
end
function one()
    return 1;
end
function onePlusOptional()
    return 1;
end
function allOptional()
end
function nothing()
end
-- Extra returned values are dropped at call sites: assigning a many-valued
-- function where fewer values are consumed is sound.
local toSingle = pair;
local toPair = triple;
-- A one-valued function satisfies a pack whose tail is optional...
local optOk = one;
-- ...but not a pack that requires two values.
local needTwo = one;
-- Zero values only satisfy packs that require none.
local voidOk = nothing;
local voidBad = nothing;
local voidBadRequired = nothing;
-- The first value must still match.
local firstMismatch = pair;

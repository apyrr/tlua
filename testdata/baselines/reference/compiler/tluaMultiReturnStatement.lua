//// [tests/cases/compiler/tluaMultiReturnStatement.tlua] ////

//// [tluaMultiReturnStatement.tlua]
function pair(): (number, string)
  return 1, "a";
end

// Inference: the return list infers a pack without an annotation.
function inferredPair()
  return 1, "a";
end

// Inference across mixed arities pads with nil.
function mixed(cond: boolean)
  if (cond) then
    return 1, "a";
  end
  return 2;
end

// A bare return contributes zero values.
function maybeNothing(cond: boolean)
  if (cond) then
    return;
  end
  return 1, "a";
end

// `return f()` forwards the whole pack.
function forward(): (number, string)
  return pair();
end

function forwardInferred()
  return pair();
end

// A call anywhere in a single-value position truncates to the first value.
local first = pair();
local sum = pair() + 1;

// The value list evaluates non-tail calls to a single value.
function headTruncates()
  return pair(), true;
end

local i = inferredPair();
local m = mixed(true);
local n = maybeNothing(false);
local f = forward();
local g = forwardInferred();
local h = headTruncates();


//// [tluaMultiReturnStatement.lua]
function pair()
    return 1, "a";
end
-- Inference: the return list infers a pack without an annotation.
function inferredPair()
    return 1, "a";
end
-- Inference across mixed arities pads with nil.
function mixed(cond)
    if (cond) then
        return 1, "a";
    end
    return 2;
end
-- A bare return contributes zero values.
function maybeNothing(cond)
    if (cond) then
        return;
    end
    return 1, "a";
end
-- `return f()` forwards the whole pack.
function forward()
    return pair();
end
function forwardInferred()
    return pair();
end
-- A call anywhere in a single-value position truncates to the first value.
local first = pair();
local sum = pair() + 1;
-- The value list evaluates non-tail calls to a single value.
function headTruncates()
    return pair(), true;
end
local i = inferredPair();
local m = mixed(true);
local n = maybeNothing(false);
local f = forward();
local g = forwardInferred();
local h = headTruncates();

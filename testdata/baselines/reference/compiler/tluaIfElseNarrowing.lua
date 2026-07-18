//// [tests/cases/compiler/tluaIfElseNarrowing.tlua] ////

//// [tluaIfElseNarrowing.tlua]
function narrowNil(x: number | nil): number
  if x ~= nil then
    return x;
  end
  return 0;
end

function narrowElse(x: number | nil): number
  if x == nil then
    return 0;
  else
    return x;
  end
end

function narrowElseifChain(x: number | string | nil): string
  if x == nil then
    return "nil";
  elseif type(x) == "number" then
    return "" + x;
  else
    return x;
  end
end

// A function-typed condition is always truthy — existing diagnostic fires.
function alwaysTruthy(f: () => void): number
  if f then
    return 1;
  end
  return 0;
end

// Narrowing persists into while bodies.
function whileNarrow(x: number | nil): number
  local total = 0;
  while x ~= nil do
    total = total + x;
    x = nil;
  end
  return total;
end


//// [tluaIfElseNarrowing.lua]
function narrowNil(x)
    if x ~= nil then
        return x;
    end
    return 0;
end
function narrowElse(x)
    if x == nil then
        return 0;
    else
        return x;
    end
end
function narrowElseifChain(x)
    if x == nil then
        return "nil";
    elseif type(x) == "number" then
        return "" + x;
    else
        return x;
    end
end
-- A function-typed condition is always truthy — existing diagnostic fires.
function alwaysTruthy(f)
    if f then
        return 1;
    end
    return 0;
end
-- Narrowing persists into while bodies.
function whileNarrow(x)
    local total = 0;
    while x ~= nil do
        total = total + x;
        x = nil;
    end
    return total;
end

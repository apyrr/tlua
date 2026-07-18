//// [tests/cases/compiler/tluaMultiReturnTypeOptionalVariadic.tlua] ////

//// [tluaMultiReturnTypeOptionalVariadic.tlua]
// Optional elements: an absent value reads as nil.
function opt(): (number, string?)
  return 1;
end

function optBoth(): (number, string?)
  return 1, "a";
end

// The trailing `?` must parse before a brace body (eager `?` consumption).
function optBrace(): (number, string?) {
  return 2;
}

// ... and before a same-line Lua statement.
function optSameLine(cond: boolean): (number, boolean?) local x = cond;
  return 3, x;
end

// Variadic tail: any number of trailing values.
function variadic(): (number, ...string)
  return 1, "a", "b", "c";
end

function variadicEmptyTail(): (number, ...string)
  return 1;
end

// Optional element in the middle position, before a required-of-rest layout.
function optFirstOfTail(): (boolean, number?, ...string)
  return true, 1, "x";
end

local a = opt();
local b = optBoth();
local c = optBrace();
local d = optSameLine(true);
local e = variadic();


//// [tluaMultiReturnTypeOptionalVariadic.lua]
-- Optional elements: an absent value reads as nil.
function opt()
    return 1;
end
function optBoth()
    return 1, "a";
end
-- The trailing `?` must parse before a brace body (eager `?` consumption).
function optBrace() {
    return 2;
}
-- ... and before a same-line Lua statement.
function optSameLine(cond)
    local x = cond;
    return 3, x;
end
-- Variadic tail: any number of trailing values.
function variadic()
    return 1, "a", "b", "c";
end
function variadicEmptyTail()
    return 1;
end
-- Optional element in the middle position, before a required-of-rest layout.
function optFirstOfTail()
    return true, 1, "x";
end
local a = opt();
local b = optBoth();
local c = optBrace();
local d = optSameLine(true);
local e = variadic();

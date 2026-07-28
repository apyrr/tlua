//// [tests/cases/compiler/tluaDefaultParameterLowering.tlua] ////

//// [tluaDefaultParameterLowering.tlua]
// Lua has no default parameters, so `x = e` in a parameter list is lowered to a
// nil-guarded prologue assignment. The guard is `x == nil` rather than the
// idiomatic `x = x or e`, because `false` is falsy in Lua and must not be
// replaced by the default.

function simple(x = 10)
    return x
end

function multiple(a: string, b = "b", c = 3)
    return a
end

-- a later default may read an earlier parameter, which by then holds its default
function earlier(a = 1, b = a)
    return b
end

-- a default may not reference a parameter declared after it: the checker rejects
-- this with TLUA2373, so the prologue's parameter-order evaluation is never
-- observable as a silent nil
function later(a = b, b = 2)
    return a
end

-- `false` must survive: the nil guard leaves it alone
function falsy(flag = true)
    return flag
end

-- defaults coexist with a vararg and with an explicitly optional parameter
function mixed(a = 1, b?: number, ...: number)
    return a
end

-- a default on a function expression is lowered the same way
local f = function(x = 5)
    return x
end

local _ = simple()
local _ = simple(0)
local _ = falsy(false)
local _ = f()


//// [tluaDefaultParameterLowering.lua]
-- Lua has no default parameters, so `x = e` in a parameter list is lowered to a
-- nil-guarded prologue assignment. The guard is `x == nil` rather than the
-- idiomatic `x = x or e`, because `false` is falsy in Lua and must not be
-- replaced by the default.
function simple(x)
  if x == nil then
    x = 10;
  end
  return x;
end
function multiple(a, b, c)
  if b == nil then
    b = "b";
  end
  if c == nil then
    c = 3;
  end
  return a;
end
-- a later default may read an earlier parameter, which by then holds its default
function earlier(a, b)
  if a == nil then
    a = 1;
  end
  if b == nil then
    b = a;
  end
  return b;
end
-- a default may not reference a parameter declared after it: the checker rejects
-- this with TLUA2373, so the prologue's parameter-order evaluation is never
-- observable as a silent nil
function later(a, b)
  if a == nil then
    a = b;
  end
  if b == nil then
    b = 2;
  end
  return a;
end
-- `false` must survive: the nil guard leaves it alone
function falsy(flag)
  if flag == nil then
    flag = true;
  end
  return flag;
end
-- defaults coexist with a vararg and with an explicitly optional parameter
function mixed(a, b, ...)
  if a == nil then
    a = 1;
  end
  return a;
end
-- a default on a function expression is lowered the same way
local f = function(x)
  if x == nil then
    x = 5;
  end
  return x;
end;
local _ = simple();
local _ = simple(0);
local _ = falsy(false);
local _ = f();

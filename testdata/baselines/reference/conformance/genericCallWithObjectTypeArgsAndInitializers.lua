//// [tests/cases/conformance/ported/genericCallWithObjectTypeArgsAndInitializers.tlua] ////

//// [genericCallWithObjectTypeArgsAndInitializers.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/assignmentCompatibility/genericCallWithObjectTypeArgsAndInitializers.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: upstream's `T extends Number` wrapper constraint is written as tlua's primitive
-- `number` constraint, which keeps the subject: the initializer satisfies the constraint
-- but is still unsafe for an arbitrary subtype of it.
-- dropped: upstream's separate `foo2<T>(x: T = undefined)` case -- tlua has a single nil
-- value, so that case is textually identical to the `x: T = nil` case above it.

-- Generic typed parameters with initializers
function foo<T>(x: T = nil)
    return x
end -- error

function foo3<T extends number>(x: T = 1)
end -- error

function foo4<T, U extends T>(x: T, y: U = x)
end -- error

function foo5<T, U extends T>(x: U, y: T = x)
end -- ok

function foo6<T, U extends T, V extends U>(x: T, y: U, z: V = y)
end -- error

function foo7<T, U extends T, V extends U>(x: V, y: U = x)
end -- should be ok


//// [genericCallWithObjectTypeArgsAndInitializers.lua]
-- ported from tests/cases/conformance/types/typeRelationships/assignmentCompatibility/genericCallWithObjectTypeArgsAndInitializers.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: upstream's `T extends Number` wrapper constraint is written as tlua's primitive
-- `number` constraint, which keeps the subject: the initializer satisfies the constraint
-- but is still unsafe for an arbitrary subtype of it.
-- dropped: upstream's separate `foo2<T>(x: T = undefined)` case -- tlua has a single nil
-- value, so that case is textually identical to the `x: T = nil` case above it.
-- Generic typed parameters with initializers
function foo(x)
  if x == nil then
    x = nil;
  end
  return x;
end -- error
function foo3(x)
  if x == nil then
    x = 1;
  end
end -- error
function foo4(x, y)
  if y == nil then
    y = x;
  end
end -- error
function foo5(x, y)
  if y == nil then
    y = x;
  end
end -- ok
function foo6(x, y, z)
  if z == nil then
    z = y;
  end
end -- error
function foo7(x, y)
  if y == nil then
    y = x;
  end
end -- should be ok

//// [tests/cases/conformance/ported/genericCallWithObjectTypeArgsAndInitializers.tlua] ////

//// [genericCallWithObjectTypeArgsAndInitializers.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/assignmentCompatibility/genericCallWithObjectTypeArgsAndInitializers.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- dropped: Number constraint rewritten as table because the JavaScript Number library type is unavailable in tlua

-- Generic typed parameters with initializers
function foo<T>(x: T = nil)
    return x
end

function foo2<T>(x: T = nil)
    return x
end

function foo3<T extends table>(x: T = 1)
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
-- dropped: Number constraint rewritten as table because the JavaScript Number library type is unavailable in tlua
-- Generic typed parameters with initializers
function foo(x = nil)
  return x;
end
function foo2(x = nil)
  return x;
end
function foo3(x = 1)
end -- error
function foo4(x, y = x)
end -- error
function foo5(x, y = x)
end -- ok
function foo6(x, y, z = y)
end -- error
function foo7(x, y = x)
end -- should be ok

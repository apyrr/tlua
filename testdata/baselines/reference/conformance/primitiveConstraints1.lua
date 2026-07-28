//// [tests/cases/conformance/ported/primitiveConstraints1.tlua] ////

//// [primitiveConstraints1.tlua]
-- ported from tests/cases/compiler/primitiveConstraints1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo1<T extends U, U>(t: T, u: U) end
foo1<string, number>('hm', 1) -- no error

function foo2<T, U extends T>(t: T, u: U) end
foo2<number, string>(1, 'hm') -- error


//// [primitiveConstraints1.lua]
-- ported from tests/cases/compiler/primitiveConstraints1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo1(t, u)
end
foo1('hm', 1); -- no error
function foo2(t, u)
end
foo2(1, 'hm'); -- error

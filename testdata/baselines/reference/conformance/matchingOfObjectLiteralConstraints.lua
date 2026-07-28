//// [tests/cases/conformance/ported/matchingOfObjectLiteralConstraints.tlua] ////

//// [matchingOfObjectLiteralConstraints.tlua]
-- ported from tests/cases/compiler/matchingOfObjectLiteralConstraints.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo2<T, U extends { y: T }>(x: U, z: T)
end
foo2({ y = "foo" }, "foo")


//// [matchingOfObjectLiteralConstraints.lua]
-- ported from tests/cases/compiler/matchingOfObjectLiteralConstraints.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo2(x, z)
end
foo2({ y = "foo" }, "foo");

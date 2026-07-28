//// [tests/cases/conformance/ported/restParamsWithNonRestParams.tlua] ////

//// [restParamsWithNonRestParams.tlua]
-- ported from tests/cases/compiler/restParamsWithNonRestParams.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(...: number)
end
foo()

function foo2(a: string, ...: number)
end
foo2()

function foo3(a?: string, ...: number)
end
foo3()


//// [restParamsWithNonRestParams.lua]
-- ported from tests/cases/compiler/restParamsWithNonRestParams.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(...)
end
foo();
function foo2(a, ...)
end
foo2();
function foo3(a, ...)
end
foo3();

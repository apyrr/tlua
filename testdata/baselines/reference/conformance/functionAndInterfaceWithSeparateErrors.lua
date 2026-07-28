//// [tests/cases/conformance/ported/functionAndInterfaceWithSeparateErrors.tlua] ////

//// [functionAndInterfaceWithSeparateErrors.tlua]
-- ported from tests/cases/compiler/functionAndInterfaceWithSeparateErrors.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function Foo(s: string);
function Foo(n: number)
end

interface Foo {
    [s: string]: string
    prop: number
}


//// [functionAndInterfaceWithSeparateErrors.lua]
-- ported from tests/cases/compiler/functionAndInterfaceWithSeparateErrors.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function Foo(n)
end

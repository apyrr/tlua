//// [tests/cases/conformance/ported/unusedMethodsInInterface.tlua] ////

//// [unusedMethodsInInterface.tlua]
-- ported from tests/cases/compiler/unusedMethodsInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)


interface I1 {
    f1()
    f2(x: number, y: string)
}


//// [unusedMethodsInInterface.lua]
-- ported from tests/cases/compiler/unusedMethodsInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

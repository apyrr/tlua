//// [tests/cases/conformance/ported/overloadRet.tlua] ////

//// [overloadRet.tlua]
-- ported from tests/cases/compiler/overloadRet.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I {
    f(s: string): number
    f(n: number): string
    g(n: number): any
    g(n: number, m: number): string
    h(n: number): I
    h(b: boolean): number
    i(b: boolean): number
    i(b: boolean): any
}


//// [overloadRet.lua]
-- ported from tests/cases/compiler/overloadRet.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/specializedSignatureOverloadReturnTypeWithIndexers.tlua] ////

//// [specializedSignatureOverloadReturnTypeWithIndexers.tlua]
-- ported from tests/cases/compiler/specializedSignatureOverloadReturnTypeWithIndexers.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A {
    f(p: string): { [p: string]: string }
    f(p: "spec"): { [p: string]: any } -- Should be ok
}
interface B {
    f(p: string): { [p: number]: string }
    f(p: "spec"): { [p: string]: any } -- Should be ok
}
interface C {
    f(p: string): { [p: number]: string }
    f(p: "spec"): { [p: number]: any } -- Should be ok
}
interface D {
    f(p: string): { [p: string]: string }
    f(p: "spec"): { [p: number]: any } -- legacy upstream comment said "Should be error",
                                       -- but the current upstream baseline (and tlua) report
                                       -- no diagnostic here; kept for overload-return coverage.
}


//// [specializedSignatureOverloadReturnTypeWithIndexers.lua]
-- ported from tests/cases/compiler/specializedSignatureOverloadReturnTypeWithIndexers.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

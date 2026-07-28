//// [tests/cases/conformance/ported/specializedSignatureInInterface.tlua] ////

//// [specializedSignatureInInterface.tlua]
-- ported from tests/cases/compiler/specializedSignatureInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A {
    (key: string): void
}

interface B extends A {
    (key: "foo"): string
    (key: "bar"): string
}


//// [specializedSignatureInInterface.lua]
-- ported from tests/cases/compiler/specializedSignatureInInterface.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

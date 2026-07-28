//// [tests/cases/conformance/ported/genericSignatureInheritance2.tlua] ////

//// [genericSignatureInheritance2.tlua]
-- ported from tests/cases/compiler/genericSignatureInheritance2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I {
    <T>(x: T): string
}

interface I2 extends I {
    <T>(x: T): void
}


//// [genericSignatureInheritance2.lua]
-- ported from tests/cases/compiler/genericSignatureInheritance2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

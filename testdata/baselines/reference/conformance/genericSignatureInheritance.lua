//// [tests/cases/conformance/ported/genericSignatureInheritance.tlua] ////

//// [genericSignatureInheritance.tlua]
-- ported from tests/cases/compiler/genericSignatureInheritance.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I {
    <T>(x: T): string
}

interface I2 extends I {
}


//// [genericSignatureInheritance.lua]
-- ported from tests/cases/compiler/genericSignatureInheritance.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

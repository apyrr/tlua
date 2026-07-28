//// [tests/cases/conformance/ported/recursiveTypeParameterReferenceError2.tlua] ////

//// [recursiveTypeParameterReferenceError2.tlua]
-- ported from tests/cases/compiler/recursiveTypeParameterReferenceError2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A {
    <T extends A>(x: T): void
}

interface B {
    <T extends B>(x: T): void
}

interface C {
    (x: A): void
    (x: B): void
}


//// [recursiveTypeParameterReferenceError2.lua]
-- ported from tests/cases/compiler/recursiveTypeParameterReferenceError2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/duplicateConstructSignature2.tlua] ////

//// [duplicateConstructSignature2.tlua]
-- ported from tests/cases/compiler/duplicateConstructSignature2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I<T> {
    (x: T): number
    (x: T): string
}


//// [duplicateConstructSignature2.lua]
-- ported from tests/cases/compiler/duplicateConstructSignature2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/addMoreOverloadsToBaseSignature.tlua] ////

//// [addMoreOverloadsToBaseSignature.tlua]
-- ported from tests/cases/compiler/addMoreOverloadsToBaseSignature.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo {
    f(): string
}

interface Bar extends Foo {
    f(key: string): string
}


//// [addMoreOverloadsToBaseSignature.lua]
-- ported from tests/cases/compiler/addMoreOverloadsToBaseSignature.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

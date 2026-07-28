//// [tests/cases/conformance/ported/interfaceThatHidesBaseProperty.tlua] ////

//// [interfaceThatHidesBaseProperty.tlua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatHidesBaseProperty.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Base {
    x: { a: number }
}

interface Derived extends Base {
    x: {
        a: number
        b: number
    }
}


//// [interfaceThatHidesBaseProperty.lua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatHidesBaseProperty.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

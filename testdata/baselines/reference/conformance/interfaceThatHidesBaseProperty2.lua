//// [tests/cases/conformance/ported/interfaceThatHidesBaseProperty2.tlua] ////

//// [interfaceThatHidesBaseProperty2.tlua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatHidesBaseProperty2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Base {
    x: { a: number }
}

interface Derived extends Base { -- error
    x: {
        a: string
    }
}


//// [interfaceThatHidesBaseProperty2.lua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatHidesBaseProperty2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

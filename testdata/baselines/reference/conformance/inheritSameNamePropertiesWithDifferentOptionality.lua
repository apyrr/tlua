//// [tests/cases/conformance/ported/inheritSameNamePropertiesWithDifferentOptionality.tlua] ////

//// [inheritSameNamePropertiesWithDifferentOptionality.tlua]
-- ported from tests/cases/compiler/inheritSameNamePropertiesWithDifferentOptionality.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface C {
    x?: number
}

interface C2 {
    x: number
}

interface A extends C, C2 { -- error
    y: string
}


//// [inheritSameNamePropertiesWithDifferentOptionality.lua]
-- ported from tests/cases/compiler/inheritSameNamePropertiesWithDifferentOptionality.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

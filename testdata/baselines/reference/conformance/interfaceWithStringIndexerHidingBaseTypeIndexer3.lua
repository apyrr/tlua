//// [tests/cases/conformance/ported/interfaceWithStringIndexerHidingBaseTypeIndexer3.tlua] ////

//// [interfaceWithStringIndexerHidingBaseTypeIndexer3.tlua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceWithStringIndexerHidingBaseTypeIndexer3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Base {
    [x: number]: { a: number }
    1: {
        a: number; b: number;
    }
}

interface Derived extends Base {
    [x: number]: {
        a: number; b: number
    };
    2: {
        a: number;
    } -- error
}


//// [interfaceWithStringIndexerHidingBaseTypeIndexer3.lua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceWithStringIndexerHidingBaseTypeIndexer3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

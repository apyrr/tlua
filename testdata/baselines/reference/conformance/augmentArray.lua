//// [tests/cases/conformance/ported/augmentArray.tlua] ////

//// [augmentArray.tlua]
-- ported from tests/cases/compiler/augmentArray.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Array<T> {
    (): any[]
}


//// [augmentArray.lua]
-- ported from tests/cases/compiler/augmentArray.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

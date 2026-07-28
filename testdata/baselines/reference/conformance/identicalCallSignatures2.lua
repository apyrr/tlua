//// [tests/cases/conformance/ported/identicalCallSignatures2.tlua] ////

//// [identicalCallSignatures2.tlua]
-- ported from tests/cases/conformance/types/objectTypeLiteral/callSignatures/identicalCallSignatures2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua uses strict checking)

-- Normally it is an error to have multiple overloads with identical signatures in a single type declaration.
-- Here the multiple overloads come from multiple bases.

interface Base<T> {
    (x: number): string
}

interface I extends Base<string>, Base<number> {}

interface I2<T> extends Base<string>, Base<number> {}


//// [identicalCallSignatures2.lua]
-- ported from tests/cases/conformance/types/objectTypeLiteral/callSignatures/identicalCallSignatures2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua uses strict checking)

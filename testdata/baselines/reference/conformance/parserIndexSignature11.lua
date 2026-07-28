//// [tests/cases/conformance/ported/parserIndexSignature11.tlua] ////

//// [parserIndexSignature11.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/IndexSignatures/parserIndexSignature11.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)

interface I {
        [p]; -- Used to be indexer, now it is a computed property
        [p1: string];
        [p2: string, p3: number];
}


//// [parserIndexSignature11.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/IndexSignatures/parserIndexSignature11.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)

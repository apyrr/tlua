//// [tests/cases/conformance/ported/parserPropertySignature8.tlua] ////

//// [parserPropertySignature8.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/PropertySignatures/parserPropertySignature8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I {
    "H"?: any
}


//// [parserPropertySignature8.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/PropertySignatures/parserPropertySignature8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

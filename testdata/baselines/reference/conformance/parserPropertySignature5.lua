//// [tests/cases/conformance/ported/parserPropertySignature5.tlua] ////

//// [parserPropertySignature5.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/PropertySignatures/parserPropertySignature5.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua uses strict checking by default)

interface I {
    "E"
}


//// [parserPropertySignature5.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/PropertySignatures/parserPropertySignature5.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua uses strict checking by default)

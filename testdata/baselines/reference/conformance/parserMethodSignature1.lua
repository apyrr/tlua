//// [tests/cases/conformance/ported/parserMethodSignature1.tlua] ////

//// [parserMethodSignature1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature1.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

interface I {
    A()
}


//// [parserMethodSignature1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature1.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

//// [tests/cases/conformance/ported/parserMethodSignature11.tlua] ////

//// [parserMethodSignature11.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature11.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

interface I {
    2<T>()
}


//// [parserMethodSignature11.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature11.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

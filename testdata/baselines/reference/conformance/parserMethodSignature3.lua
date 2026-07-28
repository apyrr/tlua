//// [tests/cases/conformance/ported/parserMethodSignature3.tlua] ////

//// [parserMethodSignature3.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature3.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

interface I {
    C<T>()
}


//// [parserMethodSignature3.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature3.ts
-- dropped: @target: es2015 (tlua targets latest; the ES target is not a tlua concept)

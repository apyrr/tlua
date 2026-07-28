//// [tests/cases/conformance/ported/parserMethodSignature5.tlua] ////

//// [parserMethodSignature5.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature5.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

interface I {
    "E"()
}


//// [parserMethodSignature5.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature5.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

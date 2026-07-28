//// [tests/cases/conformance/ported/parserMethodSignature1.tlua] ////

//// [parserMethodSignature1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

interface I {
    A()
}


//// [parserMethodSignature1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

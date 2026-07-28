//// [tests/cases/conformance/ported/parserMethodSignature11.tlua] ////

//// [parserMethodSignature11.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature11.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

interface I {
    2<T>()
}


//// [parserMethodSignature11.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature11.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

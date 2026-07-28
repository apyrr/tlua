//// [tests/cases/conformance/ported/parserMethodSignature9.tlua] ////

//// [parserMethodSignature9.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature9.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

interface I {
    0()
}


//// [parserMethodSignature9.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/MethodSignatures/parserMethodSignature9.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

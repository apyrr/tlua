//// [tests/cases/conformance/ported/externFunc.tlua] ////

//// [externFunc.tlua]
-- ported from tests/cases/compiler/externFunc.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

declare function parseInt(s: string): number

parseInt("2")


//// [externFunc.lua]
-- ported from tests/cases/compiler/externFunc.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
parseInt("2");

//// [tests/cases/conformance/ported/typePredicatesInUnion2.tlua] ////

//// [typePredicatesInUnion2.tlua]
-- ported from tests/cases/compiler/typePredicatesInUnion2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

declare function isString(x: any): x is string
declare function isNumber(x: any): x is number
declare function f(p: typeof isString | typeof isNumber): void

f(isString)
f(isNumber)


//// [typePredicatesInUnion2.lua]
-- ported from tests/cases/compiler/typePredicatesInUnion2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
f(isString);
f(isNumber);

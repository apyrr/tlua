//// [tests/cases/conformance/ported/selfReferentialFunctionType.tlua] ////

//// [selfReferentialFunctionType.tlua]
-- ported from tests/cases/compiler/selfReferentialFunctionType.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua declaration emit is unsupported for this ambient generic-function test.
-- dropped: declaration-output coverage; checker coverage is retained.

declare function f<T>(args: typeof f<T>): T
declare function g<T = typeof g>(args: T): T
declare function h<T>(): typeof h<T>


//// [selfReferentialFunctionType.lua]
-- ported from tests/cases/compiler/selfReferentialFunctionType.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua declaration emit is unsupported for this ambient generic-function test.
-- dropped: declaration-output coverage; checker coverage is retained.

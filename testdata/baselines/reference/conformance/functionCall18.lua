//// [tests/cases/conformance/ported/functionCall18.tlua] ////

//// [functionCall18.tlua]
-- ported from tests/cases/compiler/functionCall18.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- Repro from #26835: generic and non-generic overload resolution.
declare function foo<T>(a: T, b: T): any
declare function foo(a: table): any
foo<string>("hello")


//// [functionCall18.lua]
foo("hello");

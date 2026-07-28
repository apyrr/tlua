//// [tests/cases/conformance/ported/tupleTypeInference2.tlua] ////

//// [tupleTypeInference2.tlua]
-- ported from tests/cases/compiler/tupleTypeInference2.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to latest target and strict checking)

-- Repro from #22564

type A<R> = [R] | [R, string]
declare function f<T>(x: A<T>): T
f({ nil, '' } as [never, string]) -- T: never
f({ nil, '' } as [void, string]) -- T: void

-- Repro from #22563

type B<R, S> = [R] | [R, S]
declare function g<T, U>(f: B<T, U>): U
g({ {} } as [void[]]) -- U: {}

type C<R, S> = [R[]] | [R[], S]
declare function h<T, U>(f: C<T, U>): U
h({ {} } as [void[]]) -- U: {}

-- Repro from #22562

type C2<R> = [R[]] | [R[], void]
declare function h2<T>(f: C2<T>): T
h2({ {} } as [never[]]) -- T: never
h2({ {} } as [void[]]) -- T: void


//// [tupleTypeInference2.lua]
-- ported from tests/cases/compiler/tupleTypeInference2.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to latest target and strict checking)
f({ nil, '' }); -- T: never
f({ nil, '' }); -- T: void
g({ {} }); -- U: {}
h({ {} }); -- U: {}
h2({ {} }); -- T: never
h2({ {} }); -- T: void

//// [tests/cases/conformance/ported/keyofIntersection.tlua] ////

//// [keyofIntersection.tlua]
-- ported from tests/cases/conformance/types/keyof/keyofIntersection.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (tlua keeps strict checking enabled)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.


type A = { a: string }
type B = { b: string }

type T01 = keyof (A & B) -- "a" | "b"
type T02<T> = keyof (T & B) -- "b" | keyof T
type T03<U> = keyof (A & U) -- "a" | keyof U
type T04<T, U> = keyof (T & U) -- keyof T | keyof U
type T05 = T02<A> -- "a" | "b"
type T06 = T03<B> -- "a" | "b"
type T07 = T04<A, B> -- "a" | "b"

-- Repros from #22291

type Example1<T extends string, U extends string> = keyof (Record<T, any> & Record<U, any>)
type Result1 = Example1<'x', 'y'> -- "x" | "y"

type Result2 = keyof (Record<'x', any> & Record<'y', any>) -- "x" | "y"

type Example3<T extends string> = keyof (Record<T, any>)
type Result3 = Example3<'x' | 'y'> -- "x" | "y"

type Example4<T extends string, U extends string> = (Record<T, any> & Record<U, any>)
type Result4 = keyof Example4<'x', 'y'> -- "x" | "y"

type Example5<T, U> = keyof (T & U)
type Result5 = Example5<Record<'x', any>, Record<'y', any>> -- "x" | "y"


//// [keyofIntersection.lua]
-- ported from tests/cases/conformance/types/keyof/keyofIntersection.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (tlua keeps strict checking enabled)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.

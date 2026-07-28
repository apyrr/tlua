//// [tests/cases/conformance/ported/deferredLookupTypeResolution2.tlua] ////

//// [deferredLookupTypeResolution2.tlua]
-- ported from tests/cases/compiler/deferredLookupTypeResolution2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output (unsupported for Lua modules)
-- compiler gap: the checker can report a wrong-false result for this tuple mapped/indexed lookup.

-- Repro from #17456

type StringContains<S extends string, L extends string> = ({ [K in S]: 'true' } & { [key: string]: 'false' })[L]

type ObjectHasKey<O, L extends string> = StringContains<Extract<keyof O, string>, L>

type A<T> = ObjectHasKey<T, '1'>

type B = ObjectHasKey<[string, number], '1'>
type C = ObjectHasKey<[string, number], '2'>
type D = A<[string]>

-- Error, "false" not handled
type E<T> = { true: 'true' }[ObjectHasKey<T, '1'>]

type Juxtapose<T> = ({ true: 'otherwise' } & { [k: string]: 'true' })[ObjectHasKey<T, '1'>]

-- Error, "otherwise" is missing
type DeepError<T> = { true: 'true' }[Juxtapose<T>]

type DeepOK<T> = { true: 'true', otherwise: 'false' }[Juxtapose<T>]


//// [deferredLookupTypeResolution2.lua]
-- ported from tests/cases/compiler/deferredLookupTypeResolution2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output (unsupported for Lua modules)
-- compiler gap: the checker can report a wrong-false result for this tuple mapped/indexed lookup.

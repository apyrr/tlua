//// [tests/cases/conformance/ported/anyMappedTypesError.tlua] ////

//// [anyMappedTypesError.tlua]
-- ported from tests/cases/compiler/anyMappedTypesError.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

type Foo = { [P in "bar"] }


//// [anyMappedTypesError.lua]
-- ported from tests/cases/compiler/anyMappedTypesError.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/genericNumberIndex.tlua] ////

//// [genericNumberIndex.tlua]
-- ported from tests/cases/compiler/genericNumberIndex.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

type X<I extends number> = ['a'][I]


//// [genericNumberIndex.lua]
-- ported from tests/cases/compiler/genericNumberIndex.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

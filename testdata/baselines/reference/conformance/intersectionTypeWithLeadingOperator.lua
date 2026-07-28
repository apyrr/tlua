//// [tests/cases/conformance/ported/intersectionTypeWithLeadingOperator.tlua] ////

//// [intersectionTypeWithLeadingOperator.tlua]
-- ported from tests/cases/compiler/intersectionTypeWithLeadingOperator.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

type A = & string
type B =
  & { foo: string }
  & { bar: number }

type C = [& { foo: 1 } & { bar: 2 }, & { foo: 3 } & { bar: 4 }]


//// [intersectionTypeWithLeadingOperator.lua]
-- ported from tests/cases/compiler/intersectionTypeWithLeadingOperator.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

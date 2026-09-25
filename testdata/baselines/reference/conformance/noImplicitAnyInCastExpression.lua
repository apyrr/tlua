//// [tests/cases/conformance/ported/noImplicitAnyInCastExpression.tlua] ////

//// [noImplicitAnyInCastExpression.tlua]
-- ported from tests/cases/compiler/noImplicitAnyInCastExpression.ts
-- dropped: @target: es2015 and @noImplicitAny: true directives (tlua defaults to esnext and strict checking)
-- dropped: bare cast expression statements (bound to local `_` because tlua permits only calls and assignments as bare expressions)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)

-- verify no implicit-any errors are reported with cast expressions

interface IFoo {
    a: number
    b: string
}

-- Expr type not assignable to target type
local _ = ({ a = nil } as IFoo)

-- Expr type assignable to target type
local _ = ({ a = 2, b = nil } as IFoo)

-- Neither type is assignable to the other
local _ = ({ c = nil } as IFoo)


//// [noImplicitAnyInCastExpression.lua]
-- ported from tests/cases/compiler/noImplicitAnyInCastExpression.ts
-- dropped: @target: es2015 and @noImplicitAny: true directives (tlua defaults to esnext and strict checking)
-- dropped: bare cast expression statements (bound to local `_` because tlua permits only calls and assignments as bare expressions)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)
-- Expr type not assignable to target type
local _ = { a = nil };
-- Expr type assignable to target type
local _ = { a = 2, b = nil };
-- Neither type is assignable to the other
local _ = { c = nil };

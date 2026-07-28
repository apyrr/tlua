//// [tests/cases/conformance/ported/parserMemberAccessExpression1.tlua] ////

//// [parserMemberAccessExpression1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/Generics/parserMemberAccessExpression1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

local _ = Foo<T>()
local _ = Foo.Bar<T>()
local _ = Foo<T>.Bar()
local _ = Foo<T>.Bar<T>()


//// [parserMemberAccessExpression1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/Generics/parserMemberAccessExpression1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
local _ = Foo();
local _ = Foo.Bar();
local _ = Foo.Bar();
local _ = Foo.Bar();

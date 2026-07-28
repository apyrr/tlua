//// [tests/cases/conformance/ported/unionTypeParameterInference.tlua] ////

//// [unionTypeParameterInference.tlua]
-- ported from tests/cases/compiler/unionTypeParameterInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo<T> { prop: T }

declare lift: <U>(value: U | Foo<U>) => Foo<U>

function unlift<U>(value: U | Foo<U>): U
    return lift(value).prop
end


//// [unionTypeParameterInference.lua]
-- ported from tests/cases/compiler/unionTypeParameterInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function unlift(value)
  return lift(value).prop;
end

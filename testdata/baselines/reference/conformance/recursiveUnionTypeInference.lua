//// [tests/cases/conformance/ported/recursiveUnionTypeInference.tlua] ////

//// [recursiveUnionTypeInference.tlua]
-- ported from tests/cases/compiler/recursiveUnionTypeInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo<T> {
    x: T
}

function bar<T>(x: Foo<T> | string): T
    return bar(x)
end


//// [recursiveUnionTypeInference.lua]
-- ported from tests/cases/compiler/recursiveUnionTypeInference.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function bar(x)
  return bar(x);
end

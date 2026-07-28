//// [tests/cases/conformance/ported/typeParameterHasSelfAsConstraint.tlua] ////

//// [typeParameterHasSelfAsConstraint.tlua]
-- ported from tests/cases/compiler/typeParameterHasSelfAsConstraint.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo<T extends T>(x: T): number
    return x
end


//// [typeParameterHasSelfAsConstraint.lua]
-- ported from tests/cases/compiler/typeParameterHasSelfAsConstraint.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(x)
  return x;
end

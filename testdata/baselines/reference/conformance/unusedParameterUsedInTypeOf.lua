//// [tests/cases/conformance/ported/unusedParameterUsedInTypeOf.tlua] ////

//// [unusedParameterUsedInTypeOf.tlua]
-- ported from tests/cases/compiler/unusedParameterUsedInTypeOf.ts
-- dropped: @target: ES5, ES2015 directive (tlua defaults to latest target)

-- @noUnusedLocals: true
-- @noUnusedParameters: true

function f1(a: number, b: typeof a)
    return b
end


//// [unusedParameterUsedInTypeOf.lua]
-- ported from tests/cases/compiler/unusedParameterUsedInTypeOf.ts
-- dropped: @target: ES5, ES2015 directive (tlua defaults to latest target)
-- @noUnusedLocals: true
-- @noUnusedParameters: true
function f1(a, b)
  return b;
end

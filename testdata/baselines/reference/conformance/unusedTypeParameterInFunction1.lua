//// [tests/cases/conformance/ported/unusedTypeParameterInFunction1.tlua] ////

//// [unusedTypeParameterInFunction1.tlua]
-- ported from tests/cases/compiler/unusedTypeParameterInFunction1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

-- @noUnusedLocals: true
-- @noUnusedParameters: true
-- compiler gap: tlua does not report the upstream noUnusedLocals/noUnusedParameters diagnostic for the unused type parameter

function f1<T>()
end


//// [unusedTypeParameterInFunction1.lua]
-- ported from tests/cases/compiler/unusedTypeParameterInFunction1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- @noUnusedLocals: true
-- @noUnusedParameters: true
-- compiler gap: tlua does not report the upstream noUnusedLocals/noUnusedParameters diagnostic for the unused type parameter
function f1()
end

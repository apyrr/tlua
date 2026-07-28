//// [tests/cases/conformance/ported/noUnusedLocals_typeParameterMergedWithParameter.tlua] ////

//// [noUnusedLocals_typeParameterMergedWithParameter.tlua]
-- ported from tests/cases/compiler/noUnusedLocals_typeParameterMergedWithParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

-- @noUnusedLocals: true
-- @noUnusedParameters: true
-- compiler gap: noUnusedLocals/noUnusedParameters do not report the four upstream TS6133 unused-declaration diagnostics in these type-parameter/value-name collision cases.

function useNone<T>(T: number) end

function useParam<T>(T: number)
    return T
end

function useTypeParam<T>(T: T) end

function useBoth<T>(T: T)
    return T
end


//// [noUnusedLocals_typeParameterMergedWithParameter.lua]
-- ported from tests/cases/compiler/noUnusedLocals_typeParameterMergedWithParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- @noUnusedLocals: true
-- @noUnusedParameters: true
-- compiler gap: noUnusedLocals/noUnusedParameters do not report the four upstream TS6133 unused-declaration diagnostics in these type-parameter/value-name collision cases.
function useNone(T)
end
function useParam(T)
  return T;
end
function useTypeParam(T)
end
function useBoth(T)
  return T;
end

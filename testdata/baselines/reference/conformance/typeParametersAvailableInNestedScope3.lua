//// [tests/cases/conformance/ported/typeParametersAvailableInNestedScope3.tlua] ////

//// [typeParametersAvailableInNestedScope3.tlua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParametersAvailableInNestedScope3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.




function foo<T>(v: T)
    local function a<T>(a: T)
        return a
    end
    local function b(): T
        return v
    end

    local function c<T>(v: T)
        local function a<T>(a: T)
            return a
        end
        local function b(): T
            return v
        end
        return { a = a, b = b }
    end

    return { a = a, b = b, c = c }
end


//// [typeParametersAvailableInNestedScope3.lua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParametersAvailableInNestedScope3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true -- tlua has declaration-emit machinery, but it reports
--   TLUA100054 for every .tlua source, so restoring the directive would replace this test's
--   subject with that one diagnostic. tluaModuleDeclarationEmitUnsupported.tlua covers it.
function foo(v)
  local function a(a)
    return a;
  end
  local function b()
    return v;
  end
  local function c(v)
    local function a(a)
      return a;
    end
    local function b()
      return v;
    end
    return { a = a, b = b };
  end
  return { a = a, b = b, c = c };
end

//// [tests/cases/conformance/ported/typeParametersAvailableInNestedScope3.tlua] ////

//// [typeParametersAvailableInNestedScope3.tlua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParametersAvailableInNestedScope3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true (tlua does not support declaration emit for this test)



function foo<T>(v: T)
    function a<T>(a: T)
        return a
    end
    function b(): T
        return v
    end

    function c<T>(v: T)
        function a<T>(a: T)
            return a
        end
        function b(): T
            return v
        end
        return { a = a, b = b }
    end

    return { a = a, b = b, c = c }
end


//// [typeParametersAvailableInNestedScope3.lua]
-- ported from tests/cases/conformance/types/typeParameters/typeParameterLists/typeParametersAvailableInNestedScope3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @declaration: true (tlua does not support declaration emit for this test)
function foo(v)
  function a(a)
    return a;
  end
  function b()
    return v;
  end
  function c(v)
    function a(a)
      return a;
    end
    function b()
      return v;
    end
    return { a = a, b = b };
  end
  return { a = a, b = b, c = c };
end

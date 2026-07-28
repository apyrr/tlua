//// [tests/cases/conformance/ported/truthinessCallExpressionCoercion3.tlua] ////

//// [truthinessCallExpressionCoercion3.tlua]
-- ported from tests/cases/compiler/truthinessCallExpressionCoercion3.ts
-- dropped: @target: es2015, @strictNullChecks: true, and @lib: esnext,dom directives (tlua defaults to latest target, strict checking, and luajit)

-- from #41640, based on an example in ant-design
interface I {
    always(): void
}

function f(result: unknown)
    if (result as I).always then
        return result
    end
end
function g(result: unknown)
    if ((result as I)).always then
        return result
    end
end


//// [truthinessCallExpressionCoercion3.lua]
-- ported from tests/cases/compiler/truthinessCallExpressionCoercion3.ts
-- dropped: @target: es2015, @strictNullChecks: true, and @lib: esnext,dom directives (tlua defaults to latest target, strict checking, and luajit)
function f(result)
  if result.always then
    return result;
  end
end
function g(result)
  if result.always then
    return result;
  end
end

//// [tests/cases/conformance/ported/invocationExpressionInFunctionParameter.tlua] ////

//// [invocationExpressionInFunctionParameter.tlua]
-- ported from tests/cases/compiler/invocationExpressionInFunctionParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: valid default-parameter source would otherwise emit an initializer in the Lua parameter list


function foo1(val: string)
end

function foo3(x = foo1(123))
end


//// [invocationExpressionInFunctionParameter.lua]
-- ported from tests/cases/compiler/invocationExpressionInFunctionParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: valid default-parameter source would otherwise emit an initializer in the Lua parameter list
function foo1(val)
end
function foo3(x)
  if x == nil then
    x = foo1(123);
  end
end

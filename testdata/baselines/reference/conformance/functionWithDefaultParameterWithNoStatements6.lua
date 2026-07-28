//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements6.tlua] ////

//// [functionWithDefaultParameterWithNoStatements6.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a = true)
end

function bar(a = true)
end


//// [functionWithDefaultParameterWithNoStatements6.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a)
  if a == nil then
    a = true;
  end
end
function bar(a)
  if a == nil then
    a = true;
  end
end

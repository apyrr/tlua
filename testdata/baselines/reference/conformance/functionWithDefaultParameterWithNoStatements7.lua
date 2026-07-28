//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements7.tlua] ////

//// [functionWithDefaultParameterWithNoStatements7.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements7.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(a = false)
end

function bar(a = false)
end


//// [functionWithDefaultParameterWithNoStatements7.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements7.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a)
  if a == nil then
    a = false;
  end
end
function bar(a)
  if a == nil then
    a = false;
  end
end

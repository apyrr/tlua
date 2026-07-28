//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements4.tlua] ////

//// [functionWithDefaultParameterWithNoStatements4.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)


function foo(a = ``)
end

function bar(a = ``)
end


//// [functionWithDefaultParameterWithNoStatements4.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(a)
  if a == nil then
    a = "";
  end
end
function bar(a)
  if a == nil then
    a = "";
  end
end

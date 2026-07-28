//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements8.tlua] ////

//// [functionWithDefaultParameterWithNoStatements8.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements8.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: upstream's `a = undefined` initializer maps to tlua's single nil value.


function foo(a = nil)
end

function bar(a = nil)
end


//// [functionWithDefaultParameterWithNoStatements8.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements8.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- note: upstream's `a = undefined` initializer maps to tlua's single nil value.
function foo(a)
  if a == nil then
    a = nil;
  end
end
function bar(a)
  if a == nil then
    a = nil;
  end
end

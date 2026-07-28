//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements9.tlua] ////

//// [functionWithDefaultParameterWithNoStatements9.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements9.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript's `console.log` global; replaced with Lua's `print` while preserving a function-valued default expression

function foo(a = print)
end

function bar(a = print)
end


//// [functionWithDefaultParameterWithNoStatements9.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements9.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript's `console.log` global; replaced with Lua's `print` while preserving a function-valued default expression
function foo(a)
  if a == nil then
    a = print;
  end
end
function bar(a)
  if a == nil then
    a = print;
  end
end

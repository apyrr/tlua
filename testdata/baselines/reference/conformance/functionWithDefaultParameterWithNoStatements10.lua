//// [tests/cases/conformance/ported/functionWithDefaultParameterWithNoStatements10.tlua] ////

//// [functionWithDefaultParameterWithNoStatements10.tlua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements10.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: array literal defaults (rewritten as positional table literals; array literals are deleted in tlua)

function foo(a = { 0 })
end

function bar(a = { 0 })
end


//// [functionWithDefaultParameterWithNoStatements10.lua]
-- ported from tests/cases/compiler/functionWithDefaultParameterWithNoStatements10.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: array literal defaults (rewritten as positional table literals; array literals are deleted in tlua)
function foo(a)
  if a == nil then
    a = { 0 };
  end
end
function bar(a)
  if a == nil then
    a = { 0 };
  end
end

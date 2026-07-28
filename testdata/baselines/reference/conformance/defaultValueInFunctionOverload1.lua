//// [tests/cases/conformance/ported/defaultValueInFunctionOverload1.tlua] ////

//// [defaultValueInFunctionOverload1.tlua]
-- ported from tests/cases/compiler/defaultValueInFunctionOverload1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: a bodiless overload signature must be terminated by `;` in tlua, since
--   `end` is what otherwise closes a function body


function foo(x: string = '');
function foo(x = '')
end


//// [defaultValueInFunctionOverload1.lua]
-- ported from tests/cases/compiler/defaultValueInFunctionOverload1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: a bodiless overload signature must be terminated by `;` in tlua, since
--   `end` is what otherwise closes a function body
function foo(x)
  if x == nil then
    x = '';
  end
end

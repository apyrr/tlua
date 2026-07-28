//// [tests/cases/conformance/ported/emitDefaultParametersFunction.tlua] ////

//// [emitDefaultParametersFunction.tlua]
-- ported from tests/cases/conformance/es6/defaultParameters/emitDefaultParametersFunction.ts
-- dropped: dual ES5/ES2015 down-level emit comparison (tlua has a single Lua
--   emit target, no ES5 transform); kept the default-parameter TYPING only
-- dropped: named rest parameter `...rest` (the legacy TS rest form is rejected
--   in tlua); rewritten as the unnamed vararg `...`


function foo(x: string, y = 10): void
end

function baz(x: string, y = 5, ...): void
end

function bar(y = 10): void
end

function bar1(y = 10, ...): void
end


//// [emitDefaultParametersFunction.lua]
-- ported from tests/cases/conformance/es6/defaultParameters/emitDefaultParametersFunction.ts
-- dropped: dual ES5/ES2015 down-level emit comparison (tlua has a single Lua
--   emit target, no ES5 transform); kept the default-parameter TYPING only
-- dropped: named rest parameter `...rest` (the legacy TS rest form is rejected
--   in tlua); rewritten as the unnamed vararg `...`
function foo(x, y)
  if y == nil then
    y = 10;
  end
end
function baz(x, y, ...)
  if y == nil then
    y = 5;
  end
end
function bar(y)
  if y == nil then
    y = 10;
  end
end
function bar1(y, ...)
  if y == nil then
    y = 10;
  end
end

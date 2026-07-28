//// [tests/cases/conformance/ported/functionOverloads8.tlua] ////

//// [functionOverloads8.tlua]
-- ported from tests/cases/compiler/functionOverloads8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo();
function foo(foo: string);
function foo(foo?: any)
    return ""
end


//// [functionOverloads8.lua]
-- ported from tests/cases/compiler/functionOverloads8.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(foo)
  return "";
end

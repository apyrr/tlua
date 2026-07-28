//// [tests/cases/conformance/ported/functionOverloads18.tlua] ////

//// [functionOverloads18.tlua]
-- ported from tests/cases/compiler/functionOverloads18.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(bar: { a: number });
function foo(bar: { a: string })
    return { a = "" }
end


//// [functionOverloads18.lua]
-- ported from tests/cases/compiler/functionOverloads18.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(bar)
  return { a = "" };
end

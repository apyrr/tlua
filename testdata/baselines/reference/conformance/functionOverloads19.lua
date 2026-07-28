//// [tests/cases/conformance/ported/functionOverloads19.tlua] ////

//// [functionOverloads19.tlua]
-- ported from tests/cases/compiler/functionOverloads19.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(bar: { b: string });
function foo(bar: { a: string });
function foo(bar: { a: any })
    return { a = "" }
end


//// [functionOverloads19.lua]
-- ported from tests/cases/compiler/functionOverloads19.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(bar)
  return { a = "" };
end

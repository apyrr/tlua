//// [tests/cases/conformance/ported/functionOverloads22.tlua] ////

//// [functionOverloads22.tlua]
-- ported from tests/cases/compiler/functionOverloads22.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(bar: number): { a: number }[];
function foo(bar: string): { a: number, b: string }[];
function foo(bar: any): { a: any, b?: any }[]
    return { { a = "" } }
end


//// [functionOverloads22.lua]
-- ported from tests/cases/compiler/functionOverloads22.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(bar)
  return { { a = "" } };
end

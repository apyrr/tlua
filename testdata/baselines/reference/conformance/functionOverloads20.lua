//// [tests/cases/conformance/ported/functionOverloads20.tlua] ////

//// [functionOverloads20.tlua]
-- ported from tests/cases/compiler/functionOverloads20.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(bar: { a: number }): number;
function foo(bar: { a: string }): string;
function foo(bar: { a: any }): string
    return ""
end


//// [functionOverloads20.lua]
-- ported from tests/cases/compiler/functionOverloads20.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(bar)
  return "";
end

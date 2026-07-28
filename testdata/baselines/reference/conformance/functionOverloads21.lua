//// [tests/cases/conformance/ported/functionOverloads21.tlua] ////

//// [functionOverloads21.tlua]
-- ported from tests/cases/compiler/functionOverloads21.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

function foo(bar: { a: number }[]);
function foo(bar: { a: number, b: string }[]);
function foo(bar: { a: any, b?: string }[]): any
    return 0
end


//// [functionOverloads21.lua]
-- ported from tests/cases/compiler/functionOverloads21.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)
function foo(bar)
  return 0;
end

//// [tests/cases/conformance/ported/functionOverloads16.tlua] ////

//// [functionOverloads16.tlua]
-- ported from tests/cases/compiler/functionOverloads16.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)

function foo(foo: { a: string }): string;
function foo(foo: { a: string }): number;
function foo(foo: { a: string, b?: number }): any
    return ""
end


//// [functionOverloads16.lua]
-- ported from tests/cases/compiler/functionOverloads16.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
function foo(foo)
  return "";
end

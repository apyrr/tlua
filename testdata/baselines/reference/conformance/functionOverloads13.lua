//// [tests/cases/conformance/ported/functionOverloads13.tlua] ////

//// [functionOverloads13.tlua]
-- ported from tests/cases/compiler/functionOverloads13.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)

function foo(bar: number): string;
function foo(bar: number): number;
function foo(bar?: number): any
    return ""
end


//// [functionOverloads13.lua]
-- ported from tests/cases/compiler/functionOverloads13.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
function foo(bar)
  return "";
end

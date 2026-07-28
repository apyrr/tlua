//// [tests/cases/conformance/ported/functionOverloads14.tlua] ////

//// [functionOverloads14.tlua]
-- ported from tests/cases/compiler/functionOverloads14.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)

function foo(): { a: number };
function foo(): { a: string };
function foo(): { a: any }
    return { a = 1 }
end


//// [functionOverloads14.lua]
-- ported from tests/cases/compiler/functionOverloads14.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict checking)
function foo()
  return { a = 1 };
end

//// [tests/cases/conformance/ported/functionOverloads17.tlua] ////

//// [functionOverloads17.tlua]
-- ported from tests/cases/compiler/functionOverloads17.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)

function foo(): { a: number };
function foo(): { a: string }
    return { a = "" }
end


//// [functionOverloads17.lua]
-- ported from tests/cases/compiler/functionOverloads17.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
function foo()
  return { a = "" };
end

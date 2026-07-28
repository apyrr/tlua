//// [tests/cases/conformance/ported/functionOverloads8.tlua] ////

//// [functionOverloads8.tlua]
-- ported from tests/cases/compiler/functionOverloads8.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)

function foo();
function foo(foo: string);
function foo(foo?: any)
    return ""
end


//// [functionOverloads8.lua]
-- ported from tests/cases/compiler/functionOverloads8.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
function foo(foo)
  return "";
end

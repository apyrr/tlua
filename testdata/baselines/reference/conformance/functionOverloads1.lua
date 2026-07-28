//// [tests/cases/conformance/ported/functionOverloads1.tlua] ////

//// [functionOverloads1.tlua]
-- ported from tests/cases/compiler/functionOverloads1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
-- dropped: the incidental bare-expression statement, rewritten as a discarded local value because tlua only permits call and assignment expression statements.

function foo();
local _ = 1 + 1;
function foo(): string
    return "a"
end


//// [functionOverloads1.lua]
-- ported from tests/cases/compiler/functionOverloads1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to esnext and strict)
-- dropped: the incidental bare-expression statement, rewritten as a discarded local value because tlua only permits call and assignment expression statements.
local _ = 1 + 1;
function foo()
  return "a";
end

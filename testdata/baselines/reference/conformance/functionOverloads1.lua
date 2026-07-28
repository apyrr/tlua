//// [tests/cases/conformance/ported/functionOverloads1.tlua] ////

//// [functionOverloads1.tlua]
-- ported from tests/cases/compiler/functionOverloads1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's incidental bare-expression statement `1+1;` is rewritten as a
-- discarded local because tlua only permits call and assignment expression statements
-- (TLUA100057). It still separates the overload signature from the implementation,
-- which is the subject of the test.

function foo();
local _ = 1 + 1;
function foo(): string
    return "a"
end


//// [functionOverloads1.lua]
-- ported from tests/cases/compiler/functionOverloads1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's incidental bare-expression statement `1+1;` is rewritten as a
-- discarded local because tlua only permits call and assignment expression statements
-- (TLUA100057). It still separates the overload signature from the implementation,
-- which is the subject of the test.
local _ = 1 + 1;
function foo()
  return "a";
end

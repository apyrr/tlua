//// [tests/cases/conformance/ported/contextualTyping30.tlua] ////

//// [contextualTyping30.tlua]
-- ported from tests/cases/compiler/contextualTyping30.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(param: number[])
end

foo({ 1, "a" })


//// [contextualTyping30.lua]
-- ported from tests/cases/compiler/contextualTyping30.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(param)
end
foo({ 1, "a" });

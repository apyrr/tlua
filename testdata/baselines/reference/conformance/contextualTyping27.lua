//// [tests/cases/conformance/ported/contextualTyping27.tlua] ////

//// [contextualTyping27.tlua]
-- ported from tests/cases/compiler/contextualTyping27.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(param: { id: number })
end

foo(<{ id: number }>({}))


//// [contextualTyping27.lua]
-- ported from tests/cases/compiler/contextualTyping27.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo(param)
end
foo(({}));

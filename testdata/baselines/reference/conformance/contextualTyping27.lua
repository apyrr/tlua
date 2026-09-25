//// [tests/cases/conformance/ported/contextualTyping27.tlua] ////

//// [contextualTyping27.tlua]
-- ported from tests/cases/compiler/contextualTyping27.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)

function foo(param: { id: number })
end

foo(({}) as { id: number })


//// [contextualTyping27.lua]
-- ported from tests/cases/compiler/contextualTyping27.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)
function foo(param)
end
foo(({}));

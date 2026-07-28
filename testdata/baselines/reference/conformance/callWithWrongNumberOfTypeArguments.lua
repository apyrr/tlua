//// [tests/cases/conformance/ported/callWithWrongNumberOfTypeArguments.tlua] ////

//// [callWithWrongNumberOfTypeArguments.tlua]
-- ported from tests/cases/compiler/callWithWrongNumberOfTypeArguments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f<T, U>()
end

f<number>()
f<number, string>()
f<number, string, number>()


//// [callWithWrongNumberOfTypeArguments.lua]
-- ported from tests/cases/compiler/callWithWrongNumberOfTypeArguments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f()
end
f();
f();
f();

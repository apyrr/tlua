//// [tests/cases/conformance/ported/fixTypeParameterInSignatureWithRestParameters.tlua] ////

//// [fixTypeParameterInSignatureWithRestParameters.tlua]
-- ported from tests/cases/compiler/fixTypeParameterInSignatureWithRestParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function bar<T>(item1: T, item2: T)
end

bar(1, "") -- Should be ok


//// [fixTypeParameterInSignatureWithRestParameters.lua]
-- ported from tests/cases/compiler/fixTypeParameterInSignatureWithRestParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function bar(item1, item2)
end
bar(1, ""); -- Should be ok

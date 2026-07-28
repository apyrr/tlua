//// [tests/cases/conformance/ported/duplicateTypeParameters1.tlua] ////

//// [duplicateTypeParameters1.tlua]
-- ported from tests/cases/compiler/duplicateTypeParameters1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function A<X, X>()
end


//// [duplicateTypeParameters1.lua]
-- ported from tests/cases/compiler/duplicateTypeParameters1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function A()
end

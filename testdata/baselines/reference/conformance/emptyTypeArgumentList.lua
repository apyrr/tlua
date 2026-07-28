//// [tests/cases/conformance/ported/emptyTypeArgumentList.tlua] ////

//// [emptyTypeArgumentList.tlua]
-- ported from tests/cases/compiler/emptyTypeArgumentList.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo<T>()
end
foo<>()

-- https://github.com/microsoft/TypeScript/issues/33041
function noParams()
end
noParams<>()


//// [emptyTypeArgumentList.lua]
-- ported from tests/cases/compiler/emptyTypeArgumentList.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo()
end
foo();
-- https://github.com/microsoft/TypeScript/issues/33041
function noParams()
end
noParams();

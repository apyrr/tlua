//// [tests/cases/conformance/ported/FunctionDeclaration4.tlua] ////

//// [FunctionDeclaration4.tlua]
-- ported from tests/cases/compiler/FunctionDeclaration4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo();
function bar()
end


//// [FunctionDeclaration4.lua]
-- ported from tests/cases/compiler/FunctionDeclaration4.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function bar()
end

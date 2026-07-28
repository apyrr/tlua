//// [tests/cases/conformance/ported/functionOverloadImplementationOfWrongName.tlua] ////

//// [functionOverloadImplementationOfWrongName.tlua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo(x);
function foo(x, y);
function bar()
end


//// [functionOverloadImplementationOfWrongName.lua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function bar()
end

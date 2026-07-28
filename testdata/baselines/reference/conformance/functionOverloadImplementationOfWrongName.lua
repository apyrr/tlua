//// [tests/cases/conformance/ported/functionOverloadImplementationOfWrongName.tlua] ////

//// [functionOverloadImplementationOfWrongName.tlua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)

function foo(x: any): any;
function foo(x: any, y: any): any;
function bar(): any
end


//// [functionOverloadImplementationOfWrongName.lua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)
function bar()
end

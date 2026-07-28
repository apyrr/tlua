//// [tests/cases/conformance/ported/functionOverloadImplementationOfWrongName2.tlua] ////

//// [functionOverloadImplementationOfWrongName2.tlua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName2.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)

function foo(x: any): any;
function bar(): any
end
function foo(x: any, y: any): any;


//// [functionOverloadImplementationOfWrongName2.lua]
-- ported from tests/cases/compiler/functionOverloadImplementationOfWrongName2.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)
function bar()
end

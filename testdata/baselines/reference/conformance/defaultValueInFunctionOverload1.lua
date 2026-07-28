//// [tests/cases/conformance/ported/defaultValueInFunctionOverload1.tlua] ////

//// [defaultValueInFunctionOverload1.tlua]
-- ported from tests/cases/compiler/defaultValueInFunctionOverload1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)

function foo(x: string = ''): void;
function foo(x = ''): void
end


//// [defaultValueInFunctionOverload1.lua]
-- ported from tests/cases/compiler/defaultValueInFunctionOverload1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua uses latest target and strict checking)
function foo(x = '')
end

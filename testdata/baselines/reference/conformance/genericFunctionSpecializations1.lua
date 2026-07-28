//// [tests/cases/conformance/ported/genericFunctionSpecializations1.tlua] ////

//// [genericFunctionSpecializations1.tlua]
-- ported from tests/cases/compiler/genericFunctionSpecializations1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)
-- dropped: the JavaScript wrapper type `String`, rewritten to tlua's primitive `string` constraint.

function foo3<T>(test: string); -- error
function foo3<T>(test: T)
end

function foo4<T>(test: string); -- valid
function foo4<T extends string>(test: T)
end


//// [genericFunctionSpecializations1.lua]
-- ported from tests/cases/compiler/genericFunctionSpecializations1.ts
-- dropped: @target: es2015 and @strict: false directives (tlua defaults to latest target and strict checking)
-- dropped: the JavaScript wrapper type `String`, rewritten to tlua's primitive `string` constraint.
function foo3(test)
end
function foo4(test)
end

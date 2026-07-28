//// [tests/cases/conformance/ported/genericFunctionSpecializations1.tlua] ////

//// [genericFunctionSpecializations1.tlua]
-- ported from tests/cases/compiler/genericFunctionSpecializations1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's JavaScript wrapper type `String` is written as tlua's primitive
-- `string` constraint.
-- note: the inline `-- error` / `-- valid` markers are upstream's and are stale there
-- too -- the current upstream baseline reports no diagnostics for either pair.

function foo3<T>(test: string); -- error
function foo3<T>(test: T)
end

function foo4<T>(test: string); -- valid
function foo4<T extends string>(test: T)
end


//// [genericFunctionSpecializations1.lua]
-- ported from tests/cases/compiler/genericFunctionSpecializations1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's JavaScript wrapper type `String` is written as tlua's primitive
-- `string` constraint.
-- note: the inline `-- error` / `-- valid` markers are upstream's and are stale there
-- too -- the current upstream baseline reports no diagnostics for either pair.
function foo3(test)
end
function foo4(test)
end

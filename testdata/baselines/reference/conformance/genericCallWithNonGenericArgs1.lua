//// [tests/cases/conformance/ported/genericCallWithNonGenericArgs1.tlua] ////

//// [genericCallWithNonGenericArgs1.tlua]
-- ported from tests/cases/compiler/genericCallWithNonGenericArgs1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f<T>(x: any)
end

f<any>(nil)


//// [genericCallWithNonGenericArgs1.lua]
-- ported from tests/cases/compiler/genericCallWithNonGenericArgs1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f(x)
end
f(nil);

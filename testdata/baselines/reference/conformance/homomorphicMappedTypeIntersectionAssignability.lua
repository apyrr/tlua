//// [tests/cases/conformance/ported/homomorphicMappedTypeIntersectionAssignability.tlua] ////

//// [homomorphicMappedTypeIntersectionAssignability.tlua]
-- ported from tests/cases/compiler/homomorphicMappedTypeIntersectionAssignability.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to latest target and strict checking)

function f<TType>(
    a: { weak?: string } & Readonly<TType> & { name: "ok" },
    b: Readonly<TType & { name: string }>,
    c: Readonly<TType> & { name: string })
    c = a -- Works
    b = a -- Should also work
end


//// [homomorphicMappedTypeIntersectionAssignability.lua]
-- ported from tests/cases/compiler/homomorphicMappedTypeIntersectionAssignability.ts
-- dropped: @target: es2015 and @strict: true directives (tlua defaults to latest target and strict checking)
function f(a, b, c)
  c = a; -- Works
  b = a; -- Should also work
end

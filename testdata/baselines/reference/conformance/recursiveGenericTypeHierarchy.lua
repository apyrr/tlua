//// [tests/cases/conformance/ported/recursiveGenericTypeHierarchy.tlua] ////

//// [recursiveGenericTypeHierarchy.tlua]
-- ported from tests/cases/compiler/recursiveGenericTypeHierarchy.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

-- used to ICE
interface A<T extends A<T, S>, S extends A<T, S>> {}

interface B<T extends B<T, S>, S extends B<T, S>> extends A<B<T, S>, B<T, S>> {}


//// [recursiveGenericTypeHierarchy.lua]
-- ported from tests/cases/compiler/recursiveGenericTypeHierarchy.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

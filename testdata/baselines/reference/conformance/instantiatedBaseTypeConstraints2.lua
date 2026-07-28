//// [tests/cases/conformance/ported/instantiatedBaseTypeConstraints2.tlua] ////

//// [instantiatedBaseTypeConstraints2.tlua]
-- ported from tests/cases/compiler/instantiatedBaseTypeConstraints2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A<T extends A<T, S>, S extends A<T, S>> { }
interface B<U> extends A<B<U>, B<U>> { }


//// [instantiatedBaseTypeConstraints2.lua]
-- ported from tests/cases/compiler/instantiatedBaseTypeConstraints2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/unusedTypeParameterInInterface2.tlua] ////

//// [unusedTypeParameterInInterface2.tlua]
-- ported from tests/cases/compiler/unusedTypeParameterInInterface2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: noUnusedLocals does not report the unused interface type parameter U


interface int<T, U, V> {
    f1(a: T): string
    c: V
}


//// [unusedTypeParameterInInterface2.lua]
-- ported from tests/cases/compiler/unusedTypeParameterInInterface2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: noUnusedLocals does not report the unused interface type parameter U

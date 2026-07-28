//// [tests/cases/conformance/ported/nominalSubtypeCheckOfTypeParameter.tlua] ////

//// [nominalSubtypeCheckOfTypeParameter.tlua]
-- ported from tests/cases/conformance/types/typeRelationships/recursiveTypes/nominalSubtypeCheckOfTypeParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface BinaryTuple<T, S> {
    first: T
    second: S
}

interface Sequence<T> {
    hasNext(): boolean
    pop(): T
    zip<S>(seq: Sequence<S>): Sequence<BinaryTuple<T, S>>
}

-- error: nominal subtype checks for infinitely expanding type references do not
-- allow nominal subtyping for the generic type itself.
interface List<T> extends Sequence<T> {
    getLength(): number
    zip<S>(seq: Sequence<S>): List<BinaryTuple<T, S>>
}


//// [nominalSubtypeCheckOfTypeParameter.lua]
-- ported from tests/cases/conformance/types/typeRelationships/recursiveTypes/nominalSubtypeCheckOfTypeParameter.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

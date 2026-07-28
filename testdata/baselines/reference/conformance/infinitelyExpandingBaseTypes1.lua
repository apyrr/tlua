//// [tests/cases/conformance/ported/infinitelyExpandingBaseTypes1.tlua] ////

//// [infinitelyExpandingBaseTypes1.tlua]
-- ported from tests/cases/compiler/infinitelyExpandingBaseTypes1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A<T> {
    x: A<A<T>>
}

interface B<T> {
    x: B<T>
}

interface C<T> extends A<T>, B<T> {}


//// [infinitelyExpandingBaseTypes1.lua]
-- ported from tests/cases/compiler/infinitelyExpandingBaseTypes1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

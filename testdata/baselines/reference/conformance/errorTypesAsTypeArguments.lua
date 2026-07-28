//// [tests/cases/conformance/ported/errorTypesAsTypeArguments.tlua] ////

//// [errorTypesAsTypeArguments.tlua]
-- ported from tests/cases/compiler/errorTypesAsTypeArguments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo<A> {
    bar(baz: Foo<B>): Foo<C>
}


//// [errorTypesAsTypeArguments.lua]
-- ported from tests/cases/compiler/errorTypesAsTypeArguments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

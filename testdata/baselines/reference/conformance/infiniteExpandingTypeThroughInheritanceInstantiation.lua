//// [tests/cases/conformance/ported/infiniteExpandingTypeThroughInheritanceInstantiation.tlua] ////

//// [infiniteExpandingTypeThroughInheritanceInstantiation.tlua]
-- ported from tests/cases/compiler/infiniteExpandingTypeThroughInheritanceInstantiation.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's "// error" comment on B<T> is stale; the current upstream reference has no errors
--   baseline. Both compilers accept these declarations and resolve x as A<B<T>> / B<A<T>> without
--   nontermination, which is what this test covers.

interface A<T> {
    x: A<B<T>>
}

interface B<T> extends A<T> {
    x: B<A<T>>
}


//// [infiniteExpandingTypeThroughInheritanceInstantiation.lua]
-- ported from tests/cases/compiler/infiniteExpandingTypeThroughInheritanceInstantiation.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: upstream's "// error" comment on B<T> is stale; the current upstream reference has no errors
--   baseline. Both compilers accept these declarations and resolve x as A<B<T>> / B<A<T>> without
--   nontermination, which is what this test covers.

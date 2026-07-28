//// [tests/cases/conformance/ported/infiniteExpandingTypeThroughInheritanceInstantiation.tlua] ////

//// [infiniteExpandingTypeThroughInheritanceInstantiation.tlua]
-- ported from tests/cases/compiler/infiniteExpandingTypeThroughInheritanceInstantiation.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua does not report the upstream circular inheritance-instantiation diagnostic on B<T>

interface A<T> {
    x: A<B<T>>
}

interface B<T> extends A<T> {
    x: B<A<T>>
}


//// [infiniteExpandingTypeThroughInheritanceInstantiation.lua]
-- ported from tests/cases/compiler/infiniteExpandingTypeThroughInheritanceInstantiation.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- compiler gap: tlua does not report the upstream circular inheritance-instantiation diagnostic on B<T>

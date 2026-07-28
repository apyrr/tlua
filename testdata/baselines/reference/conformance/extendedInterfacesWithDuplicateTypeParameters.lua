//// [tests/cases/conformance/ported/extendedInterfacesWithDuplicateTypeParameters.tlua] ////

//// [extendedInterfacesWithDuplicateTypeParameters.tlua]
-- ported from tests/cases/compiler/extendedInterfacesWithDuplicateTypeParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface InterfaceWithMultipleTypars<A, A> { -- should error
    bar(): void
}

interface InterfaceWithSomeTypars<B> { -- should not error
    bar(): void
}

interface InterfaceWithSomeTypars<C, C> { -- should error
    bar2(): void
}


//// [extendedInterfacesWithDuplicateTypeParameters.lua]
-- ported from tests/cases/compiler/extendedInterfacesWithDuplicateTypeParameters.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/multipleBaseInterfaesWithIncompatibleProperties.tlua] ////

//// [multipleBaseInterfaesWithIncompatibleProperties.tlua]
-- ported from tests/cases/compiler/multipleBaseInterfaesWithIncompatibleProperties.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A<T> {
    x: T
}

interface C extends A<string>, A<number> { }


//// [multipleBaseInterfaesWithIncompatibleProperties.lua]
-- ported from tests/cases/compiler/multipleBaseInterfaesWithIncompatibleProperties.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

//// [tests/cases/conformance/ported/genericAndNonGenericInheritedSignature2.tlua] ////

//// [genericAndNonGenericInheritedSignature2.tlua]
-- ported from tests/cases/compiler/genericAndNonGenericInheritedSignature2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo {
    f(x: any): any
}

interface Bar {
    f<T>(x: T): T
}

interface Hello extends Bar, Foo {
}


//// [genericAndNonGenericInheritedSignature2.lua]
-- ported from tests/cases/compiler/genericAndNonGenericInheritedSignature2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

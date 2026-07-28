//// [tests/cases/conformance/ported/genericAndNonGenericInheritedSignature1.tlua] ////

//// [genericAndNonGenericInheritedSignature1.tlua]
-- ported from tests/cases/compiler/genericAndNonGenericInheritedSignature1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo {
    f(x: any): any
}

interface Bar {
    f<T>(x: T): T
}

interface Hello extends Foo, Bar {
}


//// [genericAndNonGenericInheritedSignature1.lua]
-- ported from tests/cases/compiler/genericAndNonGenericInheritedSignature1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

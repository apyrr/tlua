//// [tests/cases/conformance/ported/interfaceThatInheritsFromItself.tlua] ////

//// [interfaceThatInheritsFromItself.tlua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatInheritsFromItself.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Foo extends Foo { -- error
}

interface Foo2<T> extends Foo2<T> { -- error
}

interface Foo3<T> extends Foo3<string> { -- error
}

interface Bar implements Bar { -- error
}


//// [interfaceThatInheritsFromItself.lua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatInheritsFromItself.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

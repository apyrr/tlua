//// [tests/cases/conformance/ported/interfaceThatInheritsFromItself.tlua] ////

//// [interfaceThatInheritsFromItself.tlua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatInheritsFromItself.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: upstream's `interface Bar implements Bar` case — `implements` is a TypeScript class-only
--   clause that tlua has no declaration form for, so the case would only assert that removed syntax
--   is rejected rather than exercise recursive base types

interface Foo extends Foo { -- error
}

interface Foo2<T> extends Foo2<T> { -- error
}

interface Foo3<T> extends Foo3<string> { -- error
}


//// [interfaceThatInheritsFromItself.lua]
-- ported from tests/cases/conformance/interfaces/interfaceDeclarations/interfaceThatInheritsFromItself.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: upstream's `interface Bar implements Bar` case — `implements` is a TypeScript class-only
--   clause that tlua has no declaration form for, so the case would only assert that removed syntax
--   is rejected rather than exercise recursive base types

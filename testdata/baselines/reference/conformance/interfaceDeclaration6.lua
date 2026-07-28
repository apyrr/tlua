//// [tests/cases/conformance/ported/interfaceDeclaration6.tlua] ////

//// [interfaceDeclaration6.tlua]
-- ported from tests/cases/compiler/interfaceDeclaration6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface i1 { foo: number }
interface i2 extends i1 { foo: number }
interface i3 extends i1 { foo: string }
interface i4 {
    bar(): any
    bar(): any
}


//// [interfaceDeclaration6.lua]
-- ported from tests/cases/compiler/interfaceDeclaration6.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

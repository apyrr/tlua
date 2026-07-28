//// [tests/cases/conformance/ported/parserInterfaceDeclaration1.tlua] ////

//// [parserInterfaceDeclaration1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/InterfaceDeclarations/parserInterfaceDeclaration1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface I extends A extends B {
}


//// [parserInterfaceDeclaration1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/InterfaceDeclarations/parserInterfaceDeclaration1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

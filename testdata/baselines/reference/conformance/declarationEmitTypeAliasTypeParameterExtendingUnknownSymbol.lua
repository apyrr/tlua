//// [tests/cases/conformance/ported/declarationEmitTypeAliasTypeParameterExtendingUnknownSymbol.tlua] ////

//// [declarationEmitTypeAliasTypeParameterExtendingUnknownSymbol.tlua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasTypeParameterExtendingUnknownSymbol.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output (unsupported for Lua modules)

type A<T extends Unknown> = {}


//// [declarationEmitTypeAliasTypeParameterExtendingUnknownSymbol.lua]
-- ported from tests/cases/compiler/declarationEmitTypeAliasTypeParameterExtendingUnknownSymbol.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: declaration output (unsupported for Lua modules)

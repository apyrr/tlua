//// [tests/cases/compiler/tluaModuleDeclarationEmitUnsupported.tlua] ////

//// [tluaModuleDeclarationEmitUnsupported.tlua]
// Declaration emit for a Lua module has no ambient form yet; it must error
// with 100054 and emit an empty declaration file rather than panic.
local x = 1;
return { x = x };


//// [tluaModuleDeclarationEmitUnsupported.lua]
-- Declaration emit for a Lua module has no ambient form yet; it must error
-- with 100054 and emit an empty declaration file rather than panic.
local x = 1;
return { x = x };

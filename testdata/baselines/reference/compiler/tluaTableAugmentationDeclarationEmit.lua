//// [tests/cases/compiler/tluaTableAugmentationDeclarationEmit.tlua] ////

//// [tluaTableAugmentationDeclarationEmit.tlua]
local M = {};
M.value = 1;
M.label = "label";

Implicit = {};
Implicit.hidden = true;



//// [tluaTableAugmentationDeclarationEmit.lua]
local M = {};
M.value = 1;
M.label = "label";
Implicit = {};
Implicit.hidden = true;

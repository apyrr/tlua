//// [tests/cases/compiler/tluaTableAugmentationDeclareGlobal.tlua] ////

//// [decl.tlua]
declare global {
  interface Shape {
    x: number;
  }
}

//// [use.tlua]
Shape = {};
Shape.tag = 1;
local s: Shape = { x = 1 };


//// [decl.lua]
//// [use.lua]
Shape = {};
Shape.tag = 1;
local s = { x = 1 };

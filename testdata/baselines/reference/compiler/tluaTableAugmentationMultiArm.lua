//// [tests/cases/compiler/tluaTableAugmentationMultiArm.tlua] ////

//// [a.tlua]
V = {};
V.m = "s";
local va: string = V.m;

//// [b.tlua]
V = { b = 2 };
local vm: string = V.m;
local vb: number = V.b;

//// [c.tlua]
// Hover-visible union: records V's union type in the .types baseline.
local v = V;
local cm: string = V.m;

// `b` exists in only one arm's constructor. Read here on the un-narrowed
// union to record whether a direct read is an error or union-with-nil.
local cb = V.b;
V.b = 3;


//// [a.lua]
V = {};
V.m = "s";
local va = V.m;
//// [b.lua]
V = { b = 2 };
local vm = V.m;
local vb = V.b;
//// [c.lua]
-- Hover-visible union: records V's union type in the .types baseline.
local v = V;
local cm = V.m;
-- `b` exists in only one arm's constructor. Read here on the un-narrowed
-- union to record whether a direct read is an error or union-with-nil.
local cb = V.b;
V.b = 3;

//// [tests/cases/compiler/tluaTableAugmentationDefaultedGuard.tlua] ////

//// [a.tlua]
G = G or {};
G.fromA = 1;
local fromB: string = G.fromB;

local M = {};
M.sub = M.sub or {};
M.sub.value = true;
local value: boolean = M.sub.value;

local A = {};
A.k = A["k"] or {};
A.k.value = 1;
local mixedKey: number = A["k"].value;

Mixed = Mixed or {};
Mixed.extra = 1;

// An implicit global as the base of a nested member guard: the base read
// `GNested` inside `GNested.sub` must resolve like any other global, not
// trip used-before-assigned.
GNested = {};
GNested.sub = GNested.sub or {};
GNested.sub.deep = 1;
local deep: number = GNested.sub.deep;

//// [b.tlua]
G = G ?? {};
G.fromB = "b";
local fromA: number = G.fromA;

local function make(): { fixed: number }
  return { fixed = 1 };
end

Closed = Closed or make();
Closed.extra = true;

Mixed = make();

// Parentheses make this an ordinary self-referential assignment rather than
// the exact defaulted-guard idiom.
Paren = (Paren) or {};

local Numeric = {};
Numeric[1] = Numeric["1"] or {};

local Other = {};
Chain = Chain or Other or {};

Different = Other or {};
Different.extra = true;

local function keep(value: any): {}
  return {};
end
ReadInside = ReadInside or keep(ReadInside);

local LocalSelf = LocalSelf or {};

local Declared: {};
Declared = Declared or {};

local Outer = { base = 1 };
do
  local Outer = Outer or {};
  Outer.extra = true;
end


//// [a.lua]
G = G or {};
G.fromA = 1;
local fromB = G.fromB;
local M = {};
M.sub = M.sub or {};
M.sub.value = true;
local value = M.sub.value;
local A = {};
A.k = A["k"] or {};
A.k.value = 1;
local mixedKey = A["k"].value;
Mixed = Mixed or {};
Mixed.extra = 1;
-- An implicit global as the base of a nested member guard: the base read
-- `GNested` inside `GNested.sub` must resolve like any other global, not
-- trip used-before-assigned.
GNested = {};
GNested.sub = GNested.sub or {};
GNested.sub.deep = 1;
local deep = GNested.sub.deep;
//// [b.lua]
G = G ?? {};
G.fromB = "b";
local fromA = G.fromA;
local function make()
    return { fixed = 1 };
end
Closed = Closed or make();
Closed.extra = true;
Mixed = make();
-- Parentheses make this an ordinary self-referential assignment rather than
-- the exact defaulted-guard idiom.
Paren = (Paren) or {};
local Numeric = {};
Numeric[1] = Numeric["1"] or {};
local Other = {};
Chain = Chain or Other or {};
Different = Other or {};
Different.extra = true;
local function keep(value)
    return {};
end
ReadInside = ReadInside or keep(ReadInside);
local LocalSelf = LocalSelf or {};
local Declared;
Declared = Declared or {};
local Outer = { base = 1 };
do
    local Outer = Outer or {};
    Outer.extra = true;
end

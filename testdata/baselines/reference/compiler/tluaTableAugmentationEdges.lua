//// [tests/cases/compiler/tluaTableAugmentationEdges.tlua] ////

//// [a.tlua]
interface Shape {
  x: number;
}
Shape = {};
Shape.tag = 1;
local s: Shape = { x = 1 };

// A non-constructor arm in any file closes an implicit global program-wide.
//// [b.tlua]
G = {};
G.x = 1;

U = {};

//// [c.tlua]
local function make(): { y: number }
  return { y = 2 };
end
G = make();

// A member written in only one arm's constructor needs narrowing: before the
// assignment below narrows U, the read sees the whole union.
local ua: number = U.a;

U = { a = 1 };
local narrowed: number = U.a;

// A dotted function colliding with a constructor field is a duplicate no
// matter where it sits in the group; assignments to that field stay ordinary
// checked writes.
//// [d.tlua]
local t = { member = 1 };
function t.member(): void
end
t.member = 2;
t.member = "no";


//// [a.lua]
Shape = {};
Shape.tag = 1;
local s = { x = 1 };
-- A non-constructor arm in any file closes an implicit global program-wide.
//// [b.lua]
G = {};
G.x = 1;
U = {};
//// [c.lua]
local function make()
    return { y = 2 };
end
G = make();
-- A member written in only one arm's constructor needs narrowing: before the
-- assignment below narrows U, the read sees the whole union.
local ua = U.a;
U = { a = 1 };
local narrowed = U.a;
-- A dotted function colliding with a constructor field is a duplicate no
-- matter where it sits in the group; assignments to that field stay ordinary
-- checked writes.
//// [d.lua]
local t = { member = 1 };
function t.member()
end
t.member = 2;
t.member = "no";

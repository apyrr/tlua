//// [tests/cases/compiler/tluaLocalTypeInconsistentMerge.tlua] ////

//// [iface.tlua]
// A `local interface` merging with a bare `interface` of the same name is an
// inconsistent-modifier error: the `local` prefix decides whether the whole
// merged symbol is module-private, so a mismatched prefix is ambiguous. The error
// points at the inconsistent `local` declaration, not whichever came first.
interface Mixed { a: number };
local interface Mixed { b: number };

//// [iface-local-first.tlua]
// The same, with the `local` declaration first: the error now anchors on the bare
// declaration (the one that differs from the local primary), regardless of order.
local interface Flipped { a: number };
interface Flipped { b: number };

//// [shadow-a.tlua]
// A module-private `local type` and a bare global `type` of the same name live
// in different scopes across modules and do NOT collide: each module sees its own.
local type Name = number;
local va: Name = 1;

//// [shadow-b.tlua]
type Name = string;
local vb: Name = "x";


//// [iface.lua]
;
;
//// [iface-local-first.lua]
;
;
//// [shadow-a.lua]
local va = 1;
//// [shadow-b.lua]
local vb = "x";

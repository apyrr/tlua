//// [tests/cases/compiler/tluaWordOperatorPropertyNames.tlua] ////

//// [tluaWordOperatorPropertyNames.tlua]
// A property may be *named* `and`, `or` or `not` -- only the bare member syntax
// is unavailable, exactly as in Lua, where `t["and"]` works and `t.and` does not.
// So the type printer must quote these names: an emitted `and: number` would not
// parse back in. Ordinary keywords like `while` stay bare when *printed* as a
// member, even though a reserved word still needs the bracket key form in a
// table literal.

local t = { ["and"] = 1, ["or"] = 2, ["not"] = 3, ["while"] = 4, plain = 5 };

// Reading them back requires the bracket form.
local readAnd = t["and"];
local readOr = t["or"];
local readNot = t["not"];
local readWhile = t.while;

// Interfaces declare them the same way.
interface Bits {
  "and"(v: number): number;
  "or"(v: number): number;
  while: number;
}
declare bits: Bits;
local banded = bits["and"](1);

// An expando member named `and` is skipped by declaration emit rather than
// emitted as the bare alias `export { ... as and }`, which would not parse.
function A(): string { return "A"; }
A["and"] = t;
A["ok"] = t;


//// [tluaWordOperatorPropertyNames.lua]
-- A property may be *named* `and`, `or` or `not` -- only the bare member syntax
-- is unavailable, exactly as in Lua, where `t["and"]` works and `t.and` does not.
-- So the type printer must quote these names: an emitted `and: number` would not
-- parse back in. Ordinary keywords like `while` stay bare when *printed* as a
-- member, even though a reserved word still needs the bracket key form in a
-- table literal.
local t = { ["and"] = 1, ["or"] = 2, ["not"] = 3, ["while"] = 4, plain = 5 };
-- Reading them back requires the bracket form.
local readAnd = t["and"];
local readOr = t["or"];
local readNot = t["not"];
local readWhile = t.while;
local banded = bits["and"](1);
-- An expando member named `and` is skipped by declaration emit rather than
-- emitted as the bare alias `export { ... as and }`, which would not parse.
function A() { return "A"; }
A["and"] = t;
A["ok"] = t;

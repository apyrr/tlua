//// [tests/cases/compiler/tluaNilRename.tlua] ////

//// [tluaNilRename.tlua]
// `nil` works as a type and a value.
local a: nil = nil;
local b = nil;

// A type query can name it too.
type Nil = typeof nil;
local c: Nil = nil;

// Optional properties surface as `nil`.
interface Box {
	value?: number;
}
local g: Box = {};
local h = g.value;


//// [tluaNilRename.lua]
-- `nil` works as a type and a value.
local a = nil;
local b = nil;
local c = nil;
local g = {};
local h = g.value;

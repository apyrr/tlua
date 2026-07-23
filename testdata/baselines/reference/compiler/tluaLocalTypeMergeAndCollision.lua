//// [tests/cases/compiler/tluaLocalTypeMergeAndCollision.tlua] ////

//// [merge.tlua]
// Two `local interface` with the same name in one module merge, like ordinary
// interface declaration merging -- both stay module-private.
local interface Box { width: number };
local interface Box { height: number };

local b: Box = { width = 1, height = 2 };

//// [collision.tlua]
// A `local type` colliding with a bare `type` of the same name in one module is
// a duplicate identifier: both land in the module's local scope.
local type Dup = number;
type Dup = string;

//// [nested.tlua]
// `local type` nested in a block is redundant with bare nested types (both scope
// lexically) but is legal for spelling consistency with `local function`.
do
  local type Inner = { v: number };
  local x: Inner = { v = 1 };
end


//// [merge.lua]
;
;
local b = { width = 1, height = 2 };
//// [collision.lua]
//// [nested.lua]
-- `local type` nested in a block is redundant with bare nested types (both scope
-- lexically) but is legal for spelling consistency with `local function`.
do
  local x = { v = 1 };
end

//// [tests/cases/compiler/tluaColonFunctionErrors.tlua] ////

//// [tluaColonFunctionErrors.tlua]
local M = {};

function M:bar(v: number): number
  return v;
end

// The colon call form supplies the receiver implicitly.
M:bar(1);

// `self` is already declared by the colon.
function M:clash(self: number): void
end

// At most one colon segment, and nothing may follow it.
function M:deep.f(): void
end

function M:a:b(): void
end

// `local function` takes a plain name.
local function Bad:f(): void
end

// Unknown base.
function Unknown:f(): void
end


//// [tluaColonFunctionErrors.lua]
local M = {};
function M:bar(v)
  return v;
end
-- The colon call form supplies the receiver implicitly.
M:bar(1);
-- `self` is already declared by the colon.
function M:clash(self)
end
-- At most one colon segment, and nothing may follow it.
function M:deep()
  f();
  void ;
end
function M:a()
  ();
  void ;
end
-- `local function` takes a plain name.
local function Bad()
  ();
  void ;
end
-- Unknown base.
function Unknown:f()
end

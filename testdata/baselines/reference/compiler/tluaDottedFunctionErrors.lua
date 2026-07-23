//// [tests/cases/compiler/tluaDottedFunctionErrors.tlua] ////

//// [tluaDottedFunctionErrors.tlua]
// Unknown base.
function Unknown.f(): void
end

// A function is not a table and cannot carry members.
local function g(): void
end

function g.h(): void
end

// Duplicate members.
local Dup = {};

function Dup.f(): void
end

function Dup.f(): void
end

// `local function` takes a plain name.
local function Bad.f(): void
end


//// [tluaDottedFunctionErrors.lua]
-- Unknown base.
function Unknown.f()
end
-- A function is not a table and cannot carry members.
local function g()
end
function g.h()
end
-- Duplicate members.
local Dup = {};
function Dup.f()
end
function Dup.f()
end
-- `local function` takes a plain name.
local function Bad()
  f();
  void ;
end

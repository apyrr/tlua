//// [tests/cases/compiler/tluaDottedFunctionDeclarationEmit.tlua] ////

//// [tluaDottedFunctionDeclarationEmit.tlua]
local Exported = {};

function Exported.visible(a: number): number
  return a;
end

local Hidden = {};

function Hidden.invisible(): void
end


//// [tluaDottedFunctionDeclarationEmit.lua]
local Exported = {};
function Exported.visible(a)
  return a;
end
local Hidden = {};
function Hidden.invisible()
end

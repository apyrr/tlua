//// [tests/cases/compiler/tluaColonFunctionDeclarationEmit.tlua] ////

//// [tluaColonFunctionDeclarationEmit.tlua]
local Exported = {};

function Exported:visible(a: number): number
  return a;
end

local Hidden = {};

function Hidden:invisible(): void
  self.invisible(self);
end


//// [tluaColonFunctionDeclarationEmit.lua]
local Exported = {};
function Exported:visible(a)
    return a;
end
local Hidden = {};
function Hidden:invisible()
    self.invisible(self);
end

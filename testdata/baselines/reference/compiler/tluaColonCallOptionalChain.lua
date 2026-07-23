//// [tests/cases/compiler/tluaColonCallOptionalChain.tlua] ////

//// [tluaColonCallOptionalChain.tlua]
local W = { n = 1 };

function W:bump(): number
  return self.n;
end

local maybe: { inner: typeof W } | nil = { inner = W };

// The `?.` guard narrows the receiver, so the implicit first argument is not
// nil — and the downleveled chain keeps the colon on the call segment.
local y = maybe?.inner:bump();


//// [tluaColonCallOptionalChain.lua]
local W = { n = 1 };
function W:bump()
  return self.n;
end
local maybe = { inner = W };
-- The `?.` guard narrows the receiver, so the implicit first argument is not
-- nil — and the downleveled chain keeps the colon on the call segment.
local y = maybe and maybe.inner:bump();

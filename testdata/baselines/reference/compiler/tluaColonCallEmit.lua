//// [tests/cases/compiler/tluaColonCallEmit.tlua] ////

//// [tluaColonCallEmit.tlua]
local T = { n = 0 };

function T:bump(): number
  self.n = self.n + 1;
  return self.n;
end

// Colon calls round-trip through the emitter.
T:bump();
local a = T:bump() + T:bump();

// A parenthesized receiver keeps its parentheses.
local b = (T):bump();

// Parentheses around the whole call truncate it to one value.
local c = (T:bump());

// Chains keep every colon.
local d = { u = T };
d.u:bump();

// Return- and throw-position colon calls keep their colon: the no-ASI
// parenthesizer must not rebuild the callee as a dot access.
function T:twice(): number
  return self:bump() + self:bump();
end

function raise(): never
  throw T:bump();
end


//// [tluaColonCallEmit.lua]
local T = { n = 0 };
function T:bump()
  self.n = self.n + 1;
  return self.n;
end
-- Colon calls round-trip through the emitter.
T:bump();
local a = T:bump() + T:bump();
-- A parenthesized receiver keeps its parentheses.
local b = (T):bump();
-- Parentheses around the whole call truncate it to one value.
local c = (T:bump());
-- Chains keep every colon.
local d = { u = T };
d.u:bump();
-- Return- and throw-position colon calls keep their colon: the no-ASI
-- parenthesizer must not rebuild the callee as a dot access.
function T:twice()
  return self:bump() + self:bump();
end
function raise()
  throw T:bump();
end

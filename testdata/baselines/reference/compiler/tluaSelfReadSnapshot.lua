//// [tests/cases/compiler/tluaSelfReadSnapshot.tlua] ////

//// [tluaSelfReadSnapshot.tlua]
// A self-read nested inside its own store's value list reads the pre-store
// snapshot instead of reporting a false circularity: nil before the first
// store, the union the other stores install otherwise. The exact defaulted
// guard and whole-value capture keep their own paths; this pins the nested
// shapes that used to report TLUA7022.

count = (count or 0) + 1;
local fromGuarded: number = count;

total = 0;
total = total + 1;
local fromSequenced: number = total;

t = {};
t.n = (t.n or 0) + 1;
local fromMember: number = t.n;

_G.hits = (_G.hits or 0) + 1;
local fromEnvironment: number = hits;

p, q = (p or 0) + 1, (q or "") .. "!";
local fromList: number = p;
local fromListSecond: string = q;

// A literal element key names the same property either spelling reads, so the
// element and mixed spellings type like the dot access above.
e = {};
e[1] = (e[1] or 0) + 1;
local fromElement: number = e[1];

s = {};
s["k"] = (s["k"] or 0) + 1;
local fromStringKey: number = s["k"];

m = {};
m.k = (m["k"] or 0) + 1;
local fromMixed: number = m.k;

// An unguarded self-read is diagnosed against the snapshot: the only store
// sees nil and never handles it.
bare = bare + 1;

// Stores may sit in any container; the read still sees its own store's
// snapshot, here the union of the other stores.
steps = 0;
local function advance()
  steps = steps + 1;
end

// The snapshot is statement-ordered: a later store has not run when the first
// statement reads, so the self-read still sees nil.
early = early + 1;
early = 0;

// A store nested in a conditional may have been skipped, so nil stays in the
// union even though the store precedes the read.
if math.random() > 0.5 then
  maybe = 0;
end
maybe = maybe + 1;

// A store inside a function declared later cannot have run before this
// statement executes.
late = late + 1;
local function setLate()
  late = 0;
end

// Within one invocation a body runs its statements in order: the read in the
// first statement happens before the later store.
local function firstCall()
  seen = seen + 1;
  seen = 0;
end

// A store earlier in the same body has definitely run by the read.
local function orderedBody()
  ready = 0;
  ready = ready + 1;
end

// A self-preserving store re-stores whatever was there, nil included, so it
// cannot discharge nil from the snapshot.
if math.random() > 0.5 then
  held = 0;
end
held = held;
held = held + 1;

// An immediately invoked value runs before the store completes, so its read
// is part of the captured snapshot rather than a circularity.
iife = (function() return (iife or 0) + 1 end)();
local fromIife: number = iife;

// Flow facts at the reference narrow the snapshot: the guard removes nil from
// the self-read. (The guard's own read keeps the ordinary path's
// used-before-assigned strictness; only the self-read types via the snapshot.)
if math.random() > 0.5 then
  guarded = 0;
end
if guarded ~= nil then
  guarded = guarded + 1;
end

// A sibling store in the same block has definitely run by the time the next
// statement reads: entering the block passed it.
if math.random() > 0.5 then
  sibling = 0;
  sibling = sibling + 1;
end


//// [tluaSelfReadSnapshot.lua]
-- A self-read nested inside its own store's value list reads the pre-store
-- snapshot instead of reporting a false circularity: nil before the first
-- store, the union the other stores install otherwise. The exact defaulted
-- guard and whole-value capture keep their own paths; this pins the nested
-- shapes that used to report TLUA7022.
count = (count or 0) + 1;
local fromGuarded = count;
total = 0;
total = total + 1;
local fromSequenced = total;
t = {};
t.n = (t.n or 0) + 1;
local fromMember = t.n;
_G.hits = (_G.hits or 0) + 1;
local fromEnvironment = hits;
p, q = (p or 0) + 1, (q or "") .. "!";
local fromList = p;
local fromListSecond = q;
-- A literal element key names the same property either spelling reads, so the
-- element and mixed spellings type like the dot access above.
e = {};
e[1] = (e[1] or 0) + 1;
local fromElement = e[1];
s = {};
s["k"] = (s["k"] or 0) + 1;
local fromStringKey = s["k"];
m = {};
m.k = (m["k"] or 0) + 1;
local fromMixed = m.k;
-- An unguarded self-read is diagnosed against the snapshot: the only store
-- sees nil and never handles it.
bare = bare + 1;
-- Stores may sit in any container; the read still sees its own store's
-- snapshot, here the union of the other stores.
steps = 0;
local function advance()
  steps = steps + 1;
end
-- The snapshot is statement-ordered: a later store has not run when the first
-- statement reads, so the self-read still sees nil.
early = early + 1;
early = 0;
-- A store nested in a conditional may have been skipped, so nil stays in the
-- union even though the store precedes the read.
if math.random() > 0.5 then
  maybe = 0;
end
maybe = maybe + 1;
-- A store inside a function declared later cannot have run before this
-- statement executes.
late = late + 1;
local function setLate()
  late = 0;
end
-- Within one invocation a body runs its statements in order: the read in the
-- first statement happens before the later store.
local function firstCall()
  seen = seen + 1;
  seen = 0;
end
-- A store earlier in the same body has definitely run by the read.
local function orderedBody()
  ready = 0;
  ready = ready + 1;
end
-- A self-preserving store re-stores whatever was there, nil included, so it
-- cannot discharge nil from the snapshot.
if math.random() > 0.5 then
  held = 0;
end
held = held;
held = held + 1;
-- An immediately invoked value runs before the store completes, so its read
-- is part of the captured snapshot rather than a circularity.
iife = (function()
  return (iife or 0) + 1;
end)();
local fromIife = iife;
-- Flow facts at the reference narrow the snapshot: the guard removes nil from
-- the self-read. (The guard's own read keeps the ordinary path's
-- used-before-assigned strictness; only the self-read types via the snapshot.)
if math.random() > 0.5 then
  guarded = 0;
end
if guarded ~= nil then
  guarded = guarded + 1;
end
-- A sibling store in the same block has definitely run by the time the next
-- statement reads: entering the block passed it.
if math.random() > 0.5 then
  sibling = 0;
  sibling = sibling + 1;
end

//// [tests/cases/compiler/tluaSuspendAmbientDeclaration.tlua] ////

//// [tluaSuspendAmbientDeclaration.tlua]
// The ambient form of the suspend contract: `declare suspend function` is
// what declaration files emit, so the checker must both accept it and keep
// enforcing caller-context coloring through it.

declare suspend function fetchOk(url: string): (boolean, any);

// Error: direct call from a sync function.
function syncCaller(): void
  fetchOk("nope");
end

// OK: call from a suspend context.
suspend function suspendCaller(): number
  local ok, res = fetchOk("https://example.com");
  if ok then
    return 1;
  end
  return 0;
end

// OK: passing the ambient suspend function as a value is not a call.
coroutine.create(fetchOk);

// Error: a concrete sync signature still rejects the ambient suspend function.
local concrete: (url: string) => (boolean, any) = fetchOk;

// Error: `declare` must precede `suspend`.
suspend declare function backwards(): void;


//// [tluaSuspendAmbientDeclaration.lua]
-- The ambient form of the suspend contract: `declare suspend function` is
-- what declaration files emit, so the checker must both accept it and keep
-- enforcing caller-context coloring through it.
-- Error: direct call from a sync function.
function syncCaller()
  fetchOk("nope");
end
-- OK: call from a suspend context.
function suspendCaller()
  local ok, res = fetchOk("https://example.com");
  if ok then
    return 1;
  end
  return 0;
end
-- OK: passing the ambient suspend function as a value is not a call.
coroutine.create(fetchOk);
-- Error: a concrete sync signature still rejects the ambient suspend function.
local concrete = fetchOk;

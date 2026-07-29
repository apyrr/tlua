//// [tests/cases/compiler/tluaSuspendContract.tlua] ////

//// [tluaSuspendContract.tlua]
// The tlua coroutine contract: a suspend function may only be called from a
// suspend context. Suspend functions return their plain type (or pack), never
// Promise<T>.

declare function request(url: string): (boolean, any);

suspend function fetchOk(url: string): (boolean, any)
  local ok, res = request(url);
  return ok, res;
end

// OK: direct call from a suspend context; the result is the plain pack.
suspend function handler(): number
  local ok, res = fetchOk("https://example.com");
  if ok then
    return 1;
  end
  return 0;
end

// Error: direct call from a sync function.
function syncCaller(): void
  fetchOk("nope");
end

// Error: call at top level.
fetchOk("top");

// Error: inferred alias keeps the contract (the flag rides the signature).
function aliasCaller(): void
  local f = fetchOk;
  f("aliased");
end

// OK: alias called from a suspend context.
suspend function aliasSuspendCaller(): void
  local f = fetchOk;
  f("aliased");
end

// OK: passing a suspend function as a value is not a call. `coroutine.create`
// and `coroutine.wrap` take the `function` top type, and a suspend function is
// a function -- starting one as a coroutine is the whole point of the modifier.
coroutine.create(fetchOk);
local wrapped = coroutine.wrap(fetchOk);

// OK: `function` is the top of every Lua function, so it must behave exactly
// like the structurally identical `(...: any) => any` in target position.
local topKeyword: function = fetchOk;
local topWritten: (...: any) => any = fetchOk;

// Error: a top target is the only escape hatch -- a concrete sync signature
// still rejects, including through a contravariant callback flip.
local concrete: (url: string) => (boolean, any) = fetchOk;
declare function takesSync(cb: (url: string) => (boolean, any)): void;
takesSync(fetchOk);

// OK: suspend recursion.
suspend function pingpong(n: number): number
  if n <= 0 then
    return 0;
  end
  return pingpong(n - 1);
end

// Suspend arrow expressions participate too.
local suspendArrow = suspend function(): number return 42 end;
function arrowSyncCaller(): void
  suspendArrow(); // error
end
suspend function arrowSuspendCaller(): number
  return suspendArrow(); // OK
end

// Generic suspend functions keep the flag through instantiation.
suspend function identitySuspend<T>(x: T): T
  return x;
end
function genericSyncCaller(): void
  identitySuspend(1); // error
end
suspend function genericSuspendCaller(): number
  return identitySuspend(2); // OK
end

// The suspend return type is the plain type, not Promise<T>.
suspend function plain(): number
  return 3;
end
suspend function usePlain(): number
  local n = plain();
  return n + 1;
end


//// [tluaSuspendContract.lua]
-- The tlua coroutine contract: a suspend function may only be called from a
-- suspend context. Suspend functions return their plain type (or pack), never
-- Promise<T>.
function fetchOk(url)
  local ok, res = request(url);
  return ok, res;
end
-- OK: direct call from a suspend context; the result is the plain pack.
function handler()
  local ok, res = fetchOk("https://example.com");
  if ok then
    return 1;
  end
  return 0;
end
-- Error: direct call from a sync function.
function syncCaller()
  fetchOk("nope");
end
-- Error: call at top level.
fetchOk("top");
-- Error: inferred alias keeps the contract (the flag rides the signature).
function aliasCaller()
  local f = fetchOk;
  f("aliased");
end
-- OK: alias called from a suspend context.
function aliasSuspendCaller()
  local f = fetchOk;
  f("aliased");
end
-- OK: passing a suspend function as a value is not a call. `coroutine.create`
-- and `coroutine.wrap` take the `function` top type, and a suspend function is
-- a function -- starting one as a coroutine is the whole point of the modifier.
coroutine.create(fetchOk);
local wrapped = coroutine.wrap(fetchOk);
-- OK: `function` is the top of every Lua function, so it must behave exactly
-- like the structurally identical `(...: any) => any` in target position.
local topKeyword = fetchOk;
local topWritten = fetchOk;
-- Error: a top target is the only escape hatch -- a concrete sync signature
-- still rejects, including through a contravariant callback flip.
local concrete = fetchOk;
takesSync(fetchOk);
-- OK: suspend recursion.
function pingpong(n)
  if n <= 0 then
    return 0;
  end
  return pingpong(n - 1);
end
-- Suspend arrow expressions participate too.
local suspendArrow = function()
  return 42;
end;
function arrowSyncCaller()
  suspendArrow(); -- error
end
function arrowSuspendCaller()
  return suspendArrow(); -- OK
end
-- Generic suspend functions keep the flag through instantiation.
function identitySuspend(x)
  return x;
end
function genericSyncCaller()
  identitySuspend(1); -- error
end
function genericSuspendCaller()
  return identitySuspend(2); -- OK
end
-- The suspend return type is the plain type, not Promise<T>.
function plain()
  return 3;
end
function usePlain()
  local n = plain();
  return n + 1;
end

//// [tests/cases/compiler/tluaAsyncFunctionType.tlua] ////

//// [tluaAsyncFunctionType.tlua]
// An `async` modifier is allowed on function *type* nodes (the parenthesized
// `async (params) => T` form). It carries the coroutine contract: a call
// through an async-typed slot is checked, and async-ness is visible in the
// type display.

declare function request(url: string): boolean;

async function fetchOk(url: string): boolean
  return request(url);
end

declare function plainOk(url: string): boolean;

// OK: async fn into an async-typed slot; the slot keeps the contract.
local g: async (url: string) => boolean = fetchOk;

// OK: sync fn into an async-typed slot (a sync function may run in a coroutine).
local h: async (url: string) => boolean = plainOk;

// Calling through the async-typed slot: error from a sync function, OK from an
// async function.
function sync1(): void
  g("a");
end
async function ok(): void
  g("b");
end

// Passing an async fn as a value into the "any function" top signature is fine
// (it is how coroutine.create/wrap receive a suspendable function).
declare namespace coroutine {
  function create(f: (...: any) => any): any;
}
coroutine.create(fetchOk);

// async function types are usable in ambient / declared positions.
declare handler: async (url: string) => boolean;

// `async` used as an ordinary type name is unaffected (no parenthesized head).
type async = number;
local n: async = 1;


//// [tluaAsyncFunctionType.lua]
-- An `async` modifier is allowed on function *type* nodes (the parenthesized
-- `async (params) => T` form). It carries the coroutine contract: a call
-- through an async-typed slot is checked, and async-ness is visible in the
-- type display.
async function fetchOk(url)
  return request(url);
end
-- OK: async fn into an async-typed slot; the slot keeps the contract.
local g = fetchOk;
-- OK: sync fn into an async-typed slot (a sync function may run in a coroutine).
local h = plainOk;
-- Calling through the async-typed slot: error from a sync function, OK from an
-- async function.
function sync1()
  g("a");
end
async function ok()
  g("b");
end
-- Passing an async fn as a value into the "any function" top signature is fine
-- (it is how coroutine.create/wrap receive a suspendable function).
declare;
namespace;
coroutine;
do
end
coroutine.create(fetchOk);
local n = 1;

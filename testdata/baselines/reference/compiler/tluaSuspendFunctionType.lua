//// [tests/cases/compiler/tluaSuspendFunctionType.tlua] ////

//// [tluaSuspendFunctionType.tlua]
// A `suspend` modifier is allowed on function *type* nodes (the parenthesized
// `suspend (params) => T` form). It carries the coroutine contract: a call
// through a suspend-typed slot is checked, and suspend-ness is visible in the
// type display.

declare function request(url: string): boolean;

suspend function fetchOk(url: string): boolean
  return request(url);
end

declare function plainOk(url: string): boolean;

// OK: suspend fn into a suspend-typed slot; the slot keeps the contract.
local g: suspend (url: string) => boolean = fetchOk;

// OK: sync fn into a suspend-typed slot (a sync function may run in a coroutine).
local h: suspend (url: string) => boolean = plainOk;

// Calling through the suspend-typed slot: error from a sync function, OK from a
// suspend function.
function sync1(): void
  g("a");
end
suspend function ok(): void
  g("b");
end

// Passing a suspend fn as a value into the "any function" top signature is fine
// (it is how coroutine.create/wrap receive a suspendable function). The global
// `coroutine` lib declares these as taking the `function` keyword type, which
// must behave the same as the written `(...: any) => any` spelling.
coroutine.create(fetchOk);
local anyFn: (...: any) => any = fetchOk;

// suspend function types are usable in ambient / declared positions.
declare handler: suspend (url: string) => boolean;

// `suspend` used as an ordinary type name is unaffected (no parenthesized head).
type suspend = number;
local n: suspend = 1;


//// [tluaSuspendFunctionType.lua]
-- A `suspend` modifier is allowed on function *type* nodes (the parenthesized
-- `suspend (params) => T` form). It carries the coroutine contract: a call
-- through a suspend-typed slot is checked, and suspend-ness is visible in the
-- type display.
suspend function fetchOk(url)
  return request(url);
end
-- OK: suspend fn into a suspend-typed slot; the slot keeps the contract.
local g = fetchOk;
-- OK: sync fn into a suspend-typed slot (a sync function may run in a coroutine).
local h = plainOk;
-- Calling through the suspend-typed slot: error from a sync function, OK from a
-- suspend function.
function sync1()
  g("a");
end
suspend function ok()
  g("b");
end
-- Passing a suspend fn as a value into the "any function" top signature is fine
-- (it is how coroutine.create/wrap receive a suspendable function). The global
-- `coroutine` lib declares these as taking the `function` keyword type, which
-- must behave the same as the written `(...: any) => any` spelling.
coroutine.create(fetchOk);
local anyFn = fetchOk;
local n = 1;

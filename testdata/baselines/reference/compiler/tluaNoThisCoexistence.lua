//// [tests/cases/compiler/tluaNoThisCoexistence.tlua] ////

//// [tluaNoThisCoexistence.tlua]
// `this` is removed, but the constructs that neighbour it must keep working.
// This file must stay zero-error.

declare function use(x: unknown): void;

// Ordinary identifier type predicates still narrow. Only the `this is T` form
// was removed, and it is a distinct branch of the predicate machinery.
function isString(v: unknown): v is string
  return type(v) == "string";
end

declare function assertString(v: unknown): asserts v is string;

function narrow(v: unknown): void
  if isString(v) then
    local s: string = v;
    use(s);
  end
  assertString(v);
  local t: string = v;
  use(t);
end

// Lua's global environment remains available independently of `this`.
local g: typeof _G = _G;

// Colon methods supply the receiver as an ordinary `self` parameter, which is
// what replaced `this`. The inert `Signature.thisParameter` never fills in.
interface Counter {
  n: number;
  bump(self: Counter, by: number): number;
}

local counter: Counter = {
  n = 0,
  bump = function(self: Counter, by: number) return self.n + by; end,
};

function counter:bump(by: number): number
  return self.n + by;
end

// Interfaces and generics still instantiate: the inert `InterfaceType.thisType`
// type parameter does not perturb signature relations or inference.
interface Box<T> {
  value: T;
  map<U>(self: Box<T>, f: (v: T) => U): Box<U>;
}

declare box: Box<number>;
declare function toStr(v: number): string;

local mapped: Box<string> = box.map(box, toStr);

// Interface `extends` heritage and construct signatures are untouched.
interface Base {
  b: number;
}
interface Derived extends Base {
  d: number;
}
declare K: { new (): Derived };
local inst: Derived = new K();

use(narrow);
use(g);
use(counter);
use(mapped);
use(inst);


//// [tluaNoThisCoexistence.lua]
-- `this` is removed, but the constructs that neighbour it must keep working.
-- This file must stay zero-error.
-- Ordinary identifier type predicates still narrow. Only the `this is T` form
-- was removed, and it is a distinct branch of the predicate machinery.
function isString(v)
  return type(v) == "string";
end
function narrow(v)
  if isString(v) then
    local s = v;
    use(s);
  end
  assertString(v);
  local t = v;
  use(t);
end
-- Lua's global environment remains available independently of `this`.
local g = _G;
local counter = {
  n = 0,
  bump = function(self, by)
    return self.n + by;
  end,
};
function counter:bump(by)
  return self.n + by;
end
local mapped = box.map(box, toStr);
local inst = new K();
use(narrow);
use(g);
use(counter);
use(mapped);
use(inst);

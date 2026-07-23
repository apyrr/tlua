//// [tests/cases/compiler/tluaOptionalChainEmit.tlua] ////

//// [tluaOptionalChainEmit.tlua]
// Optional chaining (`?.`) has no Lua equivalent. It lowers to `and`-guarded
// access (Lua `and` returns the falsy operand, and a member-bearing base can
// never be `false`, so truthiness == non-nil). Side-effecting receivers are
// captured once into a preceding `local`; lazily-evaluated positions use a
// single-eval IIFE; a discarded chain becomes `if base then access end`.

type Inner = { c: number }
type T = { b: Inner, m: () => number, get: () => (Inner | nil), [k: string]: unknown }

local function run(a: T | nil, getObj: () => (T | nil))
  // (A) copiable base, single optional link — pure `and`, no temp.
  local a1 = a?.b
  local a2 = a?.b.c
  local a3 = a?.m()
  local a4 = a?.["b"] // element access

  // Multi-name local: no single operative slot, so each chain lowers in place
  // (must not hoist — regression guard for the operativeChain nil case).
  local m1, m2 = a?.b, getObj()?.b

  // (B) side-effecting base / nested optional link, order-safe slot — hoist a
  // `local`, guard with `and`.
  local b1 = a?.b?.c
  local b2 = getObj()?.b
  local b3 = a?.get()?.c

  // (C) lazily-evaluated position (RHS of `and`) — single-eval IIFE.
  local c1 = true and getObj()?.b

  // Discarded chain in statement position — rewritten to `if ... then ... end`.
  a?.m()
  getObj()?.m()

  // Order-safe `return` slot — hoisting is unnecessary here (copiable base).
  return a?.b.c
end


//// [tluaOptionalChainEmit.lua]
-- Optional chaining (`?.`) has no Lua equivalent. It lowers to `and`-guarded
-- access (Lua `and` returns the falsy operand, and a member-bearing base can
-- never be `false`, so truthiness == non-nil). Side-effecting receivers are
-- captured once into a preceding `local`; lazily-evaluated positions use a
-- single-eval IIFE; a discarded chain becomes `if base then access end`.
local function run(a, getObj)
  -- (A) copiable base, single optional link — pure `and`, no temp.
  local a1 = a and a.b;
  local a2 = a and a.b.c;
  local a3 = a and a.m();
  local a4 = a and a["b"]; -- element access
  -- Multi-name local: no single operative slot, so each chain lowers in place
  -- (must not hoist — regression guard for the operativeChain nil case).
  local m1, m2 = a and a.b, (function()
    local _a = getObj();
    return _a and _a.b;
  end)();
  local _a = a and a.b;
  -- (B) side-effecting base / nested optional link, order-safe slot — hoist a
  -- `local`, guard with `and`.
  local b1 = _a and _a.c;
  local _b = getObj();
  local b2 = _b and _b.b;
  local _c = a and a.get();
  local b3 = _c and _c.c;
  -- (C) lazily-evaluated position (RHS of `and`) — single-eval IIFE.
  local c1 = true and (function()
    local _a = getObj();
    return _a and _a.b;
  end)();
  -- Discarded chain in statement position — rewritten to `if ... then ... end`.
  if a then
    a.m();
  end
  local _d = getObj();
  if _d then
    _d.m();
  end
  -- Order-safe `return` slot — hoisting is unnecessary here (copiable base).
  return a and a.b.c;
end

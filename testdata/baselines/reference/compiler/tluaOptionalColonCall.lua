//// [tests/cases/compiler/tluaOptionalColonCall.tlua] ////

//// [tluaOptionalColonCall.tlua]
// `obj?:m(args)` is the optional colon call: `obj and obj:m(args)`. The receiver
// is passed non-nil, since the call only happens when it is not nil, and the
// result gains `nil` like any other optional chain.

interface Box {
    value: number;
    inner: Box | nil;
    read(self: Box): number;
    set(self: Box, v: number): nil;
}

declare function maybeBox(): Box | nil;

local b = maybeBox()
local r1 = b?:read()
local r2: number = b?:read()
b?:set(1)
b?:set("x")

// A receiver that is not a plain name is evaluated once.
local t: { box?: Box } = {}
t.box?:set(2)
maybeBox()?:set(3)

// Links chain in both directions.
local r3 = b?.inner?:read()
local r4 = b?:read() == 1

// `?` and `:` are one operator; outer spaces are trivia.
local r5 = b ?: read()

// A discarded call keeps its captured receiver in a `do ... end`, so its local
// ends with the statement; a leading comment stays with the statement.
local function useMany(self: { box?: Box })
    -- first
    self.box?:set(1)
    self.box?:set(2)
end

// A chain in a non-order-safe slot runs in a function; a `...` argument is passed
// through to it.
declare function consume(a: nil, b: number): nil;
local function forward(...: number)
    consume(maybeBox()?:set(select("#", ...)), 1)
end

-- A hoisted local keeps its statement's comment above it.
local r6 = t.box?.inner?:read()

// A trailing comment stays on its statement, after the hoisted locals.
local r7 = t.box?.inner?:read() -- trailing

// `a and a:m()` keeps one value: Lua's `and` yields a single result, so an
// optional colon call of a multi-value method yields its first value (or nil).
interface Pair {
    pair(self: Pair): (number, string);
}
declare maybePair: Pair | nil;
local first, second = maybePair?:pair()
local onlyNil: nil = second

// A builtin whose result the checker refines still gains the optional nil.
declare maybeText: string | nil;
declare maybeFile: LuaFile | nil;
local words: () => string | nil = maybeText?:gmatch("%a+")
local contents: string = maybeFile?:read("*a")


//// [tluaOptionalColonCall.lua]
-- `obj?:m(args)` is the optional colon call: `obj and obj:m(args)`. The receiver
-- is passed non-nil, since the call only happens when it is not nil, and the
-- result gains `nil` like any other optional chain.
local b = maybeBox();
local r1 = b and b:read();
local r2 = b and b:read();
if b then
  b:set(1);
end
if b then
  b:set("x");
end
-- A receiver that is not a plain name is evaluated once.
local t = {};
do
  local _a = t.box;
  if _a then
    _a:set(2);
  end
end
do
  local _b = maybeBox();
  if _b then
    _b:set(3);
  end
end
-- Links chain in both directions.
local _c = b and b.inner;
local r3 = _c and _c:read();
local r4 = (b and b:read()) == 1;
-- `?` and `:` are one operator; outer spaces are trivia.
local r5 = b and b:read();
-- A discarded call keeps its captured receiver in a `do ... end`, so its local
-- ends with the statement; a leading comment stays with the statement.
local function useMany(self)
  -- first
  do
    local _a = self.box;
    if _a then
      _a:set(1);
    end
  end
  do
    local _b = self.box;
    if _b then
      _b:set(2);
    end
  end
end
local function forward(...)
  consume((function(...)
    local _a = maybeBox();
    return _a and _a:set(select("#", ...));
  end)(...), 1);
end
-- A hoisted local keeps its statement's comment above it.
local _d = t.box;
local _e = _d and _d.inner;
local r6 = _e and _e:read();
-- A trailing comment stays on its statement, after the hoisted locals.
local _f = t.box;
local _g = _f and _f.inner;
local r7 = _g and _g:read(); -- trailing
local first, second = maybePair and maybePair:pair();
local onlyNil = second;
local words = maybeText and maybeText:gmatch("%a+");
local contents = maybeFile and maybeFile:read("*a");

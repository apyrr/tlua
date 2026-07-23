//// [tests/cases/compiler/tluaMultiReturnTypeSyntax.tlua] ////

//// [tluaMultiReturnTypeSyntax.tlua]
// Bare multi-return type lists on declaration headers.
function pair(): (number, string)
  return 1, "a";
end

function triple(): (number, string, boolean)
  return 1, "a", true;
end

local function localPair(): (number, string)
  return 2, "b";
end

function bracePair(): (number, string)
  return 3, "c";
end

local exprPair = function(): (number, string)
  return 4, "d";
end;

// Single-type returns are unchanged by the list grammar.
function single(): number
  return 5;
end

function explicitPair(): (number, string)
  return 6, "f";
end

local p = pair();
local t = triple();
local l = localPair();
local b = bracePair();
local e = exprPair();
local s = single();
local x = explicitPair();

// Callable types use an explicitly parenthesized return-pack syntax.
type PairFunction = () => (number, string);
declare pairFunction: PairFunction;
local pairAlias: PairFunction = pair;

interface PairCallable {
  (): (number, string);
  pair(): (number, string);
}
declare pairCallable: PairCallable;

local function consumePair(value: number, label: string) end
consumePair(pairFunction());
consumePair(pairCallable());
consumePair(pairCallable.pair());

local function takesCallback(callback: () => (number, string), label: string)
  return callback(), label;
end

type Wrapped = Map<() => (number, string), string>;

type Nested = () => (() => (number, string), boolean);
type Complex<T> = () => (T extends string ? number : boolean, (string | nil)?, ...number);


//// [tluaMultiReturnTypeSyntax.lua]
-- Bare multi-return type lists on declaration headers.
function pair()
  return 1, "a";
end
function triple()
  return 1, "a", true;
end
local function localPair()
  return 2, "b";
end
function bracePair()
  return 3, "c";
end
local exprPair = function()
  return 4, "d";
end;
-- Single-type returns are unchanged by the list grammar.
function single()
  return 5;
end
function explicitPair()
  return 6, "f";
end
local p = pair();
local t = triple();
local l = localPair();
local b = bracePair();
local e = exprPair();
local s = single();
local x = explicitPair();
local pairAlias = pair;
local function consumePair(value, label)
end
consumePair(pairFunction());
consumePair(pairCallable());
consumePair(pairCallable.pair());
local function takesCallback(callback, label)
  return callback(), label;
end

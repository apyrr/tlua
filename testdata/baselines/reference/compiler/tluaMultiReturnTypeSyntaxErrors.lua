//// [tests/cases/compiler/tluaMultiReturnTypeSyntaxErrors.tlua] ////

//// [tluaMultiReturnTypeSyntaxErrors.tlua]
// Ordinary callable return types do not consume enclosing commas.
type H = Map<() => number, string>;
type ParenthesizedSingle = () => (string)[];
type ConditionalSingle<T> = () => (T extends string ? number : boolean);
type BadPredicate = (value: unknown) => (value is string, number);

function takesCallback(cb: () => number, more: string)
  return cb() + more.length;
end

function trailingCallback(cb: () => number,)
  return cb();
end

// An arrow return type in a call argument list does not swallow the comma.
function twoArgs(a: (x: number) => number, b: number)
  return a(b);
end
local r = twoArgs(function(x): number return x end, 1);

// A bare multiple return type must be parenthesized; the list is still
// parsed for recovery, so the pack itself checks normally.
function barePair(): number, string
  return 1, "a";
end
declare function bareAmbient(): number, ...string;
local function bareLocal(): number, string?
  return 1, nil;
end
local a, b = barePair();
local n: number = a;
local s: string = b;


//// [tluaMultiReturnTypeSyntaxErrors.lua]
function takesCallback(cb, more)
    return cb() + more.length;
end
function trailingCallback(cb)
    return cb();
end
-- An arrow return type in a call argument list does not swallow the comma.
function twoArgs(a, b)
    return a(b);
end
local r = twoArgs(function(x)
    return x;
end, 1);
-- A bare multiple return type must be parenthesized; the list is still
-- parsed for recovery, so the pack itself checks normally.
function barePair()
    return 1, "a";
end
local function bareLocal()
    return 1, nil;
end
local a, b = barePair();
local n = a;
local s = b;

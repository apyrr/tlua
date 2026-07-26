//// [tests/cases/compiler/tluaIncompleteStatement.tlua] ////

//// [tluaIncompleteStatement.tlua]
-- @strict: true
local a = 1
local b = 2
local t = { f = function() end }

-- invalid: Lua statements are calls or assignments, never bare expressions
a;
a == b;
(a);
t.f;
not a;

-- valid statement shapes
t.f();
a = b;
a, b = b, a;

-- a misspelled keyword keeps its spelling suggestion, not TLUA100057
loccl c = 1;

-- recovery fragments are not doubled up: the leftover `1` after the bad
-- binding gets no TLUA100057 on top of the declaration errors
local end = 1;

-- compile-time-only assertions erase to a bare call, which stays a valid
-- statement; a non-call under an assertion is still flagged
local function g(): number return 1 end
g() as number;
g() satisfies number;
g()!;
b as number;

-- adjacent bare statements split only by a newline are each flagged
a
b

-- a scanner error in the next statement's token does not hide this
-- statement's error (must stay last: the string swallows the rest)
a;
"unterminated


//// [tluaIncompleteStatement.lua]
-- @strict: true
local a = 1;
local b = 2;
local t = { f = function()
  end };
-- invalid: Lua statements are calls or assignments, never bare expressions
a;
a == b;
(a);
t.f;
!a;
-- valid statement shapes
t.f();
a = b;
a, b = b, a;
-- a misspelled keyword keeps its spelling suggestion, not TLUA100057
loccl;
c = 1;
-- recovery fragments are not doubled up: the leftover `1` after the bad
-- binding gets no TLUA100057 on top of the declaration errors
local ;
1;
-- compile-time-only assertions erase to a bare call, which stays a valid
-- statement; a non-call under an assertion is still flagged
local function g()
  return 1;
end
g();
g();
g();
b;
-- adjacent bare statements split only by a newline are each flagged
a;
b;
-- a scanner error in the next statement's token does not hide this
-- statement's error (must stay last: the string swallows the rest)
a;
"unterminated";

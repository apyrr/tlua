//// [tests/cases/compiler/tluaNoArray.tlua] ////

//// [tluaNoArray.tlua]
// The TypeScript array-literal/iteration syntax is removed from tlua source.
// Each form below decomposes into stock parser-recovery errors (no
// tlua-specific diagnostic), exactly like the class/generator removals.
// (Tuple *type* syntax `[A, B]` is legal — see the tluaTuple tests.)

declare function pair(): (number, string);
declare items: number[];
declare function use(x: unknown): void;

// Array literals no longer parse in expression position; use a table `{...}`.
local lit = [1, 2, 3];

// Array-literal spread. Every spread form is gone; `...` is the Lua vararg, so
// the `...` here reads as a vararg expression (legal in the main chunk) and the
// operand that follows is the error.
local spread = [0, ...items];

// Array binding patterns; use a Lua name list `local a, b = f()`.
local [a, b] = pair();

// Array destructuring assignment.
[a, b] = pair();

// `for..of`; use a Lua generic-for `for _, v in ipairs(t) do`.
for (local v of items) {
  use(v);
}

// Tuple types parse in source, and a multireturn is not a tuple value: only
// the first value survives outside a value-list tail, so this is an error.
local t: [number, string] = pair();

// `of` is an ordinary identifier again.
local of = 3;
use(of);


//// [tluaNoArray.lua]
-- The TypeScript array-literal/iteration syntax is removed from tlua source.
-- Each form below decomposes into stock parser-recovery errors (no
-- tlua-specific diagnostic), exactly like the class/generator removals.
-- (Tuple *type* syntax `[A, B]` is legal — see the tluaTuple tests.)
-- Array literals no longer parse in expression position; use a table `{...}`.
local lit = [1, 2, 3];
-- Array-literal spread. Every spread form is gone; `...` is the Lua vararg, so
-- the `...` here reads as a vararg expression (legal in the main chunk) and the
-- operand that follows is the error.
local spread = [0, ...];
items;
;
-- Array binding patterns; use a Lua name list `local a, b = f()`.
local ;
a, b;
pair();
a, b;
pair();
-- `for..of`; use a Lua generic-for `for _, v in ipairs(t) do`.
for  in () do
    local v;
    of;
    items;
    {
        use(v);
    }
    -- Tuple types parse in source, and a multireturn is not a tuple value: only
    -- the first value survives outside a value-list tail, so this is an error.
    local t = pair();
    -- `of` is an ordinary identifier again.
    local of = 3;
    use(of);
end

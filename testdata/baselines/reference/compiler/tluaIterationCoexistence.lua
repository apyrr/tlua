//// [tests/cases/compiler/tluaIterationCoexistence.tlua] ////

//// [tluaIterationCoexistence.tlua]
// TS iteration (`for..of`, array-literal spread, array destructuring) is
// removed. Lua iteration remains: the generic-for drives an iterator function,
// and a number-key table spreads and is indexed like a sequence.

declare function iterStep(t: string[], i: number): (number, string?);
declare function ipairsLike(t: string[]): typeof iterStep;

// Lua generic-for over an iterator function; the names take their positions
// from the step function's return pack.
function walk(t: string[]): number
  local count = 0;
  for i, s in ipairsLike(t) do
    count = count + i;
  end
  return count;
end

// A number-key table satisfies T[] and is indexed element-wise.
function sum(xs: number[]): number
  local total = 0;
  for i = 1, 3 do
    total = total + xs[i];
  end
  return total;
end

// A vararg pack collects into a number-key table with `{...}`.
function join(sep: string, ...: string): string {
  local parts: string[] = {...};
  return parts[1];
}
function forward(a: string, b: string): string {
  return join(",", a, b);
}


//// [tluaIterationCoexistence.lua]
-- TS iteration (`for..of`, array-literal spread, array destructuring) is
-- removed. Lua iteration remains: the generic-for drives an iterator function,
-- and a number-key table spreads and is indexed like a sequence.
-- Lua generic-for over an iterator function; the names take their positions
-- from the step function's return pack.
function walk(t)
    local count = 0;
    for i, s in ipairsLike(t) do
        count = count + i;
    end
    return count;
end
-- A number-key table satisfies T[] and is indexed element-wise.
function sum(xs)
    local total = 0;
    for i = 1, 3 do
        total = total + xs[i];
    end
    return total;
end
-- A vararg pack collects into a number-key table with `{...}`.
function join(sep, ...) {
    local parts = { ... };
    return parts[1];
}
function forward(a, b) {
    return join(",", a, b);
}

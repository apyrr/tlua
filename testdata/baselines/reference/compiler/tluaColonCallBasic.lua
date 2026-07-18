//// [tests/cases/compiler/tluaColonCallBasic.tlua] ////

//// [tluaColonCallBasic.tlua]
local M = { count = 0 };

function M:inc(by: number): number
  self.count = self.count + by;
  return self.count;
end

function M:read(): number
  return self.count;
end

// The colon call passes the receiver as the implicit first argument.
local a = M:inc(1);
local b = M:read();

// It is interchangeable with the explicit form.
local c = M.inc(M, 1);

// A dotted target chains: `a.b:c()` calls through the member table.
local outer = { inner = M };
local d = outer.inner:inc(2);

// Chained colon calls: each link is its own receiver.
interface Chain {
  next(self: Chain): Chain;
  value(self: Chain): number;
}
declare chain: Chain;
local e = chain:next():next():value();

// The colon form also fits an explicit self-typed member on a table value.
local t = {
  base = 10,
  plus = function(self: { base: number }, n: number) return self.base + n; end,
};
local f = t:plus(5);

// A receiver that survived its `?.` guards is not nil at call time: the
// implicit first argument carries the narrowed type, so no spurious mismatch.
local maybe: { inner: typeof M } | nil = { inner = M };
local g = maybe?.inner:read();


//// [tluaColonCallBasic.lua]
local M = { count = 0 };
function M:inc(by)
    self.count = self.count + by;
    return self.count;
end
function M:read()
    return self.count;
end
-- The colon call passes the receiver as the implicit first argument.
local a = M:inc(1);
local b = M:read();
-- It is interchangeable with the explicit form.
local c = M.inc(M, 1);
-- A dotted target chains: `a.b:c()` calls through the member table.
local outer = { inner = M };
local d = outer.inner:inc(2);
local e = chain:next():next():value();
-- The colon form also fits an explicit self-typed member on a table value.
local t = {
    base = 10,
    plus = function(self, n)
        return self.base + n;
    end,
};
local f = t:plus(5);
-- A receiver that survived its `?.` guards is not nil at call time: the
-- implicit first argument carries the narrowed type, so no spurious mismatch.
local maybe = { inner = M };
local g = maybe and maybe.inner:read();

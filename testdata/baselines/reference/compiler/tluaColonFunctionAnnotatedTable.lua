//// [tests/cases/compiler/tluaColonFunctionAnnotatedTable.tlua] ////

//// [tluaColonFunctionAnnotatedTable.tlua]
interface Ops {
  good(self: Ops, a: number): number;
  noSelf(a: number): number;
}

local ops: Ops = {
  good = function(self: Ops, a: number) return a; end,
  noSelf = function(a: number) return a; end,
};

// An annotated table gets an assignability check against the declared member,
// and the colon form supplies the receiver the interface asks for.
function ops:good(a: number): number
  return a;
end

// `noSelf` takes no receiver, so the colon form does not fit it.
function ops:noSelf(a: number): number
  return a;
end

function ops:absent(): void
end


//// [tluaColonFunctionAnnotatedTable.lua]
local ops = {
  good = function(self, a)
    return a;
  end,
  noSelf = function(a)
    return a;
  end,
};
-- An annotated table gets an assignability check against the declared member,
-- and the colon form supplies the receiver the interface asks for.
function ops:good(a)
  return a;
end
-- `noSelf` takes no receiver, so the colon form does not fit it.
function ops:noSelf(a)
  return a;
end
function ops:absent()
end

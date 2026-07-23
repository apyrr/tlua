//// [tests/cases/compiler/tluaDottedFunctionAnnotatedTable.tlua] ////

//// [tluaDottedFunctionAnnotatedTable.tlua]
interface Ops {
  good(a: number): number;
  wrong(a: number): number;
}

local ops: Ops = { good = function(a: number) return a; end, wrong = function(a: number) return a; end };

function ops.good(a: number): number
  return a;
end

function ops.wrong(a: string): string
  return a;
end

function ops.absent(): void
end


//// [tluaDottedFunctionAnnotatedTable.lua]
local ops = { good = function(a)
    return a;
  end, wrong = function(a)
    return a;
  end };
function ops.good(a)
  return a;
end
function ops.wrong(a)
  return a;
end
function ops.absent()
end

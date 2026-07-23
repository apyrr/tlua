//// [tests/cases/compiler/tluaGenericForPack.tlua] ////

//// [tluaGenericForPack.tlua]
// Overloaded iterators union their return packs.
declare function overloaded(s: string, c: number): (number, string);
declare function overloaded(s: string, c: string): (string, number);

function useOverloaded(): void
  for k, v in overloaded do
    k;
    v;
  end
end

// A variadic pack tail projects with nil.
declare function variadicIter(s: void, c: number): (number, ...string);

function useVariadic(): void
  for k, v, w in variadicIter do
    k;
    v;
    w;
  end
end

// Definite assignment: the names are assigned by the loop.
declare function pair(t: number[], i: number): (number, number?);

function defAssign(t: number[]): number
  local total = 0;
  for k, v in pair, t, nil do
    total = total + k;
    if v ~= nil then
      total = total + v;
    end
  end
  return total;
end


//// [tluaGenericForPack.lua]
function useOverloaded()
  for k, v in overloaded do
    k;
    v;
  end
end
function useVariadic()
  for k, v, w in variadicIter do
    k;
    v;
    w;
  end
end
function defAssign(t)
  local total = 0;
  for k, v in pair, t, nil do
    total = total + k;
    if v ~= nil then
      total = total + v;
    end
  end
  return total;
end

//// [tests/cases/compiler/tluaMultiReturnDeclareAndOverloads.tlua] ////

//// [tluaMultiReturnDeclareAndOverloads.tlua]
// Ambient declarations take bare return-type lists.
declare function ambient(): (number, string);
declare function ambientVariadic(n: number): (number, ...string);

// Overload signatures terminate with `;`, so the list grammar is unambiguous.
function over(a: number): (number, string);
function over(a: string): (string, number);
function over(a: number | string): (number | string, string | number) {
  if (type(a) == "number") {
    return a, "n";
  }
  return a, 0;
}

local a = ambient();
local v = ambientVariadic(1);
local o1 = over(1);
local o2 = over("x");


//// [tluaMultiReturnDeclareAndOverloads.lua]
function over(a) {
    if (type(a) == "number") {
        return a, "n";
    }
    return a, 0;
}
local a = ambient();
local v = ambientVariadic(1);
local o1 = over(1);
local o2 = over("x");

//// [tests/cases/compiler/tluaNoGenerators.tlua] ////

//// [tluaNoGenerators.tlua]
// Generators are removed in tlua: `function*`, `yield`, `yield*`, and async
// generators no longer parse. Use the Lua coroutine library directly instead.

function* gen(): unknown
  yield 1;
  yield* more;
end

async function* asyncGen(): unknown
  yield 4;
end

local obj = {
  *method() {
    yield 5;
  },
};

// `yield` is now an ordinary identifier.
function usesYieldName(): number
  local yield = 7;
  return yield;
end


//// [tluaNoGenerators.lua]
-- Generators are removed in tlua: `function*`, `yield`, `yield*`, and async
-- generators no longer parse. Use the Lua coroutine library directly instead.
function ()
   * gen();
  unknown;
  yield;
  1;
  yield * more;
end
async function ()
   * asyncGen();
  unknown;
  yield;
  4;
end
local obj = {
    * method(), {
    yield, 5,
  },
};
-- `yield` is now an ordinary identifier.
function usesYieldName()
  local yield = 7;
  return yield;
end

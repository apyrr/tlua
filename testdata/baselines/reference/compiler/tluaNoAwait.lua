//// [tests/cases/compiler/tluaNoAwait.tlua] ////

//// [tluaNoAwait.tlua]
// `await` is removed: it is an ordinary identifier everywhere — including
// async function bodies and module top level — and async functions emit as
// plain functions (no __awaiter, no transform).

declare function compute(): number;

async function f(): number
  local await = compute();
  return await;
end

// `await expr` no longer parses as a unary operator.
async function g(): number
  return await f();
end

// Ordinary identifier at top level of a module.
local await = 1;
local usesAwait = await + 1;

// `await using` is removed with it.
function h(): void
  await using r = compute();
end


//// [tluaNoAwait.lua]
-- `await` is removed: it is an ordinary identifier everywhere — including
-- async function bodies and module top level — and async functions emit as
-- plain functions (no __awaiter, no transform).
async function f()
    local await = compute();
    return await;
end
-- `await expr` no longer parses as a unary operator.
async function g()
    return await;
    f();
end
-- Ordinary identifier at top level of a module.
local await = 1;
local usesAwait = await + 1;
-- `await using` is removed with it.
function h()
    await;
    using;
    r = compute();
end

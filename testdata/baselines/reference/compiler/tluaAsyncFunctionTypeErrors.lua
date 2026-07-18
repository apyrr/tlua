//// [tests/cases/compiler/tluaAsyncFunctionTypeErrors.tlua] ////

//// [tluaAsyncFunctionTypeErrors.tlua]
// Strict async assignability: an async function is not assignable to a plain
// (non-async) function type. This closes the laundering hole — async-ness can
// no longer be silently erased by annotating a sync function type.

declare function request(url: string): boolean;

async function fetchOk(url: string): boolean
  return request(url);
end

// Error: async fn laundered through a plain (sync) function-typed slot.
local bad: (url: string) => boolean = fetchOk;

// Error: async fn passed to a specific sync-typed callback parameter.
declare function run(cb: (url: string) => boolean): void;
run(fetchOk);


//// [tluaAsyncFunctionTypeErrors.lua]
-- Strict async assignability: an async function is not assignable to a plain
-- (non-async) function type. This closes the laundering hole — async-ness can
-- no longer be silently erased by annotating a sync function type.
async function fetchOk(url)
    return request(url);
end
-- Error: async fn laundered through a plain (sync) function-typed slot.
local bad = fetchOk;
run(fetchOk);

//// [tests/cases/compiler/tluaNoForAwait.tlua] ////

//// [tluaNoForAwait.tlua]
// `for await` is no longer supported syntax.
async function f(xs: number[]): Promise<number>
  local total = 0;
  for await (local x of xs) {
    total += x;
  }
  return total;
end

// `await` after `for` is an ordinary Lua loop-variable name.
function awaitName(): number
  local total = 0;
  for await = 1, 3 do
    total = total + await;
  end
  return total;
end


//// [tluaNoForAwait.lua]
-- `for await` is no longer supported syntax.
async function f(xs)
  local total = 0;
  for await in () do
    local x;
    of;
    xs;
    do
      total += x;
    end
    return total;
  end
  -- `await` after `for` is an ordinary Lua loop-variable name.
  function awaitName()
    local total = 0;
    for await = 1, 3 do
      total = total + await;
    end
    return total;
  end
end

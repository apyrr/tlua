//// [tests/cases/compiler/tluaWhileFlow.tlua] ////

//// [tluaWhileFlow.tlua]
// Assignment in the body is visible after the loop (loop-label join).
function joinAfterLoop(flag: boolean): number | string
  local x: number | string = 1;
  while flag do
    x = "s";
    flag = false;
  end
  return x;
end

// Code after `while true do break end` is reachable.
function reachableAfterBreak(): number
  while true do
    break;
  end
  return 1;
end

// Code after `while true do end` (no break) is unreachable.
function unreachableAfterInfinite(): number
  while true do
  end
  return 1;
end

// Statements after break are unreachable inside the body.
function unreachableAfterBreakInBody(): number
  local n = 0;
  while true do
    break;
    n = 1;
  end
  return n;
end


//// [tluaWhileFlow.lua]
-- Assignment in the body is visible after the loop (loop-label join).
function joinAfterLoop(flag)
    local x = 1;
    while flag do
        x = "s";
        flag = false;
    end
    return x;
end
-- Code after `while true do break end` is reachable.
function reachableAfterBreak()
    while true do
        break;
    end
    return 1;
end
-- Code after `while true do end` (no break) is unreachable.
function unreachableAfterInfinite()
    while true do
    end
    return 1;
end
-- Statements after break are unreachable inside the body.
function unreachableAfterBreakInBody()
    local n = 0;
    while true do
        break;
        n = 1;
    end
    return n;
end

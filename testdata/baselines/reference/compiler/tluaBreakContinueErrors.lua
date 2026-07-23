//// [tests/cases/compiler/tluaBreakContinueErrors.tlua] ////

//// [tluaBreakContinueErrors.tlua]
// break outside any loop is a grammar error.
function breakOutside(): void
  break;
end

// continue outside any loop is a grammar error.
function continueOutside(): void
  continue;
end

// break nested in a Lua if inside a Lua while is fine.
function breakNested(flag: boolean): number
  while true do
    if flag then
      break;
    end
    flag = true;
  end
  return 1;
end

// continue keeps working inside Lua loops (TS keyword, no rejection guard).
function continueInLua(n: number): number
  local i = 0;
  local total = 0;
  while i < n do
    i = i + 1;
    if i == 2 then
      continue;
    end
    total = total + i;
  end
  return total;
end


//// [tluaBreakContinueErrors.lua]
-- break outside any loop is a grammar error.
function breakOutside()
  break;
end
-- continue outside any loop is a grammar error.
function continueOutside()
  continue;
end
-- break nested in a Lua if inside a Lua while is fine.
function breakNested(flag)
  while true do
    if flag then
      break;
    end
    flag = true;
  end
  return 1;
end
-- continue keeps working inside Lua loops (TS keyword, no rejection guard).
function continueInLua(n)
  local i = 0;
  local total = 0;
  while i < n do
    i = i + 1;
    if i == 2 then
      continue;
    end
    total = total + i;
  end
  return total;
end

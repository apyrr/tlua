//// [tests/cases/compiler/tluaWhileStatement.tlua] ////

//// [tluaWhileStatement.tlua]
function sum(limit: number): number
  local total = 0;
  local i = 0;
  while i < limit do
    total = total + i;
    i = i + 1;
  end
  return total;
end

// Parenthesized condition with `do` is Lua.
function parens(flag: boolean): number
  local n = 0;
  while (flag) do
    flag = false;
    n = 1;
  end
  return n;
end

// break exits the loop; one-liner bodies parse (ASI before `end`).
function firstOver(limit: number): number
  local i = 0;
  while true do
    if i > limit then break end
    i = i + 1;
  end
  return i;
end

// Nested while loops.
function grid(w: number, h: number): number
  local count = 0;
  local x = 0;
  while x < w do
    local y = 0;
    while y < h do
      count = count + 1;
      y = y + 1;
    end
    x = x + 1;
  end
  return count;
end


//// [tluaWhileStatement.lua]
function sum(limit)
  local total = 0;
  local i = 0;
  while i < limit do
    total = total + i;
    i = i + 1;
  end
  return total;
end
-- Parenthesized condition with `do` is Lua.
function parens(flag)
  local n = 0;
  while (flag) do
    flag = false;
    n = 1;
  end
  return n;
end
-- break exits the loop; one-liner bodies parse (ASI before `end`).
function firstOver(limit)
  local i = 0;
  while true do
    if i > limit then
      break;
    end
    i = i + 1;
  end
  return i;
end
-- Nested while loops.
function grid(w, h)
  local count = 0;
  local x = 0;
  while x < w do
    local y = 0;
    while y < h do
      count = count + 1;
      y = y + 1;
    end
    x = x + 1;
  end
  return count;
end

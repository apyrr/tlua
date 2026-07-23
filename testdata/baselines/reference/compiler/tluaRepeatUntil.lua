//// [tests/cases/compiler/tluaRepeatUntil.tlua] ////

//// [tluaRepeatUntil.tlua]
declare function poll(): number | nil;

// Body locals are visible in the until-condition.
function bodyLocal(n: number): number
  repeat
    local step = n - 1;
    n = step;
  until step <= 0;
  return n;
end

// ...but not after the loop.
function afterLoop(): number
  repeat
    local secret = 1;
  until true;
  return secret;
end

// break exits; continue targets the condition.
function loops(n: number): number
  local total = 0;
  repeat
    n = n - 1;
    if n == 3 then
      continue;
    end
    if n <= 0 then
      break;
    end
    total = total + n;
  until false;
  return total;
end

// One-liner form.
function once(): number
  local n = 0;
  repeat n = n + 1 until n > 0;
  return n;
end

// Nested repeat.
function nested(n: number): number
  repeat
    repeat
      n = n - 1;
    until n <= 10;
  until n <= 0;
  return n;
end


//// [tluaRepeatUntil.lua]
-- Body locals are visible in the until-condition.
function bodyLocal(n)
  repeat
    local step = n - 1;
    n = step;
  until step <= 0;
  return n;
end
-- ...but not after the loop.
function afterLoop()
  repeat
    local secret = 1;
  until true;
  return secret;
end
-- break exits; continue targets the condition.
function loops(n)
  local total = 0;
  repeat
    n = n - 1;
    if n == 3 then
      continue;
    end
    if n <= 0 then
      break;
    end
    total = total + n;
  until false;
  return total;
end
-- One-liner form.
function once()
  local n = 0;
  repeat
    n = n + 1;
  until n > 0;
  return n;
end
-- Nested repeat.
function nested(n)
  repeat
    repeat
      n = n - 1;
    until n <= 10;
  until n <= 0;
  return n;
end

//// [tests/cases/compiler/tluaDoBlock.tlua] ////

//// [tluaDoBlock.tlua]
// Standalone do-block: locals scoped inside, shadowing works.
function scoped(): number
  local x = 1;
  do
    local x = 2;
    local y = x + 1;
    y;
  end
  return x;
end

// Block locals are not visible outside.
function leak(): number
  do
    local hidden = 1;
  end
  return hidden;
end

// Nested do-blocks; TS statements inside.
function nested(flag: boolean): number
  local n = 0;
  do
    do
      if flag then
        n = 1;
      end
    end
  end
  return n;
end

// A repeat-until loop coexists with Lua do-blocks in one body.
function coexist(flag: boolean): number
  repeat
    flag = false;
  until not (flag);
  do
    local x = 1;
    x;
  end
  return 0;
end

// A plain Lua if coexisting with the do-blocks above.
function canonical(flag: boolean): number
  if flag then
    return 1;
  end
  return 0;
end


//// [tluaDoBlock.lua]
-- Standalone do-block: locals scoped inside, shadowing works.
function scoped()
    local x = 1;
    do
        local x = 2;
        local y = x + 1;
        y;
    end
    return x;
end
-- Block locals are not visible outside.
function leak()
    do
        local hidden = 1;
    end
    return hidden;
end
-- Nested do-blocks; TS statements inside.
function nested(flag)
    local n = 0;
    do
        do
            if flag then
                n = 1;
            end
        end
    end
    return n;
end
-- A repeat-until loop coexists with Lua do-blocks in one body.
function coexist(flag)
    repeat
        flag = false;
    until !(flag);
    do
        local x = 1;
        x;
    end
    return 0;
end
-- A plain Lua if coexisting with the do-blocks above.
function canonical(flag)
    if flag then
        return 1;
    end
    return 0;
end

//// [tests/cases/compiler/tluaControlFlowCoexistence.tlua] ////

//// [tluaControlFlowCoexistence.tlua]
// Lua control-flow forms parse and emit.
function ifForms(flag: boolean): number
  if flag then
    return 1;
  end
  if flag then return 2 end
  if flag then return 3 else return 4 end
  return 0;
end

// Loop forms, including a repeat/until standing in for the old do-while.
function loopForms(flag: boolean): number
  while flag do
    flag = false;
  end
  repeat
    flag = false;
  until not (flag)
  return 0;
end

// Nested control flow.
function nested(flag: boolean): number
  if flag then
    if flag then
      return 1;
    end
  end
  while flag do
    while flag do
      flag = false;
    end
  end
  return 0;
end


//// [tluaControlFlowCoexistence.lua]
-- Lua control-flow forms parse and emit.
function ifForms(flag)
  if flag then
    return 1;
  end
  if flag then
    return 2;
  end
  if flag then
    return 3;
  else
    return 4;
  end
  return 0;
end
-- Loop forms, including a repeat/until standing in for the old do-while.
function loopForms(flag)
  while flag do
    flag = false;
  end
  repeat
    flag = false;
  until !(flag);
  return 0;
end
-- Nested control flow.
function nested(flag)
  if flag then
    if flag then
      return 1;
    end
  end
  while flag do
    while flag do
      flag = false;
    end
  end
  return 0;
end

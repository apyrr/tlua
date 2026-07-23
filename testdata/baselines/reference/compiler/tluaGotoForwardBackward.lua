//// [tests/cases/compiler/tluaGotoForwardBackward.tlua] ////

//// [tluaGotoForwardBackward.tlua]
// A forward jump skips the statements between it and the label.
function forward(n: number): number
    local total = 0;
    if n > 0 then goto done end
    total = 1;
    ::done::
    return total;
end

// A backward jump re-enters the block, so the label needs a loop edge.
function backward(n: number): number
    ::top::
    n = n - 1;
    if n > 0 then goto top end
    return n;
end

// A label is visible in nested blocks, and an inner label shadows an outer one.
function shadow(flag: boolean): number
    ::same::
    do
        ::same::
        if flag then goto same end
    end
    return 0;
end

// A goto in a loop body targeting a label after the loop leaves the loop.
function outOfLoop(n: number): string
    while n > 0 do
        n = n - 1;
        if n == 2 then goto found end
    end
    return "none";
    ::found::
    return "found";
end


//// [tluaGotoForwardBackward.lua]
-- A forward jump skips the statements between it and the label.
function forward(n)
  local total = 0;
  if n > 0 then
    goto done;
  end
  total = 1;
  ::done::
  return total;
end
-- A backward jump re-enters the block, so the label needs a loop edge.
function backward(n)
  ::top::
  n = n - 1;
  if n > 0 then
    goto top;
  end
  return n;
end
-- A label is visible in nested blocks, and an inner label shadows an outer one.
function shadow(flag)
  ::same::
  do
    ::same::
    if flag then
      goto same;
    end
  end
  return 0;
end
-- A goto in a loop body targeting a label after the loop leaves the loop.
function outOfLoop(n)
  while n > 0 do
    n = n - 1;
    if n == 2 then
      goto found;
    end
  end
  return "none";
  ::found::
  return "found";
end

//// [tests/cases/compiler/tluaGenericFor.tlua] ////

//// [tluaGenericFor.tlua]
declare function iterStep(t: string[], i: number): (number, string?);
declare function ipairsLike(t: string[]): typeof iterStep;

// The first name drops nil (a nil first value ends the loop); the second
// keeps the iterator's declared nil-ability and narrows normally.
function walk(t: string[]): number
  local count = 0;
  for i, s in ipairsLike(t) do
    count = count + i;
    if s ~= nil then
      count = count + s.length;
    end
  end
  return count;
end

// Single-name form.
function keys(t: string[]): number
  local total = 0;
  for i in ipairsLike(t) do
    total = total + i;
  end
  return total;
end

// Explist form: iterator, state, control.
declare function next(t: string[], k: number | nil): (number, string?);

function explist(t: string[]): number
  local total = 0;
  for k, v in next, t, nil do
    total = total + k;
  end
  return total;
end

// Names past the iterator's values read as nil.
function extraName(t: string[]): void
  for i, s, extra in ipairsLike(t) do
    local x: nil = extra;
    x;
  end
end

// Annotated names check against their positional value.
function annotated(t: string[]): void
  for i: number, s: string | nil in ipairsLike(t) do
    i;
    s;
  end
end

// break and TS statements inside the body.
function firstKey(t: string[]): number
  for i in ipairsLike(t) do
    return i;
  end
  return 0;
end


//// [tluaGenericFor.lua]
-- The first name drops nil (a nil first value ends the loop); the second
-- keeps the iterator's declared nil-ability and narrows normally.
function walk(t)
  local count = 0;
  for i, s in ipairsLike(t) do
    count = count + i;
    if s ~= nil then
      count = count + s.length;
    end
  end
  return count;
end
-- Single-name form.
function keys(t)
  local total = 0;
  for i in ipairsLike(t) do
    total = total + i;
  end
  return total;
end
function explist(t)
  local total = 0;
  for k, v in next, t, nil do
    total = total + k;
  end
  return total;
end
-- Names past the iterator's values read as nil.
function extraName(t)
  for i, s, extra in ipairsLike(t) do
    local x = extra;
    x;
  end
end
-- Annotated names check against their positional value.
function annotated(t)
  for i, s in ipairsLike(t) do
    i;
    s;
  end
end
-- break and TS statements inside the body.
function firstKey(t)
  for i in ipairsLike(t) do
    return i;
  end
  return 0;
end

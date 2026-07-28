//// [tests/cases/conformance/ported/discriminantsAndTypePredicates.tlua] ////

//// [discriminantsAndTypePredicates.tlua]
-- ported from tests/cases/compiler/discriminantsAndTypePredicates.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface A { type: 'A' }
interface B { type: 'B' }

function isA(x: A | B): x is A
    return x.type == 'A'
end

function isB(x: A | B): x is B
    return x.type == 'B'
end

function foo1(x: A | B): any
    local _ = x -- A | B
    if isA(x) then
        return x -- A
    end
    local _ = x -- B
    if isB(x) then
        return x -- B
    end
    local _ = x -- never
end

function foo2(x: A | B): any
    local _ = x -- A | B
    if x.type == 'A' then
        return x -- A
    end
    local _ = x -- B
    if x.type == 'B' then
        return x -- B
    end
    local _ = x -- never
end


//// [discriminantsAndTypePredicates.lua]
-- ported from tests/cases/compiler/discriminantsAndTypePredicates.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function isA(x)
  return x.type == 'A';
end
function isB(x)
  return x.type == 'B';
end
function foo1(x)
  local _ = x; -- A | B
  if isA(x) then
    return x; -- A
  end
  local _ = x; -- B
  if isB(x) then
    return x; -- B
  end
  local _ = x; -- never
end
function foo2(x)
  local _ = x; -- A | B
  if x.type == 'A' then
    return x; -- A
  end
  local _ = x; -- B
  if x.type == 'B' then
    return x; -- B
  end
  local _ = x; -- never
end

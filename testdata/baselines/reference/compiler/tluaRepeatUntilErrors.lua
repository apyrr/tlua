//// [tests/cases/compiler/tluaRepeatUntilErrors.tlua] ////

//// [tluaRepeatUntilErrors.tlua]
// Missing `until`.
function missingUntil(): number
  repeat
    local n = 1;
    n;
  return 0;
end

// `until` outside a repeat.
function strayUntil(): number
  until true;
  return 0;
end

// A void condition cannot be tested.
declare function nothing(): void;

function voidCondition(): number
  repeat
    local n = 1;
    n;
  until nothing();
  return 0;
end


//// [tluaRepeatUntilErrors.lua]
-- Missing `until`.
function missingUntil()
  repeat
    local n = 1;
    n;
    return 0;
  until ;
end
-- `until` outside a repeat.
function strayUntil()
  true;
  return 0;
end
function voidCondition()
  repeat
    local n = 1;
    n;
  until nothing();
  return 0;
end

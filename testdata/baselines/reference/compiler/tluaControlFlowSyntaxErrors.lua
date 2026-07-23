//// [tests/cases/compiler/tluaControlFlowSyntaxErrors.tlua] ////

//// [tluaControlFlowSyntaxErrors.tlua]
// Reserved words cannot be binding names.
local then = 1;
local until = 2;
function repeat(): void
end

// Missing `then`.
function missingThen(flag: boolean): number
  if flag
    return 1;
  end
  return 0;
end

// Missing `end`.
function missingEnd(flag: boolean): number
  if flag then
    return 1;
  return 0;
end

// elseif after else.
function elseifAfterElse(n: number): number
  if n == 1 then
    return 1;
  else
    return 2;
  elseif n == 3 then
    return 3;
  end
end


//// [tluaControlFlowSyntaxErrors.lua]
-- Reserved words cannot be binding names.
local ;
1;
local ;
2;
function ()
  repeat
    ();
    void ;
  until ;
end
-- Missing `then`.
function missingThen(flag)
  if flag then
    return 1;
  end
  return 0;
end
-- Missing `end`.
function missingEnd(flag)
  if flag then
    return 1;
    return 0;
  end
  -- elseif after else.
  function elseifAfterElse(n)
    if n == 1 then
      return 1;
    else
      return 2;
      n == 3;
      return 3;
    end
  end
end

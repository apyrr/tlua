//// [tests/cases/compiler/tluaFunctionSyntax.tlua] ////

//// [tluaFunctionSyntax.tlua]
function topLua(a: number): number
  return a;
end

function emptyLua(): void
end

function braceStillWorks(a: number): number
  return a + 1;
end

function withBraceStatement(flag: boolean): number
  local result = 0;
  if (flag) then
    result = braceStillWorks(1);
  end
  return result;
end

local function localLua(a: number): number
  return a + 1;
end

function withNestedBraceFunction(): number
  function nestedBrace(): number
    return 1;
  end
  return nestedBrace();
end

function withNestedLuaFunction(): number
  function nestedLua(): number
    return 2;
  end
  return nestedLua();
end

function outerLua(value: number): number
  local function innerLua(delta: number): number
    return value + delta;
  end
  return innerLua(localLua(1));
end

function preservedOverload(x: string): string;
function preservedOverload(x: number): number;
function preservedOverload(x: any)
  return x;
end

topLua(1);
emptyLua();
braceStillWorks(1);
withBraceStatement(true);
localLua(1);
withNestedBraceFunction();
withNestedLuaFunction();
outerLua(1);
preservedOverload("x");


//// [tluaFunctionSyntax.lua]
function topLua(a)
    return a;
end
function emptyLua()
end
function braceStillWorks(a)
    return a + 1;
end
function withBraceStatement(flag)
    local result = 0;
    if (flag) then
        result = braceStillWorks(1);
    end
    return result;
end
local function localLua(a)
    return a + 1;
end
function withNestedBraceFunction()
    function nestedBrace()
        return 1;
    end
    return nestedBrace();
end
function withNestedLuaFunction()
    function nestedLua()
        return 2;
    end
    return nestedLua();
end
function outerLua(value)
    local function innerLua(delta)
        return value + delta;
    end
    return innerLua(localLua(1));
end
function preservedOverload(x)
    return x;
end
topLua(1);
emptyLua();
braceStillWorks(1);
withBraceStatement(true);
localLua(1);
withNestedBraceFunction();
withNestedLuaFunction();
outerLua(1);
preservedOverload("x");

//// [tests/cases/compiler/tluaMultiReturnDeclarationEmit.tlua] ////

//// [defs.d.tlua]
declare function fromDts(): (number, string);
declare function fromDtsVariadic(): (boolean, ...number);

//// [main.tlua]
function annotated(): (number, string)
  return 1, "a";
end

function inferred(cond: boolean)
  if cond then
    return 1, "a";
  end
  return 2;
end

function variadic(): (number, ...string)
  return 1, "a", "b";
end

function optional(): (number, string?)
  return 1;
end

type PairCallback = () => (number, string);

interface PairCallable {
  (): (number, string);
  pair(): (number, string);
}

// Consuming hand-written d.tlua packs proves the declared syntax parses.
local fromDeclaration = fromDts();
local fromDeclarationVariadic = fromDtsVariadic();


//// [main.lua]
function annotated()
    return 1, "a";
end
function inferred(cond)
    if cond then
        return 1, "a";
    end
    return 2;
end
function variadic()
    return 1, "a", "b";
end
function optional()
    return 1;
end
-- Consuming hand-written d.tlua packs proves the declared syntax parses.
local fromDeclaration = fromDts();
local fromDeclarationVariadic = fromDtsVariadic();

//// [tests/cases/compiler/tluaMultiReturnDeclarationEmit.tlua] ////

//// [defs.d.tlua]
declare function fromDts(): (number, string);
declare function fromDtsVariadic(): (boolean, ...number);

//// [main.tlua]
function annotated(): (number, string) {
  return 1, "a";
}

function inferred(cond: boolean) {
  if (cond) {
    return 1, "a";
  }
  return 2;
}

function variadic(): (number, ...string) {
  return 1, "a", "b";
}

function optional(): (number, string?) {
  return 1;
}

type PairCallback = () => (number, string);

interface PairCallable {
  (): (number, string);
  pair(): (number, string);
}

// Consuming hand-written d.tlua packs proves the declared syntax parses.
local fromDeclaration = fromDts();
local fromDeclarationVariadic = fromDtsVariadic();


//// [main.lua]
function annotated() {
    return 1, "a";
}
function inferred(cond) {
    if (cond) {
        return 1, "a";
    }
    return 2;
}
function variadic() {
    return 1, "a", "b";
}
function optional() {
    return 1;
}
-- Consuming hand-written d.tlua packs proves the declared syntax parses.
local fromDeclaration = fromDts();
local fromDeclarationVariadic = fromDtsVariadic();

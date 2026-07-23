//// [tests/cases/compiler/tluaNoTry.tlua] ////

//// [tluaNoTry.tlua]
// try/catch/finally are removed. The words are ordinary identifiers, so this
// reports ordinary syntax errors without constructing legacy AST nodes.
try {
  throw "boom";
} catch (e) {
  local caught = e;
} finally {
  local cleaned = true;
}


//// [tluaNoTry.lua]
-- try/catch/finally are removed. The words are ordinary identifiers, so this
-- reports ordinary syntax errors without constructing legacy AST nodes.
try;
do
  throw "boom";
end
catch(e);
do
  local caught = e;
end
finally;
do
  local cleaned = true;
end

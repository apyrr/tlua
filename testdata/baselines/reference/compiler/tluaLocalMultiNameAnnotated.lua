//// [tests/cases/compiler/tluaLocalMultiNameAnnotated.tlua] ////

//// [tluaLocalMultiNameAnnotated.tlua]
function pair(): (number, string)
  return 1, "a";
end

// Annotations check their positional value.
local a: number, b: string = pair();

// Position 0 mismatch.
local c: string, d = pair();

// Position 1 mismatch.
local e, f: boolean = pair();

// A name past the values must accept nil.
local g: number, h: string, i: boolean = pair();

// nil-accepting annotations are fine past the end.
local j: number, k: string, l: boolean | nil = pair();

// Mixed annotated and inferred names.
local m, n: string = pair();


//// [tluaLocalMultiNameAnnotated.lua]
function pair()
  return 1, "a";
end
-- Annotations check their positional value.
local a, b = pair();
-- Position 0 mismatch.
local c, d = pair();
-- Position 1 mismatch.
local e, f = pair();
-- A name past the values must accept nil.
local g, h, i = pair();
-- nil-accepting annotations are fine past the end.
local j, k, l = pair();
-- Mixed annotated and inferred names.
local m, n = pair();

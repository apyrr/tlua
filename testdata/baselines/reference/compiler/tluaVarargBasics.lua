//// [tests/cases/compiler/tluaVarargBasics.tlua] ////

//// [tluaVarargBasics.tlua]
// A vararg parameter is a bare `...` in final position. It has no binding name,
// and its annotation is the *element* type of the pack, not an array type.
function annotated(...: number): void
end

// An unannotated vararg is an `any` pack.
function bare(...): void
end

// Varargs coexist with fixed parameters.
function withFixed(label: string, ...: string): string
  return label;
end

// `...` expands in a value list, so both names bind. A rest-only pack may be
// empty, so each name is `number | nil`.
function twoNames(...: number): (number | nil, number | nil)
  local a, b = ...;
  return a, b;
end

// A vararg function accepts zero or more trailing arguments.
annotated();
annotated(1, 2, 3);
withFixed("x");
withFixed("x", "y", "z");
bare();
bare(1, "two", true);

// The element type is checked.
annotated("nope");
withFixed("x", 1);


//// [tluaVarargBasics.lua]
-- A vararg parameter is a bare `...` in final position. It has no binding name,
-- and its annotation is the *element* type of the pack, not an array type.
function annotated(...)
end
-- An unannotated vararg is an `any` pack.
function bare(...)
end
-- Varargs coexist with fixed parameters.
function withFixed(label, ...)
  return label;
end
-- `...` expands in a value list, so both names bind. A rest-only pack may be
-- empty, so each name is `number | nil`.
function twoNames(...)
  local a, b = ...;
  return a, b;
end
-- A vararg function accepts zero or more trailing arguments.
annotated();
annotated(1, 2, 3);
withFixed("x");
withFixed("x", "y", "z");
bare();
bare(1, "two", true);
-- The element type is checked.
annotated("nope");
withFixed("x", 1);

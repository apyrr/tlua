//// [tests/cases/conformance/ported/stringLiteralCheckedInIf02.tlua] ////

//// [stringLiteralCheckedInIf02.tlua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf02.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: the upstream first-element access `[0]` maps to Lua's 1-based first slot `[1]`, and
--   tlua enables noUncheckedIndexedAccess by default, so the indexed read (and therefore the
--   inferred return type of `f`) is `S | nil` rather than upstream's `S`.

type S = "a" | "b"
type T = S[] | S

function isS(t: T): t is S
    return t == "a" or t == "b"
end

function f(foo: T)
    if isS(foo) then
        return foo
    else
        return foo[1]
    end
end


//// [stringLiteralCheckedInIf02.lua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf02.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: the upstream first-element access `[0]` maps to Lua's 1-based first slot `[1]`, and
--   tlua enables noUncheckedIndexedAccess by default, so the indexed read (and therefore the
--   inferred return type of `f`) is `S | nil` rather than upstream's `S`.
function isS(t)
  return t == "a" or t == "b";
end
function f(foo)
  if isS(foo) then
    return foo;
  else
    return foo[1];
  end
end

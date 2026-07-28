//// [tests/cases/conformance/ported/stringLiteralCheckedInIf01.tlua] ////

//// [stringLiteralCheckedInIf01.tlua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: the upstream first-element access `[0]` maps to Lua's 1-based first slot `[1]`, and
--   tlua enables noUncheckedIndexedAccess by default, so the indexed read (and therefore the
--   inferred return type of `f`) is `S | nil` rather than upstream's `S`.

type S = "a" | "b"
type T = S[] | S

function f(foo: T)
    if foo == "a" then
        return foo
    elseif foo == "b" then
        return foo
    else
        return (foo as S[])[1]
    end
end


//// [stringLiteralCheckedInIf01.lua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- note: the upstream first-element access `[0]` maps to Lua's 1-based first slot `[1]`, and
--   tlua enables noUncheckedIndexedAccess by default, so the indexed read (and therefore the
--   inferred return type of `f`) is `S | nil` rather than upstream's `S`.
function f(foo)
  if foo == "a" then
    return foo;
  elseif foo == "b" then
    return foo;
  else
    return foo[1];
  end
end

//// [tests/cases/conformance/ported/stringLiteralCheckedInIf01.tlua] ////

//// [stringLiteralCheckedInIf01.tlua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

type S = "a" | "b"
type T = S[] | S

function f(foo: T)
    if foo == "a" then
        return foo
    elseif foo == "b" then
        return foo
    else
        return (foo as S[])[0]
    end
end


//// [stringLiteralCheckedInIf01.lua]
-- ported from tests/cases/conformance/types/stringLiteral/stringLiteralCheckedInIf01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f(foo)
  if foo == "a" then
    return foo;
  elseif foo == "b" then
    return foo;
  else
    return foo[0];
  end
end

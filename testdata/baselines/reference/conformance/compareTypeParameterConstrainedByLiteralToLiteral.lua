//// [tests/cases/conformance/ported/compareTypeParameterConstrainedByLiteralToLiteral.tlua] ////

//// [compareTypeParameterConstrainedByLiteralToLiteral.tlua]
-- ported from tests/cases/compiler/compareTypeParameterConstrainedByLiteralToLiteral.ts

-- Test for #26758

function foo<T extends "a" | "b">(t: T)
    local _ = t == "a" -- Should be allowed
    local _ = t == "x" -- Should be error
end


//// [compareTypeParameterConstrainedByLiteralToLiteral.lua]
-- ported from tests/cases/compiler/compareTypeParameterConstrainedByLiteralToLiteral.ts
-- Test for #26758
function foo(t)
  local _ = t == "a"; -- Should be allowed
  local _ = t == "x"; -- Should be error
end

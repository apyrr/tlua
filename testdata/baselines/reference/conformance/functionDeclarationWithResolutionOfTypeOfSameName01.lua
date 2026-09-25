//// [tests/cases/conformance/ported/functionDeclarationWithResolutionOfTypeOfSameName01.tlua] ////

//// [functionDeclarationWithResolutionOfTypeOfSameName01.tlua]
-- ported from tests/cases/compiler/functionDeclarationWithResolutionOfTypeOfSameName01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)

interface f {
}

function f()
    local _ = f as f
end


//// [functionDeclarationWithResolutionOfTypeOfSameName01.lua]
-- ported from tests/cases/compiler/functionDeclarationWithResolutionOfTypeOfSameName01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: `<T>x` type assertions as `x as T` (tlua has only the `as` form)
function f()
  local _ = f;
end

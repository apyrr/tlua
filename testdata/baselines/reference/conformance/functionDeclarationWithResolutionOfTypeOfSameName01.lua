//// [tests/cases/conformance/ported/functionDeclarationWithResolutionOfTypeOfSameName01.tlua] ////

//// [functionDeclarationWithResolutionOfTypeOfSameName01.tlua]
-- ported from tests/cases/compiler/functionDeclarationWithResolutionOfTypeOfSameName01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface f {
}

function f()
    local _ = <f>f
end


//// [functionDeclarationWithResolutionOfTypeOfSameName01.lua]
-- ported from tests/cases/compiler/functionDeclarationWithResolutionOfTypeOfSameName01.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f()
  local _ = f;
end

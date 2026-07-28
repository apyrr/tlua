//// [tests/cases/conformance/ported/functionOverloadCompatibilityWithVoid02.tlua] ////

//// [functionOverloadCompatibilityWithVoid02.tlua]
-- ported from tests/cases/conformance/functions/functionOverloadCompatibilityWithVoid02.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f(x: string): void;
function f(x: string): number
    return 0
end


//// [functionOverloadCompatibilityWithVoid02.lua]
-- ported from tests/cases/conformance/functions/functionOverloadCompatibilityWithVoid02.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f(x)
  return 0;
end

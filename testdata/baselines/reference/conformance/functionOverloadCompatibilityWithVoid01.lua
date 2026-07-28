//// [tests/cases/conformance/ported/functionOverloadCompatibilityWithVoid01.tlua] ////

//// [functionOverloadCompatibilityWithVoid01.tlua]
-- ported from tests/cases/conformance/functions/functionOverloadCompatibilityWithVoid01.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)

function f(x: string): number;
function f(x: string): void
    return
end


//// [functionOverloadCompatibilityWithVoid01.lua]
-- ported from tests/cases/conformance/functions/functionOverloadCompatibilityWithVoid01.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
function f(x)
  return;
end

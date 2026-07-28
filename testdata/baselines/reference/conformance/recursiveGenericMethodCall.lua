//// [tests/cases/conformance/ported/recursiveGenericMethodCall.tlua] ////

//// [recursiveGenericMethodCall.tlua]
-- ported from tests/cases/compiler/recursiveGenericMethodCall.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

interface Generator<T> { (): T }

function Generate<T>(func: Generator<T>): T
    return Generate(func)
end


//// [recursiveGenericMethodCall.lua]
-- ported from tests/cases/compiler/recursiveGenericMethodCall.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function Generate(func)
  return Generate(func);
end

//// [tests/cases/conformance/ported/noReachabilityErrorsOnEmptyStatement.tlua] ////

//// [noReachabilityErrorsOnEmptyStatement.tlua]
-- ported from tests/cases/compiler/noReachabilityErrorsOnEmptyStatement.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target); C-style function braces rewritten to Lua end-blocks

function foo()
    return 1;;
end


//// [noReachabilityErrorsOnEmptyStatement.lua]
-- ported from tests/cases/compiler/noReachabilityErrorsOnEmptyStatement.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target); C-style function braces rewritten to Lua end-blocks
function foo()
  return 1;
  ;
end

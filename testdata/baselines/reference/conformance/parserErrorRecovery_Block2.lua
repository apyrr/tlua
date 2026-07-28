//// [tests/cases/conformance/ported/parserErrorRecovery_Block2.tlua] ////

//// [parserErrorRecovery_Block2.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/Blocks/parserErrorRecovery_Block2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f()
    ¬
    return
end


//// [parserErrorRecovery_Block2.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/Blocks/parserErrorRecovery_Block2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f()
  return;
end

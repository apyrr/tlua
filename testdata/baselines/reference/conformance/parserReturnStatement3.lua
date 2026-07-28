//// [tests/cases/conformance/ported/parserReturnStatement3.tlua] ////

//// [parserReturnStatement3.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/ReturnStatements/parserReturnStatement3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function f()
   return
end


//// [parserReturnStatement3.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/ReturnStatements/parserReturnStatement3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function f()
  return;
end

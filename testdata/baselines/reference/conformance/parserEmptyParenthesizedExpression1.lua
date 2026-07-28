//// [tests/cases/conformance/ported/parserEmptyParenthesizedExpression1.tlua] ////

//// [parserEmptyParenthesizedExpression1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/parserEmptyParenthesizedExpression1.ts
-- dropped: the original JS-only `toString` member access was replaced with a generic call while preserving empty-parenthesized-expression recovery


function getObj()
   ().foo()
end


//// [parserEmptyParenthesizedExpression1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/parserEmptyParenthesizedExpression1.ts
-- dropped: the original JS-only `toString` member access was replaced with a generic call while preserving empty-parenthesized-expression recovery
function getObj()
  ().foo();
end

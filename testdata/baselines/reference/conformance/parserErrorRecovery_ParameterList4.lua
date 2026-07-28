//// [tests/cases/conformance/ported/parserErrorRecovery_ParameterList4.tlua] ////

//// [parserErrorRecovery_ParameterList4.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ParameterLists/parserErrorRecovery_ParameterList4.ts
-- dropped: @target: es2015 directive; the C-style function body was rewritten to a Lua end-block

function f(a,¬)
end


//// [parserErrorRecovery_ParameterList4.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ParameterLists/parserErrorRecovery_ParameterList4.ts
-- dropped: @target: es2015 directive; the C-style function body was rewritten to a Lua end-block
function f(a)
end

//// [tests/cases/conformance/ported/parserErrorRecovery_ArgumentList2.tlua] ////

//// [parserErrorRecovery_ArgumentList2.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ArgumentLists/parserErrorRecovery_ArgumentList2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo()
	bar(;
end


//// [parserErrorRecovery_ArgumentList2.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ArgumentLists/parserErrorRecovery_ArgumentList2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo()
  bar();
end

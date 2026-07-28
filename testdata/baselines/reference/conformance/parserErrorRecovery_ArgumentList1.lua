//// [tests/cases/conformance/ported/parserErrorRecovery_ArgumentList1.tlua] ////

//// [parserErrorRecovery_ArgumentList1.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ArgumentLists/parserErrorRecovery_ArgumentList1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function foo()
	bar(
	return x
end


//// [parserErrorRecovery_ArgumentList1.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/ErrorRecovery/ArgumentLists/parserErrorRecovery_ArgumentList1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function foo()
  bar();
  return x;
end

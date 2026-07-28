//// [tests/cases/conformance/ported/parserIfStatement2.tlua] ////

//// [parserIfStatement2.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/parserIfStatement2.ts
-- dropped: upstream @target: es2015 directive (tlua uses its default target)
if a then
end


//// [parserIfStatement2.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/parserIfStatement2.ts
-- dropped: upstream @target: es2015 directive (tlua uses its default target)
if a then
end

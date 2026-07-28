//// [tests/cases/conformance/ported/parserSkippedTokens8.tlua] ////

//// [parserSkippedTokens8.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/SkippedTokens/parserSkippedTokens8.ts
-- dropped: @target: es2015 directive; tlua uses the esnext default

;
/*foo*/ \ /*bar*/


//// [parserSkippedTokens8.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/SkippedTokens/parserSkippedTokens8.ts
-- dropped: @target: es2015 directive; tlua uses the esnext default
;

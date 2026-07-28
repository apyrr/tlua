//// [tests/cases/conformance/ported/parserSkippedTokens5.tlua] ////

//// [parserSkippedTokens5.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/SkippedTokens/parserSkippedTokens5.ts
-- dropped: @target: es2015 directive; tlua uses the esnext default

\ --[[foo]] ;


//// [parserSkippedTokens5.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/SkippedTokens/parserSkippedTokens5.ts
-- dropped: @target: es2015 directive; tlua uses the esnext default
;

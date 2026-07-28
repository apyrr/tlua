//// [tests/cases/conformance/ported/tripleSlashInCommentNotParsed.tlua] ////

//// [tripleSlashInCommentNotParsed.tlua]
-- ported from tests/cases/compiler/tripleSlashInCommentNotParsed.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: `void 0` expression (void-value expressions are deleted; replaced with an inert local value)

--[[
/// <reference path="non-existing-file.d.ts" />
]]
local _ = 0


//// [tripleSlashInCommentNotParsed.lua]
-- ported from tests/cases/compiler/tripleSlashInCommentNotParsed.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: `void 0` expression (void-value expressions are deleted; replaced with an inert local value)
--[[
/// <reference path="non-existing-file.d.ts" />
]]
local _ = 0;

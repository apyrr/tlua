//// [tests/cases/conformance/ported/parenthesizedExpressionInternalComments.tlua] ////

//// [parenthesizedExpressionInternalComments.tlua]
-- ported from tests/cases/compiler/parenthesizedExpressionInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target); JavaScript bare expression statements, rewritten as local assignments

local _ = --[[1]](--[[2]] "foo" --[[3]])--[[4]]

-- open
local _ = --[[1]](
    -- next
    --[[2]]"foo"
    -- close
    --[[3]])--[[4]]


//// [parenthesizedExpressionInternalComments.lua]
-- ported from tests/cases/compiler/parenthesizedExpressionInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target); JavaScript bare expression statements, rewritten as local assignments
local _ = --[[1]] ( --[[2]]"foo" --[[3]]); --[[4]]
-- open
local _ = --[[1]] (
-- next
--[[2]] "foo"
-- close
--[[3]] ); --[[4]]

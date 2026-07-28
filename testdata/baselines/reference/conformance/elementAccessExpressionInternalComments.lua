//// [tests/cases/conformance/ported/elementAccessExpressionInternalComments.tlua] ////

//// [elementAccessExpressionInternalComments.tlua]
-- ported from tests/cases/compiler/elementAccessExpressionInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript Array global replaced with a neutral local table shape to preserve element-access comment coverage

interface ArrayLike {
    toString?: function
}
local Array: ArrayLike = {}
local _ = --[[0]] Array --[[1]][ --[[2]] "toString" --[[3]] ] --[[4]]; --[[5]]

local _ = --[[0]] Array
    -- single line
    --[[1]][ --[[2]] "toString"
    -- single line
    --[[3]] ] --[[4]]


//// [elementAccessExpressionInternalComments.lua]
-- ported from tests/cases/compiler/elementAccessExpressionInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: JavaScript Array global replaced with a neutral local table shape to preserve element-access comment coverage
local Array = {};
local _ = --[[0]] Array --[[1]][ --[[2]]"toString" --[[3]]] --[[4]]; --[[5]]
local _ = --[[0]] Array
-- single line
--[[1]] [ --[[2]]"toString"
-- single line
--[[3]] ]; --[[4]]

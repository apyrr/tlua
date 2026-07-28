//// [tests/cases/conformance/ported/ifStatementInternalComments.tlua] ////

//// [ifStatementInternalComments.tlua]
-- ported from tests/cases/compiler/ifStatementInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

if --2
    ( --3
        true --4
    ) then --5
end

if --2
    ( --3
        true --4
    ) then --5
else --7
end --6


//// [ifStatementInternalComments.lua]
-- ported from tests/cases/compiler/ifStatementInternalComments.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
if -- 2
 ( -- 3
true -- 4
) then
end
if -- 2
 ( -- 3
true -- 4
) then
else
end -- 6

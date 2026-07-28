//// [tests/cases/conformance/ported/duplicateLabel1.tlua] ////

//// [duplicateLabel1.tlua]
-- ported from tests/cases/compiler/duplicateLabel1.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)


::target::
::target::
while true do
end


//// [duplicateLabel1.lua]
-- ported from tests/cases/compiler/duplicateLabel1.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
::target::
::target::
while true do
end

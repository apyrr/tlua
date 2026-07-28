//// [tests/cases/conformance/ported/breakNotInIterationOrSwitchStatement2.tlua] ////

//// [breakNotInIterationOrSwitchStatement2.tlua]
-- ported from tests/cases/compiler/breakNotInIterationOrSwitchStatement2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

while true do
    function f()
        break
    end
end


//// [breakNotInIterationOrSwitchStatement2.lua]
-- ported from tests/cases/compiler/breakNotInIterationOrSwitchStatement2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
while true do
  function f()
    break;
  end
end

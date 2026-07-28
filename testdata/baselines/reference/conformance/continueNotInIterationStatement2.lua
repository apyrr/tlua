//// [tests/cases/conformance/ported/continueNotInIterationStatement2.tlua] ////

//// [continueNotInIterationStatement2.tlua]
-- ported from tests/cases/compiler/continueNotInIterationStatement2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

while true do
    function f()
        continue
    end
end


//// [continueNotInIterationStatement2.lua]
-- ported from tests/cases/compiler/continueNotInIterationStatement2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
while true do
  function f()
    continue;
  end
end

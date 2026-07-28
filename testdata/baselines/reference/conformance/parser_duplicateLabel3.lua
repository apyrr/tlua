//// [tests/cases/conformance/ported/parser_duplicateLabel3.tlua] ////

//// [parser_duplicateLabel3.tlua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/LabeledStatements/parser_duplicateLabel3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewrote JavaScript labels to Lua ::label:: syntax


::target::
while true do
    function f()
        ::target::
        while true do
        end
    end
end


//// [parser_duplicateLabel3.lua]
-- ported from tests/cases/conformance/parser/ecmascript5/Statements/LabeledStatements/parser_duplicateLabel3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewrote JavaScript labels to Lua ::label:: syntax
::target::
while true do
  function f()
    ::target::
    while true do
    end
  end
end

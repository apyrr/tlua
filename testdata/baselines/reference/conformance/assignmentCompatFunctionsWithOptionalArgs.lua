//// [tests/cases/conformance/ported/assignmentCompatFunctionsWithOptionalArgs.tlua] ////

//// [assignmentCompatFunctionsWithOptionalArgs.tlua]
-- ported from tests/cases/compiler/assignmentCompatFunctionsWithOptionalArgs.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: the incidental missing-implementation diagnostic from the bodyless
--   TypeScript declaration; an empty implementation keeps the optional-field
--   assignment checks as the test subject.

function foo(x: { id: number; name?: string }): void
end

foo({ id = 1234 })                 -- Ok
foo({ id = 1234, name = "hello" })  -- Ok
foo({ id = 1234, name = false })    -- Error, name of wrong type
foo({ name = "hello" })            -- Error, id required but missing


//// [assignmentCompatFunctionsWithOptionalArgs.lua]
-- ported from tests/cases/compiler/assignmentCompatFunctionsWithOptionalArgs.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: the incidental missing-implementation diagnostic from the bodyless
--   TypeScript declaration; an empty implementation keeps the optional-field
--   assignment checks as the test subject.
function foo(x)
end
foo({ id = 1234 }); -- Ok
foo({ id = 1234, name = "hello" }); -- Ok
foo({ id = 1234, name = false }); -- Error, name of wrong type
foo({ name = "hello" }); -- Error, id required but missing

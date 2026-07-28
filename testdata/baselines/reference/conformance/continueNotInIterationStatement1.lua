//// [tests/cases/conformance/ported/continueNotInIterationStatement1.tlua] ////

//// [continueNotInIterationStatement1.tlua]
-- ported from tests/cases/compiler/continueNotInIterationStatement1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

continue


//// [continueNotInIterationStatement1.lua]
-- ported from tests/cases/compiler/continueNotInIterationStatement1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
continue;

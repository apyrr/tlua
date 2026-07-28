//// [tests/cases/conformance/ported/assignmentCompatability46.tlua] ////

//// [assignmentCompatability46.tlua]
-- ported from tests/cases/compiler/assignmentCompatability46.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

declare function fn(x: never): void

fn({ 1, 2, 3 })
fn({ a = 1, b = 2 })


//// [assignmentCompatability46.lua]
-- ported from tests/cases/compiler/assignmentCompatability46.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
fn({ 1, 2, 3 });
fn({ a = 1, b = 2 });

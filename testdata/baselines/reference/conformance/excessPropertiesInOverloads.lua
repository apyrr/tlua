//// [tests/cases/conformance/ported/excessPropertiesInOverloads.tlua] ////

//// [excessPropertiesInOverloads.tlua]
-- ported from tests/cases/compiler/excessPropertiesInOverloads.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (strict checking is enabled by default)

declare function fn(a: { x: string }): nil
declare function fn(a: { y: string }): nil
fn({ z = 3, a = 3 })


//// [excessPropertiesInOverloads.lua]
-- ported from tests/cases/compiler/excessPropertiesInOverloads.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: true directive (strict checking is enabled by default)
fn({ z = 3, a = 3 });

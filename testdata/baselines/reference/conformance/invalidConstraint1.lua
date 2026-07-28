//// [tests/cases/conformance/ported/invalidConstraint1.tlua] ////

//// [invalidConstraint1.tlua]
-- ported from tests/cases/compiler/invalidConstraint1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua strictness is fixed by the harness)
-- dropped: declaration output directive (unsupported); constraint checking is retained

function f<T, U extends { a: T }>()
    return nil
end
f<string, { a: number }>() -- should error


//// [invalidConstraint1.lua]
-- ported from tests/cases/compiler/invalidConstraint1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: @strict: false directive (tlua strictness is fixed by the harness)
-- dropped: declaration output directive (unsupported); constraint checking is retained
function f()
  return nil;
end
f(); -- should error

//// [tests/cases/conformance/ported/genericUnboundedTypeParamAssignability.tlua] ////

//// [genericUnboundedTypeParamAssignability.tlua]
-- ported from tests/cases/compiler/genericUnboundedTypeParamAssignability.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Object.prototype.toString behavior (JavaScript-only lib surface)

function f1<T>(o: T)
end

function f2<T extends {}>(o: T)
end

function f3<T extends Record<string, any>>(o: T)
end

function user<T>(t: T)
    f1(t)
    f2(t) -- error in strict, unbounded T doesn't satisfy the constraint
    f3(t) -- error in strict, unbounded T doesn't satisfy the constraint
end


//// [genericUnboundedTypeParamAssignability.lua]
-- ported from tests/cases/compiler/genericUnboundedTypeParamAssignability.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: Object.prototype.toString behavior (JavaScript-only lib surface)
function f1(o)
end
function f2(o)
end
function f3(o)
end
function user(t)
  f1(t);
  f2(t); -- error in strict, unbounded T doesn't satisfy the constraint
  f3(t); -- error in strict, unbounded T doesn't satisfy the constraint
end

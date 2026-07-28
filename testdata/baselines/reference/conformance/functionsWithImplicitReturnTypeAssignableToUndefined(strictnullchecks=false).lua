//// [tests/cases/conformance/ported/functionsWithImplicitReturnTypeAssignableToUndefined.tlua] ////

//// [functionsWithImplicitReturnTypeAssignableToUndefined.tlua]
-- ported from tests/cases/compiler/functionsWithImplicitReturnTypeAssignableToUndefined.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- dropped: Math.random() conditions, replaced with boolean parameters because Math is a JavaScript runtime global


function f1(condition: boolean): unknown
    if condition then
        return true
    end
end

type MyUnknown = unknown
function f2(condition: boolean): unknown
    if condition then
        return true
    end
end

function f3(): any
end

function f4(): void
end

function f5(condition: boolean): {}
    if condition then
        return {}
    end
end

function f6(condition: boolean): Record<string, any>
    if condition then
        return { foo = true }
    end
end

function f7(condition: boolean): nil
    if condition then
        return nil
    end
end

function f8(condition: boolean): string | nil
    if condition then
        return "foo"
    end
end


//// [functionsWithImplicitReturnTypeAssignableToUndefined.lua]
-- ported from tests/cases/compiler/functionsWithImplicitReturnTypeAssignableToUndefined.ts
-- dropped: @target: es2015 directive (tlua defaults to the latest target)
-- dropped: Math.random() conditions, replaced with boolean parameters because Math is a JavaScript runtime global
function f1(condition)
  if condition then
    return true;
  end
end
function f2(condition)
  if condition then
    return true;
  end
end
function f3()
end
function f4()
end
function f5(condition)
  if condition then
    return {};
  end
end
function f6(condition)
  if condition then
    return { foo = true };
  end
end
function f7(condition)
  if condition then
    return nil;
  end
end
function f8(condition)
  if condition then
    return "foo";
  end
end

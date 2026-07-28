//// [tests/cases/conformance/ported/noImplicitReturnsWithoutReturnExpression.tlua] ////

//// [noImplicitReturnsWithoutReturnExpression.tlua]
-- ported from tests/cases/compiler/noImplicitReturnsWithoutReturnExpression.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)

function isMissingReturnExpression(): number
    return
end

function isMissingReturnExpression2(): any
    return
end

function isMissingReturnExpression3(): number | void
    return
end

function isMissingReturnExpression4(): void
    return
end

function isMissingReturnExpression5(x)
    if x then
        return 0
    else
        return
    end
end


//// [noImplicitReturnsWithoutReturnExpression.lua]
-- ported from tests/cases/compiler/noImplicitReturnsWithoutReturnExpression.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
function isMissingReturnExpression()
  return;
end
function isMissingReturnExpression2()
  return;
end
function isMissingReturnExpression3()
  return;
end
function isMissingReturnExpression4()
  return;
end
function isMissingReturnExpression5(x)
  if x then
    return 0;
  else
    return;
  end
end

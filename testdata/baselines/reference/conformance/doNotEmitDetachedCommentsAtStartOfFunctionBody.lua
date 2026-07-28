//// [tests/cases/conformance/ported/doNotEmitDetachedCommentsAtStartOfFunctionBody.tlua] ////

//// [doNotEmitDetachedCommentsAtStartOfFunctionBody.tlua]
-- ported from tests/cases/compiler/doNotEmitDetachedCommentsAtStartOfFunctionBody.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)


function foo1()
    -- Single line comment

    return 42
end

function foo2()
    --[[

        multi line
        comment
    ]]

    return 42
end

function foo3()
    -- Single line comment with more than one blank line


    return 42
end

function foo4()
    --[[

        multi line comment with more than one blank line
    ]]

    return 42
end


//// [doNotEmitDetachedCommentsAtStartOfFunctionBody.lua]
function foo1()
  return 42;
end
function foo2()
  return 42;
end
function foo3()
  return 42;
end
function foo4()
  return 42;
end

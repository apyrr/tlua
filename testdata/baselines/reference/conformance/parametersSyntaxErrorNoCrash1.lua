//// [tests/cases/conformance/ported/parametersSyntaxErrorNoCrash1.tlua] ////

//// [parametersSyntaxErrorNoCrash1.tlua]
-- ported from tests/cases/compiler/parametersSyntaxErrorNoCrash1.ts
-- dropped: @target: es2015 and @strict directives (tlua defaults to the latest target and strict mode)

-- @noEmit: true
-- @noTypesAndSymbols: true

function identity<T>(arg: T: T
    return arg
end


//// [parametersSyntaxErrorNoCrash1.lua]
-- ported from tests/cases/compiler/parametersSyntaxErrorNoCrash1.ts
-- dropped: @target: es2015 and @strict directives (tlua defaults to the latest target and strict mode)
-- @noEmit: true
-- @noTypesAndSymbols: true
function identity(arg, T)
  return arg;
end

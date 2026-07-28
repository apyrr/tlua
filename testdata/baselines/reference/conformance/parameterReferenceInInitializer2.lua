//// [tests/cases/conformance/ported/parameterReferenceInInitializer2.tlua] ////

//// [parameterReferenceInInitializer2.tlua]
-- ported from tests/cases/compiler/parameterReferenceInInitializer2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: emit coverage suppressed because the generated function parameter is not valid Lua.
function Example(x = function(x: any) return x end) -- Error: parameter 'x' cannot be referenced in its initializer
end


//// [parameterReferenceInInitializer2.lua]
-- ported from tests/cases/compiler/parameterReferenceInInitializer2.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: emit coverage suppressed because the generated function parameter is not valid Lua.
function Example(x)
  if x == nil then
    x = function(x)
      return x;
    end;
  end
end

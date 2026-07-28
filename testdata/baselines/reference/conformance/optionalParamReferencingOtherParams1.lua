//// [tests/cases/conformance/ported/optionalParamReferencingOtherParams1.tlua] ////

//// [optionalParamReferencingOtherParams1.tlua]
-- ported from tests/cases/compiler/optionalParamReferencingOtherParams1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: emit coverage suppressed because the generated function parameters are not valid Lua.

function strange(x: number, y = x * 1, z = x + y)
    return z
end


//// [optionalParamReferencingOtherParams1.lua]
-- ported from tests/cases/compiler/optionalParamReferencingOtherParams1.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- dropped: emit coverage suppressed because the generated function parameters are not valid Lua.
function strange(x, y, z)
  if y == nil then
    y = x * 1;
  end
  if z == nil then
    z = x + y;
  end
  return z;
end

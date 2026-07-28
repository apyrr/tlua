//// [tests/cases/conformance/ported/optionalParamReferencingOtherParams3.tlua] ////

//// [optionalParamReferencingOtherParams3.tlua]
-- ported from tests/cases/compiler/optionalParamReferencingOtherParams3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: the upstream bare expression statements `a;` / `b;` became `local _ =` bindings (bare expression statements are invalid in tlua)

function right(a = b, b = a)
    local _ = a
    local _ = b
end


//// [optionalParamReferencingOtherParams3.lua]
-- ported from tests/cases/compiler/optionalParamReferencingOtherParams3.ts
-- dropped: @target: es2015 directive (tlua defaults to latest target)
-- rewritten: the upstream bare expression statements `a;` / `b;` became `local _ =` bindings (bare expression statements are invalid in tlua)
function right(a, b)
  if a == nil then
    a = b;
  end
  if b == nil then
    b = a;
  end
  local _ = a;
  local _ = b;
end

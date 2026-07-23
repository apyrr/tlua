//// [tests/cases/compiler/tluaSynthesizedLuaIf.tlua] ////

//// [tluaSynthesizedLuaIf.tlua]
function getObject(): { value: number | nil }
  return { value = nil }
end

function defaultValue(value = getObject().value ||= 1)
  return value
end


//// [tluaSynthesizedLuaIf.lua]
function getObject()
  return { value = nil };
end
function defaultValue(value)
  local _a;
  if value == void 0 then
    value = (_a = getObject()).value or (_a.value = 1);
  end
  return value;
end

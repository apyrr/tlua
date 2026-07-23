//// [tests/cases/compiler/tluaNoWithStatement.tlua] ////

//// [tluaNoWithStatement.tlua]
// With statements are removed, and `with` is an ordinary identifier.
with (scope) {
  value;
}

function usesWith(): number
  local with = 3
  return with
end


//// [tluaNoWithStatement.lua]
-- With statements are removed, and `with` is an ordinary identifier.
with(scope);
do
    value;
end
function usesWith()
    local with = 3;
    return with;
end

//// [tests/cases/compiler/tluaNoImportAttributes.tlua] ////

//// [module.tlua]
return { value = 1 }

//// [main.tlua]
// Import types accept only a module specifier; their former attributes
// argument is not part of tlua syntax.
type Module = import("./module", { with: { type: "json" } })

function usesAssert(): number
  local assert = 4
  return assert
end


//// [module.lua]
return { value = 1 };
//// [main.lua]
do
    with;
    do
        type;
        "json";
    end
end
function usesAssert()
    local assert = 4;
    return assert;
end

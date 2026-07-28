//// [tests/cases/conformance/ported/arrayLiteralInNonVarArgParameter.tlua] ////

//// [arrayLiteralInNonVarArgParameter.tlua]
-- ported from tests/cases/compiler/arrayLiteralInNonVarArgParameter.ts
-- dropped: the TypeScript rest-array annotation is tlua's scalar vararg annotation, and the [] array literal is a {} table constructor.


function panic(val: string[], ...: string)
end

panic({}, "one", "two")


//// [arrayLiteralInNonVarArgParameter.lua]
-- ported from tests/cases/compiler/arrayLiteralInNonVarArgParameter.ts
-- dropped: the TypeScript rest-array annotation is tlua's scalar vararg annotation, and the [] array literal is a {} table constructor.
function panic(val, ...)
end
panic({}, "one", "two");

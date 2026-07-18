//// [tests/cases/compiler/contextuallyTypedFunctionOptionalAndRest.tlua] ////

//// [contextuallyTypedFunctionOptionalAndRest.tlua]
local f: () => void = (a?, ...b) => {};


//// [contextuallyTypedFunctionOptionalAndRest.lua]
local f = function(a, ...b)
end;

//// [tests/cases/compiler/contextuallyTypedFunctionOptionalAndRest.tlua] ////

//// [contextuallyTypedFunctionOptionalAndRest.tlua]
local f: () => void = function(a?, ...b) end;


//// [contextuallyTypedFunctionOptionalAndRest.lua]
local f = function(a, ...b)
end;

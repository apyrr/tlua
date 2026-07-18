//// [tests/cases/compiler/jsxLibraryManagedAttributesUnexpectedType.tsx] ////

//// [jsxLibraryManagedAttributesUnexpectedType.tsx]
type JsxLibraryManagedAttributes = number;

local C = nil as any as () => any;

local x = <C />;


//// [jsxLibraryManagedAttributesUnexpectedType.lua]
local C = nil;
local x = React.createElement(C, nil);

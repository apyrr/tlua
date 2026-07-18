//// [tests/cases/compiler/jsxElementTypeUnexpectedType.tsx] ////

//// [jsxElementTypeUnexpectedType.tsx]
type JsxElementType = number;

local C = nil as any as () => any;

local x = <C />;


//// [jsxElementTypeUnexpectedType.lua]
local C = nil;
local x = React.createElement(C, nil);

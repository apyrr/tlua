//// [tests/cases/compiler/jsxIntrinsicElementsTypeArgumentErrorSkipsTrivia.tsx] ////

//// [jsxIntrinsicElementsTypeArgumentErrorSkipsTrivia.tsx]
interface JsxElement {}
interface JsxIntrinsicElements { div: { children?: any }; }
declare React: { createElement: any };

local a = <div<   number> />;

local b = <div<
    number> />;


//// [jsxIntrinsicElementsTypeArgumentErrorSkipsTrivia.lua]
local a = React.createElement("div", nil);
local b = React.createElement("div", nil);

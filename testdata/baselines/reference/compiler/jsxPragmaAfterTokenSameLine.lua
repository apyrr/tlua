//// [tests/cases/compiler/jsxPragmaAfterTokenSameLine.tsx] ////

//// [jsxPragmaAfterTokenSameLine.tsx]
/** Authored by foo@example.com @jsx h */
declare h: any;
declare React: any;
declare Fragment: any;
interface JsxElement {}

local x = <Fragment></Fragment>;


//// [jsxPragmaAfterTokenSameLine.lua]
local x = React.createElement(Fragment, nil);

//// [tests/cases/compiler/jsxPragmaAfterTags.tsx] ////

//// [jsxPragmaAfterTags.tsx]
/**
 * @fileoverview comment
 * @jsx h
 */
declare h: any;
declare Fragment: any;
interface JsxElement {}

local x = <Fragment></Fragment>;


//// [jsxPragmaAfterTags.lua]
local x = h(Fragment, nil);

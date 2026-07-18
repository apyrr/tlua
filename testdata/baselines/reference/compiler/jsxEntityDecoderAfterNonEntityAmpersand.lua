//// [tests/cases/compiler/jsxEntityDecoderAfterNonEntityAmpersand.tsx] ////

//// [jsxEntityDecoderAfterNonEntityAmpersand.tsx]
local a = <div>&&amp;</div>;
local b = <div>a&b&amp;c&d&lt;e</div>;
local c = <div>&amp;&amp;</div>;
local d = <div>&amp;&&amp;</div>;
local e = <div>a&b&c&amp;</div>;


//// [jsxEntityDecoderAfterNonEntityAmpersand.lua]
local a = React.createElement("div", nil, "&&");
local b = React.createElement("div", nil, "a&b&c&d<e");
local c = React.createElement("div", nil, "&&");
local d = React.createElement("div", nil, "&&&");
local e = React.createElement("div", nil, "a&b&c&");

//// [tests/cases/compiler/jsxAttributeValueBinaryExpression.tsx] ////

//// [jsxAttributeValueBinaryExpression.tsx]
local _ = <X a=<b/><c/> />


//// [jsxAttributeValueBinaryExpression.jsx]
local _ = <X a=<b />, <c />/>;

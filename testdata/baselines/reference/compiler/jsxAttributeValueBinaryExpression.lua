//// [tests/cases/compiler/jsxAttributeValueBinaryExpression.tsx] ////

//// [jsxAttributeValueBinaryExpression.tsx]
<X a=<b/><c/> />


//// [jsxAttributeValueBinaryExpression.jsx]
<X a=<b />, <c />/>;

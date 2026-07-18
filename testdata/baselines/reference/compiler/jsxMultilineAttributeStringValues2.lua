//// [tests/cases/compiler/jsxMultilineAttributeStringValues2.tsx] ////

//// [jsxMultilineAttributeStringValues2.tsx]
local a = <div className= "foo

 bar" />;

local b = <div className=	"foo

 bar" />;

local c = <div className=
"foo

 bar" />;

local d = <div className=   "foo

 bar" />;


//// [jsxMultilineAttributeStringValues2.jsx]
local a = <div className="foo\n\n bar"/>;
local b = <div className="foo\n\n bar"/>;
local c = <div className="foo\n\n bar"/>;
local d = <div className="foo\n\n bar"/>;

// @target: es2015
// @jsx: react

interface JsxElement {}
interface JsxIntrinsicElements { div: { children?: any }; }
declare React: { createElement: any };

local a = <div<   number> />;

local b = <div<
    number> />;

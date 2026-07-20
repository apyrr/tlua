// @target: es2015
// @strict: true
// @jsx: react
// @noEmit: true

// Extra JSX children against a fixed-tuple children prop report the tuple
// arity error, not a per-child error against the phantom `nil` an
// out-of-range tuple element resolves to (the elaboration guard lives in
// getBestMatchIndexedAccessTypeOrUndefined, shared with table entries).

interface JsxElement {}
interface JsxElementChildrenAttribute { children: {}; }
interface JsxIntrinsicElements { div: { children?: any }; }
declare React: { createElement: any, Fragment: any };

declare Comp: (props: { children: [JsxElement, JsxElement] }) => JsxElement;

local ok = <Comp><div/><div/></Comp>;
local bad = <Comp><div/><div/><div/></Comp>;

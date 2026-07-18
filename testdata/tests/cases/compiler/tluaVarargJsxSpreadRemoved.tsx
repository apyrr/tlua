// @jsx: preserve
// @strict: true

interface JsxElement {}
interface JsxElementChildrenAttribute { children: {}; }
interface JsxIntrinsicElements {
    div: { a?: number; b?: string; children?: any };
}

declare props: { a: number; b: string };
declare items: number[];

// JSX is not a reason to keep spread: attribute spread is gone.
local spreadAttribute = <div {...props} />;

// Including when mixed with regular attributes.
local spreadAndAttribute = <div a={1} {...props} />;

// Children spread is gone too.
local spreadChildren = <div>{...items}</div>;

// Regular attributes and children still fold into the attributes type.
local plain = <div a={1} b="x" />;
local withChildren = <div a={1}>{items}</div>;

//// [tests/cases/compiler/tluaWordOperatorJsxNames.tsx] ////

//// [tluaWordOperatorJsxNames.tsx]
// A JSX name is an IdentifierName position, so the word-spelled operators are
// ordinary names there. Unlike the member name `t.and`, which must be written
// `t["and"]`, a JSX tag or attribute name has no alternative spelling.

interface JsxElement {}
interface JsxIntrinsicElements {
    // Member names, so the word operators are quoted here.
    div: { "and"?: boolean; "or"?: boolean; "not"?: boolean; "not-dashed"?: number };
    "or": { "and"?: boolean };
}

declare p: boolean;
declare q: boolean;

// Attribute names.
local attrs = <div and={p} or={q} not={p} />;

// A dash resumes the name after the word, so `not-dashed` is one attribute.
local dashed = <div not-dashed={1} />;

// Tag name, including the closing tag.
local tag = <or and={p}></or>;

// The operators still operate inside an expression container.
local operators = <div and={p and q} or={p or q} not={not p} />;

// A word operator after a dot is a member name, so it stays an error, exactly
// as `t.and` does.
declare Ns: { "and": () => JsxElement };
local dotted = <Ns.and />;


//// [tluaWordOperatorJsxNames.jsx]
-- A JSX name is an IdentifierName position, so the word-spelled operators are
-- ordinary names there. Unlike the member name `t.and`, which must be written
-- `t["and"]`, a JSX tag or attribute name has no alternative spelling.
-- Attribute names.
local attrs = <div and={p} or={q} not={p}/>;
-- A dash resumes the name after the word, so `not-dashed` is one attribute.
local dashed = <div not-dashed={1}/>;
-- Tag name, including the closing tag.
local tag = <or and={p}></or>;
-- The operators still operate inside an expression container.
local operators = <div and={p and q} or={p or q} not={!p}/>;
local dotted = <Ns. and/>;

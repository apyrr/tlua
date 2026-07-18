//// [tests/cases/compiler/jsxNumericEntityLoneSurrogate.tsx] ////

//// [file.tsx]
	interface JsxElement {}
	interface JsxIntrinsicElements {
		[s: string]: any;
	}
declare React: any;

// Numeric character references for lone surrogates should be preserved, not
// corrupted to U+FFFD. "\uD800" is a lone high surrogate; "\uDC00" a lone low.
local text = <div>&#xD800;a&#xDC00;</div>;
local attr = <div title="&#xD800;"></div>;

// A non-BMP code point reference is a single supplementary character.
local supplementary = <div>&#x1F600;</div>;


//// [file.lua]
-- Numeric character references for lone surrogates should be preserved, not
-- corrupted to U+FFFD. "\uD800" is a lone high surrogate; "\uDC00" a lone low.
local text = React.createElement("div", nil, "\xed\xa0\x80a\xed\xb0\x80");
local attr = React.createElement("div", { title: "\xed\xa0\x80" });
-- A non-BMP code point reference is a single supplementary character.
local supplementary = React.createElement("div", nil, "😀");

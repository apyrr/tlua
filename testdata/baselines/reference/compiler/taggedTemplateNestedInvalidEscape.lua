//// [tests/cases/compiler/taggedTemplateNestedInvalidEscape.tlua] ////

//// [taggedTemplateNestedInvalidEscape.tlua]
declare function tag(template: TemplateStringsArray, ...substitutions: any[]): string;

// The outer tagged template should NOT be rewritten with __makeTemplateObject
// because only the inner tagged template contains an invalid escape.
tag`ok ${tag`\u`}`;

// The inner tagged template here has an invalid escape, but the outer one doesn't.
tag`fine ${tag`\x`} also fine`;

// This one should be rewritten because it has its own invalid escape.
tag`\u`;

// Nested with substitution, only inner is invalid.
tag`hello ${tag`\u{}`} world`;


//// [taggedTemplateNestedInvalidEscape.lua]
var __makeTemplateObject = (this && this.__makeTemplateObject) || function (cooked, raw) {
    if (Object.defineProperty) { Object.defineProperty(cooked, "raw", { value: raw }); } else { cooked.raw = raw; }
    return cooked;
};
-- The outer tagged template should NOT be rewritten with __makeTemplateObject
-- because only the inner tagged template contains an invalid escape.
tag ("ok " .. tostring(tag(templateObject_1 or (templateObject_1 = __makeTemplateObject([void 0], ["\\u"])))));
-- The inner tagged template here has an invalid escape, but the outer one doesn't.
tag ("fine " .. tostring(tag(templateObject_2 or (templateObject_2 = __makeTemplateObject([void 0], ["\\x"])))) .. " also fine");
-- This one should be rewritten because it has its own invalid escape.
tag(templateObject_3 or (templateObject_3 = __makeTemplateObject([void 0], ["\\u"])));
-- Nested with substitution, only inner is invalid.
tag ("hello " .. tostring(tag(templateObject_4 or (templateObject_4 = __makeTemplateObject([void 0], ["\\u{}"])))) .. " world");
local templateObject_1, templateObject_2, templateObject_3, templateObject_4;

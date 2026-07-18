//// [tests/cases/compiler/tluaStringEscapes.tlua] ////

//// [tluaStringEscapes.tlua]
// String literals emit as valid Lua: JS-only escapes (\u{..}, \uXXXX) never appear —
// control bytes use \xHH, and the author's quote style is preserved.

// \u{1b} (ESC) and \x41 (A) cook, then re-emit in Lua form: "\x1bA".
local esc = "\u{1b}\x41"

// A non-ASCII separator (U+2028) is valid UTF-8 and stays raw.
local sep = "a\u{2028}b"

// Single-quoted strings keep their quote; the embedded quote is escaped.
local sq = 'it\'s a "test"'

// Readable whitespace escapes are preserved.
local ws = "tab\tnl\ncr\r"

// A NUL byte becomes \x00 (never a bare \0 that could swallow a following digit).
local nul = "\u{0}0"


//// [tluaStringEscapes.lua]
-- String literals emit as valid Lua: JS-only escapes (\u{..}, \uXXXX) never appear —
-- control bytes use \xHH, and the author's quote style is preserved.
-- \u{1b} (ESC) and \x41 (A) cook, then re-emit in Lua form: "\x1bA".
local esc = "\x1bA";
-- A non-ASCII separator (U+2028) is valid UTF-8 and stays raw.
local sep = "a b";
-- Single-quoted strings keep their quote; the embedded quote is escaped.
local sq = 'it\'s a "test"';
-- Readable whitespace escapes are preserved.
local ws = "tab\tnl\ncr\r";
-- A NUL byte becomes \x00 (never a bare \0 that could swallow a following digit).
local nul = "\x000";

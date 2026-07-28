//// [tests/cases/conformance/ported/scannerStringLiterals.tlua] ////

//// [scannerStringLiterals.tlua]
-- ported from tests/cases/conformance/scanner/ecmascript5/scannerStringLiterals.ts
-- dropped: JavaScript-only Unicode escape and backslash line-continuation cases; Lua string escaping and continuation rules differ
-- dropped: @target: es2015 directive (tlua defaults to latest target)

local _ = ""
local _ = " "
local _ = "a"
local _ = "'"
local _ = '"'
local _ = "\""

local _ = "Should error because of newline.
local _ = "Should error because of end of file.


//// [scannerStringLiterals.lua]
-- ported from tests/cases/conformance/scanner/ecmascript5/scannerStringLiterals.ts
-- dropped: JavaScript-only Unicode escape and backslash line-continuation cases; Lua string escaping and continuation rules differ
-- dropped: @target: es2015 directive (tlua defaults to latest target)
local _ = "";
local _ = " ";
local _ = "a";
local _ = "'";
local _ = '"';
local _ = "\"";
local _ = "Should error because of newline.";
local _ = "Should error because of end of file.";

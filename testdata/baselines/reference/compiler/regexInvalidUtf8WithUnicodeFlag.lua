//// [tests/cases/compiler/regexInvalidUtf8WithUnicodeFlag.tlua] ////

//// [regexInvalidUtf8WithUnicodeFlag.tlua]
local _ = /€/u


//// [regexInvalidUtf8WithUnicodeFlag.lua]
local _ = /€/u;

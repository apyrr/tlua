//// [tests/cases/compiler/unicodeLineSeparatorInMultilineComment.tlua] ////

//// [unicodeLineSeparatorInMultilineComment.tlua]
/* a b */ local x = 1;


//// [unicodeLineSeparatorInMultilineComment.lua]
--[[ a
b ]] local x = 1;

//// [tests/cases/compiler/tluaLocalTypeDisambiguation.tlua] ////

//// [vars.tlua]
// `local type` is only a type declaration when an identifier follows on the same
// line. `type` is a contextual keyword, so it stays a plain variable name here.
// (`interface` is a reserved word and can never be a variable, so there is no
// analogous `local interface = ...` ambiguity to resolve.)
local type = 5;
local n: number = type;

// A type annotation after the name keeps it a variable, not a type alias.
local type2: number = 7;

// Names that merely start with `type`/`interface` are unaffected.
local typeName = 1;
local interfaces = 2;

//// [sameline.tlua]
// The identifier must be on the SAME line: with a line break the keyword is a
// variable name and the next line is a separate statement.
local type
Name = 3;

//// [generics.tlua]
// Generic `local type`/`local interface` carry their type parameters through.
local type Box<T> = { value: T };
local interface Pair<A, B> { first: A, second: B };
local b: Box<number> = { value = 1 };
local p: Pair<number, string> = { first = 1, second = "x" };


//// [vars.lua]
-- `local type` is only a type declaration when an identifier follows on the same
-- line. `type` is a contextual keyword, so it stays a plain variable name here.
-- (`interface` is a reserved word and can never be a variable, so there is no
-- analogous `local interface = ...` ambiguity to resolve.)
local type = 5;
local n = type;
-- A type annotation after the name keeps it a variable, not a type alias.
local type2 = 7;
-- Names that merely start with `type`/`interface` are unaffected.
local typeName = 1;
local interfaces = 2;
//// [sameline.lua]
-- The identifier must be on the SAME line: with a line break the keyword is a
-- variable name and the next line is a separate statement.
local type;
Name = 3;
//// [generics.lua]
;
local b = { value = 1 };
local p = { first = 1, second = "x" };

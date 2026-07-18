//// [tests/cases/compiler/tluaLocalDeclarations.tlua] ////

//// [tluaLocalDeclarations.tlua]
local plain;
plain = 1;

local typed: number;
typed = plain;

local initialized = typed;
local typedInitialized: number = initialized;
typedInitialized = typedInitialized + 1;

local literal = "a";
literal = "b";


//// [tluaLocalDeclarations.lua]
local plain;
plain = 1;
local typed;
typed = plain;
local initialized = typed;
local typedInitialized = initialized;
typedInitialized = typedInitialized + 1;
local literal = "a";
literal = "b";

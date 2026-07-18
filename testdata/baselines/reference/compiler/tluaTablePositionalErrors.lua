//// [tests/cases/compiler/tluaTablePositionalErrors.tlua] ////

//// [tluaTablePositionalErrors.tlua]
// A computed key without `=`/`:` is a natural error.
declare k: string;
local a = { [k] };

// A keyed form missing its value.
local b = { x = };

// A bare separator.
local c = { , };


//// [tluaTablePositionalErrors.lua]
local a = { [k] =  };
-- A keyed form missing its value.
local b = { x =  };
-- A bare separator.
local c = {};

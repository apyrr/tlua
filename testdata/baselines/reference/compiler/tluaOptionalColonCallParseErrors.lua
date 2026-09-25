//// [tests/cases/compiler/tluaOptionalColonCallParseErrors.tlua] ////

//// [tluaOptionalColonCallParseErrors.tlua]
// A space between `?` and `:` does not spell the optional colon call.

interface Box {
    read(self: Box): number;
}

declare function maybeBox(): Box | nil;

local b = maybeBox()
local r = b? :read()


//// [tluaOptionalColonCallParseErrors.lua]
-- A space between `?` and `:` does not spell the optional colon call.
local b = maybeBox();
local r = b;
read();

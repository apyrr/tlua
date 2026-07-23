//// [tests/cases/compiler/tluaColonCallParseErrors.tlua] ////

//// [tluaColonCallParseErrors.tlua]
local M = { ping = function(self: {}) return 1; end };

// A colon access exists only as a call: Lua rejects `obj:f` without arguments.
local f = M:ping;

// The colon needs a member name after it.
local g = M:(1);

// Lua reserved words cannot name a colon call: committing on `M:end(` would
// swallow the enclosing block terminator.
local h = M:end();

// An instantiation expression cannot be followed by a colon call, matching the
// dot form's diagnostic.
declare id: <T>(self: {}, x: T) => T;
local i = id<number>:ping();

// Error recovery: an annotation-shaped `ident : ident` in invalid code must
// not swallow the declarations that follow it into a bogus argument list.
local bad = (b: number);
declare fine: (n: number) => void;
fine(1);


//// [tluaColonCallParseErrors.lua]
local M = { ping = function(self)
        return 1;
    end };
-- A colon access exists only as a call: Lua rejects `obj:f` without arguments.
local f = M;
ping;
-- The colon needs a member name after it.
local g = M;
(1);
-- Lua reserved words cannot name a colon call: committing on `M:end(` would
-- swallow the enclosing block terminator.
local h = M;
();
local i = id:ping();
-- Error recovery: an annotation-shaped `ident : ident` in invalid code must
-- not swallow the declarations that follow it into a bogus argument list.
local bad = (b);
number;
;
fine(1);

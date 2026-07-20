//// [tests/cases/compiler/tluaNotCallableErrors.tlua] ////

//// [tluaNotCallableErrors.tlua]
// Calling a value with no call signatures must error. TypeScript's
// "untyped call" escape hatch keyed on assignability to the global Function
// interface — an empty stub in this fork (since deleted outright), so every
// value was Function-assignable and not-callable errors vanished.
local notCallable = 1;
notCallable();

local emptyObject = {};
emptyObject();


declare untypedValue: any;
untypedValue(); // any stays callable


//// [tluaNotCallableErrors.lua]
-- Calling a value with no call signatures must error. TypeScript's
-- "untyped call" escape hatch keyed on assignability to the global Function
-- interface — an empty stub in this fork (since deleted outright), so every
-- value was Function-assignable and not-callable errors vanished.
local notCallable = 1;
notCallable();
local emptyObject = {};
emptyObject();
untypedValue(); -- any stays callable

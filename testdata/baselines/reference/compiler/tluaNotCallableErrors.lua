//// [tests/cases/compiler/tluaNotCallableErrors.tlua] ////

//// [tluaNotCallableErrors.tlua]
// Calling a value with no call signatures must error. TypeScript's
// "untyped call" escape hatch keyed on Function-assignability, which is
// meaningless against tlua's empty Function interface — every value was
// Function-assignable and not-callable errors vanished.
local notCallable = 1;
notCallable();

local emptyObject = {};
emptyObject();

declare legacyFunction: Function;
legacyFunction();

declare untypedValue: any;
untypedValue(); // any stays callable


//// [tluaNotCallableErrors.lua]
-- Calling a value with no call signatures must error. TypeScript's
-- "untyped call" escape hatch keyed on Function-assignability, which is
-- meaningless against tlua's empty Function interface — every value was
-- Function-assignable and not-callable errors vanished.
local notCallable = 1;
notCallable();
local emptyObject = {};
emptyObject();
legacyFunction();
untypedValue(); -- any stays callable

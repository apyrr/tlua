//// [tests/cases/compiler/tluaAndOrNotReserved.tlua] ////

//// [tluaAndOrNotReserved.tlua]
// `and`, `or` and `not` are keywords, so they are reserved. Every use as a name
// falls into stock recovery, as with the class, decorator, array and this
// removals. Lua reserves them too, so `t.and` is invalid there as well.

local obj = { x = 1 };

// Member names.
local a1 = obj.and;
local a2 = obj.or;
local a3 = obj.not;

// Property names in a table literal.
local o1 = { and = 1 };
local o2 = { or = 1 };
local o3 = { not = 1 };

// Binding names.
local and = 1;
local or = 2;
local not = 3;

// Declaration names.
function or() {}

// `not` shares KindExclamationToken with `!`, but only the punctuation spelling
// is the non-null and definite-assignment operator. Accepting the word in those
// roles would silently suppress a nil check.
declare maybe: number | nil;
local n1 = maybe not;
local n2 not: number;
// The punctuation spelling keeps both roles.
local n3 = maybe!;
local n4!: number;
// And the word keeps its own: prefix negation.
local n5 = not maybe;

// The type-position `!` (the JSDoc non-nullable operator) is punctuation-only
// as well: the word `not` never silently becomes a type operator.
local n6: number not;
type N7 = not number;


//// [tluaAndOrNotReserved.lua]
-- `and`, `or` and `not` are keywords, so they are reserved. Every use as a name
-- falls into stock recovery, as with the class, decorator, array and this
-- removals. Lua reserves them too, so `t.and` is invalid there as well.
local obj = { x = 1 };
-- Member names.
local a1 = obj. and ;
local a2 = obj. or ;
local a3 = obj.;
!;
-- Property names in a table literal.
local o1 = {  and , 1 };
local o2 = {  or , 1 };
local o3 = { !, 1 };
-- Binding names.
local ;
 and ;
1;
local ;
 or ;
2;
local ;
!;
3;
-- Declaration names.
function () { }
 or ();
{ }
local n1 = maybe;
!;
local n2;
!;
number;
-- The punctuation spelling keeps both roles.
local n3 = maybe;
local n4;
!;
number;
-- And the word keeps its own: prefix negation.
local n5 = !maybe;
-- The type-position `!` (the JSDoc non-nullable operator) is punctuation-only
-- as well: the word `not` never silently becomes a type operator.
local n6;
!;
!number;

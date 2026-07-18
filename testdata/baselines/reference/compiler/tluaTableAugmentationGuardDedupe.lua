//// [tests/cases/compiler/tluaTableAugmentationGuardDedupe.tlua] ////

//// [a.tlua]
Shared = Shared or {};
Shared.fromA = 1;

Nested = Nested or {};
Nested.sub = Nested.sub or {};

//// [b.tlua]
Shared = Shared or {};
Shared.fromB = "b";
Shared.self = Shared;

Nested = Nested or {};
Nested.sub = Nested.sub or {};
Nested.sub.value = true;

// Independently declared constructor members stay as distinct arms even when
// their text is identical.
NonEmpty = NonEmpty or { seed = 1 };
NonEmpty = NonEmpty or { seed = 1 };

//// [c.tlua]
local shared = Shared;
local fromA: number = Shared.fromA;
local fromB: string = Shared.fromB;
local self = Shared.self;

local nested = Nested;
local value: boolean = Nested.sub.value;

local nonEmpty = NonEmpty;


//// [a.lua]
Shared = Shared or {};
Shared.fromA = 1;
Nested = Nested or {};
Nested.sub = Nested.sub or {};
//// [b.lua]
Shared = Shared or {};
Shared.fromB = "b";
Shared.self = Shared;
Nested = Nested or {};
Nested.sub = Nested.sub or {};
Nested.sub.value = true;
-- Independently declared constructor members stay as distinct arms even when
-- their text is identical.
NonEmpty = NonEmpty or { seed = 1 };
NonEmpty = NonEmpty or { seed = 1 };
//// [c.lua]
local shared = Shared;
local fromA = Shared.fromA;
local fromB = Shared.fromB;
local self = Shared.self;
local nested = Nested;
local value = Nested.sub.value;
local nonEmpty = NonEmpty;

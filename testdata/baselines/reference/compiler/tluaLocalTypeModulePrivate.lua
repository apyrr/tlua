//// [tests/cases/compiler/tluaLocalTypeModulePrivate.tlua] ////

//// [a.tlua]
// `local` type/interface declarations are module-private: they live in this
// file's scope but are NOT hoisted into the global type table.
local type Secret = { code: number };
local interface Hidden { tag: string };

// Bare top-level type/interface ARE global (tlua has no type-export syntax),
// so another module can name them.
type Shared = { id: number };
interface Public { name: string };

// The local types resolve fine WITHIN their own module.
local s: Secret = { code = 1 };
local h: Hidden = { tag = "x" };

//// [b.tlua]
// Global types from a.tlua are visible here.
local ok1: Shared = { id = 2 };
local ok2: Public = { name = "y" };

// Module-private types from a.tlua are NOT visible here.
local bad1: Secret = { code = 3 };
local bad2: Hidden = { tag = "z" };


//// [a.lua]
;
;
-- The local types resolve fine WITHIN their own module.
local s = { code = 1 };
local h = { tag = "x" };
//// [b.lua]
-- Global types from a.tlua are visible here.
local ok1 = { id = 2 };
local ok2 = { name = "y" };
-- Module-private types from a.tlua are NOT visible here.
local bad1 = { code = 3 };
local bad2 = { tag = "z" };

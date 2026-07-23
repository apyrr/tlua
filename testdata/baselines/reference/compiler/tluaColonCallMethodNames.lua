//// [tests/cases/compiler/tluaColonCallMethodNames.tlua] ////

//// [tluaColonCallMethodNames.tlua]
// TS-only keywords are valid Lua method names (Lua's Name grammar reserves
// only its own words), so the canonical constructor idiom parses, checks, and
// emits — symmetric with what the dot path accepts.
local Account = { balance = 0 };

function Account:new(): typeof Account
  return self;
end

function Account:delete(): void
  self.balance = 0;
end

function Account:type(): string
  return "account";
end

local a = Account:new();
Account:delete();
local kind = Account:type();

// The explicit dot form stays interchangeable.
local b = Account.new(Account);

// `switch` and `case` are ordinary identifiers since the switch statement was
// deleted, so they name methods too.
function Account:switch(): void
end

function Account:case(): void
end

Account:switch();
Account:case();


//// [tluaColonCallMethodNames.lua]
-- TS-only keywords are valid Lua method names (Lua's Name grammar reserves
-- only its own words), so the canonical constructor idiom parses, checks, and
-- emits — symmetric with what the dot path accepts.
local Account = { balance = 0 };
function Account:new()
  return self;
end
function Account:delete()
  self.balance = 0;
end
function Account:type()
  return "account";
end
local a = Account:new();
Account:delete();
local kind = Account:type();
-- The explicit dot form stays interchangeable.
local b = Account.new(Account);
-- `switch` and `case` are ordinary identifiers since the switch statement was
-- deleted, so they name methods too.
function Account:switch()
end
function Account:case()
end
Account:switch();
Account:case();

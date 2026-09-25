//// [tests/cases/compiler/tluaLuaContractMembers.tlua] ////

//// [tluaLuaContractMembers.tlua]
// `function T:m ... end` is `T.m = function(self, ...) ... end`: on a table
// whose constructor states a contract, both spellings take their contextual
// type from the contract member and are checked against it.

interface Greeter {
    greet(self: Greeter, name: string): string;
    config: { debug: boolean };
    Think: (self: Greeter) => nil;
    Gen<T>(self: Greeter, x: T): T;
    Stop(self: Greeter): nil;
}

local g = {} as unknown as Greeter
g.Count = 1

-- A body that breaks the contract is an error in either spelling.
function g.greet(self: Greeter, name: string): number
    return 42
end
g.greet = function(self: Greeter, name: string): number
    return 42
end

-- Unannotated parameters take the contract's types in either spelling.
function g:greet(name)
    local n: number = name
    return name
end
g.greet = function(self, name)
    local n: number = name
    return name
end

-- A colon body's implicit self carries the table's other members; it is
-- checked with the contract's receiver, even for a member declared as a
-- function-typed property.
function g:Think()
    local count: number = self.Count
end

-- A generic contract member binds its type parameters in the body.
function g:Gen(x)
    local s: string = x
    return x
end
local generic: string = g:Gen("a")

-- Several bodies for a contract member are stores into a declared slot.
function g:Stop() end
function g:Stop() end

-- Members below a contract member keep the contract's type.
g.config = { debug = 1 }
g.config.verbose = true
local verbose: string = g.config.verbose

-- A member initialized from a sibling member is not circular.
local counters = {} as { Name: string }
counters.Name = "c"
counters.Count = 1
counters.Double = counters.Count * 2
local double: string = counters.Double

-- A table asserted before its methods are attached converts as built: the
-- attached members complete it, and stand as the contract's own members.
interface Tx {
    undo: (() => nil)[];
    open: boolean;
    commit: (self: Tx) => nil;
}
local tx = { undo = {}, open = true } as Tx
function tx:commit()
    self.open = false
end

-- A member the contract does not declare is the table's own, so two bodies
-- for it are duplicates.
function g:Extra() end
function g:Extra() end


//// [tluaLuaContractMembers.lua]
-- `function T:m ... end` is `T.m = function(self, ...) ... end`: on a table
-- whose constructor states a contract, both spellings take their contextual
-- type from the contract member and are checked against it.
local g = {};
g.Count = 1;
-- A body that breaks the contract is an error in either spelling.
function g.greet(self, name)
  return 42;
end
g.greet = function(self, name)
  return 42;
end;
-- Unannotated parameters take the contract's types in either spelling.
function g:greet(name)
  local n = name;
  return name;
end
g.greet = function(self, name)
  local n = name;
  return name;
end;
-- A colon body's implicit self carries the table's other members; it is
-- checked with the contract's receiver, even for a member declared as a
-- function-typed property.
function g:Think()
  local count = self.Count;
end
-- A generic contract member binds its type parameters in the body.
function g:Gen(x)
  local s = x;
  return x;
end
local generic = g:Gen("a");
-- Several bodies for a contract member are stores into a declared slot.
function g:Stop()
end
function g:Stop()
end
-- Members below a contract member keep the contract's type.
g.config = { debug = 1 };
g.config.verbose = true;
local verbose = g.config.verbose;
-- A member initialized from a sibling member is not circular.
local counters = {};
counters.Name = "c";
counters.Count = 1;
counters.Double = counters.Count * 2;
local double = counters.Double;
local tx = { undo = {}, open = true };
function tx:commit()
  self.open = false;
end
-- A member the contract does not declare is the table's own, so two bodies
-- for it are duplicates.
function g:Extra()
end
function g:Extra()
end

//// [tests/cases/compiler/tluaLuaEnvironments.tlua] ////

//// [host.d.tlua]
interface Plugin {
    Name: string;
    Start(self: Plugin, reason: string): nil;
    Think: (self: Plugin) => nil;
    Primary: { Ammo: string };
    AmmoName: string;
}
interface Tool {
    Mode: number;
}
declare SERVER: boolean

// Every file under one matched folder shares that folder's table.
//// [shared.tlua]
PLUGIN.Name = "Door"
PLUGIN.Locked = false
function PLUGIN:Open(): number
    return 1
end

//// [server.tlua]
function PLUGIN:Close()
    local opened: number = self:Open()
    local locked: boolean = self.Locked
end

// Per-realm bodies implement the seed's member, so they are not duplicates.
// Parameters take the seed's types, and bodies are checked against it.
if SERVER then
    function PLUGIN:Start(reason)
        local r: string = reason
    end
else
    function PLUGIN.Start(self, reason)
        local wrong: number = reason
    end
end

// A hook declared as a function-typed property takes a colon body whose self
// carries the folder's other members.
function PLUGIN:Think()
    local locked: boolean = self.Locked
end

//// [client.tlua]
// Another file of the folder implements the same seed member; that is not a
// duplicate either.
function PLUGIN:Think() end
PLUGIN.Start = function(self, reason)
    local r: string = reason
end

// A seed member initialized from a sibling member is not circular.
PLUGIN.AmmoName = "pistol"
PLUGIN.Primary = { Ammo = PLUGIN.AmmoName }

// setmetatable pairs with the folder's table.
local Base = {}
function Base:Knock(): number
    return 1
end
setmetatable(PLUGIN, { __index = Base })
local knock: string = PLUGIN:Knock()

// A single file is a group of its own. Seed members keep the seed's type, and
// members of another group are not visible.
//// [lamp.tlua]
PLUGIN.Name = 42
PLUGIN.Count = 1
PLUGIN.Double = PLUGIN.Count * 2
local double: string = PLUGIN.Double
function PLUGIN:Toggle()
    local count: string = self.Count
    local again: number = self.Toggle
    self:Missing()
end
PLUGIN:Open()
function PLUGIN:Start(reason: number) end

// Two bodies for a member the seed does not declare are duplicates.
function PLUGIN:Extra() end
function PLUGIN:Extra() end

// The deepest matching root wins: this folder is a tool, not a plugin.
//// [paint.tlua]
TOOL.Mode = 1
TOOL.Color = "red"
local p = PLUGIN

// Outside every group the names are only the seed types.
//// [main.tlua]
local tool = TOOL
local function describe(plugin: Plugin): string
    return plugin.Name
end


//// [shared.lua]
PLUGIN.Name = "Door";
PLUGIN.Locked = false;
function PLUGIN:Open()
  return 1;
end
//// [server.lua]
function PLUGIN:Close()
  local opened = self:Open();
  local locked = self.Locked;
end
-- Per-realm bodies implement the seed's member, so they are not duplicates.
-- Parameters take the seed's types, and bodies are checked against it.
if SERVER then
  function PLUGIN:Start(reason)
    local r = reason;
  end
else
  function PLUGIN.Start(self, reason)
    local wrong = reason;
  end
end
-- A hook declared as a function-typed property takes a colon body whose self
-- carries the folder's other members.
function PLUGIN:Think()
  local locked = self.Locked;
end
//// [client.lua]
-- Another file of the folder implements the same seed member; that is not a
-- duplicate either.
function PLUGIN:Think()
end
PLUGIN.Start = function(self, reason)
  local r = reason;
end;
-- A seed member initialized from a sibling member is not circular.
PLUGIN.AmmoName = "pistol";
PLUGIN.Primary = { Ammo = PLUGIN.AmmoName };
-- setmetatable pairs with the folder's table.
local Base = {};
function Base:Knock()
  return 1;
end
setmetatable(PLUGIN, { __index = Base });
local knock = PLUGIN:Knock();
-- A single file is a group of its own. Seed members keep the seed's type, and
-- members of another group are not visible.
//// [lamp.lua]
PLUGIN.Name = 42;
PLUGIN.Count = 1;
PLUGIN.Double = PLUGIN.Count * 2;
local double = PLUGIN.Double;
function PLUGIN:Toggle()
  local count = self.Count;
  local again = self.Toggle;
  self:Missing();
end
PLUGIN:Open();
function PLUGIN:Start(reason)
end
-- Two bodies for a member the seed does not declare are duplicates.
function PLUGIN:Extra()
end
function PLUGIN:Extra()
end
-- The deepest matching root wins: this folder is a tool, not a plugin.
//// [paint.lua]
TOOL.Mode = 1;
TOOL.Color = "red";
local p = PLUGIN;
-- Outside every group the names are only the seed types.
//// [main.lua]
local tool = TOOL;
local function describe(plugin)
  return plugin.Name;
end

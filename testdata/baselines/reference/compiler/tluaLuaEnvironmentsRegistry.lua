//// [tests/cases/compiler/tluaLuaEnvironmentsRegistry.tlua] ////

//// [host.d.tlua]
interface Entity {
    GetPos(self: Entity): number;
}
interface ENT extends Entity {
    PrintName: string;
}
interface EFFECT {}
interface Effects {}
interface ScriptedEntities {
    // An explicit declaration keeps its own type.
    base_anim: ENT;
}
// One signature, so the lookup does not depend on overload order across
// merged declarations.
interface EntsLib {
    FindByClass<K extends string>(name: K): (K extends keyof ScriptedEntities ? ScriptedEntities[K] : Entity)[];
}
declare ents: EntsLib

// A folder and a single file each register their name.
//// [shared.tlua]
ENT.PrintName = "Door"
function ENT:Open(): number
    return 1
end

//// [lamp.tlua]
ENT.Brightness = 5

// Two addons shipping one class name register the union of both tables.
//// [lamp.tlua]
ENT.Color = "red"

//// [spark.tlua]
EFFECT.Size = 1

//// [main.tlua]
for _, door in ipairs(ents.FindByClass("door")) do
    local opened: number = door:Open()
    local name: string = door.PrintName
end
for _, lamp in ipairs(ents.FindByClass("lamp")) do
    local brightness = lamp.Brightness
end
for _, base in ipairs(ents.FindByClass("base_anim")) do
    base:Open()
end
for _, prop in ipairs(ents.FindByClass("prop_*")) do
    local pos: number = prop:GetPos()
    prop:Open()
end
local className: string = "door"
local anyEntity = ents.FindByClass(className)
local DOOR_CLASS = "door" as const
for _, door in ipairs(ents.FindByClass(DOOR_CLASS)) do
    door:Open()
end
local known: ScriptedEntities["door"] = ents.FindByClass("door")[1] as ScriptedEntities["door"]


//// [shared.lua]
ENT.PrintName = "Door";
function ENT:Open()
  return 1;
end
//// [lamp.lua]
ENT.Brightness = 5;
-- Two addons shipping one class name register the union of both tables.
//// [lamp.lua]
ENT.Color = "red";
//// [spark.lua]
EFFECT.Size = 1;
//// [main.lua]
for _, door in ipairs(ents.FindByClass("door")) do
  local opened = door:Open();
  local name = door.PrintName;
end
for _, lamp in ipairs(ents.FindByClass("lamp")) do
  local brightness = lamp.Brightness;
end
for _, base in ipairs(ents.FindByClass("base_anim")) do
  base:Open();
end
for _, prop in ipairs(ents.FindByClass("prop_*")) do
  local pos = prop:GetPos();
  prop:Open();
end
local className = "door";
local anyEntity = ents.FindByClass(className);
local DOOR_CLASS = "door";
for _, door in ipairs(ents.FindByClass(DOOR_CLASS)) do
  door:Open();
end
local known = ents.FindByClass("door")[1];

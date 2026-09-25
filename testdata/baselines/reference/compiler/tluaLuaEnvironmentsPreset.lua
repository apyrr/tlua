//// [tests/cases/compiler/tluaLuaEnvironmentsPreset.tlua] ////

//// [tluaconfig.json]
{
    "compilerOptions": {
        "luaEnvironments": [
            { "root": "**/lua/entities/*", "globals": { "ENT": "Entity" } }
        ]
    }
}
//// [host.d.tlua]
interface Entity {
    GetPos(self: Entity): number;
}
//// [crate.tlua]
function ENT:Weight(): number
    return self:GetPos()
end
local w: string = ENT:Weight()


//// [crate.lua]
function ENT:Weight()
  return self:GetPos();
end
local w = ENT:Weight();

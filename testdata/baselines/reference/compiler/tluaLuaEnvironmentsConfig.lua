//// [tests/cases/compiler/tluaLuaEnvironmentsConfig.tlua] ////

//// [host.d.tlua]
interface Seed {
    Name: string;
}
//// [a.tlua]
OK.Name = "one"
OK.Extra = 1
//// [a.tlua]
M.Name = "m"


//// [a.lua]
OK.Name = "one";
OK.Extra = 1;
//// [a.lua]
M.Name = "m";

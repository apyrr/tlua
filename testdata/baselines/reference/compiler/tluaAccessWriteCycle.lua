//// [tests/cases/compiler/tluaAccessWriteCycle.tlua] ////

//// [a.tlua]
AccessCycle = {};
AccessCycle.a = {};
AccessCycle.b = {};
AccessCycle.a = AccessCycle.b;

//// [b.tlua]
AccessCycle.b = AccessCycle.a;
AccessCycle.a.value = 1;
local accessCycleValue = AccessCycle.a.value;


//// [a.lua]
AccessCycle = {};
AccessCycle.a = {};
AccessCycle.b = {};
AccessCycle.a = AccessCycle.b;
//// [b.lua]
AccessCycle.b = AccessCycle.a;
AccessCycle.a.value = 1;
local accessCycleValue = AccessCycle.a.value;

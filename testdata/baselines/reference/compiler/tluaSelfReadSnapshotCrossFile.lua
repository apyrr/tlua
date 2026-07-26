//// [tests/cases/compiler/tluaSelfReadSnapshotCrossFile.tlua] ////

//// [def.tlua]
counted = 0;
local function setLate()
  late = 0;
end
local function setLenient()
  lenient = 0;
end

//// [use.tlua]
counted = counted + 1;
late = late + 1;
local function bump()
  lenient = lenient + 1;
end


//// [def.lua]
counted = 0;
local function setLate()
  late = 0;
end
local function setLenient()
  lenient = 0;
end
//// [use.lua]
counted = counted + 1;
late = late + 1;
local function bump()
  lenient = lenient + 1;
end

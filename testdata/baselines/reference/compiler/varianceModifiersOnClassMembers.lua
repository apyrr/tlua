//// [tests/cases/compiler/varianceModifiersOnClassMembers.tlua] ////

//// [varianceModifiersOnClassMembers.tlua]
// https://github.com/microsoft/typescript-go/issues/4123

class C {
  in x = 1;
  out y = 2;
}

local isIn = "x" in { x = 1 };
for k = 1, 1 do
  console.log(k);
end


//// [varianceModifiersOnClassMembers.lua]
-- https://github.com/microsoft/typescript-go/issues/4123
class;
C;
do
    in x;
  1;
  out;
  y = 2;
end
local isIn = "x" in { x = 1 };
for k = 1, 1 do
  console.log(k);
end

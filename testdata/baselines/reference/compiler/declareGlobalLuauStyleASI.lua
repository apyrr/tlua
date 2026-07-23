//// [tests/cases/compiler/declareGlobalLuauStyleASI.tlua] ////

//// [declareGlobalLuauStyleASI.tlua]
// `declare` remains a valid identifier: without a same-line `name:` it is an
// expression, and a line break after `declare` splits it into two statements.
local declare = 1;
local x = 2;
declare
x;
declare = x;
local call = function(declare: number) return declare end;
call(declare);


//// [declareGlobalLuauStyleASI.lua]
-- `declare` remains a valid identifier: without a same-line `name:` it is an
-- expression, and a line break after `declare` splits it into two statements.
local declare = 1;
local x = 2;
declare;
x;
declare = x;
local call = function(declare)
  return declare;
end;
call(declare);

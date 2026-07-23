//// [tests/cases/compiler/classFieldsAssignmentNamedEvaluation.tlua] ////

//// [classFieldsAssignmentNamedEvaluation.tlua]
local x: any;
x = class { static #foo = 1; };


//// [classFieldsAssignmentNamedEvaluation.lua]
local x;
x = class;
do
  static;
  #foo;
  1;
end
;

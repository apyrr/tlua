//// [tests/cases/compiler/duplicatePropertyAndAccessor.tlua] ////

//// [duplicatePropertyAndAccessor.tlua]
// https://github.com/microsoft/typescript-go/issues/4130

class C {
  y: number = 2;
  accessor y: number = 3;
}


//// [duplicatePropertyAndAccessor.lua]
-- https://github.com/microsoft/typescript-go/issues/4130
class;
C;
do
    y;
    number = 2;
    y;
    number = 3;
end

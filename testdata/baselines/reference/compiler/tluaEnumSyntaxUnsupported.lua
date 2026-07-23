//// [tests/cases/compiler/tluaEnumSyntaxUnsupported.tlua] ////

//// [tluaEnumSyntaxUnsupported.tlua]
enum E {
    A,
}


//// [tluaEnumSyntaxUnsupported.lua]
enum;
E;
do
  A,
  ;
end

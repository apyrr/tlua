//// [tests/cases/compiler/tluaEnumSyntaxUnsupported.tlua] ////

//// [tluaEnumSyntaxUnsupported.tlua]
enum E {
    A,
}


//// [tluaEnumSyntaxUnsupported.lua]
enum;
E;
{
    A,
    ;
}

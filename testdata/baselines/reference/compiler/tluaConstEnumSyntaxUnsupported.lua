//// [tests/cases/compiler/tluaConstEnumSyntaxUnsupported.tlua] ////

//// [tluaConstEnumSyntaxUnsupported.tlua]
local enum E {
    A,
}


//// [tluaConstEnumSyntaxUnsupported.lua]
local enum;
E;
do
    A,
    ;
end

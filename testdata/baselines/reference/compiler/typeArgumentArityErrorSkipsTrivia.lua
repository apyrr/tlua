//// [tests/cases/compiler/typeArgumentArityErrorSkipsTrivia.tlua] ////

//// [typeArgumentArityErrorSkipsTrivia.tlua]
declare function f<T>(a: T): T;

f<   string, number>("a");

f<
    string, number>("a");


//// [typeArgumentArityErrorSkipsTrivia.lua]
f("a");
f("a");

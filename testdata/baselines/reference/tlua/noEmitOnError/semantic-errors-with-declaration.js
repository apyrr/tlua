currentDirectory::/user/username/projects/noEmitOnError
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/noEmitOnError/shared/types/db.tlua] *new* 
interface A {
    name: string;
}
//// [/user/username/projects/noEmitOnError/src/main.tlua] *new* 
local a: string = 10;
//// [/user/username/projects/noEmitOnError/src/other.tlua] *new* 
console.log("hi");
//// [/user/username/projects/noEmitOnError/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./dev-build",
        "declaration": true,
        "incremental": false,
        "noEmitOnError": true,
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.

[7m1[0m local a: string = 10;
[7m [0m [91m      ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/home/src/tslibs/TS/Lib/lib.luajit.d.tlua] *Lib*
/// <reference no-default-lib="true"/>
interface Boolean {}
interface Function {}
interface CallableFunction {}
interface NewableFunction {}
interface IArguments {}
interface Number { toExponential: any; }
interface Object {}
interface RegExp {}
interface String { charAt: any; }
interface Array<T> { length: number; [n: number]: T; }
interface ReadonlyArray<T> {}
interface SymbolConstructor {
    (desc?: string | number): symbol;
    for(name: string): symbol;
    readonly toStringTag: symbol;
}
declare Symbol: SymbolConstructor;
interface Symbol {
    readonly [Symbol.toStringTag]: string;
}
declare console: { log(msg: any): void; };
declare function require(module: string): any;



Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.

[7m1[0m local a: string = 10;
[7m [0m [91m      ~[0m


Found 1 error in src/main.tlua[90m:1[0m




Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a: string = "hello";

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mshared/types/db.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface A {
[7m [0m [91m~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: string = "hello";
[7m [0m [91m~~~~~[0m

[96msrc/other.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m console.log("hi");
[7m [0m [91m~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  shared/types/db.tlua[90m:1[0m
     1  src/main.tlua[90m:1[0m
     1  src/other.tlua[90m:1[0m




Edit [2]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mshared/types/db.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface A {
[7m [0m [91m~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: string = "hello";
[7m [0m [91m~~~~~[0m

[96msrc/other.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m console.log("hi");
[7m [0m [91m~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  shared/types/db.tlua[90m:1[0m
     1  src/main.tlua[90m:1[0m
     1  src/other.tlua[90m:1[0m



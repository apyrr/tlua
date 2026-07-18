currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/a.tlua] *new* 
local a: number = "hello"
//// [/home/src/projects/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": false
    }
}

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

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

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m




Edit [1]:: Fix error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = "hello";

tlua --noEmit
ExitStatus:: Success
Output::



Edit [2]:: no change

tlua --noEmit
ExitStatus:: Success
Output::



Edit [3]:: Emit after fixing error

tlua 
ExitStatus:: Success
Output::
//// [/home/src/projects/project/a.lua] *new* 
local a = "hello";




Edit [4]:: no change

tlua --noEmit
ExitStatus:: Success
Output::



Edit [5]:: Introduce error
//// [/home/src/projects/project/a.tlua] *modified* 
local a: number = "hello"

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m




Edit [6]:: Emit when error

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/a.lua] *rewrite with same content*



Edit [7]:: no change

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m



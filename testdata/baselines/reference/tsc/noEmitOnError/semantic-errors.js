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
//// [/user/username/projects/noEmitOnError/tsconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./dev-build",
        "declaration": false,
        "incremental": false,
        "noEmitOnError": true,
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'number' is not assignable to type 'string'.

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
[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'number' is not assignable to type 'string'.

[7m1[0m local a: string = 10;
[7m [0m [91m      ~[0m


Found 1 error in src/main.tlua[90m:1[0m




Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a: string = "hello";

tlua 
ExitStatus:: Success
Output::
//// [/user/username/projects/noEmitOnError/dev-build/shared/types/db.lua] *new* 

//// [/user/username/projects/noEmitOnError/dev-build/src/main.lua] *new* 
local a = "hello";

//// [/user/username/projects/noEmitOnError/dev-build/src/other.lua] *new* 
console.log("hi");




Edit [2]:: no change

tlua 
ExitStatus:: Success
Output::
//// [/user/username/projects/noEmitOnError/dev-build/shared/types/db.lua] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/src/main.lua] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/src/other.lua] *rewrite with same content*


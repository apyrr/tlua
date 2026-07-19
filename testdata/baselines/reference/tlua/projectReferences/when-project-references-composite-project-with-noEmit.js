currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/project/index.tlua] *new* 
local utils = require("utils.index");
utils.x;
//// [/home/src/workspaces/solution/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "rootDir": "..",
    },
    "references": [
        { "path": "../utils" },
    ],
}
//// [/home/src/workspaces/solution/utils/index.tlua] *new* 
local x = 10;
return { x = x };
//// [/home/src/workspaces/solution/utils/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "noEmit": true
    }
}

tlua --p project
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mproject/tluaconfig.json[0m:[93m6[0m:[93m9[0m - [91merror[0m[90m TLUA6310: [0mReferenced project '/home/src/workspaces/solution/utils' may not disable emit.

[7m6[0m         { "path": "../utils" },
[7m [0m [91m        ~~~~~~~~~~~~~~~~~~~~~~[0m


Found 1 error in project/tluaconfig.json[90m:6[0m

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
//// [/home/src/workspaces/solution/project/index.lua] *new* 
local utils = require("utils.index");
utils.x;



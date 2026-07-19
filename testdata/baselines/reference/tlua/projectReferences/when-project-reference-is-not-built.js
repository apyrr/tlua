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
        "composite": true
    }
}

tlua --p project
ExitStatus:: Success
Output::
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



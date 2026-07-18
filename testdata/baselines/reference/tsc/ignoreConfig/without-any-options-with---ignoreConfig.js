currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/c.tlua] *new* 
local c = 10;
//// [/home/src/workspaces/project/src/a.tlua] *new* 
local a = 10;
//// [/home/src/workspaces/project/src/b.tlua] *new* 
local b = 10;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "include": ["src"],
}

tlua --ignoreConfig
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
//// [/home/src/workspaces/project/src/a.lua] *new* 
local a = 10;

//// [/home/src/workspaces/project/src/b.lua] *new* 
local b = 10;



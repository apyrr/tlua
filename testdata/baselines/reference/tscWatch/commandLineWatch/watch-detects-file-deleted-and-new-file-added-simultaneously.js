currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local b = require("b");
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 1;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{}

tlua --watch
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

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
//// [/home/src/workspaces/project/a.lua] *new* 
local b = require("b");

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 1;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/b.tlua
*refresh*    /home/src/workspaces/project/a.tlua
Signatures::


Edit [0]:: delete b.tlua and create c.tlua with updated import
//// [/home/src/workspaces/project/a.tlua] *modified* 
local c = require("c");
//// [/home/src/workspaces/project/b.tlua] *deleted*
//// [/home/src/workspaces/project/c.tlua] *new* 
local c = 2;


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/home/src/workspaces/project/a.lua] *modified* 
local c = require("c");

//// [/home/src/workspaces/project/c.lua] *new* 
local c = 2;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/c.tlua
*refresh*    /home/src/workspaces/project/a.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/c.tlua
(computed .d.ts) /home/src/workspaces/project/a.tlua

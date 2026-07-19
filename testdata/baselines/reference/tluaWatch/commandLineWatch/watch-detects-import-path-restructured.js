currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tlua] *new* 
local util = require("lib.util");
//// [/home/src/workspaces/project/lib/util.tlua] *new* 
local util = "v1";
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
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
//// [/home/src/workspaces/project/index.lua] *new* 
local util = require("lib.util");

//// [/home/src/workspaces/project/lib/util.lua] *new* 
local util = "v1";


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/lib/util.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::


Edit [0]:: move file to new path and update import
//// [/home/src/workspaces/project/index.tlua] *modified* 
local util = require("src.util");
//// [/home/src/workspaces/project/lib/util.tlua] *deleted*
//// [/home/src/workspaces/project/src/util.tlua] *new* 
local util = "v2";


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/home/src/workspaces/project/index.lua] *modified* 
local util = require("src.util");

//// [/home/src/workspaces/project/src/util.lua] *new* 
local util = "v2";


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/util.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/util.tlua
(computed .d.ts) /home/src/workspaces/project/index.tlua

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tlua] *new* 
local mylib = require("mylib");
local v: string = mylib.lib;
//// [/home/src/workspaces/project/node_modules/mylib/new.tlua] *new* 
local lib: string = "s";
return { lib = lib };
//// [/home/src/workspaces/project/node_modules/mylib/old.tlua] *new* 
local lib: number = 1;
return { lib = lib };
//// [/home/src/workspaces/project/node_modules/mylib/package.json] *new* 
{"name": "mylib", "main": "old.tlua"}
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{}

tlua --watch
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96mindex.tlua[0m:[93m2[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.

[7m2[0m local v: string = mylib.lib;
[7m [0m [91m      ~[0m


Found 1 error in index.tlua[90m:2[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

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
local mylib = require("mylib");
local v = mylib.lib;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/node_modules/mylib/old.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::


Edit [0]:: change package.json main field
//// [/home/src/workspaces/project/node_modules/mylib/package.json] *modified* 
{"name": "mylib", "main": "new.tlua"}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96mindex.tlua[0m:[93m2[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.

[7m2[0m local v: string = mylib.lib;
[7m [0m [91m      ~[0m


Found 1 error in index.tlua[90m:2[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/node_modules/mylib/new.tlua
Signatures::
(used version)   /home/src/workspaces/project/node_modules/mylib/new.tlua


Diff:: watch does not yet invalidate require resolutions when a package.json affecting location changes, so the stale main (old.tlua, number) keeps its type error; the clean build resolves the new main
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -0,0 +1,8 @@
+[96mindex.tlua[0m:[93m2[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.
+
+[7m2[0m local v: string = mylib.lib;
+[7m [0m [91m      ~[0m
+
+
+Found 1 error in index.tlua[90m:2[0m
+
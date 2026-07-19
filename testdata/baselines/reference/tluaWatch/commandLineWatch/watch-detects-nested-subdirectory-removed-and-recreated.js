currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/lib/helper.tlua] *new* 
local helper = "v1";
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
	"compilerOptions": {},
	"include": ["src/**/*.tlua"]
}

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
//// [/home/src/workspaces/project/src/lib/helper.lua] *new* 
local helper = "v1";


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
  /home/src/workspaces/project/src (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/lib/helper.tlua
Signatures::


Edit [0]:: remove nested dir
//// [/home/src/workspaces/project/src/lib/helper.tlua] *deleted*


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/workspaces/project
  /home/src/workspaces/project/src (recursive)
tluaconfig.json::
SemanticDiagnostics::
Signatures::


Diff:: incremental has prior state and does not report no-inputs error
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,4 +0,0 @@
-[91merror[0m[90m TLUA18003: [0mNo inputs were found in config file '/home/src/workspaces/project/tluaconfig.json'. Specified 'include' paths were '["src/**/*.tlua"]' and 'exclude' paths were '[]'.
-
-Found 1 error.
-

Edit [1]:: recreate nested dir with new content
//// [/home/src/workspaces/project/src/lib/helper.tlua] *new* 
local helper = "v2";


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/home/src/workspaces/project/src/lib/helper.lua] *modified* 
local helper = "v2";


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
  /home/src/workspaces/project/src (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/lib/helper.tlua
Signatures::
(used version)   /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
(computed .d.ts) /home/src/workspaces/project/src/lib/helper.tlua

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local a = { private: 10 };
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
	"compilerOptions": {
            "noEmit": true
	}
}

tlua -w
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

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

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
Signatures::


Edit [0]:: fix error
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/a.tlua


Edit [1]:: emit after fixing error
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
	"compilerOptions": {
            
	}
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/home/src/workspaces/project/a.lua] *new* 
local a = "hello";


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [2]:: no emit run after fixing error
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
	"compilerOptions": {
            "noEmit": true,
            
	}
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [3]:: introduce error
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = { private: 10 };


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/a.tlua
Signatures::


Edit [4]:: emit when error
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
	"compilerOptions": {
            
	}
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = { private, 10 };


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/a.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/a.tlua


Edit [5]:: no emit run when error
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
	"compilerOptions": {
            "noEmit": true,
            
	}
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/a.tlua
Signatures::

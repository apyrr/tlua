currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local a = 1;
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 2;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
	"compilerOptions": {},
	"files": ["a.tlua", "b.tlua"]
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
//// [/home/src/workspaces/project/a.lua] *new* 
local a = 1;

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 2;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
Signatures::


Edit [0]:: delete file listed in files array
//// [/home/src/workspaces/project/b.tlua] *deleted*


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[91merror[0m[90m TS6053: [0mFile '/home/src/workspaces/project/b.tlua' not found.
  The file is in the program because:
    Part of 'files' list in tsconfig.json
  [96mtsconfig.json[0m:[93m3[0m:[93m22[0m - File is matched by 'files' list specified here.
    [7m3[0m  "files": ["a.tlua", "b.tlua"]
    [7m [0m [96m                     ~~~~~~~~[0m


Found 1 error.

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
tsconfig.json::
SemanticDiagnostics::
Signatures::

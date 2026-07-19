currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/a.tlua] *new* 
local a = 1;
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
//// [/home/src/workspaces/project/src/a.lua] *new* 
local a = 1;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
  /home/src/workspaces/project/src (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/a.tlua
Signatures::


Edit [0]:: create multiple new subdirs with files
//// [/home/src/workspaces/project/src/models/user.tlua] *new* 
interface User { name: string; }
//// [/home/src/workspaces/project/src/utils/format.tlua] *new* 
function format(s: string): string { return s.trim(); }


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96msrc/utils/format.tlua[0m:[93m1[0m:[93m47[0m - [91merror[0m[90m TLUA2339: [0mProperty 'trim' does not exist on type 'string'.

[7m1[0m function format(s: string): string { return s.trim(); }
[7m [0m [91m                                              ~~~~[0m


Found 1 error in src/utils/format.tlua[90m:1[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/workspaces/project/src/models/user.lua] *new* 

//// [/home/src/workspaces/project/src/utils/format.lua] *new* 
function format(s) { return s.trim(); }


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project
  /home/src/workspaces/project/src (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/models/user.tlua
*refresh*    /home/src/workspaces/project/src/utils/format.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/models/user.tlua
(computed .d.ts) /home/src/workspaces/project/src/utils/format.tlua

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tlua] *new* 
local x = null; local y: string = x;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{}

tlua --watch
ExitStatus:: Success
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96mindex.tlua[0m:[93m1[0m:[93m23[0m - [91merror[0m[90m TS2322: [0mType 'nil' is not assignable to type 'string'.

[7m1[0m local x = null; local y: string = x;
[7m [0m [91m                      ~[0m


Found 1 error in index.tlua[90m:1[0m

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
local x = nil;
local y = x;


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::


Edit [0]:: enable strict mode
//// [/home/src/workspaces/project/tsconfig.json] *modified* 
{"compilerOptions": {"strict": true}}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96mindex.tlua[0m:[93m1[0m:[93m23[0m - [91merror[0m[90m TS2322: [0mType 'nil' is not assignable to type 'string'.

[7m1[0m local x = null; local y: string = x;
[7m [0m [91m                      ~[0m


Found 1 error in index.tlua[90m:1[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/project (recursive)
tsconfig.json::
SemanticDiagnostics::
Signatures::

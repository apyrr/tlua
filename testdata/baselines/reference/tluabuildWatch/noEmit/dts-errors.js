currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/a.tlua] *new* 
local a = { private: 10 };
//// [/home/src/projects/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": true
    , "noEmit": true }
}

tlua -b -verbose -w
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua"]}
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    }
  ],
  "size": 61
}
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
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [0]:: Fix error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = "hello";


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [1]:: Emit after fixing error
//// [/home/src/projects/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": true
    , "noEmit": false }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/a.lua] *new* 
local a = "hello";

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [2]:: no Emit run after fixing error
//// [/home/src/projects/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": true
    , "noEmit": true }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [3]:: Introduce error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = { private: 10 };


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [4]:: Emit when error
//// [/home/src/projects/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": true
    , "noEmit": false }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.

//// [/home/src/projects/project/a.lua] *modified* 
local a = { private, 10 };

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [5]:: no Emit run when error
//// [/home/src/projects/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": true
    , "noEmit": true }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/project (recursive)
tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::

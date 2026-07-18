currentDirectory::/home/src/projects/myproject
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/configs/first/tsconfig.json] *new* 
{
    "extends": "../second/tsconfig.json",
    "include": ["${configDir}/src"],
    "compilerOptions": {
        "typeRoots": ["root1", "${configDir}/root2", "root3"],
        "types": [],
    },
}
//// [/home/src/projects/configs/second/tsconfig.json] *new* 
{
    "files": ["${configDir}/main.tlua"],
    "compilerOptions": {
        "declarationDir": "${configDir}/decls",
    },
    "watchOptions": {
        "excludeFiles": ["${configDir}/main.tlua"],
    },
}
//// [/home/src/projects/myproject/main.tlua] *new* 
// some comment
local y = 10;
//// [/home/src/projects/myproject/tsconfig.json] *new* 
{
    "extends": "../configs/first/tsconfig.json",
    "compilerOptions": {
        "declaration": true,
        "outDir": "outDir",
        "traceResolution": true,
    },
}

tlua --b -w --explainFiles --v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'outDir/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96mmain.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local y = 10;
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
main.tlua
   Part of 'files' list in tsconfig.json
[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/myproject/outDir/main.lua] *new* 
-- some comment
local y = 10;

//// [/home/src/projects/myproject/outDir/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["../main.tlua"]}
//// [/home/src/projects/myproject/outDir/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "../main.tlua"
      ],
      "original": "../main.tlua"
    }
  ],
  "size": 65
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
  /home/src/projects/configs/first
  /home/src/projects/configs/second
  /home/src/projects/myproject
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/myproject/main.tlua
Signatures::


Edit [0]:: edit extended config file
//// [/home/src/projects/configs/first/tsconfig.json] *modified* 
{
    "extends": "../second/tsconfig.json",
    "include": ["${configDir}/src"],
    "compilerOptions": {
        "typeRoots": ["${configDir}/root2"],
        "types": [],
    },
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'outDir/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96mmain.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local y = 10;
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
main.tlua
   Part of 'files' list in tsconfig.json
[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/home/src/projects/myproject/outDir/main.lua] *rewrite with same content*
//// [/home/src/projects/myproject/outDir/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/myproject/outDir/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/projects/configs/first
  /home/src/projects/configs/second
  /home/src/projects/myproject
tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/myproject/main.tlua
Signatures::

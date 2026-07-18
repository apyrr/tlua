currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/project/src/main.tsx] *new* 
local x = 10;
//// [/home/src/workspaces/solution/project/tsconfig.json] *new* 
{
    "compilerOptions": { "declaration": true },
    "include": ["src/**/*.tsx", "src/**/*.tlua"]
}

tlua --b project -v --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project/tsconfig.json' is out of date because output file 'project/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project/tsconfig.json'...

[96mproject/src/main.tsx[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 10;
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
project/src/main.tsx
   Matched by include pattern 'src/**/*.tsx' in 'project/tsconfig.json'
[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'project/tsconfig.json'...


Found 1 error in project/src/main.tsx[90m:1[0m

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
//// [/home/src/workspaces/solution/project/src/main.lua] *new* 
local x = 10;

//// [/home/src/workspaces/solution/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["./src/main.tsx"]}
//// [/home/src/workspaces/solution/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./src/main.tsx"
      ],
      "original": "./src/main.tsx"
    }
  ],
  "size": 67
}

project/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/project/src/main.tsx
Signatures::


Edit [0]:: no change

tlua --b project -v --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project/tsconfig.json' is out of date because buildinfo file 'project/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'project/tsconfig.json'...

[96mproject/src/main.tsx[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 10;
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
project/src/main.tsx
   Matched by include pattern 'src/**/*.tsx' in 'project/tsconfig.json'
[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'project/tsconfig.json'...


Found 1 error in project/src/main.tsx[90m:1[0m

//// [/home/src/workspaces/solution/project/src/main.lua] *rewrite with same content*
//// [/home/src/workspaces/solution/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/workspaces/solution/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

project/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/project/src/main.tsx
Signatures::


Edit [1]:: clean build

tlua -b project --clean
ExitStatus:: Success
Output::
//// [/home/src/workspaces/solution/project/src/main.lua] *deleted*
//// [/home/src/workspaces/solution/project/tsconfig.tsbuildinfo] *deleted*


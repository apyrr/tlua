currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local x = 10;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "dist",
        "rootDir": "src",
    },
}
//// [/home/src/workspaces/project/types/type.tlua] *new* 
type t = string;

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[91merror[0m[90m TLUA6059: [0mFile '/home/src/workspaces/project/types/type.tlua' is not under 'rootDir' '/home/src/workspaces/project/src'. 'rootDir' is expected to contain all source files.
  The file is in the program because:
    Matched by default include pattern '**/*'

Found 1 error.

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
//// [/home/src/workspaces/project/dist/index.lua] *new* 
local x = 10;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["./src/index.tlua","./types/type.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./src/index.tlua"
      ],
      "original": "./src/index.tlua"
    },
    {
      "files": [
        "./types/type.tlua"
      ],
      "original": "./types/type.tlua"
    }
  ],
  "size": 89
}
//// [/home/src/workspaces/project/types/type.lua] *new* 


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/src/index.tlua
*not cached* /home/src/workspaces/project/types/type.tlua
Signatures::


Edit [0]:: no change

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[91merror[0m[90m TLUA6059: [0mFile '/home/src/workspaces/project/types/type.tlua' is not under 'rootDir' '/home/src/workspaces/project/src'. 'rootDir' is expected to contain all source files.
  The file is in the program because:
    Matched by default include pattern '**/*'

Found 1 error.

//// [/home/src/workspaces/project/dist/index.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*
//// [/home/src/workspaces/project/types/type.lua] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/src/index.tlua
*not cached* /home/src/workspaces/project/types/type.tlua
Signatures::


Edit [1]:: Normal build without change, that does not block emit on error to show files that get emitted

tlua -p /home/src/workspaces/project/tluaconfig.json
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[91merror[0m[90m TLUA6059: [0mFile '/home/src/workspaces/project/types/type.tlua' is not under 'rootDir' '/home/src/workspaces/project/src'. 'rootDir' is expected to contain all source files.
  The file is in the program because:
    Matched by default include pattern '**/*'

Found 1 error.

//// [/home/src/workspaces/project/dist/index.lua] *rewrite with same content*
//// [/home/src/workspaces/project/types/type.lua] *rewrite with same content*


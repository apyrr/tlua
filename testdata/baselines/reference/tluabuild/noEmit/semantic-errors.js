currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/a.tlua] *new* 
local a: number = "hello"
//// [/home/src/projects/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": false,
        "declaration": false
    }
}

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":["./a.tlua"],"semanticErrors":true}
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    }
  ],
  "size": 69,
  "semanticErrors": true
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

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [0]:: no change

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [1]:: Fix error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = "hello";

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./a.tlua"]}
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    }
  ],
  "size": 47
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [2]:: no change

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'a.lua' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [3]:: Emit after fixing error

tlua -b -v
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'a.lua' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/projects/project/a.lua] *new* 
local a = "hello";

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [4]:: no change

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is up to date because newest input 'a.tlua' is older than output 'a.lua'




Edit [5]:: Introduce error
//// [/home/src/projects/project/a.tlua] *modified* 
local a: number = "hello"

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output 'tluaconfig.tluabuildinfo' is older than input 'a.tlua'

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./a.tlua"],"semanticErrors":true}
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    }
  ],
  "size": 69,
  "semanticErrors": true
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [6]:: Emit when error

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/a.lua] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [7]:: no change

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello"
[7m [0m [91m      ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::

currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/a.tlua] *new* 
local a = { private: 10 };
//// [/home/src/projects/project/tsconfig.json] *new* 
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
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua"]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
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

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [0]:: no change

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [1]:: Fix error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = "hello";

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./a.tlua"]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [2]:: no change

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'a.lua' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [3]:: Emit after fixing error

tlua -b -v
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'a.lua' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

//// [/home/src/projects/project/a.lua] *new* 
local a = "hello";

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/a.tlua
Signatures::


Edit [4]:: no change

tlua -b -v --noEmit
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is up to date because newest input 'a.tlua' is older than output 'a.lua'




Edit [5]:: Introduce error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = { private: 10 };

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output 'tsconfig.tsbuildinfo' is older than input 'a.tlua'

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua"]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [6]:: Emit when error

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/a.lua] *modified* 
local a = { private, 10 };

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::


Edit [7]:: no change

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in a.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
Signatures::

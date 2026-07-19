currentDirectory::/user/username/projects/noEmitOnError
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/noEmitOnError/shared/types/db.tlua] *new* 
interface A {
    name: string;
}
//// [/user/username/projects/noEmitOnError/src/main.tlua] *new* 
local a = { private: 10 };
//// [/user/username/projects/noEmitOnError/src/other.tlua] *new* 
console.log("hi");
//// [/user/username/projects/noEmitOnError/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./dev-build",
        "declaration": true,
        "incremental": false,
        "noEmitOnError": true,
    },
}

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'dev-build/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in src/main.tlua[90m:1[0m

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
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"]}
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "../shared/types/db.tlua"
      ],
      "original": "../shared/types/db.tlua"
    },
    {
      "files": [
        "../src/main.tlua"
      ],
      "original": "../src/main.tlua"
    },
    {
      "files": [
        "../src/other.tlua"
      ],
      "original": "../src/other.tlua"
    }
  ],
  "size": 115
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [0]:: no change

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'dev-build/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a = { p: 10 };

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'dev-build/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m14[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { p: 10 };
[7m [0m [91m             ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [2]:: no change

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'dev-build/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m14[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m1[0m local a = { p: 10 };
[7m [0m [91m             ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::

currentDirectory::/user/username/projects/noEmitOnError
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/noEmitOnError/shared/types/db.tlua] *new* 
interface A {
    name: string;
}
//// [/user/username/projects/noEmitOnError/src/main.tlua] *new* 
local a: string = 10;
//// [/user/username/projects/noEmitOnError/src/other.tlua] *new* 
console.log("hi");
//// [/user/username/projects/noEmitOnError/tsconfig.json] *new* 
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
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'dev-build/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'number' is not assignable to type 'string'.

[7m1[0m local a: string = 10;
[7m [0m [91m      ~[0m


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
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":["../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"],"semanticErrors":true}
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
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
  "size": 123,
  "semanticErrors": true
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/noEmitOnError/shared/types/db.tlua
*refresh*    /user/username/projects/noEmitOnError/src/main.tlua
*refresh*    /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [0]:: no change

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'dev-build/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96msrc/main.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'number' is not assignable to type 'string'.

[7m1[0m local a: string = 10;
[7m [0m [91m      ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/noEmitOnError/shared/types/db.tlua
*refresh*    /user/username/projects/noEmitOnError/src/main.tlua
*refresh*    /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a: string = "hello";

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'dev-build/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96mshared/types/db.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface A {
[7m [0m [91m~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: string = "hello";
[7m [0m [91m~~~~~[0m

[96msrc/other.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m console.log("hi");
[7m [0m [91m~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  shared/types/db.tlua[90m:1[0m
     1  src/main.tlua[90m:1[0m
     1  src/other.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"]}
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/noEmitOnError/shared/types/db.tlua
*refresh*    /user/username/projects/noEmitOnError/src/main.tlua
*refresh*    /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [2]:: no change

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'dev-build/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96mshared/types/db.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface A {
[7m [0m [91m~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: string = "hello";
[7m [0m [91m~~~~~[0m

[96msrc/other.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m console.log("hi");
[7m [0m [91m~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  shared/types/db.tlua[90m:1[0m
     1  src/main.tlua[90m:1[0m
     1  src/other.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *rewrite with same content*
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/noEmitOnError/shared/types/db.tlua
*refresh*    /user/username/projects/noEmitOnError/src/main.tlua
*refresh*    /user/username/projects/noEmitOnError/src/other.tlua
Signatures::

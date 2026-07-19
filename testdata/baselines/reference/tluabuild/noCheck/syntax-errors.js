currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local a = "hello
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 10;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "declaration": true,
        "incremental": false
    }
}

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/a.lua] *new* 
local a = "hello";

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 10;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"checkPending":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "checkPending": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 92
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [0]:: no change

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [1]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [2]:: no change

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [3]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 72
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
Signatures::


Edit [4]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
Signatures::


Edit [5]:: no change

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"checkPending":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "checkPending": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 92
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [6]:: Introduce error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [7]:: no change

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [8]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 72
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [9]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"checkPending":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "checkPending": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 92
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
Signatures::


Edit [10]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    }
  ],
  "size": 72
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
Signatures::


Edit [11]:: Add file with error
//// [/home/src/workspaces/project/c.tlua] *new* 
local c: number = "hello";

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *new* 
local c = "hello";

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua","./c.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    },
    {
      "files": [
        "./c.tlua"
      ],
      "original": "./c.tlua"
    }
  ],
  "size": 83
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
*refresh*    /home/src/workspaces/project/c.tlua
Signatures::


Edit [12]:: Introduce error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m17[0m - [91merror[0m[90m TLUA1002: [0mUnterminated string literal.

[7m1[0m local a = "hello
[7m [0m [91m                ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 4 errors in 3 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"checkPending":true,"root":["./a.tlua","./b.tlua","./c.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "checkPending": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    },
    {
      "files": [
        "./c.tlua"
      ],
      "original": "./c.tlua"
    }
  ],
  "size": 103
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
*not cached* /home/src/workspaces/project/c.tlua
Signatures::


Edit [13]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
*not cached* /home/src/workspaces/project/c.tlua
Signatures::


Edit [14]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua","./c.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    },
    {
      "files": [
        "./c.tlua"
      ],
      "original": "./c.tlua"
    }
  ],
  "size": 83
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
*refresh*    /home/src/workspaces/project/c.tlua
Signatures::


Edit [15]:: no change

tlua -b -v --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"checkPending":true,"root":["./a.tlua","./b.tlua","./c.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "checkPending": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    },
    {
      "files": [
        "./c.tlua"
      ],
      "original": "./c.tlua"
    }
  ],
  "size": 103
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/a.tlua
*not cached* /home/src/workspaces/project/b.tlua
*not cached* /home/src/workspaces/project/c.tlua
Signatures::


Edit [16]:: No Change run with checking

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'tluaconfig.json'...


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":["./a.tlua","./b.tlua","./c.tlua"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "./a.tlua"
      ],
      "original": "./a.tlua"
    },
    {
      "files": [
        "./b.tlua"
      ],
      "original": "./b.tlua"
    },
    {
      "files": [
        "./c.tlua"
      ],
      "original": "./c.tlua"
    }
  ],
  "size": 83
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
*refresh*    /home/src/workspaces/project/c.tlua
Signatures::

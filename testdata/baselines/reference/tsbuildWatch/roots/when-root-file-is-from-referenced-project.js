currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/projects/server/src/server.tlua] *new* 
local myClass = require('shared.src.myClass');
console.log('Hello, world!');
//// [/home/src/workspaces/solution/projects/server/tsconfig.json] *new* 
{
    "extends": "../../tsconfig.json",
    "compilerOptions": {
        "rootDir": "..",
        "outDir": "./dist",
    },
    "include": [ "src/**/*.tlua", "../shared/src/**/*.tlua" ],
    "references": [
        { "path": "../shared" },
    ],
}
//// [/home/src/workspaces/solution/projects/shared/src/logging.tlua] *new* 
function log(str: string) {
    console.log(str);
}
//// [/home/src/workspaces/solution/projects/shared/src/myClass.tlua] *new* 
local MyClass = { };
return { MyClass = MyClass };
//// [/home/src/workspaces/solution/projects/shared/src/random.tlua] *new* 
function randomFn(str: string) {
    console.log(str);
}
//// [/home/src/workspaces/solution/projects/shared/tsconfig.json] *new* 
{
    "extends": "../../tsconfig.json",
    "compilerOptions": {
        "outDir": "./dist",
    },
    "include": ["src/**/*.tlua"],
}
//// [/home/src/workspaces/solution/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
    },
    "references": [
        { "path": "projects/server" },
        { "path": "projects/shared" },
    ],
}

tlua --b -w projects/server -v --traceResolution --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * projects/shared/tsconfig.json
    * projects/server/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'projects/shared/tsconfig.json' is out of date because output file 'projects/shared/dist/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'projects/shared/tsconfig.json'...

[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function log(str: string) {
[7m [0m [91m~~~~~~~~[0m

[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local MyClass = { };
[7m [0m [91m~~~~~[0m

[96mprojects/shared/src/random.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function randomFn(str: string) {
[7m [0m [91m~~~~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/shared/src/logging.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
projects/shared/src/myClass.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
projects/shared/src/random.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
[[90mHH:MM:SS AM[0m] Project 'projects/server/tsconfig.json' is out of date because output file 'projects/server/dist/server/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'projects/server/tsconfig.json'...

======== Resolving module 'shared.src.myClass' from '/home/src/workspaces/solution/projects/server/src/server.tlua'. ========
File '/home/src/workspaces/solution/projects/shared/src/myClass.tlua' exists - use it as a name resolution result.
======== Module name 'shared.src.myClass' was successfully resolved to '/home/src/workspaces/solution/projects/shared/src/myClass.tlua'. ========
[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClass = require('shared.src.myClass');
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/server/src/server.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/server/tsconfig.json'
[[90mHH:MM:SS AM[0m] Found 4 errors. Watching for file changes.

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
//// [/home/src/workspaces/solution/projects/server/dist/server/src/server.lua] *new* 
local myClass = require('shared.src.myClass');
console.log('Hello, world!');

//// [/home/src/workspaces/solution/projects/server/dist/server/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","../../src/server.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"04eb5063e96ae22eac5b0a5fa5481c73-local myClass = require('shared.src.myClass');\nconsole.log('Hello, world!');"],"options":{"composite":true,"outDir":"..","rootDir":"../../.."},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/workspaces/solution/projects/server/dist/server/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../src/server.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../src/server.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "lib.luajit.d.tlua",
      "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "signature": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../../src/server.tlua",
      "version": "04eb5063e96ae22eac5b0a5fa5481c73-local myClass = require('shared.src.myClass');\nconsole.log('Hello, world!');",
      "signature": "04eb5063e96ae22eac5b0a5fa5481c73-local myClass = require('shared.src.myClass');\nconsole.log('Hello, world!');",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../../.."
  },
  "emitDiagnosticsPerFile": [
    [
      "../../src/server.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "../../src/server.tlua",
      "original": 2
    }
  ],
  "size": 1263
}
//// [/home/src/workspaces/solution/projects/shared/dist/src/logging.lua] *new* 
function log(str) {
    console.log(str);
}

//// [/home/src/workspaces/solution/projects/shared/dist/src/myClass.lua] *new* 
local MyClass = {};
return { MyClass = MyClass };

//// [/home/src/workspaces/solution/projects/shared/dist/src/random.lua] *new* 
function randomFn(str) {
    console.log(str);
}

//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../src/logging.tlua","../src/myClass.tlua","../src/random.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"0971487bc21733f4a85c16f26aa085fe-function log(str: string) {\n    console.log(str);\n}","5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };","07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}"],"options":{"composite":true,"outDir":"./"},"emitDiagnosticsPerFile":[[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/logging.tlua",
        "../src/myClass.tlua",
        "../src/random.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../src/logging.tlua",
    "../src/myClass.tlua",
    "../src/random.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "lib.luajit.d.tlua",
      "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "signature": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/logging.tlua",
      "version": "0971487bc21733f4a85c16f26aa085fe-function log(str: string) {\n    console.log(str);\n}",
      "signature": "0971487bc21733f4a85c16f26aa085fe-function log(str: string) {\n    console.log(str);\n}",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/myClass.tlua",
      "version": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "signature": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/random.tlua",
      "version": "07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}",
      "signature": "07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "./"
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/logging.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../src/myClass.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../src/random.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "../src/logging.tlua",
      "original": 2
    },
    {
      "file": "../src/myClass.tlua",
      "original": 3
    },
    {
      "file": "../src/random.tlua",
      "original": 4
    }
  ],
  "size": 1688
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)
projects/shared/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/projects/shared/src/logging.tlua
*refresh*    /home/src/workspaces/solution/projects/shared/src/myClass.tlua
*refresh*    /home/src/workspaces/solution/projects/shared/src/random.tlua
Signatures::

projects/server/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/projects/server/src/server.tlua
Signatures::


Edit [0]:: no change


Output::

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)


Diff:: TS100054 declaration emit errors are emit-time diagnostics not stored in buildinfo, so incremental and clean builds report them at different times
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,20 +0,0 @@
-[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m function log(str: string) {
-[7m [0m [91m~~~~~~~~[0m
-
-[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local MyClass = { };
-[7m [0m [91m~~~~~[0m
-
-[96mprojects/shared/src/random.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m function randomFn(str: string) {
-[7m [0m [91m~~~~~~~~[0m
-
-[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local myClass = require('shared.src.myClass');
-[7m [0m [91m~~~~~[0m
-

Edit [1]:: edit logging file
//// [/home/src/workspaces/solution/projects/shared/src/logging.tlua] *modified* 
function log(str: string) {
    console.log(str);
}local x = 10;


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * projects/shared/tsconfig.json
    * projects/server/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'projects/shared/tsconfig.json' is out of date because buildinfo file 'projects/shared/dist/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'projects/shared/tsconfig.json'...

[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function log(str: string) {
[7m [0m [91m~~~~~~~~[0m

[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local MyClass = { };
[7m [0m [91m~~~~~[0m

[96mprojects/shared/src/random.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function randomFn(str: string) {
[7m [0m [91m~~~~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/shared/src/logging.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
projects/shared/src/myClass.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
projects/shared/src/random.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
[[90mHH:MM:SS AM[0m] Project 'projects/server/tsconfig.json' is out of date because buildinfo file 'projects/server/dist/server/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'projects/server/tsconfig.json'...

======== Resolving module 'shared.src.myClass' from '/home/src/workspaces/solution/projects/server/src/server.tlua'. ========
File '/home/src/workspaces/solution/projects/shared/src/myClass.tlua' exists - use it as a name resolution result.
======== Module name 'shared.src.myClass' was successfully resolved to '/home/src/workspaces/solution/projects/shared/src/myClass.tlua'. ========
[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClass = require('shared.src.myClass');
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/server/src/server.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/server/tsconfig.json'
[[90mHH:MM:SS AM[0m] Found 4 errors. Watching for file changes.

//// [/home/src/workspaces/solution/projects/shared/dist/src/logging.lua] *modified* 
function log(str) {
    console.log(str);
}
local x = 10;

//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../src/logging.tlua","../src/myClass.tlua","../src/random.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };","07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}"],"options":{"composite":true,"outDir":"./"},"emitDiagnosticsPerFile":[[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/logging.tlua",
        "../src/myClass.tlua",
        "../src/random.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../src/logging.tlua",
    "../src/myClass.tlua",
    "../src/random.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "lib.luajit.d.tlua",
      "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "signature": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/logging.tlua",
      "version": "0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/myClass.tlua",
      "version": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "signature": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/random.tlua",
      "version": "07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}",
      "signature": "07919cda4104f1ffa7361b22f3221898-function randomFn(str: string) {\n    console.log(str);\n}",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "./"
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/logging.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../src/myClass.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../src/random.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "../src/logging.tlua",
      "original": 2
    },
    {
      "file": "../src/myClass.tlua",
      "original": 3
    },
    {
      "file": "../src/random.tlua",
      "original": 4
    }
  ],
  "size": 1868
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)
projects/shared/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/solution/projects/shared/src/logging.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/solution/projects/shared/src/logging.tlua

projects/server/tsconfig.json::
SemanticDiagnostics::
Signatures::


Edit [2]:: no change


Output::

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)


Diff:: TS100054 declaration emit errors are emit-time diagnostics not stored in buildinfo, so incremental and clean builds report them at different times
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,20 +0,0 @@
-[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m function log(str: string) {
-[7m [0m [91m~~~~~~~~[0m
-
-[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local MyClass = { };
-[7m [0m [91m~~~~~[0m
-
-[96mprojects/shared/src/random.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m function randomFn(str: string) {
-[7m [0m [91m~~~~~~~~[0m
-
-[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local myClass = require('shared.src.myClass');
-[7m [0m [91m~~~~~[0m
-

Edit [3]:: delete random file
//// [/home/src/workspaces/solution/projects/shared/src/random.tlua] *deleted*


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * projects/shared/tsconfig.json
    * projects/server/tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'projects/shared/tsconfig.json' is out of date because buildinfo file 'projects/shared/dist/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'projects/shared/tsconfig.json'...

[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function log(str: string) {
[7m [0m [91m~~~~~~~~[0m

[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local MyClass = { };
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/shared/src/logging.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
projects/shared/src/myClass.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/shared/tsconfig.json'
[[90mHH:MM:SS AM[0m] Project 'projects/server/tsconfig.json' is out of date because buildinfo file 'projects/server/dist/server/tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'projects/server/tsconfig.json'...

======== Resolving module 'shared.src.myClass' from '/home/src/workspaces/solution/projects/server/src/server.tlua'. ========
File '/home/src/workspaces/solution/projects/shared/src/myClass.tlua' exists - use it as a name resolution result.
======== Module name 'shared.src.myClass' was successfully resolved to '/home/src/workspaces/solution/projects/shared/src/myClass.tlua'. ========
[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClass = require('shared.src.myClass');
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
projects/server/src/server.tlua
   Matched by include pattern 'src/**/*.tlua' in 'projects/server/tsconfig.json'
[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../src/logging.tlua","../src/myClass.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };"],"options":{"composite":true,"outDir":"./"},"emitDiagnosticsPerFile":[[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/home/src/workspaces/solution/projects/shared/dist/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/logging.tlua",
        "../src/myClass.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../src/logging.tlua",
    "../src/myClass.tlua"
  ],
  "fileInfos": [
    {
      "fileName": "lib.luajit.d.tlua",
      "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "signature": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/logging.tlua",
      "version": "0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "0154faa1467efe984aa9bfe537f2d32c-function log(str: string) {\n    console.log(str);\n}local x = 10;",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/myClass.tlua",
      "version": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "signature": "5d7430c39261c48bb27fd171ed98d7cc-local MyClass = { };\nreturn { MyClass = MyClass };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "./"
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/logging.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../src/myClass.tlua",
      [
        {
          "end": 5,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "../src/logging.tlua",
      "original": 2
    },
    {
      "file": "../src/myClass.tlua",
      "original": 3
    }
  ],
  "size": 1631
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)
projects/shared/tsconfig.json::
SemanticDiagnostics::
Signatures::

projects/server/tsconfig.json::
SemanticDiagnostics::
Signatures::


Edit [4]:: no change


Output::

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /home/src/workspaces/solution
  /home/src/workspaces/solution/projects
  /home/src/workspaces/solution/projects/server
  /home/src/workspaces/solution/projects/server/src (recursive)
  /home/src/workspaces/solution/projects/shared
  /home/src/workspaces/solution/projects/shared/dist
  /home/src/workspaces/solution/projects/shared/dist/src
  /home/src/workspaces/solution/projects/shared/src (recursive)


Diff:: TS100054 declaration emit errors are emit-time diagnostics not stored in buildinfo, so incremental and clean builds report them at different times
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,15 +0,0 @@
-[96mprojects/shared/src/logging.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m function log(str: string) {
-[7m [0m [91m~~~~~~~~[0m
-
-[96mprojects/shared/src/myClass.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local MyClass = { };
-[7m [0m [91m~~~~~[0m
-
-[96mprojects/server/src/server.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.
-
-[7m1[0m local myClass = require('shared.src.myClass');
-[7m [0m [91m~~~~~[0m
-
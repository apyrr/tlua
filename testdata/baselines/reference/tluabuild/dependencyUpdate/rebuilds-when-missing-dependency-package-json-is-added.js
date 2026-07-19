currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/node_modules/my-dep/alt.tlua] *new* 
local myValue: number = 1;
return { myValue = myValue };
//// [/home/src/workspaces/project/node_modules/my-dep/index.tlua] *new* 
local myValue: string = "";
return { myValue = myValue };
//// [/home/src/workspaces/project/src/index.tlua] *new* 
local myDep = require("my-dep");
local value: string = myDep.myValue;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "dist",
        "strict": true
    },
    "include": ["src/**/*"]
}

tlua --b --verbose
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'dist/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myDep = require("my-dep");
[7m [0m [91m~~~~~[0m


Found 1 error in src/index.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/dist/src/index.lua] *new* 
local myDep = require("my-dep");
local value = myDep.myValue;

//// [/home/src/workspaces/project/dist/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[3],"missingPackageJsons":["../node_modules/my-dep/package.json","../node_modules/package.json"],"fileNames":["lib.luajit.d.tlua","../node_modules/my-dep/index.tlua","../src/index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"e2ccd7a95d0c4daa84cfe8425b49e31b-local myValue: string = \"\";\nreturn { myValue = myValue };","9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;"],"options":{"composite":true,"outDir":"./","strict":true},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[3]}
//// [/home/src/workspaces/project/dist/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/index.tlua"
      ],
      "original": 3
    }
  ],
  "missingPackageJsons": [
    "../node_modules/my-dep/package.json",
    "../node_modules/package.json"
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../node_modules/my-dep/index.tlua",
    "../src/index.tlua"
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
      "fileName": "../node_modules/my-dep/index.tlua",
      "version": "e2ccd7a95d0c4daa84cfe8425b49e31b-local myValue: string = \"\";\nreturn { myValue = myValue };",
      "signature": "e2ccd7a95d0c4daa84cfe8425b49e31b-local myValue: string = \"\";\nreturn { myValue = myValue };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/index.tlua",
      "version": "9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;",
      "signature": "9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "./",
    "strict": true
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/index.tlua",
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
      "file": "../src/index.tlua",
      "original": 3
    }
  ],
  "size": 1472
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/node_modules/my-dep/index.tlua
*refresh*    /home/src/workspaces/project/src/index.tlua
Signatures::


Edit [0]:: add package json redirecting main to a module with a breaking type change
//// [/home/src/workspaces/project/node_modules/my-dep/package.json] *new* 
{"main":"alt.tlua"}

tlua --b --verbose
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'dist/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myDep = require("my-dep");
[7m [0m [91m~~~~~[0m


Found 1 error in src/index.tlua[90m:1[0m

//// [/home/src/workspaces/project/dist/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[3],"packageJsons":["../node_modules/my-dep/package.json"],"fileNames":["lib.luajit.d.tlua","../node_modules/my-dep/alt.tlua","../src/index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"963b2fdb2cd32ffc68000ac201b4cb97-local myValue: number = 1;\nreturn { myValue = myValue };","9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;"],"options":{"composite":true,"outDir":"./","strict":true},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[3]}
//// [/home/src/workspaces/project/dist/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/index.tlua"
      ],
      "original": 3
    }
  ],
  "packageJsons": [
    "../node_modules/my-dep/package.json"
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../node_modules/my-dep/alt.tlua",
    "../src/index.tlua"
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
      "fileName": "../node_modules/my-dep/alt.tlua",
      "version": "963b2fdb2cd32ffc68000ac201b4cb97-local myValue: number = 1;\nreturn { myValue = myValue };",
      "signature": "963b2fdb2cd32ffc68000ac201b4cb97-local myValue: number = 1;\nreturn { myValue = myValue };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/index.tlua",
      "version": "9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;",
      "signature": "9d1eaa60483e57370e19e23fdca9b979-local myDep = require(\"my-dep\");\nlocal value: string = myDep.myValue;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "./",
    "strict": true
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/index.tlua",
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
      "file": "../src/index.tlua",
      "original": 3
    }
  ],
  "size": 1429
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/node_modules/my-dep/alt.tlua
Signatures::
(used version)   /home/src/workspaces/project/node_modules/my-dep/alt.tlua


Diff:: Incremental build does not re-check dependents of a changed Lua module (declaration emit is unsupported, TS100054, so module signatures never change), so the type error appears only in the clean build
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -3,11 +3,6 @@
 [7m1[0m local myDep = require("my-dep");
 [7m [0m [91m~~~~~[0m

-[96msrc/index.tlua[0m:[93m2[0m:[93m7[0m - [91merror[0m[90m TLUA2322: [0mType 'number' is not assignable to type 'string'.
-
-[7m2[0m local value: string = myDep.myValue;
-[7m [0m [91m      ~~~~~[0m
-
-
-Found 2 errors in the same file, starting at: src/index.tlua[90m:1[0m
+
+Found 1 error in src/index.tlua[90m:1[0m

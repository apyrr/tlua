currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/shared/index.tlua] *new* 
local a: Unrestricted = 1;
//// [/home/src/workspaces/solution/shared/tsconfig-base.json] *new* 
{
    "include": ["./typings-base/"],
}
//// [/home/src/workspaces/solution/shared/tsconfig.json] *new* 
{
    "extends": "./tsconfig-base.json",
    "compilerOptions": {
        "composite": true,
        "outDir": "../target-tsc-build/",
        "rootDir": "..",
    },
    "files": ["./index.tlua"],
}
//// [/home/src/workspaces/solution/shared/typings-base/globals.d.tlua] *new* 
type Unrestricted = any;
//// [/home/src/workspaces/solution/tsconfig.json] *new* 
{
    "references": [
        { "path": "./shared/tsconfig.json" },
        { "path": "./webpack/tsconfig.json" },
    ],
    "files": [],
}
//// [/home/src/workspaces/solution/webpack/index.tlua] *new* 
local b: Unrestricted = 1;
//// [/home/src/workspaces/solution/webpack/tsconfig.json] *new* 
{
    "extends": "../shared/tsconfig-base.json",
    "compilerOptions": {
        "composite": true,
        "outDir": "../target-tsc-build/",
        "rootDir": "..",
    },
    "files": ["./index.tlua"],
    "references": [{ "path": "../shared/tsconfig.json" }],
}

tlua --b --v --listFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * shared/tsconfig.json
    * webpack/tsconfig.json
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'shared/tsconfig.json' is out of date because output file 'target-tsc-build/shared/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'shared/tsconfig.json'...

[96mshared/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: Unrestricted = 1;
[7m [0m [91m~~~~~[0m

/home/src/tslibs/TS/Lib/lib.luajit.d.tlua
/home/src/workspaces/solution/shared/index.tlua
/home/src/workspaces/solution/shared/typings-base/globals.d.tlua
[[90mHH:MM:SS AM[0m] Project 'webpack/tsconfig.json' is out of date because output file 'target-tsc-build/webpack/tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'webpack/tsconfig.json'...

[96mwebpack/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b: Unrestricted = 1;
[7m [0m [91m~~~~~[0m

/home/src/tslibs/TS/Lib/lib.luajit.d.tlua
/home/src/workspaces/solution/webpack/index.tlua
/home/src/workspaces/solution/shared/typings-base/globals.d.tlua

Found 2 errors in 2 files.

Errors  Files
     1  shared/index.tlua[90m:1[0m
     1  webpack/index.tlua[90m:1[0m

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
//// [/home/src/workspaces/solution/target-tsc-build/shared/index.lua] *new* 
local a = 1;

//// [/home/src/workspaces/solution/target-tsc-build/shared/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../../shared/index.tlua","../../shared/typings-base/globals.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"d282f93170c042f4c8172e6ab1995401-local a: Unrestricted = 1;",{"version":"0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true,"outDir":"..","rootDir":"../.."},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/workspaces/solution/target-tsc-build/shared/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../shared/index.tlua",
        "../../shared/typings-base/globals.d.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../shared/index.tlua",
    "../../shared/typings-base/globals.d.tlua"
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
      "fileName": "../../shared/index.tlua",
      "version": "d282f93170c042f4c8172e6ab1995401-local a: Unrestricted = 1;",
      "signature": "d282f93170c042f4c8172e6ab1995401-local a: Unrestricted = 1;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../shared/typings-base/globals.d.tlua",
      "version": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
      "signature": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../.."
  },
  "emitDiagnosticsPerFile": [
    [
      "../../shared/index.tlua",
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
      "file": "../../shared/index.tlua",
      "original": 2
    }
  ],
  "size": 1378
}
//// [/home/src/workspaces/solution/target-tsc-build/webpack/index.lua] *new* 
local b = 1;

//// [/home/src/workspaces/solution/target-tsc-build/webpack/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../../webpack/index.tlua","../../shared/typings-base/globals.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b1b26a62f1b425bee20b5d60993deb43-local b: Unrestricted = 1;",{"version":"0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true,"outDir":"..","rootDir":"../.."},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/workspaces/solution/target-tsc-build/webpack/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../webpack/index.tlua",
        "../../shared/typings-base/globals.d.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../webpack/index.tlua",
    "../../shared/typings-base/globals.d.tlua"
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
      "fileName": "../../webpack/index.tlua",
      "version": "b1b26a62f1b425bee20b5d60993deb43-local b: Unrestricted = 1;",
      "signature": "b1b26a62f1b425bee20b5d60993deb43-local b: Unrestricted = 1;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../shared/typings-base/globals.d.tlua",
      "version": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
      "signature": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "0818246edc003d659f6bac1bc37ad307-type Unrestricted = any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../.."
  },
  "emitDiagnosticsPerFile": [
    [
      "../../webpack/index.tlua",
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
      "file": "../../webpack/index.tlua",
      "original": 2
    }
  ],
  "size": 1379
}

shared/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/shared/index.tlua
*refresh*    /home/src/workspaces/solution/shared/typings-base/globals.d.tlua
Signatures::

webpack/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/webpack/index.tlua
*refresh*    /home/src/workspaces/solution/shared/typings-base/globals.d.tlua
Signatures::

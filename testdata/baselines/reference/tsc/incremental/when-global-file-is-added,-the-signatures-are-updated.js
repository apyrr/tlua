currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua] *new* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function anotherFileWithSameReferenes() { }
//// [/home/src/workspaces/project/src/filePresent.tlua] *new* 
function something() { return 10; }
//// [/home/src/workspaces/project/src/main.tlua] *new* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() { }
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": { "composite": true },
    "include": ["src/**/*.tlua"],
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() { }
[7m [0m [91m~~~~~~~~[0m


Found 5 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     2  src/main.tlua[90m:2[0m

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
//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.lua] *new* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function anotherFileWithSameReferenes() { }

//// [/home/src/workspaces/project/src/filePresent.lua] *new* 
function something() { return 10; }

//// [/home/src/workspaces/project/src/main.lua] *new* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main() { }

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }","ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }","d1c332e0da046a53fb3d94a7c866c1fe-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }"],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "d1c332e0da046a53fb3d94a7c866c1fe-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }",
      "signature": "d1c332e0da046a53fb3d94a7c866c1fe-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 1913
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/filePresent.tlua
*refresh*    /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::


Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() { }
[7m [0m [91m~~~~~~~~[0m


Found 5 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     2  src/main.tlua[90m:2[0m


tsconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: Modify main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() { }something();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() { }something();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m20[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m3[0m function main() { }something();
[7m [0m [91m                   ~~~~~~~~~[0m


Found 6 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     3  src/main.tlua[90m:2[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main() { }
something();

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }","ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",{"version":"ab44ea0c5c1235fe4c558cabf498d06e-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"semanticDiagnosticsPerFile":[[4,[{"pos":106,"end":115,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]}]]],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "ab44ea0c5c1235fe4c558cabf498d06e-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "ab44ea0c5c1235fe4c558cabf498d06e-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tlua",
      [
        {
          "pos": 106,
          "end": 115,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 2244
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [2]:: Modify main file again
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() { }something();something();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() { }something();something();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m20[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m3[0m function main() { }something();something();
[7m [0m [91m                   ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m32[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m3[0m function main() { }something();something();
[7m [0m [91m                               ~~~~~~~~~[0m


Found 7 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     4  src/main.tlua[90m:2[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main() { }
something();
something();

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }","ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",{"version":"3be1c2d48fff1de11d372af78322107a-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"semanticDiagnosticsPerFile":[[4,[{"pos":106,"end":115,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":118,"end":127,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]}]]],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "3be1c2d48fff1de11d372af78322107a-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3be1c2d48fff1de11d372af78322107a-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tlua",
      [
        {
          "pos": 106,
          "end": 115,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 118,
          "end": 127,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 2370
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [3]:: Add new file and update main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./newFile.tlua"/>
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() { }something();something();foo();
//// [/home/src/workspaces/project/src/newFile.tlua] *new* 
function foo() { return 20; }

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m22[0m - [91merror[0m[90m TS6053: [0mFile './fileNotFound.tlua' not found.

[7m3[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                   ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m32[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                               ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m44[0m - [91merror[0m[90m TS2304: [0mCannot find name 'foo'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                                           ~~~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { return 20; }
[7m [0m [91m~~~~~~~~[0m


Found 9 errors in 4 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     5  src/main.tlua[90m:3[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./newFile.tlua"/>
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main() { }
something();
something();
foo();

//// [/home/src/workspaces/project/src/newFile.lua] *new* 
function foo() { return 20; }

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }","ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",{"version":"cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,6],[2,4,6]],"options":{"composite":true},"referencedMap":[[3,1],[5,2]],"semanticDiagnosticsPerFile":[[5,[{"pos":145,"end":154,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":157,"end":166,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":169,"end":172,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["foo"]}]]],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/newFile.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/newFile.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tlua",
      [
        {
          "pos": 145,
          "end": 154,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 157,
          "end": 166,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 169,
          "end": 172,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "foo"
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
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
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/newFile.tlua",
      "original": 4
    },
    {
      "file": "./src/main.tlua",
      "original": 5
    }
  ],
  "size": 2912
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/newFile.tlua
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/newFile.tlua
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [4]:: Write file that could not be resolved
//// [/home/src/workspaces/project/src/fileNotFound.tlua] *new* 
function something2() { return 20; }

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/fileNotFound.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something2() { return 20; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                   ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m32[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                               ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m44[0m - [91merror[0m[90m TS2304: [0mCannot find name 'foo'.

[7m4[0m function main() { }something();something();foo();
[7m [0m [91m                                           ~~~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { return 20; }
[7m [0m [91m~~~~~~~~[0m


Found 8 errors in 5 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/fileNotFound.tlua[90m:1[0m
     1  src/filePresent.tlua[90m:1[0m
     4  src/main.tlua[90m:4[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.lua] *rewrite with same content*
//// [/home/src/workspaces/project/src/fileNotFound.lua] *new* 
function something2() { return 20; }

//// [/home/src/workspaces/project/src/main.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,6]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/fileNotFound.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",{"version":"d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,3],[2,3,5]],"options":{"composite":true},"referencedMap":[[4,1],[6,2]],"semanticDiagnosticsPerFile":[[6,[{"pos":145,"end":154,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":157,"end":166,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":169,"end":172,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["foo"]}]]],"emitDiagnosticsPerFile":[[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[6,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5,6]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/fileNotFound.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        6
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/fileNotFound.tlua",
      "version": "d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "bb70c2336cd9f34182f3920fd3f496d9-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tlua",
      [
        {
          "pos": 145,
          "end": 154,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 157,
          "end": 166,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 169,
          "end": 172,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "foo"
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/fileNotFound.tlua",
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
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
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
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/fileNotFound.tlua",
      "original": 3
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 4
    },
    {
      "file": "./src/newFile.tlua",
      "original": 5
    },
    {
      "file": "./src/main.tlua",
      "original": 6
    }
  ],
  "size": 3441
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/fileNotFound.tlua
*refresh*    /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/fileNotFound.tlua
(computed .d.ts) /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [5]:: Modify main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./newFile.tlua"/>
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() { }something();something();foo();something();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() { }
[7m [0m [91m~~~~~~~~[0m

[96msrc/fileNotFound.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something2() { return 20; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() { return 10; }
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() { }something();something();foo();something();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();something();
[7m [0m [91m                   ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m32[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();something();
[7m [0m [91m                               ~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m44[0m - [91merror[0m[90m TS2304: [0mCannot find name 'foo'.

[7m4[0m function main() { }something();something();foo();something();
[7m [0m [91m                                           ~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m50[0m - [91merror[0m[90m TS2304: [0mCannot find name 'something'.

[7m4[0m function main() { }something();something();foo();something();
[7m [0m [91m                                                 ~~~~~~~~~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { return 20; }
[7m [0m [91m~~~~~~~~[0m


Found 9 errors in 5 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/fileNotFound.tlua[90m:1[0m
     1  src/filePresent.tlua[90m:1[0m
     5  src/main.tlua[90m:4[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./newFile.tlua"/>
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main() { }
something();
something();
foo();
something();

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,6]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/fileNotFound.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",{"version":"d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"aaa6ae9158ca05946f0d83e6e589fe3a-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();something();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,3],[2,3,5]],"options":{"composite":true},"referencedMap":[[4,1],[6,2]],"semanticDiagnosticsPerFile":[[6,[{"pos":145,"end":154,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":157,"end":166,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]},{"pos":169,"end":172,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["foo"]},{"pos":175,"end":184,"code":2304,"category":1,"messageKey":"Cannot_find_name_0_2304","messageArgs":["something"]}]]],"emitDiagnosticsPerFile":[[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[6,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5,6]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/fileNotFound.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        6
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "signature": "90fb0189e81698eb72c5c92453cf2ab4-function something() { return 10; }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/fileNotFound.tlua",
      "version": "d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "d97745dab1d2c6dc05ce702bd0c7145d-function something2() { return 20; }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "ee06af9d400ba2316ef28add59039c0c-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() { }",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "cf329dc888a898a1403ba3e35c2ec68e-function foo() { return 20; }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "aaa6ae9158ca05946f0d83e6e589fe3a-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();something();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "aaa6ae9158ca05946f0d83e6e589fe3a-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() { }something();something();foo();something();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tlua",
      [
        {
          "pos": 145,
          "end": 154,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 157,
          "end": 166,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        },
        {
          "pos": 169,
          "end": 172,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "foo"
          ]
        },
        {
          "pos": 175,
          "end": 184,
          "code": 2304,
          "category": 1,
          "messageKey": "Cannot_find_name_0_2304",
          "messageArgs": [
            "something"
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/fileNotFound.tlua",
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
      "./src/filePresent.tlua",
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
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
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
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/fileNotFound.tlua",
      "original": 3
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 4
    },
    {
      "file": "./src/newFile.tlua",
      "original": 5
    },
    {
      "file": "./src/main.tlua",
      "original": 6
    }
  ],
  "size": 3567
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua

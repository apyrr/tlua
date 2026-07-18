currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/alpha/src/a.tlua] *new* 
/// <reference path="../../beta/b.tlua"/>
local m = 3;
//// [/home/src/workspaces/project/alpha/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "bin",
    },
    "references": []
}
//// [/home/src/workspaces/project/beta/b.tlua] *new* 
local beta = 0;

tlua --p alpha/tsconfig.json
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96malpha/src/a.tlua[0m:[93m1[0m:[93m22[0m - [91merror[0m[90m TS6059: [0mFile '/home/src/workspaces/project/beta/b.tlua' is not under 'rootDir' '/home/src/workspaces/project/alpha'. 'rootDir' is expected to contain all source files.

[7m1[0m /// <reference path="../../beta/b.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~[0m

[96malpha/src/a.tlua[0m:[93m1[0m:[93m22[0m - [91merror[0m[90m TS6307: [0mFile '/home/src/workspaces/project/beta/b.tlua' is not listed within the file list of project '/home/src/workspaces/project/alpha/tsconfig.json'. Projects must list all files or use an 'include' pattern.

[7m1[0m /// <reference path="../../beta/b.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~[0m

[96malpha/src/a.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local m = 3;
[7m [0m [91m~~~~~[0m

[96mbeta/b.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local beta = 0;
[7m [0m [91m~~~~~[0m


Found 4 errors in 2 files.

Errors  Files
     3  alpha/src/a.tlua[90m:1[0m
     1  beta/b.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/alpha/bin/src/a.lua] *new* 
-- / <reference path="../../beta/b.tlua"/>
local m = 3;

//// [/home/src/workspaces/project/alpha/bin/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[3],"fileNames":["lib.luajit.d.tlua","../../beta/b.tlua","../src/a.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"54e89e7cf42a554307fbb5b2208265fb-local beta = 0;","0132a56b86b48910a5c49b2c76b872ed-/// <reference path=\"../../beta/b.tlua\"/>\nlocal m = 3;"],"fileIdsList":[[2]],"options":{"composite":true,"outDir":"./"},"referencedMap":[[3,1]],"emitDiagnosticsPerFile":[[3,[{"pos":42,"end":47,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/home/src/workspaces/project/alpha/bin/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../src/a.tlua"
      ],
      "original": 3
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../beta/b.tlua",
    "../src/a.tlua"
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
      "fileName": "../../beta/b.tlua",
      "version": "54e89e7cf42a554307fbb5b2208265fb-local beta = 0;",
      "signature": "54e89e7cf42a554307fbb5b2208265fb-local beta = 0;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/a.tlua",
      "version": "0132a56b86b48910a5c49b2c76b872ed-/// <reference path=\"../../beta/b.tlua\"/>\nlocal m = 3;",
      "signature": "0132a56b86b48910a5c49b2c76b872ed-/// <reference path=\"../../beta/b.tlua\"/>\nlocal m = 3;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "fileIdsList": [
    [
      "../../beta/b.tlua"
    ]
  ],
  "options": {
    "composite": true,
    "outDir": "./"
  },
  "referencedMap": {
    "../src/a.tlua": [
      "../../beta/b.tlua"
    ]
  },
  "emitDiagnosticsPerFile": [
    [
      "../src/a.tlua",
      [
        {
          "pos": 42,
          "end": 47,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../../beta/b.tlua",
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
      "file": "../../beta/b.tlua",
      "original": 2
    },
    {
      "file": "../src/a.tlua",
      "original": 3
    }
  ],
  "size": 1461
}
//// [/home/src/workspaces/project/beta/b.lua] *new* 
local beta = 0;


alpha/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/beta/b.tlua
*refresh*    /home/src/workspaces/project/alpha/src/a.tlua
Signatures::

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/alpha/a.tlua] *new* 
local m: number = 3;
return { m = m };
//// [/home/src/workspaces/project/alpha/bin/a.d.tlua] *new* 
declare alphaBuilt: number;
//// [/home/src/workspaces/project/alpha/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "bin",
    }
}
//// [/home/src/workspaces/project/beta/b.tlua] *new* 
local a = require('alpha.a');
//// [/home/src/workspaces/project/beta/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "bin",
        "rootDir": "..",
    },
    "references": [ { "path": "../alpha" } ]
}

tlua --p beta/tluaconfig.json --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mbeta/b.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require('alpha.a');
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
alpha/bin/a.d.tlua
   Imported via 'alpha.a' from file 'beta/b.tlua'
   File is output of project reference source 'alpha/a.tlua'
beta/b.tlua
   Matched by default include pattern '**/*'

Found 1 error in beta/b.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/beta/bin/beta/b.lua] *new* 
local a = require('alpha.a');

//// [/home/src/workspaces/project/beta/bin/beta/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[3],"fileNames":["lib.luajit.d.tlua","../../../alpha/bin/a.d.tlua","../../b.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"1dd4ea8365dcf97b37175ccb76bcb5f6-declare alphaBuilt: number;","affectsGlobalScope":true,"impliedNodeFormat":1},"661652bc48688e96e7a6fbaf74fa7cf8-local a = require('alpha.a');"],"options":{"composite":true,"outDir":"..","rootDir":"../../.."},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[3]}
//// [/home/src/workspaces/project/beta/bin/beta/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../b.tlua"
      ],
      "original": 3
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../../alpha/bin/a.d.tlua",
    "../../b.tlua"
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
      "fileName": "../../../alpha/bin/a.d.tlua",
      "version": "1dd4ea8365dcf97b37175ccb76bcb5f6-declare alphaBuilt: number;",
      "signature": "1dd4ea8365dcf97b37175ccb76bcb5f6-declare alphaBuilt: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "1dd4ea8365dcf97b37175ccb76bcb5f6-declare alphaBuilt: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../../b.tlua",
      "version": "661652bc48688e96e7a6fbaf74fa7cf8-local a = require('alpha.a');",
      "signature": "661652bc48688e96e7a6fbaf74fa7cf8-local a = require('alpha.a');",
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
      "../../b.tlua",
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
      "file": "../../b.tlua",
      "original": 3
    }
  ],
  "size": 1359
}

beta/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/alpha/bin/a.d.tlua
*refresh*    /home/src/workspaces/project/beta/b.tlua
Signatures::

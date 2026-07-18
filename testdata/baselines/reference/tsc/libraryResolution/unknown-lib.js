currentDirectory::/home/src/workspace/projects
useCaseSensitiveFileNames::true
Input::
//// [/home/src/tslibs/TS/Lib/lib.webworker.d.tlua] *new* 
interface WebWorkerInterface { }
//// [/home/src/workspace/projects/project1/core.d.tlua] *new* 
declare core: number;
//// [/home/src/workspace/projects/project1/file.tlua] *new* 
local file = 10;
//// [/home/src/workspace/projects/project1/file2.tlua] *new* 
/// <reference lib="webworker2"/>
/// <reference lib="unknownlib"/>
/// <reference lib="webworker"/>
//// [/home/src/workspace/projects/project1/index.tlua] *new* 
local x = "type1";
//// [/home/src/workspace/projects/project1/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "traceResolution": true,
        "libReplacement": true
    }
}
//// [/home/src/workspace/projects/project1/utils.d.tlua] *new* 
declare y: number;

tlua -p project1 --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
======== Resolving module '@typescript/lib-luajit' from '/home/src/workspace/projects/project1/__lib_node_modules_lookup_lib.luajit.d.tlua__.ts'. ========
======== Module name '@typescript/lib-luajit' was not resolved. ========
======== Module name '@typescript/lib-luajit' was not resolved. ========
[96mproject1/file.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local file = 10;
[7m [0m [91m~~~~~[0m

[96mproject1/file2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m /// <reference lib="webworker2"/>
[7m [0m [91m~[0m

[96mproject1/file2.tlua[0m:[93m1[0m:[93m21[0m - [91merror[0m[90m TS2726: [0mCannot find lib definition for 'webworker2'.

[7m1[0m /// <reference lib="webworker2"/>
[7m [0m [91m                    ~~~~~~~~~~[0m

[96mproject1/file2.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TS2726: [0mCannot find lib definition for 'unknownlib'.

[7m2[0m /// <reference lib="unknownlib"/>
[7m [0m [91m                    ~~~~~~~~~~[0m

[96mproject1/file2.tlua[0m:[93m3[0m:[93m21[0m - [91merror[0m[90m TS2726: [0mCannot find lib definition for 'webworker'.

[7m3[0m /// <reference lib="webworker"/>
[7m [0m [91m                    ~~~~~~~~~[0m

[96mproject1/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "type1";
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
project1/core.d.tlua
   Matched by default include pattern '**/*'
project1/file.tlua
   Matched by default include pattern '**/*'
project1/file2.tlua
   Matched by default include pattern '**/*'
project1/index.tlua
   Matched by default include pattern '**/*'
project1/utils.d.tlua
   Matched by default include pattern '**/*'

Found 6 errors in 3 files.

Errors  Files
     1  project1/file.tlua[90m:1[0m
     4  project1/file2.tlua[90m:1[0m
     1  project1/index.tlua[90m:1[0m

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
//// [/home/src/workspace/projects/project1/file.lua] *new* 
local file = 10;

//// [/home/src/workspace/projects/project1/file2.lua] *new* 
-- / <reference lib="webworker2"/>
-- / <reference lib="unknownlib"/>
-- / <reference lib="webworker"/>

//// [/home/src/workspace/projects/project1/index.lua] *new* 
local x = "type1";

//// [/home/src/workspace/projects/project1/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,6]],"fileNames":["lib.luajit.d.tlua","./core.d.tlua","./file.tlua","./file2.tlua","./index.tlua","./utils.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"04f379066c1bb74390b1066587d91241-declare core: number;","affectsGlobalScope":true,"impliedNodeFormat":1},"2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;","fe56266dc1ac79ed8495cdbbc4b61a97-/// <reference lib=\"webworker2\"/>\n/// <reference lib=\"unknownlib\"/>\n/// <reference lib=\"webworker\"/>","a07894186a824f107b104601266cd424-local x = \"type1\";",{"version":"f6fcbf475027f205d43c7595f33c0601-declare y: number;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[3,4,5]}
//// [/home/src/workspace/projects/project1/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./core.d.tlua",
        "./file.tlua",
        "./file2.tlua",
        "./index.tlua",
        "./utils.d.tlua"
      ],
      "original": [
        2,
        6
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./core.d.tlua",
    "./file.tlua",
    "./file2.tlua",
    "./index.tlua",
    "./utils.d.tlua"
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
      "fileName": "./core.d.tlua",
      "version": "04f379066c1bb74390b1066587d91241-declare core: number;",
      "signature": "04f379066c1bb74390b1066587d91241-declare core: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "04f379066c1bb74390b1066587d91241-declare core: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./file.tlua",
      "version": "2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;",
      "signature": "2c71e802ca277d650ab0cd4cb0d56c21-local file = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./file2.tlua",
      "version": "fe56266dc1ac79ed8495cdbbc4b61a97-/// <reference lib=\"webworker2\"/>\n/// <reference lib=\"unknownlib\"/>\n/// <reference lib=\"webworker\"/>",
      "signature": "fe56266dc1ac79ed8495cdbbc4b61a97-/// <reference lib=\"webworker2\"/>\n/// <reference lib=\"unknownlib\"/>\n/// <reference lib=\"webworker\"/>",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./index.tlua",
      "version": "a07894186a824f107b104601266cd424-local x = \"type1\";",
      "signature": "a07894186a824f107b104601266cd424-local x = \"type1\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./utils.d.tlua",
      "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "signature": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "f6fcbf475027f205d43c7595f33c0601-declare y: number;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./file.tlua",
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
      "./file2.tlua",
      [
        {
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./index.tlua",
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
      "file": "./file.tlua",
      "original": 3
    },
    {
      "file": "./file2.tlua",
      "original": 4
    },
    {
      "file": "./index.tlua",
      "original": 5
    }
  ],
  "size": 1891
}

project1/tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspace/projects/project1/core.d.tlua
*refresh*    /home/src/workspace/projects/project1/file.tlua
*refresh*    /home/src/workspace/projects/project1/file2.tlua
*refresh*    /home/src/workspace/projects/project1/index.tlua
*refresh*    /home/src/workspace/projects/project1/utils.d.tlua
Signatures::

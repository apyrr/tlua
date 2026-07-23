currentDirectory::/user/username/projects/sample1
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/sample1/core/anotherModule.tlua] *new* 
local World = "hello";
return { World = World };
//// [/user/username/projects/sample1/core/index.tlua] *new* 
multiply();
local someString: string = "HELLO WORLD";
function leftPad(s: string, n: number) return s + n; end
function multiply(a: number, b: number) return a * b; end
return { someString = someString, leftPad = leftPad, multiply = multiply };
//// [/user/username/projects/sample1/core/some_decl.d.tlua] *new* 
declare dts: any;
//// [/user/username/projects/sample1/core/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "declarationMap": true,
        "skipDefaultLibCheck": true,
    },
}
//// [/user/username/projects/sample1/logic/index.tlua] *new* 
local c = require('core.index');
function getSecondsInDay()
    return c.multiply(10, 15);
end
local mod = require('core.anotherModule');
return { getSecondsInDay = getSecondsInDay, m = mod };
//// [/user/username/projects/sample1/logic/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "sourceMap": true,
        "skipDefaultLibCheck": true,
        "rootDir": "..",
    },
    "references": [
        { "path": "../core" },
    ],
}
//// [/user/username/projects/sample1/tests/index.tlua] *new* 
local c = require('core.index');
local logic = require('logic.index');

c.leftPad("", 10);
logic.getSecondsInDay();

local mod = require('core.anotherModule');
return { m = mod };
//// [/user/username/projects/sample1/tests/tluaconfig.json] *new* 
{
    "references": [
        { "path": "../logic" },
    ],
    "files": ["index.tlua"],
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "skipDefaultLibCheck": true,
    },
}

tlua --b tests --verbose --stopBuildOnErrors
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * logic/tluaconfig.json
    * tests/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because output file 'core/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96mcore/anotherModule.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local World = "hello";
[7m [0m [91m~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA2554: [0mExpected 2 arguments, but got 0.

[7m1[0m multiply();
[7m [0m [91m~~~~~~~~[0m

  [96mcore/index.tlua[0m:[93m4[0m:[93m19[0m - An argument for 'a' was not provided.
    [7m4[0m function multiply(a: number, b: number) return a * b; end
    [7m [0m [96m                  ~~~~~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m multiply();
[7m [0m [91m~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'logic/tluaconfig.json' can't be built because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Skipping build of project 'logic/tluaconfig.json' because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Project 'tests/tluaconfig.json' can't be built because its dependency 'logic' was not built

[[90mHH:MM:SS AM[0m] Skipping build of project 'tests/tluaconfig.json' because its dependency 'logic' was not built


Found 3 errors in 2 files.

Errors  Files
     1  core/anotherModule.tlua[90m:1[0m
     2  core/index.tlua[90m:1[0m

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
//// [/user/username/projects/sample1/core/anotherModule.lua] *new* 
local World = "hello";
return { World = World };

//// [/user/username/projects/sample1/core/index.lua] *new* 
multiply();
local someString = "HELLO WORLD";
function leftPad(s, n)
  return s + n;
end
function multiply(a, b)
  return a * b;
end
return { someString = someString, leftPad = leftPad, multiply = multiply };

//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./anotherModule.tlua","./index.tlua","./some_decl.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };","65643217db0aa3790c3dcedb789e58f1-multiply();\nlocal someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",{"version":"42d1e28e7b1a08aaac11b6695520b779-declare dts: any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true,"declaration":true,"declarationMap":true,"skipDefaultLibCheck":true},"semanticDiagnosticsPerFile":[[3,[{"end":8,"code":2554,"category":1,"messageKey":"Expected_0_arguments_but_got_1_2554","messageArgs":["2","0"],"relatedInformation":[{"pos":129,"end":138,"code":6210,"category":3,"messageKey":"An_argument_for_0_was_not_provided_6210","messageArgs":["a"]}]}]]],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./anotherModule.tlua",
        "./index.tlua",
        "./some_decl.d.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./anotherModule.tlua",
    "./index.tlua",
    "./some_decl.d.tlua"
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
      "fileName": "./anotherModule.tlua",
      "version": "de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };",
      "signature": "de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./index.tlua",
      "version": "65643217db0aa3790c3dcedb789e58f1-multiply();\nlocal someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
      "signature": "65643217db0aa3790c3dcedb789e58f1-multiply();\nlocal someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./some_decl.d.tlua",
      "version": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
      "signature": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "declarationMap": true,
    "skipDefaultLibCheck": true
  },
  "semanticDiagnosticsPerFile": [
    [
      "./index.tlua",
      [
        {
          "end": 8,
          "code": 2554,
          "category": 1,
          "messageKey": "Expected_0_arguments_but_got_1_2554",
          "messageArgs": [
            "2",
            "0"
          ],
          "relatedInformation": [
            {
              "pos": 129,
              "end": 138,
              "code": 6210,
              "category": 3,
              "messageKey": "An_argument_for_0_was_not_provided_6210",
              "messageArgs": [
                "a"
              ]
            }
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./anotherModule.tlua",
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
      "./index.tlua",
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
      "file": "./anotherModule.tlua",
      "original": 2
    },
    {
      "file": "./index.tlua",
      "original": 3
    }
  ],
  "size": 2122
}

core/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/core/anotherModule.tlua
*refresh*    /user/username/projects/sample1/core/index.tlua
*refresh*    /user/username/projects/sample1/core/some_decl.d.tlua
Signatures::


Edit [0]:: no change

tlua --b tests --verbose --stopBuildOnErrors
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * logic/tluaconfig.json
    * tests/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because buildinfo file 'core/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96mcore/anotherModule.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local World = "hello";
[7m [0m [91m~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA2554: [0mExpected 2 arguments, but got 0.

[7m1[0m multiply();
[7m [0m [91m~~~~~~~~[0m

  [96mcore/index.tlua[0m:[93m4[0m:[93m19[0m - An argument for 'a' was not provided.
    [7m4[0m function multiply(a: number, b: number) return a * b; end
    [7m [0m [96m                  ~~~~~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m multiply();
[7m [0m [91m~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'logic/tluaconfig.json' can't be built because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Skipping build of project 'logic/tluaconfig.json' because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Project 'tests/tluaconfig.json' can't be built because its dependency 'logic' was not built

[[90mHH:MM:SS AM[0m] Skipping build of project 'tests/tluaconfig.json' because its dependency 'logic' was not built


Found 3 errors in 2 files.

Errors  Files
     1  core/anotherModule.tlua[90m:1[0m
     2  core/index.tlua[90m:1[0m


core/tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: fix error
//// [/user/username/projects/sample1/core/index.tlua] *modified* 
local someString: string = "HELLO WORLD";
function leftPad(s: string, n: number) return s + n; end
function multiply(a: number, b: number) return a * b; end
return { someString = someString, leftPad = leftPad, multiply = multiply };

tlua --b tests --verbose --stopBuildOnErrors
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * logic/tluaconfig.json
    * tests/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because buildinfo file 'core/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96mcore/anotherModule.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local World = "hello";
[7m [0m [91m~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local someString: string = "HELLO WORLD";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'logic/tluaconfig.json' can't be built because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Skipping build of project 'logic/tluaconfig.json' because its dependency 'core' has errors

[[90mHH:MM:SS AM[0m] Project 'tests/tluaconfig.json' can't be built because its dependency 'logic' was not built

[[90mHH:MM:SS AM[0m] Skipping build of project 'tests/tluaconfig.json' because its dependency 'logic' was not built


Found 2 errors in 2 files.

Errors  Files
     1  core/anotherModule.tlua[90m:1[0m
     1  core/index.tlua[90m:1[0m

//// [/user/username/projects/sample1/core/index.lua] *modified* 
local someString = "HELLO WORLD";
function leftPad(s, n)
  return s + n;
end
function multiply(a, b)
  return a * b;
end
return { someString = someString, leftPad = leftPad, multiply = multiply };

//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./anotherModule.tlua","./index.tlua","./some_decl.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };",{"version":"a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"42d1e28e7b1a08aaac11b6695520b779-declare dts: any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true,"declaration":true,"declarationMap":true,"skipDefaultLibCheck":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./anotherModule.tlua",
        "./index.tlua",
        "./some_decl.d.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./anotherModule.tlua",
    "./index.tlua",
    "./some_decl.d.tlua"
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
      "fileName": "./anotherModule.tlua",
      "version": "de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };",
      "signature": "de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./index.tlua",
      "version": "a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./some_decl.d.tlua",
      "version": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
      "signature": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "42d1e28e7b1a08aaac11b6695520b779-declare dts: any;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "declarationMap": true,
    "skipDefaultLibCheck": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./anotherModule.tlua",
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
      "file": "./anotherModule.tlua",
      "original": 2
    },
    {
      "file": "./index.tlua",
      "original": 3
    }
  ],
  "size": 1984
}

core/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /user/username/projects/sample1/core/index.tlua
Signatures::
(computed .d.ts) /user/username/projects/sample1/core/index.tlua

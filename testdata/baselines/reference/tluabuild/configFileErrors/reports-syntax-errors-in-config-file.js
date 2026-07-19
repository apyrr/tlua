currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
function foo() { }
//// [/home/src/workspaces/project/b.tlua] *new* 
function bar() { }
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
    },
    "files": [
        "a.tlua"
        "b.tlua"
    ]
}

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { }
[7m [0m [91m~~~~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function bar() { }
[7m [0m [91m~~~~~~~~[0m

[96mtluaconfig.json[0m:[93m7[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m7[0m         "b.tlua"
[7m [0m [91m        ~~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  tluaconfig.json[90m:7[0m

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
function foo() { }

//// [/home/src/workspaces/project/b.lua] *new* 
function bar() { }

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"700c3fba6f483a252a9c2dd153c5c8d0-function foo() { }","fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }"],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua"
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
      "fileName": "./a.tlua",
      "version": "700c3fba6f483a252a9c2dd153c5c8d0-function foo() { }",
      "signature": "700c3fba6f483a252a9c2dd153c5c8d0-function foo() { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }",
      "signature": "fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "file": "./a.tlua",
      "original": 2
    },
    {
      "file": "./b.tlua",
      "original": 3
    }
  ],
  "size": 1347
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
Signatures::


Edit [0]:: reports syntax errors after change to config file
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "composite": true, "declaration": true
    },
    "files": [
        "a.tlua"
        "b.tlua"
    ]
}

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { }
[7m [0m [91m~~~~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function bar() { }
[7m [0m [91m~~~~~~~~[0m

[96mtluaconfig.json[0m:[93m7[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m7[0m         "b.tlua"
[7m [0m [91m        ~~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  tluaconfig.json[90m:7[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: reports syntax errors after change to ts file
//// [/home/src/workspaces/project/a.tlua] *modified* 
function foo() { }function fooBar() { }

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { }function fooBar() { }
[7m [0m [91m~~~~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function bar() { }
[7m [0m [91m~~~~~~~~[0m

[96mtluaconfig.json[0m:[93m7[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m7[0m         "b.tlua"
[7m [0m [91m        ~~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  tluaconfig.json[90m:7[0m

//// [/home/src/workspaces/project/a.lua] *modified* 
function foo() { }
function fooBar() { }

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"9913eed4c728d7f8bcd32079c829fc4c-function foo() { }function fooBar() { }","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }"],"options":{"composite":true,"declaration":true},"emitDiagnosticsPerFile":[[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua"
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
      "fileName": "./a.tlua",
      "version": "9913eed4c728d7f8bcd32079c829fc4c-function foo() { }function fooBar() { }",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9913eed4c728d7f8bcd32079c829fc4c-function foo() { }function fooBar() { }",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }",
      "signature": "fad4ba9dd8cd2eb534c23aa7c53048f3-function bar() { }",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "file": "./a.tlua",
      "original": 2
    },
    {
      "file": "./b.tlua",
      "original": 3
    }
  ],
  "size": 1554
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/a.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/a.tlua


Edit [2]:: no change

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { }function fooBar() { }
[7m [0m [91m~~~~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function bar() { }
[7m [0m [91m~~~~~~~~[0m

[96mtluaconfig.json[0m:[93m7[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m7[0m         "b.tlua"
[7m [0m [91m        ~~~~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  tluaconfig.json[90m:7[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [3]:: builds after fixing config file errors
//// [/home/src/workspaces/project/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "composite": true, "declaration": true
    },
    "files": [
        "a.tlua",
        "b.tlua"
    ]
}

tlua --b
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() { }function fooBar() { }
[7m [0m [91m~~~~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function bar() { }
[7m [0m [91m~~~~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::

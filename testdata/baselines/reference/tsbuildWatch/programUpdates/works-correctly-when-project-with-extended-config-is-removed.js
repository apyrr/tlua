currentDirectory::/user/username/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/project/alpha.tsconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true,
    },
}
//// [/user/username/projects/project/bravo.tsconfig.json] *new* 
{
    "compilerOptions": {
        "strict": true,
    },
}
//// [/user/username/projects/project/commonFile1.tlua] *new* 
local x = 1
//// [/user/username/projects/project/commonFile2.tlua] *new* 
local y = 1
//// [/user/username/projects/project/other.tlua] *new* 
local z = 0;
//// [/user/username/projects/project/project1.tsconfig.json] *new* 
{
    "extends": "./alpha.tsconfig.json",
    "compilerOptions": {
        "composite": true,
    },
    "files": ["commonFile1.tlua", "commonFile2.tlua"],
}
//// [/user/username/projects/project/project2.tsconfig.json] *new* 
{
    "extends": "./bravo.tsconfig.json",
    "compilerOptions": {
        "composite": true,
    },
    "files": ["other.tlua"],
}
//// [/user/username/projects/project/tsconfig.json] *new* 
{
    "references": [
        {
            "path": "./project1.tsconfig.json",
        },
        {
            "path": "./project2.tsconfig.json",
        },
    ],
    "files": [],
}

tlua -b -w -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because output file 'project1.tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project1.tsconfig.json'...

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project2.tsconfig.json' is out of date because output file 'project2.tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project2.tsconfig.json'...

[96mother.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local z = 0;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

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
//// [/user/username/projects/project/commonFile1.lua] *new* 
local x = 1;

//// [/user/username/projects/project/commonFile2.lua] *new* 
local y = 1;

//// [/user/username/projects/project/other.lua] *new* 
local z = 0;

//// [/user/username/projects/project/project1.tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./commonFile1.tlua","./commonFile2.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"166ad5bce10b9fcd927e672c1db52a76-local x = 1","ba886869717e91f5e0849cc5121b20f3-local y = 1"],"options":{"composite":true,"strict":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/project/project1.tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./commonFile1.tlua",
        "./commonFile2.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./commonFile1.tlua",
    "./commonFile2.tlua"
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
      "fileName": "./commonFile1.tlua",
      "version": "166ad5bce10b9fcd927e672c1db52a76-local x = 1",
      "signature": "166ad5bce10b9fcd927e672c1db52a76-local x = 1",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./commonFile2.tlua",
      "version": "ba886869717e91f5e0849cc5121b20f3-local y = 1",
      "signature": "ba886869717e91f5e0849cc5121b20f3-local y = 1",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "strict": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./commonFile1.tlua",
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
      "./commonFile2.tlua",
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
      "file": "./commonFile1.tlua",
      "original": 2
    },
    {
      "file": "./commonFile2.tlua",
      "original": 3
    }
  ],
  "size": 1367
}
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"6699453b5b570f41d04e606c493e90b9-local z = 0;"],"options":{"composite":true,"strict":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./other.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./other.tlua"
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
      "fileName": "./other.tlua",
      "version": "6699453b5b570f41d04e606c493e90b9-local z = 0;",
      "signature": "6699453b5b570f41d04e606c493e90b9-local z = 0;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "strict": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./other.tlua",
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
      "file": "./other.tlua",
      "original": 2
    }
  ],
  "size": 1168
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project1.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/commonFile1.tlua
*refresh*    /user/username/projects/project/commonFile2.tlua
Signatures::

project2.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other.tlua
Signatures::


Edit [0]:: Remove project2 from base config
//// [/user/username/projects/project/tsconfig.json] *modified* 
{
    "references": [
        {
            "path": "./project1.tsconfig.json",
        },
    ],
    "files": [],
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project

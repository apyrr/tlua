currentDirectory::/user/username/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/project/alpha.tsconfig.json] *new* 
{}
//// [/user/username/projects/project/bravo.tsconfig.json] *new* 
{
    "extends": "./alpha.tsconfig.json",
}
//// [/user/username/projects/project/commonFile1.tlua] *new* 
local x = 1
//// [/user/username/projects/project/commonFile2.tlua] *new* 
local y = 1
//// [/user/username/projects/project/extendsConfig1.tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
    },
}
//// [/user/username/projects/project/extendsConfig2.tsconfig.json] *new* 
{
    "compilerOptions": {
        "strictNullChecks": false,
    },
}
//// [/user/username/projects/project/extendsConfig3.tsconfig.json] *new* 
{
    "compilerOptions": {
        "noImplicitAny": true,
    },
}
//// [/user/username/projects/project/other.tlua] *new* 
local z = 0;
//// [/user/username/projects/project/other2.tlua] *new* 
local k = 0;
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
//// [/user/username/projects/project/project3.tsconfig.json] *new* 
{
    "extends": [
        "./extendsConfig1.tsconfig.json",
        "./extendsConfig2.tsconfig.json",
        "./extendsConfig3.tsconfig.json",
    ],
    "compilerOptions": {
        "composite": false,
    },
    "files": ["other2.tlua"],
}

tlua -b -w -v project1.tsconfig.json project2.tsconfig.json project3.tsconfig.json
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

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

[[90mHH:MM:SS AM[0m] Project 'project3.tsconfig.json' is out of date because output file 'project3.tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'project3.tsconfig.json'...

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

//// [/user/username/projects/project/other2.lua] *new* 
local k = 0;

//// [/user/username/projects/project/project1.tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./commonFile1.tlua","./commonFile2.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"166ad5bce10b9fcd927e672c1db52a76-local x = 1","ba886869717e91f5e0849cc5121b20f3-local y = 1"],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
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
    "composite": true
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
  "size": 1353
}
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"6699453b5b570f41d04e606c493e90b9-local z = 0;"],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
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
    "composite": true
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
  "size": 1154
}
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":["./other2.tlua"]}
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./other2.tlua"
      ],
      "original": "./other2.tlua"
    }
  ],
  "size": 52
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

project3.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other2.tlua
Signatures::


Edit [0]:: Modify alpha config
//// [/user/username/projects/project/alpha.tsconfig.json] *modified* 
{
    "compilerOptions": {
        "strict": true
    }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because buildinfo file 'project1.tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'project1.tsconfig.json'...

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project2.tsconfig.json' is out of date because buildinfo file 'project2.tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'project2.tsconfig.json'...

[96mother.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local z = 0;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project1.tsconfig.json::
SemanticDiagnostics::
Signatures::

project2.tsconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: change bravo config
//// [/user/username/projects/project/bravo.tsconfig.json] *modified* 
{
    "extends": "./alpha.tsconfig.json",
    "compilerOptions": { "strict": false }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project2.tsconfig.json' is out of date because buildinfo file 'project2.tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'project2.tsconfig.json'...

[96mother.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local z = 0;
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"6699453b5b570f41d04e606c493e90b9-local z = 0;"],"options":{"composite":true,"strict":false},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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
    "strict": false
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
  "size": 1169
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project2.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other.tlua
Signatures::


Edit [2]:: project 2 extends alpha
//// [/user/username/projects/project/project2.tsconfig.json] *modified* 
{
    "extends": "./alpha.tsconfig.json",
    "files": ["other.tlua"]
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project2.tsconfig.json' is out of date because output 'other.lua' is older than input 'project2.tsconfig.json'

[[90mHH:MM:SS AM[0m] Building project 'project2.tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.

//// [/user/username/projects/project/other.lua] *rewrite with same content*
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./other.tlua"]}
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./other.tlua"
      ],
      "original": "./other.tlua"
    }
  ],
  "size": 51
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project2.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other.tlua
Signatures::


Edit [3]:: update aplha config
//// [/user/username/projects/project/alpha.tsconfig.json] *modified* 
{}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because buildinfo file 'project1.tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'project1.tsconfig.json'...

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project2.tsconfig.json' is out of date because output 'project2.tsconfig.tsbuildinfo' is older than input 'alpha.tsconfig.json'

[[90mHH:MM:SS AM[0m] Building project 'project2.tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.

//// [/user/username/projects/project/other.lua] *rewrite with same content*
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo] *rewrite with same content*
//// [/user/username/projects/project/project2.tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project1.tsconfig.json::
SemanticDiagnostics::
Signatures::

project2.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other.tlua
Signatures::


Edit [4]:: Modify extendsConfigFile2
//// [/user/username/projects/project/extendsConfig2.tsconfig.json] *modified* 
{
    "compilerOptions": { "strictNullChecks": true }
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project3.tsconfig.json' is out of date because output 'project3.tsconfig.tsbuildinfo' is older than input 'extendsConfig2.tsconfig.json'

[[90mHH:MM:SS AM[0m] Building project 'project3.tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.

//// [/user/username/projects/project/other2.lua] *rewrite with same content*
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo] *rewrite with same content*
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project3.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other2.tlua
Signatures::


Edit [5]:: Modify project 3
//// [/user/username/projects/project/project3.tsconfig.json] *modified* 
{
    "extends": ["./extendsConfig1.tsconfig.json", "./extendsConfig2.tsconfig.json"],
    "compilerOptions": { "composite": false },
    "files": ["other2.tlua"],
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project3.tsconfig.json' is out of date because output 'project3.tsconfig.tsbuildinfo' is older than input 'project3.tsconfig.json'

[[90mHH:MM:SS AM[0m] Building project 'project3.tsconfig.json'...

[[90mHH:MM:SS AM[0m] Found 2 errors. Watching for file changes.

//// [/user/username/projects/project/other2.lua] *rewrite with same content*
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo] *rewrite with same content*
//// [/user/username/projects/project/project3.tsconfig.tsbuildinfo.readable.baseline.txt] *rewrite with same content*

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project
project3.tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/project/other2.tlua
Signatures::


Edit [6]:: Delete extendedConfigFile2 and report error
//// [/user/username/projects/project/extendsConfig2.tsconfig.json] *deleted*


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * project1.tsconfig.json
    * project2.tsconfig.json
    * project3.tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'project1.tsconfig.json' is out of date because it has errors.

[96mcommonFile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = 1
[7m [0m [91m~~~~~[0m

[96mcommonFile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = 1
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'project3.tsconfig.json' is up to date because newest input 'other2.tlua' is older than output 'project3.tsconfig.tsbuildinfo'

[91merror[0m[90m TS5083: [0mCannot read file '/user/username/projects/project/extendsConfig2.tsconfig.json'.
[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/project

currentDirectory::/user/username/projects/sample1
useCaseSensitiveFileNames::true
Input::
//// [/home/src/tslibs/TS/Lib/lib.d.tlua] *new* 
/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
//// [/home/src/tslibs/TS/Lib/lib.esnext.full.d.tlua] *new* 
/// <reference no-default-lib="true"/>
/// <reference lib="esnext" />
//// [/user/username/projects/sample1/core/anotherModule.tlua] *new* 
local World = "hello";
return { World = World };
//// [/user/username/projects/sample1/core/index.tlua] *new* 
local someString: string = "HELLO WORLD";
function leftPad(s: string, n: number) return s + n; end
function multiply(a: number, b: number) return a * b; end
return { someString = someString, leftPad = leftPad, multiply = multiply };
//// [/user/username/projects/sample1/core/some_decl.d.tlua] *new* 
declare dts: any;
//// [/user/username/projects/sample1/core/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true,
        "listFiles": true,
        "listEmittedFiles": true,
        "target": "esnext",
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
        { "path": "../core" },
        { "path": "../logic" },
    ],
    "files": ["index.tlua"],
    "compilerOptions": {
        "composite": true,
        "declaration": true,
        "skipDefaultLibCheck": true,
        "rootDir": "..",
        "rootDir": "..",
    },
}

tlua --b core --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because output file 'core/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

TSFILE: /user/username/projects/sample1/core/anotherModule.lua
TSFILE: /user/username/projects/sample1/core/index.lua
TSFILE: /user/username/projects/sample1/core/tluaconfig.tluabuildinfo
/home/src/tslibs/TS/Lib/lib.luajit.d.tlua
/user/username/projects/sample1/core/anotherModule.tlua
/user/username/projects/sample1/core/index.tlua
/user/username/projects/sample1/core/some_decl.d.tlua
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
local someString = "HELLO WORLD";
function leftPad(s, n)
    return s + n;
end
function multiply(a, b)
    return a * b;
end
return { someString = someString, leftPad = leftPad, multiply = multiply };

//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./anotherModule.tlua","./index.tlua","./some_decl.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };","a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",{"version":"42d1e28e7b1a08aaac11b6695520b779-declare dts: any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"target":99}}
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
      "version": "a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
      "signature": "a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
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
    "target": 99
  },
  "size": 1454
}

core/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/core/anotherModule.tlua
*refresh*    /user/username/projects/sample1/core/index.tlua
*refresh*    /user/username/projects/sample1/core/some_decl.d.tlua
Signatures::


Edit [0]:: incremental-declaration-changes
//// [/user/username/projects/sample1/core/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "incremental": true,
        "listFiles": true,
        "listEmittedFiles": true,
        "target": "es5",
    },
}

tlua --b core --verbose
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because output 'core/tluaconfig.tluabuildinfo' is older than input 'core/tluaconfig.json'

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96mcore/tluaconfig.json[0m:[93m6[0m:[93m19[0m - [91merror[0m[90m TLUA5108: [0mOption 'target=ES5' has been removed. Please remove it from your configuration.

[7m6[0m         "target": "es5",
[7m [0m [91m                  ~~~~~[0m

TSFILE: /user/username/projects/sample1/core/anotherModule.lua
TSFILE: /user/username/projects/sample1/core/index.lua
TSFILE: /user/username/projects/sample1/core/tluaconfig.tluabuildinfo
/home/src/tslibs/TS/Lib/lib.luajit.d.tlua
/user/username/projects/sample1/core/anotherModule.tlua
/user/username/projects/sample1/core/index.tlua
/user/username/projects/sample1/core/some_decl.d.tlua

Found 1 error in core/tluaconfig.json[90m:6[0m

//// [/user/username/projects/sample1/core/anotherModule.lua] *rewrite with same content*
//// [/user/username/projects/sample1/core/index.lua] *rewrite with same content*
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./anotherModule.tlua","./index.tlua","./some_decl.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };","a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",{"version":"42d1e28e7b1a08aaac11b6695520b779-declare dts: any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"target":1}}
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "signature": "a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",
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
    "target": 1
  },
  "size": 1467
}

core/tluaconfig.json::
SemanticDiagnostics::
Signatures::

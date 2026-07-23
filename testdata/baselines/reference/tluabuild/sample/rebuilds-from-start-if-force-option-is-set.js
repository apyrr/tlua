currentDirectory::/user/username/projects/sample1
useCaseSensitiveFileNames::true
Input::
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

tlua --b tests
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mcore/anotherModule.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local World = "hello";
[7m [0m [91m~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local someString: string = "HELLO WORLD";
[7m [0m [91m~~~~~[0m

[96mlogic/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c = require('core.index');
[7m [0m [91m~~~~~[0m

[96mtests/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c = require('core.index');
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  core/anotherModule.tlua[90m:1[0m
     1  core/index.tlua[90m:1[0m
     1  logic/index.tlua[90m:1[0m
     1  tests/index.tlua[90m:1[0m

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
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./anotherModule.tlua","./index.tlua","./some_decl.d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"de8b11d3974d80d7f27a2b97bf4ebc60-local World = \"hello\";\nreturn { World = World };","a949631500033123bc74fbbc35557a06-local someString: string = \"HELLO WORLD\";\nfunction leftPad(s: string, n: number) return s + n; end\nfunction multiply(a: number, b: number) return a * b; end\nreturn { someString = someString, leftPad = leftPad, multiply = multiply };",{"version":"42d1e28e7b1a08aaac11b6695520b779-declare dts: any;","affectsGlobalScope":true,"impliedNodeFormat":1}],"options":{"composite":true,"declaration":true,"declarationMap":true,"skipDefaultLibCheck":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
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
  "size": 1817
}
//// [/user/username/projects/sample1/logic/index.lua] *new* 
local c = require('core.index');
function getSecondsInDay()
  return c.multiply(10, 15);
end
local mod = require('core.anotherModule');
return { getSecondsInDay = getSecondsInDay, m = mod };
//# sourceMappingURL=index.lua.map
//// [/user/username/projects/sample1/logic/index.lua.map] *new* 
{"version":3,"file":"index.lua","sourceRoot":"","sources":["index.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,OAAO,CAAC,YAAY,CAAC,CAAC;AAChC,SAAS,eAAe;EACpB,OAAO,CAAC,CAAC,QAAQ,CAAC,EAAE,EAAE,EAAE,CAAC,CAAC;GAC3B;AACH,MAAM,GAAG,GAAG,OAAO,CAAC,oBAAoB,CAAC,CAAC;AAC1C,OAAO,EAAE,eAAe,GAAG,eAAe,EAAE,CAAC,GAAG,GAAG,EAAE,CAAC"}
//// [/user/username/projects/sample1/logic/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"58f473d4d63e304a922bb0a3fd4eedf8-local c = require('core.index');\nfunction getSecondsInDay()\n    return c.multiply(10, 15);\nend\nlocal mod = require('core.anotherModule');\nreturn { getSecondsInDay = getSecondsInDay, m = mod };"],"options":{"composite":true,"declaration":true,"rootDir":"..","skipDefaultLibCheck":true,"sourceMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/sample1/logic/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./index.tlua"
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
      "fileName": "./index.tlua",
      "version": "58f473d4d63e304a922bb0a3fd4eedf8-local c = require('core.index');\nfunction getSecondsInDay()\n    return c.multiply(10, 15);\nend\nlocal mod = require('core.anotherModule');\nreturn { getSecondsInDay = getSecondsInDay, m = mod };",
      "signature": "58f473d4d63e304a922bb0a3fd4eedf8-local c = require('core.index');\nfunction getSecondsInDay()\n    return c.multiply(10, 15);\nend\nlocal mod = require('core.anotherModule');\nreturn { getSecondsInDay = getSecondsInDay, m = mod };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "rootDir": "..",
    "skipDefaultLibCheck": true,
    "sourceMap": true
  },
  "emitDiagnosticsPerFile": [
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
      "file": "./index.tlua",
      "original": 2
    }
  ],
  "size": 1417
}
//// [/user/username/projects/sample1/tests/index.lua] *new* 
local c = require('core.index');
local logic = require('logic.index');
c.leftPad("", 10);
logic.getSecondsInDay();
local mod = require('core.anotherModule');
return { m = mod };

//// [/user/username/projects/sample1/tests/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"83d3733a851c98511df99803d53b0730-local c = require('core.index');\nlocal logic = require('logic.index');\n\nc.leftPad(\"\", 10);\nlogic.getSecondsInDay();\n\nlocal mod = require('core.anotherModule');\nreturn { m = mod };"],"options":{"composite":true,"declaration":true,"rootDir":"..","skipDefaultLibCheck":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/sample1/tests/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./index.tlua"
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
      "fileName": "./index.tlua",
      "version": "83d3733a851c98511df99803d53b0730-local c = require('core.index');\nlocal logic = require('logic.index');\n\nc.leftPad(\"\", 10);\nlogic.getSecondsInDay();\n\nlocal mod = require('core.anotherModule');\nreturn { m = mod };",
      "signature": "83d3733a851c98511df99803d53b0730-local c = require('core.index');\nlocal logic = require('logic.index');\n\nc.leftPad(\"\", 10);\nlogic.getSecondsInDay();\n\nlocal mod = require('core.anotherModule');\nreturn { m = mod };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "rootDir": "..",
    "skipDefaultLibCheck": true
  },
  "emitDiagnosticsPerFile": [
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
      "file": "./index.tlua",
      "original": 2
    }
  ],
  "size": 1391
}

core/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/core/anotherModule.tlua
*refresh*    /user/username/projects/sample1/core/index.tlua
*refresh*    /user/username/projects/sample1/core/some_decl.d.tlua
Signatures::

logic/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/logic/index.tlua
Signatures::

tests/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/tests/index.tlua
Signatures::


Edit [0]:: --force build

tlua --b tests --verbose --force
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * logic/tluaconfig.json
    * tests/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is being forcibly rebuilt

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96mcore/anotherModule.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local World = "hello";
[7m [0m [91m~~~~~[0m

[96mcore/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local someString: string = "HELLO WORLD";
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'logic/tluaconfig.json' is being forcibly rebuilt

[[90mHH:MM:SS AM[0m] Building project 'logic/tluaconfig.json'...

[96mlogic/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c = require('core.index');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'tests/tluaconfig.json' is being forcibly rebuilt

[[90mHH:MM:SS AM[0m] Building project 'tests/tluaconfig.json'...

[96mtests/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c = require('core.index');
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  core/anotherModule.tlua[90m:1[0m
     1  core/index.tlua[90m:1[0m
     1  logic/index.tlua[90m:1[0m
     1  tests/index.tlua[90m:1[0m

//// [/user/username/projects/sample1/core/anotherModule.lua] *rewrite with same content*
//// [/user/username/projects/sample1/core/index.lua] *rewrite with same content*
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/sample1/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*
//// [/user/username/projects/sample1/logic/index.lua] *rewrite with same content*
//// [/user/username/projects/sample1/logic/index.lua.map] *rewrite with same content*
//// [/user/username/projects/sample1/logic/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/sample1/logic/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*
//// [/user/username/projects/sample1/tests/index.lua] *rewrite with same content*
//// [/user/username/projects/sample1/tests/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/sample1/tests/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*

core/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/core/anotherModule.tlua
*refresh*    /user/username/projects/sample1/core/index.tlua
*refresh*    /user/username/projects/sample1/core/some_decl.d.tlua
Signatures::

logic/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/logic/index.tlua
Signatures::

tests/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/tests/index.tlua
Signatures::

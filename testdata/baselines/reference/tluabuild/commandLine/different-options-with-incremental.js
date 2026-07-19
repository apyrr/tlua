currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local a = 10;local aLocal = 10;return { a = a };
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 10;local bLocal = 10;return { b = b };
//// [/home/src/workspaces/project/c.tlua] *new* 
local a = require("a");local c = a.a;return { c = c };
//// [/home/src/workspaces/project/d.tlua] *new* 
local b = require("b");local d = b.b;return { d = d };
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true
    }
}

tlua --build --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

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
local a = 10;
local aLocal = 10;
return { a = a };

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 10;
local bLocal = 10;
return { b = b };

//// [/home/src/workspaces/project/c.lua] *new* 
local a = require("a");
local c = a.a;
return { c = c };

//// [/home/src/workspaces/project/d.lua] *new* 
local b = require("b");
local d = b.b;
return { d = d };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };","239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "signature": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "size": 1294
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/a.tlua
*refresh*    /home/src/workspaces/project/b.tlua
*refresh*    /home/src/workspaces/project/c.tlua
*refresh*    /home/src/workspaces/project/d.tlua
Signatures::


Edit [0]:: with sourceMap

tlua --build --verbose --sourceMap
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 10;
return { a = a };
//# sourceMappingURL=a.lua.map
//// [/home/src/workspaces/project/a.lua.map] *new* 
{"version":3,"file":"a.lua","sourceRoot":"","sources":["a.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,EAAE,CAAC;AAAA,MAAM,MAAM,GAAG,EAAE,CAAC;AAAA,OAAO,EAAE,CAAC,GAAG,CAAC,EAAE,CAAC"}
//// [/home/src/workspaces/project/b.lua] *modified* 
local b = 10;
local bLocal = 10;
return { b = b };
//# sourceMappingURL=b.lua.map
//// [/home/src/workspaces/project/b.lua.map] *new* 
{"version":3,"file":"b.lua","sourceRoot":"","sources":["b.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,EAAE,CAAC;AAAA,MAAM,MAAM,GAAG,EAAE,CAAC;AAAA,OAAO,EAAE,CAAC,GAAG,CAAC,EAAE,CAAC"}
//// [/home/src/workspaces/project/c.lua] *modified* 
local a = require("a");
local c = a.a;
return { c = c };
//# sourceMappingURL=c.lua.map
//// [/home/src/workspaces/project/c.lua.map] *new* 
{"version":3,"file":"c.lua","sourceRoot":"","sources":["c.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,OAAO,CAAC,GAAG,CAAC,CAAC;AAAA,MAAM,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC;AAAA,OAAO,EAAE,CAAC,GAAG,CAAC,EAAE,CAAC"}
//// [/home/src/workspaces/project/d.lua] *modified* 
local b = require("b");
local d = b.b;
return { d = d };
//# sourceMappingURL=d.lua.map
//// [/home/src/workspaces/project/d.lua.map] *new* 
{"version":3,"file":"d.lua","sourceRoot":"","sources":["d.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,OAAO,CAAC,GAAG,CAAC,CAAC;AAAA,MAAM,CAAC,GAAG,CAAC,CAAC,CAAC,CAAC;AAAA,OAAO,EAAE,CAAC,GAAG,CAAC,EAAE,CAAC"}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };","239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"sourceMap":true}}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "signature": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "sourceMap": true
  },
  "size": 1323
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: should re-emit only js so they dont contain sourcemap

tlua --build --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 10;
return { a = a };

//// [/home/src/workspaces/project/b.lua] *modified* 
local b = 10;
local bLocal = 10;
return { b = b };

//// [/home/src/workspaces/project/c.lua] *modified* 
local a = require("a");
local c = a.a;
return { c = c };

//// [/home/src/workspaces/project/d.lua] *modified* 
local b = require("b");
local d = b.b;
return { d = d };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };","239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "signature": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "size": 1294
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [2]:: with declaration, emit Dts and should not emit js

tlua --build --verbose --declaration
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = 10;local aLocal = 10;return { a = a };
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };","239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"declaration":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "signature": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 1832
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [3]:: with declaration and declarationMap

tlua --build --verbose --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = 10;local aLocal = 10;return { a = a };
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };","239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"declaration":true,"declarationMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "signature": "2257d1c9ea062ec8217a13b00e6bd915-local a = 10;local aLocal = 10;return { a = a };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 1854
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [4]:: no change

tlua --build --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is up to date because newest input 'd.tlua' is older than output 'tluaconfig.tluabuildinfo'




Edit [5]:: local change
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = 10;local aLocal = 100;return { a = a };

tlua --build --verbose
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output 'tluaconfig.tluabuildinfo' is older than input 'a.tlua'

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 100;
return { a = a };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "emitDiagnosticsPerFile": [
    [
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 1849
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/a.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/a.tlua


Diff:: TS100054 declaration emit errors are emit-time diagnostics not stored in buildinfo, so incremental and clean builds report them at different times
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -0,0 +1,23 @@
+[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.
+
+[7m1[0m local b = 10;local bLocal = 10;return { b = b };
+[7m [0m [91m~~~~~[0m
+
+[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.
+
+[7m1[0m local a = require("a");local c = a.a;return { c = c };
+[7m [0m [91m~~~~~[0m
+
+[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.
+
+[7m1[0m local b = require("b");local d = b.b;return { d = d };
+[7m [0m [91m~~~~~[0m
+
+
+Found 3 errors in 3 files.
+
+Errors  Files
+     1  b.tlua[90m:1[0m
+     1  c.tlua[90m:1[0m
+     1  d.tlua[90m:1[0m
+

Edit [6]:: with declaration and declarationMap

tlua --build --verbose --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = 10;local aLocal = 100;return { a = a };
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"declaration":true,"declarationMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 2022
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [7]:: no change

tlua --build --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is up to date because newest input 'a.tlua' is older than output 'tluaconfig.tluabuildinfo'




Edit [8]:: with inlineSourceMap

tlua --build --verbose --inlineSourceMap
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 100;
return { a = a };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYS5sdWEiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJhLnRsdWEiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUEsTUFBTSxDQUFDLEdBQUcsRUFBRSxDQUFDO0FBQUEsTUFBTSxNQUFNLEdBQUcsR0FBRyxDQUFDO0FBQUEsT0FBTyxFQUFFLENBQUMsR0FBRyxDQUFDLEVBQUUsQ0FBQyJ9
//// [/home/src/workspaces/project/b.lua] *modified* 
local b = 10;
local bLocal = 10;
return { b = b };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYi5sdWEiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJiLnRsdWEiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUEsTUFBTSxDQUFDLEdBQUcsRUFBRSxDQUFDO0FBQUEsTUFBTSxNQUFNLEdBQUcsRUFBRSxDQUFDO0FBQUEsT0FBTyxFQUFFLENBQUMsR0FBRyxDQUFDLEVBQUUsQ0FBQyJ9
//// [/home/src/workspaces/project/c.lua] *modified* 
local a = require("a");
local c = a.a;
return { c = c };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiYy5sdWEiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJjLnRsdWEiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUEsTUFBTSxDQUFDLEdBQUcsT0FBTyxDQUFDLEdBQUcsQ0FBQyxDQUFDO0FBQUEsTUFBTSxDQUFDLEdBQUcsQ0FBQyxDQUFDLENBQUMsQ0FBQztBQUFBLE9BQU8sRUFBRSxDQUFDLEdBQUcsQ0FBQyxFQUFFLENBQUMifQ==
//// [/home/src/workspaces/project/d.lua] *modified* 
local b = require("b");
local d = b.b;
return { d = d };
//# sourceMappingURL=data:application/json;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiZC5sdWEiLCJzb3VyY2VSb290IjoiIiwic291cmNlcyI6WyJkLnRsdWEiXSwibmFtZXMiOltdLCJtYXBwaW5ncyI6IkFBQUEsTUFBTSxDQUFDLEdBQUcsT0FBTyxDQUFDLEdBQUcsQ0FBQyxDQUFDO0FBQUEsTUFBTSxDQUFDLEdBQUcsQ0FBQyxDQUFDLENBQUMsQ0FBQztBQUFBLE9BQU8sRUFBRSxDQUFDLEdBQUcsQ0FBQyxFQUFFLENBQUMifQ==
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"inlineSourceMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "inlineSourceMap": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 2004
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [9]:: with sourceMap

tlua --build --verbose --sourceMap
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 100;
return { a = a };
//# sourceMappingURL=a.lua.map
//// [/home/src/workspaces/project/a.lua.map] *modified* 
{"version":3,"file":"a.lua","sourceRoot":"","sources":["a.tlua"],"names":[],"mappings":"AAAA,MAAM,CAAC,GAAG,EAAE,CAAC;AAAA,MAAM,MAAM,GAAG,GAAG,CAAC;AAAA,OAAO,EAAE,CAAC,GAAG,CAAC,EAAE,CAAC"}
//// [/home/src/workspaces/project/b.lua] *modified* 
local b = 10;
local bLocal = 10;
return { b = b };
//# sourceMappingURL=b.lua.map
//// [/home/src/workspaces/project/b.lua.map] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *modified* 
local a = require("a");
local c = a.a;
return { c = c };
//# sourceMappingURL=c.lua.map
//// [/home/src/workspaces/project/c.lua.map] *rewrite with same content*
//// [/home/src/workspaces/project/d.lua] *modified* 
local b = require("b");
local d = b.b;
return { d = d };
//# sourceMappingURL=d.lua.map
//// [/home/src/workspaces/project/d.lua.map] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"sourceMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "sourceMap": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 1998
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [10]:: emit js files

tlua --build --verbose
ExitStatus:: Success
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates there is change in compilerOptions

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

//// [/home/src/workspaces/project/a.lua] *modified* 
local a = 10;
local aLocal = 100;
return { a = a };

//// [/home/src/workspaces/project/b.lua] *modified* 
local b = 10;
local bLocal = 10;
return { b = b };

//// [/home/src/workspaces/project/c.lua] *modified* 
local a = require("a");
local c = a.a;
return { c = c };

//// [/home/src/workspaces/project/d.lua] *modified* 
local b = require("b");
local d = b.b;
return { d = d };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 1969
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [11]:: with declaration and declarationMap

tlua --build --verbose --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = 10;local aLocal = 100;return { a = a };
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };","940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };","815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };"],"options":{"declaration":true,"declarationMap":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./a.tlua",
        "./b.tlua",
        "./c.tlua",
        "./d.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
      "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3f36fd34db630e96673a5824475e8f88-local a = 10;local aLocal = 100;return { a = a };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "signature": "239e7f67fbde956c9dd7a673bca3b1a0-local b = 10;local bLocal = 10;return { b = b };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "signature": "940a72b39315e7ff83d038b19313aa16-local a = require(\"a\");local c = a.a;return { c = c };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "signature": "815394481ecaa45a4cb1ed5b87eb4a11-local b = require(\"b\");local d = b.b;return { d = d };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./a.tlua",
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
      "./b.tlua",
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
      "./c.tlua",
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
      "./d.tlua",
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
  "size": 2022
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [12]:: with declaration and declarationMap, should not re-emit

tlua --build --verbose --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = 10;local aLocal = 100;return { a = a };
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;local bLocal = 10;return { b = b };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = require("a");local c = a.a;return { c = c };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = require("b");local d = b.b;return { d = d };
[7m [0m [91m~~~~~[0m


Found 4 errors in 4 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m
     1  d.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::

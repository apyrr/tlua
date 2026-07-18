currentDirectory::/home/project
useCaseSensitiveFileNames::false
Input::
//// [/home/node_modules/lib1/index.tlua] *new* 
local someLib = require('someLib');
local foo1: string = someLib.foo.foo;
return { foo1 = foo1 };
//// [/home/node_modules/lib1/package.json] *new* 
{
    "name": "lib1"
}
//// [/home/node_modules/lib2/index.tlua] *new* 
local someLib = require('someLib');
local foo2: string = someLib.foo.foo;
return { foo2 = foo2 };
//// [/home/node_modules/lib2/package.json] *new* 
{
    "name": "lib2"
}
//// [/home/node_modules/otherLib/index.tlua] *new* 
local str: string = "s";
return { str = str };
//// [/home/node_modules/otherLib/package.json] *new* 
{
    "name": "otherlib"
}
//// [/home/node_modules/someLib/index.tlua] *new* 
local other = require('otherlib');
local foo = { foo = other.str };
return { foo = foo };
//// [/home/node_modules/someLib/package.json] *new* 
{
    "name": "somelib"
}
//// [/home/project/src/index.tlua] *new* 
local lib1 = require('lib1');
local lib2 = require('lib2');
local someLib = require('someLib');
local otherLib = require('otherlib');
local foo1: string = lib1.foo1;
local foo2: string = lib2.foo2;
//// [/home/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true
    },
}

tlua -p .
ExitStatus:: Success
Output::
//// [/home/project/src/index.lua] *new* 
local lib1 = require('lib1');
local lib2 = require('lib2');
local someLib = require('someLib');
local otherLib = require('otherlib');
local foo1 = lib1.foo1;
local foo2 = lib2.foo2;

//// [/home/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[6],"packageJsons":["../node_modules/lib1/package.json","../node_modules/lib2/package.json","../node_modules/otherLib/package.json","../node_modules/someLib/package.json"],"fileNames":["lib.luajit.d.tlua","../node_modules/otherlib/index.tlua","../node_modules/somelib/index.tlua","../node_modules/lib1/index.tlua","../node_modules/lib2/index.tlua","./src/index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"30dbc7de1f476c74bc74a1ae8005b9bb-local str: string = \"s\";\nreturn { str = str };","23a32461648fa5b357b2795807b922cc-local other = require('otherlib');\nlocal foo = { foo = other.str };\nreturn { foo = foo };","2a90a454e186105000c8c0193d416555-local someLib = require('someLib');\nlocal foo1: string = someLib.foo.foo;\nreturn { foo1 = foo1 };","ac9f78442c66b675882139805d03b966-local someLib = require('someLib');\nlocal foo2: string = someLib.foo.foo;\nreturn { foo2 = foo2 };","cc6cfaa34cb010347e24399e184266ae-local lib1 = require('lib1');\nlocal lib2 = require('lib2');\nlocal someLib = require('someLib');\nlocal otherLib = require('otherlib');\nlocal foo1: string = lib1.foo1;\nlocal foo2: string = lib2.foo2;"]}
//// [/home/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/index.tlua"
      ],
      "original": 6
    }
  ],
  "packageJsons": [
    "../node_modules/lib1/package.json",
    "../node_modules/lib2/package.json",
    "../node_modules/otherLib/package.json",
    "../node_modules/someLib/package.json"
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../node_modules/otherlib/index.tlua",
    "../node_modules/somelib/index.tlua",
    "../node_modules/lib1/index.tlua",
    "../node_modules/lib2/index.tlua",
    "./src/index.tlua"
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
      "fileName": "../node_modules/otherlib/index.tlua",
      "version": "30dbc7de1f476c74bc74a1ae8005b9bb-local str: string = \"s\";\nreturn { str = str };",
      "signature": "30dbc7de1f476c74bc74a1ae8005b9bb-local str: string = \"s\";\nreturn { str = str };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../node_modules/somelib/index.tlua",
      "version": "23a32461648fa5b357b2795807b922cc-local other = require('otherlib');\nlocal foo = { foo = other.str };\nreturn { foo = foo };",
      "signature": "23a32461648fa5b357b2795807b922cc-local other = require('otherlib');\nlocal foo = { foo = other.str };\nreturn { foo = foo };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../node_modules/lib1/index.tlua",
      "version": "2a90a454e186105000c8c0193d416555-local someLib = require('someLib');\nlocal foo1: string = someLib.foo.foo;\nreturn { foo1 = foo1 };",
      "signature": "2a90a454e186105000c8c0193d416555-local someLib = require('someLib');\nlocal foo1: string = someLib.foo.foo;\nreturn { foo1 = foo1 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../node_modules/lib2/index.tlua",
      "version": "ac9f78442c66b675882139805d03b966-local someLib = require('someLib');\nlocal foo2: string = someLib.foo.foo;\nreturn { foo2 = foo2 };",
      "signature": "ac9f78442c66b675882139805d03b966-local someLib = require('someLib');\nlocal foo2: string = someLib.foo.foo;\nreturn { foo2 = foo2 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/index.tlua",
      "version": "cc6cfaa34cb010347e24399e184266ae-local lib1 = require('lib1');\nlocal lib2 = require('lib2');\nlocal someLib = require('someLib');\nlocal otherLib = require('otherlib');\nlocal foo1: string = lib1.foo1;\nlocal foo2: string = lib2.foo2;",
      "signature": "cc6cfaa34cb010347e24399e184266ae-local lib1 = require('lib1');\nlocal lib2 = require('lib2');\nlocal someLib = require('someLib');\nlocal otherLib = require('otherlib');\nlocal foo1: string = lib1.foo1;\nlocal foo2: string = lib2.foo2;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "size": 1944
}
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

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/node_modules/otherlib/index.tlua
*refresh*    /home/node_modules/someLib/index.tlua
*refresh*    /home/node_modules/lib1/index.tlua
*refresh*    /home/node_modules/lib2/index.tlua
*refresh*    /home/project/src/index.tlua
Signatures::

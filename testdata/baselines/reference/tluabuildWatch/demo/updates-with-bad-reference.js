currentDirectory::/user/username/projects/demo
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/demo/animals/animal.tlua] *new* 
type Size = "small" | "medium" | "large";
interface Animal {
    size: Size;
}
//// [/user/username/projects/demo/animals/dog.tlua] *new* 
local utilities = require('core.utilities');

interface Dog extends Animal {
    woof(): void;
    name: string;
}

function createDog(): Dog
    return ({
        size: "medium",
        woof: function()
            console.log("Woof!");
        end,
        name: utilities.makeRandomName()
    });
end

return { createDog = createDog };
//// [/user/username/projects/demo/animals/index.tlua] *new* 
local dog = require('animals.dog');

return { createDog = dog.createDog };
//// [/user/username/projects/demo/animals/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/animals",
        "rootDir": ".."
    },
    "references": [
        { "path": "../core" }
    ]
}
//// [/user/username/projects/demo/core/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/core",
        "rootDir": ".."
    },
}
//// [/user/username/projects/demo/core/utilities.tlua] *new* 
local A = require('animals.index');
A;
function makeRandomName()
    return "Bob!?! ";
end

function lastElementOf<T>(arr: T[]): T | undefined
    if arr.length == 0 then return undefined end
    return arr[arr.length - 1];
end

return { makeRandomName = makeRandomName, lastElementOf = lastElementOf };
//// [/user/username/projects/demo/tluaconfig-base.json] *new* 
{
    "compilerOptions": {
        "declaration": true,
        "target": "es5",
        "module": "commonjs",
        "strict": true,
        "noUnusedLocals": true,
        "noUnusedParameters": true,
        "noImplicitReturns": true,
        "composite": true,
    },
}
//// [/user/username/projects/demo/tluaconfig.json] *new* 
{
    "files": [],
    "references": [
        {
            "path": "./core"
        },
        {
            "path": "./animals",
        },
        {
            "path": "./zoo",
        },
    ],
}
//// [/user/username/projects/demo/zoo/tluaconfig.json] *new* 
{
    "extends": "../tluaconfig-base.json",
    "compilerOptions": {
        "outDir": "../lib/zoo",
        "rootDir": ".."
    },
    "references": [
        {
            "path": "../animals"
        }
    ]
}
//// [/user/username/projects/demo/zoo/zoo.tlua] *new* 
local animals = require('animals.index');

function createZoo(): Array<Dog>
    return {
        animals.createDog()
    };
end

return { createZoo = createZoo };

tlua --b -w --verbose
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * animals/tluaconfig.json
    * zoo/tluaconfig.json
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because output file 'lib/core/core/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96manimals/dog.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local utilities = require('core.utilities');
[7m [0m [91m~~~~~[0m

[96manimals/dog.tlua[0m:[93m10[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m10[0m         size: "medium",
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m11[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m11[0m         woof: function()
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m14[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m14[0m         name: utilities.makeRandomName()
[7m  [0m [91m            ~[0m

[96manimals/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local dog = require('animals.dog');
[7m [0m [91m~~~~~[0m

[96mcore/utilities.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local A = require('animals.index');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'animals/tluaconfig.json' is out of date because output file 'lib/animals/animals/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'animals/tluaconfig.json'...

[96manimals/animal.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m type Size = "small" | "medium" | "large";
[7m [0m [91m~~~~[0m

[96manimals/dog.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local utilities = require('core.utilities');
[7m [0m [91m~~~~~[0m

[96manimals/dog.tlua[0m:[93m10[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m10[0m         size: "medium",
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m11[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m11[0m         woof: function()
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m14[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m14[0m         name: utilities.makeRandomName()
[7m  [0m [91m            ~[0m

[96manimals/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local dog = require('animals.dog');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'zoo/tluaconfig.json' is out of date because output file 'lib/zoo/zoo/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'zoo/tluaconfig.json'...

[96mzoo/tluaconfig.json[0m:[93m3[0m:[93m5[0m - [91merror[0m[90m TLUA5108: [0mOption 'target=ES5' has been removed. Please remove it from your configuration.

[7m3[0m     "compilerOptions": {
[7m [0m [91m    ~~~~~~~~~~~~~~~~~[0m

[96mzoo/zoo.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local animals = require('animals.index');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 14 errors. Watching for file changes.

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
//// [/user/username/projects/demo/lib/animals/animals/animal.lua] *new* 

//// [/user/username/projects/demo/lib/animals/animals/dog.lua] *new* 
local utilities = require('core.utilities');
function createDog()
  return ({
    size, "medium",
    woof, function()
      console.log("Woof!");
    end,
    name, utilities.makeRandomName()
  });
end
return { createDog = createDog };

//// [/user/username/projects/demo/lib/animals/animals/index.lua] *new* 
local dog = require('animals.dog');
return { createDog = dog.createDog };

//// [/user/username/projects/demo/lib/animals/animals/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../../../animals/animal.tlua","../../../animals/dog.tlua","../../../animals/index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"5286c0bfee39cebba10388605076e53a-type Size = \"small\" | \"medium\" | \"large\";\ninterface Animal {\n    size: Size;\n}","3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };","4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };"],"options":{"composite":true,"declaration":true,"module":1,"noImplicitReturns":true,"noUnusedLocals":true,"noUnusedParameters":true,"outDir":"..","rootDir":"../../..","strict":true,"target":1},"semanticDiagnosticsPerFile":[1,2,3,4],"emitDiagnosticsPerFile":[[2,[{"end":4,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/user/username/projects/demo/lib/animals/animals/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../../animals/animal.tlua",
        "../../../animals/dog.tlua",
        "../../../animals/index.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../../animals/animal.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua"
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
      "fileName": "../../../animals/animal.tlua",
      "version": "5286c0bfee39cebba10388605076e53a-type Size = \"small\" | \"medium\" | \"large\";\ninterface Animal {\n    size: Size;\n}",
      "signature": "5286c0bfee39cebba10388605076e53a-type Size = \"small\" | \"medium\" | \"large\";\ninterface Animal {\n    size: Size;\n}",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../animals/dog.tlua",
      "version": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "signature": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../animals/index.tlua",
      "version": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "signature": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "module": 1,
    "noImplicitReturns": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "outDir": "..",
    "rootDir": "../../..",
    "strict": true,
    "target": 1
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../../../animals/animal.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "../../../animals/animal.tlua",
      [
        {
          "end": 4,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "../../../animals/dog.tlua",
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
      "../../../animals/index.tlua",
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
      "file": "../../../animals/animal.tlua",
      "original": 2
    },
    {
      "file": "../../../animals/dog.tlua",
      "original": 3
    },
    {
      "file": "../../../animals/index.tlua",
      "original": 4
    }
  ],
  "size": 2261
}
//// [/user/username/projects/demo/lib/core/animals/dog.lua] *new* 
local utilities = require('core.utilities');
function createDog()
  return ({
    size, "medium",
    woof, function()
      console.log("Woof!");
    end,
    name, utilities.makeRandomName()
  });
end
return { createDog = createDog };

//// [/user/username/projects/demo/lib/core/animals/index.lua] *new* 
local dog = require('animals.dog');
return { createDog = dog.createDog };

//// [/user/username/projects/demo/lib/core/core/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[4],"fileNames":["lib.luajit.d.tlua","../../../animals/dog.tlua","../../../animals/index.tlua","../../../core/utilities.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };","4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };","409f2459c49278d9e0cc361002ef5e9c-local A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };"],"options":{"composite":true,"declaration":true,"module":1,"noImplicitReturns":true,"noUnusedLocals":true,"noUnusedParameters":true,"outDir":"..","rootDir":"../../..","strict":true,"target":1},"semanticDiagnosticsPerFile":[1,2,3,4],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/user/username/projects/demo/lib/core/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../../core/utilities.tlua"
      ],
      "original": 4
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua",
    "../../../core/utilities.tlua"
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
      "fileName": "../../../animals/dog.tlua",
      "version": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "signature": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../animals/index.tlua",
      "version": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "signature": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../core/utilities.tlua",
      "version": "409f2459c49278d9e0cc361002ef5e9c-local A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };",
      "signature": "409f2459c49278d9e0cc361002ef5e9c-local A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "module": 1,
    "noImplicitReturns": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "outDir": "..",
    "rootDir": "../../..",
    "strict": true,
    "target": 1
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua",
    "../../../core/utilities.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "../../../animals/dog.tlua",
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
      "../../../animals/index.tlua",
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
      "../../../core/utilities.tlua",
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
      "file": "../../../animals/dog.tlua",
      "original": 2
    },
    {
      "file": "../../../animals/index.tlua",
      "original": 3
    },
    {
      "file": "../../../core/utilities.tlua",
      "original": 4
    }
  ],
  "size": 2486
}
//// [/user/username/projects/demo/lib/core/core/utilities.lua] *new* 
local A = require('animals.index');
A;
function makeRandomName()
  return "Bob!?! ";
end
function lastElementOf(arr)
  if arr.length == 0 then
    return nil;
  end
  return arr[arr.length - 1];
end
return { makeRandomName = makeRandomName, lastElementOf = lastElementOf };

//// [/user/username/projects/demo/lib/zoo/zoo/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","../../../zoo/zoo.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"61c128a7411e433590fff653485e7bd3-local animals = require('animals.index');\n\nfunction createZoo(): Array<Dog>\n    return {\n        animals.createDog()\n    };\nend\n\nreturn { createZoo = createZoo };"],"options":{"composite":true,"declaration":true,"module":1,"noImplicitReturns":true,"noUnusedLocals":true,"noUnusedParameters":true,"outDir":"..","rootDir":"../../..","strict":true,"target":1},"semanticDiagnosticsPerFile":[1,2],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/demo/lib/zoo/zoo/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../../zoo/zoo.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../../zoo/zoo.tlua"
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
      "fileName": "../../../zoo/zoo.tlua",
      "version": "61c128a7411e433590fff653485e7bd3-local animals = require('animals.index');\n\nfunction createZoo(): Array<Dog>\n    return {\n        animals.createDog()\n    };\nend\n\nreturn { createZoo = createZoo };",
      "signature": "61c128a7411e433590fff653485e7bd3-local animals = require('animals.index');\n\nfunction createZoo(): Array<Dog>\n    return {\n        animals.createDog()\n    };\nend\n\nreturn { createZoo = createZoo };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "module": 1,
    "noImplicitReturns": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "outDir": "..",
    "rootDir": "../../..",
    "strict": true,
    "target": 1
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../../../zoo/zoo.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "../../../zoo/zoo.tlua",
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
      "file": "../../../zoo/zoo.tlua",
      "original": 2
    }
  ],
  "size": 1519
}
//// [/user/username/projects/demo/lib/zoo/zoo/zoo.lua] *new* 
local animals = require('animals.index');
function createZoo()
  return {
    animals.createDog()
  };
end
return { createZoo = createZoo };


Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/demo
  /user/username/projects/demo/animals (recursive)
  /user/username/projects/demo/core (recursive)
  /user/username/projects/demo/lib
  /user/username/projects/demo/lib/animals
  /user/username/projects/demo/lib/animals/animals
  /user/username/projects/demo/lib/core
  /user/username/projects/demo/lib/core/core
  /user/username/projects/demo/zoo (recursive)
core/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/animals/dog.tlua
*not cached* /user/username/projects/demo/animals/index.tlua
*not cached* /user/username/projects/demo/core/utilities.tlua
Signatures::

animals/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/animals/animal.tlua
*not cached* /user/username/projects/demo/animals/dog.tlua
*not cached* /user/username/projects/demo/animals/index.tlua
Signatures::

zoo/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/zoo/zoo.tlua
Signatures::


Edit [0]:: Prepend a line
//// [/user/username/projects/demo/core/utilities.tlua] *modified* 

local A = require('animals.index');
A;
function makeRandomName()
    return "Bob!?! ";
end

function lastElementOf<T>(arr: T[]): T | undefined
    if arr.length == 0 then return undefined end
    return arr[arr.length - 1];
end

return { makeRandomName = makeRandomName, lastElementOf = lastElementOf };


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * core/tluaconfig.json
    * animals/tluaconfig.json
    * zoo/tluaconfig.json
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'core/tluaconfig.json' is out of date because buildinfo file 'lib/core/core/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'core/tluaconfig.json'...

[96manimals/dog.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local utilities = require('core.utilities');
[7m [0m [91m~~~~~[0m

[96manimals/dog.tlua[0m:[93m10[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m10[0m         size: "medium",
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m11[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m11[0m         woof: function()
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m14[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m14[0m         name: utilities.makeRandomName()
[7m  [0m [91m            ~[0m

[96manimals/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local dog = require('animals.dog');
[7m [0m [91m~~~~~[0m

[96mcore/utilities.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local A = require('animals.index');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'animals/tluaconfig.json' is out of date because it has errors.

[[90mHH:MM:SS AM[0m] Building project 'animals/tluaconfig.json'...

[96manimals/animal.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m type Size = "small" | "medium" | "large";
[7m [0m [91m~~~~[0m

[96manimals/dog.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local utilities = require('core.utilities');
[7m [0m [91m~~~~~[0m

[96manimals/dog.tlua[0m:[93m10[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m10[0m         size: "medium",
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m11[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m11[0m         woof: function()
[7m  [0m [91m            ~[0m

[96manimals/dog.tlua[0m:[93m14[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m14[0m         name: utilities.makeRandomName()
[7m  [0m [91m            ~[0m

[96manimals/index.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local dog = require('animals.dog');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'animals/tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Project 'zoo/tluaconfig.json' is out of date because it has errors.

[[90mHH:MM:SS AM[0m] Building project 'zoo/tluaconfig.json'...

[96mzoo/tluaconfig.json[0m:[93m3[0m:[93m5[0m - [91merror[0m[90m TLUA5108: [0mOption 'target=ES5' has been removed. Please remove it from your configuration.

[7m3[0m     "compilerOptions": {
[7m [0m [91m    ~~~~~~~~~~~~~~~~~[0m

[96mzoo/zoo.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local animals = require('animals.index');
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Updating unchanged output timestamps of project 'zoo/tluaconfig.json'...

[[90mHH:MM:SS AM[0m] Found 14 errors. Watching for file changes.

//// [/user/username/projects/demo/lib/animals/animals/tluaconfig.tluabuildinfo] *mTime changed*
//// [/user/username/projects/demo/lib/core/core/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[4],"fileNames":["lib.luajit.d.tlua","../../../animals/dog.tlua","../../../animals/index.tlua","../../../core/utilities.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };","4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",{"version":"1cf65ddd2b20157ad0977aa6285a3c91-\nlocal A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };","signature":"91677b2b73c621a5669607b4302a506f-\n(1,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"options":{"composite":true,"declaration":true,"module":1,"noImplicitReturns":true,"noUnusedLocals":true,"noUnusedParameters":true,"outDir":"..","rootDir":"../../..","strict":true,"target":1},"semanticDiagnosticsPerFile":[1,2,3,4],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":1,"end":6,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/user/username/projects/demo/lib/core/core/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../../core/utilities.tlua"
      ],
      "original": 4
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua",
    "../../../core/utilities.tlua"
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
      "fileName": "../../../animals/dog.tlua",
      "version": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "signature": "3c00568e3c7121941e7ad954dab72c1c-local utilities = require('core.utilities');\n\ninterface Dog extends Animal {\n    woof(): void;\n    name: string;\n}\n\nfunction createDog(): Dog\n    return ({\n        size: \"medium\",\n        woof: function()\n            console.log(\"Woof!\");\n        end,\n        name: utilities.makeRandomName()\n    });\nend\n\nreturn { createDog = createDog };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../animals/index.tlua",
      "version": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "signature": "4dc4b07932e5707398b3ca3d70ea4a72-local dog = require('animals.dog');\n\nreturn { createDog = dog.createDog };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../../core/utilities.tlua",
      "version": "1cf65ddd2b20157ad0977aa6285a3c91-\nlocal A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };",
      "signature": "91677b2b73c621a5669607b4302a506f-\n(1,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "1cf65ddd2b20157ad0977aa6285a3c91-\nlocal A = require('animals.index');\nA;\nfunction makeRandomName()\n    return \"Bob!?! \";\nend\n\nfunction lastElementOf<T>(arr: T[]): T | undefined\n    if arr.length == 0 then return undefined end\n    return arr[arr.length - 1];\nend\n\nreturn { makeRandomName = makeRandomName, lastElementOf = lastElementOf };",
        "signature": "91677b2b73c621a5669607b4302a506f-\n(1,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true,
    "declaration": true,
    "module": 1,
    "noImplicitReturns": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "outDir": "..",
    "rootDir": "../../..",
    "strict": true,
    "target": 1
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../../../animals/dog.tlua",
    "../../../animals/index.tlua",
    "../../../core/utilities.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "../../../animals/dog.tlua",
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
      "../../../animals/index.tlua",
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
      "../../../core/utilities.tlua",
      [
        {
          "pos": 1,
          "end": 6,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "../../../animals/dog.tlua",
      "original": 2
    },
    {
      "file": "../../../animals/index.tlua",
      "original": 3
    },
    {
      "file": "../../../core/utilities.tlua",
      "original": 4
    }
  ],
  "size": 2663
}
//// [/user/username/projects/demo/lib/core/core/utilities.lua] *rewrite with same content*
//// [/user/username/projects/demo/lib/zoo/zoo/tluaconfig.tluabuildinfo] *mTime changed*

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/demo
  /user/username/projects/demo/animals (recursive)
  /user/username/projects/demo/core (recursive)
  /user/username/projects/demo/lib
  /user/username/projects/demo/lib/animals
  /user/username/projects/demo/lib/animals/animals
  /user/username/projects/demo/lib/core
  /user/username/projects/demo/lib/core/core
  /user/username/projects/demo/zoo (recursive)
core/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/animals/dog.tlua
*not cached* /user/username/projects/demo/animals/index.tlua
*not cached* /user/username/projects/demo/core/utilities.tlua
Signatures::
(computed .d.ts) /user/username/projects/demo/core/utilities.tlua

animals/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/animals/animal.tlua
*not cached* /user/username/projects/demo/animals/dog.tlua
*not cached* /user/username/projects/demo/animals/index.tlua
Signatures::

zoo/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/demo/zoo/zoo.tlua
Signatures::

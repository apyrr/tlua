currentDirectory::/user/username/projects/sample1
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/sample1/App/app.tlua] *new* 
local library = require("Library.library");
library.createSomeObject().message;
//// [/user/username/projects/sample1/App/tluaconfig.json] *new* 
{ 
    "compilerOptions": { "rootDir": ".." },
    "references": [{ "path": "../Library" }]

}
//// [/user/username/projects/sample1/Library/library.tlua] *new* 
interface SomeObject
{
    message: string;
}

function createSomeObject(): SomeObject
    return {
        message = "new Object"
    };
end

return { createSomeObject = createSomeObject };
//// [/user/username/projects/sample1/Library/tluaconfig.json] *new* 
{ 
    "compilerOptions": {
        "composite": true
    }
}

tlua -b -w App
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96mLibrary/library.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface SomeObject
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

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
//// [/user/username/projects/sample1/App/app.lua] *new* 
local library = require("Library.library");
library.createSomeObject().message;

//// [/user/username/projects/sample1/App/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":["./app.tlua"]}
//// [/user/username/projects/sample1/App/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./app.tlua"
      ],
      "original": "./app.tlua"
    }
  ],
  "size": 49
}
//// [/user/username/projects/sample1/Library/library.lua] *new* 
function createSomeObject()
    return {
        message = "new Object"
    };
end
return { createSomeObject = createSomeObject };

//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./library.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };"],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./library.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./library.tlua"
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
      "fileName": "./library.tlua",
      "version": "6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
      "signature": "6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./library.tlua",
      [
        {
          "end": 9,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./library.tlua",
      "original": 2
    }
  ],
  "size": 1347
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/sample1
  /user/username/projects/sample1/App (recursive)
  /user/username/projects/sample1/Library (recursive)
Library/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/Library/library.tlua
Signatures::

App/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/sample1/App/app.tlua
Signatures::


Edit [0]:: Introduce error
//// [/user/username/projects/sample1/Library/library.tlua] *modified* 
interface SomeObject
{
    message2: string;
}

function createSomeObject(): SomeObject
    return {
        message2 = "new Object"
    };
end

return { createSomeObject = createSomeObject };


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96mLibrary/library.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface SomeObject
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/user/username/projects/sample1/App/app.lua] *mTime changed*
//// [/user/username/projects/sample1/App/tluaconfig.tluabuildinfo] *mTime changed*
//// [/user/username/projects/sample1/Library/library.lua] *modified* 
function createSomeObject()
    return {
        message2 = "new Object"
    };
end
return { createSomeObject = createSomeObject };

//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./library.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"3978d0a8d27c910e6bfeffa32aa12b43-interface SomeObject\n{\n    message2: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message2 = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };","signature":"fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./library.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./library.tlua"
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
      "fileName": "./library.tlua",
      "version": "3978d0a8d27c910e6bfeffa32aa12b43-interface SomeObject\n{\n    message2: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message2 = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
      "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "3978d0a8d27c910e6bfeffa32aa12b43-interface SomeObject\n{\n    message2: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message2 = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
        "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./library.tlua",
      [
        {
          "end": 9,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./library.tlua",
      "original": 2
    }
  ],
  "size": 1516
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/sample1
  /user/username/projects/sample1/App (recursive)
  /user/username/projects/sample1/Library (recursive)
Library/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /user/username/projects/sample1/Library/library.tlua
Signatures::
(computed .d.ts) /user/username/projects/sample1/Library/library.tlua


Edit [1]:: Fix error
//// [/user/username/projects/sample1/Library/library.tlua] *modified* 
interface SomeObject
{
    message: string;
}

function createSomeObject(): SomeObject
    return {
        message = "new Object"
    };
end

return { createSomeObject = createSomeObject };


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96mLibrary/library.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface SomeObject
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

//// [/user/username/projects/sample1/App/app.lua] *mTime changed*
//// [/user/username/projects/sample1/App/tluaconfig.tluabuildinfo] *mTime changed*
//// [/user/username/projects/sample1/Library/library.lua] *modified* 
function createSomeObject()
    return {
        message = "new Object"
    };
end
return { createSomeObject = createSomeObject };

//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./library.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };","signature":"fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/user/username/projects/sample1/Library/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./library.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./library.tlua"
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
      "fileName": "./library.tlua",
      "version": "6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
      "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "6d32714bf51262ad283c5b2457db6797-interface SomeObject\n{\n    message: string;\n}\n\nfunction createSomeObject(): SomeObject\n    return {\n        message = \"new Object\"\n    };\nend\n\nreturn { createSomeObject = createSomeObject };",
        "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./library.tlua",
      [
        {
          "end": 9,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./library.tlua",
      "original": 2
    }
  ],
  "size": 1514
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/sample1
  /user/username/projects/sample1/App (recursive)
  /user/username/projects/sample1/Library (recursive)
Library/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /user/username/projects/sample1/Library/library.tlua
Signatures::
(computed .d.ts) /user/username/projects/sample1/Library/library.tlua

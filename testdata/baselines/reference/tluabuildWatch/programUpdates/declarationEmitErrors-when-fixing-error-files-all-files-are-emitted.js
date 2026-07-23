currentDirectory::/user/username/projects/solution
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/solution/app/fileWithError.tlua] *new* 
local myClassWithError = {
    tags: function() end,
    p: 12
};
//// [/user/username/projects/solution/app/fileWithoutError.tlua] *new* 
local myClass = { };
//// [/user/username/projects/solution/app/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true
    }
}

tlua -b -w app
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96mapp/fileWithError.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClassWithError = {
[7m [0m [91m~~~~~[0m

[96mapp/fileWithError.tlua[0m:[93m2[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     tags: function() end,
[7m [0m [91m        ~[0m

[96mapp/fileWithError.tlua[0m:[93m3[0m:[93m6[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m     p: 12
[7m [0m [91m     ~[0m

[96mapp/fileWithoutError.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClass = { };
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 4 errors. Watching for file changes.

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
//// [/user/username/projects/solution/app/fileWithError.lua] *new* 
local myClassWithError = {
    tags, function()
    end,
    p, 12
};

//// [/user/username/projects/solution/app/fileWithoutError.lua] *new* 
local myClass = {};

//// [/user/username/projects/solution/app/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./fileWithError.tlua","./fileWithoutError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"0353d2c16b2cbe0eeb55ba7d84752db6-local myClassWithError = {\n    tags: function() end,\n    p: 12\n};","0af062a2144241ae15d557868afc35e7-local myClass = { };"],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/solution/app/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./fileWithError.tlua",
        "./fileWithoutError.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./fileWithError.tlua",
    "./fileWithoutError.tlua"
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
      "fileName": "./fileWithError.tlua",
      "version": "0353d2c16b2cbe0eeb55ba7d84752db6-local myClassWithError = {\n    tags: function() end,\n    p: 12\n};",
      "signature": "0353d2c16b2cbe0eeb55ba7d84752db6-local myClassWithError = {\n    tags: function() end,\n    p: 12\n};",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./fileWithoutError.tlua",
      "version": "0af062a2144241ae15d557868afc35e7-local myClass = { };",
      "signature": "0af062a2144241ae15d557868afc35e7-local myClass = { };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./fileWithError.tlua",
    "./fileWithoutError.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./fileWithError.tlua",
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
      "./fileWithoutError.tlua",
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
      "file": "./fileWithError.tlua",
      "original": 2
    },
    {
      "file": "./fileWithoutError.tlua",
      "original": 3
    }
  ],
  "size": 1463
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/solution
  /user/username/projects/solution/app (recursive)
app/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/solution/app/fileWithError.tlua
*not cached* /user/username/projects/solution/app/fileWithoutError.tlua
Signatures::


Edit [0]:: Fix error
//// [/user/username/projects/solution/app/fileWithError.tlua] *modified* 
local myClassWithError = {
    tags: function() end,
    
};


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[96mapp/fileWithError.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClassWithError = {
[7m [0m [91m~~~~~[0m

[96mapp/fileWithError.tlua[0m:[93m2[0m:[93m9[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     tags: function() end,
[7m [0m [91m        ~[0m

[96mapp/fileWithoutError.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local myClass = { };
[7m [0m [91m~~~~~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

//// [/user/username/projects/solution/app/fileWithError.lua] *modified* 
local myClassWithError = {
    tags, function()
    end,
};

//// [/user/username/projects/solution/app/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./fileWithError.tlua","./fileWithoutError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"1a9b4f87d2b4f4f628e000945bb18858-local myClassWithError = {\n    tags: function() end,\n    \n};","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"0af062a2144241ae15d557868afc35e7-local myClass = { };"],"options":{"composite":true},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/solution/app/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./fileWithError.tlua",
        "./fileWithoutError.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./fileWithError.tlua",
    "./fileWithoutError.tlua"
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
      "fileName": "./fileWithError.tlua",
      "version": "1a9b4f87d2b4f4f628e000945bb18858-local myClassWithError = {\n    tags: function() end,\n    \n};",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "1a9b4f87d2b4f4f628e000945bb18858-local myClassWithError = {\n    tags: function() end,\n    \n};",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./fileWithoutError.tlua",
      "version": "0af062a2144241ae15d557868afc35e7-local myClass = { };",
      "signature": "0af062a2144241ae15d557868afc35e7-local myClass = { };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./fileWithError.tlua",
    "./fileWithoutError.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./fileWithError.tlua",
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
      "./fileWithoutError.tlua",
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
      "file": "./fileWithError.tlua",
      "original": 2
    },
    {
      "file": "./fileWithoutError.tlua",
      "original": 3
    }
  ],
  "size": 1625
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/solution
  /user/username/projects/solution/app (recursive)
app/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/solution/app/fileWithError.tlua
*not cached* /user/username/projects/solution/app/fileWithoutError.tlua
Signatures::
(computed .d.ts) /user/username/projects/solution/app/fileWithError.tlua

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/globals.d.tlua] *new* 
interface SymbolConstructor {
    (description?: string | number): symbol;
}
declare Symbol: SymbolConstructor;
//// [/home/src/workspaces/project/src/hkt.tlua] *new* 
interface HKT<T> { }
//// [/home/src/workspaces/project/src/main.tlua] *new* 
local sym = Symbol();

interface Local<T> extends HKT<T> { }
interface Local<T> {
    [sym]: { a: T }
}
local x = 10;
type A = Local<number>[typeof sym];
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "rootDir": "src",
        "incremental": true,
    },
}

tlua --b --verbose
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m18[0m declare Symbol: SymbolConstructor;
[7m  [0m [91m        ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m19[0m interface Symbol {
[7m  [0m [91m          ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m4[0m declare Symbol: SymbolConstructor;
[7m [0m [91m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m18[0m declare Symbol: SymbolConstructor;
    [7m  [0m [96m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - and here.
    [7m19[0m interface Symbol {
    [7m  [0m [96m          ~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  ../../tslibs/TS/Lib/lib.luajit.d.tlua[90m:18[0m
     1  src/globals.d.tlua[90m:4[0m

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
//// [/home/src/workspaces/project/src/hkt.lua] *new* 

//// [/home/src/workspaces/project/src/main.lua] *new* 
local sym = Symbol();
local x = 10;

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/globals.d.tlua","./src/hkt.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;","affectsGlobalScope":true,"impliedNodeFormat":1},"a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }","94f5613dbcaa4a0b9ca12c25f01836ae-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\nlocal x = 10;\ntype A = Local<number>[typeof sym];"],"options":{"rootDir":"./src"},"semanticDiagnosticsPerFile":[[1,[{"pos":508,"end":514,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]},{"pos":545,"end":551,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]}]],[2,[{"pos":85,"end":91,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":1,"pos":508,"end":514,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]},{"file":1,"pos":545,"end":551,"code":6204,"category":3,"messageKey":"and_here_6204"}]}]]]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/globals.d.tlua",
        "./src/hkt.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/globals.d.tlua",
    "./src/hkt.tlua",
    "./src/main.tlua"
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
      "fileName": "./src/globals.d.tlua",
      "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "signature": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/hkt.tlua",
      "version": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "signature": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "94f5613dbcaa4a0b9ca12c25f01836ae-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\nlocal x = 10;\ntype A = Local<number>[typeof sym];",
      "signature": "94f5613dbcaa4a0b9ca12c25f01836ae-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\nlocal x = 10;\ntype A = Local<number>[typeof sym];",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "rootDir": "./src"
  },
  "semanticDiagnosticsPerFile": [
    [
      "lib.luajit.d.tlua",
      [
        {
          "pos": 508,
          "end": 514,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        },
        {
          "pos": 545,
          "end": 551,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./src/globals.d.tlua",
      [
        {
          "pos": 85,
          "end": 91,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "lib.luajit.d.tlua",
              "pos": 508,
              "end": 514,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            },
            {
              "file": "lib.luajit.d.tlua",
              "pos": 545,
              "end": 551,
              "code": 6204,
              "category": 3,
              "messageKey": "and_here_6204"
            }
          ]
        }
      ]
    ]
  ],
  "size": 2419
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/globals.d.tlua
*refresh*    /home/src/workspaces/project/src/hkt.tlua
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::


Edit [0]:: incremental-declaration-doesnt-change
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
local sym = Symbol();

interface Local<T> extends HKT<T> { }
interface Local<T> {
    [sym]: { a: T }
}

type A = Local<number>[typeof sym];

tlua --b --verbose
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m18[0m declare Symbol: SymbolConstructor;
[7m  [0m [91m        ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m19[0m interface Symbol {
[7m  [0m [91m          ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m4[0m declare Symbol: SymbolConstructor;
[7m [0m [91m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m18[0m declare Symbol: SymbolConstructor;
    [7m  [0m [96m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - and here.
    [7m19[0m interface Symbol {
    [7m  [0m [96m          ~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  ../../tslibs/TS/Lib/lib.luajit.d.tlua[90m:18[0m
     1  src/globals.d.tlua[90m:4[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
local sym = Symbol();

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/globals.d.tlua","./src/hkt.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;","affectsGlobalScope":true,"impliedNodeFormat":1},"a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",{"version":"b20120122ecde67ad93d7f788654e8a8-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"options":{"rootDir":"./src"},"semanticDiagnosticsPerFile":[[1,[{"pos":508,"end":514,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]},{"pos":545,"end":551,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]}]],[2,[{"pos":85,"end":91,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":1,"pos":508,"end":514,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]},{"file":1,"pos":545,"end":551,"code":6204,"category":3,"messageKey":"and_here_6204"}]}]]]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/globals.d.tlua",
        "./src/hkt.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/globals.d.tlua",
    "./src/hkt.tlua",
    "./src/main.tlua"
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
      "fileName": "./src/globals.d.tlua",
      "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "signature": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/hkt.tlua",
      "version": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "signature": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "b20120122ecde67ad93d7f788654e8a8-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "b20120122ecde67ad93d7f788654e8a8-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "rootDir": "./src"
  },
  "semanticDiagnosticsPerFile": [
    [
      "lib.luajit.d.tlua",
      [
        {
          "pos": 508,
          "end": 514,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        },
        {
          "pos": 545,
          "end": 551,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./src/globals.d.tlua",
      [
        {
          "pos": 85,
          "end": 91,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "lib.luajit.d.tlua",
              "pos": 508,
              "end": 514,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            },
            {
              "file": "lib.luajit.d.tlua",
              "pos": 545,
              "end": 551,
              "code": 6204,
              "category": 3,
              "messageKey": "and_here_6204"
            }
          ]
        }
      ]
    ]
  ],
  "size": 2573
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [1]:: incremental-declaration-doesnt-change
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
local sym = Symbol();

interface Local<T> extends HKT<T> { }
interface Local<T> {
    [sym]: { a: T }
}

type A = Local<number>[typeof sym];local x = 10;

tlua --b --verbose
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m18[0m declare Symbol: SymbolConstructor;
[7m  [0m [91m        ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m19[0m interface Symbol {
[7m  [0m [91m          ~~~~~~[0m

  [96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m4[0m declare Symbol: SymbolConstructor;
    [7m [0m [96m        ~~~~~~[0m

[96msrc/globals.d.tlua[0m:[93m4[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'Symbol'.

[7m4[0m declare Symbol: SymbolConstructor;
[7m [0m [91m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m18[0m:[93m9[0m - 'Symbol' was also declared here.
    [7m18[0m declare Symbol: SymbolConstructor;
    [7m  [0m [96m        ~~~~~~[0m

  [96m../../tslibs/TS/Lib/lib.luajit.d.tlua[0m:[93m19[0m:[93m11[0m - and here.
    [7m19[0m interface Symbol {
    [7m  [0m [96m          ~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  ../../tslibs/TS/Lib/lib.luajit.d.tlua[90m:18[0m
     1  src/globals.d.tlua[90m:4[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
local sym = Symbol();
local x = 10;

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/globals.d.tlua","./src/hkt.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;","affectsGlobalScope":true,"impliedNodeFormat":1},"a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",{"version":"4360ec35d3286331c70b95f66fde1568-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];local x = 10;","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"options":{"rootDir":"./src"},"semanticDiagnosticsPerFile":[[1,[{"pos":508,"end":514,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]},{"pos":545,"end":551,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":2,"pos":85,"end":91,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]}]}]],[2,[{"pos":85,"end":91,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["Symbol"],"relatedInformation":[{"file":1,"pos":508,"end":514,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["Symbol"]},{"file":1,"pos":545,"end":551,"code":6204,"category":3,"messageKey":"and_here_6204"}]}]]]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/globals.d.tlua",
        "./src/hkt.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/globals.d.tlua",
    "./src/hkt.tlua",
    "./src/main.tlua"
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
      "fileName": "./src/globals.d.tlua",
      "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "signature": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "621aa8434d5a2c7ab3d602ad159d3a54-interface SymbolConstructor {\n    (description?: string | number): symbol;\n}\ndeclare Symbol: SymbolConstructor;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/hkt.tlua",
      "version": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "signature": "a871646a5ec28b3fc04b8a880e0dd3e3-interface HKT<T> { }",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "4360ec35d3286331c70b95f66fde1568-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];local x = 10;",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "4360ec35d3286331c70b95f66fde1568-local sym = Symbol();\n\ninterface Local<T> extends HKT<T> { }\ninterface Local<T> {\n    [sym]: { a: T }\n}\n\ntype A = Local<number>[typeof sym];local x = 10;",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "options": {
    "rootDir": "./src"
  },
  "semanticDiagnosticsPerFile": [
    [
      "lib.luajit.d.tlua",
      [
        {
          "pos": 508,
          "end": 514,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        },
        {
          "pos": 545,
          "end": 551,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "./src/globals.d.tlua",
              "pos": 85,
              "end": 91,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./src/globals.d.tlua",
      [
        {
          "pos": 85,
          "end": 91,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "Symbol"
          ],
          "relatedInformation": [
            {
              "file": "lib.luajit.d.tlua",
              "pos": 508,
              "end": 514,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "Symbol"
              ]
            },
            {
              "file": "lib.luajit.d.tlua",
              "pos": 545,
              "end": 551,
              "code": 6204,
              "category": 3,
              "messageKey": "and_here_6204"
            }
          ]
        }
      ]
    ]
  ],
  "size": 2586
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua

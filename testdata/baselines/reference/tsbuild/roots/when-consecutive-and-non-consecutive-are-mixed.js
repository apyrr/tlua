currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/anotherNonConsecutive.tlua] *new* 
/// <reference path="./random2.d.tlua"/>
local nonConsecutive = "hello";
//// [/home/src/workspaces/project/asArray1.tlua] *new* 
/// <reference path="./random1.d.tlua"/>
local x = "hello";
//// [/home/src/workspaces/project/asArray2.tlua] *new* 
local x = "hello";
//// [/home/src/workspaces/project/asArray3.tlua] *new* 
local x = "hello";
//// [/home/src/workspaces/project/file1.tlua] *new* 
local x = "hello";
//// [/home/src/workspaces/project/file2.tlua] *new* 
local y = "world";
//// [/home/src/workspaces/project/nonconsecutive.tlua] *new* 
/// <reference path="./random.d.tlua"/>
    local nonConsecutive = "hello";
//// [/home/src/workspaces/project/random.d.tlua] *new* 
declare random: string;
//// [/home/src/workspaces/project/random1.d.tlua] *new* 
declare random: string;
//// [/home/src/workspaces/project/random2.d.tlua] *new* 
declare random: string;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": { "composite": true },
    "include": ["file*.tlua", "nonconsecutive*.tlua", "asArray*.tlua", "anotherNonConsecutive.tlua"],
}

tlua --b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because output file 'tsconfig.tsbuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96manotherNonConsecutive.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local nonConsecutive = "hello";
[7m [0m [91m~~~~~[0m

[96masArray1.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96masArray2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96masArray3.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96mfile1.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96mfile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = "world";
[7m [0m [91m~~~~~[0m

[96mnonconsecutive.tlua[0m:[93m2[0m:[93m5[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m     local nonConsecutive = "hello";
[7m [0m [91m    ~~~~~[0m

[96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom1.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

  [96mrandom2.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

[96mrandom1.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

[96mrandom2.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m


Found 10 errors in 10 files.

Errors  Files
     1  anotherNonConsecutive.tlua[90m:2[0m
     1  asArray1.tlua[90m:2[0m
     1  asArray2.tlua[90m:1[0m
     1  asArray3.tlua[90m:1[0m
     1  file1.tlua[90m:1[0m
     1  file2.tlua[90m:1[0m
     1  nonconsecutive.tlua[90m:2[0m
     1  random.d.tlua[90m:1[0m
     1  random1.d.tlua[90m:1[0m
     1  random2.d.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/anotherNonConsecutive.lua] *new* 
-- / <reference path="./random2.d.tlua"/>
local nonConsecutive = "hello";

//// [/home/src/workspaces/project/asArray1.lua] *new* 
-- / <reference path="./random1.d.tlua"/>
local x = "hello";

//// [/home/src/workspaces/project/asArray2.lua] *new* 
local x = "hello";

//// [/home/src/workspaces/project/asArray3.lua] *new* 
local x = "hello";

//// [/home/src/workspaces/project/file1.lua] *new* 
local x = "hello";

//// [/home/src/workspaces/project/file2.lua] *new* 
local y = "world";

//// [/home/src/workspaces/project/nonconsecutive.lua] *new* 
-- / <reference path="./random.d.tlua"/>
local nonConsecutive = "hello";

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3],5,[7,9],11],"fileNames":["lib.luajit.d.tlua","./file1.tlua","./file2.tlua","./random.d.tlua","./nonconsecutive.tlua","./random1.d.tlua","./asArray1.tlua","./asArray2.tlua","./asArray3.tlua","./random2.d.tlua","./anotherNonConsecutive.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";","c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";","b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";","b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";"],"fileIdsList":[[10],[6],[4]],"options":{"composite":true},"referencedMap":[[11,1],[7,2],[5,3]],"semanticDiagnosticsPerFile":[[4,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":6,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]},{"file":10,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]],[6,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":4,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]],[10,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":4,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]]],"emitDiagnosticsPerFile":[[11,[{"pos":41,"end":46,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[7,[{"pos":41,"end":46,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[8,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[9,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"pos":44,"end":49,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,5,7,8,9,11]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./file1.tlua",
        "./file2.tlua"
      ],
      "original": [
        2,
        3
      ]
    },
    {
      "files": [
        "./nonconsecutive.tlua"
      ],
      "original": 5
    },
    {
      "files": [
        "./asArray1.tlua",
        "./asArray2.tlua",
        "./asArray3.tlua"
      ],
      "original": [
        7,
        9
      ]
    },
    {
      "files": [
        "./anotherNonConsecutive.tlua"
      ],
      "original": 11
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./file1.tlua",
    "./file2.tlua",
    "./random.d.tlua",
    "./nonconsecutive.tlua",
    "./random1.d.tlua",
    "./asArray1.tlua",
    "./asArray2.tlua",
    "./asArray3.tlua",
    "./random2.d.tlua",
    "./anotherNonConsecutive.tlua"
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
      "fileName": "./file1.tlua",
      "version": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "signature": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./file2.tlua",
      "version": "c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",
      "signature": "c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./nonconsecutive.tlua",
      "version": "851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",
      "signature": "851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random1.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./asArray1.tlua",
      "version": "d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";",
      "signature": "d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./asArray2.tlua",
      "version": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "signature": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./asArray3.tlua",
      "version": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "signature": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random2.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./anotherNonConsecutive.tlua",
      "version": "01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";",
      "signature": "01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "fileIdsList": [
    [
      "./random2.d.tlua"
    ],
    [
      "./random1.d.tlua"
    ],
    [
      "./random.d.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./anotherNonConsecutive.tlua": [
      "./random2.d.tlua"
    ],
    "./asArray1.tlua": [
      "./random1.d.tlua"
    ],
    "./nonconsecutive.tlua": [
      "./random.d.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./random.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random1.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            },
            {
              "file": "./random2.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./random1.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./random2.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./anotherNonConsecutive.tlua",
      [
        {
          "pos": 41,
          "end": 46,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./asArray1.tlua",
      [
        {
          "pos": 41,
          "end": 46,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./asArray2.tlua",
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
      "./asArray3.tlua",
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
      "./file1.tlua",
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
      "./file2.tlua",
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
      "./nonconsecutive.tlua",
      [
        {
          "pos": 44,
          "end": 49,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./file1.tlua",
      "original": 2
    },
    {
      "file": "./file2.tlua",
      "original": 3
    },
    {
      "file": "./nonconsecutive.tlua",
      "original": 5
    },
    {
      "file": "./asArray1.tlua",
      "original": 7
    },
    {
      "file": "./asArray2.tlua",
      "original": 8
    },
    {
      "file": "./asArray3.tlua",
      "original": 9
    },
    {
      "file": "./anotherNonConsecutive.tlua",
      "original": 11
    }
  ],
  "size": 4047
}

tsconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/file1.tlua
*refresh*    /home/src/workspaces/project/file2.tlua
*refresh*    /home/src/workspaces/project/random.d.tlua
*refresh*    /home/src/workspaces/project/nonconsecutive.tlua
*refresh*    /home/src/workspaces/project/random1.d.tlua
*refresh*    /home/src/workspaces/project/asArray1.tlua
*refresh*    /home/src/workspaces/project/asArray2.tlua
*refresh*    /home/src/workspaces/project/asArray3.tlua
*refresh*    /home/src/workspaces/project/random2.d.tlua
*refresh*    /home/src/workspaces/project/anotherNonConsecutive.tlua
Signatures::


Edit [0]:: delete file1
//// [/home/src/workspaces/project/file1.lua] *deleted*
//// [/home/src/workspaces/project/file1.tlua] *deleted*

tlua --b -v
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tsconfig.json

[[90mHH:MM:SS AM[0m] Project 'tsconfig.json' is out of date because buildinfo file 'tsconfig.tsbuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tsconfig.json'...

[96manotherNonConsecutive.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local nonConsecutive = "hello";
[7m [0m [91m~~~~~[0m

[96masArray1.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96masArray2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96masArray3.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local x = "hello";
[7m [0m [91m~~~~~[0m

[96mfile2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local y = "world";
[7m [0m [91m~~~~~[0m

[96mnonconsecutive.tlua[0m:[93m2[0m:[93m5[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m     local nonConsecutive = "hello";
[7m [0m [91m    ~~~~~[0m

[96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom1.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

  [96mrandom2.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

[96mrandom1.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m

[96mrandom2.d.tlua[0m:[93m1[0m:[93m9[0m - [91merror[0m[90m TS2451: [0mCannot redeclare block-scoped variable 'random'.

[7m1[0m declare random: string;
[7m [0m [91m        ~~~~~~[0m

  [96mrandom.d.tlua[0m:[93m1[0m:[93m9[0m - 'random' was also declared here.
    [7m1[0m declare random: string;
    [7m [0m [96m        ~~~~~~[0m


Found 9 errors in 9 files.

Errors  Files
     1  anotherNonConsecutive.tlua[90m:2[0m
     1  asArray1.tlua[90m:2[0m
     1  asArray2.tlua[90m:1[0m
     1  asArray3.tlua[90m:1[0m
     1  file2.tlua[90m:1[0m
     1  nonconsecutive.tlua[90m:2[0m
     1  random.d.tlua[90m:1[0m
     1  random1.d.tlua[90m:1[0m
     1  random2.d.tlua[90m:1[0m

//// [/home/src/workspaces/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[2,4,[6,8],10],"fileNames":["lib.luajit.d.tlua","./file2.tlua","./random.d.tlua","./nonconsecutive.tlua","./random1.d.tlua","./asArray1.tlua","./asArray2.tlua","./asArray3.tlua","./random2.d.tlua","./anotherNonConsecutive.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";","b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";","b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",{"version":"9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;","affectsGlobalScope":true,"impliedNodeFormat":1},"01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";"],"fileIdsList":[[9],[5],[3]],"options":{"composite":true},"referencedMap":[[10,1],[6,2],[4,3]],"semanticDiagnosticsPerFile":[[3,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":5,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]},{"file":9,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]],[5,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":3,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]],[9,[{"pos":8,"end":14,"code":2451,"category":1,"messageKey":"Cannot_redeclare_block_scoped_variable_0_2451","messageArgs":["random"],"relatedInformation":[{"file":3,"pos":8,"end":14,"code":6203,"category":3,"messageKey":"_0_was_also_declared_here_6203","messageArgs":["random"]}]}]]],"emitDiagnosticsPerFile":[[10,[{"pos":41,"end":46,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[6,[{"pos":41,"end":46,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[7,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[8,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":44,"end":49,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,4,6,7,8,10]}
//// [/home/src/workspaces/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./file2.tlua"
      ],
      "original": 2
    },
    {
      "files": [
        "./nonconsecutive.tlua"
      ],
      "original": 4
    },
    {
      "files": [
        "./asArray1.tlua",
        "./asArray2.tlua",
        "./asArray3.tlua"
      ],
      "original": [
        6,
        8
      ]
    },
    {
      "files": [
        "./anotherNonConsecutive.tlua"
      ],
      "original": 10
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./file2.tlua",
    "./random.d.tlua",
    "./nonconsecutive.tlua",
    "./random1.d.tlua",
    "./asArray1.tlua",
    "./asArray2.tlua",
    "./asArray3.tlua",
    "./random2.d.tlua",
    "./anotherNonConsecutive.tlua"
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
      "fileName": "./file2.tlua",
      "version": "c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",
      "signature": "c101b9fa0a877d34b841097ccc37ca35-local y = \"world\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./nonconsecutive.tlua",
      "version": "851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",
      "signature": "851d0f1479a2dac9bfe52a3cacabdb2d-/// <reference path=\"./random.d.tlua\"/>\n    local nonConsecutive = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random1.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./asArray1.tlua",
      "version": "d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";",
      "signature": "d8e6a8aaa21150d655fd3ef9b9079292-/// <reference path=\"./random1.d.tlua\"/>\nlocal x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./asArray2.tlua",
      "version": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "signature": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./asArray3.tlua",
      "version": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "signature": "b5bcf31a9fd0c61377b93a61965051b5-local x = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./random2.d.tlua",
      "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "signature": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
      "affectsGlobalScope": true,
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9e5f1d474c6df35bb6d166629c97c5cc-declare random: string;",
        "affectsGlobalScope": true,
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./anotherNonConsecutive.tlua",
      "version": "01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";",
      "signature": "01fdf04470957c26966e218904bece71-/// <reference path=\"./random2.d.tlua\"/>\nlocal nonConsecutive = \"hello\";",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "fileIdsList": [
    [
      "./random2.d.tlua"
    ],
    [
      "./random1.d.tlua"
    ],
    [
      "./random.d.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./anotherNonConsecutive.tlua": [
      "./random2.d.tlua"
    ],
    "./asArray1.tlua": [
      "./random1.d.tlua"
    ],
    "./nonconsecutive.tlua": [
      "./random.d.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    [
      "./random.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random1.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            },
            {
              "file": "./random2.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./random1.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ],
    [
      "./random2.d.tlua",
      [
        {
          "pos": 8,
          "end": 14,
          "code": 2451,
          "category": 1,
          "messageKey": "Cannot_redeclare_block_scoped_variable_0_2451",
          "messageArgs": [
            "random"
          ],
          "relatedInformation": [
            {
              "file": "./random.d.tlua",
              "pos": 8,
              "end": 14,
              "code": 6203,
              "category": 3,
              "messageKey": "_0_was_also_declared_here_6203",
              "messageArgs": [
                "random"
              ]
            }
          ]
        }
      ]
    ]
  ],
  "emitDiagnosticsPerFile": [
    [
      "./anotherNonConsecutive.tlua",
      [
        {
          "pos": 41,
          "end": 46,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./asArray1.tlua",
      [
        {
          "pos": 41,
          "end": 46,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./asArray2.tlua",
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
      "./asArray3.tlua",
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
      "./file2.tlua",
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
      "./nonconsecutive.tlua",
      [
        {
          "pos": 44,
          "end": 49,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./file2.tlua",
      "original": 2
    },
    {
      "file": "./nonconsecutive.tlua",
      "original": 4
    },
    {
      "file": "./asArray1.tlua",
      "original": 6
    },
    {
      "file": "./asArray2.tlua",
      "original": 7
    },
    {
      "file": "./asArray3.tlua",
      "original": 8
    },
    {
      "file": "./anotherNonConsecutive.tlua",
      "original": 10
    }
  ],
  "size": 3847
}

tsconfig.json::
SemanticDiagnostics::
Signatures::

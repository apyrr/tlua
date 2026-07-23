currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua] *new* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function anotherFileWithSameReferenes() end
//// [/home/src/workspaces/project/src/filePresent.tlua] *new* 
function something() return 10; end
//// [/home/src/workspaces/project/src/main.tlua] *new* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() end
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": { "composite": true },
    "include": ["src/**/*.tlua"],
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TLUA6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TLUA6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() end
[7m [0m [91m~~~~~~~~[0m


Found 5 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     2  src/main.tlua[90m:2[0m

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
//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.lua] *new* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function anotherFileWithSameReferenes()
end

//// [/home/src/workspaces/project/src/filePresent.lua] *new* 
function something()
    return 10;
end

//// [/home/src/workspaces/project/src/main.lua] *new* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main()
end

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end","546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end","6bbf107228b0793d1141673a96774ff9-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end"],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
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
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "6bbf107228b0793d1141673a96774ff9-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end",
      "signature": "6bbf107228b0793d1141673a96774ff9-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 1913
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/filePresent.tlua
*refresh*    /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
*refresh*    /home/src/workspaces/project/src/main.tlua
Signatures::


Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TLUA6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m2[0m:[93m22[0m - [91merror[0m[90m TLUA6053: [0mFile './fileNotFound.tlua' not found.

[7m2[0m /// <reference path="./fileNotFound.tlua"/>
[7m [0m [91m                     ~~~~~~~~~~~~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() end
[7m [0m [91m~~~~~~~~[0m


Found 5 errors in 3 files.

Errors  Files
     2  src/anotherFileWithSameReferenes.tlua[90m:2[0m
     1  src/filePresent.tlua[90m:1[0m
     2  src/main.tlua[90m:2[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: Modify main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() end\nsomething();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() end\nsomething();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m20[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m3[0m function main() end\nsomething();
[7m [0m [91m                   ~[0m


Found 4 errors in 3 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/filePresent.tlua[90m:1[0m
     2  src/main.tlua[90m:3[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main()
end
nsomething();

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end","546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",{"version":"fd473bdbb337f0a29801a1ab2291871d-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"semanticDiagnosticsPerFile":[4],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
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
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "fd473bdbb337f0a29801a1ab2291871d-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "fd473bdbb337f0a29801a1ab2291871d-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    "./src/main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 2129
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [2]:: Modify main file again
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() end\nsomething();\nsomething();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function main() end\nsomething();\nsomething();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m20[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m3[0m function main() end\nsomething();\nsomething();
[7m [0m [91m                   ~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m34[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m3[0m function main() end\nsomething();\nsomething();
[7m [0m [91m                                 ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/filePresent.tlua[90m:1[0m
     3  src/main.tlua[90m:3[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main()
end
nsomething();
nsomething();

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,4]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end","546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",{"version":"4c973193e96e90b6d23d6c6b1ae242c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,5]],"options":{"composite":true},"referencedMap":[[3,1],[4,1]],"semanticDiagnosticsPerFile":[4],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
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
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/main.tlua",
      "version": "4c973193e96e90b6d23d6c6b1ae242c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "4c973193e96e90b6d23d6c6b1ae242c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    "./src/main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/main.tlua",
      "original": 4
    }
  ],
  "size": 2144
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [3]:: Add new file and update main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./newFile.tlua"/>
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() end\nsomething();\nsomething();foo();
//// [/home/src/workspaces/project/src/newFile.tlua] *new* 
function foo() return 20; end

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m                   ~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m34[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m                                 ~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() return 20; end
[7m [0m [91m~~~~~~~~[0m


Found 6 errors in 4 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/filePresent.tlua[90m:1[0m
     3  src/main.tlua[90m:4[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./newFile.tlua"/>
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main()
end
nsomething();
nsomething();
foo();

//// [/home/src/workspaces/project/src/newFile.lua] *new* 
function foo()
    return 20;
end

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua","./src/fileNotFound.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end","546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",{"version":"a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,6],[2,4,6]],"options":{"composite":true},"referencedMap":[[3,1],[5,2]],"semanticDiagnosticsPerFile":[4,5],"emitDiagnosticsPerFile":[[3,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        5
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua",
    "./src/fileNotFound.tlua"
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/newFile.tlua",
      "./src/fileNotFound.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/newFile.tlua",
      "./src/fileNotFound.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    "./src/newFile.tlua",
    "./src/main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 3
    },
    {
      "file": "./src/newFile.tlua",
      "original": 4
    },
    {
      "file": "./src/main.tlua",
      "original": 5
    }
  ],
  "size": 2580
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/src/newFile.tlua
*not cached* /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/newFile.tlua
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [4]:: Write file that could not be resolved
//// [/home/src/workspaces/project/src/fileNotFound.tlua] *new* 
function something2() return 20; end

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/fileNotFound.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something2() return 20; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m                   ~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m34[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();
[7m [0m [91m                                 ~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() return 20; end
[7m [0m [91m~~~~~~~~[0m


Found 7 errors in 5 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/fileNotFound.tlua[90m:1[0m
     1  src/filePresent.tlua[90m:1[0m
     3  src/main.tlua[90m:4[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/anotherFileWithSameReferenes.lua] *rewrite with same content*
//// [/home/src/workspaces/project/src/fileNotFound.lua] *new* 
function something2()
    return 20;
end

//// [/home/src/workspaces/project/src/main.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,6]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/fileNotFound.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",{"version":"9446582c514d88490d24a79bd5220fd3-function something2() return 20; end","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,3],[2,3,5]],"options":{"composite":true},"referencedMap":[[4,1],[6,2]],"semanticDiagnosticsPerFile":[3,4,5,6],"emitDiagnosticsPerFile":[[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[6,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5,6]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/fileNotFound.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        6
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/fileNotFound.tlua",
      "version": "9446582c514d88490d24a79bd5220fd3-function something2() return 20; end",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9446582c514d88490d24a79bd5220fd3-function something2() return 20; end",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "401c31599d841ddf62fd3035de8152eb-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/fileNotFound.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/fileNotFound.tlua",
      "original": 3
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 4
    },
    {
      "file": "./src/newFile.tlua",
      "original": 5
    },
    {
      "file": "./src/main.tlua",
      "original": 6
    }
  ],
  "size": 3113
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/src/fileNotFound.tlua
*not cached* /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
*not cached* /home/src/workspaces/project/src/newFile.tlua
*not cached* /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/fileNotFound.tlua
(computed .d.ts) /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
(computed .d.ts) /home/src/workspaces/project/src/main.tlua


Edit [5]:: Modify main file
//// [/home/src/workspaces/project/src/main.tlua] *modified* 
/// <reference path="./newFile.tlua"/>
/// <reference path="./filePresent.tlua"/>
/// <reference path="./fileNotFound.tlua"/>
function main() end\nsomething();\nsomething();foo();\nsomething();

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/anotherFileWithSameReferenes.tlua[0m:[93m3[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m3[0m function anotherFileWithSameReferenes() end
[7m [0m [91m~~~~~~~~[0m

[96msrc/fileNotFound.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something2() return 20; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/filePresent.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function something() return 10; end
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m4[0m function main() end\nsomething();\nsomething();foo();\nsomething();
[7m [0m [91m~~~~~~~~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m20[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();\nsomething();
[7m [0m [91m                   ~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m34[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();\nsomething();
[7m [0m [91m                                 ~[0m

[96msrc/main.tlua[0m:[93m4[0m:[93m54[0m - [91merror[0m[90m TLUA1127: [0mInvalid character.

[7m4[0m function main() end\nsomething();\nsomething();foo();\nsomething();
[7m [0m [91m                                                     ~[0m

[96msrc/newFile.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function foo() return 20; end
[7m [0m [91m~~~~~~~~[0m


Found 8 errors in 5 files.

Errors  Files
     1  src/anotherFileWithSameReferenes.tlua[90m:3[0m
     1  src/fileNotFound.tlua[90m:1[0m
     1  src/filePresent.tlua[90m:1[0m
     4  src/main.tlua[90m:4[0m
     1  src/newFile.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/main.lua] *modified* 
-- / <reference path="./newFile.tlua"/>
-- / <reference path="./filePresent.tlua"/>
-- / <reference path="./fileNotFound.tlua"/>
function main()
end
nsomething();
nsomething();
foo();
nsomething();

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,6]],"fileNames":["lib.luajit.d.tlua","./src/filePresent.tlua","./src/fileNotFound.tlua","./src/anotherFileWithSameReferenes.tlua","./src/newFile.tlua","./src/main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",{"version":"9446582c514d88490d24a79bd5220fd3-function something2() return 20; end","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end","signature":"9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end","signature":"2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},{"version":"4a73176f81ad7ea8b9c424604e1005a1-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();\\nsomething();","signature":"d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1}],"fileIdsList":[[2,3],[2,3,5]],"options":{"composite":true},"referencedMap":[[4,1],[6,2]],"semanticDiagnosticsPerFile":[3,4,5,6],"emitDiagnosticsPerFile":[[4,[{"pos":87,"end":95,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[6,[{"pos":126,"end":134,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3,4,5,6]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/filePresent.tlua",
        "./src/fileNotFound.tlua",
        "./src/anotherFileWithSameReferenes.tlua",
        "./src/newFile.tlua",
        "./src/main.tlua"
      ],
      "original": [
        2,
        6
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/filePresent.tlua",
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
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
      "fileName": "./src/filePresent.tlua",
      "version": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "signature": "b8c2dd7f24027698a62cf3ebc2acf2f5-function something() return 10; end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/fileNotFound.tlua",
      "version": "9446582c514d88490d24a79bd5220fd3-function something2() return 20; end",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "9446582c514d88490d24a79bd5220fd3-function something2() return 20; end",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/anotherFileWithSameReferenes.tlua",
      "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
      "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "546ddf7ccc4b1a871e995040f86a43c4-/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction anotherFileWithSameReferenes() end",
        "signature": "9e5e3156e0455533ff1173fa8208b160-\n(87,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/newFile.tlua",
      "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
      "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "a57450c8b35857a16921b41c1dbe1656-function foo() return 20; end",
        "signature": "2e6f268da2e4a4fbe3cf418d538b9f36-\n(0,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/main.tlua",
      "version": "4a73176f81ad7ea8b9c424604e1005a1-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();\\nsomething();",
      "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "4a73176f81ad7ea8b9c424604e1005a1-/// <reference path=\"./newFile.tlua\"/>\n/// <reference path=\"./filePresent.tlua\"/>\n/// <reference path=\"./fileNotFound.tlua\"/>\nfunction main() end\\nsomething();\\nsomething();foo();\\nsomething();",
        "signature": "d33342899b5a8577e8529f578e012712-\n(126,8): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    }
  ],
  "fileIdsList": [
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  ],
  "options": {
    "composite": true
  },
  "referencedMap": {
    "./src/anotherFileWithSameReferenes.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua"
    ],
    "./src/main.tlua": [
      "./src/filePresent.tlua",
      "./src/fileNotFound.tlua",
      "./src/newFile.tlua"
    ]
  },
  "semanticDiagnosticsPerFile": [
    "./src/fileNotFound.tlua",
    "./src/anotherFileWithSameReferenes.tlua",
    "./src/newFile.tlua",
    "./src/main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./src/anotherFileWithSameReferenes.tlua",
      [
        {
          "pos": 87,
          "end": 95,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/fileNotFound.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/filePresent.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/main.tlua",
      [
        {
          "pos": 126,
          "end": 134,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ],
    [
      "./src/newFile.tlua",
      [
        {
          "end": 8,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/filePresent.tlua",
      "original": 2
    },
    {
      "file": "./src/fileNotFound.tlua",
      "original": 3
    },
    {
      "file": "./src/anotherFileWithSameReferenes.tlua",
      "original": 4
    },
    {
      "file": "./src/newFile.tlua",
      "original": 5
    },
    {
      "file": "./src/main.tlua",
      "original": 6
    }
  ],
  "size": 3128
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/workspaces/project/src/fileNotFound.tlua
*not cached* /home/src/workspaces/project/src/anotherFileWithSameReferenes.tlua
*not cached* /home/src/workspaces/project/src/newFile.tlua
*not cached* /home/src/workspaces/project/src/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/main.tlua

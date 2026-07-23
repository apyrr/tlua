currentDirectory::/home/src/workspaces/solution
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/solution/child/child.tlua] *new* 
local child2 = require("child2");
function child()
    child2.child2();
end
return { child = child };
//// [/home/src/workspaces/solution/child/child2.tlua] *new* 
function child2() end
return { child2 = child2 };
//// [/home/src/workspaces/solution/child/tluaconfig.json] *new* 
{
    "compilerOptions": { "composite": true }
}
//// [/home/src/workspaces/solution/main/main.tlua] *new* 
local child = require("child.child");
function main()
    child.child();
end
return { main = main };
//// [/home/src/workspaces/solution/main/tluaconfig.json] *new* 
{
    "compilerOptions": { "composite": true, "rootDir": ".." },
    "references": [{ "path": "../child" }],
}

tlua --b main/tluaconfig.json -v --traceResolution --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * child/tluaconfig.json
    * main/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'child/tluaconfig.json' is out of date because output file 'child/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'child/tluaconfig.json'...

======== Resolving module 'child2' from '/home/src/workspaces/solution/child/child.tlua'. ========
File '/home/src/workspaces/solution/child/child2.tlua' exists - use it as a name resolution result.
======== Module name 'child2' was successfully resolved to '/home/src/workspaces/solution/child/child2.tlua'. ========
[96mchild/child.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local child2 = require("child2");
[7m [0m [91m~~~~~[0m

[96mchild/child2.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function child2() end
[7m [0m [91m~~~~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
child/child2.tlua
   Imported via "child2" from file 'child/child.tlua'
   Matched by default include pattern '**/*'
child/child.tlua
   Matched by default include pattern '**/*'
[[90mHH:MM:SS AM[0m] Project 'main/tluaconfig.json' is out of date because output file 'main/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'main/tluaconfig.json'...

======== Resolving module 'child.child' from '/home/src/workspaces/solution/main/main.tlua'. ========
File '/home/src/workspaces/solution/child/child.tlua' exists - use it as a name resolution result.
======== Module name 'child.child' was successfully resolved to '/home/src/workspaces/solution/child/child.tlua'. ========
[96mmain/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local child = require("child.child");
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
main/main.tlua
   Matched by default include pattern '**/*'

Found 3 errors in 3 files.

Errors  Files
     1  child/child.tlua[90m:1[0m
     1  child/child2.tlua[90m:1[0m
     1  main/main.tlua[90m:1[0m

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
//// [/home/src/workspaces/solution/child/child.lua] *new* 
local child2 = require("child2");
function child()
  child2.child2();
end
return { child = child };

//// [/home/src/workspaces/solution/child/child2.lua] *new* 
function child2()
end
return { child2 = child2 };

//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./child2.tlua","./child.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"471b3eb6a1e8bc37213aba99b3152ef6-function child2() end\nreturn { child2 = child2 };","0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };"],"options":{"composite":true},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./child2.tlua",
        "./child.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./child2.tlua",
    "./child.tlua"
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
      "fileName": "./child2.tlua",
      "version": "471b3eb6a1e8bc37213aba99b3152ef6-function child2() end\nreturn { child2 = child2 };",
      "signature": "471b3eb6a1e8bc37213aba99b3152ef6-function child2() end\nreturn { child2 = child2 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./child.tlua",
      "version": "0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };",
      "signature": "0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./child.tlua",
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
      "./child2.tlua",
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
      "file": "./child2.tlua",
      "original": 2
    },
    {
      "file": "./child.tlua",
      "original": 3
    }
  ],
  "size": 1477
}
//// [/home/src/workspaces/solution/main/main.lua] *new* 
local child = require("child.child");
function main()
  child.child();
end
return { main = main };

//// [/home/src/workspaces/solution/main/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"10d1d3cd65977580c3d40ec5d3445f47-local child = require(\"child.child\");\nfunction main()\n    child.child();\nend\nreturn { main = main };"],"options":{"composite":true,"rootDir":".."},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/workspaces/solution/main/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./main.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./main.tlua"
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
      "fileName": "./main.tlua",
      "version": "10d1d3cd65977580c3d40ec5d3445f47-local child = require(\"child.child\");\nfunction main()\n    child.child();\nend\nreturn { main = main };",
      "signature": "10d1d3cd65977580c3d40ec5d3445f47-local child = require(\"child.child\");\nfunction main()\n    child.child();\nend\nreturn { main = main };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "rootDir": ".."
  },
  "emitDiagnosticsPerFile": [
    [
      "./main.tlua",
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
      "file": "./main.tlua",
      "original": 2
    }
  ],
  "size": 1262
}

child/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/child/child2.tlua
*refresh*    /home/src/workspaces/solution/child/child.tlua
Signatures::

main/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/solution/main/main.tlua
Signatures::


Edit [0]:: delete child2 file
//// [/home/src/workspaces/solution/child/child2.lua] *deleted*
//// [/home/src/workspaces/solution/child/child2.tlua] *deleted*

tlua --b main/tluaconfig.json -v --traceResolution --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * child/tluaconfig.json
    * main/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'child/tluaconfig.json' is out of date because buildinfo file 'child/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'child/tluaconfig.json'...

======== Resolving module 'child2' from '/home/src/workspaces/solution/child/child.tlua'. ========
File '/home/src/workspaces/solution/child/child2.tlua' does not exist.
File '/home/src/workspaces/solution/child/child2.tsx' does not exist.
File '/home/src/workspaces/solution/child/child2/init.tlua' does not exist.
File '/home/src/workspaces/solution/child/child2/init.tsx' does not exist.
File '/home/src/workspaces/solution/child/child2.lua' does not exist.
File '/home/src/workspaces/solution/child/child2/init.lua' does not exist.
======== Module name 'child2' was not resolved. ========
[96mchild/child.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local child2 = require("child2");
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
child/child.tlua
   Matched by default include pattern '**/*'
[[90mHH:MM:SS AM[0m] Project 'main/tluaconfig.json' is out of date because buildinfo file 'main/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'main/tluaconfig.json'...

======== Resolving module 'child.child' from '/home/src/workspaces/solution/main/main.tlua'. ========
File '/home/src/workspaces/solution/child/child.tlua' exists - use it as a name resolution result.
======== Module name 'child.child' was successfully resolved to '/home/src/workspaces/solution/child/child.tlua'. ========
[96mmain/main.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local child = require("child.child");
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
main/main.tlua
   Matched by default include pattern '**/*'

Found 2 errors in 2 files.

Errors  Files
     1  child/child.tlua[90m:1[0m
     1  main/main.tlua[90m:1[0m

//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./child.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };"],"options":{"composite":true},"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/workspaces/solution/child/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./child.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./child.tlua"
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
      "fileName": "./child.tlua",
      "version": "0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };",
      "signature": "0a16d4ed2a3b10ab1ac9c47e521340d2-local child2 = require(\"child2\");\nfunction child()\n    child2.child2();\nend\nreturn { child = child };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true
  },
  "emitDiagnosticsPerFile": [
    [
      "./child.tlua",
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
      "file": "./child.tlua",
      "original": 2
    }
  ],
  "size": 1249
}

child/tluaconfig.json::
SemanticDiagnostics::
Signatures::

main/tluaconfig.json::
SemanticDiagnostics::
Signatures::

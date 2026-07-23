currentDirectory::/user/username/projects/reexport
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/reexport/src/main/index.tlua] *new* 
local pure = require("pure");
pure;

local session: Session = {
    foo: 1
};
//// [/user/username/projects/reexport/src/main/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "../../out",
        "rootDir": "../",
    },
    "include": ["**/*.tlua"],
    "references": [{ "path": "../pure" }],
}
//// [/user/username/projects/reexport/src/pure/init.tlua] *new* 
local session = require("pure.session");
session;
//// [/user/username/projects/reexport/src/pure/session.tlua] *new* 
interface Session {
    foo: number;
    // bar: number;
}
//// [/user/username/projects/reexport/src/pure/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "outDir": "../../out",
        "rootDir": "../",
    },
    "include": ["**/*.tlua"],
}
//// [/user/username/projects/reexport/src/tluaconfig.json] *new* 
{
    "files": [],
    "include": [],
    "references": [{ "path": "./pure" }, { "path": "./main" }],
}

tlua -b -w -verbose src
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * src/pure/tluaconfig.json
    * src/main/tluaconfig.json
    * src/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'src/pure/tluaconfig.json' is out of date because output file 'out/pure/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'src/pure/tluaconfig.json'...

[96msrc/pure/init.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local session = require("pure.session");
[7m [0m [91m~~~~~[0m

[96msrc/pure/session.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface Session {
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'src/main/tluaconfig.json' is out of date because output file 'out/main/tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'src/main/tluaconfig.json'...

[96msrc/main/index.tlua[0m:[93m5[0m:[93m8[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m5[0m     foo: 1
[7m [0m [91m       ~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

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
//// [/user/username/projects/reexport/out/main/index.lua] *new* 
local pure = require("pure");
pure;
local session = {
  foo, 1
};

//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":["../../src/main/index.tlua"]}
//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "../../src/main/index.tlua"
      ],
      "original": "../../src/main/index.tlua"
    }
  ],
  "size": 78
}
//// [/user/username/projects/reexport/out/pure/init.lua] *new* 
local session = require("pure.session");
session;

//// [/user/username/projects/reexport/out/pure/session.lua] *new* 

//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../../src/pure/session.tlua","../../src/pure/init.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}","cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;"],"options":{"composite":true,"outDir":"..","rootDir":"../../src"},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../src/pure/session.tlua",
        "../../src/pure/init.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../src/pure/session.tlua",
    "../../src/pure/init.tlua"
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
      "fileName": "../../src/pure/session.tlua",
      "version": "389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}",
      "signature": "389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../../src/pure/init.tlua",
      "version": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "signature": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../../src"
  },
  "emitDiagnosticsPerFile": [
    [
      "../../src/pure/init.tlua",
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
      "../../src/pure/session.tlua",
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
      "file": "../../src/pure/session.tlua",
      "original": 2
    },
    {
      "file": "../../src/pure/init.tlua",
      "original": 3
    }
  ],
  "size": 1495
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/reexport
  /user/username/projects/reexport/out
  /user/username/projects/reexport/out/pure
  /user/username/projects/reexport/src
  /user/username/projects/reexport/src/main (recursive)
  /user/username/projects/reexport/src/pure (recursive)
src/pure/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/reexport/src/pure/session.tlua
*refresh*    /user/username/projects/reexport/src/pure/init.tlua
Signatures::

src/main/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/reexport/src/main/index.tlua
Signatures::


Edit [0]:: Introduce error
//// [/user/username/projects/reexport/src/pure/session.tlua] *modified* 
interface Session {
    foo: number;
    bar: number;
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * src/pure/tluaconfig.json
    * src/main/tluaconfig.json
    * src/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'src/pure/tluaconfig.json' is out of date because buildinfo file 'out/pure/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'src/pure/tluaconfig.json'...

[96msrc/pure/init.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local session = require("pure.session");
[7m [0m [91m~~~~~[0m

[96msrc/pure/session.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface Session {
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'src/main/tluaconfig.json' is out of date because it has errors.

[[90mHH:MM:SS AM[0m] Building project 'src/main/tluaconfig.json'...

[96msrc/main/index.tlua[0m:[93m5[0m:[93m8[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m5[0m     foo: 1
[7m [0m [91m       ~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

//// [/user/username/projects/reexport/out/main/index.lua] *rewrite with same content*
//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*
//// [/user/username/projects/reexport/out/pure/session.lua] *rewrite with same content*
//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../../src/pure/session.tlua","../../src/pure/init.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"7c5411f65d6041f40ef12d4731d29127-interface Session {\n    foo: number;\n    bar: number;\n}","signature":"fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;"],"options":{"composite":true,"outDir":"..","rootDir":"../../src"},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../src/pure/session.tlua",
        "../../src/pure/init.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../src/pure/session.tlua",
    "../../src/pure/init.tlua"
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
      "fileName": "../../src/pure/session.tlua",
      "version": "7c5411f65d6041f40ef12d4731d29127-interface Session {\n    foo: number;\n    bar: number;\n}",
      "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "7c5411f65d6041f40ef12d4731d29127-interface Session {\n    foo: number;\n    bar: number;\n}",
        "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../../src/pure/init.tlua",
      "version": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "signature": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../../src"
  },
  "emitDiagnosticsPerFile": [
    [
      "../../src/pure/init.tlua",
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
      "../../src/pure/session.tlua",
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
      "file": "../../src/pure/session.tlua",
      "original": 2
    },
    {
      "file": "../../src/pure/init.tlua",
      "original": 3
    }
  ],
  "size": 1659
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/reexport
  /user/username/projects/reexport/out
  /user/username/projects/reexport/out/pure
  /user/username/projects/reexport/src
  /user/username/projects/reexport/src/main (recursive)
  /user/username/projects/reexport/src/pure (recursive)
src/pure/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /user/username/projects/reexport/src/pure/session.tlua
Signatures::
(computed .d.ts) /user/username/projects/reexport/src/pure/session.tlua

src/main/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/reexport/src/main/index.tlua
Signatures::


Edit [1]:: Fix error
//// [/user/username/projects/reexport/src/pure/session.tlua] *modified* 
interface Session {
    foo: number;
    // bar: number;
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Projects in this build: 
    * src/pure/tluaconfig.json
    * src/main/tluaconfig.json
    * src/tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'src/pure/tluaconfig.json' is out of date because buildinfo file 'out/pure/tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'src/pure/tluaconfig.json'...

[96msrc/pure/init.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local session = require("pure.session");
[7m [0m [91m~~~~~[0m

[96msrc/pure/session.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m interface Session {
[7m [0m [91m~~~~~~~~~[0m

[[90mHH:MM:SS AM[0m] Project 'src/main/tluaconfig.json' is out of date because it has errors.

[[90mHH:MM:SS AM[0m] Building project 'src/main/tluaconfig.json'...

[96msrc/main/index.tlua[0m:[93m5[0m:[93m8[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m5[0m     foo: 1
[7m [0m [91m       ~[0m

[[90mHH:MM:SS AM[0m] Found 3 errors. Watching for file changes.

//// [/user/username/projects/reexport/out/main/index.lua] *rewrite with same content*
//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo] *rewrite with same content*
//// [/user/username/projects/reexport/out/main/tluaconfig.tluabuildinfo.readable.baseline.txt] *rewrite with same content*
//// [/user/username/projects/reexport/out/pure/session.lua] *rewrite with same content*
//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","../../src/pure/session.tlua","../../src/pure/init.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}","signature":"fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;"],"options":{"composite":true,"outDir":"..","rootDir":"../../src"},"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[2,[{"end":9,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2,3]}
//// [/user/username/projects/reexport/out/pure/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "../../src/pure/session.tlua",
        "../../src/pure/init.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../../src/pure/session.tlua",
    "../../src/pure/init.tlua"
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
      "fileName": "../../src/pure/session.tlua",
      "version": "389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}",
      "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "389f150219fcbca1cab654141dd634bf-interface Session {\n    foo: number;\n    // bar: number;\n}",
        "signature": "fee2aa0e86ade495c53cf57aec2202bd-\n(0,9): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../../src/pure/init.tlua",
      "version": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "signature": "cb2d4fba035d98243c4fb16380a83c47-local session = require(\"pure.session\");\nsession;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "outDir": "..",
    "rootDir": "../../src"
  },
  "emitDiagnosticsPerFile": [
    [
      "../../src/pure/init.tlua",
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
      "../../src/pure/session.tlua",
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
      "file": "../../src/pure/session.tlua",
      "original": 2
    },
    {
      "file": "../../src/pure/init.tlua",
      "original": 3
    }
  ],
  "size": 1662
}

Watch Registrations::
Directory watches::
  /home/src/tslibs/TS/Lib
  /user/username/projects/reexport
  /user/username/projects/reexport/out
  /user/username/projects/reexport/out/pure
  /user/username/projects/reexport/src
  /user/username/projects/reexport/src/main (recursive)
  /user/username/projects/reexport/src/pure (recursive)
src/pure/tluaconfig.json::
SemanticDiagnostics::
*refresh*    /user/username/projects/reexport/src/pure/session.tlua
Signatures::
(computed .d.ts) /user/username/projects/reexport/src/pure/session.tlua

src/main/tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/reexport/src/main/index.tlua
Signatures::

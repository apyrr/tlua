currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/a.tlua] *new* 
local a = { private: 10 };
//// [/home/src/projects/project/b.tlua] *new* 
local b = 10;
//// [/home/src/projects/project/c.tlua] *new* 
loccl a = { private: 10 };
//// [/home/src/projects/project/d.tlua] *new* 
locdl a = { private: 10 };
//// [/home/src/projects/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true,
    }
}

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"semanticDiagnosticsPerFile":[1,2,3,4,5],"affectedFilesPendingEmit":[2,3,4,5]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./a.tlua",
      "Js",
      2
    ],
    [
      "./b.tlua",
      "Js",
      3
    ],
    [
      "./c.tlua",
      "Js",
      4
    ],
    [
      "./d.tlua",
      "Js",
      5
    ]
  ],
  "size": 1269
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
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [0]:: no change

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m


tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [1]:: With declaration enabled noEmit - Should report errors

tlua --noEmit --declaration
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true},"semanticDiagnosticsPerFile":[1,2,3,4,5],"affectedFilesPendingEmit":[2,3,4,5]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./a.tlua",
      "Js|Dts",
      2
    ],
    [
      "./b.tlua",
      "Js|Dts",
      3
    ],
    [
      "./c.tlua",
      "Js|Dts",
      4
    ],
    [
      "./d.tlua",
      "Js|Dts",
      5
    ]
  ],
  "size": 1300
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [2]:: With declaration and declarationMap noEmit - Should report errors

tlua --noEmit --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true,"declarationMap":true},"semanticDiagnosticsPerFile":[1,2,3,4,5],"affectedFilesPendingEmit":[2,3,4,5]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./a.tlua",
      "Js|Dts|DtsMap",
      2
    ],
    [
      "./b.tlua",
      "Js|Dts|DtsMap",
      3
    ],
    [
      "./c.tlua",
      "Js|Dts|DtsMap",
      4
    ],
    [
      "./d.tlua",
      "Js|Dts|DtsMap",
      5
    ]
  ],
  "size": 1322
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [3]:: no change

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m


tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [4]:: Dts Emit with error

tlua --declaration
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 9 errors in 4 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     3  c.tlua[90m:1[0m
     3  d.tlua[90m:1[0m

//// [/home/src/projects/project/a.lua] *new* 
local a = { private, 10 };

//// [/home/src/projects/project/b.lua] *new* 
local b = 10;

//// [/home/src/projects/project/c.lua] *new* 
loccl;
a = { private, 10 };

//// [/home/src/projects/project/d.lua] *new* 
locdl;
a = { private, 10 };

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true},"semanticDiagnosticsPerFile":[1,2,3,4,5],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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
      "version": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./a.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
  "size": 1756
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [5]:: Fix the error
//// [/home/src/projects/project/a.tlua] *modified* 
local a = { public: 10 };

tlua --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { public: 10 };
[7m [0m [91m                  ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };","signature":"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","impliedNodeFormat":1},"f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"semanticDiagnosticsPerFile":[1,3,4,5],"emitDiagnosticsPerFile":[[3,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[4,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[5,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"changeFileSet":[2]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
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
      "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
        "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
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
  "changeFileSet": [
    "./a.tlua"
  ],
  "size": 1730
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [6]:: With declaration enabled noEmit

tlua --noEmit --declaration
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { public: 10 };
[7m [0m [91m                  ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };","signature":"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","impliedNodeFormat":1},"f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true},"semanticDiagnosticsPerFile":[1,3,4,5],"changeFileSet":[2],"affectedFilesPendingEmit":[[3],[4],[5]]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
        "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
  ],
  "changeFileSet": [
    "./a.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./b.tlua",
      "Dts",
      [
        3
      ]
    ],
    [
      "./c.tlua",
      "Dts",
      [
        4
      ]
    ],
    [
      "./d.tlua",
      "Dts",
      [
        5
      ]
    ]
  ],
  "size": 1429
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [7]:: With declaration and declarationMap noEmit

tlua --noEmit --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { public: 10 };
[7m [0m [91m                  ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { private: 10 };
[7m [0m [91m                   ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };","signature":"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","impliedNodeFormat":1},"f264fbb7445fd7350e2974971e7a3290-local b = 10;","b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true,"declarationMap":true},"semanticDiagnosticsPerFile":[1,3,4,5],"changeFileSet":[2],"affectedFilesPendingEmit":[[3,56],[4,56],[5,56]]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
        "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./b.tlua",
    "./c.tlua",
    "./d.tlua"
  ],
  "changeFileSet": [
    "./a.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./b.tlua",
      "Dts|DtsMap",
      [
        3,
        56
      ]
    ],
    [
      "./c.tlua",
      "Dts|DtsMap",
      [
        4,
        56
      ]
    ],
    [
      "./d.tlua",
      "Dts|DtsMap",
      [
        5,
        56
      ]
    ]
  ],
  "size": 1460
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::


Edit [8]:: Fix the another 
//// [/home/src/projects/project/c.tlua] *modified* 
loccl a = { public: 10 };

tlua --noEmit --declaration --declarationMap
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { public: 10 };
[7m [0m [91m                  ~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m loccl a = { public: 10 };
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m loccl a = { public: 10 };
[7m [0m [91m                  ~[0m

[96md.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS1435: [0mUnknown keyword or identifier. Did you mean 'local'?

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m~~~~~[0m

[96md.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m locdl a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 5 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     2  c.tlua[90m:1[0m
     2  d.tlua[90m:1[0m

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,5]],"fileNames":["lib.luajit.d.tlua","./a.tlua","./b.tlua","./c.tlua","./d.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };","signature":"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","impliedNodeFormat":1},"f264fbb7445fd7350e2974971e7a3290-local b = 10;",{"version":"07db3a4539f58b535f1dc1dcf2a83432-loccl a = { public: 10 };","signature":"b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };","impliedNodeFormat":1},"022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };"],"options":{"declaration":true,"declarationMap":true},"semanticDiagnosticsPerFile":[1,3,5],"changeFileSet":[2,4],"affectedFilesPendingEmit":[[3,56],[4,56],[5,56]]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
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
      "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8abf0a35ccc1f4adf7fe3f587e0554ce-local a = { public: 10 };",
        "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./b.tlua",
      "version": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "signature": "f264fbb7445fd7350e2974971e7a3290-local b = 10;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./c.tlua",
      "version": "07db3a4539f58b535f1dc1dcf2a83432-loccl a = { public: 10 };",
      "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "07db3a4539f58b535f1dc1dcf2a83432-loccl a = { public: 10 };",
        "signature": "b83ed208082453b32c13c2f534b4b30d-loccl a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./d.tlua",
      "version": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "signature": "022c999955420bb2c43c2f0d04ae145f-locdl a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "declarationMap": true
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./b.tlua",
    "./d.tlua"
  ],
  "changeFileSet": [
    "./a.tlua",
    "./c.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "./b.tlua",
      "Dts|DtsMap",
      [
        3,
        56
      ]
    ],
    [
      "./c.tlua",
      "Dts|DtsMap",
      [
        4,
        56
      ]
    ],
    [
      "./d.tlua",
      "Dts|DtsMap",
      [
        5,
        56
      ]
    ]
  ],
  "size": 1567
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/projects/project/a.tlua
*not cached* /home/src/projects/project/b.tlua
*not cached* /home/src/projects/project/c.tlua
*not cached* /home/src/projects/project/d.tlua
Signatures::

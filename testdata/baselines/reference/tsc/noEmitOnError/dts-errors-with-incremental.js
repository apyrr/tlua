currentDirectory::/user/username/projects/noEmitOnError
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/noEmitOnError/shared/types/db.tlua] *new* 
interface A {
    name: string;
}
//// [/user/username/projects/noEmitOnError/src/main.tlua] *new* 
local a = { private: 10 };
//// [/user/username/projects/noEmitOnError/src/other.tlua] *new* 
console.log("hi");
//// [/user/username/projects/noEmitOnError/tsconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./dev-build",
        "declaration": false,
        "incremental": true,
        "noEmitOnError": true,
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in src/main.tlua[90m:1[0m

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
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}","87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","219352914d852ae252f8f427f8a2219b-console.log(\"hi\");"],"options":{"declaration":false,"noEmitOnError":true,"outDir":"./"},"semanticDiagnosticsPerFile":[1,2,3,4],"affectedFilesPendingEmit":[2,3,4]}
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "../shared/types/db.tlua",
        "../src/main.tlua",
        "../src/other.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../shared/types/db.tlua",
    "../src/main.tlua",
    "../src/other.tlua"
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
      "fileName": "../shared/types/db.tlua",
      "version": "a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",
      "signature": "a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/main.tlua",
      "version": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/other.tlua",
      "version": "219352914d852ae252f8f427f8a2219b-console.log(\"hi\");",
      "signature": "219352914d852ae252f8f427f8a2219b-console.log(\"hi\");",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": false,
    "noEmitOnError": true,
    "outDir": "./"
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../shared/types/db.tlua",
    "../src/main.tlua",
    "../src/other.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "../shared/types/db.tlua",
      "Js",
      2
    ],
    [
      "../src/main.tlua",
      "Js",
      3
    ],
    [
      "../src/other.tlua",
      "Js",
      4
    ]
  ],
  "size": 1307
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m20[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { private: 10 };
[7m [0m [91m                   ~[0m


Found 1 error in src/main.tlua[90m:1[0m


tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a = { p: 10 };

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m14[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { p: 10 };
[7m [0m [91m             ~[0m


Found 1 error in src/main.tlua[90m:1[0m

//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",{"version":"96f3993829687baf36a7f9899f05b8a5-local a = { p: 10 };","signature":"87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };","impliedNodeFormat":1},"219352914d852ae252f8f427f8a2219b-console.log(\"hi\");"],"options":{"declaration":false,"noEmitOnError":true,"outDir":"./"},"semanticDiagnosticsPerFile":[1,2,4],"changeFileSet":[3],"affectedFilesPendingEmit":[2,3,4]}
//// [/user/username/projects/noEmitOnError/dev-build/tsconfig.tsbuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "errors": true,
  "root": [
    {
      "files": [
        "../shared/types/db.tlua",
        "../src/main.tlua",
        "../src/other.tlua"
      ],
      "original": [
        2,
        4
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "../shared/types/db.tlua",
    "../src/main.tlua",
    "../src/other.tlua"
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
      "fileName": "../shared/types/db.tlua",
      "version": "a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",
      "signature": "a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "../src/main.tlua",
      "version": "96f3993829687baf36a7f9899f05b8a5-local a = { p: 10 };",
      "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "96f3993829687baf36a7f9899f05b8a5-local a = { p: 10 };",
        "signature": "87fcdb8c49ccbb081c1d2c2261f8e674-local a = { private: 10 };",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "../src/other.tlua",
      "version": "219352914d852ae252f8f427f8a2219b-console.log(\"hi\");",
      "signature": "219352914d852ae252f8f427f8a2219b-console.log(\"hi\");",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": false,
    "noEmitOnError": true,
    "outDir": "./"
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "../shared/types/db.tlua",
    "../src/other.tlua"
  ],
  "changeFileSet": [
    "../src/main.tlua"
  ],
  "affectedFilesPendingEmit": [
    [
      "../shared/types/db.tlua",
      "Js",
      2
    ],
    [
      "../src/main.tlua",
      "Js",
      3
    ],
    [
      "../src/other.tlua",
      "Js",
      4
    ]
  ],
  "size": 1427
}

tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [2]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m1[0m:[93m14[0m - [91merror[0m[90m TS1005: [0m',' expected.

[7m1[0m local a = { p: 10 };
[7m [0m [91m             ~[0m


Found 1 error in src/main.tlua[90m:1[0m


tsconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::

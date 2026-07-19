currentDirectory::/user/username/projects/noEmitOnError
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/noEmitOnError/shared/types/db.tlua] *new* 
interface A {
    name: string;
}
//// [/user/username/projects/noEmitOnError/src/main.tlua] *new* 
local a = {
    lastName: 'sdsd'
;
//// [/user/username/projects/noEmitOnError/src/other.tlua] *new* 
console.log("hi");
//// [/user/username/projects/noEmitOnError/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "outDir": "./dev-build",
        "declaration": true,
        "incremental": true,
        "noEmitOnError": true,
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m2[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     lastName: 'sdsd'
[7m [0m [91m            ~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m2[0m - [91merror[0m[90m TLUA1005: [0m'}' expected.

[7m3[0m ;
[7m [0m [91m ~[0m

  [96msrc/main.tlua[0m:[93m1[0m:[93m11[0m - The parser expected to find a '}' to match the '{' token here.
    [7m1[0m local a = {
    [7m [0m [96m          ~[0m


Found 2 errors in the same file, starting at: src/main.tlua[90m:2[0m

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
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","errors":true,"root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}","59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;","219352914d852ae252f8f427f8a2219b-console.log(\"hi\");"],"options":{"declaration":true,"noEmitOnError":true,"outDir":"./"},"semanticDiagnosticsPerFile":[1,2,3,4],"affectedFilesPendingEmit":[2,3,4]}
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
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
      "version": "59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;",
      "signature": "59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;",
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
    "declaration": true,
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
      "Js|Dts",
      2
    ],
    [
      "../src/main.tlua",
      "Js|Dts",
      3
    ],
    [
      "../src/other.tlua",
      "Js|Dts",
      4
    ]
  ],
  "size": 1316
}

tluaconfig.json::
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
[96msrc/main.tlua[0m:[93m2[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     lastName: 'sdsd'
[7m [0m [91m            ~[0m

[96msrc/main.tlua[0m:[93m3[0m:[93m2[0m - [91merror[0m[90m TLUA1005: [0m'}' expected.

[7m3[0m ;
[7m [0m [91m ~[0m

  [96msrc/main.tlua[0m:[93m1[0m:[93m11[0m - The parser expected to find a '}' to match the '{' token here.
    [7m1[0m local a = {
    [7m [0m [96m          ~[0m


Found 2 errors in the same file, starting at: src/main.tlua[90m:2[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::


Edit [1]:: Fix error
//// [/user/username/projects/noEmitOnError/src/main.tlua] *modified* 
local a = {
    lastName: 'sdsd'
};

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tlua[0m:[93m2[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     lastName: 'sdsd'
[7m [0m [91m            ~[0m


Found 1 error in src/main.tlua[90m:2[0m

//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","errors":true,"root":[[2,4]],"fileNames":["lib.luajit.d.tlua","../shared/types/db.tlua","../src/main.tlua","../src/other.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"a46794bf89f4fcf7d03d14808013af4b-interface A {\n    name: string;\n}",{"version":"277eeeebc44df34e3cd5147c67d424c9-local a = {\n    lastName: 'sdsd'\n};","signature":"59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;","impliedNodeFormat":1},"219352914d852ae252f8f427f8a2219b-console.log(\"hi\");"],"options":{"declaration":true,"noEmitOnError":true,"outDir":"./"},"semanticDiagnosticsPerFile":[1,2,4],"changeFileSet":[3],"affectedFilesPendingEmit":[2,3,4]}
//// [/user/username/projects/noEmitOnError/dev-build/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
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
      "version": "277eeeebc44df34e3cd5147c67d424c9-local a = {\n    lastName: 'sdsd'\n};",
      "signature": "59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "277eeeebc44df34e3cd5147c67d424c9-local a = {\n    lastName: 'sdsd'\n};",
        "signature": "59296670b22d81d5794886c2c4cc0097-local a = {\n    lastName: 'sdsd'\n;",
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
    "declaration": true,
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
      "Js|Dts",
      2
    ],
    [
      "../src/main.tlua",
      "Js|Dts",
      3
    ],
    [
      "../src/other.tlua",
      "Js|Dts",
      4
    ]
  ],
  "size": 1453
}

tluaconfig.json::
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
[96msrc/main.tlua[0m:[93m2[0m:[93m13[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m2[0m     lastName: 'sdsd'
[7m [0m [91m            ~[0m


Found 1 error in src/main.tlua[90m:2[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /user/username/projects/noEmitOnError/shared/types/db.tlua
*not cached* /user/username/projects/noEmitOnError/src/main.tlua
*not cached* /user/username/projects/noEmitOnError/src/other.tlua
Signatures::

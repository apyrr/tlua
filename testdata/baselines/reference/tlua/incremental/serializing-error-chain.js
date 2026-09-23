currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tlua] *new* 
declare f: {
    (x: never): any;
    (x: { a?: number }): any;
}
f({ a = "s" })
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true,
        "strict": true,
        "module": "esnext",
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mindex.tlua[0m:[93m5[0m:[93m5[0m - [91merror[0m[90m TLUA2769: [0mNo overload matches this call.
  The last overload gave the following error.
    Type 'string' is not assignable to type 'number'.

[7m5[0m f({ a = "s" })
[7m [0m [91m    ~[0m

  [96mindex.tlua[0m:[93m3[0m:[93m11[0m - The expected type comes from property 'a' which is declared here on type '{ a?: number | nil; }'
    [7m3[0m     (x: { a?: number }): any;
    [7m [0m [96m          ~[0m

  [96mindex.tlua[0m:[93m3[0m:[93m5[0m - The last overload is declared here.
    [7m3[0m     (x: { a?: number }): any;
    [7m [0m [96m    ~~~~~~~~~~~~~~~~~~~~~~~~~[0m


Found 1 error in index.tlua[90m:5[0m

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
//// [/home/src/workspaces/project/index.lua] *new* 
f({ a = "s" });

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./index.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"d1bc5398368c44fc59077ddc45f7771f-declare f: {\n    (x: never): any;\n    (x: { a?: number }): any;\n}\nf({ a = \"s\" })"],"options":{"module":99,"strict":true},"semanticDiagnosticsPerFile":[[2,[{"pos":70,"end":71,"code":2769,"category":1,"messageKey":"No_overload_matches_this_call_2769","messageChain":[{"pos":70,"end":71,"code":2770,"category":1,"messageKey":"The_last_overload_gave_the_following_error_2770","messageChain":[{"pos":70,"end":71,"code":2322,"category":1,"messageKey":"Type_0_is_not_assignable_to_type_1_2322","messageArgs":["string","number"],"relatedInformation":[{"pos":44,"end":45,"code":6500,"category":3,"messageKey":"The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500","messageArgs":["a","{ a?: number | nil; }"]}]}],"relatedInformation":[{"pos":44,"end":45,"code":6500,"category":3,"messageKey":"The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500","messageArgs":["a","{ a?: number | nil; }"]}]}],"relatedInformation":[{"pos":44,"end":45,"code":6500,"category":3,"messageKey":"The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500","messageArgs":["a","{ a?: number | nil; }"]},{"pos":38,"end":63,"code":2771,"category":1,"messageKey":"The_last_overload_is_declared_here_2771"}]}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./index.tlua"
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
      "fileName": "./index.tlua",
      "version": "d1bc5398368c44fc59077ddc45f7771f-declare f: {\n    (x: never): any;\n    (x: { a?: number }): any;\n}\nf({ a = \"s\" })",
      "signature": "d1bc5398368c44fc59077ddc45f7771f-declare f: {\n    (x: never): any;\n    (x: { a?: number }): any;\n}\nf({ a = \"s\" })",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "module": 99,
    "strict": true
  },
  "semanticDiagnosticsPerFile": [
    [
      "./index.tlua",
      [
        {
          "pos": 70,
          "end": 71,
          "code": 2769,
          "category": 1,
          "messageKey": "No_overload_matches_this_call_2769",
          "messageChain": [
            {
              "pos": 70,
              "end": 71,
              "code": 2770,
              "category": 1,
              "messageKey": "The_last_overload_gave_the_following_error_2770",
              "messageChain": [
                {
                  "pos": 70,
                  "end": 71,
                  "code": 2322,
                  "category": 1,
                  "messageKey": "Type_0_is_not_assignable_to_type_1_2322",
                  "messageArgs": [
                    "string",
                    "number"
                  ],
                  "relatedInformation": [
                    {
                      "pos": 44,
                      "end": 45,
                      "code": 6500,
                      "category": 3,
                      "messageKey": "The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500",
                      "messageArgs": [
                        "a",
                        "{ a?: number | nil; }"
                      ]
                    }
                  ]
                }
              ],
              "relatedInformation": [
                {
                  "pos": 44,
                  "end": 45,
                  "code": 6500,
                  "category": 3,
                  "messageKey": "The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500",
                  "messageArgs": [
                    "a",
                    "{ a?: number | nil; }"
                  ]
                }
              ]
            }
          ],
          "relatedInformation": [
            {
              "pos": 44,
              "end": 45,
              "code": 6500,
              "category": 3,
              "messageKey": "The_expected_type_comes_from_property_0_which_is_declared_here_on_type_1_6500",
              "messageArgs": [
                "a",
                "{ a?: number | nil; }"
              ]
            },
            {
              "pos": 38,
              "end": 63,
              "code": 2771,
              "category": 1,
              "messageKey": "The_last_overload_is_declared_here_2771"
            }
          ]
        }
      ]
    ]
  ],
  "size": 2192
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/index.tlua
Signatures::


Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mindex.tlua[0m:[93m5[0m:[93m5[0m - [91merror[0m[90m TLUA2769: [0mNo overload matches this call.
  The last overload gave the following error.
    Type 'string' is not assignable to type 'number'.

[7m5[0m f({ a = "s" })
[7m [0m [91m    ~[0m

  [96mindex.tlua[0m:[93m3[0m:[93m11[0m - The expected type comes from property 'a' which is declared here on type '{ a?: number | nil; }'
    [7m3[0m     (x: { a?: number }): any;
    [7m [0m [96m          ~[0m

  [96mindex.tlua[0m:[93m3[0m:[93m5[0m - The last overload is declared here.
    [7m3[0m     (x: { a?: number }): any;
    [7m [0m [96m    ~~~~~~~~~~~~~~~~~~~~~~~~~[0m


Found 1 error in index.tlua[90m:5[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::

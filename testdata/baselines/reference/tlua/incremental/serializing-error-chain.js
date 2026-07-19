currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/index.tsx] *new* 
interface JsxElementChildrenAttribute { children: {}; }
interface JsxIntrinsicElements { div: {} }

declare React: any;

declare function Component(props: never): any;
declare function Component(props: { children?: number }): any;
(<Component>
    <div />
    <div />
</Component>)
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "incremental": true,
        "strict": true,
        "jsx": "react",
        "module": "esnext",
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mindex.tsx[0m:[93m8[0m:[93m3[0m - [91merror[0m[90m TLUA2746: [0mThis JSX tag's 'children' prop expects a single child of type 'number | nil', but multiple children were provided.

[7m8[0m (<Component>
[7m [0m [91m  ~~~~~~~~~[0m

[96mindex.tsx[0m:[93m8[0m:[93m3[0m - [91merror[0m[90m TLUA2769: [0mNo overload matches this call.
  The last overload gave the following error.
    This JSX tag's 'children' prop expects a single child of type 'number | nil', but multiple children were provided.

[7m8[0m (<Component>
[7m [0m [91m  ~~~~~~~~~[0m

  [96mindex.tsx[0m:[93m7[0m:[93m18[0m - The last overload is declared here.
    [7m7[0m declare function Component(props: { children?: number }): any;
    [7m [0m [96m                 ~~~~~~~~~[0m


Found 2 errors in the same file, starting at: index.tsx[90m:8[0m

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
(React.createElement(Component, nil,
    React.createElement("div", nil),
    React.createElement("div", nil)));

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./index.tsx"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"1af0770f1670b525d04410724718ce6e-interface JsxElementChildrenAttribute { children: {}; }\ninterface JsxIntrinsicElements { div: {} }\n\ndeclare React: any;\n\ndeclare function Component(props: never): any;\ndeclare function Component(props: { children?: number }): any;\n(<Component>\n    <div />\n    <div />\n</Component>)"],"options":{"jsx":3,"module":99,"strict":true},"semanticDiagnosticsPerFile":[[2,[{"pos":233,"end":242,"code":2746,"category":1,"messageKey":"This_JSX_tag_s_0_prop_expects_a_single_child_of_type_1_but_multiple_children_were_provided_2746","messageArgs":["children","number | nil"]},{"pos":233,"end":242,"code":2769,"category":1,"messageKey":"No_overload_matches_this_call_2769","messageChain":[{"pos":233,"end":242,"code":2770,"category":1,"messageKey":"The_last_overload_gave_the_following_error_2770","messageChain":[{"pos":233,"end":242,"code":2746,"category":1,"messageKey":"This_JSX_tag_s_0_prop_expects_a_single_child_of_type_1_but_multiple_children_were_provided_2746","messageArgs":["children","number | nil"]}]}],"relatedInformation":[{"pos":185,"end":194,"code":2771,"category":1,"messageKey":"The_last_overload_is_declared_here_2771"}]}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tsx"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./index.tsx"
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
      "fileName": "./index.tsx",
      "version": "1af0770f1670b525d04410724718ce6e-interface JsxElementChildrenAttribute { children: {}; }\ninterface JsxIntrinsicElements { div: {} }\n\ndeclare React: any;\n\ndeclare function Component(props: never): any;\ndeclare function Component(props: { children?: number }): any;\n(<Component>\n    <div />\n    <div />\n</Component>)",
      "signature": "1af0770f1670b525d04410724718ce6e-interface JsxElementChildrenAttribute { children: {}; }\ninterface JsxIntrinsicElements { div: {} }\n\ndeclare React: any;\n\ndeclare function Component(props: never): any;\ndeclare function Component(props: { children?: number }): any;\n(<Component>\n    <div />\n    <div />\n</Component>)",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "jsx": 3,
    "module": 99,
    "strict": true
  },
  "semanticDiagnosticsPerFile": [
    [
      "./index.tsx",
      [
        {
          "pos": 233,
          "end": 242,
          "code": 2746,
          "category": 1,
          "messageKey": "This_JSX_tag_s_0_prop_expects_a_single_child_of_type_1_but_multiple_children_were_provided_2746",
          "messageArgs": [
            "children",
            "number | nil"
          ]
        },
        {
          "pos": 233,
          "end": 242,
          "code": 2769,
          "category": 1,
          "messageKey": "No_overload_matches_this_call_2769",
          "messageChain": [
            {
              "pos": 233,
              "end": 242,
              "code": 2770,
              "category": 1,
              "messageKey": "The_last_overload_gave_the_following_error_2770",
              "messageChain": [
                {
                  "pos": 233,
                  "end": 242,
                  "code": 2746,
                  "category": 1,
                  "messageKey": "This_JSX_tag_s_0_prop_expects_a_single_child_of_type_1_but_multiple_children_were_provided_2746",
                  "messageArgs": [
                    "children",
                    "number | nil"
                  ]
                }
              ]
            }
          ],
          "relatedInformation": [
            {
              "pos": 185,
              "end": 194,
              "code": 2771,
              "category": 1,
              "messageKey": "The_last_overload_is_declared_here_2771"
            }
          ]
        }
      ]
    ]
  ],
  "size": 2084
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/index.tsx
Signatures::


Edit [0]:: no change

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mindex.tsx[0m:[93m8[0m:[93m3[0m - [91merror[0m[90m TLUA2746: [0mThis JSX tag's 'children' prop expects a single child of type 'number | nil', but multiple children were provided.

[7m8[0m (<Component>
[7m [0m [91m  ~~~~~~~~~[0m

[96mindex.tsx[0m:[93m8[0m:[93m3[0m - [91merror[0m[90m TLUA2769: [0mNo overload matches this call.
  The last overload gave the following error.
    This JSX tag's 'children' prop expects a single child of type 'number | nil', but multiple children were provided.

[7m8[0m (<Component>
[7m [0m [91m  ~~~~~~~~~[0m

  [96mindex.tsx[0m:[93m7[0m:[93m18[0m - The last overload is declared here.
    [7m7[0m declare function Component(props: { children?: number }): any;
    [7m [0m [96m                 ~~~~~~~~~[0m


Found 2 errors in the same file, starting at: index.tsx[90m:8[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::

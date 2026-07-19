currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/node_modules/@types/react/index.d.tlua] *new* 
interface JsxElement {}
interface JsxIntrinsicElements {
    div: {
        propA?: boolean;
    };
}
//// [/home/src/workspaces/project/node_modules/react/jsx-runtime.lua] *new* 

//// [/home/src/workspaces/project/src/index.tsx] *new* 
local App = () => <div propA={true}></div>;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{ 
    "compilerOptions": { 
        "module": "commonjs",
        "jsx": "react-jsx", 
        "incremental": true, 
        "jsxImportSource": "react" 
    } 
}

tlua --strict
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96msrc/index.tsx[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA100051: [0mJSX element implicitly has type 'any' because no interface 'JsxIntrinsicElements' exists.

[7m1[0m local App = () => <div propA={true}></div>;
[7m [0m [91m                  ~~~~~~~~~~~~~~~~~~[0m

[96msrc/index.tsx[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA2875: [0mThis JSX tag requires the module path 'react/jsx-runtime' to exist, but none could be found. Make sure you have types for the appropriate package installed.

[7m1[0m local App = () => <div propA={true}></div>;
[7m [0m [91m                  ~~~~~~~~~~~~~~~~~~~~~~~~[0m

[96msrc/index.tsx[0m:[93m1[0m:[93m37[0m - [91merror[0m[90m TLUA100051: [0mJSX element implicitly has type 'any' because no interface 'JsxIntrinsicElements' exists.

[7m1[0m local App = () => <div propA={true}></div>;
[7m [0m [91m                                    ~~~~~~[0m


Found 3 errors in the same file, starting at: src/index.tsx[90m:1[0m

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
//// [/home/src/workspaces/project/src/index.lua] *new* 
import { jsx as _jsx } from "react/jsx-runtime";
local App = function() return (_jsx("div", { propA: true })) end;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./src/index.tsx"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"1b0175636c32d2043ca678e69550803e-local App = () => <div propA={true}></div>;"],"options":{"jsx":4,"jsxImportSource":"react","module":1,"strict":true},"semanticDiagnosticsPerFile":[[2,[{"pos":18,"end":36,"code":100051,"category":1,"messageKey":"JSX_element_implicitly_has_type_any_because_no_interface_0_exists_100051","messageArgs":["JsxIntrinsicElements"]},{"pos":18,"end":42,"code":2875,"category":1,"messageKey":"This_JSX_tag_requires_the_module_path_0_to_exist_but_none_could_be_found_Make_sure_you_have_types_fo_2875","messageArgs":["react/jsx-runtime"]},{"pos":36,"end":42,"code":100051,"category":1,"messageKey":"JSX_element_implicitly_has_type_any_because_no_interface_0_exists_100051","messageArgs":["JsxIntrinsicElements"]}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/index.tsx"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/index.tsx"
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
      "fileName": "./src/index.tsx",
      "version": "1b0175636c32d2043ca678e69550803e-local App = () => <div propA={true}></div>;",
      "signature": "1b0175636c32d2043ca678e69550803e-local App = () => <div propA={true}></div>;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "jsx": 4,
    "jsxImportSource": "react",
    "module": 1,
    "strict": true
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/index.tsx",
      [
        {
          "pos": 18,
          "end": 36,
          "code": 100051,
          "category": 1,
          "messageKey": "JSX_element_implicitly_has_type_any_because_no_interface_0_exists_100051",
          "messageArgs": [
            "JsxIntrinsicElements"
          ]
        },
        {
          "pos": 18,
          "end": 42,
          "code": 2875,
          "category": 1,
          "messageKey": "This_JSX_tag_requires_the_module_path_0_to_exist_but_none_could_be_found_Make_sure_you_have_types_fo_2875",
          "messageArgs": [
            "react/jsx-runtime"
          ]
        },
        {
          "pos": 36,
          "end": 42,
          "code": 100051,
          "category": 1,
          "messageKey": "JSX_element_implicitly_has_type_any_because_no_interface_0_exists_100051",
          "messageArgs": [
            "JsxIntrinsicElements"
          ]
        }
      ]
    ]
  ],
  "size": 1649
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/index.tsx
Signatures::

currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/node_modules/solid-js/jsx-runtime.d.tlua] *new* 
interface JsxIntrinsicElements { div: {}; }
//// [/home/src/projects/project/node_modules/solid-js/package.json] *new* 
{
    "name": "solid-js",
    "type": "module"
}
//// [/home/src/projects/project/src/main.tsx] *new* 
return <div/>;
//// [/home/src/projects/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "composite": true,
        "module": "Node16",
        "jsx": "react-jsx",
        "jsxImportSource": "solid-js",
    },
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96msrc/main.tsx[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m return <div/>;
[7m [0m [91m~~~~~~[0m

[96msrc/main.tsx[0m:[93m1[0m:[93m8[0m - [91merror[0m[90m TS2875: [0mThis JSX tag requires the module path 'solid-js/jsx-runtime' to exist, but none could be found. Make sure you have types for the appropriate package installed.

[7m1[0m return <div/>;
[7m [0m [91m       ~~~~~~[0m

[96msrc/main.tsx[0m:[93m1[0m:[93m8[0m - [91merror[0m[90m TS100051: [0mJSX element implicitly has type 'any' because no interface 'JsxIntrinsicElements' exists.

[7m1[0m return <div/>;
[7m [0m [91m       ~~~~~~[0m


Found 3 errors in the same file, starting at: src/main.tsx[90m:1[0m

//// [/home/src/projects/project/src/main.lua] *new* 
import { jsx as _jsx } from "solid-js/jsx-runtime";
return _jsx("div", {});

//// [/home/src/projects/project/tsconfig.tsbuildinfo] *new* 
{"version":"FakeTSVersion","root":[2],"fileNames":["lib.luajit.d.tlua","./src/main.tsx"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"fa8440e41d31d81279063324e177417a-return <div/>;"],"options":{"composite":true,"jsx":4,"jsxImportSource":"solid-js","module":100},"semanticDiagnosticsPerFile":[[2,[{"pos":7,"end":13,"code":2875,"category":1,"messageKey":"This_JSX_tag_requires_the_module_path_0_to_exist_but_none_could_be_found_Make_sure_you_have_types_fo_2875","messageArgs":["solid-js/jsx-runtime"]},{"pos":7,"end":13,"code":100051,"category":1,"messageKey":"JSX_element_implicitly_has_type_any_because_no_interface_0_exists_100051","messageArgs":["JsxIntrinsicElements"]}]]],"emitDiagnosticsPerFile":[[2,[{"end":6,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]],"emitSignatures":[2]}
//// [/home/src/projects/project/tsconfig.tsbuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/main.tsx"
      ],
      "original": 2
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/main.tsx"
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
      "fileName": "./src/main.tsx",
      "version": "fa8440e41d31d81279063324e177417a-return <div/>;",
      "signature": "fa8440e41d31d81279063324e177417a-return <div/>;",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "composite": true,
    "jsx": 4,
    "jsxImportSource": "solid-js",
    "module": 100
  },
  "semanticDiagnosticsPerFile": [
    [
      "./src/main.tsx",
      [
        {
          "pos": 7,
          "end": 13,
          "code": 2875,
          "category": 1,
          "messageKey": "This_JSX_tag_requires_the_module_path_0_to_exist_but_none_could_be_found_Make_sure_you_have_types_fo_2875",
          "messageArgs": [
            "solid-js/jsx-runtime"
          ]
        },
        {
          "pos": 7,
          "end": 13,
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
  "emitDiagnosticsPerFile": [
    [
      "./src/main.tsx",
      [
        {
          "end": 6,
          "code": 100054,
          "category": 1,
          "messageKey": "Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"
        }
      ]
    ]
  ],
  "emitSignatures": [
    {
      "file": "./src/main.tsx",
      "original": 2
    }
  ],
  "size": 1622
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
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/projects/project/src/main.tsx
Signatures::

currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/tslibs/TS/Lib/lib.d.tlua] *new* 
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
type ReturnType<T extends (...args: any) => any> = T extends (...args: any) => infer R ? R : any;
type InstanceType<T extends abstract new (...args: any) => any> = T extends abstract new (...args: any) => infer R ? R : any;
//// [/home/src/workspaces/project/MessageablePerson.tlua] *new* 
local Messageable = () => {
    return {
        message: 'hello' as string,
    }
};
local wrapper = () => Messageable();
type MessageablePerson = ReturnType<typeof wrapper>;
//// [/home/src/workspaces/project/main.tlua] *new* 
function logMessage( person: MessageablePerson ) {
    console.log( person.message );
}
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{ 
    "compilerOptions": {
        "module": "esnext",
        "declaration": true
    }
}

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/MessageablePerson.lua] *new* 
local Messageable = function()
    return {
        message, 'hello' as string,
    };
end;
local wrapper = function() return (Messageable()) end;

//// [/home/src/workspaces/project/main.lua] *new* 
function logMessage(person) {
    console.log(person.message);
}

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;","0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./MessageablePerson.tlua",
        "./main.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
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
      "fileName": "./MessageablePerson.tlua",
      "version": "2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./main.tlua",
      "version": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "signature": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "module": 99
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
    "./main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./MessageablePerson.tlua",
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
      "./main.tlua",
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
  "size": 1628
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::


Edit [0]:: no change

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::


Edit [1]:: narrow message to a literal type
//// [/home/src/workspaces/project/MessageablePerson.tlua] *modified* 
local Messageable = () => {
    return {
        message: 'hello' as 'hello',
    }
};
local wrapper = () => Messageable();
type MessageablePerson = ReturnType<typeof wrapper>;

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as 'hello',
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m

//// [/home/src/workspaces/project/MessageablePerson.lua] *modified* 
local Messageable = function()
    return {
        message, 'hello' as 'hello',
    };
end;
local wrapper = function() return (Messageable()) end;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"ac174b1cd471f4fd29b46b424fd87f8e-local Messageable = () => {\n    return {\n        message: 'hello' as 'hello',\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./MessageablePerson.tlua",
        "./main.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
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
      "fileName": "./MessageablePerson.tlua",
      "version": "ac174b1cd471f4fd29b46b424fd87f8e-local Messageable = () => {\n    return {\n        message: 'hello' as 'hello',\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "ac174b1cd471f4fd29b46b424fd87f8e-local Messageable = () => {\n    return {\n        message: 'hello' as 'hello',\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./main.tlua",
      "version": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "signature": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "module": 99
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
    "./main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./MessageablePerson.tlua",
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
      "./main.tlua",
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
  "size": 1796
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/MessageablePerson.tlua


Edit [2]:: no change

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as 'hello',
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::


Edit [3]:: widen message back to string
//// [/home/src/workspaces/project/MessageablePerson.tlua] *modified* 
local Messageable = () => {
    return {
        message: 'hello' as string,
    }
};
local wrapper = () => Messageable();
type MessageablePerson = ReturnType<typeof wrapper>;

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m

//// [/home/src/workspaces/project/MessageablePerson.lua] *modified* 
local Messageable = function()
    return {
        message, 'hello' as string,
    };
end;
local wrapper = function() return (Messageable()) end;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./MessageablePerson.tlua",
        "./main.tlua"
      ],
      "original": [
        2,
        3
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
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
      "fileName": "./MessageablePerson.tlua",
      "version": "2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "2dfc7c7b63ae101bafb51e02252302ae-local Messageable = () => {\n    return {\n        message: 'hello' as string,\n    }\n};\nlocal wrapper = () => Messageable();\ntype MessageablePerson = ReturnType<typeof wrapper>;",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./main.tlua",
      "version": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "signature": "0806657dce6b2dee513d74e4380386d0-function logMessage( person: MessageablePerson ) {\n    console.log( person.message );\n}",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "options": {
    "declaration": true,
    "module": 99
  },
  "semanticDiagnosticsPerFile": [
    "lib.luajit.d.tlua",
    "./MessageablePerson.tlua",
    "./main.tlua"
  ],
  "emitDiagnosticsPerFile": [
    [
      "./MessageablePerson.tlua",
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
      "./main.tlua",
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
  "size": 1795
}

tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/MessageablePerson.tlua


Edit [4]:: no change

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = () => {
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson ) {
[7m [0m [91m~~~~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  MessageablePerson.tlua[90m:1[0m
     1  main.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
*not cached* /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*not cached* /home/src/workspaces/project/MessageablePerson.tlua
*not cached* /home/src/workspaces/project/main.tlua
Signatures::

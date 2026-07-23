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
local Messageable = function()
    return {
        message: 'hello' as string,
    }
end;
local wrapper = function() return Messageable() end;
type MessageablePerson = ReturnType<typeof wrapper>;
//// [/home/src/workspaces/project/main.tlua] *new* 
function logMessage( person: MessageablePerson )
    console.log( person.message );
end
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

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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
local wrapper = function()
  return Messageable();
end;

//// [/home/src/workspaces/project/main.lua] *new* 
function logMessage(person)
  console.log(person.message);
end

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;","8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
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
      "version": "02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./main.tlua",
      "version": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
      "signature": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
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
  "size": 1649
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

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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
local Messageable = function()
    return {
        message: 'hello' as 'hello',
    }
end;
local wrapper = function() return Messageable() end;
type MessageablePerson = ReturnType<typeof wrapper>;

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as 'hello',
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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
local wrapper = function()
  return Messageable();
end;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"46eca9f21ecb1d575fe36b6e3900bd65-local Messageable = function()\n    return {\n        message: 'hello' as 'hello',\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
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
      "version": "46eca9f21ecb1d575fe36b6e3900bd65-local Messageable = function()\n    return {\n        message: 'hello' as 'hello',\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "46eca9f21ecb1d575fe36b6e3900bd65-local Messageable = function()\n    return {\n        message: 'hello' as 'hello',\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./main.tlua",
      "version": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
      "signature": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
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
  "size": 1817
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

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as 'hello',
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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
local Messageable = function()
    return {
        message: 'hello' as string,
    }
end;
local wrapper = function() return Messageable() end;
type MessageablePerson = ReturnType<typeof wrapper>;

tlua --incremental
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mMessageablePerson.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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
local wrapper = function()
  return Messageable();
end;

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,3]],"fileNames":["lib.luajit.d.tlua","./MessageablePerson.tlua","./main.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend"],"options":{"declaration":true,"module":99},"semanticDiagnosticsPerFile":[1,2,3],"emitDiagnosticsPerFile":[[2,[{"end":5,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]],[3,[{"end":8,"code":100054,"category":1,"messageKey":"Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054"}]]]}
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
      "version": "02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "02908e1f4269b6ab7bde110f803caae1-local Messageable = function()\n    return {\n        message: 'hello' as string,\n    }\nend;\nlocal wrapper = function() return Messageable() end;\ntype MessageablePerson = ReturnType<typeof wrapper>;",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./main.tlua",
      "version": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
      "signature": "8a402e9598994cb7a501b71174803091-function logMessage( person: MessageablePerson )\n    console.log( person.message );\nend",
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
  "size": 1816
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

[7m1[0m local Messageable = function()
[7m [0m [91m~~~~~[0m

[96mMessageablePerson.tlua[0m:[93m3[0m:[93m16[0m - [91merror[0m[90m TLUA1005: [0m',' expected.

[7m3[0m         message: 'hello' as string,
[7m [0m [91m               ~[0m

[96mmain.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m function logMessage( person: MessageablePerson )
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

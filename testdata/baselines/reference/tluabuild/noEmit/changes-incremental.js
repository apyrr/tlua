currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/src/class.tlua] *new* 
local classC = {
    prop = 1,
};
return { classC = classC };
//// [/home/src/workspaces/project/src/directUse.tlua] *new* 
local indirect = require('src.indirectClass');
indirect.indirectClass.classC.prop;
//// [/home/src/workspaces/project/src/indirectClass.tlua] *new* 
local classMod = require('src.class');
local indirectClass = {
    classC = classMod.classC,
};
return { indirectClass = indirectClass };
//// [/home/src/workspaces/project/src/indirectUse.tlua] *new* 
local indirect = require('src.indirectClass');
indirect.indirectClass.classC.prop;
//// [/home/src/workspaces/project/src/noChangeFile.tlua] *new* 
function writeLog(s: string) end
//// [/home/src/workspaces/project/src/noChangeFileWithEmitSpecificError.tlua] *new* 
function someFunc(arguments: boolean, ...rest: any[]) end
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "compilerOptions":  { "incremental": true }
}

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because output file 'tluaconfig.tluabuildinfo' does not exist

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/src/class.lua] *new* 
local classC = {
  prop = 1,
};
return { classC = classC };

//// [/home/src/workspaces/project/src/directUse.lua] *new* 
local indirect = require('src.indirectClass');
indirect.indirectClass.classC.prop;

//// [/home/src/workspaces/project/src/indirectClass.lua] *new* 
local classMod = require('src.class');
local indirectClass = {
  classC = classMod.classC,
};
return { indirectClass = indirectClass };

//// [/home/src/workspaces/project/src/indirectUse.lua] *new* 
local indirect = require('src.indirectClass');
indirect.indirectClass.classC.prop;

//// [/home/src/workspaces/project/src/noChangeFile.lua] *new* 
function writeLog(s)
end

//// [/home/src/workspaces/project/src/noChangeFileWithEmitSpecificError.lua] *new* 
function someFunc(arguments, ...rest)
end

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},"01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };","45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
      "signature": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "size": 2036
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /home/src/workspaces/project/src/class.tlua
*refresh*    /home/src/workspaces/project/src/indirectClass.tlua
*refresh*    /home/src/workspaces/project/src/directUse.tlua
*refresh*    /home/src/workspaces/project/src/indirectUse.tlua
*refresh*    /home/src/workspaces/project/src/noChangeFile.tlua
*refresh*    /home/src/workspaces/project/src/noChangeFileWithEmitSpecificError.tlua
Signatures::


Edit [0]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [1]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [2]:: Introduce error but still noEmit
//// [/home/src/workspaces/project/src/class.tlua] *modified* 
local classC = {
    prop1 = 1,
};
return { classC = classC };

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]],"affectedFilesPendingEmit":[2]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "affectedFilesPendingEmit": [
    [
      "./src/class.tlua",
      "Js",
      2
    ]
  ],
  "size": 2235
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/class.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/class.tlua


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [3]:: Fix error and emit
//// [/home/src/workspaces/project/src/class.tlua] *modified* 
local classC = {
    prop = 1,
};
return { classC = classC };

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/class.lua] *rewrite with same content*
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "size": 2203
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/class.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/class.tlua


Edit [4]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [5]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [6]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [7]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [8]:: Introduce error and emit
//// [/home/src/workspaces/project/src/class.tlua] *modified* 
local classC = {
    prop1 = 1,
};
return { classC = classC };

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/class.lua] *modified* 
local classC = {
  prop1 = 1,
};
return { classC = classC };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "8be09a1392ba7e2d8eb8a82451d54596-local classC = {\n    prop1 = 1,\n};\nreturn { classC = classC };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "size": 2204
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/class.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/class.tlua


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [9]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [10]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [11]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [12]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Diff:: incremental build misses dependent-file errors because module signatures come from declaration emit, which TS100054 suppresses; the clean build re-checks everything
--- nonIncremental.output.txt
+++ incremental.output.txt
@@ -1,21 +1,3 @@
-[96msrc/directUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
-[96msrc/indirectUse.tlua[0m:[93m2[0m:[93m31[0m - [91merror[0m[90m TLUA2551: [0mProperty 'prop' does not exist on type '{ prop1: number; }'. Did you mean 'prop1'?
-
-[7m2[0m indirect.indirectClass.classC.prop;
-[7m [0m [91m                              ~~~~[0m
-
-  [96msrc/class.tlua[0m:[93m2[0m:[93m5[0m - 'prop1' is declared here.
-    [7m2[0m     prop1 = 1,
-    [7m [0m [96m    ~~~~~~~~~[0m
-
 [96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

 [7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
@@ -27,10 +9,5 @@
 [7m [0m [91m                                         ~~~~[0m


-Found 4 errors in 3 files.
-
-Errors  Files
-     1  src/directUse.tlua[90m:2[0m
-     1  src/indirectUse.tlua[90m:2[0m
-     2  src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m
+Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


Edit [13]:: Fix error and no emit
//// [/home/src/workspaces/project/src/class.tlua] *modified* 
local classC = {
    prop = 1,
};
return { classC = classC };

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]],"affectedFilesPendingEmit":[2]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "affectedFilesPendingEmit": [
    [
      "./src/class.tlua",
      "Js",
      2
    ]
  ],
  "size": 2234
}

tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/workspaces/project/src/class.tlua
Signatures::
(computed .d.ts) /home/src/workspaces/project/src/class.tlua


Edit [14]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m

//// [/home/src/workspaces/project/src/class.lua] *modified* 
local classC = {
  prop = 1,
};
return { classC = classC };

//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":[[2,7]],"fileNames":["lib.luajit.d.tlua","./src/class.tlua","./src/indirectClass.tlua","./src/directUse.tlua","./src/indirectUse.tlua","./src/noChangeFile.tlua","./src/noChangeFileWithEmitSpecificError.tlua"],"fileInfos":[{"version":"d4695a71643e88fc868e824886bcb416-/// <reference no-default-lib=\"true\"/>\ninterface Boolean {}\ninterface Function {}\ninterface CallableFunction {}\ninterface NewableFunction {}\ninterface IArguments {}\ninterface Number { toExponential: any; }\ninterface Object {}\ninterface RegExp {}\ninterface String { charAt: any; }\ninterface Array<T> { length: number; [n: number]: T; }\ninterface ReadonlyArray<T> {}\ninterface SymbolConstructor {\n    (desc?: string | number): symbol;\n    for(name: string): symbol;\n    readonly toStringTag: symbol;\n}\ndeclare Symbol: SymbolConstructor;\ninterface Symbol {\n    readonly [Symbol.toStringTag]: string;\n}\ndeclare console: { log(msg: any): void; };\ndeclare function require(module: string): any;","affectsGlobalScope":true,"impliedNodeFormat":1},{"version":"01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };","signature":"59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n","impliedNodeFormat":1},"45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;","79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end","49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end"],"semanticDiagnosticsPerFile":[[7,[{"pos":18,"end":27,"code":1215,"category":1,"messageKey":"Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215","messageArgs":["arguments"]},{"pos":41,"end":45,"code":100034,"category":1,"messageKey":"A_vararg_parameter_cannot_have_a_name_100034"}]]]}
//// [/home/src/workspaces/project/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./src/class.tlua",
        "./src/indirectClass.tlua",
        "./src/directUse.tlua",
        "./src/indirectUse.tlua",
        "./src/noChangeFile.tlua",
        "./src/noChangeFileWithEmitSpecificError.tlua"
      ],
      "original": [
        2,
        7
      ]
    }
  ],
  "fileNames": [
    "lib.luajit.d.tlua",
    "./src/class.tlua",
    "./src/indirectClass.tlua",
    "./src/directUse.tlua",
    "./src/indirectUse.tlua",
    "./src/noChangeFile.tlua",
    "./src/noChangeFileWithEmitSpecificError.tlua"
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
      "fileName": "./src/class.tlua",
      "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
      "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
      "impliedNodeFormat": "CommonJS",
      "original": {
        "version": "01b7d09c6307220002a1cfad352a5df8-local classC = {\n    prop = 1,\n};\nreturn { classC = classC };",
        "signature": "59f451fde7dea7b84d390cf390374c8b-\n(0,5): error100054: Declaration_emit_is_not_supported_for_a_Lua_module_yet_100054\n",
        "impliedNodeFormat": 1
      }
    },
    {
      "fileName": "./src/indirectClass.tlua",
      "version": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "signature": "45d4ad1225e205a8c50e34dea21f6146-local classMod = require('src.class');\nlocal indirectClass = {\n    classC = classMod.classC,\n};\nreturn { indirectClass = indirectClass };",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/directUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/indirectUse.tlua",
      "version": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "signature": "7627d2557b2559156a4f49ce9c718fa0-local indirect = require('src.indirectClass');\nindirect.indirectClass.classC.prop;",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFile.tlua",
      "version": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "signature": "79aaed0fe83f9356eae8986134f0606f-function writeLog(s: string) end",
      "impliedNodeFormat": "CommonJS"
    },
    {
      "fileName": "./src/noChangeFileWithEmitSpecificError.tlua",
      "version": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "signature": "49a2ef83f2e7401bfd97db750939d27b-function someFunc(arguments: boolean, ...rest: any[]) end",
      "impliedNodeFormat": "CommonJS"
    }
  ],
  "semanticDiagnosticsPerFile": [
    [
      "./src/noChangeFileWithEmitSpecificError.tlua",
      [
        {
          "pos": 18,
          "end": 27,
          "code": 1215,
          "category": 1,
          "messageKey": "Invalid_use_of_0_Modules_are_automatically_in_strict_mode_1215",
          "messageArgs": [
            "arguments"
          ]
        },
        {
          "pos": 41,
          "end": 45,
          "code": 100034,
          "category": 1,
          "messageKey": "A_vararg_parameter_cannot_have_a_name_100034"
        }
      ]
    ]
  ],
  "size": 2203
}

tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [15]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [16]:: No Change run with noEmit

tlua -b -v --noEmit
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::


Edit [17]:: No Change run with emit

tlua -b -v
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[[90mHH:MM:SS AM[0m] Projects in this build: 
    * tluaconfig.json

[[90mHH:MM:SS AM[0m] Project 'tluaconfig.json' is out of date because buildinfo file 'tluaconfig.tluabuildinfo' indicates that program needs to report errors.

[[90mHH:MM:SS AM[0m] Building project 'tluaconfig.json'...

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1215: [0mInvalid use of 'arguments'. Modules are automatically in strict mode.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                  ~~~~~~~~~[0m

[96msrc/noChangeFileWithEmitSpecificError.tlua[0m:[93m1[0m:[93m42[0m - [91merror[0m[90m TLUA100034: [0mA vararg parameter cannot have a name.

[7m1[0m function someFunc(arguments: boolean, ...rest: any[]) end
[7m [0m [91m                                         ~~~~[0m


Found 2 errors in the same file, starting at: src/noChangeFileWithEmitSpecificError.tlua[90m:1[0m


tluaconfig.json::
SemanticDiagnostics::
Signatures::

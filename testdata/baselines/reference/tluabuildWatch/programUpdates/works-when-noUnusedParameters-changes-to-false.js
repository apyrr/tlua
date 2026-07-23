currentDirectory::/user/username/projects/myproject
useCaseSensitiveFileNames::true
Input::
//// [/user/username/projects/myproject/index.tlua] *new* 
local fn = function(a: string, b: string) return b end;
//// [/user/username/projects/myproject/tluaconfig.json] *new* 
{
    "compilerOptions": {
        "noUnusedParameters": true,
    },
}

tlua -b -w
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] Starting compilation in watch mode...

[96mindex.tlua[0m:[93m1[0m:[93m21[0m - [91merror[0m[90m TLUA6133: [0m'a' is declared but its value is never read.

[7m1[0m local fn = function(a: string, b: string) return b end;
[7m [0m [91m                    ~[0m

[[90mHH:MM:SS AM[0m] Found 1 error. Watching for file changes.

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
//// [/user/username/projects/myproject/index.lua] *new* 
local fn = function(a, b)
    return b;
end;

//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo] *new* 
{"version":"FakeTSVersion","root":["./index.tlua"],"semanticErrors":true}
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo.readable.baseline.txt] *new* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua"
      ],
      "original": "./index.tlua"
    }
  ],
  "size": 73,
  "semanticErrors": true
}

Watch Registrations::
Directory watches::
  /user/username/projects/myproject (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/myproject/index.tlua
Signatures::


Edit [0]:: Change tsconfig to set noUnusedParameters to false
//// [/user/username/projects/myproject/tluaconfig.json] *modified* 
{
    "compilerOptions": {
        "noUnusedParameters": false,
    },
}


Output::
[2J[3J[H[[90mHH:MM:SS AM[0m] File change detected. Starting incremental compilation...

[[90mHH:MM:SS AM[0m] Found 0 errors. Watching for file changes.

//// [/user/username/projects/myproject/index.lua] *rewrite with same content*
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo] *modified* 
{"version":"FakeTSVersion","root":["./index.tlua"]}
//// [/user/username/projects/myproject/tluaconfig.tluabuildinfo.readable.baseline.txt] *modified* 
{
  "version": "FakeTSVersion",
  "root": [
    {
      "files": [
        "./index.tlua"
      ],
      "original": "./index.tlua"
    }
  ],
  "size": 51
}

Watch Registrations::
Directory watches::
  /user/username/projects/myproject (recursive)
tluaconfig.json::
SemanticDiagnostics::
*refresh*    /home/src/tslibs/TS/Lib/lib.luajit.d.tlua
*refresh*    /user/username/projects/myproject/index.tlua
Signatures::

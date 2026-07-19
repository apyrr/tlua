currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/D.tlua] *new* 
local x = 10;
//// [/home/src/workspaces/project/c.tlua] *new* 
local D = require("D");
//// [/home/src/workspaces/project/d.tlua] *new* 
local y = 20;
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{
    "files": ["c.tlua", "d.tlua"]
}

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96mc.tlua[0m:[93m1[0m:[93m19[0m - [91merror[0m[90m TLUA1261: [0mAlready included file name '/home/src/workspaces/project/D.tlua' differs from file name '/home/src/workspaces/project/d.tlua' only in casing.
  The file is in the program because:
    Imported via "D" from file '/home/src/workspaces/project/c.tlua'
    Part of 'files' list in tluaconfig.json

[7m1[0m local D = require("D");
[7m [0m [91m                  ~~~[0m

  [96mtluaconfig.json[0m:[93m2[0m:[93m25[0m - File is matched by 'files' list specified here.
    [7m2[0m     "files": ["c.tlua", "d.tlua"]
    [7m [0m [96m                        ~~~~~~~~[0m


Found 1 error in c.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/D.lua] *new* 
local x = 10;

//// [/home/src/workspaces/project/c.lua] *new* 
local D = require("D");

//// [/home/src/workspaces/project/d.lua] *new* 
local y = 20;



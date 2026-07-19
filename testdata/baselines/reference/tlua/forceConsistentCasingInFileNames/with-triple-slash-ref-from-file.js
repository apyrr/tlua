currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::false
Input::
//// [/home/src/workspaces/project/src/c.tlua] *new* 
/// <reference path="./D.tlua"/>
//// [/home/src/workspaces/project/src/d.tlua] *new* 
interface c { }
//// [/home/src/workspaces/project/tluaconfig.json] *new* 
{ }

tlua 
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96msrc/c.tlua[0m:[93m1[0m:[93m22[0m - [91merror[0m[90m TLUA1261: [0mAlready included file name '/home/src/workspaces/project/src/D.tlua' differs from file name '/home/src/workspaces/project/src/d.tlua' only in casing.
  The file is in the program because:
    Referenced via './D.tlua' from file '/home/src/workspaces/project/src/c.tlua'
    Matched by default include pattern '**/*'

[7m1[0m /// <reference path="./D.tlua"/>
[7m [0m [91m                     ~~~~~~~~[0m


Found 1 error in src/c.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/src/D.lua] *new* 

//// [/home/src/workspaces/project/src/c.lua] *new* 
-- / <reference path="./D.tlua"/>



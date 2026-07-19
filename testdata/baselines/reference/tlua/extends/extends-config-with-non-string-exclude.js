currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/project/base.json] *new* 
{
    "exclude": [1],
}
//// [/home/src/projects/project/main.tlua] *new* 
local x = 1;
//// [/home/src/projects/project/tluaconfig.json] *new* 
{
    "extends": "./base.json",
}

tlua -p tluaconfig.json --pretty false
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
base.json(2,17): error TLUA5024: Compiler option 'exclude' requires a value of type string.
//// [/home/src/projects/project/main.lua] *new* 
local x = 1;

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


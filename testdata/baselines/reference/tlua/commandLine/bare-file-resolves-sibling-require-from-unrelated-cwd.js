currentDirectory::/home/src/elsewhere
useCaseSensitiveFileNames::true
Input::
//// [/home/src/app/main.tlua] *new* 
local util = require("util");
local n: number = util.value;
n;
//// [/home/src/app/util.tlua] *new* 
local value = 42;
return { value = value };

tlua /home/src/app/main.tlua
ExitStatus:: Success
Output::
//// [/home/src/app/main.lua] *new* 
local util = require("util");
local n = util.value;
n;

//// [/home/src/app/util.lua] *new* 
local value = 42;
return { value = value };

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


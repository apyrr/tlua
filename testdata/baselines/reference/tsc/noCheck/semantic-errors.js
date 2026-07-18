currentDirectory::/home/src/workspaces/project
useCaseSensitiveFileNames::true
Input::
//// [/home/src/workspaces/project/a.tlua] *new* 
local a: number = "hello";
//// [/home/src/workspaces/project/b.tlua] *new* 
local b = 10;
//// [/home/src/workspaces/project/tsconfig.json] *new* 
{
    "compilerOptions": {
        "declaration": true,
        "incremental": false
    }
}

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

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
//// [/home/src/workspaces/project/a.lua] *new* 
local a = "hello";

//// [/home/src/workspaces/project/b.lua] *new* 
local b = 10;




Edit [0]:: no change

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [1]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [2]:: no change

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [3]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [4]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [5]:: no change

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [6]:: Introduce error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a: number = "hello";

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [7]:: no change

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [8]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96ma.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local a: number = "hello";
[7m [0m [91m      ~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 3 errors in 2 files.

Errors  Files
     2  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [9]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [10]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m


Found 2 errors in 2 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*



Edit [11]:: Add file with error
//// [/home/src/workspaces/project/c.tlua] *new* 
local c: number = "hello";

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *new* 
local c = "hello";




Edit [12]:: Introduce error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a: number = "hello";

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a: number = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*



Edit [13]:: Fix `a` error with noCheck
//// [/home/src/workspaces/project/a.tlua] *modified* 
local a = "hello";

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*



Edit [14]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*



Edit [15]:: no change

tlua --noCheck
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m


Found 3 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     1  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*



Edit [16]:: No Change run with checking

tlua 
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96ma.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local a = "hello";
[7m [0m [91m~~~~~[0m

[96mb.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local b = 10;
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m1[0m - [91merror[0m[90m TS100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m1[0m local c: number = "hello";
[7m [0m [91m~~~~~[0m

[96mc.tlua[0m:[93m1[0m:[93m7[0m - [91merror[0m[90m TS2322: [0mType 'string' is not assignable to type 'number'.

[7m1[0m local c: number = "hello";
[7m [0m [91m      ~[0m


Found 4 errors in 3 files.

Errors  Files
     1  a.tlua[90m:1[0m
     1  b.tlua[90m:1[0m
     2  c.tlua[90m:1[0m

//// [/home/src/workspaces/project/a.lua] *rewrite with same content*
//// [/home/src/workspaces/project/b.lua] *rewrite with same content*
//// [/home/src/workspaces/project/c.lua] *rewrite with same content*


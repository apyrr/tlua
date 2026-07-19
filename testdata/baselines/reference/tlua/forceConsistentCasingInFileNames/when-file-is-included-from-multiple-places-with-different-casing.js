currentDirectory::/home/src/projects/project
useCaseSensitiveFileNames::false
Input::
//// [/home/src/projects/project/node_modules/fp-ts/lib/struct.tlua] *new* 
local foo = 10;
//// [/home/src/projects/project/src/anotherFile.tlua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");
//// [/home/src/projects/project/src/oneMore.tlua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");
//// [/home/src/projects/project/src/struct.tlua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");
//// [/home/src/projects/project/tluaconfig.json] *new* 
{}

tlua --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96msrc/Struct.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/node_modules/fp-ts/lib/struct.tlua' differs from already included file name '/home/src/projects/project/node_modules/fp-ts/lib/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/oneMore.tlua'

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/Struct.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'fp-ts.lib.struct' resolves to the same file as 'fp-ts.lib.Struct'; a module has one canonical name.

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/Struct.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/src/struct.tlua' differs from already included file name '/home/src/projects/project/src/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "src.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Matched by default include pattern '**/*'

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

[96msrc/Struct.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'src.struct' resolves to the same file as 'src.Struct'; a module has one canonical name.

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

[96msrc/anotherFile.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/node_modules/fp-ts/lib/struct.tlua' differs from already included file name '/home/src/projects/project/node_modules/fp-ts/lib/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/oneMore.tlua'

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFile.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'fp-ts.lib.struct' resolves to the same file as 'fp-ts.lib.Struct'; a module has one canonical name.

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/anotherFile.tlua[0m:[93m3[0m:[93m21[0m - [91merror[0m[90m TLUA1261: [0mAlready included file name '/home/src/projects/project/src/Struct.tlua' differs from file name '/home/src/projects/project/src/struct.tlua' only in casing.
  The file is in the program because:
    Imported via "src.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Matched by default include pattern '**/*'

[7m3[0m local xs3 = require("src.Struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

[96msrc/anotherFile.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/src/struct.tlua' differs from already included file name '/home/src/projects/project/src/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "src.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Matched by default include pattern '**/*'

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

[96msrc/anotherFile.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'src.struct' resolves to the same file as 'src.Struct'; a module has one canonical name.

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

[96msrc/oneMore.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/node_modules/fp-ts/lib/struct.tlua' differs from already included file name '/home/src/projects/project/node_modules/fp-ts/lib/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "fp-ts.lib.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "fp-ts.lib.struct" from file '/home/src/projects/project/src/oneMore.tlua'

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m2[0m:[93m21[0m - File is included via import here.
    [7m2[0m local xs2 = require("fp-ts.lib.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/oneMore.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'fp-ts.lib.struct' resolves to the same file as 'fp-ts.lib.Struct'; a module has one canonical name.

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/oneMore.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/home/src/projects/project/src/struct.tlua' differs from already included file name '/home/src/projects/project/src/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "src.Struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/Struct.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/anotherFile.tlua'
    Imported via "src.Struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Imported via "src.struct" from file '/home/src/projects/project/src/oneMore.tlua'
    Matched by default include pattern '**/*'

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/Struct.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/anotherFile.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("src.struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

  [96msrc/oneMore.tlua[0m:[93m3[0m:[93m21[0m - File is included via import here.
    [7m3[0m local xs3 = require("src.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~[0m

[96msrc/oneMore.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'src.struct' resolves to the same file as 'src.Struct'; a module has one canonical name.

[7m4[0m local xs4 = require("src.struct");
[7m [0m [91m                    ~~~~~~~~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
node_modules/fp-ts/lib/Struct.tlua
   Imported via "fp-ts.lib.Struct" from file 'src/anotherFile.tlua'
   Imported via "fp-ts.lib.struct" from file 'src/anotherFile.tlua'
   Imported via "fp-ts.lib.Struct" from file 'src/Struct.tlua'
   Imported via "fp-ts.lib.struct" from file 'src/Struct.tlua'
   Imported via "fp-ts.lib.Struct" from file 'src/oneMore.tlua'
   Imported via "fp-ts.lib.struct" from file 'src/oneMore.tlua'
src/Struct.tlua
   Imported via "src.Struct" from file 'src/anotherFile.tlua'
   Imported via "src.Struct" from file 'src/Struct.tlua'
   Imported via "src.struct" from file 'src/Struct.tlua'
   Imported via "src.struct" from file 'src/anotherFile.tlua'
   Imported via "src.Struct" from file 'src/oneMore.tlua'
   Imported via "src.struct" from file 'src/oneMore.tlua'
   Matched by default include pattern '**/*'
src/anotherFile.tlua
   Matched by default include pattern '**/*'
src/oneMore.tlua
   Matched by default include pattern '**/*'

Found 13 errors in 3 files.

Errors  Files
     4  src/Struct.tlua[90m:2[0m
     5  src/anotherFile.tlua[90m:2[0m
     4  src/oneMore.tlua[90m:2[0m

//// [/home/src/projects/project/src/Struct.lua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");

//// [/home/src/projects/project/src/anotherFile.lua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");

//// [/home/src/projects/project/src/oneMore.lua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("src.Struct");
local xs4 = require("src.struct");

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


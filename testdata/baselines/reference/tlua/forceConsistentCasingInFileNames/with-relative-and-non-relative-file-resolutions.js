currentDirectory::/user/username/projects/myproject
useCaseSensitiveFileNames::false
Input::
//// [/user/username/projects/myproject/node_modules/fp-ts/lib/struct.tlua] *new* 
local foo = 10;
//// [/user/username/projects/myproject/src/struct.tlua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("Struct");
local xs4 = require("struct");

tlua /user/username/projects/myproject/src/struct.tlua --forceConsistentCasingInFileNames --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsGenerated
Output::
[96msrc/struct.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/user/username/projects/myproject/node_modules/fp-ts/lib/struct.tlua' differs from already included file name '/user/username/projects/myproject/node_modules/fp-ts/lib/Struct.tlua' only in casing.
  The file is in the program because:
    Imported via "fp-ts.lib.Struct" from file '/user/username/projects/myproject/src/struct.tlua'
    Imported via "fp-ts.lib.struct" from file '/user/username/projects/myproject/src/struct.tlua'

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

  [96msrc/struct.tlua[0m:[93m1[0m:[93m21[0m - File is included via import here.
    [7m1[0m local xs1 = require("fp-ts.lib.Struct");
    [7m [0m [96m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/struct.tlua[0m:[93m2[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'fp-ts.lib.struct' resolves to the same file as 'fp-ts.lib.Struct'; a module has one canonical name.

[7m2[0m local xs2 = require("fp-ts.lib.struct");
[7m [0m [91m                    ~~~~~~~~~~~~~~~~~~[0m

[96msrc/struct.tlua[0m:[93m3[0m:[93m21[0m - [91merror[0m[90m TLUA1149: [0mFile name '/user/username/projects/myproject/src/Struct.tlua' differs from already included file name '/user/username/projects/myproject/src/struct.tlua' only in casing.
  The file is in the program because:
    Root file specified for compilation
    Imported via "Struct" from file '/user/username/projects/myproject/src/struct.tlua'
    Imported via "struct" from file '/user/username/projects/myproject/src/struct.tlua'

[7m3[0m local xs3 = require("Struct");
[7m [0m [91m                    ~~~~~~~~[0m

  [96msrc/struct.tlua[0m:[93m4[0m:[93m21[0m - File is included via import here.
    [7m4[0m local xs4 = require("struct");
    [7m [0m [96m                    ~~~~~~~~[0m

[96msrc/struct.tlua[0m:[93m4[0m:[93m21[0m - [91merror[0m[90m TLUA100055: [0mModule name 'struct' resolves to the same file as 'Struct'; a module has one canonical name.

[7m4[0m local xs4 = require("struct");
[7m [0m [91m                    ~~~~~~~~[0m

../../../../home/src/tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
node_modules/fp-ts/lib/Struct.tlua
   Imported via "fp-ts.lib.Struct" from file 'src/struct.tlua'
   Imported via "fp-ts.lib.struct" from file 'src/struct.tlua'
src/struct.tlua
   Root file specified for compilation
   Imported via "Struct" from file 'src/struct.tlua'
   Imported via "struct" from file 'src/struct.tlua'

Found 4 errors in the same file, starting at: src/struct.tlua[90m:2[0m

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
//// [/user/username/projects/myproject/src/struct.lua] *new* 
local xs1 = require("fp-ts.lib.Struct");
local xs2 = require("fp-ts.lib.struct");
local xs3 = require("Struct");
local xs4 = require("struct");



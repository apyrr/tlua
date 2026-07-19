currentDirectory::/home/src/projects/myproject
useCaseSensitiveFileNames::true
Input::
//// [/home/src/projects/configs/first/tluaconfig.json] *new* 
{
    "extends": "../second/tluaconfig.json",
    "include": ["${configDir}/src"],
    "compilerOptions": {
        "typeRoots": ["root1", "${configDir}/root2", "root3"],
        "types": [],
    },
}
//// [/home/src/projects/configs/second/tluaconfig.json] *new* 
{
    "files": ["${configDir}/main.tlua"],
    "compilerOptions": {
        "declarationDir": "${configDir}/decls",
    },
    "watchOptions": {
        "excludeFiles": ["${configDir}/main.tlua"],
    },
}
//// [/home/src/projects/myproject/main.tlua] *new* 
// some comment
local y = 10;
//// [/home/src/projects/myproject/tluaconfig.json] *new* 
{
    "extends": "../configs/first/tluaconfig.json",
    "compilerOptions": {
        "declaration": true,
        "outDir": "outDir",
        "traceResolution": true,
    },
}

tlua --explainFiles
ExitStatus:: DiagnosticsPresent_OutputsSkipped
Output::
[96mmain.tlua[0m:[93m2[0m:[93m1[0m - [91merror[0m[90m TLUA100054: [0mDeclaration emit is not supported for a Lua module yet.

[7m2[0m local y = 10;
[7m [0m [91m~~~~~[0m

../../tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
main.tlua
   Part of 'files' list in tluaconfig.json

Found 1 error in main.tlua[90m:2[0m

//// [/home/src/projects/myproject/outDir/main.lua] *new* 
-- some comment
local y = 10;

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


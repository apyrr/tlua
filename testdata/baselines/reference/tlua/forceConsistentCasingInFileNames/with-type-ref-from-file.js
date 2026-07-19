currentDirectory::/user/username/projects/myproject
useCaseSensitiveFileNames::false
Input::
//// [/user/username/projects/myproject/src/file2.d.tlua] *new* 
/// <reference types="./fileOne.d.tlua"/>
declare y: c;
//// [/user/username/projects/myproject/src/fileOne.d.tlua] *new* 
interface c { }
//// [/user/username/projects/myproject/tluaconfig.json] *new* 
{ }

tlua -p /user/username/projects/myproject --explainFiles --traceResolution
ExitStatus:: Success
Output::
======== Resolving type reference directive './fileOne.d.tlua', containing file '/user/username/projects/myproject/src/file2.d.tlua', root directory '/user/username/projects/myproject/node_modules/@types,/user/username/projects/node_modules/@types,/user/username/node_modules/@types,/user/node_modules/@types,/node_modules/@types'. ========
Resolving with primary search path '/user/username/projects/myproject/node_modules/@types, /user/username/projects/node_modules/@types, /user/username/node_modules/@types, /user/node_modules/@types, /node_modules/@types'.
Directory '/user/username/projects/myproject/node_modules/@types' does not exist, skipping all lookups in it.
Directory '/user/username/projects/node_modules/@types' does not exist, skipping all lookups in it.
Directory '/user/username/node_modules/@types' does not exist, skipping all lookups in it.
Directory '/user/node_modules/@types' does not exist, skipping all lookups in it.
Directory '/node_modules/@types' does not exist, skipping all lookups in it.
Looking up in 'node_modules' folder, initial location '/user/username/projects/myproject/src'.
Loading module as file / folder, candidate module location '/user/username/projects/myproject/src/fileOne.d.tlua', target file types: Declaration.
File name '/user/username/projects/myproject/src/fileOne.d.tlua' has a '.d.tlua' extension - stripping it.
File '/user/username/projects/myproject/src/fileOne.d.tlua' exists - use it as a name resolution result.
Resolving real path for '/user/username/projects/myproject/src/fileOne.d.tlua', result '/user/username/projects/myproject/src/fileOne.d.tlua'.
======== Type reference directive './fileOne.d.tlua' was successfully resolved to '/user/username/projects/myproject/src/fileOne.d.tlua', primary: false. ========
../../../../home/src/tslibs/TS/Lib/lib.luajit.d.tlua
   Default library for target 'ES2025'
src/fileOne.d.tlua
   Type library referenced via './fileOne.d.tlua' from file 'src/file2.d.tlua'
   Matched by default include pattern '**/*'
src/file2.d.tlua
   Matched by default include pattern '**/*'
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

